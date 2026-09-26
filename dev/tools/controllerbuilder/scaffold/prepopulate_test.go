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

package scaffold

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func strPtr(s string) *string { return &s }
func i32Ptr(i int32) *int32   { return &i }

func fieldType(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

func behaviorOptions(bs ...annotations.FieldBehavior) *descriptorpb.FieldOptions {
	o := &descriptorpb.FieldOptions{}
	proto.SetExtension(o, annotations.E_FieldBehavior, bs)
	return o
}

// testMessage builds a proto message with a representative mix of fields: an identifier field,
// an output-only field, a required field, and a standard optional field.
func testMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    strPtr("test.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: strPtr("Widget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:    strPtr("name"),
						Number:  i32Ptr(1),
						Type:    fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: behaviorOptions(annotations.FieldBehavior_IDENTIFIER),
					},
					{
						Name:    strPtr("create_time"),
						Number:  i32Ptr(2),
						Type:    fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: behaviorOptions(annotations.FieldBehavior_OUTPUT_ONLY),
					},
					{
						Name:    strPtr("display_name"),
						Number:  i32Ptr(3),
						Type:    fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
						Options: behaviorOptions(annotations.FieldBehavior_REQUIRED),
					},
					{
						Name:   strPtr("description"),
						Number: i32Ptr(4),
						Type:   fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
					},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("Widget")
}

func TestPrepopulateSpec(t *testing.T) {
	msg := testMessage(t)

	got, err := PrepopulateSpec(msg, codegen.WriteOptions{EmitRequired: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// The rendered body has to be valid Go inside a struct.
	src := "package p\ntype S struct {\n" + got.SpecFields + "}\n"
	if _, err := parser.ParseFile(token.NewFileSet(), "s.go", src, parser.AllErrors); err != nil {
		t.Fatalf("rendered spec is not valid Go: %v\n\n%s", err, src)
	}

	for _, want := range []string{
		"DisplayName *string",
		"Description *string",
		"+kcc:proto:field=google.cloud.test.v1.Widget.display_name",
		"// +required",
	} {
		if !strings.Contains(got.SpecFields, want) {
			t.Errorf("missing %q in:\n%s", want, got.SpecFields)
		}
	}

	// Verify that identity fields (e.g. name) and output-only fields (e.g. create_time)
	// are excluded from the Spec struct.
	for _, notWant := range []string{
		"+kcc:proto:field=google.cloud.test.v1.Widget.name\n",
		"+kcc:proto:field=google.cloud.test.v1.Widget.create_time",
	} {
		if strings.Contains(got.SpecFields, notWant) {
			t.Errorf("unexpected %q in:\n%s", notWant, got.SpecFields)
		}
	}
}

func TestPrepopulateSpecAlwaysQueuesTheResource(t *testing.T) {
	msg := testMessage(t)

	got, err := PrepopulateSpec(msg, codegen.WriteOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify that a resource-level judgement entry is always generated so the resource
	// is flagged for triage even if no explicit resource reference annotations were detected.
	if len(got.Judgement) == 0 {
		t.Fatal("expected at least a resource-level judgement entry")
	}
	if got.Judgement[0].Reason != "untriaged-bulk-generation" {
		t.Errorf("first entry reason = %q, want untriaged-bulk-generation", got.Judgement[0].Reason)
	}
	if got.Judgement[0].FieldPath != "" {
		t.Errorf("resource-level entry should have no field path, got %q", got.Judgement[0].FieldPath)
	}
}

// TestPrepopulateSpecQueuesTheEmittedName pins that a queue entry names a field
// the way the Spec struct spells it. A person works the queue by finding each
// path in the CRD, so under EmitPluralAcronyms an entry for relatedUris would
// point at nothing: the struct says relatedURIs.
func TestPrepopulateSpecQueuesTheEmittedName(t *testing.T) {
	// Arrange
	refOpts := &descriptorpb.FieldOptions{}
	proto.SetExtension(refOpts, annotations.E_ResourceReference,
		&annotations.ResourceReference{Type: "test.googleapis.com/Thing"})
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("refs.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: strPtr("Widget"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:    strPtr("related_uris"),
				Number:  i32Ptr(1),
				Type:    fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
				Options: refOpts,
			}},
		}},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	msg := fd.Messages().ByName("Widget")

	for _, tc := range []struct {
		name     string
		opts     codegen.WriteOptions
		wantPath string
	}{
		{"acronym casing on", codegen.WriteOptions{EmitPluralAcronyms: true}, ".spec.relatedURIs"},
		{"acronym casing off", codegen.WriteOptions{}, ".spec.relatedUris"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got, err := PrepopulateSpec(msg, tc.opts)

			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			var paths []string
			for _, j := range got.Judgement {
				if j.Reason == "possible-reference" {
					paths = append(paths, j.FieldPath)
				}
			}
			if len(paths) != 1 || paths[0] != tc.wantPath {
				t.Errorf("possible-reference paths = %v, want [%s]", paths, tc.wantPath)
			}
			tag := `json:"` + strings.TrimPrefix(tc.wantPath, ".spec.") + `,`
			if !strings.Contains(got.SpecFields, tag) {
				t.Errorf("SpecFields has no %s, so the entry names a field the struct lacks:\n%s", tag, got.SpecFields)
			}
		})
	}
}

func TestPrepopulateSpecRequiresAMessage(t *testing.T) {
	if _, err := PrepopulateSpec(nil, codegen.WriteOptions{}); err == nil {
		t.Fatal("expected an error for a nil message")
	}
}

func TestFormatJudgementEntries(t *testing.T) {
	items := []JudgementItem{
		{Reason: "untriaged-bulk-generation", Detail: "generated mechanically"},
		{FieldPath: ".spec.network", Reason: "possible-reference", Detail: "target=compute.googleapis.com/Network"},
	}

	got := FormatJudgementEntries("ExampleWidget", "example.cnrm.cloud.google.com", items)
	lines := strings.Split(strings.TrimSpace(got), "\n")
	if len(lines) != 2 {
		t.Fatalf("got %d lines, want 2:\n%s", len(lines), got)
	}

	// Verify that entries conform to the expected format: kind=<Kind> group=<Group>: ... reason=<Reason>
	for _, l := range lines {
		for _, need := range []string{"kind=ExampleWidget", "group=example.cnrm.cloud.google.com", "reason="} {
			if !strings.Contains(l, need) {
				t.Errorf("line missing %q: %s", need, l)
			}
		}
		if head, _, _ := strings.Cut(l, ":"); strings.Contains(head, "reason=") {
			t.Errorf("reason must follow the colon so the head parses cleanly: %s", l)
		}
	}
	if !strings.Contains(lines[0], ": resource reason=") {
		t.Errorf("resource-level entry has the wrong shape: %s", lines[0])
	}
	if !strings.Contains(lines[1], `: field ".spec.network" reason=`) {
		t.Errorf("field-level entry has the wrong shape: %s", lines[1])
	}
}
