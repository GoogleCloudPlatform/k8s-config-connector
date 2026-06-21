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
	fuzztesting.RegisterKRMFuzzer(computeAutoscalerFuzzer())
}

func computeAutoscalerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Autoscaler{},
		ComputeAutoscalerSpec_v1alpha1_FromProto, ComputeAutoscalerSpec_v1alpha1_ToProto,
		ComputeAutoscalerStatus_v1alpha1_FromProto, ComputeAutoscalerStatus_v1alpha1_ToProto,
	)

	// Field comparison: ComputeAutoscaler Spec vs pb.Autoscaler Proto
	// - Spec.AutoscalingPolicy             maps to proto field .autoscaling_policy
	//   - CooldownPeriod                   maps to .cool_down_period_sec
	//   - CpuUtilization                   maps to .cpu_utilization
	//     - PredictiveMethod               maps to .predictive_method
	//     - Target                         maps to .utilization_target
	//   - LoadBalancingUtilization         maps to .load_balancing_utilization
	//     - Target                         maps to .utilization_target
	//   - MaxReplicas                      maps to .max_num_replicas
	//   - Metric                           maps to .custom_metric_utilizations
	//     - Filter                         maps to .filter
	//     - Name                           maps to .metric
	//     - SingleInstanceAssignment       maps to .single_instance_assignment
	//     - Target                         maps to .utilization_target
	//     - Type                           maps to .utilization_target_type
	//   - MinReplicas                      maps to .min_num_replicas
	//   - Mode                             maps to .mode
	//   - ScaleDownControl                 maps to .scale_in_control (proto only has scale_in_control)
	//     - MaxScaledDownReplicas          maps to .max_scaled_in_replicas
	//       - Fixed                        maps to .fixed
	//       - Percent                      maps to .percent
	//     - TimeWindowSec                  maps to .time_window_sec
	//   - ScaleInControl                   maps to .scale_in_control
	//     - MaxScaledInReplicas            maps to .max_scaled_in_replicas
	//       - Fixed                        maps to .fixed
	//       - Percent                      maps to .percent
	//     - TimeWindowSec                  maps to .time_window_sec
	//   - ScalingSchedules                 maps to .scaling_schedules (slice to map)
	//     - Description                    maps to .description
	//     - Disabled                       maps to .disabled
	//     - DurationSec                    maps to .duration_sec
	//     - MinRequiredReplicas            maps to .min_required_replicas
	//     - Name                           maps to key in .scaling_schedules
	//     - Schedule                       maps to .schedule
	//     - TimeZone                       maps to .time_zone
	// - Spec.Description                   maps to proto field .description
	// - Spec.ProjectRef                    is handled via the GCP URI/ID, not mapped to a proto body field directly
	// - Spec.ResourceID                    maps to proto field .name (handled as Unimplemented_Identity)
	// - Spec.TargetRef                     maps to proto field .target
	// - Spec.Zone                          maps to proto field .zone
	//
	// Spec fields
	f.SpecField(".autoscaling_policy")
	f.SpecField(".description")
	f.SpecField(".target")
	f.SpecField(".zone")

	f.StatusField(".creation_timestamp")
	f.StatusField(".self_link")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_Internal(".id")
	f.Unimplemented_Internal(".kind")
	f.Unimplemented_Internal(".recommended_size")
	f.Unimplemented_Internal(".region")
	f.Unimplemented_Internal(".status")
	f.Unimplemented_Internal(".status_details")
	f.Unimplemented_Internal(".scaling_schedule_status")
	f.Unimplemented_Internal(".autoscaling_policy.scale_in_control.max_scaled_in_replicas.calculated")

	return f
}
