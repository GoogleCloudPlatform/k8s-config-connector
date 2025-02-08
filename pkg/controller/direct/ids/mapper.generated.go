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

package ids

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ids/apiv1/idspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ids/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	// MISSING: State
	out.TrafficLogs = direct.LazyPtr(in.GetTrafficLogs())
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Network = direct.ValueOf(in.Network)
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	out.Description = direct.ValueOf(in.Description)
	out.Severity = direct.Enum_ToProto[pb.Endpoint_Severity](mapCtx, in.Severity)
	// MISSING: State
	out.TrafficLogs = direct.ValueOf(in.TrafficLogs)
	return out
}
func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Network
	out.EndpointForwardingRule = direct.LazyPtr(in.GetEndpointForwardingRule())
	out.EndpointIP = direct.LazyPtr(in.GetEndpointIp())
	// MISSING: Description
	// MISSING: Severity
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: TrafficLogs
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Network
	out.EndpointForwardingRule = direct.ValueOf(in.EndpointForwardingRule)
	out.EndpointIp = direct.ValueOf(in.EndpointIP)
	// MISSING: Description
	// MISSING: Severity
	out.State = direct.Enum_ToProto[pb.Endpoint_State](mapCtx, in.State)
	// MISSING: TrafficLogs
	return out
}
func IdsEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.IdsEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IdsEndpointObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	// MISSING: Description
	// MISSING: Severity
	// MISSING: State
	// MISSING: TrafficLogs
	return out
}
func IdsEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IdsEndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	// MISSING: Description
	// MISSING: Severity
	// MISSING: State
	// MISSING: TrafficLogs
	return out
}
func IdsEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.IdsEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.IdsEndpointSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	// MISSING: Description
	// MISSING: Severity
	// MISSING: State
	// MISSING: TrafficLogs
	return out
}
func IdsEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.IdsEndpointSpec) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Network
	// MISSING: EndpointForwardingRule
	// MISSING: EndpointIP
	// MISSING: Description
	// MISSING: Severity
	// MISSING: State
	// MISSING: TrafficLogs
	return out
}
