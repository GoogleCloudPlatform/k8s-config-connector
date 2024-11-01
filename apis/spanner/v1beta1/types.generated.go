// Copyright 2024 Google LLC
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

package v1beta1

// +kcc:proto=google.spanner.admin.instance.v1.AutoscalingConfig
type AutoscalingConfig struct {
	// Required. Autoscaling limits for an instance.
	AutoscalingLimits *AutoscalingConfig_AutoscalingLimits `json:"autoscalingLimits,omitempty"`

	// Required. The autoscaling targets for an instance.
	AutoscalingTargets *AutoscalingConfig_AutoscalingTargets `json:"autoscalingTargets,omitempty"`

	// Optional. Optional asymmetric autoscaling options.
	//  Replicas matching the replica selection criteria will be autoscaled
	//  independently from other replicas. The autoscaler will scale the replicas
	//  based on the utilization of replicas identified by the replica selection.
	//  Replica selections should not overlap with each other.
	//
	//  Other replicas (those do not match any replica selection) will be
	//  autoscaled together and will have the same compute capacity allocated to
	//  them.
	AsymmetricAutoscalingOptions []AutoscalingConfig_AsymmetricAutoscalingOption `json:"asymmetricAutoscalingOptions,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.AutoscalingConfig.AsymmetricAutoscalingOption
type AutoscalingConfig_AsymmetricAutoscalingOption struct {
	// Required. Selects the replicas to which this AsymmetricAutoscalingOption
	//  applies. Only read-only replicas are supported.
	ReplicaSelection *ReplicaSelection `json:"replicaSelection,omitempty"`

	// Optional. Overrides applied to the top-level autoscaling configuration
	//  for the selected replicas.
	Overrides *AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides `json:"overrides,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.AutoscalingConfig.AsymmetricAutoscalingOption.AutoscalingConfigOverrides
type AutoscalingConfig_AsymmetricAutoscalingOption_AutoscalingConfigOverrides struct {
	// Optional. If specified, overrides the min/max limit in the top-level
	//  autoscaling configuration for the selected replicas.
	AutoscalingLimits *AutoscalingConfig_AutoscalingLimits `json:"autoscalingLimits,omitempty"`

	// Optional. If specified, overrides the autoscaling target
	//  high_priority_cpu_utilization_percent in the top-level autoscaling
	//  configuration for the selected replicas.
	AutoscalingTargetHighPriorityCpuUtilizationPercent *int32 `json:"autoscalingTargetHighPriorityCpuUtilizationPercent,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.AutoscalingConfig.AutoscalingLimits
type AutoscalingConfig_AutoscalingLimits struct {
	// Minimum number of nodes allocated to the instance. If set, this number
	//  should be greater than or equal to 1.
	MinNodes *int32 `json:"minNodes,omitempty"`

	// Minimum number of processing units allocated to the instance. If set,
	//  this number should be multiples of 1000.
	MinProcessingUnits *int32 `json:"minProcessingUnits,omitempty"`

	// Maximum number of nodes allocated to the instance. If set, this number
	//  should be greater than or equal to min_nodes.
	MaxNodes *int32 `json:"maxNodes,omitempty"`

	// Maximum number of processing units allocated to the instance. If set,
	//  this number should be multiples of 1000 and be greater than or equal to
	//  min_processing_units.
	MaxProcessingUnits *int32 `json:"maxProcessingUnits,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.AutoscalingConfig.AutoscalingTargets
type AutoscalingConfig_AutoscalingTargets struct {
	// Required. The target high priority cpu utilization percentage that the
	//  autoscaler should be trying to achieve for the instance. This number is
	//  on a scale from 0 (no utilization) to 100 (full utilization). The valid
	//  range is [10, 90] inclusive.
	HighPriorityCpuUtilizationPercent *int32 `json:"highPriorityCpuUtilizationPercent,omitempty"`

	// Required. The target storage utilization percentage that the autoscaler
	//  should be trying to achieve for the instance. This number is on a scale
	//  from 0 (no utilization) to 100 (full utilization). The valid range is
	//  [10, 100] inclusive.
	StorageUtilizationPercent *int32 `json:"storageUtilizationPercent,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.Instance
type Instance struct {
	// Required. A unique identifier for the instance, which cannot be changed
	//  after the instance is created. Values are of the form
	//  `projects/<project>/instances/[a-z][-a-z0-9]*[a-z0-9]`. The final
	//  segment of the name must be between 2 and 64 characters in length.
	Name *string `json:"name,omitempty"`

	// Required. The name of the instance's configuration. Values are of the form
	//  `projects/<project>/instanceConfigs/<configuration>`. See
	//  also [InstanceConfig][google.spanner.admin.instance.v1.InstanceConfig] and
	//  [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
	Config *string `json:"config,omitempty"`

	// Required. The descriptive name for this instance as it appears in UIs.
	//  Must be unique per project and between 4 and 30 characters in length.
	DisplayName *string `json:"displayName,omitempty"`

	// The number of nodes allocated to this instance. At most, one of either
	//  `node_count` or `processing_units` should be present in the message.
	//
	//  Users can set the `node_count` field to specify the target number of nodes
	//  allocated to the instance.
	//
	//  If autoscaling is enabled, `node_count` is treated as an `OUTPUT_ONLY`
	//  field and reflects the current number of nodes allocated to the instance.
	//
	//  This might be zero in API responses for instances that are not yet in the
	//  `READY` state.
	//
	//  If the instance has varying node count across replicas (achieved by
	//  setting asymmetric_autoscaling_options in autoscaling config), the
	//  node_count here is the maximum node count across all replicas.
	//
	//  For more information, see
	//  [Compute capacity, nodes, and processing
	//  units](https://cloud.google.com/spanner/docs/compute-capacity).
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// The number of processing units allocated to this instance. At most, one of
	//  either `processing_units` or `node_count` should be present in the message.
	//
	//  Users can set the `processing_units` field to specify the target number of
	//  processing units allocated to the instance.
	//
	//  If autoscaling is enabled, `processing_units` is treated as an
	//  `OUTPUT_ONLY` field and reflects the current number of processing units
	//  allocated to the instance.
	//
	//  This might be zero in API responses for instances that are not yet in the
	//  `READY` state.
	//
	//  If the instance has varying processing units per replica
	//  (achieved by setting asymmetric_autoscaling_options in autoscaling config),
	//  the processing_units here is the maximum processing units across all
	//  replicas.
	//
	//  For more information, see
	//  [Compute capacity, nodes and processing
	//  units](https://cloud.google.com/spanner/docs/compute-capacity).
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`

	// Output only. Lists the compute capacity per ReplicaSelection. A replica
	//  selection identifies a set of replicas with common properties. Replicas
	//  identified by a ReplicaSelection are scaled with the same compute capacity.
	ReplicaComputeCapacity []ReplicaComputeCapacity `json:"replicaComputeCapacity,omitempty"`

	// Optional. The autoscaling configuration. Autoscaling is enabled if this
	//  field is set. When autoscaling is enabled, node_count and processing_units
	//  are treated as OUTPUT_ONLY fields and reflect the current compute capacity
	//  allocated to the instance.
	AutoscalingConfig *AutoscalingConfig `json:"autoscalingConfig,omitempty"`

	// Output only. The current instance state. For
	//  [CreateInstance][google.spanner.admin.instance.v1.InstanceAdmin.CreateInstance],
	//  the state must be either omitted or set to `CREATING`. For
	//  [UpdateInstance][google.spanner.admin.instance.v1.InstanceAdmin.UpdateInstance],
	//  the state must be either omitted or set to `READY`.
	State *string `json:"state,omitempty"`

	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud
	//  resources into groups that reflect a customer's organizational needs and
	//  deployment strategies. Cloud Labels can be used to filter collections of
	//  resources. They can be used to control how resource metrics are aggregated.
	//  And they can be used as arguments to policy management rules (e.g. route,
	//  firewall, load balancing, etc.).
	//
	//   * Label keys must be between 1 and 63 characters long and must conform to
	//     the following regular expression: `[a-z][a-z0-9_-]{0,62}`.
	//   * Label values must be between 0 and 63 characters long and must conform
	//     to the regular expression `[a-z0-9_-]{0,63}`.
	//   * No more than 64 labels can be associated with a given resource.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//
	//  If you plan to use labels in your own code, please note that additional
	//  characters may be allowed in the future. And so you are advised to use an
	//  internal label representation, such as JSON, which doesn't rely upon
	//  specific characters being disallowed.  For example, representing labels
	//  as the string:  name + "_" + value  would prove problematic if we were to
	//  allow "_" in a future release.
	Labels map[string]string `json:"labels,omitempty"`

	// Deprecated. This field is not populated.
	EndpointUris []string `json:"endpointUris,omitempty"`

	// Output only. The time at which the instance was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the instance was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. The `Edition` of the current instance.
	Edition *string `json:"edition,omitempty"`

	// Optional. Controls the default backup behavior for new databases within the
	//  instance.
	//
	//  Note that `AUTOMATIC` is not permitted for free instances, as backups and
	//  backup schedules are not allowed for free instances.
	//
	//  In the `GetInstance` or `ListInstances` response, if the value of
	//  default_backup_schedule_type is unset or NONE, no default backup
	//  schedule will be created for new databases within the instance.
	DefaultBackupScheduleType *string `json:"defaultBackupScheduleType,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.ReplicaComputeCapacity
type ReplicaComputeCapacity struct {
	// Required. Identifies replicas by specified properties.
	//  All replicas in the selection have the same amount of compute capacity.
	ReplicaSelection *ReplicaSelection `json:"replicaSelection,omitempty"`

	// The number of nodes allocated to each replica.
	//
	//  This may be zero in API responses for instances that are not yet in
	//  state `READY`.
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// The number of processing units allocated to each replica.
	//
	//  This may be zero in API responses for instances that are not yet in
	//  state `READY`.
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.ReplicaSelection
type ReplicaSelection struct {
	// Required. Name of the location of the replicas (e.g., "us-central1").
	Location *string `json:"location,omitempty"`
}
