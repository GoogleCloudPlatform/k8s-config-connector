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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowAgentGVK = GroupVersion.WithKind("DialogflowAgent")

// +kcc:proto=google.cloud.dialogflow.v2.Agent
type DialogflowAgentSpec struct {
	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API. Possible values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"].
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.api_version
	APIVersion *string `json:"apiVersion,omitempty"`

	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.avatar_uri
	AvatarURI *string `json:"avatarUri,omitempty"`

	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.classification_threshold
	ClassificationThreshold *float32 `json:"classificationThreshold,omitempty"`

	// Immutable. The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.default_language_code
	// +required
	DefaultLanguageCode string `json:"defaultLanguageCode"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.description
	Description *string `json:"description,omitempty"`

	// The name of this agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.display_name
	// +required
	DisplayName string `json:"displayName"`

	// Determines whether this agent should log conversation queries.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.enable_logging
	EnableLogging *bool `json:"enableLogging,omitempty"`

	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	// syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	// using @sys.any or very large developer entities. Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"].
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.match_mode
	MatchMode *string `json:"matchMode,omitempty"`

	// Immutable. Optional. The project of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.supported_language_codes
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.tier
	Tier *string `json:"tier,omitempty"`

	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Agent.time_zone
	// +required
	TimeZone string `json:"timeZone"`
}

type DialogflowAgentStatus struct {
	// Conditions represent the latest available observations of the
	// DialogflowAgent's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowAgent resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar,
	// the [avatarUri] field can be used.
	AvatarUriBackend *string `json:"avatarUriBackend,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowagent;gcpdialogflowagents
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowAgent is the Schema for the DialogflowAgent API
// +k8s:openapi-gen=true
type DialogflowAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowAgentSpec   `json:"spec,omitempty"`
	Status DialogflowAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DialogflowAgentList contains a list of DialogflowAgent
type DialogflowAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowAgent{}, &DialogflowAgentList{})
}
