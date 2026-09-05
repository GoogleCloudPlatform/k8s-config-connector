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

package predicate_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func TestUnderlyingResourceOutOfSyncPredicate_Update(t *testing.T) {
	tests := []struct {
		name           string
		oldAnnotations map[string]string
		newAnnotations map[string]string
		expected       bool
	}{
		{
			name: "no change in annotations or spec",
			oldAnnotations: map[string]string{
				"foo": "bar",
			},
			newAnnotations: map[string]string{
				"foo": "bar",
			},
			expected: false,
		},
		{
			name: "unrelated annotation change",
			oldAnnotations: map[string]string{
				"foo": "bar",
			},
			newAnnotations: map[string]string{
				"foo": "baz",
			},
			expected: false,
		},
		{
			name: "actuation mode annotation added",
			oldAnnotations: map[string]string{
				"foo": "bar",
			},
			newAnnotations: map[string]string{
				"foo":                       "bar",
				k8s.ActuationModeAnnotation: "Paused",
			},
			expected: true,
		},
		{
			name: "actuation mode annotation changed",
			oldAnnotations: map[string]string{
				k8s.ActuationModeAnnotation: "Paused",
			},
			newAnnotations: map[string]string{
				k8s.ActuationModeAnnotation: "Reconciling",
			},
			expected: true,
		},
		{
			name: "actuation mode annotation removed",
			oldAnnotations: map[string]string{
				k8s.ActuationModeAnnotation: "Paused",
			},
			newAnnotations: map[string]string{},
			expected:       true,
		},
	}

	p := predicate.UnderlyingResourceOutOfSyncPredicate{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			oldObj := &unstructured.Unstructured{}
			oldObj.SetAnnotations(tc.oldAnnotations)
			oldObj.SetGeneration(1)

			newObj := &unstructured.Unstructured{}
			newObj.SetAnnotations(tc.newAnnotations)
			newObj.SetGeneration(1)

			e := event.UpdateEvent{
				ObjectOld: oldObj,
				ObjectNew: newObj,
			}

			actual := p.Update(e)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
