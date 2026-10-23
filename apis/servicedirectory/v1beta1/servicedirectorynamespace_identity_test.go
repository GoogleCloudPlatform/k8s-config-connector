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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestServiceDirectoryNamespaceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ServiceDirectoryNamespaceIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/namespaces/my-namespace",
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Namespace: "my-namespace",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://servicedirectory.googleapis.com/v1/projects/my-project/locations/us-central1/namespaces/my-namespace",
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Namespace: "my-namespace",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ServiceDirectoryNamespaceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedParent := "projects/" + tt.want.Project + "/locations/" + tt.want.Location
				if actualParent := i.ParentString(); actualParent != expectedParent {
					t.Errorf("ParentString() = %q, want %q", actualParent, expectedParent)
				}
			}
		})
	}
}

func TestServiceDirectoryNamespace_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                 string
		obj                  *ServiceDirectoryNamespace
		namespaceAnnotations map[string]string
		wantErr              bool
		want                 *ServiceDirectoryNamespaceIdentity
	}{
		{
			name: "project from spec.projectRef",
			obj: &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ns",
					Namespace: "k8s-ns",
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location: "us-central1",
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				Namespace: "test-ns",
			},
		},
		{
			name: "project from object annotation cnrm.cloud.google.com/project-id",
			obj: &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ns",
					Namespace: "k8s-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "annotated-project",
					},
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location: "us-central1",
				},
			},
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "annotated-project",
				Location:  "us-central1",
				Namespace: "test-ns",
			},
		},
		{
			name: "spec.projectRef takes precedence over object annotation",
			obj: &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ns",
					Namespace: "k8s-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "annotated-project",
					},
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location: "us-central1",
					ProjectRef: &refs.ProjectRef{
						External: "precedent-project",
					},
				},
			},
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "precedent-project",
				Location:  "us-central1",
				Namespace: "test-ns",
			},
		},
		{
			name: "project from namespace annotation cnrm.cloud.google.com/project-id",
			obj: &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ns",
					Namespace: "k8s-ns",
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location: "us-central1",
				},
			},
			namespaceAnnotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "ns-annotated-project",
			},
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "ns-annotated-project",
				Location:  "us-central1",
				Namespace: "test-ns",
			},
		},
		{
			name: "fallback to namespace name when no projectRef or annotation",
			obj: &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ns",
					Namespace: "k8s-project",
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location: "us-central1",
				},
			},
			want: &ServiceDirectoryNamespaceIdentity{
				Project:   "k8s-project",
				Location:  "us-central1",
				Namespace: "test-ns",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			_ = AddToScheme(scheme)
			_ = corev1.AddToScheme(scheme)
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        tt.obj.Namespace,
					Annotations: tt.namespaceAnnotations,
				},
			}
			fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(ns).Build()

			idVal, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got, ok := idVal.(*ServiceDirectoryNamespaceIdentity)
				if !ok {
					t.Fatalf("unexpected identity type: %T", idVal)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
