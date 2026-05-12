// Copyright 2026 Google LLC
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
        refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1alpha1"
        krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func SecretManagerRegionalSecretStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretManagerRegionalSecretObservedState {
        if in == nil {
                return nil
        }
        out := &krm.SecretManagerRegionalSecretObservedState{}
        out.VersionAliases = MapStringInt64_ToMapStringString(mapCtx, in.VersionAliases)
        out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
        return out
}

func RegionalCustomerManagedEncryption_ToProto(mapCtx *direct.MapContext, in *krm.CustomerManagedEncryption) *pb.CustomerManagedEncryption {
        if in == nil {
                return nil
        }
        out := &pb.CustomerManagedEncryption{}
        if in.KmsKeyRef != nil {
                out.KmsKeyName = in.KmsKeyRef.External
        }
        return out
}

func RegionalRotation_ToProto(mapCtx *direct.MapContext, in *krm.Rotation) *pb.Rotation {
        if in == nil {
                return nil
        }
        out := &pb.Rotation{}
        out.NextRotationTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRotationTime)
        out.RotationPeriod = direct.StringDuration_ToProto(mapCtx, in.RotationPeriod)
        return out
}

func SecretManagerRegionalSecretSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerRegionalSecretSpec) *pb.Secret {
        if in == nil {
                return nil
        }
        out := &pb.Secret{}
        
        out.Topics = []*pb.Topic{}
        for _, topicRef := range in.TopicRefs {
                topic := &krmv1beta1.Topic{Name: direct.LazyPtr(topicRef.PubSubTopicRef.External)}
                out.Topics = append(out.Topics, Topic_ToProto(mapCtx, topic))
        }
        if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime); oneof != nil {
                out.Expiration = &pb.Secret_ExpireTime{ExpireTime: oneof}
        }
        if oneof := direct.Duration_ToProto(mapCtx, in.TTL); oneof != nil {
                out.Expiration = &pb.Secret_Ttl{Ttl: oneof}
        }
        
        out.Rotation = RegionalRotation_ToProto(mapCtx, in.Rotation)
        out.VersionAliases = MapStringString_ToMapStringInt64(mapCtx, in.VersionAliases)
        out.Annotations = in.Annotations
        out.CustomerManagedEncryption = RegionalCustomerManagedEncryption_ToProto(mapCtx, in.CustomerManagedEncryption)
        return out
}

func SecretManagerRegionalSecretSpec_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.SecretManagerRegionalSecretSpec {
        if in == nil {
                return nil
        }
        out := &krm.SecretManagerRegionalSecretSpec{}
        
        // Topics are handled in export mapping if needed.
        out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
        // Rotation mapping
        if in.Rotation != nil {
                out.Rotation = &krm.Rotation{
                        NextRotationTime: direct.StringTimestamp_FromProto(mapCtx, in.Rotation.NextRotationTime),
                        RotationPeriod:   direct.StringDuration_FromProto(mapCtx, in.Rotation.RotationPeriod),
                }
        }
        out.VersionAliases = MapStringInt64_ToMapStringString(mapCtx, in.VersionAliases)
        out.Annotations = in.Annotations
        
        if in.CustomerManagedEncryption != nil {
                out.CustomerManagedEncryption = &krm.CustomerManagedEncryption{}
                out.CustomerManagedEncryption.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.CustomerManagedEncryption.KmsKeyName}
        }
        return out
}
