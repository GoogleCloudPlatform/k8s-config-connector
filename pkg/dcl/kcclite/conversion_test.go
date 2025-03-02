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

package kcclite_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	mgr          manager.Manager
	emptyObject  = make(map[string]interface{})
	dclSchemaMap = map[string]*openapi.Schema{
		"test1_beta_foo": testSchema(),
		"test1_beta_bar": {
			Type: "object",
			Extension: map[string]interface{}{
				"x-dcl-id": "projects/{{project}}/bars/{{name}}",
			},
		},
	}

	hasErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have an error")
		}
	}
	refNotFoundErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a ReferenceNotFoundError")
		}
		if _, ok := k8s.AsReferenceNotFoundError(err); !ok {
			t.Fatalf("got %v, want to have a ReferenceNotFoundError", err)
		}
	}
	refNotReadyErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a ReferenceNotReadyError")
		}
		if _, ok := k8s.AsReferenceNotReadyError(err); !ok {
			t.Fatalf("got %v, want to have a ReferenceNotReadyError", err)
		}
	}
	secretNotFoundErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a SecretNotFoundError")
		}
		if _, ok := k8s.AsSecretNotFoundError(err); !ok {
			t.Fatalf("got %v, want to have a SecretNotFoundError", err)
		}
	}
	keyInSecretNotFoundErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a KeyInSecretNotFoundError")
		}
		if _, ok := k8s.AsKeyInSecretNotFoundError(err); !ok {
			t.Fatalf("got %v, want to have a KeyInSecretNotFoundError", err)
		}
	}
	transDepNotFoundErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a TransitiveDependencyNotFoundError")
		}
		if _, ok := k8s.AsTransitiveDependencyNotFoundError(err); !ok {
			t.Fatalf("got %v, want to have a TransitiveDependencyNotFoundError", err)
		}
	}
	transDepNotReadyErrorCheckFunc = func(t *testing.T, err error) {
		if err == nil {
			t.Fatal("got nil, want to have a TransitiveDependencyNotReadyError")
		}
		if _, ok := k8s.AsTransitiveDependencyNotReadyError(err); !ok {
			t.Fatalf("got %v, want to have a TransitiveDependencyNotReadyError", err)
		}
	}
)

func testSchema() *openapi.Schema {
	return &openapi.Schema{
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": {
				Type: "string",
			},
			"name": {
				Type: "string",
			},
			"labels": {
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "string",
				},
			},
			"stringKey": {
				Type: "string",
			},
			"intKey": {
				Type: "integer",
			},
			"boolKey": {
				Type: "boolean",
			},
			"floatKey": {
				Type: "number",
			},
			"numberKey": {
				Type: "number",
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
					"nestedReferenceKey": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Test1/Bar",
									"field":    "statusField",
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
			"sensitiveField": {
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
							"resource": "test1/Bar",
							"field":    "name",
						},
					},
				},
			},
			"multiReferenceKey": {
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "test1/Bar",
							"field":    "name",
						},
						map[interface{}]interface{}{
							"resource": "test1/Foo",
							"field":    "name",
						},
					},
				},
			},
			"mapKey": {
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "string",
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
						"objectField2": {
							Type: "string",
						},
						"state": {
							ReadOnly: true,
							Type:     "string",
						},
						"objectReferenceArrayKey": {
							Type: "array",
							Items: &openapi.Schema{
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-references": []interface{}{
										map[interface{}]interface{}{
											"resource": "test1/Bar",
											"field":    "name",
										},
									},
								},
							},
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
						"sensitiveFieldInArray": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-sensitive": true,
							},
						},
						"bar": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-references": []interface{}{
									map[interface{}]interface{}{
										"resource": "test1/Bar",
										"field":    "name",
									},
								},
							},
						},
					},
				},
			},
			"multiReferenceArrayKey": {
				Type: "array",
				Items: &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"resource": "test1/Bar",
								"field":    "name",
							},
							map[interface{}]interface{}{
								"resource": "test1/Foo",
								"field":    "name",
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
								"resource": "test1/Bar",
								"field":    "name",
							},
						},
					},
				},
			},
			"statusField": {
				Type:     "string",
				ReadOnly: true,
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	}
}

func TestToKCCLite(t *testing.T) {
	tests := []struct {
		name                  string
		prevSpec              map[string]interface{}
		hasResourceReferences bool
		hasSecretReferences   bool
		expected              map[string]interface{}
	}{
		{
			name: "keys as constant",
			prevSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   true,
			},
			expected: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   true,
			},
		},
		{
			name: "list of primitives key",
			prevSpec: map[string]interface{}{
				"primitiveArrayKey": []interface{}{
					"myString1",
					"myString2",
				},
			},
			expected: map[string]interface{}{
				"primitiveArrayKey": []interface{}{
					"myString1",
					"myString2",
				},
			},
		},
		{
			name: "map key",
			prevSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
					"myMapKey2": "MyMapValue2",
				},
			},
			expected: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
					"myMapKey2": "MyMapValue2",
				},
			},
		},
		{
			name:                  "list of objects key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
						"barRef": map[string]interface{}{
							"name": "my-ref1",
						},
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
						"barRef": map[string]interface{}{
							"name": "my-ref2",
						},
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval3",
						"barRef": map[string]interface{}{
							"external": "my-ref3",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
						"barRef": map[string]interface{}{
							"external": "projects/my-project-1/bars/my-ref1",
						},
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
						"barRef": map[string]interface{}{
							"external": "projects/my-project-1/bars/my-ref2",
						},
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval3",
						"barRef": map[string]interface{}{
							"external": "my-ref3",
						},
					},
				},
			},
		},
		{
			name:                  "simple reference key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			expected: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"external": "projects/my-project-1/bars/my-ref1",
				},
			},
		},
		{
			name:                  "multi references key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"multiReferenceKeyRef": map[string]interface{}{
					"name": "my-ref1",
					"kind": "Test1Bar",
				},
			},
			expected: map[string]interface{}{
				"multiReferenceKeyRef": map[string]interface{}{
					"external": "projects/my-project-1/bars/my-ref1",
				},
			},
		},
		{
			name:                  "sensitive field with plain-text value",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "plain-text",
				},
			},
			expected: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "plain-text",
				},
			},
		},
		{
			name:                  "sensitive field with value from secret ref",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "secret-val1",
				},
			},
			hasSecretReferences: true,
		},
		{
			name:                  "nested sensitive field with value from secret ref",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveField": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "secret-key1",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveField": map[string]interface{}{
						"value": "secret-val1",
					},
				},
			},
			hasSecretReferences: true,
		},
		{
			name:                  "sensitive field nested in list of objects with values from secret refs",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret2",
									"key":  "secret-key2",
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"value": "secret-val1",
						},
					},
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"value": "secret-val2",
						},
					},
				},
			},
			hasSecretReferences: true,
		},
		{
			name:                  "nested objects key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "strval",
					"nestedReferenceKeyRef": map[string]interface{}{
						"name": "my-ref1",
					},
				},
			},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "strval",
					"nestedReferenceKeyRef": map[string]interface{}{
						"external": "foobar",
					},
				},
			},
		},
		{
			name:                  "list of references",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"name": "my-ref1",
					},
					map[string]interface{}{
						"name": "my-ref2",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
			expected: map[string]interface{}{
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"external": "projects/my-project-1/bars/my-ref1",
					},
					map[string]interface{}{
						"external": "projects/my-project-1/bars/my-ref2",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
		},
		{
			name:                  "list of multi-type references",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"multiReferenceArrayKey": []interface{}{
					map[string]interface{}{
						"name": "my-ref1",
						"kind": "Test1Bar",
					},
					map[string]interface{}{
						"name": "my-ref2",
						"kind": "Test1Bar",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
			expected: map[string]interface{}{
				"multiReferenceArrayKey": []interface{}{
					map[string]interface{}{
						"external": "projects/my-project-1/bars/my-ref1",
					},
					map[string]interface{}{
						"external": "projects/my-project-1/bars/my-ref2",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
		},
		{
			name:                  "string-object map",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"stringObjectMapKey": map[string]interface{}{
					"someKey": map[string]interface{}{
						"objectField1": 1.0,
						"state":        "state1",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
					"someOtherKey": map[string]interface{}{
						"objectField1": 2.0,
						"state":        "state2",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"name": "my-ref2",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"stringObjectMapKey": map[string]interface{}{
					"someKey": map[string]interface{}{
						"objectField1": 1.0,
						"state":        "state1",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "projects/my-project-1/bars/my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
					"someOtherKey": map[string]interface{}{
						"objectField1": 2.0,
						"state":        "state2",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "projects/my-project-1/bars/my-ref1",
							},
							map[string]interface{}{
								"external": "projects/my-project-1/bars/my-ref2",
							},
						},
					},
				},
			},
		},
	}
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Test1Foo"}
	loader := testservicemetadataloader.NewForUnitTest()
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	serviceMappingLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			if err := testcontroller.EnsureNamespaceExists(c, testID); err != nil {
				t.Fatal(err)
			}
			if tc.hasResourceReferences {
				references := []*unstructured.Unstructured{
					newBarUnstructuredWithResourceID(t, "my-ref1", testID, corev1.ConditionTrue),
					newBarUnstructuredWithResourceID(t, "my-ref2", testID, corev1.ConditionTrue),
				}
				test.EnsureObjectsExist(t, references, c)
			}
			if tc.hasSecretReferences {
				secretsData := []map[string]interface{}{
					{
						"secret-key1": "secret-val1",
					},
					{
						"secret-key2": "secret-val2",
					},
				}
				secrets := []*unstructured.Unstructured{
					test.NewSecretUnstructured("secret1", testID, secretsData[0]),
					test.NewSecretUnstructured("secret2", testID, secretsData[1]),
				}
				test.EnsureObjectsExist(t, secrets, c)
			}
			r := &dcl.Resource{
				Schema: testSchema(),
			}
			r.SetNamespace(testID)
			r.Spec = tc.prevSpec
			r.SetGroupVersionKind(gvk)
			actual, err := kcclite.ToKCCLite(r, loader, schemaLoader, serviceMappingLoader, c)
			if err != nil {
				t.Fatalf("error converting to KCC lite: %v", err)
			}
			actualSpec, _, err := unstructured.NestedFieldNoCopy(actual.Object, "spec")
			if err != nil {
				t.Fatalf("error getting the converted spec: %v", err)
			}
			if got, want := actualSpec, tc.expected; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestToKCCLiteBestEffort(t *testing.T) {
	tests := []struct {
		name           string
		prevSpec       map[string]interface{}
		expected       map[string]interface{}
		references     []*unstructured.Unstructured
		errorCheckFunc func(t *testing.T, err error)
	}{
		{
			name: "referenced resource is not found",
			prevSpec: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
				"referenceKeyRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			expected: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
			},
			errorCheckFunc: refNotFoundErrorCheckFunc,
		},
		{
			name: "referenced resource is not ready",
			prevSpec: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
				"referenceKeyRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			expected: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
			},
			references: []*unstructured.Unstructured{
				newBarUnstructuredWithResourceID(t, "my-ref1", "", corev1.ConditionFalse),
			},
			errorCheckFunc: refNotReadyErrorCheckFunc,
		},
		{
			name: "referenced resources are not all ready",
			prevSpec: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"name": "my-ref1",
					},
					map[string]interface{}{
						"name": "my-ref2",
					},
				},
			},
			expected: map[string]interface{}{
				"intKey":   "1",
				"floatKey": "0.5",
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"external": "projects/my-project-1/bars/my-ref1",
					},
				},
			},
			references: []*unstructured.Unstructured{
				newBarUnstructuredWithResourceID(t, "my-ref1", "", corev1.ConditionTrue),
				newBarUnstructuredWithResourceID(t, "my-ref2", "", corev1.ConditionFalse),
			},
			errorCheckFunc: refNotReadyErrorCheckFunc,
		},
		{
			name: "referenced resource in nested object is not found",
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "strval",
					"nestedReferenceKeyRef": map[string]interface{}{
						"name": "my-ref1",
					},
				},
			},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "strval",
				},
			},
			errorCheckFunc: refNotFoundErrorCheckFunc,
		},
		{
			name: "the secret is not found",
			prevSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
			expected:       map[string]interface{}{},
			errorCheckFunc: secretNotFoundErrorCheckFunc,
		},
		{
			name: "the key is not found in the secret",
			prevSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
			expected:       map[string]interface{}{},
			errorCheckFunc: keyInSecretNotFoundErrorCheckFunc,
			references: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{}),
			},
		},
	}
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Foo"}
	loader := testservicemetadataloader.NewForUnitTest()
	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	serviceMappingLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			if err := testcontroller.EnsureNamespaceExists(c, testID); err != nil {
				t.Fatal(err)
			}

			if len(tc.references) != 0 {
				for _, ref := range tc.references {
					ref.SetNamespace(testID)
				}
				test.EnsureObjectsExist(t, tc.references, c)
			}
			r := &dcl.Resource{
				Schema: testSchema(),
			}
			r.SetNamespace(testID)
			r.Spec = tc.prevSpec
			r.SetGroupVersionKind(gvk)
			_, err := kcclite.ToKCCLite(r, loader, schemaLoader, serviceMappingLoader, c)
			if tc.errorCheckFunc != nil {
				tc.errorCheckFunc(t, err)
			}
			actual, err := kcclite.ToKCCLiteBestEffort(r, loader, schemaLoader, serviceMappingLoader, c)
			if err != nil {
				t.Fatalf("unexpected error converting to KCC lite: %v", err)
			}
			actualSpec, _, err := unstructured.NestedFieldNoCopy(actual.Object, "spec")
			if err != nil {
				t.Fatalf("error getting the converted spec: %v", err)
			}
			if got, want := actualSpec, tc.expected; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestToKCCLiteForHierarchicalReferences(t *testing.T) {
	tests := []struct {
		name           string
		spec           map[string]interface{}
		expected       map[string]interface{}
		reference      *unstructured.Unstructured
		schema         *openapi.Schema
		gvk            schema.GroupVersionKind
		errorCheckFunc func(t *testing.T, err error)
	}{
		{
			name: "single-parent resource with external project reference; should simply pass through the user-provided value for 'external'",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
		},
		{
			name: "single-parent resource with project reference; should resolve referenced Project and format the resolved value using Project's x-dcl-id",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "projects/project_id",
				},
			},
			reference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionTrue),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
		},
		{
			name: "single-parent resource with project reference to non-ready Project",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			reference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionFalse),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
			errorCheckFunc: refNotReadyErrorCheckFunc,
		},
		{
			name: "single-parent resource with project reference to nonexistent Project",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
			errorCheckFunc: refNotFoundErrorCheckFunc,
		},
		{
			name: "single-parent resource with folder reference; should resolve referenced Folder and format the resolved value using Folder's x-dcl-id",
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-name",
				},
			},
			expected: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id",
				},
			},
			reference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionTrue),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"folder": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
							},
						},
					},
				},
			},
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5FolderRef",
			},
		},
		{
			name: "multi-parent resource with external project reference; should simply pass through the user-provided value for 'external'",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name: "multi-parent resource with project reference; should resolve referenced Project, but shouldn't format resolved value with Project's x-dcl-id",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			expected: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
			reference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionTrue),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name: "multi-parent resource with project reference to non-ready Project",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			reference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionFalse),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
			errorCheckFunc: refNotReadyErrorCheckFunc,
		},
		{
			name: "multi-parent resource with project reference to nonexistent Project",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
			errorCheckFunc: refNotFoundErrorCheckFunc,
		},
		{
			name: "multi-parent resource with external folder reference; should simply pass through the user-provided value for 'external'",
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id",
				},
			},
			expected: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name: "multi-parent resource with folder reference",
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-name",
				},
			},
			expected: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id",
				},
			},
			reference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionTrue),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name: "multi-parent resource with folder reference to non-ready Folder",
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-name",
				},
			},
			reference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionFalse),
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
			errorCheckFunc: refNotReadyErrorCheckFunc,
		},
		{
			name: "multi-parent resource with folder reference to nonexistent Folder",
			spec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
			errorCheckFunc: refNotFoundErrorCheckFunc,
		},
		{
			name: "multi-parent resource with external organization reference; should simply pass through the user-provided value for 'external'",
			spec: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"external": "organization_id",
				},
			},
			expected: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"external": "organization_id",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name: "single-parent resource with no hierarchical reference in spec; should not error out",
			spec: map[string]interface{}{
				"field": "val",
			},
			expected: map[string]interface{}{
				"field": "val",
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
		},
		{
			name: "multi-parent resource with no hierarchical reference in spec; should not error out",
			spec: map[string]interface{}{
				"field": "val",
			},
			expected: map[string]interface{}{
				"field": "val",
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
		{
			name:     "single-parent resource with nil spec; should not error out",
			spec:     nil,
			expected: nil,
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
					"project": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
		},
		{
			name:     "multi-parent resource with nil spec; should not error out",
			spec:     nil,
			expected: nil,
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
					"parent": {
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
			gvk: schema.GroupVersionKind{
				Group:   "Test5",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	serviceMappingLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			if err := testcontroller.EnsureNamespaceExists(c, testID); err != nil {
				t.Fatal(err)
			}
			if tc.reference != nil {
				tc.reference.SetNamespace(testID)
				test.EnsureObjectExists(t, tc.reference, c)
			}
			r := &dcl.Resource{
				Schema: tc.schema,
			}
			r.SetNamespace(testID)
			r.Spec = tc.spec
			r.SetGroupVersionKind(tc.gvk)

			schemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.gvk, smLoader)
			schemaMap := map[string]*openapi.Schema{
				schemaKey: tc.schema,

				// Add the following to the list of fake DCL schemas to allow for our
				// test to test resources that reference hierarchical resources
				// (e.g. "Cloudresourcemanager/Project").
				// Note that the "x-dcl-id" values used below for testing are
				// the same ones from the real DCL schemas.
				"cloudresourcemanager_ga_project": {
					Extension: map[string]interface{}{
						"x-dcl-id": "projects/{{name}}",
					},
				},
				"cloudresourcemanager_ga_folder": {
					Extension: map[string]interface{}{
						"x-dcl-id": "{{name}}",
					},
				},
			}
			schemaLoader := testdclschemaloader.New(schemaMap)

			actual, err := kcclite.ToKCCLite(r, smLoader, schemaLoader, serviceMappingLoader, c)
			if tc.errorCheckFunc != nil {
				tc.errorCheckFunc(t, err)
				return
			}
			if err != nil {
				t.Fatalf("error converting to KCC lite: %v", err)
			}
			actualSpec, _, err := unstructured.NestedFieldNoCopy(actual.Object, "spec")
			if err != nil {
				t.Fatalf("error getting the converted spec: %v", err)
			}
			if got, want := actualSpec, tc.expected; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestResolveSpecAndStatusWithMixedSpecAndLegacyStatus(t *testing.T) {
	tests := []struct {
		name           string
		state          *unstructured.Unstructured
		dclResource    *dcl.Resource
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
	}{
		{
			name: "primitives are set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    "1",
						"floatKey":  "0.5",
						"stringKey": "StringVal",
						"boolKey":   false,
					},
				},
			},
			dclResource: &dcl.Resource{
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   false,
			},
		},
		{
			name: "status fields are set from state",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"statusField": "statusVal1",
						"nestedObjectKey": map[string]interface{}{
							"nestedStatusField": "statusVal2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Status: map[string]interface{}{
						"statusField": "statusVal2",
					},
				},
				Schema: testSchema(),
			},
			expectedStatus: map[string]interface{}{
				"statusField": "statusVal1",
				"nestedObjectKey": map[string]interface{}{
					"nestedStatusField": "statusVal2",
				},
			},
		},
		{
			name: "both spec and status fields are present",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    "1",
						"floatKey":  "0.5",
						"stringKey": "StringVal",
						"boolKey":   false,
					},
					"status": map[string]interface{}{
						"statusField": "statusVal1",
						"nestedObjectKey": map[string]interface{}{
							"nestedStatusField": "statusVal2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Status: map[string]interface{}{
						"statusField": "statusVal2",
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   false,
			},
			expectedStatus: map[string]interface{}{
				"statusField": "statusVal1",
				"nestedObjectKey": map[string]interface{}{
					"nestedStatusField": "statusVal2",
				},
			},
		},
		{
			name: "lists of objects are set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval3",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval3",
					},
				},
			},
		},
		{
			name: "lists of objects are merged with defaulted fields",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
							},
							map[string]interface{}{
								"field1": 0.9,
								"field2": "strval3",
							},
							map[string]interface{}{
								"field1": 1.2,
								"field2": "strval4",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
							},
							map[string]interface{}{
								"field1": 0.7,
							},
							map[string]interface{}{
								"field1": 0.9,
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
					},
					map[string]interface{}{
						"field1": 0.9,
						"field2": "strval3",
					},
					map[string]interface{}{
						"field1": 1.2,
						"field2": "strval4",
					},
				},
			},
		},
		{
			name: "nested objects are set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": false,
					"nestedField2": "strval2",
				},
			},
		},
		{
			name: "nested objects are merged with defaulted values",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": false,
					"nestedField2": "strval2",
				},
			},
		},
		{
			name: "individual resource references are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"referenceKeyRef": map[string]interface{}{
							"external": "my-ref1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"referenceKeyRef": map[string]interface{}{
							"name": "my-ref1",
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
		},
		{
			name: "string-object maps are merged",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"state":        "state1",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.0,
								"state":        "state2",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref2",
									},
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 3.0,
								"state":        "state2",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"name": "my-ref2",
									},
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"stringObjectMapKey": map[string]interface{}{
					"someKey": map[string]interface{}{
						"objectField1": 1.0,
						"state":        "state1",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
					"someOtherKey": map[string]interface{}{
						"objectField1": 3.0,
						"state":        "state2",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"name": "my-ref2",
							},
						},
					},
				},
			},
		},
		{
			name: "lists of resource references are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"referenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref2",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"referenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"name": "my-ref2",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"name": "my-ref1",
					},
					map[string]interface{}{
						"name": "my-ref2",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
		},
		{
			name: "resource references nested in lists of objects are preserved, default values are added",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
								"barRef": map[string]interface{}{
									"external": "my-ref1",
								},
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
								"barRef": map[string]interface{}{
									"external": "my-ref2",
								},
							},
							map[string]interface{}{
								"field1": 0.9,
								"field2": "strval3",
								"barRef": map[string]interface{}{
									"external": "my-ref3",
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field2": "strval1",
								"barRef": map[string]interface{}{
									"name": "my-ref1",
								},
							},
							map[string]interface{}{
								"field2": "strval2",
								"barRef": map[string]interface{}{
									"name": "my-ref2",
								},
							},
							map[string]interface{}{
								"field2": "strval3",
								"barRef": map[string]interface{}{
									"external": "my-ref3",
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
						"barRef": map[string]interface{}{
							"name": "my-ref1",
						},
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
						"barRef": map[string]interface{}{
							"name": "my-ref2",
						},
					},
					map[string]interface{}{
						"field1": 0.9,
						"field2": "strval3",
						"barRef": map[string]interface{}{
							"external": "my-ref3",
						},
					},
				},
			},
		},
		{
			name: "external resource reference set if no reference defined by spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"referenceKeyRef": map[string]interface{}{
							"external": "my-ref1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"referenceKeyRef": map[string]interface{}{
					"external": "my-ref1",
				},
			},
		},
		{
			name: "list of external resource references set if no list defined by spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"referenceArrayKey": []interface{}{
							map[string]interface{}{
								"external": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref2",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"referenceArrayKey": []interface{}{
					map[string]interface{}{
						"external": "my-ref1",
					},
					map[string]interface{}{
						"external": "my-ref2",
					},
					map[string]interface{}{
						"external": "my-ref3",
					},
				},
			},
		},
		{
			name: "hierarchical references for single-parent resources are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-name",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"project": {
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
			},
			expectedSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
		},
		{
			name: "external hierarchical references for single-parent resources are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id_from_state",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id_from_spec",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"project": {
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
			},
			expectedSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id_from_spec",
				},
			},
		},
		{
			name: "hierarchical references for multi-parent resources are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "folder-name",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"parent": {
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
			},
			expectedSpec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder-name",
				},
			},
		},
		{
			name: "external hierarchical references for multi-parent resources are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id_from_state",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id_from_spec",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"parent": {
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
			},
			expectedSpec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id_from_spec",
				},
			},
		},
		{
			name: "hierarchical references for multi-parent resources are preserved even if spec and state contain different types of references",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-name",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"parent": {
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
			},
			expectedSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
			},
		},
		{
			name: "external hierarchical references for multi-parent resources are preserved even if spec and state contain different types of references",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"organizationRef": map[string]interface{}{
							"external": "organization_id",
						},
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"parent": {
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
			},
			expectedSpec: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"external": "organization_id",
				},
			},
		},
		{
			name: "hierarchical references for single-parent resources are taken from state if none found in spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"project": {
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
			},
			expectedSpec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project_id",
				},
			},
		},
		{
			name: "hierarchical references for multi-parent resources are taken from state if none found in spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder_id",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"parent": {
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
			},
			expectedSpec: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder_id",
				},
			},
		},
		{
			name: "sensitive fields with plain-text values are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "secret-val1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "secret-val1",
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "secret-val1",
				},
			},
		},
		{
			name: "sensitive fields with values from secret refs are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "secret-val1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
		},
		{
			name: "sensitive fields nested in lists of objects are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"sensitiveFieldInArray": map[string]interface{}{
									"value": "secret-val1",
								},
							},
							map[string]interface{}{
								"sensitiveFieldInArray": map[string]interface{}{
									"value": "secret-val2",
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"sensitiveFieldInArray": map[string]interface{}{
									"valueFrom": map[string]interface{}{
										"secretKeyRef": map[string]interface{}{
											"name": "secret1",
											"key":  "secret-key1",
										},
									},
								},
							},
							map[string]interface{}{
								"sensitiveFieldInArray": map[string]interface{}{
									"valueFrom": map[string]interface{}{
										"secretKeyRef": map[string]interface{}{
											"name": "secret2",
											"key":  "secret-key2",
										},
									},
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldInArray": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret2",
									"key":  "secret-key2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "sensitive fields nested in objects are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedSensitiveField": map[string]interface{}{
								"value": "secret-val1",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedSensitiveField": map[string]interface{}{
								"valueFrom": map[string]interface{}{
									"secretKeyRef": map[string]interface{}{
										"name": "secret1",
										"key":  "secret-key1",
									},
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveField": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "secret-key1",
							},
						},
					},
				},
			},
		},
		{
			name: "sensitive fields set with plain-text value if not specified",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "secret-val1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{},
				Schema:   testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "secret-val1",
				},
			},
		},
		{
			name: "sensitive fields nested in lists of objects set with plain-text value if not specified",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"sensitiveFieldInArray": map[string]interface{}{
									"value": "secret-val1",
								},
							},
							map[string]interface{}{
								"field1": 0.9,
								"sensitiveFieldInArray": map[string]interface{}{
									"value": "secret-val2",
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
							},
							map[string]interface{}{
								"field1": 0.9,
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"sensitiveFieldInArray": map[string]interface{}{
							"value": "secret-val1",
						},
					},
					map[string]interface{}{
						"field1": 0.9,
						"sensitiveFieldInArray": map[string]interface{}{
							"value": "secret-val2",
						},
					},
				},
			},
		},
		{
			name: "spec-defined values are preserved",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    "2",
						"floatKey":  "1",
						"stringKey": "StringVal2",
						"boolKey":   true,
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":    "1",
						"floatKey":  "0.5",
						"stringKey": "StringVal",
						"boolKey":   false,
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   false,
			},
		},
		{
			name: "server-generated id is retrieved from state",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":     "2",
						"resourceID": "server-generated-value",
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey": "1",
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":     "1",
				"resourceID": "server-generated-value",
			},
		},
		{
			name: "maps are treated as atomic when specified by user",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"mapKey": map[string]interface{}{
							"myMapKey1": "MyMapValue1",
							"myMapKey2": "MyMapValue2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"mapKey": map[string]interface{}{
							"myMapKey1": "MyMapValue1",
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
				},
			},
		},
		// Tests surrounding managed fields
		{
			name: "values are sourced from live state when not in managed fields set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    "1",
						"floatKey":  "0.5",
						"stringKey": "StringVal",
						"boolKey":   false,
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":    "2",
						"floatKey":  "1.0",
						"stringKey": "StringVal2",
						"boolKey":   true,
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "strval1",
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:unrelated": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   false,
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": false,
					"nestedField2": "strval2",
				},
			},
		},
		{
			name: "values for sensitive fields are sourced from live state when not in managed fields set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "new-val",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "old-val",
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:unrelated": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "new-val",
				},
			},
		},
		{
			name: "values are sourced from spec when in managed fields set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    "1",
						"floatKey":  "0.5",
						"stringKey": "StringVal",
						"boolKey":   false,
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":    "2",
						"floatKey":  "1.0",
						"stringKey": "StringVal2",
						"boolKey":   true,
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "strval1",
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:intKey":    emptyObject,
						"f:floatKey":  emptyObject,
						"f:stringKey": emptyObject,
						"f:boolKey":   emptyObject,
						"f:nestedObjectKey": map[string]interface{}{
							"f:nestedField1": emptyObject,
							"f:nestedField2": emptyObject,
						},
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":    "2",
				"floatKey":  "1.0",
				"stringKey": "StringVal2",
				"boolKey":   true,
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "strval1",
				},
			},
		},
		{
			name: "maps are treated as atomic when k8s-managed",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"mapKey": map[string]interface{}{
							"myMapKey1": "MyMapValue1",
							"myMapKey2": "MyMapValue2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"mapKey": map[string]interface{}{
							"myMapKey1": "MyMapValue1",
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:mapKey": map[string]interface{}{
							"f:myMapKey1": emptyObject,
						},
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
				},
			},
		},
		{
			name: "string-object maps with k8s managed fields",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"objectField2": "str1-from-state",
								"state":        "state1-from-state",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.0,
								"objectField2": "str2-from-state",
								"state":        "state2-from-state",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref2",
									},
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.1,
								"objectField2": "str1-from-spec",
								"state":        "state1-from-spec",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.2,
								"objectField2": "str2-from-spec",
								"state":        "state2-from-spec",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"name": "my-ref2",
									},
								},
							},
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:stringObjectMapKey": map[string]interface{}{
							"f:someKey": map[string]interface{}{
								".":                         emptyObject,
								"f:objectField1":            emptyObject,
								"f:objectReferenceArrayKey": emptyObject,
							},
							"f:someOtherKey": map[string]interface{}{
								".":                         emptyObject,
								"f:objectField1":            emptyObject,
								"f:objectReferenceArrayKey": emptyObject,
							},
						},
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"stringObjectMapKey": map[string]interface{}{
					"someKey": map[string]interface{}{
						"objectField1": 1.1,
						"objectField2": "str1-from-state",
						"state":        "state1-from-state",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
					"someOtherKey": map[string]interface{}{
						"objectField1": 2.2,
						"objectField2": "str2-from-state",
						"state":        "state2-from-state",
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"name": "my-ref2",
							},
						},
					},
				},
			},
		},
		{
			name: "values in lists of objects ignore managed fields",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 1.0,
								"field2": "strval2",
							},
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:unrelated": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				// reflects the traditional fully-k8s-managed overlay of
				// the spec list on the live state list
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 1.0,
						"field2": "strval2",
					},
				},
			},
		},
		{
			name: "values in primitive lists are always sourced from state",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
							"myString2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
						},
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:primitiveArrayKey": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				// reflects solely the live state
				"primitiveArrayKey": []interface{}{
					"myString1",
					"myString2",
				},
			},
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualSpec, actualStatus, err := kcclite.ResolveSpecAndStatus(tc.state, tc.dclResource, smLoader)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := actualSpec, tc.expectedSpec; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
			if got, want := actualStatus, tc.expectedStatus; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestResolveSpecAndStatusWithDesiredStateInSpecAndObservedStatesInStatus(t *testing.T) {
	tests := []struct {
		name           string
		state          *unstructured.Unstructured
		dclResource    *dcl.Resource
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
	}{
		{
			name: "only persist specified fields in spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    int64(1),
						"floatKey":  0.5,
						"stringKey": "StringVal",
						"boolKey":   false,
						"primitiveArrayKey": []interface{}{
							"myString1",
							"myString2",
						},
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
							},
						},
						"projectRef": map[string]interface{}{
							"external": "project_id",
						},
						"sensitiveField": map[string]interface{}{
							"value": "secret-val1",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":   int64(1),
						"floatKey": 0.5,
						"projectRef": map[string]interface{}{
							"name": "project-name",
						},
						"sensitiveField": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":   int64(1),
				"floatKey": 0.5,
				"projectRef": map[string]interface{}{
					"name": "project-name",
				},
				"sensitiveField": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
		},
		{
			name: "observed states for output-only fields are persisted in status",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"statusField": "statusVal1",
						"nestedObjectKey": map[string]interface{}{
							"nestedStatusField": "statusVal2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Status: map[string]interface{}{
						"statusField": "statusVal2",
					},
				},
				Schema: testSchema(),
			},
			expectedStatus: map[string]interface{}{
				"statusField": "statusVal1",
				"nestedObjectKey": map[string]interface{}{
					"nestedStatusField": "statusVal2",
				},
			},
		},
		{
			name: "persist desired state in spec and output-only observed state in status",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    int64(1),
						"floatKey":  0.5,
						"stringKey": "StringVal",
						"boolKey":   false,
					},
					"status": map[string]interface{}{
						"statusField": "statusVal1",
						"nestedObjectKey": map[string]interface{}{
							"nestedStatusField": "statusVal2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":   int64(1),
						"floatKey": 0.5,
					},
					Status: map[string]interface{}{
						"statusField": "statusVal2",
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":   int64(1),
				"floatKey": 0.5,
			},
			expectedStatus: map[string]interface{}{
				"statusField": "statusVal1",
				"nestedObjectKey": map[string]interface{}{
					"nestedStatusField": "statusVal2",
				},
			},
		},
		{
			name: "preserve lists of objects unmodified in spec if specified",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
								"field2": "strval1",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval3",
							},
							map[string]interface{}{
								"field1": 1.0,
								"field2": "strval4",
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"objectArrayKey": []interface{}{
							map[string]interface{}{
								"field1": 0.5,
							},
							map[string]interface{}{
								"field1": 0.7,
								"field2": "strval2",
							},
							map[string]interface{}{
								"field1": 0.7,
							},
						},
					},
					Status: map[string]interface{}{
						"statusField": "statusVal2",
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"objectArrayKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
					},
					map[string]interface{}{
						"field1": 0.7,
						"field2": "strval2",
					},
					map[string]interface{}{
						"field1": 0.7,
					},
				},
			},
		},
		{
			name: "primitive lists are preserved with specified values",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
							"myString2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"primitiveArrayKey": []interface{}{
							"myString1",
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"primitiveArrayKey": []interface{}{
					"myString1",
				},
			},
		},
		{
			name: "only persist specified nested fields in spec",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField2": "strval2",
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField2": "strval2",
				},
			},
		},
		{
			name: "string-object maps",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"state":        "state1",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 2.0,
								"state":        "state2",
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref1",
									},
									map[string]interface{}{
										"external": "projects/my-project-1/bars/my-ref2",
									},
								},
							},
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"stringObjectMapKey": map[string]interface{}{
							"someKey": map[string]interface{}{
								"objectField1": 1.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"external": "my-ref3",
									},
								},
							},
							"someOtherKey": map[string]interface{}{
								"objectField1": 3.0,
								"objectReferenceArrayKey": []interface{}{
									map[string]interface{}{
										"name": "my-ref1",
									},
									map[string]interface{}{
										"name": "my-ref2",
									},
								},
							},
						},
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"stringObjectMapKey": map[string]interface{}{
					"someKey": map[string]interface{}{
						"objectField1": 1.0,
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"external": "my-ref3",
							},
						},
					},
					"someOtherKey": map[string]interface{}{
						"objectField1": 3.0,
						"objectReferenceArrayKey": []interface{}{
							map[string]interface{}{
								"name": "my-ref1",
							},
							map[string]interface{}{
								"name": "my-ref2",
							},
						},
					},
				},
			},
		},
		{
			name: "resourceID in spec will be persisted",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":     "2",
						"resourceID": "someVal",
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":     "1",
						"resourceID": "someVal",
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:intKey": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":     "1",
				"resourceID": "someVal",
			},
		},
		{
			name: "server-generated id is retrieved from state and persisted",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":     "2",
						"resourceID": "server-generated-value",
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey": "1",
					},
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":     "1",
				"resourceID": "server-generated-value",
			},
		},
		{
			name: "fields in spec are persisted even if they not in managed fields set",
			state: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"intKey":    int64(1),
						"floatKey":  0.5,
						"boolKey":   false,
						"stringKey": "someVal",
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": false,
							"nestedField2": "strval2",
						},
					},
				},
			},
			dclResource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"intKey":   int64(1),
						"floatKey": 0.5,
						"boolKey":  true,
					},
					ManagedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
						"f:unrelated": emptyObject,
					}),
				},
				Schema: testSchema(),
			},
			expectedSpec: map[string]interface{}{
				"intKey":   int64(1),
				"floatKey": 0.5,
				"boolKey":  true,
			},
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			k8s.SetAnnotation(k8s.StateIntoSpecAnnotation, k8s.StateAbsentInSpec, tc.dclResource)
			actualSpec, actualStatus, err := kcclite.ResolveSpecAndStatus(tc.state, tc.dclResource, smLoader)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := actualSpec, tc.expectedSpec; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
			if got, want := actualStatus, tc.expectedStatus; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func newBarUnstructuredWithResourceID(t *testing.T, name, ns string, readyStatus corev1.ConditionStatus) *unstructured.Unstructured {
	u := test.NewBarUnstructured(name, ns, readyStatus)
	if err := unstructured.SetNestedField(u.Object, name, "spec", k8s.ResourceIDFieldName); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return u
}

func TestMain(m *testing.M) {
	testmain.ForUnitTestsWithCRDs(m, test.FakeCRDsWithHierarchicalResources(), &mgr)
}
