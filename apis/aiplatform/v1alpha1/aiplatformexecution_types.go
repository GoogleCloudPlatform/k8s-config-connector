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

// API sources for AIPlatformExecution, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/execution.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.executions

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformExecutionGVK = GroupVersion.WithKind("AIPlatformExecution")

// AIPlatformExecutionSpec defines the desired state of AIPlatformExecution
// +kcc:spec:proto=google.cloud.aiplatform.v1.Execution
// +kcc:required-from-proto
type AIPlatformExecutionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformExecution name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User provided display name of the Execution.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The state of this Execution. This is a property of the Execution, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines) and the system does not prescribe
	//  or check the validity of state transitions.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.state
	State *string `json:"state,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Executions.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Execution (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in `schema_title` to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Execution.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.metadata
	Metadata apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Execution
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.description
	Description *string `json:"description,omitempty"`
}

// AIPlatformExecutionStatus defines the config connector machine state of AIPlatformExecution
type AIPlatformExecutionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformExecution resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformExecutionObservedState `json:"observedState,omitempty"`
}

// AIPlatformExecutionObservedState is the state of the AIPlatformExecution resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Execution
type AIPlatformExecutionObservedState struct {
	// Output only. Timestamp when this Execution was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Execution was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Execution.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformexecution;gcpaiplatformexecutions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformExecution is the Schema for the AIPlatformExecution API
// +k8s:openapi-gen=true
type AIPlatformExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformExecutionSpec   `json:"spec,omitempty"`
	Status AIPlatformExecutionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformExecutionList contains a list of AIPlatformExecution
type AIPlatformExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformExecution `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformExecution{}, &AIPlatformExecutionList{})
}
