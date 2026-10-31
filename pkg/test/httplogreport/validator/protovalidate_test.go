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

package validator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValidateProtoSchemas_SyntheticDescriptorSet(t *testing.T) {
	pbPath := writeSyntheticDescriptorSet(t)
	fix := &fixture.Data{
		PrimaryCreate: &fixture.KRMResource{
			Kind:  "ExampleWidget",
			Group: "example.cnrm.cloud.google.com",
			Name:  "widget-${uniqueId}",
			Obj:   &unstructured.Unstructured{},
		},
	}

	validLog := `# Phase: CREATE
POST https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets?%24alt=json%3Benum-encoding%3Dint&widgetId=widget-${uniqueId}
Content-Type: application/json

{
  "description": "initial",
  "state": 1
}

200 OK

{
  "done": false,
  "metadata": {
    "@type": "type.googleapis.com/google.example.v1.OperationMetadata",
    "target": "projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}",
    "verb": "create"
  },
  "name": "projects/${projectId}/locations/us-central1/operations/operation-1"
}

---

GET https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/operations/operation-1

200 OK

{
  "done": true,
  "metadata": {
    "@type": "type.googleapis.com/google.example.v1.OperationMetadata",
    "target": "projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}",
    "verb": "create"
  },
  "name": "projects/${projectId}/locations/us-central1/operations/operation-1",
  "response": {
    "@type": "type.googleapis.com/google.example.v1.Widget",
    "description": "initial",
    "name": "projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}",
    "state": 1
  }
}

---

GET https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}?%24alt=json%3Benum-encoding%3Dint

200 OK

{
  "description": "initial",
  "name": "projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}",
  "state": 1
}

---

# Phase: DELETE
DELETE https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets/widget-${uniqueId}

200 OK

{}
`

	tests := []struct {
		name          string
		rawLog        string
		wantErrSubstr string
	}{
		{
			name:          "valid log passes proto schema",
			rawLog:        validLog,
			wantErrSubstr: "",
		},
		{
			name:          "hallucinated field caught",
			rawLog:        strings.Replace(validLog, `"description": "initial"`, `"description": "initial",`+"\n"+`  "hallucinatedField": "foo"`, 1),
			wantErrSubstr: "hallucinatedField",
		},
		{
			name:          "unknown lro type caught",
			rawLog:        strings.Replace(validLog, "type.googleapis.com/google.example.v1.OperationMetadata", "type.googleapis.com/google.example.v1.FakeHallucinatedMetadata", 1),
			wantErrSubstr: "unknown protobuf `@type`",
		},
		{
			name:          "string enum with enum-encoding=int caught",
			rawLog:        strings.ReplaceAll(validLog, `"state": 1`, `"state": "ACTIVE"`),
			wantErrSubstr: "enum-encoding=int",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			rawEntries, err := logparser.Parse(tc.rawLog)
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}
			classified, _ := classifier.New(fix).Classify(rawEntries)

			// Act
			errs := ValidateProtoSchemas(classified, pbPath)

			// Assert
			if tc.wantErrSubstr == "" {
				if len(errs) > 0 {
					t.Fatalf("expected 0 proto errors, got: %v", errs)
				}
				return
			}
			if len(errs) == 0 {
				t.Fatalf("expected ValidateProtoSchemas to catch %q, but got 0 errors", tc.wantErrSubstr)
			}
			if !strings.Contains(strings.Join(errs, " "), tc.wantErrSubstr) {
				t.Fatalf("expected error message to mention %q, got: %v", tc.wantErrSubstr, errs)
			}
		})
	}
}

func writeSyntheticDescriptorSet(t *testing.T) string {
	t.Helper()

	createSvcOpts := &descriptorpb.ServiceOptions{}
	proto.SetExtension(createSvcOpts, annotations.E_DefaultHost, "example.googleapis.com")

	getOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(getOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Get{
			Get: "/v1/{name=projects/*/locations/*/widgets/*}",
		},
	})

	createOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(createOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Post{
			Post: "/v1/{parent=projects/*/locations/*}/widgets",
		},
		Body: "widget",
	})

	updateOpts := &descriptorpb.MethodOptions{}
	proto.SetExtension(updateOpts, annotations.E_Http, &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Patch{
			Patch: "/v1/{widget.name=projects/*/locations/*/widgets/*}",
		},
		Body: "widget",
	})

	fds := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{
			{
				Name:    proto.String("google/protobuf/empty.proto"),
				Package: proto.String("google.protobuf"),
				Syntax:  proto.String("proto3"),
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("Empty"),
					},
				},
			},
			{
				Name:    proto.String("google/longrunning/operations.proto"),
				Package: proto.String("google.longrunning"),
				Syntax:  proto.String("proto3"),
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("Operation"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("name"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("name"),
							},
							{
								Name:     proto.String("done"),
								Number:   proto.Int32(3),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum(),
								JsonName: proto.String("done"),
							},
						},
					},
				},
			},
			{
				Name:       proto.String("google/example/v1/widget.proto"),
				Package:    proto.String("google.example.v1"),
				Syntax:     proto.String("proto3"),
				Dependency: []string{"google/longrunning/operations.proto", "google/protobuf/empty.proto"},
				EnumType: []*descriptorpb.EnumDescriptorProto{
					{
						Name: proto.String("WidgetState"),
						Value: []*descriptorpb.EnumValueDescriptorProto{
							{Name: proto.String("WIDGET_STATE_UNSPECIFIED"), Number: proto.Int32(0)},
							{Name: proto.String("ACTIVE"), Number: proto.Int32(1)},
						},
					},
				},
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("Widget"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("name"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("name"),
							},
							{
								Name:     proto.String("description"),
								Number:   proto.Int32(2),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("description"),
							},
							{
								Name:     proto.String("state"),
								Number:   proto.Int32(3),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
								TypeName: proto.String(".google.example.v1.WidgetState"),
								JsonName: proto.String("state"),
							},
						},
					},
					{
						Name: proto.String("OperationMetadata"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("target"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("target"),
							},
							{
								Name:     proto.String("verb"),
								Number:   proto.Int32(2),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("verb"),
							},
						},
					},
					{
						Name: proto.String("GetWidgetRequest"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("name"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("name"),
							},
						},
					},
					{
						Name: proto.String("CreateWidgetRequest"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("parent"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
								JsonName: proto.String("parent"),
							},
							{
								Name:     proto.String("widget"),
								Number:   proto.Int32(2),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: proto.String(".google.example.v1.Widget"),
								JsonName: proto.String("widget"),
							},
						},
					},
					{
						Name: proto.String("UpdateWidgetRequest"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:     proto.String("widget"),
								Number:   proto.Int32(1),
								Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
								Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
								TypeName: proto.String(".google.example.v1.Widget"),
								JsonName: proto.String("widget"),
							},
						},
					},
				},
				Service: []*descriptorpb.ServiceDescriptorProto{
					{
						Name:    proto.String("WidgetService"),
						Options: createSvcOpts,
						Method: []*descriptorpb.MethodDescriptorProto{
							{
								Name:       proto.String("GetWidget"),
								InputType:  proto.String(".google.example.v1.GetWidgetRequest"),
								OutputType: proto.String(".google.example.v1.Widget"),
								Options:    getOpts,
							},
							{
								Name:       proto.String("CreateWidget"),
								InputType:  proto.String(".google.example.v1.CreateWidgetRequest"),
								OutputType: proto.String(".google.longrunning.Operation"),
								Options:    createOpts,
							},
							{
								Name:       proto.String("UpdateWidget"),
								InputType:  proto.String(".google.example.v1.UpdateWidgetRequest"),
								OutputType: proto.String(".google.longrunning.Operation"),
								Options:    updateOpts,
							},
						},
					},
				},
			},
		},
	}

	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("proto.Marshal(FileDescriptorSet): %v", err)
	}
	pbPath := filepath.Join(t.TempDir(), "synthetic-googleapis.pb")
	if err := os.WriteFile(pbPath, b, 0644); err != nil {
		t.Fatalf("writing %s: %v", pbPath, err)
	}
	return pbPath
}
