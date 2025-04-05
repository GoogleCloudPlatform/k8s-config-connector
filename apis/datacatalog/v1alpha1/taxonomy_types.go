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

var DataCatalogTaxonomyGVK = GroupVersion.WithKind("DataCatalogTaxonomy")

// DataCatalogTaxonomySpec defines the desired state of DataCatalogTaxonomy
// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy
type DataCatalogTaxonomySpec struct {
	// The DataCatalogTaxonomy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	*Parent `json:",inline"`

	// Required. User-defined name of this taxonomy.
	//
	//  The name can't start or end with spaces, must contain only Unicode letters,
	//  numbers, underscores, dashes, and spaces, and be at most 200 bytes long
	//  when encoded in UTF-8.
	//
	//  The taxonomy display name must be unique within an organization.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of this taxonomy. If not set, defaults to empty.
	//
	//  The description must contain only Unicode characters, tabs, newlines,
	//  carriage returns, and page breaks, and be at most 2000 bytes long when
	//  encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.description
	Description *string `json:"description,omitempty"`

	// Optional. A list of policy types that are activated for this taxonomy. If
	//  not set, defaults to an empty list.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.activated_policy_types
	ActivatedPolicyTypes []string `json:"activatedPolicyTypes,omitempty"`
}

// DataCatalogTaxonomyStatus defines the config connector machine state of DataCatalogTaxonomy
type DataCatalogTaxonomyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogTaxonomy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogTaxonomyObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy.Service
type Taxonomy_Service struct {
	// The Google Cloud service name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.Service.name
	Name *string `json:"name,omitempty"`

	// The service agent for the service.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.Service.identity
	Identity *string `json:"identity,omitempty"`
}

// DataCatalogTaxonomyObservedState is the state of the DataCatalogTaxonomy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datacatalog.v1.Taxonomy
type DataCatalogTaxonomyObservedState struct {
	// Output only. Number of policy tags in this taxonomy.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.policy_tag_count
	PolicyTagCount *int32 `json:"policyTagCount,omitempty"`

	// Output only. Creation and modification timestamps of this taxonomy.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.taxonomy_timestamps
	TaxonomyTimestamps *SystemTimestamps `json:"taxonomyTimestamps,omitempty"`

	// Output only. Identity of the service which owns the Taxonomy. This field is
	//  only populated when the taxonomy is created by a Google Cloud service.
	//  Currently only 'DATAPLEX' is supported.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Taxonomy.service
	Service *Taxonomy_Service `json:"service,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtaxonomy;gcpdatacatalogtaxonomies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogTaxonomy is the Schema for the DataCatalogTaxonomy API
// +k8s:openapi-gen=true
type DataCatalogTaxonomy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogTaxonomySpec   `json:"spec,omitempty"`
	Status DataCatalogTaxonomyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogTaxonomyList contains a list of DataCatalogTaxonomy
type DataCatalogTaxonomyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogTaxonomy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogTaxonomy{}, &DataCatalogTaxonomyList{})
}
