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

package common

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/sets"
)

func protoMapToMap(m protoreflect.Map) map[string]interface{} {
	res := make(map[string]interface{})
	m.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		if mv, ok := v.Interface().(protoreflect.Message); ok {
			res[k.String()] = mv.Interface()
		} else {
			res[k.String()] = v.Interface()
		}
		return true
	})
	return res
}

func protoListToList(l protoreflect.List) []interface{} {
	res := make([]interface{}, l.Len())
	for i := 0; i < l.Len(); i++ {
		v := l.Get(i)
		if m, ok := v.Interface().(protoreflect.Message); ok {
			res[i] = m.Interface()
		} else {
			res[i] = v.Interface()
		}
	}
	return res
}

func TestCompareProtoMessageStructuredDiff(t *testing.T) {
	fooVal, err := structpb.NewValue("foo")
	if err != nil {
		t.Fatal(err)
	}
	barVal, err := structpb.NewValue("bar")
	if err != nil {
		t.Fatal(err)
	}

	descA := &descriptorpb.FieldDescriptorProto{
		Name: stringPtr("field1"),
		Options: &descriptorpb.FieldOptions{
			Deprecated: boolPtr(false),
		},
	}
	descB := &descriptorpb.FieldDescriptorProto{
		Name: stringPtr("field2"),
		Options: &descriptorpb.FieldOptions{
			Deprecated: boolPtr(true),
		},
	}

	tests := []struct {
		name      string
		msgA      proto.Message
		msgB      proto.Message
		wantPaths sets.Set[string]
		wantDiffs []structuredreporting.DiffField
	}{
		{
			name:      "identical",
			msgA:      fooVal,
			msgB:      fooVal,
			wantPaths: sets.Set[string]{},
			wantDiffs: []structuredreporting.DiffField{},
		},
		{
			name:      "identical different memory",
			msgA:      fooVal,
			msgB:      &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "foo"}},
			wantPaths: sets.Set[string]{},
			wantDiffs: []structuredreporting.DiffField{},
		},
		{
			name:      "descA vs descA",
			msgA:      descA,
			msgB:      descA,
			wantPaths: sets.Set[string]{},
			wantDiffs: []structuredreporting.DiffField{},
		},
		{
			name: "descA vs descA different memory",
			msgA: descA,
			msgB: &descriptorpb.FieldDescriptorProto{
				Name: stringPtr("field1"),
				Options: &descriptorpb.FieldOptions{
					Deprecated: boolPtr(false),
				},
			},
			wantPaths: sets.Set[string]{},
			wantDiffs: []structuredreporting.DiffField{},
		},
		{
			name:      "different string value",
			msgA:      fooVal,
			msgB:      barVal,
			wantPaths: sets.New("string_value"),
			wantDiffs: []structuredreporting.DiffField{
				{ID: "string_value", Old: "bar", New: "foo"},
			},
		},
		{
			name:      "nested message value",
			msgA:      descA,
			msgB:      descB,
			wantPaths: sets.New("name", "options.deprecated"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID:  "name",
					Old: "field2",
					New: "field1",
				},
				{
					ID:  "options.deprecated",
					Old: true,
					New: false,
				},
			},
		},
		{
			name: "list change",
			msgA: &descriptorpb.DescriptorProto{
				Field: []*descriptorpb.FieldDescriptorProto{descA},
			},
			msgB: &descriptorpb.DescriptorProto{
				Field: []*descriptorpb.FieldDescriptorProto{descB},
			},
			wantPaths: sets.New("field"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID:  "field",
					Old: []interface{}{descB},
					New: []interface{}{descA},
				},
			},
		},
		{
			name: "list item added",
			msgA: &descriptorpb.DescriptorProto{
				Field: []*descriptorpb.FieldDescriptorProto{descA, descB},
			},
			msgB: &descriptorpb.DescriptorProto{
				Field: []*descriptorpb.FieldDescriptorProto{descA},
			},
			wantPaths: sets.New("field"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID:  "field",
					Old: []interface{}{descA},
					New: []interface{}{descA, descB},
				},
			},
		},
		{
			name:      "timestamp change",
			msgA:      &timestamppb.Timestamp{Seconds: 100},
			msgB:      &timestamppb.Timestamp{Seconds: 200},
			wantPaths: sets.New("seconds"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID:  "seconds",
					Old: int64(200),
					New: int64(100),
				},
			},
		},
		{
			name:      "duration change",
			msgA:      &durationpb.Duration{Seconds: 10},
			msgB:      &durationpb.Duration{Seconds: 20},
			wantPaths: sets.New("seconds"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID:  "seconds",
					Old: int64(20),
					New: int64(10),
				},
			},
		},
		{
			name: "map change",
			msgA: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"key1": fooVal,
				},
			},
			msgB: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"key1": barVal,
				},
			},
			wantPaths: sets.New("fields"),
			wantDiffs: []structuredreporting.DiffField{
				{
					ID: "fields",
					Old: map[string]interface{}{
						"key1": barVal,
					},
					New: map[string]interface{}{
						"key1": fooVal,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			paths, diff, err := CompareProtoMessageStructuredDiff(tc.msgA, tc.msgB, BasicDiff)
			if err != nil {
				t.Fatalf("CompareProtoMessageStructuredDiff() error = %v", err)
			}
			if !paths.Equal(tc.wantPaths) {
				t.Errorf("CompareProtoMessageStructuredDiff() paths = %v, want %v", paths, tc.wantPaths)
			}

			gotDiffs := []structuredreporting.DiffField{}
			if diff != nil {
				gotDiffs = diff.Fields
			}

			if len(gotDiffs) != len(tc.wantDiffs) {
				t.Errorf("CompareProtoMessageStructuredDiff() diffs len = %d, want %d", len(gotDiffs), len(tc.wantDiffs))
			}
			for i, d := range tc.wantDiffs {
				if i >= len(gotDiffs) {
					break
				}
				got := gotDiffs[i]
				if got.ID != d.ID {
					t.Errorf("diff[%d].ID = %q, want %q", i, got.ID, d.ID)
				}

				normalize := func(v interface{}) interface{} {
					if pm, ok := v.(protoreflect.Map); ok {
						return protoMapToMap(pm)
					}
					if pl, ok := v.(protoreflect.List); ok {
						return protoListToList(pl)
					}
					return v
				}

				gotOld := normalize(got.Old)
				gotNew := normalize(got.New)

				if diff := cmp.Diff(d.Old, gotOld, protocmp.Transform()); diff != "" {
					t.Errorf("diff[%d].Old mismatch (-want +got):\n%s", i, diff)
				}
				if diff := cmp.Diff(d.New, gotNew, protocmp.Transform()); diff != "" {
					t.Errorf("diff[%d].New mismatch (-want +got):\n%s", i, diff)
				}
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

type sampleKRMNestedSub struct {
	StrVal    *string           `json:"strVal,omitempty"`
	IntVal    *int64            `json:"intVal,omitempty"`
	BoolVal   *bool             `json:"boolVal,omitempty"`
	FloatVal  *float64          `json:"floatVal,omitempty"`
	MapVal    map[string]string `json:"mapVal,omitempty"`
	ListVal   []string          `json:"listVal,omitempty"`
	NonPtrStr string            `json:"nonPtrStr,omitempty"`
	NonPtrInt int64             `json:"nonPtrInt,omitempty"`
}

type sampleKRMSub struct {
	// Nested primitive pointers
	StrVal   *string  `json:"strVal,omitempty"`
	IntVal   *int64   `json:"intVal,omitempty"`
	BoolVal  *bool    `json:"boolVal,omitempty"`
	FloatVal *float64 `json:"floatVal,omitempty"`

	// Deeply nested struct pointer
	NestedSub *sampleKRMNestedSub `json:"nestedSub,omitempty"`

	// Nested maps
	StringMap map[string]string              `json:"stringMap,omitempty"`
	StructMap map[string]*sampleKRMNestedSub `json:"structMap,omitempty"`

	// Nested slices
	StringSlice []string              `json:"stringSlice,omitempty"`
	StructSlice []*sampleKRMNestedSub `json:"structSlice,omitempty"`

	// Non-pointer fields
	NonPtrStr string `json:"nonPtrStr,omitempty"`
	NonPtrInt int64  `json:"nonPtrInt,omitempty"`
}

type sampleKRMSpec struct {
	// Top-level primitive pointers
	StrField   *string  `json:"strField,omitempty"`
	IntField   *int64   `json:"intField,omitempty"`
	BoolField  *bool    `json:"boolField,omitempty"`
	FloatField *float64 `json:"floatField,omitempty"`

	// Top-level struct pointers
	SubStruct *sampleKRMSub `json:"subStruct,omitempty"`

	// Top-level maps
	StringMap map[string]string        `json:"stringMap,omitempty"`
	StructMap map[string]*sampleKRMSub `json:"structMap,omitempty"`

	// Top-level slices
	StringSlice []string        `json:"stringSlice,omitempty"`
	StructSlice []*sampleKRMSub `json:"structSlice,omitempty"`

	// Non-pointer fields
	NonPtrStr string `json:"nonPtrStr,omitempty"`
	NonPtrInt int64  `json:"nonPtrInt,omitempty"`
}

func int64Ptr(i int64) *int64 {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func TestMergeUnsetFields_ValidValues(t *testing.T) {
	tests := []struct {
		name    string
		desired *sampleKRMSpec
		actual  *sampleKRMSpec
		want    *sampleKRMSpec
	}{
		{
			name: "primitive pointers: adopts actual if nil, preserves desired if non-nil",
			desired: &sampleKRMSpec{
				StrField: stringPtr("user-specified"),
				// IntField, BoolField, FloatField left nil by user
			},
			actual: &sampleKRMSpec{
				StrField:   stringPtr("gcp-string"),
				IntField:   int64Ptr(42),
				BoolField:  boolPtr(true),
				FloatField: float64Ptr(3.14),
			},
			want: &sampleKRMSpec{
				StrField:   stringPtr("user-specified"), // Preserved user value
				IntField:   int64Ptr(42),                // Adopted actual GCP value
				BoolField:  boolPtr(true),               // Adopted actual GCP value
				FloatField: float64Ptr(3.14),            // Adopted actual GCP value
			},
		},
		{
			name: "nested struct pointers: merges recursively for unspecified subfields",
			desired: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					StrVal: stringPtr("user-nested-string"),
					// IntVal, BoolVal, FloatVal left nil by user
				},
			},
			actual: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					StrVal:   stringPtr("gcp-nested-string"),
					IntVal:   int64Ptr(100),
					BoolVal:  boolPtr(false),
					FloatVal: float64Ptr(2.718),
				},
			},
			want: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					StrVal:   stringPtr("user-nested-string"), // Preserved user value
					IntVal:   int64Ptr(100),                   // Adopted actual GCP value
					BoolVal:  boolPtr(false),                  // Adopted actual GCP value
					FloatVal: float64Ptr(2.718),               // Adopted actual GCP value
				},
			},
		},
		{
			name: "deeply nested structs (3 levels): merges sub-sub-fields",
			desired: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					NestedSub: &sampleKRMNestedSub{
						StrVal: stringPtr("user-deep-string"),
						// IntVal, BoolVal left nil
					},
				},
			},
			actual: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					NestedSub: &sampleKRMNestedSub{
						StrVal:  stringPtr("gcp-deep-string"),
						IntVal:  int64Ptr(999),
						BoolVal: boolPtr(true),
					},
				},
			},
			want: &sampleKRMSpec{
				SubStruct: &sampleKRMSub{
					NestedSub: &sampleKRMNestedSub{
						StrVal:  stringPtr("user-deep-string"), // Preserved user value
						IntVal:  int64Ptr(999),                 // Adopted actual GCP value
						BoolVal: boolPtr(true),                 // Adopted actual GCP value
					},
				},
			},
		},
		{
			name:    "maps: adopts actual map if desired map is nil",
			desired: &sampleKRMSpec{
				// StringMap left nil
			},
			actual: &sampleKRMSpec{
				StringMap: map[string]string{"env": "prod"},
			},
			want: &sampleKRMSpec{
				StringMap: map[string]string{"env": "prod"}, // Adopted actual map
			},
		},
		{
			name: "slices: adopts actual slice if desired is nil, preserves if non-nil",
			desired: &sampleKRMSpec{
				StringSlice: []string{"user-item"},
				// StructSlice left nil
			},
			actual: &sampleKRMSpec{
				StringSlice: []string{"gcp-item-1", "gcp-item-2"},
				StructSlice: []*sampleKRMSub{{StrVal: stringPtr("gcp-sub")}},
			},
			want: &sampleKRMSpec{
				StringSlice: []string{"user-item"},                           // Preserved user slice
				StructSlice: []*sampleKRMSub{{StrVal: stringPtr("gcp-sub")}}, // Adopted actual slice
			},
		},
		{
			name: "non-pointer primitive fields: remain untouched",
			desired: &sampleKRMSpec{
				NonPtrStr: "user-value",
				// NonPtrInt left zero-value 0
			},
			actual: &sampleKRMSpec{
				NonPtrStr: "gcp-value",
				NonPtrInt: 999,
			},
			want: &sampleKRMSpec{
				NonPtrStr: "user-value", // Preserved user non-pointer value
				NonPtrInt: 0,            // Zero-value left untouched (not overwritten)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			MergeUnsetFields(reflect.ValueOf(tc.desired), reflect.ValueOf(tc.actual))
			if diff := cmp.Diff(tc.want, tc.desired); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestMergeUnsetFields_InvalidValues(t *testing.T) {
	t.Run("handles invalid reflect values safely without panic", func(t *testing.T) {
		// Verify that invalid/zero reflect.Value inputs hit the top guard clause
		// (if !desired.IsValid() || !actual.IsValid() { return }) and return safely.
		MergeUnsetFields(reflect.Value{}, reflect.Value{})
	})
}

func TestSortRepeatedFields_ValidValues(t *testing.T) {
	tests := []struct {
		name  string
		input proto.Message
		want  proto.Message
	}{
		{
			name: "sorts primitive list in structpb.ListValue",
			input: func() proto.Message {
				list, err := structpb.NewList([]interface{}{"banana", "apple", "cherry"})
				if err != nil {
					t.Fatal(err)
				}
				return list
			}(),
			want: func() proto.Message {
				list, err := structpb.NewList([]interface{}{"apple", "banana", "cherry"})
				if err != nil {
					t.Fatal(err)
				}
				return list
			}(),
		},
		{
			name: "sorts repeated proto messages inside message",
			input: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: stringPtr("Zebra")},
					{Name: stringPtr("Alpha")},
				},
			},
			want: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: stringPtr("Alpha")},
					{Name: stringPtr("Zebra")},
				},
			},
		},
		{
			name: "traverses map values with nested messages across multiple unsorted map keys",
			input: func() proto.Message {
				m, err := structpb.NewStruct(map[string]interface{}{
					"z_config": map[string]interface{}{
						"items": []interface{}{"b", "a"},
					},
					"a_config": map[string]interface{}{
						"items": []interface{}{"y", "x", "z"},
					},
				})
				if err != nil {
					t.Fatal(err)
				}
				return m
			}(),
			want: func() proto.Message {
				m, err := structpb.NewStruct(map[string]interface{}{
					"z_config": map[string]interface{}{
						"items": []interface{}{"a", "b"},
					},
					"a_config": map[string]interface{}{
						"items": []interface{}{"x", "y", "z"},
					},
				})
				if err != nil {
					t.Fatal(err)
				}
				return m
			}(),
		},
		{
			name: "sorts nested lists inside elements of a list",
			input: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: stringPtr("Zebra"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: stringPtr("z_field_2")},
							{Name: stringPtr("z_field_1")},
						},
					},
					{
						Name: stringPtr("Alpha"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: stringPtr("a_field_2")},
							{Name: stringPtr("a_field_1")},
						},
					},
				},
			},
			want: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: stringPtr("Alpha"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: stringPtr("a_field_1")},
							{Name: stringPtr("a_field_2")},
						},
					},
					{
						Name: stringPtr("Zebra"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: stringPtr("z_field_1")},
							{Name: stringPtr("z_field_2")},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			SortRepeatedFields(tc.input.ProtoReflect())
			if diff := cmp.Diff(tc.want, tc.input, protocmp.Transform()); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSortRepeatedFields_InvalidValues(t *testing.T) {
	t.Run("handles nil message safely without panic", func(t *testing.T) {
		// Verify that nil protoreflect.Message hits the guard clause (if msg == nil || !msg.IsValid()) and returns safely.
		SortRepeatedFields(nil)
	})
}

// TestCompareBrownfieldSpec tests the 5-step brownfield reconciliation workflow.
// We use *descriptorpb.FileDescriptorProto as the test ProtoT message because it is a standard
// library proto.Message with a rich variety of fields (primitives, top-level slices, sub-messages, booleans, and options)
// to exercise reflection, merging, normalization, and fieldmask generation hermetically.
func TestCompareBrownfieldSpec(t *testing.T) {
	// specFromProto maps from GCP live state (*descriptorpb.FileDescriptorProto) to KRM spec (*sampleKRMSpec):
	//   - in.Name                        -> res.StrField
	//   - in.GetPackage()                -> res.NonPtrStr
	//   - in.Options.Deprecated          -> res.BoolField
	//   - in.Service[0].Name             -> res.SubStruct.StrVal
	//   - in.MessageType[i].GetName()    -> res.StringSlice[i] (top-level repeated slice entry)
	//   - in.Options.UninterpretedOption -> res.StringMap
	specFromProto := func(mapCtx *direct.MapContext, in *descriptorpb.FileDescriptorProto) *sampleKRMSpec {
		if in == nil {
			return nil
		}
		res := &sampleKRMSpec{
			StrField:  in.Name,
			NonPtrStr: in.GetPackage(),
		}
		if in.Options != nil && in.Options.Deprecated != nil {
			val := *in.Options.Deprecated
			res.BoolField = &val
		}
		if len(in.Service) > 0 {
			res.SubStruct = &sampleKRMSub{
				StrVal: in.Service[0].Name,
			}
		}
		if len(in.MessageType) > 0 {
			for _, m := range in.MessageType {
				res.StringSlice = append(res.StringSlice, m.GetName())
			}
		}
		if in.Options != nil && len(in.Options.UninterpretedOption) > 0 {
			res.StringMap = make(map[string]string)
			for _, opt := range in.Options.UninterpretedOption {
				if len(opt.Name) > 0 {
					key := opt.Name[0].GetNamePart()
					val := string(opt.GetStringValue())
					res.StringMap[key] = val
				}
			}
		}
		return res
	}

	// specToProto maps from KRM spec (*sampleKRMSpec) to GCP desired state (*descriptorpb.FileDescriptorProto):
	//   - in.StrField                    -> res.Name
	//   - in.NonPtrStr                   -> res.Package
	//   - in.BoolField                   -> res.Options.Deprecated
	//   - in.SubStruct.StrVal            -> res.Service[0].Name
	//   - in.StringSlice[i]              -> res.MessageType[i].Name (top-level repeated slice entry)
	//   - in.StringMap                   -> res.Options.UninterpretedOption
	specToProto := func(mapCtx *direct.MapContext, in *sampleKRMSpec) *descriptorpb.FileDescriptorProto {
		if in == nil {
			return nil
		}
		res := &descriptorpb.FileDescriptorProto{
			Name:    in.StrField,
			Package: stringPtr(in.NonPtrStr),
		}
		if in.SubStruct != nil && in.SubStruct.StrVal != nil {
			res.Service = []*descriptorpb.ServiceDescriptorProto{
				{Name: in.SubStruct.StrVal},
			}
		}
		if len(in.StringSlice) > 0 {
			for _, s := range in.StringSlice {
				res.MessageType = append(res.MessageType, &descriptorpb.DescriptorProto{
					Name: stringPtr(s),
				})
			}
		}
		if in.BoolField != nil || len(in.StringMap) > 0 {
			res.Options = &descriptorpb.FileOptions{}
			if in.BoolField != nil {
				res.Options.Deprecated = in.BoolField
			}
			for k, v := range in.StringMap {
				res.Options.UninterpretedOption = append(res.Options.UninterpretedOption, &descriptorpb.UninterpretedOption{
					Name: []*descriptorpb.UninterpretedOption_NamePart{
						{NamePart: stringPtr(k)},
					},
					StringValue: []byte(v),
				})
			}
		}
		return res
	}

	tests := []struct {
		name               string
		desiredKRM         *sampleKRMSpec
		actualProto        *descriptorpb.FileDescriptorProto
		specFromProtoErr   bool
		normalize          func(ctx context.Context, pb *descriptorpb.FileDescriptorProto) error
		wantHasDiff        bool
		wantFieldMaskPaths sets.Set[string]
		wantErr            bool
	}{
		{
			name: "unspecified KRM fields adopt live GCP state across primitive, nested struct, map, and top-level slice fields (no diff)",
			desiredKRM: &sampleKRMSpec{
				StrField:  stringPtr("my-resource"),
				NonPtrStr: "non-ptr-value",
				// SubStruct, StringMap, StringSlice left nil by user
			},
			actualProto: &descriptorpb.FileDescriptorProto{
				Name:    stringPtr("my-resource"),   // maps to StrField
				Package: stringPtr("non-ptr-value"), // maps to NonPtrStr
				Service: []*descriptorpb.ServiceDescriptorProto{
					{Name: stringPtr("gcp-service-name")}, // maps to SubStruct.StrVal
				},
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: stringPtr("Alpha")}, // maps to StringSlice (top-level repeated slice)
					{Name: stringPtr("Zebra")},
				},
				Options: &descriptorpb.FileOptions{
					Deprecated: boolPtr(true), // maps to BoolField
					UninterpretedOption: []*descriptorpb.UninterpretedOption{
						{
							Name: []*descriptorpb.UninterpretedOption_NamePart{
								{NamePart: stringPtr("env")}, // maps to StringMap["env"]
							},
							StringValue: []byte("prod"),
						},
					},
				},
			},
			normalize:          nil, // SortRepeatedFields runs automatically in CompareBrownfieldSpec Step 4
			wantHasDiff:        false,
			wantFieldMaskPaths: sets.New[string](),
		},
		{
			name: "user updates nested struct field in KRM, producing a diff on that field",
			desiredKRM: &sampleKRMSpec{
				StrField: stringPtr("my-resource"),
				SubStruct: &sampleKRMSub{
					StrVal: stringPtr("user-updated-service-name"), // Changed from live GCP state ("old-gcp-service-name")
				},
			},
			actualProto: &descriptorpb.FileDescriptorProto{
				Name: stringPtr("my-resource"), // maps to StrField (unchanged)
				Service: []*descriptorpb.ServiceDescriptorProto{
					{Name: stringPtr("old-gcp-service-name")}, // maps to SubStruct.StrVal (differs: "old-gcp-service-name" vs "user-updated-service-name")
				},
			},
			normalize:          nil, // Callers pass nil when no custom normalization is needed
			wantHasDiff:        true,
			wantFieldMaskPaths: sets.New("service"),
		},
		{
			name: "normalize hook normalizes specific field values (e.g. lowercasing strings), preventing false diffs",
			desiredKRM: &sampleKRMSpec{
				StrField: stringPtr("MY-RESOURCE-NAME"), // Upper-case in KRM
			},
			actualProto: &descriptorpb.FileDescriptorProto{
				Name: stringPtr("my-resource-name"), // maps to StrField (lowercased during normalize hook)
			},
			normalize: func(ctx context.Context, pb *descriptorpb.FileDescriptorProto) error {
				// Normalize hook active: lowercases field values in desired and actual proto to handle case-insensitivity
				if pb.Name != nil {
					lower := strings.ToLower(*pb.Name)
					pb.Name = &lower
				}
				return nil
			},
			wantHasDiff:        false,
			wantFieldMaskPaths: sets.New[string](),
		},
		{
			name: "automatically sorts top-level slice entries when desired KRM slice order differs from actual GCP proto order (no diff)",
			desiredKRM: &sampleKRMSpec{
				StrField:    stringPtr("my-resource"),
				StringSlice: []string{"Zebra", "Alpha"}, // Unsorted order in KRM
			},
			actualProto: &descriptorpb.FileDescriptorProto{
				Name: stringPtr("my-resource"), // maps to StrField
				MessageType: []*descriptorpb.DescriptorProto{ // maps to StringSlice (top-level repeated slice)
					{Name: stringPtr("Alpha")},
					{Name: stringPtr("Zebra")},
				},
			},
			normalize:          nil, // CompareBrownfieldSpec automatically sorts repeated fields without needing a custom normalize func
			wantHasDiff:        false,
			wantFieldMaskPaths: sets.New[string](),
		},
		{
			name: "specFromProto error is reported cleanly",
			desiredKRM: &sampleKRMSpec{
				StrField: stringPtr("my-resource"),
			},
			actualProto: &descriptorpb.FileDescriptorProto{
				Name: stringPtr("my-resource"), // maps to StrField
			},
			specFromProtoErr: true,
			normalize:        nil,
			wantErr:          true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fromProto := func(mapCtx *direct.MapContext, in *descriptorpb.FileDescriptorProto) *sampleKRMSpec {
				if tc.specFromProtoErr {
					mapCtx.Errorf("simulated specFromProto mapping failure")
					return nil
				}
				return specFromProto(mapCtx, in)
			}

			toProto := func(mapCtx *direct.MapContext, in *sampleKRMSpec) *descriptorpb.FileDescriptorProto {
				return specToProto(mapCtx, in)
			}

			diff, fieldMask, err := CompareBrownfieldSpec(
				t.Context(),
				tc.desiredKRM,
				tc.actualProto,
				fromProto,
				toProto,
				tc.normalize,
			)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.wantHasDiff {
				if diff == nil || !diff.HasDiff() {
					t.Errorf("expected diff, got no diff")
				}
			} else {
				if diff != nil && diff.HasDiff() {
					t.Errorf("expected no diff, got diff:\n%+v", diff)
				}
			}

			if fieldMask == nil {
				t.Fatalf("expected non-nil fieldMask")
			}

			if tc.wantFieldMaskPaths != nil {
				paths := sets.New(fieldMask.GetPaths()...)
				if !paths.Equal(tc.wantFieldMaskPaths) {
					t.Errorf("fieldmask paths mismatch: want %v, got %v", tc.wantFieldMaskPaths, paths)
				}
			}
		})
	}
}
