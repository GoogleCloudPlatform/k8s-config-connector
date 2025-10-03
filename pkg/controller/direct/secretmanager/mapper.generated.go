// Copyright 2024 Google LLC
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

package secretmanager

import (
	"strconv"

	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CustomerManagedEncryptionStatus_FromProto(mapCtx *direct.MapContext, in *pb.CustomerManagedEncryptionStatus) *krm.CustomerManagedEncryptionStatus {
	if in == nil {
		return nil
	}
	out := &krm.CustomerManagedEncryptionStatus{}
	out.KmsKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	return out
}
func CustomerManagedEncryptionStatus_ToProto(mapCtx *direct.MapContext, in *krm.CustomerManagedEncryptionStatus) *pb.CustomerManagedEncryptionStatus {
	if in == nil {
		return nil
	}
	out := &pb.CustomerManagedEncryptionStatus{}
	out.KmsKeyVersionName = direct.ValueOf(in.KmsKeyVersionName)
	return out
}
func ReplicationStatus_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationStatus) *krm.ReplicationStatus {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationStatus{}
	out.Automatic = ReplicationStatus_AutomaticStatus_FromProto(mapCtx, in.GetAutomatic())
	out.UserManaged = ReplicationStatus_UserManagedStatus_FromProto(mapCtx, in.GetUserManaged())
	return out
}
func ReplicationStatus_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationStatus) *pb.ReplicationStatus {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationStatus{}
	if oneof := ReplicationStatus_AutomaticStatus_ToProto(mapCtx, in.Automatic); oneof != nil {
		out.ReplicationStatus = &pb.ReplicationStatus_Automatic{Automatic: oneof}
	}
	if oneof := ReplicationStatus_UserManagedStatus_ToProto(mapCtx, in.UserManaged); oneof != nil {
		out.ReplicationStatus = &pb.ReplicationStatus_UserManaged{UserManaged: oneof}
	}
	return out
}
func ReplicationStatus_AutomaticStatus_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationStatus_AutomaticStatus) *krm.ReplicationStatus_AutomaticStatus {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationStatus_AutomaticStatus{}
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func ReplicationStatus_AutomaticStatus_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationStatus_AutomaticStatus) *pb.ReplicationStatus_AutomaticStatus {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationStatus_AutomaticStatus{}
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}
func ReplicationStatus_UserManagedStatus_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationStatus_UserManagedStatus) *krm.ReplicationStatus_UserManagedStatus {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationStatus_UserManagedStatus{}
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, ReplicationStatus_UserManagedStatus_ReplicaStatus_FromProto)
	return out
}
func ReplicationStatus_UserManagedStatus_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationStatus_UserManagedStatus) *pb.ReplicationStatus_UserManagedStatus {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationStatus_UserManagedStatus{}
	out.Replicas = direct.Slice_ToProto(mapCtx, in.Replicas, ReplicationStatus_UserManagedStatus_ReplicaStatus_ToProto)
	return out
}
func ReplicationStatus_UserManagedStatus_ReplicaStatus_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationStatus_UserManagedStatus_ReplicaStatus) *krm.ReplicationStatus_UserManagedStatus_ReplicaStatus {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationStatus_UserManagedStatus_ReplicaStatus{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func ReplicationStatus_UserManagedStatus_ReplicaStatus_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationStatus_UserManagedStatus_ReplicaStatus) *pb.ReplicationStatus_UserManagedStatus_ReplicaStatus {
	if in == nil {
		return nil
	}
	out := &pb.ReplicationStatus_UserManagedStatus_ReplicaStatus{}
	out.Location = direct.ValueOf(in.Location)
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}
func Replication_Automatic_FromProto(mapCtx *direct.MapContext, in *pb.Replication_Automatic) *krm.Replication_Automatic {
	if in == nil {
		return nil
	}
	out := &krm.Replication_Automatic{}
	out.CustomerManagedEncryption = CustomerManagedEncryption_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func Replication_Automatic_ToProto(mapCtx *direct.MapContext, in *krm.Replication_Automatic) *pb.Replication_Automatic {
	if in == nil {
		return nil
	}
	out := &pb.Replication_Automatic{}
	out.CustomerManagedEncryption = CustomerManagedEncryption_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}
func Replication_UserManaged_FromProto(mapCtx *direct.MapContext, in *pb.Replication_UserManaged) *krm.Replication_UserManaged {
	if in == nil {
		return nil
	}
	out := &krm.Replication_UserManaged{}
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, Replication_UserManaged_Replica_FromProto)
	return out
}
func Replication_UserManaged_ToProto(mapCtx *direct.MapContext, in *krm.Replication_UserManaged) *pb.Replication_UserManaged {
	if in == nil {
		return nil
	}
	out := &pb.Replication_UserManaged{}
	out.Replicas = direct.Slice_ToProto(mapCtx, in.Replicas, Replication_UserManaged_Replica_ToProto)
	return out
}
func Replication_UserManaged_Replica_FromProto(mapCtx *direct.MapContext, in *pb.Replication_UserManaged_Replica) *krm.Replication_UserManaged_Replica {
	if in == nil {
		return nil
	}
	out := &krm.Replication_UserManaged_Replica{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.CustomerManagedEncryption = CustomerManagedEncryption_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func Replication_UserManaged_Replica_ToProto(mapCtx *direct.MapContext, in *krm.Replication_UserManaged_Replica) *pb.Replication_UserManaged_Replica {
	if in == nil {
		return nil
	}
	out := &pb.Replication_UserManaged_Replica{}
	out.Location = direct.ValueOf(in.Location)
	out.CustomerManagedEncryption = CustomerManagedEncryption_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}
func Rotation_FromProto(mapCtx *direct.MapContext, in *pb.Rotation) *krm.Rotation {
	if in == nil {
		return nil
	}
	out := &krm.Rotation{}
	out.NextRotationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextRotationTime())
	out.RotationPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRotationPeriod())
	return out
}
func Rotation_ToProto(mapCtx *direct.MapContext, in *krm.Rotation) *pb.Rotation {
	if in == nil {
		return nil
	}
	out := &pb.Rotation{}
	out.NextRotationTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRotationTime)
	out.RotationPeriod = direct.StringDuration_ToProto(mapCtx, in.RotationPeriod)
	return out
}
func SecretManagerSecretObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretManagerSecretObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretManagerSecretObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Topics
	// MISSING: Ttl
	// MISSING: Etag
	// MISSING: VersionDestroyTtl
	// MISSING: CustomerManagedEncryption
	return out
}
func SecretManagerSecretObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerSecretObservedState) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: Topics
	// MISSING: Ttl
	// MISSING: Etag
	// MISSING: VersionDestroyTtl
	// MISSING: CustomerManagedEncryption
	return out
}
func SecretManagerSecretSpec_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretManagerSecretSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecretManagerSecretSpec{}
	// MISSING: Name
	out.Replication = Replication_FromProto(mapCtx, in.GetReplication())
	// MISSING: CreateTime
	// MISSING: Topics
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Ttl
	// MISSING: Etag
	out.Rotation = Rotation_FromProto(mapCtx, in.GetRotation())
	for k, v := range in.VersionAliases {
		out.VersionAliases[k] = strconv.FormatInt(v, 10)
	}
	out.Annotations = in.Annotations
	// MISSING: Labels
	// MISSING: VersionDestroyTtl
	// MISSING: CustomerManagedEncryption
	return out
}
func SecretManagerSecretVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretManagerSecretVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretManagerSecretVersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestroyTime())
	// MISSING: State
	out.ReplicationStatus = ReplicationStatus_FromProto(mapCtx, in.GetReplicationStatus())
	// MISSING: Etag
	out.ClientSpecifiedPayloadChecksum = direct.LazyPtr(in.GetClientSpecifiedPayloadChecksum())
	out.ScheduledDestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduledDestroyTime())
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func SecretManagerSecretVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerSecretVersionObservedState) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.DestroyTime)
	// MISSING: State
	out.ReplicationStatus = ReplicationStatus_ToProto(mapCtx, in.ReplicationStatus)
	// MISSING: Etag
	out.ClientSpecifiedPayloadChecksum = direct.ValueOf(in.ClientSpecifiedPayloadChecksum)
	out.ScheduledDestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduledDestroyTime)
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}

func SecretManagerSecretVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretManagerSecretVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecretManagerSecretVersionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	// MISSING: ReplicationStatus
	// MISSING: Etag
	// MISSING: ClientSpecifiedPayloadChecksum
	// MISSING: ScheduledDestroyTime
	// MISSING: CustomerManagedEncryption
	return out
}
func SecretManagerSecretVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerSecretVersionSpec) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	// MISSING: ReplicationStatus
	// MISSING: Etag
	// MISSING: ClientSpecifiedPayloadChecksum
	// MISSING: ScheduledDestroyTime
	// MISSING: CustomerManagedEncryption
	return out
}

func SecretVersion_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretVersion {
	if in == nil {
		return nil
	}
	out := &krm.SecretVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestroyTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ReplicationStatus = ReplicationStatus_FromProto(mapCtx, in.GetReplicationStatus())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ClientSpecifiedPayloadChecksum = direct.LazyPtr(in.GetClientSpecifiedPayloadChecksum())
	out.ScheduledDestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduledDestroyTime())
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}
func SecretVersion_ToProto(mapCtx *direct.MapContext, in *krm.SecretVersion) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.DestroyTime)
	out.State = direct.Enum_ToProto[pb.SecretVersion_State](mapCtx, in.State)
	out.ReplicationStatus = ReplicationStatus_ToProto(mapCtx, in.ReplicationStatus)
	out.Etag = direct.ValueOf(in.Etag)
	out.ClientSpecifiedPayloadChecksum = direct.ValueOf(in.ClientSpecifiedPayloadChecksum)
	out.ScheduledDestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduledDestroyTime)
	out.CustomerManagedEncryption = CustomerManagedEncryptionStatus_ToProto(mapCtx, in.CustomerManagedEncryption)
	return out
}
