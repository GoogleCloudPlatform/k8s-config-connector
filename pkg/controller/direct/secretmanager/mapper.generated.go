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
	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Replication = Replication_FromProto(mapCtx, in.GetReplication())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Labels = in.Labels
	out.Topics = direct.Slice_FromProto(mapCtx, in.Topics, Topic_FromProto)
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Rotation = Rotation_FromProto(mapCtx, in.GetRotation())
	out.VersionAliases = in.VersionAliases
	out.Annotations = in.Annotations
	// MISSING: VersionDestroyTtl
	out.CustomerManagedEncryption = CustomerManagedEncryption_FromProto(mapCtx, in.GetCustomerManagedEncryption())
	return out
}

func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	out.Name = direct.ValueOf(in.Name)
	out.Replication = Replication_ToProto(mapCtx, in.Replication)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Labels = in.Labels
	out.Topics = direct.Slice_ToProto(mapCtx, in.Topics, Topic_ToProto)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.Secret_ExpireTime{ExpireTime: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.Ttl); oneof != nil {
		out.Expiration = &pb.Secret_Ttl{Ttl: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.Rotation = Rotation_ToProto(mapCtx, in.Rotation)
	out.VersionAliases = in.VersionAliases
	out.Annotations = in.Annotations
	// MISSING: VersionDestroyTtl
	out.CustomerManagedEncryption = CustomerManagedEncryption_ToProto(mapCtx, in.CustomerManagedEncryption)
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
	// MISSING: Labels
	// MISSING: Topics
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	// MISSING: Ttl
	// MISSING: Etag
	out.Rotation = Rotation_FromProto(mapCtx, in.GetRotation())
	// MISSING: VersionAliases
	// out.VersionAliases = in.VersionAliases
	out.Annotations = in.Annotations
	// MISSING: VersionDestroyTtl
	// MISSING: CustomerManagedEncryption
	return out
}
func Topic_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.Topic {
	if in == nil {
		return nil
	}
	out := &krm.Topic{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}

func Topic_ToProto(mapCtx *direct.MapContext, in *krm.Topic) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
