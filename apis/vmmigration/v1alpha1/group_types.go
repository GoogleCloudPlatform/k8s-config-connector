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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMMigrationGroupGVK = GroupVersion.WithKind("VMMigrationGroup")

// VMMigrationGroupSpec defines the desired state of VMMigrationGroup
// +kcc:spec:proto=google.cloud.vmmigration.v1.Group
type VMMigrationGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VMMigrationGroup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// User-provided description of the group.
	// +optional
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.description
	Description *string `json:"description,omitempty"`

	// Display name is a user defined name for this group which can be updated.
	// +optional
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// VMMigrationGroupStatus defines the config connector machine state of VMMigrationGroup
type VMMigrationGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMMigrationGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMMigrationGroupObservedState `json:"observedState,omitempty"`
}

// VMMigrationGroupObservedState is the state of the VMMigrationGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmmigration.v1.Group
type VMMigrationGroupObservedState struct {
	// Output only. The Group name.
	// +optional
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time timestamp.
	// +optional
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time timestamp.
	// +optional
	// +kcc:proto:field=google.cloud.vmmigration.v1.Group.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmmigrationgroup;gcpvmmigrationgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMMigrationGroup is the Schema for the VMMigrationGroup API
// +k8s:openapi-gen=true
type VMMigrationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMMigrationGroupSpec   `json:"spec,omitempty"`
	Status VMMigrationGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMMigrationGroupList contains a list of VMMigrationGroup
type VMMigrationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMMigrationGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMMigrationGroup{}, &VMMigrationGroupList{})
}
