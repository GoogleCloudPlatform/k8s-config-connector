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

package bigquerydataset

import "testing"

func TestParseDatasetFullID(t *testing.T) {
	tests := []struct {
		name        string
		fullID      string
		wantProject string
		wantDataset string
		wantOK      bool
	}{
		{
			name:        "plainProjectID",
			fullID:      "my-project:my_dataset",
			wantProject: "my-project",
			wantDataset: "my_dataset",
			wantOK:      true,
		},
		{
			// A universe-qualified project ID contains a colon, so splitting on
			// every colon yields three parts and drops the dataset.
			name:        "universeQualifiedProjectID",
			fullID:      "my-universe:my-project:my_dataset",
			wantProject: "my-universe:my-project",
			wantDataset: "my_dataset",
			wantOK:      true,
		},
		{
			// The same shape has existed in the public universe for years, via
			// domain-scoped projects.
			name:        "domainScopedProjectID",
			fullID:      "example.com:my-project:my_dataset",
			wantProject: "example.com:my-project",
			wantDataset: "my_dataset",
			wantOK:      true,
		},
		{
			name:   "noColon",
			fullID: "my_dataset",
			wantOK: false,
		},
		{
			name:   "empty",
			fullID: "",
			wantOK: false,
		},
		{
			name:   "emptyProject",
			fullID: ":my_dataset",
			wantOK: false,
		},
		{
			name:   "emptyDataset",
			fullID: "my-project:",
			wantOK: false,
		},
		{
			name:   "onlyColon",
			fullID: ":",
			wantOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			project, dataset, ok := parseDatasetFullID(tc.fullID)
			if ok != tc.wantOK {
				t.Fatalf("parseDatasetFullID(%q) ok = %v, want %v", tc.fullID, ok, tc.wantOK)
			}
			if !tc.wantOK {
				return
			}
			if project != tc.wantProject {
				t.Errorf("parseDatasetFullID(%q) project = %q, want %q", tc.fullID, project, tc.wantProject)
			}
			if dataset != tc.wantDataset {
				t.Errorf("parseDatasetFullID(%q) dataset = %q, want %q", tc.fullID, dataset, tc.wantDataset)
			}
		})
	}
}
