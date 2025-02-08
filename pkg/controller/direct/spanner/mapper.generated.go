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

package spanner

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
)
func AutoscalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig) *krm.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingConfig_AutoscalingLimits_FromProto(mapCtx, in.GetAutoscalingLimits())
	out.AutoscalingTargets = AutoscalingConfig_AutoscalingTargets_FromProto(mapCtx, in.GetAutoscalingTargets())
	// MISSING: AsymmetricAutoscalingOptions
	return out
}
func AutoscalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig) *pb.AutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig{}
	out.AutoscalingLimits = AutoscalingConfig_AutoscalingLimits_ToProto(mapCtx, in.AutoscalingLimits)
	out.AutoscalingTargets = AutoscalingConfig_AutoscalingTargets_ToProto(mapCtx, in.AutoscalingTargets)
	// MISSING: AsymmetricAutoscalingOptions
	return out
}
func AutoscalingConfig_AsymmetricAutoscalingOption_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig_AsymmetricAutoscalingOption) *krm.AutoscalingConfig_AsymmetricAutoscalingOption {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig_AsymmetricAutoscalingOption{}
	out.ReplicaSelection = ReplicaSelection_FromProto(mapCtx, in.GetReplicaSelection())
	out.Overrides = AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides_FromProto(mapCtx, in.GetOverrides())
	return out
}
func AutoscalingConfig_AsymmetricAutoscalingOption_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig_AsymmetricAutoscalingOption) *pb.AutoscalingConfig_AsymmetricAutoscalingOption {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig_AsymmetricAutoscalingOption{}
	out.ReplicaSelection = ReplicaSelection_ToProto(mapCtx, in.ReplicaSelection)
	out.Overrides = AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides_ToProto(mapCtx, in.Overrides)
	return out
}
func AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides) *krm.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides{}
	out.AutoscalingLimits = AutoscalingConfig_AutoscalingLimits_FromProto(mapCtx, in.GetAutoscalingLimits())
	out.AutoscalingTargetHighPriorityCpuUtilizationPercent = direct.LazyPtr(in.GetAutoscalingTargetHighPriorityCpuUtilizationPercent())
	return out
}
func AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides) *pb.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides{}
	out.AutoscalingLimits = AutoscalingConfig_AutoscalingLimits_ToProto(mapCtx, in.AutoscalingLimits)
	out.AutoscalingTargetHighPriorityCpuUtilizationPercent = direct.ValueOf(in.AutoscalingTargetHighPriorityCpuUtilizationPercent)
	return out
}
func AutoscalingConfig_AutoscalingLimits_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig_AutoscalingLimits) *krm.AutoscalingConfig_AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig_AutoscalingLimits{}
	out.MinNodes = direct.LazyPtr(in.GetMinNodes())
	out.MinProcessingUnits = direct.LazyPtr(in.GetMinProcessingUnits())
	out.MaxNodes = direct.LazyPtr(in.GetMaxNodes())
	out.MaxProcessingUnits = direct.LazyPtr(in.GetMaxProcessingUnits())
	return out
}
func AutoscalingConfig_AutoscalingLimits_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig_AutoscalingLimits) *pb.AutoscalingConfig_AutoscalingLimits {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig_AutoscalingLimits{}
	if oneof := AutoscalingConfig_AutoscalingLimits_MinNodes_ToProto(mapCtx, in.MinNodes); oneof != nil {
		out.MinLimit = oneof
	}
	if oneof := AutoscalingConfig_AutoscalingLimits_MinProcessingUnits_ToProto(mapCtx, in.MinProcessingUnits); oneof != nil {
		out.MinLimit = oneof
	}
	if oneof := AutoscalingConfig_AutoscalingLimits_MaxNodes_ToProto(mapCtx, in.MaxNodes); oneof != nil {
		out.MaxLimit = oneof
	}
	if oneof := AutoscalingConfig_AutoscalingLimits_MaxProcessingUnits_ToProto(mapCtx, in.MaxProcessingUnits); oneof != nil {
		out.MaxLimit = oneof
	}
	return out
}
func AutoscalingConfig_AutoscalingTargets_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingConfig_AutoscalingTargets) *krm.AutoscalingConfig_AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingConfig_AutoscalingTargets{}
	out.HighPriorityCpuUtilizationPercent = direct.LazyPtr(in.GetHighPriorityCpuUtilizationPercent())
	out.StorageUtilizationPercent = direct.LazyPtr(in.GetStorageUtilizationPercent())
	return out
}
func AutoscalingConfig_AutoscalingTargets_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingConfig_AutoscalingTargets) *pb.AutoscalingConfig_AutoscalingTargets {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingConfig_AutoscalingTargets{}
	out.HighPriorityCpuUtilizationPercent = direct.ValueOf(in.HighPriorityCpuUtilizationPercent)
	out.StorageUtilizationPercent = direct.ValueOf(in.StorageUtilizationPercent)
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Config = direct.LazyPtr(in.GetConfig())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.ProcessingUnits = direct.LazyPtr(in.GetProcessingUnits())
	out.ReplicaComputeCapacity = direct.Slice_FromProto(mapCtx, in.ReplicaComputeCapacity, ReplicaComputeCapacity_FromProto)
	out.AutoscalingConfig = AutoscalingConfig_FromProto(mapCtx, in.GetAutoscalingConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Labels = in.Labels
	// MISSING: InstanceType
	out.EndpointUris = in.EndpointUris
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: FreeInstanceMetadata
	out.Edition = direct.Enum_FromProto(mapCtx, in.GetEdition())
	out.DefaultBackupScheduleType = direct.Enum_FromProto(mapCtx, in.GetDefaultBackupScheduleType())
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.Config = direct.ValueOf(in.Config)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.ProcessingUnits = direct.ValueOf(in.ProcessingUnits)
	out.ReplicaComputeCapacity = direct.Slice_ToProto(mapCtx, in.ReplicaComputeCapacity, ReplicaComputeCapacity_ToProto)
	out.AutoscalingConfig = AutoscalingConfig_ToProto(mapCtx, in.AutoscalingConfig)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.Labels = in.Labels
	// MISSING: InstanceType
	out.EndpointUris = in.EndpointUris
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: FreeInstanceMetadata
	out.Edition = direct.Enum_ToProto[pb.Instance_Edition](mapCtx, in.Edition)
	out.DefaultBackupScheduleType = direct.Enum_ToProto[pb.Instance_DefaultBackupScheduleType](mapCtx, in.DefaultBackupScheduleType)
	return out
}
func InstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: ConfigType
	out.Replicas = direct.Slice_FromProto(mapCtx, in.Replicas, ReplicaInfo_FromProto)
	// MISSING: OptionalReplicas
	out.BaseConfig = direct.LazyPtr(in.GetBaseConfig())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.LeaderOptions = in.LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
func InstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfig) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: ConfigType
	out.Replicas = direct.Slice_ToProto(mapCtx, in.Replicas, ReplicaInfo_ToProto)
	// MISSING: OptionalReplicas
	out.BaseConfig = direct.ValueOf(in.BaseConfig)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.LeaderOptions = in.LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
func InstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.InstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.ConfigType = direct.Enum_FromProto(mapCtx, in.GetConfigType())
	// MISSING: Replicas
	out.OptionalReplicas = direct.Slice_FromProto(mapCtx, in.OptionalReplicas, ReplicaInfo_FromProto)
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.FreeInstanceAvailability = direct.Enum_FromProto(mapCtx, in.GetFreeInstanceAvailability())
	out.QuorumType = direct.Enum_FromProto(mapCtx, in.GetQuorumType())
	out.StorageLimitPerProcessingUnit = direct.LazyPtr(in.GetStorageLimitPerProcessingUnit())
	return out
}
func InstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	out.ConfigType = direct.Enum_ToProto[pb.InstanceConfig_Type](mapCtx, in.ConfigType)
	// MISSING: Replicas
	out.OptionalReplicas = direct.Slice_ToProto(mapCtx, in.OptionalReplicas, ReplicaInfo_ToProto)
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.State = direct.Enum_ToProto[pb.InstanceConfig_State](mapCtx, in.State)
	out.FreeInstanceAvailability = direct.Enum_ToProto[pb.InstanceConfig_FreeInstanceAvailability](mapCtx, in.FreeInstanceAvailability)
	out.QuorumType = direct.Enum_ToProto[pb.InstanceConfig_QuorumType](mapCtx, in.QuorumType)
	out.StorageLimitPerProcessingUnit = direct.ValueOf(in.StorageLimitPerProcessingUnit)
	return out
}
func ReplicaComputeCapacity_FromProto(mapCtx *direct.MapContext, in *pb.ReplicaComputeCapacity) *krm.ReplicaComputeCapacity {
	if in == nil {
		return nil
	}
	out := &krm.ReplicaComputeCapacity{}
	out.ReplicaSelection = ReplicaSelection_FromProto(mapCtx, in.GetReplicaSelection())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.ProcessingUnits = direct.LazyPtr(in.GetProcessingUnits())
	return out
}
func ReplicaComputeCapacity_ToProto(mapCtx *direct.MapContext, in *krm.ReplicaComputeCapacity) *pb.ReplicaComputeCapacity {
	if in == nil {
		return nil
	}
	out := &pb.ReplicaComputeCapacity{}
	out.ReplicaSelection = ReplicaSelection_ToProto(mapCtx, in.ReplicaSelection)
	if oneof := ReplicaComputeCapacity_NodeCount_ToProto(mapCtx, in.NodeCount); oneof != nil {
		out.ComputeCapacity = oneof
	}
	if oneof := ReplicaComputeCapacity_ProcessingUnits_ToProto(mapCtx, in.ProcessingUnits); oneof != nil {
		out.ComputeCapacity = oneof
	}
	return out
}
func ReplicaInfo_FromProto(mapCtx *direct.MapContext, in *pb.ReplicaInfo) *krm.ReplicaInfo {
	if in == nil {
		return nil
	}
	out := &krm.ReplicaInfo{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DefaultLeaderLocation = direct.LazyPtr(in.GetDefaultLeaderLocation())
	return out
}
func ReplicaInfo_ToProto(mapCtx *direct.MapContext, in *krm.ReplicaInfo) *pb.ReplicaInfo {
	if in == nil {
		return nil
	}
	out := &pb.ReplicaInfo{}
	out.Location = direct.ValueOf(in.Location)
	out.Type = direct.Enum_ToProto[pb.ReplicaInfo_ReplicaType](mapCtx, in.Type)
	out.DefaultLeaderLocation = direct.ValueOf(in.DefaultLeaderLocation)
	return out
}
func ReplicaSelection_FromProto(mapCtx *direct.MapContext, in *pb.ReplicaSelection) *krm.ReplicaSelection {
	if in == nil {
		return nil
	}
	out := &krm.ReplicaSelection{}
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func ReplicaSelection_ToProto(mapCtx *direct.MapContext, in *krm.ReplicaSelection) *pb.ReplicaSelection {
	if in == nil {
		return nil
	}
	out := &pb.ReplicaSelection{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func SpannerInstanceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.SpannerInstanceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ConfigType
	// MISSING: Replicas
	// MISSING: OptionalReplicas
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
func SpannerInstanceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpannerInstanceConfigObservedState) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ConfigType
	// MISSING: Replicas
	// MISSING: OptionalReplicas
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
func SpannerInstanceConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.InstanceConfig) *krm.SpannerInstanceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceConfigSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ConfigType
	// MISSING: Replicas
	// MISSING: OptionalReplicas
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
func SpannerInstanceConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpannerInstanceConfigSpec) *pb.InstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ConfigType
	// MISSING: Replicas
	// MISSING: OptionalReplicas
	// MISSING: BaseConfig
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: LeaderOptions
	// MISSING: Reconciling
	// MISSING: State
	// MISSING: FreeInstanceAvailability
	// MISSING: QuorumType
	// MISSING: StorageLimitPerProcessingUnit
	return out
}
