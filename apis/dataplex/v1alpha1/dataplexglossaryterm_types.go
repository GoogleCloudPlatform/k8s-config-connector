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

// API sources for DataplexGlossaryTerm, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/business_glossary.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.glossaries.terms

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexGlossaryTermGVK = GroupVersion.WithKind("DataplexGlossaryTerm")

// DataplexGlossaryTermSpec defines the desired state of DataplexGlossaryTerm
// +kcc:spec:proto=google.cloud.dataplex.v1.GlossaryTerm
// +kcc:required-from-proto
type DataplexGlossaryTermSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexGlossaryTerm name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User friendly display name of the GlossaryTerm. This is
	//  user-mutable. This will be same as the GlossaryTermId, if not specified.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The user-mutable description of the GlossaryTerm.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.description
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels for the GlossaryTerm.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The immediate parent of the GlossaryTerm in the
	//  resource-hierarchy. It can either be a Glossary or a GlossaryCategory.
	//  Format:
	//  projects/{project_id_or_number}/locations/{location_id}/glossaries/{glossary_id}
	//  OR
	//  projects/{project_id_or_number}/locations/{location_id}/glossaries/{glossary_id}/categories/{category_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.parent
	// +required
	Parent *string `json:"parent,omitempty"`
}

// DataplexGlossaryTermStatus defines the config connector machine state of DataplexGlossaryTerm
type DataplexGlossaryTermStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexGlossaryTerm resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexGlossaryTermObservedState `json:"observedState,omitempty"`
}

// DataplexGlossaryTermObservedState is the state of the DataplexGlossaryTerm resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.GlossaryTerm
type DataplexGlossaryTermObservedState struct {
	// Output only. System generated unique id for the GlossaryTerm. This ID will
	//  be different if the GlossaryTerm is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which the GlossaryTerm was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the GlossaryTerm was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryTerm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexglossaryterm;gcpdataplexglossaryterms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexGlossaryTerm is the Schema for the DataplexGlossaryTerm API
// +k8s:openapi-gen=true
type DataplexGlossaryTerm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexGlossaryTermSpec   `json:"spec,omitempty"`
	Status DataplexGlossaryTermStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexGlossaryTermList contains a list of DataplexGlossaryTerm
type DataplexGlossaryTermList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexGlossaryTerm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexGlossaryTerm{}, &DataplexGlossaryTermList{})
}
