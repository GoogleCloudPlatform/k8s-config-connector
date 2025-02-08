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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/filestore/apiv1beta1/filestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func FilestoreShareObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Share) *krm.FilestoreShareObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreShareObservedState{}
	// MISSING: Name
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Backup
	return out
}
func FilestoreShareObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreShareObservedState) *pb.Share {
	if in == nil {
		return nil
	}
	out := &pb.Share{}
	// MISSING: Name
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Backup
	return out
}
func FilestoreShareSpec_FromProto(mapCtx *direct.MapContext, in *pb.Share) *krm.FilestoreShareSpec {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreShareSpec{}
	// MISSING: Name
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Backup
	return out
}
func FilestoreShareSpec_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreShareSpec) *pb.Share {
	if in == nil {
		return nil
	}
	out := &pb.Share{}
	// MISSING: Name
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Backup
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
	out.SecurityFlavors = direct.EnumSlice_FromProto(mapCtx, in.SecurityFlavors)
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
	out.SecurityFlavors = direct.EnumSlice_ToProto[pb.NfsExportOptions_SecurityFlavor](mapCtx, in.SecurityFlavors)
	return out
}
func Share_FromProto(mapCtx *direct.MapContext, in *pb.Share) *krm.Share {
	if in == nil {
		return nil
	}
	out := &krm.Share{}
	// MISSING: Name
	out.MountName = direct.LazyPtr(in.GetMountName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CapacityGB = direct.LazyPtr(in.GetCapacityGb())
	out.NfsExportOptions = direct.Slice_FromProto(mapCtx, in.NfsExportOptions, NfsExportOptions_FromProto)
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	out.Backup = direct.LazyPtr(in.GetBackup())
	return out
}
func Share_ToProto(mapCtx *direct.MapContext, in *krm.Share) *pb.Share {
	if in == nil {
		return nil
	}
	out := &pb.Share{}
	// MISSING: Name
	out.MountName = direct.ValueOf(in.MountName)
	out.Description = direct.ValueOf(in.Description)
	out.CapacityGb = direct.ValueOf(in.CapacityGB)
	out.NfsExportOptions = direct.Slice_ToProto(mapCtx, in.NfsExportOptions, NfsExportOptions_ToProto)
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	if oneof := Share_Backup_ToProto(mapCtx, in.Backup); oneof != nil {
		out.Source = oneof
	}
	return out
}
func ShareObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Share) *krm.ShareObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ShareObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	// MISSING: Backup
	return out
}
func ShareObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ShareObservedState) *pb.Share {
	if in == nil {
		return nil
	}
	out := &pb.Share{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: MountName
	// MISSING: Description
	// MISSING: CapacityGB
	// MISSING: NfsExportOptions
	out.State = direct.Enum_ToProto[pb.Share_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	// MISSING: Backup
	return out
}
