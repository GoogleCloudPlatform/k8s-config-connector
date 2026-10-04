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

func TestUnderlyingResourceOutOfSyncPredicate(t *testing.T) {
	pred := UnderlyingResourceOutOfSyncPredicate{}

	now := metav1.NewTime(time.Now())

	tests := []struct {
		name     string
		oldObj   *unstructured.Unstructured
		newObj   *unstructured.Unstructured
		wantSync bool
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
			wantSync: false,
		},
		{
			name: "generation changed (spec update)",
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
			wantSync: true,
		},
		{
			name: "deletion timestamp set",
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
						"deletionTimestamp": now.Format(time.RFC3339),
					},
				},
			},
			wantSync: true,
		},
		{
			name: "terminal-error-mode annotation changed from builtin to none",
			oldObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.TerminalErrorModeAnnotation: "builtin",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"metadata": map[string]interface{}{
						"generation": int64(1),
						"annotations": map[string]interface{}{
							k8s.TerminalErrorModeAnnotation: "none",
						},
					},
				},
			},
			wantSync: true,
		},
		{
			name: "terminal-error-mode annotation added",
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
							k8s.TerminalErrorModeAnnotation: "builtin",
						},
					},
				},
			},
			wantSync: true,
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
							k8s.InternalForceReconcileAnnotation: "true",
						},
					},
				},
			},
			wantSync: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pred.Update(event.UpdateEvent{
				ObjectOld: tc.oldObj,
				ObjectNew: tc.newObj,
			})
			if got != tc.wantSync {
				t.Fatalf("pred.Update() = %v, want %v", got, tc.wantSync)
			}
		})
	}
}
