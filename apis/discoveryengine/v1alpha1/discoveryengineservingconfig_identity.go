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
	_ identity.IdentityV2 = &DiscoveryEngineServingConfigIdentity{}
	_ identity.Resource   = &DiscoveryEngineServingConfig{}
)

var DiscoveryEngineServingConfigIdentityFormat = gcpurls.Template[DiscoveryEngineServingConfigIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/engines/{engine}/servingConfigs/{servingConfig}")

// DiscoveryEngineServingConfigIdentity is the identity of a GCP DiscoveryEngineServingConfig resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineServingConfigIdentity struct {
	Project       string
	Location      string
	Collection    string
	Engine        string
	ServingConfig string
}

func (i *DiscoveryEngineServingConfigIdentity) String() string {
	return DiscoveryEngineServingConfigIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineServingConfigIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", i.Project, i.Location, i.Collection, i.Engine)
}

func (i *DiscoveryEngineServingConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineServingConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineServingConfig external=%q was not known (use %s): %w", ref, DiscoveryEngineServingConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineServingConfig external=%q was not known (use %s)", ref, DiscoveryEngineServingConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineServingConfigIdentity) Host() string {
	return DiscoveryEngineServingConfigIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineServingConfigSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineServingConfig) (*DiscoveryEngineServingConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if obj.Spec.EngineRef == nil {
		return nil, fmt.Errorf("spec.engineRef is not set")
	}

	engineRef := *obj.Spec.EngineRef
	normalizedEngine, err := engineRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.engineRef: %w", err)
	}

	engineLink, err := parseDiscoveryEngineEngineExternal(normalizedEngine)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.engineRef external: %w", err)
	}

	// Validation checks: parent's project/location should match serving config's project/location
	if engineLink.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.engineRef project %q does not match spec.projectRef %q", engineLink.ProjectID, projectID)
	}
	if engineLink.Location != location {
		return nil, fmt.Errorf("resolved spec.engineRef location %q does not match spec.location %q", engineLink.Location, location)
	}

	identity := &DiscoveryEngineServingConfigIdentity{
		Project:       projectID,
		Location:      location,
		Collection:    engineLink.Collection,
		Engine:        engineLink.Engine,
		ServingConfig: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineServingConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineServingConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineServingConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineServingConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineServingConfig) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
