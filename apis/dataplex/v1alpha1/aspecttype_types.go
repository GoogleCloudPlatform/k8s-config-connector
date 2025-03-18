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

var DataplexAspectTypeGVK = GroupVersion.WithKind("DataplexAspectType")

type Parent struct {
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	Location string `json:"location"`
}

// DataplexAspectTypeSpec defines the desired state of DataplexAspectType
// +kcc:proto=google.cloud.dataplex.v1.AspectType
type DataplexAspectTypeSpec struct {
	// The DataplexAspectType name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. Description of the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.description
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the AspectType.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The service computes this checksum. The client may send it on update and
	//  delete requests to ensure it has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.etag
	Etag *string `json:"etag,omitempty"`

	// Immutable. Defines the Authorization for this type.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.authorization
	Authorization *AspectType_Authorization `json:"authorization,omitempty"`

	// Required. MetadataTemplate of the aspect.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.metadata_template
	MetadataTemplate *string `json:"metadataTemplate,omitempty"`
}

// DataplexAspectTypeStatus defines the config connector machine state of DataplexAspectType
type DataplexAspectTypeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexAspectType resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexAspectTypeObservedState `json:"observedState,omitempty"`
}

// DataplexAspectTypeObservedState is the state of the DataplexAspectType resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataplex.v1.AspectType
type DataplexAspectTypeObservedState struct {
	// Output only. System generated globally unique ID for the AspectType.
	//  If you delete and recreate the AspectType with the same name, then this ID
	//  will be different.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the AspectType was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the AspectType was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Denotes the transfer status of the Aspect Type. It is
	//  unspecified for Aspect Types created from Dataplex API.
	// +kcc:proto:field=google.cloud.dataplex.v1.AspectType.transfer_status
	TransferStatus *string `json:"transferStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexaspecttype;gcpdataplexaspecttypes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexAspectType is the Schema for the DataplexAspectType API
// +k8s:openapi-gen=true
type DataplexAspectType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexAspectTypeSpec   `json:"spec,omitempty"`
	Status DataplexAspectTypeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexAspectTypeList contains a list of DataplexAspectType
type DataplexAspectTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexAspectType `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexAspectType{}, &DataplexAspectTypeList{})
}
