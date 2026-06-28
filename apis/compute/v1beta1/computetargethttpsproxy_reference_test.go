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

func TestComputeTargetHTTPSProxyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid global external reference",
			ref:     "projects/my-project/global/targetHttpsProxies/my-proxy",
			wantErr: false,
		},
		{
			name:    "valid regional external reference",
			ref:     "projects/my-project/regions/us-central1/targetHttpsProxies/my-proxy",
			wantErr: false,
		},
		{
			name:    "invalid external reference",
			ref:     "my-proxy",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ComputeTargetHTTPSProxyRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeTargetHTTPSProxyRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()
	_ = AddToScheme(s)

	proxy := &unstructured.Unstructured{}
	proxy.SetGroupVersionKind(ComputeTargetHTTPSProxyGVK)
	proxy.SetName("my-proxy")
	proxy.SetNamespace("my-ns")
	proxy.Object["spec"] = map[string]interface{}{
		"resourceID": "my-proxy-id",
		"location":   "global",
	}
	policy := &unstructured.Unstructured{}
	policy.SetGroupVersionKind(ComputeTargetHTTPSProxyGVK)
	policy.SetName("my-proxy")
	policy.SetNamespace("my-ns")
	policy.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/project-id": "my-project",
	})
	policy.Object["spec"] = map[string]interface{}{
		"resourceID": "my-proxy-id",
		"location":   "global",
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(policy).Build()

	tests := []struct {
		name             string
		ref              *ComputeTargetHTTPSProxyRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &ComputeTargetHTTPSProxyRef{
				External: "projects/my-project/global/targetHttpsProxies/my-proxy",
			},
			want: "projects/my-project/global/targetHttpsProxies/my-proxy",
		},
		{
			name: "internal reference",
			ref: &ComputeTargetHTTPSProxyRef{
				Name:      "my-proxy",
				Namespace: "my-ns",
			},
			want: "projects/my-project/global/targetHttpsProxies/my-proxy-id",
		},
		{
			name: "internal reference with default namespace",
			ref: &ComputeTargetHTTPSProxyRef{
				Name: "my-proxy",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/global/targetHttpsProxies/my-proxy-id",
		},
		{
			name: "internal reference not found",
			ref: &ComputeTargetHTTPSProxyRef{
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
