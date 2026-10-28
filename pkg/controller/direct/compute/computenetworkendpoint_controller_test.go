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

package compute

import (
	"context"
	"errors"
	"strings"
	"testing"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestBuildDesiredEndpoint(t *testing.T) {
	tests := []struct {
		name        string
		obj         *krm.ComputeNetworkEndpoint
		id          *krm.ComputeNetworkEndpointIdentity
		wantPort    int32
		wantIP      string
		wantErr     bool
		errContains string
	}{
		{
			name: "valid port from identity",
			obj: &krm.ComputeNetworkEndpoint{
				Spec: krm.ComputeNetworkEndpointSpec{
					IpAddress: "10.0.0.1",
				},
			},
			id: &krm.ComputeNetworkEndpointIdentity{
				IpAddress: "10.0.0.1",
				Port:      "8080",
			},
			wantPort: 8080,
			wantIP:   "10.0.0.1",
			wantErr:  false,
		},
		{
			name: "invalid port fails fast",
			obj: &krm.ComputeNetworkEndpoint{
				Spec: krm.ComputeNetworkEndpointSpec{
					IpAddress: "10.0.0.1",
				},
			},
			id: &krm.ComputeNetworkEndpointIdentity{
				IpAddress: "10.0.0.1",
				Port:      "not-a-number",
			},
			wantErr:     true,
			errContains: "invalid port \"not-a-number\"",
		},
		{
			name: "empty port allowed",
			obj: &krm.ComputeNetworkEndpoint{
				Spec: krm.ComputeNetworkEndpointSpec{
					IpAddress: "10.0.0.1",
				},
			},
			id: &krm.ComputeNetworkEndpointIdentity{
				IpAddress: "10.0.0.1",
			},
			wantPort: 0,
			wantIP:   "10.0.0.1",
			wantErr:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := buildDesiredEndpoint(tc.obj, tc.id)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Fatalf("expected error containing %q, got %v", tc.errContains, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.GetPort() != tc.wantPort {
				t.Errorf("port = %v, want %v", got.GetPort(), tc.wantPort)
			}
			if got.GetIpAddress() != tc.wantIP {
				t.Errorf("ipAddress = %v, want %v", got.GetIpAddress(), tc.wantIP)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	desired := &pb.NetworkEndpoint{
		IpAddress: direct.PtrTo("10.0.0.1"),
		Port:      direct.PtrTo(int32(80)),
		Instance:  direct.PtrTo("projects/p1/zones/z1/instances/vm-1"),
	}

	adapter := &ComputeNetworkEndpointAdapter{
		desired: desired,
	}

	// Exact match
	if !adapter.matches(&pb.NetworkEndpoint{
		IpAddress: direct.PtrTo("10.0.0.1"),
		Port:      direct.PtrTo(int32(80)),
		Instance:  direct.PtrTo("projects/p1/zones/z1/instances/vm-1"),
	}) {
		t.Errorf("expected exact match to succeed")
	}

	// Match with short instance name
	if !adapter.matches(&pb.NetworkEndpoint{
		IpAddress: direct.PtrTo("10.0.0.1"),
		Port:      direct.PtrTo(int32(80)),
		Instance:  direct.PtrTo("vm-1"),
	}) {
		t.Errorf("expected short name instance match to succeed")
	}

	// Mismatched IP
	if adapter.matches(&pb.NetworkEndpoint{
		IpAddress: direct.PtrTo("10.0.0.2"),
		Port:      direct.PtrTo(int32(80)),
		Instance:  direct.PtrTo("vm-1"),
	}) {
		t.Errorf("expected mismatched IP to fail")
	}

	// Mismatched Port
	if adapter.matches(&pb.NetworkEndpoint{
		IpAddress: direct.PtrTo("10.0.0.1"),
		Port:      direct.PtrTo(int32(8080)),
		Instance:  direct.PtrTo("vm-1"),
	}) {
		t.Errorf("expected mismatched Port to fail")
	}
}

func TestExportNaming(t *testing.T) {
	adapter := &ComputeNetworkEndpointAdapter{
		id: &krm.ComputeNetworkEndpointIdentity{
			Project:                     "test-project",
			Zone:                        "us-central1-a",
			ComputeNetworkEndpointGroup: "my-neg",
			Port:                        "80",
		},
		actual: &pb.NetworkEndpoint{
			IpAddress: direct.PtrTo("10.0.0.5"),
			Port:      direct.PtrTo(int32(80)),
		},
	}

	u, err := adapter.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() returned error: %v", err)
	}

	// Resource name should combine IP and Port to avoid collisions: 10-0-0-5-80
	wantName := "10-0-0-5-80"
	if u.GetName() != wantName {
		t.Errorf("Export() name = %q, want %q", u.GetName(), wantName)
	}
}

func TestIsConflict(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "googleapi 409 error",
			err:  &googleapi.Error{Code: 409},
			want: true,
		},
		{
			name: "RESOURCE_IN_USE_BY_ANOTHER_OPERATION string error",
			err:  errors.New("googleapi: Error 400: The resource 'projects/p/zones/z/networkEndpointGroups/neg' is not ready, RESOURCE_IN_USE_BY_ANOTHER_OPERATION"),
			want: true,
		},
		{
			name: "other error",
			err:  status.Error(codes.NotFound, "not found"),
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isConflict(tc.err); got != tc.want {
				t.Errorf("isConflict(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
