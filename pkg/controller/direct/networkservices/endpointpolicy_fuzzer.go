// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.networkservices.v1.EndpointPolicy
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networkServicesEndpointPolicyFuzzer())
}

func networkServicesEndpointPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EndpointPolicy{},
		NetworkServicesEndpointPolicySpec_FromProto, NetworkServicesEndpointPolicySpec_ToProto,
		NetworkServicesEndpointPolicyObservedState_FromProto, NetworkServicesEndpointPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".authorization_policy")
	f.SpecFields.Insert(".endpoint_matcher")
	f.SpecFields.Insert(".traffic_port_selector")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".server_tls_policy")
	f.SpecFields.Insert(".client_tls_policy")

	// Status fields are not mapped in KRM observed state as it's empty
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")

	return f
}
