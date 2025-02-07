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

package alloydb

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/alloydb/apiv1/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AlloydbBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.AlloydbBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbBackupObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: Description
	// MISSING: ClusterUid
	// MISSING: ClusterName
	// MISSING: Reconciling
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloydbBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: Description
	// MISSING: ClusterUid
	// MISSING: ClusterName
	// MISSING: Reconciling
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloydbBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.AlloydbBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbBackupSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: Description
	// MISSING: ClusterUid
	// MISSING: ClusterName
	// MISSING: Reconciling
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloydbBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Type
	// MISSING: Description
	// MISSING: ClusterUid
	// MISSING: ClusterName
	// MISSING: Reconciling
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	// MISSING: Tags
	return out
}
func AlloydbClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloydbClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbClusterObservedState{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloydbClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbClusterSpec{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: BackupSource
	// MISSING: MigrationSource
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	// MISSING: DatabaseVersion
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	// MISSING: EncryptionInfo
	// MISSING: ContinuousBackupConfig
	// MISSING: ContinuousBackupInfo
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloydbConnectionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionInfo) *krm.AlloydbConnectionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbConnectionInfoObservedState{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func AlloydbConnectionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbConnectionInfoObservedState) *pb.ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionInfo{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func AlloydbConnectionInfoSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionInfo) *krm.AlloydbConnectionInfoSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbConnectionInfoSpec{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func AlloydbConnectionInfoSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbConnectionInfoSpec) *pb.ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionInfo{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func Backup_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.Backup {
	if in == nil {
		return nil
	}
	out := &krm.Backup{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: ClusterUid
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	// MISSING: Reconciling
	out.EncryptionConfig = EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	// MISSING: EncryptionInfo
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Annotations = in.Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	out.Tags = in.Tags
	return out
}
func Backup_ToProto(mapCtx *direct.MapContext, in *krm.Backup) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	out.Labels = in.Labels
	// MISSING: State
	out.Type = direct.Enum_ToProto[pb.Backup_Type](mapCtx, in.Type)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: ClusterUid
	out.ClusterName = direct.ValueOf(in.ClusterName)
	// MISSING: Reconciling
	out.EncryptionConfig = EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	// MISSING: EncryptionInfo
	out.Etag = direct.ValueOf(in.Etag)
	out.Annotations = in.Annotations
	// MISSING: SizeBytes
	// MISSING: ExpiryTime
	// MISSING: ExpiryQuantity
	// MISSING: SatisfiesPzs
	// MISSING: DatabaseVersion
	out.Tags = in.Tags
	return out
}
func BackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Type
	// MISSING: Description
	out.ClusterUid = direct.LazyPtr(in.GetClusterUid())
	// MISSING: ClusterName
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_FromProto(mapCtx, in.GetEncryptionInfo())
	// MISSING: Etag
	// MISSING: Annotations
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.ExpiryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpiryTime())
	out.ExpiryQuantity = Backup_QuantityBasedExpiry_FromProto(mapCtx, in.GetExpiryQuantity())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	// MISSING: Tags
	return out
}
func BackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	// MISSING: Type
	// MISSING: Description
	out.ClusterUid = direct.ValueOf(in.ClusterUid)
	// MISSING: ClusterName
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfo_ToProto(mapCtx, in.EncryptionInfo)
	// MISSING: Etag
	// MISSING: Annotations
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.ExpiryTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpiryTime)
	out.ExpiryQuantity = Backup_QuantityBasedExpiry_ToProto(mapCtx, in.ExpiryQuantity)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	// MISSING: Tags
	return out
}
func Backup_QuantityBasedExpiry_FromProto(mapCtx *direct.MapContext, in *pb.Backup_QuantityBasedExpiry) *krm.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &krm.Backup_QuantityBasedExpiry{}
	// MISSING: RetentionCount
	// MISSING: TotalRetentionCount
	return out
}
func Backup_QuantityBasedExpiry_ToProto(mapCtx *direct.MapContext, in *krm.Backup_QuantityBasedExpiry) *pb.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &pb.Backup_QuantityBasedExpiry{}
	// MISSING: RetentionCount
	// MISSING: TotalRetentionCount
	return out
}
func Backup_QuantityBasedExpiryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup_QuantityBasedExpiry) *krm.Backup_QuantityBasedExpiryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Backup_QuantityBasedExpiryObservedState{}
	out.RetentionCount = direct.LazyPtr(in.GetRetentionCount())
	out.TotalRetentionCount = direct.LazyPtr(in.GetTotalRetentionCount())
	return out
}
func Backup_QuantityBasedExpiryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Backup_QuantityBasedExpiryObservedState) *pb.Backup_QuantityBasedExpiry {
	if in == nil {
		return nil
	}
	out := &pb.Backup_QuantityBasedExpiry{}
	out.RetentionCount = direct.ValueOf(in.RetentionCount)
	out.TotalRetentionCount = direct.ValueOf(in.TotalRetentionCount)
	return out
}
func EncryptionConfig_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfig) *krm.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionConfig_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfig) *pb.EncryptionConfig {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	// MISSING: EncryptionType
	// MISSING: KMSKeyVersions
	return out
}
func EncryptionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfoObservedState{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	out.KMSKeyVersions = in.KmsKeyVersions
	return out
}
func EncryptionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfoObservedState) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_Type](mapCtx, in.EncryptionType)
	out.KmsKeyVersions = in.KMSKeyVersions
	return out
}
