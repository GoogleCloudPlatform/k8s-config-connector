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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Constraint_FromProto(mapCtx *direct.MapContext, in *pb.Constraint) *krm.Constraint {
	if in == nil {
		return nil
	}
	out := &krm.Constraint{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ConstraintDefault = direct.Enum_FromProto(mapCtx, in.GetConstraintDefault())
	out.ListConstraint = Constraint_ListConstraint_FromProto(mapCtx, in.GetListConstraint())
	out.BooleanConstraint = Constraint_BooleanConstraint_FromProto(mapCtx, in.GetBooleanConstraint())
	out.SupportsDryRun = direct.LazyPtr(in.GetSupportsDryRun())
	return out
}
func Constraint_ToProto(mapCtx *direct.MapContext, in *krm.Constraint) *pb.Constraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ConstraintDefault = direct.Enum_ToProto[pb.Constraint_ConstraintDefault](mapCtx, in.ConstraintDefault)
	if oneof := Constraint_ListConstraint_ToProto(mapCtx, in.ListConstraint); oneof != nil {
		out.ConstraintType = &pb.Constraint_ListConstraint_{ListConstraint: oneof}
	}
	if oneof := Constraint_BooleanConstraint_ToProto(mapCtx, in.BooleanConstraint); oneof != nil {
		out.ConstraintType = &pb.Constraint_BooleanConstraint_{BooleanConstraint: oneof}
	}
	out.SupportsDryRun = direct.ValueOf(in.SupportsDryRun)
	return out
}
func Constraint_BooleanConstraint_FromProto(mapCtx *direct.MapContext, in *pb.Constraint_BooleanConstraint) *krm.Constraint_BooleanConstraint {
	if in == nil {
		return nil
	}
	out := &krm.Constraint_BooleanConstraint{}
	return out
}
func Constraint_BooleanConstraint_ToProto(mapCtx *direct.MapContext, in *krm.Constraint_BooleanConstraint) *pb.Constraint_BooleanConstraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint_BooleanConstraint{}
	return out
}
func Constraint_ListConstraint_FromProto(mapCtx *direct.MapContext, in *pb.Constraint_ListConstraint) *krm.Constraint_ListConstraint {
	if in == nil {
		return nil
	}
	out := &krm.Constraint_ListConstraint{}
	out.SupportsIn = direct.LazyPtr(in.GetSupportsIn())
	out.SupportsUnder = direct.LazyPtr(in.GetSupportsUnder())
	return out
}
func Constraint_ListConstraint_ToProto(mapCtx *direct.MapContext, in *krm.Constraint_ListConstraint) *pb.Constraint_ListConstraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint_ListConstraint{}
	out.SupportsIn = direct.ValueOf(in.SupportsIn)
	out.SupportsUnder = direct.ValueOf(in.SupportsUnder)
	return out
}
func OrgpolicyConstraintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Constraint) *krm.OrgpolicyConstraintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgpolicyConstraintObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ConstraintDefault
	// MISSING: ListConstraint
	// MISSING: BooleanConstraint
	// MISSING: SupportsDryRun
	return out
}
func OrgpolicyConstraintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgpolicyConstraintObservedState) *pb.Constraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ConstraintDefault
	// MISSING: ListConstraint
	// MISSING: BooleanConstraint
	// MISSING: SupportsDryRun
	return out
}
func OrgpolicyConstraintSpec_FromProto(mapCtx *direct.MapContext, in *pb.Constraint) *krm.OrgpolicyConstraintSpec {
	if in == nil {
		return nil
	}
	out := &krm.OrgpolicyConstraintSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ConstraintDefault
	// MISSING: ListConstraint
	// MISSING: BooleanConstraint
	// MISSING: SupportsDryRun
	return out
}
func OrgpolicyConstraintSpec_ToProto(mapCtx *direct.MapContext, in *krm.OrgpolicyConstraintSpec) *pb.Constraint {
	if in == nil {
		return nil
	}
	out := &pb.Constraint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ConstraintDefault
	// MISSING: ListConstraint
	// MISSING: BooleanConstraint
	// MISSING: SupportsDryRun
	return out
}
