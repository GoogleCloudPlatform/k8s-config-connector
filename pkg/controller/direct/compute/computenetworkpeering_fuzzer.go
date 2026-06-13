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
// proto.message: google.cloud.compute.v1.NetworkPeering
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeNetworkPeeringFuzzer())
}

func computeNetworkPeeringFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NetworkPeering{},
		ComputeNetworkPeeringSpec_v1beta1_FromProto, ComputeNetworkPeeringSpec_v1beta1_ToProto,
		ComputeNetworkPeeringStatus_v1beta1_FromProto, ComputeNetworkPeeringStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".export_custom_routes")
	f.SpecField(".export_subnet_routes_with_public_ip")
	f.SpecField(".import_custom_routes")
	f.SpecField(".import_subnet_routes_with_public_ip")
	f.SpecField(".network") // Maps to PeerNetworkRef
	f.SpecField(".name")    // Maps to ResourceID
	f.SpecField(".stack_type")

	// Status fields
	f.StatusField(".state")
	f.StatusField(".state_details")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".auto_create_routes")
	f.Unimplemented_NotYetTriaged(".connection_status")
	f.Unimplemented_NotYetTriaged(".exchange_subnet_routes")
	f.Unimplemented_NotYetTriaged(".peer_mtu")
	f.Unimplemented_NotYetTriaged(".update_strategy")

	return f
}
