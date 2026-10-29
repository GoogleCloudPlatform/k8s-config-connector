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
	"slices"
	"strings"
	"testing"
)

func TestParseGKEContext(t *testing.T) {
	raw := "gke_my-project_us-central1_cluster-a"

	got, err := parseGKEContext(raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := gkeClusterInfo{
		contextName: "gke_my-project_us-central1_cluster-a",
		identity:    "gke-my-project-us-central1-cluster-a",
		project:     "my-project",
		location:    "us-central1",
		clusterName: "cluster-a",
	}
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestParseGKEContextErrors(t *testing.T) {
	tests := []struct {
		name              string
		contextName       string
		wantErrSubstrings []string
	}{
		{
			name:              "non-GKE context name without gke prefix",
			contextName:       "cluster-a",
			wantErrSubstrings: []string{"does not match expected GKE context format"},
		},
		{
			name:              "missing cluster segment",
			contextName:       "gke_my-project_us-central1_",
			wantErrSubstrings: []string{"does not match expected GKE context format"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctxName := tc.contextName

			_, err := parseGKEContext(ctxName)

			if err == nil {
				t.Fatalf("expected error containing %v, got nil", tc.wantErrSubstrings)
			}
			for _, sub := range tc.wantErrSubstrings {
				if !strings.Contains(err.Error(), sub) {
					t.Fatalf("expected error %q to contain %q", err.Error(), sub)
				}
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
			name:     "regional GKE location is unchanged",
			location: "us-central1",
			want:     "us-central1",
		},
		{
			name:     "zonal GKE location strips zone suffix",
			location: "us-central1-a",
			want:     "us-central1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			info := gkeClusterInfo{location: tc.location}

			got := info.gcsBucketLocation()

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
