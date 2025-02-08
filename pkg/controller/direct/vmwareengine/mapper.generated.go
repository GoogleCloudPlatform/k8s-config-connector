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

package vmwareengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutoscalingSettings_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingSettings) *krm.AutoscalingSettings {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingSettings{}
	// MISSING: AutoscalingPolicies
	out.MinClusterNodeCount = direct.LazyPtr(in.GetMinClusterNodeCount())
	out.MaxClusterNodeCount = direct.LazyPtr(in.GetMaxClusterNodeCount())
	out.CoolDownPeriod = direct.StringDuration_FromProto(mapCtx, in.GetCoolDownPeriod())
	return out
}
func AutoscalingSettings_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingSettings) *pb.AutoscalingSettings {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingSettings{}
	// MISSING: AutoscalingPolicies
	out.MinClusterNodeCount = direct.ValueOf(in.MinClusterNodeCount)
	out.MaxClusterNodeCount = direct.ValueOf(in.MaxClusterNodeCount)
	out.CoolDownPeriod = direct.StringDuration_ToProto(mapCtx, in.CoolDownPeriod)
	return out
}
func AutoscalingSettings_AutoscalingPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingSettings_AutoscalingPolicy) *krm.AutoscalingSettings_AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingSettings_AutoscalingPolicy{}
	out.NodeTypeID = direct.LazyPtr(in.GetNodeTypeId())
	out.ScaleOutSize = direct.LazyPtr(in.GetScaleOutSize())
	out.CpuThresholds = AutoscalingSettings_Thresholds_FromProto(mapCtx, in.GetCpuThresholds())
	out.GrantedMemoryThresholds = AutoscalingSettings_Thresholds_FromProto(mapCtx, in.GetGrantedMemoryThresholds())
	out.ConsumedMemoryThresholds = AutoscalingSettings_Thresholds_FromProto(mapCtx, in.GetConsumedMemoryThresholds())
	out.StorageThresholds = AutoscalingSettings_Thresholds_FromProto(mapCtx, in.GetStorageThresholds())
	return out
}
func AutoscalingSettings_AutoscalingPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingSettings_AutoscalingPolicy) *pb.AutoscalingSettings_AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingSettings_AutoscalingPolicy{}
	out.NodeTypeId = direct.ValueOf(in.NodeTypeID)
	out.ScaleOutSize = direct.ValueOf(in.ScaleOutSize)
	out.CpuThresholds = AutoscalingSettings_Thresholds_ToProto(mapCtx, in.CpuThresholds)
	out.GrantedMemoryThresholds = AutoscalingSettings_Thresholds_ToProto(mapCtx, in.GrantedMemoryThresholds)
	out.ConsumedMemoryThresholds = AutoscalingSettings_Thresholds_ToProto(mapCtx, in.ConsumedMemoryThresholds)
	out.StorageThresholds = AutoscalingSettings_Thresholds_ToProto(mapCtx, in.StorageThresholds)
	return out
}
func AutoscalingSettings_Thresholds_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingSettings_Thresholds) *krm.AutoscalingSettings_Thresholds {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingSettings_Thresholds{}
	out.ScaleOut = direct.LazyPtr(in.GetScaleOut())
	out.ScaleIn = direct.LazyPtr(in.GetScaleIn())
	return out
}
func AutoscalingSettings_Thresholds_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingSettings_Thresholds) *pb.AutoscalingSettings_Thresholds {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingSettings_Thresholds{}
	out.ScaleOut = direct.ValueOf(in.ScaleOut)
	out.ScaleIn = direct.ValueOf(in.ScaleIn)
	return out
}
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	out.AutoscalingSettings = AutoscalingSettings_FromProto(mapCtx, in.GetAutoscalingSettings())
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	out.StretchedClusterConfig = StretchedClusterConfig_FromProto(mapCtx, in.GetStretchedClusterConfig())
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	out.AutoscalingSettings = AutoscalingSettings_ToProto(mapCtx, in.AutoscalingSettings)
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	out.StretchedClusterConfig = StretchedClusterConfig_ToProto(mapCtx, in.StretchedClusterConfig)
	return out
}
func ClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClusterObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Management = direct.LazyPtr(in.GetManagement())
	// MISSING: AutoscalingSettings
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func ClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.Management = direct.ValueOf(in.Management)
	// MISSING: AutoscalingSettings
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func NodeTypeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeTypeConfig) *krm.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeTypeConfig{}
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.CustomCoreCount = direct.LazyPtr(in.GetCustomCoreCount())
	return out
}
func NodeTypeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeTypeConfig) *pb.NodeTypeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeTypeConfig{}
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.CustomCoreCount = direct.ValueOf(in.CustomCoreCount)
	return out
}
func StretchedClusterConfig_FromProto(mapCtx *direct.MapContext, in *pb.StretchedClusterConfig) *krm.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &krm.StretchedClusterConfig{}
	out.PreferredLocation = direct.LazyPtr(in.GetPreferredLocation())
	out.SecondaryLocation = direct.LazyPtr(in.GetSecondaryLocation())
	return out
}
func StretchedClusterConfig_ToProto(mapCtx *direct.MapContext, in *krm.StretchedClusterConfig) *pb.StretchedClusterConfig {
	if in == nil {
		return nil
	}
	out := &pb.StretchedClusterConfig{}
	out.PreferredLocation = direct.ValueOf(in.PreferredLocation)
	out.SecondaryLocation = direct.ValueOf(in.SecondaryLocation)
	return out
}
func VmwareengineClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.VmwareengineClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineClusterObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	// MISSING: AutoscalingSettings
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func VmwareengineClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	// MISSING: AutoscalingSettings
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func VmwareengineClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.VmwareengineClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineClusterSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	// MISSING: AutoscalingSettings
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
func VmwareengineClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Management
	// MISSING: AutoscalingSettings
	// MISSING: Uid
	// MISSING: NodeTypeConfigs
	// MISSING: StretchedClusterConfig
	return out
}
