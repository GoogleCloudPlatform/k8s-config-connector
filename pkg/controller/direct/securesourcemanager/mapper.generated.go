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

package securesourcemanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BranchRule_FromProto(mapCtx *direct.MapContext, in *pb.BranchRule) *krm.BranchRule {
	if in == nil {
		return nil
	}
	out := &krm.BranchRule{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.IncludePattern = direct.LazyPtr(in.GetIncludePattern())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.RequirePullRequest = direct.LazyPtr(in.GetRequirePullRequest())
	out.MinimumReviewsCount = direct.LazyPtr(in.GetMinimumReviewsCount())
	out.MinimumApprovalsCount = direct.LazyPtr(in.GetMinimumApprovalsCount())
	out.RequireCommentsResolved = direct.LazyPtr(in.GetRequireCommentsResolved())
	out.AllowStaleReviews = direct.LazyPtr(in.GetAllowStaleReviews())
	out.RequireLinearHistory = direct.LazyPtr(in.GetRequireLinearHistory())
	out.RequiredStatusChecks = direct.Slice_FromProto(mapCtx, in.RequiredStatusChecks, BranchRule_Check_FromProto)
	return out
}
func BranchRule_ToProto(mapCtx *direct.MapContext, in *krm.BranchRule) *pb.BranchRule {
	if in == nil {
		return nil
	}
	out := &pb.BranchRule{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Annotations = in.Annotations
	out.Etag = direct.ValueOf(in.Etag)
	out.IncludePattern = direct.ValueOf(in.IncludePattern)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.RequirePullRequest = direct.ValueOf(in.RequirePullRequest)
	out.MinimumReviewsCount = direct.ValueOf(in.MinimumReviewsCount)
	out.MinimumApprovalsCount = direct.ValueOf(in.MinimumApprovalsCount)
	out.RequireCommentsResolved = direct.ValueOf(in.RequireCommentsResolved)
	out.AllowStaleReviews = direct.ValueOf(in.AllowStaleReviews)
	out.RequireLinearHistory = direct.ValueOf(in.RequireLinearHistory)
	out.RequiredStatusChecks = direct.Slice_ToProto(mapCtx, in.RequiredStatusChecks, BranchRule_Check_ToProto)
	return out
}
func BranchRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BranchRule) *krm.BranchRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BranchRuleObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
func BranchRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BranchRuleObservedState) *pb.BranchRule {
	if in == nil {
		return nil
	}
	out := &pb.BranchRule{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
func BranchRule_Check_FromProto(mapCtx *direct.MapContext, in *pb.BranchRule_Check) *krm.BranchRule_Check {
	if in == nil {
		return nil
	}
	out := &krm.BranchRule_Check{}
	out.Context = direct.LazyPtr(in.GetContext())
	return out
}
func BranchRule_Check_ToProto(mapCtx *direct.MapContext, in *krm.BranchRule_Check) *pb.BranchRule_Check {
	if in == nil {
		return nil
	}
	out := &pb.BranchRule_Check{}
	out.Context = direct.ValueOf(in.Context)
	return out
}
func SecureSourceManagerInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateNote = direct.Enum_FromProto(mapCtx, in.GetStateNote())
	// MISSING: KMSKey
	out.HostConfig = Instance_HostConfig_FromProto(mapCtx, in.GetHostConfig())
	return out
}
func SecureSourceManagerInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateNote = direct.Enum_ToProto[pb.Instance_StateNote](mapCtx, in.StateNote)
	// MISSING: KMSKey
	out.HostConfig = Instance_HostConfig_ToProto(mapCtx, in.HostConfig)
	return out
}
func SecureSourceManagerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SecureSourceManagerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.PrivateConfig = Instance_PrivateConfig_FromProto(mapCtx, in.GetPrivateConfig())
	// MISSING: KMSKey
	return out
}
func SecureSourceManagerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.PrivateConfig = Instance_PrivateConfig_ToProto(mapCtx, in.PrivateConfig)
	// MISSING: KMSKey
	return out
}
func SecureSourceManagerRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.SecureSourceManagerRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.SecureSourceManagerRepositorySpec{}
	// MISSING: Name
	// MISSING: Description
	if in.GetInstance() != "" {
		out.InstanceRef = &refs.*SecureSourceManagerInstanceRef{External: in.GetInstance()}
	}
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Uris
	out.InitialConfig = Repository_InitialConfig_FromProto(mapCtx, in.GetInitialConfig())
	return out
}
func SecureSourceManagerRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.SecureSourceManagerRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: Description
	if in.InstanceRef != nil {
		out.Instance = in.InstanceRef.External
	}
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Uris
	out.InitialConfig = Repository_InitialConfig_ToProto(mapCtx, in.InitialConfig)
	return out
}
func SecuresourcemanagerBranchRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BranchRule) *krm.SecuresourcemanagerBranchRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuresourcemanagerBranchRuleObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
func SecuresourcemanagerBranchRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuresourcemanagerBranchRuleObservedState) *pb.BranchRule {
	if in == nil {
		return nil
	}
	out := &pb.BranchRule{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
func SecuresourcemanagerBranchRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.BranchRule) *krm.SecuresourcemanagerBranchRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuresourcemanagerBranchRuleSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
func SecuresourcemanagerBranchRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuresourcemanagerBranchRuleSpec) *pb.BranchRule {
	if in == nil {
		return nil
	}
	out := &pb.BranchRule{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Annotations
	// MISSING: Etag
	// MISSING: IncludePattern
	// MISSING: Disabled
	// MISSING: RequirePullRequest
	// MISSING: MinimumReviewsCount
	// MISSING: MinimumApprovalsCount
	// MISSING: RequireCommentsResolved
	// MISSING: AllowStaleReviews
	// MISSING: RequireLinearHistory
	// MISSING: RequiredStatusChecks
	return out
}
