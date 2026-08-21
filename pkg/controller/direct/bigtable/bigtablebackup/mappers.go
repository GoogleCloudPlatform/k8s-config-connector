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

package bigtablebackup

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	adminpb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func BigtableBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupSpec) *adminpb.Backup {
	if in == nil {
		return nil
	}
	out := &adminpb.Backup{}
	if in.ExpireTime != nil {
		out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	}
	if in.BackupType != nil {
		out.BackupType = direct.Enum_ToProto[adminpb.Backup_BackupType](mapCtx, in.BackupType)
	}
	if in.HotToStandardTime != nil {
		out.HotToStandardTime = direct.StringTimestamp_ToProto(mapCtx, in.HotToStandardTime)
	}
	return out
}

func BigtableBackupSpec_FromProto(mapCtx *direct.MapContext, in *adminpb.Backup) *krm.BigtableBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupSpec{}
	if in.ExpireTime != nil {
		out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.ExpireTime)
	}
	if in.BackupType != adminpb.Backup_BACKUP_TYPE_UNSPECIFIED {
		out.BackupType = direct.Enum_FromProto(mapCtx, in.BackupType)
	}
	if in.HotToStandardTime != nil {
		out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.HotToStandardTime)
	}
	return out
}

func BigtableBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupObservedState) *adminpb.Backup {
	if in == nil {
		return nil
	}
	out := &adminpb.Backup{}
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	if in.StartTime != nil {
		out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	}
	if in.EndTime != nil {
		out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	}
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	if in.State != nil {
		out.State = direct.Enum_ToProto[adminpb.Backup_State](mapCtx, in.State)
	}
	if in.EncryptionInfo != nil {
		out.EncryptionInfo = &adminpb.EncryptionInfo{
			KmsKeyVersion: direct.ValueOf(in.EncryptionInfo.KMSKeyVersion),
		}
		if in.EncryptionInfo.EncryptionType != nil {
			out.EncryptionInfo.EncryptionType = direct.Enum_ToProto[adminpb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionInfo.EncryptionType)
		}
	}
	return out
}

func BigtableBackupObservedState_FromProto(mapCtx *direct.MapContext, in *adminpb.Backup) *krm.BigtableBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupObservedState{}
	out.SourceBackup = direct.LazyPtr(in.SourceBackup)
	if in.StartTime != nil {
		out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.StartTime)
	}
	if in.EndTime != nil {
		out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.EndTime)
	}
	out.SizeBytes = direct.LazyPtr(in.SizeBytes)
	if in.State != adminpb.Backup_STATE_UNSPECIFIED {
		out.State = direct.Enum_FromProto(mapCtx, in.State)
	}
	if in.EncryptionInfo != nil {
		out.EncryptionInfo = &krm.EncryptionInfoObservedState{
			KMSKeyVersion: direct.LazyPtr(in.EncryptionInfo.KmsKeyVersion),
		}
		if in.EncryptionInfo.EncryptionType != adminpb.EncryptionInfo_ENCRYPTION_TYPE_UNSPECIFIED {
			out.EncryptionInfo.EncryptionType = direct.Enum_FromProto(mapCtx, in.EncryptionInfo.EncryptionType)
		}
	}
	return out
}
