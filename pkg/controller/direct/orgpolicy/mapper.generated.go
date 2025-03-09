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

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CustomConstraint_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &krm.CustomConstraint{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_FromProto(mapCtx, in.MethodTypes)
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.ActionType = direct.Enum_FromProto(mapCtx, in.GetActionType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UpdateTime
	return out
}
func CustomConstraint_ToProto(mapCtx *direct.MapContext, in *krm.CustomConstraint) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	out.Name = direct.ValueOf(in.Name)
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_ToProto[pb.CustomConstraint_MethodType](mapCtx, in.MethodTypes)
	out.Condition = direct.ValueOf(in.Condition)
	out.ActionType = direct.Enum_ToProto[pb.CustomConstraint_ActionType](mapCtx, in.ActionType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UpdateTime
	return out
}
func CustomConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.CustomConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomConstraintObservedState{}
	// MISSING: Name
	// MISSING: ResourceTypes
	// MISSING: MethodTypes
	// MISSING: Condition
	// MISSING: ActionType
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CustomConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomConstraintObservedState) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	// MISSING: ResourceTypes
	// MISSING: MethodTypes
	// MISSING: Condition
	// MISSING: ActionType
	// MISSING: DisplayName
	// MISSING: Description
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func OrgPolicyCustomConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.OrgPolicyCustomConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyCustomConstraintObservedState{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func OrgPolicyCustomConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyCustomConstraintObservedState) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func OrgPolicyCustomConstraintSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomConstraint) *krm.OrgPolicyCustomConstraintSpec {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyCustomConstraintSpec{}
	// MISSING: Name
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_FromProto(mapCtx, in.MethodTypes)
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.ActionType = direct.Enum_FromProto(mapCtx, in.GetActionType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func OrgPolicyCustomConstraintSpec_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyCustomConstraintSpec) *pb.CustomConstraint {
	if in == nil {
		return nil
	}
	out := &pb.CustomConstraint{}
	// MISSING: Name
	out.ResourceTypes = in.ResourceTypes
	out.MethodTypes = direct.EnumSlice_ToProto[pb.CustomConstraint_MethodType](mapCtx, in.MethodTypes)
	out.Condition = direct.ValueOf(in.Condition)
	out.ActionType = direct.Enum_ToProto[pb.CustomConstraint_ActionType](mapCtx, in.ActionType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
