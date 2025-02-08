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


// +kcc:proto=google.cloud.vmmigration.v1.UtilizationReport
type UtilizationReport struct {

	// The report display name, as assigned by the user.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Time frame of the report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.time_frame
	TimeFrame *string `json:"timeFrame,omitempty"`

	// List of utilization information per VM.
	//  When sent as part of the request, the "vm_id" field is used in order to
	//  specify which VMs to include in the report. In that case all other fields
	//  are ignored.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.vms
	Vms []VmUtilizationInfo `json:"vms,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmUtilizationInfo
type VmUtilizationInfo struct {
	// The description of the VM in a Source of type Vmware.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationInfo.vmware_vm_details
	VmwareVmDetails *VmwareVmDetails `json:"vmwareVmDetails,omitempty"`

	// The VM's ID in the source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationInfo.vm_id
	VmID *string `json:"vmID,omitempty"`

	// Utilization metrics for this VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationInfo.utilization
	Utilization *VmUtilizationMetrics `json:"utilization,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmUtilizationMetrics
type VmUtilizationMetrics struct {
	// Max CPU usage, percent.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.cpu_max_percent
	CpuMaxPercent *int32 `json:"cpuMaxPercent,omitempty"`

	// Average CPU usage, percent.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.cpu_average_percent
	CpuAveragePercent *int32 `json:"cpuAveragePercent,omitempty"`

	// Max memory usage, percent.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.memory_max_percent
	MemoryMaxPercent *int32 `json:"memoryMaxPercent,omitempty"`

	// Average memory usage, percent.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.memory_average_percent
	MemoryAveragePercent *int32 `json:"memoryAveragePercent,omitempty"`

	// Max disk IO rate, in kilobytes per second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.disk_io_rate_max_kbps
	DiskIoRateMaxKbps *int64 `json:"diskIoRateMaxKbps,omitempty"`

	// Average disk IO rate, in kilobytes per second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.disk_io_rate_average_kbps
	DiskIoRateAverageKbps *int64 `json:"diskIoRateAverageKbps,omitempty"`

	// Max network throughput (combined transmit-rates and receive-rates), in
	//  kilobytes per second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.network_throughput_max_kbps
	NetworkThroughputMaxKbps *int64 `json:"networkThroughputMaxKbps,omitempty"`

	// Average network throughput (combined transmit-rates and receive-rates), in
	//  kilobytes per second.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationMetrics.network_throughput_average_kbps
	NetworkThroughputAverageKbps *int64 `json:"networkThroughputAverageKbps,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmwareVmDetails
type VmwareVmDetails struct {
	// The VM's id in the source (note that this is not the MigratingVm's id).
	//  This is the moref id of the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.vm_id
	VmID *string `json:"vmID,omitempty"`

	// The id of the vCenter's datacenter this VM is contained in.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.datacenter_id
	DatacenterID *string `json:"datacenterID,omitempty"`

	// The descriptive name of the vCenter's datacenter this VM is contained in.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.datacenter_description
	DatacenterDescription *string `json:"datacenterDescription,omitempty"`

	// The unique identifier of the VM in vCenter.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.uuid
	Uuid *string `json:"uuid,omitempty"`

	// The display name of the VM. Note that this is not necessarily unique.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The power state of the VM at the moment list was taken.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.power_state
	PowerState *string `json:"powerState,omitempty"`

	// The number of cpus in the VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.cpu_count
	CpuCount *int32 `json:"cpuCount,omitempty"`

	// The size of the memory of the VM in MB.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.memory_mb
	MemoryMb *int32 `json:"memoryMb,omitempty"`

	// The number of disks the VM has.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.disk_count
	DiskCount *int32 `json:"diskCount,omitempty"`

	// The total size of the storage allocated to the VM in MB.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.committed_storage_mb
	CommittedStorageMb *int64 `json:"committedStorageMb,omitempty"`

	// The VM's OS. See for example
	//  https://vdc-repo.vmware.com/vmwb-repository/dcr-public/da47f910-60ac-438b-8b9b-6122f4d14524/16b7274a-bf8b-4b4c-a05e-746f2aa93c8c/doc/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html
	//  for types of strings this might hold.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.guest_description
	GuestDescription *string `json:"guestDescription,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.UtilizationReport
type UtilizationReportObservedState struct {
	// Output only. The report unique name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.name
	Name *string `json:"name,omitempty"`

	// Output only. Current state of the report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.state
	State *string `json:"state,omitempty"`

	// Output only. The time the state was last set.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. Provides details on the state of the report in case of an
	//  error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.error
	Error *Status `json:"error,omitempty"`

	// Output only. The time the report was created (this refers to the time of
	//  the request, not the time the report creation completed).
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The point in time when the time frame ends. Notice that the
	//  time frame is counted backwards. For instance if the "frame_end_time" value
	//  is 2021/01/20 and the time frame is WEEK then the report covers the week
	//  between 2021/01/20 and 2021/01/14.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.frame_end_time
	FrameEndTime *string `json:"frameEndTime,omitempty"`

	// Output only. Total number of VMs included in the report.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.vm_count
	VmCount *int32 `json:"vmCount,omitempty"`

	// List of utilization information per VM.
	//  When sent as part of the request, the "vm_id" field is used in order to
	//  specify which VMs to include in the report. In that case all other fields
	//  are ignored.
	// +kcc:proto:field=google.cloud.vmmigration.v1.UtilizationReport.vms
	Vms []VmUtilizationInfoObservedState `json:"vms,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmUtilizationInfo
type VmUtilizationInfoObservedState struct {
	// The description of the VM in a Source of type Vmware.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmUtilizationInfo.vmware_vm_details
	VmwareVmDetails *VmwareVmDetailsObservedState `json:"vmwareVmDetails,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmwareVmDetails
type VmwareVmDetailsObservedState struct {
	// Output only. The VM Boot Option.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareVmDetails.boot_option
	BootOption *string `json:"bootOption,omitempty"`
}
