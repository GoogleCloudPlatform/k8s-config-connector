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
// krm.group: accesscontextmanager.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.identity.accesscontextmanager.v1

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	typepb "google.golang.org/genproto/googleapis/identity/accesscontextmanager/type"
)

func AccessContextManagerAccessLevelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessLevel) *krm.AccessContextManagerAccessLevelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerAccessLevelObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AccessContextManagerAccessLevelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessContextManagerAccessLevelObservedState) *pb.AccessLevel {
	if in == nil {
		return nil
	}
	out := &pb.AccessLevel{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AccessContextManagerAccessLevelSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessLevel) *krm.AccessContextManagerAccessLevelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerAccessLevelSpec{}
	// MISSING: Name
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Basic = BasicLevel_FromProto(mapCtx, in.GetBasic())
	out.Custom = CustomLevel_FromProto(mapCtx, in.GetCustom())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AccessContextManagerAccessLevelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AccessContextManagerAccessLevelSpec) *pb.AccessLevel {
	if in == nil {
		return nil
	}
	out := &pb.AccessLevel{}
	// MISSING: Name
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	if oneof := BasicLevel_ToProto(mapCtx, in.Basic); oneof != nil {
		out.Level = &pb.AccessLevel_Basic{Basic: oneof}
	}
	if oneof := CustomLevel_ToProto(mapCtx, in.Custom); oneof != nil {
		out.Level = &pb.AccessLevel_Custom{Custom: oneof}
	}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AccessContextManagerAccessPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krm.AccessContextManagerAccessPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerAccessPolicyObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessContextManagerAccessPolicyObservedState) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessPolicy) *krm.AccessContextManagerAccessPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.AccessContextManagerAccessPolicySpec{}
	// MISSING: Name
	// MISSING: Parent
	out.Title = direct.LazyPtr(in.GetTitle())
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AccessContextManagerAccessPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.AccessContextManagerAccessPolicySpec) *pb.AccessPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AccessPolicy{}
	// MISSING: Name
	// MISSING: Parent
	out.Title = direct.ValueOf(in.Title)
	// MISSING: Scopes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func BasicLevel_FromProto(mapCtx *direct.MapContext, in *pb.BasicLevel) *krm.BasicLevel {
	if in == nil {
		return nil
	}
	out := &krm.BasicLevel{}
	out.Conditions = direct.Slice_FromProto(mapCtx, in.Conditions, Condition_FromProto)
	out.CombiningFunction = direct.Enum_FromProto(mapCtx, in.GetCombiningFunction())
	return out
}
func BasicLevel_ToProto(mapCtx *direct.MapContext, in *krm.BasicLevel) *pb.BasicLevel {
	if in == nil {
		return nil
	}
	out := &pb.BasicLevel{}
	out.Conditions = direct.Slice_ToProto(mapCtx, in.Conditions, Condition_ToProto)
	out.CombiningFunction = direct.Enum_ToProto[pb.BasicLevel_ConditionCombiningFunction](mapCtx, in.CombiningFunction)
	return out
}
func Condition_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krm.Condition {
	if in == nil {
		return nil
	}
	out := &krm.Condition{}
	out.IPSubnetworks = in.IpSubnetworks
	out.DevicePolicy = DevicePolicy_FromProto(mapCtx, in.GetDevicePolicy())
	out.RequiredAccessLevels = in.RequiredAccessLevels
	out.Negate = direct.LazyPtr(in.GetNegate())
	out.Members = in.Members
	out.Regions = in.Regions
	return out
}
func Condition_ToProto(mapCtx *direct.MapContext, in *krm.Condition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	out.IpSubnetworks = in.IPSubnetworks
	out.DevicePolicy = DevicePolicy_ToProto(mapCtx, in.DevicePolicy)
	out.RequiredAccessLevels = in.RequiredAccessLevels
	out.Negate = direct.ValueOf(in.Negate)
	out.Members = in.Members
	out.Regions = in.Regions
	return out
}
func CustomLevel_FromProto(mapCtx *direct.MapContext, in *pb.CustomLevel) *krm.CustomLevel {
	if in == nil {
		return nil
	}
	out := &krm.CustomLevel{}
	out.Expr = Expr_FromProto(mapCtx, in.GetExpr())
	return out
}
func CustomLevel_ToProto(mapCtx *direct.MapContext, in *krm.CustomLevel) *pb.CustomLevel {
	if in == nil {
		return nil
	}
	out := &pb.CustomLevel{}
	out.Expr = Expr_ToProto(mapCtx, in.Expr)
	return out
}
func DevicePolicy_FromProto(mapCtx *direct.MapContext, in *pb.DevicePolicy) *krm.DevicePolicy {
	if in == nil {
		return nil
	}
	out := &krm.DevicePolicy{}
	out.RequireScreenlock = direct.LazyPtr(in.GetRequireScreenlock())
	out.AllowedEncryptionStatuses = direct.EnumSlice_FromProto(mapCtx, in.AllowedEncryptionStatuses)
	out.OSConstraints = direct.Slice_FromProto(mapCtx, in.OSConstraints, OSConstraint_FromProto)
	out.AllowedDeviceManagementLevels = direct.EnumSlice_FromProto(mapCtx, in.AllowedDeviceManagementLevels)
	out.RequireAdminApproval = direct.LazyPtr(in.GetRequireAdminApproval())
	out.RequireCorpOwned = direct.LazyPtr(in.GetRequireCorpOwned())
	return out
}
func DevicePolicy_ToProto(mapCtx *direct.MapContext, in *krm.DevicePolicy) *pb.DevicePolicy {
	if in == nil {
		return nil
	}
	out := &pb.DevicePolicy{}
	out.RequireScreenlock = direct.ValueOf(in.RequireScreenlock)
	out.AllowedEncryptionStatuses = direct.EnumSlice_ToProto[typepb.DeviceEncryptionStatus](mapCtx, in.AllowedEncryptionStatuses)
	out.OsConstraints = direct.Slice_ToProto(mapCtx, in.OSConstraints, OSConstraint_ToProto)
	out.AllowedDeviceManagementLevels = direct.EnumSlice_ToProto[typepb.DeviceManagementLevel](mapCtx, in.AllowedDeviceManagementLevels)
	out.RequireAdminApproval = direct.ValueOf(in.RequireAdminApproval)
	out.RequireCorpOwned = direct.ValueOf(in.RequireCorpOwned)
	return out
}
func OSConstraint_FromProto(mapCtx *direct.MapContext, in *pb.OsConstraint) *krm.OSConstraint {
	if in == nil {
		return nil
	}
	out := &krm.OSConstraint{}
	out.OSType = direct.Enum_FromProto(mapCtx, in.GetOsType())
	out.MinimumVersion = direct.LazyPtr(in.GetMinimumVersion())
	out.RequireVerifiedChromeOS = direct.LazyPtr(in.GetRequireVerifiedChromeOs())
	return out
}
func OSConstraint_ToProto(mapCtx *direct.MapContext, in *krm.OSConstraint) *pb.OsConstraint {
	if in == nil {
		return nil
	}
	out := &pb.OsConstraint{}
	out.OsType = direct.Enum_ToProto[typepb.OsType](mapCtx, in.OSType)
	out.MinimumVersion = direct.ValueOf(in.MinimumVersion)
	out.RequireVerifiedChromeOs = direct.ValueOf(in.RequireVerifiedChromeOS)
	return out
}
