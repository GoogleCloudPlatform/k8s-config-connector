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

package ha

import (
	"context"
	"slices"
	"testing"
)

func TestGCPProvisionerDeduplicatesProjectsAndUsesMultiRegionBucketLocation(t *testing.T) {
	runner := &recordingCommandRunner{}
	provisioner := newGCPProvisioner(runner, "my-ha-bucket")
	clusters := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a", location: "us-central1-a", clusterName: "c1"},
		{contextName: "ctx-2", project: "proj-a", location: "us-east1-b", clusterName: "c2"},
	}

	err := provisioner.provision(context.Background(), clusters)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantCommands := []string{
		"gcloud storage buckets create gs://my-ha-bucket --project=proj-a --location=us",
		"gcloud storage buckets add-iam-policy-binding gs://my-ha-bucket --member=serviceAccount:kcc-mcl@proj-a.iam.gserviceaccount.com --role=roles/storage.objectAdmin",
		"gcloud storage buckets add-iam-policy-binding gs://my-ha-bucket --member=serviceAccount:kcc-mcl@proj-a.iam.gserviceaccount.com --role=roles/storage.legacyBucketReader",
		"gcloud iam service-accounts add-iam-policy-binding kcc-mcl@proj-a.iam.gserviceaccount.com --project=proj-a --role=roles/iam.workloadIdentityUser --member=serviceAccount:proj-a.svc.id.goog[cnrm-system/kcc-mcl]",
		"gcloud iam service-accounts add-iam-policy-binding kcc-syncer@proj-a.iam.gserviceaccount.com --project=proj-a --role=roles/iam.workloadIdentityUser --member=serviceAccount:proj-a.svc.id.goog[cnrm-system/kcc-syncer]",
		"gcloud projects add-iam-policy-binding proj-a --member=serviceAccount:kcc-syncer@proj-a.iam.gserviceaccount.com --role=roles/container.viewer",
	}
	if !slices.Equal(runner.commands, wantCommands) {
		t.Fatalf("got commands:\n%v\nwant:\n%v", runner.commands, wantCommands)
	}
}

func TestGCSBucketLocation(t *testing.T) {
	tests := []struct {
		name     string
		location string
		want     string
	}{
		{
			name:     "regional US location maps to us",
			location: "us-central1",
			want:     "us",
		},
		{
			name:     "zonal US location maps to us",
			location: "us-central1-a",
			want:     "us",
		},
		{
			name:     "regional Europe location maps to eu",
			location: "europe-west1",
			want:     "eu",
		},
		{
			name:     "zonal Europe location maps to eu",
			location: "europe-west3-b",
			want:     "eu",
		},
		{
			name:     "regional Asia location maps to asia",
			location: "asia-east1",
			want:     "asia",
		},
		{
			name:     "zonal Asia location maps to asia",
			location: "asia-northeast1-a",
			want:     "asia",
		},
		{
			name:     "preserves already multi-region location",
			location: "eu",
			want:     "eu",
		},
		{
			name:     "case insensitive matching",
			location: "US-EAST1",
			want:     "us",
		},
		{
			name:     "defaults to us for location without multi-region",
			location: "australia-southeast1",
			want:     "us",
		},
		{
			name:     "defaults to us for empty location",
			location: "",
			want:     "us",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			loc := tc.location

			got := gcsBucketLocation(loc)

			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestDedupeProjects(t *testing.T) {
	clusters := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a"},
		{contextName: "ctx-2", project: "proj-a"},
		{contextName: "ctx-3", project: "proj-b"},
	}

	got := dedupeProjects(clusters)

	want := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a"},
		{contextName: "ctx-3", project: "proj-b"},
	}
	if !slices.Equal(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}
