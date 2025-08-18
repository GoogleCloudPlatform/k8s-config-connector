// Copyright 2025 Google LLC
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
// proto.message: google.cloud.billing.v1.BillingAccount
// api.group: billing.cnrm.cloud.google.com

package billing

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(billingBillingAccountFuzzer())
}

func billingBillingAccountFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BillingAccount{},
		BillingAccountSpec_FromProto, BillingAccountSpec_ToProto,
		BillingAccountObservedState_FromProto, BillingAccountObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".parent")
	f.SpecFields.Insert(".currency_code")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".open")
	f.StatusFields.Insert(".parent")
	f.StatusFields.Insert(".master_billing_account")

	return f
}
