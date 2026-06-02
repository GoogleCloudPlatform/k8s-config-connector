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

package alloydbbackup

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AlloyDBBackupStatus_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.AlloyDBBackupStatus {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBBackupStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())

	if in.GetEncryptionInfo() != nil {
		out.EncryptionInfo = []krm.BackupEncryptionInfoStatus{
			{
				EncryptionType: direct.Enum_FromProto(mapCtx, in.GetEncryptionInfo().GetEncryptionType()),
				KmsKeyVersions: in.GetEncryptionInfo().GetKmsKeyVersions(),
			},
		}
	}
	return out
}

func AlloyDBBackupStatus_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBBackupStatus) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.Etag = direct.ValueOf(in.Etag)
	out.Reconciling = direct.ValueOf(in.Reconciling)

	if len(in.EncryptionInfo) > 0 {
		out.EncryptionInfo = &pb.EncryptionInfo{
			EncryptionType: direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionInfo[0].EncryptionType),
			KmsKeyVersions: in.EncryptionInfo[0].KmsKeyVersions,
		}
	}
	return out
}

func BackupEncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.BackupEncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackupEncryptionConfig{}
	out.KmsKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}

func BackupEncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.BackupEncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KmsKeyName)
	return out
}
