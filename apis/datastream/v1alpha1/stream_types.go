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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatastreamStreamGVK = GroupVersion.WithKind("DatastreamStream")

// DatastreamStreamSpec defines the desired state of DatastreamStream
// +kcc:proto=google.cloud.datastream.v1.Stream
type DatastreamStreamSpec struct {
	// The DatastreamStream name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. A reference to the project */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location this stream.
	Location string `json:"location"`

	// The source connection profile configuration.
	SourceConfig *SourceConfigSpec `json:"sourceConfig,omitempty"`

	// The destination connection profile configuration.
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`

	// Display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Backfill strategy.
	BackfillAll  *Stream_BackfillAllStrategy  `json:"backfillAll,omitempty"`
	BackfillNone *Stream_BackfillNoneStrategy `json:"backfillNone,omitempty"`
}

// DatastreamStreamStatus defines the config connector machine state of DatastreamStream
type DatastreamStreamStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatastreamStream resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatastreamStreamObservedState `json:"observedState,omitempty"`
}

// DatastreamStreamSpec defines the desired state of DatastreamStream
// +kcc:proto=google.cloud.datastream.v1.Stream
// DatastreamStreamObservedState is the state of the DatastreamStream resource as most recently observed in GCP.
type DatastreamStreamObservedState struct {
	// Output only. The stream's name.
	Name *string `json:"name,omitempty"`

	// Output only. Create time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the stream.
	State *string `json:"state,omitempty"`

	// Output only. A list of errors that occurred on the stream.
	Errors []Error `json:"errors,omitempty"`

	// Output only. The customer-managed encryption key's resource name, if the
	//  stream is encrypted with customer-managed encryption key.
	CustomerManagedEncryptionKey *string `json:"customerManagedEncryptionKey,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceConfig
type SourceConfigSpec struct {
	// Required. Source connection profile resoource.
	//  Format: `projects/{project}/locations/{location}/connectionProfiles/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.source_connection_profile
	SourceConnectionProfileRef *ConnectionProfileRef `json:"sourceConnectionProfileRef,omitempty"`

	// Oracle data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.oracle_source_config
	OracleSourceConfig *OracleSourceConfig `json:"oracleSourceConfig,omitempty"`

	// MySQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.mysql_source_config
	MysqlSourceConfig *MysqlSourceConfig `json:"mysqlSourceConfig,omitempty"`

	// PostgreSQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.postgresql_source_config
	PostgresqlSourceConfig *PostgresqlSourceConfig `json:"postgresqlSourceConfig,omitempty"`

	// SQLServer data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.sql_server_source_config
	// This field is no longer supported in proto.
	// SqlServerSourceConfig *SqlServerSourceConfig `json:"sqlServerSourceConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.DestinationConfig
type DestinationConfig struct {
	// Required. Destination connection profile resource.
	//  Format: `projects/{project}/locations/{location}/connectionProfiles/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.destination_connection_profile
	DestinationConnectionProfileRef *ConnectionProfileRef `json:"destinationConnectionProfileRef,omitempty"`

	// A configuration for how data should be loaded to Cloud Storage.
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.gcs_destination_config
	GcsDestinationConfig *GcsDestinationConfig `json:"gcsDestinationConfig,omitempty"`

	// BigQuery destination configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.DestinationConfig.bigquery_destination_config
	BigqueryDestinationConfig *BigQueryDestinationConfigSpec `json:"bigqueryDestinationConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig
type BigQueryDestinationConfigSpec struct {
	// Single destination dataset.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.single_target_dataset
	SingleTargetDataset *BigQueryDestinationConfig_SingleTargetDataset `json:"singleTargetDataset,omitempty"`

	// Source hierarchy datasets.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.source_hierarchy_datasets
	SourceHierarchyDatasets *BigQueryDestinationConfig_SourceHierarchyDatasets `json:"sourceHierarchyDatasets,omitempty"`

	// The guaranteed data freshness (in seconds) when querying tables created by
	//  the stream. Editing this field will only affect new tables created in the
	//  future, but existing tables will not be impacted. Lower values mean that
	//  queries will return fresher data, but may result in higher cost.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.data_freshness
	DataFreshness *string `json:"dataFreshness,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SingleTargetDataset
type BigQueryDestinationConfig_SingleTargetDataset struct {
	// The BigQuery dataset reference of the target dataset.
	//  DatasetIds allowed characters:
	//  https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#datasetreference.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SingleTargetDataset.dataset_id
	DatasetRef *refs.BigQueryDatasetRef `json:"datasetRef,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate
type BigQueryDestinationConfig_SourceHierarchyDatasets_DatasetTemplate struct {
	// Required. The geographic location where the dataset should reside. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.location
	Location *string `json:"location,omitempty"`

	// If supplied, every created dataset will have its name prefixed by the
	//  provided value. The prefix and name will be separated by an underscore.
	//  i.e. <prefix>_<dataset_name>.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.dataset_id_prefix
	DatasetIDPrefix *string `json:"datasetIDPrefix,omitempty"`

	// Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account
	//  associated with your project requires access to this encryption key.
	//  i.e.
	//  projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
	//  See https://cloud.google.com/bigquery/docs/customer-managed-encryption
	//  for more information.
	// +kcc:proto:field=google.cloud.datastream.v1.BigQueryDestinationConfig.SourceHierarchyDatasets.DatasetTemplate.kms_key_name
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamstream;gcpdatastreamstreams
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// DatastreamStream is the Schema for the DatastreamStream API
// +k8s:openapi-gen=true
type DatastreamStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatastreamStreamSpec   `json:"spec,omitempty"`
	Status DatastreamStreamStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatastreamStreamList contains a list of DatastreamStream
type DatastreamStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamStream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamStream{}, &DatastreamStreamList{})
}
