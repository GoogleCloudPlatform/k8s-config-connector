// Copyright 2025 Google LLC
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

package managementconflict

import (
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(t *testing.T) {
	tests := []struct {
		Name                                  string
		ManagementConflictNamespaceAnnotation string
		ManagementConflictObjectAnnotation    string
		MetadataMappingLabels                 string
		LabelsFieldIsMutable                  bool
		ExpectedObjectAnnotation              string
		ShouldSucceed                         bool
	}{
		{
			Name:                                  "none policy on namespace, empty on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "none policy on namespace, resource on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "resource",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  true,
			ExpectedObjectAnnotation:              "resource",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "none policy on namespace, none on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "none",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on namespace, empty on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  true,
			ExpectedObjectAnnotation:              "resource",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on namespace, resource on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "resource",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  true,
			ExpectedObjectAnnotation:              "resource",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on namespace, none on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "none",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  true,
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on namespace with no labels support should default to none",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on namespace with immutable labels should default to none",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  false,
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "resource policy on object should require labels support",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "resource",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "resource",
			ShouldSucceed:                         false,
		},
		{
			Name:                                  "resource policy on object should require mutable labels",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "resource",
			MetadataMappingLabels:                 "labels_field",
			LabelsFieldIsMutable:                  false,
			ExpectedObjectAnnotation:              "resource",
			ShouldSucceed:                         false,
		},
		{
			Name:                                  "invalid policy on namespace",
			ManagementConflictNamespaceAnnotation: "invalid",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "",
			ShouldSucceed:                         false,
		},
		{
			Name:                                  "invalid policy on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "invalid",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "invalid",
			ShouldSucceed:                         false,
		},
		{
			Name:                                  "no value on namespace or resource with no labels support (i.e. default behavior when the resource doesn't support labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "",
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "no value on namespace or resource with immutable labels (i.e. default behavior when the resource doesn't support mutable labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "",
			LabelsFieldIsMutable:                  false,
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
		{
			Name:                                  "no value on namespace or resource with mutable labels (i.e. default behavior when the resource supports mutable labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			MetadataMappingLabels:                 "labels_value",
			LabelsFieldIsMutable:                  true,
			ExpectedObjectAnnotation:              "none",
			ShouldSucceed:                         true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			ns := corev1.Namespace{}
			ns.SetName("my-namespace")
			ns.SetAnnotations(newManagementConflictAnnotations(tc.ManagementConflictNamespaceAnnotation))
			obj := unstructured.Unstructured{}
			obj.SetAnnotations(newManagementConflictAnnotations(tc.ManagementConflictObjectAnnotation))

			fakeTFResourceName := "google_fake_resource"
			fakeTFResource := &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{},
			}
			fakeTFLabelsField := tc.MetadataMappingLabels
			if fakeTFLabelsField != "" {
				fakeTFResource.Schema[fakeTFLabelsField] = &tfschema.Schema{
					ForceNew: !tc.LabelsFieldIsMutable,
				}
			}
			fakeTFProvider := &tfschema.Provider{
				ResourcesMap: map[string]*tfschema.Resource{
					fakeTFResourceName: fakeTFResource,
				},
			}
			rc := corekccv1alpha1.ResourceConfig{
				Name: fakeTFResourceName,
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: fakeTFLabelsField,
				},
			}

			err := ValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(&obj, &ns, &rc, fakeTFProvider.ResourcesMap)
			if tc.ShouldSucceed != (err == nil) {
				t.Fatalf("expected success to be '%v', instead got error mismsatch: %v", tc.ShouldSucceed, err)
			}
			value, ok := obj.GetAnnotations()[FullyQualifiedAnnotation]
			if ok || tc.ExpectedObjectAnnotation != "" {
				if value != tc.ExpectedObjectAnnotation {
					t.Fatalf("unexpected management conflict annotation value: got '%v', want '%v'", value, tc.ExpectedObjectAnnotation)
				}
			}
		})
	}
}

func TestValidateOrDefaultManagementConflictPreventionAnnotationForDCLBasedResource(t *testing.T) {
	tests := []struct {
		Name                                  string
		ManagementConflictNamespaceAnnotation string
		ManagementConflictObjectAnnotation    string
		Schema                                *openapi.Schema
		ExpectedObjectAnnotation              string
		ShouldSucceed                         bool
	}{
		{
			Name:                                  "none policy on namespace, empty on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "none policy on namespace, resource on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "resource",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "resource",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "none policy on namespace, none on object",
			ManagementConflictNamespaceAnnotation: "none",
			ManagementConflictObjectAnnotation:    "none",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on namespace, empty on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "resource",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on namespace, resource on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "resource",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "resource",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on namespace, none on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "none",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on namespace with no labels support should default to none",
			ManagementConflictNamespaceAnnotation: "resource",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on namespace with immutable labels should default to none",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "resource policy on object should require labels support",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "resource",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "resource",
			ShouldSucceed:            false,
		},
		{
			Name:                                  "resource policy on object should require mutable labels",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "resource",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "resource",
			ShouldSucceed:            false,
		},
		{
			Name:                                  "invalid policy on namespace",
			ManagementConflictNamespaceAnnotation: "invalid",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "",
			ShouldSucceed:            false,
		},
		{
			Name:                                  "invalid policy on object",
			ManagementConflictNamespaceAnnotation: "resource",
			ManagementConflictObjectAnnotation:    "invalid",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "invalid",
			ShouldSucceed:            false,
		},
		{
			Name:                                  "no value on namespace or resource with no labels support (i.e. default behavior when the resource doesn't support labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "no value on namespace or resource with immutable labels (i.e. default behavior when the resource doesn't support mutable labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
		{
			Name:                                  "no value on namespace or resource with mutable labels (i.e. default behavior when the resource supports mutable labels)",
			ManagementConflictNamespaceAnnotation: "",
			ManagementConflictObjectAnnotation:    "",
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
			ExpectedObjectAnnotation: "none",
			ShouldSucceed:            true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			ns := corev1.Namespace{}
			ns.SetName("my-namespace")
			ns.SetAnnotations(newManagementConflictAnnotations(tc.ManagementConflictNamespaceAnnotation))
			obj := unstructured.Unstructured{}
			obj.SetAnnotations(newManagementConflictAnnotations(tc.ManagementConflictObjectAnnotation))

			err := ValidateOrDefaultManagementConflictPreventionAnnotationForDCLBasedResource(&obj, &ns, tc.Schema)
			if tc.ShouldSucceed != (err == nil) {
				t.Fatalf("expected success to be '%v', instead got error mismsatch: %v", tc.ShouldSucceed, err)
			}
			value, ok := k8s.GetAnnotation(FullyQualifiedAnnotation, &obj)
			if ok || tc.ExpectedObjectAnnotation != "" {
				if value != tc.ExpectedObjectAnnotation {
					t.Fatalf("unexpected management conflict annotation value: got '%v', want '%v'", value, tc.ExpectedObjectAnnotation)
				}
			}
		})
	}
}
