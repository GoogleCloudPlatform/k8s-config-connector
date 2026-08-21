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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestCertificateManagerTrustConfigIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CertificateManagerTrustConfigIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/trustConfigs/my-trust-config",
			want: &CertificateManagerTrustConfigIdentity{
				Project:     "my-project",
				Location:    "global",
				TrustConfig: "my-trust-config",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://certificatemanager.googleapis.com/projects/my-project/locations/global/trustConfigs/my-trust-config",
			want: &CertificateManagerTrustConfigIdentity{
				Project:     "my-project",
				Location:    "global",
				TrustConfig: "my-trust-config",
			},
		},
		{
			name: "url with domain prefix",
			ref:  "//certificatemanager.googleapis.com/projects/my-project/locations/global/trustConfigs/my-trust-config",
			want: &CertificateManagerTrustConfigIdentity{
				Project:     "my-project",
				Location:    "global",
				TrustConfig: "my-trust-config",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CertificateManagerTrustConfigIdentity{}
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

func TestCertificateManagerTrustConfigIdentity_GetIdentity(t *testing.T) {
	ctx := context.Background()

	projectObj := &unstructured.Unstructured{}
	projectObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	projectObj.SetName("my-project-name")
	projectObj.SetNamespace("my-namespace")
	if err := unstructured.SetNestedField(projectObj.Object, "my-project-id", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	s := runtime.NewScheme()
	_ = AddToScheme(s)
	s.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ProjectList",
	}, &unstructured.UnstructuredList{})

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(projectObj).Build()

	tests := []struct {
		name        string
		obj         *CertificateManagerTrustConfig
		expected    string
		expectError bool
	}{
		{
			name: "basic project and location ref",
			obj: &CertificateManagerTrustConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-trust-config",
					Namespace: "my-namespace",
				},
				Spec: CertificateManagerTrustConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location: ptr.To("global"),
				},
			},
			expected: "projects/my-project/locations/global/trustConfigs/my-trust-config",
		},
		{
			name: "project and location ref with custom resourceID",
			obj: &CertificateManagerTrustConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-trust-config",
					Namespace: "my-namespace",
				},
				Spec: CertificateManagerTrustConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location:   ptr.To("global"),
					ResourceID: func() *string { s := "custom-trust-config-id"; return &s }(),
				},
			},
			expected: "projects/my-project/locations/global/trustConfigs/custom-trust-config-id",
		},
		{
			name: "project ref missing external, using name",
			obj: &CertificateManagerTrustConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-trust-config",
					Namespace: "my-namespace",
				},
				Spec: CertificateManagerTrustConfigSpec{
					ProjectRef: &refs.ProjectRef{
						Name: "my-project-name",
					},
					Location: ptr.To("us-central1"),
				},
			},
			expected: "projects/my-project-id/locations/us-central1/trustConfigs/my-trust-config",
		},
		{
			name: "status mismatch",
			obj: &CertificateManagerTrustConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-trust-config",
					Namespace: "my-namespace",
				},
				Spec: CertificateManagerTrustConfigSpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location: ptr.To("global"),
				},
				Status: CertificateManagerTrustConfigStatus{
					ExternalRef: func() *string {
						s := "projects/other-project/locations/global/trustConfigs/my-trust-config"
						return &s
					}(),
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, reader)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.String() != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, id.String())
			}
		})
	}
}

func TestCertificateManagerTrustConfigRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	config := &CertificateManagerTrustConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-trust-config",
			Namespace: "my-ns",
		},
		Spec: CertificateManagerTrustConfigSpec{
			ProjectRef: &refs.ProjectRef{
				External: "my-project",
			},
			Location:   ptr.To("global"),
			ResourceID: func() *string { s := "my-trust-config-id"; return &s }(),
		},
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(config).Build()

	tests := []struct {
		name             string
		ref              *CertificateManagerTrustConfigRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &CertificateManagerTrustConfigRef{
				External: "projects/my-project/locations/global/trustConfigs/my-trust-config",
			},
			want: "projects/my-project/locations/global/trustConfigs/my-trust-config",
		},
		{
			name: "internal reference",
			ref: &CertificateManagerTrustConfigRef{
				Name:      "my-trust-config",
				Namespace: "my-ns",
			},
			want: "projects/my-project/locations/global/trustConfigs/my-trust-config-id",
		},
		{
			name: "internal reference with default namespace",
			ref: &CertificateManagerTrustConfigRef{
				Name: "my-trust-config",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/locations/global/trustConfigs/my-trust-config-id",
		},
		{
			name: "internal reference not found",
			ref: &CertificateManagerTrustConfigRef{
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

func TestCertificateManagerTrustConfigIdentityInterfaces(t *testing.T) {
	var _ identity.IdentityV2 = &CertificateManagerTrustConfigIdentity{}
	var _ identity.Resource = &CertificateManagerTrustConfig{}
	var _ refs.Ref = &CertificateManagerTrustConfigRef{}
}
