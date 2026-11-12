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

// API sources for AIPlatformContext, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/context.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.contexts

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformContextGVK = GroupVersion.WithKind("AIPlatformContext")

// AIPlatformContextSpec defines the desired state of AIPlatformContext
// +kcc:spec:proto=google.cloud.aiplatform.v1.Context
// +kcc:required-from-proto
type AIPlatformContextSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformContext name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User provided display name of the Context.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Contexts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Context (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Context.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.metadata
	Metadata apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Description of the Context
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.description
	Description *string `json:"description,omitempty"`
}

// AIPlatformContextStatus defines the config connector machine state of AIPlatformContext
type AIPlatformContextStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformContext resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformContextObservedState `json:"observedState,omitempty"`
}

// AIPlatformContextObservedState is the state of the AIPlatformContext resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.Context
type AIPlatformContextObservedState struct {
	// Output only. Timestamp when this Context was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Context was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of resource names of Contexts that are parents of this
	//  Context. A Context may have at most 10 parent_contexts.
	// +kcc:proto:field=google.cloud.aiplatform.v1.Context.parent_contexts
	ParentContexts []string `json:"parentContexts,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformcontext;gcpaiplatformcontexts
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformContext is the Schema for the AIPlatformContext API
// +k8s:openapi-gen=true
type AIPlatformContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformContextSpec   `json:"spec,omitempty"`
	Status AIPlatformContextStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformContextList contains a list of AIPlatformContext
type AIPlatformContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformContext{}, &AIPlatformContextList{})
}
