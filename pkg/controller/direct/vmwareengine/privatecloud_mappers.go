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
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.ManagementCIDR = direct.LazyPtr(in.GetManagementCidr())
	if in.GetVmwareEngineNetwork() != "" {
		out.VMwareEngineNetworkRef = &krm.VmwareEngineNetworkRef{
			External: in.GetVmwareEngineNetwork(),
		}
	}
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.ManagementCidr = direct.ValueOf(in.ManagementCIDR)
	if in.VMwareEngineNetworkRef != nil {
		out.VmwareEngineNetwork = in.VMwareEngineNetworkRef.External
	}
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	out.VMwareEngineNetworkCanonical = direct.LazyPtr(in.GetVmwareEngineNetworkCanonical())
	out.ManagementIPAddressLayoutVersion = direct.LazyPtr(in.GetManagementIpAddressLayoutVersion())
	out.DNSServerIP = direct.LazyPtr(in.GetDnsServerIp())
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.VmwareEngineNetworkCanonical = direct.ValueOf(in.VMwareEngineNetworkCanonical)
	out.ManagementIpAddressLayoutVersion = direct.ValueOf(in.ManagementIPAddressLayoutVersion)
	out.DnsServerIp = direct.ValueOf(in.DNSServerIP)
	return out
}
func PrivateCloud_ManagementCluster_FromProto(mapCtx *direct.MapContext, in *pb.PrivateCloud_ManagementCluster) *krm.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &krm.PrivateCloud_ManagementCluster{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	// Note: in KRM we use a []NodeTypeConfig with a virtual field "NodeTypeID" as map key to the proto message
	// which is a map[string]NodeTypeConfig
	for k, v := range in.GetNodeTypeConfigs() {
		nodeTypeConfig := NodeTypeConfig_FromProto(mapCtx, v)
		nodeTypeConfig.NodeTypeID = direct.LazyPtr(k)
		out.NodeTypeConfigs = append(out.NodeTypeConfigs, nodeTypeConfig)
	}
	// MISSING: StretchedClusterConfig
	return out
}
func PrivateCloud_ManagementCluster_ToProto(mapCtx *direct.MapContext, in *krm.PrivateCloud_ManagementCluster) *pb.PrivateCloud_ManagementCluster {
	if in == nil {
		return nil
	}
	out := &pb.PrivateCloud_ManagementCluster{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.NodeTypeConfigs = make(map[string]*pb.NodeTypeConfig)
	for _, v := range in.NodeTypeConfigs {
		nodeTypeID := direct.ValueOf(v.NodeTypeID)
		out.NodeTypeConfigs[nodeTypeID] = NodeTypeConfig_ToProto(mapCtx, v)
	}
	// MISSING: StretchedClusterConfig
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
	out.HCX = Hcx_FromProto(mapCtx, in.GetHcx())
	out.NSX = Nsx_FromProto(mapCtx, in.GetNsx())
	out.Vcenter = Vcenter_FromProto(mapCtx, in.GetVcenter())
	out.UID = direct.LazyPtr(in.GetUid())
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
	out.Hcx = Hcx_ToProto(mapCtx, in.HCX)
	out.Nsx = Nsx_ToProto(mapCtx, in.NSX)
	out.Vcenter = Vcenter_ToProto(mapCtx, in.Vcenter)
	out.Uid = direct.ValueOf(in.UID)
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
