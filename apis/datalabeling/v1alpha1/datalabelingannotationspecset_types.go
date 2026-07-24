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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/kccscheme"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataLabelingAnnotationSpecSetGVK = GroupVersion.WithKind("DataLabelingAnnotationSpecSet")

// DataLabelingAnnotationSpecSetSpec defines the desired state of DataLabelingAnnotationSpecSet
// +kcc:spec:proto=google.cloud.datalabeling.v1beta1.AnnotationSpecSet
type DataLabelingAnnotationSpecSetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DataLabelingAnnotationSpecSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name for AnnotationSpecSet that you define when you
	//  create it. Maximum of 64 characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification set.
	//  The description can be up to 10,000 characters long.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Required. The array of AnnotationSpecs that you define when you create the
	//  AnnotationSpecSet. These are the possible labels for the labeling task.
	// +kubebuilder:validation:Required
	AnnotationSpecs []AnnotationSpec `json:"annotationSpecs,omitempty"`
}

// DataLabelingAnnotationSpecSetStatus defines the config connector machine state of DataLabelingAnnotationSpecSet
type DataLabelingAnnotationSpecSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataLabelingAnnotationSpecSet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataLabelingAnnotationSpecSetObservedState `json:"observedState,omitempty"`
}

// DataLabelingAnnotationSpecSetObservedState is the state of the DataLabelingAnnotationSpecSet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datalabeling.v1beta1.AnnotationSpecSet
type DataLabelingAnnotationSpecSetObservedState struct {
	// Output only. The names of any related resources that are blocking changes
	//  to the annotation spec set.
	BlockingResources []string `json:"blockingResources,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalabelingannotationspecset;gcpdatalabelingannotationspecsets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataLabelingAnnotationSpecSet is the Schema for the DataLabelingAnnotationSpecSet API
// +k8s:openapi-gen=true
type DataLabelingAnnotationSpecSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataLabelingAnnotationSpecSetSpec   `json:"spec,omitempty"`
	Status DataLabelingAnnotationSpecSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataLabelingAnnotationSpecSetList contains a list of DataLabelingAnnotationSpecSet
type DataLabelingAnnotationSpecSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLabelingAnnotationSpecSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataLabelingAnnotationSpecSet{}, &DataLabelingAnnotationSpecSetList{})
	kccscheme.RegisterType(DataLabelingAnnotationSpecSetGVK, &DataLabelingAnnotationSpecSet{})
}
