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
// krm.group: asset.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.asset.v1

package asset

import (
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AssetFeedSpec_FromProto(mapCtx *direct.MapContext, in *pb.Feed) *krm.AssetFeedSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssetFeedSpec{}
	// MISSING: Name
	out.AssetNames = in.AssetNames
	out.AssetTypes = in.AssetTypes
	out.ContentType = direct.Enum_FromProto(mapCtx, in.GetContentType())
	out.FeedOutputConfig = FeedOutputConfig_FromProto(mapCtx, in.GetFeedOutputConfig())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	out.RelationshipTypes = in.RelationshipTypes
	return out
}
func AssetFeedSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssetFeedSpec) *pb.Feed {
	if in == nil {
		return nil
	}
	out := &pb.Feed{}
	// MISSING: Name
	out.AssetNames = in.AssetNames
	out.AssetTypes = in.AssetTypes
	out.ContentType = direct.Enum_ToProto[pb.ContentType](mapCtx, in.ContentType)
	out.FeedOutputConfig = FeedOutputConfig_ToProto(mapCtx, in.FeedOutputConfig)
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	out.RelationshipTypes = in.RelationshipTypes
	return out
}
func AssetSavedQueryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.AssetSavedQueryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetSavedQueryObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.LastUpdater = direct.LazyPtr(in.GetLastUpdater())
	return out
}
func AssetSavedQueryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetSavedQueryObservedState) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Creator = direct.ValueOf(in.Creator)
	out.LastUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastUpdateTime)
	out.LastUpdater = direct.ValueOf(in.LastUpdater)
	return out
}
func AssetSavedQuerySpec_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.AssetSavedQuerySpec {
	if in == nil {
		return nil
	}
	out := &krm.AssetSavedQuerySpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Content = SavedQuery_QueryContent_FromProto(mapCtx, in.GetContent())
	return out
}
func AssetSavedQuerySpec_ToProto(mapCtx *direct.MapContext, in *krm.AssetSavedQuerySpec) *pb.SavedQuery {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Content = SavedQuery_QueryContent_ToProto(mapCtx, in.Content)
	return out
}
func FeedOutputConfig_FromProto(mapCtx *direct.MapContext, in *pb.FeedOutputConfig) *krm.FeedOutputConfig {
	if in == nil {
		return nil
	}
	out := &krm.FeedOutputConfig{}
	out.PubsubDestination = PubsubDestination_FromProto(mapCtx, in.GetPubsubDestination())
	return out
}
func FeedOutputConfig_ToProto(mapCtx *direct.MapContext, in *krm.FeedOutputConfig) *pb.FeedOutputConfig {
	if in == nil {
		return nil
	}
	out := &pb.FeedOutputConfig{}
	if oneof := PubsubDestination_ToProto(mapCtx, in.PubsubDestination); oneof != nil {
		out.Destination = &pb.FeedOutputConfig_PubsubDestination{PubsubDestination: oneof}
	}
	return out
}
func IAMPolicyAnalysisQuery_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery) *krm.IAMPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery{}
	out.Scope = direct.LazyPtr(in.GetScope())
	out.ResourceSelector = IAMPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx, in.GetResourceSelector())
	out.IdentitySelector = IAMPolicyAnalysisQuery_IdentitySelector_FromProto(mapCtx, in.GetIdentitySelector())
	out.AccessSelector = IAMPolicyAnalysisQuery_AccessSelector_FromProto(mapCtx, in.GetAccessSelector())
	out.Options = IAMPolicyAnalysisQuery_Options_FromProto(mapCtx, in.GetOptions())
	out.ConditionContext = IAMPolicyAnalysisQuery_ConditionContext_FromProto(mapCtx, in.GetConditionContext())
	return out
}
func IAMPolicyAnalysisQuery_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery) *pb.IamPolicyAnalysisQuery {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery{}
	out.Scope = direct.ValueOf(in.Scope)
	out.ResourceSelector = IAMPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx, in.ResourceSelector)
	out.IdentitySelector = IAMPolicyAnalysisQuery_IdentitySelector_ToProto(mapCtx, in.IdentitySelector)
	out.AccessSelector = IAMPolicyAnalysisQuery_AccessSelector_ToProto(mapCtx, in.AccessSelector)
	out.Options = IAMPolicyAnalysisQuery_Options_ToProto(mapCtx, in.Options)
	out.ConditionContext = IAMPolicyAnalysisQuery_ConditionContext_ToProto(mapCtx, in.ConditionContext)
	return out
}
func IAMPolicyAnalysisQuery_AccessSelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_AccessSelector) *krm.IAMPolicyAnalysisQuery_AccessSelector {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery_AccessSelector{}
	out.Roles = in.Roles
	out.Permissions = in.Permissions
	return out
}
func IAMPolicyAnalysisQuery_AccessSelector_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery_AccessSelector) *pb.IamPolicyAnalysisQuery_AccessSelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_AccessSelector{}
	out.Roles = in.Roles
	out.Permissions = in.Permissions
	return out
}
func IAMPolicyAnalysisQuery_ConditionContext_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_ConditionContext) *krm.IAMPolicyAnalysisQuery_ConditionContext {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery_ConditionContext{}
	out.AccessTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAccessTime())
	return out
}
func IAMPolicyAnalysisQuery_ConditionContext_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery_ConditionContext) *pb.IamPolicyAnalysisQuery_ConditionContext {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_ConditionContext{}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.AccessTime); oneof != nil {
		out.TimeContext = &pb.IamPolicyAnalysisQuery_ConditionContext_AccessTime{AccessTime: oneof}
	}
	return out
}
func IAMPolicyAnalysisQuery_IdentitySelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_IdentitySelector) *krm.IAMPolicyAnalysisQuery_IdentitySelector {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery_IdentitySelector{}
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func IAMPolicyAnalysisQuery_IdentitySelector_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery_IdentitySelector) *pb.IamPolicyAnalysisQuery_IdentitySelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_IdentitySelector{}
	out.Identity = direct.ValueOf(in.Identity)
	return out
}
func IAMPolicyAnalysisQuery_Options_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_Options) *krm.IAMPolicyAnalysisQuery_Options {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery_Options{}
	out.ExpandGroups = direct.LazyPtr(in.GetExpandGroups())
	out.ExpandRoles = direct.LazyPtr(in.GetExpandRoles())
	out.ExpandResources = direct.LazyPtr(in.GetExpandResources())
	out.OutputResourceEdges = direct.LazyPtr(in.GetOutputResourceEdges())
	out.OutputGroupEdges = direct.LazyPtr(in.GetOutputGroupEdges())
	out.AnalyzeServiceAccountImpersonation = direct.LazyPtr(in.GetAnalyzeServiceAccountImpersonation())
	return out
}
func IAMPolicyAnalysisQuery_Options_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery_Options) *pb.IamPolicyAnalysisQuery_Options {
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
func IAMPolicyAnalysisQuery_ResourceSelector_FromProto(mapCtx *direct.MapContext, in *pb.IamPolicyAnalysisQuery_ResourceSelector) *krm.IAMPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &krm.IAMPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = direct.LazyPtr(in.GetFullResourceName())
	return out
}
func IAMPolicyAnalysisQuery_ResourceSelector_ToProto(mapCtx *direct.MapContext, in *krm.IAMPolicyAnalysisQuery_ResourceSelector) *pb.IamPolicyAnalysisQuery_ResourceSelector {
	if in == nil {
		return nil
	}
	out := &pb.IamPolicyAnalysisQuery_ResourceSelector{}
	out.FullResourceName = direct.ValueOf(in.FullResourceName)
	return out
}
func PubsubDestination_FromProto(mapCtx *direct.MapContext, in *pb.PubsubDestination) *krm.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &krm.PubsubDestination{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	return out
}
func PubsubDestination_ToProto(mapCtx *direct.MapContext, in *krm.PubsubDestination) *pb.PubsubDestination {
	if in == nil {
		return nil
	}
	out := &pb.PubsubDestination{}
	out.Topic = direct.ValueOf(in.Topic)
	return out
}
func SavedQuery_QueryContent_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery_QueryContent) *krm.SavedQuery_QueryContent {
	if in == nil {
		return nil
	}
	out := &krm.SavedQuery_QueryContent{}
	out.IAMPolicyAnalysisQuery = IAMPolicyAnalysisQuery_FromProto(mapCtx, in.GetIamPolicyAnalysisQuery())
	return out
}
func SavedQuery_QueryContent_ToProto(mapCtx *direct.MapContext, in *krm.SavedQuery_QueryContent) *pb.SavedQuery_QueryContent {
	if in == nil {
		return nil
	}
	out := &pb.SavedQuery_QueryContent{}
	if oneof := IAMPolicyAnalysisQuery_ToProto(mapCtx, in.IAMPolicyAnalysisQuery); oneof != nil {
		out.QueryContent = &pb.SavedQuery_QueryContent_IamPolicyAnalysisQuery{IamPolicyAnalysisQuery: oneof}
	}
	return out
}
