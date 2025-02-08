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

package firestore

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
)
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.DatabaseUid = direct.LazyPtr(in.GetDatabaseUid())
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Stats = Backup_Stats_FromProto(mapCtx, in.GetStats())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.Database = direct.ValueOf(in.Database)
	out.DatabaseUid = direct.ValueOf(in.DatabaseUid)
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Stats = Backup_Stats_ToProto(mapCtx, in.Stats)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	return out
}
func Backup_Stats_FromProto(mapCtx *direct.MapContext, in *pb.Backup_Stats) *krm.Backup_Stats {
	if in == nil {
		return nil
	}
	out := &krm.Backup_Stats{}
	// MISSING: SizeBytes
	// MISSING: DocumentCount
	// MISSING: IndexCount
	return out
}
func Backup_Stats_ToProto(mapCtx *direct.MapContext, in *krm.Backup_Stats) *pb.Backup_Stats {
	if in == nil {
		return nil
	}
	out := &pb.Backup_Stats{}
	// MISSING: SizeBytes
	// MISSING: DocumentCount
	// MISSING: IndexCount
	return out
}
func Backup_StatsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup_Stats) *krm.Backup_StatsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Backup_StatsObservedState{}
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.DocumentCount = direct.LazyPtr(in.GetDocumentCount())
	out.IndexCount = direct.LazyPtr(in.GetIndexCount())
	return out
}
func Backup_StatsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Backup_StatsObservedState) *pb.Backup_Stats {
	if in == nil {
		return nil
	}
	out := &pb.Backup_Stats{}
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.DocumentCount = direct.ValueOf(in.DocumentCount)
	out.IndexCount = direct.ValueOf(in.IndexCount)
	return out
}
func FirestoreBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.FirestoreBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreBackupObservedState{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
func FirestoreBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
func FirestoreBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.FirestoreBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreBackupSpec{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
func FirestoreBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Database
	// MISSING: DatabaseUid
	// MISSING: SnapshotTime
	// MISSING: ExpireTime
	// MISSING: Stats
	// MISSING: State
	return out
}
