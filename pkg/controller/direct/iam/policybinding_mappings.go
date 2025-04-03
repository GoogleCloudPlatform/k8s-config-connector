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
// krm.group: iam.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.iam.v3

package iam

import (
	pb "cloud.google.com/go/iam/apiv3/iampb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/type/expr"
)

func Expr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) *krmv1alpha1.Expr {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Expr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}

func Expr_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Expr) *expr.Expr {
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

func PolicyBinding_Target_PrincipalSet_ToProto(mapCtx *direct.MapContext, in *string) *pb.PolicyBinding_Target_PrincipalSet {
	if in == nil {
		return nil
	}
	out := &pb.PolicyBinding_Target_PrincipalSet{
		PrincipalSet: direct.ValueOf(in),
	}
	return out
}
