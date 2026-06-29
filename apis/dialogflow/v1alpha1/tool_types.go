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
	vertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowToolGVK = GroupVersion.WithKind("DialogflowTool")

// DialogflowToolSpec defines the desired state of DialogflowTool
// +kcc:spec:proto=google.cloud.dialogflow.v2.Tool
type DialogflowToolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DialogflowTool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +required
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. High level description of the Tool and its usage.
	// +optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.description
	Description *string `json:"description,omitempty"`

	// OpenAPI specification of the Tool.
	// +optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.open_api_spec
	OpenAPISpec *Tool_OpenAPITool `json:"openAPISpec,omitempty"`

	// Vertex extension tool specification.
	// +optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.extension_spec
	ExtensionSpec *Tool_ExtensionTool `json:"extensionSpec,omitempty"`

	// Client side executed function specification.
	// +optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.function_spec
	FunctionSpec *Tool_FunctionTool `json:"functionSpec,omitempty"`

	// Integration connectors tool specification.
	// +optional
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.connector_spec
	ConnectorSpec *Tool_ConnectorTool `json:"connectorSpec,omitempty"`
}

// DialogflowToolStatus defines the config connector machine state of DialogflowTool
type DialogflowToolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowTool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowToolObservedState `json:"observedState,omitempty"`
}

// DialogflowToolObservedState is the state of the DialogflowTool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Tool
type DialogflowToolObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowtool;gcpdialogflowtools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowTool is the Schema for the DialogflowTool API
// +k8s:openapi-gen=true
type DialogflowTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowToolSpec   `json:"spec,omitempty"`
	Status DialogflowToolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowToolList contains a list of DialogflowTool
type DialogflowToolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowTool `json:"items"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Tool.ExtensionTool
type Tool_ExtensionTool struct {
	// Required. The full name of the referenced vertex extension.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.ExtensionTool.name
	ExtensionRef *vertexaiv1alpha1.VertexAIExtensionRef `json:"extensionRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Tool.Authentication.ApiKeyConfig
type Tool_Authentication_APIKeyConfig struct {
	// Required. The parameter name or the header name of the API key.
	//  E.g., If the API request is "https://example.com/act?X-Api-Key=<API
	//  KEY>", "X-Api-Key" would be the parameter name.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.ApiKeyConfig.key_name
	KeyName *string `json:"keyName,omitempty"`

	// Optional. The API key. If the `secret_version_for_api_key` field is
	//  set, this field will be ignored.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.ApiKeyConfig.api_key
	APIKey *string `json:"apiKey,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the API key. If this field is set, the `api_key` field will be ignored.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.ApiKeyConfig.secret_version_for_api_key
	SecretVersionForAPIKeyRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForAPIKeyRef,omitempty"`

	// Required. Key location in the request.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.ApiKeyConfig.request_location
	RequestLocation *string `json:"requestLocation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Tool.Authentication.BearerTokenConfig
type Tool_Authentication_BearerTokenConfig struct {
	// Optional. The text token appended to the text `Bearer` to the request
	//  Authorization header.
	//  [Session parameters
	//  reference](https://cloud.google.com/dialogflow/cx/docs/concept/parameter#session-ref)
	//  can be used to pass the token dynamically, e.g.
	//  `$session.params.parameter-id`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.BearerTokenConfig.token
	Token *string `json:"token,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the Bearer token. If this field is set, the `token` field will be
	//  ignored. Format:
	//  `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.BearerTokenConfig.secret_version_for_token
	SecretVersionForTokenRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForTokenRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig
type Tool_Authentication_OAuthConfig struct {
	// Required. OAuth grant types.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.oauth_grant_type
	OauthGrantType *string `json:"oauthGrantType,omitempty"`

	// Required. The client ID from the OAuth provider.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Optional. The client secret from the OAuth provider. If the
	//  `secret_version_for_client_secret` field is set, this field will be
	//  ignored.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the client secret. If this field is set, the `client_secret` field will
	//  be ignored. Format:
	//  `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.secret_version_for_client_secret
	SecretVersionForClientSecretRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForClientSecretRef,omitempty"`

	// Required. The token endpoint in the OAuth provider to exchange for an
	//  access token.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.token_endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Tool.Authentication.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DialogflowTool{}, &DialogflowToolList{})
}
