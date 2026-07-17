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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeBackendServiceSignedURLKeyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeBackendServiceSignedURLKeyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/global/backendServices/my-backend/signedUrlKeys/my-key",
			want: &ComputeBackendServiceSignedURLKeyIdentity{
				Project:        "my-project",
				BackendService: "my-backend",
				SignedUrlKey:   "my-key",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/global/backendServices/my-backend/signedUrlKeys/my-key",
			want: &ComputeBackendServiceSignedURLKeyIdentity{
				Project:        "my-project",
				BackendService: "my-backend",
				SignedUrlKey:   "my-key",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeBackendServiceSignedURLKeyIdentity{}
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

func TestComputeBackendServiceSignedURLKeyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/global/backendServices/my-backend/signedUrlKeys/my-key",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/global/backendServices/my-backend/signedUrlKeys/my-key",
			wantErr: true,
		},
		{
			name:    "missing backend service",
			ref:     "projects/my-project/global/signedUrlKeys/my-key",
			wantErr: true,
		},
		{
			name:    "missing key name",
			ref:     "projects/my-project/global/backendServices/my-backend/signedUrlKeys/",
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
			r := &ComputeBackendServiceSignedURLKeyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeBackendServiceSignedURLKeyRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeBackendServiceSignedURLKeyRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeBackendServiceSignedURLKeyRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeBackendServiceSignedURLKeyRef{
				External: "projects/test-project/global/backendServices/test-backend/signedUrlKeys/test-key",
			},
			wantExternal: "projects/test-project/global/backendServices/test-backend/signedUrlKeys/test-key",
		},
		{
			name: "external with invalid format",
			ref: &ComputeBackendServiceSignedURLKeyRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeBackendServiceSignedURLKey external="invalid-format" was not known (use projects/{project}/global/backendServices/{backendService}/signedUrlKeys/{signedUrlKey})`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeBackendServiceSignedURLKey{})
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

func TestParseBackendServiceID(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     string
		wantErr  bool
	}{
		{
			name:     "global backend service full url",
			external: "https://www.googleapis.com/compute/v1/projects/my-project/global/backendServices/my-backend-service",
			want:     "my-backend-service",
		},
		{
			name:     "regional backend service full url",
			external: "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/backendServices/my-backend-service",
			want:     "my-backend-service",
		},
		{
			name:     "global backend service path",
			external: "projects/my-project/global/backendServices/my-backend-service",
			want:     "my-backend-service",
		},
		{
			name:     "short name",
			external: "my-backend-service",
			want:     "my-backend-service",
		},
		{
			name:     "invalid format with slash",
			external: "invalid/my-backend-service",
			wantErr:  true,
		},
		{
			name:     "empty external reference",
			external: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBackendServiceID(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseBackendServiceID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if got != tt.want {
					t.Errorf("parseBackendServiceID() got = %q, want %q", got, tt.want)
				}
			}
		})
	}
}
