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

var MapsPlatformDatasetsDatasetGVK = GroupVersion.WithKind("MapsPlatformDatasetsDataset")

// MapsPlatformDatasetsDatasetSpec defines the desired state of MapsPlatformDatasetsDataset
// +kcc:spec:proto=google.maps.mapsplatformdatasets.v1.Dataset
type MapsPlatformDatasetsDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The MapsPlatformDatasetsDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human readable name, shown in the console UI.
	// Must be unique within a project.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`

	// A description of this dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// The version ID of the dataset.
	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionID,omitempty"`

	// Specified use case for this dataset.
	// +kubebuilder:validation:Optional
	Usage []string `json:"usage,omitempty"`

	// A local file source for the dataset for a single upload.
	// +kubebuilder:validation:Optional
	LocalFileSource *LocalFileSource `json:"localFileSource,omitempty"`

	// A Google Cloud Storage file source for the dataset for a single upload.
	// +kubebuilder:validation:Optional
	GCSSource *GCSSource `json:"gcsSource,omitempty"`
}

type GCSSource struct {
	// Source data URI. For example, `gs://my_bucket/my_object`.
	// +kubebuilder:validation:Optional
	InputURI *string `json:"inputURI,omitempty"`

	// The file format of the Google Cloud Storage object. This is used mainly for validation.
	// +kubebuilder:validation:Optional
	FileFormat *string `json:"fileFormat,omitempty"`
}

type LocalFileSource struct {
	// The file name of the uploaded file.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty"`

	// The format of the file that is being uploaded.
	// +kubebuilder:validation:Optional
	FileFormat *string `json:"fileFormat,omitempty"`
}

type Status struct {
	// State enum for status.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`

	// Error message indicating reason of failure. It is empty if the datasets is not in a failed state.
	// +kubebuilder:validation:Optional
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// MapsPlatformDatasetsDatasetStatus defines the config connector machine state of MapsPlatformDatasetsDataset
type MapsPlatformDatasetsDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MapsPlatformDatasetsDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MapsPlatformDatasetsDatasetObservedState `json:"observedState,omitempty"`
}

// MapsPlatformDatasetsDatasetObservedState is the state of the MapsPlatformDatasetsDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.maps.mapsplatformdatasets.v1.Dataset
type MapsPlatformDatasetsDatasetObservedState struct {
	// Output only. The status of this dataset version.
	// +kcc:proto:field=google.maps.mapsplatformdatasets.v1.Dataset.status
	Status *Status `json:"status,omitempty"`

	// Output only. Time when the dataset was first created.
	// +kcc:proto:field=google.maps.mapsplatformdatasets.v1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the dataset metadata was last updated.
	// +kcc:proto:field=google.maps.mapsplatformdatasets.v1.Dataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this version was created.
	// +kcc:proto:field=google.maps.mapsplatformdatasets.v1.Dataset.version_create_time
	VersionCreateTime *string `json:"versionCreateTime,omitempty"`

	// Output only. The description for this version of dataset. It is provided when importing data to the dataset.
	// +kcc:proto:field=google.maps.mapsplatformdatasets.v1.Dataset.version_description
	VersionDescription *string `json:"versionDescription,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmapsplatformdatasetsdataset;gcpmapsplatformdatasetsdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MapsPlatformDatasetsDataset is the Schema for the MapsPlatformDatasetsDataset API
// +k8s:openapi-gen=true
type MapsPlatformDatasetsDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MapsPlatformDatasetsDatasetSpec   `json:"spec,omitempty"`
	Status MapsPlatformDatasetsDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MapsPlatformDatasetsDatasetList contains a list of MapsPlatformDatasetsDataset
type MapsPlatformDatasetsDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapsPlatformDatasetsDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MapsPlatformDatasetsDataset{}, &MapsPlatformDatasetsDatasetList{})
}
