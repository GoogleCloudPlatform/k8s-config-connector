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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataCatalogTagTemplateGVK = GroupVersion.WithKind("DataCatalogTagTemplate")

// DataCatalogTagTemplateSpec defines the desired state of DataCatalogTagTemplate
// +kcc:proto=google.cloud.datacatalog.v1.TagTemplate
type DataCatalogTagTemplateSpec struct {
	// The DataCatalogTagTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Display name for this template. Defaults to an empty string.
	//
	//  The name must contain only Unicode letters, numbers (0-9), underscores (_),
	//  dashes (-), spaces ( ), and can't start or end with spaces.
	//  The maximum length is 200 characters.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates whether tags created with this template are public. Public tags
	//  do not require tag template access to appear in
	//  [ListTags][google.cloud.datacatalog.v1.DataCatalog.ListTags] API response.
	//
	//  Additionally, you can search for a public tag by value with a
	//  simple search query in addition to using a ``tag:`` predicate.
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.is_publicly_readable
	IsPubliclyReadable *bool `json:"isPubliclyReadable,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Optional. Transfer status of the TagTemplate
	// +kcc:proto:field=google.cloud.datacatalog.v1.TagTemplate.dataplex_transfer_status
	DataplexTransferStatus *string `json:"dataplexTransferStatus,omitempty"`
}

// DataCatalogTagTemplateStatus defines the config connector machine state of DataCatalogTagTemplate
type DataCatalogTagTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogTagTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogTagTemplateObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtagtemplate;gcpdatacatalogtagtemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogTagTemplate is the Schema for the DataCatalogTagTemplate API
// +k8s:openapi-gen=true
type DataCatalogTagTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogTagTemplateSpec   `json:"spec,omitempty"`
	Status DataCatalogTagTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogTagTemplateList contains a list of DataCatalogTagTemplate
type DataCatalogTagTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogTagTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogTagTemplate{}, &DataCatalogTagTemplateList{})
}
