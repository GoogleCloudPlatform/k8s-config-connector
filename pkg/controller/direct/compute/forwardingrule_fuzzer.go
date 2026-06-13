// Copyright 2025 Google LLC
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
// proto.message: google.cloud.compute.v1.ForwardingRule
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeForwardingRuleFuzzer())
}

func computeForwardingRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ForwardingRule{},
		ComputeForwardingRuleSpec_v1beta1_FromProto, ComputeForwardingRuleSpec_v1beta1_ToProto,
		ComputeForwardingRuleStatus_v1beta1_FromProto, ComputeForwardingRuleStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".ip_address")
	f.SpecField(".ip_protocol")
	f.SpecField(".all_ports")
	f.SpecField(".allow_global_access")
	f.SpecField(".allow_psc_global_access")
	f.SpecField(".backend_service")
	f.SpecField(".description")
	f.SpecField(".ip_version")
	f.SpecField(".is_mirroring_collector")
	f.SpecField(".load_balancing_scheme")
	f.SpecField(".metadata_filters")
	f.SpecField(".network")
	f.SpecField(".network_tier")
	f.SpecField(".no_automate_dns_zone")
	f.SpecField(".port_range")
	f.SpecField(".ports")
	f.SpecField(".service_label")
	f.SpecField(".source_ip_ranges")
	f.SpecField(".subnetwork")

	// Status fields
	f.StatusField(".base_forwarding_rule")
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".psc_connection_id")
	f.StatusField(".psc_connection_status")
	f.StatusField(".self_link")
	f.StatusField(".service_name")
	f.StatusField(".target")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".region")

	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".fingerprint")
	f.Unimplemented_Internal(".I_p_protocol")
	f.Unimplemented_Internal(".I_p_address")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".ip_collection")
	f.Unimplemented_NotYetTriaged(".service_directory_registrations")
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_testing_percentage")
	f.Unimplemented_NotYetTriaged(".self_link_with_id")
	f.Unimplemented_NotYetTriaged(".target")
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_state")
	return f
}
