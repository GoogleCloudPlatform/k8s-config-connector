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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexDataTaxonomyGVK = GroupVersion.WithKind("DataplexDataTaxonomy")

// DataplexDataTaxonomySpec defines the desired state of DataplexDataTaxonomy
// +kcc:spec:proto=google.cloud.dataplex.v1.DataTaxonomy
type DataplexDataTaxonomySpec struct {
	ParentRef *parent.ProjectAndLocationRef `json:",inline"`

	// The DataplexDataTaxonomy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.description
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels for the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.labels
	// Labels map[string]string `json:"labels,omitempty"`
}

// DataplexDataTaxonomyStatus defines the config connector machine state of DataplexDataTaxonomy
type DataplexDataTaxonomyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexDataTaxonomy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexDataTaxonomyObservedState `json:"observedState,omitempty"`
}

// DataplexDataTaxonomyObservedState is the state of the DataplexDataTaxonomy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataTaxonomy
type DataplexDataTaxonomyObservedState struct {
	// Output only. System generated globally unique ID for the dataTaxonomy. This
	//  ID will be different if the DataTaxonomy is deleted and re-created with the
	//  same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the DataTaxonomy was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the DataTaxonomy was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of attributes in the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.attribute_count
	AttributeCount *int32 `json:"attributeCount,omitempty"`

	// Output only. The number of classes in the DataTaxonomy.
	// +kcc:proto:field=google.cloud.dataplex.v1.DataTaxonomy.class_count
	ClassCount *int32 `json:"classCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexdatataxonomy;gcpdataplexdatataxonomys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexDataTaxonomy is the Schema for the DataplexDataTaxonomy API
// +k8s:openapi-gen=true
type DataplexDataTaxonomy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexDataTaxonomySpec   `json:"spec,omitempty"`
	Status DataplexDataTaxonomyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexDataTaxonomyList contains a list of DataplexDataTaxonomy
type DataplexDataTaxonomyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexDataTaxonomy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexDataTaxonomy{}, &DataplexDataTaxonomyList{})
}
