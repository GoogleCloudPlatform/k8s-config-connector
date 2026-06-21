// Copyright 2026 Google LLC
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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeRegionAutoscalerSpec_v1alpha1_FromProto maps a pb.Autoscaler to a krm.ComputeRegionAutoscalerSpec.
func ComputeRegionAutoscalerSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Autoscaler) *krm.ComputeRegionAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionAutoscalerSpec{}
	if in.AutoscalingPolicy != nil {
		policy := RegionautoscalerAutoscalingPolicy_v1alpha1_FromProto(mapCtx, in.AutoscalingPolicy)
		if policy != nil {
			out.AutoscalingPolicy = *policy
		}
	}
	out.Description = in.Description
	out.Region = in.GetRegion()
	out.ResourceID = in.Name
	out.Target = in.GetTarget()
	return out
}

// ComputeRegionAutoscalerSpec_v1alpha1_ToProto maps a krm.ComputeRegionAutoscalerSpec to a pb.Autoscaler.
func ComputeRegionAutoscalerSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionAutoscalerSpec) *pb.Autoscaler {
	if in == nil {
		return nil
	}
	out := &pb.Autoscaler{}
	if !RegionautoscalerAutoscalingPolicy_IsEmpty(&in.AutoscalingPolicy) {
		out.AutoscalingPolicy = RegionautoscalerAutoscalingPolicy_v1alpha1_ToProto(mapCtx, &in.AutoscalingPolicy)
	}
	out.Description = in.Description
	if in.Region != "" {
		out.Region = &in.Region
	}
	out.Name = in.ResourceID
	if in.Target != "" {
		out.Target = &in.Target
	}
	return out
}

// RegionautoscalerAutoscalingPolicy_IsEmpty checks if a RegionautoscalerAutoscalingPolicy is empty.
func RegionautoscalerAutoscalingPolicy_IsEmpty(in *krm.RegionautoscalerAutoscalingPolicy) bool {
	if in == nil {
		return true
	}
	return in.CooldownPeriod == nil &&
		in.CpuUtilization == nil &&
		in.LoadBalancingUtilization == nil &&
		in.MaxReplicas == 0 &&
		len(in.Metric) == 0 &&
		in.MinReplicas == 0 &&
		in.Mode == nil &&
		in.ScaleDownControl == nil &&
		in.ScaleInControl == nil &&
		len(in.ScalingSchedules) == 0
}

// RegionautoscalerAutoscalingPolicy_v1alpha1_FromProto maps a pb.AutoscalingPolicy to a krm.RegionautoscalerAutoscalingPolicy.
func RegionautoscalerAutoscalingPolicy_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.RegionautoscalerAutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerAutoscalingPolicy{}
	out.CooldownPeriod = in.CoolDownPeriodSec
	out.CpuUtilization = RegionautoscalerCpuUtilization_v1alpha1_FromProto(mapCtx, in.CpuUtilization)
	out.LoadBalancingUtilization = RegionautoscalerLoadBalancingUtilization_v1alpha1_FromProto(mapCtx, in.LoadBalancingUtilization)
	out.MaxReplicas = in.GetMaxNumReplicas()
	out.Metric = direct.Slice_FromProto(mapCtx, in.CustomMetricUtilizations, RegionautoscalerMetric_v1alpha1_FromProto)
	out.MinReplicas = in.GetMinNumReplicas()
	out.Mode = in.Mode
	out.ScaleInControl = RegionautoscalerScaleInControl_v1alpha1_FromProto(mapCtx, in.ScaleInControl)
	out.ScalingSchedules = ScalingSchedules_v1alpha1_FromProto(mapCtx, in.ScalingSchedules)
	return out
}

// RegionautoscalerAutoscalingPolicy_v1alpha1_ToProto maps a krm.RegionautoscalerAutoscalingPolicy to a pb.AutoscalingPolicy.
func RegionautoscalerAutoscalingPolicy_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerAutoscalingPolicy) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out.CoolDownPeriodSec = in.CooldownPeriod
	out.CpuUtilization = RegionautoscalerCpuUtilization_v1alpha1_ToProto(mapCtx, in.CpuUtilization)
	out.LoadBalancingUtilization = RegionautoscalerLoadBalancingUtilization_v1alpha1_ToProto(mapCtx, in.LoadBalancingUtilization)
	if in.MaxReplicas != 0 {
		out.MaxNumReplicas = &in.MaxReplicas
	}
	out.CustomMetricUtilizations = direct.Slice_ToProto(mapCtx, in.Metric, RegionautoscalerMetric_v1alpha1_ToProto)
	if in.MinReplicas != 0 {
		out.MinNumReplicas = &in.MinReplicas
	}
	out.Mode = in.Mode
	out.ScaleInControl = RegionautoscalerScaleInControl_v1alpha1_ToProto(mapCtx, in.ScaleInControl)
	out.ScalingSchedules = ScalingSchedules_v1alpha1_ToProto(mapCtx, in.ScalingSchedules)
	return out
}

// ScalingSchedules_v1alpha1_FromProto maps a protobuf scaling schedules map to a slice of krm.RegionautoscalerScalingSchedules.
func ScalingSchedules_v1alpha1_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AutoscalingPolicyScalingSchedule) []krm.RegionautoscalerScalingSchedules {
	if in == nil {
		return nil
	}
	var out []krm.RegionautoscalerScalingSchedules
	for name, val := range in {
		if val == nil {
			continue
		}
		item := RegionautoscalerScalingSchedules_v1alpha1_FromProto(mapCtx, val)
		if item != nil {
			item.Name = name
			out = append(out, *item)
		}
	}
	return out
}

// ScalingSchedules_v1alpha1_ToProto maps a slice of krm.RegionautoscalerScalingSchedules to a protobuf scaling schedules map.
func ScalingSchedules_v1alpha1_ToProto(mapCtx *direct.MapContext, in []krm.RegionautoscalerScalingSchedules) map[string]*pb.AutoscalingPolicyScalingSchedule {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AutoscalingPolicyScalingSchedule)
	for _, item := range in {
		val := RegionautoscalerScalingSchedules_v1alpha1_ToProto(mapCtx, &item)
		if val != nil {
			out[item.Name] = val
		}
	}
	return out
}

// RegionautoscalerScalingSchedules_v1alpha1_FromProto maps a pb.AutoscalingPolicyScalingSchedule to a krm.RegionautoscalerScalingSchedules.
func RegionautoscalerScalingSchedules_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyScalingSchedule) *krm.RegionautoscalerScalingSchedules {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerScalingSchedules{}
	out.Description = in.Description
	out.Disabled = in.Disabled
	out.DurationSec = in.GetDurationSec()
	out.MinRequiredReplicas = in.GetMinRequiredReplicas()
	if in.Schedule != nil {
		out.Schedule = *in.Schedule
	}
	out.TimeZone = in.TimeZone
	return out
}

// RegionautoscalerScalingSchedules_v1alpha1_ToProto maps a krm.RegionautoscalerScalingSchedules to a pb.AutoscalingPolicyScalingSchedule.
func RegionautoscalerScalingSchedules_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerScalingSchedules) *pb.AutoscalingPolicyScalingSchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyScalingSchedule{}
	out.Description = in.Description
	out.Disabled = in.Disabled
	if in.DurationSec != 0 {
		out.DurationSec = &in.DurationSec
	}
	if in.MinRequiredReplicas != 0 {
		out.MinRequiredReplicas = &in.MinRequiredReplicas
	}
	if in.Schedule != "" {
		out.Schedule = &in.Schedule
	}
	out.TimeZone = in.TimeZone
	return out
}

// RegionautoscalerCpuUtilization_v1alpha1_FromProto maps a pb.AutoscalingPolicyCpuUtilization to a krm.RegionautoscalerCpuUtilization.
func RegionautoscalerCpuUtilization_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyCpuUtilization) *krm.RegionautoscalerCpuUtilization {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerCpuUtilization{}
	out.PredictiveMethod = in.PredictiveMethod
	out.Target = in.GetUtilizationTarget()
	return out
}

// RegionautoscalerCpuUtilization_v1alpha1_ToProto maps a krm.RegionautoscalerCpuUtilization to a pb.AutoscalingPolicyCpuUtilization.
func RegionautoscalerCpuUtilization_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerCpuUtilization) *pb.AutoscalingPolicyCpuUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyCpuUtilization{}
	out.PredictiveMethod = in.PredictiveMethod
	if in.Target != 0 {
		out.UtilizationTarget = &in.Target
	}
	return out
}

// RegionautoscalerLoadBalancingUtilization_v1alpha1_FromProto maps a pb.AutoscalingPolicyLoadBalancingUtilization to a krm.RegionautoscalerLoadBalancingUtilization.
func RegionautoscalerLoadBalancingUtilization_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyLoadBalancingUtilization) *krm.RegionautoscalerLoadBalancingUtilization {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerLoadBalancingUtilization{}
	out.Target = in.GetUtilizationTarget()
	return out
}

// RegionautoscalerLoadBalancingUtilization_v1alpha1_ToProto maps a krm.RegionautoscalerLoadBalancingUtilization to a pb.AutoscalingPolicyLoadBalancingUtilization.
func RegionautoscalerLoadBalancingUtilization_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerLoadBalancingUtilization) *pb.AutoscalingPolicyLoadBalancingUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyLoadBalancingUtilization{}
	if in.Target != 0 {
		out.UtilizationTarget = &in.Target
	}
	return out
}

// RegionautoscalerMetric_v1alpha1_FromProto maps a pb.AutoscalingPolicyCustomMetricUtilization to a krm.RegionautoscalerMetric.
func RegionautoscalerMetric_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyCustomMetricUtilization) *krm.RegionautoscalerMetric {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerMetric{}
	out.Filter = in.Filter
	out.Name = in.GetMetric()
	out.SingleInstanceAssignment = in.SingleInstanceAssignment
	out.Target = in.UtilizationTarget
	out.Type = in.UtilizationTargetType
	return out
}

// RegionautoscalerMetric_v1alpha1_ToProto maps a krm.RegionautoscalerMetric to a pb.AutoscalingPolicyCustomMetricUtilization.
func RegionautoscalerMetric_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerMetric) *pb.AutoscalingPolicyCustomMetricUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyCustomMetricUtilization{}
	out.Filter = in.Filter
	if in.Name != "" {
		out.Metric = &in.Name
	}
	out.SingleInstanceAssignment = in.SingleInstanceAssignment
	out.UtilizationTarget = in.Target
	out.UtilizationTargetType = in.Type
	return out
}

// RegionautoscalerFixedOrPercent_v1alpha1_FromProto maps a pb.FixedOrPercent to a krm.RegionautoscalerFixedOrPercent.
func RegionautoscalerFixedOrPercent_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.RegionautoscalerFixedOrPercent {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerFixedOrPercent{}
	out.Fixed = in.Fixed
	out.Percent = in.Percent
	return out
}

// RegionautoscalerFixedOrPercent_v1alpha1_ToProto maps a krm.RegionautoscalerFixedOrPercent to a pb.FixedOrPercent.
func RegionautoscalerFixedOrPercent_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerFixedOrPercent) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Fixed = in.Fixed
	out.Percent = in.Percent
	return out
}

// RegionautoscalerScaleInControl_v1alpha1_FromProto maps a pb.AutoscalingPolicyScaleInControl to a krm.RegionautoscalerScaleInControl.
func RegionautoscalerScaleInControl_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyScaleInControl) *krm.RegionautoscalerScaleInControl {
	if in == nil {
		return nil
	}
	out := &krm.RegionautoscalerScaleInControl{}
	out.MaxScaledInReplicas = RegionautoscalerFixedOrPercent_v1alpha1_FromProto(mapCtx, in.MaxScaledInReplicas)
	out.TimeWindowSec = in.TimeWindowSec
	return out
}

// RegionautoscalerScaleInControl_v1alpha1_ToProto maps a krm.RegionautoscalerScaleInControl to a pb.AutoscalingPolicyScaleInControl.
func RegionautoscalerScaleInControl_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.RegionautoscalerScaleInControl) *pb.AutoscalingPolicyScaleInControl {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyScaleInControl{}
	out.MaxScaledInReplicas = RegionautoscalerFixedOrPercent_v1alpha1_ToProto(mapCtx, in.MaxScaledInReplicas)
	out.TimeWindowSec = in.TimeWindowSec
	return out
}

// ComputeRegionAutoscalerStatus_v1alpha1_FromProto maps a pb.Autoscaler to a krm.ComputeRegionAutoscalerStatus.
func ComputeRegionAutoscalerStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Autoscaler) *krm.ComputeRegionAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionAutoscalerStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

// ComputeRegionAutoscalerStatus_v1alpha1_ToProto maps a krm.ComputeRegionAutoscalerStatus to a pb.Autoscaler.
func ComputeRegionAutoscalerStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionAutoscalerStatus) *pb.Autoscaler {
	if in == nil {
		return nil
	}
	out := &pb.Autoscaler{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}
