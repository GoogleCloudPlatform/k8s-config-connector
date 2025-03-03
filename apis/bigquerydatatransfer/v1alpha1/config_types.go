// Copyright 2024 Google LLC
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
	bigquery "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryDataTransferConfigGVK = GroupVersion.WithKind("BigQueryDataTransferConfig")

// +kcc:proto=google.cloud.bigquery.datatransfer.v1.EncryptionConfiguration
type EncryptionConfiguration struct {
	// The KMS key used for encrypting BigQuery data.
	KmsKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// BigQueryDataTransferConfigSpec defines the desired state of BigQueryDataTransferConfig
// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferConfig
type BigQueryDataTransferConfigSpec struct {
	// The number of days to look back to automatically refresh the data.
	//  For example, if `data_refresh_window_days = 10`, then every day
	//  BigQuery reingests data for [today-10, today-1], rather than ingesting data
	//  for just [today-1].
	//  Only valid if the data source supports the feature. Set the value to 0
	//  to use the default value.
	DataRefreshWindowDays *int32 `json:"dataRefreshWindowDays,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="DataSourceID field is immutable"
	// Immutable.
	// Data source ID. This cannot be changed once data transfer is created. The
	//  full list of available data source IDs can be returned through an API call:
	//  https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	// +required
	DataSourceID *string `json:"dataSourceID,omitempty"`

	// The BigQuery target dataset id.
	DatasetRef *bigquery.DatasetRef `json:"datasetRef,omitempty"`

	// Is this config disabled. When set to true, no runs will be scheduled for
	//  this transfer config.
	Disabled *bool `json:"disabled,omitempty"`

	// User specified display name for the data transfer.
	DisplayName *string `json:"displayName,omitempty"`

	// Email notifications will be sent according to these preferences
	//  to the email address of the user who owns this transfer config.
	EmailPreferences *EmailPreferences `json:"emailPreferences,omitempty"`

	// The encryption configuration part. Currently, it is only used for the
	//  optional KMS key name. The BigQuery service account of your project must be
	//  granted permissions to use the key. Read methods will return the key name
	//  applied in effect. Write methods will apply the key if it is present, or
	//  otherwise try to apply project default keys if it is absent.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	// Pub/Sub topic where notifications will be sent after transfer runs
	//  associated with this transfer config finish.
	PubSubTopicRef *refv1beta1.PubSubTopicRef `json:"pubSubTopicRef,omitempty"`

	// Parameters specific to each data source. For more information see the
	//  bq tab in the 'Setting up a data transfer' section for each data source.
	//  For example the parameters for Cloud Storage transfers are listed here:
	//  https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	// +required
	Params map[string]string `json:"params,omitempty"`

	Parent `json:",inline"`

	// The BigQueryDataTransferConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Data transfer schedule.
	//  If the data source does not support a custom schedule, this should be
	//  empty. If it is empty, the default value for the data source will be used.
	//  The specified times are in UTC.
	//  Examples of valid format:
	//  `1st,3rd monday of month 15:30`,
	//  `every wed,fri of jan,jun 13:15`, and
	//  `first sunday of quarter 00:00`.
	//  See more explanation about the format here:
	//  https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	//
	//  NOTE: The minimum interval time between recurring transfers depends on the
	//  data source; refer to the documentation for your data source.
	Schedule *string `json:"schedule,omitempty"`

	// Options customizing the data transfer schedule.
	ScheduleOptions *ScheduleOptions `json:"scheduleOptions,omitempty"`

	// Service account email. If this field is set, the transfer config will be created with this service account's credentials.
	//  It requires that the requesting user calling this API has permissions to act as this service account.
	//  Note that not all data sources support service account credentials when creating a transfer config.
	//  For the latest list of data sources, please refer to https://cloud.google.com/bigquery/docs/use-service-accounts.
	ServiceAccountRef *refv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// BigQueryDataTransferConfigStatus defines the config connector machine state of BigQueryDataTransferConfig
type BigQueryDataTransferConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryDataTransferConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryDataTransferConfigObservedState `json:"observedState,omitempty"`
}

// BigQueryDataTransferConfigSpec defines the desired state of BigQueryDataTransferConfig
// +kcc:proto=google.cloud.bigquery.datatransfer.v1.TransferConfig
type BigQueryDataTransferConfigObservedState struct {
	// Output only. Region in which BigQuery dataset is located.
	DatasetRegion *string `json:"datasetRegion,omitempty"`

	// Identifier. The resource name of the transfer config.
	//  Transfer config names have the form either
	//  `projects/{project_id}/locations/{region}/transferConfigs/{config_id}` or
	//  `projects/{project_id}/transferConfigs/{config_id}`,
	//  where `config_id` is usually a UUID, even though it is not
	//  guaranteed or required. The name is ignored when creating a transfer
	//  config.
	Name *string `json:"name,omitempty"`

	// Output only. Next time when data transfer will run.
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. Information about the user whose credentials are used to
	//  transfer data. Populated only for `transferConfigs.get` requests. In case
	//  the user information is not available, this field will not be populated.
	OwnerInfo *UserInfo `json:"ownerInfo,omitempty"`

	// Output only. State of the most recently updated transfer run.
	State *string `json:"state,omitempty"`

	// Output only. Data transfer modification time. Ignored by server on input.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserID *int64 `json:"userID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerydatatransferconfig;gcpbigquerydatatransferconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryDataTransferConfig is the Schema for the BigQueryDataTransferConfig API
// +k8s:openapi-gen=true
type BigQueryDataTransferConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryDataTransferConfigSpec   `json:"spec,omitempty"`
	Status BigQueryDataTransferConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryDataTransferConfigList contains a list of BigQueryDataTransferConfig
type BigQueryDataTransferConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryDataTransferConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryDataTransferConfig{}, &BigQueryDataTransferConfigList{})
}
