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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfig {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfig{}
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

func AccessConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfig) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	return out
}

func AttachedDiskObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDiskObservedState) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	return out
}

func ComputeInstanceBackupPropertiesObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeInstanceBackupPropertiesObservedState) *pb.ComputeInstanceBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.ComputeInstanceBackupProperties{}
	return out
}

func BackupDRBackupObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	return out
}

func Metadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krm.Metadata {
	if in == nil {
		return nil
	}
	out := &krm.Metadata{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Entry_v1alpha1_FromProto)
	return out
}

func Metadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Metadata) *pb.Metadata {
	if in == nil {
		return nil
	}
	out := &pb.Metadata{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Entry_v1alpha1_ToProto)
	return out
}

func Entry_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Entry) *krm.Entry {
	if in == nil {
		return nil
	}
	out := &krm.Entry{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}

func Entry_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Entry) *pb.Entry {
	if in == nil {
		return nil
	}
	out := &pb.Entry{}
	out.Key = in.Key
	out.Value = in.Value
	return out
}

func DiskBackupProperties_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DiskBackupProperties) *pb.DiskBackupProperties {
	if in == nil {
		return nil
	}
	out := &pb.DiskBackupProperties{}
	return out
}

func GuestOSFeature_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.GuestOSFeature) *pb.GuestOsFeature {
	if in == nil {
		return nil
	}
	out := &pb.GuestOsFeature{}
	return out
}

func Scheduling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling) *pb.Scheduling {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling{}
	return out
}

func Scheduling_NodeAffinity_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Scheduling_NodeAffinity) *pb.Scheduling_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.Scheduling_NodeAffinity{}
	return out
}

func NetworkInterfaceObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInterface) *krm.NetworkInterfaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInterfaceObservedState{}
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

func NetworkInterfaceObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInterfaceObservedState) *pb.NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInterface{}
	return out
}
