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

package v1alpha1


// +kcc:proto=google.cloud.dataproc.v1.AutoscalingPolicy
type AutoscalingPolicy struct {
	// Required. The policy id.
	//
	//  The id must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), and hyphens (-). Cannot begin or end with underscore
	//  or hyphen. Must consist of between 3 and 50 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.id
	ID *string `json:"id,omitempty"`

	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.basic_algorithm
	BasicAlgorithm *BasicAutoscalingAlgorithm `json:"basicAlgorithm,omitempty"`

	// Required. Describes how the autoscaler will operate for primary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.worker_config
	WorkerConfig *InstanceGroupAutoscalingPolicyConfig `json:"workerConfig,omitempty"`

	// Optional. Describes how the autoscaler will operate for secondary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.secondary_worker_config
	SecondaryWorkerConfig *InstanceGroupAutoscalingPolicyConfig `json:"secondaryWorkerConfig,omitempty"`

	// Optional. The labels to associate with this autoscaling policy.
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to [RFC
	//  1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	//  associated with an autoscaling policy.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm
type BasicAutoscalingAlgorithm struct {
	// Required. YARN autoscaling configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm.yarn_config
	YarnConfig *BasicYarnAutoscalingConfig `json:"yarnConfig,omitempty"`

	// Optional. Duration between scaling events. A scaling period starts after
	//  the update operation from the previous event has completed.
	//
	//  Bounds: [2m, 1d]. Default: 2m.
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm.cooldown_period
	CooldownPeriod *string `json:"cooldownPeriod,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig
type BasicYarnAutoscalingConfig struct {
	// Required. Timeout for YARN graceful decommissioning of Node Managers.
	//  Specifies the duration to wait for jobs to complete before forcefully
	//  removing workers (and potentially interrupting jobs). Only applicable to
	//  downscaling operations.
	//
	//  Bounds: [0s, 1d].
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.graceful_decommission_timeout
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout,omitempty"`

	// Required. Fraction of average YARN pending memory in the last cooldown
	//  period for which to add workers. A scale-up factor of 1.0 will result in
	//  scaling up so that there is no pending memory remaining after the update
	//  (more aggressive scaling). A scale-up factor closer to 0 will result in a
	//  smaller magnitude of scaling up (less aggressive scaling). See [How
	//  autoscaling
	//  works](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/autoscaling#how_autoscaling_works)
	//  for more information.
	//
	//  Bounds: [0.0, 1.0].
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_up_factor
	ScaleUpFactor *float64 `json:"scaleUpFactor,omitempty"`

	// Required. Fraction of average YARN pending memory in the last cooldown
	//  period for which to remove workers. A scale-down factor of 1 will result in
	//  scaling down so that there is no available memory remaining after the
	//  update (more aggressive scaling). A scale-down factor of 0 disables
	//  removing workers, which can be beneficial for autoscaling a single job.
	//  See [How autoscaling
	//  works](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/autoscaling#how_autoscaling_works)
	//  for more information.
	//
	//  Bounds: [0.0, 1.0].
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_down_factor
	ScaleDownFactor *float64 `json:"scaleDownFactor,omitempty"`

	// Optional. Minimum scale-up threshold as a fraction of total cluster size
	//  before scaling occurs. For example, in a 20-worker cluster, a threshold of
	//  0.1 means the autoscaler must recommend at least a 2-worker scale-up for
	//  the cluster to scale. A threshold of 0 means the autoscaler will scale up
	//  on any recommended change.
	//
	//  Bounds: [0.0, 1.0]. Default: 0.0.
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_up_min_worker_fraction
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty"`

	// Optional. Minimum scale-down threshold as a fraction of total cluster size
	//  before scaling occurs. For example, in a 20-worker cluster, a threshold of
	//  0.1 means the autoscaler must recommend at least a 2 worker scale-down for
	//  the cluster to scale. A threshold of 0 means the autoscaler will scale down
	//  on any recommended change.
	//
	//  Bounds: [0.0, 1.0]. Default: 0.0.
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_down_min_worker_fraction
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig
type InstanceGroupAutoscalingPolicyConfig struct {
	// Optional. Minimum number of instances for this group.
	//
	//  Primary workers - Bounds: [2, max_instances]. Default: 2.
	//  Secondary workers - Bounds: [0, max_instances]. Default: 0.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.min_instances
	MinInstances *int32 `json:"minInstances,omitempty"`

	// Required. Maximum number of instances for this group. Required for primary
	//  workers. Note that by default, clusters will not use secondary workers.
	//  Required for secondary workers if the minimum secondary instances is set.
	//
	//  Primary workers - Bounds: [min_instances, ).
	//  Secondary workers - Bounds: [min_instances, ). Default: 0.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.max_instances
	MaxInstances *int32 `json:"maxInstances,omitempty"`

	// Optional. Weight for the instance group, which is used to determine the
	//  fraction of total workers in the cluster from this instance group.
	//  For example, if primary workers have weight 2, and secondary workers have
	//  weight 1, the cluster will have approximately 2 primary workers for each
	//  secondary worker.
	//
	//  The cluster may not reach the specified balance if constrained
	//  by min/max bounds or other autoscaling settings. For example, if
	//  `max_instances` for secondary workers is 0, then only primary workers will
	//  be added. The cluster can also be out of balance when created.
	//
	//  If weight is not set on any instance group, the cluster will default to
	//  equal weight for all groups: the cluster will attempt to maintain an equal
	//  number of workers in each group within the configured size bounds for each
	//  group. If weight is set for one group only, the cluster will default to
	//  zero weight on the unset group. For example if weight is set only on
	//  primary workers, the cluster will use primary workers only and no
	//  secondary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.weight
	Weight *int32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.AutoscalingPolicy
type AutoscalingPolicyObservedState struct {
	// Output only. The "resource name" of the autoscaling policy, as described
	//  in https://cloud.google.com/apis/design/resource_names.
	//
	//  * For `projects.regions.autoscalingPolicies`, the resource name of the
	//    policy has the following format:
	//    `projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id}`
	//
	//  * For `projects.locations.autoscalingPolicies`, the resource name of the
	//    policy has the following format:
	//    `projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}`
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.name
	Name *string `json:"name,omitempty"`
}
