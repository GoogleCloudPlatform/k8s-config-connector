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

package v1beta1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.dataform.v1beta1.Target
type WorkflowConfigTarget struct {
	/* Optional. The database (Google Cloud project ID). */
	// +optional
	Database *string `json:"database,omitempty"`

	/* Optional. The schema (BigQuery dataset ID). */
	// +optional
	Schema *string `json:"schema,omitempty"`

	/* Optional. The name of the relation (table or view). */
	// +optional
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1beta1.WorkflowConfig.InvocationConfig
type WorkflowConfigInvocationConfig struct {
	/* Optional. The set of action targets to be included. */
	// +optional
	IncludedTargets []WorkflowConfigTarget `json:"includedTargets,omitempty"`

	/* Optional. The set of tags to be included. */
	// +optional
	IncludedTags []string `json:"includedTags,omitempty"`

	/* Optional. When set to true, transitive dependencies of included actions will be executed. */
	// +optional
	TransitiveDependenciesIncluded *bool `json:"transitiveDependenciesIncluded,omitempty"`

	/* Optional. When set to true, transitive dependents of included actions will be executed. */
	// +optional
	TransitiveDependentsIncluded *bool `json:"transitiveDependentsIncluded,omitempty"`

	/* Optional. When set to true, any incremental tables will be fully refreshed. */
	// +optional
	FullyRefreshIncrementalTablesEnabled *bool `json:"fullyRefreshIncrementalTablesEnabled,omitempty"`

	/* Optional. The service account to run workflow invocations under. */
	// +optional
	ServiceAccountRef *refv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:spec:proto=google.cloud.dataform.v1beta1.WorkflowConfig
type DataformRepositoryWorkflowConfigSpec struct {
	/* Immutable. The DataformRepository that this resource belongs to. */
	RepositoryRef refv1beta1.DataformRepositoryRef `json:"repositoryRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The name of the release config that this workflow config uses. */
	// +required
	ReleaseConfigRef refv1beta1.DataformRepositoryReleaseConfigRef `json:"releaseConfigRef"`

	/* Optional. Optional invocation config. */
	// +optional
	InvocationConfig *WorkflowConfigInvocationConfig `json:"invocationConfig,omitempty"`

	/* Optional. Optional cron schedule for automated workflow invocations. */
	// +optional
	CronSchedule *string `json:"cronSchedule,omitempty"`

	/* Optional. Time zone in which the cron schedule runs. */
	// +optional
	TimeZone *string `json:"timeZone,omitempty"`

	/* Optional. If set to true, workflow invocations will be disabled for this config. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`
}

type DataformRepositoryWorkflowConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   DataformRepositoryWorkflowConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the DataformRepositoryWorkflowConfig resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *DataformRepositoryWorkflowConfigObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataform.v1beta1.WorkflowConfig
type DataformRepositoryWorkflowConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformrepositoryworkflowconfig;gcpdataformrepositoryworkflowconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformRepositoryWorkflowConfig is the Schema for the dataform API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type DataformRepositoryWorkflowConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformRepositoryWorkflowConfigSpec   `json:"spec,omitempty"`
	Status DataformRepositoryWorkflowConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataformRepositoryWorkflowConfigList contains a list of DataformRepositoryWorkflowConfig
type DataformRepositoryWorkflowConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformRepositoryWorkflowConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformRepositoryWorkflowConfig{}, &DataformRepositoryWorkflowConfigList{})
}
