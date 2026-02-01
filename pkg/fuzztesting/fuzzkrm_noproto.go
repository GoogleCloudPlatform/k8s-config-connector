// Copyright 2025 Google LLC
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
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func RegisterKRMFuzzer_NoProto(fuzzer KRMFuzzer_NoProto) {
	RegisterFuzzer(fuzzer.FuzzSpec)
	RegisterFuzzer(fuzzer.FuzzStatus)
}

func RegisterKRMSpecFuzzer_NoProto(fuzzer KRMFuzzer_NoProto) {
	RegisterFuzzer(fuzzer.FuzzSpec)
}

type KRMTypedFuzzer_NoProto[APIType any, SpecType any, StatusType any] struct {
	APIType APIType

	SpecFromAPI func(ctx *direct.MapContext, in APIType) *SpecType
	SpecToAPI   func(ctx *direct.MapContext, in *SpecType) APIType

	StatusFromAPI func(ctx *direct.MapContext, in APIType) *StatusType
	StatusToAPI   func(ctx *direct.MapContext, in *StatusType) APIType

	UnimplementedFields sets.Set[string]
	SpecFields          sets.Set[string]
	StatusFields        sets.Set[string]

	FilterSpec   func(in APIType)
	FilterStatus func(in APIType)
}

// SpecField marks the specified fieldPath as round-tripping to/from the Spec
func (f *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]) SpecField(fieldPath string) {
	f.SpecFields.Insert(fieldPath)
}

// StatusField marks the specified fieldPath as round-tripping to/from the Status
func (f *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]) StatusField(fieldPath string) {
	f.StatusFields.Insert(fieldPath)
}

// Unimplemented_NotYetTriaged marks the specified fieldPath as not round-tripped,
// and should be used for fields that are added by the service and where we haven't decided whether or not to implement them.
func (f *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]) Unimplemented_NotYetTriaged(fieldPath string) {
	f.UnimplementedFields.Insert(fieldPath)
}

type KRMFuzzer_NoProto interface {
	FuzzSpec(t *testing.T, seed int64)
	FuzzStatus(t *testing.T, seed int64)
}

func NewKRMTypedFuzzer_NoProto[APIType any, SpecType any, StatusType any](
	apiType APIType,
	specFromAPI func(ctx *direct.MapContext, in APIType) *SpecType, specToAPI func(ctx *direct.MapContext, in *SpecType) APIType,
	statusFromAPI func(ctx *direct.MapContext, in APIType) *StatusType, statusToAPI func(ctx *direct.MapContext, in *StatusType) APIType,
) *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType] {
	return &KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]{
		APIType:             apiType,
		SpecFromAPI:         specFromAPI,
		SpecToAPI:           specToAPI,
		StatusFromAPI:       statusFromAPI,
		StatusToAPI:         statusToAPI,
		UnimplementedFields: sets.New[string](),
		SpecFields:          sets.New[string](),
		StatusFields:        sets.New[string](),
	}
}

func (f *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]) FuzzSpec(t *testing.T, seed int64) {
	fuzzer := NewFuzzTest_NoProto(f.APIType, f.SpecFromAPI, f.SpecToAPI)
	fuzzer.IgnoreFields = f.StatusFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterSpec
	fuzzer.Fuzz(t, seed)
}

func (f *KRMTypedFuzzer_NoProto[APIType, SpecType, StatusType]) FuzzStatus(t *testing.T, seed int64) {
	fuzzer := NewFuzzTest_NoProto(f.APIType, f.StatusFromAPI, f.StatusToAPI)
	fuzzer.IgnoreFields = f.SpecFields
	fuzzer.UnimplementedFields = f.UnimplementedFields
	fuzzer.Filter = f.FilterStatus
	fuzzer.Fuzz(t, seed)
}

// // NewKRMTypedSpecFuzzer_NoProto is a convenience function for creating a fuzzer that only
// // fuzzes the spec fields of a KRM type.
// func NewKRMTypedSpecFuzzer_NoProto[ProtoT proto.Message, SpecType any](
// 	protoType ProtoT,
// 	specFromProto func(ctx *direct.MapContext, in ProtoT) *SpecType,
// 	specToProto func(ctx *direct.MapContext, in *SpecType) ProtoT,
// ) *KRMTypedFuzzer_NoProto[ProtoT, SpecType] {
// 	return &KRMTypedFuzzer_NoProto[ProtoT, SpecType]{
// 		ProtoType:           protoType,
// 		SpecFromProto:       specFromProto,
// 		SpecToProto:         specToProto,
// 		StatusFromProto:     nil, // No status functions
// 		StatusToProto:       nil, // No status functions
// 		UnimplementedFields: sets.New[string](),
// 		SpecFields:          sets.New[string](),
// 		StatusFields:        sets.New[string](),
// 	}
// }

type FuzzTest_NoProto[APIType any, KRMType any] struct {
	APIType APIType

	FromAPI func(ctx *direct.MapContext, in APIType) *KRMType
	ToAPI   func(ctx *direct.MapContext, in *KRMType) APIType

	UnimplementedFields sets.Set[string]
	IgnoreFields        sets.Set[string]

	Filter func(in APIType)
}

func NewFuzzTest_NoProto[APIType any, KRMType any](apiType APIType, fromAPI func(ctx *direct.MapContext, in APIType) *KRMType, toAPI func(ctx *direct.MapContext, in *KRMType) APIType) *FuzzTest_NoProto[APIType, KRMType] {
	return &FuzzTest_NoProto[APIType, KRMType]{
		APIType:             apiType,
		FromAPI:             fromAPI,
		ToAPI:               toAPI,
		UnimplementedFields: sets.New[string](),
		IgnoreFields:        sets.New[string](),
	}
}

func (f *FuzzTest_NoProto[APIType, KRMType]) Fuzz(t *testing.T, seed int64) {
	randStream := rand.New(rand.NewSource(seed))

	ignoreFields := sets.New[string]()
	ignoreFields = ignoreFields.Union(f.IgnoreFields)
	ignoreFields = ignoreFields.Union(f.UnimplementedFields)

	overrides := map[string]fuzz.OverrideFiller{}
	zeroOverrides := map[string]fuzz.OverrideFiller{}

	for ignoreField := range ignoreFields {
		overrides[ignoreField] = func(t *testing.T, fieldName string, field reflect.Value) {
			// Do nothing for ignored fields during random fill
		}
		zeroOverrides[ignoreField] = func(t *testing.T, fieldName string, field reflect.Value) {
			field.Set(reflect.Zero(field.Type()))
		}
	}

	filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: randStream, FieldOverrides: overrides})

	p1 := reflect.New(reflect.ValueOf(f.APIType).Type().Elem()).Interface().(APIType)
	filler.Fill(t, p1)

	if f.Filter != nil {
		f.Filter(p1)
	}

	ctx := &direct.MapContext{}
	krm := f.FromAPI(ctx, p1)
	if ctx.Err() != nil {
		t.Logf("p1 = %v", p1)
		t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
	}

	p2 := f.ToAPI(ctx, krm)
	if ctx.Err() != nil {
		t.Logf("p1 = %v", p1)
		t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
	}

	zeroFiller := fuzz.NewZeroFiller(&fuzz.FillerConfig{FieldOverrides: zeroOverrides})
	zeroFiller.Fill(t, p1)
	zeroFiller.Fill(t, p2)

	if diff := cmp.Diff(p1, p2); diff != "" {
		t.Logf("p1 = %v", p1)
		t.Logf("p2 = %v", p2)
		t.Errorf("roundtrip failed for KRM %T; diff:\n%s", krm, diff)
	}
}
