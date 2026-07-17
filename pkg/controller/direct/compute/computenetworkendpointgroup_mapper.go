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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeNetworkEndpointGroupSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpointGroup) *krm.ComputeNetworkEndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkEndpointGroupSpec{}
	// MISSING: Annotations
	// MISSING: AppEngine
	// MISSING: CloudFunction
	// MISSING: CloudRun
	// MISSING: CreationTimestamp
	if in.DefaultPort != nil {
		val := int64(*in.DefaultPort)
		out.DefaultPort = &val
	}
	out.Description = in.Description
	// MISSING: ID
	// MISSING: Kind
	// MISSING: Name
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkEndpointType = in.NetworkEndpointType
	// MISSING: PSCData
	// MISSING: PSCTargetService
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: Size
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &krm.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	// MISSING: Zone
	return out
}

func ComputeNetworkEndpointGroupSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkEndpointGroupSpec) *pb.NetworkEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpointGroup{}
	// MISSING: Annotations
	// MISSING: AppEngine
	// MISSING: CloudFunction
	// MISSING: CloudRun
	// MISSING: CreationTimestamp
	if in.DefaultPort != nil {
		val := int32(*in.DefaultPort)
		out.DefaultPort = &val
	}
	out.Description = in.Description
	// MISSING: ID
	// MISSING: Kind
	// MISSING: Name
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.NetworkEndpointType = in.NetworkEndpointType
	// MISSING: PSCData
	// MISSING: PSCTargetService
	// MISSING: Region
	// MISSING: SelfLink
	// MISSING: Size
	if in.SubnetworkRef != nil {
		out.Subnetwork = &in.SubnetworkRef.External
	}
	// MISSING: Zone
	return out
}

func ComputeNetworkEndpointGroupStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpointGroup) *krm.ComputeNetworkEndpointGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkEndpointGroupStatus{}
	// MISSING: Annotations
	// MISSING: AppEngine
	// MISSING: CloudFunction
	// MISSING: CloudRun
	// MISSING: CreationTimestamp
	// MISSING: DefaultPort
	// MISSING: Description
	// MISSING: ID
	// MISSING: Kind
	// MISSING: Name
	// MISSING: Network
	// MISSING: NetworkEndpointType
	// MISSING: PSCData
	// MISSING: PSCTargetService
	// MISSING: Region
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		val := int64(*in.Size)
		out.Size = &val
	}
	// MISSING: Subnetwork
	// MISSING: Zone
	return out
}

func ComputeNetworkEndpointGroupStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkEndpointGroupStatus) *pb.NetworkEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpointGroup{}
	// MISSING: Annotations
	// MISSING: AppEngine
	// MISSING: CloudFunction
	// MISSING: CloudRun
	// MISSING: CreationTimestamp
	// MISSING: DefaultPort
	// MISSING: Description
	// MISSING: ID
	// MISSING: Kind
	// MISSING: Name
	// MISSING: Network
	// MISSING: NetworkEndpointType
	// MISSING: PSCData
	// MISSING: PSCTargetService
	// MISSING: Region
	out.SelfLink = in.SelfLink
	if in.Size != nil {
		val := int32(*in.Size)
		out.Size = &val
	}
	// MISSING: Subnetwork
	// MISSING: Zone
	return out
}
