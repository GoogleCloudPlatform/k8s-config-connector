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
// proto.message: google.cloud.recaptchaenterprise.v1.FirewallPolicy
// api.group: recaptchaenterprise.cnrm.cloud.google.com

package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(recaptchaEnterpriseFirewallPolicyFuzzer())
}

func recaptchaEnterpriseFirewallPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallPolicy{},
		RecaptchaEnterpriseFirewallPolicySpec_FromProto, RecaptchaEnterpriseFirewallPolicySpec_ToProto,
		RecaptchaEnterpriseFirewallPolicyObservedState_FromProto, RecaptchaEnterpriseFirewallPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".path")
	f.SpecFields.Insert(".condition")
	f.SpecFields.Insert(".actions")

	f.UnimplementedFields.Insert(".name")

	return f
}
