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

	"github.com/google/go-cmp/cmp"
)

func TestParseDiscoveryEngineDataStoreExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		wantID     *DiscoveryEngineDataStoreID
		wantString string
	}{
		{
			name: "8 tokens canonical",
			ref:  "projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
			wantID: &DiscoveryEngineDataStoreID{
				CollectionLink: &CollectionLink{
					ProjectAndLocation: &ProjectAndLocation{
						ProjectID: "my-proj",
						Location:  "global",
					},
					Collection: "default_collection",
				},
				DataStore: "my-ds",
			},
			wantString: "projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
		},
		{
			name: "8 tokens lowercase datastores",
			ref:  "projects/my-proj/locations/global/collections/custom_col/datastores/my-ds",
			wantID: &DiscoveryEngineDataStoreID{
				CollectionLink: &CollectionLink{
					ProjectAndLocation: &ProjectAndLocation{
						ProjectID: "my-proj",
						Location:  "global",
					},
					Collection: "custom_col",
				},
				DataStore: "my-ds",
			},
			wantString: "projects/my-proj/locations/global/collections/custom_col/dataStores/my-ds",
		},
		{
			name: "6 tokens dataStores defaults collection",
			ref:  "projects/my-proj/locations/global/dataStores/my-ds",
			wantID: &DiscoveryEngineDataStoreID{
				CollectionLink: &CollectionLink{
					ProjectAndLocation: &ProjectAndLocation{
						ProjectID: "my-proj",
						Location:  "global",
					},
					Collection: "default_collection",
				},
				DataStore: "my-ds",
			},
			wantString: "projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
		},
		{
			name: "6 tokens lowercase datastores defaults collection",
			ref:  "projects/my-proj/locations/global/datastores/my-ds",
			wantID: &DiscoveryEngineDataStoreID{
				CollectionLink: &CollectionLink{
					ProjectAndLocation: &ProjectAndLocation{
						ProjectID: "my-proj",
						Location:  "global",
					},
					Collection: "default_collection",
				},
				DataStore: "my-ds",
			},
			wantString: "projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
		},
		{
			name: "with service prefix",
			ref:  "//discoveryengine.googleapis.com/projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
			wantID: &DiscoveryEngineDataStoreID{
				CollectionLink: &CollectionLink{
					ProjectAndLocation: &ProjectAndLocation{
						ProjectID: "my-proj",
						Location:  "global",
					},
					Collection: "default_collection",
				},
				DataStore: "my-ds",
			},
			wantString: "projects/my-proj/locations/global/collections/default_collection/dataStores/my-ds",
		},
		{
			name:    "invalid format",
			ref:     "projects/my-proj/locations/global",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseDiscoveryEngineDataStoreExternal(tc.ref)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ParseDiscoveryEngineDataStoreExternal(%q) err = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.wantID, got); diff != "" {
				t.Errorf("ParseDiscoveryEngineDataStoreExternal(%q) diff (-want +got):\n%s", tc.ref, diff)
			}
			if gotStr := got.String(); gotStr != tc.wantString {
				t.Errorf("got.String() = %q, want %q", gotStr, tc.wantString)
			}

			// Validate via DiscoveryEngineDataStoreRef.ValidateExternal
			ref := &DiscoveryEngineDataStoreRef{External: tc.ref}
			if err := ref.ValidateExternal(tc.ref); err != nil {
				t.Errorf("ref.ValidateExternal(%q) returned unexpected error: %v", tc.ref, err)
			}
		})
	}
}
