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

package conversion_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/constants"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	stvFoo = dclunstruct.ServiceTypeVersion{
		Service: "test1",
		Type:    "Foo",
		Version: "beta",
	}
	stvBar = dclunstruct.ServiceTypeVersion{
		Service: "test1",
		Type:    "Bar",
		Version: "beta",
	}
	stvBaz = dclunstruct.ServiceTypeVersion{
		Service: "test2",
		Type:    "Baz",
		Version: "beta",
	}
	stvNoname = dclunstruct.ServiceTypeVersion{
		Service: "test2",
		Type:    "NoName",
		Version: "beta",
	}
	stvNolabelsextension = dclunstruct.ServiceTypeVersion{
		Service: "test1",
		Type:    "NoLabelsExtension",
		Version: "beta",
	}
)

func TestConverter_KRMObjectToDCLObject(t *testing.T) {
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	smLoader := testservicemetadataloader.NewForUnitTest()
	converter := conversion.New(schemaLoader, smLoader)
	kind := "Test1Foo"
	apiVersion := "test1.cnrm.cloud.google.com/v1alpha1"
	projectID := "dcl-test"
	tests := []struct {
		name     string
		krmObj   *unstructured.Unstructured
		dclObj   *dclunstruct.Resource
		hasError bool
	}{
		{
			name: "name and project only",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
		{
			name: "default to metadata.name if user-specified name is not configured in spec",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "error out if specified resourceID is an empty string",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "",
					},
				},
			},
			hasError: true,
		},
		{
			name: "server-generated id is configured for acquisition",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test2Baz",
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "server-generated-value",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvBaz,
				Object: map[string]interface{}{
					"name":    "server-generated-value",
					"project": projectID,
				},
			},
		},
		{
			name: "server-generated id is not configured for creation",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test2Baz",
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvBaz,
				Object: map[string]interface{}{
					"project": projectID,
				},
			},
		},
		{
			name: "resource with no name field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test2NoName",
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "should-be-ignored",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvNoname,
				Object: map[string]interface{}{
					"project": projectID,
				},
			},
		},
		{
			name: "primitive fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":     int64(1),
					"numberKey":  float64(0.5),
					"booleanKey": true,
					"stringKey":  "StringVal",
					"name":       "CustomizedName",
					"project":    projectID,
				},
			},
		},
		{
			name: "numeric values canonicalization",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     float64(1),
						"numberKey":  int64(1),
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":    int64(1),
					"numberKey": float64(1),
					"name":      "CustomizedName",
					"project":   projectID,
				},
			},
		},
		{
			name: "list of primitives",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
							"myString2",
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"primitiveArrayKey": []interface{}{
						"myString1",
						"myString2",
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "nested object",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "StringVal",
							"nestedReferenceFieldRef": map[string]interface{}{
								"external": "reference-url",
							},
							"nestedSensitiveField": map[string]interface{}{
								"value": "sensitive-data",
							},
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"nestedObjectKey": map[string]interface{}{
						"nestedField1":         true,
						"nestedField2":         "StringVal",
						"nestedReferenceField": "reference-url",
						"nestedSensitiveField": "sensitive-data",
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "string-object map",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "my-ref-1",
									},
									map[string]interface{}{
										"external": "my-ref-2",
									},
								},
								"readOnlyField":          "val1",
								"readOnlySensitiveField": "sensitiveVal",
								"readOnlyReferenceField": "referenceVal",
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "my-ref-3",
									},
									map[string]interface{}{
										"external": "my-ref-4",
									},
								},
								"readOnlyField": "val2",
							},
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"stringObjectMapKey": map[string]interface{}{
						"someKey": map[string]interface{}{
							"objectField1": 1.0,
							"objectReferenceArrayKey": []interface{}{
								"my-ref-1",
								"my-ref-2",
							},
							// Read-only fields are left in for objects within
							// collections (maps, arrays)
							"readOnlyField": "val1",
							// Read-only sensitive and reference fields are
							// left in objects within collections (maps,
							// arrays) and simply treated as strings rather
							// than sensitive or reference structs
							"readOnlySensitiveField": "sensitiveVal",
							"readOnlyReferenceField": "referenceVal",
						},
						"someOtherKey": map[string]interface{}{
							"objectField1": 2.0,
							"objectReferenceArrayKey": []interface{}{
								"my-ref-3",
								"my-ref-4",
							},
							"readOnlyField": "val2",
						},
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "list of objects",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1":                 1,
								"field2":                 "item1",
								"readOnlyField":          "oldVal1",
								"readOnlySensitiveField": "sensitiveVal",
								"readOnlyReferenceField": "referenceVal",
							},
							map[string]interface{}{
								"field1":        2,
								"field2":        "item2",
								"readOnlyField": "oldVal2",
							},
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":     int64(1),
					"numberKey":  float64(0.5),
					"booleanKey": true,
					"stringKey":  "StringVal",
					"objectArrayKey": []interface{}{
						map[string]interface{}{
							"field1": float64(1),
							"field2": "item1",
							// Read-only fields are left in for objects within
							// collections (maps, arrays)
							"readOnlyField": "oldVal1",
							// Read-only sensitive and reference fields are
							// left in objects within collections (maps,
							// arrays) and simply treated as strings rather
							// than sensitive or reference structs
							"readOnlySensitiveField": "sensitiveVal",
							"readOnlyReferenceField": "referenceVal",
						},
						map[string]interface{}{
							"field1":        float64(2),
							"field2":        "item2",
							"readOnlyField": "oldVal2",
						},
					},
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
		{
			name: "sensitive field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"sensitiveFieldKey": map[string]interface{}{
							"value": "sensitive-data",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"sensitiveFieldKey": "sensitive-data",
					"name":              "CustomizedName",
					"project":           projectID,
				},
			},
		},
		{
			name: "simple reference field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"referenceKeyRef": map[string]interface{}{
							"external": "my-ref",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"referenceKey": "my-ref",
					"name":         "CustomizedName",
					"project":      projectID,
				},
			},
		},
		{
			name: "complex reference field with bar resource kind",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"complexReferenceKeyRef": map[string]interface{}{
							"external": "bar-url",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"complexReferenceKey": "bar-url",
					"name":                "CustomizedName",
					"project":             projectID,
				},
			},
		},
		{
			name: "complex reference field with baz resource kind",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"complexReferenceKeyRef": map[string]interface{}{
							"external": "baz-url",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"complexReferenceKey": "baz-url",
					"name":                "CustomizedName",
					"project":             projectID,
				},
			},
		},
		{
			name: "a list of references",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"referenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "my-ref-1",
							},
							map[string]interface{}{
								"external": "my-ref-2",
							},
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"referenceArrayKey": []interface{}{
						"my-ref-1",
						"my-ref-2",
					},
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclObj, err := converter.KRMObjectToDCLObject(tc.krmObj)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting KRM to DCL")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(dclObj, tc.dclObj) {
				t.Fatalf("unexpected converted dcl obj diff (-want +got): \n%v", cmp.Diff(tc.dclObj, dclObj))
			}
		})
	}
}

func TestConverter_KRMObjectToDCLObjectForHierarchicalReferences(t *testing.T) {
	tests := []struct {
		name   string
		schema *openapi.Schema
		krmObj *unstructured.Unstructured
		dclObj *dclunstruct.Resource
	}{
		{
			name: "single-parent resource",
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
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5ProjectRef",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "ProjectRef",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"project": "project_id",
				},
			},
		},
		{
			name: "multi-parent resource with project reference",
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "projects/project_id",
				},
			},
		},
		{
			name: "multi-parent resource with folder reference",
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "folders/folder_id",
				},
			},
		},
		{
			name: "multi-parent resource with organization reference",
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"organizationRef": map[string]interface{}{
							"external": "organization_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "organizations/organization_id",
				},
			},
		},
		{
			name: "multi-parent resource with billing account reference",
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"billingAccountRef": map[string]interface{}{
							"external": "billing_account_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "billingAccounts/billing_account_id",
				},
			},
		},
		{
			name: "multi-parent resource with reference value that already has parent prefix",
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
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name": "name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "projects/project_id",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "projects/project_id",
				},
			},
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			schemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.krmObj.GroupVersionKind(), smLoader)
			schemaMap := map[string]*openapi.Schema{
				schemaKey: tc.schema,
			}
			schemaLoader := testdclschemaloader.New(schemaMap)
			converter := conversion.New(schemaLoader, smLoader)

			dclObj, err := converter.KRMObjectToDCLObject(tc.krmObj)
			if err != nil {
				t.Fatalf("got error, but wanted none: %v", err)
			}
			if !reflect.DeepEqual(dclObj, tc.dclObj) {
				t.Fatalf("unexpected converted dcl obj diff (-want +got): \n%v", cmp.Diff(tc.dclObj, dclObj))
			}
		})
	}
}

func TestConverter_DCLObjectToKRMObject(t *testing.T) {
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	smLoader := testservicemetadataloader.NewForUnitTest()
	converter := conversion.New(schemaLoader, smLoader)
	kind := "Test1Foo"
	apiVersion := "test1.cnrm.cloud.google.com/v1alpha1"
	projectID := "dcl-test"
	tests := []struct {
		name     string
		krmObj   *unstructured.Unstructured
		dclObj   *dclunstruct.Resource
		hasError bool
	}{
		{
			name: "name and project only",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
		{
			name: "server-generated id is returned by the service",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test2Baz",
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "server-generated-value",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvBaz,
				Object: map[string]interface{}{
					"name":    "server-generated-value",
					"project": projectID,
				},
			},
		},
		{
			name: "error out if server-generated id is not returned",
			dclObj: &dclunstruct.Resource{
				STV: stvBaz,
				Object: map[string]interface{}{
					"project": projectID,
				},
			},
			hasError: true,
		},
		{
			name: "resource with no name field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test2NoName",
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvNoname,
				Object: map[string]interface{}{
					"project": projectID,
				},
			},
		},
		{
			name: "primitive fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":     1,
					"numberKey":  0.5,
					"booleanKey": true,
					"stringKey":  "StringVal",
					"name":       "CustomizedName",
					"project":    projectID,
				},
			},
		},
		{
			name: "list of primitives",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
							"myString2",
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"primitiveArrayKey": []interface{}{
						"myString1",
						"myString2",
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "nested object",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "StringVal",
							"nestedReferenceFieldRef": map[string]interface{}{
								"external": "reference-url",
							},
							"nestedSensitiveField": map[string]interface{}{
								"value": "sensitive-data",
							},
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"nestedObjectKey": map[string]interface{}{
						"nestedField1":         true,
						"nestedField2":         "StringVal",
						"nestedReferenceField": "reference-url",
						"nestedSensitiveField": "sensitive-data",
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
		{
			name: "list of objects",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 1,
								"field2": "item1",
								// Read-only fields are surfaced onto the spec
								// if they're part of objects within
								// collections (maps, arrays)
								"readOnlyField": "running",
								// Read-only sensitive and reference fields are
								// surfaced onto the spec if they're part of
								// objects within collections (maps, arrays)
								// and simply surfaced as strings rather than
								// sensitive or reference structs.
								"readOnlySensitiveField": "sensitiveVal",
								"readOnlyReferenceField": "referenceVal",
							},
							map[string]interface{}{
								"field1":        2,
								"field2":        "item2",
								"readOnlyField": "suspended",
							},
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":     1,
					"numberKey":  0.5,
					"booleanKey": true,
					"stringKey":  "StringVal",
					"objectArrayKey": []interface{}{
						map[string]interface{}{
							"field1":                 1,
							"field2":                 "item1",
							"readOnlyField":          "running",
							"readOnlySensitiveField": "sensitiveVal",
							"readOnlyReferenceField": "referenceVal",
						},
						map[string]interface{}{
							"field1":        2,
							"field2":        "item2",
							"readOnlyField": "suspended",
						},
					},
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
		{
			name: "sensitive field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"sensitiveFieldKey": map[string]interface{}{
							"value": "sensitive-data",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"sensitiveFieldKey": "sensitive-data",
					"name":              "CustomizedName",
					"project":           projectID,
				},
			},
		},
		{
			name: "single-kind reference field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"referenceKeyRef": map[string]interface{}{
							"external": "my-ref",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"referenceKey": "my-ref",
					"name":         "CustomizedName",
					"project":      projectID,
				},
			},
		},
		{
			name: "multi-kind reference field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"complexReferenceKeyRef": map[string]interface{}{
							"external": "my-ref",
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"complexReferenceKey": "my-ref",
					"name":                "CustomizedName",
					"project":             projectID,
				},
			},
		},
		{
			name: "a list of references",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"referenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "my-ref-1",
							},
							map[string]interface{}{
								"external": "my-ref-2",
							},
						},
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"referenceArrayKey": []interface{}{
						"my-ref-1",
						"my-ref-2",
					},
					"name":    "CustomizedName",
					"project": projectID,
				},
			},
		},
		{
			name: "string-object map",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "my-ref-1",
									},
									map[string]interface{}{
										"external": "my-ref-2",
									},
								},
								// Read-only fields are surfaced onto the spec
								// if they're part of objects within
								// collections (maps, arrays)
								"readOnlyField": "running",
								// Read-only sensitive and reference fields are
								// surfaced onto the spec if they're part of
								// objects within collections (maps, arrays)
								// and simply surfaced as strings rather than
								// sensitive or reference structs.
								"readOnlySensitiveField": "sensitiveVal",
								"readOnlyReferenceField": "referenceVal",
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "my-ref-3",
									},
									map[string]interface{}{
										"external": "my-ref-4",
									},
								},
								"readOnlyField": "val2",
							},
						},
						"resourceID": "foo-example",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"stringObjectMapKey": map[string]interface{}{
						"someKey": map[string]interface{}{
							"objectField1": 1.0,
							"objectReferenceArrayKey": []interface{}{
								"my-ref-1",
								"my-ref-2",
							},
							"readOnlyField":          "running",
							"readOnlySensitiveField": "sensitiveVal",
							"readOnlyReferenceField": "referenceVal",
						},
						"someOtherKey": map[string]interface{}{
							"objectField1": 2.0,
							"objectReferenceArrayKey": []interface{}{
								"my-ref-3",
								"my-ref-4",
							},
							"readOnlyField": "val2",
						},
					},
					"name":    "foo-example",
					"project": projectID,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			krmObj, err := converter.DCLObjectToKRMObject(tc.dclObj)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting DCL TO KRM")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			compareUnstructuredObjects(t, krmObj, tc.krmObj)
		})
	}
}

func TestConverter_DCLObjectToKRMObjectForHierarchicalReferences(t *testing.T) {
	tests := []struct {
		name     string
		schema   *openapi.Schema
		dclObj   *dclunstruct.Resource
		krmObj   *unstructured.Unstructured
		hasError bool
	}{
		{
			name: "single-parent resource",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "ProjectRef",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"project": "project_id",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5ProjectRef",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
					},
				},
			},
		},
		{
			name: "multi-parent resource with project parent",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "projects/project_id",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "projects/project_id",
						},
					},
				},
			},
		},
		{
			name: "multi-parent resource with folder parent",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "folders/folder_id",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folders/folder_id",
						},
					},
				},
			},
		},
		{
			name: "multi-parent resource with organization parent",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "organizations/organization_id",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"spec": map[string]interface{}{
						"organizationRef": map[string]interface{}{
							"external": "organizations/organization_id",
						},
					},
				},
			},
		},
		{
			name: "multi-parent resource with billing account parent",
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "billingAccounts/billing_account_id",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
					"spec": map[string]interface{}{
						"billingAccountRef": map[string]interface{}{
							"external": "billingAccounts/billing_account_id",
						},
					},
				},
			},
		},
		{
			name: "multi-parent resource with empty parent",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
				},
			},
			hasError: true,
		},
		{
			name: "multi-parent resource with unrecognized parent",
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
			dclObj: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "test5",
					Type:    "MultipleRefs",
					Version: "beta",
				},
				Object: map[string]interface{}{
					"parent": "val_with_no_known_parent_prefix",
				},
			},
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test5MultipleRefs",
					"apiVersion": "test5.cnrm.cloud.google.com/v1alpha1",
				},
			},
			hasError: true,
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			schemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.krmObj.GroupVersionKind(), smLoader)
			schemaMap := map[string]*openapi.Schema{
				schemaKey: tc.schema,
			}
			schemaLoader := testdclschemaloader.New(schemaMap)
			converter := conversion.New(schemaLoader, smLoader)

			krmObj, err := converter.DCLObjectToKRMObject(tc.dclObj)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got no error, but wanted one")
				}
				return
			}
			if err != nil {
				t.Fatalf("got error, but wanted none: %v", err)
			}
			compareUnstructuredObjects(t, krmObj, tc.krmObj)
		})
	}
}

func TestLabelsConversion(t *testing.T) {
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	smLoader := testservicemetadataloader.NewForUnitTest()
	converter := conversion.New(schemaLoader, smLoader)
	kind := "Test1Bar"
	apiVersion := "test1.cnrm.cloud.google.com/v1alpha1"
	projectID := "dcl-test"
	tests := []struct {
		name   string
		krmObj *unstructured.Unstructured
		dclObj *dclunstruct.Resource
	}{
		{
			name: "system label 'managed-by-cnrm' is added",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name": "foo-example",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvBar,
				Object: map[string]interface{}{
					"labels": map[string]interface{}{
						"key1":            "val1",
						"key2":            "val2",
						"managed-by-cnrm": "true",
					},
					"project": projectID,
					"name":    "CustomizedName",
				},
			},
		},
		{
			name: "k8s style labels are trimmed",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name": "foo-example",
						"labels": map[string]interface{}{
							"key1":                 "val1",
							"key2":                 "val2",
							"k8s.test.io/some-key": "some-value",
						},
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvBar,
				Object: map[string]interface{}{
					"labels": map[string]interface{}{
						"key1":            "val1",
						"key2":            "val2",
						"managed-by-cnrm": "true",
					},
					"project": projectID,
					"name":    "CustomizedName",
				},
			},
		},
		{
			name: "without labels extension metadata labels will not be converted to dcl labels field",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test1NoLabelsExtension",
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name": "foo-example",
						"labels": map[string]interface{}{
							"key1":                 "val1",
							"key2":                 "val2",
							"k8s.test.io/some-key": "some-value",
						},
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvNolabelsextension,
				Object: map[string]interface{}{
					"project": projectID,
					"name":    "CustomizedName",
				},
			},
		},
		{
			name: "without labels extension dcl labels will not be converted to metadata labels",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "Test1NoLabelsExtension",
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"name":      "foo-example",
						"labels":    map[string]interface{}{},
						"namespace": "test-system",
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"resourceID": "CustomizedName",
						"labels": map[string]interface{}{
							"key1": "val1",
							"key2": "val2",
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvNolabelsextension,
				Object: map[string]interface{}{
					"project": projectID,
					"name":    "CustomizedName",
					"labels": map[string]interface{}{
						"key1": "val1",
						"key2": "val2",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclObj, err := converter.KRMObjectToDCLObject(tc.krmObj)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(dclObj, tc.dclObj) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(tc.dclObj, dclObj))
			}
			krmObj, err := converter.DCLObjectToKRMObject(tc.dclObj)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			dclSchema, err := converter.SchemaLoader.GetDCLSchema(dclObj.STV)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			labelsField := dclSchema.Extension[constants.DCLLabelsField]
			want := interface{}(nil)
			if labelsField != nil {
				want = tc.dclObj.Object[labelsField.(string)].(map[string]interface{})
			}
			got := krmObj.Object["metadata"].(map[string]interface{})["labels"]
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected labels diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestConverter_DCLObjectToKRMObject_WithStatus(t *testing.T) {
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	smLoader := testservicemetadataloader.NewForUnitTest()
	converter := conversion.New(schemaLoader, smLoader)
	kind := "Test1Foo"
	apiVersion := "test1.cnrm.cloud.google.com/v1alpha1"
	projectID := "dcl-test"
	tests := []struct {
		name   string
		krmObj *unstructured.Unstructured
		dclObj *dclunstruct.Resource
	}{
		{
			name: "primitive fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
					"status": map[string]interface{}{
						"statusField": "message",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":      1,
					"numberKey":   0.5,
					"booleanKey":  true,
					"stringKey":   "StringVal",
					"name":        "CustomizedName",
					"project":     projectID,
					"statusField": "message",
				},
			},
		},
		{
			name: "nested status fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "StringVal",
						},
					},
					"status": map[string]interface{}{
						"statusField": "message",
						"nestedObjectKey": map[string]interface{}{
							"nestedStatusField":       "StringVal2",
							"nestedStatusArrayString": []interface{}{"status1, status2"},
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":      1,
					"numberKey":   0.5,
					"booleanKey":  true,
					"stringKey":   "StringVal",
					"name":        "CustomizedName",
					"project":     projectID,
					"statusField": "message",
					"nestedObjectKey": map[string]interface{}{
						"nestedField1":            true,
						"nestedField2":            "StringVal",
						"nestedStatusField":       "StringVal2",
						"nestedStatusArrayString": []interface{}{"status1, status2"},
					},
				},
			},
		},
		{
			name: "status array of primitives",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
					"status": map[string]interface{}{
						"statusField":          "message",
						"statusArrayPrimitive": []interface{}{"status1, status2"},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":               1,
					"numberKey":            0.5,
					"booleanKey":           true,
					"stringKey":            "StringVal",
					"name":                 "CustomizedName",
					"project":              projectID,
					"statusField":          "message",
					"statusArrayPrimitive": []interface{}{"status1, status2"},
				},
			},
		},
		{
			name: "status array of objects",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
					"status": map[string]interface{}{
						"statusField": "message",
						"statusArrayObject": []interface{}{
							map[string]interface{}{
								"nestedStatusField1": "status1Field1",
								"nestedStatusField2": "status1Field2",
							},
							map[string]interface{}{
								"nestedStatusField1": "status2Field1",
								"nestedStatusField2": "status2Field2",
							},
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":      1,
					"numberKey":   0.5,
					"booleanKey":  true,
					"stringKey":   "StringVal",
					"name":        "CustomizedName",
					"project":     projectID,
					"statusField": "message",
					"statusArrayObject": []interface{}{
						map[string]interface{}{
							"nestedStatusField1": "status1Field1",
							"nestedStatusField2": "status1Field2",
						},
						map[string]interface{}{
							"nestedStatusField1": "status2Field1",
							"nestedStatusField2": "status2Field2",
						},
					},
				},
			},
		},
		{
			name: "status string-object map",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
					"status": map[string]interface{}{
						"statusField": "message",
						"statusStringObjectMap": map[string]interface{}{
							"someKey": map[string]interface{}{
								"nestedStatusField1": true,
								"nestedStatusField2": "str1",
							},
							"someOtherKey": map[string]interface{}{
								"nestedStatusField1": false,
								"nestedStatusField2": "str2",
							},
						},
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":      1,
					"numberKey":   0.5,
					"booleanKey":  true,
					"stringKey":   "StringVal",
					"name":        "CustomizedName",
					"project":     projectID,
					"statusField": "message",
					"statusStringObjectMap": map[string]interface{}{
						"someKey": map[string]interface{}{
							"nestedStatusField1": true,
							"nestedStatusField2": "str1",
						},
						"someOtherKey": map[string]interface{}{
							"nestedStatusField1": false,
							"nestedStatusField2": "str2",
						},
					},
				},
			},
		},
		{
			name: "status with fields that collide with reserved names",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       kind,
					"apiVersion": apiVersion,
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/project-id": projectID,
						},
					},
					"spec": map[string]interface{}{
						"intKey":     1,
						"numberKey":  0.5,
						"booleanKey": true,
						"stringKey":  "StringVal",
						"resourceID": "CustomizedName",
					},
					"status": map[string]interface{}{
						"statusField":                "message",
						"resourceConditions":         "value",
						"resourceObservedGeneration": "value",
					},
				},
			},
			dclObj: &dclunstruct.Resource{
				STV: stvFoo,
				Object: map[string]interface{}{
					"intKey":             1,
					"numberKey":          0.5,
					"booleanKey":         true,
					"stringKey":          "StringVal",
					"name":               "CustomizedName",
					"project":            projectID,
					"statusField":        "message",
					"conditions":         "value",
					"observedGeneration": "value",
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			krmObj, err := converter.DCLObjectToKRMObject(tc.dclObj)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			compareUnstructuredObjects(t, krmObj, tc.krmObj)
		})
	}
}

func compareUnstructuredObjects(t *testing.T, actual, expected *unstructured.Unstructured) {
	//ignore name and namespace
	unstructured.RemoveNestedField(actual.Object, "metadata", "name")
	unstructured.RemoveNestedField(expected.Object, "metadata", "name")
	unstructured.RemoveNestedField(actual.Object, "metadata", "namespace")
	unstructured.RemoveNestedField(expected.Object, "metadata", "namespace")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(expected, actual))
	}
}

var dclSchemaMap = map[string]*openapi.Schema{
	"test1_beta_foo": &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": &openapi.Schema{
				Type: "string",
			},
			"name": &openapi.Schema{
				Type: "string",
			},
			"stringKey": {
				Type: "string",
			},
			"intKey": {
				Type: "integer",
			},
			"booleanKey": {
				Type: "boolean",
			},
			"numberKey": {
				Type: "number",
			},
			"statusField": {
				Type:     "string",
				ReadOnly: true,
			},
			"statusArrayPrimitive": {
				Type:     "array",
				ReadOnly: true,
				Items: &openapi.Schema{
					Type:     "string",
					ReadOnly: true,
				},
			},
			"statusArrayObject": {
				Type:     "array",
				ReadOnly: true,
				Items: &openapi.Schema{
					Type:     "object",
					ReadOnly: true,
					Properties: map[string]*openapi.Schema{
						"nestedStatusField1": &openapi.Schema{
							Type: "boolean",
						},
						"nestedStatusField2": &openapi.Schema{
							Type: "string",
						},
					},
				},
			},
			"statusStringObjectMap": {
				Type:     "object",
				ReadOnly: true,
				AdditionalProperties: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"nestedStatusField1": &openapi.Schema{
							Type: "boolean",
						},
						"nestedStatusField2": &openapi.Schema{
							Type: "string",
						},
					},
				},
			},
			"nestedObjectKey": {
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"nestedField1": {
						Type: "boolean",
					},
					"nestedField2": {
						Type: "string",
					},
					"nestedReferenceField": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "SomeService/SomeKind",
									"field":    "selfLink",
								},
							},
						},
					},
					"nestedSensitiveField": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-sensitive": true,
						},
					},
					"nestedStatusField": {
						Type:     "string",
						ReadOnly: true,
					},
					"nestedStatusArrayString": {
						Type:     "array",
						ReadOnly: true,
						Items: &openapi.Schema{
							Type:     "string",
							ReadOnly: true,
						},
					},
				},
				Required: []string{"nestedField1"},
			},
			"sensitiveFieldKey": {
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-sensitive": true,
				},
			},
			"referenceKey": {
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "FakeService/FakeKind",
							"field":    "name",
						},
					},
				},
			},
			"complexReferenceKey": {
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "FakeService/Bar",
							"field":    "self_link",
						},
						map[interface{}]interface{}{
							"resource": "FakeService/Baz",
							"field":    "self_link",
						},
					},
				},
			},
			"primitiveArrayKey": {
				Type: "array",
				Items: &openapi.Schema{
					Type: "string",
				},
			},
			"objectArrayKey": {
				Type: "array",
				Items: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"field1": {
							Type: "number",
						},
						"field2": {
							Type: "string",
						},
						"readOnlyField": {
							Type:     "string",
							ReadOnly: true,
						},
						"readOnlySensitiveField": {
							Type:     "string",
							ReadOnly: true,
							Extension: map[string]interface{}{
								"x-dcl-sensitive": true,
							},
						},
						"readOnlyReferenceField": {
							Type:     "string",
							ReadOnly: true,
							Extension: map[string]interface{}{
								"x-dcl-references": []interface{}{
									map[interface{}]interface{}{
										"resource": "FakeService/FakeKind",
										"field":    "name",
									},
								},
							},
						},
					},
				},
			},
			"referenceArrayKey": {
				Type: "array",
				Items: &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"resource": "FakeService/FakeKind",
								"field":    "name",
							},
						},
					},
				},
			},
			"stringObjectMapKey": {
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"objectField1": {
							Type: "number",
						},
						"objectReferenceArrayKey": {
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-references": []interface{}{
										map[interface{}]interface{}{
											"resource": "FakeService/FakeKind",
											"field":    "name",
										},
									},
								},
							},
						},
						"readOnlyField": {
							ReadOnly: true,
							Type:     "string",
						},
						"readOnlySensitiveField": {
							Type:     "string",
							ReadOnly: true,
							Extension: map[string]interface{}{
								"x-dcl-sensitive": true,
							},
						},
						"readOnlyReferenceField": {
							Type:     "string",
							ReadOnly: true,
							Extension: map[string]interface{}{
								"x-dcl-references": []interface{}{
									map[interface{}]interface{}{
										"resource": "FakeService/FakeKind",
										"field":    "name",
									},
								},
							},
						},
					},
				},
			},
			"conditions": {
				Type:     "string",
				ReadOnly: true,
			},
			"observedGeneration": {
				Type:     "string",
				ReadOnly: true,
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	},
	"test1_beta_bar": &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": &openapi.Schema{
				Type: "string",
			},
			"name": &openapi.Schema{
				Type: "string",
			},
			"labels": &openapi.Schema{
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "string",
				},
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
			"x-dcl-labels":           "labels",
		},
	},
	"test2_beta_baz": &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": &openapi.Schema{
				Type: "string",
			},
			"name": &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-server-generated-parameter": true,
				},
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	},
	"test2_beta_noname": &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": &openapi.Schema{
				Type: "string",
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	},
	"test1_beta_nolabelsextension": &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": &openapi.Schema{
				Type: "string",
			},
			"name": &openapi.Schema{
				Type: "string",
			},
			"labels": &openapi.Schema{
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "string",
				},
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	},
}
