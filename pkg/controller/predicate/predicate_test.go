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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func TestUnderlyingResourceOutOfSyncPredicate_Update(t *testing.T) {
	p := UnderlyingResourceOutOfSyncPredicate{}

	tests := []struct {
		name      string
		oldObj    *unstructured.Unstructured
		newObj    *unstructured.Unstructured
		wantEvent bool
	}{
		{
			name: "actuation-mode annotation changed from Paused to Reconciling",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Reconciling",
						},
					},
				},
			},
			wantEvent: true,
		},
		{
			name: "actuation-mode annotation changed from Reconciling to Paused",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Reconciling",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			wantEvent: true,
		},
		{
			name: "actuation-mode annotation added",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			wantEvent: true,
		},
		{
			name: "actuation-mode annotation removed",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{},
				},
			},
			wantEvent: true,
		},
		{
			name: "actuation-mode annotation unchanged",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ActuationModeAnnotation: "Paused",
						},
					},
				},
			},
			wantEvent: false,
		},
		{
			name: "deletion timestamp added",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"deletionTimestamp": metav1.Now().Format(metav1.RFC3339Micro),
					},
				},
			},
			wantEvent: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := event.UpdateEvent{
				ObjectOld: tc.oldObj,
				ObjectNew: tc.newObj,
			}
			got := p.Update(e)
			if got != tc.wantEvent {
				t.Errorf("Update() = %v, want %v", got, tc.wantEvent)
			}
		})
	}
}
