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
	"testing"
)

func TestParseTargetSiteExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *TargetSiteIdentity
	}{
		{
			name: "valid reference short",
			ref:  "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/targetSites/my-targetsite",
			want: &TargetSiteIdentity{
				DiscoveryEngineDataStoreID: &DiscoveryEngineDataStoreID{
					CollectionLink: &CollectionLink{
						ProjectAndLocation: &ProjectAndLocation{
							ProjectID: "my-project",
							Location:  "global",
						},
						Collection: "default_collection",
					},
					DataStore: "my-datastore",
				},
				TargetSite: "my-targetsite",
			},
		},
		{
			name: "valid reference full url",
			ref:  "//discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/targetSites/my-targetsite",
			want: &TargetSiteIdentity{
				DiscoveryEngineDataStoreID: &DiscoveryEngineDataStoreID{
					CollectionLink: &CollectionLink{
						ProjectAndLocation: &ProjectAndLocation{
							ProjectID: "my-project",
							Location:  "global",
						},
						Collection: "default_collection",
					},
					DataStore: "my-datastore",
				},
				TargetSite: "my-targetsite",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTargetSiteExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTargetSiteExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.String() != tt.want.String() {
					t.Errorf("got.String() = %v, want %v", got.String(), tt.want.String())
				}
				if got.DiscoveryEngineDataStoreID.ProjectID != tt.want.DiscoveryEngineDataStoreID.ProjectID {
					t.Errorf("ProjectID = %v, want %v", got.DiscoveryEngineDataStoreID.ProjectID, tt.want.DiscoveryEngineDataStoreID.ProjectID)
				}
				if got.DiscoveryEngineDataStoreID.Location != tt.want.DiscoveryEngineDataStoreID.Location {
					t.Errorf("Location = %v, want %v", got.DiscoveryEngineDataStoreID.Location, tt.want.DiscoveryEngineDataStoreID.Location)
				}
				if got.DiscoveryEngineDataStoreID.Collection != tt.want.DiscoveryEngineDataStoreID.Collection {
					t.Errorf("Collection = %v, want %v", got.DiscoveryEngineDataStoreID.Collection, tt.want.DiscoveryEngineDataStoreID.Collection)
				}
				if got.DiscoveryEngineDataStoreID.DataStore != tt.want.DiscoveryEngineDataStoreID.DataStore {
					t.Errorf("DataStore = %v, want %v", got.DiscoveryEngineDataStoreID.DataStore, tt.want.DiscoveryEngineDataStoreID.DataStore)
				}
				if got.TargetSite != tt.want.TargetSite {
					t.Errorf("TargetSite = %v, want %v", got.TargetSite, tt.want.TargetSite)
				}
			}
		})
	}
}
