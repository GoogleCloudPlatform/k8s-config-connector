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

package networkservices

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func HttpRoute_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.HttpRoute {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Hostnames = in.Hostnames
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Labels = in.Labels
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, HttpRoute_RouteRule_FromProto)
	return out
}
func HttpRoute_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Hostnames = in.Hostnames
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Labels = in.Labels
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, HttpRoute_RouteRule_ToProto)
	return out
}
func HttpRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.HttpRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HttpRouteObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
func HttpRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HttpRouteObservedState) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
func HttpRoute_CorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_CorsPolicy) *krm.HttpRoute_CorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_CorsPolicy{}
	out.AllowOrigins = in.AllowOrigins
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.MaxAge = direct.LazyPtr(in.GetMaxAge())
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func HttpRoute_CorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_CorsPolicy) *pb.HttpRoute_CorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_CorsPolicy{}
	out.AllowOrigins = in.AllowOrigins
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.MaxAge = direct.ValueOf(in.MaxAge)
	out.AllowCredentials = direct.ValueOf(in.AllowCredentials)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func HttpRoute_Destination_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krm.HttpRoute_Destination {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_Destination{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.Weight = direct.LazyPtr(in.GetWeight())
	return out
}
func HttpRoute_Destination_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_Destination) *pb.HttpRoute_Destination {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_Destination{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.Weight = direct.ValueOf(in.Weight)
	return out
}
func HttpRoute_FaultInjectionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy) *krm.HttpRoute_FaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_FaultInjectionPolicy{}
	out.Delay = HttpRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx, in.GetDelay())
	out.Abort = HttpRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx, in.GetAbort())
	return out
}
func HttpRoute_FaultInjectionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_FaultInjectionPolicy) *pb.HttpRoute_FaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy{}
	out.Delay = HttpRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx, in.Delay)
	out.Abort = HttpRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx, in.Abort)
	return out
}
func HttpRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Abort) *krm.HttpRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_FaultInjectionPolicy_Abort{}
	out.HTTPStatus = direct.LazyPtr(in.GetHttpStatus())
	out.Percentage = direct.LazyPtr(in.GetPercentage())
	return out
}
func HttpRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_FaultInjectionPolicy_Abort) *pb.HttpRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Abort{}
	out.HttpStatus = direct.ValueOf(in.HTTPStatus)
	out.Percentage = direct.ValueOf(in.Percentage)
	return out
}
func HttpRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Delay) *krm.HttpRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = direct.LazyPtr(in.GetPercentage())
	return out
}
func HttpRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_FaultInjectionPolicy_Delay) *pb.HttpRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_ToProto(mapCtx, in.FixedDelay)
	out.Percentage = direct.ValueOf(in.Percentage)
	return out
}
func HttpRoute_HeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch) *krm.HttpRoute_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_HeaderMatch{}
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	out.RangeMatch = HttpRoute_HeaderMatch_IntegerRange_FromProto(mapCtx, in.GetRangeMatch())
	out.Header = direct.LazyPtr(in.GetHeader())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	return out
}
func HttpRoute_HeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_HeaderMatch) *pb.HttpRoute_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderMatch{}
	if oneof := HttpRoute_HeaderMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_HeaderMatch_RegexMatch_ToProto(mapCtx, in.RegexMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_HeaderMatch_PrefixMatch_ToProto(mapCtx, in.PrefixMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_HeaderMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_HeaderMatch_SuffixMatch_ToProto(mapCtx, in.SuffixMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_HeaderMatch_IntegerRange_ToProto(mapCtx, in.RangeMatch); oneof != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_RangeMatch{RangeMatch: oneof}
	}
	out.Header = direct.ValueOf(in.Header)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	return out
}
func HttpRoute_HeaderMatch_IntegerRange_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch_IntegerRange) *krm.HttpRoute_HeaderMatch_IntegerRange {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_HeaderMatch_IntegerRange{}
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	return out
}
func HttpRoute_HeaderMatch_IntegerRange_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_HeaderMatch_IntegerRange) *pb.HttpRoute_HeaderMatch_IntegerRange {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderMatch_IntegerRange{}
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	return out
}
func HttpRoute_HeaderModifier_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderModifier) *krm.HttpRoute_HeaderModifier {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_HeaderModifier{}
	out.Set = in.Set
	out.Add = in.Add
	out.Remove = in.Remove
	return out
}
func HttpRoute_HeaderModifier_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_HeaderModifier) *pb.HttpRoute_HeaderModifier {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderModifier{}
	out.Set = in.Set
	out.Add = in.Add
	out.Remove = in.Remove
	return out
}
func HttpRoute_QueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_QueryParameterMatch) *krm.HttpRoute_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_QueryParameterMatch{}
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.QueryParameter = direct.LazyPtr(in.GetQueryParameter())
	return out
}
func HttpRoute_QueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_QueryParameterMatch) *pb.HttpRoute_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_QueryParameterMatch{}
	if oneof := HttpRoute_QueryParameterMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_QueryParameterMatch_RegexMatch_ToProto(mapCtx, in.RegexMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HttpRoute_QueryParameterMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	out.QueryParameter = direct.ValueOf(in.QueryParameter)
	return out
}
func HttpRoute_Redirect_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Redirect) *krm.HttpRoute_Redirect {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_Redirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRewrite = direct.LazyPtr(in.GetPrefixRewrite())
	out.ResponseCode = direct.Enum_FromProto(mapCtx, in.GetResponseCode())
	out.HTTPSRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	out.PortRedirect = direct.LazyPtr(in.GetPortRedirect())
	return out
}
func HttpRoute_Redirect_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_Redirect) *pb.HttpRoute_Redirect {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_Redirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRewrite = direct.ValueOf(in.PrefixRewrite)
	out.ResponseCode = direct.Enum_ToProto[pb.HttpRoute_Redirect_ResponseCode](mapCtx, in.ResponseCode)
	out.HttpsRedirect = direct.ValueOf(in.HTTPSRedirect)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	out.PortRedirect = direct.ValueOf(in.PortRedirect)
	return out
}
func HttpRoute_RequestMirrorPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RequestMirrorPolicy) *krm.HttpRoute_RequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_RequestMirrorPolicy{}
	out.Destination = HttpRoute_Destination_FromProto(mapCtx, in.GetDestination())
	return out
}
func HttpRoute_RequestMirrorPolicy_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_RequestMirrorPolicy) *pb.HttpRoute_RequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RequestMirrorPolicy{}
	out.Destination = HttpRoute_Destination_ToProto(mapCtx, in.Destination)
	return out
}
func HttpRoute_RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RetryPolicy) *krm.HttpRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = direct.LazyPtr(in.GetNumRetries())
	out.PerTryTimeout = direct.StringDuration_FromProto(mapCtx, in.GetPerTryTimeout())
	return out
}
func HttpRoute_RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_RetryPolicy) *pb.HttpRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = direct.ValueOf(in.NumRetries)
	out.PerTryTimeout = direct.StringDuration_ToProto(mapCtx, in.PerTryTimeout)
	return out
}
func HttpRoute_RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RouteAction) *krm.HttpRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_RouteAction{}
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, HttpRoute_Destination_FromProto)
	out.Redirect = HttpRoute_Redirect_FromProto(mapCtx, in.GetRedirect())
	out.FaultInjectionPolicy = HttpRoute_FaultInjectionPolicy_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.RequestHeaderModifier = HttpRoute_HeaderModifier_FromProto(mapCtx, in.GetRequestHeaderModifier())
	out.ResponseHeaderModifier = HttpRoute_HeaderModifier_FromProto(mapCtx, in.GetResponseHeaderModifier())
	out.URLRewrite = HttpRoute_URLRewrite_FromProto(mapCtx, in.GetUrlRewrite())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.RetryPolicy = HttpRoute_RetryPolicy_FromProto(mapCtx, in.GetRetryPolicy())
	out.RequestMirrorPolicy = HttpRoute_RequestMirrorPolicy_FromProto(mapCtx, in.GetRequestMirrorPolicy())
	out.CorsPolicy = HttpRoute_CorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	return out
}
func HttpRoute_RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_RouteAction) *pb.HttpRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RouteAction{}
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, HttpRoute_Destination_ToProto)
	out.Redirect = HttpRoute_Redirect_ToProto(mapCtx, in.Redirect)
	out.FaultInjectionPolicy = HttpRoute_FaultInjectionPolicy_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestHeaderModifier = HttpRoute_HeaderModifier_ToProto(mapCtx, in.RequestHeaderModifier)
	out.ResponseHeaderModifier = HttpRoute_HeaderModifier_ToProto(mapCtx, in.ResponseHeaderModifier)
	out.UrlRewrite = HttpRoute_URLRewrite_ToProto(mapCtx, in.URLRewrite)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.RetryPolicy = HttpRoute_RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	out.RequestMirrorPolicy = HttpRoute_RequestMirrorPolicy_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.CorsPolicy = HttpRoute_CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	return out
}
func HttpRoute_RouteMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RouteMatch) *krm.HttpRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_RouteMatch{}
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.Headers = direct.Slice_FromProto(mapCtx, in.Headers, HttpRoute_HeaderMatch_FromProto)
	out.QueryParameters = direct.Slice_FromProto(mapCtx, in.QueryParameters, HttpRoute_QueryParameterMatch_FromProto)
	return out
}
func HttpRoute_RouteMatch_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_RouteMatch) *pb.HttpRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RouteMatch{}
	if oneof := HttpRoute_RouteMatch_FullPathMatch_ToProto(mapCtx, in.FullPathMatch); oneof != nil {
		out.PathMatch = oneof
	}
	if oneof := HttpRoute_RouteMatch_PrefixMatch_ToProto(mapCtx, in.PrefixMatch); oneof != nil {
		out.PathMatch = oneof
	}
	if oneof := HttpRoute_RouteMatch_RegexMatch_ToProto(mapCtx, in.RegexMatch); oneof != nil {
		out.PathMatch = oneof
	}
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.Headers = direct.Slice_ToProto(mapCtx, in.Headers, HttpRoute_HeaderMatch_ToProto)
	out.QueryParameters = direct.Slice_ToProto(mapCtx, in.QueryParameters, HttpRoute_QueryParameterMatch_ToProto)
	return out
}
func HttpRoute_RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RouteRule) *krm.HttpRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_RouteRule{}
	out.Matches = direct.Slice_FromProto(mapCtx, in.Matches, HttpRoute_RouteMatch_FromProto)
	out.Action = HttpRoute_RouteAction_FromProto(mapCtx, in.GetAction())
	return out
}
func HttpRoute_RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_RouteRule) *pb.HttpRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RouteRule{}
	out.Matches = direct.Slice_ToProto(mapCtx, in.Matches, HttpRoute_RouteMatch_ToProto)
	out.Action = HttpRoute_RouteAction_ToProto(mapCtx, in.Action)
	return out
}
func HttpRoute_URLRewrite_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_URLRewrite) *krm.HttpRoute_URLRewrite {
	if in == nil {
		return nil
	}
	out := &krm.HttpRoute_URLRewrite{}
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	return out
}
func HttpRoute_URLRewrite_ToProto(mapCtx *direct.MapContext, in *krm.HttpRoute_URLRewrite) *pb.HttpRoute_URLRewrite {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_URLRewrite{}
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	return out
}
func NetworkservicesHttpRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.NetworkservicesHttpRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesHttpRouteObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
func NetworkservicesHttpRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesHttpRouteObservedState) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
func NetworkservicesHttpRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.NetworkservicesHttpRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesHttpRouteSpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
func NetworkservicesHttpRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesHttpRouteSpec) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Labels
	// MISSING: Rules
	return out
}
