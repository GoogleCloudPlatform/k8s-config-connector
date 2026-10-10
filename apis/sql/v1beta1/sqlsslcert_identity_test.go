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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSQLSSLCertIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                  string
		ref                   string
		wantErr               bool
		want                  *SQLSSLCertIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "valid canonical reference",
			ref:  "projects/my-project/instances/my-instance/sslCerts/abc123sha1",
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "abc123sha1",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "full url with sqladmin host and api version",
			ref:  "https://sqladmin.googleapis.com/sql/v1beta4/projects/my-project/instances/my-instance/sslCerts/abc123sha1",
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "abc123sha1",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "full url with sqladmin host without api version",
			ref:  "https://sqladmin.googleapis.com/projects/my-project/instances/my-instance/sslCerts/abc123sha1",
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "abc123sha1",
			},
			wantIdentitySpecified: true,
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty sha1 segment",
			ref:     "projects/my-project/instances/my-instance/sslCerts/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SQLSSLCertIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
				if gotHost := i.Host(); gotHost != "sqladmin.googleapis.com" {
					t.Errorf("Host() = %v, want %v", gotHost, "sqladmin.googleapis.com")
				}
				if gotParent := i.ParentString(); gotParent != "projects/my-project/instances/my-instance" {
					t.Errorf("ParentString() = %v, want %v", gotParent, "projects/my-project/instances/my-instance")
				}
				if gotStr := i.String(); gotStr != "projects/my-project/instances/my-instance/sslCerts/abc123sha1" {
					t.Errorf("String() = %v, want %v", gotStr, "projects/my-project/instances/my-instance/sslCerts/abc123sha1")
				}
			}
		})
	}
}

func TestSQLSSLCert_GetIdentity(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	sqlInstance := &unstructured.Unstructured{}
	sqlInstance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sql.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SQLInstance",
	})
	sqlInstance.SetName("my-instance")
	sqlInstance.SetNamespace("test-ns")
	sqlInstance.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(sqlInstance).Build()

	tests := []struct {
		name                  string
		obj                   *SQLSSLCert
		wantErr               bool
		want                  *SQLSSLCertIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with external instanceRef and specified resourceID",
			obj: &SQLSSLCert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
					Name:      "test-cert",
				},
				Spec: SQLSSLCertSpec{
					CommonName: "client",
					InstanceRef: refs.SQLInstanceRef{
						External: "projects/my-project/instances/my-instance",
					},
					ResourceID: common.LazyPtr("abc123sha1"),
				},
			},
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "abc123sha1",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with name instanceRef and server-generated identity from status",
			obj: &SQLSSLCert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
					Name:      "test-cert",
				},
				Spec: SQLSSLCertSpec{
					CommonName: "client",
					InstanceRef: refs.SQLInstanceRef{
						Name: "my-instance",
					},
				},
				Status: SQLSSLCertStatus{
					Sha1Fingerprint: common.LazyPtr("server-generated-sha1"),
				},
			},
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "server-generated-sha1",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID and empty status (not yet created)",
			obj: &SQLSSLCert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
					Name:      "test-cert",
				},
				Spec: SQLSSLCertSpec{
					CommonName: "client",
					InstanceRef: refs.SQLInstanceRef{
						External: "projects/my-project/instances/my-instance",
					},
				},
			},
			want: &SQLSSLCertIdentity{
				Project:         "my-project",
				Instance:        "my-instance",
				Sha1Fingerprint: "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with conflict between spec resourceID and status sha1Fingerprint",
			obj: &SQLSSLCert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
					Name:      "test-cert",
				},
				Spec: SQLSSLCertSpec{
					CommonName: "client",
					InstanceRef: refs.SQLInstanceRef{
						External: "projects/my-project/instances/my-instance",
					},
					ResourceID: common.LazyPtr("spec-sha1"),
				},
				Status: SQLSSLCertStatus{
					Sha1Fingerprint: common.LazyPtr("status-sha1"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity missing instanceRef",
			obj: &SQLSSLCert{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "test-ns",
					Name:      "test-cert",
				},
				Spec: SQLSSLCertSpec{
					CommonName: "client",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdentity, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := gotIdentity.(*SQLSSLCertIdentity)
				if !ok {
					t.Fatalf("returned identity is not *SQLSSLCertIdentity, got %T", gotIdentity)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotSpecified := got.HasIdentitySpecified(); gotSpecified != tt.wantIdentitySpecified {
					t.Errorf("got.HasIdentitySpecified() = %v, want %v", gotSpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestSQLSSLCertRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)

	sqlCertReady := &SQLSSLCert{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "cert-ready",
		},
		Spec: SQLSSLCertSpec{
			CommonName: "client",
			InstanceRef: refs.SQLInstanceRef{
				External: "projects/my-project/instances/my-instance",
			},
		},
		Status: SQLSSLCertStatus{
			Conditions: []v1alpha1.Condition{
				{
					Type:   "Ready",
					Status: corev1.ConditionTrue,
				},
			},
			Sha1Fingerprint: common.LazyPtr("cert-ready-sha1"),
		},
	}

	sqlCertNotReady := &SQLSSLCert{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "cert-not-ready",
		},
		Spec: SQLSSLCertSpec{
			CommonName: "client",
			InstanceRef: refs.SQLInstanceRef{
				External: "projects/my-project/instances/my-instance",
			},
		},
		Status: SQLSSLCertStatus{
			Conditions: []v1alpha1.Condition{
				{
					Type:   "Ready",
					Status: corev1.ConditionFalse,
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(sqlCertReady, sqlCertNotReady).Build()

	t.Run("Normalize with external reference", func(t *testing.T) {
		ref := &SQLSSLCertRef{
			External: "projects/my-project/instances/my-instance/sslCerts/abc123sha1",
		}
		if err := ref.Normalize(ctx, fakeClient, "test-ns"); err != nil {
			t.Fatalf("Normalize() error = %v", err)
		}
		if want := "projects/my-project/instances/my-instance/sslCerts/abc123sha1"; ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}
	})

	t.Run("Normalize with name reference to ready resource", func(t *testing.T) {
		ref := &SQLSSLCertRef{
			Name: "cert-ready",
		}
		if err := ref.Normalize(ctx, fakeClient, "test-ns"); err != nil {
			t.Fatalf("Normalize() error = %v", err)
		}
		if want := "projects/my-project/instances/my-instance/sslCerts/cert-ready-sha1"; ref.External != want {
			t.Errorf("ref.External = %q, want %q", ref.External, want)
		}
	})

	t.Run("Normalize with name reference to not-ready resource", func(t *testing.T) {
		ref := &SQLSSLCertRef{
			Name: "cert-not-ready",
		}
		if err := ref.Normalize(ctx, fakeClient, "test-ns"); err == nil {
			t.Fatalf("Normalize() expected error for not-ready resource, got nil")
		}
	})

	t.Run("ParseExternalToIdentity", func(t *testing.T) {
		ref := &SQLSSLCertRef{
			External: "projects/my-project/instances/my-instance/sslCerts/abc123sha1",
		}
		id, err := ref.ParseExternalToIdentity()
		if err != nil {
			t.Fatalf("ParseExternalToIdentity() error = %v", err)
		}
		want := &SQLSSLCertIdentity{
			Project:         "my-project",
			Instance:        "my-instance",
			Sha1Fingerprint: "abc123sha1",
		}
		if diff := cmp.Diff(want, id); diff != "" {
			t.Errorf("ParseExternalToIdentity() mismatch (-want +got):\n%s", diff)
		}
	})
}
