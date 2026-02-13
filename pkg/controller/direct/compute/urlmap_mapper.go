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
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapSpec{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = HTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.HostRules = direct.Slice_FromProto(mapCtx, in.HostRules, HostRule_v1beta1_FromProto)
	out.PathMatchers = direct.Slice_FromProto(mapCtx, in.PathMatchers, PathMatcher_v1beta1_FromProto)
	out.Tests = direct.Slice_FromProto(mapCtx, in.Tests, URLMapTest_v1beta1_FromProto)
	return out
}

func ComputeURLMapSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapSpec) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = HTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRules, HostRule_v1beta1_ToProto)
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatchers, PathMatcher_v1beta1_ToProto)
	out.Tests = direct.Slice_ToProto(mapCtx, in.Tests, URLMapTest_v1beta1_ToProto)
	return out
}

func PathMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.PathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.PathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = HTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.Name = in.Name
	out.PathRules = direct.Slice_FromProto(mapCtx, in.PathRules, PathRule_v1beta1_FromProto)
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, HTTPRouteRule_v1beta1_FromProto)
	return out
}

func PathMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = HTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Name = in.Name
	out.PathRules = direct.Slice_ToProto(mapCtx, in.PathRules, PathRule_v1beta1_ToProto)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, HTTPRouteRule_v1beta1_ToProto)
	return out
}

func PathRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathRule) *krm.PathRule {
	if in == nil {
		return nil
	}
	out := &krm.PathRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetCustomErrorResponsePolicy())
	out.Paths = in.Paths
	out.RouteAction = HTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx, in.Service)
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func PathRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.PathRule) *pb.PathRule {
	if in == nil {
		return nil
	}
	out := &pb.PathRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Paths = in.Paths
	out.RouteAction = HTTPRouteAction_v1beta1_ToProto(mapCtx, in.RouteAction)
	out.Service = ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx, in.Service)
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func HTTPRouteRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRule) *krm.HTTPRouteRule {
	if in == nil {
		return nil
	}
	out := &krm.HTTPRouteRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetCustomErrorResponsePolicy())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, HTTPRouteRuleMatch_v1beta1_FromProto)
	out.Priority = in.Priority
	out.RouteAction = HTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx, in.Service)
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
	return out
}

func HTTPRouteRule_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HTTPRouteRule) *pb.HttpRouteRule {
	if in == nil {
		return nil
	}
	out := &pb.HttpRouteRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.CustomErrorResponsePolicy)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.MatchRules = direct.Slice_ToProto(mapCtx, in.MatchRules, HTTPRouteRuleMatch_v1beta1_ToProto)
	out.Priority = in.Priority
	out.RouteAction = HTTPRouteAction_v1beta1_ToProto(mapCtx, in.RouteAction)
	out.Service = ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx, in.Service)
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func WeightedBackendService_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.WeightedBackendService) *krm.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &krm.WeightedBackendService{}
	if in.GetBackendService() != "" {
		out.BackendService = &krm.ComputeBackendServiceRef{External: in.GetBackendService()}
	}
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.Weight = in.Weight
	return out
}

func WeightedBackendService_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.WeightedBackendService) *pb.WeightedBackendService {
	if in == nil {
		return nil
	}
	out := &pb.WeightedBackendService{}
	if in.BackendService != nil {
		out.BackendService = &in.BackendService.External
	}
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Weight = in.Weight
	return out
}

func URLMapTest_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMapTest) *krm.URLMapTest {
	if in == nil {
		return nil
	}
	out := &krm.URLMapTest{}
	out.Description = in.Description
	out.ExpectedOutputURL = in.ExpectedOutputUrl
	out.ExpectedRedirectResponseCode = in.ExpectedRedirectResponseCode
	out.Headers = direct.Slice_FromProto(mapCtx, in.Headers, URLMapTestHeader_v1beta1_FromProto)
	out.Host = in.Host
	out.Path = in.Path
	out.Service = ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx, in.Service)
	return out
}

func URLMapTest_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.URLMapTest) *pb.UrlMapTest {
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
	out.Service = ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx, in.Service)
	return out
}
