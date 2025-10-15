// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"google.golang.org/genproto/googleapis/type/expr"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.PolicySpec) *krm.PolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.PolicySpec{}
	// MISSING: Etag
	// MISSING: UpdateTime
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, PolicySpec_PolicyRule_FromProto)
	out.InheritFromParent = direct.LazyPtr(in.GetInheritFromParent())
	// Expected code `in.GetReset()`, actual code `in.GetReset_()`
	out.Reset = direct.LazyPtr(in.GetReset_())
	return out
}
func PolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.PolicySpec) *pb.PolicySpec {
	if in == nil {
		return nil
	}
	out := &pb.PolicySpec{}
	// MISSING: Etag
	// MISSING: UpdateTime
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, PolicySpec_PolicyRule_ToProto)
	out.InheritFromParent = direct.ValueOf(in.InheritFromParent)
	// Expected code `in.GetReset()`, actual code `in.GetReset_()`
	out.Reset_ = direct.ValueOf(in.Reset)
	return out
}

func Expr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) *krm.Expr {
	if in == nil {
		return nil
	}
	out := &krm.Expr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krm.Expr) *expr.Expr {
	if in == nil {
		return nil
	}
	out := &expr.Expr{}
	out.Description = direct.ValueOf(in.Description)
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	return out
}

func PolicySpec_PolicyRule_AllowAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_AllowAll {
	if in == nil {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_AllowAll{AllowAll: *in}
}
func PolicySpec_PolicyRule_DenyAll_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_DenyAll {
	if in == nil {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_DenyAll{DenyAll: *in}
}
func PolicySpec_PolicyRule_Enforce_ToProto(mapCtx *direct.MapContext, in *bool) *pb.PolicySpec_PolicyRule_Enforce {
	if in == nil {
		return nil
	}
	return &pb.PolicySpec_PolicyRule_Enforce{Enforce: *in}
}
