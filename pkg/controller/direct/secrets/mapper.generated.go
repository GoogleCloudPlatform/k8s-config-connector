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

package secrets

import (
	pb "cloud.google.com/go/secrets/apiv1beta1/secretspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secrets/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Replication_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.Replication {
	if in == nil {
		return nil
	}
	out := &krm.Replication{}
	out.Automatic = Replication_Automatic_FromProto(mapCtx, in.GetAutomatic())
	out.UserManaged = Replication_UserManaged_FromProto(mapCtx, in.GetUserManaged())
	return out
}
func Replication_ToProto(mapCtx *direct.MapContext, in *krm.Replication) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	if oneof := Replication_Automatic_ToProto(mapCtx, in.Automatic); oneof != nil {
		out.Replication = &pb.Replication_Automatic_{Automatic: oneof}
	}
	if oneof := Replication_UserManaged_ToProto(mapCtx, in.UserManaged); oneof != nil {
		out.Replication = &pb.Replication_UserManaged_{UserManaged: oneof}
	}
	return out
}
func Replication_Automatic_FromProto(mapCtx *direct.MapContext, in *pb.Replication_Automatic) *krm.Replication_Automatic {
	if in == nil {
		return nil
	}
	out := &krm.Replication_Automatic{}
	return out
}
func Replication_Automatic_ToProto(mapCtx *direct.MapContext, in *krm.Replication_Automatic) *pb.Replication_Automatic {
	if in == nil {
		return nil
	}
	out := &pb.Replication_Automatic{}
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
	return out
}
func Replication_UserManaged_Replica_ToProto(mapCtx *direct.MapContext, in *krm.Replication_UserManaged_Replica) *pb.Replication_UserManaged_Replica {
	if in == nil {
		return nil
	}
	out := &pb.Replication_UserManaged_Replica{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	// MISSING: Name
	out.Replication = Replication_FromProto(mapCtx, in.GetReplication())
	// MISSING: CreateTime
	out.Labels = in.Labels
	return out
}
func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	// MISSING: Name
	out.Replication = Replication_ToProto(mapCtx, in.Replication)
	// MISSING: CreateTime
	out.Labels = in.Labels
	return out
}
func SecretObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Replication
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	return out
}
func SecretObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretObservedState) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Replication
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	return out
}
func SecretsSecretObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretsSecretObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretsSecretObservedState{}
	// MISSING: Name
	// MISSING: Replication
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func SecretsSecretObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretsSecretObservedState) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	// MISSING: Name
	// MISSING: Replication
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func SecretsSecretSpec_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretsSecretSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecretsSecretSpec{}
	// MISSING: Name
	// MISSING: Replication
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func SecretsSecretSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecretsSecretSpec) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	// MISSING: Name
	// MISSING: Replication
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
