// Copyright 2024 Google LLC
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

package io

import (
	"go/ast"
	"go/token"
)

type StructComment struct {
	Name    *ast.Ident
	Comment *ast.CommentGroup
	Fields  map[string]*ast.CommentGroup
}

func (s *StructComment) SetCommentPos(offset token.Pos) token.Pos {
	newComments := []*ast.Comment{}
	if s.Comment == nil {
		return offset
	}
	for _, c := range s.Comment.List {
		c.Slash = offset
		offset += c.End() + 1
		newComments = append(newComments, c)
	}
	s.Comment = &ast.CommentGroup{List: newComments}
	return offset
}

func MergeStructComments(oldComments, newComments map[string]*StructComment) map[string]*StructComment {
	merged := make(map[string]*StructComment)
	for k, v := range oldComments {
		merged[k] = v
	}

	// If the new Comment is non empty, override the old comment.
	for k, v := range newComments {
		merged[k] = v
	}
	return merged
}

// ast Docs and comments are not binding with the Decl but the relative position of the Node,
// so we have to handle it specifically.
func MapGoComments(f ast.Node) map[string]*StructComment {
	comments := make(map[string]*StructComment)

	ast.Inspect(f, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			// TODO: support IMPORT, CONST, VAR
			return true
		}
		c := mapStructComments(genDecl)
		if c == nil {
			return true
		}
		comments[c.Name.Name] = c
		return true
	})
	return comments
}

func mapStructComments(g *ast.GenDecl) *StructComment {
	if g == nil {
		return nil
	}
	c := &StructComment{
		Fields: make(map[string]*ast.CommentGroup),
	}

	if g.Doc != nil {
		c.Comment = ResetCommentsPos(g.Doc)
	}

	for _, spec := range g.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		c.Name = typeSpec.Name
		for _, field := range typeSpec.Type.(*ast.StructType).Fields.List {
			if field.Doc == nil {
				continue
			}
			var name *ast.Ident

			if field.Names == nil {
				// Anonymous name
				name = field.Type.(*ast.SelectorExpr).Sel
			} else {
				// Field could has multiple names like `A, B string`. Use the first one
				name = field.Names[0]
			}
			c.Fields[name.Name] = field.Doc
			c.Comment = ResetCommentsPos(field.Doc)
		}
	}
	if c.Name == nil {
		return nil
	}
	return c
}

func ResetCommentsPos(comments *ast.CommentGroup) *ast.CommentGroup {
	if comments == nil {
		return nil
	}
	newComments := []*ast.Comment{}
	for _, c := range comments.List {
		c.Slash = token.NoPos
		newComments = append(newComments, c)
	}
	return &ast.CommentGroup{
		List: newComments,
	}
}
