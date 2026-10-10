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
	"slices"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestReferenceHints pins the paths, reasons and details the walk files. A
// person finds each path in the CRD, so a path must use the CRD's spelling at
// every depth, and a field the Spec does not contain must not be reported.
func TestReferenceHints(t *testing.T) {
	// Arrange
	msg := referenceHintsMessage(t)
	const templateRule = `the description has the resource-name template "projects/{project}/topics/{topic}"`

	for _, tc := range []struct {
		name string
		opts codegen.WriteOptions
		want []JudgementItem
	}{
		{
			name: "message maps off leaves the map out of the Spec",
			opts: codegen.WriteOptions{},
			want: []JudgementItem{
				{FieldPath: ".spec.network", Reason: "possible-reference-by-name", Detail: "ComputeNetworkRef"},
				{FieldPath: ".spec.config.target", Reason: "possible-reference-by-description", Detail: templateRule},
				{FieldPath: ".spec.config.profile", Reason: "possible-reference-by-description-loose"},
				{FieldPath: ".spec.peers[].target", Reason: "possible-reference-by-description", Detail: templateRule},
				{FieldPath: ".spec.peers[].profile", Reason: "possible-reference-by-description-loose"},
			},
		},
		{
			name: "message maps on puts the map's fields under KEY",
			opts: codegen.WriteOptions{EmitMessageMaps: true},
			want: []JudgementItem{
				{FieldPath: ".spec.network", Reason: "possible-reference-by-name", Detail: "ComputeNetworkRef"},
				{FieldPath: ".spec.config.target", Reason: "possible-reference-by-description", Detail: templateRule},
				{FieldPath: ".spec.config.profile", Reason: "possible-reference-by-description-loose"},
				{FieldPath: ".spec.peers[].target", Reason: "possible-reference-by-description", Detail: templateRule},
				{FieldPath: ".spec.peers[].profile", Reason: "possible-reference-by-description-loose"},
				{FieldPath: ".spec.byKey.KEY.target", Reason: "possible-reference-by-description", Detail: templateRule},
				{FieldPath: ".spec.byKey.KEY.profile", Reason: "possible-reference-by-description-loose"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := ReferenceHints(msg, tc.opts)

			// Assert
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ReferenceHints() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// TestReferenceHintsQueuesNotRepresentableFields checks the entry for a field
// that Classify calls NotRepresentable. outputURI's comment also matches the
// loose rule, but a loose hint would contradict the entry, so the field gets
// just the one entry. gcsSource matches no other rule.
func TestReferenceHintsQueuesNotRepresentableFields(t *testing.T) {
	// Arrange
	msg := namedCommentedMessage(t, [][2]string{
		{"output_uri", "The resource name of the Cloud Storage object to write, such as gs://bucket/object."},
		{"gcs_source", "URI of an object in Google Cloud Storage. Format: gs://{bucket}/{object}"},
	})
	want := []JudgementItem{
		{
			FieldPath: ".spec.outputURI",
			Reason:    "reference-not-representable",
			Detail:    "gcs-path-decomposable-as-bucketref-plus-path: names another resource, but KCC cannot express it as a reference today, so it stays a string",
		},
		{
			FieldPath: ".spec.gcsSource",
			Reason:    "reference-not-representable",
			Detail:    "gcs-path-decomposable-as-bucketref-plus-path: names another resource, but KCC cannot express it as a reference today, so it stays a string",
		},
	}

	// Act
	got := ReferenceHints(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReferenceHints() mismatch (-want +got):\n%s", diff)
	}
}

// TestReferenceHintsQueuesSensitiveFields checks the sensitive-field entry.
// The rule comes from TestNoSensitiveField, so it matches at any depth and in
// any case, but only at the end of the path: passwordPolicy gets no entry.
func TestReferenceHintsQueuesSensitiveFields(t *testing.T) {
	// Arrange
	msg := scanConfigMessage(t)
	const detail = "the name says this holds a password, and it is generated as a plain value. Leave this open until we settle how new Kinds take secrets"
	want := []JudgementItem{
		{FieldPath: ".spec.authentication.googleAccount.password", Reason: "sensitive-field", Detail: detail},
		{FieldPath: ".spec.authentication.customAccount.password", Reason: "sensitive-field", Detail: detail},
		{FieldPath: ".spec.rootPassword", Reason: "sensitive-field", Detail: detail},
	}

	// Act
	got := ReferenceHints(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReferenceHints() mismatch (-want +got):\n%s", diff)
	}
}

// TestReferenceHintsQueuesNestedResourceReferences checks the entry for a
// nested field with google.api.resource_reference. It gets the same
// possible-reference detail PrepopulateSpec gives a top-level field.
// ragCorpus's description also matches Classify, and both entries stay
// because they are separate signals. The top-level corpus and parent fields
// get no entry here, since PrepopulateSpec queues them; see
// TestPrepopulateSpecQueuesResourceReferences.
func TestReferenceHintsQueuesNestedResourceReferences(t *testing.T) {
	// Arrange
	msg := ragStoreMessage(t)
	want := []JudgementItem{
		{
			FieldPath: ".spec.ragResources[].ragCorpus",
			Reason:    "possible-reference",
			Detail:    "the proto marks this field as a reference to aiplatform.googleapis.com/RagCorpus (google.api.resource_reference); confirm whether it should be a KCC reference",
		},
		{
			FieldPath: ".spec.ragResources[].ragCorpus",
			Reason:    "possible-reference-by-description",
			Detail:    `the description has the resource-name template "projects/{project}/locations/{location}/ragCorpora/{rag_corpus}"`,
		},
		{
			FieldPath: ".spec.ragResources[].parent",
			Reason:    "possible-reference",
			Detail:    "the proto marks this field as the parent of a aiplatform.googleapis.com/RagFile (google.api.resource_reference child_type); confirm whether it should be a KCC reference",
		},
	}

	// Act
	got := ReferenceHints(msg, codegen.WriteOptions{})

	// Assert
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReferenceHints() mismatch (-want +got):\n%s", diff)
	}
}

// TestReferenceHintsSkipsGeneratedReferences pins which fields the walk skips
// because the generator writes them as references. connectors' Secret becomes
// a SecretRef, so clientSecret gets no hint. A proto message that is merely
// named ModelRef is an ordinary struct, so the walk still descends into it and
// hints its network field.
func TestReferenceHintsSkipsGeneratedReferences(t *testing.T) {
	// Arrange
	message := fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("connectors.proto"),
		Package: strPtr("google.cloud.connectors.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name:  strPtr("Secret"),
				Field: []*descriptorpb.FieldDescriptorProto{{Name: strPtr("secret_version"), Number: i32Ptr(1), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)}},
			},
			{
				Name:  strPtr("ModelRef"),
				Field: []*descriptorpb.FieldDescriptorProto{{Name: strPtr("network"), Number: i32Ptr(1), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)}},
			},
			{
				Name: strPtr("Connection"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("client_secret"), Number: i32Ptr(1), Type: message, TypeName: strPtr(".google.cloud.connectors.v1.Secret")},
					{Name: strPtr("api_secret"), Number: i32Ptr(2), Type: fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)},
					{Name: strPtr("model_ref"), Number: i32Ptr(3), Type: message, TypeName: strPtr(".google.cloud.connectors.v1.ModelRef")},
				},
			},
		},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}

	// Act
	items := ReferenceHints(fd.Messages().ByName("Connection"), codegen.WriteOptions{})

	// Assert
	var got []string
	for _, it := range items {
		got = append(got, it.FieldPath)
	}
	if want := []string{".spec.apiSecret", ".spec.modelRef.network"}; !slices.Equal(got, want) {
		t.Errorf("hinted paths = %q, want %q", got, want)
	}
}

// referenceHintsMessage builds a Widget whose Spec holds a Config three ways:
// directly, in a list and in a map. Config contains itself, and carries an
// OUTPUT_ONLY field whose comment the rules would otherwise match. Widget's own
// name and create_time carry such comments too, and the Spec drops both.
func referenceHintsMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	const template = "Format: projects/{project}/topics/{topic}"
	str := fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)
	msgType := fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	repeated := descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()

	comment := func(path []int32, text string) *descriptorpb.SourceCodeInfo_Location {
		return &descriptorpb.SourceCodeInfo_Location{
			Path:            path,
			Span:            []int32{0, 0, 1},
			LeadingComments: strPtr(" " + text + "\n"),
		}
	}

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    strPtr("hints.proto"),
		Package: strPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: strPtr("Widget"),
				NestedType: []*descriptorpb.DescriptorProto{{
					Name:    strPtr("ByKeyEntry"),
					Options: &descriptorpb.MessageOptions{MapEntry: func() *bool { b := true; return &b }()},
					Field: []*descriptorpb.FieldDescriptorProto{
						{Name: strPtr("key"), Number: i32Ptr(1), Type: str},
						{Name: strPtr("value"), Number: i32Ptr(2), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config")},
					},
				}},
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("network"), Number: i32Ptr(1), Type: str},
					{Name: strPtr("name"), Number: i32Ptr(2), Type: str},
					{Name: strPtr("create_time"), Number: i32Ptr(3), Type: str, Options: behaviorOptions(annotations.FieldBehavior_OUTPUT_ONLY)},
					{Name: strPtr("config"), Number: i32Ptr(4), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config")},
					{Name: strPtr("peers"), Number: i32Ptr(5), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config"), Label: repeated},
					{Name: strPtr("by_key"), Number: i32Ptr(6), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Widget.ByKeyEntry"), Label: repeated},
				},
			},
			{
				Name: strPtr("Config"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("target"), Number: i32Ptr(1), Type: str},
					{Name: strPtr("child"), Number: i32Ptr(2), Type: msgType, TypeName: strPtr(".google.cloud.test.v1.Config")},
					{Name: strPtr("profile"), Number: i32Ptr(3), Type: str},
					{Name: strPtr("server_topic"), Number: i32Ptr(4), Type: str, Options: behaviorOptions(annotations.FieldBehavior_OUTPUT_ONLY)},
				},
			},
		},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: []*descriptorpb.SourceCodeInfo_Location{
			comment([]int32{4, 0, 2, 0}, "The network the widget joins."),
			comment([]int32{4, 0, 2, 1}, template),
			comment([]int32{4, 0, 2, 2}, template),
			comment([]int32{4, 1, 2, 0}, template),
			comment([]int32{4, 1, 2, 2}, "The resource name (URI) of the destination connection profile."),
			comment([]int32{4, 1, 2, 3}, template),
		}},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("Widget")
}

// scanConfigMessage builds a cut-down websecurityscanner ScanConfig. Each of
// its two account messages has a password, which is INPUT_ONLY and so stays
// in the Spec. Two top-level fields have password in their names, one at the
// end and one at the start.
func scanConfigMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	str := fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)
	msgType := fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE)
	account := func(name string) *descriptorpb.DescriptorProto {
		return &descriptorpb.DescriptorProto{
			Name: strPtr(name),
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: strPtr("username"), Number: i32Ptr(1), Type: str},
				{Name: strPtr("password"), Number: i32Ptr(2), Type: str, Options: behaviorOptions(annotations.FieldBehavior_REQUIRED, annotations.FieldBehavior_INPUT_ONLY)},
			},
		}
	}
	const pkg = ".google.cloud.websecurityscanner.v1.ScanConfig."
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("scan_config.proto"),
		Package: strPtr("google.cloud.websecurityscanner.v1"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: strPtr("ScanConfig"),
			NestedType: []*descriptorpb.DescriptorProto{{
				Name:       strPtr("Authentication"),
				NestedType: []*descriptorpb.DescriptorProto{account("GoogleAccount"), account("CustomAccount")},
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("google_account"), Number: i32Ptr(1), Type: msgType, TypeName: strPtr(pkg + "Authentication.GoogleAccount")},
					{Name: strPtr("custom_account"), Number: i32Ptr(2), Type: msgType, TypeName: strPtr(pkg + "Authentication.CustomAccount")},
				},
			}},
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: strPtr("authentication"), Number: i32Ptr(1), Type: msgType, TypeName: strPtr(pkg + "Authentication")},
				{Name: strPtr("root_password"), Number: i32Ptr(2), Type: str},
				{Name: strPtr("password_policy"), Number: i32Ptr(3), Type: str},
			},
		}},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("ScanConfig")
}

// ragStoreMessage builds a cut-down aiplatform VertexRagStore. At the top
// level it has one field for each form of google.api.resource_reference:
// corpus with type and parent with child_type. The nested RagResource has
// one of each as well. ragCorpus also keeps its real comment, which spells
// out a resource-name template.
func ragStoreMessage(t *testing.T) protoreflect.MessageDescriptor {
	t.Helper()
	str := fieldType(descriptorpb.FieldDescriptorProto_TYPE_STRING)
	reference := func(rr *annotations.ResourceReference) *descriptorpb.FieldOptions {
		o := &descriptorpb.FieldOptions{}
		proto.SetExtension(o, annotations.E_ResourceReference, rr)
		return o
	}
	toCorpus := &annotations.ResourceReference{Type: "aiplatform.googleapis.com/RagCorpus"}
	parentOfFile := &annotations.ResourceReference{ChildType: "aiplatform.googleapis.com/RagFile"}
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    strPtr("vertex_rag_store.proto"),
		Package: strPtr("google.cloud.aiplatform.v1"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: strPtr("VertexRagStore"),
			NestedType: []*descriptorpb.DescriptorProto{{
				Name: strPtr("RagResource"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("rag_corpus"), Number: i32Ptr(1), Type: str, Options: reference(toCorpus)},
					{Name: strPtr("parent"), Number: i32Ptr(2), Type: str, Options: reference(parentOfFile)},
				},
			}},
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: strPtr("corpus"), Number: i32Ptr(1), Type: str, Options: reference(toCorpus)},
				{Name: strPtr("parent"), Number: i32Ptr(2), Type: str, Options: reference(parentOfFile)},
				{
					Name:     strPtr("rag_resources"),
					Number:   i32Ptr(3),
					Type:     fieldType(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE),
					TypeName: strPtr(".google.cloud.aiplatform.v1.VertexRagStore.RagResource"),
					Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
				},
			},
		}},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: []*descriptorpb.SourceCodeInfo_Location{{
			// [4=message_type 0, 3=nested_type 0, 2=field 0] is RagResource.rag_corpus.
			Path:            []int32{4, 0, 3, 0, 2, 0},
			Span:            []int32{0, 0, 1},
			LeadingComments: strPtr(" Optional. RagCorpora resource name.\n Format:\n `projects/{project}/locations/{location}/ragCorpora/{rag_corpus}`\n"),
		}}},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName("VertexRagStore")
}
