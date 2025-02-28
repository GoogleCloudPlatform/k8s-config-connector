// Copyright 2025 Google LLC
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

var BigLakeTableGVK = GroupVersion.WithKind("BigLakeTable")

// BigLakeTableSpec defines the desired state of BigLakeTable
// +kcc:proto=google.cloud.bigquery.biglake.v1.Table
type BigLakeTableSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The BigLakeCatalog that this resource belongs to.
	// +required
	CatalogRef *CatalogRef `json:"catalogRef"`

	// Required. The parent resource where this table will be created.
	// Format:
	// projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	// +required
	DatabaseRef *DatabaseRef `json:"databaseRef"`

	// The BigLake Table ID. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The location where the Table should reside.
	// +required
	Location *string `json:"location,omitempty"`

	// The table type.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.type
	// +optional
	Type *string `json:"type,omitempty"`

	// Options of a Hive table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.hive_options
	// +optional
	HiveOptions *HiveTableOptions `json:"hiveOptions,omitempty"`

	// The checksum of a table object computed by the server based on the value of
	//  other fields. It may be sent on update requests to ensure the client has an
	//  up-to-date value before proceeding. It is only checked for update table
	//  operations.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.etag
	// +optional
	Etag *string `json:"etag,omitempty"`
}

// BigLakeTableStatus defines the config connector machine state of BigLakeTable
type BigLakeTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeTableObservedState `json:"observedState,omitempty"`
}

// BigLakeTableObservedState is the state of the BigLakeTable resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.biglake.v1.Table
type BigLakeTableObservedState struct {
	// Output only. The creation time of the table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time of the table. Only set after the table is
	//  deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time when this table is considered expired. Only set after
	//  the table is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglaketable;gcpbiglaketables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeTable is the Schema for the BigLakeTable API
// +k8s:openapi-gen=true
type BigLakeTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeTableSpec   `json:"spec,omitempty"`
	Status BigLakeTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeTableList contains a list of BigLakeTable
type BigLakeTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeTable{}, &BigLakeTableList{})
}
