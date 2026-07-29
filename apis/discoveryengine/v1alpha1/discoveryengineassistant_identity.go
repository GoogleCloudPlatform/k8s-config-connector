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
	_ identity.IdentityV2 = &DiscoveryEngineAssistantIdentity{}
	_ identity.Resource   = &DiscoveryEngineAssistant{}
)

// DiscoveryEngineAssistantIdentityFormat is the format template for DiscoveryEngineAssistant identity.
// The assistant ID (represented by `{assistant}`) is user-specified when creating the assistant resource.
var DiscoveryEngineAssistantIdentityFormat = gcpurls.Template[DiscoveryEngineAssistantIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/engines/{engine}/assistants/{assistant}")

// DiscoveryEngineAssistantIdentity is the identity of a GCP DiscoveryEngineAssistant resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineAssistantIdentity struct {
	Project    string
	Location   string
	Collection string
	Engine     string
	Assistant  string
}

func (i *DiscoveryEngineAssistantIdentity) String() string {
	return DiscoveryEngineAssistantIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineAssistantIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineAssistantIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineAssistant external=%q was not known (use %s): %w", ref, DiscoveryEngineAssistantIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineAssistant external=%q was not known (use %s)", ref, DiscoveryEngineAssistantIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineAssistantIdentity) Host() string {
	return DiscoveryEngineAssistantIdentityFormat.Host()
}

func ParseDiscoveryEngineAssistantExternal(external string) (*DiscoveryEngineAssistantIdentity, error) {
	identity := &DiscoveryEngineAssistantIdentity{}
	if err := identity.FromExternal(external); err != nil {
		return nil, err
	}
	return identity, nil
}

func (i *DiscoveryEngineAssistantIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", i.Project, i.Location, i.Collection, i.Engine)
}

func getIdentityFromDiscoveryEngineAssistantSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineAssistant) (*DiscoveryEngineAssistantIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if obj.Spec.EngineRef == nil {
		return nil, fmt.Errorf("spec.engineRef is required")
	}

	engineRef := obj.Spec.EngineRef
	normalizedExternal, err := engineRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("resolving spec.engineRef: %w", err)
	}

	engineLink, err := parseDiscoveryEngineEngineExternal(normalizedExternal)
	if err != nil {
		return nil, fmt.Errorf("parsing engineRef.external=%q: %w", normalizedExternal, err)
	}

	// Validation checks: parent's project/location should match assistant's project/location
	if engineLink.ProjectID != projectID {
		return nil, fmt.Errorf("resolved spec.engineRef project %q does not match spec.projectRef %q", engineLink.ProjectID, projectID)
	}
	if engineLink.Location != location {
		return nil, fmt.Errorf("resolved spec.engineRef location %q does not match spec.location %q", engineLink.Location, location)
	}

	identity := &DiscoveryEngineAssistantIdentity{
		Project:    projectID,
		Location:   location,
		Collection: engineLink.Collection,
		Engine:     engineLink.Engine,
		Assistant:  resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineAssistant) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineAssistantSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineAssistantIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineAssistant identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *DiscoveryEngineAssistant) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
