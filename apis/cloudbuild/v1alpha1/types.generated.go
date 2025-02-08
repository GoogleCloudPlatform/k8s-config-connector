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


// +kcc:proto=google.devtools.cloudbuild.v2.BitbucketCloudConfig
type BitbucketCloudConfig struct {
	// Required. The Bitbucket Cloud Workspace ID to be connected to Google Cloud
	//  Platform.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.workspace
	Workspace *string `json:"workspace,omitempty"`

	// Required. SecretManager resource containing the webhook secret used to
	//  verify webhook events, formatted as `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Required. An access token with the `repository` access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate the credentials.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. An access token with the `webhook`, `repository`,
	//  `repository:admin` and `pullrequest` scope access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate these credentials.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig
type BitbucketDataCenterConfig struct {
	// Required. The URI of the Bitbucket Data Center instance or cluster this
	//  connection is for.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. Immutable. SecretManager resource containing the webhook secret
	//  used to verify webhook events, formatted as
	//  `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Required. A http access token with the `REPO_READ` access.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A http access token with the `REPO_ADMIN` scope access.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`

	// Optional. Configuration for using Service Directory to privately connect to
	//  a Bitbucket Data Center. This should only be set if the Bitbucket Data
	//  Center is hosted on-premises and not reachable by public internet. If this
	//  field is left empty, calls to the Bitbucket Data Center will be made over
	//  the public internet.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// Optional. SSL certificate to use for requests to the Bitbucket Data Center.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.ssl_ca
	SslCa *string `json:"sslCa,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.Connection
type Connection struct {
	// Immutable. The resource name of the connection, in the format
	//  `projects/{project}/locations/{location}/connections/{connection_id}`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.name
	Name *string `json:"name,omitempty"`

	// Configuration for connections to github.com.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.github_config
	GithubConfig *GitHubConfig `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.github_enterprise_config
	GithubEnterpriseConfig *GitHubEnterpriseConfig `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com or an instance of GitLab
	//  Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.gitlab_config
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty"`

	// Configuration for connections to Bitbucket Data Center.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.bitbucket_data_center_config
	BitbucketDataCenterConfig *BitbucketDataCenterConfig `json:"bitbucketDataCenterConfig,omitempty"`

	// Configuration for connections to Bitbucket Cloud.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.bitbucket_cloud_config
	BitbucketCloudConfig *BitbucketCloudConfig `json:"bitbucketCloudConfig,omitempty"`

	// If disabled is set to true, functionality is disabled for this connection.
	//  Repository based API methods and webhooks processing for repositories in
	//  this connection will be disabled.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Allows clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitHubConfig
type GitHubConfig struct {
	// OAuth credential of the account that authorized the Cloud Build GitHub App.
	//  It is recommended to use a robot account instead of a human user account.
	//  The OAuth token must be tied to the Cloud Build GitHub App.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubConfig.authorizer_credential
	AuthorizerCredential *OAuthCredential `json:"authorizerCredential,omitempty"`

	// GitHub App installation id.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubConfig.app_installation_id
	AppInstallationID *int64 `json:"appInstallationID,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig
type GitHubEnterpriseConfig struct {
	// Required. The URI of the GitHub Enterprise host this connection is for.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. API Key used for authentication of webhook events.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.api_key
	ApiKey *string `json:"apiKey,omitempty"`

	// Id of the GitHub App created from the manifest.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_id
	AppID *int64 `json:"appID,omitempty"`

	// The URL-friendly name of the GitHub App.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_slug
	AppSlug *string `json:"appSlug,omitempty"`

	// SecretManager resource containing the private key of the GitHub App,
	//  formatted as `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.private_key_secret_version
	PrivateKeySecretVersion *string `json:"privateKeySecretVersion,omitempty"`

	// SecretManager resource containing the webhook secret of the GitHub App,
	//  formatted as `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// ID of the installation of the GitHub App.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_installation_id
	AppInstallationID *int64 `json:"appInstallationID,omitempty"`

	// Configuration for using Service Directory to privately connect to a GitHub
	//  Enterprise server. This should only be set if the GitHub Enterprise server
	//  is hosted on-premises and not reachable by public internet. If this field
	//  is left empty, calls to the GitHub Enterprise server will be made over the
	//  public internet.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitHub Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.ssl_ca
	SslCa *string `json:"sslCa,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitLabConfig
type GitLabConfig struct {
	// The URI of the GitLab Enterprise host this connection is for.
	//  If not specified, the default value is https://gitlab.com.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. Immutable. SecretManager resource containing the webhook secret
	//  of a GitLab Enterprise project, formatted as
	//  `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.webhook_secret_secret_version
	WebhookSecretSecretVersion *string `json:"webhookSecretSecretVersion,omitempty"`

	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A GitLab personal access token with the `api` scope access.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`

	// Configuration for using Service Directory to privately connect to a GitLab
	//  Enterprise server. This should only be set if the GitLab Enterprise server
	//  is hosted on-premises and not reachable by public internet. If this field
	//  is left empty, calls to the GitLab Enterprise server will be made over the
	//  public internet.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitLab Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.ssl_ca
	SslCa *string `json:"sslCa,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.InstallationState
type InstallationState struct {
}

// +kcc:proto=google.devtools.cloudbuild.v2.OAuthCredential
type OAuthCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes
	//  the Cloud Build connection. Format: `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.OAuthCredential.oauth_token_secret_version
	OauthTokenSecretVersion *string `json:"oauthTokenSecretVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// Required. The Service Directory service name.
	//  Format:
	//  projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.ServiceDirectoryConfig.service
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.UserCredential
type UserCredential struct {
	// Required. A SecretManager resource containing the user token that
	//  authorizes the Cloud Build connection. Format:
	//  `projects/*/secrets/*/versions/*`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.UserCredential.user_token_secret_version
	UserTokenSecretVersion *string `json:"userTokenSecretVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig
type BitbucketDataCenterConfigObservedState struct {
	// Output only. Version of the Bitbucket Data Center running on the
	//  `host_uri`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.Connection
type ConnectionObservedState struct {
	// Output only. Server assigned timestamp for when the connection was created.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Server assigned timestamp for when the connection was updated.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Configuration for connections to github.com.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.github_config
	GithubConfig *GitHubConfigObservedState `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.github_enterprise_config
	GithubEnterpriseConfig *GitHubEnterpriseConfigObservedState `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com or an instance of GitLab
	//  Enterprise.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.gitlab_config
	GitlabConfig *GitLabConfigObservedState `json:"gitlabConfig,omitempty"`

	// Configuration for connections to Bitbucket Data Center.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.bitbucket_data_center_config
	BitbucketDataCenterConfig *BitbucketDataCenterConfigObservedState `json:"bitbucketDataCenterConfig,omitempty"`

	// Output only. Installation state of the Connection.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.installation_state
	InstallationState *InstallationState `json:"installationState,omitempty"`

	// Output only. Set to true when the connection is being set up or updated in
	//  the background.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.Connection.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitHubConfig
type GitHubConfigObservedState struct {
	// OAuth credential of the account that authorized the Cloud Build GitHub App.
	//  It is recommended to use a robot account instead of a human user account.
	//  The OAuth token must be tied to the Cloud Build GitHub App.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubConfig.authorizer_credential
	AuthorizerCredential *OAuthCredentialObservedState `json:"authorizerCredential,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig
type GitHubEnterpriseConfigObservedState struct {
	// Output only. GitHub Enterprise version installed at the host_uri.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitLabConfig
type GitLabConfigObservedState struct {
	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredentialObservedState `json:"readAuthorizerCredential,omitempty"`

	// Output only. Version of the GitLab Enterprise server running on the
	//  `host_uri`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.InstallationState
type InstallationStateObservedState struct {
	// Output only. Current step of the installation process.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.InstallationState.stage
	Stage *string `json:"stage,omitempty"`

	// Output only. Message of what the user should do next to continue the
	//  installation. Empty string if the installation is already complete.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.InstallationState.message
	Message *string `json:"message,omitempty"`

	// Output only. Link to follow for next action. Empty string if the
	//  installation is already complete.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.InstallationState.action_uri
	ActionURI *string `json:"actionURI,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.OAuthCredential
type OAuthCredentialObservedState struct {
	// Output only. The username associated to this token.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.OAuthCredential.username
	Username *string `json:"username,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.UserCredential
type UserCredentialObservedState struct {
	// Output only. The username associated to this token.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.UserCredential.username
	Username *string `json:"username,omitempty"`
}
