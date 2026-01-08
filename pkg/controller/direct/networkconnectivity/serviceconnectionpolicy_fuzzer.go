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

package networkconnectivity

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(serviceConnectionPolicyFuzzer())
}

func serviceConnectionPolicyFuzzer() fuzztesting.KRMFuzzer {
	fuzzer := fuzztesting.NewKRMTypedFuzzer(
		&pb.ServiceConnectionPolicy{},
		NetworkConnectivityServiceConnectionPolicySpec_FromProto,
		NetworkConnectivityServiceConnectionPolicySpec_ToProto,
		NetworkConnectivityServiceConnectionPolicyObservedState_FromProto,
		NetworkConnectivityServiceConnectionPolicyObservedState_ToProto,
	)

	fuzzer.UnimplementedFields.Insert(".name")
	fuzzer.UnimplementedFields.Insert(".auto_created_subnet_info")
	fuzzer.UnimplementedFields.Insert(".labels")
	fuzzer.UnimplementedFields.Insert(".psc_config.allowed_google_producers_resource_hierarchy_level")
	fuzzer.UnimplementedFields.Insert(".psc_connections[].ip_version")
	fuzzer.UnimplementedFields.Insert(".psc_connections[].producer_instance_metadata")
	fuzzer.UnimplementedFields.Insert(".psc_connections[].service_class")

	fuzzer.SpecFields.Insert(".description")
	fuzzer.SpecFields.Insert(".network")
	fuzzer.SpecFields.Insert(".service_class")
	fuzzer.SpecFields.Insert(".psc_config")
	fuzzer.SpecFields.Insert(".psc_connections[].error.details")

	fuzzer.StatusFields.Insert(".create_time")
	fuzzer.StatusFields.Insert(".update_time")
	fuzzer.StatusFields.Insert(".infrastructure")
	fuzzer.StatusFields.Insert(".psc_connections")
	fuzzer.StatusFields.Insert(".etag")

	return fuzzer
}
