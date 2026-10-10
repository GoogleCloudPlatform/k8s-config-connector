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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// A deprecated top-level field is left out of the Spec and gets one queue
// entry. network also has a resource_reference, so without the check it would
// also get a possible-reference entry for a field the Spec does not have.
// create_time is OUTPUT_ONLY, so it goes to ObservedState as before and gets no
// entry here.
func TestPrepopulateSpecLeavesOutDeprecatedFields(t *testing.T) {
	// Arrange
	msg := deprecatedMessage(t)
	const detail = "the proto marks this field deprecated, so new Kinds leave it out of the Spec. " +
		"If this Spec still has it, remove it by hand; add it back only if users still need it"
	want := &PrepopulateResult{
		SpecFields: "\t// +kcc:proto:field=google.cloud.test.v1.Widget.display_name\n" +
			"\tDisplayName *string `json:\"displayName,omitempty\"`\n" +
			"\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Widget.config\n" +
			"\tConfig *Config `json:\"config,omitempty\"`\n" +
			"\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Widget.peers\n" +
			"\tPeers []Config `json:\"peers,omitempty\"`\n" +
			"\n" +
			"\t// Output only. The state of the widget.\n" +
			"\t// +kcc:proto:field=google.cloud.test.v1.Widget.server_state\n" +
			"\tServerState *string `json:\"serverState,omitempty\"`\n",
		Judgement: []JudgementItem{
			{Reason: "untriaged-bulk-generation", Detail: "spec was generated from proto definition; verify refs, omissions, and KRM conventions"},
			{FieldPath: ".spec.network", Reason: "deprecated-field", Detail: detail},
			{FieldPath: ".spec.legacyConfig", Reason: "deprecated-field", Detail: detail},
			{FieldPath: ".spec.serverNote", Reason: "deprecated-field", Detail: detail},
		},
	}

	// Act
	got, err := PrepopulateSpec(msg, codegen.WriteOptions{})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("PrepopulateSpec() mismatch (-want +got):\n%s", diff)
	}
}

// The Spec holds Config directly and in a list, so each deprecated field of
// Config is reported under both paths. legacy_config is also a Config, but it
// is deprecated at the top level and so not in the Spec. Nothing under it is
// reported.
func TestNestedDeprecatedFields(t *testing.T) {
	// Arrange
	msg := deprecatedMessage(t)
	const detail = "the proto marks this field deprecated. It is kept because generated structs " +
		"can be shared with other Kinds. Remove it by hand if no existing Kind uses it"
	want := []JudgementItem{
		{FieldPath: ".spec.config.topK", Reason: "deprecated-field", Detail: detail},
		{FieldPath: ".spec.config.legacyTarget", Reason: "deprecated-field", Detail: detail},
		{FieldPath: ".spec.peers[].topK", Reason: "deprecated-field", Detail: detail},
		{FieldPath: ".spec.peers[].legacyTarget", Reason: "deprecated-field", Detail: detail},
	}

	// Act
	got := NestedDeprecatedFields(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NestedDeprecatedFields() mismatch (-want +got):\n%s", diff)
	}
}

// Top-level and nested entries share the reason deprecated-field, so they must
// never share a path. If they did, the queue would keep one entry, with the
// detail of whichever came last. walkSpecFields skips deprecated top-level
// fields, which keeps the paths apart.
func TestDeprecatedEntriesHaveDistinctKeys(t *testing.T) {
	// Arrange
	msg := deprecatedMessage(t)
	spec, err := PrepopulateSpec(msg, codegen.WriteOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	items := append(spec.Judgement, NestedDeprecatedFields(msg, codegen.WriteOptions{})...)

	// Act
	q := &judgement.Queue{Entries: JudgementEntries("ExampleWidget", "example.cnrm.cloud.google.com", items)}
	err = q.Validate()

	// Assert
	if err != nil {
		t.Errorf("queue does not validate: %v", err)
	}
}

// A deprecated top-level field is not in the Spec, so the reference hints skip
// it and everything under it. Otherwise network would get a name hint and
// legacyConfig.target a description hint. A nested deprecated field stays in
// the Spec, so legacyTarget is still hinted.
func TestReferenceHintsSkipsDeprecatedTopLevelFields(t *testing.T) {
	// Arrange
	msg := deprecatedMessage(t)
	want := []JudgementItem{
		{FieldPath: ".spec.config.target", Reason: "possible-reference-by-description", Detail: `the description has the resource-name template "projects/{project}/topics/{topic}"`},
		{FieldPath: ".spec.config.legacyTarget", Reason: "possible-reference-by-description", Detail: `the description has the resource-name template "projects/{project}/topics/{topic}"`},
		{FieldPath: ".spec.peers[].target", Reason: "possible-reference-by-description", Detail: `the description has the resource-name template "projects/{project}/topics/{topic}"`},
		{FieldPath: ".spec.peers[].legacyTarget", Reason: "possible-reference-by-description", Detail: `the description has the resource-name template "projects/{project}/topics/{topic}"`},
	}

	// Act
	got := ReferenceHints(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReferenceHints() mismatch (-want +got):\n%s", diff)
	}
}

// server_note and server_state both say "Output only." server_note is
// deprecated, so PrepopulateSpec left it out of the Spec and there is nothing
// to move. Only server_state is reported.
func TestDetectOutputOnlySkipsDeprecatedFields(t *testing.T) {
	// Arrange
	msg := deprecatedMessage(t)
	want := []OutputOnlyCandidate{
		{FieldPath: ".spec.serverState", Reason: "well-known-output-only-pattern-in-comment", Comment: "Output only. The state of the widget."},
	}

	// Act
	got := DetectOutputOnlyInComments(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("DetectOutputOnlyInComments() mismatch (-want +got):\n%s", diff)
	}
}

// deprecatedMessage builds a Widget with deprecated fields at the top level and
// in Config. The Spec holds Config directly and in a list, and the deprecated
// legacy_config is a Config too. Some comments match the reference and
// output-only rules:
//
//	message Widget {
//	  string display_name = 1;
//	  // The network the widget joins.
//	  string network = 2 [deprecated = true,
//	    (google.api.resource_reference).type = "compute.googleapis.com/Network"];
//	  Config legacy_config = 3 [deprecated = true];
//	  Config config = 4;
//	  repeated Config peers = 5;
//	  string create_time = 6 [deprecated = true, (google.api.field_behavior) = OUTPUT_ONLY];
//	  // Output only. Set by the server.
//	  string server_note = 7 [deprecated = true];
//	  // Output only. The state of the widget.
//	  string server_state = 8;
//	}
//
//	message Config {
//	  // Format: projects/{project}/topics/{topic}
//	  string target = 1;
//	  int32 top_k = 2 [deprecated = true];
//	  // Format: projects/{project}/topics/{topic}
//	  string legacy_target = 3 [deprecated = true];
//	}
func deprecatedMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	const template = "Format: projects/{project}/topics/{topic}"
	str := fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)
	msgType := fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	repeated := descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	config := strPtr(".google.cloud.test.v1.Config")

	networkOpts := markDeprecated(nil)
	proto.SetExtension(networkOpts, annotations.E_ResourceReference,
		&annotations.ResourceReference{Type: "compute.googleapis.com/Network"})

	comment := func(path []int32, text string) *descriptorpb.SourceCodeInfo_Location {
		return &descriptorpb.SourceCodeInfo_Location{
			Path:            path,
			Span:            []int32{0, 0, 1},
			LeadingComments: strPtr(" " + text + "\n"),
		}
	}

	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("deprecated.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: strPtr("Widget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("display_name"), Number: i32Ptr(1), Type: str},
					{Name: strPtr("network"), Number: i32Ptr(2), Type: str, Options: networkOpts},
					{Name: strPtr("legacy_config"), Number: i32Ptr(3), Type: msgType, TypeName: config, Options: markDeprecated(nil)},
					{Name: strPtr("config"), Number: i32Ptr(4), Type: msgType, TypeName: config},
					{Name: strPtr("peers"), Number: i32Ptr(5), Type: msgType, TypeName: config, Label: repeated},
					{Name: strPtr("create_time"), Number: i32Ptr(6), Type: str, Options: markDeprecated(behaviorOptions(annotations.FieldBehavior_OUTPUT_ONLY))},
					{Name: strPtr("server_note"), Number: i32Ptr(7), Type: str, Options: markDeprecated(nil)},
					{Name: strPtr("server_state"), Number: i32Ptr(8), Type: str},
				},
			},
			{
				Name: strPtr("Config"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("target"), Number: i32Ptr(1), Type: str},
					{Name: strPtr("top_k"), Number: i32Ptr(2), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_INT32), Options: markDeprecated(nil)},
					{Name: strPtr("legacy_target"), Number: i32Ptr(3), Type: str, Options: markDeprecated(nil)},
				},
			},
		},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: []*descriptorpb.SourceCodeInfo_Location{
			comment([]int32{4, 0, 2, 1}, "The network the widget joins."),
			comment([]int32{4, 0, 2, 6}, "Output only. Set by the server."),
			comment([]int32{4, 0, 2, 7}, "Output only. The state of the widget."),
			comment([]int32{4, 1, 2, 0}, template),
			comment([]int32{4, 1, 2, 2}, template),
		}},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("Widget")
}

// markDeprecated sets [deprecated = true] on o, or on new options if o is nil.
func markDeprecated(o *descriptorpb.FieldOptions) *descriptorpb.FieldOptions {
	if o == nil {
		o = &descriptorpb.FieldOptions{}
	}
	o.Deprecated = proto.Bool(true)
	return o
}
