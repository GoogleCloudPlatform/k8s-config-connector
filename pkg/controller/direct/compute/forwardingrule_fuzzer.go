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
	f.SpecFields.Insert(".ip_address")
	f.SpecFields.Insert(".ip_protocol")
	f.SpecFields.Insert(".all_ports")
	f.SpecFields.Insert(".allow_global_access")
	f.SpecFields.Insert(".allow_psc_global_access")
	f.SpecFields.Insert(".backend_service")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".ip_version")
	f.SpecFields.Insert(".is_mirroring_collector")
	f.SpecFields.Insert(".load_balancing_scheme")
	f.SpecFields.Insert(".metadata_filters")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".network_tier")
	f.SpecFields.Insert(".no_automate_dns_zone")
	f.SpecFields.Insert(".port_range")
	f.SpecFields.Insert(".ports")
	f.SpecFields.Insert(".service_label")
	f.SpecFields.Insert(".source_ip_ranges")
	f.SpecFields.Insert(".subnetwork")
	// f.SpecFields.Insert(".target") // Remove from SpecFields

	// Status fields
	f.StatusFields.Insert(".base_forwarding_rule")
	f.StatusFields.Insert(".creation_timestamp")
	f.StatusFields.Insert(".label_fingerprint")
	f.StatusFields.Insert(".psc_connection_id")
	f.StatusFields.Insert(".psc_connection_status")
	f.StatusFields.Insert(".self_link")
	f.StatusFields.Insert(".service_name")

	// Unimplemented fields
	f.UnimplementedFields.Insert(".I_p_protocol")
	f.UnimplementedFields.Insert(".I_p_address")
	f.UnimplementedFields.Insert(".id")
	f.UnimplementedFields.Insert(".ip_collection")
	f.UnimplementedFields.Insert(".kind")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".region")
	f.UnimplementedFields.Insert(".fingerprint")
	f.UnimplementedFields.Insert(".service_directory_registrations") // Add to UnimplementedFields
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_testing_percentage")
	f.Unimplemented_NotYetTriaged(".self_link_with_id")

	f.Unimplemented_NotYetTriaged(".target")
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_state")
	return f
}
