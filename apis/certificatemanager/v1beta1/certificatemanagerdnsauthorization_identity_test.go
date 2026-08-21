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

func TestCertificateManagerDNSAuthorizationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CertificateManagerDNSAuthorizationIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/dnsAuthorizations/my-dnsauth",
			want: &CertificateManagerDNSAuthorizationIdentity{
				Project:          "my-project",
				Location:         "global",
				DNSAuthorization: "my-dnsauth",
			},
		},
		{
			name: "valid reference with regional location",
			ref:  "projects/my-project/locations/us-central1/dnsAuthorizations/my-dnsauth",
			want: &CertificateManagerDNSAuthorizationIdentity{
				Project:          "my-project",
				Location:         "us-central1",
				DNSAuthorization: "my-dnsauth",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://certificatemanager.googleapis.com/projects/my-project/locations/global/dnsAuthorizations/my-dnsauth",
			want: &CertificateManagerDNSAuthorizationIdentity{
				Project:          "my-project",
				Location:         "global",
				DNSAuthorization: "my-dnsauth",
			},
		},
		{
			name: "url with domain prefix",
			ref:  "//certificatemanager.googleapis.com/projects/my-project/locations/global/dnsAuthorizations/my-dnsauth",
			want: &CertificateManagerDNSAuthorizationIdentity{
				Project:          "my-project",
				Location:         "global",
				DNSAuthorization: "my-dnsauth",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CertificateManagerDNSAuthorizationIdentity{}
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

func TestCertificateManagerDNSAuthorizationRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	dnsAuth := &unstructured.Unstructured{}
	dnsAuth.SetGroupVersionKind(CertificateManagerDNSAuthorizationGVK)
	dnsAuth.SetName("my-dnsauth")
	dnsAuth.SetNamespace("my-ns")
	dnsAuth.Object["spec"] = map[string]interface{}{
		"resourceID": "my-dnsauth-id",
		"location":   "global",
		"projectRef": map[string]interface{}{
			"external": "my-project",
		},
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(dnsAuth).Build()

	tests := []struct {
		name             string
		ref              *CertificateManagerDNSAuthorizationRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &CertificateManagerDNSAuthorizationRef{
				External: "projects/my-project/locations/global/dnsAuthorizations/my-dnsauth",
			},
			want: "projects/my-project/locations/global/dnsAuthorizations/my-dnsauth",
		},
		{
			name: "internal reference",
			ref: &CertificateManagerDNSAuthorizationRef{
				Name:      "my-dnsauth",
				Namespace: "my-ns",
			},
			want: "projects/my-project/locations/global/dnsAuthorizations/my-dnsauth-id",
		},
		{
			name: "internal reference with default namespace",
			ref: &CertificateManagerDNSAuthorizationRef{
				Name: "my-dnsauth",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/locations/global/dnsAuthorizations/my-dnsauth-id",
		},
		{
			name: "internal reference not found",
			ref: &CertificateManagerDNSAuthorizationRef{
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
