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

package retail

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
)
func AlertConfig_FromProto(mapCtx *direct.MapContext, in *pb.AlertConfig) *krm.AlertConfig {
	if in == nil {
		return nil
	}
	out := &krm.AlertConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AlertPolicies = direct.Slice_FromProto(mapCtx, in.AlertPolicies, AlertConfig_AlertPolicy_FromProto)
	return out
}
func AlertConfig_ToProto(mapCtx *direct.MapContext, in *krm.AlertConfig) *pb.AlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlertConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.AlertPolicies = direct.Slice_ToProto(mapCtx, in.AlertPolicies, AlertConfig_AlertPolicy_ToProto)
	return out
}
func AlertConfig_AlertPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AlertConfig_AlertPolicy) *krm.AlertConfig_AlertPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AlertConfig_AlertPolicy{}
	out.AlertGroup = direct.LazyPtr(in.GetAlertGroup())
	out.EnrollStatus = direct.Enum_FromProto(mapCtx, in.GetEnrollStatus())
	out.Recipients = direct.Slice_FromProto(mapCtx, in.Recipients, AlertConfig_AlertPolicy_Recipient_FromProto)
	return out
}
func AlertConfig_AlertPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AlertConfig_AlertPolicy) *pb.AlertConfig_AlertPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AlertConfig_AlertPolicy{}
	out.AlertGroup = direct.ValueOf(in.AlertGroup)
	out.EnrollStatus = direct.Enum_ToProto[pb.AlertConfig_AlertPolicy_EnrollStatus](mapCtx, in.EnrollStatus)
	out.Recipients = direct.Slice_ToProto(mapCtx, in.Recipients, AlertConfig_AlertPolicy_Recipient_ToProto)
	return out
}
func AlertConfig_AlertPolicy_Recipient_FromProto(mapCtx *direct.MapContext, in *pb.AlertConfig_AlertPolicy_Recipient) *krm.AlertConfig_AlertPolicy_Recipient {
	if in == nil {
		return nil
	}
	out := &krm.AlertConfig_AlertPolicy_Recipient{}
	out.EmailAddress = direct.LazyPtr(in.GetEmailAddress())
	return out
}
func AlertConfig_AlertPolicy_Recipient_ToProto(mapCtx *direct.MapContext, in *krm.AlertConfig_AlertPolicy_Recipient) *pb.AlertConfig_AlertPolicy_Recipient {
	if in == nil {
		return nil
	}
	out := &pb.AlertConfig_AlertPolicy_Recipient{}
	out.EmailAddress = direct.ValueOf(in.EmailAddress)
	return out
}
func RetailAlertConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AlertConfig) *krm.RetailAlertConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RetailAlertConfigObservedState{}
	// MISSING: Name
	// MISSING: AlertPolicies
	return out
}
func RetailAlertConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RetailAlertConfigObservedState) *pb.AlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlertConfig{}
	// MISSING: Name
	// MISSING: AlertPolicies
	return out
}
func RetailAlertConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AlertConfig) *krm.RetailAlertConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.RetailAlertConfigSpec{}
	// MISSING: Name
	// MISSING: AlertPolicies
	return out
}
func RetailAlertConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.RetailAlertConfigSpec) *pb.AlertConfig {
	if in == nil {
		return nil
	}
	out := &pb.AlertConfig{}
	// MISSING: Name
	// MISSING: AlertPolicies
	return out
}
