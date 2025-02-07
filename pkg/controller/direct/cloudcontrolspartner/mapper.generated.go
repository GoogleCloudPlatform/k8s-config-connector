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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudcontrolspartnerWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadObservedState{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadSpec{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func Workload_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.Workload {
	if in == nil {
		return nil
	}
	out := &krm.Workload{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	out.WorkloadOnboardingState = WorkloadOnboardingState_FromProto(mapCtx, in.GetWorkloadOnboardingState())
	out.IsOnboarded = direct.LazyPtr(in.GetIsOnboarded())
	out.KeyManagementProjectID = direct.LazyPtr(in.GetKeyManagementProjectId())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Partner = direct.Enum_FromProto(mapCtx, in.GetPartner())
	return out
}
func Workload_ToProto(mapCtx *direct.MapContext, in *krm.Workload) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	out.WorkloadOnboardingState = WorkloadOnboardingState_ToProto(mapCtx, in.WorkloadOnboardingState)
	out.IsOnboarded = direct.ValueOf(in.IsOnboarded)
	out.KeyManagementProjectId = direct.ValueOf(in.KeyManagementProjectID)
	out.Location = direct.ValueOf(in.Location)
	out.Partner = direct.Enum_ToProto[pb.Workload_Partner](mapCtx, in.Partner)
	return out
}
func WorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.WorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadObservedState{}
	// MISSING: Name
	out.FolderID = direct.LazyPtr(in.GetFolderId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Folder = direct.LazyPtr(in.GetFolder())
	out.WorkloadOnboardingState = WorkloadOnboardingStateObservedState_FromProto(mapCtx, in.GetWorkloadOnboardingState())
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func WorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	out.FolderId = direct.ValueOf(in.FolderID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Folder = direct.ValueOf(in.Folder)
	out.WorkloadOnboardingState = WorkloadOnboardingStateObservedState_ToProto(mapCtx, in.WorkloadOnboardingState)
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func WorkloadOnboardingState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadOnboardingState) *krm.WorkloadOnboardingState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadOnboardingState{}
	out.OnboardingSteps = direct.Slice_FromProto(mapCtx, in.OnboardingSteps, WorkloadOnboardingStep_FromProto)
	return out
}
func WorkloadOnboardingState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadOnboardingState) *pb.WorkloadOnboardingState {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadOnboardingState{}
	out.OnboardingSteps = direct.Slice_ToProto(mapCtx, in.OnboardingSteps, WorkloadOnboardingStep_ToProto)
	return out
}
func WorkloadOnboardingStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadOnboardingState) *krm.WorkloadOnboardingStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadOnboardingStateObservedState{}
	out.OnboardingSteps = direct.Slice_FromProto(mapCtx, in.OnboardingSteps, WorkloadOnboardingStepObservedState_FromProto)
	return out
}
func WorkloadOnboardingStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadOnboardingStateObservedState) *pb.WorkloadOnboardingState {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadOnboardingState{}
	out.OnboardingSteps = direct.Slice_ToProto(mapCtx, in.OnboardingSteps, WorkloadOnboardingStepObservedState_ToProto)
	return out
}
func WorkloadOnboardingStep_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadOnboardingStep) *krm.WorkloadOnboardingStep {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadOnboardingStep{}
	out.Step = direct.Enum_FromProto(mapCtx, in.GetStep())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.CompletionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompletionTime())
	// MISSING: CompletionState
	return out
}
func WorkloadOnboardingStep_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadOnboardingStep) *pb.WorkloadOnboardingStep {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadOnboardingStep{}
	out.Step = direct.Enum_ToProto[pb.WorkloadOnboardingStep_Step](mapCtx, in.Step)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.CompletionTime = direct.StringTimestamp_ToProto(mapCtx, in.CompletionTime)
	// MISSING: CompletionState
	return out
}
func WorkloadOnboardingStepObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadOnboardingStep) *krm.WorkloadOnboardingStepObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadOnboardingStepObservedState{}
	// MISSING: Step
	// MISSING: StartTime
	// MISSING: CompletionTime
	out.CompletionState = direct.Enum_FromProto(mapCtx, in.GetCompletionState())
	return out
}
func WorkloadOnboardingStepObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadOnboardingStepObservedState) *pb.WorkloadOnboardingStep {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadOnboardingStep{}
	// MISSING: Step
	// MISSING: StartTime
	// MISSING: CompletionTime
	out.CompletionState = direct.Enum_ToProto[pb.CompletionState](mapCtx, in.CompletionState)
	return out
}
