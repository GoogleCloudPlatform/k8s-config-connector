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

package compute

import (
	"fmt"
	"strconv"
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapSpec{}
	out.DefaultCustomErrorResponsePolicy = UrlmapDefaultCustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = UrlmapDefaultRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = UrlmapDefaultService_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = UrlmapHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.HostRule = direct.Slice_FromProto(mapCtx, in.HostRules, UrlmapHostRule_v1beta1_FromProto)
	out.Location = direct.ValueOf(in.Region)
	if out.Location == "" {
		out.Location = "global"
	}
	out.PathMatcher = direct.Slice_FromProto(mapCtx, in.PathMatchers, UrlmapPathMatcher_v1beta1_FromProto)
	out.ResourceID = in.Name
	out.Test = direct.Slice_FromProto(mapCtx, in.Tests, UrlmapTest_v1beta1_FromProto)
	return out
}

func ComputeURLMapSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapSpec) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.DefaultCustomErrorResponsePolicy = UrlmapDefaultCustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = UrlmapDefaultRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = UrlmapDefaultService_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_ToProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = UrlmapHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRule, UrlmapHostRule_v1beta1_ToProto)
	if in.Location != "global" && in.Location != "" {
		out.Region = direct.LazyPtr(in.Location)
	}
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatcher, UrlmapPathMatcher_v1beta1_ToProto)
	out.Name = in.ResourceID
	out.Tests = direct.Slice_ToProto(mapCtx, in.Test, UrlmapTest_v1beta1_ToProto)
	return out
}

func ComputeURLMapStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		out.MapId = direct.LazyPtr(int64(*in.Id))
	}
	out.SelfLink = in.SelfLink
	return out
}

func ComputeURLMapStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapStatus) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.MapId != nil {
		out.Id = direct.LazyPtr(uint64(*in.MapId))
	}
	out.SelfLink = in.SelfLink
	return out
}

func UrlmapDefaultService_v1beta1_FromProto(mapCtx *direct.MapContext, in *string) *krm.UrlmapDefaultService {
	if in == nil || *in == "" {
		return nil
	}
	val := *in
	out := &krm.UrlmapDefaultService{}
	if strings.Contains(val, "backendBuckets") {
		out.BackendBucketRef = &krm.UrlmapResourceRef{External: val}
	} else {
		out.BackendServiceRef = &krm.UrlmapResourceRef{External: val}
	}
	return out
}

func UrlmapDefaultService_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultService) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil && in.BackendBucketRef.External != "" {
		return direct.LazyPtr(in.BackendBucketRef.External)
	}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		return direct.LazyPtr(in.BackendServiceRef.External)
	}
	return nil
}

func UrlmapService_v1beta1_FromProto(mapCtx *direct.MapContext, in *string) *krm.UrlmapService {
	if in == nil || *in == "" {
		return nil
	}
	val := *in
	out := &krm.UrlmapService{}
	if strings.Contains(val, "backendBuckets") {
		out.BackendBucketRef = &krm.UrlmapResourceRef{External: val}
	} else {
		out.BackendServiceRef = &krm.UrlmapResourceRef{External: val}
	}
	return out
}

func UrlmapService_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapService) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil && in.BackendBucketRef.External != "" {
		return direct.LazyPtr(in.BackendBucketRef.External)
	}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		return direct.LazyPtr(in.BackendServiceRef.External)
	}
	return nil
}

func UrlmapCorsPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.UrlmapCorsPolicy {
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
	if in.MaxAge != nil {
		out.MaxAge = direct.LazyPtr(int64(*in.MaxAge))
	}
	return out
}

func UrlmapCorsPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapCorsPolicy) *pb.CorsPolicy {
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
	if in.MaxAge != nil {
		out.MaxAge = direct.LazyPtr(int32(*in.MaxAge))
	}
	return out
}

func UrlmapCorsPolicyRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CorsPolicy) *krm.UrlmapCorsPolicyRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapCorsPolicyRequired{}
	out.AllowCredentials = in.AllowCredentials
	out.AllowHeaders = in.AllowHeaders
	out.AllowMethods = in.AllowMethods
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowOrigins = in.AllowOrigins
	out.Disabled = direct.ValueOf(in.Disabled)
	out.ExposeHeaders = in.ExposeHeaders
	if in.MaxAge != nil {
		out.MaxAge = direct.LazyPtr(int64(*in.MaxAge))
	}
	return out
}

func UrlmapCorsPolicyRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapCorsPolicyRequired) *pb.CorsPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CorsPolicy{}
	out.AllowCredentials = in.AllowCredentials
	out.AllowHeaders = in.AllowHeaders
	out.AllowMethods = in.AllowMethods
	out.AllowOriginRegexes = in.AllowOriginRegexes
	out.AllowOrigins = in.AllowOrigins
	out.Disabled = direct.LazyPtr(in.Disabled)
	out.ExposeHeaders = in.ExposeHeaders
	if in.MaxAge != nil {
		out.MaxAge = direct.LazyPtr(int32(*in.MaxAge))
	}
	return out
}

func UrlmapTimeout_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapTimeout {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapTimeout{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = direct.LazyPtr(fmt.Sprintf("%d", *in.Seconds))
	}
	return out
}

func UrlmapTimeout_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	if in.Seconds != nil {
		val, err := strconv.ParseInt(*in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds %q: %v", *in.Seconds, err)
		} else {
			out.Seconds = direct.LazyPtr(val)
		}
	}
	return out
}

func UrlmapTimeoutRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapTimeoutRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapTimeoutRequired{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = fmt.Sprintf("%d", *in.Seconds)
	}
	return out
}

func UrlmapTimeoutRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapTimeoutRequired) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	val, err := strconv.ParseInt(in.Seconds, 10, 64)
	if err != nil {
		mapCtx.Errorf("invalid seconds %q: %v", in.Seconds, err)
	} else {
		out.Seconds = direct.LazyPtr(val)
	}
	return out
}

func UrlmapPerTryTimeout_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapPerTryTimeout {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPerTryTimeout{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = direct.LazyPtr(fmt.Sprintf("%d", *in.Seconds))
	}
	return out
}

func UrlmapPerTryTimeout_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPerTryTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	if in.Seconds != nil {
		val, err := strconv.ParseInt(*in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds %q: %v", *in.Seconds, err)
		} else {
			out.Seconds = direct.LazyPtr(val)
		}
	}
	return out
}

func UrlmapPerTryTimeoutRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapPerTryTimeoutRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPerTryTimeoutRequired{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = fmt.Sprintf("%d", *in.Seconds)
	}
	return out
}

func UrlmapPerTryTimeoutRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPerTryTimeoutRequired) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	val, err := strconv.ParseInt(in.Seconds, 10, 64)
	if err != nil {
		mapCtx.Errorf("invalid seconds %q: %v", in.Seconds, err)
	} else {
		out.Seconds = direct.LazyPtr(val)
	}
	return out
}

func UrlmapFixedDelay_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapFixedDelay {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFixedDelay{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = direct.LazyPtr(fmt.Sprintf("%d", *in.Seconds))
	}
	return out
}

func UrlmapFixedDelay_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFixedDelay) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	if in.Seconds != nil {
		val, err := strconv.ParseInt(*in.Seconds, 10, 64)
		if err != nil {
			mapCtx.Errorf("invalid seconds %q: %v", *in.Seconds, err)
		} else {
			out.Seconds = direct.LazyPtr(val)
		}
	}
	return out
}

func UrlmapFixedDelayRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.UrlmapFixedDelayRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFixedDelayRequired{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int64(*in.Nanos))
	}
	if in.Seconds != nil {
		out.Seconds = fmt.Sprintf("%d", *in.Seconds)
	}
	return out
}

func UrlmapFixedDelayRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFixedDelayRequired) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		out.Nanos = direct.LazyPtr(int32(*in.Nanos))
	}
	val, err := strconv.ParseInt(in.Seconds, 10, 64)
	if err != nil {
		mapCtx.Errorf("invalid seconds %q: %v", in.Seconds, err)
	} else {
		out.Seconds = direct.LazyPtr(val)
	}
	return out
}

func UrlmapDefaultUrlRedirect_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRedirectAction) *krm.UrlmapDefaultUrlRedirect {
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

func UrlmapDefaultUrlRedirect_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultUrlRedirect) *pb.HttpRedirectAction {
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

func UrlmapUrlRedirect_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRedirectAction) *krm.UrlmapUrlRedirect {
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

func UrlmapUrlRedirect_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapUrlRedirect) *pb.HttpRedirectAction {
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

func UrlmapRouteActionRouteRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapRouteActionRouteRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteActionRouteRules{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_FromProto(mapCtx, in.GetCorsPolicy())
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicyRouteRules_v1beta1_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_FromProto(mapCtx, in.GetRequestMirrorPolicy())
	out.RetryPolicy = UrlmapRetryPolicyRequired_v1beta1_FromProto(mapCtx, in.GetRetryPolicy())
	out.Timeout = UrlmapTimeoutRequired_v1beta1_FromProto(mapCtx, in.GetTimeout())
	out.UrlRewrite = UrlmapUrlRewriteWithPathTemplate_v1beta1_FromProto(mapCtx, in.GetUrlRewrite())
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServicesRequired_v1beta1_FromProto)
	return out
}

func UrlmapRouteActionRouteRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteActionRouteRules) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicyRouteRules_v1beta1_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = UrlmapRetryPolicyRequired_v1beta1_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = UrlmapTimeoutRequired_v1beta1_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlmapUrlRewriteWithPathTemplate_v1beta1_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServicesRequired_v1beta1_ToProto)
	return out
}

func UrlmapRouteActionPathMatcherDefaultRouteAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapRouteActionPathMatcherDefaultRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteActionPathMatcherDefaultRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_FromProto(mapCtx, in.GetCorsPolicy())
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicy_v1beta1_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_FromProto(mapCtx, in.GetRequestMirrorPolicy())
	out.RetryPolicy = UrlmapRetryPolicy_v1beta1_FromProto(mapCtx, in.GetRetryPolicy())
	out.Timeout = UrlmapTimeout_v1beta1_FromProto(mapCtx, in.GetTimeout())
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_FromProto(mapCtx, in.GetUrlRewrite())
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServices_v1beta1_FromProto)
	return out
}

func UrlmapRouteActionPathMatcherDefaultRouteAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteActionPathMatcherDefaultRouteAction) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicy_v1beta1_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = UrlmapRetryPolicy_v1beta1_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = UrlmapTimeout_v1beta1_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServices_v1beta1_ToProto)
	return out
}

func UrlmapRouteActionPathRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapRouteActionPathRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteActionPathRule{}
	out.CorsPolicy = UrlmapCorsPolicyRequired_v1beta1_FromProto(mapCtx, in.GetCorsPolicy())
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicyRequired_v1beta1_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_FromProto(mapCtx, in.GetRequestMirrorPolicy())
	out.RetryPolicy = UrlmapRetryPolicyOptional_v1beta1_FromProto(mapCtx, in.GetRetryPolicy())
	out.Timeout = UrlmapTimeoutRequired_v1beta1_FromProto(mapCtx, in.GetTimeout())
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_FromProto(mapCtx, in.GetUrlRewrite())
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServicesRequired_v1beta1_FromProto)
	return out
}

func UrlmapRouteActionPathRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteActionPathRule) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicyRequired_v1beta1_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicyRequired_v1beta1_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicyRequired_v1beta1_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = UrlmapRetryPolicyOptional_v1beta1_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = UrlmapTimeoutRequired_v1beta1_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServicesRequired_v1beta1_ToProto)
	return out
}

func UrlmapDefaultRouteAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.UrlmapDefaultRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDefaultRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_FromProto(mapCtx, in.GetCorsPolicy())
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicy_v1beta1_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicy_v1beta1_FromProto(mapCtx, in.GetRequestMirrorPolicy())
	out.RetryPolicy = UrlmapRetryPolicy_v1beta1_FromProto(mapCtx, in.GetRetryPolicy())
	out.Timeout = UrlmapTimeout_v1beta1_FromProto(mapCtx, in.GetTimeout())
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_FromProto(mapCtx, in.GetUrlRewrite())
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServices_v1beta1_FromProto)
	return out
}

func UrlmapDefaultRouteAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultRouteAction) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = UrlmapCorsPolicy_v1beta1_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = UrlmapFaultInjectionPolicy_v1beta1_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.RequestMirrorPolicy = UrlmapRequestMirrorPolicy_v1beta1_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = UrlmapRetryPolicy_v1beta1_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = UrlmapTimeout_v1beta1_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = UrlmapUrlRewrite_v1beta1_ToProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, UrlmapWeightedBackendServices_v1beta1_ToProto)
	return out
}

func UrlmapRequestMirrorPolicyRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RequestMirrorPolicy) *krm.UrlmapRequestMirrorPolicyRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRequestMirrorPolicyRequired{}
	if in.GetBackendService() != "" {
		out.BackendServiceRef = krm.UrlmapResourceRef{External: in.GetBackendService()}
	}
	return out
}

func UrlmapRequestMirrorPolicyRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRequestMirrorPolicyRequired) *pb.RequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &pb.RequestMirrorPolicy{}
	if in.BackendServiceRef.External != "" {
		out.BackendService = &in.BackendServiceRef.External
	}
	return out
}

func UrlmapRetryPolicyRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRetryPolicy) *krm.UrlmapRetryPolicyRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRetryPolicyRequired{}
	if in.NumRetries != nil {
		out.NumRetries = int64(*in.NumRetries)
	}
	out.PerTryTimeout = UrlmapPerTryTimeoutRequired_v1beta1_FromProto(mapCtx, in.GetPerTryTimeout())
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapRetryPolicyRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRetryPolicyRequired) *pb.HttpRetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRetryPolicy{}
	out.NumRetries = direct.LazyPtr(uint32(in.NumRetries))
	out.PerTryTimeout = UrlmapPerTryTimeoutRequired_v1beta1_ToProto(mapCtx, in.PerTryTimeout)
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapRetryPolicyOptional_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRetryPolicy) *krm.UrlmapRetryPolicyOptional {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRetryPolicyOptional{}
	if in.NumRetries != nil {
		out.NumRetries = direct.LazyPtr(int64(*in.NumRetries))
	}
	out.PerTryTimeout = UrlmapPerTryTimeoutRequired_v1beta1_FromProto(mapCtx, in.GetPerTryTimeout())
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapRetryPolicyOptional_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRetryPolicyOptional) *pb.HttpRetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRetryPolicy{}
	if in.NumRetries != nil {
		out.NumRetries = direct.LazyPtr(uint32(*in.NumRetries))
	}
	out.PerTryTimeout = UrlmapPerTryTimeoutRequired_v1beta1_ToProto(mapCtx, in.PerTryTimeout)
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapWeightedBackendServices_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.WeightedBackendService) *krm.UrlmapWeightedBackendServices {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapWeightedBackendServices{}
	if in.GetBackendService() != "" {
		out.BackendServiceRef = &krm.UrlmapResourceRef{External: in.GetBackendService()}
	}
	out.HeaderAction = UrlmapHeaderActionOptional_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	if in.Weight != nil {
		out.Weight = direct.LazyPtr(int64(*in.Weight))
	}
	return out
}

func UrlmapWeightedBackendServices_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapWeightedBackendServices) *pb.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &pb.WeightedBackendService{}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		out.BackendService = &in.BackendServiceRef.External
	}
	out.HeaderAction = UrlmapHeaderActionOptional_v1beta1_ToProto(mapCtx, in.HeaderAction)
	if in.Weight != nil {
		out.Weight = direct.LazyPtr(uint32(*in.Weight))
	}
	return out
}

func UrlmapWeightedBackendServicesRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.WeightedBackendService) *krm.UrlmapWeightedBackendServicesRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapWeightedBackendServicesRequired{}
	if in.GetBackendService() != "" {
		out.BackendServiceRef = krm.UrlmapResourceRef{External: in.GetBackendService()}
	}
	out.HeaderAction = UrlmapHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	if in.Weight != nil {
		out.Weight = int64(*in.Weight)
	}
	return out
}

func UrlmapWeightedBackendServicesRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapWeightedBackendServicesRequired) *pb.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &pb.WeightedBackendService{}
	if in.BackendServiceRef.External != "" {
		out.BackendService = &in.BackendServiceRef.External
	}
	out.HeaderAction = UrlmapHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Weight = direct.LazyPtr(uint32(in.Weight))
	return out
}

func UrlmapHeaderActionOptional_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderAction) *krm.UrlmapHeaderActionOptional {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHeaderActionOptional{}
	out.RequestHeadersToAdd = direct.Slice_FromProto(mapCtx, in.RequestHeadersToAdd, UrlmapRequestHeadersToAddOptional_v1beta1_FromProto)
	out.RequestHeadersToRemove = in.RequestHeadersToRemove
	out.ResponseHeadersToAdd = direct.Slice_FromProto(mapCtx, in.ResponseHeadersToAdd, UrlmapResponseHeadersToAddOptional_v1beta1_FromProto)
	out.ResponseHeadersToRemove = in.ResponseHeadersToRemove
	return out
}

func UrlmapHeaderActionOptional_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHeaderActionOptional) *pb.HttpHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderAction{}
	out.RequestHeadersToAdd = direct.Slice_ToProto(mapCtx, in.RequestHeadersToAdd, UrlmapRequestHeadersToAddOptional_v1beta1_ToProto)
	out.RequestHeadersToRemove = in.RequestHeadersToRemove
	out.ResponseHeadersToAdd = direct.Slice_ToProto(mapCtx, in.ResponseHeadersToAdd, UrlmapResponseHeadersToAddOptional_v1beta1_ToProto)
	out.ResponseHeadersToRemove = in.ResponseHeadersToRemove
	return out
}

func UrlmapRequestHeadersToAddOptional_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapRequestHeadersToAddOptional {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRequestHeadersToAddOptional{}
	if in.HeaderName != nil {
		out.HeaderName = in.HeaderName
	}
	if in.HeaderValue != nil {
		out.HeaderValue = in.HeaderValue
	}
	if in.Replace != nil {
		out.Replace = in.Replace
	}
	return out
}

func UrlmapRequestHeadersToAddOptional_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRequestHeadersToAddOptional) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = in.Replace
	return out
}

func UrlmapResponseHeadersToAddOptional_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapResponseHeadersToAddOptional {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapResponseHeadersToAddOptional{}
	if in.HeaderName != nil {
		out.HeaderName = in.HeaderName
	}
	if in.HeaderValue != nil {
		out.HeaderValue = in.HeaderValue
	}
	if in.Replace != nil {
		out.Replace = in.Replace
	}
	return out
}

func UrlmapResponseHeadersToAddOptional_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapResponseHeadersToAddOptional) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = in.HeaderName
	out.HeaderValue = in.HeaderValue
	out.Replace = in.Replace
	return out
}

func UrlmapFaultInjectionPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultInjection) *krm.UrlmapFaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFaultInjectionPolicy{}
	out.Abort = UrlmapAbort_v1beta1_FromProto(mapCtx, in.GetAbort())
	out.Delay = UrlmapDelay_v1beta1_FromProto(mapCtx, in.GetDelay())
	return out
}

func UrlmapFaultInjectionPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFaultInjectionPolicy) *pb.HttpFaultInjection {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultInjection{}
	out.Abort = UrlmapAbort_v1beta1_ToProto(mapCtx, in.Abort)
	out.Delay = UrlmapDelay_v1beta1_ToProto(mapCtx, in.Delay)
	return out
}

func UrlmapFaultInjectionPolicyRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultInjection) *krm.UrlmapFaultInjectionPolicyRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFaultInjectionPolicyRequired{}
	out.Abort = UrlmapAbortRequired_v1beta1_FromProto(mapCtx, in.GetAbort())
	out.Delay = UrlmapDelayRequired_v1beta1_FromProto(mapCtx, in.GetDelay())
	return out
}

func UrlmapFaultInjectionPolicyRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFaultInjectionPolicyRequired) *pb.HttpFaultInjection {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultInjection{}
	out.Abort = UrlmapAbortRequired_v1beta1_ToProto(mapCtx, in.Abort)
	out.Delay = UrlmapDelayRequired_v1beta1_ToProto(mapCtx, in.Delay)
	return out
}

func UrlmapFaultInjectionPolicyRouteRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultInjection) *krm.UrlmapFaultInjectionPolicyRouteRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFaultInjectionPolicyRouteRules{}
	out.Abort = UrlmapAbort_v1beta1_FromProto(mapCtx, in.GetAbort())
	out.Delay = UrlmapDelayRouteRules_v1beta1_FromProto(mapCtx, in.GetDelay())
	return out
}

func UrlmapFaultInjectionPolicyRouteRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFaultInjectionPolicyRouteRules) *pb.HttpFaultInjection {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultInjection{}
	out.Abort = UrlmapAbort_v1beta1_ToProto(mapCtx, in.Abort)
	out.Delay = UrlmapDelayRouteRules_v1beta1_ToProto(mapCtx, in.Delay)
	return out
}

func UrlmapAbort_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultAbort) *krm.UrlmapAbort {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapAbort{}
	if in.HttpStatus != nil {
		out.HttpStatus = direct.LazyPtr(int64(*in.HttpStatus))
	}
	out.Percentage = in.Percentage
	return out
}

func UrlmapAbort_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapAbort) *pb.HttpFaultAbort {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultAbort{}
	if in.HttpStatus != nil {
		out.HttpStatus = direct.LazyPtr(uint32(*in.HttpStatus))
	}
	out.Percentage = in.Percentage
	return out
}

func UrlmapAbortRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultAbort) *krm.UrlmapAbortRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapAbortRequired{}
	if in.HttpStatus != nil {
		out.HttpStatus = int64(*in.HttpStatus)
	}
	if in.Percentage != nil {
		out.Percentage = *in.Percentage
	}
	return out
}

func UrlmapAbortRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapAbortRequired) *pb.HttpFaultAbort {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultAbort{}
	out.HttpStatus = direct.LazyPtr(uint32(in.HttpStatus))
	out.Percentage = direct.LazyPtr(in.Percentage)
	return out
}

func UrlmapDelay_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultDelay) *krm.UrlmapDelay {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDelay{}
	out.FixedDelay = UrlmapFixedDelay_v1beta1_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = in.Percentage
	return out
}

func UrlmapDelay_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDelay) *pb.HttpFaultDelay {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultDelay{}
	out.FixedDelay = UrlmapFixedDelay_v1beta1_ToProto(mapCtx, in.FixedDelay)
	out.Percentage = in.Percentage
	return out
}

func UrlmapDelayRequired_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultDelay) *krm.UrlmapDelayRequired {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDelayRequired{}
	out.FixedDelay = direct.ValueOf(UrlmapFixedDelayRequired_v1beta1_FromProto(mapCtx, in.GetFixedDelay()))
	if in.Percentage != nil {
		out.Percentage = *in.Percentage
	}
	return out
}

func UrlmapDelayRequired_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDelayRequired) *pb.HttpFaultDelay {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultDelay{}
	out.FixedDelay = UrlmapFixedDelayRequired_v1beta1_ToProto(mapCtx, &in.FixedDelay)
	out.Percentage = direct.LazyPtr(in.Percentage)
	return out
}

func UrlmapDelayRouteRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpFaultDelay) *krm.UrlmapDelayRouteRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDelayRouteRules{}
	out.FixedDelay = UrlmapFixedDelayRequired_v1beta1_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = in.Percentage
	return out
}

func UrlmapDelayRouteRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDelayRouteRules) *pb.HttpFaultDelay {
	if in == nil {
		return nil
	}
	out := &pb.HttpFaultDelay{}
	out.FixedDelay = UrlmapFixedDelayRequired_v1beta1_ToProto(mapCtx, in.FixedDelay)
	out.Percentage = in.Percentage
	return out
}

func UrlmapUrlRewriteWithPathTemplate_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.UrlmapUrlRewriteWithPathTemplate {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapUrlRewriteWithPathTemplate{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	out.PathTemplateRewrite = in.PathTemplateRewrite
	return out
}

func UrlmapUrlRewriteWithPathTemplate_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapUrlRewriteWithPathTemplate) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	out.PathTemplateRewrite = in.PathTemplateRewrite
	return out
}

func UrlmapUrlRewrite_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlRewrite) *krm.UrlmapUrlRewrite {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapUrlRewrite{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	return out
}

func UrlmapUrlRewrite_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapUrlRewrite) *pb.UrlRewrite {
	if in == nil {
		return nil
	}
	out := &pb.UrlRewrite{}
	out.HostRewrite = in.HostRewrite
	out.PathPrefixRewrite = in.PathPrefixRewrite
	return out
}

func UrlmapFilterLabels_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilterLabelMatch) *krm.UrlmapFilterLabels {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapFilterLabels{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}

func UrlmapFilterLabels_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapFilterLabels) *pb.MetadataFilterLabelMatch {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilterLabelMatch{}
	out.Name = direct.LazyPtr(in.Name)
	out.Value = direct.LazyPtr(in.Value)
	return out
}

func UrlmapMetadataFilters_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataFilter) *krm.UrlmapMetadataFilters {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapMetadataFilters{}
	out.FilterLabels = direct.Slice_FromProto(mapCtx, in.FilterLabels, UrlmapFilterLabels_v1beta1_FromProto)
	out.FilterMatchCriteria = direct.ValueOf(in.FilterMatchCriteria)
	return out
}

func UrlmapMetadataFilters_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapMetadataFilters) *pb.MetadataFilter {
	if in == nil {
		return nil
	}
	out := &pb.MetadataFilter{}
	out.FilterLabels = direct.Slice_ToProto(mapCtx, in.FilterLabels, UrlmapFilterLabels_v1beta1_ToProto)
	out.FilterMatchCriteria = direct.LazyPtr(in.FilterMatchCriteria)
	return out
}

func UrlmapHostRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.UrlmapHostRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = direct.ValueOf(in.PathMatcher)
	return out
}

func UrlmapHostRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = direct.LazyPtr(in.PathMatcher)
	return out
}

func UrlmapHeaderMatches_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderMatch) *krm.UrlmapHeaderMatches {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapHeaderMatches{}
	out.ExactMatch = in.ExactMatch
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.InvertMatch = in.InvertMatch
	out.PrefixMatch = in.PrefixMatch
	out.PresentMatch = in.PresentMatch
	out.RangeMatch = UrlmapRangeMatch_v1beta1_FromProto(mapCtx, in.GetRangeMatch())
	out.RegexMatch = in.RegexMatch
	out.SuffixMatch = in.SuffixMatch
	return out
}

func UrlmapHeaderMatches_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapHeaderMatches) *pb.HttpHeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderMatch{}
	out.ExactMatch = in.ExactMatch
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.InvertMatch = in.InvertMatch
	out.PrefixMatch = in.PrefixMatch
	out.PresentMatch = in.PresentMatch
	out.RangeMatch = UrlmapRangeMatch_v1beta1_ToProto(mapCtx, in.RangeMatch)
	out.RegexMatch = in.RegexMatch
	out.SuffixMatch = in.SuffixMatch
	return out
}

func UrlmapRequestHeadersToAdd_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapRequestHeadersToAdd {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRequestHeadersToAdd{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func UrlmapRequestHeadersToAdd_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRequestHeadersToAdd) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.HeaderValue = direct.LazyPtr(in.HeaderValue)
	out.Replace = direct.LazyPtr(in.Replace)
	return out
}

func UrlmapResponseHeadersToAdd_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpHeaderOption) *krm.UrlmapResponseHeadersToAdd {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapResponseHeadersToAdd{}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	out.HeaderValue = direct.ValueOf(in.HeaderValue)
	out.Replace = direct.ValueOf(in.Replace)
	return out
}

func UrlmapResponseHeadersToAdd_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapResponseHeadersToAdd) *pb.HttpHeaderOption {
	if in == nil {
		return nil
	}
	out := &pb.HttpHeaderOption{}
	out.HeaderName = direct.LazyPtr(in.HeaderName)
	out.HeaderValue = direct.LazyPtr(in.HeaderValue)
	out.Replace = direct.LazyPtr(in.Replace)
	return out
}

func UrlmapPathMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.UrlmapPathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPathMatcher{}
	out.DefaultCustomErrorResponsePolicy = UrlmapDefaultCustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = UrlmapRouteActionPathMatcherDefaultRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = UrlmapDefaultService_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = UrlmapHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.Name = direct.ValueOf(in.Name)
	out.PathRule = direct.Slice_FromProto(mapCtx, in.PathRules, UrlmapPathRule_v1beta1_FromProto)
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, UrlmapRouteRules_v1beta1_FromProto)
	return out
}

func UrlmapPathMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.DefaultCustomErrorResponsePolicy = UrlmapDefaultCustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = UrlmapRouteActionPathMatcherDefaultRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = UrlmapDefaultService_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_ToProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = UrlmapHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Name = direct.LazyPtr(in.Name)
	out.PathRules = direct.Slice_ToProto(mapCtx, in.PathRule, UrlmapPathRule_v1beta1_ToProto)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, UrlmapRouteRules_v1beta1_ToProto)
	return out
}

func UrlmapTest_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMapTest) *krm.UrlmapTest {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapTest{}
	out.Description = in.Description
	out.Host = direct.ValueOf(in.Host)
	out.Path = direct.ValueOf(in.Path)
	out.Service = direct.ValueOf(UrlmapService_v1beta1_FromProto(mapCtx, in.Service))
	return out
}

func UrlmapTest_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapTest) *pb.UrlMapTest {
	if in == nil {
		return nil
	}
	out := &pb.UrlMapTest{}
	out.Description = in.Description
	out.Host = direct.LazyPtr(in.Host)
	out.Path = direct.LazyPtr(in.Path)
	out.Service = UrlmapService_v1beta1_ToProto(mapCtx, &in.Service)
	return out
}

func UrlmapPathRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathRule) *krm.UrlmapPathRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapPathRule{}
	out.Paths = in.Paths
	out.RouteAction = UrlmapRouteActionPathRule_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = UrlmapService_v1beta1_FromProto(mapCtx, in.Service)
	out.UrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func UrlmapPathRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapPathRule) *pb.PathRule {
	if in == nil {
		return nil
	}
	out := &pb.PathRule{}
	out.Paths = in.Paths
	out.RouteAction = UrlmapRouteActionPathRule_v1beta1_ToProto(mapCtx, in.RouteAction)
	out.Service = UrlmapService_v1beta1_ToProto(mapCtx, in.Service)
	out.UrlRedirect = UrlmapDefaultUrlRedirect_v1beta1_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func UrlmapQueryParameterMatches_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpQueryParameterMatch) *krm.UrlmapQueryParameterMatches {
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

func UrlmapQueryParameterMatches_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapQueryParameterMatches) *pb.HttpQueryParameterMatch {
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

func UrlmapRangeMatch_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Int64RangeMatch) *krm.UrlmapRangeMatch {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRangeMatch{}
	out.RangeEnd = direct.ValueOf(in.RangeEnd)
	out.RangeStart = direct.ValueOf(in.RangeStart)
	return out
}

func UrlmapRangeMatch_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRangeMatch) *pb.Int64RangeMatch {
	if in == nil {
		return nil
	}
	out := &pb.Int64RangeMatch{}
	out.RangeEnd = direct.LazyPtr(in.RangeEnd)
	out.RangeStart = direct.LazyPtr(in.RangeStart)
	return out
}

func UrlmapRetryPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRetryPolicy) *krm.UrlmapRetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRetryPolicy{}
	if in.NumRetries != nil {
		out.NumRetries = direct.LazyPtr(int64(*in.NumRetries))
	}
	out.PerTryTimeout = UrlmapPerTryTimeout_v1beta1_FromProto(mapCtx, in.GetPerTryTimeout())
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapRetryPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRetryPolicy) *pb.HttpRetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRetryPolicy{}
	if in.NumRetries != nil {
		out.NumRetries = direct.LazyPtr(uint32(*in.NumRetries))
	}
	out.PerTryTimeout = UrlmapPerTryTimeout_v1beta1_ToProto(mapCtx, in.PerTryTimeout)
	out.RetryConditions = in.RetryConditions
	return out
}

func UrlmapRouteRules_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRule) *krm.UrlmapRouteRules {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapRouteRules{}
	out.HeaderAction = UrlmapHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, UrlmapMatchRules_v1beta1_FromProto)
	if in.Priority != nil {
		out.Priority = int64(*in.Priority)
	}
	out.RouteAction = UrlmapRouteActionRouteRules_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = in.Service
	out.UrlRedirect = UrlmapUrlRedirect_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func UrlmapRouteRules_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapRouteRules) *pb.HttpRouteRule {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteRule{}
	out.HeaderAction = UrlmapHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, UrlmapMatchRules_v1beta1_ToProto)
	out.Priority = direct.LazyPtr(int32(in.Priority))
	out.RouteAction = UrlmapRouteActionRouteRules_v1beta1_ToProto(mapCtx, in.RouteAction)
	out.Service = in.Service
	out.UrlRedirect = UrlmapUrlRedirect_v1beta1_ToProto(mapCtx, in.UrlRedirect)
	return out
}

func UrlmapDefaultCustomErrorResponsePolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomErrorResponsePolicy) *krm.UrlmapDefaultCustomErrorResponsePolicy {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapDefaultCustomErrorResponsePolicy{}
	out.ErrorResponseRule = direct.Slice_FromProto(mapCtx, in.ErrorResponseRules, UrlmapErrorResponseRule_v1beta1_FromProto)
	if in.GetErrorService() != "" {
		out.ErrorServiceRef = &krm.UrlmapResourceRef{External: in.GetErrorService()}
	}
	return out
}

func UrlmapDefaultCustomErrorResponsePolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapDefaultCustomErrorResponsePolicy) *pb.CustomErrorResponsePolicy {
	if in == nil {
		return nil
	}
	out := &pb.CustomErrorResponsePolicy{}
	out.ErrorResponseRules = direct.Slice_ToProto(mapCtx, in.ErrorResponseRule, UrlmapErrorResponseRule_v1beta1_ToProto)
	if in.ErrorServiceRef != nil && in.ErrorServiceRef.External != "" {
		out.ErrorService = &in.ErrorServiceRef.External
	}
	return out
}

func UrlmapErrorResponseRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CustomErrorResponsePolicyCustomErrorResponseRule) *krm.UrlmapErrorResponseRule {
	if in == nil {
		return nil
	}
	out := &krm.UrlmapErrorResponseRule{}
	out.MatchResponseCodes = in.MatchResponseCodes
	if in.OverrideResponseCode != nil {
		out.OverrideResponseCode = direct.LazyPtr(int64(*in.OverrideResponseCode))
	}
	out.Path = in.Path
	return out
}

func UrlmapErrorResponseRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.UrlmapErrorResponseRule) *pb.CustomErrorResponsePolicyCustomErrorResponseRule {
	if in == nil {
		return nil
	}
	out := &pb.CustomErrorResponsePolicyCustomErrorResponseRule{}
	out.MatchResponseCodes = in.MatchResponseCodes
	if in.OverrideResponseCode != nil {
		out.OverrideResponseCode = direct.LazyPtr(int32(*in.OverrideResponseCode))
	}
	out.Path = in.Path
	return out
}
