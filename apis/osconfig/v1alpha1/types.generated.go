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


// +kcc:proto=google.cloud.osconfig.v1.FixedOrPercent
type FixedOrPercent struct {
	// Specifies a fixed value.
	// +kcc:proto:field=google.cloud.osconfig.v1.FixedOrPercent.fixed
	Fixed *int32 `json:"fixed,omitempty"`

	// Specifies the relative value defined as a percentage, which will be
	//  multiplied by a reference value.
	// +kcc:proto:field=google.cloud.osconfig.v1.FixedOrPercent.percent
	Percent *int32 `json:"percent,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy
type OSPolicy struct {
	// Required. The id of the OS policy with the following restrictions:
	//
	//  * Must contain only lowercase letters, numbers, and hyphens.
	//  * Must start with a letter.
	//  * Must be between 1-63 characters.
	//  * Must end with a number or a letter.
	//  * Must be unique within the assignment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.id
	ID *string `json:"id,omitempty"`

	// Policy description.
	//  Length of the description is limited to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.description
	Description *string `json:"description,omitempty"`

	// Required. Policy mode
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.mode
	Mode *string `json:"mode,omitempty"`

	// Required. List of resource groups for the policy.
	//  For a particular VM, resource groups are evaluated in the order specified
	//  and the first resource group that is applicable is selected and the rest
	//  are ignored.
	//
	//  If none of the resource groups are applicable for a VM, the VM is
	//  considered to be non-compliant w.r.t this policy. This behavior can be
	//  toggled by the flag `allow_no_resource_group_match`
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.resource_groups
	ResourceGroups []OSPolicy_ResourceGroup `json:"resourceGroups,omitempty"`

	// This flag determines the OS policy compliance status when none of the
	//  resource groups within the policy are applicable for a VM. Set this value
	//  to `true` if the policy needs to be reported as compliant even if the
	//  policy has nothing to validate or enforce.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.allow_no_resource_group_match
	AllowNoResourceGroupMatch *bool `json:"allowNoResourceGroupMatch,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.InventoryFilter
type OSPolicy_InventoryFilter struct {
	// Required. The OS short name
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.InventoryFilter.os_short_name
	OsShortName *string `json:"osShortName,omitempty"`

	// The OS version
	//
	//  Prefix matches are supported if asterisk(*) is provided as the
	//  last character. For example, to match all versions with a major
	//  version of `7`, specify the following value for this field `7.*`
	//
	//  An empty string matches all OS versions.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.InventoryFilter.os_version
	OsVersion *string `json:"osVersion,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource
type OSPolicy_Resource struct {
	// Required. The id of the resource with the following restrictions:
	//
	//  * Must contain only lowercase letters, numbers, and hyphens.
	//  * Must start with a letter.
	//  * Must be between 1-63 characters.
	//  * Must end with a number or a letter.
	//  * Must be unique within the OS policy.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.id
	ID *string `json:"id,omitempty"`

	// Package resource
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.pkg
	Pkg *OSPolicy_Resource_PackageResource `json:"pkg,omitempty"`

	// Package repository resource
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.repository
	Repository *OSPolicy_Resource_RepositoryResource `json:"repository,omitempty"`

	// Exec resource
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.exec
	Exec *OSPolicy_Resource_ExecResource `json:"exec,omitempty"`

	// File resource
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.file
	File *OSPolicy_Resource_FileResource `json:"file,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource
type OSPolicy_Resource_ExecResource struct {
	// Required. What to run to validate this resource is in the desired
	//  state. An exit code of 100 indicates "in desired state", and exit code
	//  of 101 indicates "not in desired state". Any other exit code indicates
	//  a failure running validate.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.validate
	Validate *OSPolicy_Resource_ExecResource_Exec `json:"validate,omitempty"`

	// What to run to bring this resource into the desired state.
	//  An exit code of 100 indicates "success", any other exit code indicates
	//  a failure running enforce.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.enforce
	Enforce *OSPolicy_Resource_ExecResource_Exec `json:"enforce,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec
type OSPolicy_Resource_ExecResource_Exec struct {
	// A remote or local file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec.file
	File *OSPolicy_Resource_File `json:"file,omitempty"`

	// An inline script.
	//  The size of the script is limited to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec.script
	Script *string `json:"script,omitempty"`

	// Optional arguments to pass to the source during execution.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec.args
	Args []string `json:"args,omitempty"`

	// Required. The script interpreter to use.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec.interpreter
	Interpreter *string `json:"interpreter,omitempty"`

	// Only recorded for enforce Exec.
	//  Path to an output file (that is created by this Exec) whose
	//  content will be recorded in OSPolicyResourceCompliance after a
	//  successful run. Absence or failure to read this file will result in
	//  this ExecResource being non-compliant. Output file size is limited to
	//  100K bytes.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.ExecResource.Exec.output_file_path
	OutputFilePath *string `json:"outputFilePath,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.File
type OSPolicy_Resource_File struct {
	// A generic remote file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.remote
	Remote *OSPolicy_Resource_File_Remote `json:"remote,omitempty"`

	// A Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.gcs
	Gcs *OSPolicy_Resource_File_Gcs `json:"gcs,omitempty"`

	// A local path within the VM to use.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.local_path
	LocalPath *string `json:"localPath,omitempty"`

	// Defaults to false. When false, files are subject to validations
	//  based on the file type:
	//
	//  Remote: A checksum must be specified.
	//  Cloud Storage: An object generation number must be specified.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.allow_insecure
	AllowInsecure *bool `json:"allowInsecure,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.File.Gcs
type OSPolicy_Resource_File_Gcs struct {
	// Required. Bucket of the Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.Gcs.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Required. Name of the Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.Gcs.object
	Object *string `json:"object,omitempty"`

	// Generation number of the Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.Gcs.generation
	Generation *int64 `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.File.Remote
type OSPolicy_Resource_File_Remote struct {
	// Required. URI from which to fetch the object. It should contain both
	//  the protocol and path following the format `{protocol}://{location}`.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.Remote.uri
	URI *string `json:"uri,omitempty"`

	// SHA256 checksum of the remote file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.File.Remote.sha256_checksum
	Sha256Checksum *string `json:"sha256Checksum,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource
type OSPolicy_Resource_FileResource struct {
	// A remote or local source.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource.file
	File *OSPolicy_Resource_File `json:"file,omitempty"`

	// A a file with this content.
	//  The size of the content is limited to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource.content
	Content *string `json:"content,omitempty"`

	// Required. The absolute path of the file within the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource.path
	Path *string `json:"path,omitempty"`

	// Required. Desired state of the file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource.state
	State *string `json:"state,omitempty"`

	// Consists of three octal digits which represent, in
	//  order, the permissions of the owner, group, and other users for the
	//  file (similarly to the numeric mode used in the linux chmod
	//  utility). Each digit represents a three bit number with the 4 bit
	//  corresponding to the read permissions, the 2 bit corresponds to the
	//  write bit, and the one bit corresponds to the execute permission.
	//  Default behavior is 755.
	//
	//  Below are some examples of permissions and their associated values:
	//  read, write, and execute: 7
	//  read and execute: 5
	//  read and write: 6
	//  read only: 4
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.FileResource.permissions
	Permissions *string `json:"permissions,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource
type OSPolicy_Resource_PackageResource struct {
	// Required. The desired state the agent should maintain for this package.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.desired_state
	DesiredState *string `json:"desiredState,omitempty"`

	// A package managed by Apt.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.apt
	Apt *OSPolicy_Resource_PackageResource_APT `json:"apt,omitempty"`

	// A deb package file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.deb
	Deb *OSPolicy_Resource_PackageResource_Deb `json:"deb,omitempty"`

	// A package managed by YUM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.yum
	Yum *OSPolicy_Resource_PackageResource_YUM `json:"yum,omitempty"`

	// A package managed by Zypper.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.zypper
	Zypper *OSPolicy_Resource_PackageResource_Zypper `json:"zypper,omitempty"`

	// An rpm package file.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.rpm
	Rpm *OSPolicy_Resource_PackageResource_RPM `json:"rpm,omitempty"`

	// A package managed by GooGet.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.googet
	Googet *OSPolicy_Resource_PackageResource_GooGet `json:"googet,omitempty"`

	// An MSI package.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.msi
	Msi *OSPolicy_Resource_PackageResource_MSI `json:"msi,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.APT
type OSPolicy_Resource_PackageResource_APT struct {
	// Required. Package name.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.APT.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.Deb
type OSPolicy_Resource_PackageResource_Deb struct {
	// Required. A deb package.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.Deb.source
	Source *OSPolicy_Resource_File `json:"source,omitempty"`

	// Whether dependencies should also be installed.
	//  - install when false: `dpkg -i package`
	//  - install when true: `apt-get update && apt-get -y install
	//  package.deb`
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.Deb.pull_deps
	PullDeps *bool `json:"pullDeps,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.GooGet
type OSPolicy_Resource_PackageResource_GooGet struct {
	// Required. Package name.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.GooGet.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.MSI
type OSPolicy_Resource_PackageResource_MSI struct {
	// Required. The MSI package.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.MSI.source
	Source *OSPolicy_Resource_File `json:"source,omitempty"`

	// Additional properties to use during installation.
	//  This should be in the format of Property=Setting.
	//  Appended to the defaults of `ACTION=INSTALL
	//  REBOOT=ReallySuppress`.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.MSI.properties
	Properties []string `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.RPM
type OSPolicy_Resource_PackageResource_RPM struct {
	// Required. An rpm package.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.RPM.source
	Source *OSPolicy_Resource_File `json:"source,omitempty"`

	// Whether dependencies should also be installed.
	//  - install when false: `rpm --upgrade --replacepkgs package.rpm`
	//  - install when true: `yum -y install package.rpm` or
	//  `zypper -y install package.rpm`
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.RPM.pull_deps
	PullDeps *bool `json:"pullDeps,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.YUM
type OSPolicy_Resource_PackageResource_YUM struct {
	// Required. Package name.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.YUM.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.Zypper
type OSPolicy_Resource_PackageResource_Zypper struct {
	// Required. Package name.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.PackageResource.Zypper.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource
type OSPolicy_Resource_RepositoryResource struct {
	// An Apt Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.apt
	Apt *OSPolicy_Resource_RepositoryResource_AptRepository `json:"apt,omitempty"`

	// A Yum Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.yum
	Yum *OSPolicy_Resource_RepositoryResource_YumRepository `json:"yum,omitempty"`

	// A Zypper Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.zypper
	Zypper *OSPolicy_Resource_RepositoryResource_ZypperRepository `json:"zypper,omitempty"`

	// A Goo Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.goo
	Goo *OSPolicy_Resource_RepositoryResource_GooRepository `json:"goo,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository
type OSPolicy_Resource_RepositoryResource_AptRepository struct {
	// Required. Type of archive files in this repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository.archive_type
	ArchiveType *string `json:"archiveType,omitempty"`

	// Required. URI for this repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository.uri
	URI *string `json:"uri,omitempty"`

	// Required. Distribution of this repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository.distribution
	Distribution *string `json:"distribution,omitempty"`

	// Required. List of components for this repository. Must contain at
	//  least one item.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository.components
	Components []string `json:"components,omitempty"`

	// URI of the key file for this repository. The agent maintains a
	//  keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg`.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.AptRepository.gpg_key
	GpgKey *string `json:"gpgKey,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.GooRepository
type OSPolicy_Resource_RepositoryResource_GooRepository struct {
	// Required. The name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.GooRepository.name
	Name *string `json:"name,omitempty"`

	// Required. The url of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.GooRepository.url
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.YumRepository
type OSPolicy_Resource_RepositoryResource_YumRepository struct {
	// Required. A one word, unique name for this repository. This is  the
	//  `repo id` in the yum config file and also the `display_name` if
	//  `display_name` is omitted. This id is also used as the unique
	//  identifier when checking for resource conflicts.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.YumRepository.id
	ID *string `json:"id,omitempty"`

	// The display name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.YumRepository.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The location of the repository directory.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.YumRepository.base_url
	BaseURL *string `json:"baseURL,omitempty"`

	// URIs of GPG keys.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.YumRepository.gpg_keys
	GpgKeys []string `json:"gpgKeys,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.ZypperRepository
type OSPolicy_Resource_RepositoryResource_ZypperRepository struct {
	// Required. A one word, unique name for this repository. This is the
	//  `repo id` in the zypper config file and also the `display_name` if
	//  `display_name` is omitted. This id is also used as the unique
	//  identifier when checking for GuestPolicy conflicts.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.ZypperRepository.id
	ID *string `json:"id,omitempty"`

	// The display name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.ZypperRepository.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The location of the repository directory.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.ZypperRepository.base_url
	BaseURL *string `json:"baseURL,omitempty"`

	// URIs of GPG keys.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.Resource.RepositoryResource.ZypperRepository.gpg_keys
	GpgKeys []string `json:"gpgKeys,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicy.ResourceGroup
type OSPolicy_ResourceGroup struct {
	// List of inventory filters for the resource group.
	//
	//  The resources in this resource group are applied to the target VM if it
	//  satisfies at least one of the following inventory filters.
	//
	//  For example, to apply this resource group to VMs running either `RHEL` or
	//  `CentOS` operating systems, specify 2 items for the list with following
	//  values:
	//  inventory_filters[0].os_short_name='rhel' and
	//  inventory_filters[1].os_short_name='centos'
	//
	//  If the list is empty, this resource group will be applied to the target
	//  VM unconditionally.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.ResourceGroup.inventory_filters
	InventoryFilters []OSPolicy_InventoryFilter `json:"inventoryFilters,omitempty"`

	// Required. List of resources configured for this resource group.
	//  The resources are executed in the exact order specified here.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicy.ResourceGroup.resources
	Resources []OSPolicy_Resource `json:"resources,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment
type OSPolicyAssignment struct {
	// Resource name.
	//
	//  Format:
	//  `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}`
	//
	//  This field is ignored when you create an OS policy assignment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.name
	Name *string `json:"name,omitempty"`

	// OS policy assignment description.
	//  Length of the description is limited to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.description
	Description *string `json:"description,omitempty"`

	// Required. List of OS policies to be applied to the VMs.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.os_policies
	OsPolicies []OSPolicy `json:"osPolicies,omitempty"`

	// Required. Filter to select VMs.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.instance_filter
	InstanceFilter *OSPolicyAssignment_InstanceFilter `json:"instanceFilter,omitempty"`

	// Required. Rollout to deploy the OS policy assignment.
	//  A rollout is triggered in the following situations:
	//  1) OSPolicyAssignment is created.
	//  2) OSPolicyAssignment is updated and the update contains changes to one of
	//  the following fields:
	//     - instance_filter
	//     - os_policies
	//  3) OSPolicyAssignment is deleted.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.rollout
	Rollout *OSPolicyAssignment_Rollout `json:"rollout,omitempty"`

	// The etag for this OS policy assignment.
	//  If this is provided on update, it must match the server's etag.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter
type OSPolicyAssignment_InstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is
	//  permitted.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.all
	All *bool `json:"all,omitempty"`

	// List of label sets used for VM inclusion.
	//
	//  If the list has more than one `LabelSet`, the VM is included if any
	//  of the label sets are applicable for the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.inclusion_labels
	InclusionLabels []OSPolicyAssignment_LabelSet `json:"inclusionLabels,omitempty"`

	// List of label sets used for VM exclusion.
	//
	//  If the list has more than one label set, the VM is excluded if any
	//  of the label sets are applicable for the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.exclusion_labels
	ExclusionLabels []OSPolicyAssignment_LabelSet `json:"exclusionLabels,omitempty"`

	// List of inventories to select VMs.
	//
	//  A VM is selected if its inventory data matches at least one of the
	//  following inventories.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.inventories
	Inventories []OSPolicyAssignment_InstanceFilter_Inventory `json:"inventories,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.Inventory
type OSPolicyAssignment_InstanceFilter_Inventory struct {
	// Required. The OS short name
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.Inventory.os_short_name
	OsShortName *string `json:"osShortName,omitempty"`

	// The OS version
	//
	//  Prefix matches are supported if asterisk(*) is provided as the
	//  last character. For example, to match all versions with a major
	//  version of `7`, specify the following value for this field `7.*`
	//
	//  An empty string matches all OS versions.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.InstanceFilter.Inventory.os_version
	OsVersion *string `json:"osVersion,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment.LabelSet
type OSPolicyAssignment_LabelSet struct {
	// Labels are identified by key/value pairs in this map.
	//  A VM should contain all the key/value pairs specified in this
	//  map to be selected.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.LabelSet.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment.Rollout
type OSPolicyAssignment_Rollout struct {
	// Required. The maximum number (or percentage) of VMs per zone to disrupt
	//  at any given moment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.Rollout.disruption_budget
	DisruptionBudget *FixedOrPercent `json:"disruptionBudget,omitempty"`

	// Required. This determines the minimum duration of time to wait after the
	//  configuration changes are applied through the current rollout. A
	//  VM continues to count towards the `disruption_budget` at least
	//  until this duration of time has passed after configuration changes are
	//  applied.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.Rollout.min_wait_duration
	MinWaitDuration *string `json:"minWaitDuration,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment
type OSPolicyAssignmentObservedState struct {
	// Output only. The assignment revision ID
	//  A new revision is committed whenever a rollout is triggered for a OS policy
	//  assignment
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. OS policy assignment rollout state
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.rollout_state
	RolloutState *string `json:"rolloutState,omitempty"`

	// Output only. Indicates that this revision has been successfully rolled out
	//  in this zone and new VMs will be assigned OS policies from this revision.
	//
	//  For a given OS policy assignment, there is only one revision with a value
	//  of `true` for this field.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.baseline
	Baseline *bool `json:"baseline,omitempty"`

	// Output only. Indicates that this revision deletes the OS policy assignment.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Output only. Indicates that reconciliation is in progress for the revision.
	//  This value is `true` when the `rollout_state` is one of:
	//  * IN_PROGRESS
	//  * CANCELLING
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Server generated unique id for the OS policy assignment
	//  resource.
	// +kcc:proto:field=google.cloud.osconfig.v1.OSPolicyAssignment.uid
	Uid *string `json:"uid,omitempty"`
}
