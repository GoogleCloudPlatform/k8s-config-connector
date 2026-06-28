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
	servicedirectoryv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBuildConnectionGVK = GroupVersion.WithKind("CloudBuildConnection")

// CloudBuildConnectionSpec defines the desired state of CloudBuildConnection
// +kcc:spec:proto=google.devtools.cloudbuild.v2.Connection
type CloudBuildConnectionSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The CloudBuildConnection name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Configuration for connections to github.com.
	// +kubebuilder:validation:Optional
	GithubConfig *GitHubConfig `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	// +kubebuilder:validation:Optional
	GithubEnterpriseConfig *GitHubEnterpriseConfig `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com or an instance of GitLab
	//  Enterprise.
	// +kubebuilder:validation:Optional
	GitlabConfig *GitLabConfig `json:"gitlabConfig,omitempty"`

	// Configuration for connections to Bitbucket Data Center.
	// +kubebuilder:validation:Optional
	BitbucketDataCenterConfig *BitbucketDataCenterConfig `json:"bitbucketDataCenterConfig,omitempty"`

	// Configuration for connections to Bitbucket Cloud.
	// +kubebuilder:validation:Optional
	BitbucketCloudConfig *BitbucketCloudConfig `json:"bitbucketCloudConfig,omitempty"`

	// If disabled is set to true, functionality is disabled for this connection.
	//  Repository based API methods and webhooks processing for repositories in
	//  this connection will be disabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty"`

	// Allows clients to store small amounts of arbitrary data.
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// CloudBuildConnectionStatus defines the config connector machine state of CloudBuildConnection
type CloudBuildConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudBuildConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBuildConnectionObservedState `json:"observedState,omitempty"`
}

// CloudBuildConnectionObservedState is the state of the CloudBuildConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.Connection
type CloudBuildConnectionObservedState struct {
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Configuration for connections to github.com.
	GithubConfig *GitHubConfigObservedState `json:"githubConfig,omitempty"`

	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig *GitHubEnterpriseConfigObservedState `json:"githubEnterpriseConfig,omitempty"`

	// Configuration for connections to gitlab.com or an instance of GitLab
	//  Enterprise.
	GitlabConfig *GitLabConfigObservedState `json:"gitlabConfig,omitempty"`

	// Configuration for connections to Bitbucket Data Center.
	BitbucketDataCenterConfig *BitbucketDataCenterConfigObservedState `json:"bitbucketDataCenterConfig,omitempty"`

	// Configuration for connections to Bitbucket Cloud.
	BitbucketCloudConfig *BitbucketCloudConfigObservedState `json:"bitbucketCloudConfig,omitempty"`

	// Output only. Installation state of the Connection.
	InstallationState *InstallationStateObservedState `json:"installationState,omitempty"`

	// Output only. Set to true when the connection is being set up or updated in
	//  the background.
	Reconciling *bool `json:"reconciling,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbuildconnection;gcpcloudbuildconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudBuildConnection is the Schema for the CloudBuildConnection API
// +k8s:openapi-gen=true
type CloudBuildConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudBuildConnectionSpec   `json:"spec,omitempty"`
	Status CloudBuildConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBuildConnectionList contains a list of CloudBuildConnection
type CloudBuildConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildConnection{}, &CloudBuildConnectionList{})
}

// +kcc:proto=google.devtools.cloudbuild.v2.BitbucketCloudConfig
type BitbucketCloudConfig struct {
	// Required. The Bitbucket Cloud Workspace ID to be connected to Google Cloud
	//  Platform.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.workspace
	Workspace *string `json:"workspace,omitempty"`

	// Required. SecretManager resource containing the webhook secret used to
	//  verify webhook events, formatted as `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.webhook_secret_secret_version
	WebhookSecretSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"webhookSecretSecretVersionRef,omitempty"`

	// Required. An access token with the `repository` access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate the credentials.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. An access token with the `webhook`, `repository`,
	//  `repository:admin` and `pullrequest` scope access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate these credentials.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig
type BitbucketDataCenterConfig struct {
	// Required. The URI of the Bitbucket Data Center instance or cluster this
	//  connection is for.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. Immutable. SecretManager resource containing the webhook secret
	//  used to verify webhook events, formatted as
	//  `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.webhook_secret_secret_version
	WebhookSecretSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"webhookSecretSecretVersionRef,omitempty"`

	// Required. A http access token with the `REPO_READ` access.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A http access token with the `REPO_ADMIN` scope access.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`

	// Optional. Configuration for using Service Directory to privately connect to
	//  a Bitbucket Data Center. This should only be set if the Bitbucket Data
	//  Center is hosted on-premises and not reachable by public internet. If this
	//  field is left empty, calls to the Bitbucket Data Center will be made over
	//  the public internet.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// Optional. SSL certificate to use for requests to the Bitbucket Data Center.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.ssl_ca
	SSLCA *string `json:"sslCA,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig
type GitHubEnterpriseConfig struct {
	// Required. The URI of the GitHub Enterprise host this connection is for.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. API Key used for authentication of webhook events.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.api_key
	APIKey *string `json:"apiKey,omitempty"`

	// Id of the GitHub App created from the manifest.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_id
	AppID *int64 `json:"appID,omitempty"`

	// The URL-friendly name of the GitHub App.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_slug
	AppSlug *string `json:"appSlug,omitempty"`

	// SecretManager resource containing the private key of the GitHub App,
	//  formatted as `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.private_key_secret_version
	PrivateKeySecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"privateKeySecretVersionRef,omitempty"`

	// SecretManager resource containing the webhook secret of the GitHub App,
	//  formatted as `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.webhook_secret_secret_version
	WebhookSecretSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"webhookSecretSecretVersionRef,omitempty"`

	// ID of the installation of the GitHub App.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.app_installation_id
	AppInstallationID *int64 `json:"appInstallationID,omitempty"`

	// Configuration for using Service Directory to privately connect to a GitHub
	//  Enterprise server. This should only be set if the GitHub Enterprise server
	//  is hosted on-premises and not reachable by public internet. If this field
	//  is left empty, calls to the GitHub Enterprise server will be made over the
	//  public internet.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitHub Enterprise.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.ssl_ca
	SSLCA *string `json:"sslCA,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.GitLabConfig
type GitLabConfig struct {
	// The URI of the GitLab Enterprise host this connection is for.
	//  If not specified, the default value is https://gitlab.com.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.host_uri
	HostURI *string `json:"hostURI,omitempty"`

	// Required. Immutable. SecretManager resource containing the webhook secret
	//  of a GitLab Enterprise project, formatted as
	//  `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.webhook_secret_secret_version
	WebhookSecretSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"webhookSecretSecretVersionRef,omitempty"`

	// Required. A GitLab personal access token with the minimum `read_api` scope
	//  access.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredential `json:"readAuthorizerCredential,omitempty"`

	// Required. A GitLab personal access token with the `api` scope access.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.authorizer_credential
	AuthorizerCredential *UserCredential `json:"authorizerCredential,omitempty"`

	// Configuration for using Service Directory to privately connect to a GitLab
	//  Enterprise server. This should only be set if the GitLab Enterprise server
	//  is hosted on-premises and not reachable by public internet. If this field
	//  is left empty, calls to the GitLab Enterprise server will be made over the
	//  public internet.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.service_directory_config
	ServiceDirectoryConfig *ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`

	// SSL certificate to use for requests to GitLab Enterprise.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitLabConfig.ssl_ca
	SSLCA *string `json:"sslCA,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.OAuthCredential
type OAuthCredential struct {
	// A SecretManager resource containing the OAuth token that authorizes
	//  the Cloud Build connection. Format: `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.devtools.cloudbuild.v2.OAuthCredential.oauth_token_secret_version
	OauthTokenSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"oauthTokenSecretVersionRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// Required. The Service Directory service name.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.ServiceDirectoryConfig.service
	ServiceRef *servicedirectoryv1alpha1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}

// +kcc:proto=google.devtools.cloudbuild.v2.UserCredential
type UserCredential struct {
	// Required. A SecretManager resource containing the user token that
	//  authorizes the Cloud Build connection. Format:
	//  `projects/*/secrets/*/versions/*`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.devtools.cloudbuild.v2.UserCredential.user_token_secret_version
	UserTokenSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"userTokenSecretVersionRef,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.OAuthCredential
type OAuthCredentialObservedState struct {
	// Output only. The username associated to this token.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.OAuthCredential.username
	Username *string `json:"username,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.UserCredential
type UserCredentialObservedState struct {
	// Output only. The username associated to this token.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.UserCredential.username
	Username *string `json:"username,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig
type BitbucketDataCenterConfigObservedState struct {
	// Output only. Version of the Bitbucket Data Center running on the
	//  `host_uri`.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketDataCenterConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig
type GitHubEnterpriseConfigObservedState struct {
	// Output only. GitHub Enterprise version installed at the host_uri.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.GitHubEnterpriseConfig.server_version
	ServerVersion *string `json:"serverVersion,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.GitLabConfig
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

// +kcc:observedstate:proto=google.devtools.cloudbuild.v2.BitbucketCloudConfig
type BitbucketCloudConfigObservedState struct {
	// Required. An access token with the `repository` access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate the credentials.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.read_authorizer_credential
	ReadAuthorizerCredential *UserCredentialObservedState `json:"readAuthorizerCredential,omitempty"`

	// Required. An access token with the `webhook`, `repository`,
	//  `repository:admin` and `pullrequest` scope access. It can be either a
	//  workspace, project or repository access token. It's recommended to use a
	//  system account to generate these credentials.
	// +kcc:proto:field=google.devtools.cloudbuild.v2.BitbucketCloudConfig.authorizer_credential
	AuthorizerCredential *UserCredentialObservedState `json:"authorizerCredential,omitempty"`
}
