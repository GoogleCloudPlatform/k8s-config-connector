// Copyright 2022 Google LLC
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

package webhook

import (
	"container/list"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testutil "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	"github.com/nasa9084/go-openapi"
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/yaml"
)

func TestChangesOnImmutableFields(t *testing.T) {
	t.Parallel()
	p := provider.Provider()
	v := newImmutableFieldsValidatorHandler(t)
	for _, testcase := range TestCases {
		t.Run(testcase.Name, func(t *testing.T) {
			assertImmutableFieldsValidatorResult(t, v, p, testcase)
		})
	}
}

func TestChangesOnImmutableFieldsForDCLResource(t *testing.T) {
	tests := []struct {
		name     string
		obj      *unstructured.Unstructured
		oldObj   *unstructured.Unstructured
		spec     map[string]interface{}
		oldSpec  map[string]interface{}
		schema   *openapi.Schema
		response admission.Response
	}{
		{
			name: "changes on the base level immutable fields",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},

			oldSpec: map[string]interface{}{
				"location": "US",
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"location": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.location"})),
		},
		{
			name: "changes on the resourceID field",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"resourceID": "name1",
			},
			oldSpec: map[string]interface{}{
				"resourceID": "name2",
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"name": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.resourceID"})),
		},
		{
			name: "changes on nested immutable fields",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"location": "EU",
				"nestedObjectKey": map[string]interface{}{
					"nestedIntField":    1,
					"nestedStringField": "strval1",
				},
			},
			oldSpec: map[string]interface{}{
				"location": "EU",
				"nestedObjectKey": map[string]interface{}{
					"nestedIntField":    2,
					"nestedStringField": "strval2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"location": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
					"nestedObjectKey": &openapi.Schema{
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedIntField": &openapi.Schema{
								Type: "integer",
								Extension: map[string]interface{}{
									"x-kubernetes-immutable": true,
								},
							},
							"nestedStringField": &openapi.Schema{
								Type: "string",
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.nestedObjectKey.nestedIntField"})),
		},
		{
			name: "changes on immutable resource reference",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			spec: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"name": "pubsubtopic-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"name": "pubsubtopic-sample-2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"referenceKey": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "FakeService/FakeKind",
									"field":    "name",
								},
							},
							"x-kubernetes-immutable": true,
						},
					},
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.referenceKeyRef"})),
		},
		{
			name: "changes on immutable hierarchical reference for single-parent resource",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5ProjectRef",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5ProjectRef",
				},
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-sample-2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
							},
							"x-kubernetes-immutable": true,
						},
					},
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.projectRef"})),
		},
		{
			name: "changes on immutable hierarchical reference for multi-parent resource (value change)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
							"x-kubernetes-immutable": true,
						},
					},
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.folderRef"})),
		},
		{
			name: "changes on immutable hierarchical reference for multi-parent resource (key change)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"name": "organization-sample-1",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
							"x-kubernetes-immutable": true,
						},
					},
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.folderRef", "spec.organizationRef"})),
		},
		{
			name: "changes on mutable hierarchical reference for single-parent resource",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5ProjectRef",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5ProjectRef",
				},
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-sample-2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
							},
						},
					},
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on mutable hierarchical reference for multi-parent resource (value change)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
						},
					},
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on mutable hierarchical reference for multi-parent resource (key change)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test5MultipleRefs",
				},
			},
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-sample-1",
				},
			},
			oldSpec: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"name": "organization-sample-1",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": &openapi.Schema{
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
						},
					},
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on immutable arrays of primitives",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"stringArrayKey": []string{"val1"},
			},
			oldSpec: map[string]interface{}{
				"stringArrayKey": []string{"val1", "val2"},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"stringArrayKey": &openapi.Schema{
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.stringArrayKey"})),
		},
		{
			name: "changes on immutable maps of primitives",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"stringMapKey": map[string]interface{}{
					"foo": "foo",
				},
			},
			oldSpec: map[string]interface{}{
				"stringMapKey": map[string]interface{}{
					"foo": "bar",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"stringMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.stringMapKey"})),
		},
		{
			name: "no changes to immutable maps of objects",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on immutable maps of objects (changed a value in one of the objects in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "bar1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.objectMapKey"})),
		},
		{
			name: "changes on immutable maps of objects (changed the key of one of the objects in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1-new": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.objectMapKey"})),
		},
		{
			name: "changes on immutable maps of objects (added a new object to the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
					"obj3": map[string]interface{}{
						"objectField": "foo3",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.objectMapKey"})),
		},
		{
			name: "changes on immutable maps of objects (deleted an object from the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.objectMapKey"})),
		},
		{
			name: "changes on mutable maps of objects (changed a field in one of the objects in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "bar1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"objectMapKey": map[string]interface{}{
					"obj1": map[string]interface{}{
						"objectField": "foo1",
					},
					"obj2": map[string]interface{}{
						"objectField": "foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"objectMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField": &openapi.Schema{
									Type: "string",
								},
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "no changes to immutable maps of arrays",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"arrayMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on immutable maps of arrays (changed a value in one of the arrays in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"bar1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"arrayMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.arrayMapKey"})),
		},
		{
			name: "changes on immutable maps of arrays (added a value in one of the arrays in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
						"foo11",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"arrayMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.arrayMapKey"})),
		},
		{
			name: "changes on immutable maps of arrays (deleted a value from one of the arrays in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
						"foo22",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
						"foo11",
					},
					"arr2": []interface{}{
						"foo2",
						"foo22",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"arrayMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-kubernetes-immutable": true,
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"spec.arrayMapKey"})),
		},
		{
			name: "changes on mutable maps of arrays (changed a value in one of the arrays in the map)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"bar1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			oldSpec: map[string]interface{}{
				"arrayMapKey": map[string]interface{}{
					"arr1": []interface{}{
						"foo1",
					},
					"arr2": []interface{}{
						"foo2",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"arrayMapKey": &openapi.Schema{
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes the container annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-2",
						},
					},
				},
			},
			spec:    map[string]interface{}{},
			oldSpec: map[string]interface{}{},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: admission.Errored(http.StatusBadRequest,
				fmt.Errorf("error validating container annotations: cannot make changes to container annotation cnrm.cloud.google.com/project-id")),
		},
		{
			name: "changes on mutable fields",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "my-project-1",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"location": "EU",
				"nestedObjectKey": map[string]interface{}{
					"nestedIntField":    1,
					"nestedStringField": "strval1",
				},
			},
			oldSpec: map[string]interface{}{
				"location": "US",
				"nestedObjectKey": map[string]interface{}{
					"nestedIntField":    2,
					"nestedStringField": "strval2",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"location": &openapi.Schema{
						Type: "string",
					},
					"nestedObjectKey": &openapi.Schema{
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedIntField": &openapi.Schema{
								Type: "integer",
							},
							"nestedStringField": &openapi.Schema{
								Type: "string",
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			response: allowedResponse,
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			dclSchemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.obj.GroupVersionKind(), smLoader)
			dclSchemaMap := make(map[string]*openapi.Schema)
			dclSchemaMap[dclSchemaKey] = tc.schema
			dclSchemaLoader := testdclschemaloader.New(dclSchemaMap)
			actual := validateImmutableFieldsForDCLBasedResource(tc.obj, tc.oldObj, tc.spec, tc.oldSpec, dclSchemaLoader, smLoader)
			if !testutil.Equals(t, actual, tc.response) {
				t.Fatalf("got: %v, but want: %v", actual, tc.response)
			}
		})
	}
}

func assertImmutableFieldsValidatorResult(t *testing.T, _ HandlerFunc, provider *schema.Provider, testCase TestCase) {
	r, ok := provider.ResourcesMap[testCase.TFSchemaName]
	if !ok {
		t.Errorf("couldn't get the schema for %v", testCase.TFSchemaName)
	}
	fields := list.New()
	compareAndFindChangesOnImmutableFields(testCase.Spec, testCase.OldSpec, r.Schema, "", testCase.ResourceConfig, nil, fields)

	res := make([]string, 0)
	for e := fields.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(string))
	}
	if !reflect.DeepEqual(testCase.ExpectedResult, res) {
		t.Errorf("expected to find changes on immutable location field %v, instead get %v", testCase.ExpectedResult, res)
	}
}

func TestChangesOnImmutableLocationField(t *testing.T) {
	spec := map[string]interface{}{
		"location": "us-east1",
	}

	oldSpec := map[string]interface{}{
		"location": "us-west1",
	}

	rc := &corekccv1alpha1.ResourceConfig{
		Locationality: "regional",
	}

	found := findChangesOnImmutableLocationField(spec, oldSpec, rc)
	if !found {
		t.Errorf("expected to find changes on immutable location field")
	}
}

func TestChangesOnImmutableResourceIDField(t *testing.T) {
	tests := []struct {
		name           string
		spec           map[string]interface{}
		oldSpec        map[string]interface{}
		rc             *corekccv1alpha1.ResourceConfig
		expectedResult bool
	}{
		{
			name:    "resource ID not changed",
			spec:    map[string]interface{}{k8s.ResourceIDFieldName: "test-id"},
			oldSpec: map[string]interface{}{k8s.ResourceIDFieldName: "test-id"},
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
			},
			expectedResult: false,
		},
		{
			name:    "resource ID changed",
			spec:    map[string]interface{}{k8s.ResourceIDFieldName: "updated-id"},
			oldSpec: map[string]interface{}{k8s.ResourceIDFieldName: "test-id"},
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
			},
			expectedResult: true,
		},
		{
			name:           "resource ID not supported",
			spec:           map[string]interface{}{},
			oldSpec:        map[string]interface{}{},
			rc:             &corekccv1alpha1.ResourceConfig{},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got, want :=
				findChangesOnImmutableResourceIDField(tc.spec, tc.oldSpec, tc.rc),
				tc.expectedResult; got != want {
				t.Errorf("unexpected result finding changes on %q "+
					"field: got %t, want %t", k8s.ResourceIDFieldPath, got, want)
			}
		})
	}
}

func newImmutableFieldsValidatorHandler(t *testing.T) HandlerFunc {
	t.Helper()
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatal(err)
	}

	return NewImmutableFieldsValidatorHandler(smLoader, nil, testservicemetadataloader.NewForUnitTest())
}

func TestValidateContainerAnnotations(t *testing.T) {
	tests := []struct {
		name             string
		kind             string
		old              map[string]string
		updated          map[string]string
		containers       []corekccv1alpha1.Container
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference
		shouldErr        bool
	}{
		{
			name:    "changing project ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-2"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			shouldErr: true,
		},
		{
			name:    "changing folder ID is not allowed (in the general case)",
			kind:    "GenericKind",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing folder ID is allowed for Projects",
			kind:    "Project",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
		},
		{
			name:    "changing folder ID is allowed for Folders",
			kind:    "Folder",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
		},
		{
			name:    "changing org ID is not allowed (in the general case)",
			kind:    "GenericKind",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.OrgIDAnnotation: "1234567"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing org ID is allowed for Projects",
			kind:    "Project",
			old:     map[string]string{k8s.OrgIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
		},
		{
			name:    "changing org ID is allowed for Folders",
			kind:    "Folder",
			old:     map[string]string{k8s.OrgIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
		},
		{
			name:    "changing from project ID to folder ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{k8s.FolderIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing from project ID to org ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{k8s.OrgIDAnnotation: "0987654"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from folder ID to project ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing from folder ID to org ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "0987654"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from org ID to project ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from org ID to folder ID is not allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.FolderIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from folder ID to org ID is not allowed for Projects",
			kind:    "Project",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "0987654"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from org ID to folder ID is not allowed for Projects",
			kind:    "Project",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.FolderIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from folder ID to org ID is not allowed for Folders",
			kind:    "Folder",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "0987654"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing from org ID to folder ID is not allowed for Folders",
			kind:    "Folder",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.FolderIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing project ID is not allowed if if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-2"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeProject},
			},
			shouldErr: true,
		},
		{
			name:    "changing folder ID is not allowed if resource supports hierarchical references (in the general case)",
			kind:    "GenericKind",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing org ID is not allowed if resource supports hierarchical references (in the general case)",
			kind:    "GenericKind",
			old:     map[string]string{k8s.OrgIDAnnotation: "0987654"},
			updated: map[string]string{k8s.OrgIDAnnotation: "1234567"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing folder ID is not allowed for Projects once Project supports hierarchical references",
			kind:    "Project",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing folder ID is not allowed for Folders once Folder supports hierarchical references",
			kind:    "Folder",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{k8s.FolderIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "changing org ID is not allowed for Projects once Project supports hierarchical references",
			kind:    "Project",
			old:     map[string]string{k8s.OrgIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "changing org ID is not allowed for Folders once Folder supports hierarchical references",
			kind:    "Folder",
			old:     map[string]string{k8s.OrgIDAnnotation: "123321"},
			updated: map[string]string{k8s.OrgIDAnnotation: "321123"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name:    "removing project ID is allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeProject},
			},
		},
		{
			name:    "removing folder ID is allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{k8s.FolderIDAnnotation: "123321"},
			updated: map[string]string{},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeFolder},
			},
		},
		{
			name:    "removing org ID is allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{k8s.OrgIDAnnotation: "123321"},
			updated: map[string]string{},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization},
			},
		},
		{
			name:    "adding project ID is not allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeProject},
			},
			shouldErr: true,
		},
		{
			name:    "adding folder ID is not allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{},
			updated: map[string]string{k8s.FolderIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeFolder},
			},
			shouldErr: true,
		},
		{
			name:    "adding org ID is not allowed if resource supports hierarchical references",
			kind:    "GenericKind",
			old:     map[string]string{},
			updated: map[string]string{k8s.OrgIDAnnotation: "123321"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeOrganization},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization},
			},
			shouldErr: true,
		},
		{
			name: "adding a different annotation is allowed",
			kind: "GenericKind",
			old:  map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{
				k8s.ProjectIDAnnotation: "project-1",
				"key":                   "value",
			},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeProject},
			},
		},
		{
			name:    "changing an unrecognized container annotation is allowed",
			kind:    "GenericKind",
			old:     map[string]string{k8s.ProjectIDAnnotation: "project-1"},
			updated: map[string]string{k8s.ProjectIDAnnotation: "project-2"},
			containers: []corekccv1alpha1.Container{
				{Type: corekccv1alpha1.ContainerTypeFolder},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := validateContainerAnnotationsForResource(tc.kind, tc.updated, tc.old, tc.containers, tc.hierarchicalRefs)
			if tc.shouldErr && err == nil {
				t.Errorf("expected error but there was none")
				return
			} else if !tc.shouldErr && err != nil {
				t.Errorf("got unexpected error: %v", err)
			}
		})
	}
}

func newUnstructuredFromObject(t *testing.T, value interface{}) *unstructured.Unstructured {
	var mapResult map[string]interface{}
	if err := util.Marshal(value, &mapResult); err != nil {
		t.Fatalf("unable to marshal %v to map: %v", reflect.TypeOf(value).Name(), err)
	}
	return &unstructured.Unstructured{
		Object: mapResult,
	}
}

type TestCase struct {
	Name           string
	Spec           map[string]interface{}
	OldSpec        map[string]interface{}
	TFSchemaName   string
	ResourceConfig *corekccv1alpha1.ResourceConfig
	ExpectedResult []string
}

var TestCases = []TestCase{
	{
		Name: "changesOnBaseLevelImmutableField",
		Spec: map[string]interface{}{
			"location": "EU",
		},
		OldSpec: map[string]interface{}{
			"location": "US",
		},
		TFSchemaName:   "google_bigquery_dataset",
		ResourceConfig: &corekccv1alpha1.ResourceConfig{},
		ExpectedResult: []string{"location"},
	},
	{
		Name: "changesOnNestedImmutableField",
		Spec: map[string]interface{}{
			"databaseVersion": "MYSQL_5_7",
			"replicaConfiguration": map[string]interface{}{
				"connectRetryInterval": 2,
			},
		},
		OldSpec: map[string]interface{}{
			"databaseVersion": "MYSQL_5_7",
			"replicaConfiguration": map[string]interface{}{
				"connectRetryInterval": 4,
			},
		},
		TFSchemaName:   "google_sql_database_instance",
		ResourceConfig: &corekccv1alpha1.ResourceConfig{},
		ExpectedResult: []string{"replica_configuration.connect_retry_interval"},
	},
	{
		Name: "changesOnImmutableReferenceSingleton",
		Spec: map[string]interface{}{
			"topicRef": map[string]interface{}{
				"name": "pubsubtopic-sample-1",
			},
		},
		OldSpec: map[string]interface{}{
			"topicRef": map[string]interface{}{
				"name": "pubsubtopic-sample-2",
			},
		},
		TFSchemaName: "google_pubsub_subscription",
		ResourceConfig: &corekccv1alpha1.ResourceConfig{
			ResourceReferences: []corekccv1alpha1.ReferenceConfig{
				{
					TypeConfig: corekccv1alpha1.TypeConfig{
						Key: "topicRef",
						GVK: k8sschema.GroupVersionKind{
							Kind: "PubsubTopic",
						},
					},
					TFField: "topic",
				},
			},
		},
		ExpectedResult: []string{"topicRef"},
	},
	{
		Name: "changesOnMutableNestedReference",
		Spec: map[string]interface{}{
			"peeringConfig": map[string]interface{}{
				"targetNetwork": map[string]interface{}{
					"networkRef": map[string]interface{}{
						"name": "ref1",
					},
				},
			},
		},
		OldSpec: map[string]interface{}{
			"peeringConfig": map[string]interface{}{
				"targetNetwork": map[string]interface{}{
					"networkRef": map[string]interface{}{
						"name": "ref2",
					},
				},
			},
		},
		TFSchemaName: "google_dns_managed_zone",
		ResourceConfig: &corekccv1alpha1.ResourceConfig{
			ResourceReferences: []corekccv1alpha1.ReferenceConfig{
				{
					TypeConfig: corekccv1alpha1.TypeConfig{
						Key: "networkRef",
						GVK: k8sschema.GroupVersionKind{
							Kind: "ComputeNetwork",
						},
					},
					TFField: "peering_config.target_network.network_url",
				},
			},
		},
		ExpectedResult: []string{},
	},
	{
		Name: "changesOnImmutableComplexReference",
		Spec: map[string]interface{}{
			"ipAddress": map[string]interface{}{
				"addressRef": map[string]interface{}{
					"name": "ref1",
				},
			},
		},
		OldSpec: map[string]interface{}{
			"ipAddress": map[string]interface{}{
				"ip": "8.8.8.8",
			},
		},
		TFSchemaName: "google_compute_forwarding_rule",
		ResourceConfig: &corekccv1alpha1.ResourceConfig{
			ResourceReferences: []corekccv1alpha1.ReferenceConfig{
				{
					Types: []corekccv1alpha1.TypeConfig{
						{
							Key: "addressRef",
						},
						{
							Key: "ip",
						},
					},
					TFField: "ip_address",
				},
			},
		},
		ExpectedResult: []string{"ipAddress"},
	},
	// TODO(maqiuyu): Add a test case for resourceID once supported.
}

func TestUpdateIAMPolicy(t *testing.T) {
	policy := v1beta1.IAMPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyGVK.Kind,
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPolicySpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-pubsub-topic",
				APIVersion: "my-api-version",
			},
		},
	}
	oldPolicyUnstructured := newUnstructuredFromObject(t, &policy)
	newPolicyUnstructured := newUnstructuredFromObject(t, &policy)
	assertHandleIAMPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, true)
	copyPolicy := policy
	copyPolicy.Spec.ResourceReference.Kind = "new-resource-reference-kind"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.Namespace = "new-resource-reference-namespace"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.Name = "new-resource-reference-name"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.APIVersion = "new-resource-reference-apiversion"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
}

func assertHandleIAMPolicy(t *testing.T, old *unstructured.Unstructured, new *unstructured.Unstructured, expectedAllowedValue bool) {
	t.Helper()
	oldSpec := getSpecFromUnstructed(t, old)
	newSpec := getSpecFromUnstructed(t, new)
	response := handleIAMPolicy(oldSpec, newSpec)
	if response.Allowed != expectedAllowedValue {
		t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, expectedAllowedValue)
	}
}

func TestUpdateIAMPartialPolicy(t *testing.T) {
	policy := v1beta1.IAMPartialPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyGVK.Kind,
			APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPartialPolicySpec{
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-pubsub-topic",
				APIVersion: "my-api-version",
			},
		},
	}
	oldPolicyUnstructured := newUnstructuredFromObject(t, &policy)
	newPolicyUnstructured := newUnstructuredFromObject(t, &policy)
	assertHandleIAMPartialPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, true)
	copyPolicy := policy
	copyPolicy.Spec.ResourceReference.Kind = "new-resource-reference-kind"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPartialPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.Namespace = "new-resource-reference-namespace"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPartialPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.Name = "new-resource-reference-name"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPartialPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
	copyPolicy = policy
	copyPolicy.Spec.ResourceReference.APIVersion = "new-resource-reference-apiversion"
	newPolicyUnstructured = newUnstructuredFromObject(t, &copyPolicy)
	assertHandleIAMPartialPolicy(t, oldPolicyUnstructured, newPolicyUnstructured, false)
}

func assertHandleIAMPartialPolicy(t *testing.T, old *unstructured.Unstructured, new *unstructured.Unstructured, expectedAllowedValue bool) {
	t.Helper()
	oldSpec := getSpecFromUnstructed(t, old)
	newSpec := getSpecFromUnstructed(t, new)
	response := handleIAMPartialPolicy(oldSpec, newSpec)
	if response.Allowed != expectedAllowedValue {
		t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, expectedAllowedValue)
	}
}

func TestUpdateIAMPolicyMember(t *testing.T) {
	policyMember := v1beta1.IAMPolicyMember{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
			APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMPolicyMemberSpec{
			Member: "test@google.com",
			Role:   "roles/editor",
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-pubsub-topic",
				APIVersion: "my-api-version",
			},
		},
	}
	oldPolicyMemberUnstructured := newUnstructuredFromObject(t, &policyMember)
	newPolicyMemberUnstructured := newUnstructuredFromObject(t, &policyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, true)
	copyPolicyMember := policyMember
	copyPolicyMember.Spec.Member = "new-member"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
	copyPolicyMember = policyMember
	copyPolicyMember.Spec.Role = "new-role"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
	copyPolicyMember = policyMember
	copyPolicyMember.Spec.ResourceReference.Kind = "new-resource-reference-kind"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
	copyPolicyMember = policyMember
	copyPolicyMember.Spec.ResourceReference.Namespace = "new-resource-reference-namespace"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
	copyPolicyMember = policyMember
	copyPolicyMember.Spec.ResourceReference.Name = "new-resource-reference-name"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
	copyPolicyMember = policyMember
	copyPolicyMember.Spec.ResourceReference.APIVersion = "new-resource-reference-apiversion"
	newPolicyMemberUnstructured = newUnstructuredFromObject(t, &copyPolicyMember)
	assertHandleIAMPolicyMember(t, oldPolicyMemberUnstructured, newPolicyMemberUnstructured, false)
}

func assertHandleIAMPolicyMember(t *testing.T, old *unstructured.Unstructured, new *unstructured.Unstructured, expectedAllowedValue bool) {
	t.Helper()
	oldSpec := getSpecFromUnstructed(t, old)
	newSpec := getSpecFromUnstructed(t, new)
	response := handleIAMPolicyMember(oldSpec, newSpec)
	if response.Allowed != expectedAllowedValue {
		t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, expectedAllowedValue)
	}
}

func TestUpdateIAMAuditConfig(t *testing.T) {
	auditConfig := v1beta1.IAMAuditConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1beta1.IAMAuditConfigGVK.Kind,
			APIVersion: v1beta1.IAMAuditConfigGVK.GroupVersion().String(),
		},
		Spec: v1beta1.IAMAuditConfigSpec{
			Service: "sampleservice.googleapis.com",
			AuditLogConfigs: []v1beta1.AuditLogConfig{
				{
					LogType: "DATA_READ",
					ExemptedMembers: []v1beta1.Member{
						"test@google.com",
					},
				},
			},
			ResourceReference: v1beta1.ResourceReference{
				Kind:       "my-resource-kind",
				Namespace:  "my-namespace",
				Name:       "my-pubsub-topic",
				APIVersion: "my-api-version",
			},
		},
	}
	oldAuditConfigUnstructured := newUnstructuredFromObject(t, &auditConfig)
	newAuditConfigUnstructured := newUnstructuredFromObject(t, &auditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, true)
	copyAuditConfig := auditConfig
	copyAuditConfig.Spec.Service = "newservice.googleapis.com"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, false)
	copyAuditConfig = auditConfig
	newAuditLogConfig := v1beta1.AuditLogConfig{LogType: "DATA_WRITE"}
	copyAuditConfig.Spec.AuditLogConfigs = append(copyAuditConfig.Spec.AuditLogConfigs, newAuditLogConfig)
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, true)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.AuditLogConfigs[0].LogType = "ADMIN_READ"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, true)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.AuditLogConfigs = nil
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, true)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.ResourceReference.Kind = "new-resource-reference-kind"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, false)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.ResourceReference.Namespace = "new-resource-reference-namespace"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, false)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.ResourceReference.Name = "new-resource-reference-name"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, false)
	copyAuditConfig = auditConfig
	copyAuditConfig.Spec.ResourceReference.APIVersion = "new-resource-reference-apiversion"
	newAuditConfigUnstructured = newUnstructuredFromObject(t, &copyAuditConfig)
	assertHandleIAMAuditConfig(t, oldAuditConfigUnstructured, newAuditConfigUnstructured, false)
}

func assertHandleIAMAuditConfig(t *testing.T, old *unstructured.Unstructured, new *unstructured.Unstructured, expectedAllowedValue bool) {
	t.Helper()
	oldSpec := getSpecFromUnstructed(t, old)
	newSpec := getSpecFromUnstructed(t, new)
	response := handleIAMAuditConfig(oldSpec, newSpec)
	if response.Allowed != expectedAllowedValue {
		t.Fatalf("unexpected value for Allowed: got '%v', want '%v'", response.Allowed, expectedAllowedValue)
	}
}

func getSpecFromUnstructed(t *testing.T, u *unstructured.Unstructured) map[string]interface{} {
	spec, ok, err := unstructured.NestedMap(u.Object, "spec")
	if err != nil {
		t.Fatalf("unexpected error retrieving spec from '%v': %v", u.Object, err)
	}
	if !ok {
		t.Fatalf("unexpected false value for 'ok' when retrieving spec from '%v'", u.Object)
	}
	return spec
}

func TestUpdateLogLoggingMetric(t *testing.T) {
	tests := []struct {
		name     string
		spec     map[string]interface{}
		oldSpec  map[string]interface{}
		response admission.Response
	}{
		{
			name: "change on a mutable field",
			spec: map[string]interface{}{
				"description": "An updated sample log metric",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample log metric",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on a mutable field and an immutable field",
			spec: map[string]interface{}{
				"description": "An updated sample log metric",
				"metricDescriptor": map[string]interface{}{
					"metricKind": "DELTA",
					"valueType":  "DISTRIBUTION",
				},
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample log metric",
				"metricDescriptor": map[string]interface{}{
					"metricKind": "CUMULATIVE",
					"valueType":  "DISTRIBUTION",
				},
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"metricDescriptor.metricKind"})),
		},
		{
			name: "changes on multiple immutable fields",
			spec: map[string]interface{}{
				"description": "An updated sample log metric",
				"metricDescriptor": map[string]interface{}{
					"metricKind": "DELTA",
					"valueType":  "INT64",
				},
				"projectRef": map[string]interface{}{
					"external": "projects/test-project-update",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample log metric",
				"metricDescriptor": map[string]interface{}{
					"metricKind": "CUMULATIVE",
					"valueType":  "DISTRIBUTION",
				},
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"metricDescriptor.metricKind", "metricDescriptor.valueType", "projectRef"})),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := validateImmutableFieldsForLoggingLogMetricResource(tc.oldSpec, tc.spec)
			if !testutil.Equals(t, actual, tc.response) {
				t.Fatalf("got: %v, but want: %v", actual, tc.response)
			}
		})
	}
}

func TestUpdateGKEHubFeatureMembership(t *testing.T) {
	tests := []struct {
		name     string
		spec     map[string]interface{}
		oldSpec  map[string]interface{}
		response admission.Response
	}{
		{
			name: "no change on an immutable field",
			spec: map[string]interface{}{
				"description": "An updated sample",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			response: allowedResponse,
		},
		{
			name: "changes on a immutable field",
			spec: map[string]interface{}{
				"description": "An updated sample",
				"featureRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/features/test-feature-updated",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample",
				"featureRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/features/test-feature",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"featureRef"})),
		},
		{
			name: "changes on multiple immutable fields",
			spec: map[string]interface{}{
				"description": "An updated sample",
				"featureRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/features/test-feature-updated",
				},
				"membershipRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/memberships/test-membership-updated",
				},
				"location":           "test-location-updated",
				"membershipLocation": "test-membership-location-updated",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project-updated",
				},
			},
			oldSpec: map[string]interface{}{
				"description": "A sample",
				"featureRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/features/test-feature",
				},
				"membershipRef": map[string]interface{}{
					"external": "projects/test-project/locations/test-location/memberships/test-membership",
				},
				"location":           "test-location",
				"membershipLocation": "test-membership-location",
				"projectRef": map[string]interface{}{
					"external": "projects/test-project",
				},
			},
			response: admission.Errored(http.StatusForbidden,
				k8s.NewImmutableFieldsMutationError([]string{"featureRef", "location", "projectRef", "membershipLocation", "membershipRef"})),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := validateImmutableFieldsForGKEHubFeatureMembershipResource(tc.oldSpec, tc.spec)
			if !testutil.Equals(t, actual, tc.response) {
				t.Fatalf("got: %v, but want: %v", actual, tc.response)
			}
		})
	}
}

func TestDirectResourcesAlwaysAllowed(t *testing.T) {
	directGVKs := supportedgvks.DirectResources()
	for gvk, _ := range directGVKs {
		v := immutableFieldsValidatorHandler{}
		unstruct := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"kind":       gvk.Kind,
				"apiVersion": fmt.Sprintf("%s", gvk.GroupVersion()),
			},
		}

		yamlData, err := yaml.Marshal(unstruct)
		if err != nil {
			t.Fatalf("Error marshaling YAML: %v", err)
		}
		req := admission.Request{
			AdmissionRequest: admissionv1.AdmissionRequest{
				Object: runtime.RawExtension{
					Raw: yamlData,
				},
				OldObject: runtime.RawExtension{
					Raw: yamlData,
				},
			},
		}

		response := v.Handle(nil, req)
		if response.Allowed != true {
			t.Fatalf("unexpected value for Allowed: got '%v', want 'true'", response.Allowed)
		}
	}
}
