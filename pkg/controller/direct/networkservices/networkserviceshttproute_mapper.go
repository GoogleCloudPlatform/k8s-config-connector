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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesHTTPRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.NetworkServicesHTTPRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesHTTPRouteSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hostnames = in.GetHostnames()

	if in.GetGateways() != nil {
		out.Gateways = make([]krm.NetworkServicesGatewayRef, len(in.GetGateways()))
		for i, g := range in.GetGateways() {
			out.Gateways[i] = krm.NetworkServicesGatewayRef{External: g}
		}
	}
	if in.GetMeshes() != nil {
		out.Meshes = make([]v1alpha1.NetworkServicesMeshRef, len(in.GetMeshes()))
		for i, m := range in.GetMeshes() {
			out.Meshes[i] = v1alpha1.NetworkServicesMeshRef{External: m}
		}
	}

	out.Rules = direct.Slice_FromProto(mapCtx, in.GetRules(), HttprouteRules_FromProto)
	return out
}

func NetworkServicesHTTPRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesHTTPRouteSpec) *pb.HttpRoute {
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

func NetworkServicesHTTPRouteStatus_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute) *krm.NetworkServicesHTTPRouteStatus {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesHTTPRouteStatus{}
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkServicesHTTPRouteStatus_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesHTTPRouteStatus) *pb.HttpRoute {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute{}
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func HttprouteDestination_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krm.HttprouteDestination {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteDestination{}
	if in.GetServiceName() != "" {
		out.ServiceRef = &computev1beta1.ComputeBackendServiceRef{
			External: in.GetServiceName(),
		}
	}
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func HttprouteDestination_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteDestination) *pb.HttpRoute_Destination {
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

func HttprouteDestinations_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Destination) *krm.HttprouteDestinations {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteDestinations{}
	if in.GetServiceName() != "" {
		out.ServiceRef = &computev1beta1.ComputeBackendServiceRef{
			External: in.GetServiceName(),
		}
	}
	out.Weight = direct.LazyPtr(int64(in.GetWeight()))
	return out
}

func HttprouteDestinations_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteDestinations) *pb.HttpRoute_Destination {
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

func HttprouteAbort_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Abort) *krm.HttprouteAbort {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteAbort{}
	out.HTTPStatus = direct.LazyPtr(int64(in.GetHttpStatus()))
	out.Percentage = direct.LazyPtr(int64(in.GetPercentage()))
	return out
}

func HttprouteAbort_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteAbort) *pb.HttpRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Abort{}
	out.HttpStatus = int32(direct.ValueOf(in.HTTPStatus))
	out.Percentage = int32(direct.ValueOf(in.Percentage))
	return out
}

func HttprouteRangeMatch_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_HeaderMatch_IntegerRange) *krm.HttprouteRangeMatch {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteRangeMatch{}
	out.Start = direct.LazyPtr(int64(in.GetStart()))
	out.End = direct.LazyPtr(int64(in.GetEnd()))
	return out
}

func HttprouteRangeMatch_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteRangeMatch) *pb.HttpRoute_HeaderMatch_IntegerRange {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_HeaderMatch_IntegerRange{}
	out.Start = int32(direct.ValueOf(in.Start))
	out.End = int32(direct.ValueOf(in.End))
	return out
}

func HttprouteRedirect_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_Redirect) *krm.HttprouteRedirect {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteRedirect{}
	out.HostRedirect = direct.LazyPtr(in.GetHostRedirect())
	out.PathRedirect = direct.LazyPtr(in.GetPathRedirect())
	out.PrefixRewrite = direct.LazyPtr(in.GetPrefixRewrite())
	out.ResponseCode = direct.Enum_FromProto(mapCtx, in.GetResponseCode())
	out.HTTPSRedirect = direct.LazyPtr(in.GetHttpsRedirect())
	out.StripQuery = direct.LazyPtr(in.GetStripQuery())
	out.PortRedirect = direct.LazyPtr(int64(in.GetPortRedirect()))
	return out
}

func HttprouteRedirect_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteRedirect) *pb.HttpRoute_Redirect {
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

func HttprouteDelay_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_FaultInjectionPolicy_Delay) *krm.HttprouteDelay {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteDelay{}
	out.FixedDelay = direct.StringDuration_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = direct.LazyPtr(int64(in.GetPercentage()))
	return out
}

func HttprouteDelay_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteDelay) *pb.HttpRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_ToProto(mapCtx, in.FixedDelay)
	out.Percentage = int32(direct.ValueOf(in.Percentage))
	return out
}

func HttprouteRetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.HttpRoute_RetryPolicy) *krm.HttprouteRetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.HttprouteRetryPolicy{}
	out.RetryConditions = in.GetRetryConditions()
	out.NumRetries = direct.LazyPtr(int64(in.GetNumRetries()))
	out.PerTryTimeout = direct.StringDuration_FromProto(mapCtx, in.GetPerTryTimeout())
	return out
}

func HttprouteRetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.HttprouteRetryPolicy) *pb.HttpRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.HttpRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = int32(direct.ValueOf(in.NumRetries))
	out.PerTryTimeout = direct.StringDuration_ToProto(mapCtx, in.PerTryTimeout)
	return out
}
