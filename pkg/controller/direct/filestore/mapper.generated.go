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
)

func FileShareConfig_FromProto(mapCtx *direct.MapContext, in *pb.FileShareConfig) *krm.FileShareConfig {
	if in == nil {
		return nil
	}
	out := &krm.FileShareConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CapacityGB = direct.LazyPtr(in.GetCapacityGb())
	// MISSING: SourceBackup
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
	// MISSING: SourceBackup
	out.NfsExportOptions = direct.Slice_ToProto(mapCtx, in.NfsExportOptions, NfsExportOptions_ToProto)
	return out
}
func FilestoreInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.FilestoreInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreInstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfigObservedState_FromProto)
	// MISSING: Etag
	out.SatisfiesPzs = direct.BoolValue_FromProto(mapCtx, in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	out.SuspensionReasons = direct.EnumSlice_FromProto(mapCtx, in.SuspensionReasons)
	return out
}
func FilestoreInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfigObservedState_ToProto)
	// MISSING: Etag
	out.SatisfiesPzs = direct.BoolValue_ToProto(mapCtx, in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	out.SuspensionReasons = direct.EnumSlice_ToProto[pb.Instance_SuspensionReason](mapCtx, in.SuspensionReasons)
	return out
}
func FilestoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.FilestoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreInstanceSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	// MISSING: Labels
	out.FileShares = direct.Slice_FromProto(mapCtx, in.FileShares, FileShareConfig_FromProto)
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_FromProto)
	// MISSING: Etag
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func FilestoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Description = direct.ValueOf(in.Description)
	out.Tier = direct.Enum_ToProto[pb.Instance_Tier](mapCtx, in.Tier)
	// MISSING: Labels
	out.FileShares = direct.Slice_ToProto(mapCtx, in.FileShares, FileShareConfig_ToProto)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_ToProto)
	// MISSING: Etag
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
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
	out.IPAddresses = in.IpAddresses
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
	out.IPRanges = in.IpRanges
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
