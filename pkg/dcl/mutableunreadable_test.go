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

package dcl

import (
	"reflect"
	"sort"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
)

var mutableButUnreadableFieldAnnotationTests = []struct {
	name                                    string
	resource                                *Resource
	expectedMutableButUnreadableFieldsState map[string]interface{}
}{
	{
		name: "top-level fields",
		resource: &Resource{
			Resource: k8s.Resource{
				Spec: map[string]interface{}{
					"fieldA": "val1",
					"fieldB": int64(2),
					"fieldC": map[string]interface{}{
						"field1": "val1",
						"field2": "val2",
					},
					"fieldD": []interface{}{
						"val1",
						"val2",
					},
				},
			},
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fieldA": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"fieldB": {
						Type: "integer",
					},
					"fieldC": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"field1": {
								Type: "string",
							},
							"field2": {
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"fieldD": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
						},
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{
			"spec": map[string]interface{}{
				"fieldA": "val1",
				"fieldC": map[string]interface{}{
					"field1": "val1",
					"field2": "val2",
				},
				"fieldD": []interface{}{
					"val1",
					"val2",
				},
			},
		},
	},
	{
		name: "nested fields",
		resource: &Resource{
			Resource: k8s.Resource{
				Spec: map[string]interface{}{
					"parentField": map[string]interface{}{
						"fieldA": "val1",
						"fieldB": int64(2),
						"fieldC": map[string]interface{}{
							"field1": "val1",
							"field2": "val2",
						},
						"fieldD": []interface{}{
							"val1",
							"val2",
						},
					},
				},
			},
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parentField": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"fieldA": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"fieldB": {
								Type: "integer",
							},
							"fieldC": {
								Type: "object",
								Properties: map[string]*openapi.Schema{
									"field1": {
										Type: "string",
									},
									"field2": {
										Type: "string",
									},
								},
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"fieldD": {
								Type: "array",
								Items: &openapi.Schema{
									Type: "string",
								},
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
						},
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{
			"spec": map[string]interface{}{
				"parentField": map[string]interface{}{
					"fieldA": "val1",
					"fieldC": map[string]interface{}{
						"field1": "val1",
						"field2": "val2",
					},
					"fieldD": []interface{}{
						"val1",
						"val2",
					},
				},
			},
		},
	},
	{
		name: "sensitive field",
		resource: &Resource{
			Resource: k8s.Resource{
				Spec: map[string]interface{}{
					"secretValueField": map[string]interface{}{
						"value": "test-1",
					},
				},
			},
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"secretValueField": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
							"x-dcl-sensitive":          true,
						},
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{
			"spec": map[string]interface{}{
				"secretValueField": map[string]interface{}{
					"value": "test-1",
				},
			},
		},
	},
	{
		name: "no mutable-but-unreadable fields set in spec",
		resource: &Resource{
			Resource: k8s.Resource{
				Spec: map[string]interface{}{
					"fieldB": int64(2),
					"parentField": map[string]interface{}{
						"fieldB": int64(2),
					},
				},
			},
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fieldA": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"fieldB": {
						Type: "integer",
					},
					"parentField": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"fieldA": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"fieldB": {
								Type: "integer",
							},
						},
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{},
	},
	{
		name: "no fields marked mutable-but-unreadable",
		resource: &Resource{
			Resource: k8s.Resource{
				Spec: map[string]interface{}{
					"fieldA": "val1",
				},
			},
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fieldA": {
						Type: "string",
					},
					"fieldB": {
						Type: "integer",
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{},
	},
	{
		name: "no spec",
		resource: &Resource{
			Schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fieldA": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"fieldB": {
						Type: "integer",
					},
					"parentField": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"fieldA": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"fieldB": {
								Type: "integer",
							},
						},
					},
				},
			},
		},
		expectedMutableButUnreadableFieldsState: map[string]interface{}{},
	},
}

func TestMutableButUnreadableFieldsAnnotationFor(t *testing.T) {
	for _, tc := range mutableButUnreadableFieldAnnotationTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			annotationInString, err := MutableButUnreadableFieldsAnnotationFor(tc.resource)
			if err != nil {
				t.Fatal(err)
			}

			expectedStateInString, err := util.MarshalToJSONString(tc.expectedMutableButUnreadableFieldsState)
			if err != nil {
				t.Fatalf("error marshaling the expected state to string: %v", err)
			}
			if got, want := annotationInString, expectedStateInString; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestGetMutableButUnreadableFieldsFromAnnotations(t *testing.T) {
	for _, tc := range mutableButUnreadableFieldAnnotationTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mutableButUnreadableFields, err := GetMutableButUnreadableFieldsFromAnnotations(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := mutableButUnreadableFields, tc.expectedMutableButUnreadableFieldsState; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected mutable-but-unreadable fields diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestGetMutableButUnreadableDCLPathsInObject(t *testing.T) {
	tests := []struct {
		name            string
		schema          *openapi.Schema
		expectedResults []string
		hasError        bool
	}{
		{
			name: "mutable-but-unreadable primitive fields",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"bar": {
						Type: "integer",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"baz": {
						Type: "boolean",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
					"quz": {
						Type: "number",
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
				},
			},
			expectedResults: []string{"foo", "bar", "baz", "quz"},
		},
		{
			name: "mutable-but-unreadable object field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField1": {
								Type: "boolean",
							},
							"nestedField2": {
								Type: "string",
							},
						},
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
				},
			},
			expectedResults: []string{"fooObj"},
		},
		{
			name: "mutable-but-unreadable map field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooMap": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
						},
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
				},
			},
			expectedResults: []string{"fooMap"},
		},
		{
			name: "mutable-but-unreadable primitive array field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooArray": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
						},
						Extension: map[string]interface{}{
							"x-dcl-mutable-unreadable": true,
						},
					},
				},
			},
			expectedResults: []string{"fooArray"},
		},
		{
			name: "nested mutable-but-unreadable primitive fields",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField1": {
								Type: "boolean",
							},
							"nestedField2": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"nestedField3": {
								Type: "object",
								Properties: map[string]*openapi.Schema{
									"bar": {
										Type: "integer",
										Extension: map[string]interface{}{
											"x-dcl-mutable-unreadable": true,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedResults: []string{"fooObj.nestedField2", "fooObj.nestedField3.bar"},
		},
		{
			name: "nested mutable-but-unreadable object fields",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField": {
								Type: "string",
							},
							"nestedObj": {
								Type: "object",
								Properties: map[string]*openapi.Schema{
									"bar": {
										Type: "integer",
									},
									"baz": {
										Type: "boolean",
									},
								},
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
						},
					},
				},
			},
			expectedResults: []string{"fooObj.nestedObj"},
		},
		{
			name: "nested mutable-but-unreadable map fields",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedMap1": {
								Type: "object",
								AdditionalProperties: &openapi.Schema{
									Type: "string",
								},
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
							"nestedObj": {
								Type: "object",
								Properties: map[string]*openapi.Schema{
									"bar": {
										Type: "integer",
									},
									"nestedMap2": {
										Type: "object",
										AdditionalProperties: &openapi.Schema{
											Type: "string",
										},
										Extension: map[string]interface{}{
											"x-dcl-mutable-unreadable": true,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedResults: []string{"fooObj.nestedMap1", "fooObj.nestedObj.nestedMap2"},
		},
		{
			name: "nested mutable-but-unreadable primitive array fields",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedArray": {
								Type: "array",
								Items: &openapi.Schema{
									Type: "string",
								},
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
						},
					},
				},
			},
			expectedResults: []string{"fooObj.nestedArray"},
		},
		{
			name: "mutable-but-unreadable subfield under read-only field should be ignored",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooObj": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField1": {
								Type: "boolean",
							},
							"nestedField2": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-mutable-unreadable": true,
								},
							},
						},
						ReadOnly: true,
					},
				},
			},
			expectedResults: []string{},
		},
		{
			name: "mutable-but-unreadable subfield in map should be ignored",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"fooMap": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-mutable-unreadable": true,
							},
						},
					},
				},
			},
			expectedResults: []string{},
		},
		{
			name: "entry level mutable-but-unreadable field should be ignored",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField1": {
								Type: "boolean",
							},
							"nestedField2": {
								Type: "string",
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-mutable-unreadable": true,
				},
			},
			expectedResults: []string{},
		},
		{
			name: "entry level non-object schema should cause an error",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-mutable-unreadable": true,
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			paths, err := getMutableButUnreadableDCLPathsInObject([]string{}, tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but want an error getting mutable-but-unreadable fields")
				}
				return
			}
			if err != nil {
				t.Fatalf("error getting mutable-but-unreadable fields: %v", err)
			}

			results := make([]string, 0)
			for _, path := range paths {
				results = append(results, pathslice.ToString(path))
			}
			sort.Strings(results)
			sort.Strings(tc.expectedResults)
			if got, want := results, tc.expectedResults; !reflect.DeepEqual(got, want) {
				t.Fatalf("mutable-but-unreadable fields mismatch: got '%v', want '%v'", got, want)
			}
		})
	}
}
