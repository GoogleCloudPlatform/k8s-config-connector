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

// API sources for DataplexEntry, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/dataplex/v1/catalog.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.entryGroups.entries

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEntryGVK = GroupVersion.WithKind("DataplexEntry")

// DataplexEntrySpec defines the desired state of DataplexEntry
// +kcc:spec:proto=google.cloud.dataplex.v1.Entry
// +kcc:required-from-proto
type DataplexEntrySpec struct {
	// The DataplexEntryGroup this DataplexEntry is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *EntryGroupRef `json:"parentRef"`

	// The DataplexEntry name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The relative resource name of the entry type that was
	//  used to create this entry, in the format
	//  `projects/{project_id_or_number}/locations/{location_id}/entryTypes/{entry_type_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_type
	// +required
	EntryTypeRef *EntryTypeRef `json:"entryTypeRef"`

	// Optional. The aspects that are attached to the entry. Depending on how the
	//  aspect is attached to the entry, the format of the aspect key can be one of
	//  the following:
	//
	//  * If the aspect is attached directly to the entry:
	//  `{project_id_or_number}.{location_id}.{aspect_type_id}`
	//  * If the aspect is attached to an entry's path:
	//  `{project_id_or_number}.{location_id}.{aspect_type_id}@{path}`
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.aspects
	Aspects map[string]Aspect `json:"aspects,omitempty"`

	// Optional. Immutable. The resource name of the parent entry, in the format
	//  `projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}/entries/{entry_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.parent_entry
	ParentEntryRef *DataplexEntryRef `json:"parentEntryRef,omitempty"`

	// Optional. A name for the entry that can be referenced by an external
	//  system. For more information, see [Fully qualified
	//  names](https://cloud.google.com/data-catalog/docs/fully-qualified-names).
	//  The maximum size of the field is 4000 characters.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.fully_qualified_name
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`

	// Optional. Information related to the source system of the data resource
	//  that is represented by the entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_source
	EntrySource *EntrySource `json:"entrySource,omitempty"`
}

// DataplexEntryStatus defines the config connector machine state of DataplexEntry
type DataplexEntryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEntry resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEntryObservedState `json:"observedState,omitempty"`
}

// DataplexEntryObservedState is the state of the DataplexEntry resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Entry
type DataplexEntryObservedState struct {
	// Output only. The time when the entry was created in Dataplex Universal
	//  Catalog.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the entry was last updated in Dataplex Universal
	//  Catalog.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Information related to the source system of the data resource
	//  that is represented by the entry.
	// +kcc:proto:field=google.cloud.dataplex.v1.Entry.entry_source
	EntrySource *EntrySourceObservedState `json:"entrySource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexentry;gcpdataplexentries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEntry is the Schema for the DataplexEntry API
// +k8s:openapi-gen=true
type DataplexEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEntrySpec   `json:"spec,omitempty"`
	Status DataplexEntryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEntryList contains a list of DataplexEntry
type DataplexEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEntry{}, &DataplexEntryList{})
}
