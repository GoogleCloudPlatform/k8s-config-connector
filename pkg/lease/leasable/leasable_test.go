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

package leasable

import (
	"testing"

	"github.com/nasa9084/go-openapi"
)

func TestDCLSchemaSupportsLeasing(t *testing.T) {
	tests := []struct {
		name     string
		schema   *openapi.Schema
		expected bool
	}{
		{
			name: "no labels field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"intKey": &openapi.Schema{
						Type: "integer",
					},
				},
			},
			expected: false,
		},
		{
			name: "immutable labels field",
			schema: &openapi.Schema{
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
			expected: false,
		},
		{
			name: "mutable labels field",
			schema: &openapi.Schema{
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
			expected: true,
		},
		{
			name: "immutable nested labels field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"setting": &openapi.Schema{
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"labels": &openapi.Schema{
								Type: "string",
								Extension: map[string]interface{}{
									"x-kubernetes-immutable": true,
								},
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "setting.labels",
				},
			},
			expected: false,
		},
		{
			name: "mutable nested labels field",
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"setting": &openapi.Schema{
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"labels": &openapi.Schema{
								Type: "string",
							},
						},
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "setting.labels",
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := DCLSchemaSupportsLeasing(tc.schema)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tc.expected {
				t.Fatalf("expect to get %v, but got %v", tc.expected, res)
			}
		})
	}
}
