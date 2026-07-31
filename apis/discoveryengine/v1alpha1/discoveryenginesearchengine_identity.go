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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DiscoveryEngineSearchEngineIdentity{}
	_ identity.Resource   = &DiscoveryEngineSearchEngine{}
)

var DiscoveryEngineSearchEngineIdentityFormat = gcpurls.Template[DiscoveryEngineSearchEngineIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/dataStores/{datastore}/siteSearchEngine")

// DiscoveryEngineSearchEngineIdentity is the identity of a GCP DiscoveryEngineSearchEngine resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineSearchEngineIdentity struct {
	Project    string
	Location   string
	Collection string
	Datastore  string
}

func (i *DiscoveryEngineSearchEngineIdentity) String() string {
	return DiscoveryEngineSearchEngineIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineSearchEngineIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineSearchEngineIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineSearchEngine external=%q was not known (use %s): %w", ref, DiscoveryEngineSearchEngineIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineSearchEngine external=%q was not known (use %s)", ref, DiscoveryEngineSearchEngineIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineSearchEngineIdentity) Host() string {
	return DiscoveryEngineSearchEngineIdentityFormat.Host()
}

func (i *DiscoveryEngineSearchEngineIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", i.Project, i.Location, i.Collection, i.Datastore)
}

func getIdentityFromDiscoveryEngineSearchEngineSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineSearchEngine) (*DiscoveryEngineSearchEngineIdentity, error) {
	if obj.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("spec.dataStoreRef is not set")
	}

	dataStoreRef := *obj.Spec.DataStoreRef
	normalizedDataStore, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}

	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(normalizedDataStore)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.dataStoreRef external: %w", err)
	}

	identity := &DiscoveryEngineSearchEngineIdentity{
		Project:    dataStoreLink.ProjectID,
		Location:   dataStoreLink.Location,
		Collection: dataStoreLink.Collection,
		Datastore:  dataStoreLink.DataStore,
	}
	return identity, nil
}

func (obj *DiscoveryEngineSearchEngine) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineSearchEngineSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineSearchEngineIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineSearchEngine identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineSearchEngine) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
