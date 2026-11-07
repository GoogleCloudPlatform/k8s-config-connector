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

package gcp

import "testing"

func TestFormatServiceAccountEmail(t *testing.T) {
	tests := []struct {
		name      string
		projectID string
		saName    string
		want      string
	}{
		{
			name:      "standard commercial project",
			projectID: "my-commercial-project",
			saName:    "app-runner",
			want:      "app-runner@my-commercial-project.iam.gserviceaccount.com",
		},
		{
			name:      "partitioned sovereign project (eu0)",
			projectID: "eu0:sample-project",
			saName:    "sovereign-app-sa",
			want:      "sovereign-app-sa@sample-project.eu0.iam.gserviceaccount.com",
		},
		{
			name:      "partitioned sovereign project (fr0)",
			projectID: "fr0:project-paris",
			saName:    "worker-sa",
			want:      "worker-sa@project-paris.fr0.iam.gserviceaccount.com",
		},
		{
			name:      "custom air-gapped partition",
			projectID: "sec1:airgap-corp",
			saName:    "reconciler",
			want:      "reconciler@airgap-corp.sec1.iam.gserviceaccount.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatServiceAccountEmail(tc.projectID, tc.saName); got != tc.want {
				t.Errorf("FormatServiceAccountEmail(%q, %q) = %q, want %q", tc.projectID, tc.saName, got, tc.want)
			}
		})
	}
}

func TestFormatWorkloadPool(t *testing.T) {
	tests := []struct {
		name      string
		projectID string
		want      string
	}{
		{
			name:      "standard commercial project",
			projectID: "my-commercial-project",
			want:      "my-commercial-project.svc.id.goog",
		},
		{
			name:      "partitioned sovereign project",
			projectID: "partition:sample-project",
			want:      "sample-project.partition.svc.id.goog",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatWorkloadPool(tc.projectID); got != tc.want {
				t.Errorf("FormatWorkloadPool(%q) = %q, want %q", tc.projectID, got, tc.want)
			}
		})
	}
}

func TestNormalizePartitionedDomains(t *testing.T) {
	input := `{"member":"serviceAccount:sovereign-app-sa@eu0:sample-project.iam.gserviceaccount.com","workloadPool":"eu0:sample-project.svc.id.goog"}`
	want := `{"member":"serviceAccount:sovereign-app-sa@sample-project.eu0.iam.gserviceaccount.com","workloadPool":"sample-project.eu0.svc.id.goog"}`
	if got := NormalizePartitionedDomains(input); got != want {
		t.Errorf("NormalizePartitionedDomains() = %s, want %s", got, want)
	}
}
