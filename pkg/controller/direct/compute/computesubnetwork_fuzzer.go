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
// proto.message: google.cloud.compute.v1.Subnetwork
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeSubnetworkFuzzer())
}

func computeSubnetworkFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Subnetwork{},
		ComputeSubnetworkSpec_v1beta1_FromProto, ComputeSubnetworkSpec_v1beta1_ToProto,
		ComputeSubnetworkStatus_v1beta1_FromProto, ComputeSubnetworkStatus_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".ip_cidr_range")
	f.SpecField(".ipv6_access_type")
	f.SpecField(".log_config")
	f.SpecField(".log_config.aggregation_interval")
	f.SpecField(".log_config.filter_expr")
	f.SpecField(".log_config.flow_sampling")
	f.SpecField(".log_config.metadata")
	f.SpecField(".log_config.metadata_fields")
	f.SpecField(".network")
	f.SpecField(".private_ip_google_access")
	f.SpecField(".private_ipv6_google_access")
	f.SpecField(".purpose")
	f.SpecField(".region")
	f.SpecField(".role")
	f.SpecField(".secondary_ip_ranges")
	f.SpecField(".secondary_ip_ranges[].ip_cidr_range")
	f.SpecField(".secondary_ip_ranges[].range_name")
	f.SpecField(".stack_type")

	// Status fields
	f.StatusField(".creation_timestamp")
	f.StatusField(".external_ipv6_prefix")
	f.StatusField(".fingerprint")
	f.StatusField(".gateway_address")
	f.StatusField(".internal_ipv6_prefix")
	f.StatusField(".ipv6_cidr_range")
	f.StatusField(".self_link")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".enable_flow_logs")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".ip_collection")
	f.Unimplemented_NotYetTriaged(".ipv6_gce_endpoint")
	f.Unimplemented_NotYetTriaged(".kind")
	f.Unimplemented_NotYetTriaged(".params")
	f.Unimplemented_NotYetTriaged(".reserved_internal_range")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".system_reserved_external_ipv6_ranges")
	f.Unimplemented_NotYetTriaged(".system_reserved_internal_ipv6_ranges")

	f.Unimplemented_NotYetTriaged(".allow_subnet_cidr_routes_overlap")
	f.Unimplemented_NotYetTriaged(".utilization_details")
	f.Unimplemented_NotYetTriaged(".log_config.enable")
	f.Unimplemented_NotYetTriaged(".secondary_ip_ranges[].reserved_internal_range")

	f.FilterSpec = func(in *pb.Subnetwork) {
		if in.LogConfig != nil && in.LogConfig.AggregationInterval == nil && in.LogConfig.FilterExpr == nil && in.LogConfig.FlowSampling == nil && in.LogConfig.Metadata == nil && len(in.LogConfig.MetadataFields) == 0 {
			in.LogConfig = nil
		}
	}

	f.FilterStatus = func(in *pb.Subnetwork) {
		in.LogConfig = nil
		in.SecondaryIpRanges = nil
	}

	return f
}
