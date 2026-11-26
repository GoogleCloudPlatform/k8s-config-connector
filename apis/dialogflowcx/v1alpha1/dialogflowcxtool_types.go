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

// API sources for DialogflowCXTool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/cx/v3/tool.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.tools

package v1alpha1

import (
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowCXToolGVK = GroupVersion.WithKind("DialogflowCXTool")

// DialogflowCXToolSpec defines the desired state of DialogflowCXTool
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3.Tool
// +kcc:required-from-proto
type DialogflowCXToolSpec struct {
	// The DialogflowCXAgent that this DialogflowCXTool belongs to.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DialogflowCXAgentRef `json:"parentRef"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the Tool and its usage.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.description
	// +required
	Description *string `json:"description,omitempty"`

	// OpenAPI specification of the Tool.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.open_api_spec
	OpenAPISpec *Tool_OpenAPITool `json:"openAPISpec,omitempty"`

	// Data store search tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.data_store_spec
	DataStoreSpec *Tool_DataStoreTool `json:"dataStoreSpec,omitempty"`

	// Client side executed function specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.function_spec
	FunctionSpec *Tool_FunctionTool `json:"functionSpec,omitempty"`
}

// DialogflowCXToolStatus defines the config connector machine state of DialogflowCXTool
type DialogflowCXToolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowCXTool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowCXToolObservedState `json:"observedState,omitempty"`
}

// DialogflowCXToolObservedState is the state of the DialogflowCXTool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.Tool
type DialogflowCXToolObservedState struct {
	// Output only. The tool type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.tool_type
	ToolType *string `json:"toolType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowcxtool;gcpdialogflowcxtools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowCXTool is the Schema for the DialogflowCXTool API
// +k8s:openapi-gen=true
type DialogflowCXTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowCXToolSpec   `json:"spec,omitempty"`
	Status DialogflowCXToolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowCXToolList contains a list of DialogflowCXTool
type DialogflowCXToolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowCXTool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowCXTool{}, &DialogflowCXToolList{})
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	//  Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The document processing mode for the data store connection. Should only be
	//  set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is
	//  considered as DOCUMENTS, as this is the legacy mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.DataStoreConnection.document_processing_mode
	DocumentProcessingMode *string `json:"documentProcessingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Tool.Authentication.ApiKeyConfig
type Tool_Authentication_APIKeyConfig struct {
	// Required. The parameter name or the header name of the API key.
	//  E.g., If the API request is "https://example.com/act?X-Api-Key=<API
	//  KEY>", "X-Api-Key" would be the parameter name.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.ApiKeyConfig.key_name
	// +required
	KeyName *string `json:"keyName,omitempty"`

	// Optional. The API key. If the `secret_version_for_api_key` field is
	//  set, this field will be ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.ApiKeyConfig.api_key
	APIKey *string `json:"apiKey,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the API key. If this field is set, the `api_key` field will be ignored.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.ApiKeyConfig.secret_version_for_api_key
	SecretVersionForAPIKeyRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForAPIKeyRef,omitempty"`

	// Required. Key location in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.ApiKeyConfig.request_location
	// +required
	RequestLocation *string `json:"requestLocation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Tool.Authentication.BearerTokenConfig
type Tool_Authentication_BearerTokenConfig struct {
	// Optional. The text token appended to the text `Bearer` to the request
	//  Authorization header.
	//  [Session parameters
	//  reference](https://cloud.google.com/dialogflow/cx/docs/concept/parameter#session-ref)
	//  can be used to pass the token dynamically, e.g.
	//  `$session.params.parameter-id`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.BearerTokenConfig.token
	Token *string `json:"token,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the Bearer token. If this field is set, the `token` field will be
	//  ignored. Format:
	//  `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.BearerTokenConfig.secret_version_for_token
	SecretVersionForTokenRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForTokenRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig
type Tool_Authentication_OAuthConfig struct {
	// Required. OAuth grant types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.oauth_grant_type
	// +required
	OauthGrantType *string `json:"oauthGrantType,omitempty"`

	// Required. The client ID from the OAuth provider.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.client_id
	// +required
	ClientID *string `json:"clientID,omitempty"`

	// Optional. The client secret from the OAuth provider. If the
	//  `secret_version_for_client_secret` field is set, this field will be
	//  ignored.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Optional. The name of the SecretManager secret version resource storing
	//  the client secret. If this field is set, the `client_secret` field will
	//  be ignored. Format:
	//  `projects/{project}/secrets/{secret}/versions/{version}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.secret_version_for_client_secret
	SecretVersionForClientSecretRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionForClientSecretRef,omitempty"`

	// Required. The token endpoint in the OAuth provider to exchange for an
	//  access token.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.token_endpoint
	// +required
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Tool.Authentication.ServiceAccountAuthConfig
type Tool_Authentication_ServiceAccountAuthConfig struct {
	// Required. The email address of the service account used to authenticate
	//  the tool call. Dialogflow uses this service account to exchange an
	//  access token and the access token is then sent in the `Authorization`
	//  header of the tool request.
	//
	//  The service account must have the
	//  `roles/iam.serviceAccountTokenCreator` role granted to the
	//  [Dialogflow service
	//  agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.Authentication.ServiceAccountAuthConfig.service_account
	// +required
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3.Tool.ServiceDirectoryConfig
type Tool_ServiceDirectoryConfig struct {
	// Required. The name of [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>`.
	//  `LocationID` of the service directory must be the same as the location
	//  of the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3.Tool.ServiceDirectoryConfig.service
	// +required
	ServiceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}
