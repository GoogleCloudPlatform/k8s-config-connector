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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"

	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SecretManagerSecretStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretManagerSecretObservedState {
	return &krm.SecretManagerSecretObservedState{}
}

func CustomerManagedEncryption_FromProto(mapCtx *direct.MapContext, in *pb.CustomerManagedEncryption) *krm.CustomerManagedEncryption {
	if in == nil {
		return nil
	}
	out := &krm.CustomerManagedEncryption{}
	if in.KmsKeyName != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		}
	}
	return out
}
func CustomerManagedEncryption_ToProto(mapCtx *direct.MapContext, in *krm.CustomerManagedEncryption) *pb.CustomerManagedEncryption {
	if in == nil {
		return nil
	}
	out := &pb.CustomerManagedEncryption{}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}
func Replication_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.Replication {
	if in == nil {
		return nil
	}
	out := &krm.Replication{}
	if in.GetAutomatic() != nil {
		out.LegacyAutomatic = Replication_Automatic_FromProto(mapCtx, in.GetAutomatic())
		out.LegacyAuto = direct.LazyPtr(true)
	}
	out.UserManaged = Replication_UserManaged_FromProto(mapCtx, in.GetUserManaged())
	return out
}

func Replication_ToProto(mapCtx *direct.MapContext, in *krm.Replication) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	if oneof := Replication_UserManaged_ToProto(mapCtx, in.UserManaged); oneof != nil {
		out.Replication = &pb.Replication_UserManaged_{UserManaged: oneof}
	}
	if oneof := Replication_Automatic_ToProto(mapCtx, in.LegacyAutomatic); oneof != nil {
		out.Replication = &pb.Replication_Automatic_{Automatic: oneof}
	}
	// fallback to legacy Auto field.
	if out.Replication == nil && *in.LegacyAuto {
		out.Replication = &pb.Replication_Automatic_{Automatic: &pb.Replication_Automatic{}}
	}
	return out
}

func SecretManagerSecretSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerSecretSpec) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	// MISSING: Name
	out.Replication = Replication_ToProto(mapCtx, in.Replication)
	out.Topics = []*pb.Topic{}
	for _, topicRef := range in.TopicRefs {
		topic := &krm.Topic{Name: direct.LazyPtr(topicRef.PubSubTopicRef.External)}
		out.Topics = append(out.Topics, Topic_ToProto(mapCtx, topic))
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
		out.Expiration = &pb.Secret_ExpireTime{ExpireTime: oneof}
	}
	// MISSING: Etag
	out.Rotation = Rotation_ToProto(mapCtx, in.Rotation)
	// MISSING: VersionAliases
	// out.VersionAliases = in.VersionAliases
	out.Annotations = in.Annotations
	// MISSING: VersionDestroyTtl
	// MISSING: CustomerManagedEncryption
	return out
}
