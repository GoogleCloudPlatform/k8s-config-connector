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
// proto.message: google.cloud.managedkafka.v1.Cluster
// api.group: managedkafka.cnrm.cloud.google.com

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(managedKafkaClusterFuzzer())
}

func managedKafkaClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		ManagedKafkaClusterSpec_FromProto, ManagedKafkaClusterSpec_ToProto,
		ManagedKafkaClusterObservedState_FromProto, ManagedKafkaClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")          // special field
	f.UnimplementedFields.Insert(".satisfies_pzi") // NOTYET
	f.UnimplementedFields.Insert(".satisfies_pzs") // NOTYET

	f.SpecFields.Insert(".gcp_config")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".capacity_config")
	f.SpecFields.Insert(".rebalance_config")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")

	return f
}
