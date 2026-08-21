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

	analyticsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestMarketingPlatformAdminAnalyticsAccountLinkIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *MarketingPlatformAdminAnalyticsAccountLinkIdentity
	}{
		{
			name: "valid reference",
			ref:  "organizations/12345/analyticsAccountLinks/67890",
			want: &MarketingPlatformAdminAnalyticsAccountLinkIdentity{
				Organization:         "12345",
				AnalyticsAccountLink: "67890",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://marketingplatformadmin.googleapis.com/organizations/12345/analyticsAccountLinks/67890",
			want: &MarketingPlatformAdminAnalyticsAccountLinkIdentity{
				Organization:         "12345",
				AnalyticsAccountLink: "67890",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &MarketingPlatformAdminAnalyticsAccountLinkIdentity{}
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

func TestMarketingPlatformAdminAnalyticsAccountLink_GetIdentity(t *testing.T) {
	ctx := context.Background()

	// Mock referenced AnalyticsAccount
	analyticsAccount := &unstructured.Unstructured{}
	analyticsAccount.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "analytics.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "AnalyticsAccount",
	})
	analyticsAccount.SetName("my-account")
	analyticsAccount.SetNamespace("my-namespace")
	if err := unstructured.SetNestedField(analyticsAccount.Object, "accounts/67890", "status", "externalRef"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	// Mock Organization
	orgObj := &unstructured.Unstructured{}
	orgObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	})
	orgObj.SetName("my-org")
	orgObj.SetNamespace("my-namespace")
	if err := unstructured.SetNestedField(orgObj.Object, "12345", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	s := runtime.NewScheme()
	_ = AddToScheme(s)
	s.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "analytics.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "AnalyticsAccountList",
	}, &unstructured.UnstructuredList{})
	s.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "OrganizationList",
	}, &unstructured.UnstructuredList{})

	fakeClient := fake.NewClientBuilder().WithScheme(s).WithObjects(analyticsAccount, orgObj).Build()

	tests := []struct {
		name        string
		obj         *MarketingPlatformAdminAnalyticsAccountLink
		want        *MarketingPlatformAdminAnalyticsAccountLinkIdentity
		expectError bool
	}{
		{
			name: "valid external analytics account ref and organization ref",
			obj: &MarketingPlatformAdminAnalyticsAccountLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "my-namespace",
				},
				Spec: MarketingPlatformAdminAnalyticsAccountLinkSpec{
					OrganizationRef: &refs.OrganizationRef{
						External: "organizations/12345",
					},
					AnalyticsAccountRef: &analyticsv1alpha1.AccountRef{
						External: "accounts/67890",
					},
				},
			},
			want: &MarketingPlatformAdminAnalyticsAccountLinkIdentity{
				Organization:         "12345",
				AnalyticsAccountLink: "67890",
			},
		},
		{
			name: "valid k8s analytics account ref and organization ref",
			obj: &MarketingPlatformAdminAnalyticsAccountLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "my-namespace",
				},
				Spec: MarketingPlatformAdminAnalyticsAccountLinkSpec{
					OrganizationRef: &refs.OrganizationRef{
						External: "organizations/12345",
					},
					AnalyticsAccountRef: &analyticsv1alpha1.AccountRef{
						Name: "my-account",
					},
				},
			},
			want: &MarketingPlatformAdminAnalyticsAccountLinkIdentity{
				Organization:         "12345",
				AnalyticsAccountLink: "67890",
			},
		},
		{
			name: "missing analyticsAccountRef",
			obj: &MarketingPlatformAdminAnalyticsAccountLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "my-namespace",
				},
				Spec: MarketingPlatformAdminAnalyticsAccountLinkSpec{
					OrganizationRef: &refs.OrganizationRef{
						External: "organizations/12345",
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.expectError {
				t.Fatalf("GetIdentity() error = %v, expectError %v", err, tt.expectError)
			}
			if !tt.expectError {
				gotIdentity := got.(*MarketingPlatformAdminAnalyticsAccountLinkIdentity)
				if diff := cmp.Diff(tt.want, gotIdentity); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
