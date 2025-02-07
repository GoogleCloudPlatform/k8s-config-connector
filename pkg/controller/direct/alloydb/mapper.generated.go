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
func AlloydbDatabaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.AlloydbDatabaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbDatabaseObservedState{}
	// MISSING: Name
	// MISSING: Charset
	// MISSING: Collation
	return out
}
func AlloydbDatabaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbDatabaseObservedState) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	// MISSING: Charset
	// MISSING: Collation
	return out
}
func AlloydbDatabaseSpec_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.AlloydbDatabaseSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbDatabaseSpec{}
	// MISSING: Name
	// MISSING: Charset
	// MISSING: Collation
	return out
}
func AlloydbDatabaseSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbDatabaseSpec) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	// MISSING: Name
	// MISSING: Charset
	// MISSING: Collation
	return out
}
func AlloydbSupportedDatabaseFlagObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SupportedDatabaseFlag) *krm.AlloydbSupportedDatabaseFlagObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbSupportedDatabaseFlagObservedState{}
	// MISSING: StringRestrictions
	// MISSING: IntegerRestrictions
	// MISSING: Name
	// MISSING: FlagName
	// MISSING: ValueType
	// MISSING: AcceptsMultipleValues
	// MISSING: SupportedDbVersions
	// MISSING: RequiresDbRestart
	return out
}
func AlloydbSupportedDatabaseFlagObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbSupportedDatabaseFlagObservedState) *pb.SupportedDatabaseFlag {
	if in == nil {
		return nil
	}
	out := &pb.SupportedDatabaseFlag{}
	// MISSING: StringRestrictions
	// MISSING: IntegerRestrictions
	// MISSING: Name
	// MISSING: FlagName
	// MISSING: ValueType
	// MISSING: AcceptsMultipleValues
	// MISSING: SupportedDbVersions
	// MISSING: RequiresDbRestart
	return out
}
func AlloydbSupportedDatabaseFlagSpec_FromProto(mapCtx *direct.MapContext, in *pb.SupportedDatabaseFlag) *krm.AlloydbSupportedDatabaseFlagSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbSupportedDatabaseFlagSpec{}
	// MISSING: StringRestrictions
	// MISSING: IntegerRestrictions
	// MISSING: Name
	// MISSING: FlagName
	// MISSING: ValueType
	// MISSING: AcceptsMultipleValues
	// MISSING: SupportedDbVersions
	// MISSING: RequiresDbRestart
	return out
}
func AlloydbSupportedDatabaseFlagSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbSupportedDatabaseFlagSpec) *pb.SupportedDatabaseFlag {
	if in == nil {
		return nil
	}
	out := &pb.SupportedDatabaseFlag{}
	// MISSING: StringRestrictions
	// MISSING: IntegerRestrictions
	// MISSING: Name
	// MISSING: FlagName
	// MISSING: ValueType
	// MISSING: AcceptsMultipleValues
	// MISSING: SupportedDbVersions
	// MISSING: RequiresDbRestart
	return out
}
func AlloydbUserObservedState_FromProto(mapCtx *direct.MapContext, in *pb.User) *krm.AlloydbUserObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbUserObservedState{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloydbUserObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbUserObservedState) *pb.User {
	if in == nil {
		return nil
	}
	out := &pb.User{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloydbUserSpec_FromProto(mapCtx *direct.MapContext, in *pb.User) *krm.AlloydbUserSpec {
	if in == nil {
		return nil
	}
	out := &krm.AlloydbUserSpec{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func AlloydbUserSpec_ToProto(mapCtx *direct.MapContext, in *krm.AlloydbUserSpec) *pb.User {
	if in == nil {
		return nil
	}
	out := &pb.User{}
	// MISSING: Name
	// MISSING: Password
	// MISSING: DatabaseRoles
	// MISSING: UserType
	// MISSING: KeepExtraRoles
	return out
}
func Database_FromProto(mapCtx *direct.MapContext, in *pb.Database) *krm.Database {
	if in == nil {
		return nil
	}
	out := &krm.Database{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Charset = direct.LazyPtr(in.GetCharset())
	out.Collation = direct.LazyPtr(in.GetCollation())
	return out
}
func Database_ToProto(mapCtx *direct.MapContext, in *krm.Database) *pb.Database {
	if in == nil {
		return nil
	}
	out := &pb.Database{}
	out.Name = direct.ValueOf(in.Name)
	out.Charset = direct.ValueOf(in.Charset)
	out.Collation = direct.ValueOf(in.Collation)
	return out
}
