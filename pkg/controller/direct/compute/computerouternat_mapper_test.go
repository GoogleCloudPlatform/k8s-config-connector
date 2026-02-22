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
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestComputeRouterNATSpec_Mapper(t *testing.T) {
	mapCtx := &direct.MapContext{}

	krmSpec := &krm.ComputeRouterNATSpec{
		NatIpAllocateOption: direct.PtrTo("MANUAL_ONLY"),
		NatIps: []refs.ComputeAddressRef{
			{External: "projects/p1/regions/r1/addresses/a1"},
		},
		DrainNatIps: []refs.ComputeAddressRef{
			{External: "projects/p1/regions/r1/addresses/d1"},
		},
		Rules: []krm.RouterNatRule{
			{
				RuleNumber: direct.PtrTo(uint32(1)),
				Match:      direct.PtrTo("inIpRange(destination.ip, '1.1.1.1/32')"),
				Action: &krm.RouterNatRuleAction{
					SourceNatActiveIpsRefs: []refs.ComputeAddressRef{
						{External: "projects/p1/regions/r1/addresses/a2"},
					},
				},
			},
		},
		SourceSubnetworkIpRangesToNat: "LIST_OF_SUBNETWORKS",
		Subnetwork: []krm.RouterNatSubnetwork{
			{
				SubnetworkRef:       refs.ComputeSubnetworkRef{External: "projects/p1/regions/r1/subnetworks/s1"},
				SourceIpRangesToNat: []string{"PRIMARY_IP_RANGE"},
			},
		},
	}

	proto := ComputeRouterNATSpec_v1beta1_ToProto(mapCtx, krmSpec)
	if mapCtx.Err() != nil {
		t.Fatalf("ToProto failed: %v", mapCtx.Err())
	}

	if proto.GetNatIpAllocateOption() != "MANUAL_ONLY" {
		t.Errorf("expected NatIpAllocateOption MANUAL_ONLY, got %v", proto.GetNatIpAllocateOption())
	}

	if !reflect.DeepEqual(proto.NatIps, []string{"projects/p1/regions/r1/addresses/a1"}) {
		t.Errorf("unexpected NatIps: %v", proto.NatIps)
	}

	if len(proto.Rules) != 1 || proto.Rules[0].Action.GetSourceNatActiveIps()[0] != "projects/p1/regions/r1/addresses/a2" {
		t.Errorf("unexpected Rules: %v", proto.Rules)
	}

	back := ComputeRouterNATSpec_v1beta1_FromProto(mapCtx, proto)
	if mapCtx.Err() != nil {
		t.Fatalf("FromProto failed: %v", mapCtx.Err())
	}

	if back.SourceSubnetworkIpRangesToNat != "LIST_OF_SUBNETWORKS" {
		t.Errorf("expected SourceSubnetworkIpRangesToNat LIST_OF_SUBNETWORKS, got %v", back.SourceSubnetworkIpRangesToNat)
	}

	if len(back.Subnetwork) != 1 || back.Subnetwork[0].SubnetworkRef.External != "projects/p1/regions/r1/subnetworks/s1" {
		t.Errorf("unexpected Subnetwork: %v", back.Subnetwork)
	}
}
