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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/tpu/v2"
)

func AcceleratorConfig_FromProto(mapCtx *direct.MapContext, in *pb.AcceleratorConfig) *krm.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &krm.AcceleratorConfig{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Topology = direct.LazyPtr(in.GetTopology())
	return out
}
func AcceleratorConfig_ToProto(mapCtx *direct.MapContext, in *krm.AcceleratorConfig) *pb.AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := &pb.AcceleratorConfig{}
	out.Type = direct.Enum_ToProto[pb.AcceleratorConfig_Type](mapCtx, in.Type)
	out.Topology = direct.ValueOf(in.Topology)
	return out
}
func AccessConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessConfig) *krm.AccessConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessConfigObservedState{}
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	return out
}
func AccessConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessConfigObservedState) *pb.AccessConfig {
	if in == nil {
		return nil
	}
	out := &pb.AccessConfig{}
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	return out
}
func AttachedDisk_FromProto(mapCtx *direct.MapContext, in *pb.AttachedDisk) *krm.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &krm.AttachedDisk{}
	out.SourceDisk = direct.LazyPtr(in.GetSourceDisk())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func AttachedDisk_ToProto(mapCtx *direct.MapContext, in *krm.AttachedDisk) *pb.AttachedDisk {
	if in == nil {
		return nil
	}
	out := &pb.AttachedDisk{}
	out.SourceDisk = direct.ValueOf(in.SourceDisk)
	out.Mode = direct.Enum_ToProto[pb.AttachedDisk_DiskMode](mapCtx, in.Mode)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refs.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.EnableExternalIps = direct.LazyPtr(in.GetEnableExternalIps())
	out.CanIPForward = direct.LazyPtr(in.GetCanIpForward())
	out.QueueCount = direct.LazyPtr(in.GetQueueCount())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.EnableExternalIps = direct.ValueOf(in.EnableExternalIps)
	out.CanIpForward = direct.ValueOf(in.CanIPForward)
	out.QueueCount = direct.ValueOf(in.QueueCount)
	return out
}
func NetworkEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.NetworkEndpoint{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	// MISSING: AccessConfig
	return out
}
func NetworkEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.NetworkEndpoint) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Port = direct.ValueOf(in.Port)
	// MISSING: AccessConfig
	return out
}
func NetworkEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.NetworkEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkEndpointObservedState{}
	// MISSING: IPAddress
	// MISSING: Port
	out.AccessConfig = AccessConfigObservedState_FromProto(mapCtx, in.GetAccessConfig())
	return out
}
func NetworkEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkEndpointObservedState) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	// MISSING: IPAddress
	// MISSING: Port
	out.AccessConfig = AccessConfigObservedState_ToProto(mapCtx, in.AccessConfig)
	return out
}
func SchedulingConfig_FromProto(mapCtx *direct.MapContext, in *pb.SchedulingConfig) *krm.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &krm.SchedulingConfig{}
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	out.Reserved = direct.LazyPtr(in.GetReserved())
	out.Spot = direct.LazyPtr(in.GetSpot())
	return out
}
func SchedulingConfig_ToProto(mapCtx *direct.MapContext, in *krm.SchedulingConfig) *pb.SchedulingConfig {
	if in == nil {
		return nil
	}
	out := &pb.SchedulingConfig{}
	out.Preemptible = direct.ValueOf(in.Preemptible)
	out.Reserved = direct.ValueOf(in.Reserved)
	out.Spot = direct.ValueOf(in.Spot)
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
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
func TPUVirtualMachineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.TPUVirtualMachineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TPUVirtualMachineObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.HealthDescription = direct.LazyPtr(in.GetHealthDescription())
	// MISSING: CreateTime
	out.NetworkEndpoints = direct.Slice_FromProto(mapCtx, in.NetworkEndpoints, NetworkEndpointObservedState_FromProto)
	out.Health = direct.Enum_FromProto(mapCtx, in.GetHealth())
	// MISSING: Labels
	// MISSING: ID
	// MISSING: APIVersion
	out.Symptoms = direct.Slice_FromProto(mapCtx, in.Symptoms, Symptom_FromProto)
	out.QueuedResource = direct.LazyPtr(in.GetQueuedResource())
	out.MultisliceNode = direct.LazyPtr(in.GetMultisliceNode())
	return out
}
func TPUVirtualMachineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TPUVirtualMachineObservedState) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Node_State](mapCtx, in.State)
	out.HealthDescription = direct.ValueOf(in.HealthDescription)
	// MISSING: CreateTime
	out.NetworkEndpoints = direct.Slice_ToProto(mapCtx, in.NetworkEndpoints, NetworkEndpointObservedState_ToProto)
	out.Health = direct.Enum_ToProto[pb.Node_Health](mapCtx, in.Health)
	// MISSING: Labels
	// MISSING: ID
	// MISSING: APIVersion
	out.Symptoms = direct.Slice_ToProto(mapCtx, in.Symptoms, Symptom_ToProto)
	out.QueuedResource = direct.ValueOf(in.QueuedResource)
	out.MultisliceNode = direct.ValueOf(in.MultisliceNode)
	return out
}
func TPUVirtualMachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.Node) *krm.TPUVirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.TPUVirtualMachineSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AcceleratorType = direct.LazyPtr(in.GetAcceleratorType())
	out.RuntimeVersion = direct.LazyPtr(in.GetRuntimeVersion())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.NetworkConfigs = direct.Slice_FromProto(mapCtx, in.NetworkConfigs, NetworkConfig_FromProto)
	out.CIDRBlock = direct.LazyPtr(in.GetCidrBlock())
	out.ServiceAccount = ServiceAccount_FromProto(mapCtx, in.GetServiceAccount())
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_FromProto(mapCtx, in.GetSchedulingConfig())
	// MISSING: Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	// MISSING: ID
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, AttachedDisk_FromProto)
	// MISSING: APIVersion
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.AcceleratorConfig = AcceleratorConfig_FromProto(mapCtx, in.GetAcceleratorConfig())
	return out
}
func TPUVirtualMachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.TPUVirtualMachineSpec) *pb.Node {
	if in == nil {
		return nil
	}
	out := &pb.Node{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.AcceleratorType = direct.ValueOf(in.AcceleratorType)
	out.RuntimeVersion = direct.ValueOf(in.RuntimeVersion)
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.NetworkConfigs = direct.Slice_ToProto(mapCtx, in.NetworkConfigs, NetworkConfig_ToProto)
	out.CidrBlock = direct.ValueOf(in.CIDRBlock)
	out.ServiceAccount = ServiceAccount_ToProto(mapCtx, in.ServiceAccount)
	// MISSING: CreateTime
	out.SchedulingConfig = SchedulingConfig_ToProto(mapCtx, in.SchedulingConfig)
	// MISSING: Labels
	out.Metadata = in.Metadata
	out.Tags = in.Tags
	// MISSING: ID
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, AttachedDisk_ToProto)
	// MISSING: APIVersion
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.AcceleratorConfig = AcceleratorConfig_ToProto(mapCtx, in.AcceleratorConfig)
	return out
}
