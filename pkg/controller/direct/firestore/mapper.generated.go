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
// krm.group: firestore.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.firestore.admin.v1
// proto.service: google.firestore.v1

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	firestorepb "cloud.google.com/go/firestore/apiv1/firestorepb"
	krmfirestorev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	dayofweekpb "google.golang.org/genproto/googleapis/type/dayofweek"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func ArrayValue_v1alpha1_FromProto(mapCtx *direct.MapContext, in *firestorepb.ArrayValue) *krmfirestorev1alpha1.ArrayValue {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.ArrayValue{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, Value_v1alpha1_FromProto)
	return out
}
func ArrayValue_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.ArrayValue) *firestorepb.ArrayValue {
	if in == nil {
		return nil
	}
	out := &firestorepb.ArrayValue{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, Value_v1alpha1_ToProto)
	return out
}
func DailyRecurrence_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.DailyRecurrence) *krmfirestorev1alpha1.DailyRecurrence {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.DailyRecurrence{}
	return out
}
func DailyRecurrence_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.DailyRecurrence) *pb.DailyRecurrence {
	if in == nil {
		return nil
	}
	out := &pb.DailyRecurrence{}
	return out
}
func Database_CmekConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database_CmekConfig) *krm.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &krm.Database_CmekConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: ActiveKeyVersion
	return out
}
func Database_CmekConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Database_CmekConfig) *pb.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &pb.Database_CmekConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: ActiveKeyVersion
	return out
}
func Database_CmekConfigObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database_CmekConfig) *krm.Database_CmekConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Database_CmekConfigObservedState{}
	// MISSING: KMSKeyName
	out.ActiveKeyVersion = in.ActiveKeyVersion
	return out
}
func Database_CmekConfigObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Database_CmekConfigObservedState) *pb.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &pb.Database_CmekConfig{}
	// MISSING: KMSKeyName
	out.ActiveKeyVersion = in.ActiveKeyVersion
	return out
}
func Database_SourceInfo_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database_SourceInfo) *krm.Database_SourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.Database_SourceInfo{}
	out.Backup = Database_SourceInfo_BackupSource_v1beta1_FromProto(mapCtx, in.GetBackup())
	out.Operation = direct.LazyPtr(in.GetOperation())
	return out
}
func Database_SourceInfo_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Database_SourceInfo) *pb.Database_SourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.Database_SourceInfo{}
	if oneof := Database_SourceInfo_BackupSource_v1beta1_ToProto(mapCtx, in.Backup); oneof != nil {
		out.Source = &pb.Database_SourceInfo_Backup{Backup: oneof}
	}
	out.Operation = direct.ValueOf(in.Operation)
	return out
}
func Database_SourceInfo_BackupSource_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database_SourceInfo_BackupSource) *krm.Database_SourceInfo_BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.Database_SourceInfo_BackupSource{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	return out
}
func Database_SourceInfo_BackupSource_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Database_SourceInfo_BackupSource) *pb.Database_SourceInfo_BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.Database_SourceInfo_BackupSource{}
	out.Backup = direct.ValueOf(in.Backup)
	return out
}
func Field_IndexConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field_IndexConfig) *krmfirestorev1alpha1.Field_IndexConfig {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Field_IndexConfig{}
	out.Indexes = direct.Slice_FromProto(mapCtx, in.Indexes, Index_v1alpha1_FromProto)
	// MISSING: UsesAncestorConfig
	// MISSING: AncestorField
	// MISSING: Reverting
	return out
}
func Field_IndexConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Field_IndexConfig) *pb.Field_IndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_IndexConfig{}
	out.Indexes = direct.Slice_ToProto(mapCtx, in.Indexes, Index_v1alpha1_ToProto)
	// MISSING: UsesAncestorConfig
	// MISSING: AncestorField
	// MISSING: Reverting
	return out
}
func Field_IndexConfig_ObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field_IndexConfig) *krmfirestorev1alpha1.Field_IndexConfig_ObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Field_IndexConfig_ObservedState{}
	out.Indexes = direct.Slice_FromProto(mapCtx, in.Indexes, Index_ObservedState_v1alpha1_FromProto)
	out.UsesAncestorConfig = direct.LazyPtr(in.GetUsesAncestorConfig())
	out.AncestorField = direct.LazyPtr(in.GetAncestorField())
	out.Reverting = direct.LazyPtr(in.GetReverting())
	return out
}
func Field_IndexConfig_ObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Field_IndexConfig_ObservedState) *pb.Field_IndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_IndexConfig{}
	out.Indexes = direct.Slice_ToProto(mapCtx, in.Indexes, Index_ObservedState_v1alpha1_ToProto)
	out.UsesAncestorConfig = direct.ValueOf(in.UsesAncestorConfig)
	out.AncestorField = direct.ValueOf(in.AncestorField)
	out.Reverting = direct.ValueOf(in.Reverting)
	return out
}
func Field_TTLConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field_TtlConfig) *krmfirestorev1alpha1.Field_TTLConfig {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Field_TTLConfig{}
	// MISSING: State
	return out
}
func Field_TTLConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Field_TTLConfig) *pb.Field_TtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_TtlConfig{}
	// MISSING: State
	return out
}
func Field_TTLConfigObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field_TtlConfig) *krmfirestorev1alpha1.Field_TTLConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Field_TTLConfigObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Field_TTLConfigObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Field_TTLConfigObservedState) *pb.Field_TtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_TtlConfig{}
	out.State = direct.Enum_ToProto[pb.Field_TtlConfig_State](mapCtx, in.State)
	return out
}
func FirestoreBackupScheduleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krmfirestorev1alpha1.FirestoreBackupScheduleObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.FirestoreBackupScheduleObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FirestoreBackupScheduleObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.FirestoreBackupScheduleObservedState) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func FirestoreBackupScheduleSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupSchedule) *krmfirestorev1alpha1.FirestoreBackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.FirestoreBackupScheduleSpec{}
	out.Retention = direct.StringDuration_FromProto(mapCtx, in.GetRetention())
	out.DailyRecurrence = DailyRecurrence_v1alpha1_FromProto(mapCtx, in.GetDailyRecurrence())
	out.WeeklyRecurrence = WeeklyRecurrence_v1alpha1_FromProto(mapCtx, in.GetWeeklyRecurrence())
	return out
}
func FirestoreBackupScheduleSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.FirestoreBackupScheduleSpec) *pb.BackupSchedule {
	if in == nil {
		return nil
	}
	out := &pb.BackupSchedule{}
	out.Retention = direct.StringDuration_ToProto(mapCtx, in.Retention)
	if oneof := DailyRecurrence_v1alpha1_ToProto(mapCtx, in.DailyRecurrence); oneof != nil {
		out.Recurrence = &pb.BackupSchedule_DailyRecurrence{DailyRecurrence: oneof}
	}
	if oneof := WeeklyRecurrence_v1alpha1_ToProto(mapCtx, in.WeeklyRecurrence); oneof != nil {
		out.Recurrence = &pb.BackupSchedule_WeeklyRecurrence{WeeklyRecurrence: oneof}
	}
	return out
}
func FirestoreDatabaseObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.FirestoreDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDatabaseObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DeleteTime
	// MISSING: Type
	out.VersionRetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetVersionRetentionPeriod())
	out.EarliestVersionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestVersionTime())
	// MISSING: AppEngineIntegrationMode
	out.KeyPrefix = direct.LazyPtr(in.GetKeyPrefix())
	// MISSING: DeleteProtectionState
	// MISSING: CmekConfig
	// MISSING: PreviousID
	// MISSING: SourceInfo
	// MISSING: Tags
	// MISSING: FreeTier
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: DatabaseEdition
	return out
}
func FirestoreDatabaseObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DeleteTime
	// MISSING: Type
	out.VersionRetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.VersionRetentionPeriod)
	out.EarliestVersionTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestVersionTime)
	// MISSING: AppEngineIntegrationMode
	out.KeyPrefix = direct.ValueOf(in.KeyPrefix)
	// MISSING: DeleteProtectionState
	// MISSING: CmekConfig
	// MISSING: PreviousID
	// MISSING: SourceInfo
	// MISSING: Tags
	// MISSING: FreeTier
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: DatabaseEdition
	return out
}
func FirestoreDatabaseSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.FirestoreDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreDatabaseSpec{}
	// MISSING: Name
	// MISSING: DeleteTime
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	// MISSING: Type
	out.ConcurrencyMode = direct.Enum_FromProto(mapCtx, in.GetConcurrencyMode())
	out.PointInTimeRecoveryEnablement = direct.Enum_FromProto(mapCtx, in.GetPointInTimeRecoveryEnablement())
	// MISSING: AppEngineIntegrationMode
	// MISSING: DeleteProtectionState
	// MISSING: CmekConfig
	// MISSING: PreviousID
	// MISSING: SourceInfo
	// MISSING: Tags
	// MISSING: FreeTier
	// MISSING: DatabaseEdition
	return out
}
func FirestoreDatabaseSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreDatabaseSpec) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	// MISSING: DeleteTime
	out.LocationId = direct.ValueOf(in.LocationID)
	// MISSING: Type
	out.ConcurrencyMode = direct.Enum_ToProto[pb.Database_ConcurrencyMode](mapCtx, in.ConcurrencyMode)
	out.PointInTimeRecoveryEnablement = direct.Enum_ToProto[pb.Database_PointInTimeRecoveryEnablement](mapCtx, in.PointInTimeRecoveryEnablement)
	// MISSING: AppEngineIntegrationMode
	// MISSING: DeleteProtectionState
	// MISSING: CmekConfig
	// MISSING: PreviousID
	// MISSING: SourceInfo
	// MISSING: Tags
	// MISSING: FreeTier
	// MISSING: DatabaseEdition
	return out
}
func FirestoreDocumentObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *firestorepb.Document) *krmfirestorev1alpha1.FirestoreDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.FirestoreDocumentObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FirestoreDocumentObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.FirestoreDocumentObservedState) *firestorepb.Document {
	if in == nil {
		return nil
	}
	out := &firestorepb.Document{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func FirestoreFieldObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krmfirestorev1alpha1.FirestoreFieldObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.FirestoreFieldObservedState{}
	// MISSING: Name
	out.IndexConfig = Field_IndexConfig_ObservedState_v1alpha1_FromProto(mapCtx, in.GetIndexConfig())
	out.TTLConfig = Field_TTLConfigObservedState_v1alpha1_FromProto(mapCtx, in.GetTtlConfig())
	return out
}
func FirestoreFieldObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.FirestoreFieldObservedState) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	// MISSING: Name
	out.IndexConfig = Field_IndexConfig_ObservedState_v1alpha1_ToProto(mapCtx, in.IndexConfig)
	out.TtlConfig = Field_TTLConfigObservedState_v1alpha1_ToProto(mapCtx, in.TTLConfig)
	return out
}
func FirestoreFieldSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krmfirestorev1alpha1.FirestoreFieldSpec {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.FirestoreFieldSpec{}
	// MISSING: Name
	out.IndexConfig = Field_IndexConfig_v1alpha1_FromProto(mapCtx, in.GetIndexConfig())
	return out
}
func FirestoreFieldSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.FirestoreFieldSpec) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	// MISSING: Name
	out.IndexConfig = Field_IndexConfig_v1alpha1_ToProto(mapCtx, in.IndexConfig)
	return out
}
func FirestoreIndexSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexSpec{}
	// MISSING: Name
	out.QueryScope = direct.Enum_FromProto(mapCtx, in.GetQueryScope())
	// MISSING: APIScope
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, IndexFields_v1beta1_FromProto)
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	out.QueryScope = direct.Enum_ToProto[pb.Index_QueryScope](mapCtx, in.QueryScope)
	// MISSING: APIScope
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, IndexFields_v1beta1_ToProto)
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexStatus) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func Index_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krmfirestorev1alpha1.Index {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Index{}
	// MISSING: Name
	out.QueryScope = direct.Enum_FromProto(mapCtx, in.GetQueryScope())
	out.APIScope = direct.Enum_FromProto(mapCtx, in.GetApiScope())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Index_IndexField_v1alpha1_FromProto)
	// MISSING: State
	out.Density = direct.Enum_FromProto(mapCtx, in.GetDensity())
	out.Multikey = direct.LazyPtr(in.GetMultikey())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	return out
}
func Index_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Index) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	out.QueryScope = direct.Enum_ToProto[pb.Index_QueryScope](mapCtx, in.QueryScope)
	out.ApiScope = direct.Enum_ToProto[pb.Index_ApiScope](mapCtx, in.APIScope)
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Index_IndexField_v1alpha1_ToProto)
	// MISSING: State
	out.Density = direct.Enum_ToProto[pb.Index_Density](mapCtx, in.Density)
	out.Multikey = direct.ValueOf(in.Multikey)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	return out
}
func IndexFields_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField) *krm.IndexFields {
	if in == nil {
		return nil
	}
	out := &krm.IndexFields{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Order = direct.Enum_FromProto(mapCtx, in.GetOrder())
	out.ArrayConfig = direct.Enum_FromProto(mapCtx, in.GetArrayConfig())
	// MISSING: VectorConfig
	return out
}
func IndexFields_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.IndexFields) *pb.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	if oneof := IndexFields_Order_ToProto(mapCtx, in.Order); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := IndexFields_ArrayConfig_ToProto(mapCtx, in.ArrayConfig); oneof != nil {
		out.ValueMode = oneof
	}
	// MISSING: VectorConfig
	return out
}
func Index_IndexField_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField) *krmfirestorev1alpha1.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Index_IndexField{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Order = direct.Enum_FromProto(mapCtx, in.GetOrder())
	out.ArrayConfig = direct.Enum_FromProto(mapCtx, in.GetArrayConfig())
	out.VectorConfig = Index_IndexField_VectorConfig_v1alpha1_FromProto(mapCtx, in.GetVectorConfig())
	return out
}
func Index_IndexField_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Index_IndexField) *pb.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	if oneof := Index_IndexField_Order_ToProto(mapCtx, in.Order); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := Index_IndexField_ArrayConfig_ToProto(mapCtx, in.ArrayConfig); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := Index_IndexField_VectorConfig_v1alpha1_ToProto(mapCtx, in.VectorConfig); oneof != nil {
		out.ValueMode = &pb.Index_IndexField_VectorConfig_{VectorConfig: oneof}
	}
	return out
}
func Index_IndexField_Order_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_Order_ {
	if in == nil {
		return nil
	}
	return &pb.Index_IndexField_Order_{Order: direct.Enum_ToProto[pb.Index_IndexField_Order](mapCtx, in)}
}
func Index_IndexField_ArrayConfig_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_ArrayConfig_ {
	if in == nil {
		return nil
	}
	return &pb.Index_IndexField_ArrayConfig_{ArrayConfig: direct.Enum_ToProto[pb.Index_IndexField_ArrayConfig](mapCtx, in)}
}
func Index_IndexField_VectorConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig) *krmfirestorev1alpha1.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Index_IndexField_VectorConfig{}
	out.Dimension = direct.LazyPtr(in.GetDimension())
	out.Flat = Index_IndexField_VectorConfig_FlatIndex_v1alpha1_FromProto(mapCtx, in.GetFlat())
	return out
}
func Index_IndexField_VectorConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Index_IndexField_VectorConfig) *pb.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig{}
	out.Dimension = direct.ValueOf(in.Dimension)
	if oneof := Index_IndexField_VectorConfig_FlatIndex_v1alpha1_ToProto(mapCtx, in.Flat); oneof != nil {
		out.Type = &pb.Index_IndexField_VectorConfig_Flat{Flat: oneof}
	}
	return out
}
func Index_IndexField_VectorConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig) *krm.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField_VectorConfig{}
	out.Dimension = direct.LazyPtr(in.GetDimension())
	out.Flat = Index_IndexField_VectorConfig_FlatIndex_v1beta1_FromProto(mapCtx, in.GetFlat())
	return out
}
func Index_IndexField_VectorConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_VectorConfig) *pb.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig{}
	out.Dimension = direct.ValueOf(in.Dimension)
	if oneof := Index_IndexField_VectorConfig_FlatIndex_v1beta1_ToProto(mapCtx, in.Flat); oneof != nil {
		out.Type = &pb.Index_IndexField_VectorConfig_Flat{Flat: oneof}
	}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig_FlatIndex) *krmfirestorev1alpha1.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Index_IndexField_VectorConfig_FlatIndex) *pb.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig_FlatIndex) *krm.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_VectorConfig_FlatIndex) *pb.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
func Index_ObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krmfirestorev1alpha1.Index_ObservedState {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.Index_ObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func Index_ObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.Index_ObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	out.State = direct.Enum_ToProto[pb.Index_State](mapCtx, in.State)
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func MapValue_v1alpha1_FromProto(mapCtx *direct.MapContext, in *firestorepb.MapValue) *krmfirestorev1alpha1.MapValue {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.MapValue{}
	// MISSING: Fields
	return out
}
func MapValue_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.MapValue) *firestorepb.MapValue {
	if in == nil {
		return nil
	}
	out := &firestorepb.MapValue{}
	// MISSING: Fields
	return out
}
func Value_NullValue_ToProto(mapCtx *direct.MapContext, in *string) *firestorepb.Value_NullValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_NullValue{NullValue: direct.Enum_ToProto[structpb.NullValue](mapCtx, in)}
}
func Value_BooleanValue_ToProto(mapCtx *direct.MapContext, in *bool) *firestorepb.Value_BooleanValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_BooleanValue{BooleanValue: *in}
}
func Value_IntegerValue_ToProto(mapCtx *direct.MapContext, in *int64) *firestorepb.Value_IntegerValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_IntegerValue{IntegerValue: *in}
}
func Value_DoubleValue_ToProto(mapCtx *direct.MapContext, in *float64) *firestorepb.Value_DoubleValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_DoubleValue{DoubleValue: *in}
}
func Value_StringValue_ToProto(mapCtx *direct.MapContext, in *string) *firestorepb.Value_StringValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_StringValue{StringValue: *in}
}
func Value_ReferenceValue_ToProto(mapCtx *direct.MapContext, in *string) *firestorepb.Value_ReferenceValue {
	if in == nil {
		return nil
	}
	return &firestorepb.Value_ReferenceValue{ReferenceValue: *in}
}
func WeeklyRecurrence_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyRecurrence) *krmfirestorev1alpha1.WeeklyRecurrence {
	if in == nil {
		return nil
	}
	out := &krmfirestorev1alpha1.WeeklyRecurrence{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	return out
}
func WeeklyRecurrence_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmfirestorev1alpha1.WeeklyRecurrence) *pb.WeeklyRecurrence {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyRecurrence{}
	out.Day = direct.Enum_ToProto[dayofweekpb.DayOfWeek](mapCtx, in.Day)
	return out
}
