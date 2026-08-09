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

package common_test

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

// TestAllBrownfieldControllersPassRequiredArguments verifies that calls to common.NormalizeReferences
// in direct controllers of brownfield resources pass non-nil arguments for projectRef
// (e.g. projectRef/identity) and projectMapper (e.g. m.config.ProjectMapper).
func TestAllBrownfieldControllersPassRequiredArguments(t *testing.T) {
	repoRoot, err := repo.GetRoot()
	if err != nil {
		t.Fatalf("failed finding repo root: %v", err)
	}

	// 1. Build a set of brownfield kinds (resources that support both Direct and TF/DCL)
	brownfieldKinds := make(map[string]bool)
	for gk, config := range resourceconfig.ControllerConfigStatic {
		hasDirect := false
		hasLegacy := false
		for _, c := range config.SupportedControllers {
			if c == k8s.ReconcilerTypeDirect {
				hasDirect = true
			}
			if c == k8s.ReconcilerTypeTerraform || c == k8s.ReconcilerTypeDCL {
				hasLegacy = true
			}
		}
		if hasDirect && hasLegacy {
			brownfieldKinds[strings.ToLower(gk.Kind)] = true
		}
	}

	directDir := filepath.Join(repoRoot, "pkg", "controller", "direct")
	fset := token.NewFileSet()

	var errs []string
	err = filepath.Walk(directDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, "_controller.go") {
			return err
		}

		baseName := strings.TrimSuffix(filepath.Base(path), "_controller.go")
		if !brownfieldKinds[strings.ToLower(baseName)] {
			return nil
		}

		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		ast.Inspect(node, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if isNormalizeReferencesCall(call) {
				// Expected signature: NormalizeReferences(ctx, reader, obj, projectRef, projectMapper)
				if len(call.Args) != 5 {
					pos := fset.Position(call.Pos())
					errs = append(errs, fmt.Sprintf("%s: NormalizeReferences expected 5 arguments, got %d", pos, len(call.Args)))
					return true
				}

				if ident, ok := call.Args[3].(*ast.Ident); ok && ident.Name == "nil" {
					pos := fset.Position(call.Pos())
					errs = append(errs, fmt.Sprintf("%s: NormalizeReferences called with nil projectRef; pass projectRef instead", pos))
				}

				if ident, ok := call.Args[4].(*ast.Ident); ok && ident.Name == "nil" {
					pos := fset.Position(call.Pos())
					errs = append(errs, fmt.Sprintf("%s: NormalizeReferences called with nil projectMapper; pass m.config.ProjectMapper instead", pos))
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		t.Fatalf("failed scanning direct controllers: %v", err)
	}

	if len(errs) > 0 {
		sort.Strings(errs)
		t.Fatalf("Found %d invalid NormalizeReferences call(s):\n%s", len(errs), strings.Join(errs, "\n"))
	}
}

func isNormalizeReferencesCall(call *ast.CallExpr) bool {
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		return fn.Sel.Name == "NormalizeReferences"
	case *ast.Ident:
		return fn.Name == "NormalizeReferences"
	}
	return false
}

