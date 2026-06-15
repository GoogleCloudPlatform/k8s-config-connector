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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexGlossaryGVK = GroupVersion.WithKind("DataplexGlossary")

// DataplexGlossarySpec defines the desired state of DataplexGlossary
// +kcc:spec:proto=google.cloud.dataplex.v1.Glossary
type DataplexGlossarySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexGlossary name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User friendly display name of the Glossary. This is user-mutable.
	//  This will be same as the GlossaryId, if not specified.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The user-mutable description of the Glossary.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.description
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels for the Glossary.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// DataplexGlossaryStatus defines the config connector machine state of DataplexGlossary
type DataplexGlossaryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexGlossary resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *DataplexGlossaryObservedState `json:"observedState,omitempty"`
}

// DataplexGlossaryObservedState is the state of the DataplexGlossary resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.Glossary
type DataplexGlossaryObservedState struct {
	// Output only. System generated unique id for the Glossary. This ID will be
	//  different if the Glossary is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which the Glossary was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the Glossary was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of GlossaryTerms in the Glossary.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.term_count
	TermCount *int32 `json:"termCount,omitempty"`

	// Output only. The number of GlossaryCategories in the Glossary.
	// +kcc:proto:field=google.cloud.dataplex.v1.Glossary.category_count
	CategoryCount *int32 `json:"categoryCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexglossary;gcpdataplexglossaries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexGlossary is the Schema for the DataplexGlossary API
// +k8s:openapi-gen=true
type DataplexGlossary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexGlossarySpec   `json:"spec,omitempty"`
	Status DataplexGlossaryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexGlossaryList contains a list of DataplexGlossary
type DataplexGlossaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexGlossary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexGlossary{}, &DataplexGlossaryList{})
}
