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

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestReferenceHints pins the paths and reasons the walk files. A person finds
// each path in the CRD, so a path must use the CRD's spelling at every depth,
// and a field the Spec does not contain must not be reported.
func TestReferenceHints(t *testing.T) {
	// Arrange
	msg := referenceHintsMessage(t)

	for _, tc := range []struct {
		name string
		opts codegen.WriteOptions
		want []string
	}{
		{
			name: "message maps off leaves the map out of the Spec",
			opts: codegen.WriteOptions{},
			want: []string{
				".spec.network possible-reference-by-name ComputeNetworkRef",
				".spec.config.target possible-reference-by-description ",
				".spec.config.profile possible-reference-by-description-loose ",
				".spec.peers[].target possible-reference-by-description ",
				".spec.peers[].profile possible-reference-by-description-loose ",
			},
		},
		{
			name: "message maps on puts the map's fields under KEY",
			opts: codegen.WriteOptions{EmitMessageMaps: true},
			want: []string{
				".spec.network possible-reference-by-name ComputeNetworkRef",
				".spec.config.target possible-reference-by-description ",
				".spec.config.profile possible-reference-by-description-loose ",
				".spec.peers[].target possible-reference-by-description ",
				".spec.peers[].profile possible-reference-by-description-loose ",
				".spec.byKey.KEY.target possible-reference-by-description ",
				".spec.byKey.KEY.profile possible-reference-by-description-loose ",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			items := ReferenceHints(msg, tc.opts)

			// Assert
			var got []string
			for _, it := range items {
				got = append(got, it.FieldPath+" "+it.Reason+" "+it.Detail)
			}
			if !slices.Equal(got, tc.want) {
				t.Errorf("ReferenceHints() =\n%q\nwant\n%q", got, tc.want)
			}
		})
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
