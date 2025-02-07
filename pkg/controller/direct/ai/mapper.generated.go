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

package ai

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
)
func AiPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionObservedState{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionObservedState) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionSpec{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionSpec) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func Permission_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.Permission {
	if in == nil {
		return nil
	}
	out := &krm.Permission{}
	// MISSING: Name
	out.GranteeType = direct.Enum_FromProto(mapCtx, in.GetGranteeType())
	out.EmailAddress = in.EmailAddress
	out.Role = direct.Enum_FromProto(mapCtx, in.GetRole())
	return out
}
func Permission_ToProto(mapCtx *direct.MapContext, in *krm.Permission) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	if oneof := Permission_GranteeType_ToProto(mapCtx, in.GranteeType); oneof != nil {
		out.GranteeType = oneof
	}
	out.EmailAddress = in.EmailAddress
	if oneof := Permission_Role_ToProto(mapCtx, in.Role); oneof != nil {
		out.Role = oneof
	}
	return out
}
func PermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.PermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PermissionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func PermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PermissionObservedState) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
