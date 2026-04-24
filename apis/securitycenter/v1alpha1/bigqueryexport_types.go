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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	bigquery "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
)

var SecurityCenterBigQueryExportGVK = GroupVersion.WithKind("SecurityCenterBigQueryExport")

// SecurityCenterBigQueryExportSpec defines the desired state of SecurityCenterBigQueryExport
// +kcc:spec:proto=google.cloud.securitycenter.v1.BigQueryExport
type SecurityCenterBigQueryExportSpec struct {
	// The organization, folder, or project that this resource belongs to.
	// +optional
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// +optional
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. The SecurityCenterBigQueryExport name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The description of the export (max of 1024 characters).
	// +optional
	Description *string `json:"description,omitempty"`

	// Expression that defines the filter to apply across create/update events
	// of findings.
	// +optional
	Filter *string `json:"filter,omitempty"`

	// The dataset to write findings' updates to.
	// +required
	DatasetRef *bigquery.DatasetRef `json:"datasetRef,omitempty"`
}

// SecurityCenterBigQueryExportStatus defines the observed state of SecurityCenterBigQueryExport
type SecurityCenterBigQueryExportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SecurityCenterBigQueryExport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SecurityCenterBigQueryExportObservedState `json:"observedState,omitempty"`
}

// SecurityCenterBigQueryExportObservedState defines the observed state of SecurityCenterBigQueryExport
// +kcc:observedstate:proto=google.cloud.securitycenter.v1.BigQueryExport
type SecurityCenterBigQueryExportObservedState struct {
	// Output only. The time at which the BigQuery export was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the BigQuery export was updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Email address of the user who last edited the BigQuery export.
	// +optional
	MostRecentEditor *string `json:"mostRecentEditor,omitempty"`

	// Output only. The service account that needs permission to create table and
	// upload data to the BigQuery dataset.
	// +optional
	Principal *string `json:"principal,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuritycenterbigqueryexport;gcpsecuritycenterbigqueryexports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecurityCenterBigQueryExport is the Schema for the SecurityCenterBigQueryExport API
// +k8s:openapi-gen=true
type SecurityCenterBigQueryExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityCenterBigQueryExportSpec   `json:"spec,omitempty"`
	Status SecurityCenterBigQueryExportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SecurityCenterBigQueryExportList contains a list of SecurityCenterBigQueryExport
type SecurityCenterBigQueryExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterBigQueryExport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityCenterBigQueryExport{}, &SecurityCenterBigQueryExportList{})
}
