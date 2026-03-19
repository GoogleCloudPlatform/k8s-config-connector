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

package v1beta1_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeNodeGroupRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()
	v1beta1.AddToScheme(scheme)

	tests := []struct {
		name             string
		ref              *v1beta1.ComputeNodeGroupRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &v1beta1.ComputeNodeGroupRef{
				External: "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			},
			wantExternal: "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
		},
		{
			name: "external full URL",
			ref: &v1beta1.ComputeNodeGroupRef{
				External: "https://www.googleapis.com/compute/v1/projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			},
			wantExternal: "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
		},
		{
			name: "normalize from status.externalRef",
			ref: &v1beta1.ComputeNodeGroupRef{
				Name:      "my-node-group",
				Namespace: "my-ns",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeNodeGroup",
						"metadata": map[string]interface{}{
							"name":      "my-node-group",
							"namespace": "my-ns",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
						},
					},
				},
			},
			wantExternal: "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
		},
		{
			name: "normalize from status.selfLink fallback",
			ref: &v1beta1.ComputeNodeGroupRef{
				Name:      "my-node-group",
				Namespace: "my-ns",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
						"kind":       "ComputeNodeGroup",
						"metadata": map[string]interface{}{
							"name":      "my-node-group",
							"namespace": "my-ns",
						},
						"status": map[string]interface{}{
							"selfLink": "https://www.googleapis.com/compute/v1/projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
						},
					},
				},
			},
			wantExternal: "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
		},
		{
			name: "resource not found",
			ref: &v1beta1.ComputeNodeGroupRef{
				Name:      "not-found",
				Namespace: "my-ns",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build()
			err := tt.ref.Normalize(context.Background(), client, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("Normalize() gotExternal = %v, wantExternal %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}

func TestComputeNodeGroupRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid relative path",
			ref:     "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "projects/my-project/nodeGroups/my-node-group",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &v1beta1.ComputeNodeGroupRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNodeGroupRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
