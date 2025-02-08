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
	pb "cloud.google.com/go/paymentgateway/issuerswitch/apiv1/issuerswitchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AccountReference_FromProto(mapCtx *direct.MapContext, in *pb.AccountReference) *krm.AccountReference {
	if in == nil {
		return nil
	}
	out := &krm.AccountReference{}
	out.Ifsc = direct.LazyPtr(in.GetIfsc())
	// MISSING: AccountType
	out.AccountNumber = direct.LazyPtr(in.GetAccountNumber())
	return out
}
func AccountReference_ToProto(mapCtx *direct.MapContext, in *krm.AccountReference) *pb.AccountReference {
	if in == nil {
		return nil
	}
	out := &pb.AccountReference{}
	out.Ifsc = direct.ValueOf(in.Ifsc)
	// MISSING: AccountType
	out.AccountNumber = direct.ValueOf(in.AccountNumber)
	return out
}
func AccountReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountReference) *krm.AccountReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountReferenceObservedState{}
	// MISSING: Ifsc
	out.AccountType = direct.LazyPtr(in.GetAccountType())
	// MISSING: AccountNumber
	return out
}
func AccountReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountReferenceObservedState) *pb.AccountReference {
	if in == nil {
		return nil
	}
	out := &pb.AccountReference{}
	// MISSING: Ifsc
	out.AccountType = direct.ValueOf(in.AccountType)
	// MISSING: AccountNumber
	return out
}
func PaymentgatewayRuleMetadataValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuleMetadataValue) *krm.PaymentgatewayRuleMetadataValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayRuleMetadataValueObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AccountReference
	return out
}
func PaymentgatewayRuleMetadataValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayRuleMetadataValueObservedState) *pb.RuleMetadataValue {
	if in == nil {
		return nil
	}
	out := &pb.RuleMetadataValue{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AccountReference
	return out
}
func PaymentgatewayRuleMetadataValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.RuleMetadataValue) *krm.PaymentgatewayRuleMetadataValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayRuleMetadataValueSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AccountReference
	return out
}
func PaymentgatewayRuleMetadataValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayRuleMetadataValueSpec) *pb.RuleMetadataValue {
	if in == nil {
		return nil
	}
	out := &pb.RuleMetadataValue{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AccountReference
	return out
}
func RuleMetadataValue_FromProto(mapCtx *direct.MapContext, in *pb.RuleMetadataValue) *krm.RuleMetadataValue {
	if in == nil {
		return nil
	}
	out := &krm.RuleMetadataValue{}
	// MISSING: Name
	out.ID = direct.LazyPtr(in.GetId())
	out.AccountReference = AccountReference_FromProto(mapCtx, in.GetAccountReference())
	return out
}
func RuleMetadataValue_ToProto(mapCtx *direct.MapContext, in *krm.RuleMetadataValue) *pb.RuleMetadataValue {
	if in == nil {
		return nil
	}
	out := &pb.RuleMetadataValue{}
	// MISSING: Name
	if oneof := RuleMetadataValue_Id_ToProto(mapCtx, in.ID); oneof != nil {
		out.Value = oneof
	}
	if oneof := AccountReference_ToProto(mapCtx, in.AccountReference); oneof != nil {
		out.Value = &pb.RuleMetadataValue_AccountReference{AccountReference: oneof}
	}
	return out
}
func RuleMetadataValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RuleMetadataValue) *krm.RuleMetadataValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RuleMetadataValueObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	out.AccountReference = AccountReferenceObservedState_FromProto(mapCtx, in.GetAccountReference())
	return out
}
func RuleMetadataValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RuleMetadataValueObservedState) *pb.RuleMetadataValue {
	if in == nil {
		return nil
	}
	out := &pb.RuleMetadataValue{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	if oneof := AccountReferenceObservedState_ToProto(mapCtx, in.AccountReference); oneof != nil {
		out.Value = &pb.RuleMetadataValue_AccountReference{AccountReference: oneof}
	}
	return out
}
