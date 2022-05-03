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

package dcl_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
)

func TestTrimNilFields(t *testing.T) {
	tests := []struct {
		name     string
		in       map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "object with no nil fields",
			in: map[string]interface{}{
				"boolField":   true,
				"intField":    1,
				"stringField": "string",
				"floatField":  0.1,
				"objField": map[string]interface{}{
					"boolField":   true,
					"intField":    1,
					"stringField": "string",
					"floatField":  0.1,
					"objField": map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": "string",
						"floatField":  0.1,
					},
					"listOfObjsField": []interface{}{
						map[string]interface{}{
							"boolField":   true,
							"intField":    1,
							"stringField": "string",
							"floatField":  0.1,
						},
					},
				},
				"listOfStringsField": []interface{}{
					"stringOne",
					"stringTwo",
					"stringThree",
				},
				"listOfIntsField": []interface{}{
					1000,
					2000,
					3000,
				},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": "string",
						"floatField":  0.1,
					},
					map[string]interface{}{
						"boolField":   false,
						"intField":    2,
						"stringField": "string2",
						"floatField":  0.2,
						"objField": map[string]interface{}{
							"boolField":   true,
							"intField":    1,
							"stringField": "string",
							"floatField":  0.1,
						},
					},
				},
			},
			expected: map[string]interface{}{
				"boolField":   true,
				"intField":    1,
				"stringField": "string",
				"floatField":  0.1,
				"objField": map[string]interface{}{
					"boolField":   true,
					"intField":    1,
					"stringField": "string",
					"floatField":  0.1,
					"objField": map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": "string",
						"floatField":  0.1,
					},
					"listOfObjsField": []interface{}{
						map[string]interface{}{
							"boolField":   true,
							"intField":    1,
							"stringField": "string",
							"floatField":  0.1,
						},
					},
				},
				"listOfStringsField": []interface{}{
					"stringOne",
					"stringTwo",
					"stringThree",
				},
				"listOfIntsField": []interface{}{
					1000,
					2000,
					3000,
				},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": "string",
						"floatField":  0.1,
					},
					map[string]interface{}{
						"boolField":   false,
						"intField":    2,
						"stringField": "string2",
						"floatField":  0.2,
						"objField": map[string]interface{}{
							"boolField":   true,
							"intField":    1,
							"stringField": "string",
							"floatField":  0.1,
						},
					},
				},
			},
		},
		{
			name: "object with nil fields",
			in: map[string]interface{}{
				"boolField":   true,
				"intField":    1,
				"stringField": nil,
				"floatField":  0.1,
				"objField":    map[string]interface{}(nil),
				"objFieldTwo": map[string]interface{}{
					"boolField":   true,
					"intField":    nil,
					"stringField": "string",
					"floatField":  0.1,
					"objField":    map[string]interface{}(nil),
					"objFieldTwo": map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": nil,
						"floatField":  0.1,
					},
					"listOfObjsField": []interface{}(nil),
				},
				"listOfStringsField": []interface{}(nil),
				"listOfIntsField": []interface{}{
					1000,
					2000,
					3000,
				},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"boolField":   true,
						"intField":    1,
						"stringField": nil,
						"floatField":  0.1,
					},
					map[string]interface{}{
						"boolField":   nil,
						"intField":    2,
						"stringField": "string2",
						"floatField":  0.2,
						"objField":    map[string]interface{}(nil),
						"objFieldTwo": map[string]interface{}{
							"boolField":   true,
							"intField":    1,
							"stringField": nil,
							"floatField":  0.1,
						},
					},
				},
			},
			expected: map[string]interface{}{
				"boolField":  true,
				"intField":   1,
				"floatField": 0.1,
				"objFieldTwo": map[string]interface{}{
					"boolField":   true,
					"stringField": "string",
					"floatField":  0.1,
					"objFieldTwo": map[string]interface{}{
						"boolField":  true,
						"intField":   1,
						"floatField": 0.1,
					},
				},
				"listOfIntsField": []interface{}{
					1000,
					2000,
					3000,
				},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"boolField":  true,
						"intField":   1,
						"floatField": 0.1,
					},
					map[string]interface{}{
						"intField":    2,
						"stringField": "string2",
						"floatField":  0.2,
						"objFieldTwo": map[string]interface{}{
							"boolField":  true,
							"intField":   1,
							"floatField": 0.1,
						},
					},
				},
			},
		},
		{
			name: "object with empty (not nil) collection fields",
			in: map[string]interface{}{
				"objField": map[string]interface{}{},
				"objFieldTwo": map[string]interface{}{
					"objField":           map[string]interface{}{},
					"listOfStringsField": []interface{}{},
				},
				"listOfStringsField": []interface{}{},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"objField":           map[string]interface{}{},
						"listOfStringsField": []interface{}{},
					},
				},
			},
			expected: map[string]interface{}{
				"objField": map[string]interface{}{},
				"objFieldTwo": map[string]interface{}{
					"objField":           map[string]interface{}{},
					"listOfStringsField": []interface{}{},
				},
				"listOfStringsField": []interface{}{},
				"listOfObjsField": []interface{}{
					map[string]interface{}{
						"objField":           map[string]interface{}{},
						"listOfStringsField": []interface{}{},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dcl.TrimNilFields(tc.in)
			if !test.Equals(t, tc.expected, tc.in) {
				diff := cmp.Diff(tc.expected, tc.in)
				t.Fatalf("unexpected diff in output map (-want +got):\n%v", diff)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		name     string
		in       interface{}
		expected bool
	}{
		{
			name:     "nil",
			in:       nil,
			expected: true,
		},
		{
			name:     "nil object",
			in:       map[string]interface{}(nil),
			expected: true,
		},
		{
			name:     "nil list",
			in:       []interface{}(nil),
			expected: true,
		},
		{
			name:     "non-nil primitive",
			in:       "stringValue",
			expected: false,
		},
		{
			name: "non-nil object",
			in: map[string]interface{}{
				"key": "val",
			},
			expected: false,
		},
		{
			name: "non-nil list",
			in: []interface{}{
				"val",
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			out := dcl.IsNil(tc.in)
			if out != tc.expected {
				t.Fatalf("got %v, want %v", out, tc.expected)
			}
		})
	}
}
