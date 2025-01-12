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
	"bytes"
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"regexp"
	"strings"

	"k8s.io/klog/v2"
)

// UpdateGoFile update the `fileName` gofile with the content in `output`.
func UpdateGoFile(ctx context.Context, file string, output *bytes.Buffer) error {
	klog := klog.FromContext(ctx)

	goCode := extractGoCode(output.Bytes())
	if goCode == "" {
		return fmt.Errorf("no go code in %s", output)
	}
	klog.Info("update go file", "path", file)
	// TODO: replace other go type like function, var, import..
	if err := rebuild(file, goCode); err != nil {
		return err
	}
	return nil
}

// The data is normally a combination of go code (with triple quote code blocks marker ```go ... ```)
// and wording explanation for the recommendation.
func extractGoCode(data []byte) string {
	re := regexp.MustCompile("(?s)```go\n(.*?)```")
	matches := re.FindStringSubmatch(string(data))
	if len(matches) == 2 {
		return matches[1]
	}
	if len(matches) > 2 {
		klog.Info("more than one gocode snippets found")
	}
	return ""
}

// Rebuild the entire go file.
// Known problem: Comments in ast.Node are relative position, and can be misplaced by go/printer when refactoring go file.
// https://github.com/golang/go/issues/20744/. Potentially lib (we'd better not use) https://github.com/dave/dst
// To fix the problem, we can calculate the go file offset and re-assign token.Pos to each comment, using Printer.CommentedNode to
// specifically write each node with its new comments back to dest file.
func rebuild(destFilePath string, src string) error {
	destFset := token.NewFileSet()
	destFile, err := parser.ParseFile(destFset, destFilePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing destination file: %w", err)
	}

	// Make sure src package exists otherwise the parser will fail
	_, _, found := strings.Cut(src, "package")
	if !found {
		src = "package " + destFile.Name.Name + "\n" + src
	}
	newFset := token.NewFileSet()
	newFile, err := parser.ParseFile(newFset, "", src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing src file: %w", err)
	}

	// TODO: Parse the entire package (rather than a single `file`) to avoid redeclaration errors for structs
	// defined in different files within the same package.
	fwriter, err := os.Create(destFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer fwriter.Close()

	newDecls := []ast.Decl{}

	oldComments := MapGoComments(destFile)
	newComments := MapGoComments(newFile)
	mergedComments := MergeStructComments(oldComments, newComments)

	newStructs := mapStruct(newFile)

	oldFileComments := destFile.Comments
	if oldFileComments == nil {
		oldFileComments = []*ast.CommentGroup{}
	}
	destFile.Comments = []*ast.CommentGroup{}
	offset := token.Pos(0)
	if oldFileComments != nil {
		if strings.Contains(oldFileComments[0].Text(), "Copyright") {
			destFile.Comments = append(destFile.Comments, oldFileComments[0])
			offset += oldFileComments[0].End() + 1
		}
	}
	offset += destFile.Package + 1

	var errs error
	visited := map[string]struct{}{}

	// TODO: add comments for VAR, CONST, FUNC and IMPORT.
	ast.Inspect(destFile, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			switch x.Tok {
			case token.VAR, token.CONST:
				/*
					newVarDecl := &ast.GenDecl{
						Tok:   x.Tok,
						Specs: []ast.Spec{},
					}
					for _, spec := range x.Specs {
						valueSpec, ok := spec.(*ast.ValueSpec)
						if !ok {
							return true
						}
						newSpec := &ast.ValueSpec{
							Names:  valueSpec.Names,
							Values: valueSpec.Values,
							Type:   valueSpec.Type,
						}
						newVarDecl.Specs = append(newVarDecl.Specs, newSpec)
					}*/
				delta := x.Specs[len(x.Specs)-1].End() - x.Specs[0].Pos()
				offset += delta + 1
				newDecls = append(newDecls, x)
			case token.IMPORT:
				/*
					newImportDecl := &ast.GenDecl{
						Tok:   token.IMPORT,
						Specs: []ast.Spec{},
					}
					for _, spec := range x.Specs {
						importSpec, ok := spec.(*ast.ImportSpec)
						if !ok {
							return true
						}
						newSpec := &ast.ImportSpec{
							Name: importSpec.Name,
							Path: importSpec.Path,
						}
						newImportDecl.Specs = append(newImportDecl.Specs, newSpec)
					}*/
				delta := x.Specs[len(x.Specs)-1].End() - x.Specs[0].Pos()
				offset += delta + 1
				newDecls = append(newDecls, x)
			case token.TYPE:
				for _, spec := range x.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						return true
					}
					var newStructDecl *ast.GenDecl
					newStruct, ok := newStructs[typeSpec.Name.Name]
					comments := mergedComments[typeSpec.Name.Name]

					if !ok {
						newStructDecl, err = copyStruct(x, comments)
						if err != nil {
							errors.Join(errs, err)
							return false
						}
					} else {
						newStructDecl, err = mergeStruct(x, newStruct, comments)
						if err != nil {
							errors.Join(errs, err)
							return false
						}
					}
					offset = comments.SetCommentPos(offset)
					if comments.Comment != nil {
						newStructDecl.Doc = comments.Comment
					}
					typeStruct := newStructDecl.Specs[0].(*ast.TypeSpec)
					for _, field := range typeStruct.Type.(*ast.StructType).Fields.List {
						// Add field comment on top of the field.
						newCommentList := []*ast.Comment{}
						if field.Doc != nil {
							for _, c := range field.Doc.List {
								c.Slash = offset
								offset += c.End() + 1
								newCommentList = append(newCommentList, c)
							}
							field.Doc = &ast.CommentGroup{List: newCommentList}
						}
						delta := field.End() - field.Pos()
						offset += delta + 1
					}

					newDecls = append(newDecls, newStructDecl)
					visited[typeSpec.Name.Name] = struct{}{}
				}

			}
		case *ast.FuncDecl:
			newFunDecl := &ast.FuncDecl{
				Name: x.Name,
				Type: x.Type,
				Body: x.Body,
			}
			newFields := []*ast.Field{}
			if x.Recv != nil {
				for _, r := range x.Recv.List {
					newField := &ast.Field{
						Names: r.Names,
						Type:  r.Type,
						Tag:   r.Tag,
					}
					newFields = append(newFields, newField)
				}
			}
			newFunDecl.Recv = &ast.FieldList{
				List: newFields,
			}
			// TODO: Add comments to funcDecl
			newDecls = append(newDecls, newFunDecl)
			delta := x.Body.End() - x.Body.Pos()
			offset += delta + 1
		case *ast.BadDecl:
			newBadDecl := &ast.BadDecl{
				From: x.From,
				To:   x.To,
			}
			newDecls = append(newDecls, newBadDecl)
		}
		return true
	})
	if errs != nil {
		return err
	}

	// TODO Add non struct type from newFile
	ast.Inspect(newFile, func(n ast.Node) bool {
		x, ok := n.(*ast.GenDecl)
		if !ok || x.Tok != token.TYPE {
			return true
		}
		for _, spec := range x.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok := visited[typeSpec.Name.Name]; ok {
				continue
			}
			newDecls = append(newDecls, x)
		}
		return true
	})

	destFile.Decls = newDecls

	/*
		if err := RebuildCommentPos(fwriter, destFset, destFile, mergedComments, offset); err != nil {
			return err
		}*/
	err = format.Node(fwriter, destFset, destFile)
	if err != nil {
		return fmt.Errorf("error formatting code: %w", err)
	}
	return nil
}

func RebuildCommentPos(fwriter io.Writer, fset *token.FileSet, f *ast.File, comments map[string]*StructComment, offset token.Pos) error {
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			switch x.Tok {
			case token.VAR, token.CONST, token.IMPORT:
				delta := x.Specs[len(x.Specs)-1].End() - x.Specs[0].Pos()
				offset += delta + 1
			case token.TYPE:
				newfieldComments := []*ast.CommentGroup{}

				for _, spec := range x.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						return true
					}
					comments := comments[typeSpec.Name.Name]
					offset = comments.SetCommentPos(offset)

					if comments.Comment != nil {
						// f.Comments = append(f.Comments, comments.Comment)
						x.Doc = comments.Comment
						newfieldComments = append(newfieldComments, x.Doc)
					}
					for _, field := range typeSpec.Type.(*ast.StructType).Fields.List {
						// Add field comment on top of the field.
						newCommentList := []*ast.Comment{}
						if field.Doc != nil {
							for _, c := range field.Doc.List {
								c.Slash = offset
								offset += c.End() + 1
								newCommentList = append(newCommentList, c)
							}
							newfieldComments = append(newfieldComments, &ast.CommentGroup{List: newCommentList})
						}

						delta := field.End() - field.Pos()
						offset += delta + 1
					}
				}
			}
		case *ast.FuncDecl:
			delta := x.Body.End() - x.Body.Pos()
			offset += delta + 1
			return true
		}
		return true
	})
	return nil
}

func mapStruct(f *ast.File) map[string]*ast.GenDecl {
	visited := map[string]*ast.GenDecl{}
	ast.Inspect(f, func(n ast.Node) bool {
		x, ok := n.(*ast.GenDecl)
		if !ok || x.Tok != token.TYPE {
			return true
		}

		for _, spec := range x.Specs {
			// StructType has is a single Spec
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				// TODO Support ImportSpec and ValueSpec
				continue
			}
			visited[typeSpec.Name.Name] = x
		}
		return true
	})
	return visited
}

func mergeStruct(oldGenDecl, newGenDecl *ast.GenDecl, comments *StructComment) (*ast.GenDecl, error) {
	if newGenDecl == nil {
		return oldGenDecl, nil
	}
	if oldGenDecl.Tok != token.TYPE {
		return nil, fmt.Errorf("oldGenDecl is not TYPE struct ")
	}
	if newGenDecl.Tok != token.TYPE {
		return nil, fmt.Errorf("newGenDecl is not TYPE struct ")
	}

	oldTypeSpec, ok := oldGenDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("currentGendDecl should contain TypeSpec")
	}
	newTypeSpec, ok := newGenDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("newGenDecl should contain TypeSpec")
	}

	mergedGenDecl := &ast.GenDecl{
		Tok: token.TYPE,
	}

	// Map the new fields
	replacements := make(map[*ast.Ident]*ast.Field)
	for _, newField := range newTypeSpec.Type.(*ast.StructType).Fields.List {
		if newField.Names == nil {
			// The `metav1.TypeMeta` and `metav1.ListMeta` fields normally follows the no-name convention.
			replacements[newField.Type.(*ast.SelectorExpr).Sel] = newField
			continue
		}
		if len(newField.Names) > 1 {
			names := []string{}
			for _, n := range newField.Names {
				names = append(names, n.Name)
			}
			fmt.Printf("struct %s has multi-name field %s", newTypeSpec.Name.Name, names)
		}
		replacements[newField.Names[0]] = newField
	}

	newFields := []*ast.Field{}

	// Merge fields
	visited := make(map[string]struct{})
	for _, oldfield := range oldTypeSpec.Type.(*ast.StructType).Fields.List {
		var key *ast.Ident
		if oldfield.Names == nil {
			// anonymous field
			key = oldfield.Type.(*ast.SelectorExpr).Sel
		} else {
			key = oldfield.Names[0]
		}

		var field *ast.Field
		if newfield, ok := replacements[key]; ok {
			field = MergeField(oldfield, newfield)
		} else {
			field = oldfield
		}
		// Add comment separately because it is a relative position to the struct.
		if comments != nil {
			field.Doc = comments.Fields[key.Name]
		}
		visited[key.Name] = struct{}{}
		newFields = append(newFields, field)
	}

	// Add new fields
	for key, newfield := range replacements {
		if _, ok := visited[key.Name]; !ok {
			field := newfield
			if comments != nil {
				field.Doc = comments.Fields[key.Name]
			}
			newFields = append(newFields, field)
		}
	}
	mergedGenDecl.Specs = []ast.Spec{
		&ast.TypeSpec{
			Name: oldTypeSpec.Name,
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: newFields,
				},
			},
		},
	}
	return mergedGenDecl, nil
}

func copyStruct(origin *ast.GenDecl, comments *StructComment) (*ast.GenDecl, error) {
	if origin == nil {
		return nil, nil
	}
	copy := &ast.GenDecl{
		Tok: token.TYPE,
	}

	newSpecs := []ast.Spec{}
	for _, spec := range origin.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			// TODO Support other types
			continue
		}
		newFields := &ast.FieldList{
			List: []*ast.Field{},
		}
		structType := typeSpec.Type.(*ast.StructType)
		for _, field := range structType.Fields.List {
			var key *ast.Ident
			if field.Names == nil {
				// anonymous field
				key = field.Type.(*ast.SelectorExpr).Sel
			} else {
				key = field.Names[0]
			}
			/*
				field.Type = &ast.SelectorExpr{
					X: field.Type.(*ast.SelectorExpr).X,
					Sel: &ast.Ident{
						Name: key.Name,
						// Assign offset to Namepos is essential to calculate the entire StructType Comment Pos.
						NamePos: offset,
						Obj:     field.Type.(*ast.SelectorExpr).Sel.Obj,
					},
				}
				offset = comments.SetFieldPos(key.Name, offset)
			*/
			if comments != nil {
				field.Doc = comments.Fields[key.Name]
			}
			newFields.List = append(newFields.List, field)

		}
		newStructType := &ast.StructType{
			Fields: newFields,
			Struct: structType.Struct,
		}
		newTypeSpec := &ast.TypeSpec{
			Name: typeSpec.Name,
			Type: newStructType,
		}
		// delta := newTypeSpec.End() - newTypeSpec.Pos()
		// offset += delta + 1
		newSpecs = append(newSpecs, newTypeSpec)
	}

	copy.Specs = newSpecs
	return copy, nil
}

func MergeField(oldfield, newfield *ast.Field) *ast.Field {
	f := &ast.Field{}
	if newfield.Type != nil {
		f.Type = newfield.Type
	} else {
		f.Type = oldfield.Type
	}
	if newfield.Names != nil {
		f.Names = newfield.Names
	} else {
		f.Names = oldfield.Names
	}
	if newfield.Tag != nil {
		f.Tag = newfield.Tag
	} else {
		f.Tag = oldfield.Tag
	}
	return f
}
