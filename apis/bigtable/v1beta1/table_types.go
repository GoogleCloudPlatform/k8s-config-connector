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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableTableGVK = GroupVersion.WithKind("BigtableTable")

// BigtableTableSpec defines the desired state of BigtableTable
// +kcc:spec:proto=google.bigtable.admin.v2.Table
type BigtableTableSpec struct {
	// Immutable. The instance to create the table in.
	// +required
	InstanceRef InstanceRef `json:"instanceRef"`

	// The BigtableTable name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The names of the column families that should be created immediately upon
	// table creation, specified by name. The values that may be set are specified
	// here. At least one column family must be specified.
	ColumnFamily []*TableColumnFamily `json:"columnFamily,omitempty"`

	//DEPRECATED: The ChangeStreamRetention no longer exists in the new proto, but has to map to ChangeStreamConfig.

	// Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days..
	ChangeStreamRetention *string `json:"changeStreamRetention,omitempty"`

	// NOTYET
	// A group of columns within a table which share a common configuration.
	// ChangeStreamConfig *ChangeStreamConfig `json:"changeStreamConfig,omitempty"`

	// NOTYET
	// The granularity (i.e. MILLIS) at which timestamps are stored in this table.
	// Timestamps not matching the granularity will be rejected. If unspecified at
	// creation time, the value will be set to MILLIS.
	// Granularity *string `json:"granularity,omitempty"`

	// A list of predefined keys to split the table on.
	SplitKeys []string `json:"splitKeys,omitempty"`

	// NOTE: DeletionProtection proto field is changed from string (1.38) to bool (1.40) in cloud.google.com/go/bigtable/admin/apiv2/adminpb
	// Set to true to make the table protected against data loss. i.e. deleting
	// the following resources through Admin APIs are prohibited:
	//
	// * The table.
	// * The column families in the table.
	// * The instance containing the table.
	//
	// Note one can still delete the data stored in the table through Data APIs.
	DeletionProtection *string `json:"deletionProtection,omitempty"`

	// NOTYET
	// If specified, automated backups are enabled for this table.
	// Otherwise, automated backups are disabled.
	// *Table_AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.ColumnFamily
type TableColumnFamily struct {
	// The name of the column family.
	FamilyID string `json:"family"`

	// NOTYET
	// Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	// GcRule *GcRule `json:"gcRule,omitempty"`

	// NOTYET
	// The type of data stored in each of this family's cell values, including its
	// full encoding. If omitted, the family only serves raw untyped bytes.
	//
	// For now, only the `Aggregate` type is supported.
	//
	// `Aggregate` can only be set at family creation and is immutable afterwards.
	//
	//
	// If `value_type` is `Aggregate`, written data must be compatible with:
	//  * `value_type.input_type` for `AddInput` mutations
	// ValueType *Type `json:"valueType,omitempty"`
}

// BigtableTableStatus defines the config connector machine state of BigtableTable
type BigtableTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *BigtableTableObservedState `json:"observedState,omitempty"`
}

// BigtableTableObservedState is the state of the BigtableTable resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.Table
type BigtableTableObservedState struct {
	// The unique name of the table. Values are of the form
	// projects/{project}/instances/{instance}/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*.
	// Views: NAME_ONLY, SCHEMA_VIEW, REPLICATION_VIEW, FULL
	Name *string `json:"name,omitempty"`

	// This field is populated only when the table is in a REPLICATION_VIEW.
	ClusterStates map[string]Table_ClusterState `json:"clusterStates,omitempty"`

	// If this table was restored from another data source (e.g. a backup), this
	// field will be populated with information about the restore.
	RestoreInfo *RestoreInfo `json:"restoreInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtabletable;gcpbigtabletables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableTable is the Schema for the BigtableTable API
// +k8s:openapi-gen=true
type BigtableTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableTableSpec   `json:"spec,omitempty"`
	Status BigtableTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableTableList contains a list of BigtableTable
type BigtableTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableTable{}, &BigtableTableList{})
}
