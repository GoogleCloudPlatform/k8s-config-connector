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

// API sources for BigLakeHiveTable, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/biglake/hive/v1/hive_metastore.proto
// +kcc:source:service-docs=https://cloud.google.com/products/lakehouse
// +kcc:source:resource-docs=https://cloud.google.com/bigquery/docs/manage-open-source-metadata

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeHiveTableGVK = GroupVersion.WithKind("BigLakeHiveTable")

// BigLakeHiveTableSpec defines the desired state of BigLakeHiveTable
// +kcc:spec:proto=google.cloud.biglake.hive.v1.HiveTable
// +kcc:required-from-proto
type BigLakeHiveTableSpec struct {
	// The BigLakeHiveDatabase this BigLakeHiveTable is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *BigLakeHiveDatabaseRef `json:"parentRef"`

	// The BigLakeHiveTable name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Description of the table. The maximum length is 4000 characters.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.description
	Description *string `json:"description,omitempty"`

	// Required. Storage descriptor of the table.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.storage_descriptor
	// +required
	StorageDescriptor *StorageDescriptor `json:"storageDescriptor,omitempty"`

	// Optional. The partition keys of the table.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.partition_keys
	PartitionKeys []FieldSchema `json:"partitionKeys,omitempty"`

	// Optional. Stores the properties associated with the table. The maximum size
	//  is 4MiB.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. The original view text. Empty for non-view. The maximum size is
	//  16MiB.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.view_original_text
	ViewOriginalText *string `json:"viewOriginalText,omitempty"`

	// Optional. The expanded view text. Empty for non-view. The maximum size is
	//  16MiB.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.view_expanded_text
	ViewExpandedText *string `json:"viewExpandedText,omitempty"`
}

// BigLakeHiveTableStatus defines the config connector machine state of BigLakeHiveTable
type BigLakeHiveTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeHiveTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeHiveTableObservedState `json:"observedState,omitempty"`
}

// BigLakeHiveTableObservedState is the state of the BigLakeHiveTable resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.biglake.hive.v1.HiveTable
type BigLakeHiveTableObservedState struct {
	// Output only. The creation time of the table.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The type of the table. This is EXTERNAL for BigLake hive
	//  tables.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.table_type
	TableType *string `json:"tableType,omitempty"`

	// Output only. The update time of the table.
	// +kcc:proto:field=google.cloud.biglake.hive.v1.HiveTable.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakehivetable;gcpbiglakehivetables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeHiveTable is the Schema for the BigLakeHiveTable API
// +k8s:openapi-gen=true
type BigLakeHiveTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeHiveTableSpec   `json:"spec,omitempty"`
	Status BigLakeHiveTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeHiveTableList contains a list of BigLakeHiveTable
type BigLakeHiveTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeHiveTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeHiveTable{}, &BigLakeHiveTableList{})
}
