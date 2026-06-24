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

func TestComputeNetworkAttachmentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeNetworkAttachmentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/networkAttachments/my-na",
			want: &ComputeNetworkAttachmentIdentity{
				Project:           "my-project",
				Region:            "us-central1",
				NetworkAttachment: "my-na",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://compute.googleapis.com/projects/my-project/regions/us-central1/networkAttachments/my-na",
			want: &ComputeNetworkAttachmentIdentity{
				Project:           "my-project",
				Region:            "us-central1",
				NetworkAttachment: "my-na",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeNetworkAttachmentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeNetworkAttachmentRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/regions/us-central1/networkAttachments/my-na",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/regions/us-central1/networkAttachments/my-na",
			wantErr: true,
		},
		{
			name:    "missing region",
			ref:     "projects/my-project/networkAttachments/my-na",
			wantErr: true,
		},
		{
			name:    "missing networkAttachment",
			ref:     "projects/my-project/regions/us-central1",
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
			r := &ComputeNetworkAttachmentRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNetworkAttachmentRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeNetworkAttachmentRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeNetworkAttachmentRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &ComputeNetworkAttachmentRef{
				External: "projects/test-project/regions/us-central1/networkAttachments/test-na",
			},
			wantExternal: "projects/test-project/regions/us-central1/networkAttachments/test-na",
		},
		{
			name: "external with invalid format",
			ref: &ComputeNetworkAttachmentRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeNetworkAttachment external="invalid-format" was not known (use projects/{project}/regions/{region}/networkAttachments/{networkattachment})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeNetworkAttachmentRef{
				Name:      "test-na",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeNetworkAttachment",
						"metadata": map[string]interface{}{
							"name":      "test-na",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/regions/us-central1/networkAttachments/test-na",
						},
					},
				},
			},
			wantExternal: "projects/test-project/regions/us-central1/networkAttachments/test-na",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeNetworkAttachmentRef{
				Name:      "test-na",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeNetworkAttachment",
						"metadata": map[string]interface{}{
							"name":      "test-na",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeNetworkAttachment my-namespace/test-na is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeNetworkAttachmentRef{
				Name:      "test-na",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeNetworkAttachment my-namespace/test-na is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeNetworkAttachment{})
			cl := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), cl, "default")
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("Normalize() expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Errorf("Normalize() error = %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Normalize() unexpected error: %v", err)
			}
			if tc.ref.External != tc.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tc.ref.External, tc.wantExternal)
			}
		})
	}
}
