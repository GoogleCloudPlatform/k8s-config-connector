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

// +tool:fuzz-gen
// proto.message: google.cloud.compute.v1.Autoscaler
// api.group: compute.cnrm.cloud.google.com

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(computeRegionAutoscalerFuzzer())
}

func computeRegionAutoscalerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Autoscaler{},
		ComputeRegionAutoscalerSpec_v1alpha1_FromProto, ComputeRegionAutoscalerSpec_v1alpha1_ToProto,
		ComputeRegionAutoscalerStatus_v1alpha1_FromProto, ComputeRegionAutoscalerStatus_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".region")
	f.SpecField(".target")

	// Fuzz individual fields inside .autoscaling_policy
	f.SpecField(".autoscaling_policy")
	f.SpecField(".autoscaling_policy.cool_down_period_sec")
	f.SpecField(".autoscaling_policy.cpu_utilization")
	f.SpecField(".autoscaling_policy.cpu_utilization.predictive_method")
	f.SpecField(".autoscaling_policy.cpu_utilization.utilization_target")
	f.SpecField(".autoscaling_policy.load_balancing_utilization")
	f.SpecField(".autoscaling_policy.load_balancing_utilization.utilization_target")
	f.SpecField(".autoscaling_policy.max_num_replicas")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations[].filter")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations[].metric")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations[].single_instance_assignment")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations[].utilization_target")
	f.SpecField(".autoscaling_policy.custom_metric_utilizations[].utilization_target_type")
	f.SpecField(".autoscaling_policy.min_num_replicas")
	f.SpecField(".autoscaling_policy.mode")
	f.SpecField(".autoscaling_policy.scale_in_control")
	f.SpecField(".autoscaling_policy.scale_in_control.max_scaled_in_replicas")
	f.SpecField(".autoscaling_policy.scale_in_control.max_scaled_in_replicas.fixed")
	f.SpecField(".autoscaling_policy.scale_in_control.max_scaled_in_replicas.percent")
	f.SpecField(".autoscaling_policy.scale_in_control.time_window_sec")
	f.SpecField(".autoscaling_policy.scaling_schedules")

	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".zone")

	f.Unimplemented_NotYetTriaged(".recommended_size")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".status_details")
	f.Unimplemented_NotYetTriaged(".scaling_schedule_status")
	f.Unimplemented_NotYetTriaged(".autoscaling_policy.scale_in_control.max_scaled_in_replicas.calculated")
	f.Unimplemented_NotYetTriaged(".autoscaling_policy.stabilization_period_sec")

	f.FilterSpec = func(in *pb.Autoscaler) {
		if in.AutoscalingPolicy == nil {
			return
		}
		p := in.AutoscalingPolicy
		if p.CpuUtilization != nil && p.CpuUtilization.PredictiveMethod == nil && p.CpuUtilization.UtilizationTarget == nil {
			p.CpuUtilization = nil
		}
		if p.LoadBalancingUtilization != nil && p.LoadBalancingUtilization.UtilizationTarget == nil {
			p.LoadBalancingUtilization = nil
		}
		if p.ScaleInControl != nil && p.ScaleInControl.MaxScaledInReplicas == nil && p.ScaleInControl.TimeWindowSec == nil {
			p.ScaleInControl = nil
		}
		if p.ScaleInControl != nil && p.ScaleInControl.MaxScaledInReplicas != nil && p.ScaleInControl.MaxScaledInReplicas.Fixed == nil && p.ScaleInControl.MaxScaledInReplicas.Percent == nil {
			p.ScaleInControl.MaxScaledInReplicas = nil
		}
		if p.CoolDownPeriodSec == nil &&
			p.CpuUtilization == nil &&
			p.LoadBalancingUtilization == nil &&
			p.MaxNumReplicas == nil &&
			len(p.CustomMetricUtilizations) == 0 &&
			p.MinNumReplicas == nil &&
			p.Mode == nil &&
			p.ScaleInControl == nil &&
			len(p.ScalingSchedules) == 0 {
			in.AutoscalingPolicy = nil
		}
	}

	f.FilterStatus = func(in *pb.Autoscaler) {
		in.AutoscalingPolicy = nil
	}

	return f
}
