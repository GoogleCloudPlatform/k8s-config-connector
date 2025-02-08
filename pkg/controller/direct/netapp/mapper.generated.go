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

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	out.SourceVolume = direct.LazyPtr(in.GetSourceVolume())
	out.SourceSnapshot = in.SourceSnapshot
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: ChainStorageBytes
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	out.Description = direct.ValueOf(in.Description)
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	out.SourceVolume = direct.ValueOf(in.SourceVolume)
	out.SourceSnapshot = in.SourceSnapshot
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: ChainStorageBytes
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Description
	out.VolumeUsageBytes = direct.LazyPtr(in.GetVolumeUsageBytes())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	out.ChainStorageBytes = direct.LazyPtr(in.GetChainStorageBytes())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	// MISSING: Description
	out.VolumeUsageBytes = direct.ValueOf(in.VolumeUsageBytes)
	out.BackupType = direct.Enum_ToProto[pb.Backup_Type](mapCtx, in.BackupType)
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	out.ChainStorageBytes = direct.ValueOf(in.ChainStorageBytes)
	return out
}
func NetappBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.NetappBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappBackupObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Description
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: ChainStorageBytes
	return out
}
func NetappBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Description
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: ChainStorageBytes
	return out
}
func NetappBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.NetappBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappBackupSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Description
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: ChainStorageBytes
	return out
}
func NetappBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Description
	// MISSING: VolumeUsageBytes
	// MISSING: BackupType
	// MISSING: SourceVolume
	// MISSING: SourceSnapshot
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: ChainStorageBytes
	return out
}
