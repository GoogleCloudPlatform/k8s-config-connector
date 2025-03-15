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

var DataCatalogEntryGroupGVK = GroupVersion.WithKind("DataCatalogEntryGroup")

// DataCatalogEntryGroupSpec defines the desired state of DataCatalogEntryGroup
// +kcc:proto=google.cloud.datacatalog.v1.EntryGroup
type DataCatalogEntryGroupSpec struct {
	// The DataCatalogEntryGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A short name to identify the entry group, for example,
	//  "analytics data - jan 2011". Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Entry group description. Can consist of several sentences or
	//  paragraphs that describe the entry group contents.
	//  Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. When set to [true], it means DataCatalog EntryGroup was
	//  transferred to Dataplex Catalog Service. It makes EntryGroup and its
	//  Entries to be read-only in DataCatalog. However, new Tags on EntryGroup and
	//  its Entries can be created. After setting the flag to [true] it cannot be
	//  unset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.transferred_to_dataplex
	TransferredToDataplex *bool `json:"transferredToDataplex,omitempty"`
}

// DataCatalogEntryGroupStatus defines the config connector machine state of DataCatalogEntryGroup
type DataCatalogEntryGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogEntryGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogEntryGroupObservedState `json:"observedState,omitempty"`
}

// DataCatalogEntryGroupObservedState is the state of the DataCatalogEntryGroup resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datacatalog.v1.EntryGroup
type DataCatalogEntryGroupObservedState struct {
	// Output only. Timestamps of the entry group. Default value is empty.
	// +kcc:proto:field=google.cloud.datacatalog.v1.EntryGroup.data_catalog_timestamps
	DataCatalogTimestamps *SystemTimestamps `json:"dataCatalogTimestamps,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogentrygroup;gcpdatacatalogentrygroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogEntryGroup is the Schema for the DataCatalogEntryGroup API
// +k8s:openapi-gen=true
type DataCatalogEntryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogEntryGroupSpec   `json:"spec,omitempty"`
	Status DataCatalogEntryGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogEntryGroupList contains a list of DataCatalogEntryGroup
type DataCatalogEntryGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogEntryGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogEntryGroup{}, &DataCatalogEntryGroupList{})
}
