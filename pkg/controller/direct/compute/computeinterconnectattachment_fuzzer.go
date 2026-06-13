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
// proto.message: google.cloud.compute.v1.InterconnectAttachment
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInterconnectAttachmentFuzzer())
}

func computeInterconnectAttachmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InterconnectAttachment{},
		ComputeInterconnectAttachmentSpec_v1beta1_FromProto, ComputeInterconnectAttachmentSpec_v1beta1_ToProto,
		ComputeInterconnectAttachmentStatus_v1beta1_FromProto, ComputeInterconnectAttachmentStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".admin_enabled")
	f.SpecField(".bandwidth")
	f.SpecField(".candidate_subnets")
	f.SpecField(".description")
	f.SpecField(".edge_availability_domain")
	f.SpecField(".encryption")
	f.SpecField(".interconnect")
	f.SpecField(".ipsec_internal_addresses")
	f.SpecField(".mtu")
	f.SpecField(".region")
	f.SpecField(".router")
	f.SpecField(".type")
	f.SpecField(".vlan_tag8021q")

	// Status fields
	f.StatusField(".cloud_router_ip_address")
	f.StatusField(".creation_timestamp")
	f.StatusField(".customer_router_ip_address")
	f.StatusField(".google_reference_id")
	f.StatusField(".pairing_key")
	f.StatusField(".partner_asn")
	f.StatusField(".private_interconnect_info")
	f.StatusField(".private_interconnect_info.tag8021q")
	f.StatusField(".self_link")
	f.StatusField(".state")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Labels & Annotations
	f.Unimplemented_LabelsAnnotations(".label_fingerprint")
	f.Unimplemented_LabelsAnnotations(".labels")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".attachment_group")
	f.Unimplemented_NotYetTriaged(".candidate_ipv6_subnets")
	f.Unimplemented_NotYetTriaged(".cloud_router_ipv6_address")
	f.Unimplemented_NotYetTriaged(".cloud_router_ipv6_interface_id")
	f.Unimplemented_NotYetTriaged(".configuration_constraints")
	f.Unimplemented_NotYetTriaged(".customer_router_ipv6_address")
	f.Unimplemented_NotYetTriaged(".customer_router_ipv6_interface_id")
	f.Unimplemented_NotYetTriaged(".dataplane_version")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".operational_status")
	f.Unimplemented_NotYetTriaged(".partner_metadata")
	f.Unimplemented_NotYetTriaged(".remote_service")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".stack_type")
	f.Unimplemented_NotYetTriaged(".subnet_length")

	return f
}
