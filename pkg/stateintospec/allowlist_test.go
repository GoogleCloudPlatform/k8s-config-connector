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

package stateintospec

import (
	"testing"

	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestSupportsStateIntoSpecMerge(t *testing.T) {
	tests := []struct {
		name           string
		gvk            k8sschema.GroupVersionKind
		expectedResult bool
	}{
		{
			name: "ComputeInstance should support 'state-into-spec: merge'",
			gvk: k8sschema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeInstance",
			},
			expectedResult: true,
		},
		{
			name: "AccessContextManagerServicePerimeterResource should not " +
				"support 'state-into-spec: merge'",
			gvk: k8sschema.GroupVersionKind{
				Group:   "accesscontextmanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "AccessContextManagerServicePerimeterResource",
			},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualResult := SupportsStateIntoSpecMerge(tc.gvk)
			if actualResult != tc.expectedResult {
				t.Fatalf("got %v, want %v", actualResult, tc.expectedResult)
			}
		})
	}
}