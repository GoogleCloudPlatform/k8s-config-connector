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

package configconnector

import (
	"context"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestGetKindFromCRD(t *testing.T) {
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "storagebuckets.storage.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"names": map[string]interface{}{
					"kind": "StorageBucket",
				},
			},
		},
	}
	obj, err := manifest.NewObject(u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	kind, err := getKindFromCRD(obj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if kind != "StorageBucket" {
		t.Errorf("expected StorageBucket, got %v", kind)
	}
}

func TestApplyResourceKindsFilter(t *testing.T) {
	r := &Reconciler{
		log: logr.Discard(),
	}

	u1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "storagebuckets.storage.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"names": map[string]interface{}{
					"kind": "StorageBucket",
				},
			},
		},
	}
	obj1, _ := manifest.NewObject(u1)

	u2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "pubsubtopics.pubsub.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"names": map[string]interface{}{
					"kind": "PubSubTopic",
				},
			},
		},
	}
	obj2, _ := manifest.NewObject(u2)

	crds := []*manifest.Object{obj1, obj2}

	tests := []struct {
		name          string
		resourceKinds *corev1beta1.ResourceKinds
		expectedKinds []string
	}{
		{
			name:          "No filter",
			resourceKinds: nil,
			expectedKinds: []string{"StorageBucket", "PubSubTopic"},
		},
		{
			name: "Allowlist only",
			resourceKinds: &corev1beta1.ResourceKinds{
				Allowlist: []string{"StorageBucket"},
			},
			expectedKinds: []string{"StorageBucket"},
		},
		{
			name: "Denylist only",
			resourceKinds: &corev1beta1.ResourceKinds{
				Denylist: []string{"StorageBucket"},
			},
			expectedKinds: []string{"PubSubTopic"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &corev1beta1.ConfigConnector{
				Spec: corev1beta1.ConfigConnectorSpec{
					ResourceKinds: tt.resourceKinds,
				},
			}
			m := &manifest.Objects{
				Items: crds,
			}

			filterFunc := r.applyResourceKindsFilter()
			err := filterFunc(context.Background(), cc, m)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(m.Items) != len(tt.expectedKinds) {
				t.Errorf("expected %v items, got %v", len(tt.expectedKinds), len(m.Items))
			}

			for i, expected := range tt.expectedKinds {
				kind, _ := getKindFromCRD(m.Items[i])
				if kind != expected {
					t.Errorf("expected kind %v at index %v, got %v", expected, i, kind)
				}
			}
		})
	}
}
