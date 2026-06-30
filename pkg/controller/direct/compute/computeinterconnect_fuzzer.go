// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.Interconnect
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeInterconnectFuzzer())
}

func computeInterconnectFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Interconnect{},
		ComputeInterconnectSpec_v1alpha1_FromProto, ComputeInterconnectSpec_v1alpha1_ToProto,
		ComputeInterconnectObservedState_v1alpha1_FromProto, ComputeInterconnectObservedState_v1alpha1_ToProto,
	)

	// Field comparison: ComputeInterconnect Spec vs pb.Interconnect Proto
	// - Spec.Parent            is handled via the GCP URI/ID, not mapped to a proto body field directly
	// - Spec.ResourceID        maps to proto field .name (handled as Unimplemented_Identity)
	// - Spec.AdminEnabled      maps to proto field .admin_enabled
	// - Spec.CustomerName      maps to proto field .customer_name
	// - Spec.Description       maps to proto field .description
	// - Spec.InterconnectType  maps to proto field .interconnect_type
	// - Spec.LabelFingerprint  maps to proto field .label_fingerprint
	// - Spec.Labels            maps to proto field .labels
	// - Spec.LinkType          maps to proto field .link_type
	// - Spec.Location          maps to proto field .location
	// - Spec.Macsec            maps to proto field .macsec
	// - Spec.MacsecEnabled     maps to proto field .macsec_enabled
	// - Spec.NocContactEmail   maps to proto field .noc_contact_email
	// - Spec.RemoteLocation    maps to proto field .remote_location
	// - Spec.RequestedFeatures maps to proto field .requested_features
	// - Spec.RequestedLinkCount maps to proto field .requested_link_count
	//
	// Spec fields
	f.SpecField(".admin_enabled")
	f.SpecField(".customer_name")
	f.SpecField(".description")
	f.SpecField(".interconnect_type")
	f.SpecField(".label_fingerprint")
	f.SpecField(".labels")
	f.SpecField(".link_type")
	f.SpecField(".location")
	f.SpecField(".macsec")
	f.SpecField(".macsec_enabled")
	f.SpecField(".noc_contact_email")
	f.SpecField(".remote_location")
	f.SpecField(".requested_features")
	f.SpecField(".requested_link_count")

	// Status fields
	f.StatusField(".available_features")
	f.StatusField(".circuit_infos")
	f.StatusField(".creation_timestamp")
	f.StatusField(".expected_outages")
	f.StatusField(".google_ip_address")
	f.StatusField(".google_reference_id")
	f.StatusField(".id")
	f.StatusField(".interconnect_attachments")
	f.StatusField(".kind")
	f.StatusField(".operational_status")
	f.StatusField(".peer_ip_address")
	f.StatusField(".provisioned_link_count")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".self_link")
	f.StatusField(".state")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".interconnect_groups")
	f.Unimplemented_NotYetTriaged(".aai_enabled")
	f.Unimplemented_NotYetTriaged(".application_aware_interconnect")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".wire_groups")
	f.Unimplemented_NotYetTriaged(".subzone")

	return f
}
