// Copyright 2025 Google LLC
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

// +generated:mapper
// krm.group: billingbudgets.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.billing.budgets.v1beta1

package billingbudgets

import (
	pb "cloud.google.com/go/billing/budgets/apiv1beta1/budgetspb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AllUpdatesRule_FromProto(mapCtx *direct.MapContext, in *pb.AllUpdatesRule) *krmv1beta1.AllUpdatesRule {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.AllUpdatesRule{}
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.MonitoringNotificationChannels = AllUpdatesRule_MonitoringNotificationChannels_FromProto(mapCtx, in.MonitoringNotificationChannels)
	out.DisableDefaultIAMRecipients = direct.LazyPtr(in.GetDisableDefaultIamRecipients())
	// MISSING: EnableProjectLevelRecipients
	return out
}
func AllUpdatesRule_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.AllUpdatesRule) *pb.AllUpdatesRule {
	if in == nil {
		return nil
	}
	out := &pb.AllUpdatesRule{}
	out.PubsubTopic = AllUpdatesRule_PubsubTopic_ToProto(mapCtx, in.PubsubTopic)
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.MonitoringNotificationChannels = AllUpdatesRule_MonitoringNotificationChannels_ToProto(mapCtx, in.MonitoringNotificationChannels)
	out.DisableDefaultIamRecipients = direct.ValueOf(in.DisableDefaultIAMRecipients)
	// MISSING: EnableProjectLevelRecipients
	return out
}
func BillingBudgetsBudgetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krmv1beta1.BillingBudgetsBudgetObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BillingBudgetsBudgetObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BillingBudgetsBudgetObservedState) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krmv1beta1.BillingBudgetsBudgetSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BillingBudgetsBudgetSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BudgetFilter = Filter_FromProto(mapCtx, in.GetBudgetFilter())
	out.Amount = BudgetAmount_FromProto(mapCtx, in.GetAmount())
	out.ThresholdRules = direct.Slice_FromProto(mapCtx, in.ThresholdRules, ThresholdRule_FromProto)
	out.AllUpdatesRule = AllUpdatesRule_FromProto(mapCtx, in.GetAllUpdatesRule())
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BillingBudgetsBudgetSpec) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BudgetFilter = Filter_ToProto(mapCtx, in.BudgetFilter)
	out.Amount = BudgetAmount_ToProto(mapCtx, in.Amount)
	out.ThresholdRules = direct.Slice_ToProto(mapCtx, in.ThresholdRules, ThresholdRule_ToProto)
	out.AllUpdatesRule = AllUpdatesRule_ToProto(mapCtx, in.AllUpdatesRule)
	// MISSING: Etag
	return out
}
func BudgetAmount_FromProto(mapCtx *direct.MapContext, in *pb.BudgetAmount) *krmv1beta1.BudgetAmount {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BudgetAmount{}
	out.SpecifiedAmount = Money_FromProto(mapCtx, in.GetSpecifiedAmount())
	out.LastPeriodAmount = LastPeriodAmount_FromProto(mapCtx, in.GetLastPeriodAmount())
	return out
}
func BudgetAmount_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BudgetAmount) *pb.BudgetAmount {
	if in == nil {
		return nil
	}
	out := &pb.BudgetAmount{}
	if oneof := Money_ToProto(mapCtx, in.SpecifiedAmount); oneof != nil {
		out.BudgetAmount = &pb.BudgetAmount_SpecifiedAmount{SpecifiedAmount: oneof}
	}
	if oneof := LastPeriodAmount_ToProto(mapCtx, in.LastPeriodAmount); oneof != nil {
		out.BudgetAmount = &pb.BudgetAmount_LastPeriodAmount{LastPeriodAmount: oneof}
	}
	return out
}
func CustomPeriod_FromProto(mapCtx *direct.MapContext, in *pb.CustomPeriod) *krmv1beta1.CustomPeriod {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.CustomPeriod{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	return out
}
func CustomPeriod_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.CustomPeriod) *pb.CustomPeriod {
	if in == nil {
		return nil
	}
	out := &pb.CustomPeriod{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	return out
}
func Filter_FromProto(mapCtx *direct.MapContext, in *pb.Filter) *krmv1beta1.Filter {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Filter{}
	out.Projects = Filter_Projects_FromProto(mapCtx, in.Projects)
	out.ResourceAncestors = in.ResourceAncestors
	out.CreditTypes = in.CreditTypes
	out.CreditTypesTreatment = direct.Enum_FromProto(mapCtx, in.GetCreditTypesTreatment())
	out.Services = in.Services
	out.Subaccounts = in.Subaccounts
	// MISSING: Labels
	out.CalendarPeriod = direct.Enum_FromProto(mapCtx, in.GetCalendarPeriod())
	out.CustomPeriod = CustomPeriod_FromProto(mapCtx, in.GetCustomPeriod())
	return out
}
func Filter_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Filter) *pb.Filter {
	if in == nil {
		return nil
	}
	out := &pb.Filter{}
	out.Projects = Filter_Projects_ToProto(mapCtx, in.Projects)
	out.ResourceAncestors = in.ResourceAncestors
	out.CreditTypes = in.CreditTypes
	out.CreditTypesTreatment = direct.Enum_ToProto[pb.Filter_CreditTypesTreatment](mapCtx, in.CreditTypesTreatment)
	out.Services = in.Services
	out.Subaccounts = in.Subaccounts
	// MISSING: Labels
	if oneof := Filter_CalendarPeriod_ToProto(mapCtx, in.CalendarPeriod); oneof != nil {
		out.UsagePeriod = oneof
	}
	if oneof := CustomPeriod_ToProto(mapCtx, in.CustomPeriod); oneof != nil {
		out.UsagePeriod = &pb.Filter_CustomPeriod{CustomPeriod: oneof}
	}
	return out
}
func LastPeriodAmount_FromProto(mapCtx *direct.MapContext, in *pb.LastPeriodAmount) *krmv1beta1.LastPeriodAmount {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.LastPeriodAmount{}
	return out
}
func LastPeriodAmount_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.LastPeriodAmount) *pb.LastPeriodAmount {
	if in == nil {
		return nil
	}
	out := &pb.LastPeriodAmount{}
	return out
}
func ThresholdRule_FromProto(mapCtx *direct.MapContext, in *pb.ThresholdRule) *krmv1beta1.ThresholdRule {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.ThresholdRule{}
	out.ThresholdPercent = direct.LazyPtr(in.GetThresholdPercent())
	out.SpendBasis = direct.Enum_FromProto(mapCtx, in.GetSpendBasis())
	return out
}
func ThresholdRule_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.ThresholdRule) *pb.ThresholdRule {
	if in == nil {
		return nil
	}
	out := &pb.ThresholdRule{}
	out.ThresholdPercent = direct.ValueOf(in.ThresholdPercent)
	out.SpendBasis = direct.Enum_ToProto[pb.ThresholdRule_Basis](mapCtx, in.SpendBasis)
	return out
}
