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
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GrpcRoute_Destination_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_Destination) *krm.GrpcRoute_Destination {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_Destination{}
	if serviceName := in.GetServiceName(); serviceName != "" {
		out.ServiceRef = &krmcomputev1beta1.ComputeBackendServiceRef{External: serviceName}
	}
	if in.Weight != nil {
		out.Weight = direct.LazyPtr(int64(*in.Weight))
	}
	return out
}

func GrpcRoute_Destination_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_Destination) *pb.GrpcRoute_Destination {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_Destination{}
	if in.ServiceRef != nil {
		out.DestinationType = &pb.GrpcRoute_Destination_ServiceName{ServiceName: in.ServiceRef.External}
	}
	if in.Weight != nil {
		out.Weight = direct.LazyPtr(int32(*in.Weight))
	}
	return out
}

func GrpcRoute_FaultInjectionPolicy_Abort_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_FaultInjectionPolicy_Abort) *krm.GrpcRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_FaultInjectionPolicy_Abort{}
	if in.HttpStatus != nil {
		out.HTTPStatus = direct.LazyPtr(int64(*in.HttpStatus))
	}
	if in.Percentage != nil {
		out.Percentage = direct.LazyPtr(int64(*in.Percentage))
	}
	return out
}

func GrpcRoute_FaultInjectionPolicy_Abort_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_FaultInjectionPolicy_Abort) *pb.GrpcRoute_FaultInjectionPolicy_Abort {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_FaultInjectionPolicy_Abort{}
	if in.HTTPStatus != nil {
		out.HttpStatus = direct.LazyPtr(int32(*in.HTTPStatus))
	}
	if in.Percentage != nil {
		out.Percentage = direct.LazyPtr(int32(*in.Percentage))
	}
	return out
}

func GrpcRoute_FaultInjectionPolicy_Delay_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_FaultInjectionPolicy_Delay) *krm.GrpcRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_FromProto(mapCtx, in.GetFixedDelay())
	if in.Percentage != nil {
		out.Percentage = direct.LazyPtr(int64(*in.Percentage))
	}
	return out
}

func GrpcRoute_FaultInjectionPolicy_Delay_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_FaultInjectionPolicy_Delay) *pb.GrpcRoute_FaultInjectionPolicy_Delay {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_FaultInjectionPolicy_Delay{}
	out.FixedDelay = direct.StringDuration_ToProto(mapCtx, in.FixedDelay)
	if in.Percentage != nil {
		out.Percentage = direct.LazyPtr(int32(*in.Percentage))
	}
	return out
}

func GrpcRoute_RetryPolicy_FromProto(mapCtx *direct.MapContext, in *pb.GrpcRoute_RetryPolicy) *krm.GrpcRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &krm.GrpcRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = direct.LazyPtr(int64(in.GetNumRetries()))
	return out
}

func GrpcRoute_RetryPolicy_ToProto(mapCtx *direct.MapContext, in *krm.GrpcRoute_RetryPolicy) *pb.GrpcRoute_RetryPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GrpcRoute_RetryPolicy{}
	out.RetryConditions = in.RetryConditions
	out.NumRetries = uint32(direct.ValueOf(in.NumRetries))
	return out
}

func NetworkServicesGRPCRouteSpec_Meshes_FromProto(mapCtx *direct.MapContext, in []string) []*refsv1beta1.NetworkServicesMeshRef {
	if in == nil {
		return nil
	}
	out := make([]*refsv1beta1.NetworkServicesMeshRef, len(in))
	for i, v := range in {
		out[i] = &refsv1beta1.NetworkServicesMeshRef{External: v}
	}
	return out
}

func NetworkServicesGRPCRouteSpec_Meshes_ToProto(mapCtx *direct.MapContext, in []*refsv1beta1.NetworkServicesMeshRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, 0, len(in))
	for _, v := range in {
		if v != nil {
			out = append(out, v.External)
		}
	}
	return out
}

func NetworkServicesGRPCRouteSpec_Gateways_FromProto(mapCtx *direct.MapContext, in []string) []*refsv1beta1.NetworkServicesGatewayRef {
	if in == nil {
		return nil
	}
	out := make([]*refsv1beta1.NetworkServicesGatewayRef, len(in))
	for i, v := range in {
		out[i] = &refsv1beta1.NetworkServicesGatewayRef{External: v}
	}
	return out
}

func NetworkServicesGRPCRouteSpec_Gateways_ToProto(mapCtx *direct.MapContext, in []*refsv1beta1.NetworkServicesGatewayRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, 0, len(in))
	for _, v := range in {
		if v != nil {
			out = append(out, v.External)
		}
	}
	return out
}
