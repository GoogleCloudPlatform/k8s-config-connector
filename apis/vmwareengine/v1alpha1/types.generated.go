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


// +kcc:proto=google.cloud.vmwareengine.v1.AutoscalingSettings
type AutoscalingSettings struct {

	// TODO: unsupported map type with key string and value message


	// Optional. Minimum number of nodes of any type in a cluster.
	//  If not specified the default limits apply.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.min_cluster_node_count
	MinClusterNodeCount *int32 `json:"minClusterNodeCount,omitempty"`

	// Optional. Maximum number of nodes of any type in a cluster.
	//  If not specified the default limits apply.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.max_cluster_node_count
	MaxClusterNodeCount *int32 `json:"maxClusterNodeCount,omitempty"`

	// Optional. The minimum duration between consecutive autoscale operations.
	//  It starts once addition or removal of nodes is fully completed.
	//  Defaults to 30 minutes if not specified. Cool down period must be in whole
	//  minutes (for example, 30, 31, 50, 180 minutes).
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.cool_down_period
	CoolDownPeriod *string `json:"coolDownPeriod,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy
type AutoscalingSettings_AutoscalingPolicy struct {
	// Required. The canonical identifier of the node type to add or remove.
	//  Corresponds to the `NodeType`.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.node_type_id
	NodeTypeID *string `json:"nodeTypeID,omitempty"`

	// Required. Number of nodes to add to a cluster during a scale-out
	//  operation. Must be divisible by 2 for stretched clusters. During a
	//  scale-in operation only one node (or 2 for stretched clusters) are
	//  removed in a single iteration.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.scale_out_size
	ScaleOutSize *int32 `json:"scaleOutSize,omitempty"`

	// Optional. Utilization thresholds pertaining to CPU utilization.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.cpu_thresholds
	CpuThresholds *AutoscalingSettings_Thresholds `json:"cpuThresholds,omitempty"`

	// Optional. Utilization thresholds pertaining to amount of granted memory.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.granted_memory_thresholds
	GrantedMemoryThresholds *AutoscalingSettings_Thresholds `json:"grantedMemoryThresholds,omitempty"`

	// Optional. Utilization thresholds pertaining to amount of consumed memory.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.consumed_memory_thresholds
	ConsumedMemoryThresholds *AutoscalingSettings_Thresholds `json:"consumedMemoryThresholds,omitempty"`

	// Optional. Utilization thresholds pertaining to amount of consumed
	//  storage.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.AutoscalingPolicy.storage_thresholds
	StorageThresholds *AutoscalingSettings_Thresholds `json:"storageThresholds,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.AutoscalingSettings.Thresholds
type AutoscalingSettings_Thresholds struct {
	// Required. The utilization triggering the scale-out operation in percent.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.Thresholds.scale_out
	ScaleOut *int32 `json:"scaleOut,omitempty"`

	// Required. The utilization triggering the scale-in operation in percent.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.AutoscalingSettings.Thresholds.scale_in
	ScaleIn *int32 `json:"scaleIn,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Cluster
type Cluster struct {

	// Optional. Configuration of the autoscaling applied to this cluster.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.autoscaling_settings
	AutoscalingSettings *AutoscalingSettings `json:"autoscalingSettings,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Optional. Configuration of a stretched cluster. Required for clusters that
	//  belong to a STRETCHED private cloud.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.stretched_cluster_config
	StretchedClusterConfig *StretchedClusterConfig `json:"stretchedClusterConfig,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.NodeTypeConfig
type NodeTypeConfig struct {
	// Required. The number of nodes of this type in the cluster
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeTypeConfig.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Optional. Customized number of cores available to each node of the type.
	//  This number must always be one of `nodeType.availableCustomCoreCounts`.
	//  If zero is provided max value from `nodeType.availableCustomCoreCounts`
	//  will be used.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NodeTypeConfig.custom_core_count
	CustomCoreCount *int32 `json:"customCoreCount,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.StretchedClusterConfig
type StretchedClusterConfig struct {
	// Required. Zone that will remain operational when connection between the two
	//  zones is lost. Specify the resource name of a zone that belongs to the
	//  region of the private cloud. For example:
	//  `projects/{project}/locations/europe-west3-a` where `{project}` can either
	//  be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.StretchedClusterConfig.preferred_location
	PreferredLocation *string `json:"preferredLocation,omitempty"`

	// Required. Additional zone for a higher level of availability and load
	//  balancing. Specify the resource name of a zone that belongs to the region
	//  of the private cloud. For example:
	//  `projects/{project}/locations/europe-west3-b` where `{project}` can either
	//  be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.StretchedClusterConfig.secondary_location
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.Cluster
type ClusterObservedState struct {
	// Output only. The resource name of this cluster.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/clusters/my-cluster`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// Output only. True if the cluster is a management cluster; false otherwise.
	//  There can only be one management cluster in a private cloud
	//  and it has to be the first one.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.management
	Management *bool `json:"management,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.Cluster.uid
	Uid *string `json:"uid,omitempty"`
}
