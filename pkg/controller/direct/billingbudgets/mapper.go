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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

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
		out[i] = r.External
	}
	return out
}
