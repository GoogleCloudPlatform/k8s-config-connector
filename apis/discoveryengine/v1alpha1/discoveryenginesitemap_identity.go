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
	return "projects/" + i.Project + "/locations/" + i.Location + "/collections/" + i.Collection + "/dataStores/" + i.DataStore + "/siteSearchEngine"
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
	if obj.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("spec.dataStoreRef not set")
	}
	dataStoreRef := *obj.Spec.DataStoreRef
	if _, err := dataStoreRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}
	dataStoreLink, err := ParseDiscoveryEngineDataStoreExternal(dataStoreRef.External)
	if err != nil {
		return nil, fmt.Errorf("parsing dataStoreRef.external=%q: %w", dataStoreRef.External, err)
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	identity := &DiscoveryEngineSitemapIdentity{
		Project:    dataStoreLink.ProjectID,
		Location:   dataStoreLink.Location,
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
