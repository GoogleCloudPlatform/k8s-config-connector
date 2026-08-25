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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DiscoveryEngineDataConnectorIdentity{}
	_ identity.Resource   = &DiscoveryEngineDataConnector{}
)

var DiscoveryEngineDataConnectorIdentityFormat = gcpurls.Template[DiscoveryEngineDataConnectorIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/dataConnector")

// DiscoveryEngineDataConnectorIdentity is the identity of a GCP DiscoveryEngineDataConnector resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineDataConnectorIdentity struct {
	Project    string
	Location   string
	Collection string
}

func (i *DiscoveryEngineDataConnectorIdentity) String() string {
	return DiscoveryEngineDataConnectorIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineDataConnectorIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineDataConnectorIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineDataConnector external=%q was not known (use %s): %w", ref, DiscoveryEngineDataConnectorIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineDataConnector external=%q was not known (use %s)", ref, DiscoveryEngineDataConnectorIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineDataConnectorIdentity) Host() string {
	return DiscoveryEngineDataConnectorIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineDataConnectorSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineDataConnector) (*DiscoveryEngineDataConnectorIdentity, error) {
	if obj.Spec.Location == nil || *obj.Spec.Location == "" {
		return nil, fmt.Errorf("spec.location is required")
	}
	location := *obj.Spec.Location

	if obj.Spec.CollectionID == nil || *obj.Spec.CollectionID == "" {
		return nil, fmt.Errorf("spec.collectionID is required")
	}
	collectionID := *obj.Spec.CollectionID

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DiscoveryEngineDataConnectorIdentity{
		Project:    projectID,
		Location:   location,
		Collection: collectionID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineDataConnector) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineDataConnectorSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineDataConnectorIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineDataConnector identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DiscoveryEngineDataConnector) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
