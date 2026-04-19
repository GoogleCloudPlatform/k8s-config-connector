// Copyright 2026 Google LLC
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

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDeployAutomationSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.CloudDeployAutomationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployAutomationSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for k, v := range in.Annotations {
			out.Annotations[k] = v
		}
	}
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Selector = AutomationResourceSelector_v1alpha1_FromProto(mapCtx, in.GetSelector())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRule_v1alpha1_FromProto)
	return out
}

func CloudDeployAutomationSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployAutomationSpec) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	out.Description = direct.ValueOf(in.Description)
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for k, v := range in.Annotations {
			out.Annotations[k] = v
		}
	}
	out.Suspended = direct.ValueOf(in.Suspended)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Selector = AutomationResourceSelector_v1alpha1_ToProto(mapCtx, in.Selector)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRule_v1alpha1_ToProto)
	return out
}

func CloudDeployAutomationObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.CloudDeployAutomationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDeployAutomationObservedState{}
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRuleObservedState_v1alpha1_FromProto)
	return out
}

func CloudDeployAutomationObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.CloudDeployAutomationObservedState) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRuleObservedState_v1alpha1_ToProto)
	return out
}

func AutomationResourceSelector_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutomationResourceSelector) *krm.AutomationResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.AutomationResourceSelector{}
	out.Targets = direct.Slice_FromProto(mapCtx, in.Targets, TargetAttribute_v1alpha1_FromProto)
	return out
}
func AutomationResourceSelector_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutomationResourceSelector) *pb.AutomationResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.AutomationResourceSelector{}
	out.Targets = direct.Slice_ToProto(mapCtx, in.Targets, TargetAttribute_v1alpha1_ToProto)
	return out
}

func TargetAttribute_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TargetAttribute) *krm.TargetAttribute {
	if in == nil {
		return nil
	}
	out := &krm.TargetAttribute{}
	if in.GetId() != "" {
		out.TargetRef = &krm.CloudDeployTargetRef{External: in.GetId()}
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for k, v := range in.Labels {
			out.Labels[k] = v
		}
	}
	return out
}
func TargetAttribute_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.TargetAttribute) *pb.TargetAttribute {
	if in == nil {
		return nil
	}
	out := &pb.TargetAttribute{}
	if in.TargetRef != nil {
		out.Id = in.TargetRef.External
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for k, v := range in.Labels {
			out.Labels[k] = v
		}
	}
	return out
}

func AutomationRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRule) *krm.AutomationRule {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRule{}
	out.PromoteReleaseRule = PromoteReleaseRule_v1alpha1_FromProto(mapCtx, in.GetPromoteReleaseRule())
	out.AdvanceRolloutRule = AdvanceRolloutRule_v1alpha1_FromProto(mapCtx, in.GetAdvanceRolloutRule())
	out.RepairRolloutRule = RepairRolloutRule_v1alpha1_FromProto(mapCtx, in.GetRepairRolloutRule())
	out.TimedPromoteReleaseRule = TimedPromoteReleaseRule_v1alpha1_FromProto(mapCtx, in.GetTimedPromoteReleaseRule())
	return out
}
func AutomationRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutomationRule) *pb.AutomationRule {
	if in == nil {
		return nil
	}
	out := &pb.AutomationRule{}
	if oneof := PromoteReleaseRule_v1alpha1_ToProto(mapCtx, in.PromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_PromoteReleaseRule{PromoteReleaseRule: oneof}
	}
	if oneof := AdvanceRolloutRule_v1alpha1_ToProto(mapCtx, in.AdvanceRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_AdvanceRolloutRule{AdvanceRolloutRule: oneof}
	}
	if oneof := RepairRolloutRule_v1alpha1_ToProto(mapCtx, in.RepairRolloutRule); oneof != nil {
		out.Rule = &pb.AutomationRule_RepairRolloutRule{RepairRolloutRule: oneof}
	}
	if oneof := TimedPromoteReleaseRule_v1alpha1_ToProto(mapCtx, in.TimedPromoteReleaseRule); oneof != nil {
		out.Rule = &pb.AutomationRule_TimedPromoteReleaseRule{TimedPromoteReleaseRule: oneof}
	}
	return out
}

func PromoteReleaseRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PromoteReleaseRule) *krm.PromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &krm.PromoteReleaseRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Wait = direct.Duration_FromProto(mapCtx, in.GetWait())
	if in.GetDestinationTargetId() != "" {
		out.DestinationTargetRef = &krm.CloudDeployTargetRef{External: in.GetDestinationTargetId()}
	}
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	return out
}
func PromoteReleaseRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PromoteReleaseRule) *pb.PromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.PromoteReleaseRule{}
	out.Id = direct.ValueOf(in.ID)
	out.Wait = direct.Duration_ToProto(mapCtx, in.Wait)
	if in.DestinationTargetRef != nil {
		out.DestinationTargetId = in.DestinationTargetRef.External
	}
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	return out
}

func AdvanceRolloutRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceRolloutRule) *krm.AdvanceRolloutRule {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceRolloutRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.SourcePhases = in.SourcePhases
	out.Wait = direct.Duration_FromProto(mapCtx, in.GetWait())
	return out
}
func AdvanceRolloutRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AdvanceRolloutRule) *pb.AdvanceRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.AdvanceRolloutRule{}
	out.Id = direct.ValueOf(in.ID)
	out.SourcePhases = in.SourcePhases
	out.Wait = direct.Duration_ToProto(mapCtx, in.Wait)
	return out
}

func RepairRolloutRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.RepairRolloutRule) *krm.RepairRolloutRule {
	if in == nil {
		return nil
	}
	out := &krm.RepairRolloutRule{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Phases = in.Phases
	out.Jobs = in.Jobs
	out.RepairPhases = direct.Slice_FromProto(mapCtx, in.RepairPhases, RepairPhaseConfig_v1alpha1_FromProto)
	return out
}
func RepairRolloutRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RepairRolloutRule) *pb.RepairRolloutRule {
	if in == nil {
		return nil
	}
	out := &pb.RepairRolloutRule{}
	out.Id = direct.ValueOf(in.ID)
	out.Phases = in.Phases
	out.Jobs = in.Jobs
	out.RepairPhases = direct.Slice_ToProto(mapCtx, in.RepairPhases, RepairPhaseConfig_v1alpha1_ToProto)
	return out
}

func RepairPhaseConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.RepairPhaseConfig) *krm.RepairPhaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.RepairPhaseConfig{}
	out.Retry = Retry_v1alpha1_FromProto(mapCtx, in.GetRetry())
	out.Rollback = Rollback_v1alpha1_FromProto(mapCtx, in.GetRollback())
	return out
}
func RepairPhaseConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RepairPhaseConfig) *pb.RepairPhaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepairPhaseConfig{}
	if in.Retry != nil {
		out.RepairPhase = &pb.RepairPhaseConfig_Retry{Retry: Retry_v1alpha1_ToProto(mapCtx, in.Retry)}
	}
	if in.Rollback != nil {
		out.RepairPhase = &pb.RepairPhaseConfig_Rollback{Rollback: Rollback_v1alpha1_ToProto(mapCtx, in.Rollback)}
	}
	return out
}

func Retry_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Retry) *krm.Retry {
	if in == nil {
		return nil
	}
	out := &krm.Retry{}
	out.Attempts = direct.LazyPtr(in.GetAttempts())
	out.Wait = direct.Duration_FromProto(mapCtx, in.GetWait())
	out.BackoffMode = direct.Enum_FromProto(mapCtx, in.GetBackoffMode())
	return out
}
func Retry_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Retry) *pb.Retry {
	if in == nil {
		return nil
	}
	out := &pb.Retry{}
	out.Attempts = direct.ValueOf(in.Attempts)
	out.Wait = direct.Duration_ToProto(mapCtx, in.Wait)
	out.BackoffMode = pb.BackoffMode(direct.Enum_ToProto[pb.BackoffMode](mapCtx, in.BackoffMode))
	return out
}

func Rollback_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Rollback) *krm.Rollback {
	if in == nil {
		return nil
	}
	out := &krm.Rollback{}
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	out.DisableRollbackIfRolloutPending = direct.LazyPtr(in.GetDisableRollbackIfRolloutPending())
	return out
}
func Rollback_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Rollback) *pb.Rollback {
	if in == nil {
		return nil
	}
	out := &pb.Rollback{}
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	out.DisableRollbackIfRolloutPending = direct.ValueOf(in.DisableRollbackIfRolloutPending)
	return out
}

func TimedPromoteReleaseRule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseRule) *krm.TimedPromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseRule{}
	out.ID = direct.LazyPtr(in.GetId())
	if in.GetDestinationTargetId() != "" {
		out.DestinationTargetRef = &krm.CloudDeployTargetRef{External: in.GetDestinationTargetId()}
	}
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.DestinationPhase = direct.LazyPtr(in.GetDestinationPhase())
	return out
}
func TimedPromoteReleaseRule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.TimedPromoteReleaseRule) *pb.TimedPromoteReleaseRule {
	if in == nil {
		return nil
	}
	out := &pb.TimedPromoteReleaseRule{}
	out.Id = direct.ValueOf(in.ID)
	if in.DestinationTargetRef != nil {
		out.DestinationTargetId = in.DestinationTargetRef.External
	}
	out.Schedule = direct.ValueOf(in.Schedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.DestinationPhase = direct.ValueOf(in.DestinationPhase)
	return out
}

func AutomationRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRule) *krm.AutomationRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRuleObservedState{}
	out.PromoteReleaseRule = PromoteReleaseRuleObservedState_v1alpha1_FromProto(mapCtx, in.GetPromoteReleaseRule())
	out.AdvanceRolloutRule = AdvanceRolloutRuleObservedState_v1alpha1_FromProto(mapCtx, in.GetAdvanceRolloutRule())
	out.RepairRolloutRule = RepairRolloutRuleObservedState_v1alpha1_FromProto(mapCtx, in.GetRepairRolloutRule())
	out.TimedPromoteReleaseRule = TimedPromoteReleaseRuleObservedState_v1alpha1_FromProto(mapCtx, in.GetTimedPromoteReleaseRule())
	return out
}

func PromoteReleaseRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PromoteReleaseRule) *krm.PromoteReleaseRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PromoteReleaseRuleObservedState{}
	out.Condition = AutomationRuleCondition_v1alpha1_FromProto(mapCtx, in.GetCondition())
	return out
}

func AdvanceRolloutRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AdvanceRolloutRule) *krm.AdvanceRolloutRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdvanceRolloutRuleObservedState{}
	out.Condition = AutomationRuleCondition_v1alpha1_FromProto(mapCtx, in.GetCondition())
	return out
}

func RepairRolloutRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.RepairRolloutRule) *krm.RepairRolloutRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RepairRolloutRuleObservedState{}
	out.Condition = AutomationRuleCondition_v1alpha1_FromProto(mapCtx, in.GetCondition())
	return out
}

func TimedPromoteReleaseRuleObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseRule) *krm.TimedPromoteReleaseRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseRuleObservedState{}
	out.Condition = AutomationRuleCondition_v1alpha1_FromProto(mapCtx, in.GetCondition())
	return out
}

func AutomationRuleCondition_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutomationRuleCondition) *krm.AutomationRuleCondition {
	if in == nil {
		return nil
	}
	out := &krm.AutomationRuleCondition{}
	out.TargetsPresentCondition = TargetsPresentCondition_v1alpha1_FromProto(mapCtx, in.GetTargetsPresentCondition())
	out.TimedPromoteReleaseCondition = TimedPromoteReleaseCondition_v1alpha1_FromProto(mapCtx, in.GetTimedPromoteReleaseCondition())
	return out
}

func TargetsPresentCondition_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TargetsPresentCondition) *krm.TargetsPresentCondition {
	if in == nil {
		return nil
	}
	out := &krm.TargetsPresentCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.MissingTargets = in.MissingTargets
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func TimedPromoteReleaseCondition_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseCondition) *krm.TimedPromoteReleaseCondition {
	if in == nil {
		return nil
	}
	return &krm.TimedPromoteReleaseCondition{
		NextPromotionTime: direct.StringTimestamp_FromProto(mapCtx, in.GetNextPromotionTime()),
		TargetsList:       direct.Slice_FromProto(mapCtx, in.TargetsList, TimedPromoteReleaseCondition_Targets_v1alpha1_FromProto),
	}
}

func TimedPromoteReleaseCondition_Targets_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TimedPromoteReleaseCondition_Targets) *krm.TimedPromoteReleaseCondition_Targets {
	if in == nil {
		return nil
	}
	out := &krm.TimedPromoteReleaseCondition_Targets{}
	out.SourceTargetID = direct.LazyPtr(in.GetSourceTargetId())
	out.DestinationTargetID = direct.LazyPtr(in.GetDestinationTargetId())
	return out
}
