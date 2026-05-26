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

var DataformFolderGVK = GroupVersion.WithKind("DataformFolder")

// DataformFolderSpec defines the desired state of DataformFolder
// +kcc:spec:proto=google.cloud.dataform.v1beta1.Folder
type DataformFolderSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DataformFolder name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The Folder's user-friendly name.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// Optional. The containing Folder resource name. This should take
	//  the format: projects/{project}/locations/{location}/folders/{folder},
	//  projects/{project}/locations/{location}/teamFolders/{teamFolder}, or just
	//  projects/{project}/locations/{location} if this is a root Folder. This
	//  field can only be updated through MoveFolder.
	ContainingFolder *string `json:"containingFolder,omitempty"`
}

// DataformFolderStatus defines the config connector machine state of DataformFolder
type DataformFolderStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataformFolder resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataformFolderObservedState `json:"observedState,omitempty"`
}

// DataformFolderObservedState is the state of the DataformFolder resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataform.v1beta1.Folder
type DataformFolderObservedState struct {
	// Output only. The resource name of the TeamFolder that this Folder is
	//  associated with. This should take the format:
	//  projects/{project}/locations/{location}/teamFolders/{teamFolder}. If this
	//  is not set, the Folder is not associated with a TeamFolder and is a
	//  UserFolder.
	TeamFolderName *string `json:"teamFolderName,omitempty"`

	// Output only. The timestamp of when the Folder was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the Folder was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. All the metadata information that is used internally to serve
	//  the resource. For example: timestamps, flags, status fields, etc. The
	//  format of this field is a JSON string.
	InternalMetadata *string `json:"internalMetadata,omitempty"`

	// Output only. The IAM principal identifier of the creator of the Folder.
	CreatorIAMPrincipal *string `json:"creatorIAMPrincipal,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformfolder;gcpdataformfolders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformFolder is the Schema for the DataformFolder API
// +k8s:openapi-gen=true
type DataformFolder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformFolderSpec   `json:"spec,omitempty"`
	Status DataformFolderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataformFolderList contains a list of DataformFolder
type DataformFolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformFolder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformFolder{}, &DataformFolderList{})
}
