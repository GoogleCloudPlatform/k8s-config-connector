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
	"testing"

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
