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
	_ identity.IdentityV2 = &DiscoveryEngineControlIdentity{}
	_ identity.Resource   = &DiscoveryEngineControl{}
)

var DiscoveryEngineControlIdentityFormat = gcpurls.Template[DiscoveryEngineControlIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/dataStores/{datastore}/controls/{control}")

// +k8s:deepcopy-gen=false
type DiscoveryEngineControlIdentity struct {
	Project   string
	Location  string
	Datastore string
	Control   string
}

func (i *DiscoveryEngineControlIdentity) String() string {
	return DiscoveryEngineControlIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineControlIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineControlIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineControl external=%q was not known (use %s): %w", ref, DiscoveryEngineControlIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineControl external=%q was not known (use %s)", ref, DiscoveryEngineControlIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineControlIdentity) Host() string {
	return DiscoveryEngineControlIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineControlSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveryEngineControlIdentity, error) {
	control, ok := obj.(*DiscoveryEngineControl)
	if !ok {
		return nil, fmt.Errorf("object is not a DiscoveryEngineControl")
	}
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if control.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("spec.dataStoreRef is not set")
	}

	dataStoreRef := *control.Spec.DataStoreRef
	normalizedDataStore, err := dataStoreRef.NormalizedExternal(ctx, reader, control.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}

	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(normalizedDataStore)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.dataStoreRef external: %w", err)
	}

	// Validation checks: parent's project/location should match control's project/location
	if dataStoreLink.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.dataStoreRef project %q does not match spec.projectRef %q", dataStoreLink.ProjectID, projectID)
	}
	if dataStoreLink.Location != location {
		return nil, fmt.Errorf("resolved spec.dataStoreRef location %q does not match spec.location %q", dataStoreLink.Location, location)
	}

	identity := &DiscoveryEngineControlIdentity{
		Project:   projectID,
		Location:  location,
		Datastore: dataStoreLink.DataStore,
		Control:   resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineControl) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineControlSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineControlIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineControl identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DiscoveryEngineControl) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
