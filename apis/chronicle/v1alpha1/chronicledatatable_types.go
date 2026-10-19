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

// API sources for ChronicleDataTable, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/chronicle/v1/data_table.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.dataTables

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleDataTableGVK = GroupVersion.WithKind("ChronicleDataTable")

// ChronicleDataTableSpec defines the desired state of ChronicleDataTable
// +kcc:spec:proto=google.cloud.chronicle.v1.DataTable
// +kcc:required-from-proto
type ChronicleDataTableSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ChronicleDataTable name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. A user-provided description of the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.description
	// +required
	Description *string `json:"description,omitempty"`

	// Immutable. Details of all the columns in the table
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.column_info
	ColumnInfo []DataTableColumnInfo `json:"columnInfo,omitempty"`

	// Optional. User-provided TTL of the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.row_time_to_live
	RowTimeToLive *string `json:"rowTimeToLive,omitempty"`

	// Optional. The scope info of the data table.
	//  During data table creation, if this field is not set, the data
	//  table without scopes (an unscoped table) will be created for a global
	//  user. For a scoped user, this field must be set. During data table
	//  update, if scope_info is requested to be updated, this field must be set.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.scope_info
	ScopeInfo *DataTableScopeInfo `json:"scopeInfo,omitempty"`
}

// ChronicleDataTableStatus defines the config connector machine state of ChronicleDataTable
type ChronicleDataTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleDataTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleDataTableObservedState `json:"observedState,omitempty"`
}

// ChronicleDataTableObservedState is the state of the ChronicleDataTable resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.DataTable
type ChronicleDataTableObservedState struct {
	// Output only. The unique display name of the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Table create time
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Table update time
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Data table unique id
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.data_table_uuid
	DataTableUuid *string `json:"dataTableUuid,omitempty"`

	// Output only. The resource names for the associated Rules that use this
	//  data table. Format:
	//  projects/{project}/locations/{location}/instances/{instance}/rules/{rule}.
	//  {rule} here refers to the rule id.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.rules
	Rules []string `json:"rules,omitempty"`

	// Output only. The count of rules using the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.rule_associations_count
	RuleAssociationsCount *int32 `json:"ruleAssociationsCount,omitempty"`

	// Output only. The count of rows in the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.approximate_row_count
	ApproximateRowCount *int64 `json:"approximateRowCount,omitempty"`

	// Output only. Source of the data table update.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.update_source
	UpdateSource *string `json:"updateSource,omitempty"`

	// Output only. Last update time of the TTL of the data table.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTable.row_time_to_live_update_time
	RowTimeToLiveUpdateTime *string `json:"rowTimeToLiveUpdateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchronicledatatable;gcpchronicledatatables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleDataTable is the Schema for the ChronicleDataTable API
// +k8s:openapi-gen=true
type ChronicleDataTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleDataTableSpec   `json:"spec,omitempty"`
	Status ChronicleDataTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleDataTableList contains a list of ChronicleDataTable
type ChronicleDataTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleDataTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleDataTable{}, &ChronicleDataTableList{})
}
