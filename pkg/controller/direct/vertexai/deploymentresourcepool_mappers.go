// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License.

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoscalingMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingMetricSpec) *krm.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingMetricSpec{}
	out.MetricName = direct.LazyPtr(in.GetMetricName())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}
func AutoscalingMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingMetricSpec) *pb.AutoscalingMetricSpec {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingMetricSpec{}
	out.MetricName = direct.ValueOf(in.MetricName)
	out.Target = direct.ValueOf(in.Target)
	return out
}
func FlexStart_FromProto(mapCtx *direct.MapContext, in *pb.FlexStart) *krm.FlexStart {
	if in == nil {
		return nil
	}
	out := &krm.FlexStart{}
	out.MaxRuntimeDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRuntimeDuration())
	return out
}
func FlexStart_ToProto(mapCtx *direct.MapContext, in *krm.FlexStart) *pb.FlexStart {
	if in == nil {
		return nil
	}
	out := &pb.FlexStart{}
	out.MaxRuntimeDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRuntimeDuration)
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_FromProto(mapCtx, in.GetReservationAffinityType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ReservationAffinityType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func MachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.MachineSpec) *krm.MachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.MachineSpec{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.GpuPartitionSize = direct.LazyPtr(in.GetGpuPartitionSize())
	out.MultihostGpuNodeCount = direct.LazyPtr(in.GetMultihostGpuNodeCount())
	return out
}
func MachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.MachineSpec) *pb.MachineSpec {
	if in == nil {
		return nil
	}
	out := &pb.MachineSpec{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorType = direct.Enum_ToProto[pb.AcceleratorType](mapCtx, in.AcceleratorType)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.GpuPartitionSize = direct.ValueOf(in.GpuPartitionSize)
	out.MultihostGpuNodeCount = direct.ValueOf(in.MultihostGpuNodeCount)
	return out
}
func DedicatedResources_FromProto(mapCtx *direct.MapContext, in *pb.DedicatedResources) *krm.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &krm.DedicatedResources{}
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.MinReplicaCount = direct.LazyPtr(in.GetMinReplicaCount())
	out.MaxReplicaCount = direct.LazyPtr(in.GetMaxReplicaCount())
	out.RequiredReplicaCount = direct.LazyPtr(in.GetRequiredReplicaCount())
	out.AutoscalingMetricSpecs = direct.Slice_FromProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_FromProto)
	out.Spot = direct.LazyPtr(in.GetSpot())
	out.FlexStart = FlexStart_FromProto(mapCtx, in.GetFlexStart())
	return out
}
func DedicatedResources_ToProto(mapCtx *direct.MapContext, in *krm.DedicatedResources) *pb.DedicatedResources {
	if in == nil {
		return nil
	}
	out := &pb.DedicatedResources{}
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.MinReplicaCount = direct.ValueOf(in.MinReplicaCount)
	out.MaxReplicaCount = direct.ValueOf(in.MaxReplicaCount)
	out.RequiredReplicaCount = direct.ValueOf(in.RequiredReplicaCount)
	out.AutoscalingMetricSpecs = direct.Slice_ToProto(mapCtx, in.AutoscalingMetricSpecs, AutoscalingMetricSpec_ToProto)
	out.Spot = direct.ValueOf(in.Spot)
	out.FlexStart = FlexStart_ToProto(mapCtx, in.FlexStart)
	return out
}
func VertexAIDeploymentResourcePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentResourcePool) *krm.VertexAIDeploymentResourcePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIDeploymentResourcePoolObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func VertexAIDeploymentResourcePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDeploymentResourcePoolObservedState) *pb.DeploymentResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentResourcePool{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func VertexAIDeploymentResourcePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentResourcePool) *krm.VertexAIDeploymentResourcePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIDeploymentResourcePoolSpec{}
	out.DedicatedResources = DedicatedResources_FromProto(mapCtx, in.GetDedicatedResources())
	out.EncryptionSpec = EncryptionSpec_v1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.DisableContainerLogging = direct.LazyPtr(in.GetDisableContainerLogging())
	return out
}
func VertexAIDeploymentResourcePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDeploymentResourcePoolSpec) *pb.DeploymentResourcePool {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentResourcePool{}
	out.DedicatedResources = DedicatedResources_ToProto(mapCtx, in.DedicatedResources)
	out.EncryptionSpec = EncryptionSpec_v1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.DisableContainerLogging = direct.ValueOf(in.DisableContainerLogging)
	return out
}
