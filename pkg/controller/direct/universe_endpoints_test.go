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

package direct

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// publicUniverseSuffix is the API host suffix of the public Google Cloud
// universe. An endpoint containing it is expressed against that universe, and
// must be translated before use.
const publicUniverseSuffix = "googleapis.com"

// TestEndpointOverridesAreUniverseAware fails when a direct controller passes
// option.WithEndpoint an endpoint built from a googleapis.com literal without
// routing it through ControllerConfig.Endpoint.
//
// Most controllers need no endpoint handling at all: RESTClientOptions and
// GRPCClientOptions pass the universe domain to the client library, which
// derives each service's endpoint itself. The controllers that *do* override
// their endpoint bypass that, so a hardcoded googleapis.com host wins over the
// configured universe and silently sends traffic to the public universe.
//
// This is a lint rather than a behavioural test because the failure is
// invisible at runtime in the public universe — the only universe CI can
// exercise — and because direct controllers are still being added in bulk
// (#10588). Without it, the count of broken overrides grows faster than anyone
// notices.
func TestEndpointOverridesAreUniverseAware(t *testing.T) {
	root := "."
	fset := token.NewFileSet()

	var findings []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			// A file we cannot parse is not this test's problem; the build
			// catches it.
			return nil //nolint:nilerr
		}

		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				return true
			}
			for _, finding := range checkFunc(fset, fn) {
				findings = append(findings, finding)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %v", root, err)
	}

	if len(findings) > 0 {
		t.Errorf("%d endpoint override(s) are not universe-aware.\n\n"+
			"Wrap the endpoint in ControllerConfig.Endpoint, which rewrites the host for the\n"+
			"configured universe and returns it unchanged in the public universe:\n\n"+
			"    option.WithEndpoint(m.config.Endpoint(\"service.googleapis.com:443\"))\n\n%s",
			len(findings), strings.Join(findings, "\n"))
	}
}

// checkFunc reports endpoint overrides in fn that are not universe-aware.
func checkFunc(fset *token.FileSet, fn *ast.FuncDecl) []string {
	// Assignments in this function, so that an endpoint built into a local
	// variable and passed by name can still be resolved.
	assigned := map[string]ast.Expr{}
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}
		for i, lhs := range assign.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok || i >= len(assign.Rhs) {
				continue
			}
			assigned[ident.Name] = assign.Rhs[i]
		}
		return true
	})

	var findings []string
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || !isSelector(call.Fun, "option", "WithEndpoint") || len(call.Args) != 1 {
			return true
		}

		arg := call.Args[0]
		// Resolve a bare identifier to the expression it was assigned.
		if ident, ok := arg.(*ast.Ident); ok {
			if rhs, found := assigned[ident.Name]; found {
				arg = rhs
			}
		}

		if !mentionsPublicUniverse(arg) || callsEndpointHelper(arg) {
			return true
		}

		pos := fset.Position(call.Pos())
		findings = append(findings, "  "+pos.String())
		return true
	})
	return findings
}

// isSelector reports whether expr is the selector pkg.name.
func isSelector(expr ast.Expr, pkg, name string) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != name {
		return false
	}
	ident, ok := sel.X.(*ast.Ident)
	return ok && ident.Name == pkg
}

// mentionsPublicUniverse reports whether expr contains a string literal naming
// the public universe host suffix.
func mentionsPublicUniverse(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		lit, ok := n.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}
		value, err := strconv.Unquote(lit.Value)
		if err != nil {
			return true
		}
		if strings.Contains(value, publicUniverseSuffix) {
			found = true
			return false
		}
		return true
	})
	return found
}

// callsEndpointHelper reports whether expr contains a call to an Endpoint
// method, i.e. ControllerConfig.Endpoint.
func callsEndpointHelper(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Endpoint" {
			found = true
			return false
		}
		return true
	})
	return found
}
