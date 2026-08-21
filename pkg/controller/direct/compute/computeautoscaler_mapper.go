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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeAutoscalerSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Autoscaler) *krm.ComputeAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeAutoscalerSpec{}
	out.AutoscalingPolicy = AutoscalerAutoscalingPolicy_v1alpha1_FromProto(mapCtx, in.AutoscalingPolicy)
	out.Description = in.Description
	if in.GetTarget() != "" {
		out.TargetRef = krm.AutoscalerTargetRef{External: in.GetTarget()}
	}
	out.Zone = in.GetZone()
	return out
}

func ComputeAutoscalerSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeAutoscalerSpec) *pb.Autoscaler {
	if in == nil {
		return nil
	}
	out := &pb.Autoscaler{}
	out.AutoscalingPolicy = AutoscalerAutoscalingPolicy_v1alpha1_ToProto(mapCtx, in.AutoscalingPolicy)
	out.Description = in.Description
	if in.TargetRef.External != "" {
		out.Target = direct.LazyPtr(in.TargetRef.External)
	}
	out.Zone = direct.LazyPtr(in.Zone)
	return out
}

func ComputeAutoscalerStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Autoscaler) *krm.ComputeAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeAutoscalerStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

func ComputeAutoscalerStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeAutoscalerStatus) *pb.Autoscaler {
	if in == nil {
		return nil
	}
	out := &pb.Autoscaler{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

func AutoscalerAutoscalingPolicy_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerAutoscalingPolicy) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out.CoolDownPeriodSec = direct.PtrInt64ToPtrInt32(in.CooldownPeriod)
	out.CpuUtilization = AutoscalerCpuUtilization_v1alpha1_ToProto(mapCtx, in.CpuUtilization)
	out.LoadBalancingUtilization = AutoscalerLoadBalancingUtilization_v1alpha1_ToProto(mapCtx, in.LoadBalancingUtilization)
	out.MaxNumReplicas = direct.LazyPtr(int32(in.MaxReplicas))
	out.MinNumReplicas = direct.LazyPtr(int32(in.MinReplicas))
	out.Mode = in.Mode

	if in.Metric != nil {
		out.CustomMetricUtilizations = make([]*pb.AutoscalingPolicyCustomMetricUtilization, len(in.Metric))
		for i, m := range in.Metric {
			out.CustomMetricUtilizations[i] = AutoscalerMetric_v1alpha1_ToProto(mapCtx, &m)
		}
	}

	if in.ScaleInControl != nil {
		out.ScaleInControl = AutoscalerScaleInControl_v1alpha1_ToProto(mapCtx, in.ScaleInControl)
	} else if in.ScaleDownControl != nil {
		out.ScaleInControl = AutoscalerScaleDownControl_v1alpha1_ToProto(mapCtx, in.ScaleDownControl)
	}

	if in.ScalingSchedules != nil {
		out.ScalingSchedules = make(map[string]*pb.AutoscalingPolicyScalingSchedule)
		for _, s := range in.ScalingSchedules {
			if s.Name != "" {
				out.ScalingSchedules[s.Name] = AutoscalerScalingSchedule_v1alpha1_ToProto(mapCtx, &s)
			}
		}
	}

	return out
}

func AutoscalerAutoscalingPolicy_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.AutoscalerAutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerAutoscalingPolicy{}
	out.CooldownPeriod = direct.PtrInt32ToPtrInt64(in.CoolDownPeriodSec)
	out.CpuUtilization = AutoscalerCpuUtilization_v1alpha1_FromProto(mapCtx, in.CpuUtilization)
	out.LoadBalancingUtilization = AutoscalerLoadBalancingUtilization_v1alpha1_FromProto(mapCtx, in.LoadBalancingUtilization)
	out.MaxReplicas = int64(in.GetMaxNumReplicas())
	out.MinReplicas = int64(in.GetMinNumReplicas())
	out.Mode = in.Mode

	if in.CustomMetricUtilizations != nil {
		out.Metric = make([]krm.AutoscalerMetric, 0, len(in.CustomMetricUtilizations))
		for _, m := range in.CustomMetricUtilizations {
			if m != nil {
				out.Metric = append(out.Metric, *AutoscalerMetric_v1alpha1_FromProto(mapCtx, m))
			}
		}
	}

	if in.ScaleInControl != nil {
		out.ScaleInControl = AutoscalerScaleInControl_v1alpha1_FromProto(mapCtx, in.ScaleInControl)
		out.ScaleDownControl = AutoscalerScaleDownControl_v1alpha1_FromProto(mapCtx, in.ScaleInControl)
	}

	if in.ScalingSchedules != nil {
		out.ScalingSchedules = make([]krm.AutoscalerScalingSchedules, 0, len(in.ScalingSchedules))
		for name, s := range in.ScalingSchedules {
			if s != nil {
				sch := AutoscalerScalingSchedule_v1alpha1_FromProto(mapCtx, s)
				sch.Name = name
				out.ScalingSchedules = append(out.ScalingSchedules, *sch)
			}
		}
	}

	return out
}

func AutoscalerCpuUtilization_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerCpuUtilization) *pb.AutoscalingPolicyCpuUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyCpuUtilization{}
	out.PredictiveMethod = in.PredictiveMethod
	out.UtilizationTarget = direct.LazyPtr(in.Target)
	return out
}

func AutoscalerCpuUtilization_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyCpuUtilization) *krm.AutoscalerCpuUtilization {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerCpuUtilization{}
	out.PredictiveMethod = in.PredictiveMethod
	out.Target = in.GetUtilizationTarget()
	return out
}

func AutoscalerLoadBalancingUtilization_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerLoadBalancingUtilization) *pb.AutoscalingPolicyLoadBalancingUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyLoadBalancingUtilization{}
	out.UtilizationTarget = direct.LazyPtr(in.Target)
	return out
}

func AutoscalerLoadBalancingUtilization_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyLoadBalancingUtilization) *krm.AutoscalerLoadBalancingUtilization {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerLoadBalancingUtilization{}
	out.Target = in.GetUtilizationTarget()
	return out
}

func AutoscalerMetric_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerMetric) *pb.AutoscalingPolicyCustomMetricUtilization {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyCustomMetricUtilization{}
	out.Filter = in.Filter
	out.Metric = direct.LazyPtr(in.Name)
	out.SingleInstanceAssignment = in.SingleInstanceAssignment
	out.UtilizationTarget = in.Target
	out.UtilizationTargetType = in.Type
	return out
}

func AutoscalerMetric_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyCustomMetricUtilization) *krm.AutoscalerMetric {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerMetric{}
	out.Filter = in.Filter
	out.Name = in.GetMetric()
	out.SingleInstanceAssignment = in.SingleInstanceAssignment
	out.Target = in.UtilizationTarget
	out.Type = in.UtilizationTargetType
	return out
}

func AutoscalerScaleInControl_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerScaleInControl) *pb.AutoscalingPolicyScaleInControl {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyScaleInControl{}
	out.MaxScaledInReplicas = FixedOrPercent_v1alpha1_ToProto(mapCtx, in.MaxScaledInReplicas)
	out.TimeWindowSec = direct.PtrInt64ToPtrInt32(in.TimeWindowSec)
	return out
}

func AutoscalerScaleInControl_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyScaleInControl) *krm.AutoscalerScaleInControl {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerScaleInControl{}
	out.MaxScaledInReplicas = FixedOrPercent_v1alpha1_FromProto(mapCtx, in.MaxScaledInReplicas)
	out.TimeWindowSec = direct.PtrInt32ToPtrInt64(in.TimeWindowSec)
	return out
}

func AutoscalerScaleDownControl_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerScaleDownControl) *pb.AutoscalingPolicyScaleInControl {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyScaleInControl{}
	out.MaxScaledInReplicas = FixedOrPercent_v1alpha1_ToProtoDown(mapCtx, in.MaxScaledDownReplicas)
	out.TimeWindowSec = direct.PtrInt64ToPtrInt32(in.TimeWindowSec)
	return out
}

func AutoscalerScaleDownControl_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyScaleInControl) *krm.AutoscalerScaleDownControl {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerScaleDownControl{}
	out.MaxScaledDownReplicas = FixedOrPercent_v1alpha1_FromProtoDown(mapCtx, in.MaxScaledInReplicas)
	out.TimeWindowSec = direct.PtrInt32ToPtrInt64(in.TimeWindowSec)
	return out
}

func FixedOrPercent_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerMaxScaledInReplicas) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Fixed = direct.PtrInt64ToPtrInt32(in.Fixed)
	out.Percent = direct.PtrInt64ToPtrInt32(in.Percent)
	return out
}

func FixedOrPercent_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.AutoscalerMaxScaledInReplicas {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerMaxScaledInReplicas{}
	out.Fixed = direct.PtrInt32ToPtrInt64(in.Fixed)
	out.Percent = direct.PtrInt32ToPtrInt64(in.Percent)
	return out
}

func FixedOrPercent_v1alpha1_ToProtoDown(mapCtx *direct.MapContext, in *krm.AutoscalerMaxScaledDownReplicas) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	out.Fixed = direct.PtrInt64ToPtrInt32(in.Fixed)
	out.Percent = direct.PtrInt64ToPtrInt32(in.Percent)
	return out
}

func FixedOrPercent_v1alpha1_FromProtoDown(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.AutoscalerMaxScaledDownReplicas {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerMaxScaledDownReplicas{}
	out.Fixed = direct.PtrInt32ToPtrInt64(in.Fixed)
	out.Percent = direct.PtrInt32ToPtrInt64(in.Percent)
	return out
}

func AutoscalerScalingSchedule_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalerScalingSchedules) *pb.AutoscalingPolicyScalingSchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicyScalingSchedule{}
	out.Description = in.Description
	out.Disabled = in.Disabled
	out.DurationSec = direct.LazyPtr(int32(in.DurationSec))
	out.MinRequiredReplicas = direct.LazyPtr(int32(in.MinRequiredReplicas))
	out.Schedule = direct.LazyPtr(in.Schedule)
	out.TimeZone = in.TimeZone
	return out
}

func AutoscalerScalingSchedule_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicyScalingSchedule) *krm.AutoscalerScalingSchedules {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalerScalingSchedules{}
	out.Description = in.Description
	out.Disabled = in.Disabled
	out.DurationSec = int64(in.GetDurationSec())
	out.MinRequiredReplicas = int64(in.GetMinRequiredReplicas())
	out.Schedule = in.GetSchedule()
	out.TimeZone = in.TimeZone
	return out
}
