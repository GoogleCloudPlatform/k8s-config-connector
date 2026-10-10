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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSaaSServiceMgmtTenantIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *SaaSServiceMgmtTenantIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/tenants/my-tenant",
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "my-tenant",
			},
		},
		{
			name: "full url",
			ref:  "https://saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/tenants/my-tenant",
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "my-tenant",
			},
		},
		{
			name: "double slash prefix",
			ref:  "//saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/tenants/my-tenant",
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "my-tenant",
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
			i := &SaaSServiceMgmtTenantIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/tenants/my-tenant" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/tenants/my-tenant")
				}
				if got := i.ParentString(); got != "projects/my-project/locations/us-central1" {
					t.Errorf("ParentString() = %q, want %q", got, "projects/my-project/locations/us-central1")
				}
			}
		})
	}
}

func TestSaaSServiceMgmtTenant_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *SaaSServiceMgmtTenant
		wantErr bool
		want    *SaaSServiceMgmtTenantIdentity
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &SaaSServiceMgmtTenant{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-tenant",
					Namespace: "default",
				},
				Spec: SaaSServiceMgmtTenantSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location:   common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("custom-tenant-id"),
				},
			},
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "custom-tenant-id",
			},
		},
		{
			name: "GetIdentity fallback to metadata.name",
			obj: &SaaSServiceMgmtTenant{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-tenant",
					Namespace: "default",
				},
				Spec: SaaSServiceMgmtTenantSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location: common.LazyPtr("us-central1"),
				},
			},
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "test-tenant",
			},
		},
		{
			name: "GetIdentity with matching status.externalRef",
			obj: &SaaSServiceMgmtTenant{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-tenant",
					Namespace: "default",
				},
				Spec: SaaSServiceMgmtTenantSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location:   common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("my-tenant"),
				},
				Status: SaaSServiceMgmtTenantStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/tenants/my-tenant"),
				},
			},
			want: &SaaSServiceMgmtTenantIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Tenant:   "my-tenant",
			},
		},
		{
			name: "GetIdentity with identity drift between spec and status",
			obj: &SaaSServiceMgmtTenant{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-tenant",
					Namespace: "default",
				},
				Spec: SaaSServiceMgmtTenantSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
					Location:   common.LazyPtr("us-central1"),
					ResourceID: common.LazyPtr("tenant-1"),
				},
				Status: SaaSServiceMgmtTenantStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/tenants/tenant-2"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := id.(*SaaSServiceMgmtTenantIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned type %T, want *SaaSServiceMgmtTenantIdentity", id)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestSaaSServiceMgmtSaaSIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *SaaSServiceMgmtSaaSIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/saas/my-saas",
			want: &SaaSServiceMgmtSaaSIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Saas:     "my-saas",
			},
		},
		{
			name: "full url",
			ref:  "https://saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/saas/my-saas",
			want: &SaaSServiceMgmtSaaSIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Saas:     "my-saas",
			},
		},
		{
			name: "double slash prefix",
			ref:  "//saasservicemgmt.googleapis.com/projects/my-project/locations/us-central1/saas/my-saas",
			want: &SaaSServiceMgmtSaaSIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Saas:     "my-saas",
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
			i := &SaaSServiceMgmtSaaSIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/saas/my-saas" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/saas/my-saas")
				}
				if got := i.ParentString(); got != "projects/my-project/locations/us-central1" {
					t.Errorf("ParentString() = %q, want %q", got, "projects/my-project/locations/us-central1")
				}
			}
		})
	}
}

func TestSaaSServiceMgmtSaaSRef_ParseExternalToIdentity(t *testing.T) {
	ref := &SaaSServiceMgmtSaaSRef{
		External: "projects/my-project/locations/us-central1/saas/my-saas",
	}

	parsedIdentity, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("ParseExternalToIdentity() error = %v", err)
	}

	expectedIdentity := &SaaSServiceMgmtSaaSIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Saas:     "my-saas",
	}

	actualIdentity, ok := parsedIdentity.(*SaaSServiceMgmtSaaSIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned wrong type %T, want *SaaSServiceMgmtSaaSIdentity", parsedIdentity)
	}

	if diff := cmp.Diff(expectedIdentity, actualIdentity); diff != "" {
		t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
	}
}
