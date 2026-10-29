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
	"strings"
	"testing"

	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeManagedSSLCertificateIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeManagedSSLCertificateIdentity
	}{
		{
			name: "valid global reference",
			ref:  "projects/my-project/global/sslCertificates/my-cert",
			want: &ComputeManagedSSLCertificateIdentity{
				Project:        "my-project",
				SslCertificate: "my-cert",
			},
		},
		{
			name: "valid full url",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project/global/sslCertificates/my-cert",
			want: &ComputeManagedSSLCertificateIdentity{
				Project:        "my-project",
				SslCertificate: "my-cert",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeManagedSSLCertificateIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != "projects/my-project/global/sslCertificates/my-cert" {
					t.Errorf("String() = %q, want %q", i.String(), "projects/my-project/global/sslCertificates/my-cert")
				}
				if i.Host() != "compute.googleapis.com" {
					t.Errorf("Host() = %q, want %q", i.Host(), "compute.googleapis.com")
				}
				if i.ParentString() != "projects/my-project/global" {
					t.Errorf("ParentString() = %q, want %q", i.ParentString(), "projects/my-project/global")
				}
			}
		})
	}
}

func TestComputeManagedSSLCertificateRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid external reference",
			ref:     "projects/my-project/global/sslCertificates/my-cert",
			wantErr: false,
		},
		{
			name:    "valid full url",
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
			r := &ComputeManagedSSLCertificateRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeManagedSSLCertificate_GetIdentity(t *testing.T) {
	ctx := context.Background()
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	client := fake.NewClientBuilder().WithScheme(s).Build()

	obj := &ComputeManagedSSLCertificate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cert",
			Namespace: "test-ns",
		},
		Spec: ComputeManagedSSLCertificateSpec{
			ProjectRef: apirefs.ProjectRef{
				External: "test-project",
			},
			ResourceID: ptr.To("test-cert"),
		},
	}

	// 1. Nominal spec identity
	id, err := obj.GetIdentity(ctx, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := &ComputeManagedSSLCertificateIdentity{
		Project:        "test-project",
		SslCertificate: "test-cert",
	}
	if diff := cmp.Diff(want, id); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}

	// 2. Status.ExternalRef matches spec
	obj.Status.ExternalRef = ptr.To("projects/test-project/global/sslCertificates/test-cert")
	id, err = obj.GetIdentity(ctx, client)
	if err != nil {
		t.Fatalf("unexpected error when status matches spec: %v", err)
	}
	if diff := cmp.Diff(want, id); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}

	// 3. Status.ExternalRef drifts/mismatches spec
	obj.Status.ExternalRef = ptr.To("projects/test-project/global/sslCertificates/other-cert")
	_, err = obj.GetIdentity(ctx, client)
	if err == nil || !strings.Contains(err.Error(), "cannot change ComputeManagedSSLCertificate identity") {
		t.Fatalf("expected identity change error, got %v", err)
	}

	// 4. Status.ObservedState.SelfLink matches spec
	obj.Status.ExternalRef = nil
	obj.Status.ObservedState = &ComputeManagedSSLCertificateObservedState{
		SelfLink: ptr.To("https://www.googleapis.com/compute/v1/projects/test-project/global/sslCertificates/test-cert"),
	}
	id, err = obj.GetIdentity(ctx, client)
	if err != nil {
		t.Fatalf("unexpected error when observedState selfLink matches spec: %v", err)
	}
	if diff := cmp.Diff(want, id); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeManagedSSLCertificateRef_Normalize(t *testing.T) {
	ctx := context.Background()
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	uCert := &unstructured.Unstructured{}
	uCert.SetGroupVersionKind(ComputeManagedSSLCertificateGVK)
	uCert.SetName("my-cert")
	uCert.SetNamespace("my-ns")
	uCert.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	uCert.Object["spec"] = map[string]interface{}{
		"resourceID": "my-cert",
		"projectRef": map[string]interface{}{
			"external": "my-project",
		},
	}
	uCert.Object["status"] = map[string]interface{}{
		"observedState": map[string]interface{}{
			"selfLink": "https://www.googleapis.com/compute/v1/projects/my-project/global/sslCertificates/my-cert",
		},
	}

	client := fake.NewClientBuilder().WithScheme(s).WithObjects(uCert).Build()

	r := &ComputeManagedSSLCertificateRef{
		Name:      "my-cert",
		Namespace: "my-ns",
	}

	if err := r.Normalize(ctx, client, "my-ns"); err != nil {
		t.Fatalf("Normalize() error = %v", err)
	}

	wantExternal := "projects/my-project/global/sslCertificates/my-cert"
	if r.External != wantExternal {
		t.Errorf("Normalize() External = %q, want %q", r.External, wantExternal)
	}
}
