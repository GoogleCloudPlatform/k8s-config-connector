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

// API sources for DataplexEntryLink, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/dataplex/v1/catalog.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.entryGroups.entryLinks

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEntryLinkGVK = GroupVersion.WithKind("DataplexEntryLink")

// DataplexEntryLinkSpec defines the desired state of DataplexEntryLink
// +kcc:spec:proto=google.cloud.dataplex.v1.EntryLink
// +kcc:required-from-proto
type DataplexEntryLinkSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/entryGroups/{entry_group} this resource belongs to.
	// +kcc:guess
	// EntryGroupRef *EntryGroupRef `json:"entryGroupRef,omitempty"`

	// The DataplexEntryLink name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Immutable. Relative resource name of the Entry Link Type used to
	//  create this Entry Link. For example:
	//
	//  - Entry link between synonym terms in a glossary:
	//  `projects/dataplex-types/locations/global/entryLinkTypes/synonym`
	//  - Entry link between related terms in a glossary:
	//  `projects/dataplex-types/locations/global/entryLinkTypes/related`
	//  - Entry link between glossary terms and data assets:
	//  `projects/dataplex-types/locations/global/entryLinkTypes/definition`
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryLink.entry_link_type
	// +required
	EntryLinkType *string `json:"entryLinkType,omitempty"`

	// Optional. The aspects that are attached to the entry link.
	//  The format of the aspect key has to be the following:
	//  `{project_id_or_number}.{location_id}.{aspect_type_id}`
	//  Currently, only a single aspect of a Dataplex-owned Aspect Type is allowed.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryLink.aspects
	Aspects map[string]Aspect `json:"aspects,omitempty"`

	// Required. Immutable. Specifies the Entries referenced in the Entry Link.
	//  There should be exactly two entry references.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryLink.entry_references
	// +required
	EntryReferences []EntryLink_EntryReference `json:"entryReferences,omitempty"`
}

// DataplexEntryLinkStatus defines the config connector machine state of DataplexEntryLink
type DataplexEntryLinkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEntryLink resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEntryLinkObservedState `json:"observedState,omitempty"`
}

// DataplexEntryLinkObservedState is the state of the DataplexEntryLink resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.EntryLink
type DataplexEntryLinkObservedState struct {
	// Output only. The time when the Entry Link was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryLink.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the Entry Link was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EntryLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexentrylink;gcpdataplexentrylinks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEntryLink is the Schema for the DataplexEntryLink API
// +k8s:openapi-gen=true
type DataplexEntryLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEntryLinkSpec   `json:"spec,omitempty"`
	Status DataplexEntryLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEntryLinkList contains a list of DataplexEntryLink
type DataplexEntryLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEntryLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEntryLink{}, &DataplexEntryLinkList{})
}
