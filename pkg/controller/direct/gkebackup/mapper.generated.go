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

package gkebackup

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func GkebackupVolumeBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VolumeBackup) *krm.GkebackupVolumeBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupVolumeBackupObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func GkebackupVolumeBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupVolumeBackupObservedState) *pb.VolumeBackup {
	if in == nil {
		return nil
	}
	out := &pb.VolumeBackup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func GkebackupVolumeBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.VolumeBackup) *krm.GkebackupVolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupVolumeBackupSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func GkebackupVolumeBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupVolumeBackupSpec) *pb.VolumeBackup {
	if in == nil {
		return nil
	}
	out := &pb.VolumeBackup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func NamespacedName_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedName) *krm.NamespacedName {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedName{}
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func NamespacedName_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedName) *pb.NamespacedName {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedName{}
	out.Namespace = direct.ValueOf(in.Namespace)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func VolumeBackup_FromProto(mapCtx *direct.MapContext, in *pb.VolumeBackup) *krm.VolumeBackup {
	if in == nil {
		return nil
	}
	out := &krm.VolumeBackup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func VolumeBackup_ToProto(mapCtx *direct.MapContext, in *krm.VolumeBackup) *pb.VolumeBackup {
	if in == nil {
		return nil
	}
	out := &pb.VolumeBackup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourcePvc
	// MISSING: VolumeBackupHandle
	// MISSING: Format
	// MISSING: StorageBytes
	// MISSING: DiskSizeBytes
	// MISSING: CompleteTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Etag
	return out
}
func VolumeBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VolumeBackup) *krm.VolumeBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VolumeBackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SourcePvc = NamespacedName_FromProto(mapCtx, in.GetSourcePvc())
	out.VolumeBackupHandle = direct.LazyPtr(in.GetVolumeBackupHandle())
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.StorageBytes = direct.LazyPtr(in.GetStorageBytes())
	out.DiskSizeBytes = direct.LazyPtr(in.GetDiskSizeBytes())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func VolumeBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VolumeBackupObservedState) *pb.VolumeBackup {
	if in == nil {
		return nil
	}
	out := &pb.VolumeBackup{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SourcePvc = NamespacedName_ToProto(mapCtx, in.SourcePvc)
	out.VolumeBackupHandle = direct.ValueOf(in.VolumeBackupHandle)
	out.Format = direct.Enum_ToProto[pb.VolumeBackup_VolumeBackupFormat](mapCtx, in.Format)
	out.StorageBytes = direct.ValueOf(in.StorageBytes)
	out.DiskSizeBytes = direct.ValueOf(in.DiskSizeBytes)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.State = direct.Enum_ToProto[pb.VolumeBackup_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
