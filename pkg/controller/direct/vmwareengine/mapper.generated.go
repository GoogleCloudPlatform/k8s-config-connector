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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VMwareEngineNetworkPeeringObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.VMwareEngineNetworkPeeringObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkPeeringObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	// MISSING: PeerMtu
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	// MISSING: VmwareEngineNetwork
	return out
}
func VMwareEngineNetworkPeeringObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkPeeringObservedState) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.NetworkPeering_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	// MISSING: PeerMtu
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	// MISSING: VmwareEngineNetwork
	return out
}
func VMwareEngineNetworkPeeringSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.VMwareEngineNetworkPeeringSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkPeeringSpec{}
	// MISSING: Name
	out.PeerNetwork = direct.LazyPtr(in.GetPeerNetwork())
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIP = in.ExportCustomRoutesWithPublicIp
	out.ImportCustomRoutesWithPublicIP = in.ImportCustomRoutesWithPublicIp
	// MISSING: PeerMtu
	// (near miss): "PeerMtu" vs "PeerMTU"
	out.PeerNetworkType = direct.Enum_FromProto(mapCtx, in.GetPeerNetworkType())
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func VMwareEngineNetworkPeeringSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkPeeringSpec) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	out.PeerNetwork = VMwareEngineNetworkPeeringSpec_PeerNetwork_ToProto(mapCtx, in.PeerNetwork)
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIp = in.ExportCustomRoutesWithPublicIP
	out.ImportCustomRoutesWithPublicIp = in.ImportCustomRoutesWithPublicIP
	// MISSING: PeerMtu
	// (near miss): "PeerMtu" vs "PeerMTU"
	out.PeerNetworkType = direct.Enum_ToProto[pb.NetworkPeering_PeerNetworkType](mapCtx, in.PeerNetworkType)
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	out.Description = direct.ValueOf(in.Description)
	return out
}
func VMwareEngineNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krm.VMwareEngineNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VpcNetworks
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VpcNetworks
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_Type](mapCtx, in.Type)
	// MISSING: Uid
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func VmwareEngineNetwork_VpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetwork_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VpcNetwork) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VpcNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VpcNetworkObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VpcNetworkObservedState) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_VpcNetwork_Type](mapCtx, in.Type)
	out.Network = direct.ValueOf(in.Network)
	return out
}
