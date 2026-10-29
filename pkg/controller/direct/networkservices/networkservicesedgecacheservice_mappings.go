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
	"maps"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Description = direct.ValueOf(in.Description)
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	out.RequireTls = direct.ValueOf(in.RequireTls)
	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)
	out.SslPolicy = direct.ValueOf(in.SslPolicy)
	out.EdgeSslCertificates = in.EdgeSslCertificates
	out.LogConfig = EdgeCacheServiceLogConfig_ToProto(mapCtx, in.LogConfig)
	out.Routing = EdgeCacheServiceRouting_ToProto(mapCtx, &in.Routing)
	return out
}

func EdgeCacheServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisableQuic = direct.LazyPtr(in.GetDisableQuic())
	out.DisableHttp2 = direct.LazyPtr(in.GetDisableHttp2())
	out.RequireTls = direct.LazyPtr(in.GetRequireTls())
	out.EdgeSecurityPolicy = direct.LazyPtr(in.GetEdgeSecurityPolicy())
	out.SslPolicy = direct.LazyPtr(in.GetSslPolicy())
	out.EdgeSslCertificates = in.GetEdgeSslCertificates()
	out.LogConfig = EdgeCacheServiceLogConfig_FromProto(mapCtx, in.GetLogConfig())
	if in.GetRouting() != nil {
		routing := EdgeCacheServiceRouting_FromProto(mapCtx, in.GetRouting())
		if routing != nil {
			out.Routing = *routing
		}
	}
	return out
}

func EdgeCacheServiceStatus_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceStatus) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	return out
}

func EdgeCacheServiceStatus_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceStatus {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceStatus{}
	out.Ipv4Addresses = in.GetIpv4Addresses()
	out.Ipv6Addresses = in.GetIpv6Addresses()
	return out
}

func EdgeCacheServiceLogConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceLogConfig) *pb.EdgeCacheServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceLogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	out.SampleRate = direct.ValueOf(in.SampleRate)
	return out
}

func EdgeCacheServiceLogConfig_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceLogConfig) *krm.EdgecacheserviceLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceLogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	out.SampleRate = direct.LazyPtr(in.GetSampleRate())
	return out
}

func EdgeCacheServiceRouting_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouting) *pb.EdgeCacheServiceRouting {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceRouting{}
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRule, EdgeCacheServiceHostRule_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatcher, EdgeCacheServicePathMatcher_ToProto)
	return out
}

func EdgeCacheServiceRouting_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceRouting) *krm.EdgecacheserviceRouting {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRouting{}
	out.HostRule = direct.Slice_FromProto(mapCtx, in.GetHostRules(), EdgeCacheServiceHostRule_FromProto)
	out.PathMatcher = direct.Slice_FromProto(mapCtx, in.GetPathMatchers(), EdgeCacheServicePathMatcher_FromProto)
	return out
}

func EdgeCacheServiceHostRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceHostRule) *pb.EdgeCacheServiceHostRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHostRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Hosts = in.Hosts
	out.PathMatcher = in.PathMatcher
	return out
}

func EdgeCacheServiceHostRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHostRule) *krm.EdgecacheserviceHostRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceHostRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hosts = in.GetHosts()
	out.PathMatcher = in.GetPathMatcher()
	return out
}

func EdgeCacheServicePathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheservicePathMatcher) *pb.EdgeCacheServicePathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServicePathMatcher{}
	out.Description = direct.ValueOf(in.Description)
	out.Name = in.Name
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRule, EdgeCacheServiceRouteRule_ToProto)
	return out
}

func EdgeCacheServicePathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServicePathMatcher) *krm.EdgecacheservicePathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheservicePathMatcher{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Name = in.GetName()
	out.RouteRule = direct.Slice_FromProto(mapCtx, in.GetRouteRules(), EdgeCacheServiceRouteRule_FromProto)
	return out
}

func EdgeCacheServiceRouteRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouteRule) *pb.EdgeCacheServiceRouteRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceRouteRule{}
	out.Priority = in.Priority
	out.Description = direct.ValueOf(in.Description)
	out.Origin = direct.ValueOf(in.Origin)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRule, EdgeCacheServiceMatchRule_ToProto)
	out.HeaderAction = EdgeCacheServiceHeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.RouteAction = EdgeCacheServiceRouteAction_ToProto(mapCtx, in.RouteAction)
	out.RouteMethods = EdgeCacheServiceRouteMethods_ToProto(mapCtx, in.RouteMethods)
	out.UrlRedirect = EdgeCacheServiceUrlRedirect_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func EdgeCacheServiceRouteRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceRouteRule) *krm.EdgecacheserviceRouteRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRouteRule{}
	out.Priority = in.GetPriority()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.MatchRule = direct.Slice_FromProto(mapCtx, in.GetMatchRules(), EdgeCacheServiceMatchRule_FromProto)
	out.HeaderAction = EdgeCacheServiceHeaderAction_FromProto(mapCtx, in.GetHeaderAction())
	out.RouteAction = EdgeCacheServiceRouteAction_FromProto(mapCtx, in.GetRouteAction())
	out.RouteMethods = EdgeCacheServiceRouteMethods_FromProto(mapCtx, in.GetRouteMethods())
	out.UrlRedirect = EdgeCacheServiceUrlRedirect_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func EdgeCacheServiceMatchRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceMatchRule) *pb.EdgeCacheServiceMatchRule {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceMatchRule{}
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)
	out.PathTemplateMatch = direct.ValueOf(in.PathTemplateMatch)
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.HeaderMatches = direct.Slice_ToProto(mapCtx, in.HeaderMatch, EdgeCacheServiceHeaderMatch_ToProto)
	out.QueryParameterMatches = direct.Slice_ToProto(mapCtx, in.QueryParameterMatch, EdgeCacheServiceQueryParameterMatch_ToProto)
	return out
}

func EdgeCacheServiceMatchRule_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceMatchRule) *krm.EdgecacheserviceMatchRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceMatchRule{}
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	out.PathTemplateMatch = direct.LazyPtr(in.GetPathTemplateMatch())
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.HeaderMatch = direct.Slice_FromProto(mapCtx, in.GetHeaderMatches(), EdgeCacheServiceHeaderMatch_FromProto)
	out.QueryParameterMatch = direct.Slice_FromProto(mapCtx, in.GetQueryParameterMatches(), EdgeCacheServiceQueryParameterMatch_FromProto)
	return out
}

func EdgeCacheServiceHeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceHeaderMatch) *pb.EdgeCacheServiceHeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeaderMatch{}
	out.HeaderName = in.HeaderName
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.SuffixMatch = direct.ValueOf(in.SuffixMatch)
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	return out
}

func EdgeCacheServiceHeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeaderMatch) *krm.EdgecacheserviceHeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceHeaderMatch{}
	out.HeaderName = in.GetHeaderName()
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	return out
}

func EdgeCacheServiceQueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceQueryParameterMatch) *pb.EdgeCacheServiceQueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceQueryParameterMatch{}
	out.Name = in.Name
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	return out
}

func EdgeCacheServiceQueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceQueryParameterMatch) *krm.EdgecacheserviceQueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceQueryParameterMatch{}
	out.Name = in.GetName()
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	return out
}

func EdgeCacheServiceHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceHeaderAction) *pb.EdgeCacheServiceHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_ToProto(mapCtx, in.RequestHeaderToAdd, EdgeCacheServiceRequestHeaderToAdd_ToProto)
	out.RequestHeadersToRemove = direct.Slice_ToProto(mapCtx, in.RequestHeaderToRemove, EdgeCacheServiceRequestHeaderToRemove_ToProto)
	out.ResponseHeadersToAdd = direct.Slice_ToProto(mapCtx, in.ResponseHeaderToAdd, EdgeCacheServiceResponseHeaderToAdd_ToProto)
	out.ResponseHeadersToRemove = direct.Slice_ToProto(mapCtx, in.ResponseHeaderToRemove, EdgeCacheServiceResponseHeaderToRemove_ToProto)
	return out
}

func EdgeCacheServiceHeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeaderAction) *krm.EdgecacheserviceHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceHeaderAction{}
	out.RequestHeaderToAdd = direct.Slice_FromProto(mapCtx, in.GetRequestHeadersToAdd(), EdgeCacheServiceRequestHeaderToAdd_FromProto)
	out.RequestHeaderToRemove = direct.Slice_FromProto(mapCtx, in.GetRequestHeadersToRemove(), EdgeCacheServiceRequestHeaderToRemove_FromProto)
	out.ResponseHeaderToAdd = direct.Slice_FromProto(mapCtx, in.GetResponseHeadersToAdd(), EdgeCacheServiceResponseHeaderToAdd_FromProto)
	out.ResponseHeaderToRemove = direct.Slice_FromProto(mapCtx, in.GetResponseHeadersToRemove(), EdgeCacheServiceResponseHeaderToRemove_FromProto)
	return out
}

func EdgeCacheServiceRequestHeaderToAdd_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRequestHeaderToAdd) *pb.EdgeCacheServiceHeader {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeader{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func EdgeCacheServiceRequestHeaderToAdd_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeader) *krm.EdgecacheserviceRequestHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRequestHeaderToAdd{}
	out.HeaderName = in.GetHeaderName()
	out.HeaderValue = in.GetHeaderValue()
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}

func EdgeCacheServiceRequestHeaderToRemove_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRequestHeaderToRemove) *pb.EdgeCacheServiceHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeaderToRemove{}
	out.HeaderName = in.HeaderName
	return out
}

func EdgeCacheServiceRequestHeaderToRemove_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeaderToRemove) *krm.EdgecacheserviceRequestHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRequestHeaderToRemove{}
	out.HeaderName = in.GetHeaderName()
	return out
}

func EdgeCacheServiceResponseHeaderToAdd_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceResponseHeaderToAdd) *pb.EdgeCacheServiceHeader {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeader{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func EdgeCacheServiceResponseHeaderToAdd_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeader) *krm.EdgecacheserviceResponseHeaderToAdd {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceResponseHeaderToAdd{}
	out.HeaderName = in.GetHeaderName()
	out.HeaderValue = in.GetHeaderValue()
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}

func EdgeCacheServiceResponseHeaderToRemove_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceResponseHeaderToRemove) *pb.EdgeCacheServiceHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceHeaderToRemove{}
	out.HeaderName = in.HeaderName
	return out
}

func EdgeCacheServiceResponseHeaderToRemove_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceHeaderToRemove) *krm.EdgecacheserviceResponseHeaderToRemove {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceResponseHeaderToRemove{}
	out.HeaderName = in.GetHeaderName()
	return out
}

func EdgeCacheServiceRouteAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouteAction) *pb.EdgeCacheServiceRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceRouteAction{}
	out.CdnPolicy = EdgeCacheServiceCdnPolicy_ToProto(mapCtx, in.CdnPolicy)
	out.CompressionMode = direct.ValueOf(in.CompressionMode)
	out.CorsPolicy = EdgeCacheServiceCorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.UrlRewrite = EdgeCacheServiceUrlRewrite_ToProto(mapCtx, in.UrlRewrite)
	return out
}

func EdgeCacheServiceRouteAction_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceRouteAction) *krm.EdgecacheserviceRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRouteAction{}
	out.CdnPolicy = EdgeCacheServiceCdnPolicy_FromProto(mapCtx, in.GetCdnPolicy())
	out.CompressionMode = direct.LazyPtr(in.GetCompressionMode())
	out.CorsPolicy = EdgeCacheServiceCorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	out.UrlRewrite = EdgeCacheServiceUrlRewrite_FromProto(mapCtx, in.GetUrlRewrite())
	return out
}

func EdgeCacheServiceCdnPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceCdnPolicy) *pb.EdgeCacheServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceCdnPolicy{}
	out.CacheMode = direct.ValueOf(in.CacheMode)
	out.ClientTtl = direct.ValueOf(in.ClientTtl)
	out.DefaultTtl = direct.ValueOf(in.DefaultTtl)
	out.MaxTtl = direct.ValueOf(in.MaxTtl)
	out.NegativeCaching = direct.ValueOf(in.NegativeCaching)
	out.NegativeCachingPolicy = maps.Clone(in.NegativeCachingPolicy)
	out.SignedRequestMode = direct.ValueOf(in.SignedRequestMode)
	out.SignedRequestKeyset = direct.ValueOf(in.SignedRequestKeyset)
	out.SignedRequestMaximumExpirationTtl = direct.ValueOf(in.SignedRequestMaximumExpirationTtl)
	out.AddSignatures = EdgeCacheServiceAddSignatures_ToProto(mapCtx, in.AddSignatures)
	out.CacheKeyPolicy = EdgeCacheServiceCacheKeyPolicy_ToProto(mapCtx, in.CacheKeyPolicy)
	out.SignedTokenOptions = EdgeCacheServiceSignedTokenOptions_ToProto(mapCtx, in.SignedTokenOptions)
	return out
}

func EdgeCacheServiceCdnPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceCdnPolicy) *krm.EdgecacheserviceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceCdnPolicy{}
	out.CacheMode = direct.LazyPtr(in.GetCacheMode())
	out.ClientTtl = direct.LazyPtr(in.GetClientTtl())
	out.DefaultTtl = direct.LazyPtr(in.GetDefaultTtl())
	out.MaxTtl = direct.LazyPtr(in.GetMaxTtl())
	out.NegativeCaching = direct.LazyPtr(in.GetNegativeCaching())
	out.NegativeCachingPolicy = maps.Clone(in.GetNegativeCachingPolicy())
	out.SignedRequestMode = direct.LazyPtr(in.GetSignedRequestMode())
	out.SignedRequestKeyset = direct.LazyPtr(in.GetSignedRequestKeyset())
	out.SignedRequestMaximumExpirationTtl = direct.LazyPtr(in.GetSignedRequestMaximumExpirationTtl())
	out.AddSignatures = EdgeCacheServiceAddSignatures_FromProto(mapCtx, in.GetAddSignatures())
	out.CacheKeyPolicy = EdgeCacheServiceCacheKeyPolicy_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.SignedTokenOptions = EdgeCacheServiceSignedTokenOptions_FromProto(mapCtx, in.GetSignedTokenOptions())
	return out
}

func EdgeCacheServiceAddSignatures_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceAddSignatures) *pb.EdgeCacheServiceAddSignatures {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceAddSignatures{}
	out.Actions = in.Actions
	out.CopiedParameters = in.CopiedParameters
	out.Keyset = direct.ValueOf(in.Keyset)
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.TokenTtl = direct.ValueOf(in.TokenTtl)
	return out
}

func EdgeCacheServiceAddSignatures_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceAddSignatures) *krm.EdgecacheserviceAddSignatures {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceAddSignatures{}
	out.Actions = in.GetActions()
	out.CopiedParameters = in.GetCopiedParameters()
	out.Keyset = direct.LazyPtr(in.GetKeyset())
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.TokenTtl = direct.LazyPtr(in.GetTokenTtl())
	return out
}

func EdgeCacheServiceCacheKeyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceCacheKeyPolicy) *pb.EdgeCacheServiceCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceCacheKeyPolicy{}
	out.IncludeProtocol = direct.ValueOf(in.IncludeProtocol)
	out.ExcludeHost = direct.ValueOf(in.ExcludeHost)
	out.ExcludeQueryString = direct.ValueOf(in.ExcludeQueryString)
	out.IncludedQueryParameters = in.IncludedQueryParameters
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedCookieNames = in.IncludedCookieNames
	return out
}

func EdgeCacheServiceCacheKeyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceCacheKeyPolicy) *krm.EdgecacheserviceCacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceCacheKeyPolicy{}
	out.IncludeProtocol = direct.LazyPtr(in.GetIncludeProtocol())
	out.ExcludeHost = direct.LazyPtr(in.GetExcludeHost())
	out.ExcludeQueryString = direct.LazyPtr(in.GetExcludeQueryString())
	out.IncludedQueryParameters = in.GetIncludedQueryParameters()
	out.ExcludedQueryParameters = in.GetExcludedQueryParameters()
	out.IncludedHeaderNames = in.GetIncludedHeaderNames()
	out.IncludedCookieNames = in.GetIncludedCookieNames()
	return out
}

func EdgeCacheServiceSignedTokenOptions_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceSignedTokenOptions) *pb.EdgeCacheServiceSignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceSignedTokenOptions{}
	out.AllowedSignatureAlgorithms = in.AllowedSignatureAlgorithms
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	return out
}

func EdgeCacheServiceSignedTokenOptions_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceSignedTokenOptions) *krm.EdgecacheserviceSignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceSignedTokenOptions{}
	out.AllowedSignatureAlgorithms = in.GetAllowedSignatureAlgorithms()
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	return out
}

func EdgeCacheServiceCorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceCorsPolicy) *pb.EdgeCacheServiceCorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceCorsPolicy{}
	out.MaxAge = in.MaxAge
	out.AllowOrigins = in.AllowOrigins
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.AllowCredentials = direct.ValueOf(in.AllowCredentials)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}

func EdgeCacheServiceCorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceCorsPolicy) *krm.EdgecacheserviceCorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceCorsPolicy{}
	out.MaxAge = in.GetMaxAge()
	out.AllowOrigins = in.GetAllowOrigins()
	out.AllowMethods = in.GetAllowMethods()
	out.AllowHeaders = in.GetAllowHeaders()
	out.ExposeHeaders = in.GetExposeHeaders()
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}

func EdgeCacheServiceUrlRewrite_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceUrlRewrite) *pb.EdgeCacheServiceUrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceUrlRewrite{}
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.PathTemplateRewrite = direct.ValueOf(in.PathTemplateRewrite)
	return out
}

func EdgeCacheServiceUrlRewrite_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceUrlRewrite) *krm.EdgecacheserviceUrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceUrlRewrite{}
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.PathTemplateRewrite = direct.LazyPtr(in.GetPathTemplateRewrite())
	return out
}

func EdgeCacheServiceRouteMethods_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouteMethods) *pb.EdgeCacheServiceRouteMethods {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceRouteMethods{}
	out.AllowedMethods = in.AllowedMethods
	return out
}

func EdgeCacheServiceRouteMethods_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceRouteMethods) *krm.EdgecacheserviceRouteMethods {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceRouteMethods{}
	out.AllowedMethods = in.GetAllowedMethods()
	return out
}

func EdgeCacheServiceUrlRedirect_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceUrlRedirect) *pb.EdgeCacheServiceUrlRedirect {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheServiceUrlRedirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)
	out.RedirectResponseCode = direct.ValueOf(in.RedirectResponseCode)
	out.HttpsRedirect = direct.ValueOf(in.HttpsRedirect)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}

func EdgeCacheServiceUrlRedirect_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheServiceUrlRedirect) *krm.EdgecacheserviceUrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.EdgecacheserviceUrlRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRedirect = direct.LazyPtr(in.GetPrefixRedirect())
	out.RedirectResponseCode = direct.LazyPtr(in.GetRedirectResponseCode())
	out.HttpsRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	return out
}
