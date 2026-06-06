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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
)

func TestPrivilegedAccessManagerEntitlementIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *PrivilegedAccessManagerEntitlementIdentity
	}{
		{
			name: "valid project-level entitlement reference",
			ref:  "projects/my-project/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Project:     "my-project",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
		},
		{
			name: "valid folder-level entitlement reference",
			ref:  "folders/123456/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Folder:      "123456",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
		},
		{
			name: "valid organization-level entitlement reference",
			ref:  "organizations/789012/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Organization: "789012",
				Location:     "global",
				Entitlement:  "my-entitlement",
			},
		},
		{
			name: "full url project-level",
			ref:  "https://privilegedaccessmanager.googleapis.com/projects/my-project/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Project:     "my-project",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
		},
		{
			name: "full url folder-level",
			ref:  "https://privilegedaccessmanager.googleapis.com/folders/123456/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Folder:      "123456",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
		},
		{
			name: "full url organization-level",
			ref:  "https://privilegedaccessmanager.googleapis.com/organizations/789012/locations/global/entitlements/my-entitlement",
			want: &PrivilegedAccessManagerEntitlementIdentity{
				Organization: "789012",
				Location:     "global",
				Entitlement:  "my-entitlement",
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
			i := &PrivilegedAccessManagerEntitlementIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestPrivilegedAccessManagerEntitlementIdentity_String(t *testing.T) {
	tests := []struct {
		name string
		id   PrivilegedAccessManagerEntitlementIdentity
		want string
	}{
		{
			name: "project-level",
			id: PrivilegedAccessManagerEntitlementIdentity{
				Project:     "my-project",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
			want: "projects/my-project/locations/global/entitlements/my-entitlement",
		},
		{
			name: "folder-level",
			id: PrivilegedAccessManagerEntitlementIdentity{
				Folder:      "123456",
				Location:    "global",
				Entitlement: "my-entitlement",
			},
			want: "folders/123456/locations/global/entitlements/my-entitlement",
		},
		{
			name: "organization-level",
			id: PrivilegedAccessManagerEntitlementIdentity{
				Organization: "789012",
				Location:     "global",
				Entitlement:  "my-entitlement",
			},
			want: "organizations/789012/locations/global/entitlements/my-entitlement",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.id.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivilegedAccessManagerEntitlementIdentity_GetIdentityValidation(t *testing.T) {
	tests := []struct {
		name    string
		obj     *PrivilegedAccessManagerEntitlement
		wantErr bool
	}{
		{
			name: "only projectRef",
			obj: &PrivilegedAccessManagerEntitlement{
				Spec: PrivilegedAccessManagerEntitlementSpec{
					ProjectRef: &refs.ProjectRef{External: "projects/my-project"},
				},
			},
			wantErr: false,
		},
		{
			name: "projectRef and folderRef",
			obj: &PrivilegedAccessManagerEntitlement{
				Spec: PrivilegedAccessManagerEntitlementSpec{
					ProjectRef: &refs.ProjectRef{External: "projects/my-project"},
					FolderRef:  &refs.FolderRef{External: "folders/123456"},
				},
			},
			wantErr: true,
		},
		{
			name: "projectRef, folderRef and organizationRef",
			obj: &PrivilegedAccessManagerEntitlement{
				Spec: PrivilegedAccessManagerEntitlementSpec{
					ProjectRef:      &refs.ProjectRef{External: "projects/my-project"},
					FolderRef:       &refs.FolderRef{External: "folders/123456"},
					OrganizationRef: &refs.OrganizationRef{External: "organizations/789012"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Give it a metadata name so refs.GetResourceID has something
			tt.obj.SetName("my-entitlement")
			_, err := getIdentityFromPrivilegedAccessManagerEntitlementSpec(nil, nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromPrivilegedAccessManagerEntitlementSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
