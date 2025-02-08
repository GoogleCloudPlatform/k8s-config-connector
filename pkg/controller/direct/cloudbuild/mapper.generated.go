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

package cloudbuild

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
)
func BitbucketCloudConfig_FromProto(mapCtx *direct.MapContext, in *pb.BitbucketCloudConfig) *krm.BitbucketCloudConfig {
	if in == nil {
		return nil
	}
	out := &krm.BitbucketCloudConfig{}
	out.Workspace = direct.LazyPtr(in.GetWorkspace())
	out.WebhookSecretSecretVersion = direct.LazyPtr(in.GetWebhookSecretSecretVersion())
	out.ReadAuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetReadAuthorizerCredential())
	out.AuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetAuthorizerCredential())
	return out
}
func BitbucketCloudConfig_ToProto(mapCtx *direct.MapContext, in *krm.BitbucketCloudConfig) *pb.BitbucketCloudConfig {
	if in == nil {
		return nil
	}
	out := &pb.BitbucketCloudConfig{}
	out.Workspace = direct.ValueOf(in.Workspace)
	out.WebhookSecretSecretVersion = direct.ValueOf(in.WebhookSecretSecretVersion)
	out.ReadAuthorizerCredential = UserCredential_ToProto(mapCtx, in.ReadAuthorizerCredential)
	out.AuthorizerCredential = UserCredential_ToProto(mapCtx, in.AuthorizerCredential)
	return out
}
func BitbucketDataCenterConfig_FromProto(mapCtx *direct.MapContext, in *pb.BitbucketDataCenterConfig) *krm.BitbucketDataCenterConfig {
	if in == nil {
		return nil
	}
	out := &krm.BitbucketDataCenterConfig{}
	out.HostURI = direct.LazyPtr(in.GetHostUri())
	out.WebhookSecretSecretVersion = direct.LazyPtr(in.GetWebhookSecretSecretVersion())
	out.ReadAuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetReadAuthorizerCredential())
	out.AuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetAuthorizerCredential())
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectoryConfig())
	out.SslCa = direct.LazyPtr(in.GetSslCa())
	// MISSING: ServerVersion
	return out
}
func BitbucketDataCenterConfig_ToProto(mapCtx *direct.MapContext, in *krm.BitbucketDataCenterConfig) *pb.BitbucketDataCenterConfig {
	if in == nil {
		return nil
	}
	out := &pb.BitbucketDataCenterConfig{}
	out.HostUri = direct.ValueOf(in.HostURI)
	out.WebhookSecretSecretVersion = direct.ValueOf(in.WebhookSecretSecretVersion)
	out.ReadAuthorizerCredential = UserCredential_ToProto(mapCtx, in.ReadAuthorizerCredential)
	out.AuthorizerCredential = UserCredential_ToProto(mapCtx, in.AuthorizerCredential)
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectoryConfig)
	out.SslCa = direct.ValueOf(in.SslCa)
	// MISSING: ServerVersion
	return out
}
func BitbucketDataCenterConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BitbucketDataCenterConfig) *krm.BitbucketDataCenterConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BitbucketDataCenterConfigObservedState{}
	// MISSING: HostURI
	// MISSING: WebhookSecretSecretVersion
	// MISSING: ReadAuthorizerCredential
	// MISSING: AuthorizerCredential
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.LazyPtr(in.GetServerVersion())
	return out
}
func BitbucketDataCenterConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BitbucketDataCenterConfigObservedState) *pb.BitbucketDataCenterConfig {
	if in == nil {
		return nil
	}
	out := &pb.BitbucketDataCenterConfig{}
	// MISSING: HostURI
	// MISSING: WebhookSecretSecretVersion
	// MISSING: ReadAuthorizerCredential
	// MISSING: AuthorizerCredential
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.ValueOf(in.ServerVersion)
	return out
}
func CloudbuildConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.CloudbuildConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: GithubConfig
	// MISSING: GithubEnterpriseConfig
	// MISSING: GitlabConfig
	// MISSING: BitbucketDataCenterConfig
	// MISSING: BitbucketCloudConfig
	// MISSING: InstallationState
	// MISSING: Disabled
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func CloudbuildConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: GithubConfig
	// MISSING: GithubEnterpriseConfig
	// MISSING: GitlabConfig
	// MISSING: BitbucketDataCenterConfig
	// MISSING: BitbucketCloudConfig
	// MISSING: InstallationState
	// MISSING: Disabled
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func CloudbuildConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.CloudbuildConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: GithubConfig
	// MISSING: GithubEnterpriseConfig
	// MISSING: GitlabConfig
	// MISSING: BitbucketDataCenterConfig
	// MISSING: BitbucketCloudConfig
	// MISSING: InstallationState
	// MISSING: Disabled
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func CloudbuildConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildConnectionSpec) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: GithubConfig
	// MISSING: GithubEnterpriseConfig
	// MISSING: GitlabConfig
	// MISSING: BitbucketDataCenterConfig
	// MISSING: BitbucketCloudConfig
	// MISSING: InstallationState
	// MISSING: Disabled
	// MISSING: Reconciling
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func Connection_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.Connection {
	if in == nil {
		return nil
	}
	out := &krm.Connection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.GithubConfig = GitHubConfig_FromProto(mapCtx, in.GetGithubConfig())
	out.GithubEnterpriseConfig = GitHubEnterpriseConfig_FromProto(mapCtx, in.GetGithubEnterpriseConfig())
	out.GitlabConfig = GitLabConfig_FromProto(mapCtx, in.GetGitlabConfig())
	out.BitbucketDataCenterConfig = BitbucketDataCenterConfig_FromProto(mapCtx, in.GetBitbucketDataCenterConfig())
	out.BitbucketCloudConfig = BitbucketCloudConfig_FromProto(mapCtx, in.GetBitbucketCloudConfig())
	// MISSING: InstallationState
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Connection_ToProto(mapCtx *direct.MapContext, in *krm.Connection) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	if oneof := GitHubConfig_ToProto(mapCtx, in.GithubConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GithubConfig{GithubConfig: oneof}
	}
	if oneof := GitHubEnterpriseConfig_ToProto(mapCtx, in.GithubEnterpriseConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GithubEnterpriseConfig{GithubEnterpriseConfig: oneof}
	}
	if oneof := GitLabConfig_ToProto(mapCtx, in.GitlabConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GitlabConfig{GitlabConfig: oneof}
	}
	if oneof := BitbucketDataCenterConfig_ToProto(mapCtx, in.BitbucketDataCenterConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_BitbucketDataCenterConfig{BitbucketDataCenterConfig: oneof}
	}
	if oneof := BitbucketCloudConfig_ToProto(mapCtx, in.BitbucketCloudConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_BitbucketCloudConfig{BitbucketCloudConfig: oneof}
	}
	// MISSING: InstallationState
	out.Disabled = direct.ValueOf(in.Disabled)
	// MISSING: Reconciling
	out.Annotations = in.Annotations
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func ConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.ConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.GithubConfig = GitHubConfigObservedState_FromProto(mapCtx, in.GetGithubConfig())
	out.GithubEnterpriseConfig = GitHubEnterpriseConfigObservedState_FromProto(mapCtx, in.GetGithubEnterpriseConfig())
	out.GitlabConfig = GitLabConfigObservedState_FromProto(mapCtx, in.GetGitlabConfig())
	out.BitbucketDataCenterConfig = BitbucketDataCenterConfigObservedState_FromProto(mapCtx, in.GetBitbucketDataCenterConfig())
	// MISSING: BitbucketCloudConfig
	out.InstallationState = InstallationState_FromProto(mapCtx, in.GetInstallationState())
	// MISSING: Disabled
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func ConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectionObservedState) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if oneof := GitHubConfigObservedState_ToProto(mapCtx, in.GithubConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GithubConfig{GithubConfig: oneof}
	}
	if oneof := GitHubEnterpriseConfigObservedState_ToProto(mapCtx, in.GithubEnterpriseConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GithubEnterpriseConfig{GithubEnterpriseConfig: oneof}
	}
	if oneof := GitLabConfigObservedState_ToProto(mapCtx, in.GitlabConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_GitlabConfig{GitlabConfig: oneof}
	}
	if oneof := BitbucketDataCenterConfigObservedState_ToProto(mapCtx, in.BitbucketDataCenterConfig); oneof != nil {
		out.ConnectionConfig = &pb.Connection_BitbucketDataCenterConfig{BitbucketDataCenterConfig: oneof}
	}
	// MISSING: BitbucketCloudConfig
	out.InstallationState = InstallationState_ToProto(mapCtx, in.InstallationState)
	// MISSING: Disabled
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: Annotations
	// MISSING: Etag
	return out
}
func GitHubConfig_FromProto(mapCtx *direct.MapContext, in *pb.GitHubConfig) *krm.GitHubConfig {
	if in == nil {
		return nil
	}
	out := &krm.GitHubConfig{}
	out.AuthorizerCredential = OAuthCredential_FromProto(mapCtx, in.GetAuthorizerCredential())
	out.AppInstallationID = direct.LazyPtr(in.GetAppInstallationId())
	return out
}
func GitHubConfig_ToProto(mapCtx *direct.MapContext, in *krm.GitHubConfig) *pb.GitHubConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubConfig{}
	out.AuthorizerCredential = OAuthCredential_ToProto(mapCtx, in.AuthorizerCredential)
	out.AppInstallationId = direct.ValueOf(in.AppInstallationID)
	return out
}
func GitHubConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitHubConfig) *krm.GitHubConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GitHubConfigObservedState{}
	out.AuthorizerCredential = OAuthCredentialObservedState_FromProto(mapCtx, in.GetAuthorizerCredential())
	// MISSING: AppInstallationID
	return out
}
func GitHubConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GitHubConfigObservedState) *pb.GitHubConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubConfig{}
	out.AuthorizerCredential = OAuthCredentialObservedState_ToProto(mapCtx, in.AuthorizerCredential)
	// MISSING: AppInstallationID
	return out
}
func GitHubEnterpriseConfig_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEnterpriseConfig{}
	out.HostURI = direct.LazyPtr(in.GetHostUri())
	out.ApiKey = direct.LazyPtr(in.GetApiKey())
	out.AppID = direct.LazyPtr(in.GetAppId())
	out.AppSlug = direct.LazyPtr(in.GetAppSlug())
	out.PrivateKeySecretVersion = direct.LazyPtr(in.GetPrivateKeySecretVersion())
	out.WebhookSecretSecretVersion = direct.LazyPtr(in.GetWebhookSecretSecretVersion())
	out.AppInstallationID = direct.LazyPtr(in.GetAppInstallationId())
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectoryConfig())
	out.SslCa = direct.LazyPtr(in.GetSslCa())
	// MISSING: ServerVersion
	return out
}
func GitHubEnterpriseConfig_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEnterpriseConfig) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	out.HostUri = direct.ValueOf(in.HostURI)
	out.ApiKey = direct.ValueOf(in.ApiKey)
	out.AppId = direct.ValueOf(in.AppID)
	out.AppSlug = direct.ValueOf(in.AppSlug)
	out.PrivateKeySecretVersion = direct.ValueOf(in.PrivateKeySecretVersion)
	out.WebhookSecretSecretVersion = direct.ValueOf(in.WebhookSecretSecretVersion)
	out.AppInstallationId = direct.ValueOf(in.AppInstallationID)
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectoryConfig)
	out.SslCa = direct.ValueOf(in.SslCa)
	// MISSING: ServerVersion
	return out
}
func GitHubEnterpriseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEnterpriseConfig) *krm.GitHubEnterpriseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEnterpriseConfigObservedState{}
	// MISSING: HostURI
	// MISSING: ApiKey
	// MISSING: AppID
	// MISSING: AppSlug
	// MISSING: PrivateKeySecretVersion
	// MISSING: WebhookSecretSecretVersion
	// MISSING: AppInstallationID
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.LazyPtr(in.GetServerVersion())
	return out
}
func GitHubEnterpriseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEnterpriseConfigObservedState) *pb.GitHubEnterpriseConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEnterpriseConfig{}
	// MISSING: HostURI
	// MISSING: ApiKey
	// MISSING: AppID
	// MISSING: AppSlug
	// MISSING: PrivateKeySecretVersion
	// MISSING: WebhookSecretSecretVersion
	// MISSING: AppInstallationID
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.ValueOf(in.ServerVersion)
	return out
}
func GitLabConfig_FromProto(mapCtx *direct.MapContext, in *pb.GitLabConfig) *krm.GitLabConfig {
	if in == nil {
		return nil
	}
	out := &krm.GitLabConfig{}
	out.HostURI = direct.LazyPtr(in.GetHostUri())
	out.WebhookSecretSecretVersion = direct.LazyPtr(in.GetWebhookSecretSecretVersion())
	out.ReadAuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetReadAuthorizerCredential())
	out.AuthorizerCredential = UserCredential_FromProto(mapCtx, in.GetAuthorizerCredential())
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectoryConfig())
	out.SslCa = direct.LazyPtr(in.GetSslCa())
	// MISSING: ServerVersion
	return out
}
func GitLabConfig_ToProto(mapCtx *direct.MapContext, in *krm.GitLabConfig) *pb.GitLabConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitLabConfig{}
	out.HostUri = direct.ValueOf(in.HostURI)
	out.WebhookSecretSecretVersion = direct.ValueOf(in.WebhookSecretSecretVersion)
	out.ReadAuthorizerCredential = UserCredential_ToProto(mapCtx, in.ReadAuthorizerCredential)
	out.AuthorizerCredential = UserCredential_ToProto(mapCtx, in.AuthorizerCredential)
	out.ServiceDirectoryConfig = ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectoryConfig)
	out.SslCa = direct.ValueOf(in.SslCa)
	// MISSING: ServerVersion
	return out
}
func GitLabConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GitLabConfig) *krm.GitLabConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GitLabConfigObservedState{}
	// MISSING: HostURI
	// MISSING: WebhookSecretSecretVersion
	out.ReadAuthorizerCredential = UserCredentialObservedState_FromProto(mapCtx, in.GetReadAuthorizerCredential())
	// MISSING: AuthorizerCredential
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.LazyPtr(in.GetServerVersion())
	return out
}
func GitLabConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GitLabConfigObservedState) *pb.GitLabConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitLabConfig{}
	// MISSING: HostURI
	// MISSING: WebhookSecretSecretVersion
	out.ReadAuthorizerCredential = UserCredentialObservedState_ToProto(mapCtx, in.ReadAuthorizerCredential)
	// MISSING: AuthorizerCredential
	// MISSING: ServiceDirectoryConfig
	// MISSING: SslCa
	out.ServerVersion = direct.ValueOf(in.ServerVersion)
	return out
}
func InstallationState_FromProto(mapCtx *direct.MapContext, in *pb.InstallationState) *krm.InstallationState {
	if in == nil {
		return nil
	}
	out := &krm.InstallationState{}
	// MISSING: Stage
	// MISSING: Message
	// MISSING: ActionURI
	return out
}
func InstallationState_ToProto(mapCtx *direct.MapContext, in *krm.InstallationState) *pb.InstallationState {
	if in == nil {
		return nil
	}
	out := &pb.InstallationState{}
	// MISSING: Stage
	// MISSING: Message
	// MISSING: ActionURI
	return out
}
func InstallationStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstallationState) *krm.InstallationStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstallationStateObservedState{}
	out.Stage = direct.Enum_FromProto(mapCtx, in.GetStage())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ActionURI = direct.LazyPtr(in.GetActionUri())
	return out
}
func InstallationStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstallationStateObservedState) *pb.InstallationState {
	if in == nil {
		return nil
	}
	out := &pb.InstallationState{}
	out.Stage = direct.Enum_ToProto[pb.InstallationState_Stage](mapCtx, in.Stage)
	out.Message = direct.ValueOf(in.Message)
	out.ActionUri = direct.ValueOf(in.ActionURI)
	return out
}
func OAuthCredential_FromProto(mapCtx *direct.MapContext, in *pb.OAuthCredential) *krm.OAuthCredential {
	if in == nil {
		return nil
	}
	out := &krm.OAuthCredential{}
	out.OauthTokenSecretVersion = direct.LazyPtr(in.GetOauthTokenSecretVersion())
	// MISSING: Username
	return out
}
func OAuthCredential_ToProto(mapCtx *direct.MapContext, in *krm.OAuthCredential) *pb.OAuthCredential {
	if in == nil {
		return nil
	}
	out := &pb.OAuthCredential{}
	out.OauthTokenSecretVersion = direct.ValueOf(in.OauthTokenSecretVersion)
	// MISSING: Username
	return out
}
func OAuthCredentialObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OAuthCredential) *krm.OAuthCredentialObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OAuthCredentialObservedState{}
	// MISSING: OauthTokenSecretVersion
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func OAuthCredentialObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OAuthCredentialObservedState) *pb.OAuthCredential {
	if in == nil {
		return nil
	}
	out := &pb.OAuthCredential{}
	// MISSING: OauthTokenSecretVersion
	out.Username = direct.ValueOf(in.Username)
	return out
}
func ServiceDirectoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServiceDirectoryConfig) *krm.ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServiceDirectoryConfig{}
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func ServiceDirectoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServiceDirectoryConfig) *pb.ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServiceDirectoryConfig{}
	out.Service = direct.ValueOf(in.Service)
	return out
}
func UserCredential_FromProto(mapCtx *direct.MapContext, in *pb.UserCredential) *krm.UserCredential {
	if in == nil {
		return nil
	}
	out := &krm.UserCredential{}
	out.UserTokenSecretVersion = direct.LazyPtr(in.GetUserTokenSecretVersion())
	// MISSING: Username
	return out
}
func UserCredential_ToProto(mapCtx *direct.MapContext, in *krm.UserCredential) *pb.UserCredential {
	if in == nil {
		return nil
	}
	out := &pb.UserCredential{}
	out.UserTokenSecretVersion = direct.ValueOf(in.UserTokenSecretVersion)
	// MISSING: Username
	return out
}
func UserCredentialObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UserCredential) *krm.UserCredentialObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UserCredentialObservedState{}
	// MISSING: UserTokenSecretVersion
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func UserCredentialObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UserCredentialObservedState) *pb.UserCredential {
	if in == nil {
		return nil
	}
	out := &pb.UserCredential{}
	// MISSING: UserTokenSecretVersion
	out.Username = direct.ValueOf(in.Username)
	return out
}
