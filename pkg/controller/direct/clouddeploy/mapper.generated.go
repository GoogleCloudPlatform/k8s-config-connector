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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
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

func DeployPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.DeployPolicy) *krm.DeployPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployPolicySpec{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	out.Selectors = direct.Slice_FromProto(mapCtx, in.Selectors, DeployPolicyResourceSelector_FromProto)
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicyRule_FromProto)
	//out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func DeployPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployPolicySpec) *pb.DeployPolicy {
	if in == nil {
		return nil
	}
	out := &pb.DeployPolicy{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Suspended = direct.ValueOf(in.Suspended)
	out.Selectors = direct.Slice_ToProto(mapCtx, in.Selectors, DeployPolicyResourceSelector_ToProto)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicyRule_ToProto)
	//out.Etag = direct.ValueOf(in.Etag)
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
	out.DaysOfWeek = direct.EnumSlice_ToProto[dayofweek.DayOfWeek](mapCtx, in.DaysOfWeek)
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

func Canary_FromProto(mapCtx *direct.MapContext, in *pb.Canary) *krm.Canary {
	if in == nil {
		return nil
	}
	out := &krm.Canary{}
	out.RuntimeConfig = RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.CanaryDeployment = CanaryDeployment_FromProto(mapCtx, in.GetCanaryDeployment())
	out.CustomCanaryDeployment = CustomCanaryDeployment_FromProto(mapCtx, in.GetCustomCanaryDeployment())
	return out
}
func Canary_ToProto(mapCtx *direct.MapContext, in *krm.Canary) *pb.Canary {
	if in == nil {
		return nil
	}
	out := &pb.Canary{}
	out.RuntimeConfig = RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	if oneof := CanaryDeployment_ToProto(mapCtx, in.CanaryDeployment); oneof != nil {
		out.Mode = &pb.Canary_CanaryDeployment{CanaryDeployment: oneof}
	}
	if oneof := CustomCanaryDeployment_ToProto(mapCtx, in.CustomCanaryDeployment); oneof != nil {
		out.Mode = &pb.Canary_CustomCanaryDeployment{CustomCanaryDeployment: oneof}
	}
	return out
}
func CanaryDeployment_FromProto(mapCtx *direct.MapContext, in *pb.CanaryDeployment) *krm.CanaryDeployment {
	if in == nil {
		return nil
	}
	out := &krm.CanaryDeployment{}
	out.Percentages = in.Percentages
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func CanaryDeployment_ToProto(mapCtx *direct.MapContext, in *krm.CanaryDeployment) *pb.CanaryDeployment {
	if in == nil {
		return nil
	}
	out := &pb.CanaryDeployment{}
	out.Percentages = in.Percentages
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}
func CloudRunConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunConfig) *krm.CloudRunConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudRunConfig{}
	out.AutomaticTrafficControl = direct.LazyPtr(in.GetAutomaticTrafficControl())
	out.CanaryRevisionTags = in.CanaryRevisionTags
	out.PriorRevisionTags = in.PriorRevisionTags
	out.StableRevisionTags = in.StableRevisionTags
	return out
}
func CloudRunConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudRunConfig) *pb.CloudRunConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudRunConfig{}
	out.AutomaticTrafficControl = direct.ValueOf(in.AutomaticTrafficControl)
	out.CanaryRevisionTags = in.CanaryRevisionTags
	out.PriorRevisionTags = in.PriorRevisionTags
	out.StableRevisionTags = in.StableRevisionTags
	return out
}
func CustomCanaryDeployment_FromProto(mapCtx *direct.MapContext, in *pb.CustomCanaryDeployment) *krm.CustomCanaryDeployment {
	if in == nil {
		return nil
	}
	out := &krm.CustomCanaryDeployment{}
	out.PhaseConfigs = direct.Slice_FromProto(mapCtx, in.PhaseConfigs, CustomCanaryDeployment_PhaseConfig_FromProto)
	return out
}
func CustomCanaryDeployment_ToProto(mapCtx *direct.MapContext, in *krm.CustomCanaryDeployment) *pb.CustomCanaryDeployment {
	if in == nil {
		return nil
	}
	out := &pb.CustomCanaryDeployment{}
	out.PhaseConfigs = direct.Slice_ToProto(mapCtx, in.PhaseConfigs, CustomCanaryDeployment_PhaseConfig_ToProto)
	return out
}
func CustomCanaryDeployment_PhaseConfig_FromProto(mapCtx *direct.MapContext, in *pb.CustomCanaryDeployment_PhaseConfig) *krm.CustomCanaryDeployment_PhaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.CustomCanaryDeployment_PhaseConfig{}
	out.PhaseID = direct.LazyPtr(in.GetPhaseId())
	out.Percentage = direct.LazyPtr(in.GetPercentage())
	out.Profiles = in.Profiles
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func CustomCanaryDeployment_PhaseConfig_ToProto(mapCtx *direct.MapContext, in *krm.CustomCanaryDeployment_PhaseConfig) *pb.CustomCanaryDeployment_PhaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomCanaryDeployment_PhaseConfig{}
	out.PhaseId = direct.ValueOf(in.PhaseID)
	out.Percentage = direct.ValueOf(in.Percentage)
	out.Profiles = in.Profiles
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}

func DeployParameters_FromProto(mapCtx *direct.MapContext, in *pb.DeployParameters) *krm.DeployParameters {
	if in == nil {
		return nil
	}
	out := &krm.DeployParameters{}
	out.Values = in.Values
	out.MatchTargetLabels = in.MatchTargetLabels
	return out
}
func DeployParameters_ToProto(mapCtx *direct.MapContext, in *krm.DeployParameters) *pb.DeployParameters {
	if in == nil {
		return nil
	}
	out := &pb.DeployParameters{}
	out.Values = in.Values
	out.MatchTargetLabels = in.MatchTargetLabels
	return out
}
func KubernetesConfig_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig) *krm.KubernetesConfig {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig{}
	out.GatewayServiceMesh = KubernetesConfig_GatewayServiceMesh_FromProto(mapCtx, in.GetGatewayServiceMesh())
	out.ServiceNetworking = KubernetesConfig_ServiceNetworking_FromProto(mapCtx, in.GetServiceNetworking())
	return out
}
func KubernetesConfig_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig) *pb.KubernetesConfig {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig{}
	if oneof := KubernetesConfig_GatewayServiceMesh_ToProto(mapCtx, in.GatewayServiceMesh); oneof != nil {
		out.ServiceDefinition = &pb.KubernetesConfig_GatewayServiceMesh_{GatewayServiceMesh: oneof}
	}
	if oneof := KubernetesConfig_ServiceNetworking_ToProto(mapCtx, in.ServiceNetworking); oneof != nil {
		out.ServiceDefinition = &pb.KubernetesConfig_ServiceNetworking_{ServiceNetworking: oneof}
	}
	return out
}
func KubernetesConfig_GatewayServiceMesh_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_GatewayServiceMesh) *krm.KubernetesConfig_GatewayServiceMesh {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_GatewayServiceMesh{}
	out.HTTPRoute = direct.LazyPtr(in.GetHttpRoute())
	out.Service = direct.LazyPtr(in.GetService())
	out.Deployment = direct.LazyPtr(in.GetDeployment())
	out.RouteUpdateWaitTime = direct.StringDuration_FromProto(mapCtx, in.GetRouteUpdateWaitTime())
	out.StableCutbackDuration = direct.StringDuration_FromProto(mapCtx, in.GetStableCutbackDuration())
	out.PodSelectorLabel = direct.LazyPtr(in.GetPodSelectorLabel())
	out.RouteDestinations = KubernetesConfig_GatewayServiceMesh_RouteDestinations_FromProto(mapCtx, in.GetRouteDestinations())
	return out
}
func KubernetesConfig_GatewayServiceMesh_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_GatewayServiceMesh) *pb.KubernetesConfig_GatewayServiceMesh {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_GatewayServiceMesh{}
	out.HttpRoute = direct.ValueOf(in.HTTPRoute)
	out.Service = direct.ValueOf(in.Service)
	out.Deployment = direct.ValueOf(in.Deployment)
	out.RouteUpdateWaitTime = direct.StringDuration_ToProto(mapCtx, in.RouteUpdateWaitTime)
	out.StableCutbackDuration = direct.StringDuration_ToProto(mapCtx, in.StableCutbackDuration)
	out.PodSelectorLabel = direct.ValueOf(in.PodSelectorLabel)
	out.RouteDestinations = KubernetesConfig_GatewayServiceMesh_RouteDestinations_ToProto(mapCtx, in.RouteDestinations)
	return out
}
func KubernetesConfig_GatewayServiceMesh_RouteDestinations_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations) *krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations{}
	out.DestinationIDs = in.DestinationIds
	out.PropagateService = direct.LazyPtr(in.GetPropagateService())
	return out
}
func KubernetesConfig_GatewayServiceMesh_RouteDestinations_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations) *pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations{}
	out.DestinationIds = in.DestinationIDs
	out.PropagateService = direct.ValueOf(in.PropagateService)
	return out
}
func KubernetesConfig_ServiceNetworking_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_ServiceNetworking) *krm.KubernetesConfig_ServiceNetworking {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_ServiceNetworking{}
	out.Service = direct.LazyPtr(in.GetService())
	out.Deployment = direct.LazyPtr(in.GetDeployment())
	out.DisablePodOverprovisioning = direct.LazyPtr(in.GetDisablePodOverprovisioning())
	out.PodSelectorLabel = direct.LazyPtr(in.GetPodSelectorLabel())
	return out
}
func KubernetesConfig_ServiceNetworking_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_ServiceNetworking) *pb.KubernetesConfig_ServiceNetworking {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_ServiceNetworking{}
	out.Service = direct.ValueOf(in.Service)
	out.Deployment = direct.ValueOf(in.Deployment)
	out.DisablePodOverprovisioning = direct.ValueOf(in.DisablePodOverprovisioning)
	out.PodSelectorLabel = direct.ValueOf(in.PodSelectorLabel)
	return out
}
func PipelineCondition_FromProto(mapCtx *direct.MapContext, in *pb.PipelineCondition) *krm.PipelineCondition {
	if in == nil {
		return nil
	}
	out := &krm.PipelineCondition{}
	out.PipelineReadyCondition = PipelineReadyCondition_FromProto(mapCtx, in.GetPipelineReadyCondition())
	out.TargetsPresentCondition = TargetsPresentCondition_FromProto(mapCtx, in.GetTargetsPresentCondition())
	out.TargetsTypeCondition = TargetsTypeCondition_FromProto(mapCtx, in.GetTargetsTypeCondition())
	return out
}
func PipelineCondition_ToProto(mapCtx *direct.MapContext, in *krm.PipelineCondition) *pb.PipelineCondition {
	if in == nil {
		return nil
	}
	out := &pb.PipelineCondition{}
	out.PipelineReadyCondition = PipelineReadyCondition_ToProto(mapCtx, in.PipelineReadyCondition)
	out.TargetsPresentCondition = TargetsPresentCondition_ToProto(mapCtx, in.TargetsPresentCondition)
	out.TargetsTypeCondition = TargetsTypeCondition_ToProto(mapCtx, in.TargetsTypeCondition)
	return out
}
func PipelineReadyCondition_FromProto(mapCtx *direct.MapContext, in *pb.PipelineReadyCondition) *krm.PipelineReadyCondition {
	if in == nil {
		return nil
	}
	out := &krm.PipelineReadyCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PipelineReadyCondition_ToProto(mapCtx *direct.MapContext, in *krm.PipelineReadyCondition) *pb.PipelineReadyCondition {
	if in == nil {
		return nil
	}
	out := &pb.PipelineReadyCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Postdeploy_FromProto(mapCtx *direct.MapContext, in *pb.Postdeploy) *krm.Postdeploy {
	if in == nil {
		return nil
	}
	out := &krm.Postdeploy{}
	out.Actions = in.Actions
	return out
}
func Postdeploy_ToProto(mapCtx *direct.MapContext, in *krm.Postdeploy) *pb.Postdeploy {
	if in == nil {
		return nil
	}
	out := &pb.Postdeploy{}
	out.Actions = in.Actions
	return out
}
func Predeploy_FromProto(mapCtx *direct.MapContext, in *pb.Predeploy) *krm.Predeploy {
	if in == nil {
		return nil
	}
	out := &krm.Predeploy{}
	out.Actions = in.Actions
	return out
}
func Predeploy_ToProto(mapCtx *direct.MapContext, in *krm.Predeploy) *pb.Predeploy {
	if in == nil {
		return nil
	}
	out := &pb.Predeploy{}
	out.Actions = in.Actions
	return out
}
func RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig{}
	out.Kubernetes = KubernetesConfig_FromProto(mapCtx, in.GetKubernetes())
	out.CloudRun = CloudRunConfig_FromProto(mapCtx, in.GetCloudRun())
	return out
}
func RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	if oneof := KubernetesConfig_ToProto(mapCtx, in.Kubernetes); oneof != nil {
		out.RuntimeConfig = &pb.RuntimeConfig_Kubernetes{Kubernetes: oneof}
	}
	if oneof := CloudRunConfig_ToProto(mapCtx, in.CloudRun); oneof != nil {
		out.RuntimeConfig = &pb.RuntimeConfig_CloudRun{CloudRun: oneof}
	}
	return out
}
func SerialPipeline_FromProto(mapCtx *direct.MapContext, in *pb.SerialPipeline) *krm.SerialPipeline {
	if in == nil {
		return nil
	}
	out := &krm.SerialPipeline{}
	out.Stages = direct.Slice_FromProto(mapCtx, in.Stages, Stage_FromProto)
	return out
}
func SerialPipeline_ToProto(mapCtx *direct.MapContext, in *krm.SerialPipeline) *pb.SerialPipeline {
	if in == nil {
		return nil
	}
	out := &pb.SerialPipeline{}
	out.Stages = direct.Slice_ToProto(mapCtx, in.Stages, Stage_ToProto)
	return out
}
func Stage_FromProto(mapCtx *direct.MapContext, in *pb.Stage) *krm.Stage {
	if in == nil {
		return nil
	}
	out := &krm.Stage{}
	out.TargetID = direct.LazyPtr(in.GetTargetId())
	out.Profiles = in.Profiles
	out.Strategy = Strategy_FromProto(mapCtx, in.GetStrategy())
	out.DeployParameters = direct.Slice_FromProto(mapCtx, in.DeployParameters, DeployParameters_FromProto)
	return out
}
func Stage_ToProto(mapCtx *direct.MapContext, in *krm.Stage) *pb.Stage {
	if in == nil {
		return nil
	}
	out := &pb.Stage{}
	out.TargetId = direct.ValueOf(in.TargetID)
	out.Profiles = in.Profiles
	out.Strategy = Strategy_ToProto(mapCtx, in.Strategy)
	out.DeployParameters = direct.Slice_ToProto(mapCtx, in.DeployParameters, DeployParameters_ToProto)
	return out
}
func Standard_FromProto(mapCtx *direct.MapContext, in *pb.Standard) *krm.Standard {
	if in == nil {
		return nil
	}
	out := &krm.Standard{}
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func Standard_ToProto(mapCtx *direct.MapContext, in *krm.Standard) *pb.Standard {
	if in == nil {
		return nil
	}
	out := &pb.Standard{}
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}
func Strategy_FromProto(mapCtx *direct.MapContext, in *pb.Strategy) *krm.Strategy {
	if in == nil {
		return nil
	}
	out := &krm.Strategy{}
	out.Standard = Standard_FromProto(mapCtx, in.GetStandard())
	out.Canary = Canary_FromProto(mapCtx, in.GetCanary())
	return out
}
func Strategy_ToProto(mapCtx *direct.MapContext, in *krm.Strategy) *pb.Strategy {
	if in == nil {
		return nil
	}
	out := &pb.Strategy{}
	if oneof := Standard_ToProto(mapCtx, in.Standard); oneof != nil {
		out.DeploymentStrategy = &pb.Strategy_Standard{Standard: oneof}
	}
	if oneof := Canary_ToProto(mapCtx, in.Canary); oneof != nil {
		out.DeploymentStrategy = &pb.Strategy_Canary{Canary: oneof}
	}
	return out
}
func TargetsPresentCondition_FromProto(mapCtx *direct.MapContext, in *pb.TargetsPresentCondition) *krm.TargetsPresentCondition {
	if in == nil {
		return nil
	}
	out := &krm.TargetsPresentCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.MissingTargets = in.MissingTargets
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func TargetsPresentCondition_ToProto(mapCtx *direct.MapContext, in *krm.TargetsPresentCondition) *pb.TargetsPresentCondition {
	if in == nil {
		return nil
	}
	out := &pb.TargetsPresentCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.MissingTargets = in.MissingTargets
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TargetsTypeCondition_FromProto(mapCtx *direct.MapContext, in *pb.TargetsTypeCondition) *krm.TargetsTypeCondition {
	if in == nil {
		return nil
	}
	out := &krm.TargetsTypeCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ErrorDetails = direct.LazyPtr(in.GetErrorDetails())
	return out
}
func TargetsTypeCondition_ToProto(mapCtx *direct.MapContext, in *krm.TargetsTypeCondition) *pb.TargetsTypeCondition {
	if in == nil {
		return nil
	}
	out := &pb.TargetsTypeCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.ErrorDetails = direct.ValueOf(in.ErrorDetails)
	return out
}

func AdvanceRolloutRule_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceRolloutRule) *krm.AdvanceRolloutRule {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceRolloutRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.SourcePhases = in.SourcePhases
	out.Wait = direct.StringDuration_FromProto(mapCtx, in.GetWait())
	// MISSING: Condition
	return out
}
func AdvanceRolloutRule_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceRolloutRule) *pb.AdvanceRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceRolloutRule{}
	out.Id = direct.ValueOf(in.ID)
	out.SourcePhases = in.SourcePhases
	out.Wait = direct.StringDuration_ToProto(mapCtx, in.Wait)
	// MISSING: Condition
	return out
}
func AdvanceRolloutRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceRolloutRule) *krm.AdvanceRolloutRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceRolloutRuleObservedState{}
	// MISSING: ID
	// MISSING: SourcePhases
	// MISSING: Wait
	out.Condition = AutomationRuleCondition_FromProto(mapCtx, in.GetCondition())
	return out
}
func AdvanceRolloutRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceRolloutRuleObservedState) *pb.AdvanceRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceRolloutRule{}
	// MISSING: ID
	// MISSING: SourcePhases
	// MISSING: Wait
	out.Condition = AutomationRuleCondition_ToProto(mapCtx, in.Condition)
	return out
}
func Automation_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.AutomationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomationSpec{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// out.Etag = direct.LazyPtr(in.GetEtag())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	out.ServiceAccountRef = &v1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	out.Selector = AutomationResourceSelector_FromProto(mapCtx, in.GetSelector())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRule_FromProto)
	return out
}
func Automation_ToProto(mapCtx *direct.MapContext, in *krm.AutomationSpec) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// out.Etag = direct.ValueOf(in.Etag)
	out.Suspended = direct.ValueOf(in.Suspended)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = direct.ValueOf(&in.ServiceAccountRef.External)
	}
	out.Selector = AutomationResourceSelector_ToProto(mapCtx, in.Selector)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRule_ToProto)
	return out
}
func AutomationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.CloudDeployAutomationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployAutomationObservedState{}
	// out.Name = direct.LazyPtr(in.GetName())
	// out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	//out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRuleObservedState_FromProto)
	return out
}
func AutomationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployAutomationObservedState) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	//out.Name = direct.ValueOf(in.Name)
	//out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	//out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRuleObservedState_ToProto)
	return out
}
func AutomationResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.AutomationResourceSelector) *krm.AutomationResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.AutomationResourceSelector{}
	out.Targets = direct.Slice_FromProto(mapCtx, in.Targets, TargetAttribute_FromProto)
	return out
}
func AutomationResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.AutomationResourceSelector) *pb.AutomationResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.AutomationResourceSelector{}
	out.Targets = direct.Slice_ToProto(mapCtx, in.Targets, TargetAttribute_ToProto)
	return out
}
func AutomationRule_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRule) *krm.AutomationRule {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRule{}
	out.PromoteReleaseRule = PromoteReleaseRule_FromProto(mapCtx, in.GetPromoteReleaseRule())
	out.AdvanceRolloutRule = AdvanceRolloutRule_FromProto(mapCtx, in.GetAdvanceRolloutRule())
	out.RepairRolloutRule = RepairRolloutRule_FromProto(mapCtx, in.GetRepairRolloutRule())
	out.TimedPromoteReleaseRule = TimedPromoteReleaseRule_FromProto(mapCtx, in.GetTimedPromoteReleaseRule())
	return out
}
func AutomationRule_ToProto(mapCtx *direct.MapContext, in *krm.AutomationRule) *pb.AutomationRule {
	if in == nil {
		return nil
	}
	out := &pb.AutomationRule{}
	if oneof := PromoteReleaseRule_ToProto(mapCtx, in.PromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_PromoteReleaseRule{PromoteReleaseRule: oneof}
	}
	if oneof := AdvanceRolloutRule_ToProto(mapCtx, in.AdvanceRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_AdvanceRolloutRule{AdvanceRolloutRule: oneof}
	}
	if oneof := RepairRolloutRule_ToProto(mapCtx, in.RepairRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_RepairRolloutRule{RepairRolloutRule: oneof}
	}
	if oneof := TimedPromoteReleaseRule_ToProto(mapCtx, in.TimedPromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_TimedPromoteReleaseRule{TimedPromoteReleaseRule: oneof}
	}
	return out
}
func AutomationRuleCondition_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRuleCondition) *krm.AutomationRuleCondition {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRuleCondition{}
	out.TargetsPresentCondition = TargetsPresentCondition_FromProto(mapCtx, in.GetTargetsPresentCondition())
	out.TimedPromoteReleaseCondition = TimedPromoteReleaseCondition_FromProto(mapCtx, in.GetTimedPromoteReleaseCondition())
	return out
}
func AutomationRuleCondition_ToProto(mapCtx *direct.MapContext, in *krm.AutomationRuleCondition) *pb.AutomationRuleCondition {
	if in == nil {
		return nil
	}
	out := &pb.AutomationRuleCondition{}
	out.TargetsPresentCondition = TargetsPresentCondition_ToProto(mapCtx, in.TargetsPresentCondition)
	if oneof := TimedPromoteReleaseCondition_ToProto(mapCtx, in.TimedPromoteReleaseCondition); oneof != nil {
		out.RuleTypeCondition = &pb.AutomationRuleCondition_TimedPromoteReleaseCondition{TimedPromoteReleaseCondition: oneof}
	}
	return out
}
func AutomationRuleConditionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRuleCondition) *krm.AutomationRuleConditionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRuleConditionObservedState{}
	// MISSING: TargetsPresentCondition
	out.TimedPromoteReleaseCondition = TimedPromoteReleaseConditionObservedState_FromProto(mapCtx, in.GetTimedPromoteReleaseCondition())
	return out
}
func AutomationRuleConditionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomationRuleConditionObservedState) *pb.AutomationRuleCondition {
	if in == nil {
		return nil
	}
	out := &pb.AutomationRuleCondition{}
	// MISSING: TargetsPresentCondition
	if oneof := TimedPromoteReleaseConditionObservedState_ToProto(mapCtx, in.TimedPromoteReleaseCondition); oneof != nil {
		out.RuleTypeCondition = &pb.AutomationRuleCondition_TimedPromoteReleaseCondition{TimedPromoteReleaseCondition: oneof}
	}
	return out
}
func AutomationRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRule) *krm.AutomationRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRuleObservedState{}
	out.PromoteReleaseRule = PromoteReleaseRuleObservedState_FromProto(mapCtx, in.GetPromoteReleaseRule())
	out.AdvanceRolloutRule = AdvanceRolloutRuleObservedState_FromProto(mapCtx, in.GetAdvanceRolloutRule())
	out.RepairRolloutRule = RepairRolloutRuleObservedState_FromProto(mapCtx, in.GetRepairRolloutRule())
	out.TimedPromoteReleaseRule = TimedPromoteReleaseRuleObservedState_FromProto(mapCtx, in.GetTimedPromoteReleaseRule())
	return out
}
func AutomationRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomationRuleObservedState) *pb.AutomationRule {
	if in == nil {
		return nil
	}
	out := &pb.AutomationRule{}
	if oneof := PromoteReleaseRuleObservedState_ToProto(mapCtx, in.PromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_PromoteReleaseRule{PromoteReleaseRule: oneof}
	}
	if oneof := AdvanceRolloutRuleObservedState_ToProto(mapCtx, in.AdvanceRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_AdvanceRolloutRule{AdvanceRolloutRule: oneof}
	}
	if oneof := RepairRolloutRuleObservedState_ToProto(mapCtx, in.RepairRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_RepairRolloutRule{RepairRolloutRule: oneof}
	}
	if oneof := TimedPromoteReleaseRuleObservedState_ToProto(mapCtx, in.TimedPromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_TimedPromoteReleaseRule{TimedPromoteReleaseRule: oneof}
	}
	return out
}
func DeployAutomationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.AutomationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomationObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	// MISSING: Rules
	return out
}
func DeployAutomationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomationObservedState) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	// MISSING: Rules
	return out
}
func DeployAutomationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.AutomationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomationSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	// MISSING: Rules
	return out
}
func DeployAutomationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomationSpec) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	// MISSING: Rules
	return out
}
func PromoteReleaseRule_FromProto(mapCtx *direct.MapContext, in *pb.PromoteReleaseRule) *krm.PromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &krm.PromoteReleaseRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Wait = direct.StringDuration_FromProto(mapCtx, in.GetWait())
	out.DestinationTargetID = direct.LazyPtr(in.GetDestinationTargetId())
	// MISSING: Condition
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	return out
}
func PromoteReleaseRule_ToProto(mapCtx *direct.MapContext, in *krm.PromoteReleaseRule) *pb.PromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.PromoteReleaseRule{}
	out.Id = direct.ValueOf(in.ID)
	out.Wait = direct.StringDuration_ToProto(mapCtx, in.Wait)
	out.DestinationTargetId = direct.ValueOf(in.DestinationTargetID)
	// MISSING: Condition
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	return out
}
func PromoteReleaseRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PromoteReleaseRule) *krm.PromoteReleaseRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PromoteReleaseRuleObservedState{}
	// MISSING: ID
	// MISSING: Wait
	// MISSING: DestinationTargetID
	out.Condition = AutomationRuleCondition_FromProto(mapCtx, in.GetCondition())
	// MISSING: DestinationPhase
	return out
}
func PromoteReleaseRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PromoteReleaseRuleObservedState) *pb.PromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.PromoteReleaseRule{}
	// MISSING: ID
	// MISSING: Wait
	// MISSING: DestinationTargetID
	out.Condition = AutomationRuleCondition_ToProto(mapCtx, in.Condition)
	// MISSING: DestinationPhase
	return out
}
func RepairPhaseConfig_FromProto(mapCtx *direct.MapContext, in *pb.RepairPhaseConfig) *krm.RepairPhaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.RepairPhaseConfig{}
	out.Retry = Retry_FromProto(mapCtx, in.GetRetry())
	out.Rollback = Rollback_FromProto(mapCtx, in.GetRollback())
	return out
}
func RepairPhaseConfig_ToProto(mapCtx *direct.MapContext, in *krm.RepairPhaseConfig) *pb.RepairPhaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepairPhaseConfig{}
	if oneof := Retry_ToProto(mapCtx, in.Retry); oneof != nil {
		out.RepairPhase = &pb.RepairPhaseConfig_Retry{Retry: oneof}
	}
	if oneof := Rollback_ToProto(mapCtx, in.Rollback); oneof != nil {
		out.RepairPhase = &pb.RepairPhaseConfig_Rollback{Rollback: oneof}
	}
	return out
}
func RepairRolloutRule_FromProto(mapCtx *direct.MapContext, in *pb.RepairRolloutRule) *krm.RepairRolloutRule {
	if in == nil {
		return nil
	}
	out := &krm.RepairRolloutRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Phases = in.Phases
	out.Jobs = in.Jobs
	// MISSING: Condition
	out.RepairPhases = direct.Slice_FromProto(mapCtx, in.RepairPhases, RepairPhaseConfig_FromProto)
	return out
}
func RepairRolloutRule_ToProto(mapCtx *direct.MapContext, in *krm.RepairRolloutRule) *pb.RepairRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.RepairRolloutRule{}
	out.Id = direct.ValueOf(in.ID)
	out.Phases = in.Phases
	out.Jobs = in.Jobs
	// MISSING: Condition
	out.RepairPhases = direct.Slice_ToProto(mapCtx, in.RepairPhases, RepairPhaseConfig_ToProto)
	return out
}
func RepairRolloutRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RepairRolloutRule) *krm.RepairRolloutRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RepairRolloutRuleObservedState{}
	// MISSING: ID
	// MISSING: Phases
	// MISSING: Jobs
	out.Condition = AutomationRuleCondition_FromProto(mapCtx, in.GetCondition())
	// MISSING: RepairPhases
	return out
}
func RepairRolloutRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RepairRolloutRuleObservedState) *pb.RepairRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.RepairRolloutRule{}
	// MISSING: ID
	// MISSING: Phases
	// MISSING: Jobs
	out.Condition = AutomationRuleCondition_ToProto(mapCtx, in.Condition)
	// MISSING: RepairPhases
	return out
}
func Retry_FromProto(mapCtx *direct.MapContext, in *pb.Retry) *krm.Retry {
	if in == nil {
		return nil
	}
	out := &krm.Retry{}
	out.Attempts = direct.LazyPtr(in.GetAttempts())
	out.Wait = direct.StringDuration_FromProto(mapCtx, in.GetWait())
	out.BackoffMode = direct.Enum_FromProto(mapCtx, in.GetBackoffMode())
	return out
}
func Retry_ToProto(mapCtx *direct.MapContext, in *krm.Retry) *pb.Retry {
	if in == nil {
		return nil
	}
	out := &pb.Retry{}
	out.Attempts = direct.ValueOf(in.Attempts)
	out.Wait = direct.StringDuration_ToProto(mapCtx, in.Wait)
	out.BackoffMode = direct.Enum_ToProto[pb.BackoffMode](mapCtx, in.BackoffMode)
	return out
}
func Rollback_FromProto(mapCtx *direct.MapContext, in *pb.Rollback) *krm.Rollback {
	if in == nil {
		return nil
	}
	out := &krm.Rollback{}
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	out.DisableRollbackIfRolloutPending = direct.LazyPtr(in.GetDisableRollbackIfRolloutPending())
	return out
}
func Rollback_ToProto(mapCtx *direct.MapContext, in *krm.Rollback) *pb.Rollback {
	if in == nil {
		return nil
	}
	out := &pb.Rollback{}
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	out.DisableRollbackIfRolloutPending = direct.ValueOf(in.DisableRollbackIfRolloutPending)
	return out
}
func TimedPromoteReleaseCondition_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseCondition) *krm.TimedPromoteReleaseCondition {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseCondition{}
	// MISSING: NextPromotionTime
	// MISSING: TargetsList
	return out
}
func TimedPromoteReleaseCondition_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseCondition) *pb.TimedPromoteReleaseCondition {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseCondition{}
	// MISSING: NextPromotionTime
	// MISSING: TargetsList
	return out
}
func TimedPromoteReleaseConditionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseCondition) *krm.TimedPromoteReleaseConditionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseConditionObservedState{}
	out.NextPromotionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextPromotionTime())
	out.TargetsList = direct.Slice_FromProto(mapCtx, in.TargetsList, TimedPromoteReleaseCondition_Targets_FromProto)
	return out
}
func TimedPromoteReleaseConditionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseConditionObservedState) *pb.TimedPromoteReleaseCondition {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseCondition{}
	out.NextPromotionTime = direct.StringTimestamp_ToProto(mapCtx, in.NextPromotionTime)
	out.TargetsList = direct.Slice_ToProto(mapCtx, in.TargetsList, TimedPromoteReleaseCondition_Targets_ToProto)
	return out
}
func TimedPromoteReleaseCondition_Targets_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseCondition_Targets) *krm.TimedPromoteReleaseCondition_Targets {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseCondition_Targets{}
	out.SourceTargetID = direct.LazyPtr(in.GetSourceTargetId())
	out.DestinationTargetID = direct.LazyPtr(in.GetDestinationTargetId())
	return out
}
func TimedPromoteReleaseCondition_Targets_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseCondition_Targets) *pb.TimedPromoteReleaseCondition_Targets {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseCondition_Targets{}
	out.SourceTargetId = direct.ValueOf(in.SourceTargetID)
	out.DestinationTargetId = direct.ValueOf(in.DestinationTargetID)
	return out
}
func TimedPromoteReleaseRule_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseRule) *krm.TimedPromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DestinationTargetID = direct.LazyPtr(in.GetDestinationTargetId())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	// MISSING: Condition
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	return out
}
func TimedPromoteReleaseRule_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseRule) *pb.TimedPromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseRule{}
	out.Id = direct.ValueOf(in.ID)
	out.DestinationTargetId = direct.ValueOf(in.DestinationTargetID)
	out.Schedule = direct.ValueOf(in.Schedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	// MISSING: Condition
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	return out
}
func TimedPromoteReleaseRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseRule) *krm.TimedPromoteReleaseRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseRuleObservedState{}
	// MISSING: ID
	// MISSING: DestinationTargetID
	// MISSING: Schedule
	// MISSING: TimeZone
	out.Condition = AutomationRuleCondition_FromProto(mapCtx, in.GetCondition())
	// MISSING: DestinationPhase
	return out
}
func TimedPromoteReleaseRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseRuleObservedState) *pb.TimedPromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseRule{}
	// MISSING: ID
	// MISSING: DestinationTargetID
	// MISSING: Schedule
	// MISSING: TimeZone
	out.Condition = AutomationRuleCondition_ToProto(mapCtx, in.Condition)
	// MISSING: DestinationPhase
	return out
}
