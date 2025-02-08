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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AdvanceChildRolloutJobRun_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceChildRolloutJobRun) *krm.AdvanceChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceChildRolloutJobRun{}
	// MISSING: Rollout
	// MISSING: RolloutPhaseID
	return out
}
func AdvanceChildRolloutJobRun_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceChildRolloutJobRun) *pb.AdvanceChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceChildRolloutJobRun{}
	// MISSING: Rollout
	// MISSING: RolloutPhaseID
	return out
}
func AdvanceChildRolloutJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceChildRolloutJobRun) *krm.AdvanceChildRolloutJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceChildRolloutJobRunObservedState{}
	out.Rollout = direct.LazyPtr(in.GetRollout())
	out.RolloutPhaseID = direct.LazyPtr(in.GetRolloutPhaseId())
	return out
}
func AdvanceChildRolloutJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceChildRolloutJobRunObservedState) *pb.AdvanceChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceChildRolloutJobRun{}
	out.Rollout = direct.ValueOf(in.Rollout)
	out.RolloutPhaseId = direct.ValueOf(in.RolloutPhaseID)
	return out
}
func CreateChildRolloutJobRun_FromProto(mapCtx *direct.MapContext, in *pb.CreateChildRolloutJobRun) *krm.CreateChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &krm.CreateChildRolloutJobRun{}
	// MISSING: Rollout
	// MISSING: RolloutPhaseID
	return out
}
func CreateChildRolloutJobRun_ToProto(mapCtx *direct.MapContext, in *krm.CreateChildRolloutJobRun) *pb.CreateChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &pb.CreateChildRolloutJobRun{}
	// MISSING: Rollout
	// MISSING: RolloutPhaseID
	return out
}
func CreateChildRolloutJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CreateChildRolloutJobRun) *krm.CreateChildRolloutJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CreateChildRolloutJobRunObservedState{}
	out.Rollout = direct.LazyPtr(in.GetRollout())
	out.RolloutPhaseID = direct.LazyPtr(in.GetRolloutPhaseId())
	return out
}
func CreateChildRolloutJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CreateChildRolloutJobRunObservedState) *pb.CreateChildRolloutJobRun {
	if in == nil {
		return nil
	}
	out := &pb.CreateChildRolloutJobRun{}
	out.Rollout = direct.ValueOf(in.Rollout)
	out.RolloutPhaseId = direct.ValueOf(in.RolloutPhaseID)
	return out
}
func DeployArtifact_FromProto(mapCtx *direct.MapContext, in *pb.DeployArtifact) *krm.DeployArtifact {
	if in == nil {
		return nil
	}
	out := &krm.DeployArtifact{}
	// MISSING: ArtifactURI
	// MISSING: ManifestPaths
	return out
}
func DeployArtifact_ToProto(mapCtx *direct.MapContext, in *krm.DeployArtifact) *pb.DeployArtifact {
	if in == nil {
		return nil
	}
	out := &pb.DeployArtifact{}
	// MISSING: ArtifactURI
	// MISSING: ManifestPaths
	return out
}
func DeployArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployArtifact) *krm.DeployArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployArtifactObservedState{}
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.ManifestPaths = in.ManifestPaths
	return out
}
func DeployArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployArtifactObservedState) *pb.DeployArtifact {
	if in == nil {
		return nil
	}
	out := &pb.DeployArtifact{}
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.ManifestPaths = in.ManifestPaths
	return out
}
func DeployJobRun_FromProto(mapCtx *direct.MapContext, in *pb.DeployJobRun) *krm.DeployJobRun {
	if in == nil {
		return nil
	}
	out := &krm.DeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	// MISSING: Metadata
	// MISSING: Artifact
	return out
}
func DeployJobRun_ToProto(mapCtx *direct.MapContext, in *krm.DeployJobRun) *pb.DeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.DeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	// MISSING: Metadata
	// MISSING: Artifact
	return out
}
func DeployJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobRun) *krm.DeployJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployJobRunObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func DeployJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployJobRunObservedState) *pb.JobRun {
	if in == nil {
		return nil
	}
	out := &pb.JobRun{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func DeployJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployJobRun) *krm.DeployJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployJobRunObservedState{}
	out.Build = direct.LazyPtr(in.GetBuild())
	out.FailureCause = direct.Enum_FromProto(mapCtx, in.GetFailureCause())
	out.FailureMessage = direct.LazyPtr(in.GetFailureMessage())
	out.Metadata = DeployJobRunMetadata_FromProto(mapCtx, in.GetMetadata())
	out.Artifact = DeployArtifact_FromProto(mapCtx, in.GetArtifact())
	return out
}
func DeployJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployJobRunObservedState) *pb.DeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.DeployJobRun{}
	out.Build = direct.ValueOf(in.Build)
	out.FailureCause = direct.Enum_ToProto[pb.DeployJobRun_FailureCause](mapCtx, in.FailureCause)
	out.FailureMessage = direct.ValueOf(in.FailureMessage)
	out.Metadata = DeployJobRunMetadata_ToProto(mapCtx, in.Metadata)
	out.Artifact = DeployArtifact_ToProto(mapCtx, in.Artifact)
	return out
}
func DeployJobRunSpec_FromProto(mapCtx *direct.MapContext, in *pb.JobRun) *krm.DeployJobRunSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployJobRunSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func DeployJobRunSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployJobRunSpec) *pb.JobRun {
	if in == nil {
		return nil
	}
	out := &pb.JobRun{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func JobRun_FromProto(mapCtx *direct.MapContext, in *pb.JobRun) *krm.JobRun {
	if in == nil {
		return nil
	}
	out := &krm.JobRun{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func JobRun_ToProto(mapCtx *direct.MapContext, in *krm.JobRun) *pb.JobRun {
	if in == nil {
		return nil
	}
	out := &pb.JobRun{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: PhaseID
	// MISSING: JobID
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	// MISSING: DeployJobRun
	// MISSING: VerifyJobRun
	// MISSING: PredeployJobRun
	// MISSING: PostdeployJobRun
	// MISSING: CreateChildRolloutJobRun
	// MISSING: AdvanceChildRolloutJobRun
	// MISSING: Etag
	return out
}
func JobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobRun) *krm.JobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobRunObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.PhaseID = direct.LazyPtr(in.GetPhaseId())
	out.JobID = direct.LazyPtr(in.GetJobId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DeployJobRun = DeployJobRun_FromProto(mapCtx, in.GetDeployJobRun())
	out.VerifyJobRun = VerifyJobRun_FromProto(mapCtx, in.GetVerifyJobRun())
	out.PredeployJobRun = PredeployJobRun_FromProto(mapCtx, in.GetPredeployJobRun())
	out.PostdeployJobRun = PostdeployJobRun_FromProto(mapCtx, in.GetPostdeployJobRun())
	out.CreateChildRolloutJobRun = CreateChildRolloutJobRun_FromProto(mapCtx, in.GetCreateChildRolloutJobRun())
	out.AdvanceChildRolloutJobRun = AdvanceChildRolloutJobRun_FromProto(mapCtx, in.GetAdvanceChildRolloutJobRun())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func JobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobRunObservedState) *pb.JobRun {
	if in == nil {
		return nil
	}
	out := &pb.JobRun{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.PhaseId = direct.ValueOf(in.PhaseID)
	out.JobId = direct.ValueOf(in.JobID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.JobRun_State](mapCtx, in.State)
	if oneof := DeployJobRun_ToProto(mapCtx, in.DeployJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_DeployJobRun{DeployJobRun: oneof}
	}
	if oneof := VerifyJobRun_ToProto(mapCtx, in.VerifyJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_VerifyJobRun{VerifyJobRun: oneof}
	}
	if oneof := PredeployJobRun_ToProto(mapCtx, in.PredeployJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_PredeployJobRun{PredeployJobRun: oneof}
	}
	if oneof := PostdeployJobRun_ToProto(mapCtx, in.PostdeployJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_PostdeployJobRun{PostdeployJobRun: oneof}
	}
	if oneof := CreateChildRolloutJobRun_ToProto(mapCtx, in.CreateChildRolloutJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_CreateChildRolloutJobRun{CreateChildRolloutJobRun: oneof}
	}
	if oneof := AdvanceChildRolloutJobRun_ToProto(mapCtx, in.AdvanceChildRolloutJobRun); oneof != nil {
		out.JobRun = &pb.JobRun_AdvanceChildRolloutJobRun{AdvanceChildRolloutJobRun: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func PostdeployJobRun_FromProto(mapCtx *direct.MapContext, in *pb.PostdeployJobRun) *krm.PostdeployJobRun {
	if in == nil {
		return nil
	}
	out := &krm.PostdeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func PostdeployJobRun_ToProto(mapCtx *direct.MapContext, in *krm.PostdeployJobRun) *pb.PostdeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.PostdeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func PostdeployJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PostdeployJobRun) *krm.PostdeployJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PostdeployJobRunObservedState{}
	out.Build = direct.LazyPtr(in.GetBuild())
	out.FailureCause = direct.Enum_FromProto(mapCtx, in.GetFailureCause())
	out.FailureMessage = direct.LazyPtr(in.GetFailureMessage())
	return out
}
func PostdeployJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PostdeployJobRunObservedState) *pb.PostdeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.PostdeployJobRun{}
	out.Build = direct.ValueOf(in.Build)
	out.FailureCause = direct.Enum_ToProto[pb.PostdeployJobRun_FailureCause](mapCtx, in.FailureCause)
	out.FailureMessage = direct.ValueOf(in.FailureMessage)
	return out
}
func PredeployJobRun_FromProto(mapCtx *direct.MapContext, in *pb.PredeployJobRun) *krm.PredeployJobRun {
	if in == nil {
		return nil
	}
	out := &krm.PredeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func PredeployJobRun_ToProto(mapCtx *direct.MapContext, in *krm.PredeployJobRun) *pb.PredeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.PredeployJobRun{}
	// MISSING: Build
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func PredeployJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PredeployJobRun) *krm.PredeployJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PredeployJobRunObservedState{}
	out.Build = direct.LazyPtr(in.GetBuild())
	out.FailureCause = direct.Enum_FromProto(mapCtx, in.GetFailureCause())
	out.FailureMessage = direct.LazyPtr(in.GetFailureMessage())
	return out
}
func PredeployJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PredeployJobRunObservedState) *pb.PredeployJobRun {
	if in == nil {
		return nil
	}
	out := &pb.PredeployJobRun{}
	out.Build = direct.ValueOf(in.Build)
	out.FailureCause = direct.Enum_ToProto[pb.PredeployJobRun_FailureCause](mapCtx, in.FailureCause)
	out.FailureMessage = direct.ValueOf(in.FailureMessage)
	return out
}
func VerifyJobRun_FromProto(mapCtx *direct.MapContext, in *pb.VerifyJobRun) *krm.VerifyJobRun {
	if in == nil {
		return nil
	}
	out := &krm.VerifyJobRun{}
	// MISSING: Build
	// MISSING: ArtifactURI
	// MISSING: EventLogPath
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func VerifyJobRun_ToProto(mapCtx *direct.MapContext, in *krm.VerifyJobRun) *pb.VerifyJobRun {
	if in == nil {
		return nil
	}
	out := &pb.VerifyJobRun{}
	// MISSING: Build
	// MISSING: ArtifactURI
	// MISSING: EventLogPath
	// MISSING: FailureCause
	// MISSING: FailureMessage
	return out
}
func VerifyJobRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VerifyJobRun) *krm.VerifyJobRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VerifyJobRunObservedState{}
	out.Build = direct.LazyPtr(in.GetBuild())
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.EventLogPath = direct.LazyPtr(in.GetEventLogPath())
	out.FailureCause = direct.Enum_FromProto(mapCtx, in.GetFailureCause())
	out.FailureMessage = direct.LazyPtr(in.GetFailureMessage())
	return out
}
func VerifyJobRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VerifyJobRunObservedState) *pb.VerifyJobRun {
	if in == nil {
		return nil
	}
	out := &pb.VerifyJobRun{}
	out.Build = direct.ValueOf(in.Build)
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.EventLogPath = direct.ValueOf(in.EventLogPath)
	out.FailureCause = direct.Enum_ToProto[pb.VerifyJobRun_FailureCause](mapCtx, in.FailureCause)
	out.FailureMessage = direct.ValueOf(in.FailureMessage)
	return out
}
