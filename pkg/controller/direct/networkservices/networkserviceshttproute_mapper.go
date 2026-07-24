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

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krmnetworkservicesv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	krmnetworkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesHTTPRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteSpec {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.NetworkServicesHTTPRouteSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hostnames = in.GetHostnames()

	if in.GetGateways() != nil {
		out.Gateways = make([]krmnetworkservicesv1beta1.NetworkServicesGatewayRef, len(in.GetGateways()))
		for i, g := range in.GetGateways() {
			out.Gateways[i] = krmnetworkservicesv1beta1.NetworkServicesGatewayRef{External: g}
		}
	}
	if in.GetMeshes() != nil {
		out.Meshes = make([]krmnetworkservicesv1alpha1.NetworkServicesMeshRef, len(in.GetMeshes()))
		for i, m := range in.GetMeshes() {
			out.Meshes[i] = krmnetworkservicesv1alpha1.NetworkServicesMeshRef{External: m}
		}
	}

	out.Rules = direct.Slice_FromProto(mapCtx, in.GetRules(), HttprouteRules_FromProto)
	return out
}

func NetworkServicesHTTPRouteSpec_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteSpec) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	out.Description = direct.ValueOf(in.Description)
	out.Hostnames = in.Hostnames

	if in.Gateways != nil {
		out.Gateways = make([]string, len(in.Gateways))
		for i, g := range in.Gateways {
			out.Gateways[i] = g.External
		}
	}
	if in.Meshes != nil {
		out.Meshes = make([]string, len(in.Meshes))
		for i, m := range in.Meshes {
			out.Meshes[i] = m.External
		}
	}

	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, HttprouteRules_ToProto)
	return out
}

func NetworkServicesHTTPRouteStatus_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteStatus {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.NetworkServicesHTTPRouteStatus{}
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkServicesHTTPRouteStatus_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteStatus) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func NetworkServicesHTTPRouteSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteSpec {
	return NetworkServicesHTTPRouteSpec_FromProto(mapCtx, in)
}
func NetworkServicesHTTPRouteSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteSpec) *pb.HttpRoute {
	return NetworkServicesHTTPRouteSpec_ToProto(mapCtx, in)
}
func NetworkServicesHTTPRouteStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteStatus {
	return NetworkServicesHTTPRouteStatus_FromProto(mapCtx, in)
}
func NetworkServicesHTTPRouteStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.NetworkServicesHTTPRouteStatus) *pb.HttpRoute {
	return NetworkServicesHTTPRouteStatus_ToProto(mapCtx, in)
}

func HttprouteDestination_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krmnetworkservicesv1beta1.HttprouteDestination {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteDestination{}
	if in.GetServiceName() != "" {
		out.ServiceRef = &krmnetworkservicesv1beta1.BackendServiceRef{
			External: in.GetServiceName(),
		}
	}
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func HttprouteDestination_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDestination) *pb.HttpRoute_Destination {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_Destination{}
	if in.ServiceRef != nil {
		out.ServiceName = in.ServiceRef.External
	}
	out.Weight = int32(direct.ValueOf(in.Weight))
	return out
}

func HttprouteDestinations_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krmnetworkservicesv1beta1.HttprouteDestinations {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteDestinations{}
	if in.GetServiceName() != "" {
		out.ServiceRef = &krmnetworkservicesv1beta1.BackendServiceRef{
			External: in.GetServiceName(),
		}
	}
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func HttprouteDestinations_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDestinations) *pb.HttpRoute_Destination {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_Destination{}
	if in.ServiceRef != nil {
		out.ServiceName = in.ServiceRef.External
	}
	out.Weight = int32(direct.ValueOf(in.Weight))
	return out
}

func HttprouteAbort_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Abort) *krmnetworkservicesv1beta1.HttprouteAbort {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteAbort{}
	out.HTTPStatus = direct.LazyPtr(int64(in.GetHttpStatus()))
	out.Percentage = direct.LazyPtr(int64(in.GetPercentage()))
	return out
}

func HttprouteAbort_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteAbort) *pb.HttpRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Abort{}
	out.HttpStatus = int32(direct.ValueOf(in.HTTPStatus))
	out.Percentage = int32(direct.ValueOf(in.Percentage))
	return out
}

func HttprouteRangeMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch_IntegerRange) *krmnetworkservicesv1beta1.HttprouteRangeMatch {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteRangeMatch{}
	out.Start = direct.LazyPtr(int64(in.GetStart()))
	out.End = direct.LazyPtr(int64(in.GetEnd()))
	return out
}

func HttprouteRangeMatch_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRangeMatch) *pb.HttpRoute_HeaderMatch_IntegerRange {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderMatch_IntegerRange{}
	out.Start = int32(direct.ValueOf(in.Start))
	out.End = int32(direct.ValueOf(in.End))
	return out
}

func HttprouteRedirect_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Redirect) *krmnetworkservicesv1beta1.HttprouteRedirect {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRewrite = direct.LazyPtr(in.GetPrefixRewrite())
	out.ResponseCode = direct.Enum_FromProto(mapCtx, in.GetResponseCode())
	out.HTTPSRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	out.PortRedirect = direct.LazyPtr(int64(in.GetPortRedirect()))
	return out
}

func HttprouteRedirect_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRedirect) *pb.HttpRoute_Redirect {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_Redirect{}
	out.HostRedirect = direct.ValueOf(in.HostRedirect)
	out.PathRedirect = direct.ValueOf(in.PathRedirect)
	out.PrefixRewrite = direct.ValueOf(in.PrefixRewrite)
	out.ResponseCode = direct.Enum_ToProto[pb.HttpRoute_Redirect_ResponseCode](mapCtx, in.ResponseCode)
	out.HttpsRedirect = direct.ValueOf(in.HTTPSRedirect)
	out.StripQuery = direct.ValueOf(in.StripQuery)
	out.PortRedirect = int32(direct.ValueOf(in.PortRedirect))
	return out
}

func HttprouteDelay_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Delay) *krmnetworkservicesv1beta1.HttprouteDelay {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteDelay{}
	out.FixedDelay = direct.StringDuration_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = direct.LazyPtr(int64(in.GetPercentage()))
	return out
}

func HttprouteDelay_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDelay) *pb.HttpRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_ToProto(mapCtx, in.FixedDelay)
	out.Percentage = int32(direct.ValueOf(in.Percentage))
	return out
}

func HttprouteRetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RetryPolicy) *krmnetworkservicesv1beta1.HttprouteRetryPolicy {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteRetryPolicy{}
	out.RetryConditions = in.GetRetryConditions()
	out.NumRetries = direct.LazyPtr(int64(in.GetNumRetries()))
	out.PerTryTimeout = direct.StringDuration_FromProto(mapCtx, in.GetPerTryTimeout())
	return out
}

func HttprouteRetryPolicy_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRetryPolicy) *pb.HttpRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = int32(direct.ValueOf(in.NumRetries))
	out.PerTryTimeout = direct.StringDuration_ToProto(mapCtx, in.PerTryTimeout)
	return out
}

func HttprouteQueryParameters_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_QueryParameterMatch) *krmnetworkservicesv1beta1.HttprouteQueryParameters {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteQueryParameters{}
	out.QueryParameter = direct.LazyPtr(in.GetQueryParameter())

	switch m := in.MatchType.(type) {
	case *pb.HttpRoute_QueryParameterMatch_ExactMatch:
		out.ExactMatch = &m.ExactMatch
	case *pb.HttpRoute_QueryParameterMatch_RegexMatch:
		out.RegexMatch = &m.RegexMatch
	case *pb.HttpRoute_QueryParameterMatch_PresentMatch:
		out.PresentMatch = &m.PresentMatch
	}
	return out
}

func HttprouteQueryParameters_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteQueryParameters) *pb.HttpRoute_QueryParameterMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_QueryParameterMatch{}
	if in.ExactMatch != nil {
		out.MatchType = &pb.HttpRoute_QueryParameterMatch_ExactMatch{ExactMatch: *in.ExactMatch}
	}
	if in.RegexMatch != nil {
		out.MatchType = &pb.HttpRoute_QueryParameterMatch_RegexMatch{RegexMatch: *in.RegexMatch}
	}
	if in.PresentMatch != nil {
		out.MatchType = &pb.HttpRoute_QueryParameterMatch_PresentMatch{PresentMatch: *in.PresentMatch}
	}
	out.QueryParameter = direct.ValueOf(in.QueryParameter)
	return out
}

func HttprouteHeaders_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch) *krmnetworkservicesv1beta1.HttprouteHeaders {
	if in == nil {
		return nil
	}
	out := &krmnetworkservicesv1beta1.HttprouteHeaders{}
	out.Header = direct.LazyPtr(in.GetHeader())
	out.InvertMatch = direct.LazyPtr(in.GetInvertMatch())

	switch m := in.MatchType.(type) {
	case *pb.HttpRoute_HeaderMatch_ExactMatch:
		out.ExactMatch = &m.ExactMatch
	case *pb.HttpRoute_HeaderMatch_RegexMatch:
		out.RegexMatch = &m.RegexMatch
	case *pb.HttpRoute_HeaderMatch_PrefixMatch:
		out.PrefixMatch = &m.PrefixMatch
	case *pb.HttpRoute_HeaderMatch_PresentMatch:
		out.PresentMatch = &m.PresentMatch
	case *pb.HttpRoute_HeaderMatch_SuffixMatch:
		out.SuffixMatch = &m.SuffixMatch
	case *pb.HttpRoute_HeaderMatch_RangeMatch:
		out.RangeMatch = HttprouteRangeMatch_FromProto(mapCtx, m.RangeMatch)
	}
	return out
}

func HttprouteHeaders_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteHeaders) *pb.HttpRoute_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderMatch{}
	if in.ExactMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_ExactMatch{ExactMatch: *in.ExactMatch}
	}
	if in.RegexMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_RegexMatch{RegexMatch: *in.RegexMatch}
	}
	if in.PrefixMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_PrefixMatch{PrefixMatch: *in.PrefixMatch}
	}
	if in.PresentMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_PresentMatch{PresentMatch: *in.PresentMatch}
	}
	if in.SuffixMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_SuffixMatch{SuffixMatch: *in.SuffixMatch}
	}
	if in.RangeMatch != nil {
		out.MatchType = &pb.HttpRoute_HeaderMatch_RangeMatch{RangeMatch: HttprouteRangeMatch_ToProto(mapCtx, in.RangeMatch)}
	}
	out.Header = direct.ValueOf(in.Header)
	out.InvertMatch = direct.ValueOf(in.InvertMatch)
	return out
}

func HttpRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Abort) *krmnetworkservicesv1beta1.HttprouteAbort {
	return HttprouteAbort_FromProto(mapCtx, in)
}
func HttpRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteAbort) *pb.HttpRoute_FaultInjectionPolicy_Abort {
	return HttprouteAbort_ToProto(mapCtx, in)
}
func HttprouteAbort_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Abort) *krmnetworkservicesv1beta1.HttprouteAbort {
	return HttprouteAbort_FromProto(mapCtx, in)
}
func HttprouteAbort_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteAbort) *pb.HttpRoute_FaultInjectionPolicy_Abort {
	return HttprouteAbort_ToProto(mapCtx, in)
}

func HttpRoute_Destination_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krmnetworkservicesv1beta1.HttprouteDestination {
	return HttprouteDestination_FromProto(mapCtx, in)
}
func HttpRoute_Destination_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDestination) *pb.HttpRoute_Destination {
	return HttprouteDestination_ToProto(mapCtx, in)
}
func HttprouteDestination_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krmnetworkservicesv1beta1.HttprouteDestination {
	return HttprouteDestination_FromProto(mapCtx, in)
}
func HttprouteDestination_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDestination) *pb.HttpRoute_Destination {
	return HttprouteDestination_ToProto(mapCtx, in)
}
func HttprouteDestinations_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krmnetworkservicesv1beta1.HttprouteDestinations {
	return HttprouteDestinations_FromProto(mapCtx, in)
}
func HttprouteDestinations_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDestinations) *pb.HttpRoute_Destination {
	return HttprouteDestinations_ToProto(mapCtx, in)
}

func HttpRoute_HeaderMatch_IntegerRange_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch_IntegerRange) *krmnetworkservicesv1beta1.HttprouteRangeMatch {
	return HttprouteRangeMatch_FromProto(mapCtx, in)
}
func HttpRoute_HeaderMatch_IntegerRange_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRangeMatch) *pb.HttpRoute_HeaderMatch_IntegerRange {
	return HttprouteRangeMatch_ToProto(mapCtx, in)
}
func HttprouteRangeMatch_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch_IntegerRange) *krmnetworkservicesv1beta1.HttprouteRangeMatch {
	return HttprouteRangeMatch_FromProto(mapCtx, in)
}
func HttprouteRangeMatch_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRangeMatch) *pb.HttpRoute_HeaderMatch_IntegerRange {
	return HttprouteRangeMatch_ToProto(mapCtx, in)
}

func HttpRoute_Redirect_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Redirect) *krmnetworkservicesv1beta1.HttprouteRedirect {
	return HttprouteRedirect_FromProto(mapCtx, in)
}
func HttpRoute_Redirect_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRedirect) *pb.HttpRoute_Redirect {
	return HttprouteRedirect_ToProto(mapCtx, in)
}
func HttprouteRedirect_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Redirect) *krmnetworkservicesv1beta1.HttprouteRedirect {
	return HttprouteRedirect_FromProto(mapCtx, in)
}
func HttprouteRedirect_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRedirect) *pb.HttpRoute_Redirect {
	return HttprouteRedirect_ToProto(mapCtx, in)
}

func HttpRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Delay) *krmnetworkservicesv1beta1.HttprouteDelay {
	return HttprouteDelay_FromProto(mapCtx, in)
}
func HttpRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDelay) *pb.HttpRoute_FaultInjectionPolicy_Delay {
	return HttprouteDelay_ToProto(mapCtx, in)
}
func HttprouteDelay_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Delay) *krmnetworkservicesv1beta1.HttprouteDelay {
	return HttprouteDelay_FromProto(mapCtx, in)
}
func HttprouteDelay_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteDelay) *pb.HttpRoute_FaultInjectionPolicy_Delay {
	return HttprouteDelay_ToProto(mapCtx, in)
}

func HttpRoute_RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RetryPolicy) *krmnetworkservicesv1beta1.HttprouteRetryPolicy {
	return HttprouteRetryPolicy_FromProto(mapCtx, in)
}
func HttpRoute_RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRetryPolicy) *pb.HttpRoute_RetryPolicy {
	return HttprouteRetryPolicy_ToProto(mapCtx, in)
}
func HttprouteRetryPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RetryPolicy) *krmnetworkservicesv1beta1.HttprouteRetryPolicy {
	return HttprouteRetryPolicy_FromProto(mapCtx, in)
}
func HttprouteRetryPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krmnetworkservicesv1beta1.HttprouteRetryPolicy) *pb.HttpRoute_RetryPolicy {
	return HttprouteRetryPolicy_ToProto(mapCtx, in)
}
