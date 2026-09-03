// Copyright 2026 Google LLC
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

package errorhandling

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestShouldSkip(t *testing.T) {
	tests := []struct {
		name     string
		resource *unstructured.Unstructured
		want     bool
	}{
		{
			name: "normal resource - should not skip",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
					"kind":       "APIGatewayAPI",
					"metadata": map[string]interface{}{
						"name":      "test-resource",
						"namespace": "test-namespace",
					},
				},
			},
			want: false,
		},
		{
			name: "resource with UpdateFailedTerminalError and matching generation - should skip",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
					"kind":       "APIGatewayAPI",
					"metadata": map[string]interface{}{
						"name":       "test-resource",
						"namespace":  "test-namespace",
						"generation": int64(2),
					},
					"status": map[string]interface{}{
						"observedGeneration": int64(2),
						"conditions": []interface{}{
							map[string]interface{}{
								"type":   "Ready",
								"status": "False",
								"reason": "UpdateFailedTerminalError",
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "resource with UpdateFailedTerminalError and matching generation but ContinuousRetry annotation - should not skip",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
					"kind":       "APIGatewayAPI",
					"metadata": map[string]interface{}{
						"name":       "test-resource",
						"namespace":  "test-namespace",
						"generation": int64(2),
						"annotations": map[string]interface{}{
							"cnrm.cloud.google.com/error-handling-mode": "ContinuousRetry",
						},
					},
					"status": map[string]interface{}{
						"observedGeneration": int64(2),
						"conditions": []interface{}{
							map[string]interface{}{
								"type":   "Ready",
								"status": "False",
								"reason": "UpdateFailedTerminalError",
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "resource with UpdateFailedTerminalError but generation progressed - should not skip",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
					"kind":       "APIGatewayAPI",
					"metadata": map[string]interface{}{
						"name":       "test-resource",
						"namespace":  "test-namespace",
						"generation": int64(3),
					},
					"status": map[string]interface{}{
						"observedGeneration": int64(2),
						"conditions": []interface{}{
							map[string]interface{}{
								"type":   "Ready",
								"status": "False",
								"reason": "UpdateFailedTerminalError",
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "resource with UpdateFailedTerminalError but being deleted - should not skip",
			resource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "apigateway.cnrm.cloud.google.com/v1beta1",
					"kind":       "APIGatewayAPI",
					"metadata": map[string]interface{}{
						"name":              "test-resource",
						"namespace":         "test-namespace",
						"generation":        int64(2),
						"deletionTimestamp": "2026-09-03T00:00:00Z",
						"finalizers": []interface{}{
							"cnrm.cloud.google.com/finalizer",
						},
					},
					"status": map[string]interface{}{
						"observedGeneration": int64(2),
						"conditions": []interface{}{
							map[string]interface{}{
								"type":   "Ready",
								"status": "False",
								"reason": "UpdateFailedTerminalError",
							},
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ShouldSkip(tc.resource)
			if got != tc.want {
				t.Errorf("ShouldSkip() = %v, want %v", got, tc.want)
			}
		})
	}
}
