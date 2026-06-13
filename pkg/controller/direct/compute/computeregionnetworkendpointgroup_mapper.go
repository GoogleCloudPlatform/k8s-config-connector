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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeRegionNetworkEndpointGroupSpec_v1beta1_FromProto maps Region (which is string in KRM but *string in proto) and references.
func ComputeRegionNetworkEndpointGroupSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpointGroup) *krm.ComputeRegionNetworkEndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionNetworkEndpointGroupSpec{}
	out.CloudFunction = RegionnetworkendpointgroupCloudFunction_v1beta1_FromProto(mapCtx, in.GetCloudFunction())
	out.CloudRun = RegionnetworkendpointgroupCloudRun_v1beta1_FromProto(mapCtx, in.GetCloudRun())
	out.Description = in.Description
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkEndpointType = in.NetworkEndpointType
	out.PSCTargetService = in.PscTargetService
	out.Region = direct.ValueOf(in.Region)
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &krm.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}

// ComputeRegionNetworkEndpointGroupSpec_v1beta1_ToProto maps Region (which is string in KRM but *string in proto) and references.
func ComputeRegionNetworkEndpointGroupSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionNetworkEndpointGroupSpec) *pb.NetworkEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpointGroup{}
	out.CloudFunction = RegionnetworkendpointgroupCloudFunction_v1beta1_ToProto(mapCtx, in.CloudFunction)
	out.CloudRun = RegionnetworkendpointgroupCloudRun_v1beta1_ToProto(mapCtx, in.CloudRun)
	out.Description = in.Description
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.NetworkEndpointType = in.NetworkEndpointType
	out.PscTargetService = in.PSCTargetService
	out.Region = direct.LazyPtr(in.Region)
	if in.SubnetworkRef != nil {
		out.Subnetwork = &in.SubnetworkRef.External
	}
	return out
}
