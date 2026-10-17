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

// API sources for ChronicleDataTableRow, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/chronicle/v1/data_table.proto
// +kcc:source:service-docs=https://cloud.google.com/chronicle/docs/secops/secops-overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/chronicle/docs/reference/rest/v1/projects.locations.instances.dataTables.dataTableRows

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ChronicleDataTableRowGVK = GroupVersion.WithKind("ChronicleDataTableRow")

// ChronicleDataTableRowSpec defines the desired state of ChronicleDataTableRow
// +kcc:spec:proto=google.cloud.chronicle.v1.DataTableRow
// +kcc:required-from-proto
type ChronicleDataTableRowSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/instances/{instance}/dataTables/{data_table} this resource belongs to.
	// +kcc:guess
	// DataTableRef *ChronicleDataTableRef `json:"dataTableRef,omitempty"`

	// The ChronicleDataTableRow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. All column values for a single row. The values should be in the
	//  same order as the columns of the data tables.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTableRow.values
	// +required
	Values []string `json:"values,omitempty"`

	// Optional. User-provided TTL of the data table row.
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTableRow.row_time_to_live
	RowTimeToLive *string `json:"rowTimeToLive,omitempty"`
}

// ChronicleDataTableRowStatus defines the config connector machine state of ChronicleDataTableRow
type ChronicleDataTableRowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ChronicleDataTableRow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ChronicleDataTableRowObservedState `json:"observedState,omitempty"`
}

// ChronicleDataTableRowObservedState is the state of the ChronicleDataTableRow resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.chronicle.v1.DataTableRow
type ChronicleDataTableRowObservedState struct {
	// Output only. DataTableRow create time
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTableRow.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. DataTableRow update time
	// +kcc:proto:field=google.cloud.chronicle.v1.DataTableRow.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpchronicledatatablerow;gcpchronicledatatablerows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ChronicleDataTableRow is the Schema for the ChronicleDataTableRow API
// +k8s:openapi-gen=true
type ChronicleDataTableRow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ChronicleDataTableRowSpec   `json:"spec,omitempty"`
	Status ChronicleDataTableRowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ChronicleDataTableRowList contains a list of ChronicleDataTableRow
type ChronicleDataTableRowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChronicleDataTableRow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChronicleDataTableRow{}, &ChronicleDataTableRowList{})
}
