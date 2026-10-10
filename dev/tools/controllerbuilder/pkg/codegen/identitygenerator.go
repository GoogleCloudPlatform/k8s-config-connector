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
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	// IdentityFileAnnotationKey marks a deterministically generated _identity.generated.go file.
	IdentityFileAnnotationKey = "+generated:identity"
	// ReferenceFileAnnotationKey marks a deterministically generated _reference.generated.go file.
	ReferenceFileAnnotationKey = "+generated:reference"
)

// IdentityGeneratorOptions configures deterministic identity and reference generation
// for a single KRM API GroupVersion.
type IdentityGeneratorOptions struct {
	Group            string
	Version          string
	ProtoService     string
	APIDirectory     string
	Proto            *protoapi.Proto
	CAIMetadataPath  string
	EmitTests        bool
	PatternOverrides map[string]string
}

// IdentityGenerator deterministically plans and emits <kind>_identity.generated.go,
// <kind>_reference.generated.go, and <kind>_identity_generated_test.go files.
type IdentityGenerator struct {
	opts IdentityGeneratorOptions
}

// IdentitySegment represents one "/"-delimited segment in a GCP resource name pattern.
type IdentitySegment struct {
	Literal    string
	IsVariable bool
	Field      IdentityField
}

// IdentityField describes one variable segment in a GCP resource name pattern.
type IdentityField struct {
	RawPlaceholder string
	TemplateVar    string
	StructField    string
	Collection     string
	DocPlaceholder string
	SampleValue    string
}

// AncestorSourceKind identifies how a non-leaf variable segment is resolved from <Kind>Spec.
type AncestorSourceKind string

const (
	AncestorProjectRef      AncestorSourceKind = "projectRef"
	AncestorOrganizationRef AncestorSourceKind = "organizationRef"
	AncestorFolderRef       AncestorSourceKind = "folderRef"
	AncestorLocation        AncestorSourceKind = "location"
	AncestorSpecString      AncestorSourceKind = "specString"
	AncestorParentRef       AncestorSourceKind = "parentRef"
	AncestorUnknown         AncestorSourceKind = "unknown"
)

// AncestorResolution records how a single ancestor variable segment is resolved.
type AncestorResolution struct {
	Field         IdentityField
	Source        AncestorSourceKind
	SpecFieldName string
	SpecJSONName  string
	IsPointer     bool
}

// ParentRefResolution records when a single parent reference field on <Kind>Spec
// (such as SpannerInstanceRef or BigtableTableRef) resolves a multi-segment ancestor prefix.
type ParentRefResolution struct {
	SpecFieldName string
	SpecJSONName  string
	IsPointer     bool
	CoveredFields []IdentityField
	TailSegments  []IdentitySegment
}

// IdentityPlan is the deterministic plan for generating a resource's identity,
// reference, and conformance test files.
type IdentityPlan struct {
	Kind              string
	Group             string
	Version           string
	PackageName       string
	PackageDir        string
	ProtoMessage      string
	ProtoService      string
	Host              string
	RawPattern        string
	TemplatePattern   string
	DocFormat         string
	Segments          []IdentitySegment
	Fields            []IdentityField
	ParentSegments    []IdentitySegment
	ResourceIDField   *IdentityField
	ServerGeneratedID bool
	SpecIDIsPointer   bool
	StatusField       string
	StatusFieldIsPtr  bool
	Ancestors              []AncestorResolution
	HasUnresolvedAncestors bool
	ParentRef         *ParentRefResolution
	Judgement         []judgement.Entry
}

// IdentityGenerateResult summarizes the files written or skipped for one resource.
type IdentityGenerateResult struct {
	Plan             *IdentityPlan
	IdentityPath     string
	ReferencePath    string
	TestPath         string
	IdentitySkipped  bool
	ReferenceSkipped bool
	TestSkipped      bool
}

// HandWrittenOverrides records which declarations already exist in hand-written
// (non-generated) Go files in the KRM package directory.
type HandWrittenOverrides struct {
	HasIdentityFormat          bool
	HasIdentityStruct          bool
	IdentityMethods            map[string]bool
	HasGetIdentityFromSpec     bool
	ResourceMethods            map[string]bool
	HasGVKVar                  bool
	HasRefStruct               bool
	HasRefRegister             bool
	RefMethods                 map[string]bool
	HasHandWrittenIdentityTest bool
}

// NewIdentityGenerator creates a new IdentityGenerator.
func NewIdentityGenerator(opts IdentityGeneratorOptions) *IdentityGenerator {
	return &IdentityGenerator{opts: opts}
}

// Generate writes <kind>_identity.generated.go, <kind>_reference.generated.go,
// and (when EmitTests is true) <kind>_identity_generated_test.go into the KRM
// API package directory, respecting hand-written overrides function by function.
func (g *IdentityGenerator) Generate(kind string, msg protoreflect.MessageDescriptor) (*IdentityGenerateResult, error) {
	plan, err := g.Plan(kind, msg)
	if err != nil {
		return nil, err
	}
	typeInfo, err := inspectKRMPackage(plan.PackageDir, kind, g.opts.Version)
	if err != nil {
		return nil, err
	}

	res := &IdentityGenerateResult{
		Plan:          plan,
		IdentityPath:  filepath.Join(plan.PackageDir, strings.ToLower(kind)+"_identity.generated.go"),
		ReferencePath: filepath.Join(plan.PackageDir, strings.ToLower(kind)+"_reference.generated.go"),
		TestPath:      filepath.Join(plan.PackageDir, strings.ToLower(kind)+"_identity_generated_test.go"),
	}
	if plan.RawPattern == "" {
		res.IdentitySkipped = true
		res.ReferenceSkipped = true
		res.TestSkipped = true
		return res, nil
	}

	if err := g.writeIdentityArtifact(res, plan, &typeInfo.Overrides); err != nil {
		return nil, err
	}
	if err := g.writeReferenceArtifact(res, plan, &typeInfo.Overrides); err != nil {
		return nil, err
	}
	if err := g.writeTestArtifact(res, plan, &typeInfo.Overrides); err != nil {
		return nil, err
	}
	return res, nil
}

// Plan builds the deterministic IdentityPlan for kind and msg without writing any files.
func (g *IdentityGenerator) Plan(kind string, msg protoreflect.MessageDescriptor) (*IdentityPlan, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil proto MessageDescriptor for kind %s", kind)
	}
	if strings.HasPrefix(string(msg.FullName()), "mockgcp.") {
		return nil, fmt.Errorf("generate-identity does not support mockgcp proto %s; use a googleapis proto service", msg.FullName())
	}

	pkgDir := g.packageDir()
	typeInfo, err := inspectKRMPackage(pkgDir, kind, g.opts.Version)
	if err != nil {
		return nil, fmt.Errorf("inspecting KRM package %s: %w", pkgDir, err)
	}

	rd := getResourceDescriptorProto(msg)
	defaultHost, serverGenID := g.inspectProtoService(msg)
	plan := &IdentityPlan{
		Kind:             kind,
		Group:            g.opts.Group,
		Version:          g.opts.Version,
		PackageName:      typeInfo.PackageName,
		PackageDir:       pkgDir,
		ProtoMessage:     string(msg.FullName()),
		ProtoService:     g.opts.ProtoService,
		SpecIDIsPointer:  typeInfo.ResourceIDIsPtr,
		StatusField:      typeInfo.StatusField,
		StatusFieldIsPtr: typeInfo.StatusFieldIsPtr,
	}

	patterns, hostOverride, patOverride := g.resolveCandidatePatterns(kind, msg)
	if len(patterns) == 0 {
		plan.Judgement = append(plan.Judgement, judgement.Entry{
			Kind:   kind,
			Group:  g.opts.Group,
			Reason: judgement.ReasonIdentityRootUnknown,
			Detail: fmt.Sprintf("proto message %s declares no google.api.resource pattern; pass --pattern %s=<pattern> to generate-identity", msg.FullName(), kind),
			Status: judgement.StatusOpen,
		})
		return plan, nil
	}

	rawPattern, unique := selectBestPattern(patterns, msg, typeInfo)
	if !unique {
		plan.Judgement = append(plan.Judgement, judgement.Entry{
			Kind:   kind,
			Group:  g.opts.Group,
			Reason: judgement.ReasonIdentityMultiPattern,
			Detail: fmt.Sprintf("proto message %s declares %d resource name patterns (%s) that cannot be uniquely disambiguated from %sSpec; pass --pattern %s=<pattern> to generate-identity",
				msg.FullName(), len(patterns), strings.Join(patterns, ", "), kind, kind),
			Status: judgement.StatusOpen,
		})
		return plan, nil
	}

	plan.RawPattern = rawPattern
	plan.Host = g.resolveHost(rd, defaultHost, hostOverride)
	if patOverride == "" && len(patterns) > 1 {
		plan.Judgement = append(plan.Judgement, judgement.Entry{
			Kind:   plan.Kind,
			Group:  g.opts.Group,
			Reason: judgement.ReasonIdentityMultiPattern,
			Detail: fmt.Sprintf("proto message %s declares %d resource name patterns (%s); generated identity uses %q (pass --pattern %s=<pattern> to override)",
				msg.FullName(), len(patterns), strings.Join(patterns, ", "), plan.RawPattern, plan.Kind),
			Status: judgement.StatusOpen,
		})
	}

	populatePatternSegments(plan, rawPattern)
	populateParentSegments(plan)
	plan.ServerGeneratedID = (serverGenID || typeInfo.ResourceIDServer) && plan.ResourceIDField != nil
	g.resolveSpecAncestors(plan, typeInfo)
	return plan, nil
}

// SortIdentityJudgement orders judgement entries deterministically by key.
func SortIdentityJudgement(entries []judgement.Entry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Key() < entries[j].Key()
	})
}

func (g *IdentityGenerator) writeIdentityArtifact(res *IdentityGenerateResult, plan *IdentityPlan, ov *HandWrittenOverrides) error {
	idBytes, emitted, err := g.RenderIdentityFile(plan, ov)
	if err != nil {
		return fmt.Errorf("rendering identity for %s: %w", plan.Kind, err)
	}
	if !emitted {
		res.IdentitySkipped = true
		return nil
	}
	return os.WriteFile(res.IdentityPath, idBytes, 0644)
}

func (g *IdentityGenerator) writeReferenceArtifact(res *IdentityGenerateResult, plan *IdentityPlan, ov *HandWrittenOverrides) error {
	refBytes, emitted, err := g.RenderReferenceFile(plan, ov)
	if err != nil {
		return fmt.Errorf("rendering reference for %s: %w", plan.Kind, err)
	}
	if !emitted {
		res.ReferenceSkipped = true
		return nil
	}
	return os.WriteFile(res.ReferencePath, refBytes, 0644)
}

func (g *IdentityGenerator) writeTestArtifact(res *IdentityGenerateResult, plan *IdentityPlan, ov *HandWrittenOverrides) error {
	if !g.opts.EmitTests || ov.HasHandWrittenIdentityTest || len(plan.Fields) == 0 {
		res.TestSkipped = true
		return nil
	}
	testBytes, err := g.RenderIdentityTestFile(plan)
	if err != nil {
		return fmt.Errorf("rendering identity test for %s: %w", plan.Kind, err)
	}
	return os.WriteFile(res.TestPath, testBytes, 0644)
}

func (g *IdentityGenerator) packageDir() string {
	groupPrefix := strings.TrimSuffix(g.opts.Group, ".cnrm.cloud.google.com")
	parts := append([]string{g.opts.APIDirectory}, strings.Split(groupPrefix, ".")...)
	parts = append(parts, g.opts.Version)
	return filepath.Join(parts...)
}

func (g *IdentityGenerator) resolveCandidatePatterns(kind string, msg protoreflect.MessageDescriptor) ([]string, string, string) {
	if rawOv, ok := g.opts.PatternOverrides[kind]; ok && strings.TrimSpace(rawOv) != "" {
		hostOv, patOv := parsePatternOverride(rawOv)
		return []string{patOv}, hostOv, patOv
	}
	var patterns []string
	if md := protoapi.GetResourceMetadata(msg); md != nil {
		patterns = appendUniquePatterns(patterns, md.Patterns)
	}
	return patterns, "", ""
}

func appendUniquePatterns(dst, src []string) []string {
	for _, candidate := range src {
		normCandidate := tmplVarNormalizeRe.ReplaceAllString(candidate, "{}")
		if !containsNormalizedPattern(dst, normCandidate) {
			dst = append(dst, candidate)
		}
	}
	return dst
}

func containsNormalizedPattern(patterns []string, normTarget string) bool {
	for _, existing := range patterns {
		if tmplVarNormalizeRe.ReplaceAllString(existing, "{}") == normTarget {
			return true
		}
	}
	return false
}

func parsePatternOverride(raw string) (host, pattern string) {
	raw = strings.TrimSpace(raw)
	if !strings.HasPrefix(raw, "//") {
		return "", strings.TrimPrefix(raw, "/")
	}
	rest := strings.TrimPrefix(raw, "//")
	if slash := strings.IndexByte(rest, '/'); slash > 0 {
		return rest[:slash], rest[slash+1:]
	}
	return "", strings.TrimPrefix(raw, "/")
}

func (g *IdentityGenerator) resolveHost(rd *annotations.ResourceDescriptor, defaultHost, hostOverride string) string {
	if hostOverride != "" {
		return hostOverride
	}
	if rd != nil && rd.GetType() != "" {
		if svc, _ := protoapi.ParseResourceTarget(rd.GetType()); svc != "" {
			return svc + ".googleapis.com"
		}
	}
	if defaultHost != "" {
		return defaultHost
	}
	groupPrefix := strings.Split(g.opts.Group, ".")[0]
	return groupPrefix + ".googleapis.com"
}

var tmplVarNormalizeRe = regexp.MustCompile(`\{[^}]+\}`)

func selectBestPattern(patterns []string, msg protoreflect.MessageDescriptor, info *krmTypeInfo) (string, bool) {
	if len(patterns) == 1 {
		return patterns[0], true
	}
	best := patterns[0]
	bestScore := scorePatternAgainstSpec(patterns[0], msg, info)
	tiedCount := 1
	for _, p := range patterns[1:] {
		sc := scorePatternAgainstSpec(p, msg, info)
		if sc > bestScore {
			best = p
			bestScore = sc
			tiedCount = 1
		} else if sc == bestScore {
			tiedCount++
		}
	}
	if tiedCount > 1 || bestScore <= 0 {
		return "", false
	}
	return best, true
}

func scorePatternAgainstSpec(pattern string, msg protoreflect.MessageDescriptor, info *krmTypeInfo) int {
	if info == nil || !strings.Contains(pattern, "{") {
		return -10
	}
	segs := strings.Split(pattern, "/")
	score := scoreRootSegment(segs, info)
	score += scoreLocationSegments(segs, info)
	score += scoreIntermediateParentSegments(segs, info)
	score += scoreAlternatingSegmentsAndLeaf(segs, msg)
	return score
}

func scoreRootSegment(segs []string, info *krmTypeInfo) int {
	if len(segs) < 2 {
		return -10
	}
	switch segs[0] {
	case "projects":
		if info.HasProjectRef {
			return 4
		}
		if !info.HasOrgRef && !info.HasFolderRef {
			return 2
		}
		return -5
	case "organizations":
		if info.HasOrgRef {
			return 4
		}
		return -10
	case "folders":
		if info.HasFolderRef {
			return 4
		}
		return -10
	default:
		if strings.HasPrefix(segs[1], "{") {
			ph := strings.Trim(segs[1], "{}")
			af := IdentityField{RawPlaceholder: ph, StructField: snakeToPascal(ph), Collection: segs[0]}
			if matchParentRefField(af, info.SpecFields) != nil || matchPlainSpecField(af, info.SpecFields) != nil {
				return 4
			}
		}
		return -10
	}
}

func scoreLocationSegments(segs []string, info *krmTypeInfo) int {
	hasLocInPattern := false
	bonus := 0
	for i := 0; i+2 < len(segs); i++ {
		if (segs[i] != "locations" && segs[i] != "regions" && segs[i] != "zones") || !strings.HasPrefix(segs[i+1], "{") {
			continue
		}
		hasLocInPattern = true
		if info.HasLocation && segs[i] == "regions" && strings.EqualFold(info.LocationJSONName, "region") {
			bonus = 2
		} else if info.HasLocation && segs[i] == "locations" && strings.EqualFold(info.LocationJSONName, "location") {
			bonus = 1
		}
		break
	}
	if hasLocInPattern == info.HasLocation {
		return bonus + 3
	}
	if hasLocInPattern && !info.HasLocation {
		return bonus - 3
	}
	return bonus
}

func scoreIntermediateParentSegments(segs []string, info *krmTypeInfo) int {
	score := 0
	for i := 0; i+1 < len(segs)-2; i += 2 {
		col := segs[i]
		if isStandardRootOrLocationCollection(col) {
			continue
		}
		ph := strings.Trim(segs[i+1], "{}")
		af := IdentityField{
			RawPlaceholder: ph,
			StructField:    snakeToPascal(ph),
			Collection:     col,
		}
		if matchParentRefField(af, info.SpecFields) != nil || matchPlainSpecField(af, info.SpecFields) != nil {
			score += 3
		} else {
			score -= 2
		}
	}
	return score
}

func scoreAlternatingSegmentsAndLeaf(segs []string, msg protoreflect.MessageDescriptor) int {
	score := 0
	for i := 0; i+1 < len(segs); i++ {
		if !strings.HasPrefix(segs[i], "{") && !strings.HasPrefix(segs[i+1], "{") {
			score -= 2
		}
	}
	if msg != nil && len(segs) >= 2 && strings.HasPrefix(segs[len(segs)-1], "{") {
		leafVar := strings.ReplaceAll(strings.ToLower(strings.Trim(segs[len(segs)-1], "{}")), "_", "")
		if leafVar == strings.ToLower(string(msg.Name())) {
			score += 1
		}
	}
	return score
}

func isStandardRootOrLocationCollection(col string) bool {
	switch col {
	case "projects", "organizations", "folders", "locations", "regions", "zones":
		return true
	default:
		return false
	}
}

func populatePatternSegments(plan *IdentityPlan, rawPattern string) {
	usedStructFields := map[string]int{}
	var tmplSegs []string
	var docSegs []string
	lastLiteral := ""

	for _, seg := range strings.Split(rawPattern, "/") {
		if !strings.HasPrefix(seg, "{") || !strings.HasSuffix(seg, "}") {
			lastLiteral = seg
			plan.Segments = append(plan.Segments, IdentitySegment{Literal: seg})
			tmplSegs = append(tmplSegs, seg)
			docSegs = append(docSegs, seg)
			continue
		}
		f := buildIdentityField(seg, lastLiteral, usedStructFields)
		plan.Segments = append(plan.Segments, IdentitySegment{IsVariable: true, Field: f})
		plan.Fields = append(plan.Fields, f)
		tmplSegs = append(tmplSegs, "{"+f.TemplateVar+"}")
		docSegs = append(docSegs, f.DocPlaceholder)
	}
	plan.TemplatePattern = strings.Join(tmplSegs, "/")
	plan.DocFormat = strings.Join(docSegs, "/")
}

func buildIdentityField(seg, lastLiteral string, usedStructFields map[string]int) IdentityField {
	rawVar := strings.TrimSpace(seg[1 : len(seg)-1])
	if idx := strings.Index(rawVar, "="); idx != -1 {
		rawVar = rawVar[:idx]
	}
	structField := snakeToPascal(rawVar)
	tmplVar := snakeToLowerCamel(rawVar)
	if count := usedStructFields[strings.ToLower(structField)]; count > 0 {
		suffix := fmt.Sprintf("%d", count+1)
		structField += suffix
		tmplVar += suffix
	}
	usedStructFields[strings.ToLower(structField)]++
	return IdentityField{
		RawPlaceholder: rawVar,
		TemplateVar:    tmplVar,
		StructField:    structField,
		Collection:     lastLiteral,
		DocPlaceholder: docPlaceholderFor(rawVar, tmplVar),
		SampleValue:    sampleValueFor(rawVar),
	}
}

func populateParentSegments(plan *IdentityPlan) {
	if len(plan.Segments) == 0 {
		return
	}
	if plan.Segments[len(plan.Segments)-1].IsVariable {
		idField := plan.Segments[len(plan.Segments)-1].Field
		plan.ResourceIDField = &idField
		if len(plan.Segments) > 2 {
			plan.ParentSegments = plan.Segments[:len(plan.Segments)-2]
		}
		return
	}
	if len(plan.Segments) > 1 {
		plan.ParentSegments = plan.Segments[:len(plan.Segments)-1]
	}
}

func snakeToPascal(s string) string {
	if s == "" {
		return ""
	}
	if !strings.Contains(s, "_") && !strings.Contains(s, "-") {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})
	var b strings.Builder
	for _, p := range parts {
		if p != "" {
			b.WriteString(strings.ToUpper(p[:1]) + strings.ToLower(p[1:]))
		}
	}
	return b.String()
}

func snakeToLowerCamel(s string) string {
	return lowerFirst(snakeToPascal(s))
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func docPlaceholderFor(rawVar, tmplVar string) string {
	switch rawVar {
	case "project", "project_id":
		return "{{projectID}}"
	case "location", "region", "zone":
		return "{{" + rawVar + "}}"
	default:
		if strings.HasSuffix(tmplVar, "ID") || strings.HasSuffix(tmplVar, "Id") {
			return "{{" + tmplVar + "}}"
		}
		return "{{" + tmplVar + "ID}}"
	}
}

func sampleValueFor(rawVar string) string {
	switch rawVar {
	case "project", "project_id":
		return "my-project"
	case "location", "region":
		return "us-central1"
	case "zone":
		return "us-central1-a"
	case "organization", "organization_id", "folder", "folder_id":
		return "123456789"
	case "billing_account":
		return "012345-6789AB-CDEF01"
	default:
		return "my-" + strings.ReplaceAll(strings.ToLower(rawVar), "_", "-")
	}
}
