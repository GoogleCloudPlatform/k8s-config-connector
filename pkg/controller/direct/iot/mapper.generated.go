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

package iot

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iot/v1alpha1"
)
func Device_FromProto(mapCtx *direct.MapContext, in *pb.Device) *krm.Device {
	if in == nil {
		return nil
	}
	out := &krm.Device{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	out.NumID = direct.LazyPtr(in.GetNumId())
	out.Credentials = direct.Slice_FromProto(mapCtx, in.Credentials, DeviceCredential_FromProto)
	out.LastHeartbeatTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastHeartbeatTime())
	out.LastEventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastEventTime())
	out.LastStateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastStateTime())
	out.LastConfigAckTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastConfigAckTime())
	out.LastConfigSendTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastConfigSendTime())
	out.Blocked = direct.LazyPtr(in.GetBlocked())
	out.LastErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastErrorTime())
	out.LastErrorStatus = Status_FromProto(mapCtx, in.GetLastErrorStatus())
	out.Config = DeviceConfig_FromProto(mapCtx, in.GetConfig())
	out.State = DeviceState_FromProto(mapCtx, in.GetState())
	out.LogLevel = direct.Enum_FromProto(mapCtx, in.GetLogLevel())
	out.Metadata = in.Metadata
	out.GatewayConfig = GatewayConfig_FromProto(mapCtx, in.GetGatewayConfig())
	return out
}
func Device_ToProto(mapCtx *direct.MapContext, in *krm.Device) *pb.Device {
	if in == nil {
		return nil
	}
	out := &pb.Device{}
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	out.NumId = direct.ValueOf(in.NumID)
	out.Credentials = direct.Slice_ToProto(mapCtx, in.Credentials, DeviceCredential_ToProto)
	out.LastHeartbeatTime = direct.StringTimestamp_ToProto(mapCtx, in.LastHeartbeatTime)
	out.LastEventTime = direct.StringTimestamp_ToProto(mapCtx, in.LastEventTime)
	out.LastStateTime = direct.StringTimestamp_ToProto(mapCtx, in.LastStateTime)
	out.LastConfigAckTime = direct.StringTimestamp_ToProto(mapCtx, in.LastConfigAckTime)
	out.LastConfigSendTime = direct.StringTimestamp_ToProto(mapCtx, in.LastConfigSendTime)
	out.Blocked = direct.ValueOf(in.Blocked)
	out.LastErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.LastErrorTime)
	out.LastErrorStatus = Status_ToProto(mapCtx, in.LastErrorStatus)
	out.Config = DeviceConfig_ToProto(mapCtx, in.Config)
	out.State = DeviceState_ToProto(mapCtx, in.State)
	out.LogLevel = direct.Enum_ToProto[pb.LogLevel](mapCtx, in.LogLevel)
	out.Metadata = in.Metadata
	out.GatewayConfig = GatewayConfig_ToProto(mapCtx, in.GatewayConfig)
	return out
}
func DeviceConfig_FromProto(mapCtx *direct.MapContext, in *pb.DeviceConfig) *krm.DeviceConfig {
	if in == nil {
		return nil
	}
	out := &krm.DeviceConfig{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.CloudUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCloudUpdateTime())
	out.DeviceAckTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeviceAckTime())
	out.BinaryData = in.GetBinaryData()
	return out
}
func DeviceConfig_ToProto(mapCtx *direct.MapContext, in *krm.DeviceConfig) *pb.DeviceConfig {
	if in == nil {
		return nil
	}
	out := &pb.DeviceConfig{}
	out.Version = direct.ValueOf(in.Version)
	out.CloudUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.CloudUpdateTime)
	out.DeviceAckTime = direct.StringTimestamp_ToProto(mapCtx, in.DeviceAckTime)
	out.BinaryData = in.BinaryData
	return out
}
func DeviceCredential_FromProto(mapCtx *direct.MapContext, in *pb.DeviceCredential) *krm.DeviceCredential {
	if in == nil {
		return nil
	}
	out := &krm.DeviceCredential{}
	out.PublicKey = PublicKeyCredential_FromProto(mapCtx, in.GetPublicKey())
	out.ExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpirationTime())
	return out
}
func DeviceCredential_ToProto(mapCtx *direct.MapContext, in *krm.DeviceCredential) *pb.DeviceCredential {
	if in == nil {
		return nil
	}
	out := &pb.DeviceCredential{}
	if oneof := PublicKeyCredential_ToProto(mapCtx, in.PublicKey); oneof != nil {
		out.Credential = &pb.DeviceCredential_PublicKey{PublicKey: oneof}
	}
	out.ExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpirationTime)
	return out
}
func DeviceState_FromProto(mapCtx *direct.MapContext, in *pb.DeviceState) *krm.DeviceState {
	if in == nil {
		return nil
	}
	out := &krm.DeviceState{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.BinaryData = in.GetBinaryData()
	return out
}
func DeviceState_ToProto(mapCtx *direct.MapContext, in *krm.DeviceState) *pb.DeviceState {
	if in == nil {
		return nil
	}
	out := &pb.DeviceState{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.BinaryData = in.BinaryData
	return out
}
func GatewayConfig_FromProto(mapCtx *direct.MapContext, in *pb.GatewayConfig) *krm.GatewayConfig {
	if in == nil {
		return nil
	}
	out := &krm.GatewayConfig{}
	out.GatewayType = direct.Enum_FromProto(mapCtx, in.GetGatewayType())
	out.GatewayAuthMethod = direct.Enum_FromProto(mapCtx, in.GetGatewayAuthMethod())
	out.LastAccessedGatewayID = direct.LazyPtr(in.GetLastAccessedGatewayId())
	out.LastAccessedGatewayTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastAccessedGatewayTime())
	return out
}
func GatewayConfig_ToProto(mapCtx *direct.MapContext, in *krm.GatewayConfig) *pb.GatewayConfig {
	if in == nil {
		return nil
	}
	out := &pb.GatewayConfig{}
	out.GatewayType = direct.Enum_ToProto[pb.GatewayType](mapCtx, in.GatewayType)
	out.GatewayAuthMethod = direct.Enum_ToProto[pb.GatewayAuthMethod](mapCtx, in.GatewayAuthMethod)
	out.LastAccessedGatewayId = direct.ValueOf(in.LastAccessedGatewayID)
	out.LastAccessedGatewayTime = direct.StringTimestamp_ToProto(mapCtx, in.LastAccessedGatewayTime)
	return out
}
func IotDeviceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Device) *krm.IotDeviceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IotDeviceObservedState{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: NumID
	// MISSING: Credentials
	// MISSING: LastHeartbeatTime
	// MISSING: LastEventTime
	// MISSING: LastStateTime
	// MISSING: LastConfigAckTime
	// MISSING: LastConfigSendTime
	// MISSING: Blocked
	// MISSING: LastErrorTime
	// MISSING: LastErrorStatus
	// MISSING: Config
	// MISSING: State
	// MISSING: LogLevel
	// MISSING: Metadata
	// MISSING: GatewayConfig
	return out
}
func IotDeviceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IotDeviceObservedState) *pb.Device {
	if in == nil {
		return nil
	}
	out := &pb.Device{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: NumID
	// MISSING: Credentials
	// MISSING: LastHeartbeatTime
	// MISSING: LastEventTime
	// MISSING: LastStateTime
	// MISSING: LastConfigAckTime
	// MISSING: LastConfigSendTime
	// MISSING: Blocked
	// MISSING: LastErrorTime
	// MISSING: LastErrorStatus
	// MISSING: Config
	// MISSING: State
	// MISSING: LogLevel
	// MISSING: Metadata
	// MISSING: GatewayConfig
	return out
}
func IotDeviceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Device) *krm.IotDeviceSpec {
	if in == nil {
		return nil
	}
	out := &krm.IotDeviceSpec{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: NumID
	// MISSING: Credentials
	// MISSING: LastHeartbeatTime
	// MISSING: LastEventTime
	// MISSING: LastStateTime
	// MISSING: LastConfigAckTime
	// MISSING: LastConfigSendTime
	// MISSING: Blocked
	// MISSING: LastErrorTime
	// MISSING: LastErrorStatus
	// MISSING: Config
	// MISSING: State
	// MISSING: LogLevel
	// MISSING: Metadata
	// MISSING: GatewayConfig
	return out
}
func IotDeviceSpec_ToProto(mapCtx *direct.MapContext, in *krm.IotDeviceSpec) *pb.Device {
	if in == nil {
		return nil
	}
	out := &pb.Device{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: NumID
	// MISSING: Credentials
	// MISSING: LastHeartbeatTime
	// MISSING: LastEventTime
	// MISSING: LastStateTime
	// MISSING: LastConfigAckTime
	// MISSING: LastConfigSendTime
	// MISSING: Blocked
	// MISSING: LastErrorTime
	// MISSING: LastErrorStatus
	// MISSING: Config
	// MISSING: State
	// MISSING: LogLevel
	// MISSING: Metadata
	// MISSING: GatewayConfig
	return out
}
func PublicKeyCredential_FromProto(mapCtx *direct.MapContext, in *pb.PublicKeyCredential) *krm.PublicKeyCredential {
	if in == nil {
		return nil
	}
	out := &krm.PublicKeyCredential{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Key = direct.LazyPtr(in.GetKey())
	return out
}
func PublicKeyCredential_ToProto(mapCtx *direct.MapContext, in *krm.PublicKeyCredential) *pb.PublicKeyCredential {
	if in == nil {
		return nil
	}
	out := &pb.PublicKeyCredential{}
	out.Format = direct.Enum_ToProto[pb.PublicKeyFormat](mapCtx, in.Format)
	out.Key = direct.ValueOf(in.Key)
	return out
}
