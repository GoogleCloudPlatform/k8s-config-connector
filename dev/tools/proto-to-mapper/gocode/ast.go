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

package gocode

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

type Package struct {
	GoPackage string
	SourceDir string

	Structs []*GoStruct

	Comments []string
}

func (p *Package) GetAnnotation(key string) string {
	for _, c := range p.Comments {
		for _, line := range strings.Split(c, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, key+"=") {
				v := strings.TrimPrefix(line, key+"=")

				return v
			}
		}
	}
	return ""
}

type GoStruct struct {
	GoPackage string
	Name      string
	Fields    []*StructField

	Comments []string
}

type StructField struct {
	Name string
	Type string
}

func LoadPackage(goPackage string, path string) (*Package, error) {
	fileSet := token.NewFileSet()
	mode := parser.ParseComments
	var filter func(fs.FileInfo) bool
	packages, err := parser.ParseDir(fileSet, path, filter, mode)
	if err != nil {
		return nil, fmt.Errorf("parsing directory %q: %w", path, err)
	}
	if len(packages) == 0 {
		return nil, nil
	}
	if len(packages) != 1 {
		return nil, fmt.Errorf("parsing directory %q: found %d packages; want 1", path, len(packages))
	}
	out := &Package{GoPackage: goPackage}
	for packageName, p := range packages {
		if err := out.inspect(packageName, p); err != nil {
			return nil, fmt.Errorf("inspecting package %q: %w", packageName, err)
		}
	}
	out.SourceDir = path
	return out, nil
}

func LoadPackageTree(goPackage string, basePath string) ([]*Package, error) {
	var packages []*Package

	if err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				return err
			}
			pkg, err := LoadPackage(filepath.Join(goPackage, relPath), path)
			if err != nil {
				return fmt.Errorf("loading package %q: %w", path, err)
			}
			if pkg != nil {
				packages = append(packages, pkg)
			}
			return nil
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return packages, nil
}

func (p *Package) inspect(packageName string, pkg *ast.Package) error {
	var errs []error
	var comments []ast.Node

	for _, pkgFile := range pkg.Files {
		var pkgComments []ast.Node
		for _, comment := range pkgFile.Comments {
			pkgComments = append(pkgComments, comment)
		}
		comments, err := parseComments(pkgComments)
		if err != nil {
			return err
		}
		p.Comments = append(p.Comments, comments...)
	}

	ast.Inspect(pkg, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch n := n.(type) {
		case *ast.TypeSpec:
			switch def := n.Type.(type) {
			case *ast.StructType:
				if err := p.addStruct(n.Name, def, comments); err != nil {
					errs = append(errs, err)
				}
			default:
				errs = append(errs, fmt.Errorf("unhandled type spec in %q: %T, %+v", n.Name, n.Type, n.Type))
			}
			comments = nil

		case *ast.Comment:
			comments = append(comments, n)
		case *ast.CommentGroup:
			// A CommentGroup contains a list of comments
			// Do not truncate comments when we encounter the group
		default:
			if n != nil {
				comments = nil
			}
		}

		return true
	})

	return errors.Join(errs...)
}

func (p *Package) addStruct(name *ast.Ident, def *ast.StructType, comments []ast.Node) error {
	goStruct := &GoStruct{
		GoPackage: p.GoPackage,
		Name:      name.String(),
	}

	for _, field := range def.Fields.List {
		structField := &StructField{}
		for _, name := range field.Names {
			structField.Name += name.String()
		}
		goType, err := toGoType(field.Type)
		if err != nil {
			return err
		}
		structField.Type = goType

		goStruct.Fields = append(goStruct.Fields, structField)
	}

	{
		comments, err := parseComments(comments)
		if err != nil {
			return err
		}
		goStruct.Comments = append(goStruct.Comments, comments...)
	}

	p.Structs = append(p.Structs, goStruct)
	return nil
}

func parseComments(comments []ast.Node) ([]string, error) {
	var out []string
	for _, comment := range comments {
		switch comment := comment.(type) {
		case *ast.Comment:
			text := comment.Text
			text = strings.TrimSpace(text)
			text = strings.TrimPrefix(text, "//")
			text = strings.TrimPrefix(text, "/*")
			text = strings.TrimSuffix(text, "*/")
			out = append(out, text)
		case *ast.CommentGroup:
			out = append(out, comment.Text())
		default:
			return nil, fmt.Errorf("unexpected comment node type %T", comment)
		}
	}
	return out, nil
}

func toGoType(t ast.Expr) (string, error) {
	switch t := t.(type) {
	case *ast.Ident:
		return t.String(), nil
	case *ast.StarExpr:
		s, err := toGoType(t.X)
		if err != nil {
			return "", err
		}
		return "*" + s, nil
	case *ast.ArrayType:
		s, err := toGoType(t.Elt)
		if err != nil {
			return "", err
		}
		return "[]" + s, nil
	case *ast.SelectorExpr:
		s, err := toGoType(t.X)
		if err != nil {
			return "", err
		}
		return s + "." + t.Sel.String(), nil

	case *ast.MapType:
		k, err := toGoType(t.Key)
		if err != nil {
			return "", err
		}
		v, err := toGoType(t.Value)
		if err != nil {
			return "", err
		}
		return "map[" + k + "]" + v, nil

	default:
		return "", fmt.Errorf("unhandled field type %T, %+v", t, t)
	}
}
