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

package contactcenterinsights

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactcenterinsightsPhraseMatcherObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatcher) *krm.ContactcenterinsightsPhraseMatcherObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsPhraseMatcherObservedState{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: VersionTag
	// MISSING: RevisionCreateTime
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	// MISSING: ActivationUpdateTime
	// MISSING: RoleMatch
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsPhraseMatcherObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsPhraseMatcherObservedState) *pb.PhraseMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatcher{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: VersionTag
	// MISSING: RevisionCreateTime
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	// MISSING: ActivationUpdateTime
	// MISSING: RoleMatch
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsPhraseMatcherSpec_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatcher) *krm.ContactcenterinsightsPhraseMatcherSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsPhraseMatcherSpec{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: VersionTag
	// MISSING: RevisionCreateTime
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	// MISSING: ActivationUpdateTime
	// MISSING: RoleMatch
	// MISSING: UpdateTime
	return out
}
func ContactcenterinsightsPhraseMatcherSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsPhraseMatcherSpec) *pb.PhraseMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatcher{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: VersionTag
	// MISSING: RevisionCreateTime
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	// MISSING: ActivationUpdateTime
	// MISSING: RoleMatch
	// MISSING: UpdateTime
	return out
}
func ExactMatchConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExactMatchConfig) *krm.ExactMatchConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExactMatchConfig{}
	out.CaseSensitive = direct.LazyPtr(in.GetCaseSensitive())
	return out
}
func ExactMatchConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExactMatchConfig) *pb.ExactMatchConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExactMatchConfig{}
	out.CaseSensitive = direct.ValueOf(in.CaseSensitive)
	return out
}
func PhraseMatchRule_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatchRule) *krm.PhraseMatchRule {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatchRule{}
	out.Query = direct.LazyPtr(in.GetQuery())
	out.Negated = direct.LazyPtr(in.GetNegated())
	out.Config = PhraseMatchRuleConfig_FromProto(mapCtx, in.GetConfig())
	return out
}
func PhraseMatchRule_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatchRule) *pb.PhraseMatchRule {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatchRule{}
	out.Query = direct.ValueOf(in.Query)
	out.Negated = direct.ValueOf(in.Negated)
	out.Config = PhraseMatchRuleConfig_ToProto(mapCtx, in.Config)
	return out
}
func PhraseMatchRuleConfig_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatchRuleConfig) *krm.PhraseMatchRuleConfig {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatchRuleConfig{}
	out.ExactMatchConfig = ExactMatchConfig_FromProto(mapCtx, in.GetExactMatchConfig())
	return out
}
func PhraseMatchRuleConfig_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatchRuleConfig) *pb.PhraseMatchRuleConfig {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatchRuleConfig{}
	if oneof := ExactMatchConfig_ToProto(mapCtx, in.ExactMatchConfig); oneof != nil {
		out.Config = &pb.PhraseMatchRuleConfig_ExactMatchConfig{ExactMatchConfig: oneof}
	}
	return out
}
func PhraseMatchRuleGroup_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatchRuleGroup) *krm.PhraseMatchRuleGroup {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatchRuleGroup{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.PhraseMatchRules = direct.Slice_FromProto(mapCtx, in.PhraseMatchRules, PhraseMatchRule_FromProto)
	return out
}
func PhraseMatchRuleGroup_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatchRuleGroup) *pb.PhraseMatchRuleGroup {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatchRuleGroup{}
	out.Type = direct.Enum_ToProto[pb.PhraseMatchRuleGroup_PhraseMatchRuleGroupType](mapCtx, in.Type)
	out.PhraseMatchRules = direct.Slice_ToProto(mapCtx, in.PhraseMatchRules, PhraseMatchRule_ToProto)
	return out
}
func PhraseMatcher_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatcher) *krm.PhraseMatcher {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatcher{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RevisionID
	out.VersionTag = direct.LazyPtr(in.GetVersionTag())
	// MISSING: RevisionCreateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Active = direct.LazyPtr(in.GetActive())
	out.PhraseMatchRuleGroups = direct.Slice_FromProto(mapCtx, in.PhraseMatchRuleGroups, PhraseMatchRuleGroup_FromProto)
	// MISSING: ActivationUpdateTime
	out.RoleMatch = direct.Enum_FromProto(mapCtx, in.GetRoleMatch())
	// MISSING: UpdateTime
	return out
}
func PhraseMatcher_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatcher) *pb.PhraseMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatcher{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RevisionID
	out.VersionTag = direct.ValueOf(in.VersionTag)
	// MISSING: RevisionCreateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Type = direct.Enum_ToProto[pb.PhraseMatcher_PhraseMatcherType](mapCtx, in.Type)
	out.Active = direct.ValueOf(in.Active)
	out.PhraseMatchRuleGroups = direct.Slice_ToProto(mapCtx, in.PhraseMatchRuleGroups, PhraseMatchRuleGroup_ToProto)
	// MISSING: ActivationUpdateTime
	out.RoleMatch = direct.Enum_ToProto[pb.ConversationParticipant_Role](mapCtx, in.RoleMatch)
	// MISSING: UpdateTime
	return out
}
func PhraseMatcherObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhraseMatcher) *krm.PhraseMatcherObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PhraseMatcherObservedState{}
	// MISSING: Name
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: VersionTag
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	out.ActivationUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetActivationUpdateTime())
	// MISSING: RoleMatch
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PhraseMatcherObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PhraseMatcherObservedState) *pb.PhraseMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PhraseMatcher{}
	// MISSING: Name
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: VersionTag
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: Active
	// MISSING: PhraseMatchRuleGroups
	out.ActivationUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.ActivationUpdateTime)
	// MISSING: RoleMatch
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
