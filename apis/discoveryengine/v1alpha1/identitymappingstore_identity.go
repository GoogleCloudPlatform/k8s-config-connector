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
	_ identity.IdentityV2 = &DiscoveryEngineIdentityMappingStoreIdentity{}
	_ identity.Resource   = &DiscoveryEngineIdentityMappingStore{}
)

var DiscoveryEngineIdentityMappingStoreIdentityFormat = gcpurls.Template[DiscoveryEngineIdentityMappingStoreIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/identityMappingStores/{identitymappingstore}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineIdentityMappingStoreIdentity struct {
	Project              string
	Location             string
	IdentityMappingStore string
}

func (i *DiscoveryEngineIdentityMappingStoreIdentity) String() string {
	return DiscoveryEngineIdentityMappingStoreIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineIdentityMappingStoreIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineIdentityMappingStoreIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineIdentityMappingStore external=%q was not known (use %s): %w", ref, DiscoveryEngineIdentityMappingStoreIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineIdentityMappingStore external=%q was not known (use %s)", ref, DiscoveryEngineIdentityMappingStoreIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineIdentityMappingStoreIdentity) Host() string {
	return DiscoveryEngineIdentityMappingStoreIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineIdentityMappingStoreSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveryEngineIdentityMappingStoreIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &DiscoveryEngineIdentityMappingStoreIdentity{
		Project:              projectID,
		Location:             location,
		IdentityMappingStore: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineIdentityMappingStore) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineIdentityMappingStoreSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineIdentityMappingStoreIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineIdentityMappingStore identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements ExternalIdentifier
func (obj *DiscoveryEngineIdentityMappingStore) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
