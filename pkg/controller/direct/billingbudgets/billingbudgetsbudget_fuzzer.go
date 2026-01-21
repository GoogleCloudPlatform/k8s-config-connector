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
// proto.message: google.cloud.billing.budgets.v1beta1.Budget
// crd.group: billingbudgets.cnrm.cloud.google.com
// crd.kind: BillingBudgetsBudget

package billingbudgets

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	pb "cloud.google.com/go/billing/budgets/apiv1beta1/budgetspb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(billingBudgetsBudgetFuzzer())
}

func billingBudgetsBudgetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Budget{},
		BillingBudgetsBudgetSpec_FromProto, BillingBudgetsBudgetSpec_ToProto,
		BillingBudgetsBudgetObservedState_FromProto, BillingBudgetsBudgetObservedState_ToProto,
	)

	f.Unimplemented_Etag()

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".budget_filter")
	f.SpecField(".amount")
	f.SpecField(".threshold_rules")
	f.SpecField(".all_updates_rule")

	f.Unimplemented_NotYetTriaged(".budget_filter.calendar_period")
	f.Unimplemented_NotYetTriaged(".budget_filter.labels")
	f.Unimplemented_NotYetTriaged(".budget_filter.resource_ancestors")
	f.Unimplemented_NotYetTriaged(".all_updates_rule.enable_project_level_recipients")

	return f
}
