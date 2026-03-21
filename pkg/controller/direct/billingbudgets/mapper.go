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

package billingbudgets

import (
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datepb "google.golang.org/genproto/googleapis/type/date"
	moneypb "google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/types/known/structpb"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func BillingBudgetsBudgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetsBudgetSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetsBudgetSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BudgetFilter = BudgetFilter_FromProto(mapCtx, in.GetBudgetFilter())
	out.Amount = BudgetAmount_FromProto(mapCtx, in.GetAmount())
	out.ThresholdRules = direct.Slice_FromProto(mapCtx, in.ThresholdRules, BudgetThresholdRule_FromProto)
	out.AllUpdatesRule = BudgetNotificationsRule_FromProto(mapCtx, in.GetNotificationsRule())
	return out
}

func BillingBudgetsBudgetSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetsBudgetSpec) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	if in.BillingAccountRef.Name != "" {
		mapCtx.Errorf("BillingAccountRef: Name references are not yet supported, use External instead")
		return nil
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BudgetFilter = BudgetFilter_ToProto(mapCtx, in.BudgetFilter)
	out.Amount = BudgetAmount_ToProto(mapCtx, in.Amount)
	out.ThresholdRules = direct.Slice_ToProto(mapCtx, in.ThresholdRules, BudgetThresholdRule_ToProto)
	out.NotificationsRule = BudgetNotificationsRule_ToProto(mapCtx, in.AllUpdatesRule)
	return out
}

func BillingBudgetsBudgetStatus_FromProto(mapCtx *direct.MapContext, in *pb.Budget) *krm.BillingBudgetsBudgetStatus {
	if in == nil {
		return nil
	}
	out := &krm.BillingBudgetsBudgetStatus{}
	if in.GetEtag() != "" {
		out.Etag = direct.LazyPtr(in.GetEtag())
	}
	return out
}

func BillingBudgetsBudgetStatus_ToProto(mapCtx *direct.MapContext, in *krm.BillingBudgetsBudgetStatus) *pb.Budget {
	if in == nil {
		return nil
	}
	out := &pb.Budget{}
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func BudgetNotificationsRule_FromProto(mapCtx *direct.MapContext, in *pb.NotificationsRule) *krm.BudgetNotificationsRule {
	if in == nil {
		return nil
	}
	out := &krm.BudgetNotificationsRule{}
	if in.GetPubsubTopic() != "" {
		out.PubsubTopicRef = &krm.PubSubTopicRef{External: in.GetPubsubTopic()}
	}
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.MonitoringNotificationChannels = BudgetNotificationsRule_MonitoringNotificationChannels_FromProto(mapCtx, in.MonitoringNotificationChannels)
	out.DisableDefaultIAMRecipients = direct.LazyPtr(in.GetDisableDefaultIamRecipients())
	return out
}

func BudgetNotificationsRule_ToProto(mapCtx *direct.MapContext, in *krm.BudgetNotificationsRule) *pb.NotificationsRule {
	if in == nil {
		return nil
	}
	out := &pb.NotificationsRule{}
	if in.PubsubTopicRef != nil {
		if in.PubsubTopicRef.Name != "" {
			mapCtx.Errorf("PubsubTopicRef: Name references are not yet supported, use External instead")
			return nil
		}
		out.PubsubTopic = in.PubsubTopicRef.External
	}
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.MonitoringNotificationChannels = BudgetNotificationsRule_MonitoringNotificationChannels_ToProto(mapCtx, in.MonitoringNotificationChannels)
	out.DisableDefaultIamRecipients = direct.ValueOf(in.DisableDefaultIAMRecipients)
	return out
}

func BudgetAmount_FromProto(mapCtx *direct.MapContext, in *pb.BudgetAmount) *krm.BudgetAmount {
	if in == nil {
		return nil
	}
	out := &krm.BudgetAmount{}
	out.SpecifiedAmount = Money_FromProto(mapCtx, in.GetSpecifiedAmount())
	if in.GetLastPeriodAmount() != nil {
		out.LastPeriodAmount = &apiextensionsv1.JSON{Raw: []byte("{}")}
	}
	return out
}

func BudgetAmount_ToProto(mapCtx *direct.MapContext, in *krm.BudgetAmount) *pb.BudgetAmount {
	if in == nil {
		return nil
	}
	out := &pb.BudgetAmount{}
	if in.SpecifiedAmount != nil && in.LastPeriodAmount != nil {
		mapCtx.Errorf("only one of SpecifiedAmount or LastPeriodAmount may be set")
		return nil
	}
	if in.SpecifiedAmount != nil {
		out.BudgetAmount = &pb.BudgetAmount_SpecifiedAmount{
			SpecifiedAmount: Money_ToProto(mapCtx, in.SpecifiedAmount),
		}
	}
	if in.LastPeriodAmount != nil {
		out.BudgetAmount = &pb.BudgetAmount_LastPeriodAmount{
			LastPeriodAmount: &pb.LastPeriodAmount{},
		}
	}
	return out
}

func BudgetFilter_FromProto(mapCtx *direct.MapContext, in *pb.Filter) *krm.BudgetFilter {
	if in == nil {
		return nil
	}
	out := &krm.BudgetFilter{}
	out.Projects = BudgetFilter_Projects_FromProto(mapCtx, in.Projects)
	out.CreditTypes = in.CreditTypes
	out.CreditTypesTreatment = direct.Enum_FromProto(mapCtx, in.GetCreditTypesTreatment())
	out.Services = in.Services
	out.Subaccounts = BudgetFilter_Subaccounts_FromProto(mapCtx, in.Subaccounts)

	if in.GetLabels() != nil {
		out.Labels = make(map[string]krm.BudgetLabels)
		for k, v := range in.GetLabels() {
			var vals []string
			if v != nil {
				for _, val := range v.GetValues() {
					vals = append(vals, val.GetStringValue())
				}
			}
			out.Labels[k] = krm.BudgetLabels{Values: vals}
		}
	}

	out.CalendarPeriod = direct.Enum_FromProto(mapCtx, in.GetCalendarPeriod())
	out.CustomPeriod = BudgetCustomPeriod_FromProto(mapCtx, in.GetCustomPeriod())
	return out
}

func BudgetFilter_ToProto(mapCtx *direct.MapContext, in *krm.BudgetFilter) *pb.Filter {
	if in == nil {
		return nil
	}
	out := &pb.Filter{}
	out.Projects = BudgetFilter_Projects_ToProto(mapCtx, in.Projects)
	out.CreditTypes = in.CreditTypes
	out.CreditTypesTreatment = direct.Enum_ToProto[pb.Filter_CreditTypesTreatment](mapCtx, in.CreditTypesTreatment)
	out.Services = in.Services
	out.Subaccounts = BudgetFilter_Subaccounts_ToProto(mapCtx, in.Subaccounts)

	if in.Labels != nil {
		out.Labels = make(map[string]*structpb.ListValue)
		for k, v := range in.Labels {
			lv := &structpb.ListValue{}
			for _, val := range v.Values {
				lv.Values = append(lv.Values, structpb.NewStringValue(val))
			}
			out.Labels[k] = lv
		}
	}

	if direct.ValueOf(in.CalendarPeriod) != "" && in.CustomPeriod != nil {
		mapCtx.Errorf("only one of calendarPeriod or customPeriod may be set")
		return nil
	}

	if direct.ValueOf(in.CalendarPeriod) != "" {
		out.UsagePeriod = &pb.Filter_CalendarPeriod{CalendarPeriod: direct.Enum_ToProto[pb.CalendarPeriod](mapCtx, in.CalendarPeriod)}
	} else if in.CustomPeriod != nil {
		out.UsagePeriod = &pb.Filter_CustomPeriod{CustomPeriod: BudgetCustomPeriod_ToProto(mapCtx, in.CustomPeriod)}
	}
	return out
}

func Money_FromProto(mapCtx *direct.MapContext, in *moneypb.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(int64(in.GetNanos()))
	return out
}

func Money_ToProto(mapCtx *direct.MapContext, in *krm.Money) *moneypb.Money {
	if in == nil {
		return nil
	}
	out := &moneypb.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = int32(direct.ValueOf(in.Nanos))
	return out
}

func Date_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *krm.Date {
	if in == nil {
		return nil
	}
	out := &krm.Date{}
	out.Year = direct.LazyPtr(int64(in.GetYear()))
	out.Month = direct.LazyPtr(int64(in.GetMonth()))
	out.Day = direct.LazyPtr(int64(in.GetDay()))
	return out
}

func Date_ToProto(mapCtx *direct.MapContext, in *krm.Date) *datepb.Date {
	if in == nil {
		return nil
	}
	out := &datepb.Date{}
	out.Year = int32(direct.ValueOf(in.Year))
	out.Month = int32(direct.ValueOf(in.Month))
	out.Day = int32(direct.ValueOf(in.Day))
	return out
}

func BudgetFilter_Projects_FromProto(mapCtx *direct.MapContext, in []string) []krm.ProjectRef {
	if in == nil {
		return nil
	}
	out := make([]krm.ProjectRef, len(in))
	for i, s := range in {
		out[i] = krm.ProjectRef{External: s}
	}
	return out
}

func BudgetFilter_Projects_ToProto(mapCtx *direct.MapContext, in []krm.ProjectRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.Name != "" {
			mapCtx.Errorf("ProjectRef: Name references are not yet supported, use External instead")
			return nil
		}
		out[i] = r.External
	}
	return out
}

func BudgetFilter_Subaccounts_FromProto(mapCtx *direct.MapContext, in []string) []krm.CloudBillingBillingAccountRef {
	if in == nil {
		return nil
	}
	out := make([]krm.CloudBillingBillingAccountRef, len(in))
	for i, s := range in {
		out[i] = krm.CloudBillingBillingAccountRef{External: s}
	}
	return out
}

func BudgetFilter_Subaccounts_ToProto(mapCtx *direct.MapContext, in []krm.CloudBillingBillingAccountRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.Name != "" {
			mapCtx.Errorf("CloudBillingBillingAccountRef: Name references are not yet supported, use External instead")
			return nil
		}
		out[i] = r.External
	}
	return out
}

func BudgetNotificationsRule_MonitoringNotificationChannels_FromProto(mapCtx *direct.MapContext, in []string) []krm.MonitoringNotificationChannelRef {
	if in == nil {
		return nil
	}
	out := make([]krm.MonitoringNotificationChannelRef, len(in))
	for i, s := range in {
		out[i] = krm.MonitoringNotificationChannelRef{External: s}
	}
	return out
}

func BudgetNotificationsRule_MonitoringNotificationChannels_ToProto(mapCtx *direct.MapContext, in []krm.MonitoringNotificationChannelRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.Name != "" {
			mapCtx.Errorf("MonitoringNotificationChannelRef: Name references are not yet supported, use External instead")
			return nil
		}
		out[i] = r.External
	}
	return out
}
