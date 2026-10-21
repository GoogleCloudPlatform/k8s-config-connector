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
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// goPackageIndex holds the parsed Go files of a direct controller package.
type goPackageIndex struct {
	dir         string
	packageName string
	files       []*goIndexedFile
	funcs       map[string]*goIndexedFunc
}

type goIndexedFile struct {
	path string
	// generated is set for files whose name ends in generated.go, like findFuncDeclaration.
	generated bool
	syntax    *ast.File
	// imports maps each import's name (its alias, or the last element of its path) to its path.
	imports map[string]string
}

type goIndexedFunc struct {
	decl *ast.FuncDecl
	file *goIndexedFile
}

// fuzzerMapperPair is a <Type>_FromProto / <Type>_ToProto pair found in the direct package.
type fuzzerMapperPair struct {
	FromProto string
	ToProto   string
	// KRMType and KRMGoPackage are the struct that FromProto returns.
	KRMType      string
	KRMGoPackage string
	// ProtoType and ProtoGoPackage are the Go type of the proto message that FromProto takes.
	ProtoType      string
	ProtoGoPackage string
	// Handwritten is set when FromProto or ToProto is declared in a file that isn't generated.
	Handwritten bool
}

// loadGoPackageIndex parses the non-test Go files of dir. A missing directory gives an empty index.
func loadGoPackageIndex(dir string) (*goPackageIndex, error) {
	idx := &goPackageIndex{dir: dir, funcs: make(map[string]*goIndexedFunc)}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return idx, nil
		}
		return nil, fmt.Errorf("reading directory %q: %w", dir, err)
	}
	fset := token.NewFileSet()
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		if err := idx.indexGoFile(fset, filepath.Join(dir, name), name); err != nil {
			return nil, err
		}
	}
	return idx, nil
}

func (idx *goPackageIndex) indexGoFile(fset *token.FileSet, filePath, fileName string) error {
	syntax, err := parser.ParseFile(fset, filePath, nil, parser.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("parsing %q: %w", filePath, err)
	}
	file := &goIndexedFile{
		path:      filePath,
		generated: strings.HasSuffix(fileName, "generated.go"),
		syntax:    syntax,
		imports:   make(map[string]string),
	}
	for _, imp := range syntax.Imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return fmt.Errorf("parsing import %s in %q: %w", imp.Path.Value, filePath, err)
		}
		importName := path.Base(importPath)
		if imp.Name != nil {
			importName = imp.Name.Name
		}
		file.imports[importName] = importPath
	}
	if idx.packageName == "" {
		idx.packageName = syntax.Name.Name
	}
	for _, decl := range syntax.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok && fd.Recv == nil {
			if _, found := idx.funcs[fd.Name.Name]; !found {
				idx.funcs[fd.Name.Name] = &goIndexedFunc{decl: fd, file: file}
			}
		}
	}
	idx.files = append(idx.files, file)
	return nil
}

// handwrittenFunc returns the non-generated file that declares function name, or "".
func (idx *goPackageIndex) handwrittenFunc(name string) string {
	if f := idx.funcs[name]; f != nil && !f.file.generated {
		return f.file.path
	}
	for _, file := range idx.files {
		if file.generated {
			continue
		}
		for _, decl := range file.syntax.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok && fd.Recv == nil && fd.Name.Name == name {
				return file.path
			}
		}
	}
	return ""
}

// handwrittenFuzzerUsing returns a non-generated file that imports fuzztesting and refers to
// the function name, such as a fuzzer written by hand or by an LLM, or "".
func (idx *goPackageIndex) handwrittenFuzzerUsing(name string) string {
	for _, file := range idx.files {
		if file.generated || !file.importsPackage(fuzztestingGoPackage) {
			continue
		}
		if file.referencesIdent(name) {
			return file.path
		}
	}
	return ""
}

func (f *goIndexedFile) importsPackage(pkgPath string) bool {
	for _, importPath := range f.imports {
		if importPath == pkgPath {
			return true
		}
	}
	return false
}

func (f *goIndexedFile) referencesIdent(name string) bool {
	found := false
	ast.Inspect(f.syntax, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == name {
			found = true
		}
		return !found
	})
	return found
}

// handwrittenFuzzerBelow is handwrittenFuzzerUsing for the packages in the subdirectories of
// dir. For example the fuzzer of automl/automldataset uses the mappers of the automl package.
func handwrittenFuzzerBelow(dir, name string) (string, error) {
	found := ""
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() || p == dir {
			return nil
		}
		idx, err := loadGoPackageIndex(p)
		if err != nil {
			return err
		}
		if found = idx.handwrittenFuzzerUsing(name); found != "" {
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return "", fmt.Errorf("looking for fuzzers below %q: %w", dir, err)
	}
	return found, nil
}

// findMapperPair returns the mappers that a fuzzer uses for one side, choosing among the
// pairs named by mapperPairBases. Generated and hand-written functions both count, but pairs
// that convert a proto type other than protoType, that return a type of another KRM version,
// or whose signature isn't that of a mapper, are skipped and described in skipped. A
// hand-written pair wins over a generated one, because it is the one that the controller
// uses; otherwise the order of mapperPairBases decides. It returns nil when no pair is left.
func (idx *goPackageIndex) findMapperPair(protoType, version string, typeNames ...string) (pair *fuzzerMapperPair, skipped []string) {
	for _, base := range mapperPairBases(version, typeNames...) {
		candidate, err := idx.mapperPair(base)
		if err != nil {
			skipped = append(skipped, err.Error())
			continue
		}
		if candidate == nil {
			continue
		}
		if candidate.ProtoType != protoType {
			skipped = append(skipped, fmt.Sprintf("%s converts *%s, not *%s", candidate.FromProto, candidate.ProtoType, protoType))
			continue
		}
		if version != "" && candidate.KRMGoPackage != "" && path.Base(candidate.KRMGoPackage) != version {
			skipped = append(skipped, fmt.Sprintf("%s returns a type of %s, not of version %s", candidate.FromProto, candidate.KRMGoPackage, version))
			continue
		}
		if pair == nil || (candidate.Handwritten && !pair.Handwritten) {
			pair = candidate
		}
	}
	return pair, skipped
}

// mapperPairBases returns the names, without the _FromProto and _ToProto suffixes, that the
// mappers of the given KRM types can have, in order of preference: <typeName>_<version>, as
// generate-mapper --multiversion names them, then <typeName>.
func mapperPairBases(version string, typeNames ...string) []string {
	var bases []string
	for _, typeName := range typeNames {
		if version != "" {
			bases = append(bases, typeName+"_"+version)
		}
		bases = append(bases, typeName)
	}
	return bases
}

// mapperPair returns <base>_FromProto and <base>_ToProto, or nil when the package doesn't
// declare both.
func (idx *goPackageIndex) mapperPair(base string) (*fuzzerMapperPair, error) {
	from, to := idx.funcs[base+"_FromProto"], idx.funcs[base+"_ToProto"]
	if from == nil || to == nil {
		return nil, nil
	}
	pair := &fuzzerMapperPair{
		FromProto:   base + "_FromProto",
		ToProto:     base + "_ToProto",
		Handwritten: !from.file.generated || !to.file.generated,
	}
	if err := parseFromProtoSignature(from, pair); err != nil {
		return nil, err
	}
	if err := validateToProtoSignature(to, pair); err != nil {
		return nil, err
	}
	return pair, nil
}

func parseFromProtoSignature(from *goIndexedFunc, pair *fuzzerMapperPair) error {
	ft := from.decl.Type
	if fieldListArity(ft.Params) != 2 || fieldListArity(ft.Results) != 1 {
		return fmt.Errorf("%s in %s: want func(*direct.MapContext, *<proto>) *<KRM type>", pair.FromProto, from.file.path)
	}
	var ok bool
	params := ft.Params.List
	if pair.ProtoGoPackage, pair.ProtoType, ok = from.file.typeRef(params[len(params)-1].Type); !ok {
		return fmt.Errorf("%s in %s: cannot tell the Go type of its proto parameter", pair.FromProto, from.file.path)
	}
	if pair.KRMGoPackage, pair.KRMType, ok = from.file.typeRef(ft.Results.List[0].Type); !ok {
		return fmt.Errorf("%s in %s: cannot tell the KRM type it returns", pair.FromProto, from.file.path)
	}
	return nil
}

func validateToProtoSignature(to *goIndexedFunc, pair *fuzzerMapperPair) error {
	toFT := to.decl.Type
	if fieldListArity(toFT.Params) != 2 || fieldListArity(toFT.Results) != 1 {
		return fmt.Errorf("%s in %s: want func(*direct.MapContext, *%s) *%s", pair.ToProto, to.file.path, pair.KRMType, pair.ProtoType)
	}
	toParams := toFT.Params.List
	if _, krmType, ok := to.file.typeRef(toParams[len(toParams)-1].Type); !ok || krmType != pair.KRMType {
		return fmt.Errorf("%s in %s: want *%s as its second parameter", pair.ToProto, to.file.path, pair.KRMType)
	}
	if _, protoType, ok := to.file.typeRef(toFT.Results.List[0].Type); !ok || protoType != pair.ProtoType {
		return fmt.Errorf("%s in %s: want *%s as its result", pair.ToProto, to.file.path, pair.ProtoType)
	}
	return nil
}

func fieldListArity(fl *ast.FieldList) int {
	if fl == nil {
		return 0
	}
	n := 0
	for _, f := range fl.List {
		if len(f.Names) == 0 {
			n++
		} else {
			n += len(f.Names)
		}
	}
	return n
}

// typeRef resolves a pointer type expression such as *krm.FooSpec to its import path and type name.
// The import path is empty for a type of the same package.
func (f *goIndexedFile) typeRef(expr ast.Expr) (goPackage, name string, ok bool) {
	star, isStar := expr.(*ast.StarExpr)
	if !isStar {
		return "", "", false
	}
	switch t := star.X.(type) {
	case *ast.Ident:
		return "", t.Name, true
	case *ast.SelectorExpr:
		x, isIdent := t.X.(*ast.Ident)
		if !isIdent {
			return "", "", false
		}
		importPath, found := f.imports[x.Name]
		if !found {
			return "", "", false
		}
		return importPath, t.Sel.Name, true
	}
	return "", "", false
}
