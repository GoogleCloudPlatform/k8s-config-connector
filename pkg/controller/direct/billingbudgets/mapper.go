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

package billingbudgets

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datepb "google.golang.org/genproto/googleapis/type/date"
	moneypb "google.golang.org/genproto/googleapis/type/money"
)

func Money_FromProto(mapCtx *direct.MapContext, in *moneypb.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_ToProto(mapCtx *direct.MapContext, in *krm.Money) *moneypb.Money {
	if in == nil {
		return nil
	}
	out := &moneypb.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func Date_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *krm.Date {
	if in == nil {
		return nil
	}
	out := &krm.Date{}
	out.Year = direct.LazyPtr(in.GetYear())
	out.Month = direct.LazyPtr(in.GetMonth())
	out.Day = direct.LazyPtr(in.GetDay())
	return out
}

func Date_ToProto(mapCtx *direct.MapContext, in *krm.Date) *datepb.Date {
	if in == nil {
		return nil
	}
	out := &datepb.Date{}
	out.Year = direct.ValueOf(in.Year)
	out.Month = direct.ValueOf(in.Month)
	out.Day = direct.ValueOf(in.Day)
	return out
}
