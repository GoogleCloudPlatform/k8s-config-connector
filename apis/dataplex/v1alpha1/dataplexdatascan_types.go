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

var DataplexDataScanGVK = GroupVersion.WithKind("DataplexDataScan")

// DataplexDataScanSpec defines the desired state of DataplexDataScan
// +kcc:spec:proto=google.cloud.dataplex.v1.DataScan
type DataplexDataScanSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The DataplexDataScan name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the scan.
	//
	// * Must be between 1-1024 characters.
	Description *string `json:"description,omitempty"`

	// Optional. User friendly display name.
	//
	// * Must be between 1-256 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the scan.
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The data source for DataScan.
	// +kubebuilder:validation:Required
	Data *DataSource `json:"data"`

	// Optional. DataScan execution settings.
	//
	// If not specified, the fields in it will use their default values.
	ExecutionSpec *DataScan_ExecutionSpec `json:"executionSpec,omitempty"`

	// Settings for a data quality scan.
	DataQualitySpec *DataQualitySpec `json:"dataQualitySpec,omitempty"`

	// Settings for a data profile scan.
	DataProfileSpec *DataProfileSpec `json:"dataProfileSpec,omitempty"`

	// Settings for a data discovery scan.
	DataDiscoverySpec *DataDiscoverySpec `json:"dataDiscoverySpec,omitempty"`
}

// DataplexDataScanStatus defines the config connector machine state of DataplexDataScan
type DataplexDataScanStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexDataScan resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexDataScanObservedState `json:"observedState,omitempty"`
}

// DataplexDataScanObservedState is the state of the DataplexDataScan resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataScan
type DataplexDataScanObservedState struct {
	// Output only. System generated globally unique ID for the scan. This ID will
	// be different if the scan is deleted and re-created with the same name.
	Uid *string `json:"uid,omitempty"`

	// Output only. Current state of the DataScan.
	State *string `json:"state,omitempty"`

	// Output only. The time when the scan was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the scan was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Status of the data scan execution.
	ExecutionStatus *DataScan_ExecutionStatus `json:"executionStatus,omitempty"`

	// Output only. The type of DataScan.
	Type *string `json:"type,omitempty"`

	// Output only. The result of a data quality scan.
	DataQualityResult *DataQualityResultObservedState `json:"dataQualityResult,omitempty"`

	// Output only. The result of a data profile scan.
	DataProfileResult *DataProfileResultObservedState `json:"dataProfileResult,omitempty"`

	// Output only. The result of a data discovery scan.
	DataDiscoveryResult *DataDiscoveryResultObservedState `json:"dataDiscoveryResult,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexdatascan;gcpdataplexdatascans
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexDataScan is the Schema for the DataplexDataScan API
// +k8s:openapi-gen=true
type DataplexDataScan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexDataScanSpec   `json:"spec,omitempty"`
	Status DataplexDataScanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexDataScanList contains a list of DataplexDataScan
type DataplexDataScanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexDataScan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexDataScan{}, &DataplexDataScanList{})
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.JobEndTrigger
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualitySpec_PostScanActions_JobEndTrigger struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.JobFailureTrigger
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualitySpec_PostScanActions_JobFailureTrigger struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataProfileResult.PostScanActionsResult
// +kubebuilder:validation:XPreserveUnknownFields
type DataProfileResult_PostScanActionsResultObservedState struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataQualityResult.PostScanActionsResult
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityResult_PostScanActionsResultObservedState struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.NonNullExpectation
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityRule_NonNullExpectation struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataDiscoveryResult.BigQueryPublishing
// +kubebuilder:validation:XPreserveUnknownFields
type DataDiscoveryResult_BigQueryPublishingObservedState struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataQualityRuleResult
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityRuleResultObservedState struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Trigger.OnDemand
// +kubebuilder:validation:XPreserveUnknownFields
type Trigger_OnDemand struct {
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualityRule.UniquenessExpectation
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityRule_UniquenessExpectation struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataQualityColumnResult
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityColumnResultObservedState struct {
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.DataQualityDimensionResult
// +kubebuilder:validation:XPreserveUnknownFields
type DataQualityDimensionResultObservedState struct {
}

type DataplexEntityRef struct {
	/* A reference to an externally managed Dataplex Entity.
	   Should be of the format `projects/{projectID}/locations/{location}/lakes/{lakeID}/zones/{zoneID}/entities/{entityID}`. */
	External string `json:"external,omitempty"`
}

type BigQueryConnectionRef struct {
	/* The `name` of a `BigQueryConnectionConnection` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` of a `BigQueryConnectionConnection` resource. */
	Namespace string `json:"namespace,omitempty"`
	/* A reference to an externally managed BigQuery Connection.
	   Should be of the format `projects/{projectID}/locations/{location}/connections/{connectionID}`. */
	External string `json:"external,omitempty"`
}

type BigQueryTableRef struct {
	/* The `name` of a `BigQueryTable` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` of a `BigQueryTable` resource. */
	Namespace string `json:"namespace,omitempty"`
	/* A reference to an externally managed BigQuery Table.
	   Should be of the format `//bigquery.googleapis.com/projects/{projectID}/datasets/{datasetID}/tables/{tableID}` or `projects/{projectID}/datasets/{datasetID}/tables/{tableID}`. */
	External string `json:"external,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataSource
type DataSource struct {
	// Immutable. The Dataplex entity that represents the data source (e.g.
	//  BigQuery table) for DataScan, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/entities/{entity_id}`.
	EntityRef *DataplexEntityRef `json:"entityRef,omitempty"`

	// Immutable. The service-qualified full resource name of the cloud resource
	//  for a DataScan job to scan against. The field could be: BigQuery table of
	//  type "TABLE" for DataProfileScan/DataQualityScan Format:
	//  //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	Resource *string `json:"resource,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataDiscoverySpec.BigQueryPublishingConfig
type DataDiscoverySpec_BigQueryPublishingConfig struct {
	// Optional. Determines whether to  publish discovered tables as BigLake
	//  external tables or non-BigLake external tables.
	TableType *string `json:"tableType,omitempty"`

	// Optional. The BigQuery connection used to create BigLake tables.
	//  Must be in the form
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	ConnectionRef *BigQueryConnectionRef `json:"connectionRef,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataProfileSpec.PostScanActions.BigQueryExport
type DataProfileSpec_PostScanActions_BigQueryExport struct {
	// Optional. The BigQuery table to export DataProfileScan results to.
	//  Format:
	//  //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	//  or
	//  projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	ResultsTableRef *BigQueryTableRef `json:"resultsTableRef,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.DataQualitySpec.PostScanActions.BigQueryExport
type DataQualitySpec_PostScanActions_BigQueryExport struct {
	// Optional. The BigQuery table to export DataQualityScan results to.
	//  Format:
	//  //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	//  or
	//  projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
	ResultsTableRef *BigQueryTableRef `json:"resultsTableRef,omitempty"`
}
