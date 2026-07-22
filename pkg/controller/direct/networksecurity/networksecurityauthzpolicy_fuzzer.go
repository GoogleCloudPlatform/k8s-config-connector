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
// proto.message: google.cloud.networksecurity.v1.AuthzPolicy
// api.group: networksecurity.cnrm.cloud.google.com

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(authzPolicyFuzzer())
}

func authzPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AuthzPolicy{},
		NetworkSecurityAuthzPolicySpec_v1alpha1_FromProto, NetworkSecurityAuthzPolicySpec_v1alpha1_ToProto,
		NetworkSecurityAuthzPolicyObservedState_v1alpha1_FromProto, NetworkSecurityAuthzPolicyObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".target")
	f.SpecField(".http_rules")
	f.SpecField(".network_rules")
	f.SpecField(".action")
	f.SpecField(".custom_provider")
	f.SpecField(".policy_profile")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".labels")

	return f
}
