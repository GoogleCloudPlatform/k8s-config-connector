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

var DataLabelingDatasetGVK = GroupVersion.WithKind("DataLabelingDataset")

// DataLabelingDatasetSpec defines the desired state of DataLabelingDataset
// +kcc:spec:proto=google.cloud.datalabeling.v1beta1.Dataset
type DataLabelingDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DataLabelingDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the dataset. Maximum of 64 characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification set.
	// The description can be up to 10000 characters long.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`
}

// DataLabelingDatasetStatus defines the config connector machine state of DataLabelingDataset
type DataLabelingDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataLabelingDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataLabelingDatasetObservedState `json:"observedState,omitempty"`
}

// DataLabelingDatasetObservedState is the state of the DataLabelingDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datalabeling.v1beta1.Dataset
type DataLabelingDatasetObservedState struct {
	// Output only. Time the dataset is created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. This is populated with the original input configs
	// where ImportData is called. It is available only after the clients
	// import data to this dataset.
	InputConfigs []InputConfig `json:"inputConfigs,omitempty"`

	// Output only. The names of any related resources that are blocking changes
	// to the dataset.
	BlockingResources []string `json:"blockingResources,omitempty"`

	// Output only. The number of data items in the dataset.
	DataItemCount *int64 `json:"dataItemCount,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalabelingdataset;gcpdatalabelingdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataLabelingDataset is the Schema for the DataLabelingDataset API
// +k8s:openapi-gen=true
type DataLabelingDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataLabelingDatasetSpec   `json:"spec,omitempty"`
	Status DataLabelingDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataLabelingDatasetList contains a list of DataLabelingDataset
type DataLabelingDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLabelingDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataLabelingDataset{}, &DataLabelingDatasetList{})
}
