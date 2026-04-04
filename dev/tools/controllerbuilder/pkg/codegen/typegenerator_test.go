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
	"testing"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestGoTypeForField_Map(t *testing.T) {
	// Define google.protobuf.Timestamp
	googleProtoDescriptor := &descriptorpb.FileDescriptorProto{
		Name:    directPtr("google/protobuf/timestamp.proto"),
		Package: directPtr("google.protobuf"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: directPtr("Timestamp"),
			},
		},
	}

	// Define a minimal FileDescriptor containing a message with a map field
	fileDescriptorProto := &descriptorpb.FileDescriptorProto{
		Name:    directPtr("test.proto"),
		Package: directPtr("test"),
		Syntax:  directPtr("proto3"),
		Dependency: []string{"google/protobuf/timestamp.proto"},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: directPtr("ValueMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   directPtr("name"),
						Number: directInt32(1),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
			{
				Name: directPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   directPtr("labels"),
						Number: directInt32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: directPtr(".test.TestMessage.LabelsEntry"),
					},
					{
						Name:   directPtr("message_map"),
						Number: directInt32(2),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: directPtr(".test.TestMessage.MessageMapEntry"),
					},
					{
						Name:   directPtr("timestamp_map"),
						Number: directInt32(3),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: directPtr(".test.TestMessage.TimestampMapEntry"),
					},
					{
						Name:   directPtr("int64_map"),
						Number: directInt32(4),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: directPtr(".test.TestMessage.Int64MapEntry"),
					},
					{
						Name:   directPtr("enum_map"),
						Number: directInt32(5),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
						TypeName: directPtr(".test.TestMessage.EnumMapEntry"),
					},
				},
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: directPtr("LabelsEntry"),
						Options: &descriptorpb.MessageOptions{
							MapEntry: directBool(true),
						},
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   directPtr("key"),
								Number: directInt32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:   directPtr("value"),
								Number: directInt32(2),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
						},
					},
					{
						Name: directPtr("MessageMapEntry"),
						Options: &descriptorpb.MessageOptions{
							MapEntry: directBool(true),
						},
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   directPtr("key"),
								Number: directInt32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:   directPtr("value"),
								Number: directInt32(2),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: directPtr(".test.ValueMessage"),
							},
						},
					},
					{
						Name: directPtr("TimestampMapEntry"),
						Options: &descriptorpb.MessageOptions{
							MapEntry: directBool(true),
						},
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   directPtr("key"),
								Number: directInt32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:   directPtr("value"),
								Number: directInt32(2),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: directPtr(".google.protobuf.Timestamp"),
							},
						},
					},
					{
						Name: directPtr("Int64MapEntry"),
						Options: &descriptorpb.MessageOptions{
							MapEntry: directBool(true),
						},
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   directPtr("key"),
								Number: directInt32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:   directPtr("value"),
								Number: directInt32(2),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
							},
						},
					},
					{
						Name: directPtr("EnumMapEntry"),
						Options: &descriptorpb.MessageOptions{
							MapEntry: directBool(true),
						},
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   directPtr("key"),
								Number: directInt32(1),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							},
							{
								Name:   directPtr("value"),
								Number: directInt32(2),
								Type:   descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
								TypeName: directPtr(".test.TestEnum"),
							},
						},
					},
				},
			},
		},
		EnumType: []*descriptorpb.EnumDescriptorProto{
			{
				Name: directPtr("TestEnum"),
				Value: []*descriptorpb.EnumValueDescriptorProto{
					{
						Name:   directPtr("UNKNOWN"),
						Number: directInt32(0),
					},
				},
			},
		},
	}

	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{
			googleProtoDescriptor,
			fileDescriptorProto,
		},
	})
	if err != nil {
		t.Fatalf("Failed to create FileDescriptor: %v", err)
	}
	testMessageDesc, err := files.FindDescriptorByName("test.TestMessage")
	if err != nil {
		t.Fatalf("Failed to find TestMessage: %v", err)
	}
	testMessage := testMessageDesc.(protoreflect.MessageDescriptor)

	tests := []struct {
		fieldName string
		expected  string
	}{
		{
			fieldName: "labels",
			expected:  "map[string]string",
		},
		{
			fieldName: "message_map",
			expected:  "map[string]*ValueMessage",
		},
		{
			fieldName: "timestamp_map",
			expected:  "map[string]string",
		},
		{
			fieldName: "int64_map",
			expected:  "map[string]int64",
		},
		{
			fieldName: "enum_map",
			expected:  "map[string]string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.fieldName, func(t *testing.T) {
			field := testMessage.Fields().ByName(protoreflect.Name(tt.fieldName))
			if field == nil {
				t.Fatalf("Field %s not found", tt.fieldName)
			}
			got, err := GoTypeForField(field, false)
			if err != nil {
				t.Fatalf("GoTypeForField failed: %v", err)
			}
			if got != tt.expected {
				t.Errorf("GoTypeForField() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func directPtr[T any](v T) *T {
	return &v
}

func directInt32(v int32) *int32 {
	return &v
}

func directBool(v bool) *bool {
	return &v
}
