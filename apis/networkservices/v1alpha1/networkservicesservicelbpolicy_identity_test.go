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

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNetworkServicesServiceLBPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *NetworkServicesServiceLBPolicyIdentity
		wantErr bool
	}{
		{
			name: "valid serviceLbPolicy reference",
			ref:  "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			want: &NetworkServicesServiceLBPolicyIdentity{
				Project:         "my-project",
				Location:        "us-central1",
				ServiceLbPolicy: "my-policy",
			},
			wantErr: false,
		},
		{
			name:    "invalid format (missing serviceLbPolicies)",
			ref:     "projects/my-project/locations/us-central1/my-policy",
			wantErr: true,
		},
		{
			name:    "invalid format (wrong prefix)",
			ref:     "locations/us-central1/serviceLbPolicies/my-policy",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &NetworkServicesServiceLBPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetworkServicesServiceLBPolicyIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NetworkServicesServiceLBPolicyIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestNetworkServicesServiceLBPolicyRef_Normalize(t *testing.T) {
	scheme := runtime.NewScheme()

	tests := []struct {
		name             string
		ref              *NetworkServicesServiceLBPolicyRef
		objects          []runtime.Object
		defaultNamespace string
		wantExternal     string
		wantErr          bool
	}{
		{
			name: "external already set",
			ref: &NetworkServicesServiceLBPolicyRef{
				External: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			},
			wantExternal: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr:      false,
		},
		{
			name: "both name and external set",
			ref: &NetworkServicesServiceLBPolicyRef{
				Name:     "my-policy",
				External: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			},
			wantErr: true,
		},
		{
			name: "resolve from status.externalRef",
			ref: &NetworkServicesServiceLBPolicyRef{
				Name: "my-policy",
			},
			defaultNamespace: "test-ns",
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "networkservices.cnrm.cloud.google.com/v1alpha1",
						"kind":       "NetworkServicesServiceLBPolicy",
						"metadata": map[string]interface{}{
							"name":      "my-policy",
							"namespace": "test-ns",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/serviceLbPolicies/my-policy",
			wantErr:      false,
		},
		{
			name: "missing reference",
			ref: &NetworkServicesServiceLBPolicyRef{
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
				t.Errorf("NetworkServicesServiceLBPolicyRef.Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("NetworkServicesServiceLBPolicyRef.Normalize() got = %v, want %v", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
