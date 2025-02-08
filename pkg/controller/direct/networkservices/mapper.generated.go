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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
)
func GrpcRoute_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute) *krm.GrpcRoute {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Hostnames = in.Hostnames
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GrpcRoute_RouteRule_FromProto)
	return out
}
func GrpcRoute_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute) *pb.GrpcRoute {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Hostnames = in.Hostnames
	out.Meshes = in.Meshes
	out.Gateways = in.Gateways
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GrpcRoute_RouteRule_ToProto)
	return out
}
func GrpcRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute) *krm.GrpcRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRouteObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
func GrpcRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRouteObservedState) *pb.GrpcRoute {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
func GrpcRoute_Destination_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_Destination) *krm.GrpcRoute_Destination {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_Destination{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.Weight = in.Weight
	return out
}
func GrpcRoute_Destination_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_Destination) *pb.GrpcRoute_Destination {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_Destination{}
	if oneof := GrpcRoute_Destination_ServiceName_ToProto(mapCtx, in.ServiceName); oneof != nil {
		out.DestinationType = oneof
	}
	out.Weight = in.Weight
	return out
}
func GrpcRoute_FaultInjectionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_FaultInjectionPolicy) *krm.GrpcRoute_FaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_FaultInjectionPolicy{}
	out.Delay = GrpcRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx, in.GetDelay())
	out.Abort = GrpcRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx, in.GetAbort())
	return out
}
func GrpcRoute_FaultInjectionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_FaultInjectionPolicy) *pb.GrpcRoute_FaultInjectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_FaultInjectionPolicy{}
	if oneof := GrpcRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx, in.Delay); oneof != nil {
		out.Delay = &pb.GrpcRoute_FaultInjectionPolicy_Delay_{Delay: oneof}
	}
	if oneof := GrpcRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx, in.Abort); oneof != nil {
		out.Abort = &pb.GrpcRoute_FaultInjectionPolicy_Abort_{Abort: oneof}
	}
	return out
}
func GrpcRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_FaultInjectionPolicy_Abort) *krm.GrpcRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_FaultInjectionPolicy_Abort{}
	out.HTTPStatus = in.HttpStatus
	out.Percentage = in.Percentage
	return out
}
func GrpcRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_FaultInjectionPolicy_Abort) *pb.GrpcRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_FaultInjectionPolicy_Abort{}
	out.HttpStatus = in.HTTPStatus
	out.Percentage = in.Percentage
	return out
}
func GrpcRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_FaultInjectionPolicy_Delay) *krm.GrpcRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_FromProto(mapCtx, in.GetFixedDelay())
	out.Percentage = in.Percentage
	return out
}
func GrpcRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_FaultInjectionPolicy_Delay) *pb.GrpcRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_FaultInjectionPolicy_Delay{}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.FixedDelay); oneof != nil {
		out.FixedDelay = &pb.GrpcRoute_FaultInjectionPolicy_Delay_FixedDelay{FixedDelay: oneof}
	}
	out.Percentage = in.Percentage
	return out
}
func GrpcRoute_HeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_HeaderMatch) *krm.GrpcRoute_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_HeaderMatch{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func GrpcRoute_HeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_HeaderMatch) *pb.GrpcRoute_HeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_HeaderMatch{}
	out.Type = direct.Enum_ToProto[pb.GrpcRoute_HeaderMatch_Type](mapCtx, in.Type)
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func GrpcRoute_MethodMatch_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_MethodMatch) *krm.GrpcRoute_MethodMatch {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_MethodMatch{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.GrpcService = direct.LazyPtr(in.GetGrpcService())
	out.GrpcMethod = direct.LazyPtr(in.GetGrpcMethod())
	out.CaseSensitive = in.CaseSensitive
	return out
}
func GrpcRoute_MethodMatch_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_MethodMatch) *pb.GrpcRoute_MethodMatch {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_MethodMatch{}
	out.Type = direct.Enum_ToProto[pb.GrpcRoute_MethodMatch_Type](mapCtx, in.Type)
	out.GrpcService = direct.ValueOf(in.GrpcService)
	out.GrpcMethod = direct.ValueOf(in.GrpcMethod)
	out.CaseSensitive = in.CaseSensitive
	return out
}
func GrpcRoute_RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_RetryPolicy) *krm.GrpcRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = direct.LazyPtr(in.GetNumRetries())
	return out
}
func GrpcRoute_RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_RetryPolicy) *pb.GrpcRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = direct.ValueOf(in.NumRetries)
	return out
}
func GrpcRoute_RouteAction_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_RouteAction) *krm.GrpcRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_RouteAction{}
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, GrpcRoute_Destination_FromProto)
	out.FaultInjectionPolicy = GrpcRoute_FaultInjectionPolicy_FromProto(mapCtx, in.GetFaultInjectionPolicy())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.RetryPolicy = GrpcRoute_RetryPolicy_FromProto(mapCtx, in.GetRetryPolicy())
	return out
}
func GrpcRoute_RouteAction_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_RouteAction) *pb.GrpcRoute_RouteAction {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_RouteAction{}
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, GrpcRoute_Destination_ToProto)
	out.FaultInjectionPolicy = GrpcRoute_FaultInjectionPolicy_ToProto(mapCtx, in.FaultInjectionPolicy)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.RetryPolicy = GrpcRoute_RetryPolicy_ToProto(mapCtx, in.RetryPolicy)
	return out
}
func GrpcRoute_RouteMatch_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_RouteMatch) *krm.GrpcRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_RouteMatch{}
	out.Method = GrpcRoute_MethodMatch_FromProto(mapCtx, in.GetMethod())
	out.Headers = direct.Slice_FromProto(mapCtx, in.Headers, GrpcRoute_HeaderMatch_FromProto)
	return out
}
func GrpcRoute_RouteMatch_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_RouteMatch) *pb.GrpcRoute_RouteMatch {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_RouteMatch{}
	if oneof := GrpcRoute_MethodMatch_ToProto(mapCtx, in.Method); oneof != nil {
		out.Method = &pb.GrpcRoute_RouteMatch_Method{Method: oneof}
	}
	out.Headers = direct.Slice_ToProto(mapCtx, in.Headers, GrpcRoute_HeaderMatch_ToProto)
	return out
}
func GrpcRoute_RouteRule_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_RouteRule) *krm.GrpcRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_RouteRule{}
	out.Matches = direct.Slice_FromProto(mapCtx, in.Matches, GrpcRoute_RouteMatch_FromProto)
	out.Action = GrpcRoute_RouteAction_FromProto(mapCtx, in.GetAction())
	return out
}
func GrpcRoute_RouteRule_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_RouteRule) *pb.GrpcRoute_RouteRule {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_RouteRule{}
	out.Matches = direct.Slice_ToProto(mapCtx, in.Matches, GrpcRoute_RouteMatch_ToProto)
	out.Action = GrpcRoute_RouteAction_ToProto(mapCtx, in.Action)
	return out
}
func NetworkservicesGrpcRouteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute) *krm.NetworkservicesGrpcRouteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesGrpcRouteObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
func NetworkservicesGrpcRouteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesGrpcRouteObservedState) *pb.GrpcRoute {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
func NetworkservicesGrpcRouteSpec_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute) *krm.NetworkservicesGrpcRouteSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesGrpcRouteSpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
func NetworkservicesGrpcRouteSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesGrpcRouteSpec) *pb.GrpcRoute {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: Hostnames
	// MISSING: Meshes
	// MISSING: Gateways
	// MISSING: Rules
	return out
}
