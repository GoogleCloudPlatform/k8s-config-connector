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
	agentsearchv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/agentsearch/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIExtensionGVK = GroupVersion.WithKind("VertexAIExtension")

// VertexAIExtensionSpec defines the desired state of VertexAIExtension
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Extension
type VertexAIExtensionSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIExtension name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the Extension.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.display_name
	DisplayName *string `json:"displayName"`

	// Optional. The description of the Extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.description
	Description *string `json:"description,omitempty"`

	// Required. Manifest of the Extension.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.manifest
	Manifest *ExtensionManifest `json:"manifest"`

	// Optional. Runtime config of the Extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.runtime_config
	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Optional. Tool use examples of the Extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.tool_use_examples
	ToolUseExamples []ToolUseExample `json:"toolUseExamples,omitempty"`

	// Optional. The Private Service Connect config of the Extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.private_service_connect_config
	PrivateServiceConnectConfig *ExtensionPrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`
}

// VertexAIExtensionStatus defines the config connector machine state of VertexAIExtension
type VertexAIExtensionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIExtension resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIExtensionObservedState `json:"observedState,omitempty"`
}

// VertexAIExtensionObservedState is the state of the VertexAIExtension resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Extension
type VertexAIExtensionObservedState struct {
	// Output only. Timestamp when this Extension was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Extension was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Supported operations of the Extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.extension_operations
	ExtensionOperations []ExtensionOperationObservedState `json:"extensionOperations,omitempty"`

	// Output only. Used to perform consistent read-modify-write updates.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Extension.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiextension;gcpvertexaiextensions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIExtension is the Schema for the VertexAIExtension API
// +k8s:openapi-gen=true
type VertexAIExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIExtensionSpec   `json:"spec,omitempty"`
	Status VertexAIExtensionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIExtensionList contains a list of VertexAIExtension
type VertexAIExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIExtension{}, &VertexAIExtensionList{})
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.AuthConfig.ApiKeyConfig
type AuthConfig_APIKeyConfig struct {
	// Required. The parameter name of the API key.
	// E.g. If the API request is "https://example.com/act?api_key=<API KEY>",
	// "api_key" would be the parameter name.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.AuthConfig.ApiKeyConfig.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the SecretManager secret version resource storing
	// the API key. Format:
	// `projects/{project}/secrets/{secrete}/versions/{version}`
	//
	// - If specified, the `secretmanager.versions.access` permission should be
	// granted to Vertex AI Extension Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	// on the specified resource.
	APIKeySecretRef *refsv1beta1.SecretManagerSecretVersionRef `json:"apiKeySecretRef,omitempty"`

	// Required. The location of the API key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.AuthConfig.ApiKeyConfig.http_element_location
	HTTPElementLocation *string `json:"httpElementLocation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.AuthConfig.GoogleServiceAccountConfig
type AuthConfig_GoogleServiceAccountConfig struct {
	// Optional. The service account that the extension execution service runs
	// as.
	//
	// - If the service account is specified,
	// the `iam.serviceAccounts.getAccessToken` permission should be granted to
	// Vertex AI Extension Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	// on the specified service account.
	//
	// - If not specified, the Vertex AI Extension Service Agent
	// will be used to execute the Extension.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.AuthConfig.HttpBasicAuthConfig
type AuthConfig_HTTPBasicAuthConfig struct {
	// Required. The name of the SecretManager secret version resource storing
	// the base64 encoded credentials. Format:
	// `projects/{project}/secrets/{secrete}/versions/{version}`
	//
	// - If specified, the `secretmanager.versions.access` permission should be
	// granted to Vertex AI Extension Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	// on the specified resource.
	CredentialSecretRef *refsv1beta1.SecretManagerSecretVersionRef `json:"credentialSecretRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.AuthConfig.OauthConfig
type AuthConfig_OauthConfig struct {
	// Access token for extension endpoint.
	// Only used to propagate token from
	// [[ExecuteExtensionRequest.runtime_auth_config]] at request time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.AuthConfig.OauthConfig.access_token
	AccessToken *string `json:"accessToken,omitempty"`

	// The service account used to generate access tokens for executing the
	// Extension.
	//
	// - If the service account is specified,
	// the `iam.serviceAccounts.getAccessToken` permission should be granted
	// to Vertex AI Extension Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	// on the provided service account.
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.AuthConfig.OidcConfig
type AuthConfig_OIDCConfig struct {
	// OpenID Connect formatted ID token for extension endpoint.
	// Only used to propagate token from
	// [[ExecuteExtensionRequest.runtime_auth_config]] at request time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.AuthConfig.OidcConfig.id_token
	IDToken *string `json:"idToken,omitempty"`

	// The service account used to generate an OpenID Connect
	// (OIDC)-compatible JWT token signed by the Google OIDC Provider
	// (accounts.google.com) for extension endpoint
	// (https://cloud.google.com/iam/docs/create-short-lived-credentials-direct#sa-credentials-oidc).
	//
	// - The audience for the token will be set to the URL in the server url
	// defined in the OpenApi spec.
	//
	// - If the service account is provided, the service account should grant
	// `iam.serviceAccounts.getOpenIdToken` permission to Vertex AI Extension
	// Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents).
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExtensionPrivateServiceConnectConfig
type ExtensionPrivateServiceConnectConfig struct {
	// Required. The Service Directory resource name in which the service
	// endpoints associated to the extension are registered. Format:
	// `projects/{project_id}/locations/{location_id}/namespaces/{namespace_id}/services/{service_id}`
	//
	// - The Vertex AI Extension Service Agent
	// (https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents)
	// should be granted `servicedirectory.viewer` and
	// `servicedirectory.pscAuthorizedService` roles on the resource.
	ServiceDirectoryRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceDirectoryRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FunctionDeclaration
type FunctionDeclaration struct {
	// Required. The name of the function to call.
	// Must start with a letter or an underscore.
	// Must be a-z, A-Z, 0-9, or contain underscores, dots and dashes, with a
	// maximum length of 64.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.name
	Name *string `json:"name,omitempty"`

	// Optional. Description and purpose of the function.
	// Model uses it to decide how and whether to call the function.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	/*
		// Optional. Describes the parameters to the function in JSON Schema format.
		// The schema must describe an object where the properties are the parameters
		// to the function. For example:
		//
		// ```
		// {
		// "type": "object",
		// "properties": {
		// "name": { "type": "string" },
		// "age": { "type": "integer" }
		// },
		// "additionalProperties": false,
		// "required": ["name", "age"],
		// "propertyOrdering": ["name", "age"]
		// }
		// ```
		//
		// This field is mutually exclusive with `parameters`.
		// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.parameters_json_schema
		ParametersJsonSchema *Value `json:"parametersJsonSchema,omitempty"`

		// Optional. Describes the output from this function in JSON Schema format.
		// The value specified by the schema is the response value of the function.
		//
		// This field is mutually exclusive with `response`.
		// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FunctionDeclaration.response_json_schema
		ResponseJsonSchema *Value `json:"responseJsonSchema,omitempty"`
	*/
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RuntimeConfig
type RuntimeConfig struct {
	// Code execution runtime configurations for code interpreter extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeConfig.code_interpreter_runtime_config
	CodeInterpreterRuntimeConfig *RuntimeConfig_CodeInterpreterRuntimeConfig `json:"codeInterpreterRuntimeConfig,omitempty"`

	// Runtime configuration for Vertex AI Search extension.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeConfig.vertex_ai_search_runtime_config
	VertexAISearchRuntimeConfig *RuntimeConfig_VertexAISearchRuntimeConfig `json:"vertexAISearchRuntimeConfig,omitempty"`

	// Optional. Default parameters that will be set for all the execution of this
	// extension. If specified, the parameter values can be overridden by values
	// in [[ExecuteExtensionRequest.operation_params]] at request time.
	//
	// The struct should be in a form of map with param name as the key and actual
	// param value as the value.
	// E.g. If this operation requires a param "name" to be set to "abc". you can
	// set this to something like {"name": "abc"}.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeConfig.default_params
	DefaultParams apiextensionsv1.JSON `json:"defaultParams,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RuntimeConfig.CodeInterpreterRuntimeConfig
type RuntimeConfig_CodeInterpreterRuntimeConfig struct {
	// Optional. The Cloud Storage bucket for file input of this Extension.
	// If specified, support input from the Cloud Storage bucket.
	// Vertex Extension Custom Code Service Agent should be granted
	// file reader to this bucket.
	// If not specified, the extension will only accept file contents from
	// request body and reject Cloud Storage file inputs.
	FileInputGCSBucketRef *storagev1beta1.StorageBucketRef `json:"fileInputGCSBucketRef,omitempty"`

	// Optional. The Cloud Storage bucket for file output of this Extension.
	// If specified, write all output files to the Cloud Storage bucket.
	// Vertex Extension Custom Code Service Agent should be granted
	// file writer to this bucket.
	// If not specified, the file content will be output in response body.
	FileOutputGCSBucketRef *storagev1beta1.StorageBucketRef `json:"fileOutputGCSBucketRef,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RuntimeConfig.VertexAISearchRuntimeConfig
type RuntimeConfig_VertexAISearchRuntimeConfig struct {
	// Optional. Vertex AI Search serving config name. Format:
	// `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}/servingConfigs/{serving_config}`
	ServingConfigRef *agentsearchv1alpha1.AgentSearchServingConfigRef `json:"servingConfigRef,omitempty"`

	// Optional. Vertex AI Search engine ID. This is used to construct the
	// search request. By setting this engine_id, API will construct the serving
	// config using the default value to call search API for the user. The
	// engine_id and serving_config_name cannot both be empty at the same time.
	EngineRef *agentsearchv1alpha1.AgentSearchEngineRef `json:"engineRef,omitempty"`
}
