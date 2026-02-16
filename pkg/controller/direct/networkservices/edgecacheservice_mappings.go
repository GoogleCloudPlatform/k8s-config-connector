// Copyright 2024 Google LLC
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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *EdgeCacheService {

	if in == nil {

		return nil

	}

	out := &EdgeCacheService{}

	out.Description = direct.ValueOf(in.Description)

	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)

	out.DisableQuic = direct.ValueOf(in.DisableQuic)

	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)

	out.EdgeSslCertificates = in.EdgeSslCertificates

	out.LogConfig = EdgeCacheServiceLogConfig_ToProto(mapCtx, in.LogConfig)

	out.RequireTls = direct.ValueOf(in.RequireTls)

	out.Routing = EdgeCacheServiceRouting_ToProto(mapCtx, &in.Routing)

	out.SslPolicy = direct.ValueOf(in.SslPolicy)

	return out

}

func EdgeCacheServiceLogConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceLogConfig) *EdgeCacheLogConfig {

	if in == nil {

		return nil

	}

	out := &EdgeCacheLogConfig{}

	out.Enable = direct.ValueOf(in.Enable)

	out.SampleRate = direct.ValueOf(in.SampleRate)

	return out

}

func EdgeCacheServiceRouting_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouting) *EdgeCacheRouting {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRouting{}

	out.HostRules = make([]*EdgeCacheRoutingHostRule, len(in.HostRule))

	for i, v := range in.HostRule {

		out.HostRules[i] = EdgeCacheRoutingHostRule_ToProto(mapCtx, &v)

	}

	out.PathMatchers = make([]*EdgeCacheRoutingPathMatcher, len(in.PathMatcher))

	for i, v := range in.PathMatcher {

		out.PathMatchers[i] = EdgeCacheRoutingPathMatcher_ToProto(mapCtx, &v)

	}

	return out

}

func EdgeCacheRoutingHostRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceHostRule) *EdgeCacheRoutingHostRule {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingHostRule{}

	out.Description = direct.ValueOf(in.Description)

	out.Hosts = in.Hosts

	out.PathMatcher = in.PathMatcher

	return out

}

func EdgeCacheRoutingPathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheservicePathMatcher) *EdgeCacheRoutingPathMatcher {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingPathMatcher{}

	out.Description = direct.ValueOf(in.Description)

	out.Name = in.Name

	out.RouteRules = make([]*EdgeCacheRoutingRouteRule, len(in.RouteRule))

	for i, v := range in.RouteRule {

		out.RouteRules[i] = EdgeCacheRoutingRouteRule_ToProto(mapCtx, &v)

	}

	return out

}

func EdgeCacheRoutingRouteRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouteRule) *EdgeCacheRoutingRouteRule {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingRouteRule{}

	out.Description = direct.ValueOf(in.Description)

	out.HeaderAction = EdgeCacheRoutingHeaderAction_ToProto(mapCtx, in.HeaderAction)

	out.MatchRules = make([]*EdgeCacheRoutingRouteMatch, len(in.MatchRule))

	for i, v := range in.MatchRule {

		out.MatchRules[i] = EdgeCacheRoutingRouteMatch_ToProto(mapCtx, &v)

	}

	out.Origin = direct.ValueOf(in.Origin)

	out.Priority = in.Priority

	out.RouteAction = EdgeCacheRoutingRouteAction_ToProto(mapCtx, in.RouteAction)

	out.UrlRedirect = EdgeCacheRoutingUrlRedirect_ToProto(mapCtx, in.UrlRedirect)

	return out

}

func EdgeCacheRoutingHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceHeaderAction) *EdgeCacheRoutingHeaderAction {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingHeaderAction{}

	out.RequestHeaderToAdd = make([]*EdgeCacheRoutingHeaderAdd, len(in.RequestHeaderToAdd))

	for i, v := range in.RequestHeaderToAdd {

		out.RequestHeaderToAdd[i] = &EdgeCacheRoutingHeaderAdd{

			HeaderName: v.HeaderName,

			HeaderValue: v.HeaderValue,

			Replace: direct.ValueOf(v.Replace),
		}

	}

	out.RequestHeaderToRemove = make([]*EdgeCacheRoutingHeaderRemove, len(in.RequestHeaderToRemove))

	for i, v := range in.RequestHeaderToRemove {

		out.RequestHeaderToRemove[i] = &EdgeCacheRoutingHeaderRemove{

			HeaderName: v.HeaderName,
		}

	}

	out.ResponseHeaderToAdd = make([]*EdgeCacheRoutingHeaderAdd, len(in.ResponseHeaderToAdd))

	for i, v := range in.ResponseHeaderToAdd {

		out.ResponseHeaderToAdd[i] = &EdgeCacheRoutingHeaderAdd{

			HeaderName: v.HeaderName,

			HeaderValue: v.HeaderValue,

			Replace: direct.ValueOf(v.Replace),
		}

	}

	out.ResponseHeaderToRemove = make([]*EdgeCacheRoutingHeaderRemove, len(in.ResponseHeaderToRemove))

	for i, v := range in.ResponseHeaderToRemove {

		out.ResponseHeaderToRemove[i] = &EdgeCacheRoutingHeaderRemove{

			HeaderName: v.HeaderName,
		}

	}

	return out

}

func EdgeCacheRoutingRouteMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceMatchRule) *EdgeCacheRoutingRouteMatch {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingRouteMatch{}

	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)

	out.HeaderMatches = make([]*EdgeCacheRoutingHeaderMatch, len(in.HeaderMatch))

	for i, v := range in.HeaderMatch {

		out.HeaderMatches[i] = &EdgeCacheRoutingHeaderMatch{

			ExactMatch: direct.ValueOf(v.ExactMatch),

			HeaderName: v.HeaderName,

			InvertMatch: direct.ValueOf(v.InvertMatch),

			PrefixMatch: direct.ValueOf(v.PrefixMatch),

			PresentMatch: direct.ValueOf(v.PresentMatch),

			SuffixMatch: direct.ValueOf(v.SuffixMatch),
		}

	}

	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)

	out.PathTemplateMatch = direct.ValueOf(in.PathTemplateMatch)

	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)

	out.QueryParameterMatches = make([]*EdgeCacheRoutingQueryParameterMatch, len(in.QueryParameterMatch))

	for i, v := range in.QueryParameterMatch {

		out.QueryParameterMatches[i] = &EdgeCacheRoutingQueryParameterMatch{

			ExactMatch: direct.ValueOf(v.ExactMatch),

			Name: v.Name,

			PresentMatch: direct.ValueOf(v.PresentMatch),
		}

	}

	return out

}

func EdgeCacheRoutingRouteAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceRouteAction) *EdgeCacheRoutingRouteAction {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingRouteAction{}

	out.CdnPolicy = EdgeCacheRoutingCdnPolicy_ToProto(mapCtx, in.CdnPolicy)

	out.CorsPolicy = EdgeCacheRoutingCorsPolicy_ToProto(mapCtx, in.CorsPolicy)

	out.UrlRewrite = EdgeCacheRoutingUrlRewrite_ToProto(mapCtx, in.UrlRewrite)

	return out

}

func EdgeCacheRoutingCdnPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceCdnPolicy) *EdgeCacheRoutingCdnPolicy {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingCdnPolicy{}

	if in.AddSignatures != nil {

		out.AddSignatures = &EdgeCacheRoutingAddSignatures{

			Actions: in.AddSignatures.Actions,

			CopiedParameters: in.AddSignatures.CopiedParameters,

			Keyset: direct.ValueOf(in.AddSignatures.Keyset),

			TokenQueryParameter: direct.ValueOf(in.AddSignatures.TokenQueryParameter),

			TokenTtl: direct.ValueOf(in.AddSignatures.TokenTtl),
		}

	}

	if in.CacheKeyPolicy != nil {

		out.CacheKeyPolicy = &EdgeCacheRoutingCacheKeyPolicy{

			ExcludeHost: direct.ValueOf(in.CacheKeyPolicy.ExcludeHost),

			ExcludeQueryString: direct.ValueOf(in.CacheKeyPolicy.ExcludeQueryString),

			ExcludedQueryParameters: in.CacheKeyPolicy.ExcludedQueryParameters,

			IncludeProtocol: direct.ValueOf(in.CacheKeyPolicy.IncludeProtocol),

			IncludedCookieNames: in.CacheKeyPolicy.IncludedCookieNames,

			IncludedHeaderNames: in.CacheKeyPolicy.IncludedHeaderNames,

			IncludedQueryParameters: in.CacheKeyPolicy.IncludedQueryParameters,
		}

	}

	out.CacheMode = direct.ValueOf(in.CacheMode)

	out.ClientTtl = direct.ValueOf(in.ClientTtl)

	out.DefaultTtl = direct.ValueOf(in.DefaultTtl)

	out.MaxTtl = direct.ValueOf(in.MaxTtl)

	out.NegativeCaching = direct.ValueOf(in.NegativeCaching)

	out.NegativeCachingPolicy = in.NegativeCachingPolicy

	out.SignedRequestMode = direct.ValueOf(in.SignedRequestMode)

	out.SignedRequestKeyset = direct.ValueOf(in.SignedRequestKeyset)

	out.SignedRequestMaximumExpirationTtl = direct.ValueOf(in.SignedRequestMaximumExpirationTtl)

	if in.SignedTokenOptions != nil {

		out.SignedTokenOptions = &EdgeCacheRoutingSignedTokenOptions{

			AllowedSignatureAlgorithms: in.SignedTokenOptions.AllowedSignatureAlgorithms,

			TokenQueryParameter: direct.ValueOf(in.SignedTokenOptions.TokenQueryParameter),
		}

	}

	return out

}

func EdgeCacheRoutingCorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceCorsPolicy) *EdgeCacheRoutingCorsPolicy {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingCorsPolicy{}

	out.AllowCredentials = direct.ValueOf(in.AllowCredentials)

	out.AllowHeaders = in.AllowHeaders

	out.AllowMethods = in.AllowMethods

	out.AllowOrigins = in.AllowOrigins

	out.Disabled = direct.ValueOf(in.Disabled)

	out.ExposeHeaders = in.ExposeHeaders

	out.MaxAge = in.MaxAge

	return out

}

func EdgeCacheRoutingUrlRewrite_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceUrlRewrite) *EdgeCacheRoutingUrlRewrite {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingUrlRewrite{}

	out.HostRewrite = direct.ValueOf(in.HostRewrite)

	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)

	out.PathTemplateRewrite = direct.ValueOf(in.PathTemplateRewrite)

	return out

}

func EdgeCacheRoutingUrlRedirect_ToProto(mapCtx *direct.MapContext, in *krm.EdgecacheserviceUrlRedirect) *EdgeCacheRoutingUrlRedirect {

	if in == nil {

		return nil

	}

	out := &EdgeCacheRoutingUrlRedirect{}

	out.HostRedirect = direct.ValueOf(in.HostRedirect)

	out.HttpsRedirect = direct.ValueOf(in.HttpsRedirect)

	out.PathRedirect = direct.ValueOf(in.PathRedirect)

	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)

	out.RedirectResponseCode = direct.ValueOf(in.RedirectResponseCode)

	out.StripQuery = direct.ValueOf(in.StripQuery)

	return out

}

func EdgeCacheServiceStatus_FromProto(mapCtx *direct.MapContext, in *EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceStatus {

	if in == nil {

		return nil

	}

	out := &krm.NetworkServicesEdgeCacheServiceStatus{}

	out.Ipv4Addresses = in.Ipv4Addresses

	out.Ipv6Addresses = in.Ipv6Addresses

	out.ObservedState = EdgeCacheServiceObservedState_FromProto(mapCtx, in)

	return out

}

func EdgeCacheServiceObservedState_FromProto(mapCtx *direct.MapContext, in *EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceObservedState {

	if in == nil {

		return nil

	}

	out := &krm.NetworkServicesEdgeCacheServiceObservedState{}

	out.Ipv4Addresses = in.Ipv4Addresses

	out.Ipv6Addresses = in.Ipv6Addresses

	return out

}
