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

// API sources for DataplexGlossaryCategory, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/business_glossary.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/projects.locations.glossaries.categories

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexGlossaryCategoryGVK = GroupVersion.WithKind("DataplexGlossaryCategory")

// DataplexGlossaryCategorySpec defines the desired state of DataplexGlossaryCategory
// +kcc:spec:proto=google.cloud.dataplex.v1.GlossaryCategory
// +kcc:required-from-proto
type DataplexGlossaryCategorySpec struct {
	// Required. The Glossary that this GlossaryCategory belongs to.
	// +required
	GlossaryRef *GlossaryRef `json:"glossaryRef"`

	// The DataplexGlossaryCategory name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User friendly display name of the GlossaryCategory. This is
	//  user-mutable. This will be same as the GlossaryCategoryId, if not
	//  specified.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The user-mutable description of the GlossaryCategory.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.description
	Description *string `json:"description,omitempty"`

	// Required. The immediate parent of the GlossaryCategory in the
	//  resource-hierarchy.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.parent
	// +required
	ParentRef *GlossaryRef `json:"parentRef"`
}

// DataplexGlossaryCategoryStatus defines the config connector machine state of DataplexGlossaryCategory
type DataplexGlossaryCategoryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexGlossaryCategory resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexGlossaryCategoryObservedState `json:"observedState,omitempty"`
}

// DataplexGlossaryCategoryObservedState is the state of the DataplexGlossaryCategory resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.GlossaryCategory
type DataplexGlossaryCategoryObservedState struct {
	// Output only. System generated unique id for the GlossaryCategory. This ID
	//  will be different if the GlossaryCategory is deleted and re-created with
	//  the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time at which the GlossaryCategory was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the GlossaryCategory was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.GlossaryCategory.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexglossarycategory;gcpdataplexglossarycategories
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexGlossaryCategory is the Schema for the DataplexGlossaryCategory API
// +k8s:openapi-gen=true
type DataplexGlossaryCategory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexGlossaryCategorySpec   `json:"spec,omitempty"`
	Status DataplexGlossaryCategoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexGlossaryCategoryList contains a list of DataplexGlossaryCategory
type DataplexGlossaryCategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexGlossaryCategory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexGlossaryCategory{}, &DataplexGlossaryCategoryList{})
}
