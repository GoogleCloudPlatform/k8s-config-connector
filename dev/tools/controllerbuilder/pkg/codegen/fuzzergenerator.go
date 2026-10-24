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
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

const (
	fuzztestingGoPackage    = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	FuzzerFileAnnotationKey = "+generated:fuzzer"
	fuzzerMaxDepth          = 11
)

type FuzzerMethod string

const (
	FuzzerSpecField                      FuzzerMethod = "SpecField"
	FuzzerStatusField                    FuzzerMethod = "StatusField"
	FuzzerUnimplementedIdentity          FuzzerMethod = "Unimplemented_Identity"
	FuzzerUnimplementedEtag              FuzzerMethod = "Unimplemented_Etag"
	FuzzerUnimplementedLabelsAnnotations FuzzerMethod = "Unimplemented_LabelsAnnotations"
	FuzzerUnimplementedNotYetTriaged     FuzzerMethod = "Unimplemented_NotYetTriaged"
)

type FuzzerEntry struct {
	Method FuzzerMethod
	Path   string
	Hint   string
}

var fuzzerEntryGroups = []struct {
	comment string
	methods []FuzzerMethod
}{
	{"Spec fields", []FuzzerMethod{FuzzerSpecField}},
	{"Status fields", []FuzzerMethod{FuzzerStatusField}},
	{"Identity / Special fields", []FuzzerMethod{FuzzerUnimplementedIdentity, FuzzerUnimplementedEtag, FuzzerUnimplementedLabelsAnnotations}},
	{"Unimplemented / Not Yet Triaged fields", []FuzzerMethod{FuzzerUnimplementedNotYetTriaged}},
}

type FuzzerGeneratorOptions struct {
	Group            string
	Version          string
	ProtoService     string
	APIGoPackagePath string
	APIDirectory     string
	DirectDirectory  string
	Year             int
}

type FuzzerGenerator struct {
	opts FuzzerGeneratorOptions
	krm  *fuzzerKRMLoader
}

type FuzzerResult struct {
	Path       string
	Skipped    bool
	SkipReason string
	Plan       *FuzzerPlan
}

type FuzzerPlan struct {
	Kind         string
	Message      protoreflect.MessageDescriptor
	PackageName  string
	FuncName     string
	Spec         *fuzzerMapperPair
	Status       *fuzzerMapperPair
	Entries      []FuzzerEntry
	FilterSpec   string
	FilterStatus string
}

type fuzzerPlanInput struct {
	kind    string
	version string
	message protoreflect.MessageDescriptor
	direct  *goPackageIndex
	krm     *fuzzerKRMLoader
}

func NewFuzzerGenerator(opts FuzzerGeneratorOptions) *FuzzerGenerator {
	return &FuzzerGenerator{opts: opts, krm: newFuzzerKRMLoader(opts.APIGoPackagePath, opts.APIDirectory)}
}

func (g *FuzzerGenerator) KRMGoPackage() string {
	service := strings.TrimSuffix(g.opts.Group, ".cnrm.cloud.google.com")
	return strings.TrimSuffix(g.opts.APIGoPackagePath, "/") + "/" + service + "/" + g.opts.Version
}

func FuzzerFileName(kind string) string { return strings.ToLower(kind) + "_fuzzer.generated.go" }

func (g *FuzzerGenerator) Generate(kind string, msg protoreflect.MessageDescriptor) (*FuzzerResult, error) {
	direct, err := loadGoPackageIndex(g.opts.DirectDirectory)
	if err != nil {
		return nil, err
	}
	result := &FuzzerResult{Path: filepath.Join(g.opts.DirectDirectory, FuzzerFileName(kind))}
	funcName := fuzzerFuncName(kind)
	if p := direct.handwrittenFunc(funcName); p != "" {
		result.Skipped, result.SkipReason = true, fmt.Sprintf("%s declares %s", p, funcName)
		klog.Infof("not generating %s: %s", result.Path, result.SkipReason)
		return result, removeStaleFile(result.Path)
	}

	plan, err := planFuzzer(fuzzerPlanInput{kind: kind, version: g.opts.Version, message: msg, direct: direct, krm: g.krm})
	if err != nil {
		return nil, err
	}
	if want := g.KRMGoPackage(); plan.Spec.KRMGoPackage != want {
		return nil, fmt.Errorf("%s returns a type of %s, but --api-version %s/%s is %s", plan.Spec.FromProto, plan.Spec.KRMGoPackage, g.opts.Group, g.opts.Version, want)
	}
	result.Plan = plan

	year, err := g.resolveCopyrightYear(result.Path)
	if err != nil {
		return nil, err
	}
	src, err := renderFuzzer(plan, g.fileAnnotation(msg), year)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(result.Path, src, 0644); err != nil {
		return nil, fmt.Errorf("writing %q: %w", result.Path, err)
	}
	return result, nil
}

func removeStaleFile(filePath string) error {
	if err := os.Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing stale %q: %w", filePath, err)
	}
	return nil
}

func (g *FuzzerGenerator) resolveCopyrightYear(outPath string) (int, error) {
	year := g.opts.Year
	if year == 0 {
		year = time.Now().Year()
	}
	existing, err := os.ReadFile(outPath)
	if err != nil {
		if os.IsNotExist(err) {
			return year, nil
		}
		return 0, fmt.Errorf("reading %q: %w", outPath, err)
	}
	if existingYear, ok := extractCopyrightYear(existing); ok {
		return existingYear, nil
	}
	return year, nil
}

func (g *FuzzerGenerator) fileAnnotation(msg protoreflect.MessageDescriptor) *annotations.FileAnnotation {
	attrs := map[string][]string{
		"krm.group":     {g.opts.Group},
		"krm.version":   {g.opts.Version},
		"proto.message": {string(msg.FullName())},
	}
	if g.opts.ProtoService != "" {
		attrs["proto.service"] = []string{g.opts.ProtoService}
	}
	return &annotations.FileAnnotation{Key: FuzzerFileAnnotationKey, Attributes: attrs}
}

func planFuzzer(in fuzzerPlanInput) (*FuzzerPlan, error) {
	protoType := protoNameForType(in.message)
	spec, specSide, err := in.resolveSpecSide(protoType)
	if err != nil {
		return nil, err
	}
	status, statusSide, err := in.resolveStatusSide(protoType, spec, specSide)
	if err != nil {
		return nil, err
	}

	funcName := fuzzerFuncName(in.kind)
	plan := &FuzzerPlan{
		Kind:        in.kind,
		Message:     in.message,
		PackageName: in.direct.packageName,
		FuncName:    funcName,
		Spec:        spec,
		Status:      status,
		Entries:     classifyFuzzerFields(in.message, specSide, statusSide),
	}
	if in.direct.handwrittenFunc(funcName+"FilterSpec") != "" {
		plan.FilterSpec = funcName + "FilterSpec"
	}
	if in.direct.handwrittenFunc(funcName+"FilterStatus") != "" {
		plan.FilterStatus = funcName + "FilterStatus"
	}
	return plan, nil
}

func (in fuzzerPlanInput) resolveSpecSide(protoType string) (*fuzzerMapperPair, fuzzerSide, error) {
	spec, skipped := in.direct.findMapperPair(protoType, in.version, in.kind+"Spec")
	if spec == nil {
		msg := fmt.Sprintf("no Spec mappers for kind %s in %s: want %sSpec_%s_FromProto and _ToProto, or %sSpec_FromProto and _ToProto, converting *%s",
			in.kind, in.direct.dir, in.kind, in.version, in.kind, protoType)
		if len(skipped) > 0 {
			msg += "; " + strings.Join(skipped, "; ")
		}
		return nil, fuzzerSide{}, errors.New(msg)
	}
	if spec.ProtoGoPackage == "" {
		return nil, fuzzerSide{}, fmt.Errorf("%s: cannot tell the Go package of *%s", spec.FromProto, spec.ProtoType)
	}
	specSide, err := in.krmSide(spec)
	if err != nil {
		return nil, fuzzerSide{}, err
	}
	specSide.rootMsg = true
	return spec, specSide, nil
}

func (in fuzzerPlanInput) resolveStatusSide(protoType string, spec *fuzzerMapperPair, specSide fuzzerSide) (*fuzzerMapperPair, fuzzerSide, error) {
	statusNames := statusCandidateTypeNames(in.kind, spec.KRMType, specSide)
	status, _ := in.direct.findMapperPair(protoType, in.version, statusNames...)
	var statusSide fuzzerSide
	if status != nil {
		var err error
		if statusSide, err = in.krmSide(status); err != nil {
			return nil, fuzzerSide{}, err
		}
	} else {
		statusSide = in.fallbackStatusSide(statusNames, specSide.pkg)
	}
	statusSide.rootMsg = true
	return status, statusSide, nil
}

func (in fuzzerPlanInput) fallbackStatusSide(statusNames []string, pkg *fuzzerKRMPackage) fuzzerSide {
	for _, name := range statusNames {
		if st := pkg.structs[name]; st != nil && len(st.Fields) > 0 {
			return fuzzerSide{
				coverage: fuzzerCoveredStruct,
				goStruct: unwrapObservedStateStruct(pkg, name, st),
				pkg:      pkg,
				direct:   in.direct,
				version:  in.version,
			}
		}
	}
	return fuzzerSide{}
}

func (in fuzzerPlanInput) krmSide(pair *fuzzerMapperPair) (fuzzerSide, error) {
	if pair.KRMGoPackage == "" {
		return fuzzerSide{}, fmt.Errorf("%s returns *%s from the direct package itself, not a KRM package", pair.FromProto, pair.KRMType)
	}
	pkg, err := in.krm.load(pair.KRMGoPackage)
	if err != nil {
		return fuzzerSide{}, err
	}
	st := pkg.structs[pair.KRMType]
	if st == nil {
		return fuzzerSide{}, fmt.Errorf("%s returns %s.%s, but %s declares no such struct", pair.FromProto, pair.KRMGoPackage, pair.KRMType, pkg.importPath)
	}
	return fuzzerSide{
		coverage: fuzzerCoveredStruct,
		goStruct: unwrapObservedStateStruct(pkg, pair.KRMType, st),
		pkg:      pkg,
		direct:   in.direct,
		version:  in.version,
		mapper:   in.direct.inspectMapperPair(pair),
	}, nil
}

func statusCandidateTypeNames(kind, specKRMType string, specSide fuzzerSide) []string {
	names := []string{kind + "ObservedState"}
	if st := specSide.pkg.structs[strings.TrimSuffix(specKRMType, "Spec")+"Status"]; st != nil {
		if f := findStructField(st, "ObservedState"); f != nil {
			if obs := strings.TrimPrefix(f.Type, "*"); obs != "" && obs != names[0] {
				names = append(names, obs)
			}
		}
	}
	if hasDirectStatusFields(specSide.pkg.structs[kind+"Status"]) {
		return append([]string{kind + "Status"}, names...)
	}
	return append(names, kind+"Status")
}

func unwrapObservedStateStruct(pkg *fuzzerKRMPackage, typeName string, st *gocode.GoStruct) *gocode.GoStruct {
	if !strings.HasSuffix(typeName, "Status") {
		return st
	}
	f := findStructField(st, "ObservedState")
	if f == nil {
		return st
	}
	obs := pkg.structs[strings.TrimPrefix(f.Type, "*")]
	if obs == nil || len(obs.Fields) == 0 {
		return st
	}
	if !hasDirectStatusFields(st) {
		return obs
	}
	merged := *st
	merged.Fields = append(append([]*gocode.StructField(nil), st.Fields...), obs.Fields...)
	return &merged
}

func hasDirectStatusFields(st *gocode.GoStruct) bool {
	if st == nil {
		return false
	}
	for _, f := range st.Fields {
		switch f.Name {
		case "", "CommonStatus", "Conditions", "ObservedGeneration", "ExternalRef", "ObservedState":
		default:
			return true
		}
	}
	return false
}

func renderFuzzer(plan *FuzzerPlan, annotation *annotations.FileAnnotation, year int) ([]byte, error) {
	var b bytes.Buffer
	writeCopyright(&b, year)
	b.WriteString("// Code generated by dev/tasks/generate-all. DO NOT EDIT.\n//go:build !ignore_autogenerated\n// +build !ignore_autogenerated\n\n")
	fmt.Fprintf(&b, "%s\npackage %s\n\nimport (\n\tpb %q\n\t%q\n)\n\n", annotation.FormatGo(), plan.PackageName, plan.Spec.ProtoGoPackage, fuzztestingGoPackage)

	register := "RegisterKRMFuzzer"
	if plan.Status == nil {
		register = "RegisterKRMSpecFuzzer"
	}
	fmt.Fprintf(&b, "func init() {\n\tfuzztesting.%s(%s())\n}\n\nfunc %s() fuzztesting.KRMFuzzer {\n", register, plan.FuncName, plan.FuncName)
	if plan.Status != nil {
		fmt.Fprintf(&b, "\tf := fuzztesting.NewKRMTypedFuzzer(&pb.%s{},\n\t\t%s, %s,\n\t\t%s, %s,\n\t)\n",
			plan.Spec.ProtoType, plan.Spec.FromProto, plan.Spec.ToProto, plan.Status.FromProto, plan.Status.ToProto)
	} else {
		fmt.Fprintf(&b, "\tf := fuzztesting.NewKRMTypedSpecFuzzer(&pb.%s{},\n\t\t%s, %s,\n\t)\n",
			plan.Spec.ProtoType, plan.Spec.FromProto, plan.Spec.ToProto)
	}

	for _, group := range fuzzerEntryGroups {
		if lines := formatEntryGroup(plan.Entries, group.methods); len(lines) > 0 {
			fmt.Fprintf(&b, "\n\t// %s\n%s", group.comment, strings.Join(lines, ""))
		}
	}
	if plan.FilterSpec != "" || plan.FilterStatus != "" {
		b.WriteString("\n")
	}
	if plan.FilterSpec != "" {
		fmt.Fprintf(&b, "\tf.FilterSpec = %s\n", plan.FilterSpec)
	}
	if plan.FilterStatus != "" {
		fmt.Fprintf(&b, "\tf.FilterStatus = %s\n", plan.FilterStatus)
	}
	b.WriteString("\n\treturn f\n}\n")
	return format.Source(b.Bytes())
}

func formatEntryGroup(entries []FuzzerEntry, methods []FuzzerMethod) []string {
	var lines []string
	for _, e := range entries {
		for _, m := range methods {
			if e.Method == m {
				lines = append(lines, formatFuzzerEntry(e))
			}
		}
	}
	return lines
}

func formatFuzzerEntry(e FuzzerEntry) string {
	comment := ""
	if e.Hint != "" {
		comment = " // HINT: " + e.Hint
	}
	if e.Method == FuzzerUnimplementedEtag {
		return fmt.Sprintf("\tf.Unimplemented_Etag()%s\n", comment)
	}
	return fmt.Sprintf("\tf.%s(%q)%s\n", e.Method, e.Path, comment)
}

// --- Field Classification ---

type fuzzerCoverage int

const (
	fuzzerNotCovered fuzzerCoverage = iota
	fuzzerCoveredOpaque
	fuzzerCoveredStruct
)

type fuzzerSide struct {
	coverage fuzzerCoverage
	goStruct *gocode.GoStruct
	pkg      *fuzzerKRMPackage
	direct   *goPackageIndex
	version  string
	mapper   *fuzzerMapperCoverage
	rootMsg  bool
	hint     string
}

func (s fuzzerSide) covered() bool { return s.coverage != fuzzerNotCovered }

var commonStatusGoStruct = &gocode.GoStruct{
	Name:   "Status",
	Fields: []*gocode.StructField{{Name: "Code", Type: "*int32"}, {Name: "Message", Type: "*string"}},
}

type fuzzerClass uint8

const (
	fuzzerClassSpec fuzzerClass = 1 << iota
	fuzzerClassStatus
	fuzzerClassUnimplemented
)

type fuzzerNode struct {
	path     string
	field    protoreflect.FieldDescriptor
	spec     fuzzerSide
	status   fuzzerSide
	children []*fuzzerNode
	classes  fuzzerClass
}

type fuzzerPathKey struct {
	message protoreflect.FullName
	spec    *gocode.GoStruct
	status  *gocode.GoStruct
}

func classifyFuzzerFields(msg protoreflect.MessageDescriptor, spec, status fuzzerSide) []FuzzerEntry {
	onPath := map[fuzzerPathKey]bool{{message: msg.FullName(), spec: spec.goStruct, status: status.goStruct}: true}
	var entries []FuzzerEntry
	for _, n := range buildFuzzerNodes("", msg, spec, status, 1, onPath) {
		entries = append(entries, emitFuzzerEntries(n, /* needSpec= */ true, /* needStatus= */ true, /* topLevel= */ true)...)
	}
	sort.Slice(entries, func(i, j int) bool {
		if oi, oj := fuzzerMethodOrder(entries[i].Method), fuzzerMethodOrder(entries[j].Method); oi != oj {
			return oi < oj
		}
		return entries[i].Path < entries[j].Path
	})
	return entries
}

func buildFuzzerNodes(prefix string, msg protoreflect.MessageDescriptor, spec, status fuzzerSide, depth int, onPath map[fuzzerPathKey]bool) []*fuzzerNode {
	fields := msg.Fields()
	nodes := make([]*fuzzerNode, 0, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		nodes = append(nodes, buildFuzzerNode(prefix, fields.Get(i), spec, status, depth, onPath))
	}
	return nodes
}

func buildFuzzerNode(prefix string, fd protoreflect.FieldDescriptor, spec, status fuzzerSide, depth int, onPath map[fuzzerPathKey]bool) *fuzzerNode {
	n := &fuzzerNode{
		path:   prefix + "." + string(fd.Name()),
		field:  fd,
		spec:   pairFuzzerField(fd, spec, depth, onPath),
		status: pairFuzzerField(fd, status, depth, onPath),
	}
	if depth < fuzzerMaxDepth && (n.spec.coverage == fuzzerCoveredStruct || n.status.coverage == fuzzerCoveredStruct) {
		key := fuzzerPathKey{message: fd.Message().FullName(), spec: n.spec.goStruct, status: n.status.goStruct}
		if onPath[key] {
			n.spec, n.status = fuzzerSide{}, fuzzerSide{}
		} else {
			onPath[key] = true
			childPrefix := n.path
			if fd.IsList() {
				childPrefix += "[]"
			}
			n.children = buildFuzzerNodes(childPrefix, fd.Message(), n.spec, n.status, depth+1, onPath)
			delete(onPath, key)
		}
	}
	n.classes = computeNodeClasses(n)
	return n
}

func computeNodeClasses(n *fuzzerNode) fuzzerClass {
	if len(n.children) > 0 {
		var classes fuzzerClass
		for _, child := range n.children {
			classes |= child.classes
		}
		return classes
	}
	switch {
	case n.spec.covered():
		return fuzzerClassSpec
	case n.status.covered():
		return fuzzerClassStatus
	default:
		return fuzzerClassUnimplemented
	}
}

func pairFuzzerField(fd protoreflect.FieldDescriptor, parent fuzzerSide, depth int, onPath map[fuzzerPathKey]bool) fuzzerSide {
	if parent.coverage == fuzzerNotCovered {
		return fuzzerSide{}
	}
	if parent.coverage == fuzzerCoveredOpaque {
		return fuzzerSide{coverage: fuzzerCoveredOpaque}
	}
	if parent.mapper != nil && !parent.mapper.roundTrips(fd) {
		return fuzzerSide{}
	}
	var hint string
	if parent.mapper != nil {
		hint = parent.mapper.hintFor(fd)
	}

	krmField, isRef := lookupKRMFieldForProto(fd, parent)
	if isRef {
		if fd.Kind() == protoreflect.MessageKind && fd.Message().Fields().Len() > 1 {
			return fuzzerSide{}
		}
		return fuzzerSide{coverage: fuzzerCoveredOpaque, hint: hint}
	}
	if krmField == nil {
		return fuzzerSide{}
	}

	isKRMSlice := strings.HasPrefix(krmField.Type, "[]") && krmField.Type != "[]byte"
	if fd.IsList() && !isKRMSlice {
		return fuzzerSide{}
	}
	typeName := strings.TrimLeft(krmField.Type, "[]*")
	if fd.Kind() != protoreflect.MessageKind || fd.IsMap() {
		if !fd.IsMap() {
			if st, _ := parent.pkg.lookupStruct(krmField, typeName); st != nil && findStructField(st, goFieldName(fd)) == nil {
				return fuzzerSide{}
			}
		} else if fd.MapValue().Kind() == protoreflect.MessageKind && !isOpaqueFuzzerMessage(fd.MapValue().Message()) {
			if idx := strings.Index(krmField.Type, "]"); idx >= 0 {
				valType := strings.TrimLeft(krmField.Type[idx+1:], "*")
				if st, stPkg := parent.pkg.lookupStruct(krmField, valType); st != nil {
					valKey := fuzzerPathKey{message: fd.MapValue().Message().FullName(), spec: st}
					if depth >= fuzzerMaxDepth || onPath[valKey] {
						return fuzzerSide{}
					}
					var valMapper *fuzzerMapperCoverage
					if parent.direct != nil {
						if pair := parent.direct.findStructMapperPair(protoNameForType(fd.MapValue().Message()), parent.version, st.Name); pair != nil {
							valMapper = parent.direct.inspectMapperPair(pair)
						}
					}
					valSide := fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: stPkg, direct: parent.direct, version: parent.version, mapper: valMapper}
					onPath[valKey] = true
					children := buildFuzzerNodes("", fd.MapValue().Message(), valSide, fuzzerSide{}, depth+1, onPath)
					delete(onPath, valKey)
					for _, child := range children {
						if child.classes&fuzzerClassUnimplemented != 0 {
							return fuzzerSide{}
						}
					}
				}
			}
		}
		return fuzzerSide{coverage: fuzzerCoveredOpaque, hint: hint}
	}
	if st, stPkg := parent.pkg.lookupStruct(krmField, typeName); st != nil && fd.Message().Fields().Len() > 0 {
		var childMapper *fuzzerMapperCoverage
		if parent.direct != nil {
			if pair := parent.direct.findStructMapperPair(protoNameForType(fd.Message()), parent.version, st.Name); pair != nil {
				childMapper = parent.direct.inspectMapperPair(pair)
			}
		}
		return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: stPkg, direct: parent.direct, version: parent.version, mapper: childMapper, hint: hint}
	}
	if string(fd.Message().FullName()) == "google.rpc.Status" {
		return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: commonStatusGoStruct, pkg: parent.pkg, hint: hint}
	}
	if isUnresolvedLocalKRMType(fd.Message(), typeName) {
		return fuzzerSide{}
	}
	return fuzzerSide{coverage: fuzzerCoveredOpaque, hint: hint}
}

func lookupKRMFieldForProto(fd protoreflect.FieldDescriptor, parent fuzzerSide) (field *gocode.StructField, isRef bool) {
	if parent.rootMsg && parent.mapper == nil && string(fd.Name()) == "name" {
		return nil, false
	}
	krmFieldName := goFieldName(fd)
	if f := findStructField(parent.goStruct, krmFieldName); f != nil {
		return f, isKRMReferenceField(f)
	}
	for _, suffix := range []string{"", "ID", "Id", "IDs", "Ids", "URI", "Uri", "Name"} {
		if base := strings.TrimSuffix(krmFieldName, suffix); base != "" {
			if findStructField(parent.goStruct, base+"Ref") != nil || findStructField(parent.goStruct, strings.TrimSuffix(base, "s")+"Refs") != nil {
				return nil, true
			}
		}
	}
	return nil, false
}

func isKRMReferenceField(f *gocode.StructField) bool {
	if f == nil || f.Name == "ExternalRef" {
		return false
	}
	if strings.HasSuffix(f.Name, "Ref") || strings.HasSuffix(f.Name, "Refs") {
		return true
	}
	return strings.HasSuffix(strings.TrimLeft(f.Type, "[]*"), "Ref")
}

func isUnresolvedLocalKRMType(msg protoreflect.MessageDescriptor, typeName string) bool {
	if strings.Contains(typeName, ".") {
		return false
	}
	return !isOpaqueFuzzerMessage(msg)
}

func isOpaqueFuzzerMessage(msg protoreflect.MessageDescriptor) bool {
	_, ok := protoMessagesNotMappedToGoStruct[string(msg.FullName())]
	return ok || strings.HasPrefix(string(msg.FullName()), "google.protobuf.") || msg.Fields().Len() == 0
}

func emitFuzzerEntries(n *fuzzerNode, needSpec, needStatus, topLevel bool) []FuzzerEntry {
	if !needSpec && !needStatus {
		return nil
	}
	switch n.classes {
	case fuzzerClassSpec:
		if needSpec {
			return []FuzzerEntry{{Method: FuzzerSpecField, Path: n.path, Hint: n.specHint()}}
		}
		return nil
	case fuzzerClassStatus:
		if needStatus {
			return []FuzzerEntry{{Method: FuzzerStatusField, Path: n.path, Hint: n.statusHint()}}
		}
		return nil
	case fuzzerClassUnimplemented:
		return []FuzzerEntry{unimplementedFuzzerEntry(n, topLevel)}
	}

	var entries []FuzzerEntry
	if !n.spec.covered() && needStatus {
		entries = append(entries, FuzzerEntry{Method: FuzzerStatusField, Path: n.path, Hint: n.status.hint})
		needStatus = false
	}
	if !n.status.covered() && needSpec {
		entries = append(entries, FuzzerEntry{Method: FuzzerSpecField, Path: n.path, Hint: n.spec.hint})
		needSpec = false
	}
	for _, child := range n.children {
		entries = append(entries, emitFuzzerEntries(child, needSpec, needStatus, /* topLevel= */ false)...)
	}
	return entries
}

func (n *fuzzerNode) specHint() string {
	if n.spec.hint != "" {
		return n.spec.hint
	}
	for _, child := range n.children {
		if h := child.specHint(); h != "" {
			return h
		}
	}
	return ""
}

func (n *fuzzerNode) statusHint() string {
	if n.status.hint != "" {
		return n.status.hint
	}
	for _, child := range n.children {
		if h := child.statusHint(); h != "" {
			return h
		}
	}
	return ""
}

func unimplementedFuzzerEntry(n *fuzzerNode, topLevel bool) FuzzerEntry {
	if topLevel {
		switch n.path {
		case ".name":
			return FuzzerEntry{Method: FuzzerUnimplementedIdentity, Path: n.path}
		case ".etag":
			return FuzzerEntry{Method: FuzzerUnimplementedEtag, Path: n.path}
		case ".labels", ".annotations":
			if isStringToStringMap(n.field) {
				return FuzzerEntry{Method: FuzzerUnimplementedLabelsAnnotations, Path: n.path}
			}
		}
	}
	return FuzzerEntry{Method: FuzzerUnimplementedNotYetTriaged, Path: n.path}
}

func isStringToStringMap(fd protoreflect.FieldDescriptor) bool {
	return fd.IsMap() && fd.MapKey().Kind() == protoreflect.StringKind && fd.MapValue().Kind() == protoreflect.StringKind
}

// --- General Helpers ---

func fuzzerFuncName(kind string) string {
	runes := []rune(kind)
	upper := 0
	for upper < len(runes) && unicode.IsUpper(runes[upper]) {
		upper++
	}
	if upper > 1 && upper < len(runes) && unicode.IsLower(runes[upper]) {
		upper--
	}
	for i := 0; i < upper; i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes) + "Fuzzer"
}

func fuzzerMethodOrder(m FuzzerMethod) int {
	for i, method := range []FuzzerMethod{
		FuzzerSpecField, FuzzerStatusField, FuzzerUnimplementedIdentity,
		FuzzerUnimplementedEtag, FuzzerUnimplementedLabelsAnnotations, FuzzerUnimplementedNotYetTriaged,
	} {
		if method == m {
			return i
		}
	}
	return 99
}

func findStructField(s *gocode.GoStruct, name string) *gocode.StructField {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	for _, f := range s.Fields {
		if strings.EqualFold(f.Name, name) {
			return f
		}
	}
	return nil
}
