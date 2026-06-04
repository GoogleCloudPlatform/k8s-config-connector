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
// proto.message: google.cloud.billing.budgets.v1.Budget
// api.group: billingbudgets.cnrm.cloud.google.com

package billingbudgets

import (
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	gdate "google.golang.org/genproto/googleapis/type/date"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(billingbudgetsBudgetFuzzer())
}

func billingbudgetsBudgetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Budget{},
		BillingBudgetsBudgetSpec_v1beta1_FromProto, BillingBudgetsBudgetSpec_v1beta1_ToProto,
		BillingBudgetsBudgetStatus_FromProto, BillingBudgetsBudgetStatus_ToProto,
	)

	f.FilterSpec = func(in *pb.Budget) {
		if in.BudgetFilter != nil {
			if cpOneof, ok := in.BudgetFilter.UsagePeriod.(*pb.Filter_CustomPeriod); ok && cpOneof != nil && cpOneof.CustomPeriod != nil {
				cp := cpOneof.CustomPeriod
				if cp.StartDate == nil {
					cp.StartDate = &gdate.Date{}
				}
			}
			if cpOneof, ok := in.BudgetFilter.UsagePeriod.(*pb.Filter_CalendarPeriod); ok && cpOneof != nil {
				if cpOneof.CalendarPeriod == pb.CalendarPeriod_CALENDAR_PERIOD_UNSPECIFIED {
					in.BudgetFilter.UsagePeriod = nil
				}
			}
			if in.BudgetFilter.Labels != nil {
				for _, v := range in.BudgetFilter.Labels {
					if v == nil {
						continue
					}
					for i, val := range v.Values {
						if val == nil {
							v.Values[i] = structpb.NewStringValue("")
						} else if _, ok := val.GetKind().(*structpb.Value_StringValue); !ok {
							v.Values[i] = structpb.NewStringValue(val.GetStringValue())
						}
					}
				}
			}
		}
		if in.Amount != nil && in.Amount.BudgetAmount == nil {
			in.Amount = nil
		}
	}

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".budget_filter")
	f.SpecField(".amount")
	f.SpecField(".threshold_rules")
	f.SpecField(".notifications_rule")

	f.StatusField(".etag")

	// Unimplemented fields
	f.Unimplemented_NotYetTriaged(".budget_filter.resource_ancestors")
	f.Unimplemented_NotYetTriaged(".notifications_rule.enable_project_level_recipients")

	return f
}
