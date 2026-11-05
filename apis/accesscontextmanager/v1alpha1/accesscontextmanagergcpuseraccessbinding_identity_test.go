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

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestAccessContextManagerGCPUserAccessBinding_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *AccessContextManagerGCPUserAccessBinding
		want    *AccessContextManagerGCPUserAccessBindingIdentity
		wantErr bool
	}{
		{
			name: "organizationRef with organizations/ prefix and explicit resourceID",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					OrganizationRef: &OrganizationRef{
						External: "organizations/123456789",
					},
					ResourceID: ptr.To("binding-1"),
				},
			},
			want: &AccessContextManagerGCPUserAccessBindingIdentity{
				Organization:         "123456789",
				GcpUserAccessBinding: "binding-1",
			},
			wantErr: false,
		},
		{
			name: "organizationRef with numeric only external",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					OrganizationRef: &OrganizationRef{
						External: "123456789",
					},
					ResourceID: ptr.To("binding-1"),
				},
			},
			want: &AccessContextManagerGCPUserAccessBindingIdentity{
				Organization:         "123456789",
				GcpUserAccessBinding: "binding-1",
			},
			wantErr: false,
		},
		{
			name: "organization from annotation without prefix",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/organization-id": "987654321",
					},
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					ResourceID: ptr.To("binding-1"),
				},
			},
			want: &AccessContextManagerGCPUserAccessBindingIdentity{
				Organization:         "987654321",
				GcpUserAccessBinding: "binding-1",
			},
			wantErr: false,
		},
		{
			name: "organization from annotation with organizations/ prefix",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/organization-id": "organizations/987654321",
					},
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					ResourceID: ptr.To("binding-1"),
				},
			},
			want: &AccessContextManagerGCPUserAccessBindingIdentity{
				Organization:         "987654321",
				GcpUserAccessBinding: "binding-1",
			},
			wantErr: false,
		},
		{
			name: "server-generated identity resolved from status.name",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					OrganizationRef: &OrganizationRef{
						External: "organizations/123456789",
					},
				},
				Status: AccessContextManagerGCPUserAccessBindingStatus{
					Name: ptr.To("organizations/123456789/gcpUserAccessBindings/server-assigned-id"),
				},
			},
			want: &AccessContextManagerGCPUserAccessBindingIdentity{
				Organization:         "123456789",
				GcpUserAccessBinding: "server-assigned-id",
			},
			wantErr: false,
		},
		{
			name: "status.name cross-check drift error",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					OrganizationRef: &OrganizationRef{
						External: "organizations/123456789",
					},
					ResourceID: ptr.To("binding-1"),
				},
				Status: AccessContextManagerGCPUserAccessBindingStatus{
					Name: ptr.To("organizations/123456789/gcpUserAccessBindings/different-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing organizationRef and annotation",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					ResourceID: ptr.To("binding-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid organizationRef format",
			obj: &AccessContextManagerGCPUserAccessBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-binding",
					Namespace: "default",
				},
				Spec: AccessContextManagerGCPUserAccessBindingSpec{
					OrganizationRef: &OrganizationRef{
						External: "projects/my-project",
					},
					ResourceID: ptr.To("binding-1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.obj.GetIdentity(ctx, nil)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			gotIdentity, ok := got.(*AccessContextManagerGCPUserAccessBindingIdentity)
			if !ok {
				t.Fatalf("GetIdentity() returned %T, want *AccessContextManagerGCPUserAccessBindingIdentity", got)
			}
			if diff := cmp.Diff(tc.want, gotIdentity); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAccessContextManagerGCPUserAccessBindingIdentity_Methods(t *testing.T) {
	id := &AccessContextManagerGCPUserAccessBindingIdentity{
		Organization:         "123456789",
		GcpUserAccessBinding: "binding-1",
	}

	if got, want := id.ParentString(), "organizations/123456789"; got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}

	if got, want := id.Host(), "accesscontextmanager.googleapis.com"; got != want {
		t.Errorf("Host() = %q, want %q", got, want)
	}

	if got, want := id.String(), "organizations/123456789/gcpUserAccessBindings/binding-1"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}

	if !id.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = false, want true")
	}

	emptyID := &AccessContextManagerGCPUserAccessBindingIdentity{
		Organization: "123456789",
	}
	if emptyID.HasIdentitySpecified() {
		t.Errorf("emptyID.HasIdentitySpecified() = true, want false")
	}
}
