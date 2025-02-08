// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	//  Formats:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`
	//  `projects/{project}/locations/{location}/dataStores/{data_store}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store
	DataStore *string `json:"dataStore,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool
type Tool struct {
	// The unique identifier of the Tool.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/tools/<ToolID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the Tool and its usage.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.description
	Description *string `json:"description,omitempty"`

	// OpenAPI specification of the Tool.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.open_api_spec
	OpenApiSpec *Tool_OpenApiTool `json:"openApiSpec,omitempty"`

	// Data store search tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.data_store_spec
	DataStoreSpec *Tool_DataStoreTool `json:"dataStoreSpec,omitempty"`

	// Vertex extension tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.extension_spec
	ExtensionSpec *Tool_ExtensionTool `json:"extensionSpec,omitempty"`

	// Client side executed function specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.function_spec
	FunctionSpec *Tool_FunctionTool `json:"functionSpec,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication
type Tool_Authentication struct {
	// Config for API key auth.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.api_key_config
	ApiKeyConfig *Tool_Authentication_ApiKeyConfig `json:"apiKeyConfig,omitempty"`

	// Config for OAuth.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.oauth_config
	OauthConfig *Tool_Authentication_OAuthConfig `json:"oauthConfig,omitempty"`

	// Config for [Diglogflow service
	//  agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent)
	//  auth.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.service_agent_auth_config
	ServiceAgentAuthConfig *Tool_Authentication_ServiceAgentAuthConfig `json:"serviceAgentAuthConfig,omitempty"`

	// Config for bearer token auth.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.bearer_token_config
	BearerTokenConfig *Tool_Authentication_BearerTokenConfig `json:"bearerTokenConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ApiKeyConfig
type Tool_Authentication_ApiKeyConfig struct {
	// Required. The parameter name or the header name of the API key.
	//  E.g., If the API request is "https://example.com/act?X-Api-Key=<API
	//  KEY>", "X-Api-Key" would be the parameter name.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ApiKeyConfig.key_name
	KeyName *string `json:"keyName,omitempty"`

	// Required. The API key.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ApiKeyConfig.api_key
	ApiKey *string `json:"apiKey,omitempty"`

	// Required. Key location in the request.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ApiKeyConfig.request_location
	RequestLocation *string `json:"requestLocation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.BearerTokenConfig
type Tool_Authentication_BearerTokenConfig struct {
	// Required. The text token appended to the text `Bearer` to the request
	//  Authorization header.
	//  [Session parameters
	//  reference](https://cloud.google.com/dialogflow/cx/docs/concept/parameter#session-ref)
	//  can be used to pass the token dynamically, e.g.
	//  `$session.params.parameter-id`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.BearerTokenConfig.token
	Token *string `json:"token,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig
type Tool_Authentication_OAuthConfig struct {
	// Required. OAuth grant types.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig.oauth_grant_type
	OauthGrantType *string `json:"oauthGrantType,omitempty"`

	// Required. The client ID from the OAuth provider.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Required. The client secret from the OAuth provider.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Required. The token endpoint in the OAuth provider to exchange for an
	//  access token.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig.token_endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ServiceAgentAuthConfig
type Tool_Authentication_ServiceAgentAuthConfig struct {
	// Optional. Indicate the auth token type generated from the [Diglogflow
	//  service
	//  agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).
	//  The generated token is sent in the Authorization header.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.Authentication.ServiceAgentAuthConfig.service_agent_auth
	ServiceAgentAuth *string `json:"serviceAgentAuth,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool
type Tool_DataStoreTool struct {
	// Required. List of data stores to search.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.data_store_connections
	DataStoreConnections []DataStoreConnection `json:"dataStoreConnections,omitempty"`

	// Required. Fallback prompt configurations to use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.fallback_prompt
	FallbackPrompt *Tool_DataStoreTool_FallbackPrompt `json:"fallbackPrompt,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.FallbackPrompt
type Tool_DataStoreTool_FallbackPrompt struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool
type Tool_ExtensionTool struct {
	// Required. The full name of the referenced vertex extension.
	//  Formats:
	//  `projects/{project}/locations/{location}/extensions/{extension}`
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.FunctionTool
type Tool_FunctionTool struct {
	// Optional. The JSON schema is encapsulated in a
	//  [google.protobuf.Struct][google.protobuf.Struct] to describe the input of
	//  the function. This input is a JSON object that contains the function's
	//  parameters as properties of the object.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.FunctionTool.input_schema
	InputSchema map[string]string `json:"inputSchema,omitempty"`

	// Optional. The JSON schema is encapsulated in a
	//  [google.protobuf.Struct][google.protobuf.Struct] to describe the output
	//  of the function. This output is a JSON object that contains the
	//  function's parameters as properties of the object.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.FunctionTool.output_schema
	OutputSchema map[string]string `json:"outputSchema,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool
type Tool_OpenApiTool struct {
	// Required. The OpenAPI schema specified as a text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.text_schema
	TextSchema *string `json:"textSchema,omitempty"`

	// Optional. Authentication information required by the API.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.authentication
	Authentication *Tool_Authentication `json:"authentication,omitempty"`

	// Optional. TLS configuration for the HTTPS verification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.tls_config
	TlsConfig *Tool_TLSConfig `json:"tlsConfig,omitempty"`

	// Optional. Service Directory configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.service_directory_config
	ServiceDirectoryConfig *Tool_ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ServiceDirectoryConfig
type Tool_ServiceDirectoryConfig struct {
	// Required. The name of [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>`.
	//  `LocationID` of the service directory must be the same as the location
	//  of the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ServiceDirectoryConfig.service
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.TLSConfig
type Tool_TLSConfig struct {
	// Required. Specifies a list of allowed custom CA certificates for HTTPS
	//  verification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.TLSConfig.ca_certs
	CaCerts []Tool_TLSConfig_CACert `json:"caCerts,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.TLSConfig.CACert
type Tool_TLSConfig_CACert struct {
	// Required. The name of the allowed custom CA certificates. This
	//  can be used to disambiguate the custom CA certificates.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.TLSConfig.CACert.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The allowed custom CA certificates (in DER format) for
	//  HTTPS verification. This overrides the default SSL trust store. If this
	//  is empty or unspecified, Dialogflow will use Google's default trust
	//  store to verify certificates. N.B. Make sure the HTTPS server
	//  certificates are signed with "subject alt name". For instance a
	//  certificate can be self-signed using the following command,
	//     openssl x509 -req -days 200 -in example.com.csr \
	//       -signkey example.com.key \
	//       -out example.com.crt \
	//       -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.TLSConfig.CACert.cert
	Cert []byte `json:"cert,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool
type ToolObservedState struct {
	// Output only. The tool type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.tool_type
	ToolType *string `json:"toolType,omitempty"`
}
