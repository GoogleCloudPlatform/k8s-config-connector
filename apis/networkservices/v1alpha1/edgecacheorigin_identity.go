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

// type: "networkservices.googleapis.com/EdgeCacheOrigin"
// pattern: "projects/{project}/locations/global/edgeCacheOrigins/{edge_cache_origin_id}"

var edgeCacheOriginURL = gcpurls.Template[NetworkServicesEdgeCacheOriginIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/global/edgeCacheOrigins/{edgeCacheOriginID}",
)

// NetworkServicesEdgeCacheOriginIdentity defines the resource reference to NetworkServicesEdgeCacheOrigin, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type NetworkServicesEdgeCacheOriginIdentity struct {
	ProjectID         string
	EdgeCacheOriginID string
}

func (i *NetworkServicesEdgeCacheOriginIdentity) String() string {
	return edgeCacheOriginURL.ToString(*i)
}

func (i *NetworkServicesEdgeCacheOriginIdentity) ID() string {
	return i.EdgeCacheOriginID
}

func (i *NetworkServicesEdgeCacheOriginIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/global", i.ProjectID)
}

func (i *NetworkServicesEdgeCacheOriginIdentity) FromExternal(external string) error {
	out, match, err := edgeCacheOriginURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesEdgeCacheOrigin external=%q was not known (use %s)", external, edgeCacheOriginURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewNetworkServicesEdgeCacheOriginIdentity(project, edgeCacheOriginID string) *NetworkServicesEdgeCacheOriginIdentity {
	return &NetworkServicesEdgeCacheOriginIdentity{
		ProjectID:         project,
		EdgeCacheOriginID: edgeCacheOriginID,
	}
}

// Common functions
func (i *NetworkServicesEdgeCacheOriginIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}
