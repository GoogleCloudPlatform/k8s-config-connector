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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(managedKafkaConsumerGroupFuzzer())
}

func managedKafkaConsumerGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConsumerGroup{},
		ManagedKafkaConsumerGroupSpec_FromProto, ManagedKafkaConsumerGroupSpec_ToProto,
		ManagedKafkaConsumerGroupObservedState_FromProto, ManagedKafkaConsumerGroupObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field, managed by KCC infrastructure

	f.StatusFields.Insert(".topics")

	return f
}
