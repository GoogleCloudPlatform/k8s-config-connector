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

package e2e

import "testing"

func TestIsListOrOperationResponse(t *testing.T) {
	tests := []struct {
		name string
		body map[string]any
		want bool
	}{
		{
			name: "nil body",
			body: nil,
			want: true,
		},
		{
			name: "operations list",
			body: map[string]any{
				"operations": []any{},
			},
			want: true,
		},
		{
			name: "compute list with kind suffix List",
			body: map[string]any{
				"kind":  "compute#networkList",
				"items": []any{},
			},
			want: true,
		},
		{
			name: "list response with items and no top-level name",
			body: map[string]any{
				"items": []any{
					map[string]any{"name": "item-1"},
				},
			},
			want: true,
		},
		{
			name: "single resource with items field and top-level name (e.g. NetworkSecurityAddressGroup)",
			body: map[string]any{
				"name":     "projects/test-proj/locations/global/addressGroups/ag-1",
				"capacity": 100,
				"items":    []any{"10.0.0.0/24"},
			},
			want: false,
		},
		{
			name: "standard single resource without items",
			body: map[string]any{
				"name": "projects/test-proj/locations/us-central1/clusters/c-1",
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isListOrOperationResponse(tc.body); got != tc.want {
				t.Errorf("isListOrOperationResponse(%v) = %v, want %v", tc.body, got, tc.want)
			}
		})
	}
}
