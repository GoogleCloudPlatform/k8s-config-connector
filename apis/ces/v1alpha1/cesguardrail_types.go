// Copyright 2025 Google LLC
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

// API sources for CESGuardrail, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/guardrail.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.guardrails

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESGuardrailGVK = GroupVersion.WithKind("CESGuardrail")

// CESGuardrailSpec defines the desired state of CESGuardrail
// +kcc:spec:proto=google.cloud.ces.v1beta.Guardrail
// +kcc:required-from-proto
type CESGuardrailSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESGuardrail name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Guardrail that bans certain content from being used in the
	//  conversation.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.content_filter
	ContentFilter *Guardrail_ContentFilter `json:"contentFilter,omitempty"`

	// Optional. Guardrail that blocks the conversation if the prompt is
	//  considered unsafe based on the LLM classification.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.llm_prompt_security
	LlmPromptSecurity *Guardrail_LlmPromptSecurity `json:"llmPromptSecurity,omitempty"`

	// Optional. Guardrail that blocks the conversation if the LLM response is
	//  considered violating the policy based on the LLM classification.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.llm_policy
	LlmPolicy *Guardrail_LlmPolicy `json:"llmPolicy,omitempty"`

	// Optional. Guardrail that blocks the conversation if the LLM response is
	//  considered unsafe based on the model safety settings.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.model_safety
	ModelSafety *Guardrail_ModelSafety `json:"modelSafety,omitempty"`

	// Optional. Guardrail that potentially blocks the conversation based on the
	//  result of the callback execution.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.code_callback
	CodeCallback *Guardrail_CodeCallback `json:"codeCallback,omitempty"`

	// Required. Display name of the guardrail.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the guardrail.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.description
	Description *string `json:"description,omitempty"`

	// Optional. Whether the guardrail is enabled.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. Action to take when the guardrail is triggered.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.action
	Action *TriggerAction `json:"action,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.etag
	Etag *string `json:"etag,omitempty"`
}

// CESGuardrailStatus defines the config connector machine state of CESGuardrail
type CESGuardrailStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESGuardrail resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESGuardrailObservedState `json:"observedState,omitempty"`
}

// CESGuardrailObservedState is the state of the CESGuardrail resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.Guardrail
type CESGuardrailObservedState struct {
	// Optional. Guardrail that blocks the conversation if the prompt is
	//  considered unsafe based on the LLM classification.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.llm_prompt_security
	LlmPromptSecurity *Guardrail_LlmPromptSecurityObservedState `json:"llmPromptSecurity,omitempty"`

	// Output only. Timestamp when the guardrail was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the guardrail was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.Guardrail.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesguardrail;gcpcesguardrails
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESGuardrail is the Schema for the CESGuardrail API
// +k8s:openapi-gen=true
type CESGuardrail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESGuardrailSpec   `json:"spec,omitempty"`
	Status CESGuardrailStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESGuardrailList contains a list of CESGuardrail
type CESGuardrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESGuardrail `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESGuardrail{}, &CESGuardrailList{})
}
