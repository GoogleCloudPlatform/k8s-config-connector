// Copyright 2026 Google LLC
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
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesEdgeCacheServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisableHttp2 = direct.LazyPtr(in.GetDisableHttp2())
	out.DisableQuic = direct.LazyPtr(in.GetDisableQuic())
	if in.GetEdgeSecurityPolicy() != "" {
		out.EdgeSecurityPolicyRef = &krm.ResourceRef{External: in.GetEdgeSecurityPolicy()}
	}
	if len(in.GetEdgeSslCertificates()) > 0 {
		out.EdgeSslCertificates = make([]krm.ResourceRef, len(in.GetEdgeSslCertificates()))
		for i, v := range in.GetEdgeSslCertificates() {
			out.EdgeSslCertificates[i] = krm.ResourceRef{External: v}
		}
	}
	out.LogConfig = EdgeCacheServiceLogConfig_FromProto(mapCtx, in.GetLogConfig())
	out.RequireTls = direct.LazyPtr(in.GetRequireTls())
	out.Routing = *EdgeCacheServiceRouting_FromProto(mapCtx, in.GetRouting())
	if in.GetSslPolicy() != "" {
		out.SslPolicyRef = &krm.ResourceRef{External: in.GetSslPolicy()}
	}
	return out
}

func NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Description = direct.ValueOf(in.Description)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	if in.EdgeSecurityPolicyRef != nil {
		out.EdgeSecurityPolicy = in.EdgeSecurityPolicyRef.External
	}
	if len(in.EdgeSslCertificates) > 0 {
		out.EdgeSslCertificates = make([]string, len(in.EdgeSslCertificates))
		for i, v := range in.EdgeSslCertificates {
			out.EdgeSslCertificates[i] = v.External
		}
	}
	out.LogConfig = EdgeCacheServiceLogConfig_ToProto(mapCtx, in.LogConfig)
	out.RequireTls = direct.ValueOf(in.RequireTls)
	if in.Routing.HostRule != nil || in.Routing.PathMatcher != nil { // Check if Routing is empty? struct is not pointer.
		out.Routing = EdgeCacheServiceRouting_ToProto(mapCtx, &in.Routing)
	}
	if in.SslPolicyRef != nil {
		out.SslPolicy = in.SslPolicyRef.External
	}
	return out
}

func NetworkServicesEdgeCacheServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceObservedState{}
	out.Ipv4Addresses = in.GetIpv4Addresses()
	out.Ipv6Addresses = in.GetIpv6Addresses()
	return out
}

func NetworkServicesEdgeCacheServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceObservedState) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	return out
}

// Helpers

func EdgeCacheServiceLogConfig_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_LogConfig) *krm.EdgeCacheServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceLogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	out.SampleRate = direct.LazyPtr(in.GetSampleRate())
	return out
}

func EdgeCacheServiceLogConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceLogConfig) *pb.EdgeCacheService_LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_LogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	out.SampleRate = direct.ValueOf(in.SampleRate)
	return out
}

func EdgeCacheServiceRouting_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing) *krm.EdgeCacheServiceRouting {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceRouting{}
	if in.GetHostRule() != nil {
		out.HostRule = make([]krm.EdgeCacheServiceHostRule, len(in.GetHostRule()))
		for i, v := range in.GetHostRule() {
			out.HostRule[i] = *EdgeCacheServiceHostRule_FromProto(mapCtx, v)
		}
	}
	if in.GetPathMatcher() != nil {
		out.PathMatcher = make([]krm.EdgeCacheServicePathMatcher, len(in.GetPathMatcher()))
		for i, v := range in.GetPathMatcher() {
			out.PathMatcher[i] = *EdgeCacheServicePathMatcher_FromProto(mapCtx, v)
		}
	}
	return out
}

func EdgeCacheServiceRouting_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceRouting) *pb.EdgeCacheService_Routing {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing{}
	if in.HostRule != nil {
		out.HostRule = make([]*pb.EdgeCacheService_Routing_HostRule, len(in.HostRule))
		for i, v := range in.HostRule {
			out.HostRule[i] = EdgeCacheServiceHostRule_ToProto(mapCtx, &v)
		}
	}
	if in.PathMatcher != nil {
		out.PathMatcher = make([]*pb.EdgeCacheService_Routing_PathMatcher, len(in.PathMatcher))
		for i, v := range in.PathMatcher {
			out.PathMatcher[i] = EdgeCacheServicePathMatcher_ToProto(mapCtx, &v)
		}
	}
	return out
}

func EdgeCacheServiceHostRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_HostRule) *krm.EdgeCacheServiceHostRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceHostRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hosts = in.GetHosts()
	out.PathMatcher = in.GetPathMatcher()
	return out
}

func EdgeCacheServiceHostRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceHostRule) *pb.EdgeCacheService_Routing_HostRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_HostRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Hosts = in.Hosts
	out.PathMatcher = in.PathMatcher
	return out
}

func EdgeCacheServicePathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_PathMatcher) *krm.EdgeCacheServicePathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServicePathMatcher{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Name = in.GetName()
	if in.GetRouteRule() != nil {
		out.RouteRule = make([]krm.EdgeCacheServiceRouteRule, len(in.GetRouteRule()))
		for i, v := range in.GetRouteRule() {
			out.RouteRule[i] = *EdgeCacheServiceRouteRule_FromProto(mapCtx, v)
		}
	}
	return out
}

func EdgeCacheServicePathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServicePathMatcher) *pb.EdgeCacheService_Routing_PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_PathMatcher{}
	out.Description = direct.ValueOf(in.Description)
	out.Name = in.Name
	if in.RouteRule != nil {
		out.RouteRule = make([]*pb.EdgeCacheService_Routing_RouteRule, len(in.RouteRule))
		for i, v := range in.RouteRule {
			out.RouteRule[i] = EdgeCacheServiceRouteRule_ToProto(mapCtx, &v)
		}
	}
	return out
}

func EdgeCacheServiceRouteRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule) *krm.EdgeCacheServiceRouteRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceRouteRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.HeaderAction = EdgeCacheServiceHeaderAction_FromProto(mapCtx, in.GetHeaderAction())
	if in.GetMatchRule() != nil {
		out.MatchRule = make([]krm.EdgeCacheServiceMatchRule, len(in.GetMatchRule()))
		for i, v := range in.GetMatchRule() {
			out.MatchRule[i] = *EdgeCacheServiceMatchRule_FromProto(mapCtx, v)
		}
	}
	if in.GetOrigin() != "" {
		out.OriginRef = &krm.ResourceRef{External: in.GetOrigin()}
	}
	out.Priority = in.GetPriority()
	out.RouteAction = EdgeCacheServiceRouteAction_FromProto(mapCtx, in.GetRouteAction())
	out.UrlRedirect = EdgeCacheServiceUrlRedirect_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func EdgeCacheServiceRouteRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceRouteRule) *pb.EdgeCacheService_Routing_RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule{}
	out.Description = direct.ValueOf(in.Description)
	out.HeaderAction = EdgeCacheServiceHeaderAction_ToProto(mapCtx, in.HeaderAction)
	if in.MatchRule != nil {
		out.MatchRule = make([]*pb.EdgeCacheService_Routing_RouteRule_MatchRule, len(in.MatchRule))
		for i, v := range in.MatchRule {
			out.MatchRule[i] = EdgeCacheServiceMatchRule_ToProto(mapCtx, &v)
		}
	}
	if in.OriginRef != nil {
		out.Origin = in.OriginRef.External
	}
	out.Priority = in.Priority
	out.RouteAction = EdgeCacheServiceRouteAction_ToProto(mapCtx, in.RouteAction)
	out.UrlRedirect = EdgeCacheServiceUrlRedirect_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func EdgeCacheServiceHeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_HeaderAction) *krm.EdgeCacheServiceHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceHeaderAction{}
	if in.GetRequestHeaderToAdd() != nil {
		out.RequestHeaderToAdd = make([]krm.EdgeCacheServiceRequestHeaderToAdd, len(in.GetRequestHeaderToAdd()))
		for i, v := range in.GetRequestHeaderToAdd() {
			out.RequestHeaderToAdd[i] = *EdgeCacheServiceRequestHeaderToAdd_FromProto(mapCtx, v)
		}
	}
	if in.GetRequestHeaderToRemove() != nil {
		out.RequestHeaderToRemove = make([]krm.EdgeCacheServiceRequestHeaderToRemove, len(in.GetRequestHeaderToRemove()))
		for i, v := range in.GetRequestHeaderToRemove() {
			out.RequestHeaderToRemove[i] = *EdgeCacheServiceRequestHeaderToRemove_FromProto(mapCtx, v)
		}
	}
	if in.GetResponseHeaderToAdd() != nil {
		out.ResponseHeaderToAdd = make([]krm.EdgeCacheServiceResponseHeaderToAdd, len(in.GetResponseHeaderToAdd()))
		for i, v := range in.GetResponseHeaderToAdd() {
			out.ResponseHeaderToAdd[i] = *EdgeCacheServiceResponseHeaderToAdd_FromProto(mapCtx, v)
		}
	}
	if in.GetResponseHeaderToRemove() != nil {
		out.ResponseHeaderToRemove = make([]krm.EdgeCacheServiceResponseHeaderToRemove, len(in.GetResponseHeaderToRemove()))
		for i, v := range in.GetResponseHeaderToRemove() {
			out.ResponseHeaderToRemove[i] = *EdgeCacheServiceResponseHeaderToRemove_FromProto(mapCtx, v)
		}
	}
	return out
}

func EdgeCacheServiceHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceHeaderAction) *pb.EdgeCacheService_Routing_RouteRule_HeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_HeaderAction{}
	if in.RequestHeaderToAdd != nil {
		out.RequestHeaderToAdd = make([]*pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToAdd, len(in.RequestHeaderToAdd))
		for i, v := range in.RequestHeaderToAdd {
			out.RequestHeaderToAdd[i] = EdgeCacheServiceRequestHeaderToAdd_ToProto(mapCtx, &v)
		}
	}
	if in.RequestHeaderToRemove != nil {
		out.RequestHeaderToRemove = make([]*pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToRemove, len(in.RequestHeaderToRemove))
		for i, v := range in.RequestHeaderToRemove {
			out.RequestHeaderToRemove[i] = EdgeCacheServiceRequestHeaderToRemove_ToProto(mapCtx, &v)
		}
	}
	if in.ResponseHeaderToAdd != nil {
		out.ResponseHeaderToAdd = make([]*pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToAdd, len(in.ResponseHeaderToAdd))
		for i, v := range in.ResponseHeaderToAdd {
			out.ResponseHeaderToAdd[i] = EdgeCacheServiceResponseHeaderToAdd_ToProto(mapCtx, &v)
		}
	}
	if in.ResponseHeaderToRemove != nil {
		out.ResponseHeaderToRemove = make([]*pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToRemove, len(in.ResponseHeaderToRemove))
		for i, v := range in.ResponseHeaderToRemove {
			out.ResponseHeaderToRemove[i] = EdgeCacheServiceResponseHeaderToRemove_ToProto(mapCtx, &v)
		}
	}
	return out
}

func EdgeCacheServiceRequestHeaderToAdd_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToAdd) *krm.EdgeCacheServiceRequestHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceRequestHeaderToAdd{}
	out.HeaderName = in.GetHeaderName()
	out.HeaderValue = in.GetHeaderValue()
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}

func EdgeCacheServiceRequestHeaderToAdd_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceRequestHeaderToAdd) *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToAdd{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func EdgeCacheServiceRequestHeaderToRemove_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToRemove) *krm.EdgeCacheServiceRequestHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceRequestHeaderToRemove{}
	out.HeaderName = in.GetHeaderName()
	return out
}

func EdgeCacheServiceRequestHeaderToRemove_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceRequestHeaderToRemove) *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_HeaderAction_RequestHeaderToRemove{}
	out.HeaderName = in.HeaderName
	return out
}

func EdgeCacheServiceResponseHeaderToAdd_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToAdd) *krm.EdgeCacheServiceResponseHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceResponseHeaderToAdd{}
	out.HeaderName = in.GetHeaderName()
	out.HeaderValue = in.GetHeaderValue()
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}

func EdgeCacheServiceResponseHeaderToAdd_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceResponseHeaderToAdd) *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToAdd{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func EdgeCacheServiceResponseHeaderToRemove_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToRemove) *krm.EdgeCacheServiceResponseHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceResponseHeaderToRemove{}
	out.HeaderName = in.GetHeaderName()
	return out
}

func EdgeCacheServiceResponseHeaderToRemove_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceResponseHeaderToRemove) *pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_HeaderAction_ResponseHeaderToRemove{}
	out.HeaderName = in.HeaderName
	return out
}

func EdgeCacheServiceMatchRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_MatchRule) *krm.EdgeCacheServiceMatchRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceMatchRule{}
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	if in.GetHeaderMatch() != nil {
		out.HeaderMatch = make([]krm.EdgeCacheServiceHeaderMatch, len(in.GetHeaderMatch()))
		for i, v := range in.GetHeaderMatch() {
			out.HeaderMatch[i] = *EdgeCacheServiceHeaderMatch_FromProto(mapCtx, v)
		}
	}
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.PathTemplateMatch = direct.LazyPtr(in.GetPathTemplateMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	if in.GetQueryParameterMatch() != nil {
		out.QueryParameterMatch = make([]krm.EdgeCacheServiceQueryParameterMatch, len(in.GetQueryParameterMatch()))
		for i, v := range in.GetQueryParameterMatch() {
			out.QueryParameterMatch[i] = *EdgeCacheServiceQueryParameterMatch_FromProto(mapCtx, v)
		}
	}
	return out
}

func EdgeCacheServiceMatchRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceMatchRule) *pb.EdgeCacheService_Routing_RouteRule_MatchRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_MatchRule{}
	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)
	if in.HeaderMatch != nil {
		out.HeaderMatch = make([]*pb.EdgeCacheService_Routing_RouteRule_MatchRule_HeaderMatch, len(in.HeaderMatch))
		for i, v := range in.HeaderMatch {
			out.HeaderMatch[i] = EdgeCacheServiceHeaderMatch_ToProto(mapCtx, &v)
		}
	}
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.PathTemplateMatch = direct.ValueOf(in.PathTemplateMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	if in.QueryParameterMatch != nil {
		out.QueryParameterMatch = make([]*pb.EdgeCacheService_Routing_RouteRule_MatchRule_QueryParameterMatch, len(in.QueryParameterMatch))
		for i, v := range in.QueryParameterMatch {
			out.QueryParameterMatch[i] = EdgeCacheServiceQueryParameterMatch_ToProto(mapCtx, &v)
		}
	}
	return out
}

func EdgeCacheServiceHeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_MatchRule_HeaderMatch) *krm.EdgeCacheServiceHeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceHeaderMatch{}
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.HeaderName = in.GetHeaderName()
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	return out
}

func EdgeCacheServiceHeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceHeaderMatch) *pb.EdgeCacheService_Routing_RouteRule_MatchRule_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_MatchRule_HeaderMatch{}
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.HeaderName = in.HeaderName
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	out.SuffixMatch = direct.ValueOf(in.SuffixMatch)
	return out
}

func EdgeCacheServiceQueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_MatchRule_QueryParameterMatch) *krm.EdgeCacheServiceQueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceQueryParameterMatch{}
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.Name = in.GetName()
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	return out
}

func EdgeCacheServiceQueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceQueryParameterMatch) *pb.EdgeCacheService_Routing_RouteRule_MatchRule_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_MatchRule_QueryParameterMatch{}
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.Name = in.Name
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	return out
}

func EdgeCacheServiceRouteAction_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction) *krm.EdgeCacheServiceRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceRouteAction{}
	out.CdnPolicy = EdgeCacheServiceCdnPolicy_FromProto(mapCtx, in.GetCdnPolicy())
	out.CorsPolicy = EdgeCacheServiceCorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	out.UrlRewrite = EdgeCacheServiceUrlRewrite_FromProto(mapCtx, in.GetUrlRewrite())
	return out
}

func EdgeCacheServiceRouteAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceRouteAction) *pb.EdgeCacheService_Routing_RouteRule_RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction{}
	out.CdnPolicy = EdgeCacheServiceCdnPolicy_ToProto(mapCtx, in.CdnPolicy)
	out.CorsPolicy = EdgeCacheServiceCorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.UrlRewrite = EdgeCacheServiceUrlRewrite_ToProto(mapCtx, in.UrlRewrite)
	return out
}

func EdgeCacheServiceCdnPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy) *krm.EdgeCacheServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceCdnPolicy{}
	out.AddSignatures = EdgeCacheServiceAddSignatures_FromProto(mapCtx, in.GetAddSignatures())
	out.CacheKeyPolicy = EdgeCacheServiceCacheKeyPolicy_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.CacheMode = direct.Enum_FromProto(mapCtx, in.GetCacheMode())
	out.ClientTtl = direct.StringDuration_FromProto(mapCtx, in.GetClientTtl())
	out.DefaultTtl = direct.StringDuration_FromProto(mapCtx, in.GetDefaultTtl())
	out.MaxTtl = direct.StringDuration_FromProto(mapCtx, in.GetMaxTtl())
	out.NegativeCaching = direct.LazyPtr(in.GetNegativeCaching())
	out.NegativeCachingPolicy = in.GetNegativeCachingPolicy()
	if in.GetSignedRequestKeyset() != "" {
		out.SignedRequestKeysetRef = &krm.ResourceRef{External: in.GetSignedRequestKeyset()}
	}
	out.SignedRequestMaximumExpirationTtl = direct.StringDuration_FromProto(mapCtx, in.GetSignedRequestMaximumExpirationTtl())
	out.SignedRequestMode = direct.Enum_FromProto(mapCtx, in.GetSignedRequestMode())
	out.SignedTokenOptions = EdgeCacheServiceSignedTokenOptions_FromProto(mapCtx, in.GetSignedTokenOptions())
	return out
}

func EdgeCacheServiceCdnPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceCdnPolicy) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy{}
	out.AddSignatures = EdgeCacheServiceAddSignatures_ToProto(mapCtx, in.AddSignatures)
	out.CacheKeyPolicy = EdgeCacheServiceCacheKeyPolicy_ToProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = direct.Enum_ToProto[pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_CacheMode](mapCtx, in.CacheMode)
	out.ClientTtl = direct.StringDuration_ToProto(mapCtx, in.ClientTtl)
	out.DefaultTtl = direct.StringDuration_ToProto(mapCtx, in.DefaultTtl)
	out.MaxTtl = direct.StringDuration_ToProto(mapCtx, in.MaxTtl)
	out.NegativeCaching = direct.ValueOf(in.NegativeCaching)
	out.NegativeCachingPolicy = in.NegativeCachingPolicy
	if in.SignedRequestKeysetRef != nil {
		out.SignedRequestKeyset = in.SignedRequestKeysetRef.External
	}
	out.SignedRequestMaximumExpirationTtl = direct.StringDuration_ToProto(mapCtx, in.SignedRequestMaximumExpirationTtl)
	out.SignedRequestMode = direct.Enum_ToProto[pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_SignedRequestMode](mapCtx, in.SignedRequestMode)
	out.SignedTokenOptions = EdgeCacheServiceSignedTokenOptions_ToProto(mapCtx, in.SignedTokenOptions)
	return out
}

func EdgeCacheServiceAddSignatures_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_AddSignatures) *krm.EdgeCacheServiceAddSignatures {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceAddSignatures{}
	out.Actions = direct.EnumSlice_FromProto(mapCtx, in.GetActions())
	out.CopiedParameters = in.GetCopiedParameters()
	if in.GetKeyset() != "" {
		out.KeysetRef = &krm.ResourceRef{External: in.GetKeyset()}
	}
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.TokenTtl = direct.StringDuration_FromProto(mapCtx, in.GetTokenTtl())
	return out
}

func EdgeCacheServiceAddSignatures_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceAddSignatures) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_AddSignatures {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_AddSignatures{}
	out.Actions = direct.EnumSlice_ToProto[pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_AddSignatures_GatewayAction](mapCtx, in.Actions)
	out.CopiedParameters = in.CopiedParameters
	if in.KeysetRef != nil {
		out.Keyset = in.KeysetRef.External
	}
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.TokenTtl = direct.StringDuration_ToProto(mapCtx, in.TokenTtl)
	return out
}

func EdgeCacheServiceCacheKeyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_CacheKeyPolicy) *krm.EdgeCacheServiceCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceCacheKeyPolicy{}
	out.ExcludeHost = direct.LazyPtr(in.GetExcludeHost())
	out.ExcludeQueryString = direct.LazyPtr(in.GetExcludeQueryString())
	out.ExcludedQueryParameters = in.GetExcludedQueryParameters()
	out.IncludeProtocol = direct.LazyPtr(in.GetIncludeProtocol())
	out.IncludedCookieNames = in.GetIncludedCookieNames()
	out.IncludedHeaderNames = in.GetIncludedHeaderNames()
	out.IncludedQueryParameters = in.GetIncludedQueryParameters()
	return out
}

func EdgeCacheServiceCacheKeyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceCacheKeyPolicy) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_CacheKeyPolicy{}
	out.ExcludeHost = direct.ValueOf(in.ExcludeHost)
	out.ExcludeQueryString = direct.ValueOf(in.ExcludeQueryString)
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludeProtocol = direct.ValueOf(in.IncludeProtocol)
	out.IncludedCookieNames = in.IncludedCookieNames
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedQueryParameters = in.IncludedQueryParameters
	return out
}

func EdgeCacheServiceSignedTokenOptions_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_SignedTokenOptions) *krm.EdgeCacheServiceSignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceSignedTokenOptions{}
	out.AllowedSignatureAlgorithms = direct.EnumSlice_FromProto(mapCtx, in.GetAllowedSignatureAlgorithms())
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	return out
}

func EdgeCacheServiceSignedTokenOptions_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceSignedTokenOptions) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_SignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_SignedTokenOptions{}
	out.AllowedSignatureAlgorithms = direct.EnumSlice_ToProto[pb.EdgeCacheService_Routing_RouteRule_RouteAction_CdnPolicy_SignedTokenOptions_SignatureAlgorithm](mapCtx, in.AllowedSignatureAlgorithms)
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	return out
}

func EdgeCacheServiceCorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CorsPolicy) *krm.EdgeCacheServiceCorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceCorsPolicy{}
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.AllowHeaders = in.GetAllowHeaders()
	out.AllowMethods = in.GetAllowMethods()
	out.AllowOrigins = in.GetAllowOrigins()
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.ExposeHeaders = in.GetExposeHeaders()
	out.MaxAge = in.GetMaxAge()
	return out
}

func EdgeCacheServiceCorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceCorsPolicy) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_CorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_CorsPolicy{}
	out.AllowCredentials = direct.ValueOf(in.AllowCredentials)
	out.AllowHeaders = in.AllowHeaders
	out.AllowMethods = in.AllowMethods
	out.AllowOrigins = in.AllowOrigins
	out.Disabled = direct.ValueOf(in.Disabled)
	out.ExposeHeaders = in.ExposeHeaders
	out.MaxAge = in.MaxAge
	return out
}

func EdgeCacheServiceUrlRewrite_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_RouteAction_UrlRewrite) *krm.EdgeCacheServiceUrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceUrlRewrite{}
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.PathTemplateRewrite = direct.LazyPtr(in.GetPathTemplateRewrite())
	return out
}

func EdgeCacheServiceUrlRewrite_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceUrlRewrite) *pb.EdgeCacheService_Routing_RouteRule_RouteAction_UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_RouteAction_UrlRewrite{}
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.PathTemplateRewrite = direct.ValueOf(in.PathTemplateRewrite)
	return out
}

func EdgeCacheServiceUrlRedirect_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService_Routing_RouteRule_UrlRedirect) *krm.EdgeCacheServiceUrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceUrlRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.HttpsRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRedirect = direct.LazyPtr(in.GetPrefixRedirect())
	out.RedirectResponseCode = direct.Enum_FromProto(mapCtx, in.GetRedirectResponseCode())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	return out
}

func EdgeCacheServiceUrlRedirect_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceUrlRedirect) *pb.EdgeCacheService_Routing_RouteRule_UrlRedirect {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService_Routing_RouteRule_UrlRedirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.HttpsRedirect = direct.ValueOf(in.HttpsRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)
	out.RedirectResponseCode = direct.Enum_ToProto[pb.EdgeCacheService_Routing_RouteRule_UrlRedirect_RedirectResponseCode](mapCtx, in.RedirectResponseCode)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}
