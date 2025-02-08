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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func NetworkPeering_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPeering{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.PeerNetwork = direct.LazyPtr(in.GetPeerNetwork())
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIP = in.ExportCustomRoutesWithPublicIp
	out.ImportCustomRoutesWithPublicIP = in.ImportCustomRoutesWithPublicIp
	// MISSING: State
	// MISSING: StateDetails
	out.PeerMtu = direct.LazyPtr(in.GetPeerMtu())
	out.PeerNetworkType = direct.Enum_FromProto(mapCtx, in.GetPeerNetworkType())
	// MISSING: Uid
	out.VmwareEngineNetwork = direct.LazyPtr(in.GetVmwareEngineNetwork())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func NetworkPeering_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPeering) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.PeerNetwork = direct.ValueOf(in.PeerNetwork)
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIp = in.ExportCustomRoutesWithPublicIP
	out.ImportCustomRoutesWithPublicIp = in.ImportCustomRoutesWithPublicIP
	// MISSING: State
	// MISSING: StateDetails
	out.PeerMtu = direct.ValueOf(in.PeerMtu)
	out.PeerNetworkType = direct.Enum_ToProto[pb.NetworkPeering_PeerNetworkType](mapCtx, in.PeerNetworkType)
	// MISSING: Uid
	out.VmwareEngineNetwork = direct.ValueOf(in.VmwareEngineNetwork)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func NetworkPeeringObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.NetworkPeeringObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPeeringObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
func NetworkPeeringObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPeeringObservedState) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	out.State = direct.Enum_ToProto[pb.NetworkPeering_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
func VmwareengineNetworkPeeringObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.VmwareengineNetworkPeeringObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNetworkPeeringObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
func VmwareengineNetworkPeeringObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNetworkPeeringObservedState) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
func VmwareengineNetworkPeeringSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.VmwareengineNetworkPeeringSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineNetworkPeeringSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
func VmwareengineNetworkPeeringSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineNetworkPeeringSpec) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PeerNetwork
	// MISSING: ExportCustomRoutes
	// MISSING: ImportCustomRoutes
	// MISSING: ExchangeSubnetRoutes
	// MISSING: ExportCustomRoutesWithPublicIP
	// MISSING: ImportCustomRoutesWithPublicIP
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: PeerMtu
	// MISSING: PeerNetworkType
	// MISSING: Uid
	// MISSING: VmwareEngineNetwork
	// MISSING: Description
	return out
}
