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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OSConfigGuestPolicyGVK = GroupVersion.WithKind("OSConfigGuestPolicy")

// OSConfigGuestPolicySpec defines the desired state of OSConfigGuestPolicy
// +kcc:spec:proto=google.cloud.osconfig.v1beta.GuestPolicy
type OSConfigGuestPolicySpec struct {
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Description of the guest policy. Length of the description is limited
	// to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GuestPolicy.description
	Description *string `json:"description,omitempty"`

	// Required. Specifies the VM instances that are assigned to this policy. This allows
	// you to target sets or groups of VM instances by different parameters such
	// as labels, names, OS, or zones.
	//
	// If left empty, all VM instances underneath this policy are targeted.
	//
	// At the same level in the resource hierarchy (that is within a project), the
	// service prevents the creation of multiple policies that conflict with
	// each other. For more information, see how the service [handles assignment
	// conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GuestPolicy.assignment
	Assignment *Assignment `json:"assignment,omitempty"`

	// The software packages to be managed by this policy.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GuestPolicy.packages
	Packages []Package `json:"packages,omitempty"`

	// A list of package repositories to configure on the VM instance. This is
	// done before any other configs are applied so they can use these repos.
	// Package repositories are only configured if the corresponding package
	// manager(s) are available.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GuestPolicy.package_repositories
	PackageRepositories []PackageRepository `json:"packageRepositories,omitempty"`

	// A list of Recipes to install on the VM instance.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GuestPolicy.recipes
	Recipes []SoftwareRecipe `json:"recipes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.AptRepository
type AptRepository struct {
	// Type of archive files in this repository. The default behavior is DEB.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptRepository.archive_type
	ArchiveType *string `json:"archiveType,omitempty"`

	// Required. URI for this repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptRepository.uri
	// +required
	URI *string `json:"uri,omitempty"`

	// Required. Distribution of this repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptRepository.distribution
	// +required
	Distribution *string `json:"distribution,omitempty"`

	// Required. List of components for this repository. Must contain at least one item.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptRepository.components
	Components []string `json:"components,omitempty"`

	// URI of the key file for this repository. The agent maintains
	// a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg` containing
	// all the keys in any applied guest policy.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptRepository.gpg_key
	GpgKey *string `json:"gpgKey,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.Assignment
type Assignment struct {
	// Targets instances matching at least one of these label sets. This allows
	// an assignment to target disparate groups, for example "env=prod or
	// env=staging".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.group_labels
	GroupLabels []Assignment_GroupLabel `json:"groupLabels,omitempty"`

	// Targets instances in any of these zones. Leave empty to target instances
	// in any zone.
	//
	// Zonal targeting is uncommon and is supported to facilitate the management
	// of changes by zone.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.zones
	Zones []string `json:"zones,omitempty"`

	// Targets any of the instances specified. Instances are specified by their
	// URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`.
	//
	// Instance targeting is uncommon and is supported to facilitate the
	// management of changes by the instance or to target specific VM instances
	// for development and testing.
	//
	// Only supported for project-level policies and must reference instances
	// within this project.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.instances
	Instances []computev1beta1.InstanceRef `json:"instances,omitempty"`

	// Targets VM instances whose name starts with one of these prefixes.
	//
	// Like labels, this is another way to group VM instances when targeting
	// configs, for example prefix="prod-".
	//
	// Only supported for project-level policies.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.instance_name_prefixes
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty"`

	// Targets VM instances matching at least one of the following OS types.
	//
	// VM instances must match all supplied criteria for a given OsType to be
	// included.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.os_types
	OSTypes []Assignment_OSType `json:"osTypes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.Assignment.GroupLabel
type Assignment_GroupLabel struct {
	// Google Compute Engine instance labels that must be present for an
	// instance to be included in this assignment group.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.GroupLabel.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.Assignment.OsType
type Assignment_OSType struct {
	// Targets VM instances with OS Inventory enabled and having the following
	// OS short name, for example "debian" or "windows".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.OsType.os_short_name
	OSShortName *string `json:"osShortName,omitempty"`

	// Targets VM instances with OS Inventory enabled and having the following
	// following OS version.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.OsType.os_version
	OSVersion *string `json:"osVersion,omitempty"`

	// Targets VM instances with OS Inventory enabled and having the following
	// OS architecture.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Assignment.OsType.os_architecture
	OSArchitecture *string `json:"osArchitecture,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.GooRepository
type GooRepository struct {
	// Required. The name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GooRepository.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. The url of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GooRepository.url
	// +required
	URL *string `json:"url,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.Package
type Package struct {
	// Required. The name of the package. A package is uniquely identified for conflict
	// validation by checking the package name and the manager(s) that the
	// package targets.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Package.name
	Name *string `json:"name,omitempty"`

	// The desired_state the agent should maintain for this package. The
	// default is to ensure the package is installed.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Package.desired_state
	DesiredState *string `json:"desiredState,omitempty"`

	// Type of package manager that can be used to install this package.
	// If a system does not have the package manager, the package is not
	// installed or removed no error message is returned. By default,
	// or if you specify `ANY`,
	// the agent attempts to install and remove this package using the default
	// package manager. This is useful when creating a policy that applies to
	// different types of systems.
	//
	// The default behavior is ANY.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.Package.manager
	Manager *string `json:"manager,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PackageRepository
type PackageRepository struct {
	// An Apt Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PackageRepository.apt
	Apt *AptRepository `json:"apt,omitempty"`

	// A Yum Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PackageRepository.yum
	Yum *YumRepository `json:"yum,omitempty"`

	// A Zypper Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PackageRepository.zypper
	Zypper *ZypperRepository `json:"zypper,omitempty"`

	// A Goo Repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PackageRepository.goo
	Goo *GooRepository `json:"goo,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe
type SoftwareRecipe struct {
	// Required. Unique identifier for the recipe. Only one recipe with a given name is
	// installed on an instance.
	//
	// Names are also used to identify resources which helps to determine whether
	// guest policies have conflicts. This means that requests to create multiple
	// recipes with the same name and version are rejected since they
	// could potentially have conflicting assignments.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.name
	Name *string `json:"name,omitempty"`

	// The version of this software recipe. Version can be up to 4 period
	// separated numbers (e.g. 12.34.56.78).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.version
	Version *string `json:"version,omitempty"`

	// Resources available to be used in the steps in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.artifacts
	Artifacts []SoftwareRecipe_Artifact `json:"artifacts,omitempty"`

	// Actions to be taken for installing this recipe. On failure it stops
	// executing steps and does not attempt another installation. Any steps taken
	// (including partially completed steps) are not rolled back.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.install_steps
	InstallSteps []SoftwareRecipe_Step `json:"installSteps,omitempty"`

	// Actions to be taken for updating this recipe. On failure it stops
	// executing steps and does not attempt another update for this recipe. Any
	// steps taken (including partially completed steps) are not rolled back.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.update_steps
	UpdateSteps []SoftwareRecipe_Step `json:"updateSteps,omitempty"`

	// Default is INSTALLED. The desired state the agent should maintain for this
	// recipe.
	//
	// INSTALLED: The software recipe is installed on the instance but
	// won't be updated to new versions.
	// UPDATED: The software recipe is installed on the instance. The recipe is
	// updated to a higher version, if a higher version of the recipe is
	// assigned to this instance.
	// REMOVE: Remove is unsupported for software recipes and attempts to
	// create or update a recipe to the REMOVE state is rejected.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.desired_state
	DesiredState *string `json:"desiredState,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact
type SoftwareRecipe_Artifact struct {
	// Required. Id of the artifact, which the installation and update steps of this
	// recipe can reference. Artifacts in a recipe cannot have the same id.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.id
	ID *string `json:"id,omitempty"`

	// A generic remote artifact.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.remote
	Remote *SoftwareRecipe_Artifact_Remote `json:"remote,omitempty"`

	// A Google Cloud Storage artifact.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.gcs
	GCS *SoftwareRecipe_Artifact_GCS `json:"gcs,omitempty"`

	// Defaults to false. When false, recipes are subject to validations
	// based on the artifact type:
	//
	// Remote: A checksum must be specified, and only protocols with
	// transport-layer security are permitted.
	// GCS: An object generation number must be specified.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.allow_insecure
	AllowInsecure *bool `json:"allowInsecure,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Gcs
type SoftwareRecipe_Artifact_GCS struct {
	// Bucket of the Google Cloud Storage object.
	// Given an example URL:
	// `https://storage.googleapis.com/my-bucket/foo/bar#1234567`
	// this value would be `my-bucket`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Gcs.bucket
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Name of the Google Cloud Storage object.
	// As specified [here]
	// (https://cloud.google.com/storage/docs/naming#objectnames)
	// Given an example URL:
	// `https://storage.googleapis.com/my-bucket/foo/bar#1234567`
	// this value would be `foo/bar`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Gcs.object
	Object *string `json:"object,omitempty"`

	// Must be provided if allow_insecure is false.
	// Generation number of the Google Cloud Storage object.
	// `https://storage.googleapis.com/my-bucket/foo/bar#1234567`
	// this value would be `1234567`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Gcs.generation
	Generation *int64 `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Remote
type SoftwareRecipe_Artifact_Remote struct {
	// URI from which to fetch the object. It should contain both the protocol
	// and path following the format {protocol}://{location}.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Remote.uri
	URI *string `json:"uri,omitempty"`

	// Must be provided if `allow_insecure` is `false`.
	// SHA256 checksum in hex format, to compare to the checksum of the
	// artifact. If the checksum is not empty and it doesn't match the
	// artifact then the recipe installation fails before running any of the
	// steps.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Artifact.Remote.checksum
	Checksum *string `json:"checksum,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step
type SoftwareRecipe_Step struct {
	// Copies a file onto the instance.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.file_copy
	FileCopy *SoftwareRecipe_Step_CopyFile `json:"fileCopy,omitempty"`

	// Extracts an archive into the specified directory.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.archive_extraction
	ArchiveExtraction *SoftwareRecipe_Step_ExtractArchive `json:"archiveExtraction,omitempty"`

	// Installs an MSI file.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.msi_installation
	MsiInstallation *SoftwareRecipe_Step_InstallMsi `json:"msiInstallation,omitempty"`

	// Installs a deb file via dpkg.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.dpkg_installation
	DpkgInstallation *SoftwareRecipe_Step_InstallDpkg `json:"dpkgInstallation,omitempty"`

	// Installs an rpm file via the rpm utility.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.rpm_installation
	RpmInstallation *SoftwareRecipe_Step_InstallRpm `json:"rpmInstallation,omitempty"`

	// Executes an artifact or local file.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.file_exec
	FileExec *SoftwareRecipe_Step_ExecFile `json:"fileExec,omitempty"`

	// Runs commands in a shell.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.script_run
	ScriptRun *SoftwareRecipe_Step_RunScript `json:"scriptRun,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.CopyFile
type SoftwareRecipe_Step_CopyFile struct {
	// Required. The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.CopyFile.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`

	// Required. The absolute path on the instance to put the file.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.CopyFile.destination
	Destination *string `json:"destination,omitempty"`

	// Whether to allow this step to overwrite existing files. If this is
	// false and the file already exists the file is not overwritten
	// and the step is considered a success. Defaults to false.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.CopyFile.overwrite
	Overwrite *bool `json:"overwrite,omitempty"`

	// Consists of three octal digits which represent, in
	// order, the permissions of the owner, group, and other users for the
	// file (similarly to the numeric mode used in the linux chmod utility).
	// Each digit represents a three bit number with the 4 bit
	// corresponding to the read permissions, the 2 bit corresponds to the
	// write bit, and the one bit corresponds to the execute permission.
	// Default behavior is 755.
	//
	// Below are some examples of permissions and their associated values:
	// read, write, and execute: 7
	// read and execute: 5
	// read and write: 6
	// read only: 4
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.CopyFile.permissions
	Permissions *string `json:"permissions,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExecFile
type SoftwareRecipe_Step_ExecFile struct {
	// The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExecFile.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`

	// The absolute path of the file on the local filesystem.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExecFile.local_path
	LocalPath *string `json:"localPath,omitempty"`

	// Arguments to be passed to the provided executable.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExecFile.args
	Args []string `json:"args,omitempty"`

	// Defaults to [0]. A list of possible return values that the program
	// can return to indicate a success.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExecFile.allowed_exit_codes
	AllowedExitCodes []int64 `json:"allowedExitCodes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExtractArchive
type SoftwareRecipe_Step_ExtractArchive struct {
	// Required. The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExtractArchive.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`

	// Directory to extract archive to.
	// Defaults to `/` on Linux or `C:\` on Windows.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExtractArchive.destination
	Destination *string `json:"destination,omitempty"`

	// Required. The type of the archive to extract.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.ExtractArchive.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallDpkg
type SoftwareRecipe_Step_InstallDpkg struct {
	// Required. The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallDpkg.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallMsi
type SoftwareRecipe_Step_InstallMsi struct {
	// Required. The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallMsi.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`

	// The flags to use when installing the MSI
	// defaults to ["/i"] (i.e. the install flag).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallMsi.flags
	Flags []string `json:"flags,omitempty"`

	// Return codes that indicate that the software installed or updated
	// successfully. Behaviour defaults to [0]
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallMsi.allowed_exit_codes
	AllowedExitCodes []int64 `json:"allowedExitCodes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallRpm
type SoftwareRecipe_Step_InstallRpm struct {
	// Required. The id of the relevant artifact in the recipe.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.InstallRpm.artifact_id
	ArtifactID *string `json:"artifactId,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.RunScript
type SoftwareRecipe_Step_RunScript struct {
	// Required. The shell script to be executed.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.RunScript.script
	Script *string `json:"script,omitempty"`

	// Return codes that indicate that the software installed or updated
	// successfully. Behaviour defaults to [0]
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.RunScript.allowed_exit_codes
	AllowedExitCodes []int64 `json:"allowedExitCodes,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is
	// specified the script is executed directly, which likely
	// only succeed for scripts with
	// [shebang lines](https://en.wikipedia.org/wiki/Shebang_\(Unix\)).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.SoftwareRecipe.Step.RunScript.interpreter
	Interpreter *string `json:"interpreter,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.YumRepository
type YumRepository struct {
	// Required. A one word, unique name for this repository. This is
	// the `repo id` in the Yum config file and also the `display_name` if
	// `display_name` is omitted. This id is also used as the unique identifier
	// when checking for guest policy conflicts.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumRepository.id
	// +required
	ID *string `json:"id,omitempty"`

	// The display name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumRepository.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The location of the repository directory.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumRepository.base_url
	// +required
	BaseURL *string `json:"baseUrl,omitempty"`

	// URIs of GPG keys.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumRepository.gpg_keys
	GpgKeys []string `json:"gpgKeys,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ZypperRepository
type ZypperRepository struct {
	// Required. A one word, unique name for this repository. This is
	// the `repo id` in the zypper config file and also the `display_name` if
	// `display_name` is omitted. This id is also used as the unique identifier
	// when checking for guest policy conflicts.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperRepository.id
	// +required
	ID *string `json:"id,omitempty"`

	// The display name of the repository.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperRepository.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The location of the repository directory.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperRepository.base_url
	// +required
	BaseURL *string `json:"baseUrl,omitempty"`

	// URIs of GPG keys.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperRepository.gpg_keys
	GpgKeys []string `json:"gpgKeys,omitempty"`
}

// OSConfigGuestPolicyStatus defines the config connector machine state of OSConfigGuestPolicy
type OSConfigGuestPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The etag for this GuestPolicy. If this is provided on update, it must match the server's etag.
	Etag *string `json:"etag,omitempty"`

	// Output only. Time this GuestPolicy was created.
	// +kubebuilder:validation:Format="date-time"
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last time this GuestPolicy was updated.
	// +kubebuilder:validation:Format="date-time"
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcposconfigguestpolicy;gcposconfigguestpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OSConfigGuestPolicy is the Schema for the OSConfigGuestPolicy API
// +k8s:openapi-gen=true
type OSConfigGuestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigGuestPolicySpec   `json:"spec,omitempty"`
	Status OSConfigGuestPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OSConfigGuestPolicyList contains a list of OSConfigGuestPolicy
type OSConfigGuestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSConfigGuestPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSConfigGuestPolicy{}, &OSConfigGuestPolicyList{})
}
