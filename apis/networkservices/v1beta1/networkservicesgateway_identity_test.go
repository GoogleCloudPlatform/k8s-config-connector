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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesGatewayIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesGatewayIdentity
		wantErr bool
	}{
		{
			name: "valid gateway reference",
			ref:  "projects/my-project/locations/us-central1/gateways/my-gateway",
			want: &NetworkServicesGatewayIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Gateway:  "my-gateway",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing gateways)",
			ref:     "projects/my-project/locations/us-central1/my-gateway",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/us-central1/gateways/my-gateway",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesGatewayIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesGatewayIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NetworkServicesGatewayIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesGatewayRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesGatewayRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesGatewayRef{
				External: "projects/my-project/locations/us-central1/gateways/my-gateway",
			},
			wantExternal: "projects/my-project/locations/us-central1/gateways/my-gateway",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesGatewayRef{
				Name:     "my-gateway",
				External: "projects/my-project/locations/us-central1/gateways/my-gateway",
			},
			wantErr: true,
		},
		{
			name: "resolve from fallback (metadata.name)",
			ref: &NetworkServicesGatewayRef{
				Name: "my-gateway",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesGateway",
						"metadata": map[string]interface{}{
							"name":      "my-gateway",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"location": "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/gateways/my-gateway",
			wantErr:      false,
		},
		{
			name: "resolve from fallback (resourceID)",
			ref: &NetworkServicesGatewayRef{
				Name: "k8s-name",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1beta1",
						"kind":       "NetworkServicesGateway",
						"metadata": map[string]interface{}{
							"name":      "k8s-name",
							"namespace": "test-ns",
						},
						"spec": map[string]interface{}{
							"resourceID": "my-gateway",
							"location":   "us-central1",
							"projectRef": map[string]interface{}{
								"external": "my-project",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/gateways/my-gateway",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesGatewayRef{
				Name: "non-existent",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build()
			err := tt.ref.Normalize(context.Background(), client, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesGatewayRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesGatewayRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
