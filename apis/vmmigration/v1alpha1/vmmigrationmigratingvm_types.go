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

// API sources for VMMigrationMigratingVM, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmmigration/v1/vmmigration.proto
// +kcc:source:service-docs=https://cloud.google.com/migrate/virtual-machines
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/migrate/virtual-machines/docs/reference/rest/v1/projects.locations.sources.migratingVms

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMMigrationMigratingVMGVK = GroupVersion.WithKind("VMMigrationMigratingVM")

// VMMigrationMigratingVMSpec defines the desired state of VMMigrationMigratingVM
// +kcc:spec:proto=google.cloud.vmmigration.v1.MigratingVm
// +kcc:required-from-proto
type VMMigrationMigratingVMSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VMMigrationMigratingVM name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.compute_engine_target_defaults
	ComputeEngineTargetDefaults *ComputeEngineTargetDefaults `json:"computeEngineTargetDefaults,omitempty"`

	// The unique ID of the VM in the source.
	//  The VM's name in vSphere can be changed, so this is not the VM's name but
	//  rather its moRef id. This id is of the form vm-<num>.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.source_vm_id
	SourceVMID *string `json:"sourceVMID,omitempty"`

	// The display name attached to the MigratingVm by the user.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description attached to the migrating VM by the user.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.description
	Description *string `json:"description,omitempty"`

	// The replication schedule policy.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.policy
	Policy *SchedulePolicy `json:"policy,omitempty"`

	// The labels of the migrating VM.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// VMMigrationMigratingVMStatus defines the config connector machine state of VMMigrationMigratingVM
type VMMigrationMigratingVMStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMMigrationMigratingVM resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMMigrationMigratingVMObservedState `json:"observedState,omitempty"`
}

// VMMigrationMigratingVMObservedState is the state of the VMMigrationMigratingVM resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmmigration.v1.MigratingVm
type VMMigrationMigratingVMObservedState struct {
	// Details of the target VM in Compute Engine.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.compute_engine_target_defaults
	ComputeEngineTargetDefaults *ComputeEngineTargetDefaultsObservedState `json:"computeEngineTargetDefaults,omitempty"`

	// Output only. Details of the VM from an AWS source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.aws_source_vm_details
	AwsSourceVMDetails *AwsSourceVMDetails `json:"awsSourceVMDetails,omitempty"`

	// Output only. The time the migrating VM was created (this refers to this
	//  resource and not to the time it was installed in the source).
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the migrating VM resource was updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The most updated snapshot created time in the source that
	//  finished replication.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.last_sync
	LastSync *ReplicationSync `json:"lastSync,omitempty"`

	// Output only. State of the MigratingVm.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.state
	State *string `json:"state,omitempty"`

	// Output only. The last time the migrating VM state was updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. The percentage progress of the current running replication
	//  cycle.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.current_sync_info
	CurrentSyncInfo *ReplicationCycle `json:"currentSyncInfo,omitempty"`

	// Output only. The group this migrating vm is included in, if any. The group
	//  is represented by the full path of the appropriate
	//  [Group][google.cloud.vmmigration.v1.Group] resource.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.group
	Group *string `json:"group,omitempty"`

	// Output only. The recent [clone jobs][google.cloud.vmmigration.v1.CloneJob]
	//  performed on the migrating VM. This field holds the vm's last completed
	//  clone job and the vm's running clone job, if one exists.
	//  Note: To have this field populated you need to explicitly request it via
	//  the "view" parameter of the Get/List request.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.recent_clone_jobs
	RecentCloneJobs []CloneJobObservedState `json:"recentCloneJobs,omitempty"`

	// Output only. Provides details on the state of the Migrating VM in case of
	//  an error in replication.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. The recent cutover jobs performed on the migrating VM.
	//  This field holds the vm's last completed cutover job and the vm's
	//  running cutover job, if one exists.
	//  Note: To have this field populated you need to explicitly request it via
	//  the "view" parameter of the Get/List request.
	// +kcc:proto:field=google.cloud.vmmigration.v1.MigratingVm.recent_cutover_jobs
	RecentCutoverJobs []CutoverJobObservedState `json:"recentCutoverJobs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmmigrationmigratingvm;gcpvmmigrationmigratingvms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMMigrationMigratingVM is the Schema for the VMMigrationMigratingVM API
// +k8s:openapi-gen=true
type VMMigrationMigratingVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMMigrationMigratingVMSpec   `json:"spec,omitempty"`
	Status VMMigrationMigratingVMStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMMigrationMigratingVMList contains a list of VMMigrationMigratingVM
type VMMigrationMigratingVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMMigrationMigratingVM `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMMigrationMigratingVM{}, &VMMigrationMigratingVMList{})
}
