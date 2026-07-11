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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestPrivateCACertificateIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *PrivateCACertificateIdentity
		wantErr bool
	}{
		{
			name: "valid certificate reference",
			ref:  "projects/test-project/locations/us-central1/caPools/test-pool/certificates/test-cert",
			want: &PrivateCACertificateIdentity{
				Project:     "test-project",
				Location:    "us-central1",
				CAPool:      "test-pool",
				Certificate: "test-cert",
			},
			wantErr: false,
		},
		{
			name:    "invalid format - extra segment",
			ref:     "projects/test-project/locations/us-central1/caPools/test-pool/certificates/test-cert/extra",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/test-project/caPools/test-pool/certificates/test-cert",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &PrivateCACertificateIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("PrivateCACertificateIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("PrivateCACertificateIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestPrivateCACertificateRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/us-central1/caPools/my-pool/certificates/my-cert",
			wantErr:  false,
		},
		{
			name:     "invalid external",
			external: "projects/my-project/locations/us-central1/certificates/my-cert",
			wantErr:  true,
		},
		{
			name:     "empty external",
			external: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PrivateCACertificateRef{}
			if err := r.ValidateExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("PrivateCACertificateRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrivateCACertificateRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *PrivateCACertificateRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &PrivateCACertificateRef{
				External: "projects/test-project/locations/us-central1/caPools/test-pool/certificates/test-cert",
			},
			wantExternal: "projects/test-project/locations/us-central1/caPools/test-pool/certificates/test-cert",
		},
		{
			name: "external with invalid format",
			ref: &PrivateCACertificateRef{
				External: "invalid-format",
			},
			wantErr: `format of PrivateCACertificate external="invalid-format" was not known (use projects/{project}/locations/{location}/caPools/{caPool}/certificates/{certificate})`,
		},
		{
			name: "name specified, with true Ready condition",
			ref: &PrivateCACertificateRef{
				Name:      "test-cert",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "privateca.cnrm.cloud.google.com/v1beta1",
						"kind":       "PrivateCACertificate",
						"metadata": map[string]interface{}{
							"name":      "test-cert",
							"namespace": "my-namespace",
						},
						"spec": map[string]interface{}{
							"location":   "us-central1",
							"projectRef": map[string]interface{}{"external": "projects/test-project"},
							"caPoolRef":  map[string]interface{}{"external": "projects/test-project/locations/us-central1/caPools/test-pool"},
							"resourceID": "test-cert",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": "True",
								},
							},
						},
					},
				},
			},
			wantExternal: "projects/test-project/locations/us-central1/caPools/test-pool/certificates/test-cert",
		},
		{
			name: "name specified, but ready condition is False",
			ref: &PrivateCACertificateRef{
				Name:      "test-cert",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "privateca.cnrm.cloud.google.com/v1beta1",
						"kind":       "PrivateCACertificate",
						"metadata": map[string]interface{}{
							"name":      "test-cert",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": "False",
								},
							},
						},
					},
				},
			},
			wantErr: `reference PrivateCACertificate my-namespace/test-cert is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &PrivateCACertificateRef{
				Name:      "test-cert",
				Namespace: "my-namespace",
			},
			wantErr: `reference PrivateCACertificate my-namespace/test-cert is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(PrivateCACertificateGVK.GroupVersion(), &unstructured.Unstructured{})
			s.AddKnownTypes(PrivateCACertificateGVK.GroupVersion(), &PrivateCACertificate{})
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

func TestPrivateCACertificateImplementsIdentity(t *testing.T) {
	var _ identity.Identity = &PrivateCACertificateIdentity{}
	var _ identity.IdentityV2 = &PrivateCACertificateIdentity{}
}
