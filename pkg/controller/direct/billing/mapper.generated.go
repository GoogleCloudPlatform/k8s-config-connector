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

package billing

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BillingBudgetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func BillingBudgetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetObservedState) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func BillingBudgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func BillingBudgetSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetSpec) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func Budget_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.Budget {
	if in == nil {
		return nil
	}
	out := &krm.Budget{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BudgetFilter = Filter_FromProto(mapCtx, in.GetBudgetFilter())
	out.Amount = BudgetAmount_FromProto(mapCtx, in.GetAmount())
	out.ThresholdRules = direct.Slice_FromProto(mapCtx, in.ThresholdRules, ThresholdRule_FromProto)
	out.NotificationsRule = NotificationsRule_FromProto(mapCtx, in.GetNotificationsRule())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Budget_ToProto(mapCtx *direct.MapContext, in *krm.Budget) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BudgetFilter = Filter_ToProto(mapCtx, in.BudgetFilter)
	out.Amount = BudgetAmount_ToProto(mapCtx, in.Amount)
	out.ThresholdRules = direct.Slice_ToProto(mapCtx, in.ThresholdRules, ThresholdRule_ToProto)
	out.NotificationsRule = NotificationsRule_ToProto(mapCtx, in.NotificationsRule)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func BudgetAmount_FromProto(mapCtx *direct.MapContext, in *pb.BudgetAmount) *krm.BudgetAmount {
	if in == nil {
		return nil
	}
	out := &krm.BudgetAmount{}
	out.SpecifiedAmount = Money_FromProto(mapCtx, in.GetSpecifiedAmount())
	out.LastPeriodAmount = LastPeriodAmount_FromProto(mapCtx, in.GetLastPeriodAmount())
	return out
}
func BudgetAmount_ToProto(mapCtx *direct.MapContext, in *krm.BudgetAmount) *pb.BudgetAmount {
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
func BudgetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BudgetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BudgetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func BudgetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BudgetObservedState) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: BudgetFilter
	// MISSING: Amount
	// MISSING: ThresholdRules
	// MISSING: NotificationsRule
	// MISSING: Etag
	return out
}
func CustomPeriod_FromProto(mapCtx *direct.MapContext, in *pb.CustomPeriod) *krm.CustomPeriod {
	if in == nil {
		return nil
	}
	out := &krm.CustomPeriod{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	return out
}
func CustomPeriod_ToProto(mapCtx *direct.MapContext, in *krm.CustomPeriod) *pb.CustomPeriod {
	if in == nil {
		return nil
	}
	out := &pb.CustomPeriod{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	return out
}
func Filter_FromProto(mapCtx *direct.MapContext, in *pb.Filter) *krm.Filter {
	if in == nil {
		return nil
	}
	out := &krm.Filter{}
	out.Projects = in.Projects
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
func Filter_ToProto(mapCtx *direct.MapContext, in *krm.Filter) *pb.Filter {
	if in == nil {
		return nil
	}
	out := &pb.Filter{}
	out.Projects = in.Projects
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
func LastPeriodAmount_FromProto(mapCtx *direct.MapContext, in *pb.LastPeriodAmount) *krm.LastPeriodAmount {
	if in == nil {
		return nil
	}
	out := &krm.LastPeriodAmount{}
	return out
}
func LastPeriodAmount_ToProto(mapCtx *direct.MapContext, in *krm.LastPeriodAmount) *pb.LastPeriodAmount {
	if in == nil {
		return nil
	}
	out := &pb.LastPeriodAmount{}
	return out
}
func NotificationsRule_FromProto(mapCtx *direct.MapContext, in *pb.NotificationsRule) *krm.NotificationsRule {
	if in == nil {
		return nil
	}
	out := &krm.NotificationsRule{}
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.MonitoringNotificationChannels = in.MonitoringNotificationChannels
	out.DisableDefaultIamRecipients = direct.LazyPtr(in.GetDisableDefaultIamRecipients())
	out.EnableProjectLevelRecipients = direct.LazyPtr(in.GetEnableProjectLevelRecipients())
	return out
}
func NotificationsRule_ToProto(mapCtx *direct.MapContext, in *krm.NotificationsRule) *pb.NotificationsRule {
	if in == nil {
		return nil
	}
	out := &pb.NotificationsRule{}
	out.PubsubTopic = direct.ValueOf(in.PubsubTopic)
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.MonitoringNotificationChannels = in.MonitoringNotificationChannels
	out.DisableDefaultIamRecipients = direct.ValueOf(in.DisableDefaultIamRecipients)
	out.EnableProjectLevelRecipients = direct.ValueOf(in.EnableProjectLevelRecipients)
	return out
}
func ThresholdRule_FromProto(mapCtx *direct.MapContext, in *pb.ThresholdRule) *krm.ThresholdRule {
	if in == nil {
		return nil
	}
	out := &krm.ThresholdRule{}
	out.ThresholdPercent = direct.LazyPtr(in.GetThresholdPercent())
	out.SpendBasis = direct.Enum_FromProto(mapCtx, in.GetSpendBasis())
	return out
}
func ThresholdRule_ToProto(mapCtx *direct.MapContext, in *krm.ThresholdRule) *pb.ThresholdRule {
	if in == nil {
		return nil
	}
	out := &pb.ThresholdRule{}
	out.ThresholdPercent = direct.ValueOf(in.ThresholdPercent)
	out.SpendBasis = direct.Enum_ToProto[pb.ThresholdRule_Basis](mapCtx, in.SpendBasis)
	return out
}
