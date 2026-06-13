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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeNetworkPeeringSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.ComputeNetworkPeeringSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkPeeringSpec{}
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ExportSubnetRoutesWithPublicIP = in.ExportSubnetRoutesWithPublicIp
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ImportSubnetRoutesWithPublicIP = in.ImportSubnetRoutesWithPublicIp

	if in.GetNetwork() != "" {
		out.PeerNetworkRef = krm.ComputeNetworkRef{External: in.GetNetwork()}
	}

	out.ResourceID = in.Name
	out.StackType = in.StackType
	return out
}

func ComputeNetworkPeeringSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkPeeringSpec) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ExportSubnetRoutesWithPublicIp = in.ExportSubnetRoutesWithPublicIP
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ImportSubnetRoutesWithPublicIp = in.ImportSubnetRoutesWithPublicIP

	if in.PeerNetworkRef.External != "" {
		out.Network = &in.PeerNetworkRef.External
	}

	out.Name = in.ResourceID
	out.StackType = in.StackType
	return out
}

func ComputeNetworkPeeringStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.ComputeNetworkPeeringStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkPeeringStatus{}
	out.State = in.State
	out.StateDetails = in.StateDetails
	return out
}

func ComputeNetworkPeeringStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkPeeringStatus) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	out.State = in.State
	out.StateDetails = in.StateDetails
	return out
}
