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
	pb "cloud.google.com/go/alloydb/apiv1/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
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
func ConnectionInfo_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionInfo) *krm.ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func ConnectionInfo_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionInfo) *pb.ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionInfo{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: IPAddress
	// MISSING: PublicIPAddress
	// MISSING: InstanceUid
	return out
}
func ConnectionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionInfo) *krm.ConnectionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionInfoObservedState{}
	// MISSING: Name
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.PublicIPAddress = direct.LazyPtr(in.GetPublicIpAddress())
	out.InstanceUid = direct.LazyPtr(in.GetInstanceUid())
	return out
}
func ConnectionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionInfoObservedState) *pb.ConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionInfo{}
	// MISSING: Name
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.PublicIpAddress = direct.ValueOf(in.PublicIPAddress)
	out.InstanceUid = direct.ValueOf(in.InstanceUid)
	return out
}
