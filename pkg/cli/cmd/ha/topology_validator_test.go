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
	"testing"
)

func TestTopologyValidatorValidateClusterTopology(t *testing.T) {
	reports := []clusterStatusReport{
		{contextName: "gke_proj_us-central1_cluster-a", candidateIdentity: "cluster-a", leaseNamespace: "cnrm-system", leaseName: "lease-a"},
		{contextName: "gke_proj_us-east1_cluster-b", candidateIdentity: "cluster-b", leaseNamespace: "cnrm-system", leaseName: "lease-b"},
	}
	validStates := []*clusterHAState{
		{contextName: "gke_proj_us-central1_cluster-a", configConnector: newConfigConnectorWithMCL("lease-a", "cnrm-system", "cluster-a")},
		{contextName: "gke_proj_us-east1_cluster-b", configConnector: newConfigConnectorWithMCL("lease-b", "cnrm-system", "cluster-b")},
	}
	validator := newTopologyValidator(reports, validStates)

	warnings := validator.validateClusterTopology()

	wantWarnings := []string{
		`mismatched leaseName across clusters (gke_proj_us-central1_cluster-a="lease-a", gke_proj_us-east1_cluster-b="lease-b"): set spec.experiments.multiClusterLease.leaseName to the same value on every cluster because the GCS lock object is leases/<namespace>/<leaseName>`,
	}
	if !slices.Equal(warnings, wantWarnings) {
		t.Fatalf("got warnings %v, want %v", warnings, wantWarnings)
	}
}

func TestTopologyValidatorCheckLeaderConsistency(t *testing.T) {
	reports := []clusterStatusReport{
		{contextName: "gke_proj_us-central1_cluster-a", candidateIdentity: "cluster-a", role: "LEADER", observedLeader: "cluster-a"},
		{contextName: "gke_proj_us-east1_cluster-b", candidateIdentity: "cluster-b", role: "LEADER", observedLeader: "cluster-b"},
	}
	validator := newTopologyValidator(reports, nil)

	warnings := validator.checkLeaderConsistency(nil)

	wantWarnings := []string{
		`clusters disagree on global leader (gke_proj_us-central1_cluster-a="cluster-a", gke_proj_us-east1_cluster-b="cluster-b")`,
		`multiple clusters report role LEADER simultaneously: [gke_proj_us-central1_cluster-a (cluster-a), gke_proj_us-east1_cluster-b (cluster-b)]`,
	}
	if !slices.Equal(warnings, wantWarnings) {
		t.Fatalf("got warnings %v, want %v", warnings, wantWarnings)
	}
}
