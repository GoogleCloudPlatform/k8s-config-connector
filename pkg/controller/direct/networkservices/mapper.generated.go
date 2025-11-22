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
// proto.service: google.cloud.networkservices.v1

package networkservices

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AddHeader_FromProto(mapCtx *direct.MapContext, in *pb.AddHeader) *krm.EdgeCacheService_AddHeader {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_AddHeader{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.HeaderValue = direct.LazyPtr(in.GetHeaderValue())
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}
func AddHeader_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_AddHeader) *pb.AddHeader {
	if in == nil {
		return nil
	}
	out := &pb.AddHeader{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}
func AddSignatures_FromProto(mapCtx *direct.MapContext, in *pb.AddSignatures) *krm.EdgeCacheService_AddSignatures {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_AddSignatures{}
	out.Actions = direct.EnumSlice_FromProto(mapCtx, in.Actions)
	out.Keyset = direct.LazyPtr(in.GetKeyset())
	out.TokenTtl = direct.StringDuration_FromProto(mapCtx, in.GetTokenTtl())
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.CopiedParameters = in.CopiedParameters
	return out
}
func AddSignatures_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_AddSignatures) *pb.AddSignatures {
	if in == nil {
		return nil
	}
	out := &pb.AddSignatures{}
	out.Actions = direct.EnumSlice_ToProto[pb.SignatureAction](mapCtx, in.Actions)
	out.Keyset = direct.ValueOf(in.Keyset)
	out.TokenTtl = direct.StringDuration_ToProto(mapCtx, in.TokenTtl)
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.CopiedParameters = in.CopiedParameters
	return out
}
func CDNPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CdnPolicy) *krm.EdgeCacheService_CdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CdnPolicy{}
	out.CacheMode = direct.Enum_FromProto(mapCtx, in.GetCacheMode())
	out.ClientTtl = direct.StringDuration_FromProto(mapCtx, in.GetClientTtl())
	out.DefaultTtl = direct.StringDuration_FromProto(mapCtx, in.GetDefaultTtl())
	out.MaxTtl = direct.StringDuration_FromProto(mapCtx, in.GetMaxTtl())
	out.NegativeCaching = direct.LazyPtr(in.GetNegativeCaching())
	// MISSING: NegativeCachingPolicy
	out.CacheKeyPolicy = CacheKeyPolicy_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.SignedRequestMode = direct.Enum_FromProto(mapCtx, in.GetSignedRequestMode())
	out.SignedRequestKeyset = direct.LazyPtr(in.GetSignedRequestKeyset())
	out.SignedTokenOptions = SignedTokenOptions_FromProto(mapCtx, in.GetSignedTokenOptions())
	out.AddSignatures = AddSignatures_FromProto(mapCtx, in.GetAddSignatures())
	out.SignedRequestMaximumExpirationTtl = direct.StringDuration_FromProto(mapCtx, in.GetSignedRequestMaximumExpirationTtl())
	return out
}
func CDNPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CdnPolicy) *pb.CdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CdnPolicy{}
	out.CacheMode = direct.Enum_ToProto[pb.CacheMode](mapCtx, in.CacheMode)
	out.ClientTtl = direct.StringDuration_ToProto(mapCtx, in.ClientTtl)
	out.DefaultTtl = direct.StringDuration_ToProto(mapCtx, in.DefaultTtl)
	out.MaxTtl = direct.StringDuration_ToProto(mapCtx, in.MaxTtl)
	out.NegativeCaching = direct.ValueOf(in.NegativeCaching)
	// MISSING: NegativeCachingPolicy
	out.CacheKeyPolicy = CacheKeyPolicy_ToProto(mapCtx, in.CacheKeyPolicy)
	out.SignedRequestMode = direct.Enum_ToProto[pb.SignedRequestMode](mapCtx, in.SignedRequestMode)
	out.SignedRequestKeyset = direct.ValueOf(in.SignedRequestKeyset)
	out.SignedTokenOptions = SignedTokenOptions_ToProto(mapCtx, in.SignedTokenOptions)
	out.AddSignatures = AddSignatures_ToProto(mapCtx, in.AddSignatures)
	out.SignedRequestMaximumExpirationTtl = direct.StringDuration_ToProto(mapCtx, in.SignedRequestMaximumExpirationTtl)
	return out
}
func CacheKeyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CacheKeyPolicy) *krm.EdgeCacheService_CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CacheKeyPolicy{}
	out.IncludeProtocol = direct.LazyPtr(in.GetIncludeProtocol())
	out.ExcludeHost = direct.LazyPtr(in.GetExcludeHost())
	out.ExcludeQueryString = direct.LazyPtr(in.GetExcludeQueryString())
	out.IncludedQueryParameters = in.IncludedQueryParameters
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedCookieNames = in.IncludedCookieNames
	return out
}
func CacheKeyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CacheKeyPolicy) *pb.CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CacheKeyPolicy{}
	out.IncludeProtocol = direct.ValueOf(in.IncludeProtocol)
	out.ExcludeHost = direct.ValueOf(in.ExcludeHost)
	out.ExcludeQueryString = direct.ValueOf(in.ExcludeQueryString)
	out.IncludedQueryParameters = in.IncludedQueryParameters
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedCookieNames = in.IncludedCookieNames
	return out
}
func CorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.EdgeCacheService_CorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CorsPolicy{}
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.AllowOrigins = in.AllowOrigins
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func CorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CorsPolicy) *pb.CorsPolicy {
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

// TODO: EdgeCacheService_FromProto/ToProto commented out - they map to full CRD object not Spec
// The controller uses NetworkServicesEdgeCacheServiceSpec_FromProto/ToProto instead
/*
func EdgeCacheService_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheService {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheService{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Routing = Routing_FromProto(mapCtx, in.GetRouting())
	out.EdgeSSLCertificates = in.EdgeSslCertificates
	out.DisableQuic = direct.LazyPtr(in.GetDisableQuic())
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	out.LogConfig = LogConfig_FromProto(mapCtx, in.GetLogConfig())
	out.DisableHttp2 = direct.LazyPtr(in.GetDisableHttp2())
	out.RequireTLS = direct.LazyPtr(in.GetRequireTls())
	out.EdgeSecurityPolicy = direct.LazyPtr(in.GetEdgeSecurityPolicy())
	return out
}
func EdgeCacheService_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheService) *pb.EdgeCacheService {
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
	out.EdgeSslCertificates = in.EdgeSSLCertificates
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	out.LogConfig = LogConfig_ToProto(mapCtx, in.LogConfig)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	out.RequireTls = direct.ValueOf(in.RequireTLS)
	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)
	return out
}
*/
func EdgeCacheServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Routing
	// MISSING: EdgeSSLCertificates
	// MISSING: DisableQuic
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	// MISSING: LogConfig
	// MISSING: DisableHttp2
	// MISSING: RequireTLS
	// MISSING: EdgeSecurityPolicy
	return out
}
func EdgeCacheServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceObservedState) *pb.EdgeCacheService {
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
	// MISSING: EdgeSSLCertificates
	// MISSING: DisableQuic
	out.Ipv4Addresses = in.Ipv4Addresses
	out.Ipv6Addresses = in.Ipv6Addresses
	// MISSING: LogConfig
	// MISSING: DisableHttp2
	// MISSING: RequireTLS
	// MISSING: EdgeSecurityPolicy
	return out
}
func EdgeCacheService_AddHeader_FromProto(mapCtx *direct.MapContext, in *pb.AddHeader) *krm.EdgeCacheService_AddHeader {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_AddHeader{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.HeaderValue = direct.LazyPtr(in.GetHeaderValue())
	out.Replace = direct.LazyPtr(in.GetReplace())
	return out
}
func EdgeCacheService_AddHeader_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_AddHeader) *pb.AddHeader {
	if in == nil {
		return nil
	}
	out := &pb.AddHeader{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}
func EdgeCacheService_AddSignatures_FromProto(mapCtx *direct.MapContext, in *pb.AddSignatures) *krm.EdgeCacheService_AddSignatures {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_AddSignatures{}
	out.Actions = direct.EnumSlice_FromProto(mapCtx, in.Actions)
	out.Keyset = direct.LazyPtr(in.GetKeyset())
	out.TokenTtl = direct.StringDuration_FromProto(mapCtx, in.GetTokenTtl())
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.CopiedParameters = in.CopiedParameters
	return out
}
func EdgeCacheService_AddSignatures_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_AddSignatures) *pb.AddSignatures {
	if in == nil {
		return nil
	}
	out := &pb.AddSignatures{}
	out.Actions = direct.EnumSlice_ToProto[pb.SignatureAction](mapCtx, in.Actions)
	out.Keyset = direct.ValueOf(in.Keyset)
	out.TokenTtl = direct.StringDuration_ToProto(mapCtx, in.TokenTtl)
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.CopiedParameters = in.CopiedParameters
	return out
}
func EdgeCacheService_CacheKeyPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CacheKeyPolicy) *krm.EdgeCacheService_CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CacheKeyPolicy{}
	out.IncludeProtocol = direct.LazyPtr(in.GetIncludeProtocol())
	out.ExcludeHost = direct.LazyPtr(in.GetExcludeHost())
	out.ExcludeQueryString = direct.LazyPtr(in.GetExcludeQueryString())
	out.IncludedQueryParameters = in.IncludedQueryParameters
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedCookieNames = in.IncludedCookieNames
	return out
}
func EdgeCacheService_CacheKeyPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CacheKeyPolicy) *pb.CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CacheKeyPolicy{}
	out.IncludeProtocol = direct.ValueOf(in.IncludeProtocol)
	out.ExcludeHost = direct.ValueOf(in.ExcludeHost)
	out.ExcludeQueryString = direct.ValueOf(in.ExcludeQueryString)
	out.IncludedQueryParameters = in.IncludedQueryParameters
	out.ExcludedQueryParameters = in.ExcludedQueryParameters
	out.IncludedHeaderNames = in.IncludedHeaderNames
	out.IncludedCookieNames = in.IncludedCookieNames
	return out
}
func EdgeCacheService_CdnPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CdnPolicy) *krm.EdgeCacheService_CdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CdnPolicy{}
	out.CacheMode = direct.Enum_FromProto(mapCtx, in.GetCacheMode())
	// MISSING: ClientTTL
	// (near miss): "ClientTTL" vs "ClientTtl"
	// MISSING: DefaultTTL
	// (near miss): "DefaultTTL" vs "DefaultTtl"
	// MISSING: MaxTTL
	// (near miss): "MaxTTL" vs "MaxTtl"
	out.NegativeCaching = direct.LazyPtr(in.GetNegativeCaching())
	// TODO: map type string message for field NegativeCachingPolicy
	out.CacheKeyPolicy = EdgeCacheService_CacheKeyPolicy_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.SignedRequestMode = direct.Enum_FromProto(mapCtx, in.GetSignedRequestMode())
	out.SignedRequestKeyset = direct.LazyPtr(in.GetSignedRequestKeyset())
	out.SignedTokenOptions = EdgeCacheService_SignedTokenOptions_FromProto(mapCtx, in.GetSignedTokenOptions())
	out.AddSignatures = EdgeCacheService_AddSignatures_FromProto(mapCtx, in.GetAddSignatures())
	// MISSING: SignedRequestMaximumExpirationTTL
	// (near miss): "SignedRequestMaximumExpirationTTL" vs "SignedRequestMaximumExpirationTtl"
	return out
}
func EdgeCacheService_CdnPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CdnPolicy) *pb.CdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CdnPolicy{}
	out.CacheMode = direct.Enum_ToProto[pb.CacheMode](mapCtx, in.CacheMode)
	// MISSING: ClientTTL
	// (near miss): "ClientTTL" vs "ClientTtl"
	// MISSING: DefaultTTL
	// (near miss): "DefaultTTL" vs "DefaultTtl"
	// MISSING: MaxTTL
	// (near miss): "MaxTTL" vs "MaxTtl"
	out.NegativeCaching = direct.ValueOf(in.NegativeCaching)
	// TODO: map type string message for field NegativeCachingPolicy
	out.CacheKeyPolicy = EdgeCacheService_CacheKeyPolicy_ToProto(mapCtx, in.CacheKeyPolicy)
	out.SignedRequestMode = direct.Enum_ToProto[pb.SignedRequestMode](mapCtx, in.SignedRequestMode)
	out.SignedRequestKeyset = direct.ValueOf(in.SignedRequestKeyset)
	out.SignedTokenOptions = EdgeCacheService_SignedTokenOptions_ToProto(mapCtx, in.SignedTokenOptions)
	out.AddSignatures = EdgeCacheService_AddSignatures_ToProto(mapCtx, in.AddSignatures)
	// MISSING: SignedRequestMaximumExpirationTTL
	// (near miss): "SignedRequestMaximumExpirationTTL" vs "SignedRequestMaximumExpirationTtl"
	return out
}
func EdgeCacheService_CorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.EdgeCacheService_CorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_CorsPolicy{}
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.AllowOrigins = in.AllowOrigins
	out.AllowMethods = in.AllowMethods
	out.AllowHeaders = in.AllowHeaders
	out.ExposeHeaders = in.ExposeHeaders
	out.AllowCredentials = direct.LazyPtr(in.GetAllowCredentials())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}
func EdgeCacheService_CorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_CorsPolicy) *pb.CorsPolicy {
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
func EdgeCacheService_HeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.HeaderAction) *krm.EdgeCacheService_HeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_FromProto(mapCtx, in.RequestHeadersToAdd, EdgeCacheService_AddHeader_FromProto)
	out.RequestHeadersToRemove = direct.Slice_FromProto(mapCtx, in.RequestHeadersToRemove, EdgeCacheService_RemoveHeader_FromProto)
	out.ResponseHeadersToAdd = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToAdd, EdgeCacheService_AddHeader_FromProto)
	out.ResponseHeadersToRemove = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToRemove, EdgeCacheService_RemoveHeader_FromProto)
	return out
}
func EdgeCacheService_HeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HeaderAction) *pb.HeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.HeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_ToProto(mapCtx, in.RequestHeadersToAdd, EdgeCacheService_AddHeader_ToProto)
	out.RequestHeadersToRemove = direct.Slice_ToProto(mapCtx, in.RequestHeadersToRemove, EdgeCacheService_RemoveHeader_ToProto)
	out.ResponseHeadersToAdd = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToAdd, EdgeCacheService_AddHeader_ToProto)
	out.ResponseHeadersToRemove = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToRemove, EdgeCacheService_RemoveHeader_ToProto)
	return out
}
func EdgeCacheService_HeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.HeaderMatch) *krm.EdgeCacheService_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HeaderMatch{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	return out
}
func EdgeCacheService_HeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HeaderMatch) *pb.HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HeaderMatch{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	if oneof := EdgeCacheService_HeaderMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := EdgeCacheService_HeaderMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := EdgeCacheService_HeaderMatch_PrefixMatch_ToProto(mapCtx, in.PrefixMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := EdgeCacheService_HeaderMatch_SuffixMatch_ToProto(mapCtx, in.SuffixMatch); oneof != nil {
		out.MatchType = oneof
	}
	return out
}
func EdgeCacheService_HeaderMatch_PresentMatch_ToProto(mapCtx *direct.MapContext, in *bool) *pb.HeaderMatch_PresentMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_PresentMatch{PresentMatch: *in}
}
func EdgeCacheService_HeaderMatch_ExactMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_ExactMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_ExactMatch{ExactMatch: *in}
}
func EdgeCacheService_HeaderMatch_PrefixMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_PrefixMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_PrefixMatch{PrefixMatch: *in}
}
func EdgeCacheService_HeaderMatch_SuffixMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_SuffixMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_SuffixMatch{SuffixMatch: *in}
}
func EdgeCacheService_HostRule_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.EdgeCacheService_HostRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HostRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hosts = in.Hosts
	out.PathMatcher = direct.LazyPtr(in.GetPathMatcher())
	return out
}
func EdgeCacheService_HostRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Hosts = in.Hosts
	out.PathMatcher = direct.ValueOf(in.PathMatcher)
	return out
}
func EdgeCacheService_LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.LogConfig) *krm.EdgeCacheService_LogConfig {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_LogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	if in.GetSampleRate() != 0 {
		rate := float64(in.GetSampleRate())
		out.SampleRate = &rate
	}
	return out
}
func EdgeCacheService_LogConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_LogConfig) *pb.LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.LogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	if in.SampleRate != nil {
		rate := float32(*in.SampleRate)
		out.SampleRate = rate
	}
	return out
}
func EdgeCacheService_MatchRule_FromProto(mapCtx *direct.MapContext, in *pb.MatchRule) *krm.EdgeCacheService_MatchRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_MatchRule{}
	out.PathTemplateMatch = direct.LazyPtr(in.GetPathTemplateMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.HeaderMatches = direct.Slice_FromProto(mapCtx, in.HeaderMatches, EdgeCacheService_HeaderMatch_FromProto)
	out.QueryParameterMatches = direct.Slice_FromProto(mapCtx, in.QueryParameterMatches, EdgeCacheService_QueryParameterMatch_FromProto)
	return out
}
func EdgeCacheService_MatchRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_MatchRule) *pb.MatchRule {
	if in == nil {
		return nil
	}
	out := &pb.MatchRule{}
	out.PathTemplateMatch = direct.ValueOf(in.PathTemplateMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.HeaderMatches = direct.Slice_ToProto(mapCtx, in.HeaderMatches, EdgeCacheService_HeaderMatch_ToProto)
	out.QueryParameterMatches = direct.Slice_ToProto(mapCtx, in.QueryParameterMatches, EdgeCacheService_QueryParameterMatch_ToProto)
	return out
}
func EdgeCacheService_PathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.EdgeCacheService_PathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_PathMatcher{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, EdgeCacheService_RouteRule_FromProto)
	return out
}
func EdgeCacheService_PathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_PathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, EdgeCacheService_RouteRule_ToProto)
	return out
}
func EdgeCacheService_QueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameterMatch) *krm.EdgeCacheService_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_QueryParameterMatch{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	return out
}
func EdgeCacheService_QueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_QueryParameterMatch) *pb.QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameterMatch{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := EdgeCacheService_QueryParameterMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := EdgeCacheService_QueryParameterMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	return out
}
func EdgeCacheService_QueryParameterMatch_PresentMatch_ToProto(mapCtx *direct.MapContext, in *bool) *pb.QueryParameterMatch_PresentMatch {
	if in == nil {
		return nil
	}
	return &pb.QueryParameterMatch_PresentMatch{PresentMatch: *in}
}
func EdgeCacheService_QueryParameterMatch_ExactMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.QueryParameterMatch_ExactMatch {
	if in == nil {
		return nil
	}
	return &pb.QueryParameterMatch_ExactMatch{ExactMatch: *in}
}
func EdgeCacheService_RemoveHeader_FromProto(mapCtx *direct.MapContext, in *pb.RemoveHeader) *krm.EdgeCacheService_RemoveHeader {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RemoveHeader{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	return out
}
func EdgeCacheService_RemoveHeader_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RemoveHeader) *pb.RemoveHeader {
	if in == nil {
		return nil
	}
	out := &pb.RemoveHeader{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	return out
}
func EdgeCacheService_RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.RouteAction) *krm.EdgeCacheService_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteAction{}
	// MISSING: CDNPolicy
	// (near miss): "CDNPolicy" vs "CdnPolicy"
	// MISSING: URLRewrite
	// (near miss): "URLRewrite" vs "UrlRewrite"
	out.CorsPolicy = EdgeCacheService_CorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	out.CompressionMode = direct.Enum_FromProto(mapCtx, in.GetCompressionMode())
	return out
}
func EdgeCacheService_RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteAction) *pb.RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.RouteAction{}
	// MISSING: CDNPolicy
	// (near miss): "CDNPolicy" vs "CdnPolicy"
	// MISSING: URLRewrite
	// (near miss): "URLRewrite" vs "UrlRewrite"
	out.CorsPolicy = EdgeCacheService_CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.CompressionMode = direct.Enum_ToProto[pb.CompressionMode](mapCtx, in.CompressionMode)
	return out
}
func EdgeCacheService_RouteMethods_FromProto(mapCtx *direct.MapContext, in *pb.RouteMethods) *krm.EdgeCacheService_RouteMethods {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteMethods{}
	out.AllowedMethods = in.AllowedMethods
	return out
}
func EdgeCacheService_RouteMethods_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteMethods) *pb.RouteMethods {
	if in == nil {
		return nil
	}
	out := &pb.RouteMethods{}
	out.AllowedMethods = in.AllowedMethods
	return out
}
func EdgeCacheService_RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.RouteRule) *krm.EdgeCacheService_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteRule{}
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, EdgeCacheService_MatchRule_FromProto)
	out.HeaderAction = EdgeCacheService_HeaderAction_FromProto(mapCtx, in.GetHeaderAction())
	out.RouteAction = EdgeCacheService_RouteAction_FromProto(mapCtx, in.GetRouteAction())
	// MISSING: URLRedirect
	// (near miss): "URLRedirect" vs "UrlRedirect"
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.RouteMethods = EdgeCacheService_RouteMethods_FromProto(mapCtx, in.GetRouteMethods())
	return out
}
func EdgeCacheService_RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteRule) *pb.RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.RouteRule{}
	out.Priority = direct.ValueOf(in.Priority)
	out.Description = direct.ValueOf(in.Description)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, EdgeCacheService_MatchRule_ToProto)
	out.HeaderAction = EdgeCacheService_HeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.RouteAction = EdgeCacheService_RouteAction_ToProto(mapCtx, in.RouteAction)
	// MISSING: URLRedirect
	// (near miss): "URLRedirect" vs "UrlRedirect"
	out.Origin = direct.ValueOf(in.Origin)
	out.RouteMethods = EdgeCacheService_RouteMethods_ToProto(mapCtx, in.RouteMethods)
	return out
}
func EdgeCacheService_Routing_FromProto(mapCtx *direct.MapContext, in *pb.Routing) *krm.EdgeCacheService_Routing {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_Routing{}
	out.HostRules = direct.Slice_FromProto(mapCtx, in.HostRules, EdgeCacheService_HostRule_FromProto)
	out.PathMatchers = direct.Slice_FromProto(mapCtx, in.PathMatchers, EdgeCacheService_PathMatcher_FromProto)
	return out
}
func EdgeCacheService_Routing_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_Routing) *pb.Routing {
	if in == nil {
		return nil
	}
	out := &pb.Routing{}
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRules, EdgeCacheService_HostRule_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatchers, EdgeCacheService_PathMatcher_ToProto)
	return out
}
func EdgeCacheService_SignedTokenOptions_FromProto(mapCtx *direct.MapContext, in *pb.SignedTokenOptions) *krm.EdgeCacheService_SignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_SignedTokenOptions{}
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.AllowedSignatureAlgorithms = direct.EnumSlice_FromProto(mapCtx, in.AllowedSignatureAlgorithms)
	return out
}
func EdgeCacheService_SignedTokenOptions_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_SignedTokenOptions) *pb.SignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &pb.SignedTokenOptions{}
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.AllowedSignatureAlgorithms = direct.EnumSlice_ToProto[pb.SignatureAlgorithm](mapCtx, in.AllowedSignatureAlgorithms)
	return out
}
func EdgeCacheService_UrlRedirect_FromProto(mapCtx *direct.MapContext, in *pb.UrlRedirect) *krm.EdgeCacheService_UrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_UrlRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRedirect = direct.LazyPtr(in.GetPrefixRedirect())
	out.RedirectResponseCode = direct.Enum_FromProto(mapCtx, in.GetRedirectResponseCode())
	// MISSING: HTTPSRedirect
	// (near miss): "HTTPSRedirect" vs "HttpsRedirect"
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	return out
}
func EdgeCacheService_UrlRedirect_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_UrlRedirect) *pb.UrlRedirect {
	if in == nil {
		return nil
	}
	out := &pb.UrlRedirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)
	out.RedirectResponseCode = direct.Enum_ToProto[pb.RedirectResponseCode](mapCtx, in.RedirectResponseCode)
	// MISSING: HTTPSRedirect
	// (near miss): "HTTPSRedirect" vs "HttpsRedirect"
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}
func EdgeCacheService_UrlRewrite_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.EdgeCacheService_UrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_UrlRewrite{}
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.PathTemplateRewrite = direct.LazyPtr(in.GetPathTemplateRewrite())
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	return out
}
func EdgeCacheService_UrlRewrite_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_UrlRewrite) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.PathTemplateRewrite = direct.ValueOf(in.PathTemplateRewrite)
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	return out
}
func HeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.HeaderAction) *krm.EdgeCacheService_HeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_FromProto(mapCtx, in.RequestHeadersToAdd, AddHeader_FromProto)
	out.RequestHeadersToRemove = direct.Slice_FromProto(mapCtx, in.RequestHeadersToRemove, RemoveHeader_FromProto)
	out.ResponseHeadersToAdd = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToAdd, AddHeader_FromProto)
	out.ResponseHeadersToRemove = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToRemove, RemoveHeader_FromProto)
	return out
}
func HeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HeaderAction) *pb.HeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.HeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_ToProto(mapCtx, in.RequestHeadersToAdd, AddHeader_ToProto)
	out.RequestHeadersToRemove = direct.Slice_ToProto(mapCtx, in.RequestHeadersToRemove, RemoveHeader_ToProto)
	out.ResponseHeadersToAdd = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToAdd, AddHeader_ToProto)
	out.ResponseHeadersToRemove = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToRemove, RemoveHeader_ToProto)
	return out
}
func HeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.HeaderMatch) *krm.EdgeCacheService_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HeaderMatch{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.SuffixMatch = direct.LazyPtr(in.GetSuffixMatch())
	return out
}
func HeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HeaderMatch) *pb.HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HeaderMatch{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	if oneof := HeaderMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HeaderMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HeaderMatch_PrefixMatch_ToProto(mapCtx, in.PrefixMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := HeaderMatch_SuffixMatch_ToProto(mapCtx, in.SuffixMatch); oneof != nil {
		out.MatchType = oneof
	}
	return out
}
func HeaderMatch_PresentMatch_ToProto(mapCtx *direct.MapContext, in *bool) *pb.HeaderMatch_PresentMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_PresentMatch{PresentMatch: *in}
}
func HeaderMatch_ExactMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_ExactMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_ExactMatch{ExactMatch: *in}
}
func HeaderMatch_PrefixMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_PrefixMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_PrefixMatch{PrefixMatch: *in}
}
func HeaderMatch_SuffixMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.HeaderMatch_SuffixMatch {
	if in == nil {
		return nil
	}
	return &pb.HeaderMatch_SuffixMatch{SuffixMatch: *in}
}
func HostRule_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.EdgeCacheService_HostRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_HostRule{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hosts = in.Hosts
	out.PathMatcher = direct.LazyPtr(in.GetPathMatcher())
	return out
}
func HostRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_HostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = direct.ValueOf(in.Description)
	out.Hosts = in.Hosts
	out.PathMatcher = direct.ValueOf(in.PathMatcher)
	return out
}
func LogConfig_FromProto(mapCtx *direct.MapContext, in *pb.LogConfig) *krm.EdgeCacheService_LogConfig {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_LogConfig{}
	out.Enable = direct.LazyPtr(in.GetEnable())
	if in.GetSampleRate() != 0 {
		rate := float64(in.GetSampleRate())
		out.SampleRate = &rate
	}
	return out
}
func LogConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_LogConfig) *pb.LogConfig {
	if in == nil {
		return nil
	}
	out := &pb.LogConfig{}
	out.Enable = direct.ValueOf(in.Enable)
	if in.SampleRate != nil {
		rate := float32(*in.SampleRate)
		out.SampleRate = rate
	}
	return out
}
func MatchRule_FromProto(mapCtx *direct.MapContext, in *pb.MatchRule) *krm.EdgeCacheService_MatchRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_MatchRule{}
	out.PathTemplateMatch = direct.LazyPtr(in.GetPathTemplateMatch())
	out.PrefixMatch = direct.LazyPtr(in.GetPrefixMatch())
	out.FullPathMatch = direct.LazyPtr(in.GetFullPathMatch())
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	out.HeaderMatches = direct.Slice_FromProto(mapCtx, in.HeaderMatches, HeaderMatch_FromProto)
	out.QueryParameterMatches = direct.Slice_FromProto(mapCtx, in.QueryParameterMatches, QueryParameterMatch_FromProto)
	return out
}
func MatchRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_MatchRule) *pb.MatchRule {
	if in == nil {
		return nil
	}
	out := &pb.MatchRule{}
	out.PathTemplateMatch = direct.ValueOf(in.PathTemplateMatch)
	out.PrefixMatch = direct.ValueOf(in.PrefixMatch)
	out.FullPathMatch = direct.ValueOf(in.FullPathMatch)
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	out.HeaderMatches = direct.Slice_ToProto(mapCtx, in.HeaderMatches, HeaderMatch_ToProto)
	out.QueryParameterMatches = direct.Slice_ToProto(mapCtx, in.QueryParameterMatches, QueryParameterMatch_ToProto)
	return out
}
func NetworkServicesEdgeCacheServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: EdgeSSLCertificates
	// MISSING: IPV4Addresses
	// (near miss): "IPV4Addresses" vs "Ipv4Addresses"
	// MISSING: IPV6Addresses
	// (near miss): "IPV6Addresses" vs "Ipv6Addresses"
	// MISSING: RequireTLS
	return out
}
func NetworkServicesEdgeCacheServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceObservedState) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: EdgeSSLCertificates
	// MISSING: IPV4Addresses
	// (near miss): "IPV4Addresses" vs "Ipv4Addresses"
	// MISSING: IPV6Addresses
	// (near miss): "IPV6Addresses" vs "Ipv6Addresses"
	// MISSING: RequireTLS
	return out
}
func NetworkServicesEdgeCacheServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeCacheService) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesEdgeCacheServiceSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Routing = EdgeCacheService_Routing_FromProto(mapCtx, in.GetRouting())
	// MISSING: EdgeSSLCertificates
	out.DisableQuic = direct.LazyPtr(in.GetDisableQuic())
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	out.LogConfig = EdgeCacheService_LogConfig_FromProto(mapCtx, in.GetLogConfig())
	out.DisableHttp2 = direct.LazyPtr(in.GetDisableHttp2())
	// MISSING: RequireTLS
	// (near miss): "RequireTLS" vs "RequireTls"
	out.EdgeSecurityPolicy = direct.LazyPtr(in.GetEdgeSecurityPolicy())
	return out
}
func NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *pb.EdgeCacheService {
	if in == nil {
		return nil
	}
	out := &pb.EdgeCacheService{}
	// MISSING: Name
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Routing = EdgeCacheService_Routing_ToProto(mapCtx, in.Routing)
	// MISSING: EdgeSSLCertificates
	out.DisableQuic = direct.ValueOf(in.DisableQuic)
	// MISSING: IPV4Addresses
	// MISSING: IPV6Addresses
	out.LogConfig = EdgeCacheService_LogConfig_ToProto(mapCtx, in.LogConfig)
	out.DisableHttp2 = direct.ValueOf(in.DisableHttp2)
	// MISSING: RequireTLS
	// (near miss): "RequireTLS" vs "RequireTls"
	out.EdgeSecurityPolicy = direct.ValueOf(in.EdgeSecurityPolicy)
	return out
}

// TODO: ServiceBinding mapper functions commented out temporarily
// ServiceBinding protos are not in the mockgcp-generated package yet
// These will need to be re-enabled when ServiceBinding is migrated to direct controller

/*
	func NetworkServicesServiceBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingObservedState {
		if in == nil {
			return nil
		}
		out := &krm.NetworkServicesServiceBindingObservedState{}
		// MISSING: Name
		out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
		out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
		// MISSING: ServiceID
		return out
	}

	func NetworkServicesServiceBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingObservedState) *pb.ServiceBinding {
		if in == nil {
			return nil
		}
		out := &pb.ServiceBinding{}
		// MISSING: Name
		out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
		out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
		// MISSING: ServiceID
		return out
	}

	func NetworkServicesServiceBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingSpec {
		if in == nil {
			return nil
		}
		out := &krm.NetworkServicesServiceBindingSpec{}
		// MISSING: Name
		out.Description = direct.LazyPtr(in.GetDescription())
		if in.GetService() != "" {
			out.ServiceRef = &krmservicedirectoryv1alpha1.ServiceDirectoryServiceRef{External: in.GetService()}
		}
		// MISSING: ServiceID
		out.Labels = in.Labels
		return out
	}

	func NetworkServicesServiceBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingSpec) *pb.ServiceBinding {
		if in == nil {
			return nil
		}
		out := &pb.ServiceBinding{}
		// MISSING: Name
		out.Description = direct.ValueOf(in.Description)
		if in.ServiceRef != nil {
			out.Service = in.ServiceRef.External
		}
		// MISSING: ServiceID
		out.Labels = in.Labels
		return out
	}
*/
func PathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.EdgeCacheService_PathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_PathMatcher{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, RouteRule_FromProto)
	return out
}
func PathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_PathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, RouteRule_ToProto)
	return out
}
func QueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.QueryParameterMatch) *krm.EdgeCacheService_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_QueryParameterMatch{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PresentMatch = direct.LazyPtr(in.GetPresentMatch())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	return out
}
func QueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_QueryParameterMatch) *pb.QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.QueryParameterMatch{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := QueryParameterMatch_PresentMatch_ToProto(mapCtx, in.PresentMatch); oneof != nil {
		out.MatchType = oneof
	}
	if oneof := QueryParameterMatch_ExactMatch_ToProto(mapCtx, in.ExactMatch); oneof != nil {
		out.MatchType = oneof
	}
	return out
}
func QueryParameterMatch_PresentMatch_ToProto(mapCtx *direct.MapContext, in *bool) *pb.QueryParameterMatch_PresentMatch {
	if in == nil {
		return nil
	}
	return &pb.QueryParameterMatch_PresentMatch{PresentMatch: *in}
}
func QueryParameterMatch_ExactMatch_ToProto(mapCtx *direct.MapContext, in *string) *pb.QueryParameterMatch_ExactMatch {
	if in == nil {
		return nil
	}
	return &pb.QueryParameterMatch_ExactMatch{ExactMatch: *in}
}
func RemoveHeader_FromProto(mapCtx *direct.MapContext, in *pb.RemoveHeader) *krm.EdgeCacheService_RemoveHeader {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RemoveHeader{}
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	return out
}
func RemoveHeader_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RemoveHeader) *pb.RemoveHeader {
	if in == nil {
		return nil
	}
	out := &pb.RemoveHeader{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	return out
}
func RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.RouteAction) *krm.EdgeCacheService_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteAction{}
	out.CdnPolicy = CDNPolicy_FromProto(mapCtx, in.GetCdnPolicy())
	out.UrlRewrite = URLRewrite_FromProto(mapCtx, in.GetUrlRewrite())
	out.CorsPolicy = CorsPolicy_FromProto(mapCtx, in.GetCorsPolicy())
	out.CompressionMode = direct.Enum_FromProto(mapCtx, in.GetCompressionMode())
	return out
}
func RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteAction) *pb.RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.RouteAction{}
	out.CdnPolicy = CDNPolicy_ToProto(mapCtx, in.CdnPolicy)
	out.UrlRewrite = URLRewrite_ToProto(mapCtx, in.UrlRewrite)
	out.CorsPolicy = CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.CompressionMode = direct.Enum_ToProto[pb.CompressionMode](mapCtx, in.CompressionMode)
	return out
}
func RouteMethods_FromProto(mapCtx *direct.MapContext, in *pb.RouteMethods) *krm.EdgeCacheService_RouteMethods {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteMethods{}
	out.AllowedMethods = in.AllowedMethods
	return out
}
func RouteMethods_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteMethods) *pb.RouteMethods {
	if in == nil {
		return nil
	}
	out := &pb.RouteMethods{}
	out.AllowedMethods = in.AllowedMethods
	return out
}
func RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.RouteRule) *krm.EdgeCacheService_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_RouteRule{}
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, MatchRule_FromProto)
	out.HeaderAction = HeaderAction_FromProto(mapCtx, in.GetHeaderAction())
	out.RouteAction = RouteAction_FromProto(mapCtx, in.GetRouteAction())
	out.UrlRedirect = URLRedirect_FromProto(mapCtx, in.GetUrlRedirect())
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.RouteMethods = RouteMethods_FromProto(mapCtx, in.GetRouteMethods())
	return out
}
func RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_RouteRule) *pb.RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.RouteRule{}
	out.Priority = direct.ValueOf(in.Priority)
	out.Description = direct.ValueOf(in.Description)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, MatchRule_ToProto)
	out.HeaderAction = HeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.RouteAction = RouteAction_ToProto(mapCtx, in.RouteAction)
	out.UrlRedirect = URLRedirect_ToProto(mapCtx, in.UrlRedirect)
	out.Origin = direct.ValueOf(in.Origin)
	out.RouteMethods = RouteMethods_ToProto(mapCtx, in.RouteMethods)
	return out
}
func Routing_FromProto(mapCtx *direct.MapContext, in *pb.Routing) *krm.EdgeCacheService_Routing {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_Routing{}
	out.HostRules = direct.Slice_FromProto(mapCtx, in.HostRules, HostRule_FromProto)
	out.PathMatchers = direct.Slice_FromProto(mapCtx, in.PathMatchers, PathMatcher_FromProto)
	return out
}
func Routing_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_Routing) *pb.Routing {
	if in == nil {
		return nil
	}
	out := &pb.Routing{}
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRules, HostRule_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatchers, PathMatcher_ToProto)
	return out
}
func SignedTokenOptions_FromProto(mapCtx *direct.MapContext, in *pb.SignedTokenOptions) *krm.EdgeCacheService_SignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_SignedTokenOptions{}
	out.TokenQueryParameter = direct.LazyPtr(in.GetTokenQueryParameter())
	out.AllowedSignatureAlgorithms = direct.EnumSlice_FromProto(mapCtx, in.AllowedSignatureAlgorithms)
	return out
}
func SignedTokenOptions_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_SignedTokenOptions) *pb.SignedTokenOptions {
	if in == nil {
		return nil
	}
	out := &pb.SignedTokenOptions{}
	out.TokenQueryParameter = direct.ValueOf(in.TokenQueryParameter)
	out.AllowedSignatureAlgorithms = direct.EnumSlice_ToProto[pb.SignatureAlgorithm](mapCtx, in.AllowedSignatureAlgorithms)
	return out
}
func URLRedirect_FromProto(mapCtx *direct.MapContext, in *pb.UrlRedirect) *krm.EdgeCacheService_UrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_UrlRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRedirect = direct.LazyPtr(in.GetPrefixRedirect())
	out.RedirectResponseCode = direct.Enum_FromProto(mapCtx, in.GetRedirectResponseCode())
	out.HttpsRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	return out
}
func URLRedirect_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_UrlRedirect) *pb.UrlRedirect {
	if in == nil {
		return nil
	}
	out := &pb.UrlRedirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRedirect = direct.ValueOf(in.PrefixRedirect)
	out.RedirectResponseCode = direct.Enum_ToProto[pb.RedirectResponseCode](mapCtx, in.RedirectResponseCode)
	out.HttpsRedirect = direct.ValueOf(in.HttpsRedirect)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}
func URLRewrite_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.EdgeCacheService_UrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.EdgeCacheService_UrlRewrite{}
	out.PathPrefixRewrite = direct.LazyPtr(in.GetPathPrefixRewrite())
	out.PathTemplateRewrite = direct.LazyPtr(in.GetPathTemplateRewrite())
	out.HostRewrite = direct.LazyPtr(in.GetHostRewrite())
	return out
}
func URLRewrite_ToProto(mapCtx *direct.MapContext, in *krm.EdgeCacheService_UrlRewrite) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.PathPrefixRewrite = direct.ValueOf(in.PathPrefixRewrite)
	out.PathTemplateRewrite = direct.ValueOf(in.PathTemplateRewrite)
	out.HostRewrite = direct.ValueOf(in.HostRewrite)
	return out
}
