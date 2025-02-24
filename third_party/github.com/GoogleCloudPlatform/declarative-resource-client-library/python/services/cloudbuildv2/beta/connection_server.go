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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuildv2/beta/cloudbuildv2_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/beta"
)

// ConnectionServer implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnectionInstallationStateStageEnum converts a ConnectionInstallationStateStageEnum enum from its proto representation.
func ProtoToCloudbuildv2BetaConnectionInstallationStateStageEnum(e betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum) *beta.ConnectionInstallationStateStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum_name[int32(e)]; ok {
		e := beta.ConnectionInstallationStateStageEnum(n[len("Cloudbuildv2BetaConnectionInstallationStateStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectionGithubConfig converts a ConnectionGithubConfig object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGithubConfig(p *betapb.Cloudbuildv2BetaConnectionGithubConfig) *beta.ConnectionGithubConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGithubConfig{
		AuthorizerCredential: ProtoToCloudbuildv2BetaConnectionGithubConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		AppInstallationId:    dcl.Int64OrNil(p.GetAppInstallationId()),
	}
	return obj
}

// ProtoToConnectionGithubConfigAuthorizerCredential converts a ConnectionGithubConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGithubConfigAuthorizerCredential(p *betapb.Cloudbuildv2BetaConnectionGithubConfigAuthorizerCredential) *beta.ConnectionGithubConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGithubConfigAuthorizerCredential{
		OAuthTokenSecretVersion: dcl.StringOrNil(p.GetOauthTokenSecretVersion()),
		Username:                dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfig converts a ConnectionGithubEnterpriseConfig object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGithubEnterpriseConfig(p *betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfig) *beta.ConnectionGithubEnterpriseConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGithubEnterpriseConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		AppId:                      dcl.Int64OrNil(p.GetAppId()),
		AppSlug:                    dcl.StringOrNil(p.GetAppSlug()),
		PrivateKeySecretVersion:    dcl.StringOrNil(p.GetPrivateKeySecretVersion()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		AppInstallationId:          dcl.Int64OrNil(p.GetAppInstallationId()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfigServiceDirectoryConfig converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfig(p *betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfig) *beta.ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGithubEnterpriseConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionGitlabConfig converts a ConnectionGitlabConfig object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGitlabConfig(p *betapb.Cloudbuildv2BetaConnectionGitlabConfig) *beta.ConnectionGitlabConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGitlabConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		ReadAuthorizerCredential:   ProtoToCloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredential(p.GetReadAuthorizerCredential()),
		AuthorizerCredential:       ProtoToCloudbuildv2BetaConnectionGitlabConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
		ServerVersion:              dcl.StringOrNil(p.GetServerVersion()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigReadAuthorizerCredential converts a ConnectionGitlabConfigReadAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredential(p *betapb.Cloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredential) *beta.ConnectionGitlabConfigReadAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGitlabConfigReadAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigAuthorizerCredential converts a ConnectionGitlabConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGitlabConfigAuthorizerCredential(p *betapb.Cloudbuildv2BetaConnectionGitlabConfigAuthorizerCredential) *beta.ConnectionGitlabConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGitlabConfigAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigServiceDirectoryConfig converts a ConnectionGitlabConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfig(p *betapb.Cloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfig) *beta.ConnectionGitlabConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionGitlabConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionInstallationState converts a ConnectionInstallationState object from its proto representation.
func ProtoToCloudbuildv2BetaConnectionInstallationState(p *betapb.Cloudbuildv2BetaConnectionInstallationState) *beta.ConnectionInstallationState {
	if p == nil {
		return nil
	}
	obj := &beta.ConnectionInstallationState{
		Stage:     ProtoToCloudbuildv2BetaConnectionInstallationStateStageEnum(p.GetStage()),
		Message:   dcl.StringOrNil(p.GetMessage()),
		ActionUri: dcl.StringOrNil(p.GetActionUri()),
	}
	return obj
}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *betapb.Cloudbuildv2BetaConnection) *beta.Connection {
	obj := &beta.Connection{
		Name:                   dcl.StringOrNil(p.GetName()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		GithubConfig:           ProtoToCloudbuildv2BetaConnectionGithubConfig(p.GetGithubConfig()),
		GithubEnterpriseConfig: ProtoToCloudbuildv2BetaConnectionGithubEnterpriseConfig(p.GetGithubEnterpriseConfig()),
		GitlabConfig:           ProtoToCloudbuildv2BetaConnectionGitlabConfig(p.GetGitlabConfig()),
		InstallationState:      ProtoToCloudbuildv2BetaConnectionInstallationState(p.GetInstallationState()),
		Disabled:               dcl.Bool(p.GetDisabled()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ConnectionInstallationStateStageEnumToProto converts a ConnectionInstallationStateStageEnum enum to its proto representation.
func Cloudbuildv2BetaConnectionInstallationStateStageEnumToProto(e *beta.ConnectionInstallationStateStageEnum) betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum {
	if e == nil {
		return betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum(0)
	}
	if v, ok := betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum_value["ConnectionInstallationStateStageEnum"+string(*e)]; ok {
		return betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum(v)
	}
	return betapb.Cloudbuildv2BetaConnectionInstallationStateStageEnum(0)
}

// ConnectionGithubConfigToProto converts a ConnectionGithubConfig object to its proto representation.
func Cloudbuildv2BetaConnectionGithubConfigToProto(o *beta.ConnectionGithubConfig) *betapb.Cloudbuildv2BetaConnectionGithubConfig {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGithubConfig{}
	p.SetAuthorizerCredential(Cloudbuildv2BetaConnectionGithubConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	return p
}

// ConnectionGithubConfigAuthorizerCredentialToProto converts a ConnectionGithubConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2BetaConnectionGithubConfigAuthorizerCredentialToProto(o *beta.ConnectionGithubConfigAuthorizerCredential) *betapb.Cloudbuildv2BetaConnectionGithubConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGithubConfigAuthorizerCredential{}
	p.SetOauthTokenSecretVersion(dcl.ValueOrEmptyString(o.OAuthTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGithubEnterpriseConfigToProto converts a ConnectionGithubEnterpriseConfig object to its proto representation.
func Cloudbuildv2BetaConnectionGithubEnterpriseConfigToProto(o *beta.ConnectionGithubEnterpriseConfig) *betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfig {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetAppId(dcl.ValueOrEmptyInt64(o.AppId))
	p.SetAppSlug(dcl.ValueOrEmptyString(o.AppSlug))
	p.SetPrivateKeySecretVersion(dcl.ValueOrEmptyString(o.PrivateKeySecretVersion))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	p.SetServiceDirectoryConfig(Cloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	return p
}

// ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o *beta.ConnectionGithubEnterpriseConfigServiceDirectoryConfig) *betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGithubEnterpriseConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionGitlabConfigToProto converts a ConnectionGitlabConfig object to its proto representation.
func Cloudbuildv2BetaConnectionGitlabConfigToProto(o *beta.ConnectionGitlabConfig) *betapb.Cloudbuildv2BetaConnectionGitlabConfig {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGitlabConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetReadAuthorizerCredential(Cloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredentialToProto(o.ReadAuthorizerCredential))
	p.SetAuthorizerCredential(Cloudbuildv2BetaConnectionGitlabConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetServiceDirectoryConfig(Cloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	p.SetServerVersion(dcl.ValueOrEmptyString(o.ServerVersion))
	return p
}

// ConnectionGitlabConfigReadAuthorizerCredentialToProto converts a ConnectionGitlabConfigReadAuthorizerCredential object to its proto representation.
func Cloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredentialToProto(o *beta.ConnectionGitlabConfigReadAuthorizerCredential) *betapb.Cloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGitlabConfigReadAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigAuthorizerCredentialToProto converts a ConnectionGitlabConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2BetaConnectionGitlabConfigAuthorizerCredentialToProto(o *beta.ConnectionGitlabConfigAuthorizerCredential) *betapb.Cloudbuildv2BetaConnectionGitlabConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGitlabConfigAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigServiceDirectoryConfigToProto converts a ConnectionGitlabConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfigToProto(o *beta.ConnectionGitlabConfigServiceDirectoryConfig) *betapb.Cloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionGitlabConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionInstallationStateToProto converts a ConnectionInstallationState object to its proto representation.
func Cloudbuildv2BetaConnectionInstallationStateToProto(o *beta.ConnectionInstallationState) *betapb.Cloudbuildv2BetaConnectionInstallationState {
	if o == nil {
		return nil
	}
	p := &betapb.Cloudbuildv2BetaConnectionInstallationState{}
	p.SetStage(Cloudbuildv2BetaConnectionInstallationStateStageEnumToProto(o.Stage))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetActionUri(dcl.ValueOrEmptyString(o.ActionUri))
	return p
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *beta.Connection) *betapb.Cloudbuildv2BetaConnection {
	p := &betapb.Cloudbuildv2BetaConnection{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGithubConfig(Cloudbuildv2BetaConnectionGithubConfigToProto(resource.GithubConfig))
	p.SetGithubEnterpriseConfig(Cloudbuildv2BetaConnectionGithubEnterpriseConfigToProto(resource.GithubEnterpriseConfig))
	p.SetGitlabConfig(Cloudbuildv2BetaConnectionGitlabConfigToProto(resource.GitlabConfig))
	p.SetInstallationState(Cloudbuildv2BetaConnectionInstallationStateToProto(resource.InstallationState))
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
func (s *ConnectionServer) applyConnection(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudbuildv2BetaConnectionRequest) (*betapb.Cloudbuildv2BetaConnection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// applyCloudbuildv2BetaConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyCloudbuildv2BetaConnection(ctx context.Context, request *betapb.ApplyCloudbuildv2BetaConnectionRequest) (*betapb.Cloudbuildv2BetaConnection, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteCloudbuildv2BetaConnection(ctx context.Context, request *betapb.DeleteCloudbuildv2BetaConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListCloudbuildv2BetaConnection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListCloudbuildv2BetaConnection(ctx context.Context, request *betapb.ListCloudbuildv2BetaConnectionRequest) (*betapb.ListCloudbuildv2BetaConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.Cloudbuildv2BetaConnection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudbuildv2BetaConnectionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
