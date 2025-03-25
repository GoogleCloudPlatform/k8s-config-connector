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

// +generated:mapper
// proto.service: google.bigtable.admin.v2
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1alpha1

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BigtableBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupObservedState{}
	// MISSING: Name
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EncryptionInfo = EncryptionInfoObservedState_FromProto(mapCtx, in.GetEncryptionInfo())
	return out
}
func BigtableBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.EncryptionInfo = EncryptionInfoObservedState_ToProto(mapCtx, in.EncryptionInfo)
	return out
}
func BigtableBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BigtableBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableBackupSpec{}
	// MISSING: Name
	if in.GetSourceTable() != "" {
		out.SourceTableRef = &bigtablev1beta1.TableRef{External: in.GetSourceTable()}
	}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.GetHotToStandardTime())
	return out
}
func BigtableBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	if in.SourceTableRef != nil {
		out.SourceTable = in.SourceTableRef.External
	}
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.HotToStandardTime = direct.StringTimestamp_ToProto(mapCtx, in.HotToStandardTime)
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	out.KMSKeyVersion = direct.LazyPtr(in.GetKmsKeyVersion())
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	out.KmsKeyVersion = direct.ValueOf(in.KMSKeyVersion)
	return out
}
