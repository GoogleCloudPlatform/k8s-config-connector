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
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.bigtable.admin.v2

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppProfile_DataBoostIsolationReadOnly_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_DataBoostIsolationReadOnly) *krmv1beta1.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_DataBoostIsolationReadOnly{}
	out.ComputeBillingOwner = direct.Enum_FromProto(mapCtx, in.GetComputeBillingOwner())
	return out
}
func AppProfile_DataBoostIsolationReadOnly_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_DataBoostIsolationReadOnly) *pb.AppProfile_DataBoostIsolationReadOnly {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_DataBoostIsolationReadOnly{}
	if oneof := AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx, in.ComputeBillingOwner); oneof != nil {
		out.ComputeBillingOwner = oneof
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny) *krmv1beta1.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	out.RowAffinity = AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx, in.GetRowAffinity())
	return out
}
func AppProfile_MultiClusterRoutingUseAny_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_MultiClusterRoutingUseAny) *pb.AppProfile_MultiClusterRoutingUseAny {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny{}
	out.ClusterIds = in.ClusterIds
	if oneof := AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx, in.RowAffinity); oneof != nil {
		out.Affinity = &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity_{RowAffinity: oneof}
	}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_MultiClusterRoutingUseAny_RowAffinity_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_MultiClusterRoutingUseAny_RowAffinity) *pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_MultiClusterRoutingUseAny_RowAffinity{}
	return out
}
func AppProfile_SingleClusterRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_SingleClusterRouting) *krmv1beta1.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_SingleClusterRouting{}
	out.ClusterID = direct.LazyPtr(in.GetClusterId())
	out.AllowTransactionalWrites = direct.LazyPtr(in.GetAllowTransactionalWrites())
	return out
}
func AppProfile_SingleClusterRouting_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_SingleClusterRouting) *pb.AppProfile_SingleClusterRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_SingleClusterRouting{}
	out.ClusterId = direct.ValueOf(in.ClusterID)
	out.AllowTransactionalWrites = direct.ValueOf(in.AllowTransactionalWrites)
	return out
}
func AppProfile_StandardIsolation_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile_StandardIsolation) *krmv1beta1.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func AppProfile_StandardIsolation_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AppProfile_StandardIsolation) *pb.AppProfile_StandardIsolation {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile_StandardIsolation{}
	out.Priority = direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in.Priority)
	return out
}
func AuthorizedView_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmv1alpha1.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.AuthorizedView{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SubsetView = AuthorizedView_SubsetView_FromProto(mapCtx, in.GetSubsetView())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	return out
}
func AuthorizedView_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.AuthorizedView) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := AuthorizedView_SubsetView_ToProto(mapCtx, in.SubsetView); oneof != nil {
		out.AuthorizedView = &pb.AuthorizedView_SubsetView_{SubsetView: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	return out
}
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krmv1alpha1.Backup {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Backup{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	// MISSING: SourceBackup
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SizeBytes
	// MISSING: State
	// MISSING: EncryptionInfo
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.HotToStandardTime = direct.StringTimestamp_FromProto(mapCtx, in.GetHotToStandardTime())
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	// MISSING: SourceBackup
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: SizeBytes
	// MISSING: State
	// MISSING: EncryptionInfo
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	out.HotToStandardTime = direct.StringTimestamp_ToProto(mapCtx, in.HotToStandardTime)
	return out
}
func BigtableAppProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppProfile) *krmv1beta1.BigtableAppProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigtableAppProfileObservedState{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAppProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigtableAppProfileObservedState) *pb.AppProfile {
	if in == nil {
		return nil
	}
	out := &pb.AppProfile{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Priority
	return out
}
func BigtableAuthorizedViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizedView) *krmv1alpha1.BigtableAuthorizedViewObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigtableAuthorizedViewObservedState{}
	// MISSING: Name
	return out
}
func BigtableAuthorizedViewObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigtableAuthorizedViewObservedState) *pb.AuthorizedView {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizedView{}
	// MISSING: Name
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krmv1beta1.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	// MISSING: KMSKeyVersion
	// (near miss): "KMSKeyVersion" vs "KmsKeyVersion"
	return out
}
