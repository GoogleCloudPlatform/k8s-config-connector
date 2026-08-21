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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Result struct {
	Added     []string
	Removed   []string
	Changed   []string
	Unchanged []string
}

func (r *Result) Accuracy() float64 {
	total := len(r.Added) + len(r.Removed) + len(r.Changed) + len(r.Unchanged)
	return float64(len(r.Unchanged)) / float64(total)
}

func (r *Result) String() string {
	return "total: " + fmt.Sprint(len(r.Added)+len(r.Removed)+len(r.Changed)+len(r.Unchanged)) +
		", added: " + fmt.Sprint(r.Added) +
		", removed: " + fmt.Sprint(r.Removed) +
		", changed: " + fmt.Sprint(r.Changed) +
		", unchanged: " + fmt.Sprint(r.Unchanged)
}

type structPair struct {
	A, B *ast.StructType
}

type Results []Result

func (r Results) Accuracy() float64 {
	total := 0
	unchanged := 0
	for _, result := range r {
		total += len(result.Added) + len(result.Removed) + len(result.Changed) + len(result.Unchanged)
		unchanged += len(result.Unchanged)
	}
	return float64(unchanged) / float64(total)
}

// If codefactors is given, only compare the go structs whose name match the codefactor list.
func CompareStruct(code, otherCode string, factors ...string) (*Results, error) {
	aFset := token.NewFileSet()
	aFile, err := parser.ParseFile(aFset, "", code, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing first go code snippet: %w", err)
	}

	bFset := token.NewFileSet()
	// Make sure src package exists otherwise the parser will fail
	_, _, found := strings.Cut(otherCode, "package")
	if !found {
		otherCode = "package main" + "\n" + otherCode
	}
	bFile, err := parser.ParseFile(bFset, "", otherCode, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing second go code snippet: %w", err)
	}

	compareAll := false
	if factors == nil {
		compareAll = true
	}

	pairMap := make(map[string]*structPair)
	for _, c := range factors {
		pairMap[c] = &structPair{}
	}

	ast.Inspect(aFile, func(n ast.Node) bool {
		x, ok := n.(*ast.GenDecl)
		if !ok || x.Tok != token.TYPE {
			return true
		}

		for _, spec := range x.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// record A for comparison
			if pair, ok := pairMap[typeSpec.Name.Name]; compareAll || ok {
				structType, typeOk := typeSpec.Type.(*ast.StructType)
				if !typeOk {
					continue
				}
				if pair == nil {
					pair = &structPair{}
				}
				pair.A = structType
				pairMap[typeSpec.Name.Name] = pair
			}
		}
		return true
	})

	ast.Inspect(bFile, func(n ast.Node) bool {
		x, ok := n.(*ast.GenDecl)
		if !ok || x.Tok != token.TYPE {
			return true
		}

		for _, spec := range x.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// record B for comparison
			if pair, ok := pairMap[typeSpec.Name.Name]; compareAll || ok {
				structType, typeOk := typeSpec.Type.(*ast.StructType)
				if !typeOk {
					continue
				}
				if pair == nil {
					pair = &structPair{}
				}
				pair.B = structType
				pairMap[typeSpec.Name.Name] = pair
			}
		}
		return true
	})

	results := &Results{}
	for name, pair := range pairMap {
		r := compareFields(pair)
		fmt.Println("Compare: ", name, "Result: ", r)
		*results = append(*results, *r)
	}
	return results, nil
}

func compareFields(pair *structPair) *Result {
	r := &Result{}

	fieldsMap := make(map[string]*ast.Field)
	if pair.A != nil {
		for _, f := range pair.A.Fields.List {
			// todo: handle multi names
			fieldsMap[fieldName(f)] = f
		}
	}

	visited := make(map[string]struct{})
	if pair.B != nil {
		for _, bField := range pair.B.Fields.List {
			fname := fieldName(bField)
			aField, ok := fieldsMap[fname]
			if !ok {
				r.Added = append(r.Added, fname)
			} else {
				d := diff(aField, bField)
				if d != "" {
					r.Changed = append(r.Changed, fname)
				} else {
					r.Unchanged = append(r.Unchanged, fname)
				}
			}
			visited[fname] = struct{}{}
		}
	}

	if pair.A != nil {
		for _, f := range pair.A.Fields.List {
			if _, ok := visited[fieldName(f)]; !ok {
				r.Removed = append(r.Removed, fieldName(f))
			}
		}
	}
	return r
}

func diff(a, b *ast.Field) string {
	// compare "Names"
	if len(a.Names) != len(b.Names) {
		return fmt.Sprintf("different number of Names: %s, %s", a.Names, b.Names)
	}

	for _, n := range a.Names {
		match := false
		for _, m := range b.Names {
			if n.Name == m.Name {
				match = true
				break
			}
		}
		if !match {
			return fmt.Sprintf("mismatch name: %s, nil", n)
		}
	}

	// compare "Type"
	switch a.Type.(type) {
	case *ast.Ident:
		aType := a.Type.(*ast.Ident).Name
		bType := b.Type.(*ast.Ident).Name
		if aType != bType {
			return fmt.Sprintf("different Ident type: %s, %s", aType, bType)
		}
	case *ast.SelectorExpr:
		aType := a.Type.(*ast.SelectorExpr).Sel.Name
		bType := b.Type.(*ast.SelectorExpr).Sel.Name
		if aType != bType {
			return fmt.Sprintf("different SelectorExpr type: %s, %s", aType, bType)
		}
	}

	// compare Tag
	if a.Tag.Kind != b.Tag.Kind {
		return fmt.Sprintf("different tag kind: %s, %s", a.Tag.Kind, b.Tag.Kind)
	}
	if a.Tag.Value != b.Tag.Value {
		return fmt.Sprintf("different tag value: %s, %s", a.Tag.Value, b.Tag.Value)
	}
	return ""
}

func fieldName(f *ast.Field) string {
	if len(f.Names) == 0 {
		switch f.Type.(type) {
		case *ast.Ident:
			return f.Type.(*ast.Ident).Name
		case *ast.SelectorExpr:
			return f.Type.(*ast.SelectorExpr).Sel.Name
		}
	}
	return f.Names[0].Name
}
