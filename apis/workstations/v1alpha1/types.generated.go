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


// +kcc:proto=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory
type WorkstationConfig_EphemeralDirectory struct {
	// An EphemeralDirectory backed by a Compute Engine persistent disk.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.gce_pd
	GCEPD *WorkstationConfig_EphemeralDirectory_GcePersistentDisk `json:"gcePD,omitempty"`

	// Required. Location of this directory in the running workstation.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.mount_path
	MountPath *string `json:"mountPath,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk
type WorkstationConfig_EphemeralDirectory_GcePersistentDisk struct {
	// Optional. Type of the disk to use. Defaults to `"pd-standard"`.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Optional. Name of the snapshot to use as the source for the disk. Must
	//  be empty if
	//  [source_image][google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_image]
	//  is set. Updating
	//  [source_snapshot][google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_snapshot]
	//  will update content in the ephemeral directory after the workstation is
	//  restarted. This field is mutable.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_snapshot
	SourceSnapshot *string `json:"sourceSnapshot,omitempty"`

	// Optional. Name of the disk image to use as the source for the disk.
	//  Must be empty if
	//  [source_snapshot][google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_snapshot]
	//  is set. Updating
	//  [source_image][google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_image]
	//  will update content in the ephemeral directory after the workstation is
	//  restarted. This field is mutable.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_image
	SourceImage *string `json:"sourceImage,omitempty"`

	// Optional. Whether the disk is read only. If true, the disk may be
	//  shared by multiple VMs and
	//  [source_snapshot][google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.source_snapshot]
	//  must be set.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.EphemeralDirectory.GcePersistentDisk.read_only
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1beta.WorkstationConfig.Host.GceInstance.Accelerator
type WorkstationConfig_Host_GceInstance_Accelerator struct {
	// Optional. Type of accelerator resource to attach to the instance, for
	//  example,
	//  `"nvidia-tesla-p100"`.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.Host.GceInstance.Accelerator.type
	Type *string `json:"type,omitempty"`

	// Optional. Number of accelerator cards exposed to the instance.
	// +kcc:proto:field=google.cloud.workstations.v1beta.WorkstationConfig.Host.GceInstance.Accelerator.count
	Count *int32 `json:"count,omitempty"`
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
