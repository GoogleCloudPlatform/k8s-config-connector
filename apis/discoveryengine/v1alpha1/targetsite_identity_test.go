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

func TestDiscoveryEngineDataStoreTargetSiteIdentity_ParseTargetSiteExternal(t *testing.T) {
	tests := []struct {
		name           string
		ref            string
		wantErr        bool
		wantProject    string
		wantLocation   string
		wantCollection string
		wantDataStore  string
		wantTargetSite string
	}{
		{
			name:           "valid reference",
			ref:            "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/targetSites/my-target-site",
			wantProject:    "my-project",
			wantLocation:   "global",
			wantCollection: "default_collection",
			wantDataStore:  "my-datastore",
			wantTargetSite: "my-target-site",
		},
		{
			name:           "valid reference with slash prefix",
			ref:            "/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/targetSites/my-target-site",
			wantProject:    "my-project",
			wantLocation:   "global",
			wantCollection: "default_collection",
			wantDataStore:  "my-datastore",
			wantTargetSite: "my-target-site",
		},
		{
			name:           "valid reference with domain prefix",
			ref:            "//discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore/siteSearchEngine/targetSites/my-target-site",
			wantProject:    "my-project",
			wantLocation:   "global",
			wantCollection: "default_collection",
			wantDataStore:  "my-datastore",
			wantTargetSite: "my-target-site",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
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
				if got.ProjectID != tt.wantProject {
					t.Errorf("Project = %v, want %v", got.ProjectID, tt.wantProject)
				}
				if got.Location != tt.wantLocation {
					t.Errorf("Location = %v, want %v", got.Location, tt.wantLocation)
				}
				if got.Collection != tt.wantCollection {
					t.Errorf("Collection = %v, want %v", got.Collection, tt.wantCollection)
				}
				if got.DataStore != tt.wantDataStore {
					t.Errorf("DataStore = %v, want %v", got.DataStore, tt.wantDataStore)
				}
				if got.TargetSite != tt.wantTargetSite {
					t.Errorf("TargetSite = %v, want %v", got.TargetSite, tt.wantTargetSite)
				}
			}
		})
	}
}
