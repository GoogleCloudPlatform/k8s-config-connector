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
	cloudbuildv2pb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuildv2/cloudbuildv2_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2"
)

// ConnectionServer implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnectionInstallationStateStageEnum converts a ConnectionInstallationStateStageEnum enum from its proto representation.
func ProtoToCloudbuildv2ConnectionInstallationStateStageEnum(e cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum) *cloudbuildv2.ConnectionInstallationStateStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum_name[int32(e)]; ok {
		e := cloudbuildv2.ConnectionInstallationStateStageEnum(n[len("Cloudbuildv2ConnectionInstallationStateStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectionGithubConfig converts a ConnectionGithubConfig object from its proto representation.
func ProtoToCloudbuildv2ConnectionGithubConfig(p *cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfig) *cloudbuildv2.ConnectionGithubConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGithubConfig{
		AuthorizerCredential: ProtoToCloudbuildv2ConnectionGithubConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		AppInstallationId:    dcl.Int64OrNil(p.GetAppInstallationId()),
	}
	return obj
}

// ProtoToConnectionGithubConfigAuthorizerCredential converts a ConnectionGithubConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2ConnectionGithubConfigAuthorizerCredential(p *cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfigAuthorizerCredential) *cloudbuildv2.ConnectionGithubConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGithubConfigAuthorizerCredential{
		OAuthTokenSecretVersion: dcl.StringOrNil(p.GetOauthTokenSecretVersion()),
		Username:                dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfig converts a ConnectionGithubEnterpriseConfig object from its proto representation.
func ProtoToCloudbuildv2ConnectionGithubEnterpriseConfig(p *cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfig) *cloudbuildv2.ConnectionGithubEnterpriseConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGithubEnterpriseConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		AppId:                      dcl.Int64OrNil(p.GetAppId()),
		AppSlug:                    dcl.StringOrNil(p.GetAppSlug()),
		PrivateKeySecretVersion:    dcl.StringOrNil(p.GetPrivateKeySecretVersion()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		AppInstallationId:          dcl.Int64OrNil(p.GetAppInstallationId()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
	}
	return obj
}

// ProtoToConnectionGithubEnterpriseConfigServiceDirectoryConfig converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig(p *cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig) *cloudbuildv2.ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGithubEnterpriseConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionGitlabConfig converts a ConnectionGitlabConfig object from its proto representation.
func ProtoToCloudbuildv2ConnectionGitlabConfig(p *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfig) *cloudbuildv2.ConnectionGitlabConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGitlabConfig{
		HostUri:                    dcl.StringOrNil(p.GetHostUri()),
		WebhookSecretSecretVersion: dcl.StringOrNil(p.GetWebhookSecretSecretVersion()),
		ReadAuthorizerCredential:   ProtoToCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential(p.GetReadAuthorizerCredential()),
		AuthorizerCredential:       ProtoToCloudbuildv2ConnectionGitlabConfigAuthorizerCredential(p.GetAuthorizerCredential()),
		ServiceDirectoryConfig:     ProtoToCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig(p.GetServiceDirectoryConfig()),
		SslCa:                      dcl.StringOrNil(p.GetSslCa()),
		ServerVersion:              dcl.StringOrNil(p.GetServerVersion()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigReadAuthorizerCredential converts a ConnectionGitlabConfigReadAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential(p *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential) *cloudbuildv2.ConnectionGitlabConfigReadAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGitlabConfigReadAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigAuthorizerCredential converts a ConnectionGitlabConfigAuthorizerCredential object from its proto representation.
func ProtoToCloudbuildv2ConnectionGitlabConfigAuthorizerCredential(p *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential) *cloudbuildv2.ConnectionGitlabConfigAuthorizerCredential {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGitlabConfigAuthorizerCredential{
		UserTokenSecretVersion: dcl.StringOrNil(p.GetUserTokenSecretVersion()),
		Username:               dcl.StringOrNil(p.GetUsername()),
	}
	return obj
}

// ProtoToConnectionGitlabConfigServiceDirectoryConfig converts a ConnectionGitlabConfigServiceDirectoryConfig object from its proto representation.
func ProtoToCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig(p *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig) *cloudbuildv2.ConnectionGitlabConfigServiceDirectoryConfig {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionGitlabConfigServiceDirectoryConfig{
		Service: dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToConnectionInstallationState converts a ConnectionInstallationState object from its proto representation.
func ProtoToCloudbuildv2ConnectionInstallationState(p *cloudbuildv2pb.Cloudbuildv2ConnectionInstallationState) *cloudbuildv2.ConnectionInstallationState {
	if p == nil {
		return nil
	}
	obj := &cloudbuildv2.ConnectionInstallationState{
		Stage:     ProtoToCloudbuildv2ConnectionInstallationStateStageEnum(p.GetStage()),
		Message:   dcl.StringOrNil(p.GetMessage()),
		ActionUri: dcl.StringOrNil(p.GetActionUri()),
	}
	return obj
}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *cloudbuildv2pb.Cloudbuildv2Connection) *cloudbuildv2.Connection {
	obj := &cloudbuildv2.Connection{
		Name:                   dcl.StringOrNil(p.GetName()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:             dcl.StringOrNil(p.GetUpdateTime()),
		GithubConfig:           ProtoToCloudbuildv2ConnectionGithubConfig(p.GetGithubConfig()),
		GithubEnterpriseConfig: ProtoToCloudbuildv2ConnectionGithubEnterpriseConfig(p.GetGithubEnterpriseConfig()),
		GitlabConfig:           ProtoToCloudbuildv2ConnectionGitlabConfig(p.GetGitlabConfig()),
		InstallationState:      ProtoToCloudbuildv2ConnectionInstallationState(p.GetInstallationState()),
		Disabled:               dcl.Bool(p.GetDisabled()),
		Reconciling:            dcl.Bool(p.GetReconciling()),
		Etag:                   dcl.StringOrNil(p.GetEtag()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ConnectionInstallationStateStageEnumToProto converts a ConnectionInstallationStateStageEnum enum to its proto representation.
func Cloudbuildv2ConnectionInstallationStateStageEnumToProto(e *cloudbuildv2.ConnectionInstallationStateStageEnum) cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum {
	if e == nil {
		return cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum(0)
	}
	if v, ok := cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum_value["ConnectionInstallationStateStageEnum"+string(*e)]; ok {
		return cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum(v)
	}
	return cloudbuildv2pb.Cloudbuildv2ConnectionInstallationStateStageEnum(0)
}

// ConnectionGithubConfigToProto converts a ConnectionGithubConfig object to its proto representation.
func Cloudbuildv2ConnectionGithubConfigToProto(o *cloudbuildv2.ConnectionGithubConfig) *cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfig{}
	p.SetAuthorizerCredential(Cloudbuildv2ConnectionGithubConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	return p
}

// ConnectionGithubConfigAuthorizerCredentialToProto converts a ConnectionGithubConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2ConnectionGithubConfigAuthorizerCredentialToProto(o *cloudbuildv2.ConnectionGithubConfigAuthorizerCredential) *cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGithubConfigAuthorizerCredential{}
	p.SetOauthTokenSecretVersion(dcl.ValueOrEmptyString(o.OAuthTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGithubEnterpriseConfigToProto converts a ConnectionGithubEnterpriseConfig object to its proto representation.
func Cloudbuildv2ConnectionGithubEnterpriseConfigToProto(o *cloudbuildv2.ConnectionGithubEnterpriseConfig) *cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetAppId(dcl.ValueOrEmptyInt64(o.AppId))
	p.SetAppSlug(dcl.ValueOrEmptyString(o.AppSlug))
	p.SetPrivateKeySecretVersion(dcl.ValueOrEmptyString(o.PrivateKeySecretVersion))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetAppInstallationId(dcl.ValueOrEmptyInt64(o.AppInstallationId))
	p.SetServiceDirectoryConfig(Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	return p
}

// ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto converts a ConnectionGithubEnterpriseConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfigToProto(o *cloudbuildv2.ConnectionGithubEnterpriseConfigServiceDirectoryConfig) *cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGithubEnterpriseConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionGitlabConfigToProto converts a ConnectionGitlabConfig object to its proto representation.
func Cloudbuildv2ConnectionGitlabConfigToProto(o *cloudbuildv2.ConnectionGitlabConfig) *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfig{}
	p.SetHostUri(dcl.ValueOrEmptyString(o.HostUri))
	p.SetWebhookSecretSecretVersion(dcl.ValueOrEmptyString(o.WebhookSecretSecretVersion))
	p.SetReadAuthorizerCredential(Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialToProto(o.ReadAuthorizerCredential))
	p.SetAuthorizerCredential(Cloudbuildv2ConnectionGitlabConfigAuthorizerCredentialToProto(o.AuthorizerCredential))
	p.SetServiceDirectoryConfig(Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigToProto(o.ServiceDirectoryConfig))
	p.SetSslCa(dcl.ValueOrEmptyString(o.SslCa))
	p.SetServerVersion(dcl.ValueOrEmptyString(o.ServerVersion))
	return p
}

// ConnectionGitlabConfigReadAuthorizerCredentialToProto converts a ConnectionGitlabConfigReadAuthorizerCredential object to its proto representation.
func Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredentialToProto(o *cloudbuildv2.ConnectionGitlabConfigReadAuthorizerCredential) *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigAuthorizerCredentialToProto converts a ConnectionGitlabConfigAuthorizerCredential object to its proto representation.
func Cloudbuildv2ConnectionGitlabConfigAuthorizerCredentialToProto(o *cloudbuildv2.ConnectionGitlabConfigAuthorizerCredential) *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigAuthorizerCredential{}
	p.SetUserTokenSecretVersion(dcl.ValueOrEmptyString(o.UserTokenSecretVersion))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	return p
}

// ConnectionGitlabConfigServiceDirectoryConfigToProto converts a ConnectionGitlabConfigServiceDirectoryConfig object to its proto representation.
func Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfigToProto(o *cloudbuildv2.ConnectionGitlabConfigServiceDirectoryConfig) *cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ConnectionInstallationStateToProto converts a ConnectionInstallationState object to its proto representation.
func Cloudbuildv2ConnectionInstallationStateToProto(o *cloudbuildv2.ConnectionInstallationState) *cloudbuildv2pb.Cloudbuildv2ConnectionInstallationState {
	if o == nil {
		return nil
	}
	p := &cloudbuildv2pb.Cloudbuildv2ConnectionInstallationState{}
	p.SetStage(Cloudbuildv2ConnectionInstallationStateStageEnumToProto(o.Stage))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetActionUri(dcl.ValueOrEmptyString(o.ActionUri))
	return p
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *cloudbuildv2.Connection) *cloudbuildv2pb.Cloudbuildv2Connection {
	p := &cloudbuildv2pb.Cloudbuildv2Connection{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetGithubConfig(Cloudbuildv2ConnectionGithubConfigToProto(resource.GithubConfig))
	p.SetGithubEnterpriseConfig(Cloudbuildv2ConnectionGithubEnterpriseConfigToProto(resource.GithubEnterpriseConfig))
	p.SetGitlabConfig(Cloudbuildv2ConnectionGitlabConfigToProto(resource.GitlabConfig))
	p.SetInstallationState(Cloudbuildv2ConnectionInstallationStateToProto(resource.InstallationState))
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
func (s *ConnectionServer) applyConnection(ctx context.Context, c *cloudbuildv2.Client, request *cloudbuildv2pb.ApplyCloudbuildv2ConnectionRequest) (*cloudbuildv2pb.Cloudbuildv2Connection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// applyCloudbuildv2Connection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyCloudbuildv2Connection(ctx context.Context, request *cloudbuildv2pb.ApplyCloudbuildv2ConnectionRequest) (*cloudbuildv2pb.Cloudbuildv2Connection, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteCloudbuildv2Connection(ctx context.Context, request *cloudbuildv2pb.DeleteCloudbuildv2ConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListCloudbuildv2Connection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListCloudbuildv2Connection(ctx context.Context, request *cloudbuildv2pb.ListCloudbuildv2ConnectionRequest) (*cloudbuildv2pb.ListCloudbuildv2ConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*cloudbuildv2pb.Cloudbuildv2Connection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudbuildv2pb.ListCloudbuildv2ConnectionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*cloudbuildv2.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudbuildv2.NewClient(conf), nil
}
