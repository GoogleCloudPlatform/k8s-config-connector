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

func ComputeURLMapHTTPRouteRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRouteRule) *krm.ComputeURLMapHTTPRouteRule {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapHTTPRouteRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetCustomErrorResponsePolicy())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.MatchRules = direct.Slice_FromProto(mapCtx, in.MatchRules, HTTPRouteRuleMatch_v1beta1_FromProto)
	out.Priority = in.Priority
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx, in.Service)
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
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
	out.Service = ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx, in.Service)
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func ComputeURLMapPathMatcher_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathMatcher) *krm.ComputeURLMapPathMatcher {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapPathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.Name = in.Name
	out.PathRule = direct.Slice_FromProto(mapCtx, in.PathRules, ComputeURLMapPathRule_v1beta1_FromProto)
	out.RouteRules = direct.Slice_FromProto(mapCtx, in.RouteRules, ComputeURLMapHTTPRouteRule_v1beta1_FromProto)
	return out
}
func ComputeURLMapPathMatcher_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapPathMatcher) *pb.PathMatcher {
	if in == nil {
		return nil
	}
	out := &pb.PathMatcher{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.Name = in.Name
	out.PathRules = direct.Slice_ToProto(mapCtx, in.PathRule, ComputeURLMapPathRule_v1beta1_ToProto)
	out.RouteRules = direct.Slice_ToProto(mapCtx, in.RouteRules, ComputeURLMapHTTPRouteRule_v1beta1_ToProto)
	return out
}

func ComputeURLMapPathRule_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.PathRule) *krm.ComputeURLMapPathRule {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapPathRule{}
	out.CustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetCustomErrorResponsePolicy())
	out.Paths = in.Paths
	out.RouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetRouteAction())
	out.Service = ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx, in.Service)
	out.URLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetUrlRedirect())
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
	out.Service = ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx, in.Service)
	out.UrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.URLRedirect)
	return out
}

func ComputeURLMapSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.UrlMap) *krm.ComputeURLMapSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeURLMapSpec{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_FromProto(mapCtx, in.GetDefaultCustomErrorResponsePolicy())
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_FromProto(mapCtx, in.GetDefaultRouteAction())
	out.DefaultService = ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx, in.DefaultService)
	out.DefaultURLRedirect = HTTPRedirectAction_v1beta1_FromProto(mapCtx, in.GetDefaultUrlRedirect())
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_FromProto(mapCtx, in.GetHeaderAction())
	out.HostRule = direct.Slice_FromProto(mapCtx, in.HostRules, HostRule_v1beta1_FromProto)
	// MISSING: Kind
	// MISSING: Name
	out.PathMatcher = direct.Slice_FromProto(mapCtx, in.PathMatchers, ComputeURLMapPathMatcher_v1beta1_FromProto)
	// MISSING: Region
	out.Test = direct.Slice_FromProto(mapCtx, in.Tests, ComputeURLMapTest_v1beta1_FromProto)
	return out
}
func ComputeURLMapSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapSpec) *pb.UrlMap {
	if in == nil {
		return nil
	}
	out := &pb.UrlMap{}
	out.DefaultCustomErrorResponsePolicy = CustomErrorResponsePolicy_v1beta1_ToProto(mapCtx, in.DefaultCustomErrorResponsePolicy)
	out.DefaultRouteAction = ComputeURLMapHTTPRouteAction_v1beta1_ToProto(mapCtx, in.DefaultRouteAction)
	out.DefaultService = ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx, in.DefaultService)
	out.DefaultUrlRedirect = HTTPRedirectAction_v1beta1_ToProto(mapCtx, in.DefaultURLRedirect)
	out.Description = in.Description
	out.HeaderAction = HTTPHeaderAction_v1beta1_ToProto(mapCtx, in.HeaderAction)
	out.HostRules = direct.Slice_ToProto(mapCtx, in.HostRule, HostRule_v1beta1_ToProto)
	// MISSING: Kind
	// MISSING: Name
	out.PathMatchers = direct.Slice_ToProto(mapCtx, in.PathMatcher, ComputeURLMapPathMatcher_v1beta1_ToProto)
	// MISSING: Region
	out.Tests = direct.Slice_ToProto(mapCtx, in.Test, ComputeURLMapTest_v1beta1_ToProto)
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
	out.Service = ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx, in.Service)
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
	out.Service = ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx, in.Service)
	return out
}
