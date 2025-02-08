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

package paymentgateway

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/paymentgateway/issuerswitch/accountmanager/apiv1/accountmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ManagedAccount_FromProto(mapCtx *direct.MapContext, in *pb.ManagedAccount) *krm.ManagedAccount {
	if in == nil {
		return nil
	}
	out := &krm.ManagedAccount{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AccountReference = AccountReference_FromProto(mapCtx, in.GetAccountReference())
	// MISSING: State
	out.Balance = Money_FromProto(mapCtx, in.GetBalance())
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ManagedAccount_ToProto(mapCtx *direct.MapContext, in *krm.ManagedAccount) *pb.ManagedAccount {
	if in == nil {
		return nil
	}
	out := &pb.ManagedAccount{}
	out.Name = direct.ValueOf(in.Name)
	out.AccountReference = AccountReference_ToProto(mapCtx, in.AccountReference)
	// MISSING: State
	out.Balance = Money_ToProto(mapCtx, in.Balance)
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ManagedAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedAccount) *krm.ManagedAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedAccountObservedState{}
	// MISSING: Name
	out.AccountReference = AccountReferenceObservedState_FromProto(mapCtx, in.GetAccountReference())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Balance
	out.LastReconciliationState = direct.Enum_FromProto(mapCtx, in.GetLastReconciliationState())
	out.LastReconciliationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastReconciliationTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ManagedAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedAccountObservedState) *pb.ManagedAccount {
	if in == nil {
		return nil
	}
	out := &pb.ManagedAccount{}
	// MISSING: Name
	out.AccountReference = AccountReferenceObservedState_ToProto(mapCtx, in.AccountReference)
	out.State = direct.Enum_ToProto[pb.ManagedAccount_State](mapCtx, in.State)
	// MISSING: Balance
	out.LastReconciliationState = direct.Enum_ToProto[pb.ManagedAccount_AccountReconciliationState](mapCtx, in.LastReconciliationState)
	out.LastReconciliationTime = direct.StringTimestamp_ToProto(mapCtx, in.LastReconciliationTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func PaymentgatewayManagedAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedAccount) *krm.PaymentgatewayManagedAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayManagedAccountObservedState{}
	// MISSING: Name
	// MISSING: AccountReference
	// MISSING: State
	// MISSING: Balance
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PaymentgatewayManagedAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayManagedAccountObservedState) *pb.ManagedAccount {
	if in == nil {
		return nil
	}
	out := &pb.ManagedAccount{}
	// MISSING: Name
	// MISSING: AccountReference
	// MISSING: State
	// MISSING: Balance
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PaymentgatewayManagedAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagedAccount) *krm.PaymentgatewayManagedAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayManagedAccountSpec{}
	// MISSING: Name
	// MISSING: AccountReference
	// MISSING: State
	// MISSING: Balance
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PaymentgatewayManagedAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayManagedAccountSpec) *pb.ManagedAccount {
	if in == nil {
		return nil
	}
	out := &pb.ManagedAccount{}
	// MISSING: Name
	// MISSING: AccountReference
	// MISSING: State
	// MISSING: Balance
	// MISSING: LastReconciliationState
	// MISSING: LastReconciliationTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
