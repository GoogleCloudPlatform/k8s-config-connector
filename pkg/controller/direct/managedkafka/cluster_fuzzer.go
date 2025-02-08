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

// +tool:fuzz-gen
// proto.message: google.cloud.managedkafka.v1.Cluster
// krm.kind: ManagedKafkaCluster

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"k8s.io/apimachinery/pkg/util/sets"
)

func init() {
	fuzztesting.RegisterFuzzer(managedKafkaClusterSpecFuzzer().FuzzSpec)
	fuzztesting.RegisterFuzzer(managedKafkaClusterObservedStateFuzzer().FuzzObservedState)
}

var managedKafkaClusterKrmFields = fuzztesting.KRMFields{
	UnimplementedFields: sets.New(".name",
		".satisfies_pzi",
		".satisfies_pzs"),
	SpecFields: sets.New(".labels",
		".gcp_config",
		".capacity_config",
		".rebalance_config"),
	ObservedStateFields: sets.New(".create_time",
		".create_time",
		".update_time",
		".state"),
}

func managedKafkaClusterSpecFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		ManagedKafkaClusterSpec_FromProto, ManagedKafkaClusterSpec_ToProto,
	)
	f.KRMFields = managedKafkaClusterKrmFields
	return f
}

func managedKafkaClusterObservedStateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		ManagedKafkaClusterObservedState_FromProto, ManagedKafkaClusterObservedState_ToProto,
	)
	f.KRMFields = managedKafkaClusterKrmFields
	return f
}
