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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	krmmonitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	krmpubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AllUpdatesRule_FromProto(mapCtx *direct.MapContext, in *pb.AllUpdatesRule) *krm.AllUpdatesRule {
	if in == nil {
		return nil
	}
	out := &krm.AllUpdatesRule{}
	if in.GetPubsubTopic() != "" {
		out.PubsubTopicRef = &krmpubsubv1beta1.PubSubTopicRef{External: in.GetPubsubTopic()}
	}
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())

	if v := in.GetMonitoringNotificationChannels(); len(v) != 0 {
		for i := range v {
			out.MonitoringNotificationChannelRefs = append(out.MonitoringNotificationChannelRefs, &krmmonitoringv1beta1.NotificationChannelRef{External: v[i]})
		}
	}

	out.DisableDefaultIAMRecipients = direct.LazyPtr(in.GetDisableDefaultIamRecipients())
	// MISSING: EnableProjectLevelRecipients
	return out
}
func AllUpdatesRule_ToProto(mapCtx *direct.MapContext, in *krm.AllUpdatesRule) *pb.AllUpdatesRule {
	if in == nil {
		return nil
	}
	out := &pb.AllUpdatesRule{}
	if in.PubsubTopicRef != nil {
		out.PubsubTopic = in.PubsubTopicRef.External
	}
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)

	if v := in.MonitoringNotificationChannelRefs; len(v) != 0 {
		for i := range v {
			out.MonitoringNotificationChannels = append(out.MonitoringNotificationChannels, v[i].External)
		}
	}

	out.DisableDefaultIamRecipients = direct.ValueOf(in.DisableDefaultIAMRecipients)
	// MISSING: EnableProjectLevelRecipients
	return out
}
func BillingBudgetsBudgetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetsBudgetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetsBudgetObservedState{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetsBudgetObservedState) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	// MISSING: Name
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetsBudgetSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetsBudgetSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BudgetFilter = Filter_FromProto(mapCtx, in.GetBudgetFilter())
	out.Amount = BudgetAmount_FromProto(mapCtx, in.GetAmount())
	out.ThresholdRules = direct.Slice_FromProto(mapCtx, in.ThresholdRules, ThresholdRule_FromProto)
	out.AllUpdatesRule = AllUpdatesRule_FromProto(mapCtx, in.GetAllUpdatesRule())
	// MISSING: Etag
	return out
}
func BillingBudgetsBudgetSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetsBudgetSpec) *pb.Budget {
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

	if v := in.GetProjects(); len(v) != 0 {
		for i := range v {
			out.ProjectRefs = append(out.ProjectRefs, &refsv1beta1.ProjectRef{External: v[i]})
		}
	}

	// MISSING: ResourceAncestors
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

	if v := in.ProjectRefs; len(v) != 0 {
		for i := range v {
			out.Projects = append(out.Projects, v[i].External)
		}
	}

	// MISSING: ResourceAncestors
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
func Filter_CalendarPeriod_ToProto(mapCtx *direct.MapContext, in *string) *pb.Filter_CalendarPeriod {
	if in == nil {
		return nil
	}
	return &pb.Filter_CalendarPeriod{CalendarPeriod: direct.Enum_ToProto[pb.CalendarPeriod](mapCtx, in)}
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
