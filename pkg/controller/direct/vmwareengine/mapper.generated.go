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

// +generated:mapper
// krm.group: vmwareengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.vmwareengine.v1

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	krmvmwareenginev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Hcx_FromProto(mapCtx *direct.MapContext, in *pb.Hcx) *krm.Hcx {
	if in == nil {
		return nil
	}
	out := &krm.Hcx{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FQDN = direct.LazyPtr(in.GetFqdn())
	return out
}
func Hcx_ToProto(mapCtx *direct.MapContext, in *krm.Hcx) *pb.Hcx {
	if in == nil {
		return nil
	}
	out := &pb.Hcx{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.Hcx_State](mapCtx, in.State)
	out.Fqdn = direct.ValueOf(in.FQDN)
	return out
}
func NetworkPolicy_NetworkService_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krm.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicy_NetworkService{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkService_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicy_NetworkService) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	out.Enabled = direct.ValueOf(in.Enabled)
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krm.NetworkPolicy_NetworkServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkPolicy_NetworkServiceObservedState{}
	// MISSING: Enabled
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkPolicy_NetworkServiceObservedState) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	// MISSING: Enabled
	out.State = direct.Enum_ToProto[pb.NetworkPolicy_NetworkService_State](mapCtx, in.State)
	return out
}
func NodeTypeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeTypeConfig) *krm.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeTypeConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.CustomCoreCount = direct.LazyPtr(in.GetCustomCoreCount())
	return out
}
func NodeTypeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeTypeConfig) *pb.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeTypeConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.CustomCoreCount = direct.ValueOf(in.CustomCoreCount)
	return out
}
func Nsx_FromProto(mapCtx *direct.MapContext, in *pb.Nsx) *krm.Nsx {
	if in == nil {
		return nil
	}
	out := &krm.Nsx{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FQDN = direct.LazyPtr(in.GetFqdn())
	return out
}
func Nsx_ToProto(mapCtx *direct.MapContext, in *krm.Nsx) *pb.Nsx {
	if in == nil {
		return nil
	}
	out := &pb.Nsx{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.Nsx_State](mapCtx, in.State)
	out.Fqdn = direct.ValueOf(in.FQDN)
	return out
}
func StretchedClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.StretchedClusterConfig) *krm.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.StretchedClusterConfig{}
	out.PreferredLocation = direct.LazyPtr(in.GetPreferredLocation())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	return out
}
func StretchedClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.StretchedClusterConfig) *pb.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.StretchedClusterConfig{}
	out.PreferredLocation = direct.ValueOf(in.PreferredLocation)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	return out
}
func VMwareEngineExternalAccessRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineExternalAccessRuleSpec) *pb.ExternalAccessRule {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAccessRule{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Priority = direct.ValueOf(in.Priority)
	out.Action = direct.Enum_ToProto[pb.ExternalAccessRule_Action](mapCtx, in.Action)
	out.IpProtocol = direct.ValueOf(in.IPProtocol)
	out.SourceIpRanges = direct.Slice_ToProto(mapCtx, in.SourceIPRanges, ExternalAccessRule_IPRange_ToProto)
	out.SourcePorts = in.SourcePorts
	out.DestinationIpRanges = direct.Slice_ToProto(mapCtx, in.DestinationIPRanges, ExternalAccessRule_IPRange_ToProto)
	out.DestinationPorts = in.DestinationPorts
	// MISSING: Uid
	return out
}
func VMwareEngineExternalAddressSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krmvmwareenginev1beta1.VMwareEngineExternalAddressSpec {
	if in == nil {
		return nil
	}
	out := &krmvmwareenginev1beta1.VMwareEngineExternalAddressSpec{}
	// MISSING: Name
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func VMwareEngineExternalAddressSpec_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1beta1.VMwareEngineExternalAddressSpec) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	out.InternalIp = direct.ValueOf(in.InternalIP)
	// MISSING: Uid
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
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Uid
	return out
}
func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_Type](mapCtx, in.Type)
	// MISSING: Uid
	return out
}
func VMwareEnginePrivateCloudSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.VMwareEnginePrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEnginePrivateCloudSpec{}
	// MISSING: Name
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.ManagementCluster = PrivateCloud_ManagementCluster_FromProto(mapCtx, in.GetManagementCluster())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Uid
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func VMwareEnginePrivateCloudSpec_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEnginePrivateCloudSpec) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	// MISSING: Name
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.ManagementCluster = PrivateCloud_ManagementCluster_ToProto(mapCtx, in.ManagementCluster)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Uid
	out.Type = direct.Enum_ToProto[pb.PrivateCloud_Type](mapCtx, in.Type)
	return out
}
func Vcenter_FromProto(mapCtx *direct.MapContext, in *pb.Vcenter) *krm.Vcenter {
	if in == nil {
		return nil
	}
	out := &krm.Vcenter{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FQDN = direct.LazyPtr(in.GetFqdn())
	return out
}
func Vcenter_ToProto(mapCtx *direct.MapContext, in *krm.Vcenter) *pb.Vcenter {
	if in == nil {
		return nil
	}
	out := &pb.Vcenter{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.Vcenter_State](mapCtx, in.State)
	out.Fqdn = direct.ValueOf(in.FQDN)
	return out
}
func VmwareEngineNetwork_VPCNetwork_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VPCNetwork {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VPCNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VPCNetwork_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VPCNetwork) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VPCNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krm.VmwareEngineNetwork_VPCNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareEngineNetwork_VPCNetworkObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func VmwareEngineNetwork_VPCNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareEngineNetwork_VPCNetworkObservedState) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_VpcNetwork_Type](mapCtx, in.Type)
	out.Network = direct.ValueOf(in.Network)
	return out
}
