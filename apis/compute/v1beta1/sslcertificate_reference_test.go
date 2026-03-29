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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeSSLCertificateRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid external reference (projects/)",
			ref:     "projects/my-project/global/sslCertificates/my-cert",
			wantErr: false,
		},
		{
			name:    "valid external reference (https://)",
			ref:     "https://www.googleapis.com/compute/v1/projects/my-project/global/sslCertificates/my-cert",
			wantErr: false,
		},
		{
			name:    "invalid external reference",
			ref:     "my-cert",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeSSLCertificateRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeSSLCertificateRef_NormalizedExternal(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	cert := &unstructured.Unstructured{}
	cert.SetGroupVersionKind(ComputeSSLCertificateGVK)
	cert.SetName("my-cert")
	cert.SetNamespace("my-ns")
	cert.Object["status"] = map[string]interface{}{
		"selfLink": "https://www.googleapis.com/compute/v1/projects/my-project/global/sslCertificates/my-cert",
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(cert).Build()

	tests := []struct {
		name           string
		ref            *ComputeSSLCertificateRef
		otherNamespace string
		want           string
		wantErr        bool
	}{
		{
			name: "external reference",
			ref: &ComputeSSLCertificateRef{
				External: "projects/my-project/global/sslCertificates/my-cert",
			},
			want: "projects/my-project/global/sslCertificates/my-cert",
		},
		{
			name: "internal reference",
			ref: &ComputeSSLCertificateRef{
				Name:      "my-cert",
				Namespace: "my-ns",
			},
			want: "https://www.googleapis.com/compute/v1/projects/my-project/global/sslCertificates/my-cert",
		},
		{
			name: "both name and external",
			ref: &ComputeSSLCertificateRef{
				Name:     "my-cert",
				External: "projects/my-project/global/sslCertificates/my-cert",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ref.NormalizedExternal(context.Background(), reader, tt.otherNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizedExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizedExternal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
