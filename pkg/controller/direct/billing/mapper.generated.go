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
// krm.group: billing.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.billing.v1

package billing

import (
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccountObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Open = direct.LazyPtr(in.GetOpen())
	out.MasterBillingAccount = direct.LazyPtr(in.GetMasterBillingAccount())
	// MISSING: Parent
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	return out
}
func BillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	out.Name = direct.ValueOf(in.Name)
	out.Open = direct.ValueOf(in.Open)
	out.MasterBillingAccount = direct.ValueOf(in.MasterBillingAccount)
	// MISSING: Parent
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	return out
}
func BillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccountSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Parent
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	return out
}
func BillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Parent
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	return out
}
