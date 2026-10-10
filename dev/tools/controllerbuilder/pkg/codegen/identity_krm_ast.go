// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
)

type krmSpecField struct {
	GoName    string
	JSONName  string
	TypeExpr  string
	IsPointer bool
	IsRef     bool
}

type krmTypeInfo struct {
	PackageName      string
	SpecFields       []krmSpecField
	HasProjectRef    bool
	HasOrgRef        bool
	OrgRefIsPtr      bool
	HasFolderRef     bool
	FolderRefIsPtr   bool
	HasLocation      bool
	LocationGoName   string
	LocationJSONName string
	LocationIsPtr    bool
	HasResourceID    bool
	ResourceIDIsPtr  bool
	ResourceIDServer bool
	StatusField      string
	StatusFieldIsPtr bool
	Overrides        HandWrittenOverrides
}

func inspectKRMPackage(pkgDir, kind, defaultPkgName string) (*krmTypeInfo, error) {
	info := &krmTypeInfo{
		PackageName: defaultPkgName,
		Overrides: HandWrittenOverrides{
			IdentityMethods: make(map[string]bool),
			ResourceMethods: make(map[string]bool),
			RefMethods:      make(map[string]bool),
		},
	}

	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return info, nil
		}
		return nil, err
	}

	fset := token.NewFileSet()
	structDecls := map[string]*ast.StructType{}
	structComments := map[string]map[string]string{}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		if err := inspectKRMFile(fset, pkgDir, entry.Name(), kind, info, structDecls, structComments); err != nil {
			return nil, err
		}
	}

	populateSpecFieldFlags(kind, info, fset, structDecls, structComments)
	populateStatusFieldFlags(kind, info, fset, structDecls)
	return info, nil
}

func inspectKRMFile(fset *token.FileSet, pkgDir, name, kind string, info *krmTypeInfo, structDecls map[string]*ast.StructType, structComments map[string]map[string]string) error {
	isGenerated := strings.HasSuffix(name, ".generated.go") || strings.HasSuffix(name, "_generated_test.go") || strings.HasPrefix(name, "zz_generated.")
	isTest := strings.HasSuffix(name, "_test.go")

	fullPath := filepath.Join(pkgDir, name)
	fileAST, err := parser.ParseFile(fset, fullPath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", fullPath, err)
	}
	if fileAST.Name != nil && fileAST.Name.Name != "" && !strings.HasSuffix(fileAST.Name.Name, "_test") {
		info.PackageName = fileAST.Name.Name
	}
	if isTest {
		if !isGenerated && name == strings.ToLower(kind)+"_identity_test.go" {
			info.Overrides.HasHandWrittenIdentityTest = true
		}
		return nil
	}

	for _, decl := range fileAST.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			inspectKRMGenDecl(d, kind, isGenerated, info, structDecls, structComments)
		case *ast.FuncDecl:
			if !isGenerated {
				inspectKRMFuncDecl(fset, d, kind, info)
			}
		}
	}
	return nil
}

func inspectKRMGenDecl(d *ast.GenDecl, kind string, isGenerated bool, info *krmTypeInfo, structDecls map[string]*ast.StructType, structComments map[string]map[string]string) {
	for _, spec := range d.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			inspectKRMTypeSpec(s, kind, isGenerated, info, structDecls, structComments)
		case *ast.ValueSpec:
			if !isGenerated {
				inspectKRMValueSpec(s, kind, info)
			}
		}
	}
}

func inspectKRMTypeSpec(s *ast.TypeSpec, kind string, isGenerated bool, info *krmTypeInfo, structDecls map[string]*ast.StructType, structComments map[string]map[string]string) {
	if st, ok := s.Type.(*ast.StructType); ok {
		structDecls[s.Name.Name] = st
		structComments[s.Name.Name] = extractFieldComments(st)
	}
	if isGenerated {
		return
	}
	switch s.Name.Name {
	case kind + "Identity":
		info.Overrides.HasIdentityStruct = true
	case kind + "Ref":
		info.Overrides.HasRefStruct = true
	}
}

func inspectKRMValueSpec(s *ast.ValueSpec, kind string, info *krmTypeInfo) {
	for _, ident := range s.Names {
		switch ident.Name {
		case kind + "IdentityFormat":
			info.Overrides.HasIdentityFormat = true
		case kind + "GVK":
			info.Overrides.HasGVKVar = true
		}
	}
}

func inspectKRMFuncDecl(fset *token.FileSet, d *ast.FuncDecl, kind string, info *krmTypeInfo) {
	idType := kind + "Identity"
	refType := kind + "Ref"
	if d.Recv == nil || len(d.Recv.List) == 0 {
		if d.Name.Name == "getIdentityFrom"+kind+"Spec" {
			info.Overrides.HasGetIdentityFromSpec = true
		}
		if d.Name.Name == "init" && d.Body != nil {
			bodyText := renderExpr(fset, d.Body)
			if strings.Contains(bodyText, "&"+refType) && strings.Contains(bodyText, "Register") {
				info.Overrides.HasRefRegister = true
			}
		}
		return
	}
	switch exprTypeName(d.Recv.List[0].Type) {
	case idType:
		info.Overrides.IdentityMethods[d.Name.Name] = true
	case kind:
		info.Overrides.ResourceMethods[d.Name.Name] = true
	case refType:
		info.Overrides.RefMethods[d.Name.Name] = true
	}
}

func populateSpecFieldFlags(kind string, info *krmTypeInfo, fset *token.FileSet, structDecls map[string]*ast.StructType, structComments map[string]map[string]string) {
	specStruct, ok := structDecls[kind+"Spec"]
	if !ok {
		return
	}
	info.SpecFields = collectSpecFields(fset, specStruct, structDecls)
	comments := structComments[kind+"Spec"]
	for _, sf := range info.SpecFields {
		switch sf.GoName {
		case "ProjectRef":
			info.HasProjectRef = true
		case "OrganizationRef":
			info.HasOrgRef = true
			info.OrgRefIsPtr = sf.IsPointer
		case "FolderRef":
			info.HasFolderRef = true
			info.FolderRefIsPtr = sf.IsPointer
		case "Location", "Region", "Zone":
			if !info.HasLocation || sf.GoName == "Location" {
				info.HasLocation = true
				info.LocationGoName = sf.GoName
				info.LocationJSONName = sf.JSONName
				info.LocationIsPtr = sf.IsPointer
			}
		case "ResourceID":
			info.HasResourceID = true
			info.ResourceIDIsPtr = sf.IsPointer
			if strings.Contains(strings.ToLower(comments["ResourceID"]), "server-generated") {
				info.ResourceIDServer = true
			}
		}
	}
}

func populateStatusFieldFlags(kind string, info *krmTypeInfo, fset *token.FileSet, structDecls map[string]*ast.StructType) {
	statusStruct, ok := structDecls[kind+"Status"]
	if !ok {
		return
	}
	for _, sf := range collectSpecFields(fset, statusStruct, structDecls) {
		if sf.GoName == "ExternalRef" {
			info.StatusField = "ExternalRef"
			info.StatusFieldIsPtr = sf.IsPointer
			return
		}
		if sf.GoName == "Name" && info.StatusField == "" {
			info.StatusField = "Name"
			info.StatusFieldIsPtr = sf.IsPointer
		}
	}
}

func extractFieldComments(st *ast.StructType) map[string]string {
	out := map[string]string{}
	if st == nil || st.Fields == nil {
		return out
	}
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			continue
		}
		var text string
		if f.Doc != nil {
			text += f.Doc.Text()
		}
		if f.Comment != nil {
			text += " " + f.Comment.Text()
		}
		for _, n := range f.Names {
			out[n.Name] = strings.TrimSpace(text)
		}
	}
	return out
}

func collectSpecFields(fset *token.FileSet, st *ast.StructType, allStructs map[string]*ast.StructType) []krmSpecField {
	if st == nil || st.Fields == nil {
		return nil
	}
	var out []krmSpecField
	for _, f := range st.Fields.List {
		typeStr := renderExpr(fset, f.Type)
		baseType := strings.TrimPrefix(typeStr, "*")
		if len(f.Names) == 0 || strings.HasSuffix(baseType, "ProjectAndLocationRef") || strings.HasSuffix(baseType, "ProjectAndLocationParent") {
			out = append(out, collectEmbeddedStructFields(fset, baseType, allStructs)...)
			if len(f.Names) == 0 {
				continue
			}
		}
		isPtr := strings.HasPrefix(typeStr, "*")
		jsonName := parseJSONTagName(f.Tag)
		for _, ident := range f.Names {
			jn := jsonName
			if jn == "" {
				jn = lowerFirst(ident.Name)
			}
			isRef := strings.HasSuffix(ident.Name, "Ref") || strings.HasSuffix(baseType, "Ref")
			out = append(out, krmSpecField{
				GoName:    ident.Name,
				JSONName:  jn,
				TypeExpr:  typeStr,
				IsPointer: isPtr,
				IsRef:     isRef,
			})
		}
	}
	return out
}

func collectEmbeddedStructFields(fset *token.FileSet, baseType string, allStructs map[string]*ast.StructType) []krmSpecField {
	switch {
	case strings.HasSuffix(baseType, "ProjectAndLocationRef") || strings.HasSuffix(baseType, "ProjectAndLocationParent") || strings.HasSuffix(baseType, ".Parent"):
		return []krmSpecField{
			{GoName: "ProjectRef", JSONName: "projectRef", TypeExpr: "*refs.ProjectRef", IsPointer: true, IsRef: true},
			{GoName: "Location", JSONName: "location", TypeExpr: "string", IsPointer: false},
		}
	case strings.HasSuffix(baseType, "ProjectParent") || strings.HasSuffix(baseType, "ProjectRef"):
		return []krmSpecField{
			{GoName: "ProjectRef", JSONName: "projectRef", TypeExpr: "*refs.ProjectRef", IsPointer: true, IsRef: true},
		}
	case strings.HasSuffix(baseType, "OrganizationParent"):
		return []krmSpecField{
			{GoName: "OrganizationRef", JSONName: "organizationRef", TypeExpr: "*refs.OrganizationRef", IsPointer: true, IsRef: true},
		}
	case strings.HasSuffix(baseType, "FolderParent"):
		return []krmSpecField{
			{GoName: "FolderRef", JSONName: "folderRef", TypeExpr: "*refs.FolderRef", IsPointer: true, IsRef: true},
		}
	default:
		if emb, ok := allStructs[baseType]; ok {
			return collectSpecFields(fset, emb, allStructs)
		}
		return nil
	}
}

func renderExpr(fset *token.FileSet, expr ast.Node) string {
	var buf bytes.Buffer
	_ = format.Node(&buf, fset, expr)
	return buf.String()
}

func exprTypeName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.StarExpr:
		return exprTypeName(e.X)
	case *ast.Ident:
		return e.Name
	case *ast.IndexExpr:
		return exprTypeName(e.X)
	default:
		return ""
	}
}

func parseJSONTagName(tagLit *ast.BasicLit) string {
	if tagLit == nil {
		return ""
	}
	st := reflect.StructTag(strings.Trim(tagLit.Value, "`\""))
	val := st.Get("json")
	if val == "" {
		return ""
	}
	return strings.Split(val, ",")[0]
}

func (g *IdentityGenerator) resolveSpecAncestors(plan *IdentityPlan, typeInfo *krmTypeInfo) {
	ancestorEnd := ancestorSliceEnd(plan)
	if ancestorEnd <= 0 {
		return
	}
	var ancestorFields []IdentityField
	for _, s := range plan.Segments[:ancestorEnd] {
		if s.IsVariable {
			ancestorFields = append(ancestorFields, s.Field)
		}
	}
	if len(ancestorFields) == 0 {
		return
	}

	lastAncestor := ancestorFields[len(ancestorFields)-1]
	if !isStandardRootAndLocationTuple(ancestorFields) {
		if parentRef := matchParentRefField(lastAncestor, typeInfo.SpecFields); parentRef != nil {
			plan.ParentRef = &ParentRefResolution{
				SpecFieldName: parentRef.GoName,
				SpecJSONName:  parentRef.JSONName,
				IsPointer:     parentRef.IsPointer,
				CoveredFields: ancestorFields,
				TailSegments:  plan.Segments[ancestorEnd:],
			}
			return
		}
	}

	var unknownAncestors []string
	for _, af := range ancestorFields {
		res, known := resolveSingleAncestor(af, typeInfo)
		plan.Ancestors = append(plan.Ancestors, res)
		if !known {
			unknownAncestors = append(unknownAncestors, af.Collection+"/{"+af.RawPlaceholder+"}")
		}
	}
	if len(unknownAncestors) > 0 {
		plan.Judgement = append(plan.Judgement, judgement.Entry{
			Kind:   plan.Kind,
			Group:  plan.Group,
			Reason: judgement.ReasonIdentityRootUnknown,
			Detail: fmt.Sprintf("pattern %q contains ancestor segment(s) %s with no matching reference or field on %sSpec",
				plan.RawPattern, strings.Join(unknownAncestors, ", "), plan.Kind),
			Status: judgement.StatusOpen,
		})
		plan.HasUnresolvedAncestors = true
	}
}

func ancestorSliceEnd(plan *IdentityPlan) int {
	if plan.ResourceIDField != nil && len(plan.Segments) >= 2 {
		return len(plan.Segments) - 2
	}
	if plan.ResourceIDField == nil && len(plan.Segments) >= 1 {
		return len(plan.Segments) - 1
	}
	return len(plan.Segments)
}

func isStandardRootAndLocationTuple(ancestorFields []IdentityField) bool {
	switch len(ancestorFields) {
	case 1:
		return isRootCollection(ancestorFields[0].Collection)
	case 2:
		return isRootCollection(ancestorFields[0].Collection) && isLocationCollection(ancestorFields[1].Collection)
	default:
		return false
	}
}

func isRootCollection(col string) bool {
	return col == "projects" || col == "organizations" || col == "folders"
}

func isLocationCollection(col string) bool {
	return col == "locations" || col == "regions" || col == "zones"
}

func resolveSingleAncestor(af IdentityField, typeInfo *krmTypeInfo) (AncestorResolution, bool) {
	switch af.Collection {
	case "projects":
		known := typeInfo.HasProjectRef || len(typeInfo.SpecFields) > 0
		return AncestorResolution{Field: af, Source: AncestorProjectRef}, known
	case "organizations":
		if !typeInfo.HasOrgRef {
			return AncestorResolution{Field: af, Source: AncestorOrganizationRef}, false
		}
		return AncestorResolution{
			Field:         af,
			Source:        AncestorOrganizationRef,
			SpecFieldName: "OrganizationRef",
			SpecJSONName:  "organizationRef",
			IsPointer:     typeInfo.OrgRefIsPtr,
		}, true
	case "folders":
		if !typeInfo.HasFolderRef {
			return AncestorResolution{Field: af, Source: AncestorFolderRef}, false
		}
		return AncestorResolution{
			Field:         af,
			Source:        AncestorFolderRef,
			SpecFieldName: "FolderRef",
			SpecJSONName:  "folderRef",
			IsPointer:     typeInfo.FolderRefIsPtr,
		}, true
	case "locations", "regions", "zones":
		if !typeInfo.HasLocation {
			return AncestorResolution{Field: af, Source: AncestorLocation}, false
		}
		return AncestorResolution{
			Field:         af,
			Source:        AncestorLocation,
			SpecFieldName: typeInfo.LocationGoName,
			SpecJSONName:  typeInfo.LocationJSONName,
			IsPointer:     typeInfo.LocationIsPtr,
		}, true
	default:
		return resolveCustomCollectionAncestor(af, typeInfo.SpecFields)
	}
}

func resolveCustomCollectionAncestor(af IdentityField, fields []krmSpecField) (AncestorResolution, bool) {
	if refField := matchParentRefField(af, fields); refField != nil {
		return AncestorResolution{
			Field:         af,
			Source:        AncestorParentRef,
			SpecFieldName: refField.GoName,
			SpecJSONName:  refField.JSONName,
			IsPointer:     refField.IsPointer,
		}, true
	}
	if strField := matchPlainSpecField(af, fields); strField != nil {
		return AncestorResolution{
			Field:         af,
			Source:        AncestorSpecString,
			SpecFieldName: strField.GoName,
			SpecJSONName:  strField.JSONName,
			IsPointer:     strField.IsPointer,
		}, true
	}
	return AncestorResolution{Field: af, Source: AncestorUnknown}, false
}

func matchParentRefField(af IdentityField, fields []krmSpecField) *krmSpecField {
	singular := strings.ToLower(Singular(af.Collection))
	wantNames := []string{
		strings.ToLower(af.StructField) + "ref",
		singular + "ref",
		"parentref",
	}
	for _, want := range wantNames {
		for i := range fields {
			f := &fields[i]
			if !f.IsRef {
				continue
			}
			lowerGo := strings.ToLower(f.GoName)
			lowerJSON := strings.ToLower(f.JSONName)
			lowerType := strings.ToLower(strings.TrimPrefix(f.TypeExpr, "*"))
			if lowerGo == want || lowerJSON == want || (want != "parentref" && (strings.HasSuffix(lowerGo, want) || strings.HasSuffix(lowerType, want))) {
				return f
			}
		}
	}
	return nil
}

func matchPlainSpecField(af IdentityField, fields []krmSpecField) *krmSpecField {
	singular := strings.ToLower(Singular(af.Collection))
	wantNames := []string{
		strings.ToLower(af.StructField),
		singular,
		strings.ToLower(af.StructField) + "id",
	}
	for _, want := range wantNames {
		for i := range fields {
			f := &fields[i]
			if !f.IsRef && (f.TypeExpr == "string" || f.TypeExpr == "*string") && strings.ToLower(f.GoName) == want {
				return f
			}
		}
	}
	return nil
}
