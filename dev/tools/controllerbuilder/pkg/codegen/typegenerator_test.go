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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
					{Name: protoPtr("related_uris"), Number: protoPtr(int32(4)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
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
	acronyms := WriteOptions{EmitPluralAcronyms: true}

	// The flag-on case is the one the judgement queue relies on; see
	// GetJSONForKRM.
	tests := []struct {
		fieldName string
		opts      WriteOptions
		expected  string
	}{
		{"project_id", WriteOptions{}, "projectID"},
		{"display_name", WriteOptions{}, "displayName"},
		{"http_header", WriteOptions{}, "httpHeader"},
		{"related_uris", WriteOptions{}, "relatedUris"},
		{"related_uris", acronyms, "relatedURIs"},
	}

	for _, tt := range tests {
		field := fields.ByName(protoreflect.Name(tt.fieldName))
		if field == nil {
			t.Errorf("could not find field %q", tt.fieldName)
			continue
		}

		if got := GetJSONForKRM(field, tt.opts); got != tt.expected {
			t.Errorf("GetJSONForKRM(%q, %+v) = %q, want %q", tt.fieldName, tt.opts, got, tt.expected)
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

		got, err := GoTypeForField(field, tt.isTransitiveOutput, WriteOptions{})
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

func TestWriteFieldRequiredMarker(t *testing.T) {
	grid := []struct {
		name       string
		behaviors  []annotations.FieldBehavior
		opts       WriteOptions
		wantMarker bool
	}{
		{
			name:       "REQUIRED with the flag on emits the marker",
			behaviors:  []annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED},
			opts:       WriteOptions{EmitRequired: true},
			wantMarker: true,
		},
		{
			// Services that have not opted in have to keep generating byte-identical
			// output, which is the whole reason for the flag.
			name:      "REQUIRED with the flag off emits nothing",
			behaviors: []annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED},
			opts:      WriteOptions{},
		},
		{
			name:      "OPTIONAL never emits the marker",
			behaviors: []annotations.FieldBehavior{annotations.FieldBehavior_OPTIONAL},
			opts:      WriteOptions{EmitRequired: true},
		},
		{
			// About half of proto fields carry no field_behavior at all; they must stay
			// optional, which is what omitempty already gives us.
			name:      "no field_behavior stays optional",
			behaviors: nil,
			opts:      WriteOptions{EmitRequired: true},
		},
		{
			// An output field is never user-supplied, and the API server validates
			// status, so requiring it would let GCP's response fail our own schema.
			name: "OUTPUT_ONLY wins over REQUIRED",
			behaviors: []annotations.FieldBehavior{
				annotations.FieldBehavior_REQUIRED,
				annotations.FieldBehavior_OUTPUT_ONLY,
			},
			opts: WriteOptions{EmitRequired: true},
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			fieldOpts := &descriptorpb.FieldOptions{}
			if g.behaviors != nil {
				proto.SetExtension(fieldOpts, annotations.E_FieldBehavior, g.behaviors)
			}

			fdp := &descriptorpb.FileDescriptorProto{
				Name:    protoPtr("test.proto"),
				Package: protoPtr("google.cloud.test.v1"),
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: protoPtr("TestMessage"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:    protoPtr("project_id"),
								Number:  protoPtr(int32(1)),
								Type:    typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
								Options: fieldOpts,
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
			WriteField(&buf, msg.Fields().Get(0), msg, 0, false, g.opts, "")

			got := strings.Contains(buf.String(), "// +required")
			if got != g.wantMarker {
				t.Errorf("+required present = %v, want %v\noutput:\n%s", got, g.wantMarker, buf.String())
			}
		})
	}
}

// TestWriteMessage checks the rendered output rather than type strings,
// because a field GoTypeForField cannot type does not fail generation.
// WriteField writes a "// TODO:" comment in its place and the field is missing
// from the CRD. The expected output below has both kinds of field.
func TestWriteMessage(t *testing.T) {
	mapEntry := func(name string, value *descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
		return &descriptorpb.DescriptorProto{
			Name: protoPtr(name),
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
				value,
			},
			Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
		}
	}
	mapField := func(name string, num int32, entry string) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: protoPtr(name), Number: protoPtr(num),
			Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
			TypeName: protoPtr(".google.cloud.test.v1.TestMessage." + entry),
			Label:    labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED),
		}
	}

	fdp := &descriptorpb.FileDescriptorProto{
		Name:       protoPtr("test.proto"),
		Package:    protoPtr("google.cloud.test.v1"),
		Dependency: []string{"google/protobuf/timestamp.proto"},
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("TargetMessage")},
			{
				Name: protoPtr("TestMessage"),
				NestedType: []*descriptorpb.DescriptorProto{
					mapEntry("LabelsEntry", &descriptorpb.FieldDescriptorProto{
						Name: protoPtr("value"), Number: protoPtr(int32(2)),
						Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					}),
					mapEntry("TasksEntry", &descriptorpb.FieldDescriptorProto{
						Name: protoPtr("value"), Number: protoPtr(int32(2)),
						Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
						TypeName: protoPtr(".google.cloud.test.v1.TargetMessage"),
					}),
					mapEntry("SeenEntry", &descriptorpb.FieldDescriptorProto{
						Name: protoPtr("value"), Number: protoPtr(int32(2)),
						Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
						TypeName: protoPtr(".google.protobuf.Timestamp"),
					}),
					{
						Name: protoPtr("ByIndexEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_INT32)},
							{Name: protoPtr("value"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
						},
						Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
					},
				},
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   protoPtr("project_id"),
						Number: protoPtr(int32(1)),
						Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					},
					mapField("labels", 2, "LabelsEntry"),
					mapField("tasks", 3, "TasksEntry"),
					mapField("seen", 4, "SeenEntry"),
					mapField("by_index", 5, "ByIndexEntry"),
				},
			},
		},
	}

	deps := new(protoregistry.Files)
	if err := deps.RegisterFile(timestamppb.Now().ProtoReflect().Descriptor().ParentFile()); err != nil {
		t.Fatalf("registering timestamp.proto: %v", err)
	}
	fd, err := protodesc.NewFile(fdp, deps)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}

	msg := fd.Messages().ByName("TestMessage")
	var buf bytes.Buffer
	// Prepopulating is set because the marker only names its field for a
	// service that opted in, and the assertions below tell the fields apart by
	// name. TestWriteFieldMarkerIsScopedToPrepopulating pins that scoping.
	WriteMessage(&buf, msg, WriteOptions{EmitMessageMaps: true, Prepopulating: true})

	got := buf.String()
	// labels, tasks and seen cover the supported map values: a scalar, a
	// message, and a message with a special-cased Go type. by_index is still
	// left out, because a CRD keys additionalProperties by string, and its
	// TODO line is the only sign of it.
	expected := strings.Join([]string{
		"",
		"// +kcc:proto=google.cloud.test.v1.TestMessage",
		"type TestMessage struct {",
		"\t// +kcc:proto:field=google.cloud.test.v1.TestMessage.project_id",
		"\tProjectID *string `json:\"projectID,omitempty\"`",
		"",
		"\t// +kcc:proto:field=google.cloud.test.v1.TestMessage.labels",
		"\tLabels map[string]string `json:\"labels,omitempty\"`",
		"",
		"\t// +kcc:proto:field=google.cloud.test.v1.TestMessage.tasks",
		"\tTasks map[string]TargetMessage `json:\"tasks,omitempty\"`",
		"",
		"\t// +kcc:proto:field=google.cloud.test.v1.TestMessage.seen",
		"\tSeen map[string]string `json:\"seen,omitempty\"`",
		"",
		"\t// TODO: byIndex: unsupported map type with key int32 and value string",
		"",
		"}",
		"",
	}, "\n")

	if got != expected {
		t.Errorf("WriteMessage output mismatch.\nGot:\n%q\nWant:\n%q", got, expected)
	}

	// With EmitMessageMaps off, message-valued maps are left out as they were
	// before the flag existed, and a "// TODO:" marker names each one.
	var off bytes.Buffer
	WriteMessage(&off, msg, WriteOptions{Prepopulating: true})
	for _, want := range []string{
		"\t// TODO: tasks: unsupported map type with key string and value message",
		"\t// TODO: seen: unsupported map type with key string and value message",
	} {
		if !strings.Contains(off.String(), want) {
			t.Errorf("with EmitMessageMaps off, want %q in:\n%s", want, off.String())
		}
	}
	if strings.Contains(off.String(), "Tasks map[string]") {
		t.Errorf("with EmitMessageMaps off, Tasks should not be generated:\n%s", off.String())
	}
}

func labelDescriptor(l descriptorpb.FieldDescriptorProto_Label) *descriptorpb.FieldDescriptorProto_Label {
	return &l
}

func typeDescriptor(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

// TestNeedsObservedStateRequired covers the message shape that needs an ObservedState struct of its
// own, and the two hand-written shapes that have to be left alone.
//
// SharedConfig has a REQUIRED field and no OUTPUT_ONLY field anywhere. Without the split it
// deduplicates down to a single struct shared by the spec and the observed state, and WriteMessage
// stamps "// +required" on it, which lands required: in the status schema. redis PscConfig is a
// real case of exactly this.
//
// Splitting is safe only when nothing hand-written claims the message. There are two ways that
// goes wrong: a hand-written plain type (aiplatform FunctionDeclaration) and a hand-written
// ObservedState type (dataplex DataQualityDimensionResult).
func TestNeedsObservedStateRequired(t *testing.T) {
	requiredOptions := &descriptorpb.FieldOptions{}
	proto.SetExtension(requiredOptions, annotations.E_FieldBehavior,
		[]annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED})

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("SharedConfig"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:    protoPtr("key"),
						Number:  protoPtr(int32(1)),
						Type:    typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: requiredOptions,
					},
					{
						Name:   protoPtr("value"),
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
	shared := fd.Messages().ByName("SharedConfig")
	fqn := "google.cloud.test.v1.SharedConfig"

	empty := &nonAutogeneratedTypes{typeNames: map[string]bool{}, protoTagged: map[string]bool{}}

	grid := []struct {
		name             string
		emitRequired     bool
		nonAutogenerated *nonAutogeneratedTypes
		want             bool
	}{
		{
			name:             "generator owns the message: split it",
			emitRequired:     true,
			nonAutogenerated: empty,
			want:             true,
		},
		{
			name:             "flag off: do not split message (preserve deduplication)",
			emitRequired:     false,
			nonAutogenerated: empty,
			// Splitting with the flag off would change generated output for every service that
			// has not opted into --emit-required-from-proto.
			want: false,
		},
		{
			name:             "non-autogenerated plain type claims the message: do not split",
			emitRequired:     true,
			nonAutogenerated: &nonAutogeneratedTypes{
				typeNames:   map[string]bool{"SharedConfig": true},
				protoTagged: map[string]bool{fqn: true},
			},
			// E.g. aiplatform's FunctionDeclaration. The generator emits no struct here because
			// the package already declares it, so splitting would create an undefined reference.
			want: false,
		},
		{
			name:             "non-autogenerated ObservedState type claims the message: do not split",
			emitRequired:     true,
			nonAutogenerated: &nonAutogeneratedTypes{
				typeNames:   map[string]bool{"SharedConfigObservedState": true},
				protoTagged: map[string]bool{fqn: true},
			},
			// E.g. dataplex's DataQualityDimensionResult, where *SharedConfigObservedState
			// is already declared in the package.
			want: false,
		},
	}

	for _, tc := range grid {
		t.Run(tc.name, func(t *testing.T) {
			g := &TypeGenerator{}
			g.writeOptions = WriteOptions{EmitRequired: tc.emitRequired}
			g.nonAutogenerated = tc.nonAutogenerated
			if got := g.needsObservedState(shared, make(map[string]bool)); got != tc.want {
				t.Errorf("needsObservedState() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestScanNonAutogeneratedTypes verifies that scanNonAutogeneratedTypes accurately records
// both Go type declarations and proto message annotations (+kcc:*proto) across non-autogenerated files.
func TestScanNonAutogeneratedTypes(t *testing.T) {
	dir := t.TempDir()

	// Non-autogenerated: one plain type, one ObservedState type for a different message.
	nonAutogenerated := `package v1alpha1

// +kcc:proto=google.cloud.test.v1.FunctionDeclaration
type FunctionDeclaration struct {
	Name *string ` + "`json:\"name,omitempty\"`" + `
}

// +kcc:observedstate:proto=google.cloud.test.v1.DimensionResult
type DimensionResultObservedState struct {
}
`
	// Generated files have to be ignored, because everything in them is ours to reshape.
	generated := `package v1alpha1

// +kcc:proto=google.cloud.test.v1.PscConfig
type PSCConfig struct {
}
`
	if err := os.WriteFile(filepath.Join(dir, "thing_types.go"), []byte(nonAutogenerated), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "types.generated.go"), []byte(generated), 0o644); err != nil {
		t.Fatal(err)
	}

	g := &generatorBase{}
	g.init(dir)
	got, err := g.scanNonAutogeneratedTypes(dir)
	if err != nil {
		t.Fatalf("scanNonAutogeneratedTypes: %v", err)
	}

	for _, name := range []string{"FunctionDeclaration", "DimensionResultObservedState"} {
		if !got.typeNames[name] {
			t.Errorf("expected non-autogenerated type %q to be recorded", name)
		}
	}
	if got.typeNames["PSCConfig"] {
		t.Error("PSCConfig comes from types.generated.go and must not count as non-autogenerated")
	}
	if got.protoTagged["google.cloud.test.v1.PscConfig"] {
		t.Error("a generated file's proto tag must not mark the message as claimed")
	}

	if got.ownedByGenerator("google.cloud.test.v1.FunctionDeclaration", "FunctionDeclaration", "FunctionDeclarationObservedState") {
		t.Error("a message with a non-autogenerated plain type is not owned by the generator")
	}
	if got.ownedByGenerator("google.cloud.test.v1.DimensionResult", "DimensionResult", "DimensionResultObservedState") {
		t.Error("a message with a non-autogenerated ObservedState type is not owned by the generator")
	}
	if !got.ownedByGenerator("google.cloud.test.v1.PscConfig", "PSCConfig", "PSCConfigObservedState") {
		t.Error("a message only present in generated output is owned by the generator")
	}

	// A service generated for the first time has no package directory yet.
	missing, err := g.scanNonAutogeneratedTypes(filepath.Join(dir, "does-not-exist"))
	if err != nil {
		t.Fatalf("missing directory should not be an error: %v", err)
	}
	if !missing.ownedByGenerator("google.cloud.test.v1.Anything", "Anything", "AnythingObservedState") {
		t.Error("everything is generator-owned when the package does not exist yet")
	}
}

// TestReservedTypeNamesSkipsAndComments verifies that reserved type names are not emitted
// as active Go types, and are written as comments when includeSkippedOutput is enabled.
func TestReservedTypeNamesSkipsAndComments(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("BillingAccount"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("name"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	msg := fd.Messages().ByName("BillingAccount")
	goTypeName := "BillingAccount"

	for _, tc := range []struct {
		name                 string
		includeSkippedOutput bool
		reserved             bool
		wantComment          bool
		wantActive           bool
	}{
		{
			name:                 "not reserved, generated actively",
			includeSkippedOutput: false,
			reserved:             false,
			wantComment:          false,
			wantActive:           true,
		},
		{
			name:                 "reserved, skipped silently when includeSkippedOutput is false",
			includeSkippedOutput: false,
			reserved:             true,
			wantComment:          false,
			wantActive:           false,
		},
		{
			name:                 "reserved, generated as comment when includeSkippedOutput is true",
			includeSkippedOutput: true,
			reserved:             true,
			wantComment:          true,
			wantActive:           false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			pkgDir := filepath.Join(dir, "test")
			if err := os.MkdirAll(pkgDir, 0755); err != nil {
				t.Fatalf("MkdirAll: %v", err)
			}
			g := NewTypeGenerator("test", dir, nil)
			g.visitedMessages = []protoreflect.MessageDescriptor{msg}
			g.includeSkippedOutput = tc.includeSkippedOutput
			if tc.reserved {
				g.WithReservedTypeNames(goTypeName)
			}
			if err := g.WriteVisitedMessages(); err != nil {
				t.Fatalf("WriteVisitedMessages: %v", err)
			}
			f := g.generatedFiles[generatedFileKey{GoPackage: "test", FileName: "types.generated.go"}]
			body := ""
			if f != nil {
				body = f.body.String()
			}
			hasComment := strings.Contains(body, "/* go type "+`"`+goTypeName+`"`+" is a Kind the scaffolder will declare, skipping")
			hasActive := strings.Contains(body, "type "+goTypeName+" struct") && !hasComment
			if hasActive != tc.wantActive {
				t.Errorf("hasActive = %v, want %v; body = %s", hasActive, tc.wantActive, body)
			}
			if hasComment != tc.wantComment {
				t.Errorf("hasComment = %v, want %v; body = %s", hasComment, tc.wantComment, body)
			}
		})
	}
}

// TestWriteObservedStateFieldsNotes verifies that WriteObservedStateFields generates
// expected notes: standard fields render successfully into the struct, while fields
// that cannot be mapped generate TODO markers.
func TestWriteObservedStateFieldsNotes(t *testing.T) {
	msg := observedStateTestMessage(t)

	for _, tc := range []struct {
		name         string
		field        string
		wantRendered string
	}{
		{
			name:         "a field with a Go type renders into the struct",
			field:        "create_time",
			wantRendered: `json:"createTime`,
		},
		{
			name:         "a field WriteField cannot type leaves a marker",
			field:        "by_index",
			wantRendered: "// TODO:",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			details := &OutputMessageDetails{
				Message:      msg,
				OutputFields: []protoreflect.FieldDescriptor{msg.Fields().ByName(protoreflect.Name(tc.field))},
			}

			// Act
			var buf bytes.Buffer
			notes := WriteObservedStateFields(&buf, details, sets.NewString(), nil, WriteOptions{})

			// Assert
			if len(notes) != 1 {
				t.Fatalf("got %d notes, want one", len(notes))
			}
			if notes[0].Skipped {
				t.Errorf("Skipped = true, want false for a field the skip map does not name")
			}
			if !strings.Contains(notes[0].Rendered, tc.wantRendered) {
				t.Errorf("Rendered = %q, want it to contain %q", notes[0].Rendered, tc.wantRendered)
			}
			if !strings.Contains(buf.String(), tc.wantRendered) {
				t.Errorf("struct body = %q, want it to contain %q", buf.String(), tc.wantRendered)
			}
		})
	}
}

// TestWriteObservedStateFieldsSkips verifies that fields in the skip map produce
// no struct output and are marked as Skipped in the returned notes.
func TestWriteObservedStateFieldsSkips(t *testing.T) {
	// Arrange
	msg := observedStateTestMessage(t)
	details := &OutputMessageDetails{
		Message:      msg,
		OutputFields: []protoreflect.FieldDescriptor{msg.Fields().ByName("name")},
	}

	// Act
	var buf bytes.Buffer
	notes := WriteObservedStateFields(&buf, details, sets.NewString(), map[string]bool{"name": true}, WriteOptions{})

	// Assert
	if len(notes) != 1 {
		t.Fatalf("got %d notes, want one", len(notes))
	}
	if !notes[0].Skipped {
		t.Errorf("Skipped = false, want true")
	}
	if notes[0].Rendered != "" {
		t.Errorf("Rendered = %q, want empty", notes[0].Rendered)
	}
	if buf.String() != "" {
		t.Errorf("struct body = %q, want empty", buf.String())
	}
}

// TestWriteObservedStateFieldsHonoursWriteOptions verifies that WriteObservedStateFields
// respects WriteOptions while explicitly omitting +required markers (which are
// invalid on status / observed-state fields).
func TestWriteObservedStateFieldsHonoursWriteOptions(t *testing.T) {
	msg := observedStateTestMessage(t)

	for _, tc := range []struct {
		name         string
		field        string
		opts         WriteOptions
		wantRendered string
		wantJSONName string
	}{
		{
			name:         "EmitRequired is the one flag that does not carry over",
			field:        "required_field",
			opts:         WriteOptions{EmitRequired: true},
			wantRendered: `json:"requiredField,omitempty"`,
			wantJSONName: "requiredField",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			details := &OutputMessageDetails{
				Message:      msg,
				OutputFields: []protoreflect.FieldDescriptor{msg.Fields().ByName(protoreflect.Name(tc.field))},
			}

			// Act
			var buf bytes.Buffer
			notes := WriteObservedStateFields(&buf, details, sets.NewString(), nil, tc.opts)

			// Assert
			if len(notes) != 1 {
				t.Fatalf("got %d notes, want one", len(notes))
			}
			if !strings.Contains(buf.String(), tc.wantRendered) {
				t.Errorf("struct body = %q, want it to contain %q", buf.String(), tc.wantRendered)
			}
			if notes[0].JSONName != tc.wantJSONName {
				t.Errorf("note JSONName = %q, want %q", notes[0].JSONName, tc.wantJSONName)
			}
			// A required marker here would make the API server reject a status
			// KCC itself wrote.
			if strings.Contains(buf.String(), "+required") {
				t.Errorf("struct body = %q, want no +required marker", buf.String())
			}
		})
	}
}

// observedStateTestMessage constructs a mock proto message descriptor with various
// field behaviors (supported types, unsupported map key types, and required fields)
// for testing observed-state field emission.
func observedStateTestMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	requiredOpts := &descriptorpb.FieldOptions{}
	proto.SetExtension(requiredOpts, annotations.E_FieldBehavior,
		[]annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED})

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("obs.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("TestMessage"),
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: protoPtr("ByIndexEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_INT32)},
							{Name: protoPtr("value"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
						},
						Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
					},
				},
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("name"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: protoPtr("create_time"), Number: protoPtr(int32(2)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{
						Name: protoPtr("by_index"), Number: protoPtr(int32(3)),
						Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
						TypeName: protoPtr(".google.cloud.test.v1.TestMessage.ByIndexEntry"),
						Label:    labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED),
					},
					{
						Name: protoPtr("required_field"), Number: protoPtr(int32(6)),
						Type:    typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: requiredOpts,
					},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	return fd.Messages().ByName("TestMessage")
}

func TestAcronymCasing(t *testing.T) {
	tests := []struct {
		token   string
		plurals bool
		want    string
		wantOK  bool
	}{
		// Singular has always worked, with or without the option.
		{token: "url", plurals: false, want: "URL", wantOK: true},
		{token: "url", plurals: true, want: "URL", wantOK: true},
		{token: "uri", plurals: false, want: "URI", wantOK: true},
		{token: "id", plurals: true, want: "ID", wantOK: true},

		// Plurals are the gap: EqualFold("uris", "URI") is false, so these fell
		// through to strings.Title and became Uris and Ids.
		{token: "uris", plurals: false, wantOK: false},
		{token: "uris", plurals: true, want: "URIs", wantOK: true},
		{token: "ids", plurals: false, wantOK: false},
		{token: "ids", plurals: true, want: "IDs", wantOK: true},
		{token: "fqdns", plurals: true, want: "FQDNs", wantOK: true},
		{token: "cpus", plurals: true, want: "CPUs", wantOK: true},

		// These are not acronyms either way. "s" must not be stripped down to
		// nothing, and a word that merely ends in s stays a word.
		{token: "s", plurals: true, wantOK: false},
		{token: "values", plurals: true, wantOK: false},
		{token: "description", plurals: true, wantOK: false},
	}
	for _, tt := range tests {
		got, ok := AcronymCasing(tt.token, tt.plurals)
		if ok != tt.wantOK {
			t.Errorf("AcronymCasing(%q, %v) ok = %v, want %v", tt.token, tt.plurals, ok, tt.wantOK)
			continue
		}
		if ok && got != tt.want {
			t.Errorf("AcronymCasing(%q, %v) = %q, want %q", tt.token, tt.plurals, got, tt.want)
		}
	}
}

// Maps had no tests before this, including the two value types that already
// worked. A regression here would not fail generation: WriteField replaces a
// field it cannot type with a "// TODO:" comment, and the field is missing
// from the CRD.
func TestGoTypeForFieldMaps(t *testing.T) {
	mapEntry := func(name string, key, value *descriptorpb.FieldDescriptorProto) *descriptorpb.DescriptorProto {
		return &descriptorpb.DescriptorProto{
			Name:    protoPtr(name),
			Field:   []*descriptorpb.FieldDescriptorProto{key, value},
			Options: &descriptorpb.MessageOptions{MapEntry: protoPtr(true)},
		}
	}
	strKey := func() *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)}
	}
	val := func(t descriptorpb.FieldDescriptorProto_Type, typeName string) *descriptorpb.FieldDescriptorProto {
		f := &descriptorpb.FieldDescriptorProto{Name: protoPtr("value"), Number: protoPtr(int32(2)), Type: typeDescriptor(t)}
		if typeName != "" {
			f.TypeName = protoPtr(typeName)
		}
		return f
	}
	mapField := func(name string, num int32, entry string) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: protoPtr(name), Number: protoPtr(num),
			Type:     typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
			TypeName: protoPtr(".google.cloud.test.v1.TestMessage." + entry),
			Label:    labelDescriptor(descriptorpb.FieldDescriptorProto_LABEL_REPEATED),
		}
	}

	fdp := &descriptorpb.FileDescriptorProto{
		Name:       protoPtr("maps.proto"),
		Package:    protoPtr("google.cloud.test.v1"),
		Dependency: []string{"google/protobuf/struct.proto", "google/protobuf/timestamp.proto"},
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("TargetMessage")},
			{
				Name: protoPtr("TestMessage"),
				NestedType: []*descriptorpb.DescriptorProto{
					mapEntry("StringMapEntry", strKey(), val(descriptorpb.FieldDescriptorProto_TYPE_STRING, "")),
					mapEntry("Int64MapEntry", strKey(), val(descriptorpb.FieldDescriptorProto_TYPE_INT64, "")),
					mapEntry("MessageMapEntry", strKey(), val(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".google.cloud.test.v1.TargetMessage")),
					mapEntry("StructMapEntry", strKey(), val(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".google.protobuf.Struct")),
					mapEntry("TimestampMapEntry", strKey(), val(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, ".google.protobuf.Timestamp")),
					mapEntry("IntKeyMapEntry",
						&descriptorpb.FieldDescriptorProto{Name: protoPtr("key"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_INT32)},
						val(descriptorpb.FieldDescriptorProto_TYPE_STRING, "")),
				},
				Field: []*descriptorpb.FieldDescriptorProto{
					mapField("string_map", 1, "StringMapEntry"),
					mapField("int64_map", 2, "Int64MapEntry"),
					mapField("message_map", 3, "MessageMapEntry"),
					mapField("struct_map", 4, "StructMapEntry"),
					mapField("int_key_map", 5, "IntKeyMapEntry"),
					mapField("timestamp_map", 6, "TimestampMapEntry"),
				},
			},
		},
	}

	// struct_map refers to google.protobuf.Struct, so struct.proto has to
	// resolve. The descriptor comes from the structpb package, because the
	// global registry only holds files that some linked package registered.
	deps := new(protoregistry.Files)
	for _, f := range []protoreflect.FileDescriptor{
		(&structpb.Struct{}).ProtoReflect().Descriptor().ParentFile(),
		timestamppb.Now().ProtoReflect().Descriptor().ParentFile(),
	} {
		if err := deps.RegisterFile(f); err != nil {
			t.Fatalf("registering %s: %v", f.Path(), err)
		}
	}
	fd, err := protodesc.NewFile(fdp, deps)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	fields := fd.Messages().ByName("TestMessage").Fields()

	tests := []struct {
		field   string
		want    string
		wantErr bool
		why     string
	}{
		{field: "string_map", want: "map[string]string", why: "worked before this change, without a test"},
		{field: "int64_map", want: "map[string]int64", why: "worked before this change, without a test"},
		{
			field: "message_map", want: "map[string]TargetMessage",
			why: "the value's message generates as a struct, and the map holds values " +
				"rather than pointers, as 16 of the 23 message maps in the corpus do",
		},
		{
			field: "int_key_map", wantErr: true,
			why: "a CRD keys additionalProperties by string",
		},
		{
			field: "struct_map", want: "map[string]apiextensionsv1.JSON",
			why: "a message with a special-cased Go type uses that type",
		},
		{
			field: "timestamp_map", want: "map[string]string",
			why: "the same holds when the special-cased type is a scalar",
		},
	}

	for _, tt := range tests {
		f := fields.ByName(protoreflect.Name(tt.field))
		if f == nil {
			t.Fatalf("could not find field %q", tt.field)
		}
		if !f.IsMap() {
			t.Fatalf("%q is not a map field; the fixture is wrong", tt.field)
		}
		got, err := GoTypeForField(f, false, WriteOptions{EmitMessageMaps: true})
		if tt.wantErr {
			if err == nil {
				t.Errorf("GoTypeForField(%q) = %q, want an error (%s)", tt.field, got, tt.why)
			}
			continue
		}
		if err != nil {
			t.Errorf("GoTypeForField(%q) returned error: %v (%s)", tt.field, err, tt.why)
			continue
		}
		if got != tt.want {
			t.Errorf("GoTypeForField(%q) = %q, want %q (%s)", tt.field, got, tt.want, tt.why)
		}
	}

	// With EmitMessageMaps off, every message-valued map is declined, as before
	// the flag existed. Scalar-valued maps are unaffected.
	for _, name := range []string{"message_map", "struct_map", "timestamp_map"} {
		if got, err := GoTypeForField(fields.ByName(protoreflect.Name(name)), false, WriteOptions{}); err == nil {
			t.Errorf("GoTypeForField(%q) with EmitMessageMaps off = %q, want an error", name, got)
		}
	}
	if got, err := GoTypeForField(fields.ByName("string_map"), false, WriteOptions{}); err != nil || got != "map[string]string" {
		t.Errorf("GoTypeForField(%q) with EmitMessageMaps off = %q, %v, want map[string]string", "string_map", got, err)
	}
}

// TestReservedTypeNamesExistingDeclarationPrecedence verifies that when a type already
// exists in a non-autogenerated file on disk, it is skipped as "found existing non-generated go type"
// rather than being attributed to the scaffolder reservation.
func TestReservedTypeNamesExistingDeclarationPrecedence(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("test.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protoPtr("BillingAccount"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: protoPtr("name"), Number: protoPtr(int32(1)), Type: typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("failed to create file descriptor: %v", err)
	}
	msg := fd.Messages().ByName("BillingAccount")
	goTypeName := "BillingAccount"

	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "test")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}

	// Write an existing non-generated type declaration.
	existingCode := `package test

type BillingAccount struct {
	Name string
}
`
	if err := os.WriteFile(filepath.Join(pkgDir, "billingaccount_types.go"), []byte(existingCode), 0o644); err != nil {
		t.Fatal(err)
	}

	g := NewTypeGenerator("test", dir, nil)
	g.visitedMessages = []protoreflect.MessageDescriptor{msg}
	g.includeSkippedOutput = true
	g.WithReservedTypeNames(goTypeName)

	if err := g.WriteVisitedMessages(); err != nil {
		t.Fatalf("WriteVisitedMessages: %v", err)
	}

	f := g.generatedFiles[generatedFileKey{GoPackage: "test", FileName: "types.generated.go"}]
	if f == nil {
		t.Fatal("expected generated file")
	}
	body := f.body.String()
	wantExisting := `/* found existing non-generated go type "BillingAccount", skipping`
	if !strings.Contains(body, wantExisting) {
		t.Errorf("body missing %q; got:\n%s", wantExisting, body)
	}
	dontWantReserved := `is a Kind the scaffolder will declare`
	if strings.Contains(body, dontWantReserved) {
		t.Errorf("body should not contain %q when type exists on disk; got:\n%s", dontWantReserved, body)
	}
}
