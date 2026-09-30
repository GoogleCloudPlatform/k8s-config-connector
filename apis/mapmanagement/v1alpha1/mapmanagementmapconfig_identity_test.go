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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestMapManagementMapConfigIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *MapManagementMapConfigIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/mapConfigs/94208ff807038e1575d64665",
			want: &MapManagementMapConfigIdentity{
				Project:   "my-project",
				MapConfig: "94208ff807038e1575d64665",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://mapmanagement.googleapis.com/projects/my-project/mapConfigs/94208ff807038e1575d64665",
			want: &MapManagementMapConfigIdentity{
				Project:   "my-project",
				MapConfig: "94208ff807038e1575d64665",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &MapManagementMapConfigIdentity{}
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

func TestMapManagementMapConfig_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).Build()
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *MapManagementMapConfig
		want    *MapManagementMapConfigIdentity
		wantErr bool
	}{
		{
			name: "server-assigned ID unset initially",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "test-project"},
				},
			},
			want: &MapManagementMapConfigIdentity{
				Project:   "test-project",
				MapConfig: "",
			},
		},
		{
			name: "user-specified resourceID initially",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "test-project"},
					ResourceID: direct.PtrTo("existing-map-config-id"),
				},
			},
			want: &MapManagementMapConfigIdentity{
				Project:   "test-project",
				MapConfig: "existing-map-config-id",
			},
		},
		{
			name: "status externalRef populated for server-assigned ID",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "test-project"},
				},
				Status: MapManagementMapConfigStatus{
					ExternalRef: direct.PtrTo("projects/test-project/mapConfigs/server-assigned-12345"),
				},
			},
			want: &MapManagementMapConfigIdentity{
				Project:   "test-project",
				MapConfig: "server-assigned-12345",
			},
		},
		{
			name: "status externalRef matching specified resourceID",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "test-project"},
					ResourceID: direct.PtrTo("server-assigned-12345"),
				},
				Status: MapManagementMapConfigStatus{
					ExternalRef: direct.PtrTo("projects/test-project/mapConfigs/server-assigned-12345"),
				},
			},
			want: &MapManagementMapConfigIdentity{
				Project:   "test-project",
				MapConfig: "server-assigned-12345",
			},
		},
		{
			name: "conflict between specified resourceID and status externalRef",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "test-project"},
					ResourceID: direct.PtrTo("different-id"),
				},
				Status: MapManagementMapConfigStatus{
					ExternalRef: direct.PtrTo("projects/test-project/mapConfigs/server-assigned-12345"),
				},
			},
			wantErr: true,
		},
		{
			name: "conflict between projectRef and status externalRef project",
			obj: &MapManagementMapConfig{
				Spec: MapManagementMapConfigSpec{
					ProjectRef: &refs.ProjectRef{External: "project-b"},
				},
				Status: MapManagementMapConfigStatus{
					ExternalRef: direct.PtrTo("projects/project-a/mapConfigs/server-assigned-12345"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, reader)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
