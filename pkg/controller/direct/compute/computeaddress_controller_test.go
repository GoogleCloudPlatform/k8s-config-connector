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
	"testing"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestCompareAddress(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		actual   *pb.Address
		desired  *pb.Address
		wantDiff bool
	}{
		{
			name: "identical objects have no diff",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Address:     direct.PtrTo("10.0.0.1"),
				Description: direct.PtrTo("my address"),
				Purpose:     direct.PtrTo("GCE_ENDPOINT"),
				Subnetwork:  direct.PtrTo("projects/p1/regions/us-central1/subnetworks/default"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Address:     direct.PtrTo("10.0.0.1"),
				Description: direct.PtrTo("my address"),
				Purpose:     direct.PtrTo("GCE_ENDPOINT"),
				Subnetwork:  direct.PtrTo("projects/p1/regions/us-central1/subnetworks/default"),
			},
			wantDiff: false,
		},
		{
			name: "server-allocated address ignored when omitted in desired",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Address:     direct.PtrTo("10.128.0.90"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
			},
			wantDiff: false,
		},
		{
			name: "server-defaulted purpose ignored when omitted in desired",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Purpose:     direct.PtrTo("GCE_ENDPOINT"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
			},
			wantDiff: false,
		},
		{
			name: "subnetwork canonical URL matching relative path",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Subnetwork:  direct.PtrTo("https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/subnetworks/default"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Subnetwork:  direct.PtrTo("projects/my-project/regions/us-central1/subnetworks/default"),
			},
			wantDiff: false,
		},
		{
			name: "network canonical URL matching relative path",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Network:     direct.PtrTo("https://www.googleapis.com/compute/v1/projects/my-project/global/networks/default"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Network:     direct.PtrTo("projects/my-project/global/networks/default"),
			},
			wantDiff: false,
		},
		{
			name: "description omitted in desired ignored",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("EXTERNAL"),
				Description: direct.PtrTo("server or default desc"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("EXTERNAL"),
			},
			wantDiff: false,
		},
		{
			name: "different address when explicitly specified reported as diff",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Address:     direct.PtrTo("10.0.0.1"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Address:     direct.PtrTo("10.0.0.2"),
			},
			wantDiff: true,
		},
		{
			name: "different purpose when explicitly specified reported as diff",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Purpose:     direct.PtrTo("GCE_ENDPOINT"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("INTERNAL"),
				Purpose:     direct.PtrTo("VPC_PEERING"),
			},
			wantDiff: true,
		},
		{
			name: "different description when explicitly specified reported as diff",
			actual: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("EXTERNAL"),
				Description: direct.PtrTo("old description"),
			},
			desired: &pb.Address{
				Name:        direct.PtrTo("addr-1"),
				AddressType: direct.PtrTo("EXTERNAL"),
				Description: direct.PtrTo("new description"),
			},
			wantDiff: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diff, _, err := compareAddress(ctx, tc.actual, tc.desired)
			if err != nil {
				t.Fatalf("unexpected error comparing address: %v", err)
			}
			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("expected diff=%v, got=%v (fields: %v)", tc.wantDiff, diff.HasDiff(), diff.Fields)
			}
		})
	}
}
