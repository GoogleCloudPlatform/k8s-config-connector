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

package cloudcontrolspartner

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1beta/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
)
func Customer_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.Customer {
	if in == nil {
		return nil
	}
	out := &krm.Customer{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func Customer_ToProto(mapCtx *direct.MapContext, in *krm.Customer) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CustomerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CustomerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CustomerOnboardingState = CustomerOnboardingState_FromProto(mapCtx, in.GetCustomerOnboardingState())
	out.IsOnboarded = direct.LazyPtr(in.GetIsOnboarded())
	return out
}
func CustomerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerObservedState) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CustomerOnboardingState = CustomerOnboardingState_ToProto(mapCtx, in.CustomerOnboardingState)
	out.IsOnboarded = direct.ValueOf(in.IsOnboarded)
	return out
}
func CustomerOnboardingState_FromProto(mapCtx *direct.MapContext, in *pb.CustomerOnboardingState) *krm.CustomerOnboardingState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerOnboardingState{}
	out.OnboardingSteps = direct.Slice_FromProto(mapCtx, in.OnboardingSteps, CustomerOnboardingStep_FromProto)
	return out
}
func CustomerOnboardingState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerOnboardingState) *pb.CustomerOnboardingState {
	if in == nil {
		return nil
	}
	out := &pb.CustomerOnboardingState{}
	out.OnboardingSteps = direct.Slice_ToProto(mapCtx, in.OnboardingSteps, CustomerOnboardingStep_ToProto)
	return out
}
func CustomerOnboardingStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomerOnboardingState) *krm.CustomerOnboardingStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerOnboardingStateObservedState{}
	out.OnboardingSteps = direct.Slice_FromProto(mapCtx, in.OnboardingSteps, CustomerOnboardingStepObservedState_FromProto)
	return out
}
func CustomerOnboardingStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerOnboardingStateObservedState) *pb.CustomerOnboardingState {
	if in == nil {
		return nil
	}
	out := &pb.CustomerOnboardingState{}
	out.OnboardingSteps = direct.Slice_ToProto(mapCtx, in.OnboardingSteps, CustomerOnboardingStepObservedState_ToProto)
	return out
}
func CustomerOnboardingStep_FromProto(mapCtx *direct.MapContext, in *pb.CustomerOnboardingStep) *krm.CustomerOnboardingStep {
	if in == nil {
		return nil
	}
	out := &krm.CustomerOnboardingStep{}
	out.Step = direct.Enum_FromProto(mapCtx, in.GetStep())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.CompletionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompletionTime())
	// MISSING: CompletionState
	return out
}
func CustomerOnboardingStep_ToProto(mapCtx *direct.MapContext, in *krm.CustomerOnboardingStep) *pb.CustomerOnboardingStep {
	if in == nil {
		return nil
	}
	out := &pb.CustomerOnboardingStep{}
	out.Step = direct.Enum_ToProto[pb.CustomerOnboardingStep_Step](mapCtx, in.Step)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.CompletionTime = direct.StringTimestamp_ToProto(mapCtx, in.CompletionTime)
	// MISSING: CompletionState
	return out
}
func CustomerOnboardingStepObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomerOnboardingStep) *krm.CustomerOnboardingStepObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomerOnboardingStepObservedState{}
	// MISSING: Step
	// MISSING: StartTime
	// MISSING: CompletionTime
	out.CompletionState = direct.Enum_FromProto(mapCtx, in.GetCompletionState())
	return out
}
func CustomerOnboardingStepObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomerOnboardingStepObservedState) *pb.CustomerOnboardingStep {
	if in == nil {
		return nil
	}
	out := &pb.CustomerOnboardingStep{}
	// MISSING: Step
	// MISSING: StartTime
	// MISSING: CompletionTime
	out.CompletionState = direct.Enum_ToProto[pb.CompletionState](mapCtx, in.CompletionState)
	return out
}
