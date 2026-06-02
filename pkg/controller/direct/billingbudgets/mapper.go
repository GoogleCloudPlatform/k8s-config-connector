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
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gdate "google.golang.org/genproto/googleapis/type/date"
	money "google.golang.org/genproto/googleapis/type/money"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func AllUpdatesRule_MonitoringNotificationChannels_FromProto(mapCtx *direct.MapContext, in []string) []monitoringv1beta1.MonitoringNotificationChannelRef {
	out := make([]monitoringv1beta1.MonitoringNotificationChannelRef, len(in))
	for i, v := range in {
		out[i] = monitoringv1beta1.MonitoringNotificationChannelRef{External: v}
	}
	return out
}

func AllUpdatesRule_MonitoringNotificationChannels_ToProto(mapCtx *direct.MapContext, in []monitoringv1beta1.MonitoringNotificationChannelRef) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func BudgetFilter_Projects_FromProto(mapCtx *direct.MapContext, in []string) []refsv1beta1.ProjectRef {
	out := make([]refsv1beta1.ProjectRef, len(in))
	for i, v := range in {
		out[i] = refsv1beta1.ProjectRef{External: v}
	}
	return out
}

func BudgetFilter_Projects_ToProto(mapCtx *direct.MapContext, in []refsv1beta1.ProjectRef) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func BudgetFilter_Subaccounts_FromProto(mapCtx *direct.MapContext, in []string) []refs.BillingAccountRef {
	out := make([]refs.BillingAccountRef, len(in))
	for i, v := range in {
		out[i] = refs.BillingAccountRef{External: v}
	}
	return out
}

func BudgetFilter_Subaccounts_ToProto(mapCtx *direct.MapContext, in []refs.BillingAccountRef) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func BudgetSpecifiedAmount_FromProto(mapCtx *direct.MapContext, in *money.Money) *krm.BudgetSpecifiedAmount {
	if in == nil {
		return nil
	}
	out := &krm.BudgetSpecifiedAmount{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	units := in.GetUnits()
	out.Units = &units
	nanos := int64(in.GetNanos())
	out.Nanos = &nanos
	return out
}

func BudgetSpecifiedAmount_ToProto(mapCtx *direct.MapContext, in *krm.BudgetSpecifiedAmount) *money.Money {
	if in == nil {
		return nil
	}
	out := &money.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = int32(direct.ValueOf(in.Nanos))
	return out
}

func BudgetDate_FromProto(mapCtx *direct.MapContext, in *gdate.Date) *krm.BudgetDate {
	if in == nil {
		return nil
	}
	out := &krm.BudgetDate{}
	year := int64(in.GetYear())
	out.Year = &year
	month := int64(in.GetMonth())
	out.Month = &month
	day := int64(in.GetDay())
	out.Day = &day
	return out
}

func BudgetDate_ToProto(mapCtx *direct.MapContext, in *krm.BudgetDate) *gdate.Date {
	if in == nil {
		return nil
	}
	out := &gdate.Date{}
	out.Year = int32(direct.ValueOf(in.Year))
	out.Month = int32(direct.ValueOf(in.Month))
	out.Day = int32(direct.ValueOf(in.Day))
	return out
}

func BudgetFilter_Labels_FromProto(mapCtx *direct.MapContext, in map[string]*structpb.ListValue) map[string]krm.FilterLabels {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.FilterLabels)
	for k, v := range in {
		var values []string
		for _, val := range v.GetValues() {
			values = append(values, val.GetStringValue())
		}
		out[k] = krm.FilterLabels{Values: values}
	}
	return out
}

func BudgetFilter_Labels_ToProto(mapCtx *direct.MapContext, in map[string]krm.FilterLabels) map[string]*structpb.ListValue {
	if in == nil {
		return nil
	}
	out := make(map[string]*structpb.ListValue)
	for k, v := range in {
		var list []*structpb.Value
		for _, s := range v.Values {
			list = append(list, structpb.NewStringValue(s))
		}
		out[k] = &structpb.ListValue{Values: list}
	}
	return out
}

func BudgetAmount_FromProto(mapCtx *direct.MapContext, in *pb.BudgetAmount) krm.BudgetAmount {
	if in == nil {
		return krm.BudgetAmount{}
	}
	out := krm.BudgetAmount{}
	out.SpecifiedAmount = BudgetSpecifiedAmount_FromProto(mapCtx, in.GetSpecifiedAmount())
	out.LastPeriodAmount = BudgetLastPeriodAmount_FromProto(mapCtx, in.GetLastPeriodAmount())
	return out
}

func BudgetAmount_ToProto(mapCtx *direct.MapContext, in krm.BudgetAmount) *pb.BudgetAmount {
	out := &pb.BudgetAmount{}
	if oneof := BudgetSpecifiedAmount_ToProto(mapCtx, in.SpecifiedAmount); oneof != nil {
		out.BudgetAmount = &pb.BudgetAmount_SpecifiedAmount{SpecifiedAmount: oneof}
	}
	if oneof := BudgetLastPeriodAmount_ToProto(mapCtx, in.LastPeriodAmount); oneof != nil {
		out.BudgetAmount = &pb.BudgetAmount_LastPeriodAmount{LastPeriodAmount: oneof}
	}
	return out
}

func BudgetCustomPeriod_FromProto(mapCtx *direct.MapContext, in *pb.CustomPeriod) *krm.BudgetCustomPeriod {
	if in == nil {
		return nil
	}
	out := &krm.BudgetCustomPeriod{}
	if in.GetStartDate() != nil {
		out.StartDate = *BudgetDate_FromProto(mapCtx, in.GetStartDate())
	}
	out.EndDate = BudgetDate_FromProto(mapCtx, in.GetEndDate())
	return out
}

func BudgetCustomPeriod_ToProto(mapCtx *direct.MapContext, in *krm.BudgetCustomPeriod) *pb.CustomPeriod {
	if in == nil {
		return nil
	}
	out := &pb.CustomPeriod{}
	out.StartDate = BudgetDate_ToProto(mapCtx, &in.StartDate)
	out.EndDate = BudgetDate_ToProto(mapCtx, in.EndDate)
	return out
}

func BudgetThresholdRule_FromProto(mapCtx *direct.MapContext, in *pb.ThresholdRule) *krm.BudgetThresholdRule {
	if in == nil {
		return nil
	}
	out := &krm.BudgetThresholdRule{}
	out.ThresholdPercent = in.GetThresholdPercent()
	out.SpendBasis = direct.Enum_FromProto(mapCtx, in.GetSpendBasis())
	return out
}

func BudgetThresholdRule_ToProto(mapCtx *direct.MapContext, in *krm.BudgetThresholdRule) *pb.ThresholdRule {
	if in == nil {
		return nil
	}
	out := &pb.ThresholdRule{}
	out.ThresholdPercent = in.ThresholdPercent
	out.SpendBasis = direct.Enum_ToProto[pb.ThresholdRule_Basis](mapCtx, in.SpendBasis)
	return out
}
