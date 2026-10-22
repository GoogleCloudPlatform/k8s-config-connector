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
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type fuzzerKRMPackage struct {
	importPath string
	structs    map[string]*gocode.GoStruct
	loader     *fuzzerKRMLoader
}

func (p *fuzzerKRMPackage) lookupStruct(krmField *gocode.StructField, typeName string) (*gocode.GoStruct, *fuzzerKRMPackage) {
	if p == nil {
		return nil, nil
	}
	if st := p.structs[typeName]; st != nil {
		return st, p
	}
	if krmField == nil || krmField.GoPackage == "" || p.loader == nil {
		return nil, p
	}
	otherPkg, err := p.loader.load(krmField.GoPackage)
	if err != nil {
		return nil, p
	}
	if _, shortName, ok := strings.Cut(typeName, "."); ok {
		if st := otherPkg.structs[shortName]; st != nil {
			return st, otherPkg
		}
	}
	return nil, p
}

type fuzzerKRMLoader struct {
	apiGoPackagePath string
	apiDirectory     string
	packages         map[string]*fuzzerKRMPackage
}

func newFuzzerKRMLoader(apiGoPackagePath, apiDirectory string) *fuzzerKRMLoader {
	return &fuzzerKRMLoader{
		apiGoPackagePath: strings.TrimSuffix(apiGoPackagePath, "/"),
		apiDirectory:     apiDirectory,
		packages:         make(map[string]*fuzzerKRMPackage),
	}
}

func (l *fuzzerKRMLoader) load(importPath string) (*fuzzerKRMPackage, error) {
	if pkg, ok := l.packages[importPath]; ok {
		return pkg, nil
	}
	rel, ok := strings.CutPrefix(importPath, l.apiGoPackagePath+"/")
	if !ok {
		return nil, fmt.Errorf("KRM package %q is not under %q", importPath, l.apiGoPackagePath)
	}
	dir := filepath.Join(l.apiDirectory, filepath.FromSlash(rel))
	pkg, err := gocode.LoadPackage(importPath, dir)
	if err != nil || pkg == nil {
		return nil, fmt.Errorf("loading KRM package %q from %s: %w", importPath, dir, err)
	}
	out := &fuzzerKRMPackage{importPath: pkg.GoPackage, structs: make(map[string]*gocode.GoStruct), loader: l}
	for _, s := range pkg.Structs {
		out.structs[s.Name] = s
	}
	l.packages[importPath] = out
	return out, nil
}

type goPackageIndex struct {
	dir         string
	packageName string
	files       []*goIndexedFile
	funcs       map[string]*goIndexedFunc
}

type goIndexedFile struct {
	path      string
	generated bool
	syntax    *ast.File
	imports   map[string]string
}

type goIndexedFunc struct {
	decl *ast.FuncDecl
	file *goIndexedFile
}

type fuzzerMapperPair struct {
	FromProto      string
	ToProto        string
	KRMType        string
	KRMGoPackage   string
	ProtoType      string
	ProtoGoPackage string
	Handwritten    bool
}

type fuzzerMapperCoverage struct {
	isStub           bool
	fromReads        map[string]bool
	toWrites         map[string]bool
	conditionalGuard map[string]bool
}

func (c *fuzzerMapperCoverage) roundTrips(fd protoreflect.FieldDescriptor) bool {
	if c == nil || c.isStub {
		return true
	}
	norm := normProtoIdent(string(fd.Name()))
	return c.fromReads[norm] && c.toWrites[norm]
}

func (c *fuzzerMapperCoverage) hintFor(fd protoreflect.FieldDescriptor) string {
	if c != nil && !c.isStub && c.conditionalGuard[normProtoIdent(string(fd.Name()))] {
		return "mapper guards this field behind a value-dependent 'if' condition; map unconditionally or configure f.FilterSpec / f.FilterStatus"
	}
	return ""
}

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
		file, err := parseIndexedGoFile(fset, filepath.Join(dir, name))
		if err != nil {
			return nil, err
		}
		if idx.packageName == "" {
			idx.packageName = file.syntax.Name.Name
		}
		if !isFuzzerSourceFile(name) {
			idx.indexTopLevelFuncs(file)
		}
		idx.files = append(idx.files, file)
	}
	return idx, nil
}

func parseIndexedGoFile(fset *token.FileSet, filePath string) (*goIndexedFile, error) {
	syntax, err := parser.ParseFile(fset, filePath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %w", filePath, err)
	}
	imports, err := extractFileImports(syntax, filePath)
	if err != nil {
		return nil, err
	}
	return &goIndexedFile{
		path:      filePath,
		generated: strings.HasSuffix(filePath, "generated.go"),
		syntax:    syntax,
		imports:   imports,
	}, nil
}

func extractFileImports(syntax *ast.File, filePath string) (map[string]string, error) {
	imports := make(map[string]string, len(syntax.Imports))
	for _, imp := range syntax.Imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return nil, fmt.Errorf("parsing import %s in %q: %w", imp.Path.Value, filePath, err)
		}
		importName := path.Base(importPath)
		if imp.Name != nil {
			importName = imp.Name.Name
		}
		imports[importName] = importPath
	}
	return imports, nil
}

func isFuzzerSourceFile(name string) bool {
	return strings.HasSuffix(name, "_fuzzer.go") || strings.HasSuffix(name, "_fuzzer.generated.go")
}

func (idx *goPackageIndex) indexTopLevelFuncs(file *goIndexedFile) {
	for _, decl := range file.syntax.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Recv != nil {
			continue
		}
		existing, found := idx.funcs[fd.Name.Name]
		if !found || (existing.file.generated && !file.generated) {
			idx.funcs[fd.Name.Name] = &goIndexedFunc{decl: fd, file: file}
		}
	}
}

func (idx *goPackageIndex) handwrittenFunc(name string) string {
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

func (idx *goPackageIndex) findMapperPair(protoType, version string, typeNames ...string) (pair *fuzzerMapperPair, skipped []string) {
	for _, raw := range typeNames {
		t := strings.ToUpper(raw[:1]) + raw[1:]
		for _, base := range []string{t + "_" + version, t} {
			if base == "_" {
				continue
			}
			candidate, reason := idx.matchMapperCandidate(base, protoType, version)
			if reason != "" {
				skipped = append(skipped, reason)
				continue
			}
			if candidate != nil && (pair == nil || (candidate.Handwritten && !pair.Handwritten)) {
				pair = candidate
			}
		}
	}
	return pair, skipped
}

func (idx *goPackageIndex) matchMapperCandidate(base, protoType, version string) (candidate *fuzzerMapperPair, skipReason string) {
	candidate, err := idx.mapperPair(base)
	if err != nil {
		return nil, err.Error()
	}
	if candidate == nil {
		return nil, ""
	}
	if candidate.ProtoType != protoType {
		return nil, fmt.Sprintf("%s converts *%s, not *%s", candidate.FromProto, candidate.ProtoType, protoType)
	}
	if version != "" && candidate.KRMGoPackage != "" && path.Base(candidate.KRMGoPackage) != version {
		return nil, fmt.Sprintf("%s returns a type of %s, not of version %s", candidate.FromProto, candidate.KRMGoPackage, version)
	}
	return candidate, ""
}

func (idx *goPackageIndex) findStructMapperPair(protoType, version, krmType string) *fuzzerMapperPair {
	if pair, _ := idx.findMapperPair(protoType, version, krmType); pair != nil {
		return pair
	}
	for _, name := range idx.sortedFuncNames() {
		base, ok := strings.CutSuffix(name, "_FromProto")
		if !ok {
			continue
		}
		if c, err := idx.mapperPair(base); err == nil && c != nil && c.ProtoType == protoType && c.KRMType == krmType {
			return c
		}
	}
	return nil
}

func (idx *goPackageIndex) sortedFuncNames() []string {
	names := make([]string, 0, len(idx.funcs))
	for name := range idx.funcs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

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
	ft, toFT := from.decl.Type, to.decl.Type
	if fieldListArity(ft.Params) != 2 || fieldListArity(ft.Results) != 1 {
		return nil, fmt.Errorf("%s in %s: want func(*direct.MapContext, *<proto>) *<KRM type>", pair.FromProto, from.file.path)
	}
	var ok bool
	if pair.ProtoGoPackage, pair.ProtoType, ok = from.file.typeRef(ft.Params.List[len(ft.Params.List)-1].Type); !ok {
		return nil, fmt.Errorf("%s in %s: cannot tell the Go type of its proto parameter", pair.FromProto, from.file.path)
	}
	if pair.KRMGoPackage, pair.KRMType, ok = from.file.typeRef(ft.Results.List[0].Type); !ok {
		return nil, fmt.Errorf("%s in %s: cannot tell the KRM type it returns", pair.FromProto, from.file.path)
	}
	if fieldListArity(toFT.Params) != 2 || fieldListArity(toFT.Results) != 1 {
		return nil, fmt.Errorf("%s in %s: want func(*direct.MapContext, *%s) *%s", pair.ToProto, to.file.path, pair.KRMType, pair.ProtoType)
	}
	if _, krmType, ok := to.file.typeRef(toFT.Params.List[len(toFT.Params.List)-1].Type); !ok || krmType != pair.KRMType {
		return nil, fmt.Errorf("%s in %s: want *%s as its second parameter", pair.ToProto, to.file.path, pair.KRMType)
	}
	if _, protoType, ok := to.file.typeRef(toFT.Results.List[0].Type); !ok || protoType != pair.ProtoType {
		return nil, fmt.Errorf("%s in %s: want *%s as its result", pair.ToProto, to.file.path, pair.ProtoType)
	}
	return pair, nil
}

func (idx *goPackageIndex) inspectMapperPair(pair *fuzzerMapperPair) *fuzzerMapperCoverage {
	if idx == nil || pair == nil {
		return nil
	}
	from, to := idx.funcs[pair.FromProto], idx.funcs[pair.ToProto]
	if from == nil || to == nil || from.decl.Body == nil || to.decl.Body == nil {
		return nil
	}
	cov := &fuzzerMapperCoverage{
		fromReads:        make(map[string]bool),
		toWrites:         make(map[string]bool),
		conditionalGuard: make(map[string]bool),
	}
	fromAlloc := cov.inspectFromProtoBody(from)
	toAlloc := cov.inspectToProtoBody(to, idx)
	cov.isStub = !fromAlloc && !toAlloc && len(cov.fromReads) == 0 && len(cov.toWrites) == 0
	return cov
}

func (c *fuzzerMapperCoverage) inspectFromProtoBody(from *goIndexedFunc) (hasAlloc bool) {
	ast.Inspect(from.decl.Body, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.AssignStmt:
			if assignsIdent(s, "out") {
				hasAlloc = true
			}
		case *ast.SelectorExpr:
			if isProtoInputVar(s.X, from.file) {
				c.fromReads[normProtoIdent(trimProtoGetter(s.Sel.Name))] = true
			}
		case *ast.IfStmt:
			if isValueComparisonGuard(s.Cond) {
				c.recordGuardedProtoReads(s.Body, from.file)
				if s.Else != nil {
					c.recordGuardedProtoReads(s.Else, from.file)
				}
			}
		}
		return true
	})
	return hasAlloc
}

func (c *fuzzerMapperCoverage) recordGuardedProtoReads(node ast.Node, file *goIndexedFile) {
	ast.Inspect(node, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok && isProtoInputVar(sel.X, file) {
			c.conditionalGuard[normProtoIdent(trimProtoGetter(sel.Sel.Name))] = true
		}
		return true
	})
}

func (c *fuzzerMapperCoverage) inspectToProtoBody(to *goIndexedFunc, idx *goPackageIndex) (hasAlloc bool) {
	ast.Inspect(to.decl.Body, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.AssignStmt:
			if assignsIdent(s, "out") {
				hasAlloc = true
			}
			for i, lhs := range s.Lhs {
				if sel, ok := lhs.(*ast.SelectorExpr); ok && (i >= len(s.Rhs) || !idx.isNilOrStubCallExpr(s.Rhs[i])) {
					c.toWrites[normProtoIdent(sel.Sel.Name)] = true
				}
			}
		case *ast.CompositeLit:
			for _, elt := range s.Elts {
				if kv, ok := elt.(*ast.KeyValueExpr); ok && !idx.isNilOrStubCallExpr(kv.Value) {
					if key, ok := kv.Key.(*ast.Ident); ok {
						c.toWrites[normProtoIdent(key.Name)] = true
					}
				}
			}
		}
		return true
	})
	return hasAlloc
}

func (idx *goPackageIndex) isNilOrStubCallExpr(expr ast.Expr) bool {
	if isNilIdent(expr) {
		return true
	}
	call, ok := expr.(*ast.CallExpr)
	if !ok || idx == nil {
		return false
	}
	id, ok := call.Fun.(*ast.Ident)
	if !ok {
		return false
	}
	fn := idx.funcs[id.Name]
	if fn == nil || fn.decl.Body == nil {
		return false
	}
	hasNonNilReturn := false
	ast.Inspect(fn.decl.Body, func(n ast.Node) bool {
		if ret, ok := n.(*ast.ReturnStmt); ok {
			for _, r := range ret.Results {
				if !isNilIdent(r) {
					hasNonNilReturn = true
				}
			}
		}
		return true
	})
	return !hasNonNilReturn
}

func assignsIdent(s *ast.AssignStmt, name string) bool {
	for _, lhs := range s.Lhs {
		if id, ok := lhs.(*ast.Ident); ok && id.Name == name {
			return true
		}
	}
	return false
}

func isProtoInputVar(expr ast.Expr, file *goIndexedFile) bool {
	id, ok := expr.(*ast.Ident)
	if !ok {
		return false
	}
	if id.Name == "out" || id.Name == "mapCtx" {
		return false
	}
	return file.imports[id.Name] == ""
}

func isValueComparisonGuard(cond ast.Expr) bool {
	var found bool
	ast.Inspect(cond, func(n ast.Node) bool {
		bin, ok := n.(*ast.BinaryExpr)
		if !ok {
			return true
		}
		switch bin.Op {
		case token.EQL:
			if !isNilIdent(bin.X) && !isNilIdent(bin.Y) {
				found = true
			}
		case token.NEQ:
			if !isNilOrEmptyString(bin.X) && !isNilOrEmptyString(bin.Y) {
				found = true
			}
		}
		return true
	})
	return found
}

func isNilOrEmptyString(expr ast.Expr) bool {
	if isNilIdent(expr) {
		return true
	}
	lit, ok := expr.(*ast.BasicLit)
	return ok && lit.Kind == token.STRING && lit.Value == `""`
}

func isNilIdent(expr ast.Expr) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == "nil"
}

func normProtoIdent(s string) string { return strings.ToLower(strings.ReplaceAll(s, "_", "")) }

func trimProtoGetter(name string) string {
	if strings.HasPrefix(name, "Get") && len(name) > 3 && unicode.IsUpper(rune(name[3])) {
		return name[3:]
	}
	return name
}

func fieldListArity(fl *ast.FieldList) int {
	if fl == nil {
		return 0
	}
	n := 0
	for _, f := range fl.List {
		n += max(1, len(f.Names))
	}
	return n
}

func (f *goIndexedFile) typeRef(expr ast.Expr) (goPackage, name string, ok bool) {
	star, isStar := expr.(*ast.StarExpr)
	if !isStar {
		return "", "", false
	}
	switch t := star.X.(type) {
	case *ast.Ident:
		return "", t.Name, true
	case *ast.SelectorExpr:
		if x, isIdent := t.X.(*ast.Ident); isIdent {
			if importPath, found := f.imports[x.Name]; found {
				return importPath, t.Sel.Name, true
			}
		}
	}
	return "", "", false
}
