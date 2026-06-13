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
// proto.message: google.cloud.compute.v1.RouterNat
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(computeRouterNATFuzzer())
}

func computeRouterNATFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.RouterNat, krm.ComputeRouterNATSpec](&pb.RouterNat{},
		ComputeRouterNATSpec_v1beta1_FromProto, ComputeRouterNATSpec_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".drain_nat_ips")
	f.SpecField(".enable_dynamic_port_allocation")
	f.SpecField(".enable_endpoint_independent_mapping")
	f.SpecField(".icmp_idle_timeout_sec")
	f.SpecField(".log_config")
	f.SpecField(".log_config.enable")
	f.SpecField(".log_config.filter")
	f.SpecField(".max_ports_per_vm")
	f.SpecField(".min_ports_per_vm")
	f.SpecField(".nat_ip_allocate_option")
	f.SpecField(".nat_ips")
	f.SpecField(".rules")
	f.SpecField(".rules[].action")
	f.SpecField(".rules[].action.source_nat_active_ips")
	f.SpecField(".rules[].action.source_nat_drain_ips")
	f.SpecField(".rules[].description")
	f.SpecField(".rules[].match")
	f.SpecField(".rules[].rule_number")
	f.SpecField(".source_subnetwork_ip_ranges_to_nat")
	f.SpecField(".subnetworks")
	f.SpecField(".subnetworks[].secondary_ip_range_names")
	f.SpecField(".subnetworks[].source_ip_ranges_to_nat")
	f.SpecField(".subnetworks[].name")
	f.SpecField(".tcp_established_idle_timeout_sec")
	f.SpecField(".tcp_time_wait_timeout_sec")
	f.SpecField(".tcp_transitory_idle_timeout_sec")
	f.SpecField(".udp_idle_timeout_sec")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields in the Proto message that are not in KRM schema
	f.Unimplemented_NotYetTriaged(".auto_network_tier")
	f.Unimplemented_NotYetTriaged(".endpoint_types")
	f.Unimplemented_NotYetTriaged(".nat64_subnetworks")
	f.Unimplemented_NotYetTriaged(".rules[].action.source_nat_active_ranges")
	f.Unimplemented_NotYetTriaged(".rules[].action.source_nat_drain_ranges")
	f.Unimplemented_NotYetTriaged(".source_subnetwork_ip_ranges_to_nat64")
	f.Unimplemented_NotYetTriaged(".type")

	f.FilterSpec = func(in *pb.RouterNat) {
		if in.LogConfig != nil {
			if in.LogConfig.Enable == nil {
				val := false
				in.LogConfig.Enable = &val
			}
			if in.LogConfig.Filter == nil {
				val := ""
				in.LogConfig.Filter = &val
			}
		}
		if in.NatIpAllocateOption == nil {
			val := ""
			in.NatIpAllocateOption = &val
		}
		if in.SourceSubnetworkIpRangesToNat == nil {
			val := ""
			in.SourceSubnetworkIpRangesToNat = &val
		}
		for _, rule := range in.Rules {
			if rule.RuleNumber == nil {
				val := uint32(0)
				rule.RuleNumber = &val
			}
			if rule.Match == nil {
				val := ""
				rule.Match = &val
			}
		}
	}

	return f
}
