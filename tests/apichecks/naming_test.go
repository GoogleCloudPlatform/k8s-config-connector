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

package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func findKindsInDir(dir string) ([]string, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)
	if err != nil {
		return nil, err
	}

	var kinds []string
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
					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok {
						continue
					}
					// Check if this struct embeds TypeMeta and ObjectMeta
					hasTypeMeta := false
					hasObjectMeta := false
					for _, field := range structType.Fields.List {
						// Embed fields don't have Names
						if len(field.Names) == 0 {
							if isTypeMeta(field.Type) {
								hasTypeMeta = true
							}
							if isObjectMeta(field.Type) {
								hasObjectMeta = true
							}
						}
					}
					if hasTypeMeta && hasObjectMeta {
						kinds = append(kinds, typeSpec.Name.Name)
					}
				}
			}
		}
	}
	return kinds, nil
}

func isTypeMeta(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name == "metav1" && t.Sel.Name == "TypeMeta"
		}
	case *ast.Ident:
		return t.Name == "TypeMeta"
	}
	return false
}

func isObjectMeta(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name == "metav1" && t.Sel.Name == "ObjectMeta"
		}
	case *ast.Ident:
		return t.Name == "ObjectMeta"
	}
	return false
}

func isMatchingPrefix(prefix, kind, service string) bool {
	// 1. Exact match
	if prefix == kind {
		return true
	}
	// 2. Suffix match (e.g. "bigquerytable" ends with "table")
	if strings.HasSuffix(kind, prefix) {
		return true
	}
	// 3. Strip service prefix from the kind name (e.g. kind "bigquerytable", service "bigquery" -> "table")
	if service != "" {
		if strings.HasPrefix(kind, service) {
			stripped := strings.TrimPrefix(kind, service)
			if prefix == stripped {
				return true
			}
		}
	}
	return false
}

func TestDirectResourceFileNaming(t *testing.T) {
	apisDir := "../../apis"
	controllerDir := "../../pkg/controller/direct"

	// packageKindsCache maps directory path -> list of lowercased kind names defined in that directory
	packageKindsCache := make(map[string][]string)
	allKinds := make(map[string]bool)

	// Pre-pass: Walk apis/ to find all packages and their defined Kinds
	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			kinds, err := findKindsInDir(path)
			if err != nil {
				return err
			}
			if len(kinds) > 0 {
				var lowerKinds []string
				for _, k := range kinds {
					lowerLower := strings.ToLower(k)
					lowerKinds = append(lowerKinds, lowerLower)
					allKinds[lowerLower] = true
				}
				packageKindsCache[path] = lowerKinds
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error scanning apis directory: %v", err)
	}

	suffixes := []string{
		"_types.go",
		"_types_test.go",
		"_identity.go",
		"_identity_test.go",
		"_reference.go",
		"_reference_test.go",
		"_mapper.go",
		"_mapper_test.go",
		"_fuzzer.go",
		"_fuzzer_test.go",
		"_controller.go",
		"_controller_test.go",
	}

	var errors []string

	// Walk apis directory to check naming matching kinds defined within package
	err = filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		filename := info.Name()
		var matchedSuffix string
		for _, s := range suffixes {
			if strings.HasSuffix(filename, s) {
				matchedSuffix = s
				break
			}
		}
		if matchedSuffix == "" {
			return nil
		}

		dir := filepath.Dir(path)
		kinds, hasKinds := packageKindsCache[dir]
		if !hasKinds {
			return nil
		}

		relDir, err := filepath.Rel(apisDir, dir)
		if err != nil {
			return err
		}
		parts := strings.Split(filepath.ToSlash(relDir), "/")
		service := ""
		if len(parts) > 0 {
			service = parts[0]
		}

		prefix := strings.TrimSuffix(filename, matchedSuffix)
		found := false
		for _, k := range kinds {
			if isMatchingPrefix(prefix, k, service) {
				found = true
				break
			}
		}
		if !found {
			rel, err := filepath.Rel(apisDir, path)
			if err != nil {
				return err
			}
			normalizedPath := filepath.ToSlash(filepath.Join("apis", rel))
			sort.Strings(kinds)
			errors = append(errors, fmt.Sprintf("[naming_violation] file=%s prefix=%s (expected one of %v)", normalizedPath, prefix, kinds))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking apis dir: %v", err)
	}

	// Walk pkg/controller/direct directory to check naming matching any kind in the repository
	err = filepath.Walk(controllerDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		filename := info.Name()
		var matchedSuffix string
		for _, s := range suffixes {
			if strings.HasSuffix(filename, s) {
				matchedSuffix = s
				break
			}
		}
		if matchedSuffix == "" {
			return nil
		}

		relPath, err := filepath.Rel(controllerDir, path)
		if err != nil {
			return err
		}
		parts := strings.Split(filepath.ToSlash(relPath), "/")
		service := ""
		if len(parts) > 0 {
			service = parts[0]
		}

		prefix := strings.TrimSuffix(filename, matchedSuffix)
		found := false
		for k := range allKinds {
			if isMatchingPrefix(prefix, k, service) {
				found = true
				break
			}
		}
		if !found {
			rel, err := filepath.Rel(controllerDir, path)
			if err != nil {
				return err
			}
			normalizedPath := filepath.ToSlash(filepath.Join("pkg/controller/direct", rel))
			errors = append(errors, fmt.Sprintf("[naming_violation] file=%s prefix=%s (expected direct resource kind prefix)", normalizedPath, prefix))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking controller dir: %v", err)
	}

	sort.Strings(errors)
	want := strings.Join(errors, "\n")
	if want != "" {
		want += "\n"
	}

	test.CompareGoldenFile(t, "testdata/exceptions/naming_violations.txt", want)
}
