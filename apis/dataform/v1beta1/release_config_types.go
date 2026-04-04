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

// +kcc:proto=google.cloud.dataform.v1beta1.CodeCompilationConfig
type ReleaseConfigCodeCompilationConfig struct {
	/* Optional. The default database (Google Cloud project ID). */
	// +optional
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`

	/* Optional. The default schema (BigQuery dataset ID). */
	// +optional
	DefaultSchema *string `json:"defaultSchema,omitempty"`

	/* Optional. The default BigQuery location to use. */
	// +optional
	DefaultLocation *string `json:"defaultLocation,omitempty"`

	/* Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names. */
	// +optional
	SchemaSuffix *string `json:"schemaSuffix,omitempty"`

	/* Optional. The prefix that should be prepended to all table names. */
	// +optional
	TablePrefix *string `json:"tablePrefix,omitempty"`

	/* Optional. The dependency graph (i.e. assertion) schema name. */
	// +optional
	AssertionSchema *string `json:"assertionSchema,omitempty"`

	/* Optional. User-defined variables that can be read during compilation. */
	// +optional
	Vars map[string]string `json:"vars,omitempty"`

	/* Optional. The suffix that should be appended to all database (Google Cloud project ID) names. */
	// +optional
	DatabaseSuffix *string `json:"databaseSuffix,omitempty"`

	/* Optional. The prefix to prepend to built-in assertion names. */
	// +optional
	BuiltinAssertionNamePrefix *string `json:"builtinAssertionNamePrefix,omitempty"`
}

// +kcc:spec:proto=google.cloud.dataform.v1beta1.ReleaseConfig
type DataformRepositoryReleaseConfigSpec struct {
	/* Immutable. The DataformRepository that this resource belongs to. */
	RepositoryRef refv1beta1.DataformRepositoryRef `json:"repositoryRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Git commitish to be released. */
	// +required
	GitCommitish string `json:"gitCommitish"`

	/* Optional. Optional code compilation config. */
	// +optional
	CodeCompilationConfig *ReleaseConfigCodeCompilationConfig `json:"codeCompilationConfig,omitempty"`

	/* Optional. Optional cron schedule for automated releases. */
	// +optional
	CronSchedule *string `json:"cronSchedule,omitempty"`

	/* Optional. Time zone in which the cron schedule runs. */
	// +optional
	TimeZone *string `json:"timeZone,omitempty"`

	/* Optional. If set to true, releases will be disabled for this config. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`
}

type DataformRepositoryReleaseConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   DataformRepositoryReleaseConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the DataformRepositoryReleaseConfig resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *DataformRepositoryReleaseConfigObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataform.v1beta1.ReleaseConfig
type DataformRepositoryReleaseConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformrepositoryreleaseconfig;gcpdataformrepositoryreleaseconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformRepositoryReleaseConfig is the Schema for the dataform API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type DataformRepositoryReleaseConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformRepositoryReleaseConfigSpec   `json:"spec,omitempty"`
	Status DataformRepositoryReleaseConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataformRepositoryReleaseConfigList contains a list of DataformRepositoryReleaseConfig
type DataformRepositoryReleaseConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformRepositoryReleaseConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformRepositoryReleaseConfig{}, &DataformRepositoryReleaseConfigList{})
}
