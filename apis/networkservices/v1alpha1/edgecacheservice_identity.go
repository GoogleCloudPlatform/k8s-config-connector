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

var edgeCacheServiceURL = gcpurls.Template[EdgeCacheServiceIdentity](
	"networkservices.googleapis.com",
	"projects/{projectID}/locations/global/edgeCacheServices/{edgeCacheServiceID}",
)

// EdgeCacheServiceIdentity represents the identity of a NetworkServicesEdgeCacheService.
// +k8s:deepcopy-gen=false
type EdgeCacheServiceIdentity struct {
	ProjectID          string
	EdgeCacheServiceID string
}

func (i *EdgeCacheServiceIdentity) String() string {
	return edgeCacheServiceURL.ToString(*i)
}

func (i *EdgeCacheServiceIdentity) ID() string {
	return i.EdgeCacheServiceID
}

func (i *EdgeCacheServiceIdentity) Parent() string {
	return fmt.Sprintf("projects/%s", i.ProjectID)
}

func (i *EdgeCacheServiceIdentity) FromExternal(external string) error {
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

var _ identity.Resource = &NetworkServicesEdgeCacheService{}

// GetIdentity builds an EdgeCacheServiceIdentity from the Config Connector EdgeCacheService object.
func (obj *NetworkServicesEdgeCacheService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	id := &EdgeCacheServiceIdentity{}

	// Resolve user-configured Project
	pRef := &refsv1beta1.ProjectRef{
		External:  obj.Spec.ProjectRef.External,
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
	}
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), pRef)
	if err != nil {
		return nil, err
	}
	id.ProjectID = projectRef.ProjectID

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
		statusIdentity := &EdgeCacheServiceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != id.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, id.String())
		}
	}
	return id, nil
}

func ParseEdgeCacheServiceExternal(external string) (projectID string, resourceID string, err error) {
	id := &EdgeCacheServiceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return "", "", err
	}
	return id.ProjectID, id.EdgeCacheServiceID, nil
}
