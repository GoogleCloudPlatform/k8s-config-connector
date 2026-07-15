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

var DataformTeamFolderGVK = GroupVersion.WithKind("DataformTeamFolder")

// DataformTeamFolderSpec defines the desired state of DataformTeamFolder
// +kcc:spec:proto=google.cloud.dataform.v1.TeamFolder
type DataformTeamFolderSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataformTeamFolder name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The TeamFolder's user-friendly name.
	// +required
	DisplayName *string `json:"displayName,omitempty"`
}

// DataformTeamFolderStatus defines the config connector machine state of DataformTeamFolder
type DataformTeamFolderStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataformTeamFolder resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataformTeamFolderObservedState `json:"observedState,omitempty"`
}

// DataformTeamFolderObservedState is the state of the DataformTeamFolder resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataform.v1.TeamFolder
type DataformTeamFolderObservedState struct {
	// Output only. The timestamp of when the TeamFolder was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the TeamFolder was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. All the metadata information that is used internally to serve the resource.
	InternalMetadata *string `json:"internalMetadata,omitempty"`

	// Output only. The IAM principal identifier of the creator of the TeamFolder.
	CreatorIAMPrincipal *string `json:"creatorIAMPrincipal,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformteamfolder;gcpdataformteamfolders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformTeamFolder is the Schema for the DataformTeamFolder API
// +k8s:openapi-gen=true
type DataformTeamFolder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformTeamFolderSpec   `json:"spec,omitempty"`
	Status DataformTeamFolderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataformTeamFolderList contains a list of DataformTeamFolder
type DataformTeamFolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformTeamFolder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformTeamFolder{}, &DataformTeamFolderList{})
}
