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
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BillingAccount_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccount {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.MasterBillingAccount = direct.LazyPtr(in.GetMasterBillingAccount())
	// MISSING: Parent
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	return out
}
func BillingAccount_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccount) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.MasterBillingAccount = direct.ValueOf(in.MasterBillingAccount)
	// MISSING: Parent
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	return out
}
func BillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccountObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Open = direct.LazyPtr(in.GetOpen())
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	out.Parent = direct.LazyPtr(in.GetParent())
	// MISSING: CurrencyCode
	return out
}
func BillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	out.Name = direct.ValueOf(in.Name)
	out.Open = direct.ValueOf(in.Open)
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	out.Parent = direct.ValueOf(in.Parent)
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountObservedState{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountSpec{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
