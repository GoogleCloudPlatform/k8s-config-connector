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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// type: "networkservices.googleapis.com/EdgeCacheService"
// pattern: "projects/{project}/locations/global/edgeCacheServices/{edge_cache_service_id}"

var edgeCacheServiceURL = gcpurls.Template[NetworkServicesEdgeCacheServiceIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/global/edgeCacheServices/{edgeCacheServiceID}",
)

// NetworkServicesEdgeCacheServiceIdentity defines the resource reference to NetworkServicesEdgeCacheService, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type NetworkServicesEdgeCacheServiceIdentity struct {
	ProjectID          string
	EdgeCacheServiceID string
}

func (i *NetworkServicesEdgeCacheServiceIdentity) String() string {
	return edgeCacheServiceURL.ToString(*i)
}

func (i *NetworkServicesEdgeCacheServiceIdentity) ID() string {
	return i.EdgeCacheServiceID
}

func (i *NetworkServicesEdgeCacheServiceIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/global", i.ProjectID)
}

func (i *NetworkServicesEdgeCacheServiceIdentity) FromExternal(external string) error {
	out, match, err := edgeCacheServiceURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesEdgeCacheService external=%q was not known (use %s)", external, edgeCacheServiceURL.CanonicalForm())
	}
	*i = *out
	return nil
}

// Helper to construct Identity from components
func NewNetworkServicesEdgeCacheServiceIdentity(project, edgeCacheServiceID string) *NetworkServicesEdgeCacheServiceIdentity {
	return &NetworkServicesEdgeCacheServiceIdentity{
		ProjectID:          project,
		EdgeCacheServiceID: edgeCacheServiceID,
	}
}

// Common functions
func (i *NetworkServicesEdgeCacheServiceIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

var _ identity.Resource = &NetworkServicesEdgeCacheService{}

func (obj *NetworkServicesEdgeCacheService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &NetworkServicesEdgeCacheServiceIdentity{}

	// Resolve user-configured Project
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	id.ProjectID = projectID

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	id.EdgeCacheServiceID = resourceID

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkServicesEdgeCacheServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}
