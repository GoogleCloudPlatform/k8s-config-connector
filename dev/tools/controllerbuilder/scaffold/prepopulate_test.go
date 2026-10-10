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
	"fmt"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/util/sets"
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

// TestPrepopulateSpecQueuesResourceReferences checks the possible-reference
// entry for a top-level field with google.api.resource_reference, once with
// type and once with child_type. ReferenceHints queues the nested ones; see
// TestReferenceHintsQueuesNestedResourceReferences.
func TestPrepopulateSpecQueuesResourceReferences(t *testing.T) {
	// Arrange
	msg := ragStoreMessage(t)
	want := []JudgementItem{
		{
			Reason: "untriaged-bulk-generation",
			Detail: "spec was generated from proto definition; verify refs, omissions, and KRM conventions",
		},
		{
			FieldPath: ".spec.corpus",
			Reason:    "possible-reference",
			Detail:    "the proto marks this field as a reference to aiplatform.googleapis.com/RagCorpus (google.api.resource_reference); confirm whether it should be a KCC reference",
		},
		{
			FieldPath: ".spec.parent",
			Reason:    "possible-reference",
			Detail:    "the proto marks this field as the parent of a aiplatform.googleapis.com/RagFile (google.api.resource_reference child_type); confirm whether it should be a KCC reference",
		},
	}

	// Act
	got, err := PrepopulateSpec(msg, codegen.WriteOptions{})

	// Assert
	if err != nil {
		t.Fatalf("PrepopulateSpec() error: %v", err)
	}
	if diff := cmp.Diff(want, got.Judgement); diff != "" {
		t.Errorf("PrepopulateSpec() Judgement mismatch (-want +got):\n%s", diff)
	}
}

func TestPrepopulateSpecRequiresAMessage(t *testing.T) {
	if _, err := PrepopulateSpec(nil, codegen.WriteOptions{}); err == nil {
		t.Fatal("expected an error for a nil message")
	}
}

func TestJudgementEntries(t *testing.T) {
	items := []JudgementItem{
		{Reason: "untriaged-bulk-generation", Detail: "generated mechanically"},
		{FieldPath: ".spec.network", Reason: "possible-reference", Detail: "target=compute.googleapis.com/Network"},
	}

	got := JudgementEntries("ExampleWidget", "example.cnrm.cloud.google.com", items)

	want := []judgement.Entry{
		{Kind: "ExampleWidget", Group: "example.cnrm.cloud.google.com", Reason: "untriaged-bulk-generation", Detail: "generated mechanically", Status: judgement.StatusOpen},
		{Kind: "ExampleWidget", Group: "example.cnrm.cloud.google.com", Field: ".spec.network", Reason: "possible-reference", Detail: "target=compute.googleapis.com/Network", Status: judgement.StatusOpen},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v\nwant %+v", got, want)
	}
	for _, e := range got {
		if err := e.Validate(); err != nil {
			t.Errorf("entry %+v is not valid: %v", e, err)
		}
	}
}

func testMessageWithMap(t *testing.T) (protoreflect.MessageDescriptor, protoreflect.FieldDescriptor) {
	t.Helper()
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("maps.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: strPtr("TargetMessage")},
			{
				Name: strPtr("Widget"),
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: strPtr("TasksEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: strPtr("key"), Number: i32Ptr(1), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
							{Name: strPtr("value"), Number: i32Ptr(2), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: strPtr(".google.cloud.test.v1.TargetMessage")},
						},
						Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
					},
				},
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     strPtr("tasks"),
						Number:   i32Ptr(1),
						Type:     fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
						TypeName: strPtr(".google.cloud.test.v1.Widget.TasksEntry"),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	msg := fd.Messages().ByName("Widget")
	return msg, msg.Fields().ByName("tasks")
}

func TestPrepopulateSpecQueuesDerivedMessageMaps(t *testing.T) {
	msg, _ := testMessageWithMap(t)

	// EmitMessageMaps: true -> should emit message-map-derived
	got, err := PrepopulateSpec(msg, codegen.WriteOptions{EmitMessageMaps: true})
	if err != nil {
		t.Fatalf("PrepopulateSpec failed: %v", err)
	}
	var derived []JudgementItem
	for _, j := range got.Judgement {
		if j.Reason == "message-map-derived" {
			derived = append(derived, j)
		}
	}
	if len(derived) != 1 {
		t.Fatalf("expected 1 message-map-derived item, got %d: %+v", len(derived), got.Judgement)
	}
	if derived[0].FieldPath != ".spec.tasks" {
		t.Errorf("FieldPath = %q, want .spec.tasks", derived[0].FieldPath)
	}
	if !strings.Contains(derived[0].Detail, "TargetMessage") {
		t.Errorf("Detail = %q, want mention of TargetMessage", derived[0].Detail)
	}

	// EmitMessageMaps: false -> should not emit message-map-derived
	gotOff, err := PrepopulateSpec(msg, codegen.WriteOptions{EmitMessageMaps: false})
	if err != nil {
		t.Fatalf("PrepopulateSpec failed: %v", err)
	}
	for _, j := range gotOff.Judgement {
		if j.Reason == "message-map-derived" {
			t.Errorf("unexpected message-map-derived item when EmitMessageMaps is false: %+v", j)
		}
	}
}

func TestPrepopulateObservedStateQueuesDerivedMessageMaps(t *testing.T) {
	msg, field := testMessageWithMap(t)
	details := &codegen.OutputMessageDetails{
		Message:      msg,
		OutputFields: []protoreflect.FieldDescriptor{field},
	}

	// EmitMessageMaps: true -> should emit message-map-derived
	_, jOn := PrepopulateObservedState(details, sets.NewString(), codegen.WriteOptions{EmitMessageMaps: true})
	var derived []JudgementItem
	for _, j := range jOn {
		if j.Reason == "message-map-derived" {
			derived = append(derived, j)
		}
	}
	if len(derived) != 1 {
		t.Fatalf("expected 1 message-map-derived item, got %d: %+v", len(derived), jOn)
	}
	if derived[0].FieldPath != ".status.observedState.tasks" {
		t.Errorf("FieldPath = %q, want .status.observedState.tasks", derived[0].FieldPath)
	}

	// EmitMessageMaps: false -> should not emit message-map-derived
	_, jOff := PrepopulateObservedState(details, sets.NewString(), codegen.WriteOptions{EmitMessageMaps: false})
	for _, j := range jOff {
		if j.Reason == "message-map-derived" {
			t.Errorf("unexpected message-map-derived item when EmitMessageMaps is false: %+v", j)
		}
	}
}

func TestDetectOutputOnlyInComments(t *testing.T) {
	// Arrange
	msg := commentedMessage(t,
		"Output only. Set by the server.",                   // field_0, the long-standing spelling
		"[Output Only] IP address on the Google side.",      // field_1, how Compute writes it
		"The display name of the widget.",                   // field_2, no signal
		"Set by the user. Output only in some other sense.", // field_3, marker not at the front, so the weaker reason
	)
	want := []OutputOnlyCandidate{
		{FieldPath: ".spec.field0", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output only. Set by the server."},
		{FieldPath: ".spec.field1", Reason: "well-known-output-only-pattern-in-comment", Comment: "[Output Only] IP address on the Google side."},
		{FieldPath: ".spec.field3", Reason: "possible-output-only-pattern-in-comment", Comment: "Set by the user. Output only in some other sense."},
	}

	// Act
	got := DetectOutputOnlyInComments(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("DetectOutputOnlyInComments() mismatch (-want +got):\n%s", diff)
	}
}

// Any other mention of "output only" gets the weaker reason, in any case, in
// brackets or with a hyphen. That covers the words after other text, as in
// websecurityscanner's managed_scan, and a comment that starts with the words
// but not with a marker: a typo, or a condition. A sentence break between the
// words is not a match.
func TestDetectOutputOnlyMentionedInComment(t *testing.T) {
	// Arrange
	msg := commentedMessage(t,
		"Whether the scan config is managed by Web Security Scanner, output only.",      // field_0, websecurityscanner
		"When the version was created. Output-only field, populated by the system.",     // field_1, dlp's hyphen
		"Required. [Output Only] Set by the server.",                                    // field_2, a marker further on
		"The job output. Only the last run is kept.",                                    // field_3, a sentence break
		"Output only . The translated content.",                                         // field_4, automl's typo
		"Output only\n When true, the index configuration is being reverted.",           // field_5, firestore, the words on their own line
		"[Output only for type PARTNER. Input only for PARTNER_PROVIDER.] Pairing key.", // field_6, compute, conditional
		"Output only for the create operation. Required for update.",                    // field_7, spanner-style, conditional
	)
	want := []OutputOnlyCandidate{
		{FieldPath: ".spec.field0", Reason: "possible-output-only-pattern-in-comment", Comment: "Whether the scan config is managed by Web Security Scanner, output only."},
		{FieldPath: ".spec.field1", Reason: "possible-output-only-pattern-in-comment", Comment: "When the version was created. Output-only field, populated by the system."},
		{FieldPath: ".spec.field2", Reason: "possible-output-only-pattern-in-comment", Comment: "Required. [Output Only] Set by the server."},
		{FieldPath: ".spec.field4", Reason: "possible-output-only-pattern-in-comment", Comment: "Output only . The translated content."},
		{FieldPath: ".spec.field5", Reason: "possible-output-only-pattern-in-comment", Comment: "Output only When true, the index configuration is being reverted."},
		{FieldPath: ".spec.field6", Reason: "possible-output-only-pattern-in-comment", Comment: "[Output only for type PARTNER. Input only for PARTNER_PROVIDER.] Pairing key."},
		{FieldPath: ".spec.field7", Reason: "possible-output-only-pattern-in-comment", Comment: "Output only for the create operation. Required for update."},
	}

	// Act
	got := DetectOutputOnlyInComments(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("DetectOutputOnlyInComments() mismatch (-want +got):\n%s", diff)
	}
}

// TestDetectOutputOnlyWalksNestedFields checks that both rules reach nested
// Spec fields, including ones in a list. TestOutputOnlyCandidateItem checks
// the path the queue entry gives a nested field.
func TestDetectOutputOnlyWalksNestedFields(t *testing.T) {
	// Arrange
	msg := nestedConfigMessage(t)
	want := []OutputOnlyCandidate{
		{FieldPath: ".spec.config.state", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output only. The state of the config."},
		{FieldPath: ".spec.config.managed", Reason: "possible-output-only-pattern-in-comment", Comment: "Whether the config is managed by the service, output only."},
		{FieldPath: ".spec.peers[].state", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output only. The state of the config."},
		{FieldPath: ".spec.peers[].managed", Reason: "possible-output-only-pattern-in-comment", Comment: "Whether the config is managed by the service, output only."},
	}

	// Act
	got := DetectOutputOnlyInComments(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("DetectOutputOnlyInComments() mismatch (-want +got):\n%s", diff)
	}
}

// TestOutputOnlyCandidateItem checks the entry for each reason. The entry
// names the field's place under ObservedState and keeps the rest of a nested
// path. The detail ends with the comment, so a reviewer can decide without
// opening the proto.
func TestOutputOnlyCandidateItem(t *testing.T) {
	for _, tc := range []struct {
		name string
		cand OutputOnlyCandidate
		want JudgementItem
	}{
		{
			name: "comment starts with the marker",
			cand: OutputOnlyCandidate{FieldPath: ".spec.createTime", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output only. When it was created."},
			want: JudgementItem{
				FieldPath: ".status.observedState.createTime",
				Reason:    "well-known-output-only-pattern-in-comment",
				Detail:    "proto comment says output only but there is no field_behavior annotation, so it was generated into the Spec. Move the field to status.observedState if it is confirmed output only. Proto comment: Output only. When it was created.",
			},
		},
		{
			name: "comment mentions the words in a nested field",
			cand: OutputOnlyCandidate{FieldPath: ".spec.config.managedScan", Reason: "possible-output-only-pattern-in-comment", Comment: "Whether the scan config is managed by Web Security Scanner, output only."},
			want: JudgementItem{
				FieldPath: ".status.observedState.config.managedScan",
				Reason:    "possible-output-only-pattern-in-comment",
				Detail: "proto comment mentions output only but doesn't start with \"Output only.\" or \"[Output Only]\", and there is no field_behavior annotation, so it was generated into the Spec. " +
					"The comment may be a typo, apply only some of the time, or mean something else. Move the field to status.observedState if it is confirmed output only. Proto comment: Whether the scan config is managed by Web Security Scanner, output only.",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := tc.cand.Item()

			// Assert
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Item() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Compute writes both "[Output Only]" and "[Output only]", so the prefixes
// match in any case.
func TestDetectOutputOnlyIgnoresCase(t *testing.T) {
	// Arrange
	msg := commentedMessage(t,
		"[Output only] Number of network endpoints in the group.", // field_0, compute's other spelling
		"Output Only. The overall outcome of the test.",           // field_1, devtools.testing
	)
	want := []OutputOnlyCandidate{
		{FieldPath: ".spec.field0", Reason: "well-known-output-only-pattern-in-comment", Comment: "[Output only] Number of network endpoints in the group."},
		{FieldPath: ".spec.field1", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output Only. The overall outcome of the test."},
	}

	// Act
	got := DetectOutputOnlyInComments(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("DetectOutputOnlyInComments() mismatch (-want +got):\n%s", diff)
	}
}

// A field the proto already annotates needs no prose detection. Reporting it
// would queue a field the generator has already placed correctly.
func TestDetectOutputOnlySkipsAnnotatedFields(t *testing.T) {
	// Arrange
	unannotated := commentedMessage(t, "[Output Only] Set by the server.")
	annotated := testMessage(t)

	// Act
	fromUnannotated := DetectOutputOnlyInComments(unannotated, codegen.WriteOptions{})
	fromAnnotated := DetectOutputOnlyInComments(annotated, codegen.WriteOptions{})

	// Assert
	if len(fromUnannotated) != 1 {
		t.Errorf("unannotated field should be reported, got %d", len(fromUnannotated))
	}
	if len(fromAnnotated) != 0 {
		t.Errorf("annotated fields should not be reported, got %v", fromAnnotated)
	}
}

// With --place-server-set-fields, a field on the server-set allowlist is
// already in ObservedState and has its own server-set-field-placed entry.
// Reporting it here too would add an entry saying it was left in the Spec.
func TestDetectOutputOnlySkipsServerSetFields(t *testing.T) {
	// Arrange: no field_behavior anywhere, like compute's Network.
	msg := namedCommentedMessage(t, [][2]string{
		{"creation_timestamp", "[Output Only] Creation timestamp in RFC3339 text format."},              // on the allowlist
		{"firewall_policy", "[Output Only] URL of the firewall policy the network is associated with."}, // not on it
	})

	tests := []struct {
		name string
		opts codegen.WriteOptions
		want []string
	}{
		{"server-set placement on", codegen.WriteOptions{PlaceServerSetFields: true}, []string{".spec.firewallPolicy"}},
		{"server-set placement off", codegen.WriteOptions{}, []string{".spec.creationTimestamp", ".spec.firewallPolicy"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := DetectOutputOnlyInComments(msg, tc.opts)

			// Assert
			var paths []string
			for _, c := range got {
				paths = append(paths, c.FieldPath)
			}
			if strings.Join(paths, ",") != strings.Join(tc.want, ",") {
				t.Errorf("got %v, want %v", paths, tc.want)
			}
		})
	}
}

// TestSpecFieldsMatchesPrepopulateSpec checks that SpecFields lists the fields
// PrepopulateSpec writes, in the same order. generate-types uses SpecFields to
// plan a new Kind's structs before PrepopulateSpec runs.
func TestSpecFieldsMatchesPrepopulateSpec(t *testing.T) {
	for _, tc := range []struct {
		name string
		msg  protoreflect.MessageDescriptor
		opts codegen.WriteOptions
	}{
		{name: "field_behavior annotations", msg: testMessage(t)},
		{
			name: "server-set fields placed in status",
			msg:  namedCommentedMessage(t, [][2]string{{"etag", ""}, {"description", ""}, {"self_link", ""}}),
			opts: codegen.WriteOptions{PlaceServerSetFields: true},
		},
		{name: "deprecated top-level fields", msg: deprecatedMessage(t)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			fields := SpecFields(tc.msg, tc.opts)
			got, err := PrepopulateSpec(tc.msg, tc.opts)
			if err != nil {
				t.Fatalf("PrepopulateSpec: %v", err)
			}

			// Assert
			var want []string
			for _, line := range strings.Split(got.SpecFields, "\n") {
				if v, ok := strings.CutPrefix(strings.TrimSpace(line), "// "+codegen.KCCProtoFieldAnnotation+"="); ok {
					want = append(want, v)
				}
			}
			var names []string
			for _, f := range fields {
				names = append(names, string(f.FullName()))
			}
			if !reflect.DeepEqual(names, want) {
				t.Errorf("SpecFields() = %v, PrepopulateSpec wrote %v", names, want)
			}
		})
	}
}

// commentedMessage builds a message whose fields carry leading comments, which
// is what DetectOutputOnlyInComments reads. The fields are named field_0,
// field_1 and so on.
func commentedMessage(t *testing.T, comments ...string) protoreflect.MessageDescriptor {
	t.Helper()
	var fields [][2]string
	for i, c := range comments {
		fields = append(fields, [2]string{fmt.Sprintf("field_%d", i), c})
	}
	return namedCommentedMessage(t, fields)
}

// namedCommentedMessage is commentedMessage for a test that needs real field
// names. Each entry is {name, comment}. SourceCodeInfo paths are
// [4=message_type, msgIndex, 2=field, fieldIndex].
func namedCommentedMessage(t *testing.T, named [][2]string) protoreflect.MessageDescriptor {
	t.Helper()
	var fields []*descriptorpb.FieldDescriptorProto
	var locs []*descriptorpb.SourceCodeInfo_Location
	for i, nc := range named {
		fields = append(fields, &descriptorpb.FieldDescriptorProto{
			Name:   strPtr(nc[0]),
			Number: i32Ptr(int32(i + 1)),
			Type:   fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING),
		})
		locs = append(locs, &descriptorpb.SourceCodeInfo_Location{
			Path:            []int32{4, 0, 2, int32(i)},
			Span:            []int32{int32(i), 0, 1},
			LeadingComments: strPtr(" " + nc[1] + "\n"),
		})
	}
	fdp := &descriptorpb.FileDescriptorProto{
		Name:           strPtr("commented.proto"),
		Package:        strPtr("google.cloud.test.v1"),
		MessageType:    []*descriptorpb.DescriptorProto{{Name: strPtr("Widget"), Field: fields}},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: locs},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("Widget")
}

// nestedConfigMessage builds a Widget that holds a Config directly and in a
// list. In Config, state's comment starts with "Output only.", managed's says
// it later on, and target's doesn't mention it.
func nestedConfigMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	str := fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)
	msgType := fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	comment := func(path []int32, text string) *descriptorpb.SourceCodeInfo_Location {
		return &descriptorpb.SourceCodeInfo_Location{Path: path, Span: []int32{0, 0, 1}, LeadingComments: strPtr(" " + text + "\n")}
	}
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("nested.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: strPtr("Widget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("config"), Number: i32Ptr(1), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config")},
					{Name: strPtr("peers"), Number: i32Ptr(2), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config"), Label: descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()},
				},
			},
			{
				Name: strPtr("Config"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("state"), Number: i32Ptr(1), Type: str},
					{Name: strPtr("managed"), Number: i32Ptr(2), Type: str},
					{Name: strPtr("target"), Number: i32Ptr(3), Type: str},
				},
			},
		},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: []*descriptorpb.SourceCodeInfo_Location{
			comment([]int32{4, 1, 2, 0}, "Output only. The state of the config."),
			comment([]int32{4, 1, 2, 1}, "Whether the config is managed by the service, output only."),
			comment([]int32{4, 1, 2, 2}, "The target the config points at."),
		}},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("Widget")
}
