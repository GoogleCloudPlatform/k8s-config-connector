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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	container "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEBackupBackupPlanGVK = GroupVersion.WithKind("GKEBackupBackupPlan")

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// GKEBackupBackupPlanSpec defines the desired state of GKEBackupBackupPlan
// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan
type GKEBackupBackupPlanSpec struct {
	// The GKEBackupBackupPlan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User specified descriptive string for this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The source cluster from which Backups will be created
	//  via this BackupPlan. Valid formats:
	//
	//  - `projects/*/locations/*/clusters/*`
	//  - `projects/*/zones/*/clusters/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.cluster
	// +required
	ClusterRef *container.ContainerClusterRef `json:"clusterRef,omitempty"`

	// Optional. RetentionPolicy governs lifecycle of Backups created under this
	//  plan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.retention_policy
	RetentionPolicy *BackupPlan_RetentionPolicy `json:"retentionPolicy,omitempty"`

	// Optional. A set of custom labels supplied by user.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Defines a schedule for automatic Backup creation via this
	//  BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_schedule
	BackupSchedule *BackupPlan_Schedule `json:"backupSchedule,omitempty"`

	// Optional. This flag indicates whether this BackupPlan has been deactivated.
	//  Setting this field to True locks the BackupPlan such that no further
	//  updates will be allowed (except deletes), including the deactivated field
	//  itself. It also prevents any new Backups from being created via this
	//  BackupPlan (including scheduled Backups).
	//
	//  Default: False
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.deactivated
	Deactivated *bool `json:"deactivated,omitempty"`

	// Optional. Defines the configuration of Backups created via this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_config
	BackupConfig *BackupPlan_BackupConfig `json:"backupConfig,omitempty"`
}

// GKEBackupBackupPlanStatus defines the config connector machine state of GKEBackupBackupPlan
type GKEBackupBackupPlanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEBackupBackupPlan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEBackupBackupPlanObservedState `json:"observedState,omitempty"`
}

// GKEBackupBackupPlanObservedState is the state of the GKEBackupBackupPlan resource as most recently observed in GCP.
// +kcc:proto=google.cloud.gkebackup.v1.BackupPlan
type GKEBackupBackupPlanObservedState struct {
	// Output only. The full name of the BackupPlan resource.
	//  Format: `projects/*/locations/*/backupPlans/*`
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Server generated global unique identifier of
	//  [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The timestamp when this BackupPlan resource was created.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when this BackupPlan resource was last
	//  updated.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Defines a schedule for automatic Backup creation via this
	//  BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.backup_schedule
	BackupSchedule *BackupPlan_ScheduleObservedState `json:"backupSchedule,omitempty"`

	// Output only. `etag` is used for optimistic concurrency control as a way to
	//  help prevent simultaneous updates of a backup plan from overwriting each
	//  other. It is strongly suggested that systems make use of the 'etag' in the
	//  read-modify-write cycle to perform BackupPlan updates in order to avoid
	//  race conditions: An `etag` is returned in the response to `GetBackupPlan`,
	//  and systems are expected to put that etag in the request to
	//  `UpdateBackupPlan` or `DeleteBackupPlan` to ensure that their change
	//  will be applied to the same version of the resource.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The number of Kubernetes Pods backed up in the
	//  last successful Backup created via this BackupPlan.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.protected_pod_count
	ProtectedPodCount *int32 `json:"protectedPodCount,omitempty"`

	// Output only. State of the BackupPlan. This State field reflects the
	//  various stages a BackupPlan can be in
	//  during the Create operation. It will be set to "DEACTIVATED"
	//  if the BackupPlan is deactivated on an Update
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.state
	State *string `json:"state,omitempty"`

	// Output only. Human-readable description of why BackupPlan is in the current
	//  `state`
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.state_reason
	StateReason *string `json:"stateReason,omitempty"`

	// Output only. A number that represents the current risk level of this
	//  BackupPlan from RPO perspective with 1 being no risk and 5 being highest
	//  risk.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.rpo_risk_level
	RpoRiskLevel *int32 `json:"rpoRiskLevel,omitempty"`

	// Output only. Human-readable description of why the BackupPlan is in the
	//  current rpo_risk_level and action items if any.
	// +kcc:proto:field=google.cloud.gkebackup.v1.BackupPlan.rpo_risk_reason
	RpoRiskReason *string `json:"rpoRiskReason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkebackupbackupplan;gcpgkebackupbackupplans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEBackupBackupPlan is the Schema for the GKEBackupBackupPlan API
// +k8s:openapi-gen=true
type GKEBackupBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEBackupBackupPlanSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupPlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEBackupBackupPlanList contains a list of GKEBackupBackupPlan
type GKEBackupBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEBackupBackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEBackupBackupPlan{}, &GKEBackupBackupPlanList{})
}
