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
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"sigs.k8s.io/yaml"
)

type FixtureInfo struct {
	Dir        string
	FolderName string
	Kind       string
	CreateYAML []byte
	UpdateYAML []byte
	CreateObj  map[string]any
	UpdateObj  map[string]any
}

type TargetResolver struct {
	RepoRoot string
	Fixtures []*FixtureInfo

	// KindToFixtures maps Kind -> list of FixtureInfo
	KindToFixtures map[string][]*FixtureInfo

	// PackageStructs: pkg path -> map[structName]*ast.StructType
	PackageStructs map[string]map[string]*ast.StructType
	PackageAliases map[string]map[string]string // struct name -> target type name
	KindSpecMap    map[string]KindSpecInfo
}

type KindTarget struct {
	All   bool     `json:"all"`
	Tests []string `json:"tests"`
}

type KindSpecInfo struct {
	PkgPath  string
	SpecType string
}

func main() {
	var baseRef string
	var jsonOutput bool

	flag.StringVar(&baseRef, "base", "", "Base git revision to compare against (default: auto-detected merge-base with master)")
	flag.BoolVar(&jsonOutput, "json", false, "Output results as JSON instead of YAML")
	flag.Parse()

	repoRoot, err := findRepoRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding repo root: %v\n", err)
		os.Exit(1)
	}

	resolver, err := NewTargetResolver(repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing target resolver: %v\n", err)
		os.Exit(1)
	}

	changedFiles, err := resolver.GetChangedFiles(baseRef)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting changed files: %v\n", err)
		os.Exit(1)
	}

	targets, err := resolver.ResolveTargets(changedFiles)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving test targets: %v\n", err)
		os.Exit(1)
	}

	if jsonOutput {
		data, err := yaml.Marshal(targets)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error formatting output: %v\n", err)
			os.Exit(1)
		}
		var obj any
		_ = yaml.Unmarshal(data, &obj)
		jsonBytes, _ := yaml.Marshal(obj)
		fmt.Println(string(jsonBytes))
	} else {
		data, err := yaml.Marshal(targets)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error formatting output: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(data))
	}
}

func findRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func NewTargetResolver(repoRoot string) (*TargetResolver, error) {
	r := &TargetResolver{
		RepoRoot:       repoRoot,
		KindToFixtures: make(map[string][]*FixtureInfo),
		PackageStructs: make(map[string]map[string]*ast.StructType),
		PackageAliases: make(map[string]map[string]string),
		KindSpecMap:    make(map[string]KindSpecInfo),
	}

	if err := r.loadFixtures(); err != nil {
		return nil, fmt.Errorf("loading fixtures: %w", err)
	}

	if err := r.indexAPITypes(); err != nil {
		return nil, fmt.Errorf("indexing API types: %w", err)
	}

	return r, nil
}

func (r *TargetResolver) loadFixtures() error {
	baseDir := filepath.Join(r.RepoRoot, "pkg", "test", "resourcefixture", "testdata", "basic")
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return nil
	}

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if d.Name() != "create.yaml" {
			return nil
		}

		fixtureDir := filepath.Dir(path)
		createBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var createObj map[string]any
		if err := yaml.Unmarshal(createBytes, &createObj); err != nil {
			return nil // Skip unparseable yaml
		}

		kind, _ := createObj["kind"].(string)
		if kind == "" {
			return nil
		}

		folderName := filepath.Base(fixtureDir)
		info := &FixtureInfo{
			Dir:        fixtureDir,
			FolderName: folderName,
			Kind:       kind,
			CreateYAML: createBytes,
			CreateObj:  createObj,
		}

		updatePath := filepath.Join(fixtureDir, "update.yaml")
		if updateBytes, err := os.ReadFile(updatePath); err == nil {
			info.UpdateYAML = updateBytes
			var updateObj map[string]any
			if err := yaml.Unmarshal(updateBytes, &updateObj); err == nil {
				info.UpdateObj = updateObj
			}
		}

		r.Fixtures = append(r.Fixtures, info)
		r.KindToFixtures[kind] = append(r.KindToFixtures[kind], info)
		return nil
	})

	return err
}

func (r *TargetResolver) indexAPITypes() error {
	apisDir := filepath.Join(r.RepoRoot, "apis")
	fset := token.NewFileSet()

	err := filepath.WalkDir(apisDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}

		pkgs, err := parser.ParseDir(fset, path, func(fi fs.FileInfo) bool {
			name := fi.Name()
			if strings.HasSuffix(name, "_test.go") || strings.HasPrefix(name, "zz_generated.") {
				return false
			}
			return strings.HasSuffix(name, ".go")
		}, 0)
		if err != nil || len(pkgs) == 0 {
			return nil
		}

		pkgStructs := make(map[string]*ast.StructType)
		pkgAliases := make(map[string]string)

		for _, pkg := range pkgs {
			for _, file := range pkg.Files {
				for _, decl := range file.Decls {
					genDecl, ok := decl.(*ast.GenDecl)
					if !ok || genDecl.Tok != token.TYPE {
						continue
					}
					for _, spec := range genDecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}
						typeName := typeSpec.Name.Name
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							pkgStructs[typeName] = structType

							// Check if this struct is a Kind definition (has Spec field and TypeMeta)
							if isKindStruct(structType) {
								specFieldType := getSpecFieldTypeName(structType)
								if specFieldType != "" {
									r.KindSpecMap[typeName] = KindSpecInfo{
										PkgPath:  path,
										SpecType: specFieldType,
									}
								}
							}
						} else if ident, ok := typeSpec.Type.(*ast.Ident); ok {
							pkgAliases[typeName] = ident.Name
						} else if sel, ok := typeSpec.Type.(*ast.SelectorExpr); ok {
							pkgAliases[typeName] = sel.Sel.Name
						}
					}
				}
			}
		}

		if len(pkgStructs) > 0 {
			r.PackageStructs[path] = pkgStructs
			r.PackageAliases[path] = pkgAliases
		}

		return nil
	})

	return err
}

func isKindStruct(st *ast.StructType) bool {
	hasTypeMeta := false
	hasSpec := false
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			// Embedded field
			if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "TypeMeta" {
				hasTypeMeta = true
			} else if sel, ok := field.Type.(*ast.SelectorExpr); ok && sel.Sel.Name == "TypeMeta" {
				hasTypeMeta = true
			}
		} else {
			for _, name := range field.Names {
				if name.Name == "Spec" {
					hasSpec = true
				}
			}
		}
	}
	return hasTypeMeta && hasSpec
}

func getSpecFieldTypeName(st *ast.StructType) string {
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			if name.Name == "Spec" {
				if ident, ok := field.Type.(*ast.Ident); ok {
					return ident.Name
				}
				if star, ok := field.Type.(*ast.StarExpr); ok {
					if ident, ok := star.X.(*ast.Ident); ok {
						return ident.Name
					}
				}
			}
		}
	}
	return ""
}

// GetChangedFiles discovers modified, staged, and untracked files in the repository.
// If baseRef is provided, it diffs against that commit/branch in addition to uncommitted changes.
// By default (when baseRef is empty), it inspects uncommitted working tree modifications (staged and unstaged)
// against HEAD and includes untracked files, designed to be run before committing updates.
func (r *TargetResolver) GetChangedFiles(baseRef string) ([]string, error) {
	fileSet := make(map[string]bool)

	// 1. If baseRef is provided, diff against baseRef
	if baseRef != "" {
		cmd := exec.Command("git", "diff", "--name-only", baseRef)
		if out, err := cmd.Output(); err == nil {
			for _, line := range strings.Split(string(out), "\n") {
				line = strings.TrimSpace(line)
				if line != "" {
					fileSet[line] = true
				}
			}
		}
	}

	// 2. Uncommitted changes (working tree & staged against HEAD)
	cmd := exec.Command("git", "diff", "--name-only", "HEAD")
	if out, err := cmd.Output(); err == nil {
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				fileSet[line] = true
			}
		}
	}

	// 3. Untracked files
	cmd = exec.Command("git", "ls-files", "--others", "--exclude-standard")
	if out, err := cmd.Output(); err == nil {
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				fileSet[line] = true
			}
		}
	}

	var files []string
	for f := range fileSet {
		files = append(files, f)
	}
	sort.Strings(files)
	return files, nil
}

func (r *TargetResolver) ResolveTargets(changedFiles []string) (map[string]KindTarget, error) {
	result := make(map[string]map[string]bool)
	allKind := make(map[string]bool)

	addResult := func(kind, folder string, isWholeKind bool) {
		if kind == "" || folder == "" {
			return
		}
		if result[kind] == nil {
			result[kind] = make(map[string]bool)
		}
		result[kind][folder] = true
		if isWholeKind {
			allKind[kind] = true
		}
	}

	var refFiles []string

	for _, file := range changedFiles {
		// Rule 1: Manifest / Fixture changes (including create.yaml, update.yaml, dependencies.yaml, _http.log)
		if strings.HasPrefix(file, "pkg/test/resourcefixture/testdata/") {
			absPath := filepath.Join(r.RepoRoot, file)
			fixtureDir := findFixtureDir(absPath)
			if fixtureDir != "" {
				createPath := filepath.Join(fixtureDir, "create.yaml")
				if data, err := os.ReadFile(createPath); err == nil {
					var obj map[string]any
					if err := yaml.Unmarshal(data, &obj); err == nil {
						if kind, ok := obj["kind"].(string); ok && kind != "" {
							addResult(kind, filepath.Base(fixtureDir), false)
						}
					}
				}
			}
			continue
		}

		// Rule 3 Check: Is this a reference file?
		if isReferenceFile(file) {
			refFiles = append(refFiles, file)
			continue
		}

		// Rule 2: Controller / Mapper / Types / CRD / Identity / Fuzzer changes (Whole-Kind)
		if kind := r.extractKindFromResourceFile(file); kind != "" {
			for _, fix := range r.KindToFixtures[kind] {
				addResult(kind, fix.FolderName, true)
			}
		}
	}

	// Process Rule 3: Reference Changes (Selective)
	if len(refFiles) > 0 {
		refTypes := r.extractRefTypesFromFiles(refFiles)
		for _, refType := range refTypes {
			kindFieldPaths := r.findFieldPathsForRef(refType)
			for kind, paths := range kindFieldPaths {
				for _, fix := range r.KindToFixtures[kind] {
					if fixtureHasRefField(fix, paths) {
						addResult(kind, fix.FolderName, false)
					}
				}
			}
		}
	}

	// Format output as map[string]KindTarget with sorted slices
	finalMap := make(map[string]KindTarget)
	for kind, folders := range result {
		var list []string
		for f := range folders {
			list = append(list, f)
		}
		sort.Strings(list)
		finalMap[kind] = KindTarget{
			All:   allKind[kind],
			Tests: list,
		}
	}

	return finalMap, nil
}

func findFixtureDir(path string) string {
	dir := path
	fi, err := os.Stat(path)
	if err == nil && !fi.IsDir() {
		dir = filepath.Dir(path)
	}

	for {
		createPath := filepath.Join(dir, "create.yaml")
		if _, err := os.Stat(createPath); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir || !strings.Contains(dir, "testdata") {
			break
		}
		dir = parent
	}
	return ""
}

func isReferenceFile(file string) bool {
	if strings.HasSuffix(file, "_reference.go") {
		return true
	}
	if strings.HasPrefix(file, "apis/refs/") {
		return true
	}
	if strings.Contains(file, "/refs/") && strings.HasSuffix(file, ".go") {
		return true
	}
	return false
}

func (r *TargetResolver) extractRefTypesFromFiles(files []string) []string {
	refSet := make(map[string]bool)
	fset := token.NewFileSet()

	for _, file := range files {
		absPath := filepath.Join(r.RepoRoot, file)
		node, err := parser.ParseFile(fset, absPath, nil, 0)
		if err != nil {
			continue
		}

		for _, decl := range node.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				name := typeSpec.Name.Name
				if strings.HasSuffix(name, "Ref") {
					refSet[name] = true
				}
			}
		}
	}

	var list []string
	for ref := range refSet {
		list = append(list, ref)
	}
	sort.Strings(list)
	return list
}

func (r *TargetResolver) extractKindFromResourceFile(file string) string {
	// 1. CRD file: config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<service>.cnrm.cloud.google.com.yaml
	if strings.HasPrefix(file, "config/crds/resources/") && strings.HasSuffix(file, ".yaml") {
		absPath := filepath.Join(r.RepoRoot, file)
		if data, err := os.ReadFile(absPath); err == nil {
			var crd struct {
				Spec struct {
					Names struct {
						Kind string `json:"kind"`
					} `json:"names"`
				} `json:"spec"`
			}
			if err := yaml.Unmarshal(data, &crd); err == nil && crd.Spec.Names.Kind != "" {
				return crd.Spec.Names.Kind
			}
		}
	}

	// 2. Direct controller / mapper / fuzzer: pkg/controller/direct/<service>/<kind_lowercase>_{controller,mappings,fuzzer}.go
	if strings.HasPrefix(file, "pkg/controller/direct/") && strings.HasSuffix(file, ".go") {
		base := filepath.Base(file)
		for _, suffix := range []string{"_controller.go", "_mappings.go", "_fuzzer.go", "_types.go"} {
			if strings.HasSuffix(base, suffix) {
				prefix := strings.TrimSuffix(base, suffix)
				return r.findKindByPrefix(prefix)
			}
		}
	}

	// 3. API types / identity: apis/<service>/<version>/<kind_lowercase>_{types,identity,identity_test,mappings}.go
	if strings.HasPrefix(file, "apis/") && strings.HasSuffix(file, ".go") {
		base := filepath.Base(file)
		for _, suffix := range []string{"_types.go", "_identity.go", "_identity_test.go", "_mappings.go"} {
			if strings.HasSuffix(base, suffix) {
				prefix := strings.TrimSuffix(base, suffix)
				return r.findKindByPrefix(prefix)
			}
		}
	}

	return ""
}

func (r *TargetResolver) findKindByPrefix(prefix string) string {
	cleanPrefix := strings.ToLower(strings.ReplaceAll(prefix, "_", ""))
	// First check known kinds in fixtures
	for kind := range r.KindToFixtures {
		if strings.ToLower(kind) == cleanPrefix {
			return kind
		}
	}
	// Check known kinds in KindSpecMap
	for kind := range r.KindSpecMap {
		if strings.ToLower(kind) == cleanPrefix {
			return kind
		}
	}
	return ""
}

func (r *TargetResolver) findFieldPathsForRef(refType string) map[string][][]string {
	result := make(map[string][][]string)

	for kind, specInfo := range r.KindSpecMap {
		pkgStructs := r.PackageStructs[specInfo.PkgPath]
		if pkgStructs == nil {
			continue
		}

		visited := make(map[string]bool)
		paths := r.findPathsInStruct(pkgStructs, specInfo.SpecType, refType, []string{"spec"}, visited)
		if len(paths) > 0 {
			result[kind] = paths
		}
	}

	return result
}

func (r *TargetResolver) findPathsInStruct(pkgStructs map[string]*ast.StructType, structName, targetRef string, currentPath []string, visited map[string]bool) [][]string {
	if visited[structName] {
		return nil
	}
	visited[structName] = true
	defer func() { visited[structName] = false }()

	st := pkgStructs[structName]
	if st == nil {
		return nil
	}

	var results [][]string

	for _, field := range st.Fields.List {
		jsonName := getJSONFieldName(field)
		if jsonName == "-" || jsonName == "" {
			continue
		}

		fieldPath := append(append([]string{}, currentPath...), jsonName)
		fieldTypeName := extractTypeName(field.Type)

		if fieldTypeName == targetRef || strings.HasSuffix(fieldTypeName, "."+targetRef) {
			results = append(results, fieldPath)
		} else if innerStruct := pkgStructs[fieldTypeName]; innerStruct != nil {
			nested := r.findPathsInStruct(pkgStructs, fieldTypeName, targetRef, fieldPath, visited)
			results = append(results, nested...)
		}
	}

	return results
}

func getJSONFieldName(field *ast.Field) string {
	if field.Tag != nil {
		tagVal := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		jsonTag := tagVal.Get("json")
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			return parts[0]
		}
	}
	if len(field.Names) > 0 {
		return strings.ToLower(field.Names[0].Name[:1]) + field.Names[0].Name[1:]
	}
	return ""
}

func extractTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return extractTypeName(t.X)
	case *ast.ArrayType:
		return extractTypeName(t.Elt)
	case *ast.SelectorExpr:
		return t.Sel.Name
	default:
		return ""
	}
}

func fixtureHasRefField(fix *FixtureInfo, paths [][]string) bool {
	for _, path := range paths {
		if hasNonEmptyPath(fix.CreateObj, path) {
			return true
		}
		if fix.UpdateObj != nil && hasNonEmptyPath(fix.UpdateObj, path) {
			return true
		}
	}
	return false
}

func hasNonEmptyPath(obj any, path []string) bool {
	if len(path) == 0 {
		if obj == nil {
			return false
		}
		switch v := obj.(type) {
		case string:
			return v != ""
		case map[string]any:
			return len(v) > 0
		case []any:
			return len(v) > 0
		default:
			return true
		}
	}

	if m, ok := obj.(map[string]any); ok {
		val, exists := m[path[0]]
		if !exists || val == nil {
			return false
		}
		return hasNonEmptyPath(val, path[1:])
	}

	if list, ok := obj.([]any); ok {
		for _, elem := range list {
			if hasNonEmptyPath(elem, path) {
				return true
			}
		}
		return false
	}

	return false
}
