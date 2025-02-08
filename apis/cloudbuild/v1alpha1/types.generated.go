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


// +kcc:proto=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig
type GitHubEnterpriseConfig struct {
	// Optional. The full resource name for the GitHubEnterpriseConfig
	//  For example:
	//  "projects/{$project_id}/locations/{$location_id}/githubEnterpriseConfigs/{$config_id}"
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.name
	Name *string `json:"name,omitempty"`

	// The URL of the github enterprise host the configuration is for.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.host_url
	HostURL *string `json:"hostURL,omitempty"`

	// Required. The GitHub app id of the Cloud Build app on the GitHub Enterprise
	//  server.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.app_id
	AppID *int64 `json:"appID,omitempty"`

	// The key that should be attached to webhook calls to the ReceiveWebhook
	//  endpoint.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.webhook_key
	WebhookKey *string `json:"webhookKey,omitempty"`

	// Optional. The network to be used when reaching out to the GitHub
	//  Enterprise server. The VPC network must be enabled for private
	//  service connection. This should be set if the GitHub Enterprise server is
	//  hosted on-premises and not reachable by public internet.
	//  If this field is left empty, no network peering will occur and calls to
	//  the GitHub Enterprise server will be made over the public internet.
	//  Must be in the format
	//  `projects/{project}/global/networks/{network}`, where {project}
	//  is a project number or id and {network} is the name of a
	//  VPC network in the project.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.peered_network
	PeeredNetwork *string `json:"peeredNetwork,omitempty"`

	// Names of secrets in Secret Manager.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.secrets
	Secrets *GitHubEnterpriseSecrets `json:"secrets,omitempty"`

	// Name to display for this config.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. SSL certificate to use for requests to GitHub Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.ssl_ca
	SslCa *string `json:"sslCa,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitHubEnterpriseSecrets
type GitHubEnterpriseSecrets struct {
	// The resource name for the private key secret version.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseSecrets.private_key_version_name
	PrivateKeyVersionName *string `json:"privateKeyVersionName,omitempty"`

	// The resource name for the webhook secret secret version in Secret Manager.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseSecrets.webhook_secret_version_name
	WebhookSecretVersionName *string `json:"webhookSecretVersionName,omitempty"`

	// The resource name for the OAuth secret secret version in Secret Manager.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseSecrets.oauth_secret_version_name
	OauthSecretVersionName *string `json:"oauthSecretVersionName,omitempty"`

	// The resource name for the OAuth client ID secret version in Secret Manager.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseSecrets.oauth_client_id_version_name
	OauthClientIDVersionName *string `json:"oauthClientIDVersionName,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig
type GitHubEnterpriseConfigObservedState struct {
	// Output only. Time when the installation was associated with the project.
	// +kcc:proto:field=google.devtools.cloudbuild.v1.GitHubEnterpriseConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
