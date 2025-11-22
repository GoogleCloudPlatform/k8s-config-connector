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

// +tool:fuzz-gen
// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package accesscontextmanager

import (
	//pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"google.golang.org/genproto/googleapis/type/expr"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1beta1"
	oldrefs "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

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

func Condition_RequiredAccessLevels_FromProto(mapCtx *direct.MapContext, in []string) []oldrefs.ResourceRef {
	if in == nil {
		return nil
	}
	out := make([]oldrefs.ResourceRef, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element oldrefs.ResourceRef
		element.External = item
		/*
			out.Kind = direct.LazyPtr(in.GetKind())
			out.Name = direct.LazyPtr(in.GetName())
			out.Namespace = direct.LazyPtr(in.GetNamespace())
		*/
		out[index] = element
	}
	return out
}

func Condition_RequiredAccessLevels_ToProto(mapCtx *direct.MapContext, in []oldrefs.ResourceRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element string
		element = item.External
		/*
			element.Kind = direct.ValueOf(in.Kind)
			element.Name = direct.ValueOf(in.Name)
			element.Namespace = direct.ValueOf(in.Namespace)
		*/
		out[index] = element
	}
	return out
}

func Condition_Members_FromProto(mapCtx *direct.MapContext, in []string) []krm.Member {
	if in == nil {
		return nil
	}
	out := make([]krm.Member, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element krm.Member
		element.ServiceAccountRef.External = item
		/*
			element.ServiceAccountRef.Name
			element.ServiceAccountRef.Namespace
			element.User = direct.LazyPtr(in.GetUser())
		*/
		out[index] = element
	}
	return out
}

func Condition_Members_ToProto(mapCtx *direct.MapContext, in []krm.Member) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element string
		element = item.ServiceAccountRef.External
		/*
			element.ServiceAccountRef.Name
			element.ServiceAccountRef.Namespace
			out.User = direct.ValueOf(in.User)
		*/
		out[index] = element
	}
	return out
}
