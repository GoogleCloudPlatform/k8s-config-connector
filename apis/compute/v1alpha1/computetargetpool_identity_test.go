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
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeTargetPoolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeTargetPoolIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/targetPools/my-tp",
			want: &ComputeTargetPoolIdentity{
				Project:    "my-project",
				Region:     "us-central1",
				TargetPool: "my-tp",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/regions/us-central1/targetPools/my-tp",
			want: &ComputeTargetPoolIdentity{
				Project:    "my-project",
				Region:     "us-central1",
				TargetPool: "my-tp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeTargetPoolIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeTargetPoolRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/regions/us-central1/targetPools/my-tp",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/regions/us-central1/targetPools/my-tp",
			wantErr: true,
		},
		{
			name:    "missing region",
			ref:     "projects/my-project/targetPools/my-tp",
			wantErr: true,
		},
		{
			name:    "missing targetPool",
			ref:     "projects/my-project/regions/us-central1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeTargetPoolRef{}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeTargetPool_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add to scheme: %v", err)
	}

	tests := []struct {
		name    string
		obj     *ComputeTargetPool
		wantErr bool
		want    string
	}{
		{
			name: "valid object specs",
			obj: &ComputeTargetPool{
				Spec: ComputeTargetPoolSpec{
					Location:   "us-central1",
					ResourceID: ptrTo("my-tp"),
				},
			},
			want: "projects/test-project/regions/us-central1/targetPools/my-tp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			tt.obj.SetNamespace("test-namespace")
			tt.obj.SetName("my-tp")
			tt.obj.SetAnnotations(map[string]string{
				"cnrm.cloud.google.com/project-id": "test-project",
			})

			cl := fake.NewClientBuilder().WithScheme(scheme).Build()

			got, err := tt.obj.GetIdentity(ctx, cl)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if got.String() != tt.want {
					t.Errorf("GetIdentity() got = %q, want %q", got.String(), tt.want)
				}
			}
		})
	}
}

func ptrTo[T any](v T) *T {
	return &v
}
