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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestWebSecurityScannerScanConfigIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *WebSecurityScannerScanConfigIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/scanConfigs/my-scan-config",
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "my-scan-config",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://websecurityscanner.googleapis.com/projects/my-project/scanConfigs/my-scan-config",
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "my-scan-config",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &WebSecurityScannerScanConfigIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != "projects/my-project/scanConfigs/my-scan-config" {
					t.Errorf("String() = %v, want %v", i.String(), "projects/my-project/scanConfigs/my-scan-config")
				}
				if i.ParentString() != "projects/my-project" {
					t.Errorf("ParentString() = %v, want %v", i.ParentString(), "projects/my-project")
				}
			}
		})
	}
}

func TestWebSecurityScannerScanConfig_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *WebSecurityScannerScanConfig
		wantErr               bool
		want                  *WebSecurityScannerScanConfigIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &WebSecurityScannerScanConfig{
				Spec: WebSecurityScannerScanConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					ResourceID: common.LazyPtr("my-scan-config"),
				},
			},
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "my-scan-config",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity (empty spec resourceID, defaulted from status)",
			obj: &WebSecurityScannerScanConfig{
				Spec: WebSecurityScannerScanConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
				},
				Status: WebSecurityScannerScanConfigStatus{
					ExternalRef: common.LazyPtr("projects/my-project/scanConfigs/12345"),
				},
			},
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "12345",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID and empty status (not yet created)",
			obj: &WebSecurityScannerScanConfig{
				Spec: WebSecurityScannerScanConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with matching spec resourceID and status externalRef",
			obj: &WebSecurityScannerScanConfig{
				Spec: WebSecurityScannerScanConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					ResourceID: common.LazyPtr("12345"),
				},
				Status: WebSecurityScannerScanConfigStatus{
					ExternalRef: common.LazyPtr("projects/my-project/scanConfigs/12345"),
				},
			},
			want: &WebSecurityScannerScanConfigIdentity{
				Project:    "my-project",
				ScanConfig: "12345",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity conflict with spec resourceID",
			obj: &WebSecurityScannerScanConfig{
				Spec: WebSecurityScannerScanConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					ResourceID: common.LazyPtr("conflict-id"),
				},
				Status: WebSecurityScannerScanConfigStatus{
					ExternalRef: common.LazyPtr("projects/my-project/scanConfigs/12345"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdentity, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := gotIdentity.(*WebSecurityScannerScanConfigIdentity)
				if !ok {
					t.Fatalf("returned identity is not *WebSecurityScannerScanConfigIdentity, got %T", gotIdentity)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotSpecified := got.HasIdentitySpecified(); gotSpecified != tt.wantIdentitySpecified {
					t.Errorf("got.HasIdentitySpecified() = %v, want %v", gotSpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}
