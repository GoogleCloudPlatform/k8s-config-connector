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
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	// fuzzerAgreementTarget is the share of fields that the design asks the deterministic
	// classification to classify the way the existing fuzzers do.
	fuzzerAgreementTarget = 0.90

	// fuzzerAgreementLeafDepth is the deepest proto field, in path segments, that counts as a
	// leaf in the leaf-weighted agreement.
	fuzzerAgreementLeafDepth = 6

	kccGoModule = "github.com/GoogleCloudPlatform/k8s-config-connector"
)

// TestFuzzerAgreement is the offline agreement check of phase 1 of the deterministic
// generation design. For every fuzzer that generate-fuzzer wrote with an LLM (the
// *_fuzzer.go files under pkg/controller/direct with a "+tool:fuzz-gen" header) it
// classifies the fields of the resource message the way generate-fuzzer --deterministic
// does, and compares the result with the fuzzer's entries. The design's target is that at
// least 90% of the fields are classified the same way.
//
// It reads the googleapis descriptor set pinned in apis/git.versions, which
// generate-proto.sh writes to .build/googleapis-<sha>.pb (.build/googleapis.pb is
// overwritten by the generate.sh scripts that pin another version), and the descriptor sets
// that those scripts pin for their own service. It only runs when asked:
//
//	KCC_FUZZER_AGREEMENT=1 go test ./pkg/codegen -run TestFuzzerAgreement -v
//
// KCC_FUZZER_AGREEMENT_REPORT=<file> also writes the full report, with every kind, there.
func TestFuzzerAgreement(t *testing.T) {
	if os.Getenv("KCC_FUZZER_AGREEMENT") != "1" {
		t.Skip("set KCC_FUZZER_AGREEMENT=1 to compare the deterministic fuzzer classification with the existing fuzzers")
	}
	root, err := filepath.Abs(filepath.Join("..", "..", "..", "..", ".."))
	if err != nil {
		t.Fatalf("finding the repo root: %v", err)
	}
	sha, err := pinnedGoogleapisSHA(root)
	if err != nil {
		t.Fatalf("reading the pinned googleapis version: %v", err)
	}
	defaultPB := filepath.Join(root, ".build", "googleapis-"+sha+".pb")
	if _, err := os.Stat(defaultPB); err != nil {
		t.Skipf("the pinned googleapis descriptor set %s is missing; dev/tools/controllerbuilder/generate-proto.sh writes it", defaultPB)
	}

	check := &fuzzerAgreementCheck{
		root:      root,
		defaultPB: defaultPB,
		protos:    make(map[string]*protoapi.Proto),
		directs:   make(map[string]*goPackageIndex),
		krm:       newFuzzerKRMLoader(kccGoModule+"/apis", filepath.Join(root, "apis")),
	}
	files, err := check.findLLMFuzzers()
	if err != nil {
		t.Fatalf("finding the fuzzers: %v", err)
	}
	for _, file := range files {
		check.compareFile(file)
	}

	summary, full := check.report()
	t.Log("\n" + summary)
	if p := os.Getenv("KCC_FUZZER_AGREEMENT_REPORT"); p != "" {
		if err := os.WriteFile(p, []byte(full), 0644); err != nil {
			t.Errorf("writing the report: %v", err)
		}
	}
	if agreed, total, _, _ := check.totals(); total == 0 || float64(agreed)/float64(total) < fuzzerAgreementTarget {
		t.Errorf("%d of %d fields are classified the way the existing fuzzers do; the target is %.0f%%", agreed, total, 100*fuzzerAgreementTarget)
	}
}

func pinnedGoogleapisSHA(root string) (string, error) {
	b, err := os.ReadFile(filepath.Join(root, "apis", "git.versions"))
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(b), "\n") {
		if fields := strings.Fields(line); len(fields) == 2 && fields[0] == "https://github.com/googleapis/googleapis" {
			return fields[1], nil
		}
	}
	return "", errors.New("apis/git.versions does not pin github.com/googleapis/googleapis")
}

// fieldClass is what the fuzz tests do with a field.
type fieldClass string

const (
	// fieldSpec is cleared by FuzzStatus only.
	fieldSpec fieldClass = "spec"
	// fieldStatus is cleared by FuzzSpec only.
	fieldStatus fieldClass = "status"
	// fieldUnimplemented is cleared by both.
	fieldUnimplemented fieldClass = "unimplemented"
	// fieldUnlisted is cleared by neither: it has to round-trip through both mappers.
	fieldUnlisted fieldClass = "unlisted"
)

// clears reports whether FuzzSpec and FuzzStatus clear a field of the class.
func (c fieldClass) clears() (inSpec, inStatus bool) {
	switch c {
	case fieldSpec:
		return false, true
	case fieldStatus:
		return true, false
	case fieldUnimplemented:
		return true, true
	}
	return false, false
}

// disagreementDirection says whether the deterministic class tests a field in more fuzz tests
// than the LLM class (stricter), in fewer (looser), or in different ones (swapped). FuzzStatus
// doesn't run for a spec-only fuzzer.
func disagreementDirection(llm, det fieldClass, llmSpecOnly, detSpecOnly bool) string {
	llmSpec, llmStatus := llm.clears()
	detSpec, detStatus := det.clears()
	if llmSpecOnly {
		llmStatus = true
	}
	if detSpecOnly {
		detStatus = true
	}
	switch {
	case llmSpec == detSpec && llmStatus == detStatus:
		return "same tests"
	case (detSpec == llmSpec || !detSpec) && (detStatus == llmStatus || !detStatus):
		return "stricter"
	case (detSpec == llmSpec || detSpec) && (detStatus == llmStatus || detStatus):
		return "looser"
	}
	return "swapped"
}

// fuzzerListing holds the entries of a fuzzer, as field path -> the method that lists it.
type fuzzerListing struct {
	spec, status, unimplemented map[string]string
	// specOnly is set for spec-only fuzzers, whose FuzzStatus never runs.
	specOnly bool
}

func newFuzzerListing(specOnly bool) *fuzzerListing {
	return &fuzzerListing{
		spec:          make(map[string]string),
		status:        make(map[string]string),
		unimplemented: make(map[string]string),
		specOnly:      specOnly,
	}
}

func (l *fuzzerListing) add(method, fieldPath string) {
	switch method {
	case "SpecField", "SpecFields.Insert":
		l.spec[fieldPath] = method
	case "StatusField", "StatusFields.Insert":
		l.status[fieldPath] = method
	default:
		l.unimplemented[fieldPath] = method
	}
}

func (l *fuzzerListing) paths() []string {
	var out []string
	for _, m := range []map[string]string{l.spec, l.status, l.unimplemented} {
		for p := range m {
			out = append(out, p)
		}
	}
	sort.Strings(out)
	return out
}

// class returns what the fuzz tests do with the field at fieldPath, given the entries for
// the field and for the fields that contain it, and the method of the nearest entry.
// FuzzSpec clears status and unimplemented fields, FuzzStatus clears spec and unimplemented
// fields, and clearing a field clears the fields below it.
func (l *fuzzerListing) class(fieldPath string) (fieldClass, string) {
	clearedInSpec, clearedInStatus := false, false
	method := ""
	for _, p := range fieldPathAndParents(fieldPath) {
		if m, ok := l.unimplemented[p]; ok {
			clearedInSpec, clearedInStatus = true, true
			if method == "" {
				method = m
			}
		}
		if m, ok := l.status[p]; ok {
			clearedInSpec = true
			if method == "" {
				method = m
			}
		}
		if m, ok := l.spec[p]; ok {
			clearedInStatus = true
			if method == "" {
				method = m
			}
		}
	}
	switch {
	case clearedInSpec && clearedInStatus:
		return fieldUnimplemented, method
	case clearedInStatus:
		return fieldSpec, method
	case clearedInSpec:
		return fieldStatus, method
	case l.specOnly:
		// Only FuzzSpec runs, and it round-trips the field through the Spec mappers.
		return fieldSpec, method
	default:
		return fieldUnlisted, method
	}
}

// fieldPathAndParents returns the path and the paths of the fields that contain it, for
// example .a.b[].c, .a.b and .a.
func fieldPathAndParents(fieldPath string) []string {
	var out []string
	for fieldPath != "" {
		out = append(out, fieldPath)
		i := strings.LastIndex(fieldPath, ".")
		if i <= 0 {
			break
		}
		fieldPath = strings.TrimSuffix(fieldPath[:i], "[]")
	}
	return out
}

// resolveFuzzerPath reports whether fieldPath names a field of msg in the fuzz.Visit grammar.
func resolveFuzzerPath(msg protoreflect.MessageDescriptor, fieldPath string) bool {
	if !strings.HasPrefix(fieldPath, ".") {
		return false
	}
	segments := strings.Split(fieldPath[1:], ".")
	for i, segment := range segments {
		name, elements := strings.CutSuffix(segment, "[]")
		fd := msg.Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			return false
		}
		if i == len(segments)-1 {
			return !elements
		}
		if fd.Kind() != protoreflect.MessageKind || fd.IsMap() || fd.IsList() != elements {
			return false
		}
		msg = fd.Message()
	}
	return false
}

// fuzzerAgreementUnit is a field that both fuzzers treat as a whole: neither lists a field
// below it.
type fuzzerAgreementUnit struct {
	path string
	// leaves is the number of proto leaves at or below the field, down to fuzzerAgreementLeafDepth.
	leaves               int
	llm, det             fieldClass
	llmMethod, detMethod string
}

// fuzzerAgreementUnits splits the fields of msg until neither listing lists a field below them.
func fuzzerAgreementUnits(msg protoreflect.MessageDescriptor, llm, det *fuzzerListing) []fuzzerAgreementUnit {
	split := make(map[string]bool)
	for _, l := range []*fuzzerListing{llm, det} {
		for _, p := range l.paths() {
			for _, parent := range fieldPathAndParents(p)[1:] {
				split[parent] = true
			}
		}
	}
	var units []fuzzerAgreementUnit
	var walk func(prefix string, msg protoreflect.MessageDescriptor, depth int)
	walk = func(prefix string, msg protoreflect.MessageDescriptor, depth int) {
		fields := msg.Fields()
		for i := 0; i < fields.Len(); i++ {
			fd := fields.Get(i)
			fieldPath := prefix + "." + string(fd.Name())
			if split[fieldPath] && fd.Kind() == protoreflect.MessageKind && !fd.IsMap() {
				childPrefix := fieldPath
				if fd.IsList() {
					childPrefix += "[]"
				}
				walk(childPrefix, fd.Message(), depth+1)
				continue
			}
			u := fuzzerAgreementUnit{path: fieldPath, leaves: countFuzzerLeaves(fd, depth, make(map[protoreflect.FullName]bool))}
			u.llm, u.llmMethod = llm.class(fieldPath)
			u.det, u.detMethod = det.class(fieldPath)
			units = append(units, u)
		}
	}
	walk("", msg, 1)
	return units
}

func countFuzzerLeaves(fd protoreflect.FieldDescriptor, depth int, onPath map[protoreflect.FullName]bool) int {
	if fd.Kind() != protoreflect.MessageKind || fd.IsMap() || isOpaqueFuzzerMessage(fd.Message()) ||
		depth >= fuzzerAgreementLeafDepth || onPath[fd.Message().FullName()] {
		return 1
	}
	msg := fd.Message()
	onPath[msg.FullName()] = true
	defer delete(onPath, msg.FullName())
	n := 0
	for i := 0; i < msg.Fields().Len(); i++ {
		n += countFuzzerLeaves(msg.Fields().Get(i), depth+1, onPath)
	}
	return n
}

// llmFuzzer is what the agreement check reads from a fuzzer written by an LLM.
type llmFuzzer struct {
	// protoMessage is the proto.message attribute of the +tool:fuzz-gen header.
	protoMessage string
	constructor  string
	// protoType is the Go type of the proto message, X in &pb.X{}.
	protoType string
	// specFromProto and statusFromProto are the mappers the fuzzer is built from, without a
	// package qualifier. mapperGoPackage is the import path of their package, or "" for the
	// fuzzer's own package.
	specFromProto, statusFromProto string
	mapperGoPackage                string
	entries                        [][2]string // method, path
	// nonLiteral counts the entries whose path is not a string literal.
	nonLiteral int
}

// fuzzerAgreementSkip is why a fuzzer could not be compared.
type fuzzerAgreementSkip struct {
	reason string
	detail string
}

func (s *fuzzerAgreementSkip) Error() string { return s.reason + ": " + s.detail }

func skipAgreement(reason, format string, args ...any) error {
	return &fuzzerAgreementSkip{reason: reason, detail: fmt.Sprintf(format, args...)}
}

func parseLLMFuzzer(p string) (*llmFuzzer, error) {
	src, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	out := &llmFuzzer{}
	inHeader := false
	scanner := bufio.NewScanner(bytes.NewReader(src))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "// +tool:fuzz-gen" {
			inHeader = true
			continue
		}
		if !inHeader {
			continue
		}
		attr, ok := strings.CutPrefix(line, "//")
		if !ok {
			break
		}
		if k, v, ok := strings.Cut(attr, ":"); ok && strings.TrimSpace(k) == "proto.message" {
			out.protoMessage = strings.TrimSpace(v)
		}
	}

	file, err := parser.ParseFile(token.NewFileSet(), p, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	imports := make(map[string]string)
	for _, imp := range file.Imports {
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return nil, err
		}
		name := path.Base(importPath)
		if imp.Name != nil {
			name = imp.Name.Name
		}
		imports[name] = importPath
	}

	mapperName := func(expr ast.Expr) (name, goPackage string) {
		switch e := expr.(type) {
		case *ast.Ident:
			return e.Name, ""
		case *ast.SelectorExpr:
			if x, ok := e.X.(*ast.Ident); ok {
				return e.Sel.Name, imports[x.Name]
			}
		}
		return "", ""
	}

	// stringLists holds the variables assigned a []string of literals. An LLM sometimes lists
	// the same fields below several paths with a loop over such a variable, as the
	// MonitoringDashboard fuzzer does for its widget paths.
	stringLists := make(map[string][]string)
	var inspect func(root ast.Node, env map[string]string)
	inspect = func(root ast.Node, env map[string]string) {
		ast.Inspect(root, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.AssignStmt:
				if len(n.Lhs) == 1 && len(n.Rhs) == 1 {
					if id, ok := n.Lhs[0].(*ast.Ident); ok {
						if list, ok := stringListLiteral(n.Rhs[0]); ok {
							stringLists[id.Name] = list
						}
					}
				}
				return true
			case *ast.RangeStmt:
				value, ok := n.Value.(*ast.Ident)
				if !ok {
					return true
				}
				list, ok := stringListLiteral(n.X)
				if id, isIdent := n.X.(*ast.Ident); isIdent {
					list, ok = stringLists[id.Name]
				}
				if !ok {
					return true
				}
				for _, element := range list {
					inner := map[string]string{value.Name: element}
					for k, v := range env {
						if k != value.Name {
							inner[k] = v
						}
					}
					inspect(n.Body, inner)
				}
				return false
			case *ast.CallExpr:
				parseLLMFuzzerCall(out, n, env, mapperName)
				return true
			}
			return true
		})
	}
	inspect(file, nil)
	return out, nil
}

// parseLLMFuzzerCall records the fuzzer constructor or entry that call makes, if any. env
// holds the values of the loop variables around the call.
func parseLLMFuzzerCall(out *llmFuzzer, call *ast.CallExpr, env map[string]string, mapperName func(ast.Expr) (string, string)) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		// Generic instantiations such as fuzztesting.NewKRMTypedFuzzer[...](...).
		if index, isIndex := call.Fun.(*ast.IndexListExpr); isIndex {
			sel, ok = index.X.(*ast.SelectorExpr)
		}
		if !ok {
			return
		}
	}
	method := sel.Sel.Name
	switch {
	case strings.HasPrefix(method, "NewKRMTyped"):
		if out.constructor != "" {
			out.constructor = "more than one fuzzer"
			return
		}
		out.constructor = method
		if method != "NewKRMTypedFuzzer" && method != "NewKRMTypedSpecFuzzer" || len(call.Args) < 3 {
			return
		}
		if unary, ok := call.Args[0].(*ast.UnaryExpr); ok {
			if lit, ok := unary.X.(*ast.CompositeLit); ok {
				if typeSel, ok := lit.Type.(*ast.SelectorExpr); ok {
					out.protoType = typeSel.Sel.Name
				}
			}
		}
		out.specFromProto, out.mapperGoPackage = mapperName(call.Args[1])
		if method == "NewKRMTypedFuzzer" && len(call.Args) >= 5 {
			out.statusFromProto, _ = mapperName(call.Args[3])
		}
	case method == "Insert":
		inner, ok := sel.X.(*ast.SelectorExpr)
		if !ok {
			return
		}
		switch inner.Sel.Name {
		case "SpecFields", "StatusFields", "UnimplementedFields":
		default:
			return
		}
		for _, arg := range call.Args {
			if s, ok := evalStringExpr(arg, env); ok {
				out.entries = append(out.entries, [2]string{inner.Sel.Name + ".Insert", s})
			} else {
				out.nonLiteral++
			}
		}
	case method == "Unimplemented_Etag":
		out.entries = append(out.entries, [2]string{method, fuzzerEtagPath})
	case method == "SpecField" || method == "StatusField" || method == "IdentityField" ||
		method == "Ignore_JSONBookkeeping" || strings.HasPrefix(method, "Unimplemented_"):
		if len(call.Args) == 1 {
			if s, ok := evalStringExpr(call.Args[0], env); ok {
				out.entries = append(out.entries, [2]string{method, s})
				return
			}
		}
		out.nonLiteral++
	}
}

// evalStringExpr evaluates a string literal, a loop variable in env, or a concatenation of those.
func evalStringExpr(expr ast.Expr, env map[string]string) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind != token.STRING {
			return "", false
		}
		s, err := strconv.Unquote(e.Value)
		return s, err == nil
	case *ast.Ident:
		s, ok := env[e.Name]
		return s, ok
	case *ast.ParenExpr:
		return evalStringExpr(e.X, env)
	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", false
		}
		x, ok := evalStringExpr(e.X, env)
		if !ok {
			return "", false
		}
		y, ok := evalStringExpr(e.Y, env)
		return x + y, ok
	}
	return "", false
}

// stringListLiteral evaluates a []string composite literal whose elements are string literals.
func stringListLiteral(expr ast.Expr) ([]string, bool) {
	lit, ok := expr.(*ast.CompositeLit)
	if !ok {
		return nil, false
	}
	arr, ok := lit.Type.(*ast.ArrayType)
	if !ok || arr.Len != nil {
		return nil, false
	}
	if elt, ok := arr.Elt.(*ast.Ident); !ok || elt.Name != "string" {
		return nil, false
	}
	var out []string
	for _, e := range lit.Elts {
		s, ok := evalStringExpr(e, nil)
		if !ok {
			return nil, false
		}
		out = append(out, s)
	}
	return out, true
}

// fuzzerAgreementKind is the comparison for one fuzzer.
type fuzzerAgreementKind struct {
	file string
	kind string
	// specOnly and detSpecOnly are set when the LLM fuzzer, or the deterministic one, is spec-only.
	specOnly, detSpecOnly bool
	units                 []fuzzerAgreementUnit
	// llmUnknown and detUnknown are the entries that name no field of the message.
	llmUnknown, detUnknown []string
	nonLiteral             int
	notes                  []string
}

func (k *fuzzerAgreementKind) counts() (agreed, total, agreedLeaves, totalLeaves int) {
	for _, u := range k.units {
		total++
		totalLeaves += u.leaves
		if u.llm == u.det {
			agreed++
			agreedLeaves += u.leaves
		}
	}
	return agreed, total, agreedLeaves, totalLeaves
}

type fuzzerAgreementCheck struct {
	root      string
	defaultPB string
	protos    map[string]*protoapi.Proto
	directs   map[string]*goPackageIndex
	krm       *fuzzerKRMLoader

	found    int
	kinds    []*fuzzerAgreementKind
	skipped  []fuzzerAgreementSkipped
	unmarked []string
}

type fuzzerAgreementSkipped struct {
	file string
	skip *fuzzerAgreementSkip
}

// findLLMFuzzers returns the *_fuzzer.go files with a +tool:fuzz-gen header, relative to the repo root.
func (c *fuzzerAgreementCheck) findLLMFuzzers() ([]string, error) {
	var out []string
	err := filepath.WalkDir(filepath.Join(c.root, "pkg", "controller", "direct"), func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if !bytes.Contains(b, []byte("\n// +tool:fuzz-gen\n")) {
			return nil
		}
		rel, err := filepath.Rel(c.root, p)
		if err != nil {
			return err
		}
		if !strings.HasSuffix(p, "_fuzzer.go") {
			c.unmarked = append(c.unmarked, rel)
			return nil
		}
		out = append(out, rel)
		return nil
	})
	sort.Strings(out)
	c.found = len(out)
	return out, err
}

func (c *fuzzerAgreementCheck) compareFile(rel string) {
	kind, err := c.compare(rel)
	if err != nil {
		var skip *fuzzerAgreementSkip
		if !errors.As(err, &skip) {
			skip = &fuzzerAgreementSkip{reason: "error", detail: err.Error()}
		}
		c.skipped = append(c.skipped, fuzzerAgreementSkipped{file: rel, skip: skip})
		return
	}
	c.kinds = append(c.kinds, kind)
}

var specFromProtoPattern = regexp.MustCompile(`^(\w+)Spec(?:_(v\d+(?:alpha\d+|beta\d+)?))?_FromProto$`)

func (c *fuzzerAgreementCheck) compare(rel string) (result *fuzzerAgreementKind, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = skipAgreement("panic", "%v", r)
		}
	}()

	llm, err := parseLLMFuzzer(filepath.Join(c.root, rel))
	if err != nil {
		return nil, skipAgreement("cannot parse the fuzzer", "%v", err)
	}
	if llm.constructor != "NewKRMTypedFuzzer" && llm.constructor != "NewKRMTypedSpecFuzzer" {
		return nil, skipAgreement("not a proto KRMTypedFuzzer", "constructor %q", llm.constructor)
	}
	m := specFromProtoPattern.FindStringSubmatch(llm.specFromProto)
	if m == nil {
		return nil, skipAgreement("Spec mapper is not named <Kind>Spec[_<version>]_FromProto", "%s", llm.specFromProto)
	}
	kindName := m[1]

	directDir := filepath.Dir(filepath.Join(c.root, rel))
	if llm.mapperGoPackage != "" {
		sub, ok := strings.CutPrefix(llm.mapperGoPackage, kccGoModule+"/")
		if !ok {
			return nil, skipAgreement("mappers outside the repo", "%s", llm.mapperGoPackage)
		}
		directDir = filepath.Join(c.root, filepath.FromSlash(sub))
	}
	direct := c.directs[directDir]
	if direct == nil {
		if direct, err = loadGoPackageIndex(directDir); err != nil {
			return nil, skipAgreement("cannot parse the mapper package", "%v", err)
		}
		c.directs[directDir] = direct
	}

	// The KRM package of the Spec mapper gives the API version and the service that
	// generate-fuzzer --deterministic would be run with.
	fileSpec, err := direct.mapperPair(strings.TrimSuffix(llm.specFromProto, "_FromProto"))
	if err != nil {
		return nil, skipAgreement("cannot read the Spec mappers", "%v", err)
	}
	if fileSpec == nil {
		return nil, skipAgreement("Spec mappers not found", "%s in %s", llm.specFromProto, directDir)
	}
	krmPath, ok := strings.CutPrefix(fileSpec.KRMGoPackage, kccGoModule+"/apis/")
	if !ok || strings.Count(krmPath, "/") != 1 {
		return nil, skipAgreement("Spec mapper does not return a type of apis/<service>/<version>", "%s returns %s.%s", fileSpec.FromProto, fileSpec.KRMGoPackage, fileSpec.KRMType)
	}
	service, version, _ := strings.Cut(krmPath, "/")

	if llm.protoMessage == "" {
		return nil, skipAgreement("no proto.message in the +tool:fuzz-gen header", "")
	}
	msg, descriptor, err := c.findMessage(service, llm.protoMessage)
	if err != nil {
		return nil, err
	}

	plan, err := planFuzzer(fuzzerPlanInput{kind: kindName, version: version, message: msg, direct: direct, krm: c.krm})
	if err != nil {
		return nil, skipAgreement("deterministic classification failed", "%v", err)
	}

	result = &fuzzerAgreementKind{
		file:        rel,
		kind:        kindName,
		specOnly:    llm.constructor == "NewKRMTypedSpecFuzzer",
		detSpecOnly: plan.Status == nil,
		nonLiteral:  llm.nonLiteral,
	}
	if descriptor != c.defaultPB {
		result.notes = append(result.notes, "descriptor "+filepath.Base(descriptor))
	}
	if plan.Spec.FromProto != llm.specFromProto {
		result.notes = append(result.notes, fmt.Sprintf("Spec mapper %s, the fuzzer uses %s", plan.Spec.FromProto, llm.specFromProto))
	}
	switch {
	case plan.Status == nil && llm.statusFromProto != "":
		result.notes = append(result.notes, fmt.Sprintf("spec-only, the fuzzer uses %s", llm.statusFromProto))
	case plan.Status != nil && llm.statusFromProto == "":
		result.notes = append(result.notes, fmt.Sprintf("uses %s, the fuzzer is spec-only", plan.Status.FromProto))
	case plan.Status != nil && plan.Status.FromProto != llm.statusFromProto:
		result.notes = append(result.notes, fmt.Sprintf("Status mapper %s, the fuzzer uses %s", plan.Status.FromProto, llm.statusFromProto))
	}

	llmListing := newFuzzerListing(result.specOnly)
	for _, e := range llm.entries {
		llmListing.add(e[0], e[1])
		if !resolveFuzzerPath(msg, e[1]) {
			result.llmUnknown = append(result.llmUnknown, e[1])
		}
	}
	detListing := newFuzzerListing(plan.Status == nil)
	for _, e := range plan.Entries {
		detListing.add(string(e.Method), e.Path)
		if !resolveFuzzerPath(msg, e.Path) {
			result.detUnknown = append(result.detUnknown, e.Path)
		}
	}
	result.units = fuzzerAgreementUnits(msg, llmListing, detListing)
	return result, nil
}

var (
	protoSHAPattern        = regexp.MustCompile(`PROTO_SHA="?([0-9a-f]{40})"?`)
	protoSourcePathPattern = regexp.MustCompile(`--proto-source-path\s+"?\$\{REPO_ROOT\}/\.build/([A-Za-z0-9._-]+\.pb)"?`)
)

// descriptorFor returns the descriptor set that apis/<service>/generate.sh compiles the
// service with, when it pins its own googleapis version and the file exists, else the default.
func (c *fuzzerAgreementCheck) descriptorFor(service string) string {
	b, err := os.ReadFile(filepath.Join(c.root, "apis", service, "generate.sh"))
	if err != nil {
		return c.defaultPB
	}
	var candidates []string
	if m := protoSHAPattern.FindSubmatch(b); m != nil {
		candidates = append(candidates, filepath.Join(c.root, ".build", "googleapis-"+string(m[1])+".pb"))
	}
	if m := protoSourcePathPattern.FindSubmatch(b); m != nil {
		candidates = append(candidates, filepath.Join(c.root, ".build", string(m[1])))
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return c.defaultPB
}

func (c *fuzzerAgreementCheck) loadProto(p string) (*protoapi.Proto, error) {
	if api := c.protos[p]; api != nil {
		return api, nil
	}
	api, err := protoapi.LoadProto(p, "")
	if err != nil {
		return nil, err
	}
	c.protos[p] = api
	return api, nil
}

// findMessage looks the message up in the service's descriptor set, then in the default one.
func (c *fuzzerAgreementCheck) findMessage(service, fullName string) (protoreflect.MessageDescriptor, string, error) {
	for _, p := range []string{c.descriptorFor(service), c.defaultPB} {
		api, err := c.loadProto(p)
		if err != nil {
			return nil, "", skipAgreement("cannot load the descriptor set", "%s: %v", p, err)
		}
		desc, err := api.Files().FindDescriptorByName(protoreflect.FullName(fullName))
		if err != nil {
			continue
		}
		if msg, ok := desc.(protoreflect.MessageDescriptor); ok {
			return msg, p, nil
		}
	}
	return nil, "", skipAgreement("proto message not in the descriptor set", "%s", fullName)
}

func (c *fuzzerAgreementCheck) totals() (agreed, total, agreedLeaves, totalLeaves int) {
	for _, k := range c.kinds {
		a, t, al, tl := k.counts()
		agreed += a
		total += t
		agreedLeaves += al
		totalLeaves += tl
	}
	return agreed, total, agreedLeaves, totalLeaves
}

func percent(n, d int) string {
	if d == 0 {
		return "n/a"
	}
	return fmt.Sprintf("%.1f%%", 100*float64(n)/float64(d))
}

type countedKey struct {
	key   string
	count int
	kinds map[string]bool
	// examples holds up to a few kind:path examples.
	examples []string
}

// counter counts occurrences of keys, and the kinds they occur in.
type counter map[string]*countedKey

func (c counter) add(key, kind, example string) {
	e := c[key]
	if e == nil {
		e = &countedKey{key: key, kinds: make(map[string]bool)}
		c[key] = e
	}
	e.count++
	e.kinds[kind] = true
	if len(e.examples) < 4 {
		e.examples = append(e.examples, kind+" "+example)
	}
}

func (c counter) sorted() []*countedKey {
	var out []*countedKey
	for _, e := range c {
		out = append(out, e)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].count != out[j].count {
			return out[i].count > out[j].count
		}
		return out[i].key < out[j].key
	})
	return out
}

// report returns a summary, and the full report with a row per kind.
func (c *fuzzerAgreementCheck) report() (string, string) {
	var b strings.Builder
	agreed, total, agreedLeaves, totalLeaves := c.totals()
	fmt.Fprintf(&b, "# Deterministic fuzzer classification vs. the fuzzers written by an LLM\n\n")
	fmt.Fprintf(&b, "Fuzzers with a +tool:fuzz-gen header: %d; compared: %d; not compared: %d.\n", c.found, len(c.kinds), len(c.skipped))
	if len(c.unmarked) > 0 {
		fmt.Fprintf(&b, "Other files with the header, not fuzzers, ignored: %s.\n", strings.Join(c.unmarked, ", "))
	}
	fmt.Fprintf(&b, "\nA field is compared at the finest level that either fuzzer lists; its class is what the\n")
	fmt.Fprintf(&b, "fuzz tests do with it (spec: cleared by FuzzStatus only; status: cleared by FuzzSpec only;\n")
	fmt.Fprintf(&b, "unimplemented: cleared by both; unlisted: cleared by neither, so it must round-trip through\n")
	fmt.Fprintf(&b, "both mappers; in a spec-only fuzzer an unlisted field counts as spec).\n\n")
	fmt.Fprintf(&b, "- Fields: %d of %d classified the same way (%s); the target is %.0f%%.\n", agreed, total, percent(agreed, total), 100*fuzzerAgreementTarget)
	fmt.Fprintf(&b, "- Proto leaves (depth <= %d), weighting each field by its leaves: %d of %d (%s).\n", fuzzerAgreementLeafDepth, agreedLeaves, totalLeaves, percent(agreedLeaves, totalLeaves))
	allAgree := 0
	for _, k := range c.kinds {
		if a, t, _, _ := k.counts(); a == t {
			allAgree++
		}
	}
	fmt.Fprintf(&b, "- Kinds whose fields all agree: %d of %d.\n", allAgree, len(c.kinds))

	// Distribution of the per-kind agreement.
	buckets := []struct {
		name string
		min  float64
	}{{"100%", 1}, {"90-99%", 0.9}, {"75-89%", 0.75}, {"50-74%", 0.5}, {"<50%", -1}}
	counts := make([]int, len(buckets))
	for _, k := range c.kinds {
		a, t, _, _ := k.counts()
		share := 1.0
		if t > 0 {
			share = float64(a) / float64(t)
		}
		for i, bucket := range buckets {
			if share >= bucket.min {
				counts[i]++
				break
			}
		}
	}
	fmt.Fprintf(&b, "\n## Per-kind agreement\n\n| fields agreeing | kinds |\n|---|---|\n")
	for i, bucket := range buckets {
		fmt.Fprintf(&b, "| %s | %d |\n", bucket.name, counts[i])
	}

	// Disagreement patterns.
	byClass := counter{}
	byField := counter{}
	byDirection := counter{}
	for _, k := range c.kinds {
		for _, u := range k.units {
			if u.llm == u.det {
				continue
			}
			pattern := fmt.Sprintf("LLM %s, deterministic %s", u.llm, u.det)
			byClass.add(pattern, k.kind, u.path)
			byDirection.add(disagreementDirection(u.llm, u.det, k.specOnly, k.detSpecOnly), k.kind, u.path)
			name := u.path[strings.LastIndex(u.path, ".")+1:]
			if !strings.Contains(u.path[1:], ".") {
				name = "." + name + " (top level)"
			}
			byField.add(fmt.Sprintf("%s: %s", pattern, name), k.kind, u.path)
		}
	}
	fmt.Fprintf(&b, "\n## Disagreements by direction\n\n")
	fmt.Fprintf(&b, "Stricter: the deterministic fuzzer round-trips the field in a fuzz test that clears it in the\n")
	fmt.Fprintf(&b, "LLM fuzzer. Looser: the other way around. Swapped: each round-trips it in a test the other clears.\n\n")
	fmt.Fprintf(&b, "| direction | fields | kinds |\n|---|---|---|\n")
	for _, e := range byDirection.sorted() {
		fmt.Fprintf(&b, "| %s | %d | %d |\n", e.key, e.count, len(e.kinds))
	}
	fmt.Fprintf(&b, "\n## Disagreements by class\n\n| LLM -> deterministic | fields | kinds | examples |\n|---|---|---|---|\n")
	for _, e := range byClass.sorted() {
		fmt.Fprintf(&b, "| %s | %d | %d | %s |\n", e.key, e.count, len(e.kinds), strings.Join(e.examples, "; "))
	}

	fmt.Fprintf(&b, "\n## Most frequent disagreeing fields\n\n| pattern: field | fields | kinds |\n|---|---|---|\n")
	for i, e := range byField.sorted() {
		if i == 25 {
			break
		}
		fmt.Fprintf(&b, "| %s | %d | %d |\n", e.key, e.count, len(e.kinds))
	}

	// Methods of the special fields.
	special := counter{}
	for _, k := range c.kinds {
		for _, u := range k.units {
			switch u.path {
			case fuzzerIdentityPath, fuzzerEtagPath, ".labels", ".annotations":
			default:
				continue
			}
			special.add(fmt.Sprintf("%s: LLM %s %s, deterministic %s %s", u.path, u.llm, u.llmMethod, u.det, u.detMethod), k.kind, u.path)
		}
	}
	fmt.Fprintf(&b, "\n## Top-level .name, .etag, .labels and .annotations\n\n| field: LLM class method, deterministic class method | kinds |\n|---|---|\n")
	for _, e := range special.sorted() {
		fmt.Fprintf(&b, "| %s | %d |\n", e.key, e.count)
	}

	// Entries that are not fields of the message.
	unknown := counter{}
	llmUnknownKinds, detUnknown, nonLiteral := 0, 0, 0
	for _, k := range c.kinds {
		if len(k.llmUnknown) > 0 {
			llmUnknownKinds++
		}
		for _, p := range k.llmUnknown {
			unknown.add(p[strings.LastIndex(p, ".")+1:], k.kind, p)
		}
		detUnknown += len(k.detUnknown)
		nonLiteral += k.nonLiteral
	}
	fmt.Fprintf(&b, "\n## Entries that name no field of the message\n\n")
	fmt.Fprintf(&b, "- LLM entries: %d, in %d kinds. They are fields of the Go client library that the pinned\n", len(unknownList(c.kinds)), llmUnknownKinds)
	fmt.Fprintf(&b, "  descriptor set doesn't have, map keys, or mistakes; they are not compared.\n")
	fmt.Fprintf(&b, "- Deterministic entries: %d.\n", detUnknown)
	fmt.Fprintf(&b, "- LLM entries whose path is not a string literal (not compared): %d.\n", nonLiteral)
	if len(unknown) > 0 {
		fmt.Fprintf(&b, "\n| field name | entries | kinds | examples |\n|---|---|---|---|\n")
		for i, e := range unknown.sorted() {
			if i == 15 {
				break
			}
			fmt.Fprintf(&b, "| %s | %d | %d | %s |\n", e.key, e.count, len(e.kinds), strings.Join(e.examples, "; "))
		}
	}

	// Fuzzers that could not be compared.
	reasons := counter{}
	for _, s := range c.skipped {
		reasons.add(s.skip.reason, s.file, s.skip.detail)
	}
	fmt.Fprintf(&b, "\n## Fuzzers not compared\n\n| reason | fuzzers | examples |\n|---|---|---|\n")
	for _, e := range reasons.sorted() {
		fmt.Fprintf(&b, "| %s | %d | %s |\n", e.key, e.count, strings.Join(e.examples, "; "))
	}
	summary := b.String()

	// The full report adds every kind.
	fmt.Fprintf(&b, "\n## Kinds\n\n| kind | fuzzer | fields | agreeing | leaves agreeing | notes |\n|---|---|---|---|---|---|\n")
	sorted := append([]*fuzzerAgreementKind(nil), c.kinds...)
	sort.Slice(sorted, func(i, j int) bool {
		ai, ti, _, _ := sorted[i].counts()
		aj, tj, _, _ := sorted[j].counts()
		si, sj := float64(ai)/float64(max(ti, 1)), float64(aj)/float64(max(tj, 1))
		if si != sj {
			return si < sj
		}
		return sorted[i].file < sorted[j].file
	})
	for _, k := range sorted {
		a, t, al, tl := k.counts()
		notes := append([]string(nil), k.notes...)
		if k.specOnly {
			notes = append(notes, "spec-only fuzzer")
		}
		var disagreements []string
		for _, u := range k.units {
			if u.llm != u.det {
				disagreements = append(disagreements, fmt.Sprintf("%s %s->%s", u.path, u.llm, u.det))
			}
		}
		if len(disagreements) > 0 {
			notes = append(notes, "differ: "+strings.Join(disagreements, ", "))
		}
		fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %s |\n", k.kind, k.file, t, percent(a, t), percent(al, tl), strings.Join(notes, "; "))
	}
	fmt.Fprintf(&b, "\n## Fuzzers not compared, in full\n\n")
	for _, s := range c.skipped {
		fmt.Fprintf(&b, "- %s: %s: %s\n", s.file, s.skip.reason, s.skip.detail)
	}
	return summary, b.String()
}

func unknownList(kinds []*fuzzerAgreementKind) []string {
	var out []string
	for _, k := range kinds {
		out = append(out, k.llmUnknown...)
	}
	return out
}

func TestFuzzerListingClass(t *testing.T) {
	// Arrange
	l := newFuzzerListing(false)
	l.add("SpecField", ".a")
	l.add("Unimplemented_NotYetTriaged", ".a.b")
	l.add("StatusFields.Insert", ".c[].d")
	l.add("SpecField", ".e")
	l.add("StatusField", ".e")
	specOnly := newFuzzerListing(true)
	specOnly.add("StatusField", ".s")

	for _, tc := range []struct {
		listing *fuzzerListing
		path    string
		want    fieldClass
	}{
		{listing: l, path: ".a", want: fieldSpec},
		{listing: l, path: ".a.x", want: fieldSpec},
		{listing: l, path: ".a.b.y", want: fieldUnimplemented},
		{listing: l, path: ".c[].d", want: fieldStatus},
		{listing: l, path: ".c[].z", want: fieldUnlisted},
		// Listed as both a spec field and a status field, both fuzz tests clear it.
		{listing: l, path: ".e", want: fieldUnimplemented},
		{listing: specOnly, path: ".s", want: fieldStatus},
		{listing: specOnly, path: ".t", want: fieldSpec},
	} {
		// Act
		got, _ := tc.listing.class(tc.path)

		// Assert
		if got != tc.want {
			t.Errorf("class(%q) = %s, want %s", tc.path, got, tc.want)
		}
	}
}

func TestFuzzerAgreementUnits(t *testing.T) {
	// Arrange
	msg := fuzzerTestMessage(t, "Widget")
	llm := newFuzzerListing(false)
	llm.add("SpecField", ".config")
	llm.add("StatusField", ".rules")
	det := newFuzzerListing(false)
	det.add("SpecField", ".config.size")
	det.add("StatusField", ".config.state")
	det.add("SpecField", ".rules[].action")
	det.add("StatusField", ".rules[].hit_count")

	// Act
	units := fuzzerAgreementUnits(msg, llm, det)

	// Assert
	got := make(map[string]string)
	for _, u := range units {
		got[u.path] = fmt.Sprintf("%s/%s", u.llm, u.det)
	}
	want := map[string]string{
		".name":              "unlisted/unlisted",
		".etag":              "unlisted/unlisted",
		".labels":            "unlisted/unlisted",
		".display_name":      "unlisted/unlisted",
		".create_time":       "unlisted/unlisted",
		".config.size":       "spec/spec",
		".config.state":      "spec/status",
		".config.extra":      "spec/unlisted",
		".rules[].action":    "status/spec",
		".rules[].hit_count": "status/status",
		".network":           "unlisted/unlisted",
		".unmapped":          "unlisted/unlisted",
	}
	if len(got) != len(want) {
		t.Errorf("got %d units, want %d: %v", len(got), len(want), got)
	}
	for p, w := range want {
		if got[p] != w {
			t.Errorf("unit %s = %q, want %q", p, got[p], w)
		}
	}
	if !resolveFuzzerPath(msg, ".rules[].hit_count") || resolveFuzzerPath(msg, ".rules.hit_count") || resolveFuzzerPath(msg, ".config[].size") {
		t.Errorf("resolveFuzzerPath does not follow the fuzz.Visit grammar")
	}
}
