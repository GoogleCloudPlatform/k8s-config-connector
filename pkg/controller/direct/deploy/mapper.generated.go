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
