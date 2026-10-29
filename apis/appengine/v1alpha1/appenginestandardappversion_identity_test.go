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
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestAppEngineStandardAppVersion_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *AppEngineStandardAppVersion
		wantErr bool
		want    *AppEngineStandardAppVersionIdentity
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &AppEngineStandardAppVersion{
				Spec: AppEngineStandardAppVersionSpec{
					ServiceRef: &AppEngineServiceRef{
						External: "apps/my-app/services/my-service",
					},
					ResourceID: common.LazyPtr("v1"),
				},
			},
			want: &AppEngineStandardAppVersionIdentity{
				App:     "my-app",
				Service: "my-service",
				Version: "v1",
			},
		},
		{
			name: "GetIdentity defaulting resourceID to metadata.name when unset",
			obj: &AppEngineStandardAppVersion{
				ObjectMeta: metav1.ObjectMeta{
					Name: "defaulted-v1",
				},
				Spec: AppEngineStandardAppVersionSpec{
					ServiceRef: &AppEngineServiceRef{
						External: "apps/my-app/services/my-service",
					},
				},
			},
			want: &AppEngineStandardAppVersionIdentity{
				App:     "my-app",
				Service: "my-service",
				Version: "defaulted-v1",
			},
		},
		{
			name: "GetIdentity error on identity change (spec resourceID does not match status.externalRef)",
			obj: &AppEngineStandardAppVersion{
				Spec: AppEngineStandardAppVersionSpec{
					ServiceRef: &AppEngineServiceRef{
						External: "apps/my-app/services/my-service",
					},
					ResourceID: common.LazyPtr("v1"),
				},
				Status: AppEngineStandardAppVersionStatus{
					ExternalRef: common.LazyPtr("apps/my-app/services/my-service/versions/v2"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity error when serviceRef is nil",
			obj: &AppEngineStandardAppVersion{
				Spec: AppEngineStandardAppVersionSpec{
					ResourceID: common.LazyPtr("v1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, fakeClient)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			typedID, ok := id.(*AppEngineStandardAppVersionIdentity)
			if !ok {
				t.Fatalf("expected *AppEngineStandardAppVersionIdentity, got %T", id)
			}

			if diff := cmp.Diff(tc.want, typedID); diff != "" {
				t.Errorf("GetIdentity mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
