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

package filestore

import (
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func InstanceFileShares_FromProto(mapCtx *direct.MapContext, in *pb.FileShareConfig) *krm.InstanceFileShares {
	if in == nil {
		return nil
	}
	out := &krm.InstanceFileShares{}
	out.Name = direct.LazyPtr(in.GetName())
	if in.GetCapacityGb() != 0 {
		out.CapacityGb = direct.LazyPtr(in.GetCapacityGb())
	}
	if in.GetSourceBackup() != "" {
		out.SourceBackupRef = &krm.FilestoreBackupRef{External: in.GetSourceBackup()}
	}
	out.NfsExportOptions = direct.Slice_FromProto(mapCtx, in.NfsExportOptions, InstanceNfsExportOptions_FromProto)
	return out
}

func InstanceFileShares_ToProto(mapCtx *direct.MapContext, in *krm.InstanceFileShares) *pb.FileShareConfig {
	if in == nil {
		return nil
	}
	out := &pb.FileShareConfig{}
	out.Name = direct.ValueOf(in.Name)
	if in.CapacityGb != nil {
		out.CapacityGb = *in.CapacityGb
	}
	if in.SourceBackupRef != nil {
		out.Source = &pb.FileShareConfig_SourceBackup{SourceBackup: in.SourceBackupRef.External}
	}
	out.NfsExportOptions = direct.Slice_ToProto(mapCtx, in.NfsExportOptions, InstanceNfsExportOptions_ToProto)
	return out
}

func InstanceNetworks_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.InstanceNetworks {
	if in == nil {
		return nil
	}
	out := &krm.InstanceNetworks{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computerefs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Modes = direct.EnumSlice_FromProto(mapCtx, in.Modes)
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	out.IpAddresses = in.IpAddresses
	return out
}

func InstanceNetworks_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNetworks) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.Modes = direct.EnumSlice_ToProto[pb.NetworkConfig_AddressMode](mapCtx, in.Modes)
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	out.IpAddresses = in.IpAddresses
	return out
}

func InstanceNfsExportOptions_FromProto(mapCtx *direct.MapContext, in *pb.NfsExportOptions) *krm.InstanceNfsExportOptions {
	if in == nil {
		return nil
	}
	out := &krm.InstanceNfsExportOptions{}
	out.IpRanges = in.IpRanges
	out.AccessMode = direct.Enum_FromProto(mapCtx, in.GetAccessMode())
	out.SquashMode = direct.Enum_FromProto(mapCtx, in.GetSquashMode())
	out.AnonUid = direct.LazyPtr(in.GetAnonUid())
	out.AnonGid = direct.LazyPtr(in.GetAnonGid())
	return out
}

func InstanceNfsExportOptions_ToProto(mapCtx *direct.MapContext, in *krm.InstanceNfsExportOptions) *pb.NfsExportOptions {
	if in == nil {
		return nil
	}
	out := &pb.NfsExportOptions{}
	out.IpRanges = in.IpRanges
	out.AccessMode = direct.Enum_ToProto[pb.NfsExportOptions_AccessMode](mapCtx, in.AccessMode)
	out.SquashMode = direct.Enum_ToProto[pb.NfsExportOptions_SquashMode](mapCtx, in.SquashMode)
	out.AnonUid = direct.ValueOf(in.AnonUid)
	out.AnonGid = direct.ValueOf(in.AnonGid)
	return out
}

func FilestoreInstanceStatus_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.FilestoreInstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreInstanceStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	return out
}

func FilestoreInstanceStatus_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreInstanceStatus) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	return out
}
