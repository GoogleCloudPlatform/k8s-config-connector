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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Organization resource ID that the source projects should belong
	//  to. Projects that do not belong to the provided organization are not
	//  considered when creating the dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.organization_number
	OrganizationNumber *int64 `json:"organizationNumber,omitempty"`

	// Defines the options for providing source projects for the dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.source_projects
	SourceProjects *DatasetConfig_SourceProjects `json:"sourceProjects,omitempty"`

	// Defines the options for providing source folders for the dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.source_folders
	SourceFolders *DatasetConfig_SourceFolders `json:"sourceFolders,omitempty"`

	// Defines the options for providing a source organization for the dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.organization_scope
	OrganizationScope *bool `json:"organizationScope,omitempty"`

	// Input only. Cloud Storage object path containing a list of
	//  project or folder numbers to include in the dataset;
	//  it cannot contain a mix of project and folders.
	//
	//  The object must be a text file where each line has one of the following
	//  entries:
	//
	//  - Project number, formatted as `projects/{project_number}`, for example,
	//  `projects/1234567890`.
	//  - Folder identifier, formatted as `folders/{folder_number}`, for example,
	//  `folders/9876543210`.
	//  Path must be in the format `gs://{bucket_name}/{object_name}`.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.cloud_storage_object_path
	CloudStorageObjectPath *string `json:"cloudStorageObjectPath,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.include_cloud_storage_locations
	IncludeCloudStorageLocations *DatasetConfig_CloudStorageLocations `json:"includeCloudStorageLocations,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.exclude_cloud_storage_locations
	ExcludeCloudStorageLocations *DatasetConfig_CloudStorageLocations `json:"excludeCloudStorageLocations,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.include_cloud_storage_buckets
	IncludeCloudStorageBuckets *DatasetConfig_CloudStorageBuckets `json:"includeCloudStorageBuckets,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.exclude_cloud_storage_buckets
	ExcludeCloudStorageBuckets *DatasetConfig_CloudStorageBuckets `json:"excludeCloudStorageBuckets,omitempty"`

	// If set to `true`, the request includes all the newly created buckets in the
	//  dataset that meet the inclusion and exclusion rules.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.include_newly_created_buckets
	IncludeNewlyCreatedBuckets *bool `json:"includeNewlyCreatedBuckets,omitempty"`

	// Optional. If set to `false`, then all the permission checks must be
	//  successful before the system can start ingesting data. This field can only
	//  be updated before the system ingests data for the first time. Any attempt
	//  to modify the field after data ingestion starts results in an error.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.skip_verification_and_ingest
	SkipVerificationAndIngest *bool `json:"skipVerificationAndIngest,omitempty"`

	// Number of days of history that must be retained.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.retention_period_days
	RetentionPeriodDays *int32 `json:"retentionPeriodDays,omitempty"`

	// Details of the linked dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.link
	Link *DatasetConfig_Link `json:"link,omitempty"`

	// Identity used by this `datasetConfig`.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.identity
	Identity *Identity `json:"identity,omitempty"`

	// Optional. A user-provided description for the dataset configuration.
	//
	//  Maximum length: 256 characters.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.SourceProjects
type DatasetConfig_SourceProjects struct {
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.SourceProjects.project_numbers
	ProjectNumbers []int64 `json:"projectNumbers,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.SourceFolders
type DatasetConfig_SourceFolders struct {
	// Optional. The list of folder numbers to include in the dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.SourceFolders.folder_numbers
	FolderNumbers []int64 `json:"folderNumbers,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageLocations
type DatasetConfig_CloudStorageLocations struct {
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageLocations.locations
	Locations []string `json:"locations,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageBuckets
type DatasetConfig_CloudStorageBuckets struct {
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageBuckets.cloud_storage_buckets
	CloudStorageBuckets []DatasetConfig_CloudStorageBuckets_CloudStorageBucket `json:"cloudStorageBuckets,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageBuckets.CloudStorageBucket
type DatasetConfig_CloudStorageBuckets_CloudStorageBucket struct {
	// Cloud Storage bucket name.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageBuckets.CloudStorageBucket.bucket_name
	BucketName *string `json:"bucketName,omitempty"`

	// A regex pattern for bucket names matching the regex. Regex should
	//  follow the syntax specified in `google/re2` on GitHub.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.CloudStorageBuckets.CloudStorageBucket.bucket_prefix_regex
	BucketPrefixRegex *string `json:"bucketPrefixRegex,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.DatasetConfig.Link
type DatasetConfig_Link struct {
}

// +kcc:proto=google.cloud.storageinsights.v1.Identity
type Identity struct {
	// Type of identity to use for the datasetConfig.
	// +optional
	// +kubebuilder:validation:Enum=IDENTITY_TYPE_UNSPECIFIED;IDENTITY_TYPE_PER_CONFIG;IDENTITY_TYPE_PER_PROJECT
	// +kcc:proto:field=google.cloud.storageinsights.v1.Identity.type
	Type *string `json:"type,omitempty"`
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
	//  This is auto-populated.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The UTC time at which the dataset configuration was last
	//  updated. This is auto-populated.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. System generated unique identifier for the resource.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.uid
	Uid *string `json:"uid,omitempty"`

	// Details of the linked dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.link
	Link *DatasetConfig_LinkObservedState `json:"link,omitempty"`

	// Identity used by this `datasetConfig`.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.identity
	Identity *IdentityObservedState `json:"identity,omitempty"`

	// Output only. Status of the `datasetConfig`.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.status
	Status *common.Status `json:"status,omitempty"`

	// Output only. State of the `datasetConfig`.
	// +optional
	// +kubebuilder:validation:Enum=CONFIG_STATE_UNSPECIFIED;CONFIG_STATE_ACTIVE;CONFIG_STATE_VERIFICATION_IN_PROGRESS;CONFIG_STATE_CREATED;CONFIG_STATE_PROCESSING
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.dataset_config_state
	DatasetConfigState *string `json:"datasetConfigState,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.storageinsights.v1.DatasetConfig.Link
type DatasetConfig_LinkObservedState struct {
	// Output only. Dataset name for linked dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.Link.dataset
	Dataset *string `json:"dataset,omitempty"`

	// Output only. State of the linked dataset.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.DatasetConfig.Link.linked
	Linked *bool `json:"linked,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.storageinsights.v1.Identity
type IdentityObservedState struct {
	// Output only. Name of the identity.
	// +optional
	// +kcc:proto:field=google.cloud.storageinsights.v1.Identity.name
	Name *string `json:"name,omitempty"`
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
