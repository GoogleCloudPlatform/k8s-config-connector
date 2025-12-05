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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecacheservice/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EdgeCacheServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceSpec{}
	out.Description = direct.LazyPtr(in.Description)
	out.Labels = in.Labels
	out.DisableQuic = direct.LazyPtr(in.DisableQuic)
	out.DisableHttp2 = direct.LazyPtr(in.DisableHttp2)
	out.RequireTLS = direct.LazyPtr(in.RequireTls)
	out.EdgeSSLCertificates = in.EdgeSslCertificates
	out.EdgeSecurityPolicy = direct.LazyPtr(in.EdgeSecurityPolicy)

	if in.LogConfig != nil {
		out.LogConfig = &krm.LogConfig{
			Enable:     direct.LazyPtr(in.LogConfig.Enable),
			SampleRate: direct.LazyPtr(in.LogConfig.SampleRate),
		}
	}

	if in.Routing != nil {
		out.Routing = &krm.Routing{}
		for _, hr := range in.Routing.HostRules {
			out.Routing.HostRules = append(out.Routing.HostRules, krm.HostRule{
				Description: direct.LazyPtr(hr.Description),
				Hosts:       hr.Hosts,
				PathMatcher: direct.LazyPtr(hr.PathMatcher),
			})
		}
		for _, pm := range in.Routing.PathMatchers {
			krmPm := krm.PathMatcher{
				Name:        direct.LazyPtr(pm.Name),
				Description: direct.LazyPtr(pm.Description),
			}
			for _, rr := range pm.RouteRules {
				krmRr := krm.RouteRule{
					Priority:     direct.LazyPtr(rr.Priority),
					Description:  direct.LazyPtr(rr.Description),
					Origin:       direct.LazyPtr(rr.Origin),
					RouteMethods: rr.RouteMethods,
				}
				if rr.HeaderAction != nil {
					krmRr.HeaderAction = &krm.HeaderAction{}
					for _, hta := range rr.HeaderAction.HeadersToAdd {
						krmRr.HeaderAction.HeadersToAdd = append(krmRr.HeaderAction.HeadersToAdd, krm.HeaderToAdd{
							HeaderName:  direct.LazyPtr(hta.HeaderName),
							HeaderValue: direct.LazyPtr(hta.HeaderValue),
							Replace:     direct.LazyPtr(hta.Replace),
						})
					}
					for _, htr := range rr.HeaderAction.HeadersToRemove {
						krmRr.HeaderAction.HeadersToRemove = append(krmRr.HeaderAction.HeadersToRemove, krm.HeaderToRemove{
							HeaderName: direct.LazyPtr(htr.HeaderName),
						})
					}
				}
				if rr.RouteAction != nil {
					krmRr.RouteAction = &krm.RouteAction{
						CompressionMode: direct.LazyPtr(rr.RouteAction.CompressionMode),
					}
					if rr.RouteAction.CdnPolicy != nil {
						krmRr.RouteAction.CDNPolicy = &krm.CDNPolicy{
							CacheMode:              direct.LazyPtr(rr.RouteAction.CdnPolicy.CacheMode),
							ClientTTL:              rr.RouteAction.CdnPolicy.ClientTtl,
							DefaultTTL:             direct.LazyPtr(rr.RouteAction.CdnPolicy.DefaultTtl.String()),
							MaxTTL:                 direct.LazyPtr(rr.RouteAction.CdnPolicy.MaxTtl.String()),
							CacheNegativeCallbacks: direct.LazyPtr(rr.RouteAction.CdnPolicy.CacheNegativeCallbacks),
							SignedRequestKeyset:    rr.RouteAction.CdnPolicy.SignedRequestKeyset,
							SignedRequestMode:      direct.LazyPtr(rr.RouteAction.CdnPolicy.SignedRequestMode),
						}
					}
					if rr.RouteAction.UrlRewrite != nil {
						krmRr.RouteAction.URLRewrite = &krm.URLRewrite{
							PathPrefixRewrite: direct.LazyPtr(rr.RouteAction.UrlRewrite.PathPrefixRewrite),
							HostRewrite:       direct.LazyPtr(rr.RouteAction.UrlRewrite.HostRewrite),
						}
					}
					if rr.RouteAction.CorsPolicy != nil {
						krmRr.RouteAction.CorsPolicy = &krm.CorsPolicy{
							MaxAge:           direct.LazyPtr(rr.RouteAction.CorsPolicy.MaxAge.String()),
							AllowOrigins:     rr.RouteAction.CorsPolicy.AllowOrigins,
							AllowMethods:     rr.RouteAction.CorsPolicy.AllowMethods,
							AllowHeaders:     rr.RouteAction.CorsPolicy.AllowHeaders,
							ExposeHeaders:    rr.RouteAction.CorsPolicy.ExposeHeaders,
							AllowCredentials: direct.LazyPtr(rr.RouteAction.CorsPolicy.AllowCredentials),
							Disabled:         direct.LazyPtr(rr.RouteAction.CorsPolicy.Disabled),
						}
					}
				}
				if rr.UrlRedirect != nil {
					krmRr.URLRedirect = &krm.URLRedirect{
						HostRedirect:         direct.LazyPtr(rr.UrlRedirect.HostRedirect),
						PathRedirect:         direct.LazyPtr(rr.UrlRedirect.PathRedirect),
						PrefixRedirect:       direct.LazyPtr(rr.UrlRedirect.PrefixRedirect),
						RedirectResponseCode: direct.LazyPtr(rr.UrlRedirect.RedirectResponseCode),
						HTTPSRedirect:        direct.LazyPtr(rr.UrlRedirect.HttpsRedirect),
						StripQuery:           direct.LazyPtr(rr.UrlRedirect.StripQuery),
					}
				}
				for _, mr := range rr.MatchRules {
					krmMr := krm.MatchRule{
						PrefixMatch:   direct.LazyPtr(mr.PrefixMatch),
						FullPathMatch: direct.LazyPtr(mr.FullPathMatch),
						RegexMatch:    direct.LazyPtr(mr.RegexMatch),
						IgnoreCase:    direct.LazyPtr(mr.IgnoreCase),
					}
					for _, hm := range mr.HeaderMatches {
						krmMr.HeaderMatches = append(krmMr.HeaderMatches, krm.HeaderMatch{
							HeaderName:   direct.LazyPtr(hm.HeaderName),
							PresentMatch: direct.LazyPtr(hm.PresentMatch),
							ExactMatch:   direct.LazyPtr(hm.ExactMatch),
							PrefixMatch:  direct.LazyPtr(hm.PrefixMatch),
							SuffixMatch:  direct.LazyPtr(hm.SuffixMatch),
							RegexMatch:   direct.LazyPtr(hm.RegexMatch),
							InvertMatch:  direct.LazyPtr(hm.InvertMatch),
						})
					}
					for _, qpm := range mr.QueryParamMatches {
						krmMr.QueryParamMatches = append(krmMr.QueryParamMatches, krm.QueryParamMatch{
							QueryParam:   direct.LazyPtr(qpm.QueryParam),
							PresentMatch: direct.LazyPtr(qpm.PresentMatch),
							ExactMatch:   direct.LazyPtr(qpm.ExactMatch),
							RegexMatch:   direct.LazyPtr(qpm.RegexMatch),
						})
					}
					krmRr.MatchRules = append(krmRr.MatchRules, krmMr)
				}
				krmPm.RouteRules = append(krmPm.RouteRules, krmRr)
			}
			out.Routing.PathMatchers = append(out.Routing.PathMatchers, krmPm)
		}
	}

	return out
}

func EdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	out.RequireTls = direct.ValueOf(in.RequireTLS)
	out.EdgeSslCertificates = in.EdgeSSLCertificates
	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)
	out.LogConfig = LogConfig_ToProto(mapCtx, in.LogConfig)
	out.Routing = Routing_ToProto(mapCtx, in.Routing)
	return out
}
