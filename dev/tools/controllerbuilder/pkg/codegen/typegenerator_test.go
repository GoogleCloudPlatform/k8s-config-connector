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
			WriteField(&buf, msg.Fields().Get(0), msg, 0, false, g.opts)

			got := strings.Contains(buf.String(), "// +required")
			if got != g.wantMarker {
				t.Errorf("+required present = %v, want %v\noutput:\n%s", got, g.wantMarker, buf.String())
			}
		})
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
	WriteMessage(&buf, msg, WriteOptions{})

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
