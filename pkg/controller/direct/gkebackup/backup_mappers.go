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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkebackup/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GKEBackupBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.GKEBackupBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GKEBackupBackupObservedState{}
	// MISSING: Name
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Manual = direct.LazyPtr(in.GetManual())
	out.DeleteLockExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteLockExpireTime())
	out.RetainExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRetainExpireTime())
	out.EncryptionKey = EncryptionKey_FromProto(mapCtx, in.GetEncryptionKey())
	// out.AllNamespaces = direct.LazyPtr(in.GetAllNamespaces())
	if _, ok := in.BackupScope.(*pb.Backup_AllNamespaces); ok {
		// special handling for oneof bool field to ensure it is round-trippable
		out.AllNamespaces = direct.PtrTo(in.GetAllNamespaces())
	}
	out.SelectedNamespaces = Namespaces_FromProto(mapCtx, in.GetSelectedNamespaces())
	out.SelectedApplications = NamespacedNames_FromProto(mapCtx, in.GetSelectedApplications())
	out.ContainsVolumeData = direct.LazyPtr(in.GetContainsVolumeData())
	out.ContainsSecrets = direct.LazyPtr(in.GetContainsSecrets())
	out.ClusterMetadata = Backup_ClusterMetadataObservedState_FromProto(mapCtx, in.GetClusterMetadata())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateReason = direct.LazyPtr(in.GetStateReason())
	out.CompleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompleteTime())
	out.ResourceCount = direct.LazyPtr(in.GetResourceCount())
	out.VolumeCount = direct.LazyPtr(in.GetVolumeCount())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.PodCount = direct.LazyPtr(in.GetPodCount())
	out.ConfigBackupSizeBytes = direct.LazyPtr(in.GetConfigBackupSizeBytes())
	out.PermissiveMode = direct.LazyPtr(in.GetPermissiveMode())
	return out
}
func GKEBackupBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GKEBackupBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Manual = direct.ValueOf(in.Manual)
	out.DeleteLockExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteLockExpireTime)
	out.RetainExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.RetainExpireTime)
	out.EncryptionKey = EncryptionKey_ToProto(mapCtx, in.EncryptionKey)
	if in.AllNamespaces != nil {
		out.BackupScope = &pb.Backup_AllNamespaces{AllNamespaces: direct.ValueOf(in.AllNamespaces)}
	}
	if oneof := Namespaces_ToProto(mapCtx, in.SelectedNamespaces); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedNamespaces{SelectedNamespaces: oneof}
	}
	if oneof := NamespacedNames_ToProto(mapCtx, in.SelectedApplications); oneof != nil {
		out.BackupScope = &pb.Backup_SelectedApplications{SelectedApplications: oneof}
	}
	out.ContainsVolumeData = direct.ValueOf(in.ContainsVolumeData)
	out.ContainsSecrets = direct.ValueOf(in.ContainsSecrets)
	out.ClusterMetadata = Backup_ClusterMetadataObservedState_ToProto(mapCtx, in.ClusterMetadata)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.StateReason = direct.ValueOf(in.StateReason)
	out.CompleteTime = direct.StringTimestamp_ToProto(mapCtx, in.CompleteTime)
	out.ResourceCount = direct.ValueOf(in.ResourceCount)
	out.VolumeCount = direct.ValueOf(in.VolumeCount)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.Etag = direct.ValueOf(in.Etag)
	out.PodCount = direct.ValueOf(in.PodCount)
	out.ConfigBackupSizeBytes = direct.ValueOf(in.ConfigBackupSizeBytes)
	out.PermissiveMode = direct.ValueOf(in.PermissiveMode)
	return out
}
func Backup_ClusterMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup_ClusterMetadata) *krm.Backup_ClusterMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Backup_ClusterMetadataObservedState{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.K8sVersion = direct.LazyPtr(in.GetK8SVersion())
	out.BackupCRDVersions = in.GetBackupCrdVersions()
	out.GKEVersion = direct.LazyPtr(in.GetGkeVersion())
	out.AnthosVersion = direct.LazyPtr(in.GetAnthosVersion())
	return out
}
func Backup_ClusterMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Backup_ClusterMetadataObservedState) *pb.Backup_ClusterMetadata {
	if in == nil {
		return nil
	}
	out := &pb.Backup_ClusterMetadata{}
	out.Cluster = direct.ValueOf(in.Cluster)
	out.K8SVersion = direct.ValueOf(in.K8sVersion)
	out.BackupCrdVersions = in.BackupCRDVersions
	if oneof := Backup_ClusterMetadataObservedState_GkeVersion_ToProto(mapCtx, in.GKEVersion); oneof != nil {
		out.PlatformVersion = oneof
	}
	if oneof := Backup_ClusterMetadataObservedState_AnthosVersion_ToProto(mapCtx, in.AnthosVersion); oneof != nil {
		out.PlatformVersion = oneof
	}
	return out
}
func Backup_ClusterMetadataObservedState_GkeVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.Backup_ClusterMetadata_GkeVersion {
	if in == nil {
		return nil
	}
	out := &pb.Backup_ClusterMetadata_GkeVersion{}
	out.GkeVersion = direct.ValueOf(in)
	return out
}
func Backup_ClusterMetadataObservedState_AnthosVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.Backup_ClusterMetadata_AnthosVersion {
	if in == nil {
		return nil
	}
	out := &pb.Backup_ClusterMetadata_AnthosVersion{}
	out.AnthosVersion = direct.ValueOf(in)
	return out
}
