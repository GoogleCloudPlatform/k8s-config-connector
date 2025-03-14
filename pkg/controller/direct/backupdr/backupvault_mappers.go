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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRBackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupDRBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupVaultObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Deletable = in.Deletable
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.BackupCount = direct.LazyPtr(in.GetBackupCount())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.TotalStoredBytes = direct.LazyPtr(in.GetTotalStoredBytes())
	out.UID = direct.LazyPtr(in.GetUid())
	return out
}
func BackupDRBackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Deletable = in.Deletable
	out.State = direct.Enum_ToProto[pb.BackupVault_State](mapCtx, in.State)
	out.BackupCount = direct.ValueOf(in.BackupCount)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.TotalStoredBytes = direct.ValueOf(in.TotalStoredBytes)
	out.Uid = direct.ValueOf(in.UID)
	return out
}
func BackupDRBackupVaultSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	out.Description = in.Description
	out.Labels = in.Labels
	out.BackupMinimumEnforcedRetentionDuration = direct.StringDuration_ToProto(mapCtx, in.BackupMinimumEnforcedRetentionDuration)
	out.Etag = in.Etag
	out.EffectiveTime = direct.StringTimestamp_ToProto(mapCtx, in.EffectiveTime)
	// MISSING: Uid
	out.Annotations = in.Annotations
	out.AccessRestriction = direct.Enum_ToProto[pb.BackupVault_AccessRestriction](mapCtx, in.AccessRestriction)
	return out
}
