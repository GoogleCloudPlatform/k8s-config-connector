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

package deploy

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
)
func AdvanceChildRolloutJob_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceChildRolloutJob) *krm.AdvanceChildRolloutJob {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceChildRolloutJob{}
	return out
}
func AdvanceChildRolloutJob_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceChildRolloutJob) *pb.AdvanceChildRolloutJob {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceChildRolloutJob{}
	return out
}
func ChildRolloutJobs_FromProto(mapCtx *direct.MapContext, in *pb.ChildRolloutJobs) *krm.ChildRolloutJobs {
	if in == nil {
		return nil
	}
	out := &krm.ChildRolloutJobs{}
	// MISSING: CreateRolloutJobs
	// MISSING: AdvanceRolloutJobs
	return out
}
func ChildRolloutJobs_ToProto(mapCtx *direct.MapContext, in *krm.ChildRolloutJobs) *pb.ChildRolloutJobs {
	if in == nil {
		return nil
	}
	out := &pb.ChildRolloutJobs{}
	// MISSING: CreateRolloutJobs
	// MISSING: AdvanceRolloutJobs
	return out
}
func ChildRolloutJobsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ChildRolloutJobs) *krm.ChildRolloutJobsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChildRolloutJobsObservedState{}
	out.CreateRolloutJobs = direct.Slice_FromProto(mapCtx, in.CreateRolloutJobs, Job_FromProto)
	out.AdvanceRolloutJobs = direct.Slice_FromProto(mapCtx, in.AdvanceRolloutJobs, Job_FromProto)
	return out
}
func ChildRolloutJobsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChildRolloutJobsObservedState) *pb.ChildRolloutJobs {
	if in == nil {
		return nil
	}
	out := &pb.ChildRolloutJobs{}
	out.CreateRolloutJobs = direct.Slice_ToProto(mapCtx, in.CreateRolloutJobs, Job_ToProto)
	out.AdvanceRolloutJobs = direct.Slice_ToProto(mapCtx, in.AdvanceRolloutJobs, Job_ToProto)
	return out
}
func CreateChildRolloutJob_FromProto(mapCtx *direct.MapContext, in *pb.CreateChildRolloutJob) *krm.CreateChildRolloutJob {
	if in == nil {
		return nil
	}
	out := &krm.CreateChildRolloutJob{}
	return out
}
func CreateChildRolloutJob_ToProto(mapCtx *direct.MapContext, in *krm.CreateChildRolloutJob) *pb.CreateChildRolloutJob {
	if in == nil {
		return nil
	}
	out := &pb.CreateChildRolloutJob{}
	return out
}
func DeployJob_FromProto(mapCtx *direct.MapContext, in *pb.DeployJob) *krm.DeployJob {
	if in == nil {
		return nil
	}
	out := &krm.DeployJob{}
	return out
}
func DeployJob_ToProto(mapCtx *direct.MapContext, in *krm.DeployJob) *pb.DeployJob {
	if in == nil {
		return nil
	}
	out := &pb.DeployJob{}
	return out
}
func DeployRolloutObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.DeployRolloutObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployRolloutObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	// MISSING: TargetID
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	// MISSING: Etag
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func DeployRolloutObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployRolloutObservedState) *pb.Rollout {
	if in == nil {
		return nil
	}
	out := &pb.Rollout{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	// MISSING: TargetID
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	// MISSING: Etag
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func DeployRolloutSpec_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.DeployRolloutSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployRolloutSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	// MISSING: TargetID
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	// MISSING: Etag
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func DeployRolloutSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployRolloutSpec) *pb.Rollout {
	if in == nil {
		return nil
	}
	out := &pb.Rollout{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	// MISSING: TargetID
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	// MISSING: Etag
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func DeploymentJobs_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentJobs) *krm.DeploymentJobs {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentJobs{}
	// MISSING: DeployJob
	// MISSING: VerifyJob
	// MISSING: PredeployJob
	// MISSING: PostdeployJob
	return out
}
func DeploymentJobs_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentJobs) *pb.DeploymentJobs {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentJobs{}
	// MISSING: DeployJob
	// MISSING: VerifyJob
	// MISSING: PredeployJob
	// MISSING: PostdeployJob
	return out
}
func DeploymentJobsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentJobs) *krm.DeploymentJobsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentJobsObservedState{}
	out.DeployJob = Job_FromProto(mapCtx, in.GetDeployJob())
	out.VerifyJob = Job_FromProto(mapCtx, in.GetVerifyJob())
	out.PredeployJob = Job_FromProto(mapCtx, in.GetPredeployJob())
	out.PostdeployJob = Job_FromProto(mapCtx, in.GetPostdeployJob())
	return out
}
func DeploymentJobsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentJobsObservedState) *pb.DeploymentJobs {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentJobs{}
	out.DeployJob = Job_ToProto(mapCtx, in.DeployJob)
	out.VerifyJob = Job_ToProto(mapCtx, in.VerifyJob)
	out.PredeployJob = Job_ToProto(mapCtx, in.PredeployJob)
	out.PostdeployJob = Job_ToProto(mapCtx, in.PostdeployJob)
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	// MISSING: ID
	// MISSING: State
	// MISSING: SkipMessage
	// MISSING: JobRun
	// MISSING: DeployJob
	// MISSING: VerifyJob
	// MISSING: PredeployJob
	// MISSING: PostdeployJob
	// MISSING: CreateChildRolloutJob
	// MISSING: AdvanceChildRolloutJob
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: ID
	// MISSING: State
	// MISSING: SkipMessage
	// MISSING: JobRun
	// MISSING: DeployJob
	// MISSING: VerifyJob
	// MISSING: PredeployJob
	// MISSING: PostdeployJob
	// MISSING: CreateChildRolloutJob
	// MISSING: AdvanceChildRolloutJob
	return out
}
func JobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.JobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SkipMessage = direct.LazyPtr(in.GetSkipMessage())
	out.JobRun = direct.LazyPtr(in.GetJobRun())
	out.DeployJob = DeployJob_FromProto(mapCtx, in.GetDeployJob())
	out.VerifyJob = VerifyJob_FromProto(mapCtx, in.GetVerifyJob())
	out.PredeployJob = PredeployJob_FromProto(mapCtx, in.GetPredeployJob())
	out.PostdeployJob = PostdeployJob_FromProto(mapCtx, in.GetPostdeployJob())
	out.CreateChildRolloutJob = CreateChildRolloutJob_FromProto(mapCtx, in.GetCreateChildRolloutJob())
	out.AdvanceChildRolloutJob = AdvanceChildRolloutJob_FromProto(mapCtx, in.GetAdvanceChildRolloutJob())
	return out
}
func JobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Id = direct.ValueOf(in.ID)
	out.State = direct.Enum_ToProto[pb.Job_State](mapCtx, in.State)
	out.SkipMessage = direct.ValueOf(in.SkipMessage)
	out.JobRun = direct.ValueOf(in.JobRun)
	if oneof := DeployJob_ToProto(mapCtx, in.DeployJob); oneof != nil {
		out.JobType = &pb.Job_DeployJob{DeployJob: oneof}
	}
	if oneof := VerifyJob_ToProto(mapCtx, in.VerifyJob); oneof != nil {
		out.JobType = &pb.Job_VerifyJob{VerifyJob: oneof}
	}
	if oneof := PredeployJob_ToProto(mapCtx, in.PredeployJob); oneof != nil {
		out.JobType = &pb.Job_PredeployJob{PredeployJob: oneof}
	}
	if oneof := PostdeployJob_ToProto(mapCtx, in.PostdeployJob); oneof != nil {
		out.JobType = &pb.Job_PostdeployJob{PostdeployJob: oneof}
	}
	if oneof := CreateChildRolloutJob_ToProto(mapCtx, in.CreateChildRolloutJob); oneof != nil {
		out.JobType = &pb.Job_CreateChildRolloutJob{CreateChildRolloutJob: oneof}
	}
	if oneof := AdvanceChildRolloutJob_ToProto(mapCtx, in.AdvanceChildRolloutJob); oneof != nil {
		out.JobType = &pb.Job_AdvanceChildRolloutJob{AdvanceChildRolloutJob: oneof}
	}
	return out
}
func Phase_FromProto(mapCtx *direct.MapContext, in *pb.Phase) *krm.Phase {
	if in == nil {
		return nil
	}
	out := &krm.Phase{}
	// MISSING: ID
	// MISSING: State
	// MISSING: SkipMessage
	// MISSING: DeploymentJobs
	// MISSING: ChildRolloutJobs
	return out
}
func Phase_ToProto(mapCtx *direct.MapContext, in *krm.Phase) *pb.Phase {
	if in == nil {
		return nil
	}
	out := &pb.Phase{}
	// MISSING: ID
	// MISSING: State
	// MISSING: SkipMessage
	// MISSING: DeploymentJobs
	// MISSING: ChildRolloutJobs
	return out
}
func PhaseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Phase) *krm.PhaseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PhaseObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SkipMessage = direct.LazyPtr(in.GetSkipMessage())
	out.DeploymentJobs = DeploymentJobs_FromProto(mapCtx, in.GetDeploymentJobs())
	out.ChildRolloutJobs = ChildRolloutJobs_FromProto(mapCtx, in.GetChildRolloutJobs())
	return out
}
func PhaseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PhaseObservedState) *pb.Phase {
	if in == nil {
		return nil
	}
	out := &pb.Phase{}
	out.Id = direct.ValueOf(in.ID)
	out.State = direct.Enum_ToProto[pb.Phase_State](mapCtx, in.State)
	out.SkipMessage = direct.ValueOf(in.SkipMessage)
	if oneof := DeploymentJobs_ToProto(mapCtx, in.DeploymentJobs); oneof != nil {
		out.Jobs = &pb.Phase_DeploymentJobs{DeploymentJobs: oneof}
	}
	if oneof := ChildRolloutJobs_ToProto(mapCtx, in.ChildRolloutJobs); oneof != nil {
		out.Jobs = &pb.Phase_ChildRolloutJobs{ChildRolloutJobs: oneof}
	}
	return out
}
func PostdeployJob_FromProto(mapCtx *direct.MapContext, in *pb.PostdeployJob) *krm.PostdeployJob {
	if in == nil {
		return nil
	}
	out := &krm.PostdeployJob{}
	// MISSING: Actions
	return out
}
func PostdeployJob_ToProto(mapCtx *direct.MapContext, in *krm.PostdeployJob) *pb.PostdeployJob {
	if in == nil {
		return nil
	}
	out := &pb.PostdeployJob{}
	// MISSING: Actions
	return out
}
func PostdeployJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PostdeployJob) *krm.PostdeployJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PostdeployJobObservedState{}
	out.Actions = in.Actions
	return out
}
func PostdeployJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PostdeployJobObservedState) *pb.PostdeployJob {
	if in == nil {
		return nil
	}
	out := &pb.PostdeployJob{}
	out.Actions = in.Actions
	return out
}
func PredeployJob_FromProto(mapCtx *direct.MapContext, in *pb.PredeployJob) *krm.PredeployJob {
	if in == nil {
		return nil
	}
	out := &krm.PredeployJob{}
	// MISSING: Actions
	return out
}
func PredeployJob_ToProto(mapCtx *direct.MapContext, in *krm.PredeployJob) *pb.PredeployJob {
	if in == nil {
		return nil
	}
	out := &pb.PredeployJob{}
	// MISSING: Actions
	return out
}
func PredeployJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PredeployJob) *krm.PredeployJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PredeployJobObservedState{}
	out.Actions = in.Actions
	return out
}
func PredeployJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PredeployJobObservedState) *pb.PredeployJob {
	if in == nil {
		return nil
	}
	out := &pb.PredeployJob{}
	out.Actions = in.Actions
	return out
}
func Rollout_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.Rollout {
	if in == nil {
		return nil
	}
	out := &krm.Rollout{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	out.TargetID = direct.LazyPtr(in.GetTargetId())
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func Rollout_ToProto(mapCtx *direct.MapContext, in *krm.Rollout) *pb.Rollout {
	if in == nil {
		return nil
	}
	out := &pb.Rollout{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: ApproveTime
	// MISSING: EnqueueTime
	// MISSING: DeployStartTime
	// MISSING: DeployEndTime
	out.TargetId = direct.ValueOf(in.TargetID)
	// MISSING: ApprovalState
	// MISSING: State
	// MISSING: FailureReason
	// MISSING: DeployingBuild
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: DeployFailureCause
	// MISSING: Phases
	// MISSING: Metadata
	// MISSING: ControllerRollout
	// MISSING: RollbackOfRollout
	// MISSING: RolledBackByRollouts
	// MISSING: ActiveRepairAutomationRun
	return out
}
func RolloutObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.RolloutObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RolloutObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ApproveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetApproveTime())
	out.EnqueueTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnqueueTime())
	out.DeployStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeployStartTime())
	out.DeployEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeployEndTime())
	// MISSING: TargetID
	out.ApprovalState = direct.Enum_FromProto(mapCtx, in.GetApprovalState())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FailureReason = direct.LazyPtr(in.GetFailureReason())
	out.DeployingBuild = direct.LazyPtr(in.GetDeployingBuild())
	// MISSING: Etag
	out.DeployFailureCause = direct.Enum_FromProto(mapCtx, in.GetDeployFailureCause())
	out.Phases = direct.Slice_FromProto(mapCtx, in.Phases, Phase_FromProto)
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	out.ControllerRollout = direct.LazyPtr(in.GetControllerRollout())
	out.RollbackOfRollout = direct.LazyPtr(in.GetRollbackOfRollout())
	out.RolledBackByRollouts = in.RolledBackByRollouts
	out.ActiveRepairAutomationRun = direct.LazyPtr(in.GetActiveRepairAutomationRun())
	return out
}
func RolloutObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RolloutObservedState) *pb.Rollout {
	if in == nil {
		return nil
	}
	out := &pb.Rollout{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ApproveTime = direct.StringTimestamp_ToProto(mapCtx, in.ApproveTime)
	out.EnqueueTime = direct.StringTimestamp_ToProto(mapCtx, in.EnqueueTime)
	out.DeployStartTime = direct.StringTimestamp_ToProto(mapCtx, in.DeployStartTime)
	out.DeployEndTime = direct.StringTimestamp_ToProto(mapCtx, in.DeployEndTime)
	// MISSING: TargetID
	out.ApprovalState = direct.Enum_ToProto[pb.Rollout_ApprovalState](mapCtx, in.ApprovalState)
	out.State = direct.Enum_ToProto[pb.Rollout_State](mapCtx, in.State)
	out.FailureReason = direct.ValueOf(in.FailureReason)
	out.DeployingBuild = direct.ValueOf(in.DeployingBuild)
	// MISSING: Etag
	out.DeployFailureCause = direct.Enum_ToProto[pb.Rollout_FailureCause](mapCtx, in.DeployFailureCause)
	out.Phases = direct.Slice_ToProto(mapCtx, in.Phases, Phase_ToProto)
	out.Metadata = Metadata_ToProto(mapCtx, in.Metadata)
	out.ControllerRollout = direct.ValueOf(in.ControllerRollout)
	out.RollbackOfRollout = direct.ValueOf(in.RollbackOfRollout)
	out.RolledBackByRollouts = in.RolledBackByRollouts
	out.ActiveRepairAutomationRun = direct.ValueOf(in.ActiveRepairAutomationRun)
	return out
}
func VerifyJob_FromProto(mapCtx *direct.MapContext, in *pb.VerifyJob) *krm.VerifyJob {
	if in == nil {
		return nil
	}
	out := &krm.VerifyJob{}
	return out
}
func VerifyJob_ToProto(mapCtx *direct.MapContext, in *krm.VerifyJob) *pb.VerifyJob {
	if in == nil {
		return nil
	}
	out := &pb.VerifyJob{}
	return out
}
