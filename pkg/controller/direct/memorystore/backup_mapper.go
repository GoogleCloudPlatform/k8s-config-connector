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

// +generated:mapper
// krm.group: memorystore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.memorystore.v1

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
func MemorystoreInstanceBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MemorystoreInstanceBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceBackupObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.InstanceUid = direct.LazyPtr(in.GetInstanceUid())
	out.TotalSizeBytes = direct.LazyPtr(in.GetTotalSizeBytes())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.BackupFiles = direct.Slice_FromProto(mapCtx, in.BackupFiles, BackupFileObservedState_FromProto)
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.ReplicaCount = direct.LazyPtr(in.GetReplicaCount())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func MemorystoreInstanceBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Instance = direct.ValueOf(in.Instance)
	out.InstanceUid = direct.ValueOf(in.InstanceUid)
	out.TotalSizeBytes = direct.ValueOf(in.TotalSizeBytes)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.BackupFiles = direct.Slice_ToProto(mapCtx, in.BackupFiles, BackupFileObservedState_ToProto)
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.ReplicaCount = direct.ValueOf(in.ReplicaCount)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func MemorystoreInstanceBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MemorystoreInstanceBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceBackupSpec{}
	backupCollection, backupName := parseBackupExternal(mapCtx, in.Name)
	out.BackupCollection = direct.LazyPtr(backupCollection)
	out.ResourceID = direct.LazyPtr(backupName)
	if in.Instance != "" {
		out.InstanceRef = &refsv1beta1.MemorystoreInstanceRef{
			External: in.Instance,
		}
	}
	return out
}
func MemorystoreInstanceBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	if in.BackupCollection != nil && in.ResourceID != nil {
		out.Name = direct.ValueOf(in.BackupCollection) + "/" + direct.ValueOf(in.ResourceID)
	}
	if in.InstanceRef != nil {
		out.Instance = in.InstanceRef.External
	}
	return out
}
func MemorystoreInstanceBackupSpec_ToProtoBackupInstanceRequest(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceBackupSpec) *pb.BackupInstanceRequest {
	if in == nil {
		return nil
	}
	out := &pb.BackupInstanceRequest{}
	out.BackupId = in.ResourceID
	if in.InstanceRef != nil {
		out.Name = in.InstanceRef.External
	}
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}
func parseBackupExternal(mapCtx *direct.MapContext, in string) (string, string) {
	if in == "" {
		return "", ""
	}
	backupCollection, backupName, err := krm.ParseBackupExternal(in)
	if err != nil {
		mapCtx.Errorf("invalid backup name: %s", in)
		return "", ""
	}
	return backupCollection, backupName
}
