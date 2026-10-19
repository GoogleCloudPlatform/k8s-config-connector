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

// API sources for AIPlatformDatasetVersion, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/dataset_version.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.datasets.datasetVersions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformDatasetVersionGVK = GroupVersion.WithKind("AIPlatformDatasetVersion")

// AIPlatformDatasetVersionSpec defines the desired state of AIPlatformDatasetVersion
// +kcc:spec:proto=google.cloud.aiplatform.v1.DatasetVersion
// +kcc:required-from-proto
type AIPlatformDatasetVersionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformDatasetVersion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.etag
	Etag *string `json:"etag,omitempty"`

	// The user-defined name of the DatasetVersion.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// AIPlatformDatasetVersionStatus defines the config connector machine state of AIPlatformDatasetVersion
type AIPlatformDatasetVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformDatasetVersion resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformDatasetVersionObservedState `json:"observedState,omitempty"`
}

// AIPlatformDatasetVersionObservedState is the state of the AIPlatformDatasetVersion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.DatasetVersion
type AIPlatformDatasetVersionObservedState struct {
	// Output only. Timestamp when this DatasetVersion was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this DatasetVersion was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Name of the associated BigQuery dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.big_query_dataset_name
	BigQueryDatasetName *string `json:"bigQueryDatasetName,omitempty"`

	// Required. Output only. Additional information about the DatasetVersion.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.metadata
	Metadata *Value `json:"metadata,omitempty"`

	// Output only. Reference to the public base model last used by the dataset
	//  version. Only set for prompt dataset versions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.model_reference
	ModelReference *string `json:"modelReference,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.DatasetVersion.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformdatasetversion;gcpaiplatformdatasetversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformDatasetVersion is the Schema for the AIPlatformDatasetVersion API
// +k8s:openapi-gen=true
type AIPlatformDatasetVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformDatasetVersionSpec   `json:"spec,omitempty"`
	Status AIPlatformDatasetVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformDatasetVersionList contains a list of AIPlatformDatasetVersion
type AIPlatformDatasetVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformDatasetVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformDatasetVersion{}, &AIPlatformDatasetVersionList{})
}
