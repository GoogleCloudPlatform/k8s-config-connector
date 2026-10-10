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

package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
)

// TypeDecl is a type that a Go file declares.
type TypeDecl struct {
	// File is the base name of the file.
	File string
	Spec *ast.TypeSpec
	// Doc is the type's doc comment, or nil.
	Doc *ast.CommentGroup
}

// Name returns the name of the type.
func (d TypeDecl) Name() string {
	return d.Spec.Name.Name
}

// ParseTypeDecls returns the types declared in the Go files of dir that
// include accepts. An error reading dir is returned as is, so callers can
// check os.IsNotExist.
func ParseTypeDecls(dir string, include func(fs.DirEntry) bool) ([]TypeDecl, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	var out []TypeDecl
	for _, e := range entries {
		if !include(e) {
			continue
		}
		decls, err := parseTypeDeclsInFile(fset, dir, e.Name())
		if err != nil {
			return nil, err
		}
		out = append(out, decls...)
	}
	return out, nil
}

// parseTypeDeclsInFile returns the types that a Go file declares.
func parseTypeDeclsInFile(fset *token.FileSet, dir, file string) ([]TypeDecl, error) {
	p := filepath.Join(dir, file)
	f, err := parser.ParseFile(fset, p, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", p, err)
	}
	var out []TypeDecl
	for _, decl := range f.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok == token.TYPE {
			out = append(out, typesIn(gd, file)...)
		}
	}
	return out, nil
}

// typesIn returns the types that a type declaration in file declares.
func typesIn(gd *ast.GenDecl, file string) []TypeDecl {
	var out []TypeDecl
	for _, spec := range gd.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		doc := ts.Doc
		if doc == nil && len(gd.Specs) == 1 {
			// A declaration of one type holds that type's doc comment.
			doc = gd.Doc
		}
		out = append(out, TypeDecl{File: file, Spec: ts, Doc: doc})
	}
	return out
}
