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

// +generated:types
// krm.group: spanner.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.spanner.admin.instance.v1
// resource: SpannerInstanceConfig:InstanceConfig

package v1alpha1

// +kcc:proto=google.spanner.admin.instance.v1.InstanceConfig
type InstanceConfig struct {
	// A unique identifier for the instance configuration.  Values
	//  are of the form
	//  `projects/<project>/instanceConfigs/[a-z][-a-z0-9]*`.
	//
	//  User instance configuration must start with `custom-`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.name
	Name *string `json:"name,omitempty"`

	// The name of this instance configuration as it appears in UIs.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The geographic placement of nodes in this instance configuration and their
	//  replication properties.
	//
	//  To create user-managed configurations, input
	//  `replicas` must include all replicas in `replicas` of the `base_config`
	//  and include one or more replicas in the `optional_replicas` of the
	//  `base_config`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.replicas
	Replicas []ReplicaInfo `json:"replicas,omitempty"`

	// Base configuration name, e.g. projects/<project_name>/instanceConfigs/nam3,
	//  based on which this configuration is created. Only set for user-managed
	//  configurations. `base_config` must refer to a configuration of type
	//  `GOOGLE_MANAGED` in the same project as this configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.base_config
	BaseConfig *string `json:"baseConfig,omitempty"`

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
	//  characters may be allowed in the future. Therefore, you are advised to use
	//  an internal label representation, such as JSON, which doesn't rely upon
	//  specific characters being disallowed.  For example, representing labels
	//  as the string:  name + "_" + value  would prove problematic if we were to
	//  allow "_" in a future release.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// etag is used for optimistic concurrency control as a way
	//  to help prevent simultaneous updates of a instance configuration from
	//  overwriting each other. It is strongly suggested that systems make use of
	//  the etag in the read-modify-write cycle to perform instance configuration
	//  updates in order to avoid race conditions: An etag is returned in the
	//  response which contains instance configurations, and systems are expected
	//  to put that etag in the request to update instance configuration to ensure
	//  that their change is applied to the same version of the instance
	//  configuration. If no etag is provided in the call to update the instance
	//  configuration, then the existing instance configuration is overwritten
	//  blindly.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.etag
	Etag *string `json:"etag,omitempty"`

	// Allowed values of the "default_leader" schema option for databases in
	//  instances that use this instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.leader_options
	LeaderOptions []string `json:"leaderOptions,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.ReplicaInfo
type ReplicaInfo struct {
	// The location of the serving resources, e.g., "us-central1".
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.location
	Location *string `json:"location,omitempty"`

	// The type of replica.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.type
	Type *string `json:"type,omitempty"`

	// If true, this location is designated as the default leader location where
	//  leader replicas are placed. See the [region types
	//  documentation](https://cloud.google.com/spanner/docs/instances#region_types)
	//  for more details.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.default_leader_location
	DefaultLeaderLocation *bool `json:"defaultLeaderLocation,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.InstanceConfig
type InstanceConfigObservedState struct {
	// Output only. Whether this instance configuration is a Google-managed or
	//  user-managed configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.config_type
	ConfigType *string `json:"configType,omitempty"`

	// Output only. The available optional replicas to choose from for
	//  user-managed configurations. Populated for Google-managed configurations.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.optional_replicas
	OptionalReplicas []ReplicaInfo `json:"optionalReplicas,omitempty"`

	// Output only. If true, the instance configuration is being created or
	//  updated. If false, there are no ongoing operations for the instance
	//  configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The current instance configuration state. Applicable only for
	//  `USER_MANAGED` configurations.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.state
	State *string `json:"state,omitempty"`

	// Output only. Describes whether free instances are available to be created
	//  in this instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.free_instance_availability
	FreeInstanceAvailability *string `json:"freeInstanceAvailability,omitempty"`

	// Output only. The `QuorumType` of the instance configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.quorum_type
	QuorumType *string `json:"quorumType,omitempty"`

	// Output only. The storage limit in bytes per processing unit.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstanceConfig.storage_limit_per_processing_unit
	StorageLimitPerProcessingUnit *int64 `json:"storageLimitPerProcessingUnit,omitempty"`
}
