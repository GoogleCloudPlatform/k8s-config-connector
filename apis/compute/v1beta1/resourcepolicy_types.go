// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeResourcePolicyGVK = GroupVersion.WithKind("ComputeResourcePolicy")

// ComputeResourcePolicySpec defines the desired state of ComputeResourcePolicy
// +kcc:spec:proto=google.cloud.compute.v1.ResourcePolicy
type ComputeResourcePolicySpec struct {
	// Immutable. Region where resource policy resides.
	// +required
	Region *string `json:"region"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.description
	Description *string `json:"description,omitempty"`

	// Immutable. Replication consistency group for asynchronous disk replication.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.disk_consistency_group_policy
	DiskConsistencyGroupPolicy *ResourcePolicyDiskConsistencyGroupPolicy `json:"diskConsistencyGroupPolicy,omitempty"`

	// Immutable. Resource policy for instances used for placement configuration.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.group_placement_policy
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `json:"groupPlacementPolicy,omitempty"`

	// Immutable. Resource policy for scheduling instance operations.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.instance_schedule_policy
	InstanceSchedulePolicy *ResourcePolicyInstanceSchedulePolicy `json:"instanceSchedulePolicy,omitempty"`

	// Immutable. Policy for creating snapshots of persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.snapshot_schedule_policy
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `json:"snapshotSchedulePolicy,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyDiskConsistencyGroupPolicy
type ResourcePolicyDiskConsistencyGroupPolicy struct {
	// Immutable. Enable disk consistency on the resource policy.
	// +required
	Enabled *bool `json:"enabled"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyGroupPlacementPolicy
type ResourcePolicyGroupPlacementPolicy struct {
	// Immutable. The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyGroupPlacementPolicy.availability_domain_count
	AvailabilityDomainCount *int `json:"availabilityDomainCount,omitempty"`

	// Immutable. Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify 'COLLOCATED' to enable collocation. Can only be specified with 'vm_count'. If compute instances are created
	// with a COLLOCATED policy, then exactly 'vm_count' instances must be created at the same time with the resource policy
	// attached. Possible values: ["COLLOCATED"].
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyGroupPlacementPolicy.collocation
	Collocation *string `json:"collocation,omitempty"`

	// Immutable. Specifies the number of max logical switches.
	MaxDistance *int `json:"maxDistance,omitempty"`

	// Immutable. Number of VMs in this placement group. Google does not recommend that you use this field
	// unless you use a compact policy and you want your policy to work only if it contains this
	// exact number of VMs.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyGroupPlacementPolicy.vm_count
	VmCount *int `json:"vmCount,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy
type ResourcePolicyInstanceSchedulePolicy struct {
	// Immutable. The expiration time of the schedule. The timestamp is an RFC3339 string.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy.expiration_time
	ExpirationTime *string `json:"expirationTime,omitempty"`

	// Immutable. The start time of the schedule. The timestamp is an RFC3339 string.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Immutable. Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy.time_zone
	TimeZone *string `json:"timeZone"`

	// Immutable. Specifies the schedule for starting instances.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy.vm_start_schedule
	VmStartSchedule *ResourcePolicyInstanceSchedulePolicySchedule `json:"vmStartSchedule,omitempty"`

	// Immutable. Specifies the schedule for stopping instances.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicy.vm_stop_schedule
	VmStopSchedule *ResourcePolicyInstanceSchedulePolicySchedule `json:"vmStopSchedule,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicySchedule
type ResourcePolicyInstanceSchedulePolicySchedule struct {
	// Immutable. Specifies the frequency for the operation, using the unix-cron format.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyInstanceSchedulePolicySchedule.schedule
	Schedule *string `json:"schedule"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicy
type ResourcePolicySnapshotSchedulePolicy struct {
	// Immutable. Retention policy applied to snapshots created by this resource policy.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicy.retention_policy
	RetentionPolicy *ResourcePolicySnapshotSchedulePolicyRetentionPolicy `json:"retentionPolicy,omitempty"`

	// Immutable. Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicy.schedule
	Schedule *ResourcePolicySnapshotSchedulePolicySchedule `json:"schedule"`

	// Immutable. Properties with which the snapshots are created, such as labels.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicy.snapshot_properties
	SnapshotProperties *ResourcePolicySnapshotSchedulePolicySnapshotProperties `json:"snapshotProperties,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicyRetentionPolicy
type ResourcePolicySnapshotSchedulePolicyRetentionPolicy struct {
	// Immutable. Maximum age of the snapshot that is allowed to be kept.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicyRetentionPolicy.max_retention_days
	MaxRetentionDays *int `json:"maxRetentionDays"`

	// Immutable. Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted. Default value: "KEEP_AUTO_SNAPSHOTS" Possible values: ["KEEP_AUTO_SNAPSHOTS", "APPLY_RETENTION_POLICY"].
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicyRetentionPolicy.on_source_disk_delete
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySchedule
type ResourcePolicySnapshotSchedulePolicySchedule struct {
	// Immutable. The policy will execute every nth day at the specified time.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySchedule.daily_schedule
	DailySchedule *ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `json:"dailySchedule,omitempty"`

	// Immutable. The policy will execute every nth hour starting at the specified time.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySchedule.hourly_schedule
	HourlySchedule *ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `json:"hourlySchedule,omitempty"`

	// Immutable. Allows specifying a snapshot time for each day of the week.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySchedule.weekly_schedule
	WeeklySchedule *ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `json:"weeklySchedule,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyDailyCycle
type ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule struct {
	// Immutable. Defines a schedule with units measured in days. The value determines how many days
	// pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyDailyCycle.days_in_cycle
	DaysInCycle *int `json:"daysInCycle"`

	// Immutable. This must be in UTC format that resolves to one of
	// 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,
	// both 13:00-5 and 08:00 are valid.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyDailyCycle.start_time
	StartTime *string `json:"startTime"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyHourlyCycle
type ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule struct {
	// Immutable. The number of hours between snapshots.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyHourlyCycle.hours_in_cycle
	HoursInCycle *int `json:"hoursInCycle"`

	// Immutable. Time within the window to start the operations.
	// It must be in an hourly format "HH:MM", where HH : [00-23] and MM : [00] GMT. eg: 21:00.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyHourlyCycle.start_time
	StartTime *string `json:"startTime"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyWeeklyCycle
type ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	// Immutable. May contain up to seven (one for each day of the week) snapshot times.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyWeeklyCycle.day_of_weeks
	DayOfWeeks []ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek `json:"dayOfWeeks"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicyWeeklyCycleDayOfWeek
type ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek struct {
	// Immutable. The day of the week to create the snapshot. e.g. MONDAY Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyWeeklyCycleDayOfWeek.day
	Day *string `json:"day"`

	// Immutable. Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicyWeeklyCycleDayOfWeek.start_time
	StartTime *string `json:"startTime"`
}

// +kcc:proto=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySnapshotProperties
type ResourcePolicySnapshotSchedulePolicySnapshotProperties struct {
	// Immutable. Chain name that the snapshot is created in.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySnapshotProperties.chain_name
	ChainName *string `json:"chainName,omitempty"`

	// Immutable. Whether to perform a 'guest aware' snapshot.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySnapshotProperties.guest_flush
	GuestFlush *bool `json:"guestFlush,omitempty"`

	// Immutable. A set of key-value pairs.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySnapshotProperties.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Cloud Storage bucket location to store the auto snapshot (regional or multi-regional).
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicySnapshotSchedulePolicySnapshotProperties.storage_locations
	StorageLocations []string `json:"storageLocations,omitempty"`
}

// ComputeResourcePolicyStatus defines the config connector machine state of ComputeResourcePolicy
type ComputeResourcePolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeResourcePolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeResourcePolicyObservedState `json:"observedState,omitempty"`

	// The server-defined URL of this resource.
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeResourcePolicyObservedState is the state of the ComputeResourcePolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.ResourcePolicy
type ComputeResourcePolicyObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] The status of resource policy creation.
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeresourcepolicy;gcpcomputeresourcepolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeResourcePolicy is the Schema for the ComputeResourcePolicy API
// +k8s:openapi-gen=true
type ComputeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeResourcePolicySpec   `json:"spec,omitempty"`
	Status ComputeResourcePolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeResourcePolicyList contains a list of ComputeResourcePolicy
type ComputeResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeResourcePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeResourcePolicy{}, &ComputeResourcePolicyList{})
}
