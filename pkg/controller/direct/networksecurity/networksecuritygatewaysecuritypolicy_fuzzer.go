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
// proto.message: google.cloud.networksecurity.v1.GatewaySecurityPolicy
// api.group: networksecurity.cnrm.cloud.google.com

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(gatewaySecurityPolicyFuzzer())
}

func gatewaySecurityPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.GatewaySecurityPolicy{},
		NetworkSecurityGatewaySecurityPolicySpec_FromProto, NetworkSecurityGatewaySecurityPolicySpec_ToProto,
		NetworkSecurityGatewaySecurityPolicyObservedState_FromProto, NetworkSecurityGatewaySecurityPolicyObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".tls_inspection_policy")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")

	return f
}
