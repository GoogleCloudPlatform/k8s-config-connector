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

// API sources for FinancialServicesDataset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/financialservices/v1/dataset.proto
// +kcc:source:service-docs=https://cloud.google.com/financial-services/anti-money-laundering/docs/concepts/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/financial-services/anti-money-laundering/docs/reference/rest/v1/projects.locations.instances.datasets

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FinancialServicesDatasetGVK = GroupVersion.WithKind("FinancialServicesDataset")

// FinancialServicesDatasetSpec defines the desired state of FinancialServicesDataset
// +kcc:spec:proto=google.cloud.financialservices.v1.Dataset
// +kcc:required-from-proto
type FinancialServicesDatasetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project_num}/locations/{location}/instances/{instance} this resource belongs to.
	// +kcc:guess
	// InstanceRef *FinancialServicesInstanceRef `json:"instanceRef,omitempty"`

	// The FinancialServicesDataset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The set of BigQuery tables in the dataset.  The key should be the
	//  table type and the value should be the BigQuery tables in the format
	//  `bq://{project}.{dataset}.{table}`.
	//  Current table types are:
	//
	//    * `party`
	//    * `account_party_link`
	//    * `transaction`
	//    * `risk_case_event`
	//    * `party_supplementary_data`
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.table_specs
	// +required
	TableSpecs map[string]string `json:"tableSpecs,omitempty"`

	// Required. Core time window of the dataset. All tables should have complete
	//  data covering this period.
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.date_range
	// +required
	DateRange *Interval `json:"dateRange,omitempty"`

	// The timezone of the data, default will act as UTC.
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// FinancialServicesDatasetStatus defines the config connector machine state of FinancialServicesDataset
type FinancialServicesDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FinancialServicesDataset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FinancialServicesDatasetObservedState `json:"observedState,omitempty"`
}

// FinancialServicesDatasetObservedState is the state of the FinancialServicesDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.financialservices.v1.Dataset
type FinancialServicesDatasetObservedState struct {
	// Output only. The timestamp of creation of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of the most recent update of this resource.
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the dataset (creating, active, deleting, etc.)
	// +kcc:proto:field=google.cloud.financialservices.v1.Dataset.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfinancialservicesdataset;gcpfinancialservicesdatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FinancialServicesDataset is the Schema for the FinancialServicesDataset API
// +k8s:openapi-gen=true
type FinancialServicesDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FinancialServicesDatasetSpec   `json:"spec,omitempty"`
	Status FinancialServicesDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FinancialServicesDatasetList contains a list of FinancialServicesDataset
type FinancialServicesDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinancialServicesDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FinancialServicesDataset{}, &FinancialServicesDatasetList{})
}
