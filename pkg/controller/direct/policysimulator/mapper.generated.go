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

package policysimulator

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/policysimulator/apiv1/policysimulatorpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/policysimulator/v1alpha1"
)
func AccessStateDiff_FromProto(mapCtx *direct.MapContext, in *pb.AccessStateDiff) *krm.AccessStateDiff {
	if in == nil {
		return nil
	}
	out := &krm.AccessStateDiff{}
	out.Baseline = ExplainedAccess_FromProto(mapCtx, in.GetBaseline())
	out.Simulated = ExplainedAccess_FromProto(mapCtx, in.GetSimulated())
	out.AccessChange = direct.Enum_FromProto(mapCtx, in.GetAccessChange())
	return out
}
func AccessStateDiff_ToProto(mapCtx *direct.MapContext, in *krm.AccessStateDiff) *pb.AccessStateDiff {
	if in == nil {
		return nil
	}
	out := &pb.AccessStateDiff{}
	out.Baseline = ExplainedAccess_ToProto(mapCtx, in.Baseline)
	out.Simulated = ExplainedAccess_ToProto(mapCtx, in.Simulated)
	out.AccessChange = direct.Enum_ToProto[pb.AccessStateDiff_AccessChangeType](mapCtx, in.AccessChange)
	return out
}
func AccessTuple_FromProto(mapCtx *direct.MapContext, in *pb.AccessTuple) *krm.AccessTuple {
	if in == nil {
		return nil
	}
	out := &krm.AccessTuple{}
	out.Principal = direct.LazyPtr(in.GetPrincipal())
	out.FullResourceName = direct.LazyPtr(in.GetFullResourceName())
	out.Permission = direct.LazyPtr(in.GetPermission())
	return out
}
func AccessTuple_ToProto(mapCtx *direct.MapContext, in *krm.AccessTuple) *pb.AccessTuple {
	if in == nil {
		return nil
	}
	out := &pb.AccessTuple{}
	out.Principal = direct.ValueOf(in.Principal)
	out.FullResourceName = direct.ValueOf(in.FullResourceName)
	out.Permission = direct.ValueOf(in.Permission)
	return out
}
func BindingExplanation_FromProto(mapCtx *direct.MapContext, in *pb.BindingExplanation) *krm.BindingExplanation {
	if in == nil {
		return nil
	}
	out := &krm.BindingExplanation{}
	out.Access = direct.Enum_FromProto(mapCtx, in.GetAccess())
	out.Role = direct.LazyPtr(in.GetRole())
	out.RolePermission = direct.Enum_FromProto(mapCtx, in.GetRolePermission())
	out.RolePermissionRelevance = direct.Enum_FromProto(mapCtx, in.GetRolePermissionRelevance())
	// MISSING: Memberships
	out.Relevance = direct.Enum_FromProto(mapCtx, in.GetRelevance())
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	return out
}
func BindingExplanation_ToProto(mapCtx *direct.MapContext, in *krm.BindingExplanation) *pb.BindingExplanation {
	if in == nil {
		return nil
	}
	out := &pb.BindingExplanation{}
	out.Access = direct.Enum_ToProto[pb.AccessState](mapCtx, in.Access)
	out.Role = direct.ValueOf(in.Role)
	out.RolePermission = direct.Enum_ToProto[pb.BindingExplanation_RolePermission](mapCtx, in.RolePermission)
	out.RolePermissionRelevance = direct.Enum_ToProto[pb.HeuristicRelevance](mapCtx, in.RolePermissionRelevance)
	// MISSING: Memberships
	out.Relevance = direct.Enum_ToProto[pb.HeuristicRelevance](mapCtx, in.Relevance)
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	return out
}
func BindingExplanation_AnnotatedMembership_FromProto(mapCtx *direct.MapContext, in *pb.BindingExplanation_AnnotatedMembership) *krm.BindingExplanation_AnnotatedMembership {
	if in == nil {
		return nil
	}
	out := &krm.BindingExplanation_AnnotatedMembership{}
	out.Membership = direct.Enum_FromProto(mapCtx, in.GetMembership())
	out.Relevance = direct.Enum_FromProto(mapCtx, in.GetRelevance())
	return out
}
func BindingExplanation_AnnotatedMembership_ToProto(mapCtx *direct.MapContext, in *krm.BindingExplanation_AnnotatedMembership) *pb.BindingExplanation_AnnotatedMembership {
	if in == nil {
		return nil
	}
	out := &pb.BindingExplanation_AnnotatedMembership{}
	out.Membership = direct.Enum_ToProto[pb.BindingExplanation_Membership](mapCtx, in.Membership)
	out.Relevance = direct.Enum_ToProto[pb.HeuristicRelevance](mapCtx, in.Relevance)
	return out
}
func ExplainedAccess_FromProto(mapCtx *direct.MapContext, in *pb.ExplainedAccess) *krm.ExplainedAccess {
	if in == nil {
		return nil
	}
	out := &krm.ExplainedAccess{}
	out.AccessState = direct.Enum_FromProto(mapCtx, in.GetAccessState())
	out.Policies = direct.Slice_FromProto(mapCtx, in.Policies, ExplainedPolicy_FromProto)
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Status_FromProto)
	return out
}
func ExplainedAccess_ToProto(mapCtx *direct.MapContext, in *krm.ExplainedAccess) *pb.ExplainedAccess {
	if in == nil {
		return nil
	}
	out := &pb.ExplainedAccess{}
	out.AccessState = direct.Enum_ToProto[pb.AccessState](mapCtx, in.AccessState)
	out.Policies = direct.Slice_ToProto(mapCtx, in.Policies, ExplainedPolicy_ToProto)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Status_ToProto)
	return out
}
func ExplainedPolicy_FromProto(mapCtx *direct.MapContext, in *pb.ExplainedPolicy) *krm.ExplainedPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ExplainedPolicy{}
	out.Access = direct.Enum_FromProto(mapCtx, in.GetAccess())
	out.FullResourceName = direct.LazyPtr(in.GetFullResourceName())
	out.Policy = Policy_FromProto(mapCtx, in.GetPolicy())
	out.BindingExplanations = direct.Slice_FromProto(mapCtx, in.BindingExplanations, BindingExplanation_FromProto)
	out.Relevance = direct.Enum_FromProto(mapCtx, in.GetRelevance())
	return out
}
func ExplainedPolicy_ToProto(mapCtx *direct.MapContext, in *krm.ExplainedPolicy) *pb.ExplainedPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ExplainedPolicy{}
	out.Access = direct.Enum_ToProto[pb.AccessState](mapCtx, in.Access)
	out.FullResourceName = direct.ValueOf(in.FullResourceName)
	out.Policy = Policy_ToProto(mapCtx, in.Policy)
	out.BindingExplanations = direct.Slice_ToProto(mapCtx, in.BindingExplanations, BindingExplanation_ToProto)
	out.Relevance = direct.Enum_ToProto[pb.HeuristicRelevance](mapCtx, in.Relevance)
	return out
}
func PolicysimulatorReplayResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReplayResult) *krm.PolicysimulatorReplayResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicysimulatorReplayResultObservedState{}
	// MISSING: Diff
	// MISSING: Error
	// MISSING: Name
	// MISSING: Parent
	// MISSING: AccessTuple
	// MISSING: LastSeenDate
	return out
}
func PolicysimulatorReplayResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicysimulatorReplayResultObservedState) *pb.ReplayResult {
	if in == nil {
		return nil
	}
	out := &pb.ReplayResult{}
	// MISSING: Diff
	// MISSING: Error
	// MISSING: Name
	// MISSING: Parent
	// MISSING: AccessTuple
	// MISSING: LastSeenDate
	return out
}
func PolicysimulatorReplayResultSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReplayResult) *krm.PolicysimulatorReplayResultSpec {
	if in == nil {
		return nil
	}
	out := &krm.PolicysimulatorReplayResultSpec{}
	// MISSING: Diff
	// MISSING: Error
	// MISSING: Name
	// MISSING: Parent
	// MISSING: AccessTuple
	// MISSING: LastSeenDate
	return out
}
func PolicysimulatorReplayResultSpec_ToProto(mapCtx *direct.MapContext, in *krm.PolicysimulatorReplayResultSpec) *pb.ReplayResult {
	if in == nil {
		return nil
	}
	out := &pb.ReplayResult{}
	// MISSING: Diff
	// MISSING: Error
	// MISSING: Name
	// MISSING: Parent
	// MISSING: AccessTuple
	// MISSING: LastSeenDate
	return out
}
func ReplayDiff_FromProto(mapCtx *direct.MapContext, in *pb.ReplayDiff) *krm.ReplayDiff {
	if in == nil {
		return nil
	}
	out := &krm.ReplayDiff{}
	out.AccessDiff = AccessStateDiff_FromProto(mapCtx, in.GetAccessDiff())
	return out
}
func ReplayDiff_ToProto(mapCtx *direct.MapContext, in *krm.ReplayDiff) *pb.ReplayDiff {
	if in == nil {
		return nil
	}
	out := &pb.ReplayDiff{}
	out.AccessDiff = AccessStateDiff_ToProto(mapCtx, in.AccessDiff)
	return out
}
func ReplayResult_FromProto(mapCtx *direct.MapContext, in *pb.ReplayResult) *krm.ReplayResult {
	if in == nil {
		return nil
	}
	out := &krm.ReplayResult{}
	out.Diff = ReplayDiff_FromProto(mapCtx, in.GetDiff())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.AccessTuple = AccessTuple_FromProto(mapCtx, in.GetAccessTuple())
	out.LastSeenDate = Date_FromProto(mapCtx, in.GetLastSeenDate())
	return out
}
func ReplayResult_ToProto(mapCtx *direct.MapContext, in *krm.ReplayResult) *pb.ReplayResult {
	if in == nil {
		return nil
	}
	out := &pb.ReplayResult{}
	if oneof := ReplayDiff_ToProto(mapCtx, in.Diff); oneof != nil {
		out.Result = &pb.ReplayResult_Diff{Diff: oneof}
	}
	if oneof := Status_ToProto(mapCtx, in.Error); oneof != nil {
		out.Result = &pb.ReplayResult_Error{Error: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.AccessTuple = AccessTuple_ToProto(mapCtx, in.AccessTuple)
	out.LastSeenDate = Date_ToProto(mapCtx, in.LastSeenDate)
	return out
}
