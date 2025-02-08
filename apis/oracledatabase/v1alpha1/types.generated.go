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


// +kcc:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructure
type CloudExadataInfrastructure struct {
	// Identifier. The name of the Exadata Infrastructure resource with the
	//  format:
	//  projects/{project}/locations/{region}/cloudExadataInfrastructures/{cloud_exadata_infrastructure}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.name
	Name *string `json:"name,omitempty"`

	// Optional. User friendly name for this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Google Cloud Platform location where Oracle Exadata is hosted.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.gcp_oracle_zone
	GcpOracleZone *string `json:"gcpOracleZone,omitempty"`

	// Optional. Various properties of the infra.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.properties
	Properties *CloudExadataInfrastructureProperties `json:"properties,omitempty"`

	// Optional. Labels or tags associated with the resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties
type CloudExadataInfrastructureProperties struct {

	// Optional. The number of compute servers for the Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.compute_count
	ComputeCount *int32 `json:"computeCount,omitempty"`

	// Optional. The number of Cloud Exadata storage servers for the Exadata
	//  Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.storage_count
	StorageCount *int32 `json:"storageCount,omitempty"`

	// Optional. The total storage allocated to the Exadata Infrastructure
	//  resource, in gigabytes (GB).
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.total_storage_size_gb
	TotalStorageSizeGB *int32 `json:"totalStorageSizeGB,omitempty"`

	// Optional. Maintenance window for repair.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.maintenance_window
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Required. The shape of the Exadata Infrastructure. The shape determines the
	//  amount of CPU, storage, and memory resources allocated to the instance.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.shape
	Shape *string `json:"shape,omitempty"`

	// Optional. The list of customer contacts.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.customer_contacts
	CustomerContacts []CustomerContact `json:"customerContacts,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CustomerContact
type CustomerContact struct {
	// Required. The email address used by Oracle to send notifications regarding
	//  databases and infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CustomerContact.email
	Email *string `json:"email,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.MaintenanceWindow
type MaintenanceWindow struct {
	// Optional. The maintenance window scheduling preference.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.preference
	Preference *string `json:"preference,omitempty"`

	// Optional. Months during the year when maintenance should be performed.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.months
	Months []string `json:"months,omitempty"`

	// Optional. Weeks during the month when maintenance should be performed.
	//  Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a
	//  duration of 7 days. Weeks start and end based on calendar dates, not days
	//  of the week.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.weeks_of_month
	WeeksOfMonth []int32 `json:"weeksOfMonth,omitempty"`

	// Optional. Days during the week when maintenance should be performed.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.days_of_week
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	// Optional. The window of hours during the day when maintenance should be
	//  performed. The window is a 4 hour slot. Valid values are:
	//    0 - represents time slot 0:00 - 3:59 UTC
	//    4 - represents time slot 4:00 - 7:59 UTC
	//    8 - represents time slot 8:00 - 11:59 UTC
	//    12 - represents time slot 12:00 - 15:59 UTC
	//    16 - represents time slot 16:00 - 19:59 UTC
	//    20 - represents time slot 20:00 - 23:59 UTC
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.hours_of_day
	HoursOfDay []int32 `json:"hoursOfDay,omitempty"`

	// Optional. Lead time window allows user to set a lead time to prepare for a
	//  down time. The lead time is in weeks and valid value is between 1 to 4.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.lead_time_week
	LeadTimeWeek *int32 `json:"leadTimeWeek,omitempty"`

	// Optional. Cloud CloudExadataInfrastructure node patching method, either
	//  "ROLLING"
	//   or "NONROLLING". Default value is ROLLING.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.patching_mode
	PatchingMode *string `json:"patchingMode,omitempty"`

	// Optional. Determines the amount of time the system will wait before the
	//  start of each database server patching operation. Custom action timeout is
	//  in minutes and valid value is between 15 to 120 (inclusive).
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.custom_action_timeout_mins
	CustomActionTimeoutMins *int32 `json:"customActionTimeoutMins,omitempty"`

	// Optional. If true, enables the configuration of a custom action timeout
	//  (waiting period) between database server patching operations.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.MaintenanceWindow.is_custom_action_timeout_enabled
	IsCustomActionTimeoutEnabled *bool `json:"isCustomActionTimeoutEnabled,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructure
type CloudExadataInfrastructureObservedState struct {
	// Output only. Entitlement ID of the private offer against which this
	//  infrastructure resource is provisioned.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Optional. Various properties of the infra.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.properties
	Properties *CloudExadataInfrastructurePropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the Exadata Infrastructure was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties
type CloudExadataInfrastructurePropertiesObservedState struct {
	// Output only. OCID of created infra.
	//  https://docs.oracle.com/en-us/iaas/Content/General/Concepts/identifiers.htm#Oracle
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.ocid
	Ocid *string `json:"ocid,omitempty"`

	// Output only. The available storage can be allocated to the Exadata
	//  Infrastructure resource, in gigabytes (GB).
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.available_storage_size_gb
	AvailableStorageSizeGB *int32 `json:"availableStorageSizeGB,omitempty"`

	// Output only. The current lifecycle state of the Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.state
	State *string `json:"state,omitempty"`

	// Output only. Deep link to the OCI console to view this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.oci_url
	OciURL *string `json:"ociURL,omitempty"`

	// Output only. The number of enabled CPU cores.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.cpu_count
	CpuCount *int32 `json:"cpuCount,omitempty"`

	// Output only. The total number of CPU cores available.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.max_cpu_count
	MaxCpuCount *int32 `json:"maxCpuCount,omitempty"`

	// Output only. The memory allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Output only. The total memory available in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.max_memory_gb
	MaxMemoryGB *int32 `json:"maxMemoryGB,omitempty"`

	// Output only. The local node storage allocated in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.db_node_storage_size_gb
	DbNodeStorageSizeGB *int32 `json:"dbNodeStorageSizeGB,omitempty"`

	// Output only. The total local node storage available in GBs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.max_db_node_storage_size_gb
	MaxDbNodeStorageSizeGB *int32 `json:"maxDbNodeStorageSizeGB,omitempty"`

	// Output only. Size, in terabytes, of the DATA disk group.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.data_storage_size_tb
	DataStorageSizeTb *float64 `json:"dataStorageSizeTb,omitempty"`

	// Output only. The total available DATA disk group size.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.max_data_storage_tb
	MaxDataStorageTb *float64 `json:"maxDataStorageTb,omitempty"`

	// Output only. The requested number of additional storage servers activated
	//  for the Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.activated_storage_count
	ActivatedStorageCount *int32 `json:"activatedStorageCount,omitempty"`

	// Output only. The requested number of additional storage servers for the
	//  Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.additional_storage_count
	AdditionalStorageCount *int32 `json:"additionalStorageCount,omitempty"`

	// Output only. The software version of the database servers (dom0) in the
	//  Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.db_server_version
	DbServerVersion *string `json:"dbServerVersion,omitempty"`

	// Output only. The software version of the storage servers (cells) in the
	//  Exadata Infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.storage_server_version
	StorageServerVersion *string `json:"storageServerVersion,omitempty"`

	// Output only. The OCID of the next maintenance run.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.next_maintenance_run_id
	NextMaintenanceRunID *string `json:"nextMaintenanceRunID,omitempty"`

	// Output only. The time when the next maintenance run will occur.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.next_maintenance_run_time
	NextMaintenanceRunTime *string `json:"nextMaintenanceRunTime,omitempty"`

	// Output only. The time when the next security maintenance run will occur.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.next_security_maintenance_run_time
	NextSecurityMaintenanceRunTime *string `json:"nextSecurityMaintenanceRunTime,omitempty"`

	// Output only. The monthly software version of the storage servers (cells)
	//  in the Exadata Infrastructure. Example: 20.1.15
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.monthly_storage_server_version
	MonthlyStorageServerVersion *string `json:"monthlyStorageServerVersion,omitempty"`

	// Output only. The monthly software version of the database servers (dom0)
	//  in the Exadata Infrastructure. Example: 20.1.15
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructureProperties.monthly_db_server_version
	MonthlyDbServerVersion *string `json:"monthlyDbServerVersion,omitempty"`
}
