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

// +generated:mapper
// krm.group: memorystore.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.memorystore.v1

package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	connections := slice_FromProto(mapCtx, in.Connections, Endpoint_ConnectionDetail_FromProto)
	if connections == nil {
		return nil
	}
	out := &krm.Endpoint{}
	out.Connections = connections
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Endpoint_ConnectionDetail_ToProto)
	return out
}
func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_InstanceEndpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	connections := slice_FromProto(mapCtx, in.Connections, Endpoint_ConnectionDetailObservedState_FromProto)
	if connections == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	out.Connections = connections
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Instance_InstanceEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Instance_InstanceEndpoint{}
	out.Connections = direct.Slice_ToProto(mapCtx, in.Connections, Endpoint_ConnectionDetailObservedState_ToProto)
	return out
}
func Endpoint_ConnectionDetail_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krm.Endpoint_ConnectionDetail {
	if in == nil {
		return nil
	}
	userConnection := in.GetPscConnection()
	if userConnection == nil {
		return nil
	}
	out := &krm.Endpoint_ConnectionDetail{}
	out.PscConnection = PscConnection_FromProto(mapCtx, userConnection)
	return out
}
func Endpoint_ConnectionDetail_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_ConnectionDetail) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	if oneof := PscConnection_ToProto(mapCtx, in.PscConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func Endpoint_ConnectionDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_ConnectionDetail) *krm.Endpoint_ConnectionDetailObservedState {
	if in == nil {
		return nil
	}
	userConnection := in.GetPscConnection()
	if userConnection == nil {
		return nil
	}
	out := &krm.Endpoint_ConnectionDetailObservedState{}
	out.PscConnection = PscConnectionObservedState_v1alpha1_FromProto(mapCtx, userConnection)
	return out
}
func Endpoint_ConnectionDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_ConnectionDetailObservedState) *pb.Instance_ConnectionDetail {
	if in == nil {
		return nil
	}
	out := &pb.Instance_ConnectionDetail{}
	if oneof := PscConnectionObservedState_v1alpha1_ToProto(mapCtx, in.PscConnection); oneof != nil {
		out.Connection = &pb.Instance_ConnectionDetail_PscConnection{PscConnection: oneof}
	}
	return out
}
func MemorystoreInstanceEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceEndpointObservedState{}
	out.Endpoints = slice_FromProto(mapCtx, in.Endpoints, EndpointObservedState_FromProto)
	return out
}
func MemorystoreInstanceEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceEndpointObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, EndpointObservedState_ToProto)
	return out
}
func MemorystoreInstanceEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceEndpointSpec{}
	out.Endpoints = slice_FromProto(mapCtx, in.Endpoints, Endpoint_FromProto)
	return out
}
func MemorystoreInstanceEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceEndpointSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, Endpoint_ToProto)
	return out
}
func PscConnection_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krm.PscConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscConnection{}
	if in.Ports != nil {
		out.Port = direct.LazyPtr(in.GetPort())
	}
	out.ForwardingRuleRef = &computev1beta1.ForwardingRuleRef{External: in.GetForwardingRule()}
	return out
}
func PscConnection_ToProto(mapCtx *direct.MapContext, in *krm.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	if in.Port != nil {
		out.Ports = &pb.PscConnection_Port{Port: direct.ValueOf(in.Port)}
	}
	if in.ForwardingRuleRef != nil {
		out.ForwardingRule = in.ForwardingRuleRef.External
	}
	return out
}
func PscConnectionObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krm.PscConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PscConnectionObservedState{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.PscConnectionStatus = direct.Enum_FromProto(mapCtx, in.GetPscConnectionStatus())
	out.ConnectionType = direct.Enum_FromProto(mapCtx, in.GetConnectionType())
	return out
}
func PscConnectionObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.PscConnectionObservedState) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.PscConnectionStatus = direct.Enum_ToProto[pb.PscConnectionStatus](mapCtx, in.PscConnectionStatus)
	out.ConnectionType = direct.Enum_ToProto[pb.ConnectionType](mapCtx, in.ConnectionType)
	return out
}
func slice_FromProto[T, U any](mapCtx *direct.MapContext, in []*T, mapper func(mapCtx *direct.MapContext, in *T) *U) []U {
	if in == nil {
		return nil
	}

	outSlice := make([]U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, inItem)
		if outItem != nil {
			outSlice = append(outSlice, *outItem)
		}
	}
	if len(outSlice) == 0 {
		return nil
	}
	return outSlice
}
