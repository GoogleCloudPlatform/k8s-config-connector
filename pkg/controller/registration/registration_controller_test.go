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
	"testing"
)

func TestMergeConcurrencyOverrides(t *testing.T) {
	testGK := "ComputeForwardingRule.compute.cnrm.cloud.google.com"
	otherGK := "BigQueryDataset.bigquery.cnrm.cloud.google.com"

	tests := []struct {
		name                    string
		ccConcurrencyOverrides  map[string]int
		cccConcurrencyOverrides map[string]int
		want                    int
	}{
		{
			name: "no overrides: returns 0 (defaulted by controller)",
			want: 0,
		},
		{
			name: "CC override only (cluster mode or global default): returns CC value",
			ccConcurrencyOverrides: map[string]int{
				testGK: 5,
			},
			want: 5,
		},
		{
			name: "CCC override only: returns CCC value",
			cccConcurrencyOverrides: map[string]int{
				testGK: 8,
			},
			want: 8,
		},
		{
			name: "CCC overrides another GK: falls back to global CC value for target GK",
			ccConcurrencyOverrides: map[string]int{
				testGK: 4,
			},
			cccConcurrencyOverrides: map[string]int{
				otherGK: 12,
			},
			want: 4,
		},
		{
			name: "both CCC and CC override same GK: CCC takes precedence",
			ccConcurrencyOverrides: map[string]int{
				testGK: 3,
			},
			cccConcurrencyOverrides: map[string]int{
				testGK: 15,
			},
			want: 15,
		},
		{
			name: "non-positive value in CCC falls back to CC override",
			ccConcurrencyOverrides: map[string]int{
				testGK: 7,
			},
			cccConcurrencyOverrides: map[string]int{
				testGK: 0,
			},
			want: 7,
		},
		{
			name: "non-positive values in both are ignored",
			ccConcurrencyOverrides: map[string]int{
				testGK: -1,
			},
			cccConcurrencyOverrides: map[string]int{
				testGK: 0,
			},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			merged := mergeConcurrencyOverrides(tc.ccConcurrencyOverrides, tc.cccConcurrencyOverrides)
			if got := merged[testGK]; got != tc.want {
				t.Errorf("mergeConcurrencyOverrides()[%q] = %d, want %d", testGK, got, tc.want)
			}
		})
	}
}
