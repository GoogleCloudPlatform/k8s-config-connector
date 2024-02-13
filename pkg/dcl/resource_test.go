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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/nasa9084/go-openapi"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidateResourceIDIfSupported(t *testing.T) {
	meta := v1.ObjectMeta{
		Namespace: "test-ns",
		Name:      "test-resource",
	}
	tests := []struct {
		name         string
		resource     dcl.Resource
		expectedSpec map[string]interface{}
		hasError     bool
	}{
		{
			name: "spec.resourceID field not supported",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
				},
				Schema: &openapi.Schema{},
			},
		},
		{
			name: "non server-generated ID, spec.resourceID field set",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
					Spec: map[string]interface{}{
						"resourceID": "test-id",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
						},
					},
				},
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "test-id",
			},
			hasError: false,
		},
		{
			name: "non server-generated ID, spec.resourceID field unset",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
						},
					},
				},
			},
			hasError: false,
		},
		{
			name: "non server-generated ID, spec.resourceID empty",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
					Spec: map[string]interface{}{
						"resourceID": "",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
						},
					},
				},
			},
			hasError: true,
		},
		{
			name: "server-generated ID, spec.resourceID field set",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
					Spec: map[string]interface{}{
						"resourceID": "test-id",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
				},
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "test-id",
			},
			hasError: false,
		},
		{
			name: "server-generated ID, spec.resourceID field unset",
			resource: dcl.Resource{
				Resource: k8s.Resource{
					ObjectMeta: meta,
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
				},
			},
			hasError: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.resource.ValidateResourceIDIfSupported()
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("error setting resource ID: %v", err)
			}
			if got, want := tc.resource.Spec, tc.expectedSpec; !test.Equals(t, got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestHasServerGeneratedIDButNotConfigured(t *testing.T) {
	tests := []struct {
		name     string
		resource *dcl.Resource
		result   bool
	}{
		{
			name: "user specified name",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "my-name",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
						},
					},
					Type: "object",
				},
			},
			result: false,
		},
		{
			name: "configured server-generated id",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "server-generated-value",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
					Type: "object",
				},
			},
			result: false,
		},
		{
			name: "non-configured server-generated id",
			resource: &dcl.Resource{
				Resource: k8s.Resource{},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
					Type: "object",
				},
			},
			result: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tc.resource.HasServerGeneratedIDButNotConfigured()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tc.result {
				t.Fatalf("got: %v, want: %v", actual, tc.result)
			}
		})
	}
}

func TestHasServerGeneratedIDAndConfigured(t *testing.T) {
	tests := []struct {
		name     string
		resource *dcl.Resource
		result   bool
	}{
		{
			name: "user specified name",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "my-name",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
						},
					},
					Type: "object",
				},
			},
			result: false,
		},
		{
			name: "configured server-generated id",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "server-generated-value",
					},
				},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
					Type: "object",
				},
			},
			result: true,
		},
		{
			name: "non-configured server-generated id",
			resource: &dcl.Resource{
				Resource: k8s.Resource{},
				Schema: &openapi.Schema{
					Properties: map[string]*openapi.Schema{
						"name": {
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-server-generated-parameter": true,
							},
						},
					},
					Type: "object",
				},
			},
			result: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tc.resource.HasServerGeneratedIDAndConfigured()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != tc.result {
				t.Fatalf("got: %v, want: %v", actual, tc.result)
			}
		})
	}
}

func TestHasMutableButUnreadableFields(t *testing.T) {
	tests := []struct {
		name           string
		resource       *dcl.Resource
		expectedResult bool
		hasError       bool
	}{
		{
			name: "has mutable-but-unreadable fields",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "my-name",
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Extension: map[string]interface{}{
						"x-dcl-uses-state-hint": true,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "has no mutable-but-unreadable fields",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "my-name",
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
				},
			},
			expectedResult: false,
		},
		{
			name: "error our when checking for mutable-but-unreadable fields",
			resource: &dcl.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"resourceID": "my-name",
					},
				},
				Schema: &openapi.Schema{
					Type: "object",
					Extension: map[string]interface{}{
						"x-dcl-uses-state-hint": "true",
					},
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := tc.resource.HasMutableButUnreadableFields()
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
