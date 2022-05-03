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
	"strings"
	"testing"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAllowAllKnownFields(t *testing.T) {
	obj := map[string]interface{}{
		"integer-key": 2,
		"map-key": map[string]interface{}{
			"nested-key": false,
		},
		"unvalidated-map-key": map[string]interface{}{
			"unvalidated-key": true,
		},
		"array-key": []interface{}{
			map[string]interface{}{
				"map-inside-array-key": "map-inside-array-val",
			},
		},
	}
	schema := &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"integer-key": {Type: "integer"},
			"unused-key":  {Type: "boolean"},
			"map-key": {
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"nested-key": {Type: "boolean"},
				},
			},
			"unvalidated-map-key": {Type: "object"},
			"array-key": {
				Type: "array",
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"unused-key":           {Type: "integer"},
							"map-inside-array-key": {Type: "string"},
						},
					},
				},
			},
		},
	}
	if err := validateNoUnknownFields(schema, obj, ""); err != nil {
		t.Fatalf("expected object to be valid but got error: %v", err)
	}
}

func TestDenyUnknownKeyInMap(t *testing.T) {
	obj := map[string]interface{}{
		"valid-key":   "valid-val",
		"unknown-key": 0,
	}
	schema := &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"valid-key": {Type: "string"},
		},
	}
	err := validateNoUnknownFields(schema, obj, "")
	if err == nil {
		t.Fatalf("expected object to not be allowed")
	}
	if !strings.Contains(err.Error(), `"unknown-key"`) {
		t.Fatalf("expected error to contain unknown key; error: %v", err)
	}
}

func TestDenyUnknownKeyInNestedMap(t *testing.T) {
	obj := map[string]interface{}{
		"map-key": map[string]interface{}{
			"valid-key":   "valid-val",
			"unknown-key": 0,
		},
	}
	schema := &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"map-key": {
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"valid-key": {Type: "string"},
				},
			},
		},
	}
	err := validateNoUnknownFields(schema, obj, "")
	if err == nil {
		t.Fatalf("expected object to not be allowed")
	}
	if !strings.Contains(err.Error(), `"map-key.unknown-key"`) {
		t.Fatalf("expected error to contain unknown key; error: %v", err)
	}
}

func TestDenyUnknownKeyInMapInsideArray(t *testing.T) {
	obj := map[string]interface{}{
		"array-key": []interface{}{
			map[string]interface{}{
				"valid-key":   "valid-val",
				"unknown-key": 0,
			},
		},
	}
	schema := &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"array-key": {
				Type: "array",
				Items: &apiextensions.JSONSchemaPropsOrArray{
					Schema: &apiextensions.JSONSchemaProps{
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"valid-key": {Type: "string"},
						},
					},
				},
			},
		},
	}
	err := validateNoUnknownFields(schema, obj, "")
	if err == nil {
		t.Fatalf("expected object to not be allowed")
	}
	if !strings.Contains(err.Error(), `"array-key[0].unknown-key"`) {
		t.Fatalf("expected error to contain unknown key; error: %v", err)
	}
}
