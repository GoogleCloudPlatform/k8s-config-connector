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

package extension_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/nasa9084/go-openapi"
)

func TestGetReferenceFieldName(t *testing.T) {
	tests := []struct {
		name               string
		path               []string
		convertedFieldName string
		schema             *openapi.Schema
		hasError           bool
	}{
		{
			name:               "append Ref to the original field name",
			path:               []string{"foo"},
			convertedFieldName: "fooRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "append Ref to the original nested field name",
			path:               []string{"object", "foo"},
			convertedFieldName: "fooRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "serviceAccountEmail is converted to serviceAccountRef",
			path:               []string{"serviceAccountEmail"},
			convertedFieldName: "serviceAccountRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "nested serviceAccountEmail is converted to serviceAccountRef",
			path:               []string{"object", "serviceAccountEmail"},
			convertedFieldName: "serviceAccountRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "kmsKeyName is converted to kmsKeyRef",
			path:               []string{"kmsKeyName"},
			convertedFieldName: "kmsKeyRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "nested kmsKeyName is converted to kmsKeyRef",
			path:               []string{"object", "kmsKeyName"},
			convertedFieldName: "kmsKeyRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "groupId is converted to groupRef",
			path:               []string{"groupId"},
			convertedFieldName: "groupRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "nested groupId is converted to groupRef",
			path:               []string{"object", "groupId"},
			convertedFieldName: "groupRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "resourceLink is converted to resourceRef",
			path:               []string{"resourceLink"},
			convertedFieldName: "resourceRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "nested resourceLink is converted to resourceRef",
			path:               []string{"object", "resourceLink"},
			convertedFieldName: "resourceRef",
			schema: &openapi.Schema{
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
		},
		{
			name:               "field name is preserved if the field is an array of references",
			path:               []string{"networks"},
			convertedFieldName: "networks",
			schema: &openapi.Schema{
				Type: "array",
				Items: &openapi.Schema{
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
			},
		},
		{
			name:               "nested field name is preserved if the field is an array of references",
			path:               []string{"object", "networks"},
			convertedFieldName: "networks",
			schema: &openapi.Schema{
				Type: "array",
				Items: &openapi.Schema{
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
			},
		},
		{
			name:               "field name is converted if the field can take the reference of different resource kinds",
			path:               []string{"group"},
			convertedFieldName: "groupRef",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "SomeService/SomeKind",
							"field":    "selfLink",
						},
						map[interface{}]interface{}{
							"resource": "SomeOtherService/SomeOtherKind",
							"field":    "selfLink",
						},
					},
				},
			},
		},
		{
			name:               "nested field name is converted if the field can take the reference of different resource kinds",
			path:               []string{"object", "group"},
			convertedFieldName: "groupRef",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "SomeService/SomeKind",
							"field":    "selfLink",
						},
						map[interface{}]interface{}{
							"resource": "SomeOtherService/SomeOtherKind",
							"field":    "selfLink",
						},
					},
				},
			},
		},
		{
			name:     "cannot get reference field name for top-level parent field",
			path:     []string{"parent"},
			hasError: true,
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Project",
							"field":    "name",
							"parent":   true,
						},
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Folder",
							"field":    "name",
							"parent":   true,
						},
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Organization",
							"field":    "name",
							"parent":   true,
						},
					},
				},
			},
		},
		{
			name:               "field name is updated for nested parent field that can take the reference of different resource kinds",
			path:               []string{"object", "parent"},
			convertedFieldName: "parentRef",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Project",
							"field":    "name",
							"parent":   true,
						},
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Folder",
							"field":    "name",
							"parent":   true,
						},
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Organization",
							"field":    "name",
							"parent":   true,
						},
					},
				},
			},
		},
		{
			name:               "append Ref to nested parent field that can take the reference of only one kind",
			path:               []string{"object", "parent"},
			convertedFieldName: "parentRef",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Cloudresourcemanager/Project",
							"field":    "name",
							"parent":   true,
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
			if !extension.IsReferenceField(tc.schema) {
				t.Fatalf("expect the field to be a reference field")
			}
			actual, err := extension.GetReferenceFieldName(tc.path, tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got no error, wanted one")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tc.convertedFieldName {
				t.Fatalf("got the converted reference field name as %v, want %v", actual, tc.convertedFieldName)
			}
		})
	}
}

func TestIsResourceIDFieldServerGenerated(t *testing.T) {
	tests := []struct {
		name           string
		schema         *openapi.Schema
		expectedResult bool
	}{
		{
			name: "missing 'x-dcl-server-generated-parameter' extension is treated as user-specified name",
			schema: &openapi.Schema{
				Type: "string",
			},
			expectedResult: false,
		},
		{
			name: "server-generated id",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-server-generated-parameter": true,
				},
			},
			expectedResult: true,
		},
		{
			name: "the name field is specifically marked as non-server-generated",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-server-generated-parameter": false,
				},
			},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := extension.IsResourceIDFieldServerGenerated(tc.schema)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tc.expectedResult {
				t.Fatalf("got the result as %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}

func TestHasSensitiveFields(t *testing.T) {
	tests := []struct {
		name           string
		schema         *openapi.Schema
		expectedResult bool
		hasError       bool
	}{
		{
			name: "has sensitive string field",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-sensitive": true,
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive string field in string array",
			schema: &openapi.Schema{
				Type: "array",
				Items: &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-sensitive": true,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive string field in object array",
			schema: &openapi.Schema{
				Type: "array",
				Items: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"foo": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-sensitive": true,
							},
						},
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive string field in object",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-sensitive": true,
						},
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive string field in string map",
			schema: &openapi.Schema{
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-sensitive": true,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive string field in object map",
			schema: &openapi.Schema{
				Type: "object",
				AdditionalProperties: &openapi.Schema{
					Type: "object",
					Properties: map[string]*openapi.Schema{
						"foo": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-sensitive": true,
							},
						},
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has sensitive number field",
			schema: &openapi.Schema{
				Type: "number",
				Extension: map[string]interface{}{
					"x-dcl-sensitive": true,
				},
			},
			hasError: true,
		},
		{
			name: "has sensitive array field",
			schema: &openapi.Schema{
				Type: "array",
				Items: &openapi.Schema{
					Type: "string",
				},
				Extension: map[string]interface{}{
					"x-dcl-sensitive": true,
				},
			},
			hasError: true,
		},
		{
			name: "has sensitive object field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-sensitive": true,
				},
			},
			hasError: true,
		},
		{
			name: "has no sensitive field",
			schema: &openapi.Schema{
				Type: "string",
			},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := extension.HasSensitiveFields(tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got no error, but want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHasStateHint(t *testing.T) {
	tests := []struct {
		name           string
		schema         *openapi.Schema
		expectedResult bool
		hasError       bool
	}{
		{
			name: "has state hint",
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-uses-state-hint": true,
				},
			},
			expectedResult: true,
		},
		{
			name: "has no state hint",
			schema: &openapi.Schema{
				Type: "object",
			},
			expectedResult: false,
		},
		{
			name: "wrong type for x-dcl-uses-state-hint extension",
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-uses-state-hint": "true",
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := extension.HasStateHint(tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatal("got no error, but want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestIsMutableButUnreadableField(t *testing.T) {
	tests := []struct {
		name           string
		schema         *openapi.Schema
		expectedResult bool
		hasError       bool
	}{
		{
			name: "mutable-but-unreadable field",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-mutable-unreadable": true,
				},
			},
			expectedResult: true,
		},
		{
			name: "mutable and readable field",
			schema: &openapi.Schema{
				Type: "string",
			},
			expectedResult: false,
		},
		{
			name: "immutable and unreadable field",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-mutable-unreadable": true,
					"x-kubernetes-immutable":   true,
				},
			},
			expectedResult: false,
		},
		{
			name: "wrong type for x-dcl-mutable-unreadable extension",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-mutable-unreadable": "true",
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := extension.IsMutableButUnreadableField(tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatal("got no error, but want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHasIam(t *testing.T) {
	tests := []struct {
		name           string
		schema         *openapi.Schema
		expectedResult bool
		hasError       bool
	}{
		{
			name: "has iam policy",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-has-iam": true,
				},
			},
			expectedResult: true,
		},
		{
			name: "has no iam policy",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-has-iam": false,
				},
			},
			expectedResult: false,
		},
		{
			name: "wrong type for x-dcl-has-iam extension",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-has-iam": "true",
				},
			},
			expectedResult: false,
			hasError:       true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := extension.HasIam(tc.schema)
			if tc.hasError {
				if err == nil {
					t.Fatal("got no error, but want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
