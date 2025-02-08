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

package identity

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/identity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AccessLevel_FromProto(mapCtx *direct.MapContext, in *pb.AccessLevel) *krm.AccessLevel {
	if in == nil {
		return nil
	}
	out := &krm.AccessLevel{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Basic = BasicLevel_FromProto(mapCtx, in.GetBasic())
	out.Custom = CustomLevel_FromProto(mapCtx, in.GetCustom())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func AccessLevel_ToProto(mapCtx *direct.MapContext, in *krm.AccessLevel) *pb.AccessLevel {
	if in == nil {
		return nil
	}
	out := &pb.AccessLevel{}
	out.Name = direct.ValueOf(in.Name)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	if oneof := BasicLevel_ToProto(mapCtx, in.Basic); oneof != nil {
		out.Level = &pb.AccessLevel_Basic{Basic: oneof}
	}
	if oneof := CustomLevel_ToProto(mapCtx, in.Custom); oneof != nil {
		out.Level = &pb.AccessLevel_Custom{Custom: oneof}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
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
	out.OsConstraints = direct.Slice_FromProto(mapCtx, in.OsConstraints, OsConstraint_FromProto)
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
	out.AllowedEncryptionStatuses = direct.EnumSlice_ToProto[pb.DeviceEncryptionStatus](mapCtx, in.AllowedEncryptionStatuses)
	out.OsConstraints = direct.Slice_ToProto(mapCtx, in.OsConstraints, OsConstraint_ToProto)
	out.AllowedDeviceManagementLevels = direct.EnumSlice_ToProto[pb.DeviceManagementLevel](mapCtx, in.AllowedDeviceManagementLevels)
	out.RequireAdminApproval = direct.ValueOf(in.RequireAdminApproval)
	out.RequireCorpOwned = direct.ValueOf(in.RequireCorpOwned)
	return out
}
func IdentityAccessLevelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessLevel) *krm.IdentityAccessLevelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IdentityAccessLevelObservedState{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Basic
	// MISSING: Custom
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func IdentityAccessLevelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IdentityAccessLevelObservedState) *pb.AccessLevel {
	if in == nil {
		return nil
	}
	out := &pb.AccessLevel{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Basic
	// MISSING: Custom
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func IdentityAccessLevelSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessLevel) *krm.IdentityAccessLevelSpec {
	if in == nil {
		return nil
	}
	out := &krm.IdentityAccessLevelSpec{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Basic
	// MISSING: Custom
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func IdentityAccessLevelSpec_ToProto(mapCtx *direct.MapContext, in *krm.IdentityAccessLevelSpec) *pb.AccessLevel {
	if in == nil {
		return nil
	}
	out := &pb.AccessLevel{}
	// MISSING: Name
	// MISSING: Title
	// MISSING: Description
	// MISSING: Basic
	// MISSING: Custom
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func OsConstraint_FromProto(mapCtx *direct.MapContext, in *pb.OsConstraint) *krm.OsConstraint {
	if in == nil {
		return nil
	}
	out := &krm.OsConstraint{}
	out.OsType = direct.Enum_FromProto(mapCtx, in.GetOsType())
	out.MinimumVersion = direct.LazyPtr(in.GetMinimumVersion())
	out.RequireVerifiedChromeOs = direct.LazyPtr(in.GetRequireVerifiedChromeOs())
	return out
}
func OsConstraint_ToProto(mapCtx *direct.MapContext, in *krm.OsConstraint) *pb.OsConstraint {
	if in == nil {
		return nil
	}
	out := &pb.OsConstraint{}
	out.OsType = direct.Enum_ToProto[pb.OsType](mapCtx, in.OsType)
	out.MinimumVersion = direct.ValueOf(in.MinimumVersion)
	out.RequireVerifiedChromeOs = direct.ValueOf(in.RequireVerifiedChromeOs)
	return out
}
