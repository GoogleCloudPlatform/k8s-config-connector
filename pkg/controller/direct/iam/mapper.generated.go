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

package iam

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iam/apiv3beta/iampb"
)
func PolicyBinding_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding) *krm.PolicyBinding {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBinding{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = in.Annotations
	out.Target = PolicyBinding_Target_FromProto(mapCtx, in.GetTarget())
	out.PolicyKind = direct.Enum_FromProto(mapCtx, in.GetPolicyKind())
	out.Policy = direct.LazyPtr(in.GetPolicy())
	// MISSING: PolicyUid
	out.Condition = Expr_FromProto(mapCtx, in.GetCondition())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PolicyBinding_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBinding) *pb.PolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	out.Etag = direct.ValueOf(in.Etag)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = in.Annotations
	out.Target = PolicyBinding_Target_ToProto(mapCtx, in.Target)
	out.PolicyKind = direct.Enum_ToProto[pb.PolicyBinding_PolicyKind](mapCtx, in.PolicyKind)
	out.Policy = direct.ValueOf(in.Policy)
	// MISSING: PolicyUid
	out.Condition = Expr_ToProto(mapCtx, in.Condition)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PolicyBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding) *krm.PolicyBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBindingObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: Target
	// MISSING: PolicyKind
	// MISSING: Policy
	out.PolicyUid = direct.LazyPtr(in.GetPolicyUid())
	// MISSING: Condition
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PolicyBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBindingObservedState) *pb.PolicyBinding {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Etag
	// MISSING: DisplayName
	// MISSING: Annotations
	// MISSING: Target
	// MISSING: PolicyKind
	// MISSING: Policy
	out.PolicyUid = direct.ValueOf(in.PolicyUid)
	// MISSING: Condition
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func PolicyBinding_Target_FromProto(mapCtx *direct.MapContext, in *pb.PolicyBinding_Target) *krm.PolicyBinding_Target {
	if in == nil {
		return nil
	}
	out := &krm.PolicyBinding_Target{}
	out.PrincipalSet = direct.LazyPtr(in.GetPrincipalSet())
	return out
}
func PolicyBinding_Target_ToProto(mapCtx *direct.MapContext, in *krm.PolicyBinding_Target) *pb.PolicyBinding_Target {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding_Target{}
	if oneof := PolicyBinding_Target_PrincipalSet_ToProto(mapCtx, in.PrincipalSet); oneof != nil {
		out.Target = oneof
	}
	return out
}
