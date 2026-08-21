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

package v1beta1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDNSRecordSetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DNSRecordSetIdentity
	}{
		{
			name: "valid reference - primary (with location)",
			ref:  "projects/my-project/locations/global/managedZones/my-zone/rrsets/my-record.example.com.",
			want: &DNSRecordSetIdentity{
				Project:     "my-project",
				Location:    "global",
				ManagedZone: "my-zone",
				Name:        "my-record.example.com.",
			},
		},
		{
			name: "valid reference - fallback (without location)",
			ref:  "projects/my-project/managedZones/my-zone/rrsets/my-record.example.com.",
			want: &DNSRecordSetIdentity{
				Project:     "my-project",
				ManagedZone: "my-zone",
				Name:        "my-record.example.com.",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url - primary (with location)",
			ref:  "https://dns.googleapis.com/projects/my-project/locations/global/managedZones/my-zone/rrsets/www.example.com.",
			want: &DNSRecordSetIdentity{
				Project:     "my-project",
				Location:    "global",
				ManagedZone: "my-zone",
				Name:        "www.example.com.",
			},
		},
		{
			name: "full url - fallback (without location)",
			ref:  "https://dns.googleapis.com/projects/my-project/managedZones/my-zone/rrsets/www.example.com.",
			want: &DNSRecordSetIdentity{
				Project:     "my-project",
				ManagedZone: "my-zone",
				Name:        "www.example.com.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DNSRecordSetIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
