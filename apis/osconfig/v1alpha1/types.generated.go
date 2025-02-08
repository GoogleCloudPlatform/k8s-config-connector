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


// +kcc:proto=google.cloud.osconfig.v1beta.AptSettings
type AptSettings struct {
	// By changing the type to DIST, the patching is performed
	//  using `apt-get dist-upgrade` instead.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.type
	Type *string `json:"type,omitempty"`

	// List of packages to exclude from update. These packages will be excluded
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages
	//  that will be updated. If these packages are not installed, they will be
	//  ignored. This field cannot be specified with any other patch configuration
	//  fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.AptSettings.exclusive_packages
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ExecStep
type ExecStep struct {
	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStep.linux_exec_step_config
	LinuxExecStepConfig *ExecStepConfig `json:"linuxExecStepConfig,omitempty"`

	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStep.windows_exec_step_config
	WindowsExecStepConfig *ExecStepConfig `json:"windowsExecStepConfig,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ExecStepConfig
type ExecStepConfig struct {
	// An absolute path to the executable on the VM.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.local_path
	LocalPath *string `json:"localPath,omitempty"`

	// A Google Cloud Storage object containing the executable.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.gcs_object
	GcsObject *GcsObject `json:"gcsObject,omitempty"`

	// Defaults to [0]. A list of possible return values that the
	//  execution can return to indicate a success.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.allowed_success_codes
	AllowedSuccessCodes []int32 `json:"allowedSuccessCodes,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is
	//  specified the script will be executed directly, which will likely
	//  only succeed for scripts with [shebang lines]
	//  (https://en.wikipedia.org/wiki/Shebang_\(Unix\)).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ExecStepConfig.interpreter
	Interpreter *string `json:"interpreter,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.FixedOrPercent
type FixedOrPercent struct {
	// Specifies a fixed value.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.FixedOrPercent.fixed
	Fixed *int32 `json:"fixed,omitempty"`

	// Specifies the relative value defined as a percentage, which will be
	//  multiplied by a reference value.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.FixedOrPercent.percent
	Percent *int32 `json:"percent,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.GcsObject
type GcsObject struct {
	// Required. Bucket of the Google Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Required. Name of the Google Cloud Storage object.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.object
	Object *string `json:"object,omitempty"`

	// Required. Generation number of the Google Cloud Storage object. This is used to
	//  ensure that the ExecStep specified by this PatchJob does not change.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.GcsObject.generation_number
	GenerationNumber *int64 `json:"generationNumber,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.GooSettings
type GooSettings struct {
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchConfig
type PatchConfig struct {
	// Post-patch reboot settings.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.reboot_config
	RebootConfig *string `json:"rebootConfig,omitempty"`

	// Apt update settings. Use this setting to override the default `apt` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.apt
	Apt *AptSettings `json:"apt,omitempty"`

	// Yum update settings. Use this setting to override the default `yum` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.yum
	Yum *YumSettings `json:"yum,omitempty"`

	// Goo update settings. Use this setting to override the default `goo` patch
	//  rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.goo
	Goo *GooSettings `json:"goo,omitempty"`

	// Zypper update settings. Use this setting to override the default `zypper`
	//  patch rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.zypper
	Zypper *ZypperSettings `json:"zypper,omitempty"`

	// Windows update settings. Use this override the default windows patch rules.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.windows_update
	WindowsUpdate *WindowsUpdateSettings `json:"windowsUpdate,omitempty"`

	// The `ExecStep` to run before the patch update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.pre_step
	PreStep *ExecStep `json:"preStep,omitempty"`

	// The `ExecStep` to run after the patch update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.post_step
	PostStep *ExecStep `json:"postStep,omitempty"`

	// Allows the patch job to run on Managed instance groups (MIGs).
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchConfig.mig_instances_allowed
	MigInstancesAllowed *bool `json:"migInstancesAllowed,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchInstanceFilter
type PatchInstanceFilter struct {
	// Target all VM instances in the project. If true, no other criteria is
	//  permitted.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.all
	All *bool `json:"all,omitempty"`

	// Targets VM instances matching at least one of these label sets. This allows
	//  targeting of disparate groups, for example "env=prod or env=staging".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.group_labels
	GroupLabels []PatchInstanceFilter_GroupLabel `json:"groupLabels,omitempty"`

	// Targets VM instances in ANY of these zones. Leave empty to target VM
	//  instances in any zone.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.zones
	Zones []string `json:"zones,omitempty"`

	// Targets any of the VM instances specified. Instances are specified by their
	//  URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`,
	//  `projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`, or
	//  `https://www.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.instances
	Instances []string `json:"instances,omitempty"`

	// Targets VMs whose name starts with one of these prefixes. Similar to
	//  labels, this is another way to group VMs when targeting configs, for
	//  example prefix="prod-".
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.instance_name_prefixes
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchInstanceFilter.GroupLabel
type PatchInstanceFilter_GroupLabel struct {
	// Compute Engine instance labels that must be present for a VM instance to
	//  be targeted by this filter.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchInstanceFilter.GroupLabel.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchJob
type PatchJob struct {
	// Unique identifier for this patch job in the form
	//  `projects/*/patchJobs/*`
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.name
	Name *string `json:"name,omitempty"`

	// Display name for this patch job. This is not a unique identifier.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description of the patch job. Length of the description is limited
	//  to 1024 characters.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.description
	Description *string `json:"description,omitempty"`

	// Time this patch job was created.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Last time this patch job was updated.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The current state of the PatchJob.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.state
	State *string `json:"state,omitempty"`

	// Instances to patch.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.instance_filter
	InstanceFilter *PatchInstanceFilter `json:"instanceFilter,omitempty"`

	// Patch configuration being applied.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.patch_config
	PatchConfig *PatchConfig `json:"patchConfig,omitempty"`

	// Duration of the patch job. After the duration ends, the
	//  patch job times out.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.duration
	Duration *string `json:"duration,omitempty"`

	// Summary of instance details.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.instance_details_summary
	InstanceDetailsSummary *PatchJob_InstanceDetailsSummary `json:"instanceDetailsSummary,omitempty"`

	// If this patch job is a dry run, the agent reports that it has
	//  finished without running any updates on the VM instance.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.dry_run
	DryRun *bool `json:"dryRun,omitempty"`

	// If this patch job failed, this message provides information about the
	//  failure.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Reflects the overall progress of the patch job in the range of
	//  0.0 being no progress to 100.0 being complete.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.percent_complete
	PercentComplete *float64 `json:"percentComplete,omitempty"`

	// Rollout strategy being applied.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.rollout
	Rollout *PatchRollout `json:"rollout,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary
type PatchJob_InstanceDetailsSummary struct {
	// Number of instances pending patch job.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.pending_instance_count
	PendingInstanceCount *int64 `json:"pendingInstanceCount,omitempty"`

	// Number of instances that are inactive.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.inactive_instance_count
	InactiveInstanceCount *int64 `json:"inactiveInstanceCount,omitempty"`

	// Number of instances notified about patch job.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.notified_instance_count
	NotifiedInstanceCount *int64 `json:"notifiedInstanceCount,omitempty"`

	// Number of instances that have started.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.started_instance_count
	StartedInstanceCount *int64 `json:"startedInstanceCount,omitempty"`

	// Number of instances that are downloading patches.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.downloading_patches_instance_count
	DownloadingPatchesInstanceCount *int64 `json:"downloadingPatchesInstanceCount,omitempty"`

	// Number of instances that are applying patches.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.applying_patches_instance_count
	ApplyingPatchesInstanceCount *int64 `json:"applyingPatchesInstanceCount,omitempty"`

	// Number of instances rebooting.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.rebooting_instance_count
	RebootingInstanceCount *int64 `json:"rebootingInstanceCount,omitempty"`

	// Number of instances that have completed successfully.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.succeeded_instance_count
	SucceededInstanceCount *int64 `json:"succeededInstanceCount,omitempty"`

	// Number of instances that require reboot.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.succeeded_reboot_required_instance_count
	SucceededRebootRequiredInstanceCount *int64 `json:"succeededRebootRequiredInstanceCount,omitempty"`

	// Number of instances that failed.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.failed_instance_count
	FailedInstanceCount *int64 `json:"failedInstanceCount,omitempty"`

	// Number of instances that have acked and will start shortly.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.acked_instance_count
	AckedInstanceCount *int64 `json:"ackedInstanceCount,omitempty"`

	// Number of instances that exceeded the time out while applying the patch.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.timed_out_instance_count
	TimedOutInstanceCount *int64 `json:"timedOutInstanceCount,omitempty"`

	// Number of instances that are running the pre-patch step.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.pre_patch_step_instance_count
	PrePatchStepInstanceCount *int64 `json:"prePatchStepInstanceCount,omitempty"`

	// Number of instances that are running the post-patch step.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.post_patch_step_instance_count
	PostPatchStepInstanceCount *int64 `json:"postPatchStepInstanceCount,omitempty"`

	// Number of instances that do not appear to be running the agent. Check to
	//  ensure that the agent is installed, running, and able to communicate with
	//  the service.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.InstanceDetailsSummary.no_agent_detected_instance_count
	NoAgentDetectedInstanceCount *int64 `json:"noAgentDetectedInstanceCount,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchRollout
type PatchRollout struct {
	// Mode of the patch rollout.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchRollout.mode
	Mode *string `json:"mode,omitempty"`

	// The maximum number (or percentage) of VMs per zone to disrupt at any given
	//  moment. The number of VMs calculated from multiplying the percentage by the
	//  total number of VMs in a zone is rounded up.
	//
	//  During patching, a VM is considered disrupted from the time the agent is
	//  notified to begin until patching has completed. This disruption time
	//  includes the time to complete reboot and any post-patch steps.
	//
	//  A VM contributes to the disruption budget if its patching operation fails
	//  either when applying the patches, running pre or post patch steps, or if it
	//  fails to respond with a success notification before timing out. VMs that
	//  are not running or do not have an active agent do not count toward this
	//  disruption budget.
	//
	//  For zone-by-zone rollouts, if the disruption budget in a zone is exceeded,
	//  the patch job stops, because continuing to the next zone requires
	//  completion of the patch process in the previous zone.
	//
	//  For example, if the disruption budget has a fixed value of `10`, and 8 VMs
	//  fail to patch in the current zone, the patch job continues to patch 2 VMs
	//  at a time until the zone is completed. When that zone is completed
	//  successfully, patching begins with 10 VMs at a time in the next zone. If 10
	//  VMs in the next zone fail to patch, the patch job stops.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchRollout.disruption_budget
	DisruptionBudget *FixedOrPercent `json:"disruptionBudget,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.WindowsUpdateSettings
type WindowsUpdateSettings struct {
	// Only apply updates of these windows update classifications. If empty, all
	//  updates are applied.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.classifications
	Classifications []string `json:"classifications,omitempty"`

	// List of KBs to exclude from update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of kbs to be updated. These are the only patches
	//  that will be updated. This field must not be used with other
	//  patch configurations.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.WindowsUpdateSettings.exclusive_patches
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.YumSettings
type YumSettings struct {
	// Adds the `--security` flag to `yum update`. Not supported on
	//  all platforms.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.security
	Security *bool `json:"security,omitempty"`

	// Will cause patch to run `yum update-minimal` instead.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.minimal
	Minimal *bool `json:"minimal,omitempty"`

	// List of packages to exclude from update. These packages are excluded by
	//  using the yum `--exclude` flag.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages
	//  that will be updated. If these packages are not installed, they will be
	//  ignored. This field must not be specified with any other patch
	//  configuration fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.YumSettings.exclusive_packages
	ExclusivePackages []string `json:"exclusivePackages,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.ZypperSettings
type ZypperSettings struct {
	// Adds the `--with-optional` flag to `zypper patch`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.with_optional
	WithOptional *bool `json:"withOptional,omitempty"`

	// Adds the `--with-update` flag, to `zypper patch`.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.with_update
	WithUpdate *bool `json:"withUpdate,omitempty"`

	// Install only patches with these categories.
	//  Common categories include security, recommended, and feature.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.categories
	Categories []string `json:"categories,omitempty"`

	// Install only patches with these severities.
	//  Common severities include critical, important, moderate, and low.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.severities
	Severities []string `json:"severities,omitempty"`

	// List of patches to exclude from update.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.excludes
	Excludes []string `json:"excludes,omitempty"`

	// An exclusive list of patches to be updated. These are the only patches
	//  that will be installed using 'zypper patch patch:<patch_name>' command.
	//  This field must not be used with any other patch configuration fields.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.ZypperSettings.exclusive_patches
	ExclusivePatches []string `json:"exclusivePatches,omitempty"`
}

// +kcc:proto=google.cloud.osconfig.v1beta.PatchJob
type PatchJobObservedState struct {
	// Output only. Name of the patch deployment that created this patch job.
	// +kcc:proto:field=google.cloud.osconfig.v1beta.PatchJob.patch_deployment
	PatchDeployment *string `json:"patchDeployment,omitempty"`
}
