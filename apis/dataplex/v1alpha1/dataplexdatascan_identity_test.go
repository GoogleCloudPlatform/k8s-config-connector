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

func TestDataScanIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *DataScanIdentity
		wantErr bool
	}{
		{
			name: "valid fully qualified",
			ref:  "projects/12345/locations/us-central1/dataScans/my-scan",
			want: &DataScanIdentity{
				Project:  "12345",
				Location: "us-central1",
				DataScan: "my-scan",
			},
			wantErr: false,
		},
		{
			name: "valid with service domain",
			ref:  "//dataplex.googleapis.com/projects/my-project/locations/europe-west1/dataScans/scan2",
			want: &DataScanIdentity{
				Project:  "my-project",
				Location: "europe-west1",
				DataScan: "scan2",
			},
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "projects/12345/locations/us-central1/lakes/my-lake/dataScans/my-scan",
			wantErr: true,
		},
		{
			name:    "wrong resource type",
			ref:     "projects/12345/locations/us-central1/zones/my-zone",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataScanIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && i.String() != tt.want.String() {
				t.Errorf("FromExternal() got = %v, want %v", i.String(), tt.want.String())
			}
		})
	}
}
