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
	"math/rand"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
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

	// MutateBeforeRoundtrip is used to change the generated proto to be valid,
	// where the round-tripping performs some validation
	MutateBeforeRoundtrip func(v ProtoT) error

	// CompareOptions is used to set options for cmp.Diff, typically to ignore
	// fields we set in MutateBeforeRoundtrip
	CompareOptions []cmp.Option
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
	fuzzer.MutateBeforeRoundtrip = f.MutateBeforeRoundtrip
	fuzzer.CompareOptions = f.CompareOptions
	fuzzer.Fuzz(t, seed)
}

func (f *KRMTypedFuzzer[ProtoT, SpecType, StatusType]) FuzzStatus(t *testing.T, seed int64) {
	fuzzer := NewFuzzTest(f.ProtoType, f.StatusFromProto, f.StatusToProto)
	fuzzer.IgnoreFields = f.SpecFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.MutateBeforeRoundtrip = f.MutateBeforeRoundtrip
	fuzzer.CompareOptions = f.CompareOptions
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

	// MutateBeforeRoundtrip is used to change the generated proto to be valid,
	// where the round-tripping performs some validation
	MutateBeforeRoundtrip func(v ProtoT) error

	// CompareOptions is used to set options for cmp.Diff, typically to ignore
	// fields we set in MutateBeforeRoundtrip
	CompareOptions []cmp.Option
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

	p1 := proto.Clone(f.ProtoType).(ProtoT)
	fuzz.FillWithRandom(t, randStream, p1)

	ignoreFields := sets.New[string]()
	ignoreFields = ignoreFields.Union(f.IgnoreFields)
	ignoreFields = ignoreFields.Union(f.UnimplementedFields)

	// Remove any output only or known-unimplemented fields
	clearFields := &fuzz.ClearFields{
		Paths: ignoreFields,
	}
	fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

	if f.MutateBeforeRoundtrip != nil {
		if err := f.MutateBeforeRoundtrip(p1); err != nil {
			t.Fatalf("error replacing fields with MutateBeforeRoundtrip: %v", err)
		}
	}

	ctx := &direct.MapContext{}
	krm := f.FromProto(ctx, p1)
	if ctx.Err() != nil {
		t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
	}

	p2 := f.ToProto(ctx, krm)
	if ctx.Err() != nil {
		t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
	}

	cmpOptions := []cmp.Option{
		protocmp.Transform(),
	}
	cmpOptions = append(cmpOptions, f.CompareOptions...)
	if diff := cmp.Diff(p1, p2, cmpOptions...); diff != "" {
		t.Logf("p1 = %v", prototext.Format(p1))
		t.Logf("p2 = %v", prototext.Format(p2))
		t.Errorf("roundtrip failed for KRM %T; diff:\n%s", krm, diff)
	}
}
