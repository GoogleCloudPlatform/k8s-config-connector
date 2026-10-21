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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

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
	tests := []struct {
		name                 string
		objectAnnotations    map[string]string
		namespaceAnnotations map[string]string
		wantProject          string
	}{
		{
			name: "respects object-level project-id annotation over namespace name",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "partition:sample-project",
			},
			wantProject: "partition:sample-project",
		},
		{
			name: "respects object-level project-id annotation over namespace annotation",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "partition:sample-project",
			},
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			wantProject: "partition:sample-project",
		},
		{
			name: "falls back to namespace annotation when object annotation is absent",
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			wantProject: "ns-project",
		},
		{
			name:        "falls back to namespace name when no annotations are present",
			wantProject: "kcc-audit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "kcc-audit",
					Annotations: tt.namespaceAnnotations,
				},
			}
			s := runtime.NewScheme()
			if err := corev1.AddToScheme(s); err != nil {
				t.Fatalf("failed to add corev1 to scheme: %v", err)
			}
			reader := fake.NewClientBuilder().WithScheme(s).WithObjects(ns).Build()

			resourceID := "sample-sd-namespace"
			obj := &ServiceDirectoryNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "sample-sd-namespace",
					Namespace:   "kcc-audit",
					Annotations: tt.objectAnnotations,
				},
				Spec: ServiceDirectoryNamespaceSpec{
					Location:   "us-central1",
					ResourceID: &resourceID,
				},
			}

			idVal, err := obj.GetIdentity(context.Background(), reader)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			id, ok := idVal.(*ServiceDirectoryNamespaceIdentity)
			if !ok {
				t.Fatalf("unexpected identity type: %T", idVal)
			}
			if id.Project != tt.wantProject {
				t.Errorf("got project %q, want %q", id.Project, tt.wantProject)
			}
			if id.Location != "us-central1" {
				t.Errorf("got location %q, want %q", id.Location, "us-central1")
			}
			if id.Namespace != "sample-sd-namespace" {
				t.Errorf("got namespace %q, want %q", id.Namespace, "sample-sd-namespace")
			}
		})
	}
}
