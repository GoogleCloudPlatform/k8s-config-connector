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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FolderGVK = GroupVersion.WithKind("Folder")

// FolderSpec defines the desired state of Folder
// +kcc:spec:proto=google.cloud.resourcemanager.v3.Folder
type FolderSpec struct {
	// The folder's display name. A folder's display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The folder that this resource belongs to. Changing this forces the resource to be migrated to the newly specified folder. Only one of folderRef or organizationRef may be specified.
	// +optional
	FolderRef *refsv1beta1.FolderRef `json:"folderRef,omitempty"`

	// The organization that this resource belongs to. Changing this forces the resource to be migrated to the newly specified organization. Only one of folderRef or organizationRef may be specified.
	// +optional
	OrganizationRef *OrganizationRef `json:"organizationRef,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// FolderStatus defines the config connector machine state of Folder
type FolderStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// The folder id from the name "folders/{folder_id}".
	// +optional
	FolderId *string `json:"folderId,omitempty"`

	// The lifecycle state of the folder such as ACTIVE or DELETE_REQUESTED.
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// The resource name of the Folder. Its format is folders/{folder_id}.
	// +optional
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfolder;gcpfolders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// Folder is the Schema for the Folder API
// +k8s:openapi-gen=true
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FolderSpec   `json:"spec,omitempty"`
	Status FolderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FolderList contains a list of Folder
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Folder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Folder{}, &FolderList{})
}

type OrganizationRef struct {
	/* The 'name' field of an Organization resource. */
	// +optional
	Name string `json:"name,omitempty"`

	/* The 'namespace' field of an Organization resource. */
	// +optional
	Namespace string `json:"namespace,omitempty"`

	/* The 'name' field of an Organization resource, when not managed by Config Connector. */
	// +optional
	External string `json:"external,omitempty"`
}
