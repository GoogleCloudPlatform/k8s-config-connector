// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocode

import (
	"go/ast"
	"go/token"
)

// NewDocMap creates a map associating AST nodes (structs and fields) with their
// documentation comments that are immediately above them.
func NewDocMap(fset *token.FileSet, file *ast.File) map[ast.Node]*ast.CommentGroup {
	docMap := make(map[ast.Node]*ast.CommentGroup)

	// Traverse the AST and map comments to structs and fields
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.GenDecl:
			if node.Tok != token.TYPE || len(node.Specs) == 0 {
				return true
			}

			for _, spec := range node.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					return true
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					return true
				}

				// associate the comment group immediately above the struct
				if node.Doc != nil {
					docMap[ts] = node.Doc
				}

				// handle fields within the struct
				for _, field := range st.Fields.List {
					if field.Doc != nil {
						docMap[field] = field.Doc
					}
				}
			}
		}
		return true
	})

	return docMap
}
