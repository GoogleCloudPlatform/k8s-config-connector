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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCertificateManagerCertificateMapIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CertificateManagerCertificateMapIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/certificateMaps/my-certmap",
			want: &CertificateManagerCertificateMapIdentity{
				Project:        "my-project",
				CertificateMap: "my-certmap",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://certificatemanager.googleapis.com/projects/my-project/locations/global/certificateMaps/my-certmap",
			want: &CertificateManagerCertificateMapIdentity{
				Project:        "my-project",
				CertificateMap: "my-certmap",
			},
		},
		{
			name: "url with domain prefix",
			ref:  "//certificatemanager.googleapis.com/projects/my-project/locations/global/certificateMaps/my-certmap",
			want: &CertificateManagerCertificateMapIdentity{
				Project:        "my-project",
				CertificateMap: "my-certmap",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CertificateManagerCertificateMapIdentity{}
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

func TestCertificateManagerCertificateMapRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	certMap := &unstructured.Unstructured{}
	certMap.SetGroupVersionKind(CertificateManagerCertificateMapGVK)
	certMap.SetName("my-certmap")
	certMap.SetNamespace("my-ns")
	certMap.Object["spec"] = map[string]interface{}{
		"resourceID": "my-certmap-id",
		"projectRef": map[string]interface{}{
			"external": "my-project",
		},
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(certMap).Build()

	tests := []struct {
		name             string
		ref              *CertificateManagerCertificateMapRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &CertificateManagerCertificateMapRef{
				External: "projects/my-project/locations/global/certificateMaps/my-certmap",
			},
			want: "projects/my-project/locations/global/certificateMaps/my-certmap",
		},
		{
			name: "internal reference",
			ref: &CertificateManagerCertificateMapRef{
				Name:      "my-certmap",
				Namespace: "my-ns",
			},
			want: "projects/my-project/locations/global/certificateMaps/my-certmap-id",
		},
		{
			name: "internal reference with default namespace",
			ref: &CertificateManagerCertificateMapRef{
				Name: "my-certmap",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/locations/global/certificateMaps/my-certmap-id",
		},
		{
			name: "internal reference not found",
			ref: &CertificateManagerCertificateMapRef{
				Name:      "non-existent",
				Namespace: "my-ns",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), reader, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.ref.External != tt.want {
					t.Errorf("Normalize() got = %v, want %v", tt.ref.External, tt.want)
				}
			}
		})
	}
}
