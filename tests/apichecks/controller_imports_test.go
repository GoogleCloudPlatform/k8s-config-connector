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
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestAPIsNoControllerImports ensures that packages under apis/ do not import
// internal controller implementation packages (such as pkg/controller,
// pkg/krmtotf, or pkg/dcl).
func TestAPIsNoControllerImports(t *testing.T) {
	t.Parallel()

	forbiddenPrefixes := []string{
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller",
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf",
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl",
	}

	apisDir := "../../apis"
	fset := token.NewFileSet()

	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			t.Errorf("error parsing %s: %v", path, err)
			return nil
		}

		relPath, relErr := filepath.Rel("../..", path)
		if relErr != nil {
			relPath = path
		}

		for _, imp := range file.Imports {
			importPath, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				continue
			}
			for _, forbidden := range forbiddenPrefixes {
				if importPath == forbidden || strings.HasPrefix(importPath, forbidden+"/") {
					pos := fset.Position(imp.Pos())
					t.Errorf("%s:%d: forbidden import %q in apis/ package (apis/ must not depend on %s)", relPath, pos.Line, importPath, forbidden)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking apis dir: %v", err)
	}
}

// TestPtrPackageZeroExternalImports ensures that the top-level ptr/ package
// only imports Go standard library packages and never becomes a grab-bag of
// external or internal KCC dependencies.
func TestPtrPackageZeroExternalImports(t *testing.T) {
	t.Parallel()

	ptrDir := "../../ptr"
	fset := token.NewFileSet()

	err := filepath.Walk(ptrDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			t.Errorf("error parsing %s: %v", path, err)
			return nil
		}

		relPath, relErr := filepath.Rel("../..", path)
		if relErr != nil {
			relPath = path
		}

		for _, imp := range file.Imports {
			importPath, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				continue
			}
			// Any non-stdlib package has a dot in its first path element (e.g. github.com, k8s.io).
			firstSeg, _, _ := strings.Cut(importPath, "/")
			if strings.Contains(firstSeg, ".") {
				pos := fset.Position(imp.Pos())
				t.Errorf("%s:%d: forbidden non-stdlib import %q in ptr/ package (ptr/ must remain zero-dependency)", relPath, pos.Line, importPath)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking ptr dir: %v", err)
	}
}
