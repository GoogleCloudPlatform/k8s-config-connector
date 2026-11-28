// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestProjectIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *ProjectIdentity
		wantErr bool
	}{
		{
			name: "valid project reference (long)",
			ref:  "projects/my-project",
			want: &ProjectIdentity{
				ProjectID: "my-project",
			},
			wantErr: false,
		},
		{
			name: "valid project reference (short)",
			ref:  "my-project",
			want: &ProjectIdentity{
				ProjectID: "my-project",
			},
			wantErr: false,
		},
		{
			name: "valid project reference (host prefixed)",
			ref:  "//cloudresourcemanager.googleapis.com/projects/my-project",
			want: &ProjectIdentity{
				ProjectID: "my-project",
			},
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "projects/my-project/other",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			want:    nil,
			wantErr: true, // Split returns [""] which len is 1, so it might take "" as projectID currently?
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ProjectIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.ProjectID != tt.want.ProjectID {
					t.Errorf("ProjectIdentity.FromExternal() = %v, want %v", i, tt.want)
				}
			}
		})
	}
}

func TestProjectRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external (long)",
			external: "projects/my-project",
			wantErr:  false,
		},
		{
			name:     "valid external (short)",
			external: "my-project",
			wantErr:  false,
		},
		{
			name:     "valid external (host prefixed short)",
			external: "//cloudresourcemanager.googleapis.com/my-project",
			wantErr:  true, // Expect failure with gcpurls if not handled
		},
		{
			name:     "invalid external",
			external: "invalid/my-project/extra",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProjectRef{}
			if err := r.ValidateExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("ProjectRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectImplementsIdentity(t *testing.T) {
	var _ identity.Identity = &ProjectIdentity{}
}

func TestResolveProject(t *testing.T) {
	tests := []struct {
		name    string
		ref     *ProjectRef
		wantID  string
		wantErr bool
	}{
		{
			name: "valid ref (long)",
			ref: &ProjectRef{
				External: "projects/my-project",
			},
			wantID:  "my-project",
			wantErr: false,
		},
		{
			name: "valid ref (short)",
			ref: &ProjectRef{
				External: "my-project",
			},
			wantID:  "my-project",
			wantErr: false,
		},
		{
			name: "valid ref (host prefixed short)",
			ref: &ProjectRef{
				External: "//cloudresourcemanager.googleapis.com/my-project",
			},
			wantID:  "my-project",
			wantErr: false,
		},
		{
			name: "invalid ref",
			ref: &ProjectRef{
				External: "bad/format/extra",
			},
			wantErr: true,
		},
		{
			name:    "nil ref",
			ref:     nil,
			wantErr: false, // returns nil, nil
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveProject(context.Background(), nil, "default", tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.ref == nil {
				if got != nil {
					t.Errorf("ResolveProject() got = %v, want nil", got)
				}
				return
			}
			if !tt.wantErr {
				if got.ProjectID != tt.wantID {
					t.Errorf("ResolveProject() got ID = %v, want %v", got.ProjectID, tt.wantID)
				}
			}
		})
	}
}

func TestProjectRef_SetExternal(t *testing.T) {
	r := &ProjectRef{
		Name:      "my-name",
		Namespace: "my-namespace",
	}
	r.SetExternal("projects/my-project")
	if r.External != "projects/my-project" {
		t.Errorf("expected External to be %q, got %q", "projects/my-project", r.External)
	}
	if r.Name != "" {
		t.Errorf("expected Name to be cleared, got %q", r.Name)
	}
	if r.Namespace != "" {
		t.Errorf("expected Namespace to be cleared, got %q", r.Namespace)
	}
}

func TestResolveProjectID(t *testing.T) {
	tests := []struct {
		name                 string
		spec                 map[string]interface{}
		objectAnnotations    map[string]string
		namespaceAnnotations map[string]string
		wantProjectID        string
		wantErr              bool
	}{
		{
			name: "1. spec.projectRef external takes precedence",
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "ref-project",
				},
			},
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "obj-project",
			},
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			wantProjectID: "ref-project",
		},
		{
			name: "2. object annotation takes precedence over namespace annotation and namespace name",
			objectAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "partition:sample-project",
			},
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			wantProjectID: "partition:sample-project",
		},
		{
			name: "3. namespace annotation takes precedence over namespace name when object annotation is absent",
			namespaceAnnotations: map[string]string{
				k8s.ProjectIDAnnotation: "ns-project",
			},
			wantProjectID: "ns-project",
		},
		{
			name:          "4. falls back to namespace name when no ref or annotations are set",
			wantProjectID: "test-ns",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "test-ns",
					Annotations: tt.namespaceAnnotations,
				},
			}
			s := runtime.NewScheme()
			if err := corev1.AddToScheme(s); err != nil {
				t.Fatalf("failed to add corev1 to scheme: %v", err)
			}
			reader := fake.NewClientBuilder().WithScheme(s).WithObjects(ns).Build()

			u := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test.cnrm.cloud.google.com/v1beta1",
					"kind":       "TestResource",
					"metadata": map[string]interface{}{
						"name":      "test-resource",
						"namespace": "test-ns",
					},
				},
			}
			if tt.spec != nil {
				u.Object["spec"] = tt.spec
			}
			if tt.objectAnnotations != nil {
				u.SetAnnotations(tt.objectAnnotations)
			}

			got, err := ResolveProjectID(context.Background(), reader, u)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ResolveProjectID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.wantProjectID {
				t.Errorf("ResolveProjectID() = %q, want %q", got, tt.wantProjectID)
			}
		})
	}
}
