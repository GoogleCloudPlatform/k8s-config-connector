// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuildv2/alpha/cloudbuildv2_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/alpha"
)

// ConnectionServer implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnectionInstallationStateStageEnum converts a ConnectionInstallationStateStageEnum enum from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionInstallationStateStageEnum(e alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum) *alpha.ConnectionInstallationStateStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum_name[int32(e)]; ok {
		e := alpha.ConnectionInstallationStateStageEnum(n[len("Cloudbuildv2AlphaConnectionInstallationStateStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectionGithubConfig converts a ConnectionGithubConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubConfig) *alpha.ConnectionGithubConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubConfig{
		AuthorizerCredential: ProtoToCloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		AppInstallationId:    dcl.Int64OrNil(p.GetAppInstallationId()),
	}
	return obj
}

// ProtoToConnectionGithubConfigAuthorizerCredential converts a ConnectionGithubConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential) *alpha.ConnectionGithubConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubConfigAuthorizerCredential{
		OAuthTokenSecretVersion: dcl.StringOrNil(p.GetOauthTokenSecretVersion()),
		Username:                dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfig converts a ConnectionGithubEnterpriseConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig) *alpha.ConnectionGithubEnterpriseConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubEnterpriseConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		AppId:                      dcl.Int64OrNil(p.GetAppId()),
		AppSlug:                    dcl.StringOrNil(p.GetAppSlug()),
		PrivateKeySecretVersion:    dcl.StringOrNil(p.GetPrivateKeySecretVersion()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		AppInstallationId:          dcl.Int64OrNil(p.GetAppInstallationId()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfigServiceDirectoryConfig converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig) *alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionGitlabConfig converts a ConnectionGitlabConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGitlabConfig(p *alphapb.Cloudbuildv2AlphaConnectionGitlabConfig) *alpha.ConnectionGitlabConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGitlabConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		ReadAuthorizerCredential:   ProtoToCloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredential(p.GetReadAuthorizerCredential()),
		AuthorizerCredential:       ProtoToCloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
		ServerVersion:              dcl.StringOrNil(p.GetServerVersion()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigReadAuthorizerCredential converts a ConnectionGitlabConfigReadAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredential) *alpha.ConnectionGitlabConfigReadAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGitlabConfigReadAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigAuthorizerCredential converts a ConnectionGitlabConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredential) *alpha.ConnectionGitlabConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGitlabConfigAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigServiceDirectoryConfig converts a ConnectionGitlabConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfig(p *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfig) *alpha.ConnectionGitlabConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionGitlabConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionBitbucketDataCenterConfig converts a ConnectionBitbucketDataCenterConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfig(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfig) *alpha.ConnectionBitbucketDataCenterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketDataCenterConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		ReadAuthorizerCredential:   ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredential(p.GetReadAuthorizerCredential()),
		AuthorizerCredential:       ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
		ServerVersion:              dcl.StringOrNil(p.GetServerVersion()),
	}
	return obj
}

// ProtoToConnectionBitbucketDataCenterConfigReadAuthorizerCredential converts a ConnectionBitbucketDataCenterConfigReadAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredential) *alpha.ConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketDataCenterConfigReadAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionBitbucketDataCenterConfigAuthorizerCredential converts a ConnectionBitbucketDataCenterConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredential) *alpha.ConnectionBitbucketDataCenterConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketDataCenterConfigAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionBitbucketDataCenterConfigServiceDirectoryConfig converts a ConnectionBitbucketDataCenterConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfig(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfig) *alpha.ConnectionBitbucketDataCenterConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketDataCenterConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionBitbucketCloudConfig converts a ConnectionBitbucketCloudConfig object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfig(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfig) *alpha.ConnectionBitbucketCloudConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketCloudConfig{
		Workspace:                  dcl.StringOrNil(p.GetWorkspace()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		ReadAuthorizerCredential:   ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredential(p.GetReadAuthorizerCredential()),
		AuthorizerCredential:       ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredential(p.GetAuthorizerCredential()),
	}
	return obj
}

// ProtoToConnectionBitbucketCloudConfigReadAuthorizerCredential converts a ConnectionBitbucketCloudConfigReadAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredential) *alpha.ConnectionBitbucketCloudConfigReadAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketCloudConfigReadAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionBitbucketCloudConfigAuthorizerCredential converts a ConnectionBitbucketCloudConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredential(p *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredential) *alpha.ConnectionBitbucketCloudConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionBitbucketCloudConfigAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionInstallationState converts a ConnectionInstallationState object from its proto representation.
func ProtoToCloudbuildv2AlphaConnectionInstallationState(p *alphapb.Cloudbuildv2AlphaConnectionInstallationState) *alpha.ConnectionInstallationState {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectionInstallationState{
		Stage:     ProtoToCloudbuildv2AlphaConnectionInstallationStateStageEnum(p.GetStage()),
		Message:   dcl.StringOrNil(p.GetMessage()),
		ActionUri: dcl.StringOrNil(p.GetActionUri()),
	}
	return obj
}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *alphapb.Cloudbuildv2AlphaConnection) *alpha.Connection {
	obj := &alpha.Connection{
		Name:                      dcl.StringOrNil(p.GetName()),
		CreateTime:                dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                dcl.StringOrNil(p.GetUpdateTime()),
		GithubConfig:              ProtoToCloudbuildv2AlphaConnectionGithubConfig(p.GetGithubConfig()),
		GithubEnterpriseConfig:    ProtoToCloudbuildv2AlphaConnectionGithubEnterpriseConfig(p.GetGithubEnterpriseConfig()),
		GitlabConfig:              ProtoToCloudbuildv2AlphaConnectionGitlabConfig(p.GetGitlabConfig()),
		BitbucketDataCenterConfig: ProtoToCloudbuildv2AlphaConnectionBitbucketDataCenterConfig(p.GetBitbucketDataCenterConfig()),
		BitbucketCloudConfig:      ProtoToCloudbuildv2AlphaConnectionBitbucketCloudConfig(p.GetBitbucketCloudConfig()),
		InstallationState:         ProtoToCloudbuildv2AlphaConnectionInstallationState(p.GetInstallationState()),
		Disabled:                  dcl.Bool(p.GetDisabled()),
		Reconciling:               dcl.Bool(p.GetReconciling()),
		Etag:                      dcl.StringOrNil(p.GetEtag()),
		Project:                   dcl.StringOrNil(p.GetProject()),
		Location:                  dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ConnectionInstallationStateStageEnumToProto converts a ConnectionInstallationStateStageEnum enum to its proto representation.
func Cloudbuildv2AlphaConnectionInstallationStateStageEnumToProto(e *alpha.ConnectionInstallationStateStageEnum) alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum {
	if e == nil {
		return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(0)
	}
	if v, ok := alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum_value["ConnectionInstallationStateStageEnum"+string(*e)]; ok {
		return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(v)
	}
	return alphapb.Cloudbuildv2AlphaConnectionInstallationStateStageEnum(0)
}

// ConnectionGithubConfigToProto converts a ConnectionGithubConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubConfigToProto(o *alpha.ConnectionGithubConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubConfig{}
	p.SetAuthorizerCredential(Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	return p
}

// ConnectionGithubConfigAuthorizerCredentialToProto converts a ConnectionGithubConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredentialToProto(o *alpha.ConnectionGithubConfigAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubConfigAuthorizerCredential{}
	p.SetOauthTokenSecretVersion(dcl.ValueOrEmptyString(o.OAuthTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGithubEnterpriseConfigToProto converts a ConnectionGithubEnterpriseConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubEnterpriseConfigToProto(o *alpha.ConnectionGithubEnterpriseConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetAppId(dcl.ValueOrEmptyInt64(o.AppId))
	p.SetAppSlug(dcl.ValueOrEmptyString(o.AppSlug))
	p.SetPrivateKeySecretVersion(dcl.ValueOrEmptyString(o.PrivateKeySecretVersion))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	p.SetServiceDirectoryConfig(Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	return p
}

// ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o *alpha.ConnectionGithubEnterpriseConfigServiceDirectoryConfig) *alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGithubEnterpriseConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionGitlabConfigToProto converts a ConnectionGitlabConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGitlabConfigToProto(o *alpha.ConnectionGitlabConfig) *alphapb.Cloudbuildv2AlphaConnectionGitlabConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGitlabConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetReadAuthorizerCredential(Cloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredentialToProto(o.ReadAuthorizerCredential))
	p.SetAuthorizerCredential(Cloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetServiceDirectoryConfig(Cloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	p.SetServerVersion(dcl.ValueOrEmptyString(o.ServerVersion))
	return p
}

// ConnectionGitlabConfigReadAuthorizerCredentialToProto converts a ConnectionGitlabConfigReadAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredentialToProto(o *alpha.ConnectionGitlabConfigReadAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGitlabConfigReadAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigAuthorizerCredentialToProto converts a ConnectionGitlabConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredentialToProto(o *alpha.ConnectionGitlabConfigAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGitlabConfigAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigServiceDirectoryConfigToProto converts a ConnectionGitlabConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfigToProto(o *alpha.ConnectionGitlabConfigServiceDirectoryConfig) *alphapb.Cloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionGitlabConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionBitbucketDataCenterConfigToProto converts a ConnectionBitbucketDataCenterConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigToProto(o *alpha.ConnectionBitbucketDataCenterConfig) *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetReadAuthorizerCredential(Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredentialToProto(o.ReadAuthorizerCredential))
	p.SetAuthorizerCredential(Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetServiceDirectoryConfig(Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	p.SetServerVersion(dcl.ValueOrEmptyString(o.ServerVersion))
	return p
}

// ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialToProto converts a ConnectionBitbucketDataCenterConfigReadAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredentialToProto(o *alpha.ConnectionBitbucketDataCenterConfigReadAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigReadAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionBitbucketDataCenterConfigAuthorizerCredentialToProto converts a ConnectionBitbucketDataCenterConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredentialToProto(o *alpha.ConnectionBitbucketDataCenterConfigAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionBitbucketDataCenterConfigServiceDirectoryConfigToProto converts a ConnectionBitbucketDataCenterConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfigToProto(o *alpha.ConnectionBitbucketDataCenterConfigServiceDirectoryConfig) *alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionBitbucketCloudConfigToProto converts a ConnectionBitbucketCloudConfig object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketCloudConfigToProto(o *alpha.ConnectionBitbucketCloudConfig) *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfig{}
	p.SetWorkspace(dcl.ValueOrEmptyString(o.Workspace))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetReadAuthorizerCredential(Cloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredentialToProto(o.ReadAuthorizerCredential))
	p.SetAuthorizerCredential(Cloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	return p
}

// ConnectionBitbucketCloudConfigReadAuthorizerCredentialToProto converts a ConnectionBitbucketCloudConfigReadAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredentialToProto(o *alpha.ConnectionBitbucketCloudConfigReadAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigReadAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionBitbucketCloudConfigAuthorizerCredentialToProto converts a ConnectionBitbucketCloudConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredentialToProto(o *alpha.ConnectionBitbucketCloudConfigAuthorizerCredential) *alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionBitbucketCloudConfigAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionInstallationStateToProto converts a ConnectionInstallationState object to its proto representation.
func Cloudbuildv2AlphaConnectionInstallationStateToProto(o *alpha.ConnectionInstallationState) *alphapb.Cloudbuildv2AlphaConnectionInstallationState {
	if o == nil {
		return nil
	}
	p := &alphapb.Cloudbuildv2AlphaConnectionInstallationState{}
	p.SetStage(Cloudbuildv2AlphaConnectionInstallationStateStageEnumToProto(o.Stage))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetActionUri(dcl.ValueOrEmptyString(o.ActionUri))
	return p
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *alpha.Connection) *alphapb.Cloudbuildv2AlphaConnection {
	p := &alphapb.Cloudbuildv2AlphaConnection{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGithubConfig(Cloudbuildv2AlphaConnectionGithubConfigToProto(resource.GithubConfig))
	p.SetGithubEnterpriseConfig(Cloudbuildv2AlphaConnectionGithubEnterpriseConfigToProto(resource.GithubEnterpriseConfig))
	p.SetGitlabConfig(Cloudbuildv2AlphaConnectionGitlabConfigToProto(resource.GitlabConfig))
	p.SetBitbucketDataCenterConfig(Cloudbuildv2AlphaConnectionBitbucketDataCenterConfigToProto(resource.BitbucketDataCenterConfig))
	p.SetBitbucketCloudConfig(Cloudbuildv2AlphaConnectionBitbucketCloudConfigToProto(resource.BitbucketCloudConfig))
	p.SetInstallationState(Cloudbuildv2AlphaConnectionInstallationStateToProto(resource.InstallationState))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) applyConnection(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudbuildv2AlphaConnectionRequest) (*alphapb.Cloudbuildv2AlphaConnection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// applyCloudbuildv2AlphaConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.ApplyCloudbuildv2AlphaConnectionRequest) (*alphapb.Cloudbuildv2AlphaConnection, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.DeleteCloudbuildv2AlphaConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListCloudbuildv2AlphaConnection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListCloudbuildv2AlphaConnection(ctx context.Context, request *alphapb.ListCloudbuildv2AlphaConnectionRequest) (*alphapb.ListCloudbuildv2AlphaConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.Cloudbuildv2AlphaConnection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudbuildv2AlphaConnectionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
