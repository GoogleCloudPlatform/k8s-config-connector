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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataCatalogTaxonomyGVK = GroupVersion.WithKind("DataCatalogTaxonomy")

type DataCatalogTaxonomySpec struct {
	/* A list of policy types that are activated for this taxonomy. If not set,
	defaults to an empty list. Possible values: ["POLICY_TYPE_UNSPECIFIED", "FINE_GRAINED_ACCESS_CONTROL"]. */
	// +optional
	ActivatedPolicyTypes []string `json:"activatedPolicyTypes,omitempty"`

	/* Description of this taxonomy. It must: contain only unicode characters,
	tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
	long when encoded in UTF-8. If not set, defaults to an empty description. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* User defined name of this taxonomy.
	It must: contain only unicode letters, numbers, underscores, dashes
	and spaces; not start or end with spaces; and be at most 200 bytes
	long when encoded in UTF-8. */
	DisplayName string `json:"displayName"`

	/* The project that this resource belongs to. */
	ProjectRef *v1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Taxonomy location region. */
	// +optional
	Region *string `json:"region,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DataCatalogTaxonomyStatus struct {
	/* Conditions represent the latest available observations of the
	   DataCatalogTaxonomy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Resource name of this taxonomy, whose format is:
	"projects/{project}/locations/{region}/taxonomies/{taxonomy}". */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogTaxonomy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogtaxonomy;gcpdatacatalogtaxonomies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// +k8s:openapi-gen=true
// DataCatalogTaxonomy is the Schema for the datacatalog API
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
