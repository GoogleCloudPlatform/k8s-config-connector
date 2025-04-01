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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(managedKafkaConsumerGroupFuzzer())
}

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
	// KRM ObservedState is empty. The 'topics' field from proto is not mapped here.
	return out
}

func ManagedKafkaConsumerGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedKafkaConsumerGroupObservedState) *pb.ConsumerGroup {
	if in == nil {
		return nil
	}
	// KRM ObservedState is empty, so nothing maps back to the proto.
	return &pb.ConsumerGroup{}
}

func managedKafkaConsumerGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConsumerGroup{},
		ManagedKafkaConsumerGroupSpec_FromProto, ManagedKafkaConsumerGroupSpec_ToProto,
		ManagedKafkaConsumerGroupObservedState_FromProto, ManagedKafkaConsumerGroupObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")   // special field, managed by KCC infrastructure
	f.UnimplementedFields.Insert(".topics") // Optional field in proto, but not present in KRM spec or status

	// No SpecFields mapped from proto
	// No StatusFields mapped from proto

	return f
}
