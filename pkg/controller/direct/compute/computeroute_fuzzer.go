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
// proto.message: google.cloud.compute.v1.Route
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeRouteFuzzer())
}

func computeRouteFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Route{},
		ComputeRouteSpec_v1beta1_FromProto, ComputeRouteSpec_v1beta1_ToProto,
		ComputeRouteStatus_v1beta1_FromProto, ComputeRouteStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".dest_range")
	f.SpecField(".network")
	f.SpecField(".next_hop_gateway")
	f.SpecField(".next_hop_instance")
	f.SpecField(".next_hop_ip")
	f.SpecField(".next_hop_vpn_tunnel")
	f.SpecField(".priority")
	f.SpecField(".tags")

	// Status fields
	f.StatusField(".next_hop_network")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".as_paths")
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".next_hop_hub")
	f.Unimplemented_NotYetTriaged(".next_hop_ilb")
	f.Unimplemented_NotYetTriaged(".next_hop_inter_region_cost")
	f.Unimplemented_NotYetTriaged(".next_hop_interconnect_attachment")
	f.Unimplemented_NotYetTriaged(".next_hop_med")
	f.Unimplemented_NotYetTriaged(".next_hop_origin")
	f.Unimplemented_NotYetTriaged(".next_hop_peering")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".route_status")
	f.Unimplemented_NotYetTriaged(".route_type")
	f.Unimplemented_NotYetTriaged(".warnings")

	return f
}
