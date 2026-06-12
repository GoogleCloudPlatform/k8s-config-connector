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

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	krmmanagedkafkav1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedkafka/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// Shim functions to allow mapper.generated.go to compile since ManagedKafkaConsumerGroup
// is implemented in an isolated sub-package.

func Topics_FromProto(mapCtx *direct.MapContext, in map[string]*pb.ConsumerTopicMetadata) map[string]*krmmanagedkafkav1alpha1.ConsumerTopicMetadata {
	return nil
}

func Topics_ToProto(mapCtx *direct.MapContext, in map[string]*krmmanagedkafkav1alpha1.ConsumerTopicMetadata) map[string]*pb.ConsumerTopicMetadata {
	return nil
}
