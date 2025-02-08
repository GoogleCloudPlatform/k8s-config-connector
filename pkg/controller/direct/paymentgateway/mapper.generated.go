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
func PaymentgatewayRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Rule) *krm.PaymentgatewayRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayRuleObservedState{}
	// MISSING: Name
	// MISSING: RuleDescription
	// MISSING: ApiType
	// MISSING: TransactionType
	return out
}
func PaymentgatewayRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayRuleObservedState) *pb.Rule {
	if in == nil {
		return nil
	}
	out := &pb.Rule{}
	// MISSING: Name
	// MISSING: RuleDescription
	// MISSING: ApiType
	// MISSING: TransactionType
	return out
}
func PaymentgatewayRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.Rule) *krm.PaymentgatewayRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayRuleSpec{}
	// MISSING: Name
	// MISSING: RuleDescription
	// MISSING: ApiType
	// MISSING: TransactionType
	return out
}
func PaymentgatewayRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayRuleSpec) *pb.Rule {
	if in == nil {
		return nil
	}
	out := &pb.Rule{}
	// MISSING: Name
	// MISSING: RuleDescription
	// MISSING: ApiType
	// MISSING: TransactionType
	return out
}
func Rule_FromProto(mapCtx *direct.MapContext, in *pb.Rule) *krm.Rule {
	if in == nil {
		return nil
	}
	out := &krm.Rule{}
	out.Name = direct.LazyPtr(in.GetName())
	out.RuleDescription = direct.LazyPtr(in.GetRuleDescription())
	out.ApiType = direct.Enum_FromProto(mapCtx, in.GetApiType())
	out.TransactionType = direct.Enum_FromProto(mapCtx, in.GetTransactionType())
	return out
}
func Rule_ToProto(mapCtx *direct.MapContext, in *krm.Rule) *pb.Rule {
	if in == nil {
		return nil
	}
	out := &pb.Rule{}
	out.Name = direct.ValueOf(in.Name)
	out.RuleDescription = direct.ValueOf(in.RuleDescription)
	out.ApiType = direct.Enum_ToProto[pb.ApiType](mapCtx, in.ApiType)
	out.TransactionType = direct.Enum_ToProto[pb.TransactionType](mapCtx, in.TransactionType)
	return out
}
