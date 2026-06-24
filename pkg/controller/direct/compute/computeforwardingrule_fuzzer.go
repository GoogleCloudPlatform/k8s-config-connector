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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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

	// Field Comparison & Mapping Documentation:
	//
	// KRM Spec Fields:
	// - allPorts                      -> f.SpecField(".all_ports")
	// - allowGlobalAccess             -> f.SpecField(".allow_global_access")
	// - allowPscGlobalAccess          -> f.SpecField(".allow_psc_global_access")
	// - backendServiceRef             -> f.SpecField(".backend_service")
	// - description                   -> f.SpecField(".description")
	// - ipAddress                     -> f.SpecField(".I_p_address") (Protobuf field name is I_p_address)
	// - ipProtocol                    -> f.SpecField(".I_p_protocol") (Protobuf field name is I_p_protocol)
	// - ipVersion                     -> f.SpecField(".ip_version")
	// - isMirroringCollector          -> f.SpecField(".is_mirroring_collector")
	// - loadBalancingScheme           -> f.SpecField(".load_balancing_scheme")
	// - location                      -> f.Unimplemented_Identity(".region") (part of Resource Identity/URL)
	// - metadataFilters               -> f.SpecField(".metadata_filters")
	// - networkRef                    -> f.SpecField(".network")
	// - networkTier                   -> f.SpecField(".network_tier")
	// - noAutomateDnsZone             -> f.SpecField(".no_automate_dns_zone")
	// - portRange                     -> f.SpecField(".port_range")
	// - ports                         -> f.SpecField(".ports")
	// - resourceID                    -> f.Unimplemented_Identity(".name") (part of Resource Identity/URL)
	// - serviceDirectoryRegistrations -> f.SpecField(".service_directory_registrations") (ServiceDirectoryRegion is filtered/cleared because it's not in KRM)
	// - serviceLabel                  -> f.SpecField(".service_label")
	// - sourceIpRanges                -> f.SpecField(".source_ip_ranges")
	// - subnetworkRef                 -> f.SpecField(".subnetwork")
	// - target                        -> f.SpecField(".target") (fuzzed as SpecField, filtered to valid target URL, ignored in Status to avoid collision)
	//
	// KRM Status Fields:
	// - baseForwardingRule            -> f.StatusField(".base_forwarding_rule")
	// - creationTimestamp             -> f.StatusField(".creation_timestamp")
	// - labelFingerprint              -> f.StatusField(".label_fingerprint")
	// - pscConnectionId               -> f.StatusField(".psc_connection_id")
	// - pscConnectionStatus           -> f.StatusField(".psc_connection_status")
	// - selfLink                      -> f.StatusField(".self_link")
	// - serviceName                   -> f.StatusField(".service_name")
	// - target                        -> Ignored in Status (listed as f.SpecField to allow correct Spec roundtrip)

	// Spec fields
	f.SpecField(".I_p_address")
	f.SpecField(".I_p_protocol")
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
	f.SpecField(".service_directory_registrations")
	f.SpecField(".service_label")
	f.SpecField(".source_ip_ranges")
	f.SpecField(".subnetwork")
	f.SpecField(".target")

	// Status fields
	f.StatusField(".base_forwarding_rule")
	f.StatusField(".creation_timestamp")
	f.StatusField(".label_fingerprint")
	f.StatusField(".psc_connection_id")
	f.StatusField(".psc_connection_status")
	f.StatusField(".self_link")
	f.StatusField(".service_name")

	// Unimplemented fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".region")

	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".fingerprint")

	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".ip_collection")
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_testing_percentage")
	f.Unimplemented_NotYetTriaged(".self_link_with_id")
	f.Unimplemented_NotYetTriaged(".external_managed_backend_bucket_migration_state")

	// Fuzzer filters for non-trivial fields
	f.FilterSpec = func(in *pb.ForwardingRule) {
		if in.Target != nil {
			// Set target to a valid target URL so reference mapper doesn't fail
			in.Target = direct.LazyPtr("projects/p/global/targetHttpProxies/t")
		}
		for _, reg := range in.ServiceDirectoryRegistrations {
			if reg != nil {
				// Clear field not mapped in KRM
				reg.ServiceDirectoryRegion = nil
			}
		}
	}
	f.FilterStatus = f.FilterSpec

	return f
}
