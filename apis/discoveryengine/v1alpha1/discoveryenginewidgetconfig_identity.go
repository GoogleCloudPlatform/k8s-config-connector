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
	_ identity.IdentityV2 = &DiscoveryEngineWidgetConfigIdentity{}
	_ identity.Resource   = &DiscoveryEngineWidgetConfig{}
)

var DiscoveryEngineWidgetConfigIdentityFormat = gcpurls.Template[DiscoveryEngineWidgetConfigIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/engines/{engine}/widgetConfigs/{widgetConfig}")

// DiscoveryEngineWidgetConfigIdentity is the identity of a GCP DiscoveryEngineWidgetConfig resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineWidgetConfigIdentity struct {
	Project      string
	Location     string
	Collection   string
	Engine       string
	WidgetConfig string
}

func (i *DiscoveryEngineWidgetConfigIdentity) String() string {
	return DiscoveryEngineWidgetConfigIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineWidgetConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineWidgetConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineWidgetConfig external=%q was not known (use %s): %w", ref, DiscoveryEngineWidgetConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineWidgetConfig external=%q was not known (use %s)", ref, DiscoveryEngineWidgetConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineWidgetConfigIdentity) Host() string {
	return DiscoveryEngineWidgetConfigIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineWidgetConfigSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DiscoveryEngineWidgetConfigIdentity, error) {
	widgetConfig, ok := obj.(*DiscoveryEngineWidgetConfig)
	if !ok {
		return nil, fmt.Errorf("object is not a DiscoveryEngineWidgetConfig")
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

	collectionID := "default_collection"
	if widgetConfig.Spec.CollectionID != nil {
		collectionID = *widgetConfig.Spec.CollectionID
	}

	if widgetConfig.Spec.EngineRef == nil {
		return nil, fmt.Errorf("spec.engineRef is not set")
	}

	engineRef := *widgetConfig.Spec.EngineRef
	normalizedEngine, err := engineRef.NormalizedExternal(ctx, reader, widgetConfig.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.engineRef: %w", err)
	}

	engineLink, err := parseDiscoveryEngineEngineExternal(normalizedEngine)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.engineRef external: %w", err)
	}

	// Validation checks: parent's project/location should match widget config's project/location
	if engineLink.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.engineRef project %q does not match spec.projectRef %q", engineLink.ProjectID, projectID)
	}
	if engineLink.Location != location {
		return nil, fmt.Errorf("resolved spec.engineRef location %q does not match spec.location %q", engineLink.Location, location)
	}
	if engineLink.Collection != collectionID {
		return nil, fmt.Errorf("resolved spec.engineRef collection %q does not match spec.collectionID %q", engineLink.Collection, collectionID)
	}

	identity := &DiscoveryEngineWidgetConfigIdentity{
		Project:      projectID,
		Location:     location,
		Collection:   collectionID,
		Engine:       engineLink.Engine,
		WidgetConfig: resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineWidgetConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineWidgetConfigSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineWidgetConfigIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineWidgetConfig identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *DiscoveryEngineWidgetConfig) ExternalIdentifier() *string {
	return c.Status.ExternalRef
}
