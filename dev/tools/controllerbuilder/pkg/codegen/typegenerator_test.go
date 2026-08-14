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
	"bytes"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/util/sets"
)

func TestAsSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"CamelCase", "camel_case"},
		{"camelCase", "camel_case"},
		{"Snake_Case", "snake__case"}, // Note: current implementation produces double underscores
		{"snake_case", "snake_case"},
		{"ID", "id"},
		{"HTTPResponse", "http_response"},
	}

	for _, tt := range tests {
		if got := AsSnakeCase(tt.input); got != tt.expected {
			t.Errorf("AsSnakeCase(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestIdentifyOutputs(t *testing.T) {
	// Setup FieldBehavior options
	outputOnly := []annotations.FieldBehavior{annotations.FieldBehavior_OUTPUT_ONLY}
	outputOnlyOptions := &descriptorpb.FieldOptions{}
	proto.SetExtension(outputOnlyOptions, annotations.E_FieldBehavior, outputOnly)

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("SubMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:    protoPtr("output_field"),
						Number:  protoPtr(int32(1)),
						Type:    typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: outputOnlyOptions,
					},
					{
						Name:   protoPtr("input_field"),
						Number: protoPtr(int32(2)),
						Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					},
				},
			},
			{
				Name: protoPtr("RootMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     protoPtr("sub_message"),
						Number:   protoPtr(int32(1)),
						Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
						TypeName: protoPtr(".google.cloud.test.v1.SubMessage"),
					},
					{
						Name:   protoPtr("regular_field"),
						Number: protoPtr(int32(2)),
						Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	rootMsg := fd.Messages().ByName("RootMessage")
	g := &TypeGenerator{}

	outputDeps := make(map[string]*OutputMessageDetails)
	hasOutputs := g.identifyOutputs(rootMsg, make(map[string]string), outputDeps, false)

	if !hasOutputs {
		t.Errorf("expected RootMessage to have outputs")
	}

	if _, ok := outputDeps["google.cloud.test.v1.RootMessage"]; !ok {
		t.Errorf("expected RootMessage in outputDeps")
	}

	if _, ok := outputDeps["google.cloud.test.v1.SubMessage"]; !ok {
		t.Errorf("expected SubMessage in outputDeps")
	}

	rootDetails := outputDeps["google.cloud.test.v1.RootMessage"]
	if len(rootDetails.OutputFields) != 1 || string(rootDetails.OutputFields[0].Name()) != "sub_message" {
		t.Errorf("expected root output field to be 'sub_message', got %v", rootDetails.OutputFields)
	}

	subDetails := outputDeps["google.cloud.test.v1.SubMessage"]
	if len(subDetails.OutputFields) != 1 || string(subDetails.OutputFields[0].Name()) != "output_field" {
		t.Errorf("expected sub output field to be 'output_field', got %v", subDetails.OutputFields)
	}
}

func TestIsAcronym(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ID", true},
		{"id", true},
		{"Id", true},
		{"HTTP", true},
		{"NotAnAcronym", false},
	}

	for _, tt := range tests {
		if got := IsAcronym(tt.input); got != tt.expected {
			t.Errorf("IsAcronym(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestGoNameForProtoMessage(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TestMessage"),
			},
			{
				Name: protoPtr("HTTPResponse"),
			},
			{
				Name: protoPtr("ProjectID"),
			},
			{
				Name: protoPtr("NestedMessage"),
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: protoPtr("InnerMessage"),
					},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	tests := []struct {
		msgName  string
		expected string
	}{
		{"google.cloud.test.v1.TestMessage", "TestMessage"},
		{"google.cloud.test.v1.HTTPResponse", "HTTPResponse"},
		{"google.cloud.test.v1.ProjectID", "ProjectID"},
		{"google.cloud.test.v1.NestedMessage", "NestedMessage"},
		{"google.cloud.test.v1.NestedMessage.InnerMessage", "NestedMessage_InnerMessage"},
	}

	for _, tt := range tests {
		msg := fd.Messages().ByName(protoreflect.Name(lastPart(tt.msgName)))
		if tt.msgName == "google.cloud.test.v1.NestedMessage.InnerMessage" {
			msg = fd.Messages().ByName("NestedMessage").Messages().ByName("InnerMessage")
		}

		if msg == nil {
			t.Errorf("could not find message %q", tt.msgName)
			continue
		}

		if got := GoNameForProtoMessage(msg); got != tt.expected {
			t.Errorf("GoNameForProtoMessage(%q) = %q, want %q", tt.msgName, got, tt.expected)
		}
	}
}

func lastPart(s string) string {
	parts := strings.Split(s, ".")
	return parts[len(parts)-1]
}

func protoPtr[T any](v T) *T {
	return &v
}

func TestGoFieldName(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("project_id"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("display_name"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("http_header"), Number: protoPtr(int32(3)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msg := fd.Messages().ByName("TestMessage")
	fields := msg.Fields()

	tests := []struct {
		fieldName string
		expected  string
	}{
		{"project_id", "ProjectID"},
		{"display_name", "DisplayName"},
		{"http_header", "HTTPHeader"},
	}

	for _, tt := range tests {
		field := fields.ByName(protoreflect.Name(tt.fieldName))
		if field == nil {
			t.Errorf("could not find field %q", tt.fieldName)
			continue
		}

		if got := goFieldName(field); got != tt.expected {
			t.Errorf("goFieldName(%q) = %q, want %q", tt.fieldName, got, tt.expected)
		}
	}
}

func TestGetJSONForKRM(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("project_id"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("display_name"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("http_header"), Number: protoPtr(int32(3)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msg := fd.Messages().ByName("TestMessage")
	fields := msg.Fields()

	tests := []struct {
		fieldName string
		expected  string
	}{
		{"project_id", "projectID"},
		{"display_name", "displayName"},
		{"http_header", "httpHeader"},
	}

	for _, tt := range tests {
		field := fields.ByName(protoreflect.Name(tt.fieldName))
		if field == nil {
			t.Errorf("could not find field %q", tt.fieldName)
			continue
		}

		if got := GetJSONForKRM(field); got != tt.expected {
			t.Errorf("GetJSONForKRM(%q) = %q, want %q", tt.fieldName, got, tt.expected)
		}
	}
}

func TestGoTypeForField(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TargetMessage"),
			},
			{
				Name: protoPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("string_field"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("int64_field"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_INT64)},
					{Name: protoPtr("bool_field"), Number: protoPtr(int32(3)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_BOOL)},
					{Name: protoPtr("bytes_field"), Number: protoPtr(int32(4)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_BYTES)},
					{Name: protoPtr("repeated_string"), Number: protoPtr(int32(5)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING), Label: labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED)},
					{Name: protoPtr("message_field"), Number: protoPtr(int32(6)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: protoPtr(".google.cloud.test.v1.TargetMessage")},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msg := fd.Messages().ByName("TestMessage")
	fields := msg.Fields()

	tests := []struct {
		fieldName          string
		isTransitiveOutput bool
		expected           string
	}{
		{"string_field", false, "*string"},
		{"int64_field", false, "*int64"},
		{"bool_field", false, "*bool"},
		{"bytes_field", false, "[]byte"},
		{"repeated_string", false, "[]string"},
		{"message_field", false, "*TargetMessage"},
		{"message_field", true, "*TargetMessageObservedState"},
	}

	for _, tt := range tests {
		field := fields.ByName(protoreflect.Name(tt.fieldName))
		if field == nil {
			t.Fatalf("could not find field %q", tt.fieldName)
		}

		got, err := GoTypeForField(field, tt.isTransitiveOutput)
		if err != nil {
			t.Errorf("GoTypeForField(%q, %v) returned error: %v", tt.fieldName, tt.isTransitiveOutput, err)
			continue
		}

		if got != tt.expected {
			t.Errorf("GoTypeForField(%q, %v) = %q, want %q", tt.fieldName, tt.isTransitiveOutput, got, tt.expected)
		}
	}
}

func TestFindDependenciesForMessage(t *testing.T) {
	fdpTS := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("google/protobuf/timestamp.proto"),
		Package: protoPtr("google.protobuf"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("Timestamp")},
		},
	}
	fdp := &descriptorpb.FileDescriptorProto{
		Name:       protoPtr("test.proto"),
		Package:    protoPtr("google.cloud.test.v1"),
		Dependency: []string{"google/protobuf/timestamp.proto"},
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("DepMessage"),
			},
			{
				Name: protoPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("dep_field"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: protoPtr(".google.cloud.test.v1.DepMessage")},
					{Name: protoPtr("ts_field"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: protoPtr(".google.protobuf.Timestamp")},
				},
			},
		},
	}

	fds := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{fdpTS, fdp},
	}
	files, err := protodesc.NewFiles(fds)
	if err != nil {
		t.Fatalf("failed to create file descriptors: %v", err)
	}

	descriptor, err := files.FindDescriptorByName("google.cloud.test.v1.TestMessage")
	if err != nil {
		t.Fatalf("failed to find TestMessage: %v", err)
	}
	msg := descriptor.(protoreflect.MessageDescriptor)

	deps, err := FindDependenciesForMessage(msg, sets.NewString())
	if err != nil {
		t.Fatalf("FindDependenciesForMessage failed: %v", err)
	}

	// Should only find DepMessage. Timestamp should be skipped.
	if len(deps) != 1 {
		t.Errorf("expected 1 dependency, got %d: %v", len(deps), deps)
	} else if string(deps[0].FullName()) != "google.cloud.test.v1.DepMessage" {
		t.Errorf("expected DepMessage, got %s", deps[0].FullName())
	}
}

func TestDeduplicateAndSort(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("B")},
			{Name: protoPtr("A")},
			{Name: protoPtr("C")},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msgs := []protoreflect.MessageDescriptor{
		fd.Messages().ByName("B"),
		fd.Messages().ByName("A"),
		fd.Messages().ByName("C"),
		fd.Messages().ByName("A"), // Duplicate
	}

	sorted := deduplicateAndSort(msgs)
	if len(sorted) != 3 {
		t.Errorf("expected 3 messages, got %d", len(sorted))
	}

	if string(sorted[0].Name()) != "A" || string(sorted[1].Name()) != "B" || string(sorted[2].Name()) != "C" {
		t.Errorf("unexpected sort order: %v, %v, %v", sorted[0].Name(), sorted[1].Name(), sorted[2].Name())
	}
}

func TestWriteMessage(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TestMessage"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   protoPtr("project_id"),
						Number: protoPtr(int32(1)),
						Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					},
				},
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msg := fd.Messages().ByName("TestMessage")
	var buf bytes.Buffer
	WriteMessage(&buf, msg)

	got := buf.String()
	expected := "\n// +kcc:proto=google.cloud.test.v1.TestMessage\ntype TestMessage struct {\n\t// +kcc:proto:field=google.cloud.test.v1.TestMessage.project_id\n\tProjectID *string `json:\"projectID,omitempty\"`\n}\n"

	if got != expected {
		t.Errorf("WriteMessage output mismatch.\nGot:\n%q\nWant:\n%q", got, expected)
	}
}

func labelDescriptor(l descriptorpb.FieldDescriptorProto_Label) *descriptorpb.FieldDescriptorProto_Label {
	return &l
}

func typeDescriptor(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}
