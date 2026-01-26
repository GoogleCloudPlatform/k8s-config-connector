// Copyright 2025 Google LLC
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

package jsonunmarshalreuse

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "jsonunmarshalreuse",
	Doc:  "checks for suboptimal JSON unmarshalling practices where a non-empty variable might be reused, leading to merging instead of overwriting",
	Run:  run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		// Check function call identifier
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		isTarget := false
		
		// Target 1: encoding/json.Unmarshal
		if sel.Sel.Name == "Unmarshal" {
			if obj := pass.TypesInfo.ObjectOf(sel.Sel); obj != nil {
				if pkg := obj.Pkg(); pkg != nil && pkg.Path() == "encoding/json" {
					isTarget = true
				}
			}
		}
		
		// Target 2: pkg/util.Marshal
		if !isTarget && sel.Sel.Name == "Marshal" {
			if obj := pass.TypesInfo.ObjectOf(sel.Sel); obj != nil {
				if pkg := obj.Pkg(); pkg != nil && pkg.Path() == "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util" {
					isTarget = true
				}
			}
		}

		if !isTarget {
			return
		}

		// Ensure there are at least two arguments (data, v)
		if len(call.Args) < 2 {
			return
		}

		// The second argument is the destination 'v'
		arg := call.Args[1]

		// Get the type of the argument after dereferencing if it's a pointer.
		var expr ast.Expr = arg
		
		// Handle &v (UnaryExpr) - this is the most common case: json.Unmarshal(data, &v)
		if unaryExpr, isUnary := expr.(*ast.UnaryExpr); isUnary && unaryExpr.Op == token.AND {
			expr = unaryExpr.X
		}
		
		// Handle *v (StarExpr)
		if starExpr, isStar := expr.(*ast.StarExpr); isStar {
			expr = starExpr.X
		}

		// Function to check if an expression is a problematic initialization
		isProblematic := func(e ast.Expr) (bool, string) {
			// Check for non-empty composite literal (struct, slice, map)
			if lit, isLit := e.(*ast.CompositeLit); isLit {
				if len(lit.Elts) > 0 {
					return true, "potential reuse of non-empty variable in json.Unmarshal/util.Marshal; consider using an empty literal or nil"
				}
			}

			// Check for make() call with length > 0
			if makeCall, isMakeCall := e.(*ast.CallExpr); isMakeCall {
				if fun, isFun := makeCall.Fun.(*ast.Ident); isFun && fun.Name == "make" {
					if len(makeCall.Args) >= 2 {
						// The second argument of make is the length
						if basicLit, isBasicLit := makeCall.Args[1].(*ast.BasicLit); isBasicLit && basicLit.Kind == token.INT {
							if length, err := strconv.Atoi(basicLit.Value); err == nil && length > 0 {
								return true, "potential reuse of variable created with non-zero length in json.Unmarshal/util.Marshal; consider using make([]T, 0) or nil"
							}
						}
					}
				}
			}
			return false, ""
		}

		// 1. Check if the argument itself is a problematic expression (e.g. inline literal)
		if problem, msg := isProblematic(expr); problem {
			pass.Reportf(call.Pos(), "%s", msg)
			return
		}

		// 2. If it's an identifier, check its declaration
		if ident, isIdent := expr.(*ast.Ident); isIdent {
			if obj := ident.Obj; obj != nil && obj.Kind == ast.Var {
				if decl, ok := obj.Decl.(*ast.ValueSpec); ok {
					// Check initialization values
					for _, value := range decl.Values {
						if problem, msg := isProblematic(value); problem {
							pass.Reportf(call.Pos(), "%s", msg)
							return
						}
					}
				}
				// Handle short variable declaration (assign statement)
				if assign, ok := obj.Decl.(*ast.AssignStmt); ok {
					for _, expr := range assign.Rhs {
						if problem, msg := isProblematic(expr); problem {
							pass.Reportf(call.Pos(), "%s", msg)
							return
						}
					}
				}
			}
		}
	})
	return nil, nil
}