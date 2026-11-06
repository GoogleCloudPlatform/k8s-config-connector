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

package fuzztesting

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

type FuzzFn func(t *testing.T, seed int64)

var (
	fuzzers                  []FuzzFn
	registeredFuzzers        []any
	registeredNoProtoFuzzers []any
)

func RegisterKRMFuzzer(fuzzer KRMFuzzer) {
	RegisterFuzzer(fuzzer.FuzzSpec)
	RegisterFuzzer(fuzzer.FuzzStatus)
	registeredFuzzers = append(registeredFuzzers, fuzzer)
}

func RegisterKRMSpecFuzzer(fuzzer KRMFuzzer) {
	RegisterFuzzer(fuzzer.FuzzSpec)
	registeredFuzzers = append(registeredFuzzers, fuzzer)
}

func RegisterFuzzer(fuzzer FuzzFn) {
	fuzzers = append(fuzzers, fuzzer)
}

func ChooseFuzzer(n int64) FuzzFn {
	all := fuzzers
	totalShards := os.Getenv("FUZZ_TOTAL_SHARDS")
	shardIdxStr := os.Getenv("FUZZ_SHARD_INDEX")
	if totalShards != "" && shardIdxStr != "" {
		total, _ := strconv.Atoi(totalShards)
		shardIdx, _ := strconv.Atoi(shardIdxStr)
		if total > 1 && shardIdx >= 0 {
			var sharded []FuzzFn
			for i, f := range all {
				if i%total == shardIdx {
					sharded = append(sharded, f)
				}
			}
			if len(sharded) > 0 {
				all = sharded
			}
		}
	}
	return all[n%int64(len(all))]
}

func GetRegisteredFuzzers() []any {
	var all []any
	all = append(all, registeredFuzzers...)
	all = append(all, registeredNoProtoFuzzers...)
	return all
}

type KRMTypedFuzzer[ProtoT proto.Message, SpecType any, StatusType any] struct {
	ProtoType ProtoT

	SpecFromProto func(ctx *direct.MapContext, in ProtoT) *SpecType
	SpecToProto   func(ctx *direct.MapContext, in *SpecType) ProtoT

	StatusFromProto func(ctx *direct.MapContext, in ProtoT) *StatusType
	StatusToProto   func(ctx *direct.MapContext, in *StatusType) ProtoT

	UnimplementedFields sets.Set[string]
	SpecFields          sets.Set[string]
	StatusFields        sets.Set[string]

	FilterSpec   func(in ProtoT)
	FilterStatus func(in ProtoT)
}

// SpecField marks the specified fieldPath as round-tripping to/from the Spec
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) SpecField(fieldPath string) {
	f.SpecFields.Insert(fieldPath)
}

// StatusField marks the specified fieldPath as round-tripping to/from the Status
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) StatusField(fieldPath string) {
	f.StatusFields.Insert(fieldPath)
}

// Unimplemented_Internal marks the specified fieldPath as not round-tripped,
// and should be used for fields that are considered internal implementation details of the service
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Unimplemented_Internal(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

// Unimplemented_Identity marks the specified fieldPath as not round-tripped,
// and should be used for fields that are considered identity (URL) rather than being part of the object itself.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Unimplemented_Identity(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

// Unimplemented_LabelsAnnotations marks the specified fieldPath as not round-tripped,
// and should be used for fields that are either labels or annotations
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Unimplemented_LabelsAnnotations(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

// Unimplemented_Etag marks the 'etag' field as not round-tripped.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Unimplemented_Etag() {
	f.UnimplementedFields.Insert(".etag")
}

// Ignore_JSONBookkeeping marks the specified fieldPath as fields we expect to be missing, such as ServerResponse or ForceSendFields.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Ignore_JSONBookkeeping(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

// Unimplemented_NotYetTriaged marks the specified fieldPath as not round-tripped,
// and should be used for fields that are added by the service and where we haven't decided whether or not to implement them.
// This should be the "starting point" for new fields added by services.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) Unimplemented_NotYetTriaged(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

// IdentityField marks a field as not supported in the mapper, because it is part of the identity (URL) rather than being part of the object itself.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) IdentityField(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

type KRMFuzzer interface {
	FuzzSpec(t *testing.T, seed int64)
	FuzzStatus(t *testing.T, seed int64)
}

func NewKRMTypedFuzzer[ProtoT proto.Message, SpecType any, StatusType any](
	protoType ProtoT,
	specFromProto func(ctx *direct.MapContext, in ProtoT) *SpecType, specToProto func(ctx *direct.MapContext, in *SpecType) ProtoT,
	statusFromProto func(ctx *direct.MapContext, in ProtoT) *StatusType, statusToProto func(ctx *direct.MapContext, in *StatusType) ProtoT,
) *KRMTypedFuzzer[ProtoT, SpecType, StatusType] {
	return &KRMTypedFuzzer[ProtoT, SpecType, StatusType]{
		ProtoType:           protoType,
		SpecFromProto:       specFromProto,
		SpecToProto:         specToProto,
		StatusFromProto:     statusFromProto,
		StatusToProto:       statusToProto,
		UnimplementedFields: sets.New[string](),
		SpecFields:          sets.New[string](),
		StatusFields:        sets.New[string](),
	}
}

func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) FuzzSpec(t *testing.T, seed int64) {
	fuzzer := NewFuzzTest(f.ProtoType, f.SpecFromProto, f.SpecToProto)
	fuzzer.TargetFields = f.SpecFields
	fuzzer.IgnoreFields = f.StatusFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterSpec
	fuzzer.Fuzz(t, seed)
}

func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) FuzzStatus(t *testing.T, seed int64) {
	if f.StatusFromProto == nil || f.StatusToProto == nil {
		t.Skip("Status FromProto or ToProto is nil, skipping status fuzz test")
		return
	}
	fuzzer := NewFuzzTest(f.ProtoType, f.StatusFromProto, f.StatusToProto)
	fuzzer.TargetFields = f.StatusFields
	fuzzer.IgnoreFields = f.SpecFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterStatus
	fuzzer.Fuzz(t, seed)
}

type NoStatus struct{}

// NewKRMTypedSpecFuzzer is a convenience function for creating a fuzzer that only
// fuzzes the spec fields of a KRM type.
func NewKRMTypedSpecFuzzer[ProtoT proto.Message, SpecType any](
	protoType ProtoT,
	specFromProto func(ctx *direct.MapContext, in ProtoT) *SpecType,
	specToProto func(ctx *direct.MapContext, in *SpecType) ProtoT,
) *KRMTypedFuzzer[ProtoT, SpecType, NoStatus] {
	return &KRMTypedFuzzer[ProtoT, SpecType, NoStatus]{
		ProtoType:           protoType,
		SpecFromProto:       specFromProto,
		SpecToProto:         specToProto,
		StatusFromProto:     nil, // No status functions
		StatusToProto:       nil, // No status functions
		UnimplementedFields: sets.New[string](),
		SpecFields:          sets.New[string](),
		StatusFields:        sets.New[string](),
	}
}

type FuzzTest[ProtoT proto.Message, KRMType any] struct {
	ProtoType ProtoT

	FromProto func(ctx *direct.MapContext, in ProtoT) *KRMType
	ToProto   func(ctx *direct.MapContext, in *KRMType) ProtoT

	TargetFields        sets.Set[string]
	UnimplementedFields sets.Set[string]
	IgnoreFields        sets.Set[string]

	Filter func(in ProtoT)
}

func NewFuzzTest[ProtoT proto.Message, KRMType any](protoType ProtoT, fromProto func(ctx *direct.MapContext, in ProtoT) *KRMType, toProto func(ctx *direct.MapContext, in *KRMType) ProtoT) *FuzzTest[ProtoT, KRMType] {
	return &FuzzTest[ProtoT, KRMType]{
		ProtoType:           protoType,
		FromProto:           fromProto,
		ToProto:             toProto,
		TargetFields:        sets.New[string](),
		UnimplementedFields: sets.New[string](),
		IgnoreFields:        sets.New[string](),
	}
}

func (f *FuzzTest[ProtoT, KRMType]) Fuzz(t *testing.T, seed int64) {
	randStream := rand.New(rand.NewSource(seed))

	p1 := proto.CloneOf(f.ProtoType)
	fuzz.FillWithRandom(t, randStream, p1)

	ignoreFields := sets.New[string]()
	ignoreFields = ignoreFields.Union(f.IgnoreFields)
	ignoreFields = ignoreFields.Union(f.UnimplementedFields)

	// Remove any output only or known-unimplemented fields
	clearFields := &fuzz.ClearFields{
		Paths: ignoreFields,
	}
	fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

	if f.Filter != nil {
		f.Filter(p1)
	}

	ctx := &direct.MapContext{}
	krm := f.FromProto(ctx, p1)
	if ctx.Err() != nil {
		t.Logf("p1 = %v", prototext.Format(p1))
		t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
	}

	p2 := f.ToProto(ctx, krm)
	if ctx.Err() != nil {
		t.Logf("p1 = %v", prototext.Format(p1))
		t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
	}

	if diff := cmp.Diff(p1, p2, protocmp.Transform()); diff != "" {
		t.Logf("p1 = %v", prototext.Format(p1))
		t.Logf("p2 = %v", prototext.Format(p2))
		t.Errorf("roundtrip failed for KRM %T; diff:\n%s", krm, diff)
		if hints := diagnoseRoundtripHints(p1, p2, fmt.Sprintf("%T", krm), f.TargetFields); len(hints) > 0 {
			for _, hint := range hints {
				t.Logf("HINT: %s", hint)
			}
			fmt.Fprintf(os.Stderr, "<hint_for_agent>%s</hint_for_agent>", strings.Join(hints, "\n"))
		}
	}
}

type roundtripDiffContext struct {
	rootProto    protoreflect.FullName
	krmType      string
	targetFields sets.Set[string]
}

func diagnoseRoundtripHints(m1, m2 proto.Message, krmType string, targetFields sets.Set[string]) []string {
	d := roundtripDiffContext{
		rootProto:    m1.ProtoReflect().Descriptor().FullName(),
		krmType:      krmType,
		targetFields: targetFields,
	}
	return deduplicateStrings(d.diffMessages("", m1.ProtoReflect(), m2.ProtoReflect()))
}

func (d roundtripDiffContext) diffMessages(prefix string, m1, m2 protoreflect.Message) []string {
	hints, reportedOneofs := d.diffOneofs(prefix, m1, m2)
	fields := m1.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if od := fd.ContainingOneof(); od != nil && reportedOneofs[od.Name()] {
			continue
		}
		hints = append(hints, d.diffField(prefix+"."+string(fd.Name()), fd, m1, m2)...)
	}
	return hints
}

func (d roundtripDiffContext) diffOneofs(prefix string, m1, m2 protoreflect.Message) ([]string, map[protoreflect.Name]bool) {
	oneofs := m1.Descriptor().Oneofs()
	reported := make(map[protoreflect.Name]bool)
	var hints []string
	for i := 0; i < oneofs.Len(); i++ {
		od := oneofs.Get(i)
		if od.IsSynthetic() {
			continue
		}
		fd1, fd2 := m1.WhichOneof(od), m2.WhichOneof(od)
		if fd2 != nil && fd1 != fd2 {
			reported[od.Name()] = true
			path := prefix + "." + string(fd2.Name())
			hints = append(hints, fmt.Sprintf(
				"Proto oneof %q at %q was overwritten by branch %q on roundtrip for %s. In _FromProto, wrap the reference or struct allocation for %q in 'if in.Get...() != \"\"' (or check WhichOneof) so an unset oneof branch does not clobber other branches in _ToProto.",
				od.Name(), path, fd2.Name(), d.krmType, fd2.Name(),
			))
		}
	}
	return hints, reported
}

func (d roundtripDiffContext) diffField(path string, fd protoreflect.FieldDescriptor, m1, m2 protoreflect.Message) []string {
	if fd.IsList() {
		return d.diffListField(path, fd, m1, m2)
	}
	if fd.IsMap() {
		if !m1.Get(fd).Equal(m2.Get(fd)) {
			return []string{d.fieldMissingOrModifiedHint(path)}
		}
		return nil
	}
	if fd.Kind() == protoreflect.MessageKind {
		return d.diffMessageField(path, fd, m1, m2)
	}
	if hint := d.diffScalarField(path, fd, m1, m2); hint != "" {
		return []string{hint}
	}
	return nil
}

func (d roundtripDiffContext) diffListField(path string, fd protoreflect.FieldDescriptor, m1, m2 protoreflect.Message) []string {
	l1, l2 := m1.Get(fd).List(), m2.Get(fd).List()
	if fd.Kind() == protoreflect.MessageKind {
		if l1.Len() == l2.Len() {
			var hints []string
			for j := 0; j < l1.Len(); j++ {
				hints = append(hints, d.diffMessages(path+"[]", l1.Get(j).Message(), l2.Get(j).Message())...)
			}
			return hints
		}
		if l1.Len() > l2.Len() && listContainsEmptyMessage(l1) {
			return []string{d.emptySubMessageDroppedHint(path + "[]")}
		}
	}
	if !m1.Get(fd).Equal(m2.Get(fd)) {
		return []string{d.fieldMissingOrModifiedHint(path)}
	}
	return nil
}

func (d roundtripDiffContext) diffMessageField(path string, fd protoreflect.FieldDescriptor, m1, m2 protoreflect.Message) []string {
	has1, has2 := m1.Has(fd), m2.Has(fd)
	switch {
	case has1 && has2:
		return d.diffMessages(path, m1.Get(fd).Message(), m2.Get(fd).Message())
	case has1 && !has2 && hasNoSetFields(m1.Get(fd).Message()):
		return []string{d.emptySubMessageDroppedHint(path)}
	case has1 != has2:
		return []string{d.fieldMissingOrModifiedHint(path)}
	default:
		return nil
	}
}

func (d roundtripDiffContext) diffScalarField(path string, fd protoreflect.FieldDescriptor, m1, m2 protoreflect.Message) string {
	has1, has2 := m1.Has(fd), m2.Has(fd)
	if has1 == has2 && (!has1 || m1.Get(fd).Equal(m2.Get(fd))) {
		return ""
	}
	if fd.HasPresence() && has1 && !has2 && isZeroProtoScalar(fd, m1.Get(fd)) {
		if fd.Kind() == protoreflect.EnumKind {
			return fmt.Sprintf(
				"Proto optional enum %q lost zero-value (0 / UNSPECIFIED) presence on roundtrip for %s because direct.Enum_FromProto maps 0 to nil. Either preserve zero-value enum presence in the mapper or mark %q with f.Unimplemented_NotYetTriaged(%q).",
				path, d.krmType, path, path,
			)
		}
		return fmt.Sprintf(
			"Proto optional field %q lost zero-value presence on roundtrip for %s. Ensure the corresponding KRM struct field is a pointer type (e.g. *string, *int64, *bool) and mapped with direct.LazyPtr / direct.ValueOf.",
			path, d.krmType,
		)
	}
	return d.fieldMissingOrModifiedHint(path)
}

func (d roundtripDiffContext) emptySubMessageDroppedHint(path string) string {
	return fmt.Sprintf(
		"Empty proto message at %q collapsed to nil on roundtrip for %s. Ensure _FromProto and _ToProto return an empty struct rather than nil when sub-fields are unset (remove any secondary 'if ... == nil { return nil }' check).",
		path, d.krmType,
	)
}

func (d roundtripDiffContext) fieldMissingOrModifiedHint(path string) string {
	if isPathCoveredBySet(path, d.targetFields) {
		return fmt.Sprintf(
			"Field %q is registered as a mapped field for %s, but its value was lost or modified during FromProto -> ToProto roundtrip. Fix _FromProto / _ToProto for %q (check for missing assignments or conditional 'if' guards), or configure f.FilterSpec / f.FilterStatus if %q requires proto state invariants.",
			path, d.krmType, path, path,
		)
	}
	return fmt.Sprintf(
		"Add `f.Unimplemented_NotYetTriaged(%q)` (or `f.SpecField(%q)` / `f.StatusField(%q)` if mapped) to the fuzzer for the proto type %v.",
		path, path, path, d.rootProto,
	)
}

func listContainsEmptyMessage(list protoreflect.List) bool {
	for i := 0; i < list.Len(); i++ {
		if hasNoSetFields(list.Get(i).Message()) {
			return true
		}
	}
	return false
}

func hasNoSetFields(m protoreflect.Message) bool {
	empty := true
	m.Range(func(_ protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		empty = false
		return false
	})
	return empty
}

func isZeroProtoScalar(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return !v.Bool()
	case protoreflect.EnumKind:
		return v.Enum() == 0
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return v.Int() == 0
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return v.Uint() == 0
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return v.Float() == 0
	case protoreflect.StringKind:
		return v.String() == ""
	case protoreflect.BytesKind:
		return len(v.Bytes()) == 0
	}
	return false
}

func isPathCoveredBySet(path string, s sets.Set[string]) bool {
	if s.Len() == 0 {
		return false
	}
	for p := path; p != ""; {
		if s.Has(p) {
			return true
		}
		for strings.HasSuffix(p, "[]") {
			p = strings.TrimSuffix(p, "[]")
			if s.Has(p) {
				return true
			}
		}
		dot := strings.LastIndexByte(p, '.')
		if dot <= 0 {
			break
		}
		p = p[:dot]
	}
	return false
}

func deduplicateStrings(in []string) []string {
	seen := make(map[string]bool, len(in))
	var out []string
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

// VisitValues is a helper function that visits all values in a proto message,
// calling the provided function for each value.
// It is useful for applying filters.
func VisitValues(m proto.Message, fn func(path string, fd protoreflect.FieldDescriptor, v protoreflect.Value) protoreflect.Value) {
	visitValues("", m.ProtoReflect(), fn)
}

func visitValues(parentPath string, m protoreflect.Message, fn func(path string, fd protoreflect.FieldDescriptor, v protoreflect.Value) protoreflect.Value) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		path := parentPath + "." + string(fd.Name())
		v2 := fn(path, fd, v)
		m.Set(fd, v2)

		switch fd.Kind() {
		case protoreflect.MessageKind, protoreflect.GroupKind:
			if fd.IsList() {
				list := v.List()
				for i := 0; i < list.Len(); i++ {
					visitValues(path+"[]", list.Get(i).Message(), fn)
				}
			} else if fd.IsMap() {
				mapField := v.Map()
				if fd.MapValue().Kind() == protoreflect.MessageKind {
					mapField.Range(func(mapKey protoreflect.MapKey, mapValue protoreflect.Value) bool {
						visitValues(path+"[key="+mapKey.String()+"]", mapValue.Message(), fn)
						return true
					})
				}
			} else {
				visitValues(path, v.Message(), fn)
			}
		}
		return true
	})
}
