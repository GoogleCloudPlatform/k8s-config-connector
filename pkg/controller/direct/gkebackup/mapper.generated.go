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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkebackup/apiv1/gkebackuppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	out.Labels = in.Labels
	out.DeleteLockDays = direct.LazyPtr(in.GetDeleteLockDays())
	// MISSING: DeleteLockExpireTime
	out.RetainDays = direct.LazyPtr(in.GetRetainDays())
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	out.Labels = in.Labels
	out.DeleteLockDays = direct.ValueOf(in.DeleteLockDays)
	// MISSING: DeleteLockExpireTime
	out.RetainDays = direct.ValueOf(in.RetainDays)
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	out.Description = direct.ValueOf(in.Description)
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Manual = direct.LazyPtr(in.GetManual())
	// MISSING: Labels
	// MISSING: DeleteLockDays
	out.DeleteLockExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteLockExpireTime())
	// MISSING: RetainDays
	out.RetainExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRetainExpireTime())
	out.EncryptionKey = EncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.ContainsVolumeData = direct.LazyPtr(in.GetContainsVolumeData())
	out.ContainsSecrets = direct.LazyPtr(in.GetContainsSecrets())
	out.ClusterMetadata = Backup_ClusterMetadata_FromProto(mapCtx, in.GetClusterMetadata())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.ResourceCount = direct.LazyPtr(in.GetResourceCount())
	out.VolumeCount = direct.LazyPtr(in.GetVolumeCount())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Description
	out.PodCount = direct.LazyPtr(in.GetPodCount())
	out.ConfigBackupSizeBytes = direct.LazyPtr(in.GetConfigBackupSizeBytes())
	out.PermissiveMode = direct.LazyPtr(in.GetPermissiveMode())
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Manual = direct.ValueOf(in.Manual)
	// MISSING: Labels
	// MISSING: DeleteLockDays
	out.DeleteLockExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteLockExpireTime)
	// MISSING: RetainDays
	out.RetainExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.RetainExpireTime)
	out.EncryptionKey = EncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	if oneof := BackupObservedState_AllNamespaces_ToProto(mapCtx, in.AllNamespaces); oneof != nil {
		out.BackupScope = oneof
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedApplications{SelectedApplications: oneof}
	}
	out.ContainsVolumeData = direct.ValueOf(in.ContainsVolumeData)
	out.ContainsSecrets = direct.ValueOf(in.ContainsSecrets)
	out.ClusterMetadata = Backup_ClusterMetadata_ToProto(mapCtx, in.ClusterMetadata)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.ResourceCount = direct.ValueOf(in.ResourceCount)
	out.VolumeCount = direct.ValueOf(in.VolumeCount)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Description
	out.PodCount = direct.ValueOf(in.PodCount)
	out.ConfigBackupSizeBytes = direct.ValueOf(in.ConfigBackupSizeBytes)
	out.PermissiveMode = direct.ValueOf(in.PermissiveMode)
	return out
}
func EncryptionKey_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionKey) *krm.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionKey{}
	out.GcpKMSEncryptionKey = direct.LazyPtr(in.GetGcpKmsEncryptionKey())
	return out
}
func EncryptionKey_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionKey) *pb.EncryptionKey {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionKey{}
	out.GcpKmsEncryptionKey = direct.ValueOf(in.GcpKMSEncryptionKey)
	return out
}
func GkebackupBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.GkebackupBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupBackupObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	// MISSING: Labels
	// MISSING: DeleteLockDays
	// MISSING: DeleteLockExpireTime
	// MISSING: RetainDays
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	// MISSING: Description
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func GkebackupBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	// MISSING: Labels
	// MISSING: DeleteLockDays
	// MISSING: DeleteLockExpireTime
	// MISSING: RetainDays
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	// MISSING: Description
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func GkebackupBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.GkebackupBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.GkebackupBackupSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	// MISSING: Labels
	// MISSING: DeleteLockDays
	// MISSING: DeleteLockExpireTime
	// MISSING: RetainDays
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	// MISSING: Description
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func GkebackupBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.GkebackupBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Manual
	// MISSING: Labels
	// MISSING: DeleteLockDays
	// MISSING: DeleteLockExpireTime
	// MISSING: RetainDays
	// MISSING: RetainExpireTime
	// MISSING: EncryptionKey
	// MISSING: AllNamespaces
	// MISSING: SelectedNamespaces
	// MISSING: SelectedApplications
	// MISSING: ContainsVolumeData
	// MISSING: ContainsSecrets
	// MISSING: ClusterMetadata
	// MISSING: State
	// MISSING: StateReason
	// MISSING: CompleteTime
	// MISSING: ResourceCount
	// MISSING: VolumeCount
	// MISSING: SizeBytes
	// MISSING: Etag
	// MISSING: Description
	// MISSING: PodCount
	// MISSING: ConfigBackupSizeBytes
	// MISSING: PermissiveMode
	return out
}
func NamespacedName_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedName) *krm.NamespacedName {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedName{}
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func NamespacedName_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedName) *pb.NamespacedName {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedName{}
	out.Namespace = direct.ValueOf(in.Namespace)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func NamespacedNames_FromProto(mapCtx *direct.MapContext, in *pb.NamespacedNames) *krm.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &krm.NamespacedNames{}
	out.NamespacedNames = direct.Slice_FromProto(mapCtx, in.NamespacedNames, NamespacedName_FromProto)
	return out
}
func NamespacedNames_ToProto(mapCtx *direct.MapContext, in *krm.NamespacedNames) *pb.NamespacedNames {
	if in == nil {
		return nil
	}
	out := &pb.NamespacedNames{}
	out.NamespacedNames = direct.Slice_ToProto(mapCtx, in.NamespacedNames, NamespacedName_ToProto)
	return out
}
func Namespaces_FromProto(mapCtx *direct.MapContext, in *pb.Namespaces) *krm.Namespaces {
	if in == nil {
		return nil
	}
	out := &krm.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
func Namespaces_ToProto(mapCtx *direct.MapContext, in *krm.Namespaces) *pb.Namespaces {
	if in == nil {
		return nil
	}
	out := &pb.Namespaces{}
	out.Namespaces = in.Namespaces
	return out
}
