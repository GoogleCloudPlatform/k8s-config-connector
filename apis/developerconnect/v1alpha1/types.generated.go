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


// +kcc:proto=google.cloud.developerconnect.v1.Connection
type Connection struct {
	// Configuration for connections to github.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_config
	GithubConfig *GitHubConfig `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_enterprise_config
	GithubEnterpriseConfig *GitHubEnterpriseConfig `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_config
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty"`

	// Configuration for connections to an instance of GitLab Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_enterprise_config
	GitlabEnterpriseConfig *GitLabEnterpriseConfig `json:"gitlabEnterpriseConfig,omitempty"`

	// Identifier. The resource name of the connection, in the format
	//  `projects/{project}/locations/{location}/connections/{connection_id}`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.name
	Name *string `json:"name,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. If disabled is set to true, functionality is disabled for this
	//  connection. Repository based API methods and webhooks processing for
	//  repositories in this connection will be disabled.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Optional. Allows clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The crypto key configuration. This field is used by the
	//  Customer-Managed Encryption Keys (CMEK) feature.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.crypto_key_config
	CryptoKeyConfig *CryptoKeyConfig `json:"cryptoKeyConfig,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.CryptoKeyConfig
type CryptoKeyConfig struct {
	// Required. The name of the key which is used to encrypt/decrypt customer
	//  data. For key in Cloud KMS, the key should be in the format of
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.CryptoKeyConfig.key_reference
	KeyReference *string `json:"keyReference,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitHubConfig
type GitHubConfig struct {
	// Required. Immutable. The GitHub Application that was installed to the
	//  GitHub user or organization.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubConfig.github_app
	GithubApp *string `json:"githubApp,omitempty"`

	// Optional. OAuth credential of the account that authorized the GitHub App.
	//  It is recommended to use a robot account instead of a human user account.
	//  The OAuth token must be tied to the GitHub App of this config.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubConfig.authorizer_credential
	AuthorizerCredential *OAuthCredential `json:"authorizerCredential,omitempty"`

	// Optional. GitHub App installation id.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubConfig.app_installation_id
	AppInstallationID *int64 `json:"appInstallationID,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitHubEnterpriseConfig
type GitHubEnterpriseConfig struct {
	// Required. The URI of the GitHub Enterprise host this connection is for.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Optional. ID of the GitHub App created from the manifest.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.app_id
	AppID *int64 `json:"appID,omitempty"`

	// Optional. SecretManager resource containing the private key of the GitHub
	//  App, formatted as `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.private_key_secret_version
	PrivateKeySecretVersion *string `json:"privateKeySecretVersion,omitempty"`

	// Optional. SecretManager resource containing the webhook secret of the
	//  GitHub App, formatted as `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Optional. ID of the installation of the GitHub App.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.app_installation_id
	AppInstallationID *int64 `json:"appInstallationID,omitempty"`

	// Optional. Configuration for using Service Directory to privately connect to
	//  a GitHub Enterprise server. This should only be set if the GitHub
	//  Enterprise server is hosted on-premises and not reachable by public
	//  internet. If this field is left empty, calls to the GitHub Enterprise
	//  server will be made over the public internet.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// Optional. SSL certificate to use for requests to GitHub Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.ssl_ca_certificate
	SslCaCertificate *string `json:"sslCaCertificate,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitLabConfig
type GitLabConfig struct {
	// Required. Immutable. SecretManager resource containing the webhook secret
	//  of a GitLab project, formatted as `projects/*/secrets/*/versions/*`. This
	//  is used to validate webhooks.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access and a minimum role of `reporter`. The GitLab Projects visible to
	//  this Personal Access Token will control which Projects Developer Connect
	//  has access to.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A GitLab personal access token with the minimum `api` scope
	//  access and a minimum role of `maintainer`. The GitLab Projects visible to
	//  this Personal Access Token will control which Projects Developer Connect
	//  has access to.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitLabEnterpriseConfig
type GitLabEnterpriseConfig struct {
	// Required. The URI of the GitLab Enterprise host this connection is for.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. Immutable. SecretManager resource containing the webhook secret
	//  of a GitLab project, formatted as `projects/*/secrets/*/versions/*`. This
	//  is used to validate webhooks.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access and a minimum role of `reporter`. The GitLab Projects visible to
	//  this Personal Access Token will control which Projects Developer Connect
	//  has access to.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A GitLab personal access token with the minimum `api` scope
	//  access and a minimum role of `maintainer`. The GitLab Projects visible to
	//  this Personal Access Token will control which Projects Developer Connect
	//  has access to.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`

	// Optional. Configuration for using Service Directory to privately connect to
	//  a GitLab Enterprise instance. This should only be set if the GitLab
	//  Enterprise server is hosted on-premises and not reachable by public
	//  internet. If this field is left empty, calls to the GitLab Enterprise
	//  server will be made over the public internet.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// Optional. SSL Certificate Authority certificate to use for requests to
	//  GitLab Enterprise instance.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.ssl_ca_certificate
	SslCaCertificate *string `json:"sslCaCertificate,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.InstallationState
type InstallationState struct {
}

// +kcc:proto=google.cloud.developerconnect.v1.OAuthCredential
type OAuthCredential struct {
	// Required. A SecretManager resource containing the OAuth token that
	//  authorizes the connection. Format: `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.OAuthCredential.oauth_token_secret_version
	OauthTokenSecretVersion *string `json:"oauthTokenSecretVersion,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// Required. The Service Directory service name.
	//  Format:
	//  projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	// +kcc:proto:field=google.cloud.developerconnect.v1.ServiceDirectoryConfig.service
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.UserCredential
type UserCredential struct {
	// Required. A SecretManager resource containing the user token that
	//  authorizes the Developer Connect connection. Format:
	//  `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.UserCredential.user_token_secret_version
	UserTokenSecretVersion *string `json:"userTokenSecretVersion,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.Connection
type ConnectionObservedState struct {
	// Configuration for connections to github.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_config
	GithubConfig *GitHubConfigObservedState `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.github_enterprise_config
	GithubEnterpriseConfig *GitHubEnterpriseConfigObservedState `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_config
	GitlabConfig *GitLabConfigObservedState `json:"gitlabConfig,omitempty"`

	// Configuration for connections to an instance of GitLab Enterprise.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.gitlab_enterprise_config
	GitlabEnterpriseConfig *GitLabEnterpriseConfigObservedState `json:"gitlabEnterpriseConfig,omitempty"`

	// Output only. [Output only] Create timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. [Output only] Delete timestamp
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. Installation state of the Connection.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.installation_state
	InstallationState *InstallationState `json:"installationState,omitempty"`

	// Output only. Set to true when the connection is being set up or updated in
	//  the background.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-assigned unique identifier for a the
	//  GitRepositoryLink.
	// +kcc:proto:field=google.cloud.developerconnect.v1.Connection.uid
	Uid *string `json:"uid,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitHubConfig
type GitHubConfigObservedState struct {
	// Optional. OAuth credential of the account that authorized the GitHub App.
	//  It is recommended to use a robot account instead of a human user account.
	//  The OAuth token must be tied to the GitHub App of this config.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubConfig.authorizer_credential
	AuthorizerCredential *OAuthCredentialObservedState `json:"authorizerCredential,omitempty"`

	// Output only. The URI to navigate to in order to manage the installation
	//  associated with this GitHubConfig.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubConfig.installation_uri
	InstallationURI *string `json:"installationURI,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitHubEnterpriseConfig
type GitHubEnterpriseConfigObservedState struct {
	// Output only. The URL-friendly name of the GitHub App.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.app_slug
	AppSlug *string `json:"appSlug,omitempty"`

	// Output only. The URI to navigate to in order to manage the installation
	//  associated with this GitHubEnterpriseConfig.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.installation_uri
	InstallationURI *string `json:"installationURI,omitempty"`

	// Output only. GitHub Enterprise version installed at the host_uri.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitHubEnterpriseConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitLabConfig
type GitLabConfigObservedState struct {
	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access and a minimum role of `reporter`. The GitLab Projects visible to
	//  this Personal Access Token will control which Projects Developer Connect
	//  has access to.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredentialObservedState `json:"readAuthorizerCredential,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.GitLabEnterpriseConfig
type GitLabEnterpriseConfigObservedState struct {
	// Output only. Version of the GitLab Enterprise server running on the
	//  `host_uri`.
	// +kcc:proto:field=google.cloud.developerconnect.v1.GitLabEnterpriseConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.InstallationState
type InstallationStateObservedState struct {
	// Output only. Current step of the installation process.
	// +kcc:proto:field=google.cloud.developerconnect.v1.InstallationState.stage
	Stage *string `json:"stage,omitempty"`

	// Output only. Message of what the user should do next to continue the
	//  installation. Empty string if the installation is already complete.
	// +kcc:proto:field=google.cloud.developerconnect.v1.InstallationState.message
	Message *string `json:"message,omitempty"`

	// Output only. Link to follow for next action. Empty string if the
	//  installation is already complete.
	// +kcc:proto:field=google.cloud.developerconnect.v1.InstallationState.action_uri
	ActionURI *string `json:"actionURI,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.OAuthCredential
type OAuthCredentialObservedState struct {
	// Output only. The username associated with this token.
	// +kcc:proto:field=google.cloud.developerconnect.v1.OAuthCredential.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.cloud.developerconnect.v1.UserCredential
type UserCredentialObservedState struct {
	// Output only. The username associated with this token.
	// +kcc:proto:field=google.cloud.developerconnect.v1.UserCredential.username
	Username *string `json:"username,omitempty"`
}
