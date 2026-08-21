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
// See the License for the_identity.go specific language governing permissions and
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

func TestComputeNetworkEndpointIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ComputeNetworkEndpointIdentity
	}{
		{
			name: "valid reference with instance",
			ref:  "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg/my-instance/10.0.0.1/80",
			want: &ComputeNetworkEndpointIdentity{
				Project:                     "my-project",
				Zone:                        "us-central1-a",
				ComputeNetworkEndpointGroup: "my-neg",
				Instance:                    "my-instance",
				IpAddress:                   "10.0.0.1",
				Port:                        "80",
			},
		},
		{
			name: "valid reference without instance",
			ref:  "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1/80",
			want: &ComputeNetworkEndpointIdentity{
				Project:                     "my-project",
				Zone:                        "us-central1-a",
				ComputeNetworkEndpointGroup: "my-neg",
				Instance:                    "",
				IpAddress:                   "10.0.0.1",
				Port:                        "80",
			},
		},
		{
			name: "full url with instance",
			ref:  "https://compute.googleapis.com/projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg/my-instance/10.0.0.1/80",
			want: &ComputeNetworkEndpointIdentity{
				Project:                     "my-project",
				Zone:                        "us-central1-a",
				ComputeNetworkEndpointGroup: "my-neg",
				Instance:                    "my-instance",
				IpAddress:                   "10.0.0.1",
				Port:                        "80",
			},
		},
		{
			name: "full url without instance",
			ref:  "https://compute.googleapis.com/projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1/80",
			want: &ComputeNetworkEndpointIdentity{
				Project:                     "my-project",
				Zone:                        "us-central1-a",
				ComputeNetworkEndpointGroup: "my-neg",
				Instance:                    "",
				IpAddress:                   "10.0.0.1",
				Port:                        "80",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeNetworkEndpointIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.ref && "https://compute.googleapis.com/"+i.String() != tt.ref {
					// Compare string format
					formatted := i.String()
					var expected string
					if tt.want.Instance != "" {
						expected = "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg/my-instance/10.0.0.1/80"
					} else {
						expected = "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1/80"
					}
					if formatted != expected {
						t.Errorf("String() mismatch: got %q, want %q", formatted, expected)
					}
				}
			}
		})
	}
}

func TestComputeNetworkEndpointRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference with instance",
			ref:     "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg/my-instance/10.0.0.1/80",
			wantErr: false,
		},
		{
			name:    "valid reference without instance",
			ref:     "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1/80",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1/80",
			wantErr: true,
		},
		{
			name:    "missing port",
			ref:     "projects/my-project/zones/us-central1-a/networkEndpointGroups/my-neg//10.0.0.1",
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
			r := &ComputeNetworkEndpointRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNetworkEndpointRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComputeNetworkEndpointRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *ComputeNetworkEndpointRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format (with instance)",
			ref: &ComputeNetworkEndpointRef{
				External: "projects/test-project/zones/us-central1-a/networkEndpointGroups/test-neg/test-instance/10.0.0.1/80",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/networkEndpointGroups/test-neg/test-instance/10.0.0.1/80",
		},
		{
			name: "external with valid format (without instance)",
			ref: &ComputeNetworkEndpointRef{
				External: "projects/test-project/zones/us-central1-a/networkEndpointGroups/test-neg//10.0.0.1/80",
			},
			wantExternal: "projects/test-project/zones/us-central1-a/networkEndpointGroups/test-neg//10.0.0.1/80",
		},
		{
			name: "external with invalid format",
			ref: &ComputeNetworkEndpointRef{
				External: "invalid-format",
			},
			wantErr: `format of ComputeNetworkEndpoint external="invalid-format" was not known (use projects/{Project}/zones/{Zone}/networkEndpointGroups/{ComputeNetworkEndpointGroup}/{Instance}/{IpAddress}/{Port} or projects/{Project}/zones/{Zone}/networkEndpointGroups/{ComputeNetworkEndpointGroup}//{IpAddress}/{Port})`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(GroupVersion, &unstructured.Unstructured{})
			s.AddKnownTypes(GroupVersion, &ComputeNetworkEndpoint{})
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
