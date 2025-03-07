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

package crdutil_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/crdutil"

	"github.com/google/go-cmp/cmp"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestGetSchemaForFieldUnderObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		field          string
		parent         *apiextensions.JSONSchemaProps
		expectedSchema *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name:  "get schema for field under object",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type: "object",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Description: "field under object",
				Type:        "string",
			},
		},
		{
			name:  "get schema for field under map",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under map",
								Type:        "string",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Description: "field under map",
				Type:        "string",
			},
		},
		{
			name:  "get schema for field under array",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under array",
								Type:        "string",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Description: "field under array",
				Type:        "string",
			},
		},
		{
			name:  "get empty schema when field doesn't exist",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "string",
					},
				},
				Type: "object",
			},
			expectedSchema: nil,
		},
		{
			name:  "can't get schema for field under string field",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Type: "string",
			},
			hasError: true,
		},
		{
			name:  "can't get schema for incorrect object schema",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Description: "an object field without properties or additionalProperties",
				Type:        "object",
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			schema, _, err := crdutil.GetSchemaForFieldUnderObjectOrArray(tc.field, tc.parent)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := schema, tc.expectedSchema; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestSetSchemaForFieldUnderObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		field          string
		parent         *apiextensions.JSONSchemaProps
		fieldSchema    *apiextensions.JSONSchemaProps
		expectedSchema *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name:  "set schema for nonexistent field under object",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "field under object",
				Type:        "string",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type: "object",
			},
		},
		{
			name:  "set schema for nonexistent field under map",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "field under map",
				Type:        "string",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under map",
								Type:        "string",
							},
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
		},
		{
			name:  "set schema for nonexistent field under array",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "field under array",
				Type:        "string",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under array",
								Type:        "string",
							},
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
		},
		{
			name:  "set schema for existing field under object",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
					"test": {
						Description: "old description",
						Type:        "bool",
					},
				},
				Type: "object",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "field under object",
				Type:        "string",
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type: "object",
			},
		},
		{
			name:  "can't set schema under a string field",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Type: "string",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "test schema",
				Type:        "string",
			},
			hasError: true,
		},
		{
			name:  "can't set schema for incorrect object schema",
			field: "test",
			parent: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			fieldSchema: &apiextensions.JSONSchemaProps{
				Description: "test schema",
				Type:        "string",
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := crdutil.SetSchemaForFieldUnderObjectOrArray(tc.field, tc.parent, tc.fieldSchema); err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := tc.parent, tc.expectedSchema; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestGetRequiredRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		schema       *apiextensions.JSONSchemaProps
		expectedRule []string
		hasError     bool
	}{
		{
			name: "get required rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type:     "object",
				Required: []string{"requiredField"},
			},
			expectedRule: []string{"requiredField"},
		},
		{
			name: "get required rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under map",
								Type:        "string",
							},
						},
						Type:     "object",
						Required: []string{"requiredField"},
					},
				},
				Type: "object",
			},
			expectedRule: []string{"requiredField"},
		},
		{
			name: "get required rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under array",
								Type:        "string",
							},
						},
						Type:     "object",
						Required: []string{"requiredField"},
					},
				},
				Type: "array",
			},
			expectedRule: []string{"requiredField"},
		},
		{
			name: "can't get rule for field under string field",
			schema: &apiextensions.JSONSchemaProps{
				Type:     "string",
				Required: []string{"nonretrievable"},
			},
			hasError: true,
		},
		{
			name: "can't get rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Description: "an object field without properties or additionalProperties",
				Type:        "object",
				Required:    []string{"nonretrievable"},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rule, err := crdutil.GetRequiredRuleForObjectOrArray(tc.schema)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := rule, tc.expectedRule; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestSetRequiredRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		schema         *apiextensions.JSONSchemaProps
		rule           []string
		expectedSchema *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name: "set required rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
			},
			rule: []string{"requiredField"},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type:     "object",
				Required: []string{"requiredField"},
			},
		},
		{
			name: "set required rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			rule: []string{"requiredField"},
			expectedSchema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type:     "object",
						Required: []string{"requiredField"},
					},
				},
				Type: "object",
			},
		},
		{
			name: "set required rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
			rule: []string{"requiredField"},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type:     "object",
						Required: []string{"requiredField"},
					},
				},
				Type: "array",
			},
		},
		{
			name: "can't set rulee under a string field",
			schema: &apiextensions.JSONSchemaProps{
				Type: "string",
			},
			rule:     []string{"unsettable"},
			hasError: true,
		},
		{
			name: "can't set rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			rule:     []string{"unsettable"},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := crdutil.SetRequiredRuleForObjectOrArray(tc.schema, tc.rule); err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := tc.schema, tc.expectedSchema; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestGetNotRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		schema       *apiextensions.JSONSchemaProps
		expectedRule *apiextensions.JSONSchemaProps
		hasError     bool
	}{
		{
			name: "get not rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type: "object",
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"requiredField"},
				},
			},
			expectedRule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
		},
		{
			name: "get not rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under map",
								Type:        "string",
							},
						},
						Type: "object",
						Not: &apiextensions.JSONSchemaProps{
							Required: []string{"requiredField"},
						},
					},
				},
				Type: "object",
			},
			expectedRule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
		},
		{
			name: "get not rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under array",
								Type:        "string",
							},
						},
						Type: "object",
						Not: &apiextensions.JSONSchemaProps{
							Required: []string{"requiredField"},
						},
					},
				},
				Type: "array",
			},
			expectedRule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
		},
		{
			name: "can't get rule for field under string field",
			schema: &apiextensions.JSONSchemaProps{
				Type: "string",
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"nonretrievable"},
				},
			},
			hasError: true,
		},
		{
			name: "can't get rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Description: "an object field without properties or additionalProperties",
				Type:        "object",
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"nonretrievable"},
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rule, err := crdutil.GetNotRuleForObjectOrArray(tc.schema)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := rule, tc.expectedRule; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestSetNotRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		schema         *apiextensions.JSONSchemaProps
		rule           *apiextensions.JSONSchemaProps
		expectedSchema *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name: "set required rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
			},
			rule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"requiredField"},
				},
			},
		},
		{
			name: "set required rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			rule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
						Not: &apiextensions.JSONSchemaProps{
							Required: []string{"requiredField"},
						},
					},
				},
				Type: "object",
			},
		},
		{
			name: "set required rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
			rule: &apiextensions.JSONSchemaProps{
				Required: []string{"requiredField"},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
						Not: &apiextensions.JSONSchemaProps{
							Required: []string{"requiredField"},
						},
					},
				},
				Type: "array",
			},
		},
		{
			name: "can't set rulee under a string field",
			schema: &apiextensions.JSONSchemaProps{
				Type: "string",
			},
			rule: &apiextensions.JSONSchemaProps{
				Required: []string{"unsettable"},
			},
			hasError: true,
		},
		{
			name: "can't set rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			rule: &apiextensions.JSONSchemaProps{
				Required: []string{"unsettable"},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := crdutil.SetNotRuleForObjectOrArray(tc.schema, tc.rule); err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := tc.schema, tc.expectedSchema; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestGetOneOfRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		schema       *apiextensions.JSONSchemaProps
		expectedRule []*apiextensions.JSONSchemaProps
		hasError     bool
	}{
		{
			name: "get oneOf rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"test": {
						Description: "field under object",
						Type:        "string",
					},
				},
				Type: "object",
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"field1"}},
					{Required: []string{"field2"}},
				},
			},
			expectedRule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
		},
		{
			name: "get oneOf rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under map",
								Type:        "string",
							},
						},
						Type: "object",
						OneOf: []apiextensions.JSONSchemaProps{
							{Required: []string{"field1"}},
							{Required: []string{"field2"}},
						},
					},
				},
				Type: "object",
			},
			expectedRule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
		},
		{
			name: "get oneOf rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"test": {
								Description: "field under array",
								Type:        "string",
							},
						},
						Type: "object",
						OneOf: []apiextensions.JSONSchemaProps{
							{Required: []string{"field1"}},
							{Required: []string{"field2"}},
						},
					},
				},
				Type: "array",
			},
			expectedRule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
		},
		{
			name: "can't get rule for field under string field",
			schema: &apiextensions.JSONSchemaProps{
				Type: "string",
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"field1"}},
					{Required: []string{"field2"}},
				},
			},
			hasError: true,
		},
		{
			name: "can't get rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Description: "an object field without properties or additionalProperties",
				Type:        "object",
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"field1"}},
					{Required: []string{"field2"}},
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rule, err := crdutil.GetOneOfRuleForObjectOrArray(tc.schema)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := rule, tc.expectedRule; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}

func TestSetOneOfRuleForObjectOrArray(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		schema         *apiextensions.JSONSchemaProps
		rule           []*apiextensions.JSONSchemaProps
		expectedSchema *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name: "set required rule under object",
			schema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
			},
			rule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"otherField": {
						Type: "bool",
					},
				},
				Type: "object",
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"field1"}},
					{Required: []string{"field2"}},
				},
			},
		},
		{
			name: "set required rule under map",
			schema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "object",
			},
			rule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
						OneOf: []apiextensions.JSONSchemaProps{
							{Required: []string{"field1"}},
							{Required: []string{"field2"}},
						},
					},
				},
				Type: "object",
			},
		},
		{
			name: "set required rule under array",
			schema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
					},
				},
				Type: "array",
			},
			rule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
			expectedSchema: &apiextensions.JSONSchemaProps{
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Properties: map[string]apiextensions.JSONSchemaProps{
							"otherField": {
								Type: "bool",
							},
						},
						Type: "object",
						OneOf: []apiextensions.JSONSchemaProps{
							{Required: []string{"field1"}},
							{Required: []string{"field2"}},
						},
					},
				},
				Type: "array",
			},
		},
		{
			name: "can't set rulee under a string field",
			schema: &apiextensions.JSONSchemaProps{
				Type: "string",
			},
			rule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
			hasError: true,
		},
		{
			name: "can't set rule for incorrect object schema",
			schema: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			rule: []*apiextensions.JSONSchemaProps{
				{Required: []string{"field1"}},
				{Required: []string{"field2"}},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := crdutil.SetOneOfRuleForObjectOrArray(tc.schema, tc.rule); err != nil {
				if !tc.hasError {
					t.Fatalf("got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := tc.schema, tc.expectedSchema; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}
