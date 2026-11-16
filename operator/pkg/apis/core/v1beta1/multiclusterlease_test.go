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

package v1beta1

import (
	"strings"
	"testing"
)

func TestGKEClusterResourceName(t *testing.T) {
	cluster := &GKECluster{Project: "my-project", Location: "us-central1", Name: "cluster-a"}
	if got, want := cluster.ResourceName(), "projects/my-project/locations/us-central1/clusters/cluster-a"; got != want {
		t.Errorf("ResourceName() = %q, want %q", got, want)
	}
}

func TestParseGKEClusterResourceName(t *testing.T) {
	got, err := ParseGKEClusterResourceName("projects/my-project/locations/us-central1-a/clusters/cluster-a")
	if err != nil {
		t.Fatalf("ParseGKEClusterResourceName: %v", err)
	}
	if want := (GKECluster{Project: "my-project", Location: "us-central1-a", Name: "cluster-a"}); *got != want {
		t.Errorf("ParseGKEClusterResourceName = %+v, want %+v", *got, want)
	}

	for _, invalid := range []string{
		"",
		"cluster-a",
		"projects/my-project/zones/us-central1-a/clusters/cluster-a",
		"projects//locations/us-central1/clusters/cluster-a",
		"projects/my-project/locations/us-central1/clusters/",
		"projects/my-project/locations/us-central1/clusters/cluster-a/nodePools/pool",
		"//container.googleapis.com/projects/my-project/locations/us-central1/clusters/cluster-a",
	} {
		if got, err := ParseGKEClusterResourceName(invalid); err == nil {
			t.Errorf("ParseGKEClusterResourceName(%q) = %+v, want an error", invalid, *got)
		}
	}
}

func TestCandidateIdentity(t *testing.T) {
	for _, tc := range []struct {
		name    string
		spec    MultiClusterLeaseSpec
		want    string
		wantErr string
	}{
		{
			name: "identity is the resource name of gkeCluster",
			spec: MultiClusterLeaseSpec{GKECluster: &GKECluster{Project: "p", Location: "us-central1", Name: "c"}},
			want: "projects/p/locations/us-central1/clusters/c",
		},
		{
			name:    "deprecated clusterCandidateIdentity is rejected",
			spec:    MultiClusterLeaseSpec{ClusterCandidateIdentity: "cluster-a", GKECluster: &GKECluster{Project: "p", Location: "us-central1", Name: "c"}},
			wantErr: "clusterCandidateIdentity is no longer supported",
		},
		{
			name:    "gkeCluster is required",
			spec:    MultiClusterLeaseSpec{},
			wantErr: "gkeCluster is required",
		},
		{
			name:    "gkeCluster parts can't be empty",
			spec:    MultiClusterLeaseSpec{GKECluster: &GKECluster{Project: "p", Name: "c"}},
			wantErr: "gkeCluster is invalid",
		},
		{
			name:    "gkeCluster parts can't contain a slash",
			spec:    MultiClusterLeaseSpec{GKECluster: &GKECluster{Project: "p", Location: "us-central1", Name: "a/b"}},
			wantErr: "gkeCluster is invalid",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.spec.CandidateIdentity()
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("CandidateIdentity() error = %v, want an error containing %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("CandidateIdentity(): %v", err)
			}
			if got != tc.want {
				t.Errorf("CandidateIdentity() = %q, want %q", got, tc.want)
			}
		})
	}
}
