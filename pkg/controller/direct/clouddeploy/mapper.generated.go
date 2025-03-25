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
// krm.group: clouddeploy.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.deploy.v1

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	cloudbuildv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DeliveryPipelineAttribute_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipelineAttribute) *krm.DeliveryPipelineAttribute {
	if in == nil {
		return nil
	}
	out := &krm.DeliveryPipelineAttribute{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Labels = in.Labels
	return out
}
func DeliveryPipelineAttribute_ToProto(mapCtx *direct.MapContext, in *krm.DeliveryPipelineAttribute) *pb.DeliveryPipelineAttribute {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipelineAttribute{}
	out.Id = direct.ValueOf(in.ID)
	out.Labels = in.Labels
	return out
}
func DeliveryPipelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeliveryPipelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeliveryPipelineObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Condition = PipelineCondition_FromProto(mapCtx, in.GetCondition())
	return out
}
func DeliveryPipelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeliveryPipelineObservedState) *pb.DeliveryPipeline {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipeline{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Condition = PipelineCondition_ToProto(mapCtx, in.Condition)
	return out
}
func DeliveryPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeliveryPipelineSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeliveryPipelineSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	// MISSING: Labels
	out.SerialPipeline = SerialPipeline_FromProto(mapCtx, in.GetSerialPipeline())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	return out
}
func DeliveryPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeliveryPipelineSpec) *pb.DeliveryPipeline {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipeline{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	// MISSING: Labels
	if oneof := SerialPipeline_ToProto(mapCtx, in.SerialPipeline); oneof != nil {
		out.Pipeline = &pb.DeliveryPipeline_SerialPipeline{SerialPipeline: oneof}
	}
	out.Etag = direct.ValueOf(in.Etag)
	out.Suspended = direct.ValueOf(in.Suspended)
	return out
}
func DeployDeployPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicy) *krm.DeployDeployPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployDeployPolicyObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployDeployPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployDeployPolicyObservedState) *pb.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployDeployPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicy) *krm.DeployDeployPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployDeployPolicySpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployDeployPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployDeployPolicySpec) *pb.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicy{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployPolicy_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicy) *krm.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &krm.DeployPolicy{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	out.Selectors = direct.Slice_FromProto(mapCtx, in.Selectors, DeployPolicyResourceSelector_FromProto)
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicyRule_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func DeployPolicy_ToProto(mapCtx *direct.MapContext, in *krm.DeployPolicy) *pb.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicy{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Suspended = direct.ValueOf(in.Suspended)
	out.Selectors = direct.Slice_ToProto(mapCtx, in.Selectors, DeployPolicyResourceSelector_ToProto)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicyRule_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func DeployPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicy) *krm.DeployPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployPolicyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployPolicyObservedState) *pb.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Suspended
	// MISSING: Selectors
	// MISSING: Rules
	// MISSING: Etag
	return out
}
func DeployPolicyResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicyResourceSelector) *krm.DeployPolicyResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.DeployPolicyResourceSelector{}
	out.DeliveryPipeline = DeliveryPipelineAttribute_FromProto(mapCtx, in.GetDeliveryPipeline())
	out.Target = TargetAttribute_FromProto(mapCtx, in.GetTarget())
	return out
}
func DeployPolicyResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.DeployPolicyResourceSelector) *pb.DeployPolicyResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicyResourceSelector{}
	out.DeliveryPipeline = DeliveryPipelineAttribute_ToProto(mapCtx, in.DeliveryPipeline)
	out.Target = TargetAttribute_ToProto(mapCtx, in.Target)
	return out
}
func OneTimeWindow_FromProto(mapCtx *direct.MapContext, in *pb.OneTimeWindow) *krm.OneTimeWindow {
	if in == nil {
		return nil
	}
	out := &krm.OneTimeWindow{}
	out.StartDate = Date_FromProto(mapCtx, in.GetStartDate())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	out.EndDate = Date_FromProto(mapCtx, in.GetEndDate())
	out.EndTime = TimeOfDay_FromProto(mapCtx, in.GetEndTime())
	return out
}
func OneTimeWindow_ToProto(mapCtx *direct.MapContext, in *krm.OneTimeWindow) *pb.OneTimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.OneTimeWindow{}
	out.StartDate = Date_ToProto(mapCtx, in.StartDate)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	out.EndDate = Date_ToProto(mapCtx, in.EndDate)
	out.EndTime = TimeOfDay_ToProto(mapCtx, in.EndTime)
	return out
}
func PolicyRule_FromProto(mapCtx *direct.MapContext, in *pb.PolicyRule) *krm.PolicyRule {
	if in == nil {
		return nil
	}
	out := &krm.PolicyRule{}
	out.RolloutRestriction = RolloutRestriction_FromProto(mapCtx, in.GetRolloutRestriction())
	return out
}
func PolicyRule_ToProto(mapCtx *direct.MapContext, in *krm.PolicyRule) *pb.PolicyRule {
	if in == nil {
		return nil
	}
	out := &pb.PolicyRule{}
	if oneof := RolloutRestriction_ToProto(mapCtx, in.RolloutRestriction); oneof != nil {
		out.Rule = &pb.PolicyRule_RolloutRestriction{RolloutRestriction: oneof}
	}
	return out
}
func RolloutRestriction_FromProto(mapCtx *direct.MapContext, in *pb.RolloutRestriction) *krm.RolloutRestriction {
	if in == nil {
		return nil
	}
	out := &krm.RolloutRestriction{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Invokers = direct.EnumSlice_FromProto(mapCtx, in.Invokers)
	out.Actions = direct.EnumSlice_FromProto(mapCtx, in.Actions)
	out.TimeWindows = TimeWindows_FromProto(mapCtx, in.GetTimeWindows())
	return out
}
func RolloutRestriction_ToProto(mapCtx *direct.MapContext, in *krm.RolloutRestriction) *pb.RolloutRestriction {
	if in == nil {
		return nil
	}
	out := &pb.RolloutRestriction{}
	out.Id = direct.ValueOf(in.ID)
	out.Invokers = direct.EnumSlice_ToProto[pb.DeployPolicy_Invoker](mapCtx, in.Invokers)
	out.Actions = direct.EnumSlice_ToProto[pb.RolloutRestriction_RolloutActions](mapCtx, in.Actions)
	out.TimeWindows = TimeWindows_ToProto(mapCtx, in.TimeWindows)
	return out
}
func TargetAttribute_FromProto(mapCtx *direct.MapContext, in *pb.TargetAttribute) *krm.TargetAttribute {
	if in == nil {
		return nil
	}
	out := &krm.TargetAttribute{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Labels = in.Labels
	return out
}
func TargetAttribute_ToProto(mapCtx *direct.MapContext, in *krm.TargetAttribute) *pb.TargetAttribute {
	if in == nil {
		return nil
	}
	out := &pb.TargetAttribute{}
	out.Id = direct.ValueOf(in.ID)
	out.Labels = in.Labels
	return out
}
func TimeWindows_FromProto(mapCtx *direct.MapContext, in *pb.TimeWindows) *krm.TimeWindows {
	if in == nil {
		return nil
	}
	out := &krm.TimeWindows{}
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.OneTimeWindows = direct.Slice_FromProto(mapCtx, in.OneTimeWindows, OneTimeWindow_FromProto)
	out.WeeklyWindows = direct.Slice_FromProto(mapCtx, in.WeeklyWindows, WeeklyWindow_FromProto)
	return out
}
func TimeWindows_ToProto(mapCtx *direct.MapContext, in *krm.TimeWindows) *pb.TimeWindows {
	if in == nil {
		return nil
	}
	out := &pb.TimeWindows{}
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.OneTimeWindows = direct.Slice_ToProto(mapCtx, in.OneTimeWindows, OneTimeWindow_ToProto)
	out.WeeklyWindows = direct.Slice_ToProto(mapCtx, in.WeeklyWindows, WeeklyWindow_ToProto)
	return out
}
func WeeklyWindow_FromProto(mapCtx *direct.MapContext, in *pb.WeeklyWindow) *krm.WeeklyWindow {
	if in == nil {
		return nil
	}
	out := &krm.WeeklyWindow{}
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = TimeOfDay_FromProto(mapCtx, in.GetEndTime())
	return out
}
func WeeklyWindow_ToProto(mapCtx *direct.MapContext, in *krm.WeeklyWindow) *pb.WeeklyWindow {
	if in == nil {
		return nil
	}
	out := &pb.WeeklyWindow{}
	out.DaysOfWeek = direct.EnumSlice_ToProto[pb.DayOfWeek](mapCtx, in.DaysOfWeek)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	out.EndTime = TimeOfDay_ToProto(mapCtx, in.EndTime)
	return out
}

func CustomTargetSkaffoldActions_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetSkaffoldActions) *krm.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.LazyPtr(in.GetRenderAction())
	out.DeployAction = direct.LazyPtr(in.GetDeployAction())
	out.IncludeSkaffoldModules = direct.Slice_FromProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_FromProto)
	return out
}
func CustomTargetSkaffoldActions_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetSkaffoldActions) *pb.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.ValueOf(in.RenderAction)
	out.DeployAction = direct.ValueOf(in.DeployAction)
	out.IncludeSkaffoldModules = direct.Slice_ToProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_ToProto)
	return out
}
func CustomTargetType_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetTypeSpec{}
	//out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	//out.Annotations = in.Annotations
	//out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	//out.Etag = direct.LazyPtr(in.GetEtag())
	out.CustomActions = CustomTargetSkaffoldActions_FromProto(mapCtx, in.GetCustomActions())
	return out
}
func CustomTargetType_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetTypeSpec) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// out.Name = direct.ValueOf(in.Name)
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// out.Etag = direct.ValueOf(in.Etag)
	if oneof := CustomTargetSkaffoldActions_ToProto(mapCtx, in.CustomActions); oneof != nil {
		out.Definition = &pb.CustomTargetType_CustomActions{CustomActions: oneof}
	}
	return out
}
func CustomTargetTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetTypeObservedState{}
	// MISSING: Name
	out.CustomTargetTypeID = direct.LazyPtr(in.GetCustomTargetTypeId())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func CustomTargetTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetTypeObservedState) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// MISSING: Name
	out.CustomTargetTypeId = direct.ValueOf(in.CustomTargetTypeID)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}

func SkaffoldModules_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules) *krm.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules{}
	out.Configs = in.Configs
	out.Git = SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx, in.GetGit())
	out.GoogleCloudStorage = SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx, in.GetGoogleCloudStorage())
	out.GoogleCloudBuildRepo = SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx, in.GetGoogleCloudBuildRepo())
	return out
}
func SkaffoldModules_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules) *pb.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules{}
	out.Configs = in.Configs
	if oneof := SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx, in.Git); oneof != nil {
		out.Source = &pb.SkaffoldModules_Git{Git: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx, in.GoogleCloudStorage); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudStorage{GoogleCloudStorage: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx, in.GoogleCloudBuildRepo); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudBuildRepo{GoogleCloudBuildRepo: oneof}
	}
	return out
}
func SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCSSource) *krm.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGCSSource) *pb.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.ValueOf(in.Source)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCBRepoSource) *krm.SkaffoldModules_SkaffoldGcbRepoSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGcbRepoSource{}
	if in.GetRepository() != "" {
		out.RepositoryRef = &cloudbuildv1alpha1.RepositoryRef{
			External: in.GetRepository(),
		}
	}

	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGcbRepoSource) *pb.SkaffoldModules_SkaffoldGCBRepoSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCBRepoSource{}
	if in.RepositoryRef != nil {
		out.Repository = in.RepositoryRef.External
	}
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
func SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGitSource) *krm.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.LazyPtr(in.GetRepo())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGitSource) *pb.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.ValueOf(in.Repo)
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
