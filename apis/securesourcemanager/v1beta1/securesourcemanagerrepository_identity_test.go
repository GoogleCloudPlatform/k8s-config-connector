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

func TestSecureSourceManagerRepositoryRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/repositories/my-repo",
			wantErr: false,
		},
		{
			name:    "valid full url",
			ref:     "//securesourcemanager.googleapis.com/projects/my-project/locations/us-central1/repositories/my-repo",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/repositories/my-repo",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/repositories/my-repo",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SecureSourceManagerRepositoryRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecureSourceManagerRepositoryRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	repo := &unstructured.Unstructured{}
	repo.SetGroupVersionKind(SecureSourceManagerRepositoryGVK)
	repo.SetName("my-repo")
	repo.SetNamespace("my-ns")
	repo.Object["spec"] = map[string]interface{}{
		"location": "us-central1",
		"projectRef": map[string]interface{}{
			"external": "my-project",
		},
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(repo).Build()

	tests := []struct {
		name             string
		ref              *SecureSourceManagerRepositoryRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &SecureSourceManagerRepositoryRef{
				External: "projects/my-project/locations/us-central1/repositories/my-repo",
			},
			want: "projects/my-project/locations/us-central1/repositories/my-repo",
		},
		{
			name: "internal reference resolved via fallback",
			ref: &SecureSourceManagerRepositoryRef{
				Name:      "my-repo",
				Namespace: "my-ns",
			},
			want: "projects/my-project/locations/us-central1/repositories/my-repo",
		},
		{
			name: "internal reference with default namespace",
			ref: &SecureSourceManagerRepositoryRef{
				Name: "my-repo",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/locations/us-central1/repositories/my-repo",
		},
		{
			name: "internal reference not found",
			ref: &SecureSourceManagerRepositoryRef{
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
			if tt.ref.External != tt.want {
				t.Errorf("Normalize() got = %v, want %v", tt.ref.External, tt.want)
			}
		})
	}
}
