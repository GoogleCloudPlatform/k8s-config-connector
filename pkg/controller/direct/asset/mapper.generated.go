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

package asset

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AssetFeedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.AssetFeedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetFeedObservedState{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetFeedObservedState) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedSpec_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.AssetFeedSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssetFeedSpec{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetFeedSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssetFeedSpec) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	// MISSING: Name
	// MISSING: AssetNames
	// MISSING: AssetTypes
	// MISSING: ContentType
	// MISSING: FeedOutputConfig
	// MISSING: Condition
	// MISSING: RelationshipTypes
	return out
}
func AssetSavedQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.AssetSavedQueryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetSavedQueryObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	// MISSING: Labels
	// MISSING: Content
	return out
}
func AssetSavedQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetSavedQueryObservedState) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	// MISSING: Labels
	// MISSING: Content
	return out
}
func AssetSavedQuerySpec_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.AssetSavedQuerySpec {
	if in == nil {
		return nil
	}
	out := &krm.AssetSavedQuerySpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	// MISSING: Labels
	// MISSING: Content
	return out
}
func AssetSavedQuerySpec_ToProto(mapCtx *direct.MapContext, in *krm.AssetSavedQuerySpec) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	// MISSING: Labels
	// MISSING: Content
	return out
}
func IamPolicyAnalysisQuery_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery) *krm.IamPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery{}
	out.Scope = direct.LazyPtr(in.GetScope())
	out.ResourceSelector = IamPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx, in.GetResourceSelector())
	out.IdentitySelector = IamPolicyAnalysisQuery_IdentitySelector_FromProto(mapCtx, in.GetIdentitySelector())
	out.AccessSelector = IamPolicyAnalysisQuery_AccessSelector_FromProto(mapCtx, in.GetAccessSelector())
	out.Options = IamPolicyAnalysisQuery_Options_FromProto(mapCtx, in.GetOptions())
	out.ConditionContext = IamPolicyAnalysisQuery_ConditionContext_FromProto(mapCtx, in.GetConditionContext())
	return out
}
func IamPolicyAnalysisQuery_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery) *pb.IamPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery{}
	out.Scope = direct.ValueOf(in.Scope)
	out.ResourceSelector = IamPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx, in.ResourceSelector)
	out.IdentitySelector = IamPolicyAnalysisQuery_IdentitySelector_ToProto(mapCtx, in.IdentitySelector)
	out.AccessSelector = IamPolicyAnalysisQuery_AccessSelector_ToProto(mapCtx, in.AccessSelector)
	out.Options = IamPolicyAnalysisQuery_Options_ToProto(mapCtx, in.Options)
	out.ConditionContext = IamPolicyAnalysisQuery_ConditionContext_ToProto(mapCtx, in.ConditionContext)
	return out
}
func IamPolicyAnalysisQuery_AccessSelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_AccessSelector) *krm.IamPolicyAnalysisQuery_AccessSelector {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery_AccessSelector{}
	out.Roles = in.Roles
	out.Permissions = in.Permissions
	return out
}
func IamPolicyAnalysisQuery_AccessSelector_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery_AccessSelector) *pb.IamPolicyAnalysisQuery_AccessSelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_AccessSelector{}
	out.Roles = in.Roles
	out.Permissions = in.Permissions
	return out
}
func IamPolicyAnalysisQuery_ConditionContext_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_ConditionContext) *krm.IamPolicyAnalysisQuery_ConditionContext {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery_ConditionContext{}
	out.AccessTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAccessTime())
	return out
}
func IamPolicyAnalysisQuery_ConditionContext_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery_ConditionContext) *pb.IamPolicyAnalysisQuery_ConditionContext {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_ConditionContext{}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.AccessTime); oneof != nil {
		out.TimeContext = &pb.IamPolicyAnalysisQuery_ConditionContext_AccessTime{AccessTime: oneof}
	}
	return out
}
func IamPolicyAnalysisQuery_IdentitySelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_IdentitySelector) *krm.IamPolicyAnalysisQuery_IdentitySelector {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery_IdentitySelector{}
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func IamPolicyAnalysisQuery_IdentitySelector_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery_IdentitySelector) *pb.IamPolicyAnalysisQuery_IdentitySelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_IdentitySelector{}
	out.Identity = direct.ValueOf(in.Identity)
	return out
}
func IamPolicyAnalysisQuery_Options_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_Options) *krm.IamPolicyAnalysisQuery_Options {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery_Options{}
	out.ExpandGroups = direct.LazyPtr(in.GetExpandGroups())
	out.ExpandRoles = direct.LazyPtr(in.GetExpandRoles())
	out.ExpandResources = direct.LazyPtr(in.GetExpandResources())
	out.OutputResourceEdges = direct.LazyPtr(in.GetOutputResourceEdges())
	out.OutputGroupEdges = direct.LazyPtr(in.GetOutputGroupEdges())
	out.AnalyzeServiceAccountImpersonation = direct.LazyPtr(in.GetAnalyzeServiceAccountImpersonation())
	return out
}
func IamPolicyAnalysisQuery_Options_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery_Options) *pb.IamPolicyAnalysisQuery_Options {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_Options{}
	out.ExpandGroups = direct.ValueOf(in.ExpandGroups)
	out.ExpandRoles = direct.ValueOf(in.ExpandRoles)
	out.ExpandResources = direct.ValueOf(in.ExpandResources)
	out.OutputResourceEdges = direct.ValueOf(in.OutputResourceEdges)
	out.OutputGroupEdges = direct.ValueOf(in.OutputGroupEdges)
	out.AnalyzeServiceAccountImpersonation = direct.ValueOf(in.AnalyzeServiceAccountImpersonation)
	return out
}
func IamPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_ResourceSelector) *krm.IamPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.IamPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = direct.LazyPtr(in.GetFullResourceName())
	return out
}
func IamPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.IamPolicyAnalysisQuery_ResourceSelector) *pb.IamPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = direct.ValueOf(in.FullResourceName)
	return out
}
func SavedQuery_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.SavedQuery {
	if in == nil {
		return nil
	}
	out := &krm.SavedQuery{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	out.Labels = in.Labels
	out.Content = SavedQuery_QueryContent_FromProto(mapCtx, in.GetContent())
	return out
}
func SavedQuery_ToProto(mapCtx *direct.MapContext, in *krm.SavedQuery) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: Creator
	// MISSING: LastUpdateTime
	// MISSING: LastUpdater
	out.Labels = in.Labels
	out.Content = SavedQuery_QueryContent_ToProto(mapCtx, in.Content)
	return out
}
func SavedQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.SavedQueryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SavedQueryObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.LastUpdater = direct.LazyPtr(in.GetLastUpdater())
	// MISSING: Labels
	// MISSING: Content
	return out
}
func SavedQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SavedQueryObservedState) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Creator = direct.ValueOf(in.Creator)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	out.LastUpdater = direct.ValueOf(in.LastUpdater)
	// MISSING: Labels
	// MISSING: Content
	return out
}
func SavedQuery_QueryContent_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery_QueryContent) *krm.SavedQuery_QueryContent {
	if in == nil {
		return nil
	}
	out := &krm.SavedQuery_QueryContent{}
	out.IamPolicyAnalysisQuery = IamPolicyAnalysisQuery_FromProto(mapCtx, in.GetIamPolicyAnalysisQuery())
	return out
}
func SavedQuery_QueryContent_ToProto(mapCtx *direct.MapContext, in *krm.SavedQuery_QueryContent) *pb.SavedQuery_QueryContent {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery_QueryContent{}
	if oneof := IamPolicyAnalysisQuery_ToProto(mapCtx, in.IamPolicyAnalysisQuery); oneof != nil {
		out.QueryContent = &pb.SavedQuery_QueryContent_IamPolicyAnalysisQuery{IamPolicyAnalysisQuery: oneof}
	}
	return out
}
