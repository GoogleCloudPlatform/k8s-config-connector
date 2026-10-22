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

// TestAPIsNoInternalControllerImports ensures that packages under apis/ remain
// lightweight API/type definitions and do not import internal controller or
// runtime packages (such as pkg/k8s, pkg/controller, pkg/krmtotf, pkg/dcl, or
// operator) that tightly couple external consumers of apis/ to KCC's internal
// controller-runtime version.
func TestAPIsNoInternalControllerImports(t *testing.T) {
	t.Parallel()

	forbiddenPrefixes := []string{
		"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s",
		"github.com/GoogleCloudPlatform/k8s-config-connector/operator",
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
