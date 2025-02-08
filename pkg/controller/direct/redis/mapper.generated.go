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

package redis

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/redis/cluster/apiv1beta1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
func BackupFile_FromProto(mapCtx *direct.MapContext, in *pb.BackupFile) *krm.BackupFile {
	if in == nil {
		return nil
	}
	out := &krm.BackupFile{}
	// MISSING: FileName
	// MISSING: SizeBytes
	// MISSING: CreateTime
	return out
}
func BackupFile_ToProto(mapCtx *direct.MapContext, in *krm.BackupFile) *pb.BackupFile {
	if in == nil {
		return nil
	}
	out := &pb.BackupFile{}
	// MISSING: FileName
	// MISSING: SizeBytes
	// MISSING: CreateTime
	return out
}
func BackupFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupFile) *krm.BackupFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupFileObservedState{}
	out.FileName = direct.LazyPtr(in.GetFileName())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func BackupFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupFileObservedState) *pb.BackupFile {
	if in == nil {
		return nil
	}
	out := &pb.BackupFile{}
	out.FileName = direct.ValueOf(in.FileName)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.ClusterUid = direct.LazyPtr(in.GetClusterUid())
	out.TotalSizeBytes = direct.LazyPtr(in.GetTotalSizeBytes())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.BackupFiles = direct.Slice_FromProto(mapCtx, in.BackupFiles, BackupFile_FromProto)
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.ClusterUid = direct.ValueOf(in.ClusterUid)
	out.TotalSizeBytes = direct.ValueOf(in.TotalSizeBytes)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.BackupFiles = direct.Slice_ToProto(mapCtx, in.BackupFiles, BackupFile_ToProto)
	out.NodeType = direct.Enum_ToProto[pb.NodeType](mapCtx, in.NodeType)
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	// MISSING: KMSKeyPrimaryState
	// MISSING: LastUpdateTime
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	// MISSING: KMSKeyPrimaryState
	// MISSING: LastUpdateTime
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.KMSKeyVersions = in.KmsKeyVersions
	out.KMSKeyPrimaryState = direct.Enum_FromProto(mapCtx, in.GetKmsKeyPrimaryState())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionType)
	out.KmsKeyVersions = in.KMSKeyVersions
	out.KmsKeyPrimaryState = direct.Enum_ToProto[pb.EncryptionInfo_KmsKeyState](mapCtx, in.KMSKeyPrimaryState)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	return out
}
func RedisBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.RedisBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisBackupObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
func RedisBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RedisBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
func RedisBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.RedisBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisBackupSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
func RedisBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Cluster
	// MISSING: ClusterUid
	// MISSING: TotalSizeBytes
	// MISSING: ExpireTime
	// MISSING: EngineVersion
	// MISSING: BackupFiles
	// MISSING: NodeType
	// MISSING: ReplicaCount
	// MISSING: ShardCount
	// MISSING: BackupType
	// MISSING: State
	// MISSING: EncryptionInfo
	// MISSING: Uid
	return out
}
