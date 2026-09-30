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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDiscoveryEngineSitemapIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DiscoveryEngineSitemapIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/my-sitemap",
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "my-sitemap",
			},
		},
		{
			name: "full url",
			ref:  "https://discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/my-sitemap",
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "my-sitemap",
			},
		},
		{
			name: "double slash prefix",
			ref:  "//discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/my-sitemap",
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "my-sitemap",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DiscoveryEngineSitemapIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/my-sitemap" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/my-sitemap")
				}
				if got := i.ParentString(); got != "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine" {
					t.Errorf("ParentString() = %q, want %q", got, "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine")
				}
			}
		})
	}
}

func TestDiscoveryEngineSitemap_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *DiscoveryEngineSitemap
		wantErr               bool
		want                  *DiscoveryEngineSitemapIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &DiscoveryEngineSitemap{
				Spec: DiscoveryEngineSitemapSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
					ResourceID: common.LazyPtr("my-sitemap"),
				},
			},
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "my-sitemap",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity (empty spec resourceID, defaulted from status)",
			obj: &DiscoveryEngineSitemap{
				Spec: DiscoveryEngineSitemapSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
				},
				Status: DiscoveryEngineSitemapStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/server-gen-id"),
				},
			},
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "server-gen-id",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID and empty status (not yet created)",
			obj: &DiscoveryEngineSitemap{
				Spec: DiscoveryEngineSitemapSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
				},
			},
			want: &DiscoveryEngineSitemapIdentity{
				Project:    "my-project",
				Location:   "global",
				Collection: "default_collection",
				DataStore:  "my-datastore",
				Sitemap:    "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with identity drift between spec and status",
			obj: &DiscoveryEngineSitemap{
				Spec: DiscoveryEngineSitemapSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
					ResourceID: common.LazyPtr("sitemap-1"),
				},
				Status: DiscoveryEngineSitemapStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/sitemaps/sitemap-2"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := id.(*DiscoveryEngineSitemapIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned type %T, want *DiscoveryEngineSitemapIdentity", id)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if got := got.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
			}
		})
	}
}
