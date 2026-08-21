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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIAPBrandIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                  string
		ref                   string
		wantErr               bool
		want                  *IAPBrandIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/brands/123456789",
			want: &IAPBrandIdentity{
				Project: "my-project",
				Brand:   "123456789",
			},
			wantIdentitySpecified: true,
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://iap.googleapis.com/projects/my-project/brands/123456789",
			want: &IAPBrandIdentity{
				Project: "my-project",
				Brand:   "123456789",
			},
			wantIdentitySpecified: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IAPBrandIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestIAPBrand_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *IAPBrand
		wantErr               bool
		want                  *IAPBrandIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &IAPBrand{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAPBrandSpec{
					ResourceID: common.LazyPtr("123456789"),
				},
			},
			want: &IAPBrandIdentity{
				Project: "my-project",
				Brand:   "123456789",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID (not yet created)",
			obj: &IAPBrand{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: IAPBrandSpec{},
			},
			want: &IAPBrandIdentity{
				Project: "my-project",
				Brand:   "",
			},
			wantIdentitySpecified: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotIdentitySpecified := got.(*IAPBrandIdentity).HasIdentitySpecified(); gotIdentitySpecified != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", gotIdentitySpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}
