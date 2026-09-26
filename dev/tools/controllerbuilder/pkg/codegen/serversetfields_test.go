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

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// serverSetTestFile builds two messages:
//
//	Discovery  no field_behavior anywhere, like compute's protos
//	Annotated  one field marked OUTPUT_ONLY, the rest bare
//
// Both carry the same field names, so a test can isolate the guard from the
// allowlist.
func serverSetTestFile(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()

	outputOnly := &descriptorpb.FieldOptions{}
	proto.SetExtension(outputOnly, annotations.E_FieldBehavior,
		[]annotations.FieldBehavior{annotations.FieldBehavior_OUTPUT_ONLY})

	names := []string{"creation_timestamp", "self_link", "etag", "state", "status", "type", "name", "description"}
	bare := func() []*descriptorpb.FieldDescriptorProto {
		var out []*descriptorpb.FieldDescriptorProto
		for i, n := range names {
			out = append(out, &descriptorpb.FieldDescriptorProto{
				Name:   protoPtr(n),
				Number: protoPtr(int32(i + 1)),
				Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
			})
		}
		return out
	}

	annotated := bare()
	// Only "description" is annotated. No allowlisted field carries an
	// annotation itself, so the guard has to come from the message, not the
	// field.
	annotated[len(annotated)-1].Options = outputOnly

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("serverset.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("Discovery"), Field: bare()},
			{Name: protoPtr("Annotated"), Field: annotated},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd
}

// TestIsServerSetField covers the allowlist and the guard: the names the rule
// moves, the names it leaves in the Spec, and the annotation that turns it off.
func TestIsServerSetField(t *testing.T) {
	fd := serverSetTestFile(t)
	discovery := fd.Messages().ByName("Discovery")
	annotated := fd.Messages().ByName("Annotated")
	on := WriteOptions{PlaceServerSetFields: true}

	tests := []struct {
		name  string
		msg   protoreflect.MessageDescriptor
		field string
		opts  WriteOptions
		want  bool
	}{
		{"allowlisted, nothing annotated", discovery, "creation_timestamp", on, true},
		{"allowlisted, nothing annotated", discovery, "self_link", on, true},
		{"etag is in the list", discovery, "etag", on, true},

		// serverSetFieldNames records why these three stay out.
		{"state stays in the Spec", discovery, "state", on, false},
		{"status stays in the Spec", discovery, "status", on, false},
		{"type stays in the Spec", discovery, "type", on, false},
		// identityFields in the scaffold package handles name instead.
		{"name is left to the identity policy", discovery, "name", on, false},
		{"an ordinary field is untouched", discovery, "description", on, false},

		// The guard: one annotation anywhere on the message turns the rule off.
		{"annotated message, allowlisted field", annotated, "creation_timestamp", on, false},
		{"annotated message, etag", annotated, "etag", on, false},

		{"off by default", discovery, "creation_timestamp", WriteOptions{}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name+"/"+tc.field, func(t *testing.T) {
			// Arrange
			f := tc.msg.Fields().ByName(protoreflect.Name(tc.field))
			if f == nil {
				t.Fatalf("no field %q on %s", tc.field, tc.msg.Name())
			}

			// Act
			got := IsServerSetField(f, tc.msg, tc.opts)

			// Assert
			if got != tc.want {
				t.Errorf("IsServerSetField(%s.%s) = %v, want %v", tc.msg.Name(), tc.field, got, tc.want)
			}
		})
	}
}

// TestIsServerSetFieldOnlyAppliesToTheRootMessage pins where the rule stops.
// identifyOutputs recurses into nested messages, and a creationTimestamp found
// there is often the user's to set.
func TestIsServerSetFieldOnlyAppliesToTheRootMessage(t *testing.T) {
	// Arrange: one field of one message, asked twice. Discovery carries no
	// annotation, so the guard cannot account for a false answer here and only
	// the root check can.
	fd := serverSetTestFile(t)
	discovery := fd.Messages().ByName("Discovery")
	field := discovery.Fields().ByName("creation_timestamp")
	opts := WriteOptions{PlaceServerSetFields: true}
	visiting := &TypeGenerator{writeOptions: opts, rootMessageFQN: string(discovery.FullName())}
	elsewhere := &TypeGenerator{writeOptions: opts, rootMessageFQN: "google.cloud.test.v1.Annotated"}

	// Act
	atRoot := visiting.isServerSet(field, discovery)
	away := elsewhere.isServerSet(field, discovery)

	// Assert
	if !atRoot {
		t.Error("creationTimestamp on the visited message is not server-set, want server-set")
	}
	if away {
		t.Error("creationTimestamp below the visited message is server-set, want not server-set")
	}
}
