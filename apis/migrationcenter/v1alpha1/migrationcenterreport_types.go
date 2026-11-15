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

// API sources for MigrationCenterReport, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/migrationcenter/v1/migrationcenter.proto
// +kcc:source:service-docs=https://cloud.google.com/migration-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/migration-center/docs/reference/rest/v1/projects.locations.reportConfigs.reports

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MigrationCenterReportGVK = GroupVersion.WithKind("MigrationCenterReport")

// MigrationCenterReportSpec defines the desired state of MigrationCenterReport
// +kcc:spec:proto=google.cloud.migrationcenter.v1.Report
// +kcc:required-from-proto
type MigrationCenterReportSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/reportConfigs/{report_config} this resource belongs to.
	// +kcc:guess
	// ReportConfigRef *MigrationCenterReportConfigRef `json:"reportConfigRef,omitempty"`

	// The MigrationCenterReport name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.description
	Description *string `json:"description,omitempty"`

	// Report type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.type
	Type *string `json:"type,omitempty"`

	// Report creation state.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.state
	State *string `json:"state,omitempty"`
}

// MigrationCenterReportStatus defines the config connector machine state of MigrationCenterReport
type MigrationCenterReportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MigrationCenterReport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MigrationCenterReportObservedState `json:"observedState,omitempty"`
}

// MigrationCenterReportObservedState is the state of the MigrationCenterReport resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.migrationcenter.v1.Report
type MigrationCenterReportObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Summary view of the Report.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Report.summary
	Summary *ReportSummary `json:"summary,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmigrationcenterreport;gcpmigrationcenterreports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MigrationCenterReport is the Schema for the MigrationCenterReport API
// +k8s:openapi-gen=true
type MigrationCenterReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MigrationCenterReportSpec   `json:"spec,omitempty"`
	Status MigrationCenterReportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MigrationCenterReportList contains a list of MigrationCenterReport
type MigrationCenterReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationCenterReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationCenterReport{}, &MigrationCenterReportList{})
}
