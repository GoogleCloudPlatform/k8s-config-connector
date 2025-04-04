// Copyright 2025 Google LLC
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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEBackupRestoreGVK = GroupVersion.WithKind("GKEBackupRestore")

// GKEBackupRestoreSpec defines the desired state of GKEBackupRestore
// +kcc:proto=google.cloud.gkebackup.v1.Restore
type GKEBackupRestoreSpec struct {
	// The GKEBackupRestore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The RestorePlan from which this Restore is created.
	// +required
	RestorePlanRef *RestorePlanRef `json:"restorePlanRef,omitempty"`

	// User specified descriptive string for this Restore.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. A reference to the
	//  [Backup][google.cloud.gkebackup.v1.Backup] used as the source from which
	//  this Restore will restore. Note that this Backup must be a sub-resource of
	//  the RestorePlan's
	//  [backup_plan][google.cloud.gkebackup.v1.RestorePlan.backup_plan]. Format:
	//  `projects/*/locations/*/backupPlans/*/backups/*`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.backup
	Backup *string `json:"backup,omitempty"`

	// A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. Filters resources for `Restore`. If not specified, the
	//  scope of the restore will remain the same as defined in the `RestorePlan`.
	//  If this is specified, and no resources are matched by the
	//  `inclusion_filters` or everyting is excluded by the `exclusion_filters`,
	//  nothing will be restored. This filter can only be specified if the value of
	//  [namespaced_resource_restore_mode][google.cloud.gkebackup.v1.RestoreConfig.namespaced_resource_restore_mode]
	//  is set to `MERGE_SKIP_ON_CONFLICT`, `MERGE_REPLACE_VOLUME_ON_CONFLICT` or
	//  `MERGE_REPLACE_ON_CONFLICT`.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.filter
	Filter *Restore_Filter `json:"filter,omitempty"`

	// Optional. Immutable. Overrides the volume data restore policies selected in
	//  the Restore Config for override-scoped resources.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.volume_data_restore_policy_overrides
	VolumeDataRestorePolicyOverrides []VolumeDataRestorePolicyOverride `json:"volumeDataRestorePolicyOverrides,omitempty"`
}

// GKEBackupRestoreStatus defines the config connector machine state of GKEBackupRestore
type GKEBackupRestoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupRestore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupRestoreObservedState `json:"observedState,omitempty"`
}

// GKEBackupRestoreObservedState is the state of the GKEBackupRestore resource as most recently observed in GCP.
// +kcc:proto=google.cloud.gkebackup.v1.Restore
type GKEBackupRestoreObservedState struct {
	// Output only. The full name of the Restore resource.
	//  Format: `projects/*/locations/*/restorePlans/*/restores/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this Restore resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this Restore resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The target cluster into which this Restore will restore data.
	//  Valid formats:
	//
	//    - `projects/*/locations/*/clusters/*`
	//    - `projects/*/zones/*/clusters/*`
	//
	//  Inherited from parent RestorePlan's
	//  [cluster][google.cloud.gkebackup.v1.RestorePlan.cluster] value.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Output only. Configuration of the Restore.  Inherited from parent
	//  RestorePlan's
	//  [restore_config][google.cloud.gkebackup.v1.RestorePlan.restore_config].
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.restore_config
	RestoreConfig *RestoreConfig `json:"restoreConfig,omitempty"`

	// Output only. The current state of the Restore.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why the Restore is in its
	//  current state.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. Timestamp of when the restore operation completed.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. Number of resources restored during the restore execution.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.resources_restored_count
	ResourcesRestoredCount *int32 `json:"resourcesRestoredCount,omitempty"`

	// Output only. Number of resources excluded during the restore execution.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.resources_excluded_count
	ResourcesExcludedCount *int32 `json:"resourcesExcludedCount,omitempty"`

	// Output only. Number of resources that failed to be restored during the
	//  restore execution.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.resources_failed_count
	ResourcesFailedCount *int32 `json:"resourcesFailedCount,omitempty"`

	// Output only. Number of volumes restored during the restore execution.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.volumes_restored_count
	VolumesRestoredCount *int32 `json:"volumesRestoredCount,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a restore from overwriting each other.
	//  It is strongly suggested that systems make use of the `etag` in the
	//  read-modify-write cycle to perform restore updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetRestore`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateRestore` or `DeleteRestore` to ensure that their change will be
	//  applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.Restore.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackuprestore;gcpgkebackuprestores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupRestore is the Schema for the GKEBackupRestore API
// +k8s:openapi-gen=true
type GKEBackupRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupRestoreSpec   `json:"spec,omitempty"`
	Status GKEBackupRestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupRestoreList contains a list of GKEBackupRestore
type GKEBackupRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupRestore{}, &GKEBackupRestoreList{})
}
