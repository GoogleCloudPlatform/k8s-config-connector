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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoscalingPolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingPolicy{}
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: Name
	out.BasicAlgorithm = BasicAutoscalingAlgorithm_FromProto(mapCtx, in.GetBasicAlgorithm())
	out.WorkerConfig = InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx, in.GetWorkerConfig())
	out.SecondaryWorkerConfig = InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx, in.GetSecondaryWorkerConfig())
	out.Labels = in.Labels
	return out
}
func AutoscalingPolicy_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingPolicy) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out.Id = direct.ValueOf(in.ID)
	// MISSING: Name
	if oneof := BasicAutoscalingAlgorithm_ToProto(mapCtx, in.BasicAlgorithm); oneof != nil {
		out.Algorithm = &pb.AutoscalingPolicy_BasicAlgorithm{BasicAlgorithm: oneof}
	}
	out.WorkerConfig = InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx, in.WorkerConfig)
	out.SecondaryWorkerConfig = InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx, in.SecondaryWorkerConfig)
	out.Labels = in.Labels
	return out
}
func AutoscalingPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.AutoscalingPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingPolicyObservedState{}
	// MISSING: ID
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func AutoscalingPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingPolicyObservedState) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	// MISSING: ID
	out.Name = direct.ValueOf(in.Name)
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func BasicAutoscalingAlgorithm_FromProto(mapCtx *direct.MapContext, in *pb.BasicAutoscalingAlgorithm) *krm.BasicAutoscalingAlgorithm {
	if in == nil {
		return nil
	}
	out := &krm.BasicAutoscalingAlgorithm{}
	out.YarnConfig = BasicYarnAutoscalingConfig_FromProto(mapCtx, in.GetYarnConfig())
	out.CooldownPeriod = direct.StringDuration_FromProto(mapCtx, in.GetCooldownPeriod())
	return out
}
func BasicAutoscalingAlgorithm_ToProto(mapCtx *direct.MapContext, in *krm.BasicAutoscalingAlgorithm) *pb.BasicAutoscalingAlgorithm {
	if in == nil {
		return nil
	}
	out := &pb.BasicAutoscalingAlgorithm{}
	if oneof := BasicYarnAutoscalingConfig_ToProto(mapCtx, in.YarnConfig); oneof != nil {
		out.Config = &pb.BasicAutoscalingAlgorithm_YarnConfig{YarnConfig: oneof}
	}
	out.CooldownPeriod = direct.StringDuration_ToProto(mapCtx, in.CooldownPeriod)
	return out
}
func BasicYarnAutoscalingConfig_FromProto(mapCtx *direct.MapContext, in *pb.BasicYarnAutoscalingConfig) *krm.BasicYarnAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &krm.BasicYarnAutoscalingConfig{}
	out.GracefulDecommissionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetGracefulDecommissionTimeout())
	out.ScaleUpFactor = direct.LazyPtr(in.GetScaleUpFactor())
	out.ScaleDownFactor = direct.LazyPtr(in.GetScaleDownFactor())
	out.ScaleUpMinWorkerFraction = direct.LazyPtr(in.GetScaleUpMinWorkerFraction())
	out.ScaleDownMinWorkerFraction = direct.LazyPtr(in.GetScaleDownMinWorkerFraction())
	return out
}
func BasicYarnAutoscalingConfig_ToProto(mapCtx *direct.MapContext, in *krm.BasicYarnAutoscalingConfig) *pb.BasicYarnAutoscalingConfig {
	if in == nil {
		return nil
	}
	out := &pb.BasicYarnAutoscalingConfig{}
	out.GracefulDecommissionTimeout = direct.StringDuration_ToProto(mapCtx, in.GracefulDecommissionTimeout)
	out.ScaleUpFactor = direct.ValueOf(in.ScaleUpFactor)
	out.ScaleDownFactor = direct.ValueOf(in.ScaleDownFactor)
	out.ScaleUpMinWorkerFraction = direct.ValueOf(in.ScaleUpMinWorkerFraction)
	out.ScaleDownMinWorkerFraction = direct.ValueOf(in.ScaleDownMinWorkerFraction)
	return out
}
func DataprocAutoscalingPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.DataprocAutoscalingPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocAutoscalingPolicyObservedState{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func DataprocAutoscalingPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocAutoscalingPolicyObservedState) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func DataprocAutoscalingPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.DataprocAutoscalingPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocAutoscalingPolicySpec{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func DataprocAutoscalingPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocAutoscalingPolicySpec) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	// MISSING: ID
	// MISSING: Name
	// MISSING: BasicAlgorithm
	// MISSING: WorkerConfig
	// MISSING: SecondaryWorkerConfig
	// MISSING: Labels
	return out
}
func InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx *direct.MapContext, in *pb.InstanceGroupAutoscalingPolicyConfig) *krm.InstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &krm.InstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = direct.LazyPtr(in.GetMinInstances())
	out.MaxInstances = direct.LazyPtr(in.GetMaxInstances())
	out.Weight = direct.LazyPtr(in.GetWeight())
	return out
}
func InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx *direct.MapContext, in *krm.InstanceGroupAutoscalingPolicyConfig) *pb.InstanceGroupAutoscalingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.InstanceGroupAutoscalingPolicyConfig{}
	out.MinInstances = direct.ValueOf(in.MinInstances)
	out.MaxInstances = direct.ValueOf(in.MaxInstances)
	out.Weight = direct.ValueOf(in.Weight)
	return out
}
