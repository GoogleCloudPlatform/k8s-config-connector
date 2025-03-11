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
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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
func Automation_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.Automation {
	if in == nil {
		return nil
	}
	out := &krm.Automation{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Selector = AutomationResourceSelector_FromProto(mapCtx, in.GetSelector())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRule_FromProto)
	return out
}
func Automation_ToProto(mapCtx *direct.MapContext, in *krm.Automation) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	// MISSING: Name
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.Suspended = direct.ValueOf(in.Suspended)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Selector = AutomationResourceSelector_ToProto(mapCtx, in.Selector)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRule_ToProto)
	return out
}
func AutomationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.AutomationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AutomationRuleObservedState_FromProto)
	return out
}
func AutomationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomationObservedState) *pb.Automation {
	if in == nil {
		return nil
	}
	out := &pb.Automation{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Suspended
	// MISSING: ServiceAccount
	// MISSING: Selector
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AutomationRuleObservedState_ToProto)
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
func DeployAutomationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.DeployAutomationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployAutomationObservedState{}
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
func DeployAutomationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployAutomationObservedState) *pb.Automation {
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
func DeployAutomationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Automation) *krm.DeployAutomationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployAutomationSpec{}
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
func DeployAutomationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployAutomationSpec) *pb.Automation {
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
