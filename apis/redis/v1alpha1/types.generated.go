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


// +kcc:proto=google.cloud.redis.v1beta1.Instance
type Instance struct {
	// Required. Unique name of the resource in this scope including project and
	//  location using the form:
	//      `projects/{project_id}/locations/{location_id}/instances/{instance_id}`
	//
	//  Note: Redis instances are managed and addressed at regional level so
	//  location_id here refers to a GCP region; however, users may choose which
	//  specific zone (or collection of zones for cross-zone instances) an instance
	//  should be provisioned in. Refer to [location_id][google.cloud.redis.v1beta1.Instance.location_id] and
	//  [alternative_location_id][google.cloud.redis.v1beta1.Instance.alternative_location_id] fields for more details.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.name
	Name *string `json:"name,omitempty"`

	// An arbitrary and optional user-provided name for the instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Resource labels to represent user provided metadata
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The zone where the instance will be provisioned. If not provided,
	//  the service will choose a zone from the specified region for the instance.
	//  For standard tier, additional nodes will be added across multiple zones for
	//  protection against zonal failures. If specified, at least one node will be
	//  provisioned in this zone.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.location_id
	LocationID *string `json:"locationID,omitempty"`

	// Optional. If specified, at least one node will be provisioned in this zone
	//  in addition to the zone specified in location_id. Only applicable to
	//  standard tier. If provided, it must be a different zone from the one
	//  provided in [location_id]. Additional nodes beyond the first 2 will be
	//  placed in zones selected by the service.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.alternative_location_id
	AlternativeLocationID *string `json:"alternativeLocationID,omitempty"`

	// Optional. The version of Redis software.
	//  If not provided, latest supported version will be used. Currently, the
	//  supported values are:
	//
	//   *   `REDIS_3_2` for Redis 3.2 compatibility
	//   *   `REDIS_4_0` for Redis 4.0 compatibility (default)
	//   *   `REDIS_5_0` for Redis 5.0 compatibility
	//   *   `REDIS_6_X` for Redis 6.x compatibility
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.redis_version
	RedisVersion *string `json:"redisVersion,omitempty"`

	// Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses
	//  that are reserved for this instance. Range must
	//  be unique and non-overlapping with existing subnets in an authorized
	//  network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP
	//  address ranges associated with this private service access connection.
	//  If not provided, the service will choose an unused /29 block, for
	//  example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED
	//  the default block size is /28.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.reserved_ip_range
	ReservedIPRange *string `json:"reservedIPRange,omitempty"`

	// Optional. Additional IP range for node placement. Required when enabling read
	//  replicas on an existing instance. For DIRECT_PEERING mode value must be a
	//  CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value
	//  must be the name of an allocated address range associated with the private
	//  service access connection, or "auto".
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.secondary_ip_range
	SecondaryIPRange *string `json:"secondaryIPRange,omitempty"`

	// Optional. Redis configuration parameters, according to
	//  http://redis.io/topics/config. Currently, the only supported parameters
	//  are:
	//
	//   Redis version 3.2 and newer:
	//
	//   *   maxmemory-policy
	//   *   notify-keyspace-events
	//
	//   Redis version 4.0 and newer:
	//
	//   *   activedefrag
	//   *   lfu-decay-time
	//   *   lfu-log-factor
	//   *   maxmemory-gb
	//
	//   Redis version 5.0 and newer:
	//
	//   *   stream-node-max-bytes
	//   *   stream-node-max-entries
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.redis_configs
	RedisConfigs map[string]string `json:"redisConfigs,omitempty"`

	// Required. The service tier of the instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.tier
	Tier *string `json:"tier,omitempty"`

	// Required. Redis memory size in GiB.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Optional. The full name of the Google Compute Engine
	//  [network](https://cloud.google.com/vpc/docs/vpc) to which the
	//  instance is connected. If left unspecified, the `default` network
	//  will be used.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.authorized_network
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty"`

	// Optional. The network connect mode of the Redis instance.
	//  If not provided, the connect mode defaults to DIRECT_PEERING.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.connect_mode
	ConnectMode *string `json:"connectMode,omitempty"`

	// Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to
	//  "true" AUTH is enabled on the instance. Default value is "false" meaning
	//  AUTH is disabled.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.auth_enabled
	AuthEnabled *bool `json:"authEnabled,omitempty"`

	// Optional. The TLS mode of the Redis instance.
	//  If not provided, TLS is disabled for the instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.transit_encryption_mode
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`

	// Optional. The maintenance policy for the instance. If not provided,
	//  maintenance events can be performed at any time.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.maintenance_policy
	MaintenancePolicy *MaintenancePolicy `json:"maintenancePolicy,omitempty"`

	// Optional. The number of replica nodes. The valid range for the Standard Tier with
	//  read replicas enabled is [1-5] and defaults to 2. If read replicas are not
	//  enabled for a Standard Tier instance, the only valid value is 1 and the
	//  default is 1. The valid value for basic tier is 0 and the default is also
	//  0.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.replica_count
	ReplicaCount *int32 `json:"replicaCount,omitempty"`

	// Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.read_replicas_mode
	ReadReplicasMode *string `json:"readReplicasMode,omitempty"`

	// Optional. Persistence configuration parameters
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.persistence_config
	PersistenceConfig *PersistenceConfig `json:"persistenceConfig,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.MaintenancePolicy
type MaintenancePolicy struct {

	// Optional. Description of what this policy is for. Create/Update methods
	//  return INVALID_ARGUMENT if the length is greater than 512.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenancePolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. Maintenance window that is applied to resources covered by this
	//  policy. Minimum 1. For the current version, the maximum number of
	//  weekly_window is expected to be one.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenancePolicy.weekly_maintenance_window
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.MaintenanceSchedule
type MaintenanceSchedule struct {

	// If the scheduled maintenance can be rescheduled, default is true.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenanceSchedule.can_reschedule
	CanReschedule *bool `json:"canReschedule,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.NodeInfo
type NodeInfo struct {
}

// +kcc:proto=google.cloud.redis.v1beta1.PersistenceConfig
type PersistenceConfig struct {
	// Optional. Controls whether Persistence features are enabled.
	//  If not provided, the existing value will be used.
	// +kcc:proto:field=google.cloud.redis.v1beta1.PersistenceConfig.persistence_mode
	PersistenceMode *string `json:"persistenceMode,omitempty"`

	// Optional. Period between RDB snapshots. Snapshots will be attempted every period
	//  starting from the provided snapshot start time. For example, a start time
	//  of 01/01/2033 06:45 and SIX_HOURS snapshot period will do nothing until
	//  01/01/2033, and then trigger snapshots every day at 06:45, 12:45, 18:45,
	//  and 00:45 the next day, and so on.
	//  If not provided, TWENTY_FOUR_HOURS will be used as default.
	// +kcc:proto:field=google.cloud.redis.v1beta1.PersistenceConfig.rdb_snapshot_period
	RdbSnapshotPeriod *string `json:"rdbSnapshotPeriod,omitempty"`

	// Optional. Date and time that the first snapshot was/will be attempted, and to which
	//  future snapshots will be aligned.
	//  If not provided, the current time will be used.
	// +kcc:proto:field=google.cloud.redis.v1beta1.PersistenceConfig.rdb_snapshot_start_time
	RdbSnapshotStartTime *string `json:"rdbSnapshotStartTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.TlsCertificate
type TlsCertificate struct {
	// Serial number, as extracted from the certificate.
	// +kcc:proto:field=google.cloud.redis.v1beta1.TlsCertificate.serial_number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// PEM representation.
	// +kcc:proto:field=google.cloud.redis.v1beta1.TlsCertificate.cert
	Cert *string `json:"cert,omitempty"`

	// Sha1 Fingerprint of the certificate.
	// +kcc:proto:field=google.cloud.redis.v1beta1.TlsCertificate.sha1_fingerprint
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.WeeklyMaintenanceWindow
type WeeklyMaintenanceWindow struct {
	// Required. The day of week that maintenance updates occur.
	// +kcc:proto:field=google.cloud.redis.v1beta1.WeeklyMaintenanceWindow.day
	Day *string `json:"day,omitempty"`

	// Required. Start time of the window in UTC time.
	// +kcc:proto:field=google.cloud.redis.v1beta1.WeeklyMaintenanceWindow.start_time
	StartTime *TimeOfDay `json:"startTime,omitempty"`
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

// +kcc:proto=google.cloud.redis.v1beta1.Instance
type InstanceObservedState struct {
	// Output only. Hostname or IP address of the exposed Redis endpoint used by
	//   clients to connect to the service.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.host
	Host *string `json:"host,omitempty"`

	// Output only. The port number of the exposed Redis endpoint.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.port
	Port *int32 `json:"port,omitempty"`

	// Output only. The current zone where the Redis primary node is located. In
	//  basic tier, this will always be the same as [location_id]. In
	//  standard tier, this can be the zone of any node in the instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.current_location_id
	CurrentLocationID *string `json:"currentLocationID,omitempty"`

	// Output only. The time the instance was created.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The current state of this instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current status of this
	//  instance, if available.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.status_message
	StatusMessage *string `json:"statusMessage,omitempty"`

	// Output only. Cloud IAM identity used by import / export operations to
	//  transfer data to/from Cloud Storage. Format is
	//  "serviceAccount:<service_account_email>". The value may change over time
	//  for a given instance so should be checked before each import/export
	//  operation.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.persistence_iam_identity
	PersistenceIamIdentity *string `json:"persistenceIamIdentity,omitempty"`

	// Output only. List of server CA certificates for the instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.server_ca_certs
	ServerCaCerts []TlsCertificate `json:"serverCaCerts,omitempty"`

	// Optional. The maintenance policy for the instance. If not provided,
	//  maintenance events can be performed at any time.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.maintenance_policy
	MaintenancePolicy *MaintenancePolicyObservedState `json:"maintenancePolicy,omitempty"`

	// Output only. Date and time of upcoming maintenance events which have been
	//  scheduled.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.maintenance_schedule
	MaintenanceSchedule *MaintenanceSchedule `json:"maintenanceSchedule,omitempty"`

	// Output only. Info per node.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.nodes
	Nodes []NodeInfo `json:"nodes,omitempty"`

	// Output only. Hostname or IP address of the exposed readonly Redis
	//  endpoint. Standard tier only. Targets all healthy replica nodes in
	//  instance. Replication is asynchronous and replica nodes will exhibit some
	//  lag behind the primary. Write requests must target 'host'.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.read_endpoint
	ReadEndpoint *string `json:"readEndpoint,omitempty"`

	// Output only. The port number of the exposed readonly redis
	//  endpoint. Standard tier only. Write requests should target 'port'.
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.read_endpoint_port
	ReadEndpointPort *int32 `json:"readEndpointPort,omitempty"`

	// Optional. Persistence configuration parameters
	// +kcc:proto:field=google.cloud.redis.v1beta1.Instance.persistence_config
	PersistenceConfig *PersistenceConfigObservedState `json:"persistenceConfig,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.MaintenancePolicy
type MaintenancePolicyObservedState struct {
	// Output only. The time when the policy was created.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenancePolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the policy was last updated.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenancePolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Maintenance window that is applied to resources covered by this
	//  policy. Minimum 1. For the current version, the maximum number of
	//  weekly_window is expected to be one.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenancePolicy.weekly_maintenance_window
	WeeklyMaintenanceWindow []WeeklyMaintenanceWindowObservedState `json:"weeklyMaintenanceWindow,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.MaintenanceSchedule
type MaintenanceScheduleObservedState struct {
	// Output only. The start time of any upcoming scheduled maintenance for this instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenanceSchedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The end time of any upcoming scheduled maintenance for this instance.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenanceSchedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The deadline that the maintenance schedule start time can not go beyond,
	//  including reschedule.
	// +kcc:proto:field=google.cloud.redis.v1beta1.MaintenanceSchedule.schedule_deadline_time
	ScheduleDeadlineTime *string `json:"scheduleDeadlineTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.NodeInfo
type NodeInfoObservedState struct {
	// Output only. Node identifying string. e.g. 'node-0', 'node-1'
	// +kcc:proto:field=google.cloud.redis.v1beta1.NodeInfo.id
	ID *string `json:"id,omitempty"`

	// Output only. Location of the node.
	// +kcc:proto:field=google.cloud.redis.v1beta1.NodeInfo.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.PersistenceConfig
type PersistenceConfigObservedState struct {
	// Output only. The next time that a snapshot attempt is scheduled to occur.
	// +kcc:proto:field=google.cloud.redis.v1beta1.PersistenceConfig.rdb_next_snapshot_time
	RdbNextSnapshotTime *string `json:"rdbNextSnapshotTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.TlsCertificate
type TlsCertificateObservedState struct {
	// Output only. The time when the certificate was created in [RFC
	//  3339](https://tools.ietf.org/html/rfc3339) format, for example
	//  `2020-05-18T00:00:00.094Z`.
	// +kcc:proto:field=google.cloud.redis.v1beta1.TlsCertificate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the certificate expires in [RFC
	//  3339](https://tools.ietf.org/html/rfc3339) format, for example
	//  `2020-05-18T00:00:00.094Z`.
	// +kcc:proto:field=google.cloud.redis.v1beta1.TlsCertificate.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.cloud.redis.v1beta1.WeeklyMaintenanceWindow
type WeeklyMaintenanceWindowObservedState struct {
	// Output only. Duration of the maintenance window. The current window is fixed at 1 hour.
	// +kcc:proto:field=google.cloud.redis.v1beta1.WeeklyMaintenanceWindow.duration
	Duration *string `json:"duration,omitempty"`
}
