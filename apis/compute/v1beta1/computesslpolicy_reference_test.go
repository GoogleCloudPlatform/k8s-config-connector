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

func TestComputeSSLPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid external reference",
			ref:     "projects/my-project/global/sslPolicies/my-policy",
			wantErr: false,
		},
		{
			name:    "invalid external reference",
			ref:     "invalid/format",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeSSLPolicyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeSSLPolicyRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	policy := &unstructured.Unstructured{}
	policy.SetGroupVersionKind(ComputeSSLPolicyGVK)
	policy.SetName("my-policy")
	policy.SetNamespace("my-ns")
	policy.Object["status"] = map[string]interface{}{
		"selfLink": "https://www.googleapis.com/compute/v1/projects/my-project/global/sslPolicies/my-policy",
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(policy).Build()

	tests := []struct {
		name             string
		ref              *ComputeSSLPolicyRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &ComputeSSLPolicyRef{
				External: "projects/my-project/global/sslPolicies/my-policy",
			},
			want: "projects/my-project/global/sslPolicies/my-policy",
		},
		{
			name: "internal reference",
			ref: &ComputeSSLPolicyRef{
				Name:      "my-policy",
				Namespace: "my-ns",
			},
			want: "projects/my-project/global/sslPolicies/my-policy",
		},
		{
			name: "internal reference with default namespace",
			ref: &ComputeSSLPolicyRef{
				Name: "my-policy",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/global/sslPolicies/my-policy",
		},
		{
			name: "internal reference not found",
			ref: &ComputeSSLPolicyRef{
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
