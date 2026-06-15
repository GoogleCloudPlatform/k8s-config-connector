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

func TestEventarcGoogleApiSourceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                string
		ref                 string
		wantProject         string
		wantLocation        string
		wantGoogleApiSource string
		wantErr             bool
	}{
		{
			name:                "valid full URL",
			ref:                 "//eventarc.googleapis.com/projects/my-project/locations/us-central1/googleApiSources/my-source",
			wantProject:         "my-project",
			wantLocation:        "us-central1",
			wantGoogleApiSource: "my-source",
		},
		{
			name:                "valid relative path",
			ref:                 "projects/my-project/locations/us-central1/googleApiSources/my-source",
			wantProject:         "my-project",
			wantLocation:        "us-central1",
			wantGoogleApiSource: "my-source",
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/my-project/googleApiSources/my-source",
			wantErr: true,
		},
		{
			name:    "invalid format - extra segments",
			ref:     "projects/my-project/locations/us-central1/googleApiSources/my-source/extra",
			wantErr: true,
		},
		{
			name:    "empty googleApiSource",
			ref:     "projects/my-project/locations/us-central1/googleApiSources/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &EventarcGoogleApiSourceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("FromExternal() Project got = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("FromExternal() Location got = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Google_api_source != tt.wantGoogleApiSource {
					t.Errorf("FromExternal() GoogleApiSource got = %v, want %v", i.Google_api_source, tt.wantGoogleApiSource)
				}
			}
		})
	}
}

func TestEventarcGoogleApiSourceIdentity_Methods(t *testing.T) {
	i := &EventarcGoogleApiSourceIdentity{
		Project:           "my-project",
		Location:          "us-central1",
		Google_api_source: "my-source",
	}

	want := "projects/my-project/locations/us-central1/googleApiSources/my-source"
	if got := i.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}

	wantParent := "projects/my-project/locations/us-central1"
	if gotParent := i.ParentString(); gotParent != wantParent {
		t.Errorf("ParentString() = %v, want %v", gotParent, wantParent)
	}
}
