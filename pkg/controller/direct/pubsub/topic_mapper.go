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

// +generated:mapper
// krm.group: pubsub.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.pubsub.v1

package pubsub

import (
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PubSubTopicSpec_FromProto(mapCtx *direct.MapContext, in *pb.Topic) *krm.PubSubTopicSpec {
	if in == nil {
		return nil
	}
	out := &krm.PubSubTopicSpec{}

	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &krm.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	out.MessageRetentionDuration = direct.StringDuration_FromProto(mapCtx, in.GetMessageRetentionDuration())
	out.MessageStoragePolicy = MessageStoragePolicy_FromProto(mapCtx, in.GetMessageStoragePolicy())
	out.SchemaSettings = SchemaSettings_FromProto(mapCtx, in.GetSchemaSettings())

	return out
}

func PubSubTopicSpec_ToProto(mapCtx *direct.MapContext, in *krm.PubSubTopicSpec) *pb.Topic {
	if in == nil {
		return nil
	}
	out := &pb.Topic{}

	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	out.MessageRetentionDuration = direct.StringDuration_ToProto(mapCtx, in.MessageRetentionDuration)
	out.MessageStoragePolicy = MessageStoragePolicy_ToProto(mapCtx, in.MessageStoragePolicy)
	out.SchemaSettings = SchemaSettings_ToProto(mapCtx, in.SchemaSettings)

	return out
}

func MessageStoragePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MessageStoragePolicy) *krm.MessageStoragePolicy {
	if in == nil {
		return nil
	}
	out := &krm.MessageStoragePolicy{}
	out.AllowedPersistenceRegions = in.AllowedPersistenceRegions
	return out
}

func MessageStoragePolicy_ToProto(mapCtx *direct.MapContext, in *krm.MessageStoragePolicy) *pb.MessageStoragePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MessageStoragePolicy{}
	out.AllowedPersistenceRegions = in.AllowedPersistenceRegions
	return out
}

func SchemaSettings_FromProto(mapCtx *direct.MapContext, in *pb.SchemaSettings) *krm.SchemaSettings {
	if in == nil {
		return nil
	}
	out := &krm.SchemaSettings{}
	if in.GetEncoding() != pb.Encoding_ENCODING_UNSPECIFIED {
		encodingStr := in.GetEncoding().String()
		out.Encoding = &encodingStr
	}
	if in.GetSchema() != "" {
		out.SchemaRef = &krm.PubSubSchemaRef{External: in.GetSchema()}
	}
	return out
}

func SchemaSettings_ToProto(mapCtx *direct.MapContext, in *krm.SchemaSettings) *pb.SchemaSettings {
	if in == nil {
		return nil
	}
	out := &pb.SchemaSettings{}
	if in.Encoding != nil {
		if val, ok := pb.Encoding_value[*in.Encoding]; ok {
			out.Encoding = pb.Encoding(val)
		} else {
			mapCtx.Errorf("unknown encoding value %q", *in.Encoding)
		}
	}
	if in.SchemaRef != nil {
		out.Schema = in.SchemaRef.External
	}
	return out
}
