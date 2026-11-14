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

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestCompareComputeRoute(t *testing.T) {
	ctx := context.Background()
	id := &krm.ComputeRouteIdentity{
		Project: "test-project",
		Route:   "test-route",
	}

	testCases := []struct {
		name       string
		actual     *computepb.Route
		desired    *computepb.Route
		expectDiff bool
		diffField  string
	}{
		{
			name: "identical routes",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			expectDiff: false,
		},
		{
			name: "network full URL vs relative path",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/beta/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			expectDiff: false,
		},
		{
			name: "network short name canonicalization",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			expectDiff: false,
		},
		{
			name: "nextHopGateway short name vs full URL",
			actual: &computepb.Route{
				Name:           direct.PtrTo("test-route"),
				DestRange:      direct.PtrTo("0.0.0.0/0"),
				Network:        direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopGateway: direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/gateways/default-internet-gateway"),
				Priority:       direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:           direct.PtrTo("test-route"),
				DestRange:      direct.PtrTo("0.0.0.0/0"),
				Network:        direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopGateway: direct.PtrTo("default-internet-gateway"),
				Priority:       direct.PtrTo[uint32](100),
			},
			expectDiff: false,
		},
		{
			name: "nextHopGateway global/gateways prefix",
			actual: &computepb.Route{
				Name:           direct.PtrTo("test-route"),
				DestRange:      direct.PtrTo("0.0.0.0/0"),
				Network:        direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopGateway: direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/gateways/default-internet-gateway"),
				Priority:       direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:           direct.PtrTo("test-route"),
				DestRange:      direct.PtrTo("0.0.0.0/0"),
				Network:        direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopGateway: direct.PtrTo("global/gateways/default-internet-gateway"),
				Priority:       direct.PtrTo[uint32](100),
			},
			expectDiff: false,
		},
		{
			name: "priority default 1000",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](1000),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  nil,
			},
			expectDiff: false,
		},
		{
			name: "different priority detects diff",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](200),
			},
			expectDiff: true,
			diffField:  "priority",
		},
		{
			name: "different destRange detects diff",
			actual: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("0.0.0.0/0"),
				Network:   direct.PtrTo("https://www.googleapis.com/compute/v1/projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			desired: &computepb.Route{
				Name:      direct.PtrTo("test-route"),
				DestRange: direct.PtrTo("192.168.0.0/16"),
				Network:   direct.PtrTo("projects/test-project/global/networks/test-network"),
				NextHopIp: direct.PtrTo("10.132.1.5"),
				Priority:  direct.PtrTo[uint32](100),
			},
			expectDiff: true,
			diffField:  "dest_range",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			diff, _, err := compareComputeRoute(ctx, tc.actual, tc.desired, id)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			hasDiff := diff.HasDiff()
			if hasDiff != tc.expectDiff {
				t.Fatalf("expected diff=%v, got %v (fields: %+v)", tc.expectDiff, hasDiff, diff.Fields)
			}
		})
	}
}
