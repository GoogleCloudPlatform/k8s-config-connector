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


// +kcc:proto=google.cloud.memcache.v1beta2.Instance
type Instance struct {
	// Required. Unique name of the resource in this scope including project and
	//  location using the form:
	//      `projects/{project_id}/locations/{location_id}/instances/{instance_id}`
	//
	//  Note: Memcached instances are managed and addressed at the regional level
	//  so `location_id` here refers to a Google Cloud region; however, users may
	//  choose which zones Memcached nodes should be provisioned in within an
	//  instance. Refer to [zones][google.cloud.memcache.v1beta2.Instance.zones] field for more details.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.name
	Name *string `json:"name,omitempty"`

	// User provided name for the instance, which is only used for display
	//  purposes. Cannot be more than 80 characters.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/vpc/docs/vpc) to which the
	//  instance is connected. If left unspecified, the `default` network
	//  will be used.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.authorized_network
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty"`

	// Zones in which Memcached nodes should be provisioned.
	//  Memcached nodes will be equally distributed across these zones. If not
	//  provided, the service will by default create nodes in all zones in the
	//  region for the instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.zones
	Zones []string `json:"zones,omitempty"`

	// Required. Number of nodes in the Memcached instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Required. Configuration for Memcached nodes.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.node_config
	NodeConfig *Instance_NodeConfig `json:"nodeConfig,omitempty"`

	// The major version of Memcached software.
	//  If not provided, latest supported version will be used. Currently the
	//  latest supported major version is `MEMCACHE_1_5`.
	//  The minor version will be automatically determined by our system based on
	//  the latest supported minor version.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.memcache_version
	MemcacheVersion *string `json:"memcacheVersion,omitempty"`

	// User defined parameters to apply to the memcached process
	//  on each node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.parameters
	Parameters *MemcacheParameters `json:"parameters,omitempty"`

	// List of messages that describe the current state of the Memcached instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.instance_messages
	InstanceMessages []Instance_InstanceMessage `json:"instanceMessages,omitempty"`

	// The maintenance policy for the instance. If not provided,
	//  the maintenance event will be performed based on Memorystore
	//  internal rollout schedule.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.maintenance_policy
	MaintenancePolicy *MaintenancePolicy `json:"maintenancePolicy,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.Instance.InstanceMessage
type Instance_InstanceMessage struct {
	// A code that correspond to one type of user-facing message.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.InstanceMessage.code
	Code *string `json:"code,omitempty"`

	// Message on memcached instance which will be exposed to users.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.InstanceMessage.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.Instance.Node
type Instance_Node struct {

	// User defined parameters currently applied to the node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.parameters
	Parameters *MemcacheParameters `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.Instance.NodeConfig
type Instance_NodeConfig struct {
	// Required. Number of cpus per Memcached node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.NodeConfig.cpu_count
	CpuCount *int32 `json:"cpuCount,omitempty"`

	// Required. Memory size in MiB for each Memcached node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.NodeConfig.memory_size_mb
	MemorySizeMb *int32 `json:"memorySizeMb,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.MaintenancePolicy
type MaintenancePolicy struct {

	// Description of what this policy is for. Create/Update methods
	//  return INVALID_ARGUMENT if the length is greater than 512.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenancePolicy.description
	Description *string `json:"description,omitempty"`

	// Required. Maintenance window that is applied to resources covered by this
	//  policy. Minimum 1. For the current version, the maximum number of
	//  weekly_maintenance_windows is expected to be one.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenancePolicy.weekly_maintenance_window
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.MaintenanceSchedule
type MaintenanceSchedule struct {
}

// +kcc:proto=google.cloud.memcache.v1beta2.MemcacheParameters
type MemcacheParameters struct {

	// User defined set of parameters to use in the memcached process.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MemcacheParameters.params
	Params map[string]string `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.WeeklyMaintenanceWindow
type WeeklyMaintenanceWindow struct {
	// Required. Allows to define schedule that runs specified day of the week.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.WeeklyMaintenanceWindow.day
	Day *string `json:"day,omitempty"`

	// Required. Start time of the window in UTC.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.WeeklyMaintenanceWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Required. Duration of the time window.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.WeeklyMaintenanceWindow.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.type.TimeOfDay
type TimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose
	//  to allow the value "24:00:00" for scenarios like business closing time.
	// +kcc:proto:field=google.type.TimeOfDay.hours
	Hours *int32 `json:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.TimeOfDay.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may
	//  allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.TimeOfDay.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kcc:proto:field=google.type.TimeOfDay.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.Instance
type InstanceObservedState struct {
	// User defined parameters to apply to the memcached process
	//  on each node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.parameters
	Parameters *MemcacheParametersObservedState `json:"parameters,omitempty"`

	// Output only. List of Memcached nodes.
	//  Refer to [Node][google.cloud.memcache.v1beta2.Instance.Node] message for more details.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.memcache_nodes
	MemcacheNodes []Instance_Node `json:"memcacheNodes,omitempty"`

	// Output only. The time the instance was created.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the instance was updated.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of this Memcached instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. The full version of memcached server running on this instance.
	//  System automatically determines the full memcached version for an instance
	//  based on the input MemcacheVersion.
	//  The full version format will be "memcached-1.5.16".
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.memcache_full_version
	MemcacheFullVersion *string `json:"memcacheFullVersion,omitempty"`

	// Output only. Endpoint for the Discovery API.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.discovery_endpoint
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty"`

	// Output only. Returns true if there is an update waiting to be applied
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.update_available
	UpdateAvailable *bool `json:"updateAvailable,omitempty"`

	// The maintenance policy for the instance. If not provided,
	//  the maintenance event will be performed based on Memorystore
	//  internal rollout schedule.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.maintenance_policy
	MaintenancePolicy *MaintenancePolicyObservedState `json:"maintenancePolicy,omitempty"`

	// Output only. Published maintenance schedule.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.maintenance_schedule
	MaintenanceSchedule *MaintenanceSchedule `json:"maintenanceSchedule,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.Instance.Node
type Instance_NodeObservedState struct {
	// Output only. Identifier of the Memcached node. The node id does not
	//  include project or location like the Memcached instance name.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.node_id
	NodeID *string `json:"nodeID,omitempty"`

	// Output only. Location (GCP Zone) for the Memcached node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.zone
	Zone *string `json:"zone,omitempty"`

	// Output only. Current state of the Memcached node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.state
	State *string `json:"state,omitempty"`

	// Output only. Hostname or IP address of the Memcached node used by the
	//  clients to connect to the Memcached server on this node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.host
	Host *string `json:"host,omitempty"`

	// Output only. The port number of the Memcached server on this node.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.port
	Port *int32 `json:"port,omitempty"`

	// Output only. Returns true if there is an update waiting to be applied
	// +kcc:proto:field=google.cloud.memcache.v1beta2.Instance.Node.update_available
	UpdateAvailable *bool `json:"updateAvailable,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.MaintenancePolicy
type MaintenancePolicyObservedState struct {
	// Output only. The time when the policy was created.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenancePolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the policy was updated.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenancePolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.MaintenanceSchedule
type MaintenanceScheduleObservedState struct {
	// Output only. The start time of any upcoming scheduled maintenance for this instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenanceSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The end time of any upcoming scheduled maintenance for this instance.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenanceSchedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The deadline that the maintenance schedule start time can not go beyond,
	//  including reschedule.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MaintenanceSchedule.schedule_deadline_time
	ScheduleDeadlineTime *string `json:"scheduleDeadlineTime,omitempty"`
}

// +kcc:proto=google.cloud.memcache.v1beta2.MemcacheParameters
type MemcacheParametersObservedState struct {
	// Output only. The unique ID associated with this set of parameters. Users
	//  can use this id to determine if the parameters associated with the instance
	//  differ from the parameters associated with the nodes. A discrepancy between
	//  parameter ids can inform users that they may need to take action to apply
	//  parameters on nodes.
	// +kcc:proto:field=google.cloud.memcache.v1beta2.MemcacheParameters.id
	ID *string `json:"id,omitempty"`
}
