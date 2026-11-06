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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krmbackupdrv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krmbackupdrv1alpha1.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krmbackupdrv1alpha1.AccessConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Name = in.Name
	out.ExternalIP = in.ExternalIp
	out.ExternalIPV6 = in.ExternalIpv6
	out.ExternalIPV6PrefixLength = in.ExternalIpv6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.NetworkTier = direct.Enum_FromProto(mapCtx, in.GetNetworkTier())
	return out
}

func AccessConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.Type = direct.EnumPtr_ToProto[pb.AccessConfig_AccessType](mapCtx, in.Type)
	out.Name = in.Name
	out.ExternalIp = in.ExternalIP
	out.ExternalIpv6 = in.ExternalIPV6
	out.ExternalIpv6PrefixLength = in.ExternalIPV6PrefixLength
	out.SetPublicPtr = in.SetPublicPtr
	out.PublicPtrDomainName = in.PublicPtrDomainName
	out.NetworkTier = direct.EnumPtr_ToProto[pb.AccessConfig_NetworkTier](mapCtx, in.NetworkTier)
	return out
}

func NetworkInterfaceObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krmbackupdrv1alpha1.NetworkInterfaceObservedState {
	if in == nil {
		return nil
	}
	out := &krmbackupdrv1alpha1.NetworkInterfaceObservedState{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IPAddress = in.IpAddress
	out.IPV6Address = in.Ipv6Address
	out.InternalIPV6PrefixLength = in.InternalIpv6PrefixLength
	out.Name = in.Name
	out.AccessConfigs = direct.Slice_FromProto(mapCtx, in.AccessConfigs, AccessConfig_v1alpha1_FromProto)
	out.IPV6AccessConfigs = direct.Slice_FromProto(mapCtx, in.Ipv6AccessConfigs, AccessConfig_v1alpha1_FromProto)
	out.AliasIPRanges = direct.Slice_FromProto(mapCtx, in.AliasIpRanges, AliasIPRange_v1alpha1_FromProto)
	out.StackType = direct.Enum_FromProto(mapCtx, in.GetStackType())
	out.IPV6AccessType = direct.Enum_FromProto(mapCtx, in.GetIpv6AccessType())
	out.QueueCount = in.QueueCount
	out.NicType = direct.Enum_FromProto(mapCtx, in.GetNicType())
	out.NetworkAttachment = in.NetworkAttachment
	return out
}

func NetworkInterfaceObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.NetworkInterfaceObservedState) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	out.Network = in.Network
	out.Subnetwork = in.Subnetwork
	out.IpAddress = in.IPAddress
	out.Ipv6Address = in.IPV6Address
	out.InternalIpv6PrefixLength = in.InternalIPV6PrefixLength
	out.Name = in.Name
	out.AccessConfigs = direct.Slice_ToProto(mapCtx, in.AccessConfigs, AccessConfig_v1alpha1_ToProto)
	out.Ipv6AccessConfigs = direct.Slice_ToProto(mapCtx, in.IPV6AccessConfigs, AccessConfig_v1alpha1_ToProto)
	out.AliasIpRanges = direct.Slice_ToProto(mapCtx, in.AliasIPRanges, AliasIPRange_v1alpha1_ToProto)
	out.StackType = direct.EnumPtr_ToProto[pb.NetworkInterface_StackType](mapCtx, in.StackType)
	out.Ipv6AccessType = direct.EnumPtr_ToProto[pb.NetworkInterface_Ipv6AccessType](mapCtx, in.IPV6AccessType)
	out.QueueCount = in.QueueCount
	out.NicType = direct.EnumPtr_ToProto[pb.NetworkInterface_NicType](mapCtx, in.NicType)
	out.NetworkAttachment = in.NetworkAttachment
	return out
}

func Metadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krmbackupdrv1alpha1.Metadata {
	if in == nil {
		return nil
	}
	out := &krmbackupdrv1alpha1.Metadata{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Entry_v1alpha1_FromProto)
	return out
}

func Metadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbackupdrv1alpha1.Metadata) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Entry_v1alpha1_ToProto)
	return out
}
