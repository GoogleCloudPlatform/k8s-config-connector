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

package dclschemaloader

import (
	"reflect"
	"testing"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/nasa9084/go-openapi"
)

func TestDCLSchemaLoader_GetDCLSchema(t *testing.T) {
	loader, err := New()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name          string
		shouldSucceed bool
		stv           dclunstruct.ServiceTypeVersion
	}{
		{
			name:          "valid ServiceTypeVersion",
			shouldSucceed: true,
			stv: dclunstruct.ServiceTypeVersion{
				Service: "billingbudgets",
				Type:    "Budget",
				Version: "beta",
			},
		},
		{
			name:          "invalid ServiceTypeVersion",
			shouldSucceed: false,
			stv: dclunstruct.ServiceTypeVersion{
				Service: "invalidService",
				Type:    "InvalidResource",
				Version: "ga",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			schema, err := loader.GetDCLSchema(tc.stv)
			if tc.shouldSucceed {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error, instead got nil")
				}
				return
			}
			if schema.Title != tc.stv.Type {
				t.Fatalf("got the schema for type %v, but want %v", schema.Title, tc.stv.Type)
			}
		})
	}
}

func TestDCLSchemaLoader_CheckPropertiesForRefs(t *testing.T) {
	doc := &openapi.Document{
		Components: &openapi.Components{
			Schemas: map[string]*openapi.Schema{
				"Bar": {
					Description: "test bar",
					Type:        "string",
				},
				"BarBar": {
					Properties: map[string]*openapi.Schema{
						"Foo": {
							Ref: "#/components/schemas/Bar",
						},
					},
					Type: "object",
				},
			},
		},
	}
	tests := []struct {
		name           string
		testField      string
		inputSchema    *openapi.Schema
		expectedOutput *openapi.Schema
	}{
		{
			name:      "simple ref",
			testField: "Foo",
			inputSchema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"Foo": {
						Ref: "#/components/schemas/Bar",
					},
				},
				Type: "object",
			},
			expectedOutput: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"Foo": {
						Description: "test bar",
						Type:        "string",
					},
				},
				Type: "object",
			},
		},
		{
			name:      "nested property ref",
			testField: "TopLevel",
			inputSchema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"TopLevel": {
						Properties: map[string]*openapi.Schema{
							"Foo": {
								Ref: "#/components/schemas/Bar",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			expectedOutput: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"TopLevel": {
						Properties: map[string]*openapi.Schema{
							"Foo": {
								Description: "test bar",
								Type:        "string",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
		},
		{
			name:      "nested reference's ref",
			testField: "FooBaz",
			inputSchema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"FooBaz": {
						Ref: "#/components/schemas/BarBar",
					},
				},
				Type: "object",
			},
			expectedOutput: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"FooBaz": {
						Properties: map[string]*openapi.Schema{
							"Foo": {
								Description: "test bar",
								Type:        "string",
							},
						},
						Type: "object",
					},
				},
			},
		},
		{
			name:      "items array ref",
			testField: "Foo",
			inputSchema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"Foo": {
						Items: &openapi.Schema{
							Ref: "#/components/schemas/Bar",
						},
						Type: "array",
					},
				},
				Type: "object",
			},
			expectedOutput: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"Foo": {
						Items: &openapi.Schema{
							Description: "test bar",
							Type:        "string",
						},
						Type: "array",
					},
				},
				Type: "object",
			},
		},
		{
			name:      "items array property ref",
			testField: "FooList",
			inputSchema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"FooList": {
						Items: &openapi.Schema{
							Properties: map[string]*openapi.Schema{
								"Foo": {
									Ref: "#/components/schemas/Bar",
								},
							},
							Type: "object",
						},
						Type: "array",
					},
				},
				Type: "object",
			},
			expectedOutput: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"FooList": {
						Items: &openapi.Schema{
							Properties: map[string]*openapi.Schema{
								"Foo": {
									Description: "test bar",
									Type:        "string",
								},
							},
							Type: "object",
						},
						Type: "array",
					},
				},
				Type: "object",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := CheckAndResolveRefs(tc.inputSchema, doc); err != nil {
				t.Fatalf("%v", err)
			}
			expectedProp := *tc.expectedOutput.Properties[tc.testField]
			actualProp := *tc.inputSchema.Properties[tc.testField]
			if !reflect.DeepEqual(expectedProp, actualProp) {
				t.Fatalf("results mismatch: got %v, want %v", actualProp, expectedProp)
			}
		})
	}
}
