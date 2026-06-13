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
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapSpec) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	if in.DefaultService != nil {
		out.DefaultService = ComputeURLMapDefaultService_ToProto(in.DefaultService)
	}
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRules, ComputeURLMapHostRule_v1beta1_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatchers, ComputeURLMapPathMatcher_v1beta1_ToProto)
	out.Tests = direct.Slice_ToProto(mapCtx, in.Tests, ComputeURLMapTest_v1beta1_ToProto)
	return out
}

func ComputeURLMapSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapSpec{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.DefaultRouteAction)
	if in.DefaultService != nil {
		out.DefaultService = ComputeURLMapDefaultService_FromProto(*in.DefaultService)
	}
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_FromProto(mapCtx, in.HostRules, ComputeURLMapHostRule_v1beta1_FromProto)
	out.PathMatchers = direct.Slice_FromProto(mapCtx, in.PathMatchers, ComputeURLMapPathMatcher_v1beta1_FromProto)
	out.Tests = direct.Slice_FromProto(mapCtx, in.Tests, ComputeURLMapTest_v1beta1_FromProto)
	return out
}

func ComputeURLMapObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapObservedState) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.Fingerprint = in.Fingerprint
	return out
}

func ComputeURLMapObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapObservedState{}
	out.Fingerprint = in.Fingerprint
	return out
}

func ComputeURLMapDefaultService_ToProto(in *krm.ComputeURLMapDefaultService) *string {
	if in == nil {
		return nil
	}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		return &in.BackendServiceRef.External
	}
	if in.BackendBucketRef != nil && in.BackendBucketRef.External != "" {
		return &in.BackendBucketRef.External
	}
	return nil
}

func ComputeURLMapDefaultService_FromProto(in string) *krm.ComputeURLMapDefaultService {
	if in == "" {
		return nil
	}
	out := &krm.ComputeURLMapDefaultService{}
	if strings.Contains(in, "/backendBuckets/") {
		out.BackendBucketRef = &krm.ComputeBackendBucketRef{External: in}
	} else {
		out.BackendServiceRef = &krm.ComputeBackendServiceRef{External: in}
	}
	return out
}

func ComputeURLMapHostRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapHostRule) *pb.HostRule {
	if in == nil {
		return nil
	}
	out := &pb.HostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = in.PathMatcher
	return out
}

func ComputeURLMapHostRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HostRule) *krm.ComputeURLMapHostRule {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapHostRule{}
	out.Description = in.Description
	out.Hosts = in.Hosts
	out.PathMatcher = in.PathMatcher
	return out
}

func ComputeURLMapPathMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapPathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	if in.DefaultService != nil {
		out.DefaultService = ComputeURLMapDefaultService_ToProto(in.DefaultService)
	}
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Name = in.Name
	out.PathRules = direct.Slice_ToProto(mapCtx, in.PathRules, ComputeURLMapPathRule_v1beta1_ToProto)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, ComputeURLMapHTTPRouteRule_v1beta1_ToProto)
	return out
}

func ComputeURLMapPathMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.ComputeURLMapPathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapPathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.DefaultRouteAction)
	if in.DefaultService != nil {
		out.DefaultService = ComputeURLMapDefaultService_FromProto(*in.DefaultService)
	}
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.DefaultUrlRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.HeaderAction)
	out.Name = in.Name
	out.PathRules = direct.Slice_FromProto(mapCtx, in.PathRules, ComputeURLMapPathRule_v1beta1_FromProto)
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, ComputeURLMapHTTPRouteRule_v1beta1_FromProto)
	return out
}

func ComputeURLMapPathRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapPathRule) *pb.PathRule {
	if in == nil {
		return nil
	}
	out := &pb.PathRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Paths = in.Paths
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.RouteAction)
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_ToProto(in.Service)
	}
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func ComputeURLMapPathRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathRule) *krm.ComputeURLMapPathRule {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapPathRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Paths = in.Paths
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.RouteAction)
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_FromProto(*in.Service)
	}
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.UrlRedirect)
	return out
}

func ComputeURLMapHTTPRouteRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapHTTPRouteRule) *pb.HttpRouteRule {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, HTTPRouteRuleMatch_v1beta1_ToProto)
	out.Priority = in.Priority
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.RouteAction)
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_ToProto(in.Service)
	}
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func ComputeURLMapHTTPRouteRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRule) *krm.ComputeURLMapHTTPRouteRule {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapHTTPRouteRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, HTTPRouteRuleMatch_v1beta1_FromProto)
	out.Priority = in.Priority
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.RouteAction)
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_FromProto(*in.Service)
	}
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.UrlRedirect)
	return out
}

func ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapHTTPRouteAction) *pb.HttpRouteAction {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteAction{}
	out.CorsPolicy = CorsPolicy_v1beta1_ToProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = HTTPFaultInjection_v1beta1_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.MaxStreamDuration = Duration_v1beta1_ToProto(mapCtx, in.MaxStreamDuration)
	out.RequestMirrorPolicy = ComputeURLMapRequestMirrorPolicy_v1beta1_ToProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = HTTPRetryPolicy_v1beta1_ToProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_v1beta1_ToProto(mapCtx, in.Timeout)
	out.UrlRewrite = URLRewrite_v1beta1_ToProto(mapCtx, in.URLRewrite)
	out.WeightedBackendServices = direct.Slice_ToProto(mapCtx, in.WeightedBackendServices, ComputeURLMapWeightedBackendService_v1beta1_ToProto)
	return out
}

func ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteAction) *krm.ComputeURLMapHTTPRouteAction {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapHTTPRouteAction{}
	out.CorsPolicy = CorsPolicy_v1beta1_FromProto(mapCtx, in.CorsPolicy)
	out.FaultInjectionPolicy = HTTPFaultInjection_v1beta1_FromProto(mapCtx, in.FaultInjectionPolicy)
	out.MaxStreamDuration = Duration_v1beta1_FromProto(mapCtx, in.MaxStreamDuration)
	out.RequestMirrorPolicy = ComputeURLMapRequestMirrorPolicy_v1beta1_FromProto(mapCtx, in.RequestMirrorPolicy)
	out.RetryPolicy = HTTPRetryPolicy_v1beta1_FromProto(mapCtx, in.RetryPolicy)
	out.Timeout = Duration_v1beta1_FromProto(mapCtx, in.Timeout)
	out.URLRewrite = URLRewrite_v1beta1_FromProto(mapCtx, in.UrlRewrite)
	out.WeightedBackendServices = direct.Slice_FromProto(mapCtx, in.WeightedBackendServices, ComputeURLMapWeightedBackendService_v1beta1_FromProto)
	return out
}

func ComputeURLMapRequestMirrorPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapRequestMirrorPolicy) *pb.RequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &pb.RequestMirrorPolicy{}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		out.BackendService = &in.BackendServiceRef.External
	}
	return out
}

func ComputeURLMapRequestMirrorPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.RequestMirrorPolicy) *krm.ComputeURLMapRequestMirrorPolicy {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapRequestMirrorPolicy{}
	if in.BackendService != nil {
		out.BackendServiceRef = &krm.ComputeBackendServiceRef{External: *in.BackendService}
	}
	return out
}

func ComputeURLMapWeightedBackendService_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapWeightedBackendService) *pb.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &pb.WeightedBackendService{}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		out.BackendService = &in.BackendServiceRef.External
	}
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Weight = in.Weight
	return out
}

func ComputeURLMapWeightedBackendService_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.WeightedBackendService) *krm.ComputeURLMapWeightedBackendService {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapWeightedBackendService{}
	if in.BackendService != nil {
		out.BackendServiceRef = &krm.ComputeBackendServiceRef{External: *in.BackendService}
	}
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.HeaderAction)
	out.Weight = in.Weight
	return out
}

func ComputeURLMapTest_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapTest) *pb.UrlMapTest {
	if in == nil {
		return nil
	}
	out := &pb.UrlMapTest{}
	out.Description = in.Description
	out.ExpectedOutputUrl = in.ExpectedOutputURL
	out.ExpectedRedirectResponseCode = in.ExpectedRedirectResponseCode
	out.Headers = direct.Slice_ToProto(mapCtx, in.Headers, URLMapTestHeader_v1beta1_ToProto)
	out.Host = in.Host
	out.Path = in.Path
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_ToProto(in.Service)
	}
	return out
}

func ComputeURLMapTest_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMapTest) *krm.ComputeURLMapTest {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapTest{}
	out.Description = in.Description
	out.ExpectedOutputURL = in.ExpectedOutputUrl
	out.ExpectedRedirectResponseCode = in.ExpectedRedirectResponseCode
	out.Headers = direct.Slice_FromProto(mapCtx, in.Headers, URLMapTestHeader_v1beta1_FromProto)
	out.Host = in.Host
	out.Path = in.Path
	if in.Service != nil {
		out.Service = ComputeURLMapDefaultService_FromProto(*in.Service)
	}
	return out
}
