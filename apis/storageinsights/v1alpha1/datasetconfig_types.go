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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageInsightsDatasetConfigGVK = GroupVersion.WithKind("StorageInsightsDatasetConfig")

type DatasetConfig_SourceProjects struct {
	// Optional. The list of project numbers to include in the dataset.
	ProjectNumbers []int64 `json:"projectNumbers,omitempty"`
}

type DatasetConfig_SourceFolders struct {
	// Optional. The list of folder numbers to include in the dataset.
	FolderNumbers []int64 `json:"folderNumbers,omitempty"`
}

type DatasetConfig_CloudStorageLocations struct {
	Locations []string `json:"locations,omitempty"`
}

type DatasetConfig_CloudStorageBuckets_CloudStorageBucket struct {
	// Cloud Storage bucket reference.
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// A regex pattern for bucket names matching the regex. Regex should
	// follow the syntax specified in `google/re2` on GitHub.
	BucketPrefixRegex *string `json:"bucketPrefixRegex,omitempty"`
}

type DatasetConfig_CloudStorageBuckets struct {
	CloudStorageBuckets []DatasetConfig_CloudStorageBuckets_CloudStorageBucket `json:"cloudStorageBuckets,omitempty"`
}

type DatasetConfig_Link struct {
}

type Identity struct {
	// Type of identity to use for the datasetConfig.
	Type *string `json:"type,omitempty"`
}

type DatasetConfig_LinkObservedState struct {
	// Output only. Dataset name for linked dataset.
	Dataset *string `json:"dataset,omitempty"`

	// Output only. State of the linked dataset.
	Linked *bool `json:"linked,omitempty"`
}

type IdentityObservedState struct {
	// Output only. Name of the identity.
	Name *string `json:"name,omitempty"`
}

// StorageInsightsDatasetConfigSpec defines the desired state of StorageInsightsDatasetConfig
// +kcc:spec:proto=google.cloud.storageinsights.v1.DatasetConfig
type StorageInsightsDatasetConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The StorageInsightsDatasetConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Labels as key value pairs
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Organization resource ID that the source projects should belong
	// to. Projects that do not belong to the provided organization are not
	// considered when creating the dataset.
	OrganizationNumber *int64 `json:"organizationNumber,omitempty"`

	// Defines the options for providing source projects for the dataset.
	SourceProjects *DatasetConfig_SourceProjects `json:"sourceProjects,omitempty"`

	// Defines the options for providing source folders for the dataset.
	SourceFolders *DatasetConfig_SourceFolders `json:"sourceFolders,omitempty"`

	// Defines the options for providing a source organization for the dataset.
	OrganizationScope *bool `json:"organizationScope,omitempty"`

	// Input only. Cloud Storage object path containing a list of
	// project or folder numbers to include in the dataset;
	// it cannot contain a mix of project and folders.
	//
	// The object must be a text file where each line has one of the following
	// entries:
	//
	// - Project number, formatted as `projects/<project_number>`, for example,
	// `projects/1234567890`.
	// - Folder identifier, formatted as `folders/<folder_number>`, for example,
	// `folders/9876543210`.
	// Path must be in the format `gs://<bucket_name>/<object_name>`.
	CloudStorageObjectPath *string `json:"cloudStorageObjectPath,omitempty"`

	IncludeCloudStorageLocations *DatasetConfig_CloudStorageLocations `json:"includeCloudStorageLocations,omitempty"`

	ExcludeCloudStorageLocations *DatasetConfig_CloudStorageLocations `json:"excludeCloudStorageLocations,omitempty"`

	IncludeCloudStorageBuckets *DatasetConfig_CloudStorageBuckets `json:"includeCloudStorageBuckets,omitempty"`

	ExcludeCloudStorageBuckets *DatasetConfig_CloudStorageBuckets `json:"excludeCloudStorageBuckets,omitempty"`

	// If set to `true`, the request includes all the newly created buckets in the
	// dataset that meet the inclusion and exclusion rules.
	IncludeNewlyCreatedBuckets *bool `json:"includeNewlyCreatedBuckets,omitempty"`

	// Optional. If set to `false`, then all the permission checks must be
	// successful before the system can start ingesting data. This field can only
	// be updated before the system ingests data for the first time. Any attempt
	// to modify the field after data ingestion starts results in an error.
	SkipVerificationAndIngest *bool `json:"skipVerificationAndIngest,omitempty"`

	// Number of days of history that must be retained.
	RetentionPeriodDays *int32 `json:"retentionPeriodDays,omitempty"`

	// Details of the linked dataset.
	// +kubebuilder:validation:XPreserveUnknownFields
	Link *DatasetConfig_Link `json:"link,omitempty"`

	// Identity used by this `datasetConfig`.
	Identity *Identity `json:"identity,omitempty"`

	// Optional. A user-provided description for the dataset configuration.
	//
	// Maximum length: 256 characters.
	Description *string `json:"description,omitempty"`
}

// StorageInsightsDatasetConfigStatus defines the config connector machine state of StorageInsightsDatasetConfig
type StorageInsightsDatasetConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageInsightsDatasetConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageInsightsDatasetConfigObservedState `json:"observedState,omitempty"`
}

// StorageInsightsDatasetConfigObservedState is the state of the StorageInsightsDatasetConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.storageinsights.v1.DatasetConfig
type StorageInsightsDatasetConfigObservedState struct {
	// Output only. The UTC time at which the dataset configuration was created.
	// This is auto-populated.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The UTC time at which the dataset configuration was last
	// updated. This is auto-populated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. System generated unique identifier for the resource.
	Uid *string `json:"uid,omitempty"`

	// Details of the linked dataset.
	Link *DatasetConfig_LinkObservedState `json:"link,omitempty"`

	// Identity used by this `datasetConfig`.
	Identity *IdentityObservedState `json:"identity,omitempty"`

	// Output only. Status of the `datasetConfig`.
	Status *common.Status `json:"status,omitempty"`

	// Output only. State of the `datasetConfig`.
	DatasetConfigState *string `json:"datasetConfigState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstorageinsightsdatasetconfig;gcpstorageinsightsdatasetconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageInsightsDatasetConfig is the Schema for the StorageInsightsDatasetConfig API
// +k8s:openapi-gen=true
type StorageInsightsDatasetConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageInsightsDatasetConfigSpec   `json:"spec,omitempty"`
	Status StorageInsightsDatasetConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageInsightsDatasetConfigList contains a list of StorageInsightsDatasetConfig
type StorageInsightsDatasetConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageInsightsDatasetConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageInsightsDatasetConfig{}, &StorageInsightsDatasetConfigList{})
}
