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
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: CapacityGB
	// MISSING: StorageBytes
	out.SourceInstance = direct.LazyPtr(in.GetSourceInstance())
	out.SourceFileShare = direct.LazyPtr(in.GetSourceFileShare())
	// MISSING: SourceInstanceTier
	// MISSING: DownloadBytes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: CapacityGB
	// MISSING: StorageBytes
	out.SourceInstance = direct.ValueOf(in.SourceInstance)
	out.SourceFileShare = direct.ValueOf(in.SourceFileShare)
	// MISSING: SourceInstanceTier
	// MISSING: DownloadBytes
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	out.CapacityGB = direct.LazyPtr(in.GetCapacityGb())
	out.StorageBytes = direct.LazyPtr(in.GetStorageBytes())
	// MISSING: SourceInstance
	// MISSING: SourceFileShare
	out.SourceInstanceTier = direct.Enum_FromProto(mapCtx, in.GetSourceInstanceTier())
	out.DownloadBytes = direct.LazyPtr(in.GetDownloadBytes())
	out.SatisfiesPzs = direct.BoolValue_FromProto(mapCtx, in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	// MISSING: KMSKeyName
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	out.CapacityGb = direct.ValueOf(in.CapacityGB)
	out.StorageBytes = direct.ValueOf(in.StorageBytes)
	// MISSING: SourceInstance
	// MISSING: SourceFileShare
	out.SourceInstanceTier = direct.Enum_ToProto[pb.Instance_Tier](mapCtx, in.SourceInstanceTier)
	out.DownloadBytes = direct.ValueOf(in.DownloadBytes)
	out.SatisfiesPzs = direct.BoolValue_ToProto(mapCtx, in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	// MISSING: KMSKeyName
	return out
}
