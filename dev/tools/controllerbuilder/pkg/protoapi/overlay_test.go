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

package protoapi

import (
	"bytes"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestApplyOverlayFromReader(t *testing.T) {
	overlay := `
package google.cloud.batch.v1;

message Task {
  TaskStatus status = 2 [(google.api.field_behavior) = OUTPUT_ONLY];
  
  message Inner {
      map<string, string> labels = 1 [(google.api.field_behavior) = OPTIONAL];
  }
}
`
	fds := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{
			{
				Package: proto.String("google.cloud.batch.v1"),
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("Task"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name: proto.String("status"),
							},
						},
						NestedType: []*descriptorpb.DescriptorProto{
							{
								Name: proto.String("Inner"),
								Field: []*descriptorpb.FieldDescriptorProto{
									{
										Name: proto.String("labels"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	err := applyOverlayFromReader(fds, bytes.NewReader([]byte(overlay)))
	if err != nil {
		t.Fatalf("applyOverlayFromReader failed: %v", err)
	}

	field := fds.File[0].MessageType[0].Field[0]
	if field.Options == nil {
		t.Fatalf("field.Options is nil")
	}

	ext := proto.GetExtension(field.Options, annotations.E_FieldBehavior)
	behaviors, ok := ext.([]annotations.FieldBehavior)
	if !ok || len(behaviors) != 1 || behaviors[0] != annotations.FieldBehavior_OUTPUT_ONLY {
		t.Fatalf("expected OUTPUT_ONLY, got %v", ext)
	}

	innerField := fds.File[0].MessageType[0].NestedType[0].Field[0]
	if innerField.Options == nil {
		t.Fatalf("innerField.Options is nil")
	}

	ext2 := proto.GetExtension(innerField.Options, annotations.E_FieldBehavior)
	behaviors2, ok := ext2.([]annotations.FieldBehavior)
	if !ok || len(behaviors2) != 1 || behaviors2[0] != annotations.FieldBehavior_OPTIONAL {
		t.Fatalf("expected OPTIONAL, got %v", ext2)
	}
}
