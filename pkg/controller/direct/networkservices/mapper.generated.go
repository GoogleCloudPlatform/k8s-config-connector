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
// krm.group: networkservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: mockgcp.cloud.edgecacheservice.v1

package networkservices

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecacheservice/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CDNPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CdnPolicy) *krm.CDNPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CDNPolicy{}
	out.CacheMode = direct.LazyPtr(in.GetCacheMode())
	out.ClientTTL = in.ClientTtl
	out.DefaultTTL = direct.StringDuration_FromProto(mapCtx, in.GetDefaultTtl())
	out.MaxTTL = direct.StringDuration_FromProto(mapCtx, in.GetMaxTtl())
	out.CacheNegativeCallbacks = direct.LazyPtr(in.GetCacheNegativeCallbacks())
	out.SignedRequestKeyset = in.SignedRequestKeyset
	out.SignedRequestMode = direct.LazyPtr(in.GetSignedRequestMode())
	return out
}
func CDNPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CDNPolicy) *pb.CdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CdnPolicy{}
	out.CacheMode = direct.ValueOf(in.CacheMode)
	out.ClientTtl = in.ClientTTL
	out.DefaultTtl = direct.StringDuration_ToProto(mapCtx, in.DefaultTTL)
	out.MaxTtl = direct.StringDuration_ToProto(mapCtx, in.MaxTTL)
	out.CacheNegativeCallbacks = direct.ValueOf(in.CacheNegativeCallbacks)
	out.SignedRequestKeyset = in.SignedRequestKeyset
	out.SignedRequestMode = direct.ValueOf(in.SignedRequestMode)
	return out
}
func CorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.CorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CorsPolicy{}
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.AllowOrigins = in.AllowOrigins
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func CorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CorsPolicy) *pb.CorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CorsPolicy{}
	out.MaxAge = direct.StringDuration_ToProto(mapCtx, in.MaxAge)
	out.AllowOrigins = in.AllowOrigins
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.AllowCredentials = direct.ValueOf(in.AllowCredentials)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
func EdgeCacheService_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Routing = Routing_FromProto(mapCtx, in.GetRouting())
	out.RequireTLS = direct.LazyPtr(in.GetRequireTls())
	out.EdgeSSLCertificates = in.EdgeSslCertificates
	out.EdgeSecurityPolicy = direct.LazyPtr(in.GetEdgeSecurityPolicy())
	out.LogConfig = LogConfig_FromProto(mapCtx, in.GetLogConfig())
	out.DisableQuic = direct.LazyPtr(in.GetDisableQuic())
	out.DisableHttp2 = direct.LazyPtr(in.GetDisableHttp2())
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	return out
}
func EdgeCacheService_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Routing = Routing_ToProto(mapCtx, in.Routing)
	out.RequireTls = direct.ValueOf(in.RequireTLS)
	out.EdgeSslCertificates = in.EdgeSSLCertificates
	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)
	out.LogConfig = LogConfig_ToProto(mapCtx, in.LogConfig)
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	return out
}
func EdgeCacheServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.EdgeCacheServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheServiceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Routing
	// MISSING: RequireTLS
	// MISSING: EdgeSSLCertificates
	// MISSING: EdgeSecurityPolicy
	// MISSING: LogConfig
	// MISSING: DisableQuic
	// MISSING: DisableHttp2
	out.IPV4Addresses = in.Ipv4Addresses
	out.IPV6Addresses = in.Ipv6Addresses
	return out
}
func EdgeCacheServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheServiceObservedState) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Routing
	// MISSING: RequireTLS
	// MISSING: EdgeSSLCertificates
	// MISSING: EdgeSecurityPolicy
	// MISSING: LogConfig
	// MISSING: DisableQuic
	// MISSING: DisableHttp2
	out.Ipv4Addresses = in.IPV4Addresses
	out.Ipv6Addresses = in.IPV6Addresses
	return out
}
func HeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.HeaderAction) *krm.HeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.HeaderAction{}
	out.HeadersToRemove = direct.Slice_FromProto(mapCtx, in.HeadersToRemove, HeaderToRemove_FromProto)
	out.HeadersToAdd = direct.Slice_FromProto(mapCtx, in.HeadersToAdd, HeaderToAdd_FromProto)
	return out
}
func HeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.HeaderAction) *pb.HeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.HeaderAction{}
	out.HeadersToRemove = direct.Slice_ToProto(mapCtx, in.HeadersToRemove, HeaderToRemove_ToProto)
	out.HeadersToAdd = direct.Slice_ToProto(mapCtx, in.HeadersToAdd, HeaderToAdd_ToProto)
	return out
}
func HeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.HeaderMatch) *krm.HeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.HeaderMatch{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	return out
}
func HeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.HeaderMatch) *pb.HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HeaderMatch{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.SuffixMatch = direct.ValueOf(in.SuffixMatch)
	out.RegexMatch = direct.ValueOf(in.RegexMatch)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	return out
}
func HeaderToAdd_FromProto(mapCtx *direct.MapContext, in *pb.HeaderToAdd) *krm.HeaderToAdd {
	if in == nil {
		return nil
	}
	out := &krm.HeaderToAdd{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.HeaderValue = direct.LazyPtr(in.GetHeaderValue())
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}
func HeaderToAdd_ToProto(mapCtx *direct.MapContext, in *krm.HeaderToAdd) *pb.HeaderToAdd {
	if in == nil {
		return nil
	}
	out := &pb.HeaderToAdd{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}
func HeaderToRemove_FromProto(mapCtx *direct.MapContext, in *pb.HeaderToRemove) *krm.HeaderToRemove {
	if in == nil {
		return nil
	}
	out := &krm.HeaderToRemove{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	return out
}
func HeaderToRemove_ToProto(mapCtx *direct.MapContext, in *krm.HeaderToRemove) *pb.HeaderToRemove {
	if in == nil {
		return nil
	}
	out := &pb.HeaderToRemove{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	return out
}
func HostRule_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.HostRule {
	if in == nil {
		return nil
	}
	out := &krm.HostRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hosts = in.Hosts
	out.PathMatcher = direct.LazyPtr(in.GetPathMatcher())
	return out
}
func HostRule_ToProto(mapCtx *direct.MapContext, in *krm.HostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Hosts = in.Hosts
	out.PathMatcher = direct.ValueOf(in.PathMatcher)
	return out
}
func LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.LogConfig) *krm.LogConfig {
	if in == nil {
		return nil
	}
	out := &krm.LogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	out.SampleRate = direct.LazyPtr(in.GetSampleRate())
	return out
}
func LogConfig_ToProto(mapCtx *direct.MapContext, in *krm.LogConfig) *pb.LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.LogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	out.SampleRate = direct.ValueOf(in.SampleRate)
	return out
}
func MatchRule_FromProto(mapCtx *direct.MapContext, in *pb.MatchRule) *krm.MatchRule {
	if in == nil {
		return nil
	}
	out := &krm.MatchRule{}
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.HeaderMatches = direct.Slice_FromProto(mapCtx, in.HeaderMatches, HeaderMatch_FromProto)
	out.QueryParamMatches = direct.Slice_FromProto(mapCtx, in.QueryParamMatches, QueryParamMatch_FromProto)
	return out
}
func MatchRule_ToProto(mapCtx *direct.MapContext, in *krm.MatchRule) *pb.MatchRule {
	if in == nil {
		return nil
	}
	out := &pb.MatchRule{}
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)
	out.RegexMatch = direct.ValueOf(in.RegexMatch)
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.HeaderMatches = direct.Slice_ToProto(mapCtx, in.HeaderMatches, HeaderMatch_ToProto)
	out.QueryParamMatches = direct.Slice_ToProto(mapCtx, in.QueryParamMatches, QueryParamMatch_ToProto)
	return out
}
func NetworkServicesEdgeCacheServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Routing
	// MISSING: RequireTLS
	// MISSING: EdgeSSLCertificates
	// MISSING: EdgeSecurityPolicy
	// MISSING: LogConfig
	// MISSING: DisableQuic
	// MISSING: DisableHttp2
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	return out
}
func NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Routing
	// MISSING: RequireTLS
	// MISSING: EdgeSSLCertificates
	// MISSING: EdgeSecurityPolicy
	// MISSING: LogConfig
	// MISSING: DisableQuic
	// MISSING: DisableHttp2
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	return out
}
func PathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.PathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.PathMatcher{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, RouteRule_FromProto)
	return out
}
func PathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.PathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, RouteRule_ToProto)
	return out
}
func QueryParamMatch_FromProto(mapCtx *direct.MapContext, in *pb.QueryParamMatch) *krm.QueryParamMatch {
	if in == nil {
		return nil
	}
	out := &krm.QueryParamMatch{}
	out.QueryParam = direct.LazyPtr(in.GetQueryParam())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	return out
}
func QueryParamMatch_ToProto(mapCtx *direct.MapContext, in *krm.QueryParamMatch) *pb.QueryParamMatch {
	if in == nil {
		return nil
	}
	out := &pb.QueryParamMatch{}
	out.QueryParam = direct.ValueOf(in.QueryParam)
	out.PresentMatch = direct.ValueOf(in.PresentMatch)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	out.RegexMatch = direct.ValueOf(in.RegexMatch)
	return out
}
func RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.RouteAction) *krm.RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.RouteAction{}
	out.CDNPolicy = CDNPolicy_FromProto(mapCtx, in.GetCdnPolicy())
	out.URLRewrite = URLRewrite_FromProto(mapCtx, in.GetUrlRewrite())
	out.CorsPolicy = CorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	out.CompressionMode = direct.LazyPtr(in.GetCompressionMode())
	return out
}
func RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.RouteAction) *pb.RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.RouteAction{}
	out.CdnPolicy = CDNPolicy_ToProto(mapCtx, in.CDNPolicy)
	out.UrlRewrite = URLRewrite_ToProto(mapCtx, in.URLRewrite)
	out.CorsPolicy = CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.CompressionMode = direct.ValueOf(in.CompressionMode)
	return out
}
func RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.RouteRule) *krm.RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.RouteRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, MatchRule_FromProto)
	out.HeaderAction = HeaderAction_FromProto(mapCtx, in.GetHeaderAction())
	out.RouteAction = RouteAction_FromProto(mapCtx, in.GetRouteAction())
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.URLRedirect = URLRedirect_FromProto(mapCtx, in.GetUrlRedirect())
	out.RouteMethods = in.RouteMethods
	return out
}
func RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.RouteRule) *pb.RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.RouteRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Priority = direct.ValueOf(in.Priority)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, MatchRule_ToProto)
	out.HeaderAction = HeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.RouteAction = RouteAction_ToProto(mapCtx, in.RouteAction)
	out.Origin = direct.ValueOf(in.Origin)
	out.UrlRedirect = URLRedirect_ToProto(mapCtx, in.URLRedirect)
	out.RouteMethods = in.RouteMethods
	return out
}
func Routing_FromProto(mapCtx *direct.MapContext, in *pb.Routing) *krm.Routing {
	if in == nil {
		return nil
	}
	out := &krm.Routing{}
	out.HostRules = direct.Slice_FromProto(mapCtx, in.HostRules, HostRule_FromProto)
	out.PathMatchers = direct.Slice_FromProto(mapCtx, in.PathMatchers, PathMatcher_FromProto)
	return out
}
func Routing_ToProto(mapCtx *direct.MapContext, in *krm.Routing) *pb.Routing {
	if in == nil {
		return nil
	}
	out := &pb.Routing{}
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRules, HostRule_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatchers, PathMatcher_ToProto)
	return out
}
func URLRedirect_FromProto(mapCtx *direct.MapContext, in *pb.UrlRedirect) *krm.URLRedirect {
	if in == nil {
		return nil
	}
	out := &krm.URLRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRedirect = direct.LazyPtr(in.GetPrefixRedirect())
	out.RedirectResponseCode = direct.LazyPtr(in.GetRedirectResponseCode())
	out.HTTPSRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	return out
}
func URLRedirect_ToProto(mapCtx *direct.MapContext, in *krm.URLRedirect) *pb.UrlRedirect {
	if in == nil {
		return nil
	}
	out := &pb.UrlRedirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)
	out.RedirectResponseCode = direct.ValueOf(in.RedirectResponseCode)
	out.HttpsRedirect = direct.ValueOf(in.HTTPSRedirect)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}
func URLRewrite_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.URLRewrite {
	if in == nil {
		return nil
	}
	out := &krm.URLRewrite{}
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	return out
}
func URLRewrite_ToProto(mapCtx *direct.MapContext, in *krm.URLRewrite) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	return out
}
