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

func TestComputeURLMapRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid external reference (projects/)",
			ref:     "projects/my-project/global/urlMaps/my-urlmap",
			wantErr: false,
		},
		{
			name:    "valid external reference (https://)",
			ref:     "https://www.googleapis.com/compute/v1/projects/my-project/global/urlMaps/my-urlmap",
			wantErr: false,
		},
		{
			name:    "invalid external reference",
			ref:     "my-urlmap",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeURLMapRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeURLMapRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	urlmap := &unstructured.Unstructured{}
	urlmap.SetGroupVersionKind(ComputeURLMapGVK)
	urlmap.SetName("my-urlmap")
	urlmap.SetNamespace("my-ns")
	urlmap.Object["status"] = map[string]interface{}{
		"selfLink": "https://www.googleapis.com/compute/v1/projects/my-project/global/urlMaps/my-urlmap",
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(urlmap).Build()

	tests := []struct {
		name             string
		ref              *ComputeURLMapRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &ComputeURLMapRef{
				External: "projects/my-project/global/urlMaps/my-urlmap",
			},
			want: "projects/my-project/global/urlMaps/my-urlmap",
		},
		{
			name: "internal reference",
			ref: &ComputeURLMapRef{
				Name:      "my-urlmap",
				Namespace: "my-ns",
			},
			want: "https://www.googleapis.com/compute/v1/projects/my-project/global/urlMaps/my-urlmap",
		},
		{
			name: "internal reference with default namespace",
			ref: &ComputeURLMapRef{
				Name: "my-urlmap",
			},
			defaultNamespace: "my-ns",
			want:             "https://www.googleapis.com/compute/v1/projects/my-project/global/urlMaps/my-urlmap",
		},
		{
			name: "internal reference not found",
			ref: &ComputeURLMapRef{
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
