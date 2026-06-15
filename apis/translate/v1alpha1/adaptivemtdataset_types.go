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

var TranslateAdaptiveMtDatasetGVK = GroupVersion.WithKind("TranslateAdaptiveMtDataset")

// TranslateAdaptiveMtDatasetSpec defines the desired state of TranslateAdaptiveMtDataset
// +kcc:spec:proto=google.cloud.translation.v3.AdaptiveMtDataset
type TranslateAdaptiveMtDatasetSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The TranslateAdaptiveMtDataset name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The name of the dataset to show in the interface.
	// +optional
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The BCP-47 language code of the source language.
	// +required
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.source_language_code
	SourceLanguageCode *string `json:"sourceLanguageCode,omitempty"`

	// Required. The BCP-47 language code of the target language.
	// +required
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.target_language_code
	TargetLanguageCode *string `json:"targetLanguageCode,omitempty"`
}

// TranslateAdaptiveMtDatasetStatus defines the config connector machine state of TranslateAdaptiveMtDataset
type TranslateAdaptiveMtDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranslateAdaptiveMtDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TranslateAdaptiveMtDatasetObservedState `json:"observedState,omitempty"`
}

// TranslateAdaptiveMtDatasetObservedState is the state of the TranslateAdaptiveMtDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.translation.v3.AdaptiveMtDataset
type TranslateAdaptiveMtDatasetObservedState struct {
	// Output only. Timestamp when this dataset was created.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this dataset was last updated.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The number of examples in the dataset.
	// +kcc:proto:field=google.cloud.translation.v3.AdaptiveMtDataset.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranslateadaptivemtdataset;gcptranslateadaptivemtdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranslateAdaptiveMtDataset is the Schema for the TranslateAdaptiveMtDataset API
// +k8s:openapi-gen=true
type TranslateAdaptiveMtDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranslateAdaptiveMtDatasetSpec   `json:"spec,omitempty"`
	Status TranslateAdaptiveMtDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranslateAdaptiveMtDatasetList contains a list of TranslateAdaptiveMtDataset
type TranslateAdaptiveMtDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranslateAdaptiveMtDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranslateAdaptiveMtDataset{}, &TranslateAdaptiveMtDatasetList{})
}
