// Copyright 2024 Google LLC
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

package resourceoverrides_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestResourceManagerLienPreActuationTransform(t *testing.T) {
	tests := []struct {
		name           string
		initialParent  string
		expectedParent string
	}{
		{
			name:           "partition prefix stripped from full projects path",
			initialParent:  "projects/partition:sample-project",
			expectedParent: "projects/sample-project",
		},
		{
			name:           "partition prefix stripped and projects prefix added",
			initialParent:  "partition:sample-project",
			expectedParent: "projects/sample-project",
		},
		{
			name:           "raw project ID gets projects prefix",
			initialParent:  "sample-project",
			expectedParent: "projects/sample-project",
		},
		{
			name:           "numeric project number gets projects prefix",
			initialParent:  "123456789012",
			expectedParent: "projects/123456789012",
		},
		{
			name:           "already normalized projects path preserved",
			initialParent:  "projects/123456789012",
			expectedParent: "projects/123456789012",
		},
	}

	handler := resourceoverrides.NewResourceOverridesHandler()
	handler.Register(resourceoverrides.GetResourceManagerLienResourceOverrides())

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "resourcemanager.cnrm.cloud.google.com/v1beta1",
					Kind:       "ResourceManagerLien",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-lien",
					Namespace: "default",
				},
				Spec: map[string]interface{}{
					"origin": "test-origin",
					"reason": "test-reason",
					"parent": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": tc.initialParent,
						},
					},
				},
			}

			if err := handler.PreActuationTransform(res); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			actualParent, _, err := unstructured.NestedString(res.Spec, "parent", "projectRef", "external")
			if err != nil {
				t.Fatalf("error reading parent: %v", err)
			}
			if actualParent != tc.expectedParent {
				t.Fatalf("expected parent %q, got %q", tc.expectedParent, actualParent)
			}
		})
	}
}
