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
	"fmt"
	"strings"
)

// ResourceName returns the GKE resource name of the cluster:
// projects/<project>/locations/<location>/clusters/<name>.
func (c *GKECluster) ResourceName() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", c.Project, c.Location, c.Name)
}

// ParseGKEClusterResourceName parses a GKE resource name of the form
// projects/<project>/locations/<location>/clusters/<name>.
func ParseGKEClusterResourceName(resourceName string) (*GKECluster, error) {
	tokens := strings.Split(resourceName, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "clusters" ||
		tokens[1] == "" || tokens[3] == "" || tokens[5] == "" {
		return nil, fmt.Errorf("%q is not a GKE cluster resource name of the form projects/<project>/locations/<location>/clusters/<name>", resourceName)
	}
	return &GKECluster{Project: tokens[1], Location: tokens[3], Name: tokens[5]}, nil
}

// CandidateIdentity returns this cluster's identity in the multi-cluster election: the
// resource name of gkeCluster. It returns an error if gkeCluster is missing or invalid, or
// if the deprecated clusterCandidateIdentity is set.
func (s *MultiClusterLeaseSpec) CandidateIdentity() (string, error) {
	if s.ClusterCandidateIdentity != "" {
		return "", fmt.Errorf("spec.experiments.multiClusterLease.clusterCandidateIdentity is no longer supported; set gkeCluster to the GKE cluster that this Config Connector runs in")
	}
	if s.GKECluster == nil {
		return "", fmt.Errorf("spec.experiments.multiClusterLease.gkeCluster is required")
	}
	identity := s.GKECluster.ResourceName()
	// The CRD rejects empty parts and parts with a "/", but check anyway so that the identity
	// always parses back to the same cluster.
	if parsed, err := ParseGKEClusterResourceName(identity); err != nil || *parsed != *s.GKECluster {
		return "", fmt.Errorf("spec.experiments.multiClusterLease.gkeCluster is invalid: project, location and name must be non-empty and must not contain \"/\"")
	}
	return identity, nil
}
