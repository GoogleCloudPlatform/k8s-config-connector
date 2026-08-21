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

package livestate_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/livestate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var testDCLSchema = &openapi.Schema{
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
								"path":     "statusField",
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
						"path":     "name",
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
										"path":     "name",
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
									"path":     "name",
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
							"resource": "test1/Bar",
							"path":     "name",
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

func TestSetMutableButUnreadableFields(t *testing.T) {
	tests := []struct {
		name                     string
		krmObj                   *unstructured.Unstructured
		mutableButUnreadableSpec map[string]interface{}
		path                     []string
		schema                   *openapi.Schema
		expectedResult           *unstructured.Unstructured
		hasError                 bool
	}{
		{
			name: "set primitive fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"boolKey":   true,
				"floatKey":  1.1,
				"intKey":    10,
				"stringKey": "testValue",
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"boolKey":   true,
						"floatKey":  1.1,
						"intKey":    10,
						"stringKey": "testValue",
					},
				},
			},
		},
		{
			name: "set primitive array",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"primitiveArrayKey": []interface{}{"test1", "test2"},
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"primitiveArrayKey": []interface{}{"test1", "test2"},
					},
				},
			},
		},
		{
			name: "set nested fields",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField1": true,
					"nestedField2": "test",
				},
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "test",
						},
					},
				},
			},
		},
		{
			name: "set map",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"label1": "value1",
					"label2": "value2",
				},
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"mapKey": map[string]interface{}{
							"label1": "value1",
							"label2": "value2",
						},
					},
				},
			},
		},
		{
			name: "non primitive array should return an error",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"objectArrayKey": map[string]interface{}{
					"field1": 10,
				},
			},
			path:     []string{"spec"},
			schema:   testDCLSchema,
			hasError: true,
		},
		{
			name: "overwrite the existent value in krmObj with the value in" +
				"mutableButUnreadableSpec",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "oldVal",
						},
					},
				},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedField2": "newVal",
				},
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"nestedObjectKey": map[string]interface{}{
							"nestedField1": true,
							"nestedField2": "newVal",
						},
					},
				},
			},
		},
		{
			name: "set sensitive path with a plain text value",
			krmObj: &unstructured.Unstructured{
				Object: map[string]interface{}{},
			},
			mutableButUnreadableSpec: map[string]interface{}{
				"sensitiveField": map[string]interface{}{
					"value": "sensitive-value",
				},
			},
			path:   []string{"spec"},
			schema: testDCLSchema,
			expectedResult: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "sensitive-value",
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := livestate.SetMutableButUnreadableFields(tc.krmObj, tc.mutableButUnreadableSpec, tc.path, tc.schema, nil, "", nil)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but want an error setting mutable-but-unreadable fields")
				}
				return
			}
			if err != nil {
				t.Fatalf("error setting mutable-but-unreadable fields: %v", err)
			}

			if got, want := result, tc.expectedResult; !test.Equals(t, got, want) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}
