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

// proto.message: google.identity.accesscontextmanager.v1.ServicePerimeter
// api.group: accesscontextmanager.cnrm.cloud.google.com

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"google.golang.org/genproto/googleapis/type/expr"

	acm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accesscontextmanager/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessLevelExpr_FromProto(mapCtx *direct.MapContext, in *expr.Expr) acm.AccessLevelExpr {
	out := acm.AccessLevelExpr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}

func AccessLevelExpr_ToProto(mapCtx *direct.MapContext, in acm.AccessLevelExpr) *expr.Expr {
	out := &expr.Expr{}
	out.Description = direct.ValueOf(in.Description)
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	return out
}

func Condition_RequiredAccessLevels_FromProto(mapCtx *direct.MapContext, input []string) []acm.AccessLevelRef {
	if input == nil {
		return nil
	}
	out := make([]acm.AccessLevelRef, len(input))
	for index, item := range input {
		// TODO: determine the format of item and deserialize to Ref
		var element acm.AccessLevelRef
		element.External = direct.LazyPtr(item)

		//	out.Kind = direct.LazyPtr(in.GetKind())
		//	out.Name = direct.LazyPtr(in.GetName())
		//	out.Namespace = direct.LazyPtr(in.GetNamespace())

		out[index] = element
	}
	return out
}

func Condition_RequiredAccessLevels_ToProto(mapCtx *direct.MapContext, in []acm.AccessLevelRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element string
		element = direct.ValueOf(item.External)

		//	element.Kind = direct.ValueOf(in.Kind)
		//	element.Name = direct.ValueOf(in.Name)
		//	element.Namespace = direct.ValueOf(in.Namespace)

		out[index] = element
	}
	return out
}

func Condition_Members_FromProto(mapCtx *direct.MapContext, in []string) []acm.Member {
	if in == nil {
		return nil
	}
	out := make([]acm.Member, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element acm.Member
		element.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: item}
		/*
			element.ServiceAccountRef.Name
			element.ServiceAccountRef.Namespace
			element.User = direct.LazyPtr(in.GetUser())
		*/
		out[index] = element
	}
	return out
}

func Condition_Members_ToProto(mapCtx *direct.MapContext, in []acm.Member) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for index, item := range in {
		// TODO: determine the format of item and deserialize to Ref
		var element string
		if item.ServiceAccountRef != nil {
			element = item.ServiceAccountRef.External
		} else if item.User != nil {
			element = *item.User
		}
		/*
			element.ServiceAccountRef.Name
			element.ServiceAccountRef.Namespace
			out.User = direct.ValueOf(in.User)
		*/
		out[index] = element
	}
	return out
}

func AccessContextManagerServicePerimeterConfig_Resources_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessContextManagerServicePerimeterResource {
	if in == nil {
		return nil
	}
	out := make([]acm.AccessContextManagerServicePerimeterResource, len(in))
	for i, r := range in {
		out[i] = acm.AccessContextManagerServicePerimeterResource{
			ProjectRef: &acm.ServicePerimeterProjectRef{
				External: r,
			},
		}
	}
	return out
}

func AccessContextManagerServicePerimeterConfig_Resources_ToProto(mapCtx *direct.MapContext, in []acm.AccessContextManagerServicePerimeterResource) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.ProjectRef != nil {
			out[i] = r.ProjectRef.External
		}
	}
	return out
}

func AccessContextManagerServicePerimeterConfig_AccessLevels_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessLevelRef {
	if in == nil {
		return nil
	}
	out := make([]acm.AccessLevelRef, len(in))
	for i, r := range in {
		out[i] = acm.AccessLevelRef{
			External: direct.LazyPtr(r),
		}
	}
	return out
}

func AccessContextManagerServicePerimeterConfig_AccessLevels_ToProto(mapCtx *direct.MapContext, in []acm.AccessLevelRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		out[i] = direct.ValueOf(r.External)
	}
	return out
}

func AccessContextManagerServicePerimeterEgressFrom_Identities_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessContextManagerServicePerimeterIdentity {
	if in == nil {
		return nil
	}
	out := make([]acm.AccessContextManagerServicePerimeterIdentity, len(in))
	for i, r := range in {
		out[i] = acm.AccessContextManagerServicePerimeterIdentity{
			ServiceAccountRef: &refsv1beta1.IAMServiceAccountRef{
				External: r,
			},
		}
	}
	return out
}

func AccessContextManagerServicePerimeterEgressFrom_Identities_ToProto(mapCtx *direct.MapContext, in []acm.AccessContextManagerServicePerimeterIdentity) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.ServiceAccountRef != nil {
			out[i] = r.ServiceAccountRef.External
		} else if r.User != nil {
			out[i] = *r.User
		}
	}
	return out
}

func AccessContextManagerServicePerimeterEgressTo_Resources_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessContextManagerServicePerimeterEgressResource {
	if in == nil {
		return nil
	}
	out := make([]acm.AccessContextManagerServicePerimeterEgressResource, len(in))
	for i, r := range in {
		out[i] = acm.AccessContextManagerServicePerimeterEgressResource{
			ProjectRef: &acm.ServicePerimeterProjectRef{
				External: r,
			},
		}
	}
	return out
}

func AccessContextManagerServicePerimeterEgressTo_Resources_ToProto(mapCtx *direct.MapContext, in []acm.AccessContextManagerServicePerimeterEgressResource) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.ProjectRef != nil {
			out[i] = r.ProjectRef.External
		}
	}
	return out
}

func AccessContextManagerServicePerimeterIngressFrom_Identities_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessContextManagerServicePerimeterIdentity {
	return AccessContextManagerServicePerimeterEgressFrom_Identities_FromProto(mapCtx, in)
}

func AccessContextManagerServicePerimeterIngressFrom_Identities_ToProto(mapCtx *direct.MapContext, in []acm.AccessContextManagerServicePerimeterIdentity) []string {
	return AccessContextManagerServicePerimeterEgressFrom_Identities_ToProto(mapCtx, in)
}

func AccessContextManagerServicePerimeterIngressTo_Resources_FromProto(mapCtx *direct.MapContext, in []string) []acm.AccessContextManagerServicePerimeterIngressResource {
	if in == nil {
		return nil
	}
	out := make([]acm.AccessContextManagerServicePerimeterIngressResource, len(in))
	for i, r := range in {
		out[i] = acm.AccessContextManagerServicePerimeterIngressResource{
			ProjectRef: &acm.ServicePerimeterProjectRef{
				External: r,
			},
		}
	}
	return out
}

func AccessContextManagerServicePerimeterIngressTo_Resources_ToProto(mapCtx *direct.MapContext, in []acm.AccessContextManagerServicePerimeterIngressResource) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, r := range in {
		if r.ProjectRef != nil {
			out[i] = r.ProjectRef.External
		}
	}
	return out
}

func AccessContextManagerServicePerimeterIngressSource_FromProto(mapCtx *direct.MapContext, in *pb.ServicePerimeterConfig_IngressSource) *acm.AccessContextManagerServicePerimeterIngressSource {
	if in == nil {
		return nil
	}
	out := &acm.AccessContextManagerServicePerimeterIngressSource{}
	if in.GetAccessLevel() != "" {
		out.AccessLevelRef = &acm.AccessLevelRef{External: direct.LazyPtr(in.GetAccessLevel())}
	}
	if in.GetResource() != "" {
		out.ProjectRef = &acm.ServicePerimeterProjectRef{External: in.GetResource()}
	}
	return out
}

func AccessContextManagerServicePerimeterIngressSource_ToProto(mapCtx *direct.MapContext, in *acm.AccessContextManagerServicePerimeterIngressSource) *pb.ServicePerimeterConfig_IngressSource {
	if in == nil {
		return nil
	}
	out := &pb.ServicePerimeterConfig_IngressSource{}
	if in.AccessLevelRef != nil {
		out.Source = &pb.ServicePerimeterConfig_IngressSource_AccessLevel{
			AccessLevel: direct.ValueOf(in.AccessLevelRef.External),
		}
	}
	if in.ProjectRef != nil {
		out.Source = &pb.ServicePerimeterConfig_IngressSource_Resource{
			Resource: in.ProjectRef.External,
		}
	}
	return out
}
