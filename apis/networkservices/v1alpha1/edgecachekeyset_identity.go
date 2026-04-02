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

package v1alpha1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

// type: "networkservices.googleapis.com/EdgeCacheKeyset"
// pattern: "projects/{project}/locations/global/edgeCacheKeysets/{edge_cache_keyset_id}"

var edgeCacheKeysetURL = gcpurls.Template[NetworkServicesEdgeCacheKeysetIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/global/edgeCacheKeysets/{edgeCacheKeysetID}",
)

// NetworkServicesEdgeCacheKeysetIdentity defines the resource reference to NetworkServicesEdgeCacheKeyset, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type NetworkServicesEdgeCacheKeysetIdentity struct {
	ProjectID         string
	EdgeCacheKeysetID string
}

func (i *NetworkServicesEdgeCacheKeysetIdentity) String() string {
	return edgeCacheKeysetURL.ToString(*i)
}

func (i *NetworkServicesEdgeCacheKeysetIdentity) ID() string {
	return i.EdgeCacheKeysetID
}

func (i *NetworkServicesEdgeCacheKeysetIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/global", i.ProjectID)
}

func (i *NetworkServicesEdgeCacheKeysetIdentity) FromExternal(external string) error {
	out, match, err := edgeCacheKeysetURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesEdgeCacheKeyset external=%q was not known (use %s)", external, edgeCacheKeysetURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewNetworkServicesEdgeCacheKeysetIdentity(project, edgeCacheKeysetID string) *NetworkServicesEdgeCacheKeysetIdentity {
	return &NetworkServicesEdgeCacheKeysetIdentity{
		ProjectID:         project,
		EdgeCacheKeysetID: edgeCacheKeysetID,
	}
}

// Common functions
func (i *NetworkServicesEdgeCacheKeysetIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}
