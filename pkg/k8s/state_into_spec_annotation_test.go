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

package k8s

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValidateOrDefaultStateIntoSpecAnnotation(t *testing.T) {
	tests := []struct {
		name         string
		obj          *unstructured.Unstructured
		defaultValue string
		hasError     bool
		expectedVal  string
	}{
		{
			name: "valid 'state-into-spec' value",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "absent",
						},
					},
				},
			},
			defaultValue: "merge",
			expectedVal:  "absent",
		},
		{
			name: "invalid 'state-into-spec' value",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "not_accepted_value",
						},
					},
				},
			},
			defaultValue: "merge",
			hasError:     true,
		},
		{
			name: "'state-into-spec' value is an empty string",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "",
						},
					},
				},
			},
			defaultValue: "merge",
			hasError:     true,
		},
		{
			name: "'state-into-spec' annotation defaulted to defaultValue if unset",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			defaultValue: "absent",
			expectedVal:  "absent",
		},
		{
			name: "BigQueryDataset kind can use 'absent' as the 'state-into-spec' value",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "bigquery.cnrm.cloud.google.com/v1beta1",
					"kind":       "BigQueryDataset",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "absent",
						},
					},
				},
			},
			defaultValue: "absent",
			expectedVal:  "absent",
		},
		{
			name: "BigQueryDataset kind can use 'merge' as the 'state-into-spec' value",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "bigquery.cnrm.cloud.google.com/v1beta1",
					"kind":       "BigQueryDataset",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "merge",
						},
					},
				},
			},
			defaultValue: "merge",
			expectedVal:  "merge",
		},
		{
			name: "error out if defaultValue is invalid",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			defaultValue: "invalid",
			hasError:     true,
		},
		{
			name: "gvk doesn't support 'merge' cannot use 'merge' as the 'state-into-spec' value",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1beta1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							StateIntoSpecAnnotation: "merge",
						},
					},
				},
			},
			defaultValue: "merge",
			hasError:     true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateOrDefaultStateIntoSpecAnnotation(tc.obj, tc.defaultValue)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actualVal, found := GetAnnotation(StateIntoSpecAnnotation, tc.obj)
			if !found {
				t.Fatalf("'%v' annotation is not found", StateIntoSpecAnnotation)
			}
			if actualVal != tc.expectedVal {
				t.Fatalf("got %v, want %v", actualVal, tc.expectedVal)
			}
		})
	}
}
