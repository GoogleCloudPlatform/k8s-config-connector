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

// API sources for AIPlatformRagCorpus, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/aiplatform/v1/vertex_rag_data.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.ragCorpora

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformRagCorpusGVK = GroupVersion.WithKind("AIPlatformRagCorpus")

// AIPlatformRagCorpusSpec defines the desired state of AIPlatformRagCorpus
// +kcc:spec:proto=google.cloud.aiplatform.v1.RagCorpus
// +kcc:required-from-proto
type AIPlatformRagCorpusSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformRagCorpus name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Immutable. The config for the Vector DBs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vector_db_config
	VectorDbConfig *RagVectorDbConfig `json:"vectorDbConfig,omitempty"`

	// Optional. Immutable. The config for the Vertex AI Search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vertex_ai_search_config
	VertexAiSearchConfig *VertexAiSearchConfig `json:"vertexAiSearchConfig,omitempty"`

	// Required. The display name of the RagCorpus.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the RagCorpus.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.description
	Description *string `json:"description,omitempty"`

	// Optional. Immutable. The CMEK key name used to encrypt at-rest data related
	//  to this Corpus. Only applicable to RagManagedDb option for Vector DB. This
	//  field can only be set at corpus creation time, and cannot be updated or
	//  deleted.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// AIPlatformRagCorpusStatus defines the config connector machine state of AIPlatformRagCorpus
type AIPlatformRagCorpusStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformRagCorpus resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformRagCorpusObservedState `json:"observedState,omitempty"`
}

// AIPlatformRagCorpusObservedState is the state of the AIPlatformRagCorpus resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.RagCorpus
type AIPlatformRagCorpusObservedState struct {
	// Optional. Immutable. The config for the Vector DBs.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.vector_db_config
	VectorDbConfig *RagVectorDbConfigObservedState `json:"vectorDbConfig,omitempty"`

	// Output only. Timestamp when this RagCorpus was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this RagCorpus was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. RagCorpus state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.RagCorpus.corpus_status
	CorpusStatus *CorpusStatusObservedState `json:"corpusStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformragcorpus;gcpaiplatformragcorpuss
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformRagCorpus is the Schema for the AIPlatformRagCorpus API
// +k8s:openapi-gen=true
type AIPlatformRagCorpus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformRagCorpusSpec   `json:"spec,omitempty"`
	Status AIPlatformRagCorpusStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformRagCorpusList contains a list of AIPlatformRagCorpus
type AIPlatformRagCorpusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformRagCorpus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformRagCorpus{}, &AIPlatformRagCorpusList{})
}
