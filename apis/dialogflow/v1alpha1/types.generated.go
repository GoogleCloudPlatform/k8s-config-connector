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


// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook
type Webhook struct {
	// The unique identifier of the webhook.
	//  Required for the
	//  [Webhooks.UpdateWebhook][google.cloud.dialogflow.cx.v3beta1.Webhooks.UpdateWebhook]
	//  method.
	//  [Webhooks.CreateWebhook][google.cloud.dialogflow.cx.v3beta1.Webhooks.CreateWebhook]
	//  populates the name automatically. Format:
	//  `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/webhooks/<WebhookID>`.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.name
	Name *string `json:"name,omitempty"`

	// Required. The human-readable name of the webhook, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration for a generic web service.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.generic_web_service
	GenericWebService *Webhook_GenericWebService `json:"genericWebService,omitempty"`

	// Configuration for a [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.service_directory
	ServiceDirectory *Webhook_ServiceDirectoryConfig `json:"serviceDirectory,omitempty"`

	// Webhook execution timeout. Execution is considered failed if Dialogflow
	//  doesn't receive a response from webhook at the end of the timeout period.
	//  Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Indicates whether the webhook is disabled.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService
type Webhook_GenericWebService struct {
	// Required. The webhook URI for receiving POST requests. It must use https
	//  protocol.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.uri
	URI *string `json:"uri,omitempty"`

	// The user name for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.username
	Username *string `json:"username,omitempty"`

	// The password for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.password
	Password *string `json:"password,omitempty"`

	// The HTTP request headers to send together with webhook
	//  requests.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.request_headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// Optional. Specifies a list of allowed custom CA certificates (in DER
	//  format) for HTTPS verification. This overrides the default SSL trust
	//  store. If this is empty or unspecified, Dialogflow will use Google's
	//  default trust store to verify certificates. N.B. Make sure the HTTPS
	//  server certificates are signed with "subject alt name". For instance a
	//  certificate can be self-signed using the following command,
	//  ```
	//     openssl x509 -req -days 200 -in example.com.csr \
	//       -signkey example.com.key \
	//       -out example.com.crt \
	//       -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.allowed_ca_certs
	AllowedCaCerts [][]byte `json:"allowedCaCerts,omitempty"`

	// Optional. The OAuth configuration of the webhook. If specified,
	//  Dialogflow will initiate the OAuth client credential flow to exchange an
	//  access token from the 3rd party platform and put it in the auth header.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.oauth_config
	OauthConfig *Webhook_GenericWebService_OAuthConfig `json:"oauthConfig,omitempty"`

	// Optional. Indicate the auth token type generated from the [Diglogflow
	//  service
	//  agent](https://cloud.google.com/iam/docs/service-agents#dialogflow-service-agent).
	//  The generated token is sent in the Authorization header.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.service_agent_auth
	ServiceAgentAuth *string `json:"serviceAgentAuth,omitempty"`

	// Optional. Type of the webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.webhook_type
	WebhookType *string `json:"webhookType,omitempty"`

	// Optional. HTTP method for the flexible webhook calls. Standard webhook
	//  always uses POST.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.http_method
	HTTPMethod *string `json:"httpMethod,omitempty"`

	// Optional. Defines a custom JSON object as request body to send to
	//  flexible webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.request_body
	RequestBody *string `json:"requestBody,omitempty"`

	// Optional. Maps the values extracted from specific fields of the flexible
	//  webhook response into session parameters.
	//  - Key: session parameter name
	//  - Value: field path in the webhook response
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.parameter_mapping
	ParameterMapping map[string]string `json:"parameterMapping,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig
type Webhook_GenericWebService_OAuthConfig struct {
	// Required. The client ID provided by the 3rd party platform.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Required. The client secret provided by the 3rd party platform.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.client_secret
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Required. The token endpoint provided by the 3rd party platform to
	//  exchange an access token.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.token_endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.GenericWebService.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig
type Webhook_ServiceDirectoryConfig struct {
	// Required. The name of [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	//  Format:
	//  `projects/<ProjectID>/locations/<LocationID>/namespaces/<NamespaceID>/services/<ServiceID>`.
	//  `Location ID` of the service directory must be the same as the location
	//  of the agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig.service
	Service *string `json:"service,omitempty"`

	// Generic Service configuration of this webhook.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Webhook.ServiceDirectoryConfig.generic_web_service
	GenericWebService *Webhook_GenericWebService `json:"genericWebService,omitempty"`
}
