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

package monitoring

import (
	"testing"
)

func TestFormatMonitoringAlertPolicyURL(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "standard commercial project create URL",
			in:   "https://monitoring.googleapis.com/v3/projects/my-project/alertPolicies",
			want: "https://monitoring.googleapis.com/v3/projects/my-project/alertPolicies",
		},
		{
			name: "sovereign colon-partitioned project create URL",
			in:   "https://monitoring.custom.universe.goog/v3/projects/partition:sample-project/alertPolicies",
			want: "https://monitoring.custom.universe.goog/v3/projects/partition%3Asample-project/alertPolicies",
		},
		{
			name: "sovereign colon-partitioned project get/update/delete URL",
			in:   "https://monitoring.custom.universe.goog/v3/projects/partition:sample-project/alertPolicies/12345",
			want: "https://monitoring.custom.universe.goog/v3/projects/partition%3Asample-project/alertPolicies/12345",
		},
		{
			name: "already escaped project URL",
			in:   "https://monitoring.custom.universe.goog/v3/projects/partition%3Asample-project/alertPolicies",
			want: "https://monitoring.custom.universe.goog/v3/projects/partition%3Asample-project/alertPolicies",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatMonitoringAlertPolicyURL(tt.in)
			if got != tt.want {
				t.Errorf("formatMonitoringAlertPolicyURL(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
