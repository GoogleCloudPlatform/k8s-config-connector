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

package scaffold

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

// KRMPackageInfo captures structural facts inspected from the KRM API package
// (apis/<service>/<version>) and existing direct controller package
// (pkg/controller/direct/<service>) to specialize the generated controller.
type KRMPackageInfo struct {
	RootPath string

	// Identity facts
	IdentityInspected         bool   // true if Identity AST was inspected
	IdentityImportPath        string // Optional extra import path if Identity is in apis/refs
	IdentityPkgPrefix         string // "krm." (default) or "refs."
	IdentityTypeName          string // e.g., "BeyondCorpAppConnectionIdentity" or "AppConnectionIdentity"
	IdentityConstructorExpr   string // e.g., "krm.NewAppConnectionIdentity(ctx, reader, obj)"
	HasGetIdentity            bool   // true if (*<Kind>) GetIdentity(ctx, reader) exists
	HasIdentitySpecified      bool   // true if (*<IdentityType>) HasIdentitySpecified() exists
	HasFromExternal           bool   // true if (*<IdentityType>) FromExternal(string) exists
	IdentityUsesMetadataName  bool   // true if Identity constructor falls back to metadata.name (GetResourceID or GetName)
	IdentityIsClientSpecified bool   // true if Identity uses metadata.name fallback and enforces strict statusIdentity == specIdentity
	HasNoParent               bool   // true if Identity is a root-level resource with no parent segment
	ParentStringExpr          string // e.g., "a.id.ParentString()" or "a.id.Parent().String()"
	IDExpr                    string // e.g., "a.id.ID()" or "a.id.Connection"

	// Spec & Status facts
	HasSpecLabels     bool   // true if <Kind>Spec has a Labels field
	HasSpecReferences bool   // true if <Kind>Spec or nested structs have *Ref / *Refs fields
	StatusInspected   bool   // true if <Kind>Status struct was found
	HasObservedState  bool   // true if <Kind>Status has an ObservedState field and mapper exists
	LocationExpr      string // e.g., "direct.ValueOf(obj.Spec.Location)" or "obj.Spec.Location"

	// Controller package & mapper facts
	PackageName                string // Go package name for the controller file
	MapperPackageImport        string // Optional import path when controller is in a sub-package
	MapperPrefix               string // Optional package qualifier (e.g. "apihub.") for mapper calls
	MapperVersionSuffix        string // "" or "_v1alpha1" / "_v1beta1"
	HasExplicitSpecSuffix      bool   // true if Spec mapper version suffix was explicitly found in AST
	ObservedStateVersionSuffix string // "" or "_v1alpha1" / "_v1beta1"
	HasExplicitObsSuffix       bool   // true if ObservedState mapper version suffix was explicitly found in AST
	DetectedPBImportPath       string // PB import path inspected from <Kind>Spec*_ToProto return type
	DetectedPBAlias            string // PB package alias inspected from <Kind>Spec*_ToProto
	DetectedRegionalEndpoint   string // "" or "prefix" / "rep" inferred from sibling controllers
	UseGRPCClient              bool   // true if sibling controllers in pkg/controller/direct/<service> use GRPCClientOptions()
}

// InspectKRMPackageForController parses the Go files in apis/<service>/<version>
// and pkg/controller/direct/<service> to extract Identity, Spec, Status, Mapper,
// and RegionalEndpoint facts for deterministic controller generation.
func InspectKRMPackageForController(rootPath, serviceName, version, kind, protoResource string) *KRMPackageInfo {
	return InspectKRMPackageForControllerWithTargetDir(rootPath, serviceName, version, kind, protoResource, "")
}

// InspectKRMPackageForControllerWithTargetDir is like InspectKRMPackageForController,
// additionally accepting the target directory where the controller file will be written
// so sub-package controllers (e.g. pkg/controller/direct/apihub/apihubapi) can qualify
// mapper calls from their parent service package.
func InspectKRMPackageForControllerWithTargetDir(rootPath, serviceName, version, kind, protoResource, targetDir string) *KRMPackageInfo {
	info := &KRMPackageInfo{
		RootPath:          rootPath,
		IdentityPkgPrefix: "krm.",
	}
	apiDir := filepath.Join(rootPath, "apis", serviceName, version)
	directDir := filepath.Join(rootPath, "pkg", "controller", "direct", serviceName)

	inspectAPIDir(rootPath, apiDir, kind, protoResource, info)
	inspectDirectDir(directDir, targetDir, serviceName, kind, version, info)
	detectMultiversionFromGenerateScripts(apiDir, directDir, version, info)

	return info
}

// --- API Package AST Inspection Helpers (Footnotes) ---

type apiDecls struct {
	files           []*ast.File
	structDecls     map[string]*ast.StructType
	methodsByRecv   map[string]map[string]*ast.FuncDecl
	funcDecls       map[string]*ast.FuncDecl
	gcpURLTemplates map[string]string
}

func inspectAPIDir(rootPath, apiDir, kind, protoResource string, info *KRMPackageInfo) {
	files := parseGoFilesInDir(apiDir)
	if len(files) == 0 {
		return
	}
	decls := newAPIDecls(files)

	if kindMethods := decls.methodsByRecv[kind]; kindMethods != nil {
		_, info.HasGetIdentity = kindMethods["GetIdentity"]
	}

	chosenIdentity := findConstructorIdentity(decls.funcDecls, kind, info)
	if chosenIdentity == "" {
		chosenIdentity = resolveIdentityTypeName(decls, kind, protoResource)
	}
	if chosenIdentity == "" && rootPath != "" {
		chosenIdentity = resolveIdentityFromRefsFallback(rootPath, kind+"Identity", decls, info)
	}
	if info.IdentityConstructorExpr == "" && !info.HasGetIdentity {
		populateIdentityConstructorFallback(decls.funcDecls, chosenIdentity, kind+"Identity", protoResource+"Identity", info)
	}
	if chosenIdentity != "" {
		populateIdentityMethodsAndFields(decls, chosenIdentity, protoResource, info)
	}

	inspectSpecStruct(decls.structDecls, kind+"Spec", info)
	if info.LocationExpr == "" && chosenIdentity != "" {
		info.LocationExpr = resolveLocationFromIdentity(decls, chosenIdentity)
	}
	inspectStatusStruct(decls.structDecls, kind+"Status", info)
}

func newAPIDecls(files []*ast.File) *apiDecls {
	d := &apiDecls{
		files:           files,
		structDecls:     make(map[string]*ast.StructType),
		methodsByRecv:   make(map[string]map[string]*ast.FuncDecl),
		funcDecls:       make(map[string]*ast.FuncDecl),
		gcpURLTemplates: make(map[string]string),
	}
	collectDeclsFromFiles(files, d.structDecls, d.methodsByRecv, d.funcDecls, d.gcpURLTemplates)
	return d
}

func findConstructorIdentity(funcDecls map[string]*ast.FuncDecl, kind string, info *KRMPackageInfo) string {
	for _, fnName := range slices.Sorted(maps.Keys(funcDecls)) {
		if !strings.HasPrefix(fnName, "New") || !strings.HasSuffix(fnName, "Identity") {
			continue
		}
		fn := funcDecls[fnName]
		if !funcAcceptsKindReceiver(fn, kind) {
			continue
		}
		info.IdentityConstructorExpr = formatIdentityConstructorCall(fnName, countFuncParams(fn))
		if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
			if retType := receiverTypeName(fn.Type.Results.List[0].Type); retType != "" {
				return retType
			}
		}
		return ""
	}
	return ""
}

func funcAcceptsKindReceiver(fn *ast.FuncDecl, kind string) bool {
	if fn.Type == nil || fn.Type.Params == nil {
		return false
	}
	for _, field := range fn.Type.Params.List {
		if receiverTypeName(field.Type) == kind {
			return true
		}
	}
	return false
}

func countFuncParams(fn *ast.FuncDecl) int {
	if fn == nil || fn.Type == nil || fn.Type.Params == nil {
		return 0
	}
	count := 0
	for _, field := range fn.Type.Params.List {
		if len(field.Names) == 0 {
			count++
		} else {
			count += len(field.Names)
		}
	}
	return count
}

func formatIdentityConstructorCall(fnName string, paramCount int) string {
	if paramCount == 4 {
		return fmt.Sprintf("krm.%s(ctx, reader, obj, u)", fnName)
	}
	return fmt.Sprintf("krm.%s(ctx, reader, obj)", fnName)
}

func resolveIdentityTypeName(decls *apiDecls, kind, protoResource string) string {
	kindIdentityName := kind + "Identity"
	protoIdentityName := protoResource + "Identity"

	for _, cand := range []string{kindIdentityName, protoIdentityName} {
		if _, ok := decls.structDecls[cand]; ok {
			return cand
		}
	}
	for _, cand := range []string{kindIdentityName, protoIdentityName} {
		if _, ok := decls.methodsByRecv[cand]; ok {
			return cand
		}
	}
	for _, name := range slices.Sorted(maps.Keys(decls.structDecls)) {
		if strings.EqualFold(name, kindIdentityName) || strings.EqualFold(name, protoIdentityName) {
			return name
		}
	}
	return ""
}

func resolveIdentityFromRefsFallback(rootPath, kindIdentityName string, decls *apiDecls, info *KRMPackageInfo) string {
	refFiles := parseGoFilesInDir(filepath.Join(rootPath, "apis", "refs"))
	if len(refFiles) == 0 {
		return ""
	}
	refDecls := newAPIDecls(refFiles)
	st, ok := refDecls.structDecls[kindIdentityName]
	if !ok {
		return ""
	}
	decls.structDecls[kindIdentityName] = st
	decls.methodsByRecv[kindIdentityName] = refDecls.methodsByRecv[kindIdentityName]
	if pat, ok := refDecls.gcpURLTemplates[kindIdentityName]; ok {
		decls.gcpURLTemplates[kindIdentityName] = pat
	}
	info.IdentityImportPath = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	info.IdentityPkgPrefix = "refs."
	return kindIdentityName
}

func populateIdentityConstructorFallback(funcDecls map[string]*ast.FuncDecl, chosenIdentity, kindIdentityName, protoIdentityName string, info *KRMPackageInfo) {
	sortedFnNames := slices.Sorted(maps.Keys(funcDecls))
	for _, candName := range []string{"New" + chosenIdentity, "New" + kindIdentityName, "New" + protoIdentityName} {
		for _, fnName := range sortedFnNames {
			if strings.EqualFold(fnName, candName) {
				info.IdentityConstructorExpr = formatIdentityConstructorCall(fnName, countFuncParams(funcDecls[fnName]))
				return
			}
		}
	}
}

func populateIdentityMethodsAndFields(decls *apiDecls, chosenIdentity, protoResource string, info *KRMPackageInfo) {
	info.IdentityInspected = true
	info.IdentityTypeName = chosenIdentity
	info.IdentityUsesMetadataName, info.IdentityIsClientSpecified = inspectIdentityFileTraits(decls.files, chosenIdentity)

	if idMethods := decls.methodsByRecv[chosenIdentity]; idMethods != nil {
		_, info.HasIdentitySpecified = idMethods["HasIdentitySpecified"]
		_, info.HasFromExternal = idMethods["FromExternal"]
		if _, ok := idMethods["ParentString"]; ok {
			info.ParentStringExpr = "a.id.ParentString()"
		} else if _, ok := idMethods["Parent"]; ok {
			info.ParentStringExpr = "a.id.Parent().String()"
		}
		if _, ok := idMethods["ID"]; ok {
			info.IDExpr = "a.id.ID()"
		} else if _, ok := idMethods[protoResource+"ID"]; ok {
			info.IDExpr = "a.id." + protoResource + "ID()"
		}
	}

	inspectIdentityStructFields(decls.structDecls[chosenIdentity], decls.gcpURLTemplates[chosenIdentity], info)
	if info.IDExpr == "" {
		info.IDExpr = "a.id.String()"
	}
}

func inspectIdentityStructFields(st *ast.StructType, urlPattern string, info *KRMPackageInfo) {
	if st == nil || st.Fields == nil {
		return
	}
	var exportedStringFields []string
	hasParentStructField := false
	parentFieldIsString := false

	for _, field := range st.Fields.List {
		ident, isIdent := field.Type.(*ast.Ident)
		for _, name := range field.Names {
			if !name.IsExported() {
				continue
			}
			if name.Name == "Parent" {
				hasParentStructField = true
				parentFieldIsString = isIdent && ident.Name == "string"
			}
			if isIdent && ident.Name == "string" {
				exportedStringFields = append(exportedStringFields, name.Name)
			}
		}
	}

	if info.IDExpr == "" && len(exportedStringFields) > 0 {
		info.IDExpr = "a.id." + exportedStringFields[len(exportedStringFields)-1]
	}
	if info.ParentStringExpr != "" {
		return
	}
	switch {
	case hasParentStructField && parentFieldIsString:
		info.ParentStringExpr = "a.id.Parent"
	case hasParentStructField:
		info.ParentStringExpr = "a.id.Parent.String()"
	case urlPattern != "":
		parentExpr, hasNoParent := deriveParentExprFromGCPURLPattern(urlPattern, exportedStringFields)
		info.HasNoParent = hasNoParent
		info.ParentStringExpr = parentExpr
	case len(exportedStringFields) == 1:
		info.HasNoParent = true
	case len(exportedStringFields) >= 3 && exportedStringFields[0] == "Project" && exportedStringFields[1] == "Location":
		info.ParentStringExpr = `fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)`
	case len(exportedStringFields) == 2 && exportedStringFields[0] == "Project":
		info.ParentStringExpr = `fmt.Sprintf("projects/%s", a.id.Project)`
	}
}

func inspectSpecStruct(structDecls map[string]*ast.StructType, specName string, info *KRMPackageInfo) {
	specStruct := structDecls[specName]
	if specStruct == nil || specStruct.Fields == nil {
		return
	}
	for _, field := range specStruct.Fields.List {
		_, isStar := field.Type.(*ast.StarExpr)
		for _, name := range field.Names {
			switch name.Name {
			case "Labels":
				info.HasSpecLabels = true
			case "Location":
				info.LocationExpr = formatSpecFieldAccess("Location", isStar)
			case "Region":
				if info.LocationExpr == "" {
					info.LocationExpr = formatSpecFieldAccess("Region", isStar)
				}
			}
		}
	}
	visited := make(map[string]bool)
	info.HasSpecReferences = structHasResourceReferences(specName, structDecls, visited)
}

func formatSpecFieldAccess(fieldName string, isPointer bool) string {
	if isPointer {
		return "direct.ValueOf(obj.Spec." + fieldName + ")"
	}
	return "obj.Spec." + fieldName
}

func resolveLocationFromIdentity(decls *apiDecls, chosenIdentity string) string {
	if structHasExportedField(decls.structDecls[chosenIdentity], "Location") {
		return "id.Location"
	}
	idMethods := decls.methodsByRecv[chosenIdentity]
	if idMethods == nil {
		return ""
	}
	if _, ok := idMethods["Location"]; ok {
		return "id.Location()"
	}
	parentFn, ok := idMethods["Parent"]
	if !ok || parentFn.Type == nil || parentFn.Type.Results == nil || len(parentFn.Type.Results.List) == 0 {
		return ""
	}
	parentType := receiverTypeName(parentFn.Type.Results.List[0].Type)
	if parentType == "ProjectAndLocationIdentity" || structHasExportedField(decls.structDecls[parentType], "Location") {
		return "id.Parent().Location"
	}
	return ""
}

func inspectStatusStruct(structDecls map[string]*ast.StructType, statusName string, info *KRMPackageInfo) {
	statusStruct := structDecls[statusName]
	if statusStruct == nil || statusStruct.Fields == nil {
		return
	}
	info.StatusInspected = true
	info.HasObservedState = structHasExportedField(statusStruct, "ObservedState")
}

func structHasExportedField(st *ast.StructType, wantName string) bool {
	if st == nil || st.Fields == nil {
		return false
	}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			if name.Name == wantName {
				return true
			}
		}
	}
	return false
}

// --- Direct Controller Package AST Inspection Helpers (Footnotes) ---

type directDirScan struct {
	targetDir               string
	serviceName             string
	kind                    string
	version                 string
	versionedSpecToProto    string
	unversionedSpecToProto  string
	versionedObsFromProto   string
	unversionedObsFromProto string

	hasGRPCClient               bool
	hasRESTClient               bool
	hasGeneratedMapperFile      bool
	foundSpecMapperInTarget     bool
	targetSpecSuffix            string
	externalMapperDir           string
	externalMapperPkg           string
	externalSpecSuffix          string
	foundObsMapperInTarget      bool
	foundHandwrittenObsInTarget bool
	targetObsSuffix             string
	foundObsMapperExternal      bool
	externalObsSuffix           string
}

func inspectDirectDir(directDir, targetDir, serviceName, kind, version string, info *KRMPackageInfo) {
	if targetDir != "" {
		info.PackageName = filepath.Base(targetDir)
	} else {
		info.PackageName = serviceName
	}

	scan := &directDirScan{
		targetDir:               targetDir,
		serviceName:             serviceName,
		kind:                    kind,
		version:                 version,
		versionedSpecToProto:    kind + "Spec_" + version + "_ToProto",
		unversionedSpecToProto:  kind + "Spec_ToProto",
		versionedObsFromProto:   kind + "ObservedState_" + version + "_FromProto",
		unversionedObsFromProto: kind + "ObservedState_FromProto",
	}

	fset := token.NewFileSet()
	_ = filepath.WalkDir(directDir, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}
		scanDirectGoFile(fset, p, scan, info)
		return nil
	})

	applyDirectScanToInfo(scan, info)
}

func scanDirectGoFile(fset *token.FileSet, filePath string, scan *directDirScan, info *KRMPackageInfo) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	content := string(data)
	isGeneratedMapper := strings.Contains(content, "// +generated:mapper")
	if isGeneratedMapper {
		scan.hasGeneratedMapperFile = true
	}
	if strings.Contains(content, "GRPCClientOptions()") {
		scan.hasGRPCClient = true
	}
	if strings.Contains(content, "RESTClientOptions()") {
		scan.hasRESTClient = true
	}
	detectRegionalEndpointFromSource(content, info)

	f, err := parser.ParseFile(fset, filePath, content, 0)
	if err != nil {
		return
	}
	fileDir := filepath.Dir(filePath)
	if scan.targetDir != "" && fileDir == scan.targetDir && f.Name != nil && f.Name.Name != "" {
		info.PackageName = f.Name.Name
	}

	importsByAlias := collectFileImportsByAlias(f)
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil || fn.Name == nil {
			continue
		}
		switch fn.Name.Name {
		case scan.versionedSpecToProto, scan.unversionedSpecToProto:
			recordSpecMapperDecl(fn, f, fileDir, importsByAlias, scan, info)
		case scan.versionedObsFromProto, scan.unversionedObsFromProto:
			recordObsMapperDecl(fn, fileDir, isGeneratedMapper, scan)
		}
	}
}

func detectRegionalEndpointFromSource(content string, info *KRMPackageInfo) {
	if strings.Contains(content, ".rep.googleapis.com") {
		info.DetectedRegionalEndpoint = "rep"
		return
	}
	if info.DetectedRegionalEndpoint == "" &&
		strings.Contains(content, "option.WithEndpoint(") &&
		strings.Contains(content, "%s-") &&
		strings.Contains(content, ".googleapis.com") {
		info.DetectedRegionalEndpoint = "prefix"
	}
}

func collectFileImportsByAlias(f *ast.File) map[string]string {
	importsByAlias := make(map[string]string)
	for _, imp := range f.Imports {
		if imp.Path == nil {
			continue
		}
		impPath := strings.Trim(imp.Path.Value, `"`)
		alias := path.Base(impPath)
		if imp.Name != nil && imp.Name.Name != "" {
			alias = imp.Name.Name
		}
		importsByAlias[alias] = impPath
	}
	return importsByAlias
}

func recordSpecMapperDecl(fn *ast.FuncDecl, f *ast.File, fileDir string, importsByAlias map[string]string, scan *directDirScan, info *KRMPackageInfo) {
	sfx := ""
	if fn.Name.Name == scan.versionedSpecToProto {
		sfx = "_" + scan.version
	}
	if scan.targetDir == "" || fileDir == scan.targetDir {
		scan.foundSpecMapperInTarget = true
		scan.targetSpecSuffix = sfx
	} else {
		scan.externalMapperDir = fileDir
		if f.Name != nil {
			scan.externalMapperPkg = f.Name.Name
		}
		scan.externalSpecSuffix = sfx
	}
	if pbPath := extractReturnProtoImport(fn, importsByAlias); pbPath != "" {
		info.DetectedPBImportPath = pbPath
		info.DetectedPBAlias = path.Base(pbPath)
	}
}

func extractReturnProtoImport(fn *ast.FuncDecl, importsByAlias map[string]string) string {
	if fn.Type == nil || fn.Type.Results == nil || len(fn.Type.Results.List) == 0 {
		return ""
	}
	star, ok := fn.Type.Results.List[0].Type.(*ast.StarExpr)
	if !ok {
		return ""
	}
	sel, ok := star.X.(*ast.SelectorExpr)
	if !ok {
		return ""
	}
	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return ""
	}
	return importsByAlias[pkgIdent.Name]
}

func recordObsMapperDecl(fn *ast.FuncDecl, fileDir string, isGeneratedMapper bool, scan *directDirScan) {
	sfx := ""
	if fn.Name.Name == scan.versionedObsFromProto {
		sfx = "_" + scan.version
	}
	if scan.targetDir == "" || fileDir == scan.targetDir {
		if !scan.foundObsMapperInTarget || (!isGeneratedMapper && !scan.foundHandwrittenObsInTarget) {
			scan.foundObsMapperInTarget = true
			scan.targetObsSuffix = sfx
			if !isGeneratedMapper {
				scan.foundHandwrittenObsInTarget = true
			}
		}
		return
	}
	scan.foundObsMapperExternal = true
	scan.externalObsSuffix = sfx
}

func applyDirectScanToInfo(scan *directDirScan, info *KRMPackageInfo) {
	if scan.foundSpecMapperInTarget {
		info.MapperVersionSuffix = scan.targetSpecSuffix
		info.HasExplicitSpecSuffix = true
		if scan.foundObsMapperInTarget {
			info.ObservedStateVersionSuffix = scan.targetObsSuffix
			info.HasExplicitObsSuffix = true
		} else if scan.foundObsMapperExternal {
			info.ObservedStateVersionSuffix = scan.externalObsSuffix
			info.HasExplicitObsSuffix = true
		}
	} else if scan.externalMapperDir != "" {
		info.MapperVersionSuffix = scan.externalSpecSuffix
		info.HasExplicitSpecSuffix = true
		if scan.foundObsMapperExternal {
			info.ObservedStateVersionSuffix = scan.externalObsSuffix
			info.HasExplicitObsSuffix = true
		}
		if scan.targetDir != "" && scan.targetDir != scan.externalMapperDir {
			populateExternalMapperImport(scan, info)
		}
	}

	foundObsMapper := scan.foundObsMapperInTarget || scan.foundObsMapperExternal
	if info.StatusInspected && info.HasObservedState && !foundObsMapper && scan.hasGeneratedMapperFile {
		info.HasObservedState = false
	}
	if scan.hasGRPCClient && !scan.hasRESTClient {
		info.UseGRPCClient = true
	}
}

func populateExternalMapperImport(scan *directDirScan, info *KRMPackageInfo) {
	if info.RootPath != "" {
		if rel, err := filepath.Rel(info.RootPath, scan.externalMapperDir); err == nil {
			info.MapperPackageImport = "github.com/GoogleCloudPlatform/k8s-config-connector/" + filepath.ToSlash(rel)
		}
	}
	if info.MapperPackageImport == "" {
		info.MapperPackageImport = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/" + scan.serviceName
	}
	pkg := scan.externalMapperPkg
	if pkg == "" {
		pkg = filepath.Base(scan.externalMapperDir)
	}
	info.MapperPrefix = pkg + "."
}

func detectMultiversionFromGenerateScripts(apiDir, directDir, version string, info *KRMPackageInfo) {
	if info.HasExplicitSpecSuffix || info.MapperVersionSuffix != "" {
		return
	}
	for _, genShPath := range []string{
		filepath.Join(apiDir, "generate.sh"),
		filepath.Join(directDir, "generate.sh"),
	} {
		if data, err := os.ReadFile(genShPath); err == nil {
			if strings.Contains(string(data), "--multiversion") {
				info.MapperVersionSuffix = "_" + version
				return
			}
		}
	}
}

// --- Shared AST Parsing & Traversal Helpers (Footnotes) ---

func parseGoFilesInDir(dir string) []*ast.File {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	fset := token.NewFileSet()
	var files []*ast.File
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		f, err := parser.ParseFile(fset, filepath.Join(dir, entry.Name()), nil, 0)
		if err == nil {
			files = append(files, f)
		}
	}
	return files
}

func inspectIdentityFileTraits(files []*ast.File, chosenIdentity string) (usesMetadataName bool, isClientSpecified bool) {
	targetFile := findFileDefiningType(files, chosenIdentity)
	if targetFile == nil {
		return false, false
	}

	hasStrictEq := false
	hasExternalRef := false
	ast.Inspect(targetFile, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.SelectorExpr:
			if x.Sel != nil && x.Sel.Name == "ExternalRef" {
				hasExternalRef = true
			}
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok && sel.Sel != nil {
				if sel.Sel.Name == "GetResourceID" || sel.Sel.Name == "GetName" {
					usesMetadataName = true
				}
			}
		case *ast.BinaryExpr:
			if x.Op == token.NEQ && isIdentityEqualityOperand(x.X) {
				hasStrictEq = true
			}
		}
		return true
	})
	return usesMetadataName, usesMetadataName && (hasStrictEq || !hasExternalRef)
}

func findFileDefiningType(files []*ast.File, typeName string) *ast.File {
	for _, f := range files {
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if ok && ts.Name != nil && ts.Name.Name == typeName {
					return f
				}
			}
		}
	}
	return nil
}

func isIdentityEqualityOperand(expr ast.Expr) bool {
	if call, ok := expr.(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel != nil && sel.Sel.Name == "String" {
			return true
		}
	}
	if ident, ok := expr.(*ast.Ident); ok && ident.Name == "actualResourceID" {
		return true
	}
	return false
}

func collectDeclsFromFiles(
	files []*ast.File,
	structDecls map[string]*ast.StructType,
	methodsByRecv map[string]map[string]*ast.FuncDecl,
	funcDecls map[string]*ast.FuncDecl,
	gcpURLTemplates map[string]string,
) {
	for _, f := range files {
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				collectGenDecl(d, structDecls, gcpURLTemplates)
			case *ast.FuncDecl:
				collectFuncDecl(d, methodsByRecv, funcDecls)
			}
		}
	}
}

func collectGenDecl(d *ast.GenDecl, structDecls map[string]*ast.StructType, gcpURLTemplates map[string]string) {
	switch d.Tok {
	case token.TYPE:
		for _, spec := range d.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if st, ok := ts.Type.(*ast.StructType); ok {
				structDecls[ts.Name.Name] = st
			}
		}
	case token.VAR:
		for _, spec := range d.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok || len(vs.Values) == 0 {
				continue
			}
			if idType, pattern := extractGCPURLTemplateCall(vs.Values[0]); idType != "" && pattern != "" {
				gcpURLTemplates[idType] = pattern
			}
		}
	}
}

func collectFuncDecl(d *ast.FuncDecl, methodsByRecv map[string]map[string]*ast.FuncDecl, funcDecls map[string]*ast.FuncDecl) {
	if d.Recv == nil || len(d.Recv.List) == 0 {
		if d.Name != nil {
			funcDecls[d.Name.Name] = d
		}
		return
	}
	recvName := receiverTypeName(d.Recv.List[0].Type)
	if recvName == "" {
		return
	}
	if methodsByRecv[recvName] == nil {
		methodsByRecv[recvName] = make(map[string]*ast.FuncDecl)
	}
	methodsByRecv[recvName][d.Name.Name] = d
}

func extractGCPURLTemplateCall(expr ast.Expr) (string, string) {
	call, ok := expr.(*ast.CallExpr)
	if !ok || len(call.Args) < 2 {
		return "", ""
	}
	idxExpr, ok := call.Fun.(*ast.IndexExpr)
	if !ok {
		return "", ""
	}
	sel, ok := idxExpr.X.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil || sel.Sel.Name != "Template" {
		return "", ""
	}
	idIdent, ok := idxExpr.Index.(*ast.Ident)
	if !ok {
		return "", ""
	}
	patLit, ok := call.Args[1].(*ast.BasicLit)
	if !ok || patLit.Kind != token.STRING {
		return "", ""
	}
	return idIdent.Name, strings.Trim(patLit.Value, `"`)
}

func deriveParentExprFromGCPURLPattern(pattern string, exportedStringFields []string) (string, bool) {
	parts := strings.Split(pattern, "/")
	if len(parts) <= 2 {
		return "", true
	}
	parentParts := parts[:len(parts)-2]
	var fmtSegments []string
	var fieldArgs []string
	fieldIdx := 0
	for _, seg := range parentParts {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			if fieldIdx >= len(exportedStringFields) {
				return "", false
			}
			fmtSegments = append(fmtSegments, "%s")
			fieldArgs = append(fieldArgs, "a.id."+exportedStringFields[fieldIdx])
			fieldIdx++
		} else {
			fmtSegments = append(fmtSegments, seg)
		}
	}
	if len(fieldArgs) == 0 {
		return fmt.Sprintf("%q", strings.Join(fmtSegments, "/")), false
	}
	return fmt.Sprintf("fmt.Sprintf(%q, %s)", strings.Join(fmtSegments, "/"), strings.Join(fieldArgs, ", ")), false
}

func structHasResourceReferences(structName string, structDecls map[string]*ast.StructType, visited map[string]bool) bool {
	if visited[structName] {
		return false
	}
	visited[structName] = true
	st := structDecls[structName]
	if st == nil || st.Fields == nil {
		return false
	}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			if isNonRootRefFieldName(name.Name) {
				return true
			}
		}
		if elemName := baseIdentTypeName(field.Type); elemName != "" {
			if structHasResourceReferences(elemName, structDecls, visited) {
				return true
			}
		}
	}
	return false
}

func isNonRootRefFieldName(name string) bool {
	switch name {
	case "ProjectRef", "FolderRef", "OrganizationRef", "BillingAccountRef":
		return false
	}
	return strings.HasSuffix(name, "Ref") || strings.HasSuffix(name, "Refs")
}

func receiverTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

func baseIdentTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return baseIdentTypeName(t.X)
	case *ast.ArrayType:
		return baseIdentTypeName(t.Elt)
	}
	return ""
}
