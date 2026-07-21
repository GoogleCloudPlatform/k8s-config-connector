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
	_ identity.IdentityV2 = &DiscoveryEngineSessionIdentity{}
	_ identity.Resource   = &DiscoveryEngineSession{}
)

var DiscoveryEngineSessionIdentityFormat = gcpurls.Template[DiscoveryEngineSessionIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/dataStores/{datastore}/sessions/{session}")

// DiscoveryEngineSessionIdentity is the identity of a GCP DiscoveryEngineSession resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineSessionIdentity struct {
	Project   string
	Location  string
	Datastore string
	Session   string
}

func (i *DiscoveryEngineSessionIdentity) String() string {
	return DiscoveryEngineSessionIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineSessionIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineSessionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineSession external=%q was not known (use %s): %w", ref, DiscoveryEngineSessionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineSession external=%q was not known (use %s)", ref, DiscoveryEngineSessionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineSessionIdentity) Host() string {
	return DiscoveryEngineSessionIdentityFormat.Host()
}

func (i *DiscoveryEngineSessionIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/dataStores/%s", i.Project, i.Location, i.Datastore)
}

func getIdentityFromDiscoveryEngineSessionSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineSession) (*DiscoveryEngineSessionIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	if obj.Spec.Location == nil || *obj.Spec.Location == "" {
		return nil, fmt.Errorf("spec.location is not set")
	}
	location := *obj.Spec.Location

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

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

	// Validation checks: parent's project/location should match session's project/location
	if dataStoreLink.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.dataStoreRef project %q does not match spec.projectRef %q", dataStoreLink.ProjectID, projectID)
	}
	if dataStoreLink.Location != location {
		return nil, fmt.Errorf("resolved spec.dataStoreRef location %q does not match spec.location %q", dataStoreLink.Location, location)
	}

	identity := &DiscoveryEngineSessionIdentity{
		Project:   projectID,
		Location:  location,
		Datastore: dataStoreLink.DataStore,
		Session:   resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineSession) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineSessionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineSessionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineSession identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineSession) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
