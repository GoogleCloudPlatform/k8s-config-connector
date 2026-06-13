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
	fuzztesting.RegisterKRMFuzzer_NoProto(dnsResponsePolicyFuzzer())
}

func dnsResponsePolicyFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.ResponsePolicy{},
		DNSResponsePolicySpec_FromAPI, DNSResponsePolicySpec_ToAPI,
		DNSResponsePolicyStatus_FromAPI, DNSResponsePolicyStatus_ToAPI,
	)

	f.SpecField(".Description")
	f.SpecField(".GkeClusters")
	f.SpecField(".Networks")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")
	f.Ignore_JSONBookkeeping(".GkeClusters[]")
	f.Ignore_JSONBookkeeping(".Networks[]")

	f.Unimplemented_NotYetTriaged(".Id")
	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Kind")
	f.Unimplemented_NotYetTriaged(".ResponsePolicyName")
	f.Unimplemented_NotYetTriaged(".GkeClusters[].Kind")
	f.Unimplemented_NotYetTriaged(".Networks[].Kind")

	return f
}
