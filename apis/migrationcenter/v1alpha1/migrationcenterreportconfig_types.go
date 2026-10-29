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

// API sources for MigrationCenterReportConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/migrationcenter/v1/migrationcenter.proto
// +kcc:source:service-docs=https://cloud.google.com/migration-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/migration-center/docs/reference/rest/v1/projects.locations.reportConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MigrationCenterReportConfigGVK = GroupVersion.WithKind("MigrationCenterReportConfig")

// MigrationCenterReportConfigSpec defines the desired state of MigrationCenterReportConfig
// +kcc:spec:proto=google.cloud.migrationcenter.v1.ReportConfig
// +kcc:required-from-proto
type MigrationCenterReportConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The MigrationCenterReportConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.description
	Description *string `json:"description,omitempty"`

	// Required. Collection of combinations of groups and preference sets.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.group_preferenceset_assignments
	// +required
	GroupPreferencesetAssignments []ReportConfig_GroupPreferenceSetAssignment `json:"groupPreferencesetAssignments,omitempty"`
}

// MigrationCenterReportConfigStatus defines the config connector machine state of MigrationCenterReportConfig
type MigrationCenterReportConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MigrationCenterReportConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MigrationCenterReportConfigObservedState `json:"observedState,omitempty"`
}

// MigrationCenterReportConfigObservedState is the state of the MigrationCenterReportConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.migrationcenter.v1.ReportConfig
type MigrationCenterReportConfigObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmigrationcenterreportconfig;gcpmigrationcenterreportconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MigrationCenterReportConfig is the Schema for the MigrationCenterReportConfig API
// +k8s:openapi-gen=true
type MigrationCenterReportConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MigrationCenterReportConfigSpec   `json:"spec,omitempty"`
	Status MigrationCenterReportConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MigrationCenterReportConfigList contains a list of MigrationCenterReportConfig
type MigrationCenterReportConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationCenterReportConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationCenterReportConfig{}, &MigrationCenterReportConfigList{})
}
