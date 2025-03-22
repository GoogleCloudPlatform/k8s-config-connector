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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Hcx_FromProto(mapCtx *direct.MapContext, in *pb.Hcx) *krm.Hcx {
	if in == nil {
		return nil
	}
	out := &krm.Hcx{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func Hcx_ToProto(mapCtx *direct.MapContext, in *krm.Hcx) *pb.Hcx {
	if in == nil {
		return nil
	}
	out := &pb.Hcx{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func HcxObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Hcx) *krm.HcxObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HcxObservedState{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Fqdn
	return out
}
func HcxObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HcxObservedState) *pb.Hcx {
	if in == nil {
		return nil
	}
	out := &pb.Hcx{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_ToProto[pb.Hcx_State](mapCtx, in.State)
	// MISSING: Fqdn
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.ManagementCIDR = direct.LazyPtr(in.GetManagementCidr())
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: ManagementIPAddressLayoutVersion
	// MISSING: DNSServerIP
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.ManagementCidr = direct.ValueOf(in.ManagementCIDR)
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: ManagementIPAddressLayoutVersion
	// MISSING: DNSServerIP
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	// MISSING: ManagementCIDR
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// (near miss): "VmwareEngineNetworkCanonical" vs "VMwareEngineNetworkCanonical"
	out.ManagementIPAddressLayoutVersion = direct.LazyPtr(in.GetManagementIpAddressLayoutVersion())
	out.DNSServerIP = direct.LazyPtr(in.GetDnsServerIp())
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: ManagementCIDR
	// MISSING: VmwareEngineNetwork
	// MISSING: VmwareEngineNetworkCanonical
	// (near miss): "VmwareEngineNetworkCanonical" vs "VMwareEngineNetworkCanonical"
	out.ManagementIpAddressLayoutVersion = direct.ValueOf(in.ManagementIPAddressLayoutVersion)
	out.DnsServerIp = direct.ValueOf(in.DNSServerIP)
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
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func Nsx_ToProto(mapCtx *direct.MapContext, in *krm.Nsx) *pb.Nsx {
	if in == nil {
		return nil
	}
	out := &pb.Nsx{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func NsxObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Nsx) *krm.NsxObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NsxObservedState{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Fqdn
	return out
}
func NsxObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NsxObservedState) *pb.Nsx {
	if in == nil {
		return nil
	}
	out := &pb.Nsx{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_ToProto[pb.Nsx_State](mapCtx, in.State)
	// MISSING: Fqdn
	return out
}
func PrivateCloud_ManagementCluster_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud_ManagementCluster) *krm.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCloud_ManagementCluster{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	// TODO: map type string message for field NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func PrivateCloud_ManagementCluster_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCloud_ManagementCluster) *pb.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud_ManagementCluster{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	// TODO: map type string message for field NodeTypeConfigs
	// MISSING: StretchedClusterConfig
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
func VMwareEnginePrivateCloudObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.VMwareEnginePrivateCloudObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VMwareEnginePrivateCloudObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.NetworkConfig = NetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	// MISSING: Hcx
	// (near miss): "Hcx" vs "HCX"
	// MISSING: Nsx
	// (near miss): "Nsx" vs "NSX"
	out.Vcenter = VcenterObservedState_FromProto(mapCtx, in.GetVcenter())
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
	return out
}
func VMwareEnginePrivateCloudObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VMwareEnginePrivateCloudObservedState) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.State = direct.Enum_ToProto[pb.PrivateCloud_State](mapCtx, in.State)
	out.NetworkConfig = NetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	// MISSING: Hcx
	// (near miss): "Hcx" vs "HCX"
	// MISSING: Nsx
	// (near miss): "Nsx" vs "NSX"
	out.Vcenter = VcenterObservedState_ToProto(mapCtx, in.Vcenter)
	// MISSING: Uid
	// (near miss): "Uid" vs "UID"
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
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func Vcenter_ToProto(mapCtx *direct.MapContext, in *krm.Vcenter) *pb.Vcenter {
	if in == nil {
		return nil
	}
	out := &pb.Vcenter{}
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: State
	// MISSING: Fqdn
	// (near miss): "Fqdn" vs "FQDN"
	return out
}
func VcenterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Vcenter) *krm.VcenterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VcenterObservedState{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Fqdn
	return out
}
func VcenterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VcenterObservedState) *pb.Vcenter {
	if in == nil {
		return nil
	}
	out := &pb.Vcenter{}
	// MISSING: InternalIP
	// MISSING: Version
	out.State = direct.Enum_ToProto[pb.Vcenter_State](mapCtx, in.State)
	// MISSING: Fqdn
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
