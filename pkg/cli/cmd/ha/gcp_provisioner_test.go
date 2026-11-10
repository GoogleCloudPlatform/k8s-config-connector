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
	"bytes"
	"context"
	"slices"
	"strings"
	"testing"
)

func TestGCPProvisionerDeduplicatesProjectsAndUsesMultiRegionBucketLocation(t *testing.T) {
	runner := &recordingCommandRunner{}
	provisioner := newGCPProvisioner(runner, "my-ha-bucket", "", nil)
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

func TestGCPProvisionerUsesExplicitBucketLocation(t *testing.T) {
	runner := &recordingCommandRunner{}
	provisioner := newGCPProvisioner(runner, "my-ha-bucket", "nam4", nil)
	clusters := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a", location: "us-central1-a", clusterName: "c1"},
		{contextName: "ctx-2", project: "proj-b", location: "europe-west1-b", clusterName: "c2"},
	}

	err := provisioner.provision(context.Background(), clusters)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runner.commands) == 0 || !strings.Contains(runner.commands[0], "--location=nam4") {
		t.Fatalf("expected first command to use --location=nam4, got %v", runner.commands)
	}
}

func TestGCPProvisionerWarnsOnRegionalFallback(t *testing.T) {
	runner := &recordingCommandRunner{}
	var stderr bytes.Buffer
	provisioner := newGCPProvisioner(runner, "my-ha-bucket", "", &stderr)
	clusters := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a", location: "australia-southeast1-a", clusterName: "c1"},
		{contextName: "ctx-2", project: "proj-a", location: "australia-southeast1-b", clusterName: "c2"},
	}

	err := provisioner.provision(context.Background(), clusters)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(runner.commands) == 0 || !strings.Contains(runner.commands[0], "--location=australia-southeast1") {
		t.Fatalf("expected first command to use --location=australia-southeast1, got %v", runner.commands)
	}
	warnMsg := stderr.String()
	if !strings.Contains(warnMsg, "Warning:") || !strings.Contains(warnMsg, "does not map to a GCS multi-region") || !strings.Contains(warnMsg, "australia-southeast1") {
		t.Fatalf("expected warning on regional fallback, got: %q", warnMsg)
	}
}

func TestGCPProvisionerFailsOnConflictingClusterLocations(t *testing.T) {
	runner := &recordingCommandRunner{}
	provisioner := newGCPProvisioner(runner, "my-ha-bucket", "", nil)
	clusters := []gkeClusterInfo{
		{contextName: "ctx-1", project: "proj-a", location: "us-central1-a", clusterName: "c1"},
		{contextName: "ctx-2", project: "proj-b", location: "europe-west1-b", clusterName: "c2"},
	}

	err := provisioner.provision(context.Background(), clusters)

	if err == nil {
		t.Fatalf("expected error on differing cluster locations, got nil")
	}
	wantErrSubstrings := []string{
		"clusters have differing bucket locations",
		`cluster "ctx-1" (location "us-central1-a") maps to "us"`,
		`cluster "ctx-2" (location "europe-west1-b") maps to "eu"`,
		"--bucket-location",
	}
	for _, sub := range wantErrSubstrings {
		if !strings.Contains(err.Error(), sub) {
			t.Errorf("expected error to contain %q, got %q", sub, err.Error())
		}
	}
}

func TestResolveBucketLocation(t *testing.T) {
	tests := []struct {
		name             string
		clusters         []gkeClusterInfo
		explicitLocation string
		wantLocation     string
		wantWarning      string
		wantErrSubstr    string
	}{
		{
			name: "explicit location takes precedence",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "us-central1-a"},
				{contextName: "ctx-2", location: "europe-west1-b"},
			},
			explicitLocation: "nam4",
			wantLocation:     "nam4",
		},
		{
			name: "clusters in same multi-region agree on us",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "us-central1-a"},
				{contextName: "ctx-2", location: "us-east1-b"},
			},
			wantLocation: "us",
		},
		{
			name: "clusters in Europe EU multi-region agree on eu",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "europe-west1-a"},
				{contextName: "ctx-2", location: "europe-west3-b"},
			},
			wantLocation: "eu",
		},
		{
			name: "clusters in Asia multi-region agree on asia",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "asia-east1-a"},
				{contextName: "ctx-2", location: "asia-northeast1-b"},
			},
			wantLocation: "asia",
		},
		{
			name: "excluded Europe region falls back to regional and warns",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "europe-west2-a"},
				{contextName: "ctx-2", location: "europe-west2-b"},
			},
			wantLocation: "europe-west2",
			wantWarning:  "Warning: cluster location \"europe-west2-a\" does not map to a GCS multi-region; using regional bucket location \"europe-west2\"",
		},
		{
			name: "excluded Europe region Zürich falls back to regional and warns",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "europe-west6-a"},
				{contextName: "ctx-2", location: "europe-west6-b"},
			},
			wantLocation: "europe-west6",
			wantWarning:  "Warning: cluster location \"europe-west6-a\" does not map to a GCS multi-region; using regional bucket location \"europe-west6\"",
		},
		{
			name: "excluded Asia region Hong Kong falls back to regional and warns",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "asia-east2-a"},
				{contextName: "ctx-2", location: "asia-east2-b"},
			},
			wantLocation: "asia-east2",
			wantWarning:  "Warning: cluster location \"asia-east2-a\" does not map to a GCS multi-region; using regional bucket location \"asia-east2\"",
		},
		{
			name: "excluded Asia region Jakarta falls back to regional and warns",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "asia-southeast2-a"},
				{contextName: "ctx-2", location: "asia-southeast2-b"},
			},
			wantLocation: "asia-southeast2",
			wantWarning:  "Warning: cluster location \"asia-southeast2-a\" does not map to a GCS multi-region; using regional bucket location \"asia-southeast2\"",
		},
		{
			name: "non-multi-region cluster location falls back to regional and warns",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "australia-southeast1-a"},
				{contextName: "ctx-2", location: "australia-southeast1-b"},
			},
			wantLocation: "australia-southeast1",
			wantWarning:  "Warning: cluster location \"australia-southeast1-a\" does not map to a GCS multi-region; using regional bucket location \"australia-southeast1\"",
		},
		{
			name: "conflicting multi-regions error out",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "us-central1-a"},
				{contextName: "ctx-2", location: "europe-west1-b"},
			},
			wantErrSubstr: "clusters have differing bucket locations",
		},
		{
			name: "conflicting multi-region and non-multi-region error out",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "europe-west1-a"},
				{contextName: "ctx-2", location: "europe-west2-b"},
			},
			wantErrSubstr: "clusters have differing bucket locations",
		},
		{
			name: "conflicting regional locations error out",
			clusters: []gkeClusterInfo{
				{contextName: "ctx-1", location: "australia-southeast1-a"},
				{contextName: "ctx-2", location: "australia-southeast2-b"},
			},
			wantErrSubstr: "clusters have differing bucket locations",
		},
		{
			name:          "empty clusters without explicit location errors out",
			clusters:      nil,
			wantErrSubstr: "no clusters provided to determine lease bucket location",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var stderr bytes.Buffer
			gotLoc, err := resolveBucketLocation(tc.clusters, tc.explicitLocation, &stderr)

			if tc.wantErrSubstr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErrSubstr) {
					t.Fatalf("expected error containing %q, got: %v", tc.wantErrSubstr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotLoc != tc.wantLocation {
				t.Errorf("got location %q, want %q", gotLoc, tc.wantLocation)
			}
			if tc.wantWarning != "" {
				if !strings.Contains(stderr.String(), tc.wantWarning) {
					t.Errorf("expected warning containing %q, got: %q", tc.wantWarning, stderr.String())
				}
			} else if stderr.Len() > 0 {
				t.Errorf("expected no warning, got: %q", stderr.String())
			}
		})
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
			name:     "regional Europe London is excluded from EU multi-region",
			location: "europe-west2",
			want:     "europe-west2",
		},
		{
			name:     "zonal Europe London is excluded from EU multi-region",
			location: "europe-west2-a",
			want:     "europe-west2",
		},
		{
			name:     "regional Europe Zürich is excluded from EU multi-region",
			location: "europe-west6",
			want:     "europe-west6",
		},
		{
			name:     "zonal Europe Zürich is excluded from EU multi-region",
			location: "europe-west6-c",
			want:     "europe-west6",
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
			name:     "regional Asia Hong Kong is excluded from ASIA multi-region",
			location: "asia-east2",
			want:     "asia-east2",
		},
		{
			name:     "zonal Asia Hong Kong is excluded from ASIA multi-region",
			location: "asia-east2-b",
			want:     "asia-east2",
		},
		{
			name:     "regional Asia Jakarta is excluded from ASIA multi-region",
			location: "asia-southeast2",
			want:     "asia-southeast2",
		},
		{
			name:     "zonal Asia Jakarta is excluded from ASIA multi-region",
			location: "asia-southeast2-c",
			want:     "asia-southeast2",
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
			name:     "falls back to cluster region for location without multi-region",
			location: "australia-southeast1",
			want:     "australia-southeast1",
		},
		{
			name:     "zonal location without multi-region strips zone",
			location: "australia-southeast1-a",
			want:     "australia-southeast1",
		},
		{
			name:     "Canada region falls back to regional",
			location: "northamerica-northeast1-a",
			want:     "northamerica-northeast1",
		},
		{
			name:     "empty location returns empty",
			location: "",
			want:     "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := gcsBucketLocation(tc.location)

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
