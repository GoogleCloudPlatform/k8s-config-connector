// Copyright 2026 Google LLC
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
// proto.message: google.cloud.networkmanagement.v1.VpcFlowLogsConfig
// krm.group: networkmanagement.cnrm.cloud.google.com
// krm.kind: NetworkManagementVPCFlowLogsConfig

package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(NetworkManagementVpcFlowLogsConfigFuzzer())
}

func NetworkManagementVpcFlowLogsConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.VpcFlowLogsConfig{},
		NetworkManagementVPCFlowLogsConfigSpec_FromProto, NetworkManagementVPCFlowLogsConfigSpec_ToProto,
		NetworkManagementVPCFlowLogsConfigObservedState_FromProto, NetworkManagementVPCFlowLogsConfigObservedState_ToProto,
	)

	f.FilterSpec = func(in *pb.VpcFlowLogsConfig) {
		if in.State == nil || *in.State == pb.VpcFlowLogsConfig_STATE_UNSPECIFIED {
			val := pb.VpcFlowLogsConfig_ENABLED
			in.State = &val
		}
		if in.AggregationInterval == nil || *in.AggregationInterval == pb.VpcFlowLogsConfig_AGGREGATION_INTERVAL_UNSPECIFIED {
			val := pb.VpcFlowLogsConfig_INTERVAL_5_SEC
			in.AggregationInterval = &val
		}
		if in.Metadata == nil || *in.Metadata == pb.VpcFlowLogsConfig_METADATA_UNSPECIFIED {
			val := pb.VpcFlowLogsConfig_INCLUDE_ALL_METADATA
			in.Metadata = &val
		}
	}

	f.FilterStatus = func(in *pb.VpcFlowLogsConfig) {
		if in.TargetResourceState != nil && *in.TargetResourceState == pb.VpcFlowLogsConfig_TARGET_RESOURCE_STATE_UNSPECIFIED {
			val := pb.VpcFlowLogsConfig_TARGET_RESOURCE_EXISTS
			in.TargetResourceState = &val
		}
	}

	f.SpecField(".description")
	f.SpecField(".state")
	f.SpecField(".aggregation_interval")
	f.SpecField(".flow_sampling")
	f.SpecField(".metadata")
	f.SpecField(".metadata_fields")
	f.SpecField(".filter_expr")
	f.SpecField(".interconnect_attachment")
	f.SpecField(".vpn_tunnel")
	f.SpecField(".labels")

	f.StatusField(".target_resource_state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")

	// Fields not exposed on KRM Spec/Status
	f.Unimplemented_NotYetTriaged(".cross_project_metadata")
	f.Unimplemented_NotYetTriaged(".subnet")
	f.Unimplemented_NotYetTriaged(".network")

	return f
}
