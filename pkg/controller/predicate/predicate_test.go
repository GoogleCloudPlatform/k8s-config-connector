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

package predicate

import (
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func TestUnderlyingResourceOutOfSyncPredicate_Update(t *testing.T) {
	pred := UnderlyingResourceOutOfSyncPredicate{}

	tests := []struct {
		name     string
		oldObj   *unstructured.Unstructured
		newObj   *unstructured.Unstructured
		expected bool
	}{
		{
			name: "no changes",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			expected: false,
		},
		{
			name: "generation changed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(2),
					},
				},
			},
			expected: true,
		},
		{
			name: "actuation mode annotation added",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "actuation mode annotation changed from Paused to Reconciling",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Reconciling",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "actuation mode annotation removed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			expected: true,
		},
		{
			name: "unrelated annotation changed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							"example.com/some-annotation": "val1",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							"example.com/some-annotation": "val2",
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "deletion timestamp added",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation":        int64(1),
						"deletionTimestamp": metav1.NewTime(time.Now()).Format(time.RFC3339),
					},
				},
			},
			expected: true,
		},
		{
			name: "deletion defender finalizer removed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"finalizers": []interface{}{k8s.DeletionDefenderFinalizerName},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"finalizers": []interface{}{},
					},
				},
			},
			expected: true,
		},
		{
			name: "labels changed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"labels": map[string]interface{}{
							"env": "dev",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"labels": map[string]interface{}{
							"env": "prod",
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "force reconcile annotation changed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.InternalForceReconcileAnnotation: "1",
						},
					},
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			evt := event.UpdateEvent{
				ObjectOld: tc.oldObj,
				ObjectNew: tc.newObj,
			}
			got := pred.Update(evt)
			if got != tc.expected {
				t.Errorf("predicate.Update() = %v, want %v", got, tc.expected)
			}
		})
	}
}
