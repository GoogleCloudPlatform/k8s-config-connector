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

// +tool:fuzz-gen
// proto.message: google.cloud.managedkafka.v1.ConnectCluster
// api.group: managedkafka.cnrm.cloud.google.com

package managedkafka

import (
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(managedKafkaConnectClusterFuzzer())
}

func managedKafkaConnectClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConnectCluster{},
		ManagedKafkaConnectClusterSpec_v1alpha1_FromProto, ManagedKafkaConnectClusterSpec_v1alpha1_ToProto,
		ManagedKafkaConnectClusterObservedState_v1alpha1_FromProto, ManagedKafkaConnectClusterObservedState_v1alpha1_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".gcp_config")
	f.SpecField(".kafka_cluster")
	f.SpecField(".labels")
	f.SpecField(".capacity_config")
	f.SpecField(".config")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")

	return f
}
