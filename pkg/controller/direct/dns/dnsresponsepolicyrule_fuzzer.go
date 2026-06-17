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

package dns

import (
	api "google.golang.org/api/dns/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(dnsResponsePolicyRuleFuzzer())
}

func dnsResponsePolicyRuleFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.ResponsePolicyRule{},
		DNSResponsePolicyRuleSpec_FromAPI, DNSResponsePolicyRuleSpec_ToAPI,
		DNSResponsePolicyRuleStatus_FromAPI, DNSResponsePolicyRuleStatus_ToAPI,
	)

	f.SpecField(".Behavior")
	f.SpecField(".DnsName")
	f.SpecField(".LocalData")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")
	f.Ignore_JSONBookkeeping(".LocalData.ForceSendFields")
	f.Ignore_JSONBookkeeping(".LocalData.NullFields")
	f.Ignore_JSONBookkeeping(".LocalData.LocalDatas[]")

	f.Unimplemented_NotYetTriaged(".RuleName")
	f.Unimplemented_NotYetTriaged(".Kind")
	f.Unimplemented_NotYetTriaged(".LocalData.LocalDatas[].RoutingPolicy")
	f.Unimplemented_NotYetTriaged(".LocalData.LocalDatas[].SignatureRrdatas")
	f.Unimplemented_NotYetTriaged(".LocalData.LocalDatas[].Kind")

	return f
}
