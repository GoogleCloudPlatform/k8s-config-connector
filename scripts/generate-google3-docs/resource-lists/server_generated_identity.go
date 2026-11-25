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
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// serverGeneratedIdentityResourcesCache caches the result of scanning the API
// packages, since it is used by multiple lists.
var serverGeneratedIdentityResourcesCache []schema.GroupVersionKind

// allResourcesWithServerGeneratedIdentity returns the resources (of all
// versions) whose identity type implements identity.ServerGeneratedIdentity.
//
// It statically analyzes the API packages under apis/<service>/<version>,
// looking for compile-time interface assertions of the form:
//
//	var (
//		_ identity.ServerGeneratedIdentity = &SitemapIdentity{}
//	)
//
// The kind is derived from the identity type name by trimming the "Identity"
// suffix, and the group is resolved from the CRDs.
func allResourcesWithServerGeneratedIdentity() ([]schema.GroupVersionKind, error) {
	if serverGeneratedIdentityResourcesCache != nil {
		return serverGeneratedIdentityResourcesCache, nil
	}

	crds, err := crdloader.LoadCRDs()
	if err != nil {
		return nil, fmt.Errorf("error loading CRDs: %w", err)
	}

	apisDir := filepath.Join(repo.GetRootOrLogFatal(), "apis")
	gvks := make([]schema.GroupVersionKind, 0)
	err = filepath.WalkDir(apisDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return nil
		}
		fileGVKs, err := serverGeneratedIdentityGVKsInFile(p, crds)
		if err != nil {
			return fmt.Errorf("error analyzing %v: %w", p, err)
		}
		gvks = append(gvks, fileGVKs...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking API packages in %v: %w", apisDir, err)
	}
	serverGeneratedIdentityResourcesCache = gvks
	return gvks, nil
}

// serverGeneratedIdentityGVKsInFile returns the GVKs of the identity types
// asserted to implement identity.ServerGeneratedIdentity in the given Go file.
func serverGeneratedIdentityGVKsInFile(filePath string, crds []apiextensions.CustomResourceDefinition) ([]schema.GroupVersionKind, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}

	identityTypes := serverGeneratedIdentityTypes(f)
	if len(identityTypes) == 0 {
		return nil, nil
	}

	// API packages are laid out as apis/<service>/<version>.
	pkgDir := filepath.ToSlash(filepath.Dir(filePath))
	gvks := make([]schema.GroupVersionKind, 0, len(identityTypes))
	for _, typeName := range identityTypes {
		kind := strings.TrimSuffix(typeName, "Identity")
		gvk, err := resolveDirectGVK(pkgDir, kind+"GVK", crds)
		if err != nil {
			log.Printf("WARNING: %v: %v implements identity.ServerGeneratedIdentity, but its GVK could not be resolved; skipping: %v", filePath, typeName, err)
			continue
		}
		gvks = append(gvks, gvk)
	}
	return gvks, nil
}

// serverGeneratedIdentityTypes returns the names of the types used in
// assertions of the form `_ identity.ServerGeneratedIdentity = &<Type>{}`.
func serverGeneratedIdentityTypes(f *ast.File) []string {
	var names []string
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.VAR {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || !isServerGeneratedIdentityType(valueSpec.Type) {
				continue
			}
			for _, value := range valueSpec.Values {
				if name, ok := compositeLitTypeName(value); ok && strings.HasSuffix(name, "Identity") {
					names = append(names, name)
				}
			}
		}
	}
	return names
}

// isServerGeneratedIdentityType returns true if expr is
// `<pkg>.ServerGeneratedIdentity`.
func isServerGeneratedIdentityType(expr ast.Expr) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "ServerGeneratedIdentity" {
		return false
	}
	_, ok = sel.X.(*ast.Ident)
	return ok
}

// compositeLitTypeName returns the type name of `&<Type>{}` or `<Type>{}`.
func compositeLitTypeName(expr ast.Expr) (string, bool) {
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		expr = unary.X
	}
	lit, ok := expr.(*ast.CompositeLit)
	if !ok {
		return "", false
	}
	ident, ok := lit.Type.(*ast.Ident)
	if !ok {
		return "", false
	}
	return ident.Name, true
}
