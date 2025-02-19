// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SecureSourceManagerRepositoryGVK = GroupVersion.WithKind("SecureSourceManagerRepository")

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SecureSourceManagerRepositorySpec defines the desired state of SecureSourceManagerRepository
// +kcc:proto=google.cloud.securesourcemanager.v1.Repository
type SecureSourceManagerRepositorySpec struct {
	/* Immutable. The Project that this resource belongs to. */
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the instance. */
	// +required
	Location string `json:"location"`

	// The SecureSourceManagerRepository name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The name of the instance in which the repository is hosted, formatted as
	// `projects/{project_number}/locations/{location_id}/instances/{instance_id}`
	// +required
	InstanceRef *SecureSourceManagerInstanceRef `json:"instanceRef,omitempty"`

	// Input only. Initial configurations for the repository.
	InitialConfig *Repository_InitialConfig `json:"initialConfig,omitempty"`

	// Optional. Description of the repository, which cannot exceed 500 characters.
	// Temporarily omitted for now.
	// Description *string `json:"description,omitempty"`
}

// SecureSourceManagerRepositoryStatus defines the config connector machine state of SecureSourceManagerRepository
type SecureSourceManagerRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecureSourceManagerRepository resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecureSourceManagerRepositoryObservedState `json:"observedState,omitempty"`
}

// SecureSourceManagerRepositoryObservedState is the state of the SecureSourceManagerRepository resource as most recently observed in GCP.
type SecureSourceManagerRepositoryObservedState struct {
	// // Output only. Create timestamp.
	// CreateTime *string `json:"createTime,omitempty"`

	// // Output only. Update timestamp.
	// UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Unique identifier of the repository.
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. URIs for the repository.
	URIs *Repository_URIs `json:"uris,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerrepository;gcpsecuresourcemanagerrepositories
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecureSourceManagerRepository is the Schema for the SecureSourceManagerRepository API
// +k8s:openapi-gen=true
type SecureSourceManagerRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SecureSourceManagerRepositorySpec   `json:"spec,omitempty"`
	Status SecureSourceManagerRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecureSourceManagerRepositoryList contains a list of SecureSourceManagerRepository
type SecureSourceManagerRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerRepository{}, &SecureSourceManagerRepositoryList{})
}
