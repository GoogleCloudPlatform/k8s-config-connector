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

package filestore

import (
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func FileShareConfig_FromProto(mapCtx *direct.MapContext, in *pb.FileShareConfig) *krm.FileShareConfig {
	if in == nil {
		return nil
	}
	out := &krm.FileShareConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CapacityGB = direct.LazyPtr(in.GetCapacityGb())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	out.NfsExportOptions = direct.Slice_FromProto(mapCtx, in.NfsExportOptions, NfsExportOptions_FromProto)
	return out
}
func FileShareConfig_ToProto(mapCtx *direct.MapContext, in *krm.FileShareConfig) *pb.FileShareConfig {
	if in == nil {
		return nil
	}
	out := &pb.FileShareConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.CapacityGb = direct.ValueOf(in.CapacityGB)
	if oneof := FileShareConfig_SourceBackup_ToProto(mapCtx, in.SourceBackup); oneof != nil {
		out.Source = oneof
	}
	out.NfsExportOptions = direct.Slice_ToProto(mapCtx, in.NfsExportOptions, NfsExportOptions_ToProto)
	return out
}
func FilestoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.FilestoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreInstanceObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: KMSKeyName
	// MISSING: SuspensionReasons
	return out
}
func FilestoreInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: KMSKeyName
	// MISSING: SuspensionReasons
	return out
}
func FilestoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.FilestoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreInstanceSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: KMSKeyName
	// MISSING: SuspensionReasons
	return out
}
func FilestoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	// MISSING: KMSKeyName
	// MISSING: SuspensionReasons
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.Labels = in.Labels
	out.FileShares = direct.Slice_FromProto(mapCtx, in.FileShares, FileShareConfig_FromProto)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: SuspensionReasons
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: CreateTime
	out.Tier = direct.Enum_ToProto[pb.Instance_Tier](mapCtx, in.Tier)
	out.Labels = in.Labels
	out.FileShares = direct.Slice_ToProto(mapCtx, in.FileShares, FileShareConfig_ToProto)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: SuspensionReasons
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfigObservedState_FromProto)
	// MISSING: Etag
	out.SatisfiesPzs = direct.BoolValue_FromProto(mapCtx, in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	// MISSING: KMSKeyName
	out.SuspensionReasons = direct.EnumSlice_FromProto(mapCtx, in.SuspensionReasons)
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Tier
	// MISSING: Labels
	// MISSING: FileShares
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfigObservedState_ToProto)
	// MISSING: Etag
	out.SatisfiesPzs = direct.BoolValue_ToProto(mapCtx, in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	// MISSING: KMSKeyName
	out.SuspensionReasons = direct.EnumSlice_ToProto[pb.Instance_SuspensionReason](mapCtx, in.SuspensionReasons)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Modes = direct.EnumSlice_FromProto(mapCtx, in.Modes)
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	// MISSING: IPAddresses
	out.ConnectMode = direct.Enum_FromProto(mapCtx, in.GetConnectMode())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.Modes = direct.EnumSlice_ToProto[pb.NetworkConfig_AddressMode](mapCtx, in.Modes)
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	// MISSING: IPAddresses
	out.ConnectMode = direct.Enum_ToProto[pb.NetworkConfig_ConnectMode](mapCtx, in.ConnectMode)
	return out
}
func NetworkConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfigObservedState{}
	// MISSING: Network
	// MISSING: Modes
	// MISSING: ReservedIPRange
	out.IPAddresses = in.IPAddresses
	// MISSING: ConnectMode
	return out
}
func NetworkConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfigObservedState) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	// MISSING: Network
	// MISSING: Modes
	// MISSING: ReservedIPRange
	out.IpAddresses = in.IPAddresses
	// MISSING: ConnectMode
	return out
}
func NfsExportOptions_FromProto(mapCtx *direct.MapContext, in *pb.NfsExportOptions) *krm.NfsExportOptions {
	if in == nil {
		return nil
	}
	out := &krm.NfsExportOptions{}
	out.IPRanges = in.IPRanges
	out.AccessMode = direct.Enum_FromProto(mapCtx, in.GetAccessMode())
	out.SquashMode = direct.Enum_FromProto(mapCtx, in.GetSquashMode())
	out.AnonUid = direct.LazyPtr(in.GetAnonUid())
	out.AnonGid = direct.LazyPtr(in.GetAnonGid())
	return out
}
func NfsExportOptions_ToProto(mapCtx *direct.MapContext, in *krm.NfsExportOptions) *pb.NfsExportOptions {
	if in == nil {
		return nil
	}
	out := &pb.NfsExportOptions{}
	out.IpRanges = in.IPRanges
	out.AccessMode = direct.Enum_ToProto[pb.NfsExportOptions_AccessMode](mapCtx, in.AccessMode)
	out.SquashMode = direct.Enum_ToProto[pb.NfsExportOptions_SquashMode](mapCtx, in.SquashMode)
	out.AnonUid = direct.ValueOf(in.AnonUid)
	out.AnonGid = direct.ValueOf(in.AnonGid)
	return out
}
