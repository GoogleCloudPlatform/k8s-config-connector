// Copyright 2024 Google LLC
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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

func TestOrganizationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *OrganizationIdentity
		wantErr bool
	}{
		{
			name: "valid organization reference",
			ref:  "organizations/1234567890",
			want: &OrganizationIdentity{
				OrganizationID: "1234567890",
			},
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "organizations/12345/other",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "missing prefix",
			ref:     "1234567890",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &OrganizationIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.OrganizationID != tt.want.OrganizationID {
					t.Errorf("OrganizationIdentity.FromExternal() = %v, want %v", i, tt.want)
				}
			}
		})
	}
}

func TestOrganizationRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "organizations/12345",
			wantErr:  false,
		},
		{
			name:     "invalid external",
			external: "invalid/12345",
			wantErr:  true,
		},
		{
			name:     "empty external",
			external: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &OrganizationRef{}
			if err := r.ValidateExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("OrganizationRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrganizationImplementsIdentity(t *testing.T) {
	var _ identity.Identity = &Organization{}
	var _ identity.Identity = &OrganizationIdentity{}
}

func TestResolveOrganization(t *testing.T) {
	tests := []struct {
		name    string
		ref     *OrganizationRef
		wantID  string
		wantErr bool
	}{
		{
			name: "valid ref",
			ref: &OrganizationRef{
				External: "organizations/12345",
			},
			wantID:  "12345",
			wantErr: false,
		},
		{
			name: "invalid ref",
			ref: &OrganizationRef{
				External: "bad/format",
			},
			wantErr: true,
		},
		{
			name:    "nil ref",
			ref:     nil,
			wantErr: false, // returns nil, nil
		},
		{
			name: "empty external",
			ref: &OrganizationRef{
				External: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveOrganization(context.Background(), nil, nil, tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.ref == nil {
				if got != nil {
					t.Errorf("ResolveOrganization() got = %v, want nil", got)
				}
				return
			}
			if !tt.wantErr {
				if got.OrganizationID != tt.wantID {
					t.Errorf("ResolveOrganization() got ID = %v, want %v", got.OrganizationID, tt.wantID)
				}
			}
		})
	}
}
