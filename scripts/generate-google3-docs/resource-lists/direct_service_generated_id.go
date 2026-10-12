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

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// directResourcesWithServiceGeneratedResourceIDCache caches the result of
// scanning the direct controllers, since it is used by multiple lists.
var directResourcesWithServiceGeneratedResourceIDCache []schema.GroupVersionKind

// allDirectResourcesWithServiceGeneratedResourceID returns the direct resources
// (of all versions) with a service-generated resource ID.
//
// It statically analyzes the direct controllers under pkg/controller/direct.
// A direct resource is considered to have a service-generated resource ID if
// its adapter's Find() method short-circuits when the resource ID is not yet
// known, e.g.:
//
//	func (a *sitemapAdapter) Find(ctx context.Context) (bool, error) {
//		if a.id.Sitemap == "" {
//			return false, nil
//		}
//		...
//	}
//
// The GVK of the resource is taken from the registry.RegisterModel() call in
// the same file.
func allDirectResourcesWithServiceGeneratedResourceID() ([]schema.GroupVersionKind, error) {
	if directResourcesWithServiceGeneratedResourceIDCache != nil {
		return directResourcesWithServiceGeneratedResourceIDCache, nil
	}

	crds, err := crdloader.LoadCRDs()
	if err != nil {
		return nil, fmt.Errorf("error loading CRDs: %w", err)
	}

	directDir := filepath.Join(repo.GetRootOrLogFatal(), "pkg", "controller", "direct")
	gvks := make([]schema.GroupVersionKind, 0)
	err = filepath.WalkDir(directDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return nil
		}
		fileGVKs, err := directGVKsWithServiceGeneratedResourceIDInFile(p, crds)
		if err != nil {
			return fmt.Errorf("error analyzing %v: %w", p, err)
		}
		gvks = append(gvks, fileGVKs...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking direct controllers in %v: %w", directDir, err)
	}
	directResourcesWithServiceGeneratedResourceIDCache = gvks
	return gvks, nil
}

// directGVKsWithServiceGeneratedResourceIDInFile returns the GVKs registered in
// the given Go file if the file contains a Find() method that indicates a
// service-generated resource ID.
func directGVKsWithServiceGeneratedResourceIDInFile(filePath string, crds []apiextensions.CustomResourceDefinition) ([]schema.GroupVersionKind, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}

	hasServiceGeneratedID := false
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "Find" || fn.Recv == nil || fn.Body == nil {
			continue
		}
		if findIndicatesServiceGeneratedID(fn) {
			hasServiceGeneratedID = true
			break
		}
	}
	if !hasServiceGeneratedID {
		return nil, nil
	}

	gvkRefs := registeredGVKRefs(f)
	if len(gvkRefs) == 0 {
		log.Printf("WARNING: %v has a Find() indicating a service-generated resource ID, but no registry.RegisterModel() call in the same file; skipping", filePath)
		return nil, nil
	}

	imports := importsByName(f)
	gvks := make([]schema.GroupVersionKind, 0, len(gvkRefs))
	for _, ref := range gvkRefs {
		importPath, ok := imports[ref.pkg]
		if !ok {
			return nil, fmt.Errorf("cannot resolve import for %v.%v", ref.pkg, ref.name)
		}
		gvk, err := resolveDirectGVK(importPath, ref.name, crds)
		if err != nil {
			return nil, err
		}
		gvks = append(gvks, gvk)
	}
	return gvks, nil
}

// findIndicatesServiceGeneratedID returns true if the given Find() method
// contains a statement of the form:
//
//	if <recv>.<...>.<field> == "" {
//		return false, nil
//	}
func findIndicatesServiceGeneratedID(fn *ast.FuncDecl) bool {
	if len(fn.Recv.List) == 0 || len(fn.Recv.List[0].Names) == 0 {
		return false
	}
	recvName := fn.Recv.List[0].Names[0].Name

	found := false
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		if found {
			return false
		}
		ifStmt, ok := n.(*ast.IfStmt)
		if !ok {
			return true
		}
		if isRecvFieldEmptyStringCheck(ifStmt.Cond, recvName) && returnsFalseNil(ifStmt.Body) {
			found = true
			return false
		}
		return true
	})
	return found
}

// isRecvFieldEmptyStringCheck returns true if expr is of the form
// `<recv>.<...>.<field> == ""` (or `"" == <recv>.<...>.<field>`), where the
// selector chain has at least two levels (e.g. `a.id.Sitemap`). The last
// element may also be a no-argument method call (e.g. `a.id.ID() == ""`).
func isRecvFieldEmptyStringCheck(expr ast.Expr, recvName string) bool {
	bin, ok := expr.(*ast.BinaryExpr)
	if !ok || bin.Op != token.EQL {
		return false
	}
	lhs, rhs := bin.X, bin.Y
	if isEmptyStringLit(lhs) {
		lhs, rhs = rhs, lhs
	}
	if !isEmptyStringLit(rhs) {
		return false
	}
	if call, ok := lhs.(*ast.CallExpr); ok && len(call.Args) == 0 {
		lhs = call.Fun
	}
	depth := 0
	for {
		sel, ok := lhs.(*ast.SelectorExpr)
		if !ok {
			break
		}
		depth++
		lhs = sel.X
	}
	ident, ok := lhs.(*ast.Ident)
	return ok && ident.Name == recvName && depth >= 2
}

func isEmptyStringLit(expr ast.Expr) bool {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return false
	}
	s, err := strconv.Unquote(lit.Value)
	return err == nil && s == ""
}

// returnsFalseNil returns true if the block contains `return false, nil`.
func returnsFalseNil(block *ast.BlockStmt) bool {
	for _, stmt := range block.List {
		ret, ok := stmt.(*ast.ReturnStmt)
		if !ok || len(ret.Results) != 2 {
			continue
		}
		first, ok1 := ret.Results[0].(*ast.Ident)
		second, ok2 := ret.Results[1].(*ast.Ident)
		if ok1 && ok2 && first.Name == "false" && second.Name == "nil" {
			return true
		}
	}
	return false
}

type gvkRef struct {
	pkg  string
	name string
}

// registeredGVKRefs returns the `<pkg>.<Name>GVK` arguments passed to
// registry.RegisterModel() in the given file.
func registeredGVKRefs(f *ast.File) []gvkRef {
	var refs []gvkRef
	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || len(call.Args) == 0 {
			return true
		}
		fun, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || fun.Sel.Name != "RegisterModel" {
			return true
		}
		arg, ok := call.Args[0].(*ast.SelectorExpr)
		if !ok {
			return true
		}
		pkg, ok := arg.X.(*ast.Ident)
		if !ok || !strings.HasSuffix(arg.Sel.Name, "GVK") {
			return true
		}
		refs = append(refs, gvkRef{pkg: pkg.Name, name: arg.Sel.Name})
		return true
	})
	return refs
}

// importsByName returns a map from the local package name to the import path.
func importsByName(f *ast.File) map[string]string {
	imports := make(map[string]string)
	for _, imp := range f.Imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}
		name := path.Base(importPath)
		if imp.Name != nil {
			name = imp.Name.Name
		}
		imports[name] = importPath
	}
	return imports
}

// resolveDirectGVK resolves the GVK for the GVK variable `gvkVarName` declared
// in the API package at `importPath` (e.g. ".../apis/compute/v1beta1").
func resolveDirectGVK(importPath, gvkVarName string, crds []apiextensions.CustomResourceDefinition) (schema.GroupVersionKind, error) {
	version := path.Base(importPath)
	service := path.Base(path.Dir(importPath))
	kind := strings.TrimSuffix(gvkVarName, "GVK")

	var matches []apiextensions.CustomResourceDefinition
	for _, crd := range crds {
		if crd.Spec.Names.Kind != kind {
			continue
		}
		for _, v := range crd.Spec.Versions {
			if v.Name == version {
				matches = append(matches, crd)
				break
			}
		}
	}
	switch len(matches) {
	case 0:
		return schema.GroupVersionKind{}, fmt.Errorf("no CRD found for kind %v with version %v (from %v.%v)", kind, version, importPath, gvkVarName)
	case 1:
		return schema.GroupVersionKind{Group: matches[0].Spec.Group, Version: version, Kind: kind}, nil
	default:
		// Disambiguate using the service name from the API package path.
		group := service + ".cnrm.cloud.google.com"
		for _, crd := range matches {
			if crd.Spec.Group == group {
				return schema.GroupVersionKind{Group: group, Version: version, Kind: kind}, nil
			}
		}
		return schema.GroupVersionKind{}, fmt.Errorf("multiple CRDs found for kind %v with version %v (from %v.%v)", kind, version, importPath, gvkVarName)
	}
}
