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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	out.UID = direct.LazyPtr(in.GetUid())
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
	out.Uid = direct.ValueOf(in.UID)
	return out
}
func VMwareEngineNetworkPeeringSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krm.VMwareEngineNetworkPeeringSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEngineNetworkPeeringSpec{}
	// MISSING: Name
	out.PeerNetwork = VMwareEngineNetworkPeering_PeerNetwork_FromProto(mapCtx, in.GetPeerNetwork(), in.GetPeerNetworkType())
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIP = in.ExportCustomRoutesWithPublicIp
	out.ImportCustomRoutesWithPublicIP = in.ImportCustomRoutesWithPublicIp
	out.PeerMTU = direct.LazyPtr(in.PeerMtu)
	out.PeerNetworkType = direct.Enum_FromProto(mapCtx, in.GetPeerNetworkType())
	if in.GetVmwareEngineNetwork() != "" {
		out.VMwareEngineNetworkRef = &krm.VmwareEngineNetworkRef{
			External: in.GetVmwareEngineNetwork(),
		}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func VMwareEngineNetworkPeeringSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkPeeringSpec) *pb.NetworkPeering {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPeering{}
	// MISSING: Name
	out.PeerNetwork = VMwareEngineNetworkPeering_PeerNetwork_ToProto(mapCtx, in.PeerNetwork, in.PeerNetworkType)
	out.ExportCustomRoutes = in.ExportCustomRoutes
	out.ImportCustomRoutes = in.ImportCustomRoutes
	out.ExchangeSubnetRoutes = in.ExchangeSubnetRoutes
	out.ExportCustomRoutesWithPublicIp = in.ExportCustomRoutesWithPublicIP
	out.ImportCustomRoutesWithPublicIp = in.ImportCustomRoutesWithPublicIP
	out.PeerMtu = direct.ValueOf(in.PeerMTU)
	out.PeerNetworkType = direct.Enum_ToProto[pb.NetworkPeering_PeerNetworkType](mapCtx, in.PeerNetworkType)
	if in.VMwareEngineNetworkRef != nil {
		out.VmwareEngineNetwork = in.VMwareEngineNetworkRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func VMwareEngineNetworkPeering_PeerNetwork_FromProto(mapCtx *direct.MapContext, in string, inType pb.NetworkPeering_PeerNetworkType) *krm.PeerNetwork {
	if in == "" {
		return nil
	}

	switch inType {
	case pb.NetworkPeering_STANDARD:
		return &krm.PeerNetwork{
			ComputeNetworkRef: &refsv1beta1.ComputeNetworkRef{
				External: in,
			},
		}
	case pb.NetworkPeering_VMWARE_ENGINE_NETWORK:
		return &krm.PeerNetwork{
			VMwareEngineNetworkRef: &krm.VmwareEngineNetworkRef{
				External: in,
			},
		}
	default:
		mapCtx.Errorf("unknown peer network type: %v", inType)
		return nil
	}
}
func VMwareEngineNetworkPeering_PeerNetwork_ToProto(mapCtx *direct.MapContext, in *krm.PeerNetwork, inType *string) string {
	if in == nil {
		return ""
	}

	peerNetworkType := direct.Enum_ToProto[pb.NetworkPeering_PeerNetworkType](mapCtx, inType)

	switch peerNetworkType {
	case pb.NetworkPeering_STANDARD:
		return in.ComputeNetworkRef.External
	case pb.NetworkPeering_VMWARE_ENGINE_NETWORK:
		return in.VMwareEngineNetworkRef.External
	default:
		mapCtx.Errorf("unknown peer network type: %v", inType)
		return ""
	}
}
