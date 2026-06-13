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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.RouterBgpPeer
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeRouterPeerFuzzer())
}

func computeRouterPeerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.RouterBgpPeer{},
		ComputeRouterPeerSpec_v1beta1_FromProto, ComputeRouterPeerSpec_v1beta1_ToProto,
		ComputeRouterPeerStatus_v1beta1_FromProto, ComputeRouterPeerStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".advertise_mode")
	f.SpecField(".advertised_groups")
	f.SpecField(".advertised_ip_ranges")
	f.SpecField(".advertised_ip_ranges[].description")
	f.SpecField(".advertised_ip_ranges[].range")
	f.SpecField(".advertised_route_priority")
	f.SpecField(".bfd")
	f.SpecField(".bfd.min_receive_interval")
	f.SpecField(".bfd.min_transmit_interval")
	f.SpecField(".bfd.multiplier")
	f.SpecField(".bfd.session_initialization_mode")
	f.SpecField(".enable")
	f.SpecField(".enable_ipv6")
	f.SpecField(".ip_address")
	f.SpecField(".ipv6_nexthop_address")
	f.SpecField(".peer_asn")
	f.SpecField(".peer_ip_address")
	f.SpecField(".peer_ipv6_nexthop_address")
	f.SpecField(".router_appliance_instance")
	f.SpecField(".interface_name")

	// Status fields
	f.StatusField(".management_type")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".custom_learned_ip_ranges")
	f.Unimplemented_NotYetTriaged(".custom_learned_route_priority")
	f.Unimplemented_NotYetTriaged(".enable_ipv4")
	f.Unimplemented_NotYetTriaged(".export_policies")
	f.Unimplemented_NotYetTriaged(".import_policies")
	f.Unimplemented_NotYetTriaged(".ipv4_nexthop_address")
	f.Unimplemented_NotYetTriaged(".md5_authentication_key_name")
	f.Unimplemented_NotYetTriaged(".peer_ipv4_nexthop_address")

	f.FilterSpec = func(in *pb.RouterBgpPeer) {
		if in.PeerAsn == nil {
			var zero uint32 = 0
			in.PeerAsn = &zero
		}
		if in.Enable != nil {
			if *in.Enable == "TRUE" {
				// keep TRUE
			} else {
				val := "FALSE"
				in.Enable = &val
			}
		}
	}

	return f
}
