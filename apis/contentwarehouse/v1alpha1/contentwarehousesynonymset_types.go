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

var ContentWarehouseSynonymSetGVK = GroupVersion.WithKind("ContentWarehouseSynonymSet")

// ContentWarehouseSynonymSetSpec defines the desired state of ContentWarehouseSynonymSet
// +kcc:spec:proto=google.cloud.contentwarehouse.v1.SynonymSet
type ContentWarehouseSynonymSetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// This is a freeform field. Example contexts can be "sales," "engineering,"
	//  "real estate," "accounting," etc.
	//  The context can be supplied during search requests.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.context
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty"`

	// List of Synonyms for the context.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.synonyms
	// +kubebuilder:validation:Optional
	Synonyms []SynonymSetSynonym `json:"synonyms,omitempty"`

	// The ContentWarehouseSynonymSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.SynonymSet.Synonym
type SynonymSetSynonym struct {
	// For example: sale, invoice, bill, order
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.SynonymSet.Synonym.words
	// +kubebuilder:validation:Optional
	Words []string `json:"words,omitempty"`
}

// ContentWarehouseSynonymSetStatus defines the config connector machine state of ContentWarehouseSynonymSet
type ContentWarehouseSynonymSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContentWarehouseSynonymSet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContentWarehouseSynonymSetObservedState `json:"observedState,omitempty"`
}

// ContentWarehouseSynonymSetObservedState is the state of the ContentWarehouseSynonymSet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.contentwarehouse.v1.SynonymSet
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
type ContentWarehouseSynonymSetObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontentwarehousesynonymset;gcpcontentwarehousesynonymsets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContentWarehouseSynonymSet is the Schema for the ContentWarehouseSynonymSet API
// +k8s:openapi-gen=true
type ContentWarehouseSynonymSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContentWarehouseSynonymSetSpec   `json:"spec,omitempty"`
	Status ContentWarehouseSynonymSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContentWarehouseSynonymSetList contains a list of ContentWarehouseSynonymSet
type ContentWarehouseSynonymSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentWarehouseSynonymSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContentWarehouseSynonymSet{}, &ContentWarehouseSynonymSetList{})
}
