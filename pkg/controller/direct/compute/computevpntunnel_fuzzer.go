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
// proto.message: google.cloud.compute.v1.VpnTunnel
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeVPNTunnelFuzzer())
}

func computeVPNTunnelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.VpnTunnel{},
		ComputeVPNTunnelSpec_v1beta1_FromProto, ComputeVPNTunnelSpec_v1beta1_ToProto,
		ComputeVPNTunnelStatus_v1beta1_FromProto, ComputeVPNTunnelStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".ike_version")
	f.SpecField(".local_traffic_selector")
	f.SpecField(".peer_external_gateway")
	f.SpecField(".peer_external_gateway_interface")
	f.SpecField(".peer_gcp_gateway")
	f.SpecField(".peer_ip")
	f.SpecField(".region")
	f.SpecField(".remote_traffic_selector")
	f.SpecField(".router")
	f.SpecField(".shared_secret")
	f.SpecField(".target_vpn_gateway")
	f.SpecField(".vpn_gateway")
	f.SpecField(".vpn_gateway_interface")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".detailed_status")
	f.StatusField(".id")
	f.StatusField(".label_fingerprint")
	f.StatusField(".self_link")
	f.StatusField(".shared_secret_hash")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".cipher_suite")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".status")

	f.FilterSpec = func(in *pb.VpnTunnel) {
		if in.Region == nil {
			in.Region = direct.LazyPtr("")
		}
	}

	return f
}
