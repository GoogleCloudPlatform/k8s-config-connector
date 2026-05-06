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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/dynamicpb"
	"k8s.io/apimachinery/pkg/util/sets"
)

type FuzzFn func(t *testing.T, seed int64)

var fuzzers []FuzzFn

func RegisterKRMFuzzer(fuzzer KRMFuzzer) {
	RegisterFuzzer(fuzzer.FuzzSpec)
	RegisterFuzzer(fuzzer.FuzzStatus)
}

func RegisterKRMSpecFuzzer(fuzzer KRMFuzzer) {
	RegisterFuzzer(fuzzer.FuzzSpec)
}

func RegisterFuzzer(fuzzer FuzzFn) {
	fuzzers = append(fuzzers, fuzzer)
}

func ChooseFuzzer(n int64) FuzzFn {
	return fuzzers[n%int64(len(fuzzers))]
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

	ProtoSourcePath string
}

// WithProtoSource sets the path to the .pb file containing the descriptors for this fuzzer.
func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) WithProtoSource(path string) *KRMTypedFuzzer[ProtoT, SpecType, StatusType] {
	f.ProtoSourcePath = path
	return f
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
	fuzzer.IgnoreFields = f.StatusFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterSpec
	fuzzer.ProtoSourcePath = f.ProtoSourcePath
	fuzzer.Fuzz(t, seed)
}

func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) FuzzStatus(t *testing.T, seed int64) {
	fuzzer := NewFuzzTest(f.ProtoType, f.StatusFromProto, f.StatusToProto)
	fuzzer.IgnoreFields = f.SpecFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterStatus
	fuzzer.ProtoSourcePath = f.ProtoSourcePath
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

	UnimplementedFields sets.Set[string]
	IgnoreFields        sets.Set[string]

	Filter func(in ProtoT)

	ProtoSourcePath string
}

func NewFuzzTest[ProtoT proto.Message, KRMType any](protoType ProtoT, fromProto func(ctx *direct.MapContext, in ProtoT) *KRMType, toProto func(ctx *direct.MapContext, in *KRMType) ProtoT) *FuzzTest[ProtoT, KRMType] {
	return &FuzzTest[ProtoT, KRMType]{
		ProtoType:           protoType,
		FromProto:           fromProto,
		ToProto:             toProto,
		UnimplementedFields: sets.New[string](),
		IgnoreFields:        sets.New[string](),
	}
}

func (f *FuzzTest[ProtoT, KRMType]) Fuzz(t *testing.T, seed int64) {
	randStream := rand.New(rand.NewSource(seed))

	var p1_comp ProtoT

	if f.ProtoSourcePath != "" {
		var md protoreflect.MessageDescriptor
		files, err := fuzz.GetProtoRegistry(f.ProtoSourcePath)
		if err != nil {
			t.Fatalf("failed to load proto source %q: %v", f.ProtoSourcePath, err)
		}
		fullName := f.ProtoType.ProtoReflect().Descriptor().FullName()
		desc, err := files.FindDescriptorByName(fullName)
		if err != nil {
			t.Fatalf("failed to find descriptor for %q in %q: %v", fullName, f.ProtoSourcePath, err)
		}
		md = desc.(protoreflect.MessageDescriptor)
		p1_dyn := dynamicpb.NewMessage(md)
		fuzz.FillWithRandom(t, randStream, p1_dyn)

		// Remove any output only or known-unimplemented fields
		ignoreFields := sets.New[string]()
		ignoreFields = ignoreFields.Union(f.IgnoreFields)
		ignoreFields = ignoreFields.Union(f.UnimplementedFields)
		clearFields := &fuzz.ClearFields{
			Paths: ignoreFields,
		}
		fuzz.Visit("", p1_dyn.ProtoReflect(), nil, clearFields)

		// Convert dynamic to compiled for the mapper
		p1_comp = proto.Clone(f.ProtoType).(ProtoT)
		b, err := proto.Marshal(p1_dyn)
		if err != nil {
			t.Fatalf("failed to marshal dynamic message: %v", err)
		}
		if err := proto.Unmarshal(b, p1_comp); err != nil {
			t.Fatalf("failed to unmarshal bytes to compiled message: %v", err)
		}

		if f.Filter != nil {
			f.Filter(p1_comp)
		}

		ctx := &direct.MapContext{}
		krm := f.FromProto(ctx, p1_comp)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2_comp := f.ToProto(ctx, krm)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		// Convert p2 back to dynamic for comparison
		p2_dyn := dynamicpb.NewMessage(md)
		b2, err := proto.Marshal(p2_comp)
		if err != nil {
			t.Fatalf("failed to marshal compiled message p2: %v", err)
		}
		if err := proto.Unmarshal(b2, p2_dyn); err != nil {
			t.Fatalf("failed to unmarshal bytes to dynamic message p2: %v", err)
		}

		if diff := cmp.Diff(p1_dyn, p2_dyn, protocmp.Transform()); diff != "" {
			t.Errorf("roundtrip failed for KRM %T; diff:\n%s", krm, diff)
			diffPaths := diffFieldPaths(p1_dyn, p2_dyn)
			for _, diffPath := range diffPaths {
				hint := fmt.Sprintf("Add `f.Unimplemented_NotYetTriaged(%q)` to the fuzzer for the proto type %v to mark this field as not yet triaged.", diffPath, f.ProtoType.ProtoReflect().Descriptor().FullName())
				fmt.Fprintf(os.Stderr, "<hint_for_agent>%s</hint_for_agent>", hint)
				break
			}
		}
	} else {
		p1_comp = proto.Clone(f.ProtoType).(ProtoT)
		fuzz.FillWithRandom(t, randStream, p1_comp)

		ignoreFields := sets.New[string]()
		ignoreFields = ignoreFields.Union(f.IgnoreFields)
		ignoreFields = ignoreFields.Union(f.UnimplementedFields)

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: ignoreFields,
		}
		fuzz.Visit("", p1_comp.ProtoReflect(), nil, clearFields)

		if f.Filter != nil {
			f.Filter(p1_comp)
		}

		ctx := &direct.MapContext{}
		krm := f.FromProto(ctx, p1_comp)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2_comp := f.ToProto(ctx, krm)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		if diff := cmp.Diff(p1_comp, p2_comp, protocmp.Transform()); diff != "" {
			t.Errorf("roundtrip failed for KRM %T; diff:\n%s", krm, diff)
			diffPaths := diffFieldPaths(p1_comp, p2_comp)
			for _, diffPath := range diffPaths {
				hint := fmt.Sprintf("Add `f.Unimplemented_NotYetTriaged(%q)` to the fuzzer for the proto type %v to mark this field as not yet triaged.", diffPath, f.ProtoType.ProtoReflect().Descriptor().FullName())
				fmt.Fprintf(os.Stderr, "<hint_for_agent>%s</hint_for_agent>", hint)
				break
			}
		}
	}
}

// diffFieldPaths returns the field paths that differ between two proto messages.
func diffFieldPaths(m1, m2 proto.Message) []string {
	paths1 := sets.New[string]()
	VisitValues(m1, func(path string, fd protoreflect.FieldDescriptor, v protoreflect.Value) protoreflect.Value {
		paths1.Insert(path)
		return v
	})
	paths2 := sets.New[string]()
	VisitValues(m2, func(path string, fd protoreflect.FieldDescriptor, v protoreflect.Value) protoreflect.Value {
		paths2.Insert(path)
		return v
	})
	return paths1.Difference(paths2).UnsortedList()
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
