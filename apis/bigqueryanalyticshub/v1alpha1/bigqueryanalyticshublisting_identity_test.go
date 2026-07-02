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

func TestBigQueryAnalyticsHubListingIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigQueryAnalyticsHubListingIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/dataExchanges/my-exchange/listings/my-listing",
			want: &BigQueryAnalyticsHubListingIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				DataExchange: "my-exchange",
				Listing:      "my-listing",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://analyticshub.googleapis.com/projects/my-project/locations/us-central1/dataExchanges/my-exchange/listings/my-listing",
			want: &BigQueryAnalyticsHubListingIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				DataExchange: "my-exchange",
				Listing:      "my-listing",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigQueryAnalyticsHubListingIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.DataExchange != tt.want.DataExchange {
					t.Errorf("DataExchange = %v, want %v", i.DataExchange, tt.want.DataExchange)
				}
				if i.Listing != tt.want.Listing {
					t.Errorf("Listing = %v, want %v", i.Listing, tt.want.Listing)
				}
			}
		})
	}
}

func TestBigQueryAnalyticsHubListingIdentity_String(t *testing.T) {
	i := &BigQueryAnalyticsHubListingIdentity{
		Project:      "my-project",
		Location:     "us-central1",
		DataExchange: "my-exchange",
		Listing:      "my-listing",
	}
	want := "projects/my-project/locations/us-central1/dataExchanges/my-exchange/listings/my-listing"
	if got := i.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestBigQueryAnalyticsHubListingIdentity_Host(t *testing.T) {
	i := &BigQueryAnalyticsHubListingIdentity{}
	want := "analyticshub.googleapis.com"
	if got := i.Host(); got != want {
		t.Errorf("Host() = %v, want %v", got, want)
	}
}
