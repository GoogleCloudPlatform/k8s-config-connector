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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
)
func Hcx_FromProto(mapCtx *direct.MapContext, in *pb.Hcx) *krm.Hcx {
	if in == nil {
		return nil
	}
	out := &krm.Hcx{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: State
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
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
	out.Fqdn = direct.ValueOf(in.Fqdn)
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
	out.ManagementCidr = direct.LazyPtr(in.GetManagementCidr())
	out.VmwareEngineNetwork = direct.LazyPtr(in.GetVmwareEngineNetwork())
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: ManagementIPAddressLayoutVersion
	// MISSING: DnsServerIP
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.ManagementCidr = direct.ValueOf(in.ManagementCidr)
	out.VmwareEngineNetwork = direct.ValueOf(in.VmwareEngineNetwork)
	// MISSING: VmwareEngineNetworkCanonical
	// MISSING: ManagementIPAddressLayoutVersion
	// MISSING: DnsServerIP
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	// MISSING: ManagementCidr
	// MISSING: VmwareEngineNetwork
	out.VmwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	out.ManagementIPAddressLayoutVersion = direct.LazyPtr(in.GetManagementIpAddressLayoutVersion())
	out.DnsServerIP = direct.LazyPtr(in.GetDnsServerIp())
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: ManagementCidr
	// MISSING: VmwareEngineNetwork
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VmwareEngineNetworkCanonical)
	out.ManagementIpAddressLayoutVersion = direct.ValueOf(in.ManagementIPAddressLayoutVersion)
	out.DnsServerIp = direct.ValueOf(in.DnsServerIP)
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
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
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
	out.Fqdn = direct.ValueOf(in.Fqdn)
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
func PrivateCloud_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCloud{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.ManagementCluster = PrivateCloud_ManagementCluster_FromProto(mapCtx, in.GetManagementCluster())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func PrivateCloud_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCloud) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.ManagementCluster = PrivateCloud_ManagementCluster_ToProto(mapCtx, in.ManagementCluster)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	out.Type = direct.Enum_ToProto[pb.PrivateCloud_Type](mapCtx, in.Type)
	return out
}
func PrivateCloudObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.PrivateCloudObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCloudObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.NetworkConfig = NetworkConfigObservedState_FromProto(mapCtx, in.GetNetworkConfig())
	// MISSING: ManagementCluster
	// MISSING: Description
	out.Hcx = Hcx_FromProto(mapCtx, in.GetHcx())
	out.Nsx = Nsx_FromProto(mapCtx, in.GetNsx())
	out.Vcenter = Vcenter_FromProto(mapCtx, in.GetVcenter())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Type
	return out
}
func PrivateCloudObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCloudObservedState) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.State = direct.Enum_ToProto[pb.PrivateCloud_State](mapCtx, in.State)
	out.NetworkConfig = NetworkConfigObservedState_ToProto(mapCtx, in.NetworkConfig)
	// MISSING: ManagementCluster
	// MISSING: Description
	out.Hcx = Hcx_ToProto(mapCtx, in.Hcx)
	out.Nsx = Nsx_ToProto(mapCtx, in.Nsx)
	out.Vcenter = Vcenter_ToProto(mapCtx, in.Vcenter)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Type
	return out
}
func PrivateCloud_ManagementCluster_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud_ManagementCluster) *krm.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCloud_ManagementCluster{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	// MISSING: NodeTypeConfigs
	out.StretchedClusterConfig = StretchedClusterConfig_FromProto(mapCtx, in.GetStretchedClusterConfig())
	return out
}
func PrivateCloud_ManagementCluster_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCloud_ManagementCluster) *pb.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud_ManagementCluster{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	// MISSING: NodeTypeConfigs
	out.StretchedClusterConfig = StretchedClusterConfig_ToProto(mapCtx, in.StretchedClusterConfig)
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
func Vcenter_FromProto(mapCtx *direct.MapContext, in *pb.Vcenter) *krm.Vcenter {
	if in == nil {
		return nil
	}
	out := &krm.Vcenter{}
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: State
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
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
	out.Fqdn = direct.ValueOf(in.Fqdn)
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
func VmwareenginePrivateCloudObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.VmwareenginePrivateCloudObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareenginePrivateCloudObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: NetworkConfig
	// MISSING: ManagementCluster
	// MISSING: Description
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	// MISSING: Type
	return out
}
func VmwareenginePrivateCloudObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareenginePrivateCloudObservedState) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: NetworkConfig
	// MISSING: ManagementCluster
	// MISSING: Description
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	// MISSING: Type
	return out
}
func VmwareenginePrivateCloudSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud) *krm.VmwareenginePrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareenginePrivateCloudSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: NetworkConfig
	// MISSING: ManagementCluster
	// MISSING: Description
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	// MISSING: Type
	return out
}
func VmwareenginePrivateCloudSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareenginePrivateCloudSpec) *pb.PrivateCloud {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: State
	// MISSING: NetworkConfig
	// MISSING: ManagementCluster
	// MISSING: Description
	// MISSING: Hcx
	// MISSING: Nsx
	// MISSING: Vcenter
	// MISSING: Uid
	// MISSING: Type
	return out
}
