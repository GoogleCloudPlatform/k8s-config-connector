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
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkPolicy_NetworkService_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krmv1alpha1.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NetworkPolicy_NetworkService{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkService_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkPolicy_NetworkService) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	out.Enabled = direct.ValueOf(in.Enabled)
	// MISSING: State
	return out
}
func NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krmv1alpha1.NetworkPolicy_NetworkServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NetworkPolicy_NetworkServiceObservedState{}
	// MISSING: Enabled
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkPolicy_NetworkServiceObservedState) *pb.NetworkPolicy_NetworkService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkPolicy_NetworkService{}
	// MISSING: Enabled
	out.State = direct.Enum_ToProto[pb.NetworkPolicy_NetworkService_State](mapCtx, in.State)
	return out
}
func NodeTypeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeTypeConfig) *krmv1alpha1.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NodeTypeConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.CustomCoreCount = direct.LazyPtr(in.GetCustomCoreCount())
	return out
}
func NodeTypeConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NodeTypeConfig) *pb.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeTypeConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.CustomCoreCount = direct.ValueOf(in.CustomCoreCount)
	return out
}
func StretchedClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.StretchedClusterConfig) *krmv1alpha1.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StretchedClusterConfig{}
	out.PreferredLocation = direct.LazyPtr(in.GetPreferredLocation())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	return out
}
func StretchedClusterConfig_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.StretchedClusterConfig) *pb.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.StretchedClusterConfig{}
	out.PreferredLocation = direct.ValueOf(in.PreferredLocation)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	return out
}
func VMwareEngineExternalAccessRuleSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VMwareEngineExternalAccessRuleSpec) *pb.ExternalAccessRule {
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
func VMwareEngineExternalAddressObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krmv1alpha1.VMwareEngineExternalAddressObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VMwareEngineExternalAddressObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	return out
}
func VMwareEngineExternalAddressObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VMwareEngineExternalAddressObservedState) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.State = direct.Enum_ToProto[pb.ExternalAddress_State](mapCtx, in.State)
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	return out
}
func VMwareEngineExternalAddressSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krmv1alpha1.VMwareEngineExternalAddressSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VMwareEngineExternalAddressSpec{}
	// MISSING: Name
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func VMwareEngineExternalAddressSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VMwareEngineExternalAddressSpec) *pb.ExternalAddress {
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
func VMwareEngineNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krmv1alpha1.VMwareEngineNetworkSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VMwareEngineNetworkSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VpcNetworks
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
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
func VMwareEnginePrivateCloudSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krmv1alpha1.VMwareEnginePrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VMwareEnginePrivateCloudSpec{}
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
func VMwareEnginePrivateCloudSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VMwareEnginePrivateCloudSpec) *pb.PrivateCloud {
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
func VmwareEngineNetwork_VpcNetwork_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krmv1alpha1.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetwork_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VmwareEngineNetwork_VpcNetwork) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	// MISSING: Type
	// MISSING: Network
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krmv1alpha1.VmwareEngineNetwork_VpcNetworkObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VmwareEngineNetwork_VpcNetworkObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Network = direct.LazyPtr(in.GetNetwork())
	return out
}
func VmwareEngineNetwork_VpcNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VmwareEngineNetwork_VpcNetworkObservedState) *pb.VmwareEngineNetwork_VpcNetwork {
	if in == nil {
		return nil
	}
	out := &pb.VmwareEngineNetwork_VpcNetwork{}
	out.Type = direct.Enum_ToProto[pb.VmwareEngineNetwork_VpcNetwork_Type](mapCtx, in.Type)
	out.Network = direct.ValueOf(in.Network)
	return out
}
