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

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DatastreamPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.DatastreamPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamPrivateConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamPrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.DatastreamPrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamPrivateConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func DatastreamPrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamPrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func Error_FromProto(mapCtx *direct.MapContext, in *pb.Error) *krm.Error {
	if in == nil {
		return nil
	}
	out := &krm.Error{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.ErrorUuid = direct.LazyPtr(in.GetErrorUuid())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetErrorTime())
	out.Details = in.Details
	return out
}
func Error_ToProto(mapCtx *direct.MapContext, in *krm.Error) *pb.Error {
	if in == nil {
		return nil
	}
	out := &pb.Error{}
	out.Reason = direct.ValueOf(in.Reason)
	out.ErrorUuid = direct.ValueOf(in.ErrorUuid)
	out.Message = direct.ValueOf(in.Message)
	out.ErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.ErrorTime)
	out.Details = in.Details
	return out
}
func PrivateConnection_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Error
	out.VpcPeeringConfig = VpcPeeringConfig_FromProto(mapCtx, in.GetVpcPeeringConfig())
	return out
}
func PrivateConnection_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnection) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: Error
	out.VpcPeeringConfig = VpcPeeringConfig_ToProto(mapCtx, in.VpcPeeringConfig)
	return out
}
func PrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Error_FromProto(mapCtx, in.GetError())
	// MISSING: VpcPeeringConfig
	return out
}
func PrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = Error_ToProto(mapCtx, in.Error)
	// MISSING: VpcPeeringConfig
	return out
}
func VpcPeeringConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConfig) *krm.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcPeeringConfig{}
	out.Vpc = direct.LazyPtr(in.GetVpc())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	return out
}
func VpcPeeringConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcPeeringConfig) *pb.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConfig{}
	out.Vpc = direct.ValueOf(in.Vpc)
	out.Subnet = direct.ValueOf(in.Subnet)
	return out
}
