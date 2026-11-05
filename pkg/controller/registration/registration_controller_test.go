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

package registration

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestMaxConcurrentReconcilesOverrides(t *testing.T) {
	testGVK := schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeForwardingRule",
	}
	testGK := testGVK.GroupKind().String()
	otherGK := "BigQueryDataset.bigquery.cnrm.cloud.google.com"

	tests := []struct {
		name         string
		ccOverrides  map[string]int
		cccOverrides map[string]int
		want         int
	}{
		{
			name: "no overrides: returns default ControllerMaxConcurrentReconciles (20)",
			want: k8s.ControllerMaxConcurrentReconciles,
		},
		{
			name: "CC override only (cluster mode or global default): returns CC value",
			ccOverrides: map[string]int{
				testGK: 5,
			},
			want: 5,
		},
		{
			name: "CCC override only: returns CCC value",
			cccOverrides: map[string]int{
				testGK: 8,
			},
			want: 8,
		},
		{
			name: "CCC overrides another GK: falls back to global CC value for target GK",
			ccOverrides: map[string]int{
				testGK: 4,
			},
			cccOverrides: map[string]int{
				otherGK: 12,
			},
			want: 4,
		},
		{
			name: "both CCC and CC override same GK: CCC takes precedence",
			ccOverrides: map[string]int{
				testGK: 3,
			},
			cccOverrides: map[string]int{
				testGK: 15,
			},
			want: 15,
		},
		{
			name: "non-positive value in CCC falls back to CC override",
			ccOverrides: map[string]int{
				testGK: 7,
			},
			cccOverrides: map[string]int{
				testGK: 0,
			},
			want: 7,
		},
		{
			name: "non-positive values in both fall back to default (20)",
			ccOverrides: map[string]int{
				testGK: -1,
			},
			cccOverrides: map[string]int{
				testGK: 0,
			},
			want: k8s.ControllerMaxConcurrentReconciles,
		},
		{
			name: "unrecognized GroupKind is ignored",
			ccOverrides: map[string]int{
				"UnknownKind.unknown.cnrm.cloud.google.com": 50,
			},
			want: k8s.ControllerMaxConcurrentReconciles,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			opts := &defaultControllerOptions{
				maxConcurrentReconcilesOverrides: mergeMaxConcurrentReconcilesOverrides(ctx, tc.ccOverrides, tc.cccOverrides),
			}
			if got := opts.maxConcurrentReconcilesForGVK(testGVK); got != tc.want {
				t.Errorf("maxConcurrentReconcilesForGVK(%v) = %d, want %d", testGVK, got, tc.want)
			}
			if _, exists := opts.maxConcurrentReconcilesOverrides[schema.ParseGroupKind("UnknownKind.unknown.cnrm.cloud.google.com")]; exists {
				t.Errorf("expected unrecognized GroupKind to be ignored")
			}
		})
	}
}
