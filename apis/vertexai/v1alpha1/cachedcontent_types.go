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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAICachedContentGVK = GroupVersion.WithKind("VertexAICachedContent")

// VertexAICachedContentSpec defines the desired state of VertexAICachedContent
// +kcc:spec:proto=google.cloud.aiplatform.v1.CachedContent
type VertexAICachedContentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAICachedContent name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Timestamp of when this resource is considered expired.
	// This is *always* provided on output, regardless of what was sent on input.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Input only. The TTL for this resource. The expiration time is computed: now + TTL.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.ttl
	TTL *string `json:"ttl,omitempty"`

	// Optional. Immutable. The user-generated meaningful display name of the cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The name of the `Model` to use for cached content. Currently,
	// only the published Gemini base models are supported, in form of
	// projects/{PROJECT}/locations/{LOCATION}/publishers/google/models/{MODEL}
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.model
	Model *string `json:"model,omitempty"`

	// Optional. Input only. Immutable. Developer set system instruction. Currently, text only
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.system_instruction
	SystemInstruction *Content `json:"systemInstruction,omitempty"`

	// Optional. Input only. Immutable. The content to cache
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.contents
	Contents []Content `json:"contents,omitempty"`

	// Optional. Input only. Immutable. A list of `Tools` the model may use to generate the next response
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.tools
	Tools []Tool `json:"tools,omitempty"`

	// Optional. Input only. Immutable. Tool config. This config is shared for all tools
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.tool_config
	ToolConfig *ToolConfig `json:"toolConfig,omitempty"`

	// Input only. Immutable. Customer-managed encryption key spec for a `CachedContent`.
	// If set, this `CachedContent` and all its sub-resources will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAICachedContentStatus defines the config connector machine state of VertexAICachedContent
type VertexAICachedContentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAICachedContent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAICachedContentObservedState `json:"observedState,omitempty"`
}

// VertexAICachedContentObservedState is the state of the VertexAICachedContent resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.CachedContent
type VertexAICachedContentObservedState struct {
	// Output only. Creation time of the cache entry.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the cache entry was last updated in UTC time.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Metadata on the usage of the cached content.
	// +kcc:proto:field=google.cloud.aiplatform.v1.CachedContent.usage_metadata
	UsageMetadata *CachedContent_UsageMetadata `json:"usageMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaicachedcontent;gcpvertexaicachedcontents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAICachedContent is the Schema for the VertexAICachedContent API
// +k8s:openapi-gen=true
type VertexAICachedContent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAICachedContentSpec   `json:"spec,omitempty"`
	Status VertexAICachedContentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAICachedContentList contains a list of VertexAICachedContent
type VertexAICachedContentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAICachedContent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAICachedContent{}, &VertexAICachedContentList{})
}
