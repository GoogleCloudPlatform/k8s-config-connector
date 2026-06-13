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
// See the License for the_identity.go specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeTargetGRPCProxyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeTargetGRPCProxyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/global/targetGrpcProxies/my-proxy",
			want: &ComputeTargetGRPCProxyIdentity{
				Project:         "my-project",
				TargetGrpcProxy: "my-proxy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/global/targetGrpcProxies/my-proxy",
			want: &ComputeTargetGRPCProxyIdentity{
				Project:         "my-project",
				TargetGrpcProxy: "my-proxy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeTargetGRPCProxyIdentity{}
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

func TestComputeTargetGRPCProxyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/global/targetGrpcProxies/my-proxy",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/global/targetGrpcProxies/my-proxy",
			wantErr: true,
		},
		{
			name:    "missing proxy",
			ref:     "projects/my-project/global/targetGrpcProxies",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeTargetGRPCProxyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeTargetGRPCProxyRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeTargetGRPCProxyRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeTargetGRPCProxyRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeTargetGRPCProxyRef{
				External: "projects/test-project/global/targetGrpcProxies/test-proxy",
			},
			wantExternal: "projects/test-project/global/targetGrpcProxies/test-proxy",
		},
		{
			name: "external with invalid format",
			ref: &ComputeTargetGRPCProxyRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeTargetGRPCProxy external="invalid-format" was not known (use projects/{project}/global/targetGrpcProxies/{targetGrpcProxy})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeTargetGRPCProxyRef{
				Name:      "test-proxy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeTargetGRPCProxy",
						"metadata": map[string]interface{}{
							"name":      "test-proxy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/global/targetGrpcProxies/test-proxy",
						},
					},
				},
			},
			wantExternal: "projects/test-project/global/targetGrpcProxies/test-proxy",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeTargetGRPCProxyRef{
				Name:      "test-proxy",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeTargetGRPCProxy",
						"metadata": map[string]interface{}{
							"name":      "test-proxy",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeTargetGRPCProxy my-namespace/test-proxy is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeTargetGRPCProxyRef{
				Name:      "test-proxy",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeTargetGRPCProxy my-namespace/test-proxy is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeTargetGRPCProxy{})
			cl := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), cl, "default")
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("Normalize() expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Errorf("Normalize() error = %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Normalize() unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}
