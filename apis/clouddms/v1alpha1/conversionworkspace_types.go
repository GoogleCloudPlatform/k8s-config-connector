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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSConversionWorkspaceGVK = GroupVersion.WithKind("CloudDMSConversionWorkspace")

// CloudDMSConversionWorkspaceSpec defines the desired state of CloudDMSConversionWorkspace
// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspace
type CloudDMSConversionWorkspaceSpec struct {
	// The CloudDMSConversionWorkspace name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The location where the alloydb cluster should reside.
	// +required
	Location string `json:"location,omitempty"`

	// Required. The source engine details.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.source
	Source *DatabaseEngineInfo `json:"source,omitempty"`

	// Required. The destination engine details.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.destination
	Destination *DatabaseEngineInfo `json:"destination,omitempty"`

	// Optional. A generic list of settings for the workspace.
	//  The settings are database pair dependant and can indicate default behavior
	//  for the mapping rules engine or turn on or off specific features.
	//  Such examples can be: convert_foreign_key_to_interleave=true,
	//  skip_triggers=false, ignore_non_table_synonyms=true
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.global_settings
	GlobalSettings map[string]string `json:"globalSettings,omitempty"`

	// Optional. The display name for the workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// CloudDMSConversionWorkspaceStatus defines the config connector machine state of CloudDMSConversionWorkspace
type CloudDMSConversionWorkspaceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSConversionWorkspace resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSConversionWorkspaceObservedState `json:"observedState,omitempty"`
}

// CloudDMSConversionWorkspaceObservedState is the state of the CloudDMSConversionWorkspace resource as most recently observed in GCP.
// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspace
type CloudDMSConversionWorkspaceObservedState struct {
	// Output only. Whether the workspace has uncommitted changes (changes which
	//  were made after the workspace was committed).
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.has_uncommitted_changes
	HasUncommittedChanges *bool `json:"hasUncommittedChanges,omitempty"`

	// Output only. The latest commit ID.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.latest_commit_id
	LatestCommitID *string `json:"latestCommitID,omitempty"`

	// Output only. The timestamp when the workspace was committed.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.latest_commit_time
	LatestCommitTime *string `json:"latestCommitTime,omitempty"`

	// Output only. The timestamp when the workspace resource was created.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the workspace resource was last updated.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsconversionworkspace;gcpclouddmsconversionworkspaces
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSConversionWorkspace is the Schema for the CloudDMSConversionWorkspace API
// +k8s:openapi-gen=true
type CloudDMSConversionWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSConversionWorkspaceSpec   `json:"spec,omitempty"`
	Status CloudDMSConversionWorkspaceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSConversionWorkspaceList contains a list of CloudDMSConversionWorkspace
type CloudDMSConversionWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSConversionWorkspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSConversionWorkspace{}, &CloudDMSConversionWorkspaceList{})
}
