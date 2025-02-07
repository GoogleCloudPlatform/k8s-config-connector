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

package apigeeregistry

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigeeregistry/v1alpha1"
)
func ApigeeregistryInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.ApigeeregistryInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApigeeregistryInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func ApigeeregistryInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApigeeregistryInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	out.Config = Instance_Config_FromProto(mapCtx, in.GetConfig())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	out.Config = Instance_Config_ToProto(mapCtx, in.Config)
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.Config = Instance_ConfigObservedState_FromProto(mapCtx, in.GetConfig())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.Config = Instance_ConfigObservedState_ToProto(mapCtx, in.Config)
	return out
}
func Instance_Config_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Config) *krm.Instance_Config {
	if in == nil {
		return nil
	}
	out := &krm.Instance_Config{}
	// MISSING: Location
	out.CmekKeyName = direct.LazyPtr(in.GetCmekKeyName())
	return out
}
func Instance_Config_ToProto(mapCtx *direct.MapContext, in *krm.Instance_Config) *pb.Instance_Config {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Config{}
	// MISSING: Location
	out.CmekKeyName = direct.ValueOf(in.CmekKeyName)
	return out
}
func Instance_ConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance_Config) *krm.Instance_ConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Instance_ConfigObservedState{}
	out.Location = direct.LazyPtr(in.GetLocation())
	// MISSING: CmekKeyName
	return out
}
func Instance_ConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Instance_ConfigObservedState) *pb.Instance_Config {
	if in == nil {
		return nil
	}
	out := &pb.Instance_Config{}
	out.Location = direct.ValueOf(in.Location)
	// MISSING: CmekKeyName
	return out
}
