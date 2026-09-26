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
	_ identity.ServerGeneratedIdentity = &DiscoveryEngineSitemapIdentity{}
	_ identity.Resource                = &DiscoveryEngineSitemap{}
)

var DiscoveryEngineSitemapIdentityFormat = gcpurls.Template[DiscoveryEngineSitemapIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}/siteSearchEngine/sitemaps/{sitemap}")

// DiscoveryEngineSitemapIdentity is the identity of a GCP DiscoveryEngineSitemap resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineSitemapIdentity struct {
	Project    string
	Location   string
	Collection string
	DataStore  string
	Sitemap    string
}

func (i *DiscoveryEngineSitemapIdentity) HasIdentitySpecified() bool {
	return i.Sitemap != ""
}

func (i *DiscoveryEngineSitemapIdentity) String() string {
	return DiscoveryEngineSitemapIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineSitemapIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/siteSearchEngine", i.Project, i.Location, i.Collection, i.DataStore)
}

func (i *DiscoveryEngineSitemapIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineSitemapIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineSitemap external=%q was not known (use %s): %w", ref, DiscoveryEngineSitemapIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineSitemap external=%q was not known (use %s)", ref, DiscoveryEngineSitemapIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineSitemapIdentity) Host() string {
	return DiscoveryEngineSitemapIdentityFormat.Host()
}

func getIdentityFromDiscoveryEngineSitemapSpec(ctx context.Context, reader client.Reader, obj *DiscoveryEngineSitemap) (*DiscoveryEngineSitemapIdentity, error) {
	// For DiscoveryEngineSitemap, resourceID is optional / server-generated.
	// We retrieve it directly from Spec.ResourceID to avoid falling back to GetName().
	resourceID := common.ValueOf(obj.Spec.ResourceID)

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

	// Validation checks: parent's project/location should match sitemap's project/location
	if !IsProjectIDMatch(dataStoreLink.ProjectID, projectID) {
		return nil, fmt.Errorf("resolved spec.dataStoreRef project %q does not match spec.projectRef %q", dataStoreLink.ProjectID, projectID)
	}
	if dataStoreLink.Location != location {
		return nil, fmt.Errorf("resolved spec.dataStoreRef location %q does not match spec.location %q", dataStoreLink.Location, location)
	}

	identity := &DiscoveryEngineSitemapIdentity{
		Project:    projectID,
		Location:   location,
		Collection: dataStoreLink.Collection,
		DataStore:  dataStoreLink.DataStore,
		Sitemap:    resourceID,
	}
	return identity, nil
}

func (obj *DiscoveryEngineSitemap) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDiscoveryEngineSitemapSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DiscoveryEngineSitemapIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.Sitemap == "" {
			specIdentity.Sitemap = statusIdentity.Sitemap
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DiscoveryEngineSitemap identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *DiscoveryEngineSitemap) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
