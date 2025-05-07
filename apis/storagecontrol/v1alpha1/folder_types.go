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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageFolderGVK = GroupVersion.WithKind("StorageFolder")

// StorageFolderSpec defines the desired state of StorageFolder
// +kcc:proto=google.storage.control.v2.Folder
type StorageFolderSpec struct {
	*StorageFolderParent `json:",inline"`

	// The StorageFolder name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type StorageFolderParent struct {
	// Required. The host project of the application.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Required. The storage bucket where the folder will be created in.
	StorageBucketRef *v1beta1.StorageBucketRef `json:"storagebucketRef,omitempty"`
}

// StorageFolderStatus defines the config connector machine state of StorageFolder
type StorageFolderStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageFolder resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageFolderObservedState `json:"observedState,omitempty"`
}

// StorageFolderObservedState is the state of the StorageFolder resource as most recently observed in GCP.
// +kcc:proto=google.storage.control.v2.Folder
type StorageFolderObservedState struct {
	// Output only. The version of the metadata for this folder. Used for
	//  preconditions and for detecting changes in metadata.
	// +kcc:proto:field=google.storage.control.v2.Folder.metageneration
	Metageneration *int64 `json:"metageneration,omitempty"`

	// Output only. The creation time of the folder.
	// +kcc:proto:field=google.storage.control.v2.Folder.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The modification time of the folder.
	// +kcc:proto:field=google.storage.control.v2.Folder.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only present if the folder is part of an ongoing RenameFolder
	//  operation. Contains information which can be used to query the operation
	//  status. The presence of this field also indicates all write operations are
	//  blocked for this folder, including folder, managed folder, and object
	//  operations.
	// +kcc:proto:field=google.storage.control.v2.Folder.pending_rename_info
	PendingRenameInfo *PendingRenameInfoObservedState `json:"pendingRenameInfo,omitempty"`
}

// +kcc:proto=google.storage.control.v2.PendingRenameInfo
type PendingRenameInfo struct {
}

// +kcc:proto=google.storage.control.v2.PendingRenameInfo
type PendingRenameInfoObservedState struct {
	// Output only. The name of the rename operation.
	// +kcc:proto:field=google.storage.control.v2.PendingRenameInfo.operation
	Operation *string `json:"operation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpstoragefolder;gcpstoragefolders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageFolder is the Schema for the StorageFolder API
// +k8s:openapi-gen=true
type StorageFolder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageFolderSpec   `json:"spec,omitempty"`
	Status StorageFolderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageFolderList contains a list of StorageFolder
type StorageFolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageFolder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageFolder{}, &StorageFolderList{})
}
