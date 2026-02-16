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

// These types are manually defined because they are missing from the GCP Go client library.
// They correspond to the EdgeCacheService REST API.

type EdgeCacheService struct {
	Name                string              `json:"name,omitempty"`
	CreateTime          string              `json:"createTime,omitempty"`
	UpdateTime          string              `json:"updateTime,omitempty"`
	Description         string              `json:"description,omitempty"`
	Labels              map[string]string   `json:"labels,omitempty"`
	DisableQuic         bool                `json:"disableQuic,omitempty"`
	DisableHttp2        bool                `json:"disableHttp2,omitempty"`
	RequireTls          bool                `json:"requireTls,omitempty"`
	EdgeSslCertificates []string            `json:"edgeSslCertificates,omitempty"`
	Ipv4Addresses       []string            `json:"ipv4Addresses,omitempty"`
	Ipv6Addresses       []string            `json:"ipv6Addresses,omitempty"`
	Routing             *EdgeCacheRouting   `json:"routing,omitempty"`
	LogConfig           *EdgeCacheLogConfig `json:"logConfig,omitempty"`
	EdgeSecurityPolicy  string              `json:"edgeSecurityPolicy,omitempty"`
	SslPolicy           string              `json:"sslPolicy,omitempty"`
}

type EdgeCacheLogConfig struct {
	Enable     bool    `json:"enable,omitempty"`
	SampleRate float64 `json:"sampleRate,omitempty"`
}

type EdgeCacheRouting struct {
	HostRules    []*EdgeCacheRoutingHostRule    `json:"hostRules,omitempty"`
	PathMatchers []*EdgeCacheRoutingPathMatcher `json:"pathMatchers,omitempty"`
}

type EdgeCacheRoutingHostRule struct {
	Description string   `json:"description,omitempty"`
	Hosts       []string `json:"hosts,omitempty"`
	PathMatcher string   `json:"pathMatcher,omitempty"`
}

type EdgeCacheRoutingPathMatcher struct {
	Description string                       `json:"description,omitempty"`
	Name        string                       `json:"name,omitempty"`
	RouteRules  []*EdgeCacheRoutingRouteRule `json:"routeRules,omitempty"`
}

type EdgeCacheRoutingRouteRule struct {
	Description  string                        `json:"description,omitempty"`
	Priority     string                        `json:"priority,omitempty"`
	MatchRules   []*EdgeCacheRoutingRouteMatch `json:"matchRules,omitempty"`
	Origin       string                        `json:"origin,omitempty"`
	UrlRedirect  *EdgeCacheRoutingUrlRedirect  `json:"urlRedirect,omitempty"`
	RouteAction  *EdgeCacheRoutingRouteAction  `json:"routeAction,omitempty"`
	HeaderAction *EdgeCacheRoutingHeaderAction `json:"headerAction,omitempty"`
}

type EdgeCacheRoutingRouteMatch struct {
	FullPathMatch         string                                 `json:"fullPathMatch,omitempty"`
	PrefixMatch           string                                 `json:"prefixMatch,omitempty"`
	PathTemplateMatch     string                                 `json:"pathTemplateMatch,omitempty"`
	IgnoreCase            bool                                   `json:"ignoreCase,omitempty"`
	HeaderMatches         []*EdgeCacheRoutingHeaderMatch         `json:"headerMatches,omitempty"`
	QueryParameterMatches []*EdgeCacheRoutingQueryParameterMatch `json:"queryParameterMatches,omitempty"`
}

type EdgeCacheRoutingHeaderMatch struct {
	HeaderName   string `json:"headerName,omitempty"`
	ExactMatch   string `json:"exactMatch,omitempty"`
	PrefixMatch  string `json:"prefixMatch,omitempty"`
	SuffixMatch  string `json:"suffixMatch,omitempty"`
	PresentMatch bool   `json:"presentMatch,omitempty"`
	InvertMatch  bool   `json:"invertMatch,omitempty"`
}

type EdgeCacheRoutingQueryParameterMatch struct {
	Name         string `json:"name,omitempty"`
	ExactMatch   string `json:"exactMatch,omitempty"`
	PresentMatch bool   `json:"presentMatch,omitempty"`
}

type EdgeCacheRoutingRouteAction struct {
	CdnPolicy  *EdgeCacheRoutingCdnPolicy  `json:"cdnPolicy,omitempty"`
	CorsPolicy *EdgeCacheRoutingCorsPolicy `json:"corsPolicy,omitempty"`
	UrlRewrite *EdgeCacheRoutingUrlRewrite `json:"urlRewrite,omitempty"`
}

type EdgeCacheRoutingCdnPolicy struct {
	CacheMode                         string                              `json:"cacheMode,omitempty"`
	ClientTtl                         string                              `json:"clientTtl,omitempty"`
	DefaultTtl                        string                              `json:"defaultTtl,omitempty"`
	MaxTtl                            string                              `json:"maxTtl,omitempty"`
	CacheKeyPolicy                    *EdgeCacheRoutingCacheKeyPolicy     `json:"cacheKeyPolicy,omitempty"`
	NegativeCaching                   bool                                `json:"negativeCaching,omitempty"`
	NegativeCachingPolicy             map[string]string                   `json:"negativeCachingPolicy,omitempty"`
	SignedRequestMode                 string                              `json:"signedRequestMode,omitempty"`
	SignedRequestKeyset               string                              `json:"signedRequestKeyset,omitempty"`
	SignedRequestMaximumExpirationTtl string                              `json:"signedRequestMaximumExpirationTtl,omitempty"`
	AddSignatures                     *EdgeCacheRoutingAddSignatures      `json:"addSignatures,omitempty"`
	SignedTokenOptions                *EdgeCacheRoutingSignedTokenOptions `json:"signedTokenOptions,omitempty"`
}

type EdgeCacheRoutingCacheKeyPolicy struct {
	IncludeProtocol         bool     `json:"includeProtocol,omitempty"`
	ExcludeHost             bool     `json:"excludeHost,omitempty"`
	ExcludeQueryString      bool     `json:"excludeQueryString,omitempty"`
	IncludedQueryParameters []string `json:"includedQueryParameters,omitempty"`
	ExcludedQueryParameters []string `json:"excludedQueryParameters,omitempty"`
	IncludedHeaderNames     []string `json:"includedHeaderNames,omitempty"`
	IncludedCookieNames     []string `json:"includedCookieNames,omitempty"`
}

type EdgeCacheRoutingAddSignatures struct {
	Actions             []string `json:"actions,omitempty"`
	Keyset              string   `json:"keyset,omitempty"`
	TokenTtl            string   `json:"tokenTtl,omitempty"`
	TokenQueryParameter string   `json:"tokenQueryParameter,omitempty"`
	CopiedParameters    []string `json:"copiedParameters,omitempty"`
}

type EdgeCacheRoutingSignedTokenOptions struct {
	AllowedSignatureAlgorithms []string `json:"allowedSignatureAlgorithms,omitempty"`
	TokenQueryParameter        string   `json:"tokenQueryParameter,omitempty"`
}

type EdgeCacheRoutingCorsPolicy struct {
	AllowOrigins     []string `json:"allowOrigins,omitempty"`
	AllowMethods     []string `json:"allowMethods,omitempty"`
	AllowHeaders     []string `json:"allowHeaders,omitempty"`
	ExposeHeaders    []string `json:"exposeHeaders,omitempty"`
	MaxAge           string   `json:"maxAge,omitempty"`
	AllowCredentials bool     `json:"allowCredentials,omitempty"`
	Disabled         bool     `json:"disabled,omitempty"`
}

type EdgeCacheRoutingUrlRewrite struct {
	PathPrefixRewrite   string `json:"pathPrefixRewrite,omitempty"`
	HostRewrite         string `json:"hostRewrite,omitempty"`
	PathTemplateRewrite string `json:"pathTemplateRewrite,omitempty"`
}

type EdgeCacheRoutingUrlRedirect struct {
	HostRedirect         string `json:"hostRedirect,omitempty"`
	PathRedirect         string `json:"pathRedirect,omitempty"`
	PrefixRedirect       string `json:"prefixRedirect,omitempty"`
	RedirectResponseCode string `json:"redirectResponseCode,omitempty"`
	HttpsRedirect        bool   `json:"httpsRedirect,omitempty"`
	StripQuery           bool   `json:"stripQuery,omitempty"`
}

type EdgeCacheRoutingHeaderAction struct {
	RequestHeaderToAdd     []*EdgeCacheRoutingHeaderAdd    `json:"requestHeaderToAdd,omitempty"`
	RequestHeaderToRemove  []*EdgeCacheRoutingHeaderRemove `json:"requestHeaderToRemove,omitempty"`
	ResponseHeaderToAdd    []*EdgeCacheRoutingHeaderAdd    `json:"responseHeaderToAdd,omitempty"`
	ResponseHeaderToRemove []*EdgeCacheRoutingHeaderRemove `json:"responseHeaderToRemove,omitempty"`
}

type EdgeCacheRoutingHeaderAdd struct {
	HeaderName  string `json:"headerName,omitempty"`
	HeaderValue string `json:"headerValue,omitempty"`
	Replace     bool   `json:"replace,omitempty"`
}

type EdgeCacheRoutingHeaderRemove struct {
	HeaderName string `json:"headerName,omitempty"`
}
