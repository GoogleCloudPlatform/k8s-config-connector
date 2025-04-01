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

// +tool:fuzz-gen
// proto.message: google.cloud.managedkafka.v1.ConsumerGroup
// api.group: managedkafka.cnrm.cloud.google.com

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// Placeholder functions until the real ones are generated or implemented
// Based on the KRM struct definitions and proto, there are no fields to map between Spec/Status and the Proto.
func ManagedKafkaConsumerGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConsumerGroup) *krm.ManagedKafkaConsumerGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaConsumerGroupSpec{}
	// No fields in KRM Spec map directly from the proto message. 'topics' is effectively read-only/status.
	return out
}

func ManagedKafkaConsumerGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaConsumerGroupSpec) *pb.ConsumerGroup {
	if in == nil {
		return nil
	}
	out := &pb.ConsumerGroup{}
	// No fields in KRM Spec map directly to the proto message.
	return out
}

func ManagedKafkaConsumerGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConsumerGroup) *krm.ManagedKafkaConsumerGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedKafkaConsumerGroupObservedState{}
	out.Topics = make(map[string]*krm.ConsumerTopicMetadata)
	for k, v := range in.GetTopics() {
		out.Topics[k] = ConsumerTopicMetadata_FromProto(mapCtx, v)
	}
	return out
}

func ManagedKafkaConsumerGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaConsumerGroupObservedState) *pb.ConsumerGroup {
	if in == nil {
		return nil
	}
	out := &pb.ConsumerGroup{}
	out.Topics = make(map[string]*pb.ConsumerTopicMetadata)
	for k, v := range in.Topics {
		out.Topics[k] = ConsumerTopicMetadata_ToProto(mapCtx, v)
	}
	return out
}

func ConsumerTopicMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ConsumerTopicMetadata) *krm.ConsumerTopicMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ConsumerTopicMetadata{}
	return out
}

func ConsumerTopicMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ConsumerTopicMetadata) *pb.ConsumerTopicMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ConsumerTopicMetadata{}
	return out
}
