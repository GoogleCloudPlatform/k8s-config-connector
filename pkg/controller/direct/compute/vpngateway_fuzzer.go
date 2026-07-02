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
// proto.message: google.cloud.compute.v1.VpnGateway
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeVPNGatewayFuzzer())
}

func computeVPNGatewayFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.VpnGateway, krm.ComputeVPNGatewaySpec](&pb.VpnGateway{},
		ComputeVPNGatewaySpec_v1beta1_FromProto, ComputeVPNGatewaySpec_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".network")
	f.SpecField(".region")
	f.SpecField(".stack_type")

	// Nested Spec fields
	f.SpecField(".vpn_interfaces[].id")
	f.SpecField(".vpn_interfaces[].interconnect_attachment")
	f.SpecField(".vpn_interfaces[].ip_address")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".creation_timestamp")
	f.Unimplemented_NotYetTriaged(".gateway_ip_version")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".vpn_interfaces[].ipv6_address")

	return f
}
