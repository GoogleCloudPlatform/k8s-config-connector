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

package tpu

import (
	pb "cloud.google.com/go/tpu/apiv1/tpupb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.NetworkEndpoint{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func NetworkEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.NetworkEndpoint) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func SchedulingConfig_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingConfig) *krm.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingConfig{}
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.Reserved = direct.LazyPtr(in.GetReserved())
	return out
}
func SchedulingConfig_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingConfig) *pb.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingConfig{}
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Reserved = direct.ValueOf(in.Reserved)
	return out
}
func Symptom_FromProto(mapCtx *direct.MapContext, in *pb.Symptom) *krm.Symptom {
	if in == nil {
		return nil
	}
	out := &krm.Symptom{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SymptomType = direct.Enum_FromProto(mapCtx, in.GetSymptomType())
	out.Details = direct.LazyPtr(in.GetDetails())
	out.WorkerID = direct.LazyPtr(in.GetWorkerId())
	return out
}
func Symptom_ToProto(mapCtx *direct.MapContext, in *krm.Symptom) *pb.Symptom {
	if in == nil {
		return nil
	}
	out := &pb.Symptom{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.SymptomType = direct.Enum_ToProto[pb.Symptom_SymptomType](mapCtx, in.SymptomType)
	out.Details = direct.ValueOf(in.Details)
	out.WorkerId = direct.ValueOf(in.WorkerID)
	return out
}
func TPUNodeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.TPUNodeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TPUNodeObservedState{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: Port
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.HealthDescription = direct.LazyPtr(in.GetHealthDescription())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: CreateTime
	out.NetworkEndpoints = direct.Slice_FromProto(mapCtx, in.NetworkEndpoints, NetworkEndpoint_FromProto)
	out.Health = direct.Enum_FromProto(mapCtx, in.GetHealth())
	// MISSING: Labels
	out.APIVersion = direct.Enum_FromProto(mapCtx, in.GetApiVersion())
	out.Symptoms = direct.Slice_FromProto(mapCtx, in.Symptoms, Symptom_FromProto)
	return out
}
func TPUNodeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TPUNodeObservedState) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	// MISSING: IPAddress
	// MISSING: Port
	out.State = direct.Enum_ToProto[pb.Node_State](mapCtx, in.State)
	out.HealthDescription = direct.ValueOf(in.HealthDescription)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: CreateTime
	out.NetworkEndpoints = direct.Slice_ToProto(mapCtx, in.NetworkEndpoints, NetworkEndpoint_ToProto)
	out.Health = direct.Enum_ToProto[pb.Node_Health](mapCtx, in.Health)
	// MISSING: Labels
	out.ApiVersion = direct.Enum_ToProto[pb.Node_ApiVersion](mapCtx, in.APIVersion)
	out.Symptoms = direct.Slice_ToProto(mapCtx, in.Symptoms, Symptom_ToProto)
	return out
}
func TPUNodeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.TPUNodeSpec {
	if in == nil {
		return nil
	}
	out := &krm.TPUNodeSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	// MISSING: IPAddress
	// MISSING: Port
	out.TensorflowVersion = direct.LazyPtr(in.GetTensorflowVersion())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.CIDRBlock = direct.LazyPtr(in.GetCidrBlock())
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_FromProto(mapCtx, in.GetSchedulingConfig())
	// MISSING: Labels
	out.UseServiceNetworking = direct.LazyPtr(in.GetUseServiceNetworking())
	return out
}
func TPUNodeSpec_ToProto(mapCtx *direct.MapContext, in *krm.TPUNodeSpec) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	// MISSING: IPAddress
	// MISSING: Port
	out.TensorflowVersion = direct.ValueOf(in.TensorflowVersion)
	out.Network = direct.ValueOf(in.Network)
	out.CidrBlock = direct.ValueOf(in.CIDRBlock)
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_ToProto(mapCtx, in.SchedulingConfig)
	// MISSING: Labels
	out.UseServiceNetworking = direct.ValueOf(in.UseServiceNetworking)
	return out
}
