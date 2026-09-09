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

package common

import "testing"

func TestExtractProjectID(t *testing.T) {
	tests := []struct {
		name         string
		resourceName string
		want         string
	}{
		{
			name:         "empty",
			resourceName: "",
			want:         "",
		},
		{
			name:         "https URL with project ID",
			resourceName: "https://www.googleapis.com/compute/v1/projects/my-project/global/forwardingRules/my-rule",
			want:         "my-project",
		},
		{
			name:         "https URL with project number",
			resourceName: "https://compute.googleapis.com/compute/v1/projects/123456789/regions/us-central1/backendServices/my-backend",
			want:         "123456789",
		},
		{
			name:         "double slash URL",
			resourceName: "//compute.googleapis.com/projects/my-project/global/forwardingRules/my-rule",
			want:         "my-project",
		},
		{
			name:         "projects prefix",
			resourceName: "projects/my-project/locations/global/lbRouteExtensions/my-extension",
			want:         "my-project",
		},
		{
			name:         "just projects",
			resourceName: "projects/my-project",
			want:         "my-project",
		},
		{
			name:         "invalid format",
			resourceName: "foo/bar",
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractProjectID(tt.resourceName); got != tt.want {
				t.Errorf("ExtractProjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}
