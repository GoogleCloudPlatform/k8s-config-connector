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

package k8s_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/google/go-cmp/cmp"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestRenameStatusFieldsWithReservedNames(t *testing.T) {
	tests := []struct {
		name           string
		resourceName   string
		originalStatus *apiextensions.JSONSchemaProps
		expectedStatus *apiextensions.JSONSchemaProps
		hasError       bool
	}{
		{
			name: "fields that don't collide with reserved names are left alone",
			originalStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"randomField": {
						Type: "string",
					},
				},
			},
			expectedStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"randomField": {
						Type: "string",
					},
				},
			},
		},
		{
			name: "fields that collide with reserved names are renamed",
			originalStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"conditions": {
						Type: "string",
					},
					"observedGeneration": {
						Type: "string",
					},
				},
			},
			expectedStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"resourceConditions": {
						Type: "string",
					},
					"resourceObservedGeneration": {
						Type: "string",
					},
				},
			},
		},
		{
			name:         "fields that collide with reserved names but resource is in exclude list",
			resourceName: "google_storage_default_object_access_control",
			originalStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"conditions": {
						Type: "string",
					},
					"observedGeneration": {
						Type: "string",
					},
				},
			},
			expectedStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"conditions": {
						Type: "string",
					},
					"observedGeneration": {
						Type: "string",
					},
				},
			},
		},
		{
			name: "error if status has fields that collide with the renames of the reserved names",
			originalStatus: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"resourceConditions": {
						Type: "string",
					},
					"resourceObservedGeneration": {
						Type: "string",
					},
				},
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualStatus, err := k8s.RenameStatusFieldsWithReservedNamesIfResourceNotExcluded(tc.resourceName, tc.originalStatus)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil error, want an error")
				}
				return
			}
			if !test.Equals(t, tc.expectedStatus, actualStatus) {
				t.Fatalf("unexpected diff in returned status (-want +got): \n%v", cmp.Diff(tc.expectedStatus, actualStatus))
			}
		})
	}
}
