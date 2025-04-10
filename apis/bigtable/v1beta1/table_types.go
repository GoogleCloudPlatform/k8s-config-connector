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
	// The table name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The BigtableInstance containing this table.
	// +required
	InstanceRef *InstanceRef `json:"instanceRef,omitempty"`

	// How long the change stream should be retained. Change stream data older
	//  than the retention period will not be returned when reading the change
	//  stream from the table.
	//  Values must be at least 1 day and at most 7 days, and will be truncated to
	//  microsecond granularity.
	// +kcc:proto:field=google.bigtable.admin.v2.ChangeStreamConfig.retention_period
	ChangeStreamRetention *string `json:"changeStreamRetention,omitempty"`

	// The column families configured for this table, mapped by column family ID.
	//  Views: `SCHEMA_VIEW`, `STATS_VIEW`, `FULL`
	// +kcc:proto:field=google.bigtable.admin.v2.Table.column_families
	ColumnFamily []Table_ColumnFamilies `json:"columnFamily,omitempty"`

	/* NOTYET: terraform backcompat
	// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored
	//  in this table. Timestamps not matching the granularity will be rejected. If
	//  unspecified at creation time, the value will be set to `MILLIS`. Views:
	//  `SCHEMA_VIEW`, `FULL`.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.granularity
	Granularity *string `json:"granularity,omitempty"`
	*/

	/* NOTYET: terraform backcompat
	// If specified, enable the change stream on this table.
	//  Otherwise, the change stream is disabled and the change stream is not
	//  retained.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.change_stream_config
	ChangeStreamConfig *ChangeStreamConfig `json:"changeStreamConfig,omitempty"`
	*/

	// Set to PROTECTED to make the table protected against data loss. i.e. deleting
	//  the following resources through Admin APIs are prohibited:
	//
	//  * The table.
	//  * The column families in the table.
	//  * The instance containing the table.
	//
	//  Note one can still delete the data stored in the table through Data APIs.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.deletion_protection
	DeletionProtection *string `json:"deletionProtection,omitempty"`

	SplitKeys []string `json:"splitKeys,omitempty"`

	/* NOTYET: terraform backcompat
	// If specified, automated backups are enabled for this table.
	//  Otherwise, automated backups are disabled.
	// +kcc:proto:field=google.bigtable.admin.v2.Table.automated_backup_policy
	AutomatedBackupPolicy *Table_AutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
	*/

}

type Table_ColumnFamilies struct {
	// The name of the column family.
	// +required
	Family *string `json:"family,omitempty"`
}

// BigtableTableStatus defines the config connector machine state of BigtableTable
// +kcc:status:proto=google.bigtable.admin.v2.Table
type BigtableTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET: terraform backcompat
	// A unique specifier for the BigtableTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigtableTableObservedState `json:"observedState,omitempty"`
	*/
}

// BigtableTableObservedState is the state of the BigtableTable resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.Table
type BigtableTableObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtabletable;gcpbigtabletables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"
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
