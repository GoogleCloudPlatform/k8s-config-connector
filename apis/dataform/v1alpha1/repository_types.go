// Copyright 2024 Google LLC
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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


type RepositoryGitRemoteSettings struct {
	/* The name of the Secret Manager secret version to use as an authentication token for Git operations. Must be in the format projects/* /secrets/* /versions/*. */
	AuthenticationTokenSecretVersion string `json:"authenticationTokenSecretVersion"`

	/* The Git remote's default branch name. */
	DefaultBranch string `json:"defaultBranch"`

	/* Indicates the status of the Git access token. https://cloud.google.com/dataform/reference/rest/v1beta1/projects.locations.repositories#TokenStatus. */
	// +optional
	TokenStatus *string `json:"tokenStatus,omitempty"`

	/* The Git remote's URL. */
	Url string `json:"url"`
}

type RepositoryWorkspaceCompilationOverrides struct {
	/* Optional. The default database (Google Cloud project ID). */
	// +optional
	DefaultDatabase *string `json:"defaultDatabase,omitempty"`

	/* Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names. */
	// +optional
	SchemaSuffix *string `json:"schemaSuffix,omitempty"`

	/* Optional. The prefix that should be prepended to all table names. */
	// +optional
	TablePrefix *string `json:"tablePrefix,omitempty"`
}

type DataformRepositorySpec struct {
	/* Optional. If set, configures this repository to be linked to a Git remote. */
	// +optional
	GitRemoteSettings *RepositoryGitRemoteSettings `json:"gitRemoteSettings,omitempty"`

	/* The project that this resource belongs to. */
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. A reference to the region. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Optional. If set, fields of workspaceCompilationOverrides override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. */
	// +optional
	WorkspaceCompilationOverrides *RepositoryWorkspaceCompilationOverrides `json:"workspaceCompilationOverrides,omitempty"`
}

type DataformRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   DataformRepository's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformrepository;gcpdataformrepositories
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformRepository is the Schema for the dataform API
// +k8s:openapi-gen=true
type DataformRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataformRepositorySpec   `json:"spec,omitempty"`
	Status DataformRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataformRepositoryList contains a list of DataformRepository
type DataformRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformRepository{}, &DataformRepositoryList{})
}
