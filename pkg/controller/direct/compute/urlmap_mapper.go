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

package compute

import (
	"strconv"
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapSpec{}
	out.DefaultRouteAction = HttpRouteAction_FromProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapSpec_DefaultService_FromProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HttpRedirectAction_FromProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HttpHeaderAction_FromProto(mapCtx, in.HeaderAction)
	out.HostRule = direct.Slice_FromProto(mapCtx, in.HostRules, HostRule_FromProto)
	// Location is handled by identity
	// out.Location = ...
	out.PathMatcher = direct.Slice_FromProto(mapCtx, in.PathMatchers, PathMatcher_FromProto)
	out.ResourceID = in.Name
	out.Test = direct.Slice_FromProto(mapCtx, in.Tests, UrlMapTest_FromProto)
	return out
}

func ComputeURLMapSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapSpec) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.DefaultRouteAction = HttpRouteAction_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapSpec_DefaultService_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HttpRedirectAction_ToProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HttpHeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRule, HostRule_ToProto)
	// Location is handled by identity
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatcher, PathMatcher_ToProto)
	out.Name = in.ResourceID
	out.Tests = direct.Slice_ToProto(mapCtx, in.Test, UrlMapTest_ToProto)
	return out
}

func ComputeURLMapSpec_DefaultService_FromProto(mapCtx *direct.MapContext, in *string) *krm.UrlmapDefaultService {
	if in == nil {
		return nil
	}
	// We can't easily distinguish between BackendService and BackendBucket just from the string
	// unless we parse the URL.
	// For now, let's try to detect based on the resource type in the URL.
	if strings.Contains(*in, "/backendBuckets/") {
		return &krm.UrlmapDefaultService{
			BackendBucketRef: &v1alpha1.ResourceRef{External: *in},
		}
	}
	return &krm.UrlmapDefaultService{
		BackendServiceRef: &v1alpha1.ResourceRef{External: *in},
	}
}

func ComputeURLMapSpec_DefaultService_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultService) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil {
		return direct.LazyPtr(in.BackendBucketRef.External)
	}
	if in.BackendServiceRef != nil {
		return direct.LazyPtr(in.BackendServiceRef.External)
	}
	return nil
}

func HttpRouteAction_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapDefaultRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDefaultRouteAction{}
	out.CorsPolicy = CorsPolicy_FromProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = FaultInjectionPolicy_FromProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = RequestMirrorPolicy_FromProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = RetryPolicy_FromProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_FromProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlRewrite_FromProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, WeightedBackendService_FromProto)
	return out
}

func HttpRouteAction_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultRouteAction) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = FaultInjectionPolicy_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = RequestMirrorPolicy_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlRewrite_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, WeightedBackendService_ToProto)
	return out
}

func HttpRedirectAction_FromProto(mapCtx *direct.MapContext, in *pb.HttpRedirectAction) *krm.UrlmapDefaultUrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDefaultUrlRedirect{}
	out.HostRedirect = in.HostRedirect
	out.HttpsRedirect = in.HttpsRedirect
	out.PathRedirect = in.PathRedirect
	out.PrefixRedirect = in.PrefixRedirect
	out.RedirectResponseCode = in.RedirectResponseCode
	out.StripQuery = direct.ValueOf(in.StripQuery)
	return out
}

func HttpRedirectAction_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultUrlRedirect) *pb.HttpRedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRedirectAction{}
	out.HostRedirect = in.HostRedirect
	out.HttpsRedirect = in.HttpsRedirect
	out.PathRedirect = in.PathRedirect
	out.PrefixRedirect = in.PrefixRedirect
	out.RedirectResponseCode = in.RedirectResponseCode
	out.StripQuery = direct.LazyPtr(in.StripQuery)
	return out
}

func HttpHeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderAction) *krm.UrlmapHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_FromProto(mapCtx, in.RequestHeadersToAdd, HttpHeaderOption_FromProto)
	out.RequestHeadersToRemove = in.RequestHeadersToRemove
	out.ResponseHeadersToAdd = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToAdd, HttpHeaderOption_Response_FromProto)
	out.ResponseHeadersToRemove = in.ResponseHeadersToRemove
	return out
}

func HttpHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHeaderAction) *pb.HttpHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_ToProto(mapCtx, in.RequestHeadersToAdd, HttpHeaderOption_ToProto)
	out.RequestHeadersToRemove = in.RequestHeadersToRemove
	out.ResponseHeadersToAdd = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToAdd, HttpHeaderOption_Response_ToProto)
	out.ResponseHeadersToRemove = in.ResponseHeadersToRemove
	return out
}

func HostRule_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.UrlmapHostRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = direct.ValueOf(in.PathMatcher)
	return out
}

func HostRule_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = direct.LazyPtr(in.PathMatcher)
	return out
}

func PathMatcher_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.UrlmapPathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPathMatcher{}
	out.DefaultRouteAction = HttpRouteAction_FromProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapSpec_DefaultService_FromProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HttpRedirectAction_FromProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HttpHeaderAction_FromProto(mapCtx, in.HeaderAction)
	out.Name = direct.ValueOf(in.Name)
	out.PathRule = direct.Slice_FromProto(mapCtx, in.PathRules, PathRule_FromProto)
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, RouteRule_FromProto)
	return out
}

func PathMatcher_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.DefaultRouteAction = HttpRouteAction_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapSpec_DefaultService_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HttpRedirectAction_ToProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HttpHeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.Name = direct.LazyPtr(in.Name)
	out.PathRules = direct.Slice_ToProto(mapCtx, in.PathRule, PathRule_ToProto)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, RouteRule_ToProto)
	return out
}

func PathRule_FromProto(mapCtx *direct.MapContext, in *pb.PathRule) *krm.UrlmapPathRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPathRule{}
	out.Paths = in.Paths
	out.RouteAction = RouteAction_FromProto(mapCtx, in.RouteAction)
	out.Service = ComputeURLMapSpec_Service_FromProto(mapCtx, in.Service)
	out.UrlRedirect = UrlRedirect_FromProto(mapCtx, in.UrlRedirect)
	return out
}

func PathRule_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPathRule) *pb.PathRule {
	if in == nil {
		return nil
	}
	out := &pb.PathRule{}
	out.Paths = in.Paths
	out.RouteAction = RouteAction_ToProto(mapCtx, in.RouteAction)
	out.Service = ComputeURLMapSpec_Service_ToProto(mapCtx, in.Service)
	out.UrlRedirect = UrlRedirect_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRule) *krm.UrlmapRouteRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteRules{}
	out.HeaderAction = HttpHeaderAction_FromProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, HttpRouteRuleMatch_FromProto)
	out.Priority = int64(direct.ValueOf(in.Priority))
	out.RouteAction = RouteAction_FromProto(mapCtx, in.RouteAction)
	out.Service = in.Service
	out.UrlRedirect = UrlRedirect_FromProto(mapCtx, in.UrlRedirect)
	return out
}

func RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteRules) *pb.HttpRouteRule {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteRule{}
	out.HeaderAction = HttpHeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, HttpRouteRuleMatch_ToProto)
	out.Priority = direct.LazyPtr(int32(in.Priority))
	out.RouteAction = RouteAction_ToProto(mapCtx, in.RouteAction)
	out.Service = in.Service
	out.UrlRedirect = UrlRedirect_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func UrlMapTest_FromProto(mapCtx *direct.MapContext, in *pb.UrlMapTest) *krm.UrlmapTest {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapTest{}
	out.Description = in.Description
	out.Host = direct.ValueOf(in.Host)
	out.Path = direct.ValueOf(in.Path)
	service := ComputeURLMapSpec_Service_FromProto(mapCtx, in.Service)
	if service != nil {
		out.Service = *service
	}
	return out
}

func UrlMapTest_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapTest) *pb.UrlMapTest {
	if in == nil {
		return nil
	}
	out := &pb.UrlMapTest{}
	out.Description = in.Description
	out.Host = direct.LazyPtr(in.Host)
	out.Path = direct.LazyPtr(in.Path)
	out.Service = ComputeURLMapSpec_Service_ToProto(mapCtx, &in.Service)
	return out
}

// Helper functions for nested structs

func CorsPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.UrlmapCorsPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapCorsPolicy{}
	out.AllowCredentials = in.AllowCredentials
	out.AllowHeaders = in.AllowHeaders
	out.AllowMethods = in.AllowMethods
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowOrigins = in.AllowOrigins
	out.Disabled = in.Disabled
	out.ExposeHeaders = in.ExposeHeaders
	out.MaxAge = int64Ptr(in.MaxAge)
	return out
}

func CorsPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapCorsPolicy) *pb.CorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CorsPolicy{}
	out.AllowCredentials = in.AllowCredentials
	out.AllowHeaders = in.AllowHeaders
	out.AllowMethods = in.AllowMethods
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowOrigins = in.AllowOrigins
	out.Disabled = in.Disabled
	out.ExposeHeaders = in.ExposeHeaders
	out.MaxAge = int32Ptr(in.MaxAge)
	return out
}

func FaultInjectionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultInjection) *krm.UrlmapFaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFaultInjectionPolicy{}
	out.Abort = HttpFaultAbort_FromProto(mapCtx, in.Abort)
	out.Delay = HttpFaultDelay_FromProto(mapCtx, in.Delay)
	return out
}

func FaultInjectionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFaultInjectionPolicy) *pb.HttpFaultInjection {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultInjection{}
	out.Abort = HttpFaultAbort_ToProto(mapCtx, in.Abort)
	out.Delay = HttpFaultDelay_ToProto(mapCtx, in.Delay)
	return out
}

func RequestMirrorPolicy_FromProto(mapCtx *direct.MapContext, in *pb.RequestMirrorPolicy) *krm.UrlmapRequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRequestMirrorPolicy{}
	out.BackendServiceRef = v1alpha1.ResourceRef{External: direct.ValueOf(in.BackendService)}
	return out
}

func RequestMirrorPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRequestMirrorPolicy) *pb.RequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &pb.RequestMirrorPolicy{}
	out.BackendService = direct.LazyPtr(in.BackendServiceRef.External)
	return out
}

func RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRetryPolicy) *krm.UrlmapRetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRetryPolicy{}
	out.NumRetries = int64(direct.ValueOf(in.NumRetries))
	out.PerTryTimeout = Duration_FromProto_PerTryTimeout(mapCtx, in.PerTryTimeout)
	out.RetryConditions = in.RetryConditions
	return out
}

func RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRetryPolicy) *pb.HttpRetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRetryPolicy{}
	out.NumRetries = direct.LazyPtr(uint32(in.NumRetries))
	out.PerTryTimeout = Duration_ToProto_PerTryTimeout(mapCtx, in.PerTryTimeout)
	out.RetryConditions = in.RetryConditions
	return out
}

func Duration_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapTimeout {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapTimeout{}
	out.Nanos = int64Ptr(in.Nanos)
	out.Seconds = strconv.FormatInt(direct.ValueOf(in.Seconds), 10)
	return out
}

func Duration_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Nanos = int32Ptr(in.Nanos)
	if in.Seconds != "" {
		s, err := strconv.ParseInt(in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds: %v", err)
		} else {
			out.Seconds = direct.LazyPtr(s)
		}
	}
	return out
}

func Duration_FromProto_PerTryTimeout(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapPerTryTimeout {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPerTryTimeout{}
	out.Nanos = int64Ptr(in.Nanos)
	out.Seconds = strconv.FormatInt(direct.ValueOf(in.Seconds), 10)
	return out
}

func Duration_ToProto_PerTryTimeout(mapCtx *direct.MapContext, in *krm.UrlmapPerTryTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Nanos = int32Ptr(in.Nanos)
	if in.Seconds != "" {
		s, err := strconv.ParseInt(in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds: %v", err)
		} else {
			out.Seconds = direct.LazyPtr(s)
		}
	}
	return out
}

func UrlRewrite_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.UrlmapUrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapUrlRewrite{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	out.PathTemplateRewrite = in.PathTemplateRewrite
	return out
}

func UrlRewrite_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapUrlRewrite) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	out.PathTemplateRewrite = in.PathTemplateRewrite
	return out
}

func WeightedBackendService_FromProto(mapCtx *direct.MapContext, in *pb.WeightedBackendService) *krm.UrlmapWeightedBackendServices {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapWeightedBackendServices{}
	out.BackendServiceRef = v1alpha1.ResourceRef{External: direct.ValueOf(in.BackendService)}
	out.HeaderAction = HttpHeaderAction_FromProto(mapCtx, in.HeaderAction)
	out.Weight = int64(direct.ValueOf(in.Weight))
	return out
}

func WeightedBackendService_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapWeightedBackendServices) *pb.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &pb.WeightedBackendService{}
	out.BackendService = direct.LazyPtr(in.BackendServiceRef.External)
	out.HeaderAction = HttpHeaderAction_ToProto(mapCtx, in.HeaderAction)
	out.Weight = direct.LazyPtr(uint32(in.Weight))
	return out
}

func HttpFaultAbort_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultAbort) *krm.UrlmapAbort {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapAbort{}
	out.HttpStatus = int64PtrFromUint32Ptr(in.HttpStatus)
	out.Percentage = in.Percentage
	return out
}

func HttpFaultAbort_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapAbort) *pb.HttpFaultAbort {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultAbort{}
	out.HttpStatus = uint32PtrFromInt64Ptr(in.HttpStatus)
	out.Percentage = in.Percentage
	return out
}

func HttpFaultDelay_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultDelay) *krm.UrlmapDelay {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDelay{}
	out.FixedDelay = Duration_FromProto_FixedDelay(mapCtx, in.FixedDelay)
	out.Percentage = in.Percentage
	return out
}

func HttpFaultDelay_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDelay) *pb.HttpFaultDelay {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultDelay{}
	out.FixedDelay = Duration_ToProto_FixedDelay(mapCtx, in.FixedDelay)
	out.Percentage = in.Percentage
	return out
}

func Duration_FromProto_FixedDelay(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapFixedDelay {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFixedDelay{}
	out.Nanos = int64Ptr(in.Nanos)
	out.Seconds = strconv.FormatInt(direct.ValueOf(in.Seconds), 10)
	return out
}

func Duration_ToProto_FixedDelay(mapCtx *direct.MapContext, in *krm.UrlmapFixedDelay) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Nanos = int32Ptr(in.Nanos)
	if in.Seconds != "" {
		s, err := strconv.ParseInt(in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds: %v", err)
		} else {
			out.Seconds = direct.LazyPtr(s)
		}
	}
	return out
}

func HttpHeaderOption_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapRequestHeadersToAdd {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRequestHeadersToAdd{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func HttpHeaderOption_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRequestHeadersToAdd) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.HeaderValue = direct.LazyPtr(in.HeaderValue)
	out.Replace = direct.LazyPtr(in.Replace)
	return out
}

func HttpHeaderOption_Response_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapResponseHeadersToAdd {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapResponseHeadersToAdd{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func HttpHeaderOption_Response_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapResponseHeadersToAdd) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.HeaderValue = direct.LazyPtr(in.HeaderValue)
	out.Replace = direct.LazyPtr(in.Replace)
	return out
}

func ComputeURLMapSpec_Service_FromProto(mapCtx *direct.MapContext, in *string) *krm.UrlmapService {
	if in == nil {
		return nil
	}
	// Similar logic to DefaultService
	if strings.Contains(*in, "/backendBuckets/") {
		return &krm.UrlmapService{
			BackendBucketRef: &v1alpha1.ResourceRef{External: *in},
		}
	}
	return &krm.UrlmapService{
		BackendServiceRef: &v1alpha1.ResourceRef{External: *in},
	}
}

func ComputeURLMapSpec_Service_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapService) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil {
		return direct.LazyPtr(in.BackendBucketRef.External)
	}
	if in.BackendServiceRef != nil {
		return direct.LazyPtr(in.BackendServiceRef.External)
	}
	return nil
}

func RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteAction{}
	out.CorsPolicy = CorsPolicy_FromProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = FaultInjectionPolicy_FromProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = RequestMirrorPolicy_FromProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = RetryPolicy_FromProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_FromProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlRewrite_FromProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, WeightedBackendService_FromProto)
	return out
}

func RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteAction) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = CorsPolicy_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = FaultInjectionPolicy_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = RequestMirrorPolicy_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlRewrite_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, WeightedBackendService_ToProto)
	return out
}

func UrlRedirect_FromProto(mapCtx *direct.MapContext, in *pb.HttpRedirectAction) *krm.UrlmapUrlRedirect {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapUrlRedirect{}
	out.HostRedirect = in.HostRedirect
	out.HttpsRedirect = in.HttpsRedirect
	out.PathRedirect = in.PathRedirect
	out.PrefixRedirect = in.PrefixRedirect
	out.RedirectResponseCode = in.RedirectResponseCode
	out.StripQuery = in.StripQuery
	return out
}

func UrlRedirect_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapUrlRedirect) *pb.HttpRedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRedirectAction{}
	out.HostRedirect = in.HostRedirect
	out.HttpsRedirect = in.HttpsRedirect
	out.PathRedirect = in.PathRedirect
	out.PrefixRedirect = in.PrefixRedirect
	out.RedirectResponseCode = in.RedirectResponseCode
	out.StripQuery = in.StripQuery
	return out
}

func HttpRouteRuleMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRuleMatch) *krm.UrlmapMatchRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapMatchRules{}
	out.FullPathMatch = in.FullPathMatch
	out.HeaderMatches = direct.Slice_FromProto(mapCtx, in.HeaderMatches, HttpHeaderMatch_FromProto)
	out.IgnoreCase = in.IgnoreCase
	out.MetadataFilters = direct.Slice_FromProto(mapCtx, in.MetadataFilters, MetadataFilter_FromProto)
	out.PathTemplateMatch = in.PathTemplateMatch
	out.PrefixMatch = in.PrefixMatch
	out.QueryParameterMatches = direct.Slice_FromProto(mapCtx, in.QueryParameterMatches, HttpQueryParameterMatch_FromProto)
	out.RegexMatch = in.RegexMatch
	return out
}

func HttpRouteRuleMatch_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapMatchRules) *pb.HttpRouteRuleMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteRuleMatch{}
	out.FullPathMatch = in.FullPathMatch
	out.HeaderMatches = direct.Slice_ToProto(mapCtx, in.HeaderMatches, HttpHeaderMatch_ToProto)
	out.IgnoreCase = in.IgnoreCase
	out.MetadataFilters = direct.Slice_ToProto(mapCtx, in.MetadataFilters, MetadataFilter_ToProto)
	out.PathTemplateMatch = in.PathTemplateMatch
	out.PrefixMatch = in.PrefixMatch
	out.QueryParameterMatches = direct.Slice_ToProto(mapCtx, in.QueryParameterMatches, HttpQueryParameterMatch_ToProto)
	out.RegexMatch = in.RegexMatch
	return out
}

func HttpHeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderMatch) *krm.UrlmapHeaderMatches {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHeaderMatches{}
	out.ExactMatch = in.ExactMatch
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.InvertMatch = in.InvertMatch
	out.PrefixMatch = in.PrefixMatch
	out.PresentMatch = in.PresentMatch
	out.RangeMatch = Int64RangeMatch_FromProto(mapCtx, in.RangeMatch)
	out.RegexMatch = in.RegexMatch
	out.SuffixMatch = in.SuffixMatch
	return out
}

func HttpHeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHeaderMatches) *pb.HttpHeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderMatch{}
	out.ExactMatch = in.ExactMatch
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.InvertMatch = in.InvertMatch
	out.PrefixMatch = in.PrefixMatch
	out.PresentMatch = in.PresentMatch
	out.RangeMatch = Int64RangeMatch_ToProto(mapCtx, in.RangeMatch)
	out.RegexMatch = in.RegexMatch
	out.SuffixMatch = in.SuffixMatch
	return out
}

func Int64RangeMatch_FromProto(mapCtx *direct.MapContext, in *pb.Int64RangeMatch) *krm.UrlmapRangeMatch {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRangeMatch{}
	out.RangeEnd = int64(direct.ValueOf(in.RangeEnd))
	out.RangeStart = int64(direct.ValueOf(in.RangeStart))
	return out
}

func Int64RangeMatch_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRangeMatch) *pb.Int64RangeMatch {
	if in == nil {
		return nil
	}
	out := &pb.Int64RangeMatch{}
	out.RangeEnd = direct.LazyPtr(int64(in.RangeEnd))
	out.RangeStart = direct.LazyPtr(int64(in.RangeStart))
	return out
}

func MetadataFilter_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.UrlmapMetadataFilters {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapMetadataFilters{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_FromProto)
	out.FilterMatchCriteria = direct.ValueOf(in.FilterMatchCriteria)
	return out
}

func MetadataFilter_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapMetadataFilters) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, MetadataFilterLabelMatch_ToProto)
	out.FilterMatchCriteria = direct.LazyPtr(in.FilterMatchCriteria)
	return out
}

func MetadataFilterLabelMatch_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.UrlmapFilterLabels {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFilterLabels{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}

func MetadataFilterLabelMatch_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFilterLabels) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = direct.LazyPtr(in.Name)
	out.Value = direct.LazyPtr(in.Value)
	return out
}

func HttpQueryParameterMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpQueryParameterMatch) *krm.UrlmapQueryParameterMatches {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapQueryParameterMatches{}
	out.ExactMatch = in.ExactMatch
	out.Name = direct.ValueOf(in.Name)
	out.PresentMatch = in.PresentMatch
	out.RegexMatch = in.RegexMatch
	return out
}

func HttpQueryParameterMatch_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapQueryParameterMatches) *pb.HttpQueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpQueryParameterMatch{}
	out.ExactMatch = in.ExactMatch
	out.Name = direct.LazyPtr(in.Name)
	out.PresentMatch = in.PresentMatch
	out.RegexMatch = in.RegexMatch
	return out
}

func int64Ptr(v *int32) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func int32Ptr(v *int64) *int32 {
	if v == nil {
		return nil
	}
	val := int32(*v)
	return &val
}

func int64PtrFromUint32Ptr(v *uint32) *int64 {
	if v == nil {
		return nil
	}
	val := int64(*v)
	return &val
}

func uint32PtrFromInt64Ptr(v *int64) *uint32 {
	if v == nil {
		return nil
	}
	val := uint32(*v)
	return &val
}
