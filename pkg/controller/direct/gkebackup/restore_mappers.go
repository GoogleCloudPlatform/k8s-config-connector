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
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEBackupRestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krmv1alpha1.GKEBackupRestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GKEBackupRestoreObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.RestoreConfig = RestoreConfig_FromProto(mapCtx, in.GetRestoreConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.ResourcesRestoredCount = direct.LazyPtr(in.GetResourcesRestoredCount())
	out.ResourcesExcludedCount = direct.LazyPtr(in.GetResourcesExcludedCount())
	out.ResourcesFailedCount = direct.LazyPtr(in.GetResourcesFailedCount())
	out.VolumesRestoredCount = direct.LazyPtr(in.GetVolumesRestoredCount())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func GKEBackupRestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GKEBackupRestoreObservedState) *pb.Restore {
	if in == nil {
		return nil
	}
	out := &pb.Restore{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.RestoreConfig = RestoreConfig_ToProto(mapCtx, in.RestoreConfig)
	out.State = direct.Enum_ToProto[pb.Restore_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.ResourcesRestoredCount = direct.ValueOf(in.ResourcesRestoredCount)
	out.ResourcesExcludedCount = direct.ValueOf(in.ResourcesExcludedCount)
	out.ResourcesFailedCount = direct.ValueOf(in.ResourcesFailedCount)
	out.VolumesRestoredCount = direct.ValueOf(in.VolumesRestoredCount)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func GKEBackupRestoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.Restore) *krmv1alpha1.GKEBackupRestoreSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GKEBackupRestoreSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetBackup() != "" {
		out.BackupRef = &krmv1alpha1.BackupRef{External: in.GetBackup()}
	}
	out.Labels = in.Labels
	out.Filter = Restore_Filter_FromProto(mapCtx, in.GetFilter())
	out.VolumeDataRestorePolicyOverrides = direct.Slice_FromProto(mapCtx, in.VolumeDataRestorePolicyOverrides, VolumeDataRestorePolicyOverride_FromProto)
	return out
}
