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
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/apimachinery/pkg/util/sets"
)

func protoMapToMap(m protoreflect.Map) map[string]interface{} {
	res := make(map[string]interface{})
	m.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		res[k.String()] = v.Message().Interface()
		return true
	})
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
				if pm, ok := got.Old.(protoreflect.Map); ok {
					if diff := cmp.Diff(d.Old, protoMapToMap(pm), protocmp.Transform()); diff != "" {
						t.Errorf("diff[%d].Old mismatch (-want +got):\n%s", i, diff)
					}
				} else {
					if diff := cmp.Diff(d.Old, got.Old, protocmp.Transform()); diff != "" {
						t.Errorf("diff[%d].Old mismatch (-want +got):\n%s", i, diff)
					}
				}
				if pm, ok := got.New.(protoreflect.Map); ok {
					if diff := cmp.Diff(d.New, protoMapToMap(pm), protocmp.Transform()); diff != "" {
						t.Errorf("diff[%d].New mismatch (-want +got):\n%s", i, diff)
					}
				} else {
					if diff := cmp.Diff(d.New, got.New, protocmp.Transform()); diff != "" {
						t.Errorf("diff[%d].New mismatch (-want +got):\n%s", i, diff)
					}
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
