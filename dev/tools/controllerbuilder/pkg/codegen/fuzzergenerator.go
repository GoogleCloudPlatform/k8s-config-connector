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

// This file implements the deterministic fuzzer generator (generate-fuzzer --prepopulate-resource).
//
// For one KRM kind it pairs every field of the resource's proto message with the kind's
// Spec and ObservedState structs, using the rules the mapper generator uses, and writes
// <kind>_fuzzer.generated.go with the SpecField, StatusField and Unimplemented_* entries
// that make both FuzzSpec and FuzzStatus pass.

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

const (
	fuzztestingGoPackage = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	// FuzzerFileAnnotationKey marks files written by the deterministic fuzzer generator.
	FuzzerFileAnnotationKey = "+generated:fuzzer"
)

// FuzzerGeneratorOptions configures a FuzzerGenerator.
type FuzzerGeneratorOptions struct {
	// Group and Version are the KRM API group and version, for example
	// compute.cnrm.cloud.google.com and v1beta1.
	Group   string
	Version string
	// ProtoService is the proto package of the resources. It is recorded in the file annotation.
	ProtoService string
	// APIGoPackagePath is the import path of the directory that holds the KRM packages,
	// github.com/GoogleCloudPlatform/k8s-config-connector/apis.
	APIGoPackagePath string
	// APIDirectory is the directory that holds the KRM packages, <repo>/apis.
	APIDirectory string
	// DirectDirectory is the direct controller package the fuzzers are written to,
	// <repo>/pkg/controller/direct/<service>.
	DirectDirectory string
	// Year is the copyright year of new files. Zero means the current year.
	Year int
}

// FuzzerGenerator writes <kind>_fuzzer.generated.go files into a direct controller package.
type FuzzerGenerator struct {
	opts FuzzerGeneratorOptions
	krm  *fuzzerKRMLoader
}

// FuzzerResult reports what the generator did for one kind.
type FuzzerResult struct {
	// Path is the generated file. It is written unless Skipped is set.
	Path string
	// Skipped is set when the package already has a fuzzer for the kind that isn't generated.
	// SkipReason says where it is.
	Skipped    bool
	SkipReason string
	// Plan is the generated fuzzer. It is nil when the kind was skipped.
	Plan *FuzzerPlan
}

// FuzzerPlan is the fuzzer the generator writes for one kind.
type FuzzerPlan struct {
	Kind    string
	Message protoreflect.MessageDescriptor
	// PackageName is the Go package of the direct controller.
	PackageName string
	// FuncName is the function that builds the fuzzer, <lowerCamelKind>Fuzzer.
	FuncName string
	Spec     *fuzzerMapperPair
	// Status is nil when the package has no ObservedState or Status mappers for the kind; the
	// fuzzer is then spec-only.
	Status  *fuzzerMapperPair
	Entries []FuzzerEntry
	// FilterSpec and FilterStatus are the hand-written hooks <FuncName>FilterSpec and
	// <FuncName>FilterStatus, when the package declares them.
	FilterSpec   string
	FilterStatus string

	// specSide and statusSide are the KRM structs the entries were classified against.
	specSide, statusSide fuzzerSide
}

type fuzzerPlanInput struct {
	kind string
	// version is the KRM version, used to pick <Type>_<version>_FromProto in packages whose
	// mappers are generated with --multiversion.
	version string
	message protoreflect.MessageDescriptor
	direct  *goPackageIndex
	krm     *fuzzerKRMLoader
}

func NewFuzzerGenerator(opts FuzzerGeneratorOptions) *FuzzerGenerator {
	return &FuzzerGenerator{
		opts: opts,
		krm:  newFuzzerKRMLoader(opts.APIGoPackagePath, opts.APIDirectory),
	}
}

// KRMGoPackage returns the import path of the KRM package for the configured group and
// version, derived the way generate-mapper derives it.
func (g *FuzzerGenerator) KRMGoPackage() string {
	service := strings.TrimSuffix(g.opts.Group, ".cnrm.cloud.google.com")
	return strings.TrimSuffix(g.opts.APIGoPackagePath, "/") + "/" + service + "/" + g.opts.Version
}

// FuzzerFileName returns the name of the generated fuzzer file for kind.
func FuzzerFileName(kind string) string {
	return strings.ToLower(kind) + "_fuzzer.generated.go"
}

// Generate writes the fuzzer for kind, whose resource message is msg. A fuzzer that isn't
// generated wins: when the package declares the fuzzer function in a file that isn't
// generated, or registers a fuzzer built from the kind's Spec mappers, nothing is written,
// and a stale generated file is removed.
func (g *FuzzerGenerator) Generate(kind string, msg protoreflect.MessageDescriptor) (*FuzzerResult, error) {
	direct, err := loadGoPackageIndex(g.opts.DirectDirectory)
	if err != nil {
		return nil, err
	}
	result := &FuzzerResult{Path: filepath.Join(g.opts.DirectDirectory, FuzzerFileName(kind))}

	skipReason, err := g.findHandwrittenFuzzerReason(direct, kind)
	if err != nil {
		return nil, err
	}
	if skipReason != "" {
		return g.skip(result, skipReason)
	}

	plan, err := planFuzzer(fuzzerPlanInput{
		kind:    kind,
		version: g.opts.Version,
		message: msg,
		direct:  direct,
		krm:     g.krm,
	})
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

func (g *FuzzerGenerator) findHandwrittenFuzzerReason(direct *goPackageIndex, kind string) (string, error) {
	funcName := fuzzerFuncName(kind)
	if p := direct.handwrittenFunc(funcName); p != "" {
		return fmt.Sprintf("%s declares %s", p, funcName), nil
	}
	for _, base := range mapperPairBases(g.opts.Version, kind+"Spec") {
		fromProto := base + "_FromProto"
		if direct.funcs[fromProto] == nil {
			continue
		}
		p := direct.handwrittenFuzzerUsing(fromProto)
		if p == "" {
			var err error
			if p, err = handwrittenFuzzerBelow(g.opts.DirectDirectory, fromProto); err != nil {
				return "", err
			}
		}
		if p != "" {
			return fmt.Sprintf("%s registers a fuzzer built from %s", p, fromProto), nil
		}
	}
	return "", nil
}

func (g *FuzzerGenerator) resolveCopyrightYear(filePath string) (int, error) {
	year := g.opts.Year
	if year == 0 {
		year = time.Now().Year()
	}
	existing, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return 0, fmt.Errorf("reading %q: %w", filePath, err)
	}
	if existingYear, ok := extractCopyrightYear(existing); ok {
		year = existingYear
	}
	return year, nil
}

func (g *FuzzerGenerator) skip(result *FuzzerResult, reason string) (*FuzzerResult, error) {
	result.Skipped = true
	result.SkipReason = reason
	klog.Infof("not generating %s: %s", result.Path, reason)
	if err := os.Remove(result.Path); err == nil {
		klog.Infof("removed stale %s", result.Path)
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("removing stale %q: %w", result.Path, err)
	}
	return result, nil
}

func (g *FuzzerGenerator) fileAnnotation(msg protoreflect.MessageDescriptor) *annotations.FileAnnotation {
	attributes := map[string][]string{
		"krm.group":     {g.opts.Group},
		"krm.version":   {g.opts.Version},
		"proto.message": {string(msg.FullName())},
	}
	if g.opts.ProtoService != "" {
		attributes["proto.service"] = []string{g.opts.ProtoService}
	}
	return &annotations.FileAnnotation{Key: FuzzerFileAnnotationKey, Attributes: attributes}
}

// planFuzzer finds the kind's mappers and the structs they convert to, and classifies the
// fields of the resource message.
func planFuzzer(in fuzzerPlanInput) (*FuzzerPlan, error) {
	protoType := protoNameForType(in.message)
	spec, skipped := in.direct.findMapperPair(protoType, in.version, in.kind+"Spec")
	if spec == nil {
		msg := fmt.Sprintf("no Spec mappers for kind %s in %s: want %sSpec_%s_FromProto and _ToProto, or %sSpec_FromProto and _ToProto, converting *%s",
			in.kind, in.direct.dir, in.kind, in.version, in.kind, protoType)
		if len(skipped) > 0 {
			msg += "; " + strings.Join(skipped, "; ")
		}
		return nil, errors.New(msg)
	}
	if spec.ProtoGoPackage == "" {
		return nil, fmt.Errorf("%s: cannot tell the Go package of *%s", spec.FromProto, spec.ProtoType)
	}

	specSide, err := in.krmSide(spec)
	if err != nil {
		return nil, err
	}

	statusTypeNames := statusCandidateTypeNames(in.kind, spec.KRMType, specSide)
	status, _ := in.direct.findMapperPair(protoType, in.version, statusTypeNames...)

	statusSide, err := in.resolveStatusSide(status, statusTypeNames, specSide)
	if err != nil {
		return nil, err
	}

	plan := &FuzzerPlan{
		Kind:        in.kind,
		Message:     in.message,
		PackageName: in.direct.packageName,
		FuncName:    fuzzerFuncName(in.kind),
		Spec:        spec,
		Status:      status,
		Entries:     classifyFuzzerFields(in.message, specSide, statusSide),
		specSide:    specSide,
		statusSide:  statusSide,
	}
	if in.direct.handwrittenFunc(plan.FuncName+"FilterSpec") != "" {
		plan.FilterSpec = plan.FuncName + "FilterSpec"
	}
	if in.direct.handwrittenFunc(plan.FuncName+"FilterStatus") != "" {
		plan.FilterStatus = plan.FuncName + "FilterStatus"
	}
	return plan, nil
}

func statusCandidateTypeNames(kind, specKRMType string, specSide fuzzerSide) []string {
	statusTypeNames := []string{kind + "ObservedState"}
	if st := specSide.pkg.structs[strings.TrimSuffix(specKRMType, "Spec")+"Status"]; st != nil {
		if f := findStructField(st, "ObservedState"); f != nil {
			if name := strings.TrimPrefix(f.Type, "*"); name != "" && name != statusTypeNames[0] {
				statusTypeNames = append(statusTypeNames, name)
			}
		}
	}
	return append(statusTypeNames, kind+"Status")
}

func (in fuzzerPlanInput) resolveStatusSide(status *fuzzerMapperPair, statusTypeNames []string, specSide fuzzerSide) (fuzzerSide, error) {
	if status != nil {
		return in.krmSide(status)
	}
	// Without ObservedState mappers the fuzzer is spec-only, and FuzzStatus never runs.
	// The ObservedState struct, or the Status struct, still tells which fields are status fields.
	for _, name := range statusTypeNames {
		if st := specSide.pkg.structs[name]; st != nil {
			return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: specSide.pkg}, nil
		}
	}
	return fuzzerSide{}, nil
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
	return fuzzerSide{coverage: fuzzerCoveredStruct, goStruct: st, pkg: pkg}, nil
}

// fuzzerFuncName returns <lowerCamelKind>Fuzzer. A leading acronym is lowercased as a whole,
// so ComputeTargetVPNGateway gives computeTargetVPNGatewayFuzzer and IAMDenyPolicy gives
// iamDenyPolicyFuzzer.
func fuzzerFuncName(kind string) string {
	runes := []rune(kind)
	upper := 0
	for upper < len(runes) && unicode.IsUpper(runes[upper]) {
		upper++
	}
	lower := upper
	if upper > 1 && upper < len(runes) && unicode.IsLower(runes[upper]) {
		// The last capital starts the next word, as in IAMDeny.
		lower = upper - 1
	}
	for i := 0; i < lower; i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes) + "Fuzzer"
}

// renderFuzzer renders the generated fuzzer file, with the header conventions of mapper.generated.go.
func renderFuzzer(plan *FuzzerPlan, annotation *annotations.FileAnnotation, year int) ([]byte, error) {
	var b bytes.Buffer
	writeFuzzerHeader(&b, plan, annotation, year)
	writeFuzzerInitAndConstructor(&b, plan)
	writeFuzzerEntryGroups(&b, plan.Entries)
	writeFuzzerFilters(&b, plan)
	b.WriteString("\n\treturn f\n}\n")

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("formatting the fuzzer for %s: %w\n%s", plan.Kind, err, b.String())
	}
	return formatted, nil
}

func writeFuzzerHeader(b *bytes.Buffer, plan *FuzzerPlan, annotation *annotations.FileAnnotation, year int) {
	writeCopyright(b, year)
	b.WriteString("// Code generated by dev/tasks/generate-all. DO NOT EDIT.\n")
	b.WriteString("//go:build !ignore_autogenerated\n")
	b.WriteString("// +build !ignore_autogenerated\n\n")
	fmt.Fprintf(b, "%s\n", annotation.FormatGo())
	fmt.Fprintf(b, "package %s\n\n", plan.PackageName)
	fmt.Fprintf(b, "import (\n\tpb %q\n\t%q\n)\n\n", plan.Spec.ProtoGoPackage, fuzztestingGoPackage)
}

func writeFuzzerInitAndConstructor(b *bytes.Buffer, plan *FuzzerPlan) {
	register := "RegisterKRMFuzzer"
	if plan.Status == nil {
		register = "RegisterKRMSpecFuzzer"
	}
	fmt.Fprintf(b, "func init() {\n\tfuzztesting.%s(%s())\n}\n\n", register, plan.FuncName)

	fmt.Fprintf(b, "func %s() fuzztesting.KRMFuzzer {\n", plan.FuncName)
	if plan.Status != nil {
		fmt.Fprintf(b, "\tf := fuzztesting.NewKRMTypedFuzzer(&pb.%s{},\n", plan.Spec.ProtoType)
		fmt.Fprintf(b, "\t\t%s, %s,\n", plan.Spec.FromProto, plan.Spec.ToProto)
		fmt.Fprintf(b, "\t\t%s, %s,\n", plan.Status.FromProto, plan.Status.ToProto)
	} else {
		fmt.Fprintf(b, "\tf := fuzztesting.NewKRMTypedSpecFuzzer(&pb.%s{},\n", plan.Spec.ProtoType)
		fmt.Fprintf(b, "\t\t%s, %s,\n", plan.Spec.FromProto, plan.Spec.ToProto)
	}
	b.WriteString("\t)\n")
}

func writeFuzzerEntryGroups(b *bytes.Buffer, entries []FuzzerEntry) {
	for _, group := range fuzzerEntryGroups {
		var lines []string
		for _, e := range entries {
			for _, m := range group.methods {
				if e.Method != m {
					continue
				}
				if m == FuzzerUnimplementedEtag {
					lines = append(lines, "\tf.Unimplemented_Etag()\n")
				} else {
					lines = append(lines, fmt.Sprintf("\tf.%s(%q)\n", e.Method, e.Path))
				}
			}
		}
		if len(lines) == 0 {
			continue
		}
		fmt.Fprintf(b, "\n\t// %s\n", group.comment)
		for _, line := range lines {
			b.WriteString(line)
		}
	}
}

func writeFuzzerFilters(b *bytes.Buffer, plan *FuzzerPlan) {
	if plan.FilterSpec != "" || plan.FilterStatus != "" {
		b.WriteString("\n")
	}
	if plan.FilterSpec != "" {
		fmt.Fprintf(b, "\tf.FilterSpec = %s\n", plan.FilterSpec)
	}
	if plan.FilterStatus != "" {
		fmt.Fprintf(b, "\tf.FilterStatus = %s\n", plan.FilterStatus)
	}
}
