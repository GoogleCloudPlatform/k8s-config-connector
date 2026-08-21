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

package main

import (
	"context"
	"flag"
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type TypeCache struct {
	packages map[string]*packages.Package
}

func NewTypeCache() *TypeCache {
	return &TypeCache{
		packages: make(map[string]*packages.Package),
	}
}

func (c *TypeCache) GetType(pkgPath, typeName string) (types.Type, error) {
	pkg, ok := c.packages[pkgPath]
	if !ok {
		cfg := &packages.Config{
			Mode: packages.NeedName | packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		}
		pkgs, err := packages.Load(cfg, pkgPath)
		if err != nil {
			return nil, fmt.Errorf("loading package %s: %w", pkgPath, err)
		}
		if len(pkgs) == 0 {
			return nil, fmt.Errorf("no package found for %s", pkgPath)
		}
		if len(pkgs[0].Errors) > 0 {
			// Log package errors as warnings instead of failing, as typechecking still resolves structs
			for _, e := range pkgs[0].Errors {
				log.Printf("Warning: package error: %v", e.Msg)
			}
		}
		pkg = pkgs[0]
		c.packages[pkgPath] = pkg
	}

	obj := pkg.Types.Scope().Lookup(typeName)
	if obj == nil {
		return nil, fmt.Errorf("type %s not found in package %s", typeName, pkgPath)
	}
	return obj.Type(), nil
}

type typePair struct {
	KRMType     *types.Named
	APIType     *types.Named
	KRMTypeName string
	APITypeName string
}

func main() {
	var mappers stringSlice
	var outputFile string

	flag.Var(&mappers, "mapper", "Fully-qualified KRM type and API type mapping, e.g. --mapper KRM_FQ_TYPE::API_FQ_TYPE")
	flag.StringVar(&outputFile, "output-file", "", "Path to the output generated mapper go file")
	flag.Parse()

	if len(mappers) == 0 || outputFile == "" {
		log.Fatalf("Flags (--mapper, --output-file) are required")
	}

	ctx := context.Background()
	if err := Run(ctx, mappers, outputFile); err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}

func Run(ctx context.Context, mappers []string, outputFile string) error {
	cache := NewTypeCache()

	var queue []typePair
	var visited = make(map[string]bool)
	var pairsToGenerate []typePair

	var krmPackagePath string
	apiPackagePaths := make(map[string]bool)

	for _, m := range mappers {
		parts := strings.Split(m, "::")
		if len(parts) != 2 {
			return fmt.Errorf("invalid mapper format %q, expected KRM_FQ_TYPE::API_FQ_TYPE", m)
		}
		krmFQ, apiFQ := parts[0], parts[1]

		krmPkg, krmName, err := parseFullyQualifiedType(krmFQ)
		if err != nil {
			return err
		}
		apiPkg, apiName, err := parseFullyQualifiedType(apiFQ)
		if err != nil {
			return err
		}

		if krmPackagePath == "" {
			krmPackagePath = krmPkg
		} else if krmPackagePath != krmPkg {
			return fmt.Errorf("all KRM types must be from the same package, but got %s and %s", krmPackagePath, krmPkg)
		}

		apiPackagePaths[apiPkg] = true

		krmType, err := cache.GetType(krmPkg, krmName)
		if err != nil {
			return err
		}
		apiType, err := cache.GetType(apiPkg, apiName)
		if err != nil {
			return err
		}

		krmNamed, ok := krmType.(*types.Named)
		if !ok {
			return fmt.Errorf("KRM type %s is not a named type", krmName)
		}
		apiNamed, ok := apiType.(*types.Named)
		if !ok {
			return fmt.Errorf("API type %s is not a named type", apiName)
		}

		key := krmFQ + "::" + apiFQ
		if !visited[key] {
			visited[key] = true
			pair := typePair{
				KRMType:     krmNamed,
				APIType:     apiNamed,
				KRMTypeName: krmName,
				APITypeName: apiName,
			}
			queue = append(queue, pair)
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		pairsToGenerate = append(pairsToGenerate, p)

		krmStruct, ok := p.KRMType.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		apiStruct, ok := p.APIType.Underlying().(*types.Struct)
		if !ok {
			continue
		}

		krmFields := getStructFields(krmStruct)
		apiFields := getStructFields(apiStruct)

		for _, kf := range krmFields {
			var matchedAPI *types.Var
			for _, af := range apiFields {
				if matchKRMAndAPIFields(kf, af) {
					matchedAPI = af
					break
				}
			}
			if matchedAPI == nil {
				continue
			}

			if isRefType(kf.Type()) {
				continue
			}

			krmSub, krmOk := resolveUnderlyingNamedStruct(kf.Type())
			apiSub, apiOk := resolveUnderlyingNamedStruct(matchedAPI.Type())

			if krmOk && apiOk {
				krmSubFQ := krmSub.Obj().Pkg().Path() + "/" + krmSub.Obj().Name()
				apiSubFQ := apiSub.Obj().Pkg().Path() + "/" + apiSub.Obj().Name()
				key := krmSubFQ + "::" + apiSubFQ
				if !visited[key] {
					visited[key] = true
					apiPackagePaths[apiSub.Obj().Pkg().Path()] = true
					queue = append(queue, typePair{
						KRMType:     krmSub,
						APIType:     apiSub,
						KRMTypeName: krmSub.Obj().Name(),
						APITypeName: apiSub.Obj().Name(),
					})
				}
			}
		}
	}

	outDir := filepath.Dir(outputFile)
	outPackageName := detectPackageName(outDir)
	outPackagePath := goPackagePathForDir(outDir)

	var buf strings.Builder
	buf.WriteString("// Copyright 2026 Google LLC\n")
	buf.WriteString("//\n")
	buf.WriteString("// Licensed under the Apache License, Version 2.0 (the \"License\");\n")
	buf.WriteString("// you may not use this file except in compliance with the License.\n")
	buf.WriteString("// You may obtain a copy of the License at\n")
	buf.WriteString("//\n")
	buf.WriteString("//      http://www.apache.org/licenses/LICENSE-2.0\n")
	buf.WriteString("//\n")
	buf.WriteString("// Unless required by applicable law or agreed to in writing, software\n")
	buf.WriteString("// distributed under the License is distributed on an \"AS IS\" BASIS,\n")
	buf.WriteString("// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n")
	buf.WriteString("// See the License for the specific language governing permissions and\n")
	buf.WriteString("// limitations under the License.\n\n")

	buf.WriteString("// Code generated by dev/tools/openapi-to-krm. DO NOT EDIT.\n")
	buf.WriteString("//go:build !ignore_autogenerated\n")
	buf.WriteString("// +build !ignore_autogenerated\n\n")
	buf.WriteString("// +generated:mapper\n\n")

	fmt.Fprintf(&buf, "package %s\n\n", outPackageName)

	buf.WriteString("import (\n")
	fmt.Fprintf(&buf, "\t\"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct\"\n")

	if krmPackagePath != outPackagePath {
		fmt.Fprintf(&buf, "\tkrm \"%s\"\n", krmPackagePath)
	}

	var sortedAPIPkgs []string
	for p := range apiPackagePaths {
		sortedAPIPkgs = append(sortedAPIPkgs, p)
	}
	sort.Strings(sortedAPIPkgs)

	apiAliasMap := make(map[string]string)
	for i, p := range sortedAPIPkgs {
		alias := "api"
		if i > 0 {
			alias = fmt.Sprintf("api%d", i)
		}
		apiAliasMap[p] = alias
		fmt.Fprintf(&buf, "\t%s \"%s\"\n", alias, p)
	}

	buf.WriteString(")\n\n")

	for _, p := range pairsToGenerate {
		writeFromAPIMapper(&buf, p, krmPackagePath, outPackagePath, apiAliasMap, outDir)
		writeToAPIMapper(&buf, p, krmPackagePath, outPackagePath, apiAliasMap, outDir)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("creating directory %s: %w", outDir, err)
	}

	if err := os.WriteFile(outputFile, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing output file: %w", err)
	}

	fmt.Printf("Successfully generated mappers to %s\n", outputFile)
	return nil
}

func parseFullyQualifiedType(fqType string) (string, string, error) {
	idx := strings.LastIndex(fqType, "/")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid fully-qualified type (no slash): %s", fqType)
	}
	pkgPath := fqType[:idx]
	typeName := fqType[idx+1:]
	return pkgPath, typeName, nil
}

func detectPackageName(dir string) string {
	files, err := os.ReadDir(dir)
	if err == nil {
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") {
				b, err := os.ReadFile(filepath.Join(dir, f.Name()))
				if err == nil {
					for _, line := range strings.Split(string(b), "\n") {
						line = strings.TrimSpace(line)
						if strings.HasPrefix(line, "package ") {
							parts := strings.Fields(line)
							if len(parts) >= 2 {
								return strings.TrimSuffix(parts[1], ";")
							}
						}
					}
				}
			}
		}
	}
	return filepath.Base(dir)
}

func goPackagePathForDir(dir string) string {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return ""
	}
	rel, err := filepath.Rel("/workspaces/k8s-config-connector", absDir)
	if err != nil {
		return ""
	}
	if rel == "." {
		return "github.com/GoogleCloudPlatform/k8s-config-connector"
	}
	return "github.com/GoogleCloudPlatform/k8s-config-connector/" + filepath.ToSlash(rel)
}

func getStructFields(s *types.Struct) []*types.Var {
	var fields []*types.Var
	for i := 0; i < s.NumFields(); i++ {
		f := s.Field(i)
		if f.Embedded() {
			if named, ok := f.Type().(*types.Named); ok {
				if str, ok := named.Underlying().(*types.Struct); ok {
					fields = append(fields, getStructFields(str)...)
				}
			}
		} else {
			fields = append(fields, f)
		}
	}
	return fields
}

func resolveUnderlyingNamedStruct(t types.Type) (*types.Named, bool) {
	curr := t
	if ptr, ok := curr.(*types.Pointer); ok {
		curr = ptr.Elem()
	}
	if slice, ok := curr.(*types.Slice); ok {
		curr = slice.Elem()
		if ptr, ok := curr.(*types.Pointer); ok {
			curr = ptr.Elem()
		}
	}
	if named, ok := curr.(*types.Named); ok {
		if _, ok := named.Underlying().(*types.Struct); ok {
			return named, true
		}
	}
	return nil, false
}

type TypeAnalysis struct {
	IsPtr        bool
	IsSlice      bool
	SliceElemPtr bool
	IsMap        bool
	MapKey       types.Type
	MapVal       types.Type
	Elem         types.Type
}

func analyzeType(t types.Type) TypeAnalysis {
	var res TypeAnalysis
	curr := t
	if ptr, ok := curr.(*types.Pointer); ok {
		res.IsPtr = true
		curr = ptr.Elem()
	}
	if slice, ok := curr.(*types.Slice); ok {
		res.IsSlice = true
		curr = slice.Elem()
		if ptr, ok := curr.(*types.Pointer); ok {
			res.SliceElemPtr = true
			curr = ptr.Elem()
		}
	}
	if m, ok := curr.(*types.Map); ok {
		res.IsMap = true
		res.MapKey = m.Key()
		res.MapVal = m.Elem()
	}
	res.Elem = curr
	return res
}

func goTypeName(t types.Type, krmPkgPath, outPkgPath string, apiAliasMap map[string]string) string {
	qualifier := func(p *types.Package) string {
		if p.Path() == outPkgPath {
			return ""
		}
		if p.Path() == krmPkgPath {
			return "krm"
		}
		if alias, ok := apiAliasMap[p.Path()]; ok {
			return alias
		}
		return p.Name()
	}
	return types.TypeString(t, qualifier)
}

func findFuncDeclaration(goFuncName string, srcDir string) bool {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return false
	}

	for _, f := range files {
		p := filepath.Join(srcDir, f.Name())
		if !strings.HasSuffix(p, ".go") {
			continue
		}
		if strings.Contains(f.Name(), "generated") || strings.HasPrefix(f.Name(), "zz_") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			continue
		}

		for _, line := range strings.Split(string(b), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "func "+goFuncName+"(") {
				return true
			}
		}
	}

	return false
}

func mapBasicType(valExpr string, fromType, toType types.Type) string {
	fromAn := analyzeType(fromType)
	toAn := analyzeType(toType)

	fromElem := fromAn.Elem
	toElem := toAn.Elem

	needsCast := fromElem.String() != toElem.String()

	expr := valExpr
	if needsCast {
		if fromAn.IsPtr {
			expr = "direct.ValueOf(" + expr + ")"
		}
		expr = toElem.String() + "(" + expr + ")"
		if toAn.IsPtr {
			expr = "direct.LazyPtr(" + expr + ")"
		}
		return expr
	}

	if toAn.IsPtr && !fromAn.IsPtr {
		return "direct.LazyPtr(" + expr + ")"
	}
	if !toAn.IsPtr && fromAn.IsPtr {
		return "direct.ValueOf(" + expr + ")"
	}
	return expr
}

func matchKRMAndAPIFields(kf *types.Var, af *types.Var) bool {
	if strings.EqualFold(kf.Name(), af.Name()) {
		return true
	}
	if strings.HasSuffix(kf.Name(), "Ref") {
		kfBase := strings.TrimSuffix(kf.Name(), "Ref")
		if strings.EqualFold(kfBase, af.Name()) {
			return true
		}
		if strings.EqualFold(kfBase, strings.TrimSuffix(af.Name(), "Url")) {
			return true
		}
		if strings.EqualFold(kfBase, strings.TrimSuffix(af.Name(), "Name")) {
			return true
		}
	}
	return false
}

func isRefType(t types.Type) bool {
	curr := t
	if ptr, ok := curr.(*types.Pointer); ok {
		curr = ptr.Elem()
	}
	named, ok := curr.(*types.Named)
	if !ok {
		return false
	}
	str, ok := named.Underlying().(*types.Struct)
	if !ok {
		return false
	}
	for i := 0; i < str.NumFields(); i++ {
		if str.Field(i).Name() == "External" {
			if basic, ok := str.Field(i).Type().(*types.Basic); ok && basic.Kind() == types.String {
				return true
			}
		}
	}
	return false
}

func writeFromAPIMapper(buf *strings.Builder, p typePair, krmPkgPath, outPkgPath string, apiAliasMap map[string]string, outDir string) {
	krmName := p.KRMTypeName

	krmTypeStr := goTypeName(p.KRMType, krmPkgPath, outPkgPath, apiAliasMap)
	apiTypeStr := goTypeName(p.APIType, krmPkgPath, outPkgPath, apiAliasMap)

	funcName := krmName + "_FromAPI"

	exists := findFuncDeclaration(funcName, outDir)
	if exists {
		fmt.Fprintf(buf, "/* found existing non-generated mapping function %q, skipping\n", funcName)
	}

	fmt.Fprintf(buf, "func %s(mapCtx *direct.MapContext, in *%s) *%s {\n", funcName, apiTypeStr, krmTypeStr)
	buf.WriteString("\tif in == nil {\n\t\treturn nil\n\t}\n")
	fmt.Fprintf(buf, "\tout := &%s{}\n", krmTypeStr)

	krmStruct := p.KRMType.Underlying().(*types.Struct)
	apiStruct := p.APIType.Underlying().(*types.Struct)

	krmFields := getStructFields(krmStruct)
	apiFields := getStructFields(apiStruct)

	mappedKRMFields := make(map[string]bool)

	for _, kf := range krmFields {
		var af *types.Var
		for _, f := range apiFields {
			if matchKRMAndAPIFields(kf, f) {
				af = f
				break
			}
		}
		if af == nil {
			continue
		}

		mappedKRMFields[kf.Name()] = true

		krmAn := analyzeType(kf.Type())
		apiAn := analyzeType(af.Type())

		if isRefType(kf.Type()) {
			valExpr := "in." + af.Name()
			if apiAn.IsPtr {
				valExpr = "direct.ValueOf(in." + af.Name() + ")"
			}
			if krmAn.IsPtr {
				refTypeName := goTypeName(krmAn.Elem, krmPkgPath, outPkgPath, apiAliasMap)
				fmt.Fprintf(buf, "\tif %s != \"\" {\n", valExpr)
				fmt.Fprintf(buf, "\t\tout.%s = &%s{External: %s}\n", kf.Name(), refTypeName, valExpr)
				fmt.Fprintf(buf, "\t}\n")
			} else {
				fmt.Fprintf(buf, "\tout.%s.External = %s\n", kf.Name(), valExpr)
			}
			continue
		}

		stmt := ""
		if krmAn.IsSlice && apiAn.IsSlice {
			krmSub, krmOk := resolveUnderlyingNamedStruct(kf.Type())
			_, apiOk := resolveUnderlyingNamedStruct(af.Type())
			if krmOk && apiOk {
				mapperFuncName := krmSub.Obj().Name() + "_FromAPI"
				stmt = fmt.Sprintf("direct.Slice_FromProto(mapCtx, in.%s, %s)", af.Name(), mapperFuncName)
			} else {
				stmt = "in." + af.Name()
			}
		} else if krmAn.IsMap && apiAn.IsMap {
			stmt = "in." + af.Name()
		} else if _, ok := krmAn.Elem.(*types.Named); ok {
			krmSub, krmOk := resolveUnderlyingNamedStruct(kf.Type())
			if krmOk {
				mapperFuncName := krmSub.Obj().Name() + "_FromAPI"
				stmt = fmt.Sprintf("%s(mapCtx, in.%s)", mapperFuncName, af.Name())
			}
		} else {
			stmt = mapBasicType("in."+af.Name(), af.Type(), kf.Type())
		}

		if stmt != "" {
			fmt.Fprintf(buf, "\tout.%s = %s\n", kf.Name(), stmt)
		}
	}

	for _, kf := range krmFields {
		if !mappedKRMFields[kf.Name()] {
			if kf.Name() == "ResourceID" {
				hasName := false
				for _, af := range apiFields {
					if af.Name() == "Name" {
						hasName = true
						break
					}
				}
				if hasName {
					fmt.Fprintf(buf, "\tout.ResourceID = direct.LazyPtr(in.Name)\n")
					continue
				}
			}
			fmt.Fprintf(buf, "\t// MISSING: %s\n", kf.Name())
		}
	}

	buf.WriteString("\treturn out\n}\n")
	if exists {
		buf.WriteString("*/\n\n")
	} else {
		buf.WriteString("\n")
	}
}

func writeToAPIMapper(buf *strings.Builder, p typePair, krmPkgPath, outPkgPath string, apiAliasMap map[string]string, outDir string) {
	krmName := p.KRMTypeName

	krmTypeStr := goTypeName(p.KRMType, krmPkgPath, outPkgPath, apiAliasMap)
	apiTypeStr := goTypeName(p.APIType, krmPkgPath, outPkgPath, apiAliasMap)

	funcName := krmName + "_ToAPI"

	exists := findFuncDeclaration(funcName, outDir)
	if exists {
		fmt.Fprintf(buf, "/* found existing non-generated mapping function %q, skipping\n", funcName)
	}

	fmt.Fprintf(buf, "func %s(mapCtx *direct.MapContext, in *%s) *%s {\n", funcName, krmTypeStr, apiTypeStr)
	buf.WriteString("\tif in == nil {\n\t\treturn nil\n\t}\n")
	fmt.Fprintf(buf, "\tout := &%s{}\n", apiTypeStr)

	krmStruct := p.KRMType.Underlying().(*types.Struct)
	apiStruct := p.APIType.Underlying().(*types.Struct)

	krmFields := getStructFields(krmStruct)
	apiFields := getStructFields(apiStruct)

	mappedAPIFields := make(map[string]bool)

	for _, kf := range krmFields {
		var af *types.Var
		for _, f := range apiFields {
			if matchKRMAndAPIFields(kf, f) {
				af = f
				break
			}
		}
		if af == nil {
			continue
		}

		mappedAPIFields[af.Name()] = true

		krmAn := analyzeType(kf.Type())
		apiAn := analyzeType(af.Type())

		if isRefType(kf.Type()) {
			valExpr := "in." + kf.Name() + ".External"
			if krmAn.IsPtr {
				fmt.Fprintf(buf, "\tif in.%s != nil {\n", kf.Name())
				if apiAn.IsPtr {
					fmt.Fprintf(buf, "\t\tout.%s = direct.LazyPtr(%s)\n", af.Name(), valExpr)
				} else {
					fmt.Fprintf(buf, "\t\tout.%s = %s\n", af.Name(), valExpr)
				}
				fmt.Fprintf(buf, "\t}\n")
			} else {
				if apiAn.IsPtr {
					fmt.Fprintf(buf, "\tout.%s = direct.LazyPtr(%s)\n", af.Name(), valExpr)
				} else {
					fmt.Fprintf(buf, "\tout.%s = %s\n", af.Name(), valExpr)
				}
			}
			continue
		}

		stmt := ""
		if krmAn.IsSlice && apiAn.IsSlice {
			krmSub, krmOk := resolveUnderlyingNamedStruct(kf.Type())
			_, apiOk := resolveUnderlyingNamedStruct(af.Type())
			if krmOk && apiOk {
				mapperFuncName := krmSub.Obj().Name() + "_ToAPI"
				stmt = fmt.Sprintf("direct.Slice_ToProto(mapCtx, in.%s, %s)", kf.Name(), mapperFuncName)
			} else {
				stmt = "in." + kf.Name()
			}
		} else if krmAn.IsMap && apiAn.IsMap {
			stmt = "in." + kf.Name()
		} else if _, ok := krmAn.Elem.(*types.Named); ok {
			krmSub, krmOk := resolveUnderlyingNamedStruct(kf.Type())
			if krmOk {
				mapperFuncName := krmSub.Obj().Name() + "_ToAPI"
				stmt = fmt.Sprintf("%s(mapCtx, in.%s)", mapperFuncName, kf.Name())
			}
		} else {
			stmt = mapBasicType("in."+kf.Name(), kf.Type(), af.Type())
		}

		if stmt != "" {
			fmt.Fprintf(buf, "\tout.%s = %s\n", af.Name(), stmt)
		}
	}

	for _, af := range apiFields {
		if af.Name() == "ForceSendFields" || af.Name() == "NullFields" || af.Name() == "ServerResponse" || af.Name() == "Kind" {
			continue
		}
		if !mappedAPIFields[af.Name()] {
			if af.Name() == "Name" {
				hasResourceID := false
				for _, kf := range krmFields {
					if kf.Name() == "ResourceID" {
						hasResourceID = true
						break
					}
				}
				if hasResourceID {
					fmt.Fprintf(buf, "\tout.Name = direct.ValueOf(in.ResourceID)\n")
					continue
				}
			}
			fmt.Fprintf(buf, "\t// MISSING: %s\n", af.Name())
		}
	}

	buf.WriteString("\treturn out\n}\n")
	if exists {
		buf.WriteString("*/\n\n")
	} else {
		buf.WriteString("\n")
	}
}
