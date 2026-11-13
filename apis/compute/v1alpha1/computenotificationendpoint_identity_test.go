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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestComputeNotificationEndpointIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeNotificationEndpointIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/notificationEndpoints/my-endpoint",
			want: &ComputeNotificationEndpointIdentity{
				Project:              "my-project",
				Region:               "us-central1",
				Notificationendpoint: "my-endpoint",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "wrong resource type",
			ref:     "projects/my-project/regions/us-central1/wrongTypes/my-endpoint",
			wantErr: true,
		},
		{
			name: "reference with host prefix",
			ref:  "//compute.googleapis.com/projects/my-project/regions/us-central1/notificationEndpoints/my-endpoint",
			want: &ComputeNotificationEndpointIdentity{
				Project:              "my-project",
				Region:               "us-central1",
				Notificationendpoint: "my-endpoint",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeNotificationEndpointIdentity{}
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

func TestComputeNotificationEndpointIdentity_ParentString(t *testing.T) {
	id := &ComputeNotificationEndpointIdentity{
		Project:              "my-project",
		Region:               "us-central1",
		Notificationendpoint: "my-endpoint",
	}
	want := "projects/my-project/regions/us-central1"
	got := id.ParentString()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}
}

func TestComputeNotificationEndpointIdentity_GetIdentity(t *testing.T) {
	s := runtime.NewScheme()
	s.AddKnownTypes(GroupVersion, &ComputeNotificationEndpoint{})
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	tests := []struct {
		name    string
		obj     *ComputeNotificationEndpoint
		want    *ComputeNotificationEndpointIdentity
		wantErr bool
	}{
		{
			name: "from spec with resourceID",
			obj: &ComputeNotificationEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-endpoint",
					Namespace: "default",
				},
				Spec: ComputeNotificationEndpointSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location:   "us-central1",
					ResourceID: ptrTo("custom-endpoint"),
				},
			},
			want: &ComputeNotificationEndpointIdentity{
				Project:              "my-project",
				Region:               "us-central1",
				Notificationendpoint: "custom-endpoint",
			},
		},
		{
			name: "from spec defaulting to metadata.name",
			obj: &ComputeNotificationEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-endpoint",
					Namespace: "default",
				},
				Spec: ComputeNotificationEndpointSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1",
				},
			},
			want: &ComputeNotificationEndpointIdentity{
				Project:              "my-project",
				Region:               "us-central1",
				Notificationendpoint: "k8s-endpoint",
			},
		},
		{
			name: "matching status.externalRef",
			obj: &ComputeNotificationEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-endpoint",
					Namespace: "default",
				},
				Spec: ComputeNotificationEndpointSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1",
				},
				Status: ComputeNotificationEndpointStatus{
					ExternalRef: ptrTo("projects/my-project/regions/us-central1/notificationEndpoints/k8s-endpoint"),
				},
			},
			want: &ComputeNotificationEndpointIdentity{
				Project:              "my-project",
				Region:               "us-central1",
				Notificationendpoint: "k8s-endpoint",
			},
		},
		{
			name: "mismatched status.externalRef",
			obj: &ComputeNotificationEndpoint{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-endpoint",
					Namespace: "default",
				},
				Spec: ComputeNotificationEndpointSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project",
					},
					Location: "us-central1",
				},
				Status: ComputeNotificationEndpointStatus{
					ExternalRef: ptrTo("projects/other-project/regions/us-central1/notificationEndpoints/k8s-endpoint"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(context.Background(), cl)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestComputeNotificationEndpointRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/regions/us-central1/notificationEndpoints/my-endpoint",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/regions/us-central1/notificationEndpoints/my-endpoint",
			wantErr: true,
		},
		{
			name:    "missing region",
			ref:     "projects/my-project/notificationEndpoints/my-endpoint",
			wantErr: true,
		},
		{
			name:    "missing endpoint id",
			ref:     "projects/my-project/regions/us-central1/notificationEndpoints",
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
			r := &ComputeNotificationEndpointRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNotificationEndpointRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeNotificationEndpointRef_Normalize(t *testing.T) {
	testCases := []struct {
		name         string
		ref          *ComputeNotificationEndpointRef
		objects      []runtime.Object
		wantExternal string
		wantErr      string
	}{
		{
			name: "external with valid format",
			ref: &ComputeNotificationEndpointRef{
				External: "projects/test-project/regions/us-central1/notificationEndpoints/test-endpoint",
			},
			wantExternal: "projects/test-project/regions/us-central1/notificationEndpoints/test-endpoint",
		},
		{
			name: "external with invalid format",
			ref: &ComputeNotificationEndpointRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeNotificationEndpoint external="invalid-format" was not known (use projects/{project}/regions/{region}/notificationEndpoints/{notificationendpoint})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &ComputeNotificationEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeNotificationEndpoint",
						"metadata": map[string]interface{}{
							"name":      "test-endpoint",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/test-project/regions/us-central1/notificationEndpoints/test-endpoint",
						},
					},
				},
			},
			wantExternal: "projects/test-project/regions/us-central1/notificationEndpoints/test-endpoint",
		},
		{
			name: "name specified, without status.externalRef",
			ref: &ComputeNotificationEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "compute.cnrm.cloud.google.com/v1alpha1",
						"kind":       "ComputeNotificationEndpoint",
						"metadata": map[string]interface{}{
							"name":      "test-endpoint",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{},
					},
				},
			},
			wantErr: `reference ComputeNotificationEndpoint my-namespace/test-endpoint is not ready`,
		},
		{
			name: "name specified, resource not found",
			ref: &ComputeNotificationEndpointRef{
				Name:      "test-endpoint",
				Namespace: "my-namespace",
			},
			wantErr: `reference ComputeNotificationEndpoint my-namespace/test-endpoint is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeNotificationEndpoint{})
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
