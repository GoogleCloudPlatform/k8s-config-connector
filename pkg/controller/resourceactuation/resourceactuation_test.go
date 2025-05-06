// Copyright 2023 Google LLC
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

package resourceactuation_test

import (
	"strconv"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceactuation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestShouldSkip(t *testing.T) {
	testcases := []struct {
		name               string
		generation         int64
		observedGeneration int64
		reconcileInterval  int32
		status             string
		reason             string
		wantSkip           bool
	}{
		{
			name:               "Should skip the resource actuation",
			generation:         1,
			observedGeneration: 1,
			reconcileInterval:  0,
			status:             "True",
			reason:             k8s.UpToDate,
			wantSkip:           true,
		},
		{
			name:               "Generation and ObservedGeneration don't match",
			generation:         2,
			observedGeneration: 1,
			reconcileInterval:  0,
			status:             "True",
			reason:             k8s.UpToDate,
			wantSkip:           false,
		},
		{
			name:               "Reconcile interval is not set to 0",
			generation:         1,
			observedGeneration: 1,
			reconcileInterval:  600,
			status:             "True",
			reason:             k8s.UpToDate,
			wantSkip:           false,
		},
		{
			name:               "Resource status is not true",
			generation:         1,
			observedGeneration: 1,
			reconcileInterval:  0,
			status:             "False",
			reason:             k8s.UpToDate,
			wantSkip:           false,
		},
		{
			name:               "Resource reason is not UpToDate",
			generation:         1,
			observedGeneration: 1,
			reconcileInterval:  0,
			status:             "True",
			reason:             k8s.Updating,
			wantSkip:           false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var u unstructured.Unstructured
			u.SetGeneration(tc.generation)
			if err := unstructured.SetNestedField(u.Object, tc.observedGeneration, "status", "observedGeneration"); err != nil {
				t.Errorf("Unable to set status.observedGeneration: %v", err)
			}
			condition := map[string]interface{}{
				"status": tc.status,
				"reason": tc.reason,
			}
			if err := unstructured.SetNestedSlice(u.Object, []interface{}{condition}, "status", "conditions"); err != nil {
				t.Errorf("Unable to set status.conditions: %v", err)
			}
			u.SetAnnotations(map[string]string{k8s.ReconcileIntervalInSecondsAnnotation: strconv.Itoa(int(tc.reconcileInterval))})
			skip, _ := resourceactuation.ShouldSkip(&u)
			if skip != tc.wantSkip {
				t.Errorf("resourceactuation.ShouldSkip returns %t, want %t", skip, tc.wantSkip)
			}
		})
	}
}
