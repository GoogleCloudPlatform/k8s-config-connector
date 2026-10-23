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

package v1alpha1_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestApigeeAddonsConfigIdentity_Conformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[v1alpha1.ApigeeAddonsConfigIdentity]{
		SampleExternal: "organizations/test-org/addonsConfig",
		WantHost:       "apigee.googleapis.com",
		WantParent:     "organizations/test-org",
		WantIdentity: &v1alpha1.ApigeeAddonsConfigIdentity{
			Organization: "test-org",
		},
		Ref: &v1alpha1.ApigeeAddonsConfigRef{},
	})
}

func TestApigeeAddonsConfigIdentity_GetIdentity(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(scheme)
	_ = corev1.AddToScheme(scheme)
	reader := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *v1alpha1.ApigeeAddonsConfig
		want    string
		wantErr bool
	}{
		{
			name: "from spec.org",
			obj: &v1alpha1.ApigeeAddonsConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-addons",
					Namespace: "test-ns",
				},
				Spec: v1alpha1.ApigeeAddonsConfigSpec{
					Org: "test-org",
				},
			},
			want: "organizations/test-org/addonsConfig",
		},
		{
			name: "with organizations/ prefix in spec.org",
			obj: &v1alpha1.ApigeeAddonsConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-addons",
					Namespace: "test-ns",
				},
				Spec: v1alpha1.ApigeeAddonsConfigSpec{
					Org: "organizations/test-org",
				},
			},
			want: "organizations/test-org/addonsConfig",
		},
		{
			name: "missing spec.org",
			obj: &v1alpha1.ApigeeAddonsConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-addons",
					Namespace: "test-ns",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, reader)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := id.String(); got != tc.want {
				t.Errorf("GetIdentity mismatch (-want +got):\n%s", cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestApigeeAddonsConfigIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *v1alpha1.ApigeeAddonsConfigIdentity
		wantErr bool
	}{
		{
			name: "canonical external ref",
			ref:  "organizations/my-org/addonsConfig",
			want: &v1alpha1.ApigeeAddonsConfigIdentity{
				Organization: "my-org",
			},
		},
		{
			name: "canonical external ref with host",
			ref:  "//apigee.googleapis.com/organizations/my-org/addonsConfig",
			want: &v1alpha1.ApigeeAddonsConfigIdentity{
				Organization: "my-org",
			},
		},
		{
			name:    "invalid format - wrong resource type",
			ref:     "organizations/my-org/environments/env1",
			wantErr: true,
		},
		{
			name:    "invalid format - missing addonsConfig suffix",
			ref:     "organizations/my-org",
			wantErr: true,
		},
		{
			name:    "invalid format - empty organization",
			ref:     "organizations//addonsConfig",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &v1alpha1.ApigeeAddonsConfigIdentity{}
			err := id.FromExternal(tc.ref)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, id); diff != "" {
				t.Errorf("FromExternal mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestApigeeAddonsConfigRef_Normalize(t *testing.T) {
	ctx := context.TODO()

	t.Run("external reference provided", func(t *testing.T) {
		reader := fake.NewClientBuilder().Build()
		ref := &v1alpha1.ApigeeAddonsConfigRef{
			External: "organizations/my-org/addonsConfig",
		}
		if err := ref.Normalize(ctx, reader, "default"); err != nil {
			t.Fatalf("Normalize failed: %v", err)
		}
		want := "organizations/my-org/addonsConfig"
		if got := ref.External; got != want {
			t.Errorf("Normalize mismatch (-want +got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("internal reference to ready resource", func(t *testing.T) {
		u := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apigee.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ApigeeAddonsConfig",
				"metadata": map[string]interface{}{
					"name":      "my-addons-config",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"org": "my-org",
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
		}

		scheme := runtime.NewScheme()
		_ = v1alpha1.AddToScheme(scheme)
		_ = corev1.AddToScheme(scheme)
		reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(u).Build()

		ref := &v1alpha1.ApigeeAddonsConfigRef{
			Name: "my-addons-config",
		}
		if err := ref.Normalize(ctx, reader, "default"); err != nil {
			t.Fatalf("Normalize failed: %v", err)
		}
		want := "organizations/my-org/addonsConfig"
		if got := ref.External; got != want {
			t.Errorf("Normalize mismatch (-want +got):\n%s", cmp.Diff(want, got))
		}
	})

	t.Run("internal reference to not ready resource", func(t *testing.T) {
		u := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apigee.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ApigeeAddonsConfig",
				"metadata": map[string]interface{}{
					"name":      "my-addons-config",
					"namespace": "default",
				},
				"spec": map[string]interface{}{
					"org": "my-org",
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
		}

		scheme := runtime.NewScheme()
		_ = v1alpha1.AddToScheme(scheme)
		_ = corev1.AddToScheme(scheme)
		reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(u).Build()

		ref := &v1alpha1.ApigeeAddonsConfigRef{
			Name: "my-addons-config",
		}
		if err := ref.Normalize(ctx, reader, "default"); err == nil {
			t.Fatalf("expected error for unready resource, got nil")
		}
	})
}
