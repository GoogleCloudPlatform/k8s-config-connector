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

package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "lintwithname",
	Doc:  "checks for usage of klog.FromContext(ctx).WithName",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if sel.Sel.Name != "WithName" {
				return true
			}

			// Check receiver is klog.FromContext(...)
			recCall, ok := sel.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			funSel, ok := recCall.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			pkgIdent, ok := funSel.X.(*ast.Ident)
			if !ok {
				return true
			}

			if pkgIdent.Name == "klog" && funSel.Sel.Name == "FromContext" {
				pass.Reportf(call.Pos(), "do not use klog.FromContext(ctx).WithName(), use klog.FromContext(ctx) instead")
			}

			return true
		})
	}
	return nil, nil
}
