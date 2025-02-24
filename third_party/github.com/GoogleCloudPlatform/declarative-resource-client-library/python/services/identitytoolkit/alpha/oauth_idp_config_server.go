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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/alpha/identitytoolkit_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
)

// OAuthIdpConfigServer implements the gRPC interface for OAuthIdpConfig.
type OAuthIdpConfigServer struct{}

// ProtoToOAuthIdpConfigResponseType converts a OAuthIdpConfigResponseType object from its proto representation.
func ProtoToIdentitytoolkitAlphaOAuthIdpConfigResponseType(p *alphapb.IdentitytoolkitAlphaOAuthIdpConfigResponseType) *alpha.OAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &alpha.OAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.GetIdToken()),
		Code:    dcl.Bool(p.GetCode()),
		Token:   dcl.Bool(p.GetToken()),
	}
	return obj
}

// ProtoToOAuthIdpConfig converts a OAuthIdpConfig resource from its proto representation.
func ProtoToOAuthIdpConfig(p *alphapb.IdentitytoolkitAlphaOAuthIdpConfig) *alpha.OAuthIdpConfig {
	obj := &alpha.OAuthIdpConfig{
		Name:         dcl.StringOrNil(p.GetName()),
		ClientId:     dcl.StringOrNil(p.GetClientId()),
		Issuer:       dcl.StringOrNil(p.GetIssuer()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		Enabled:      dcl.Bool(p.GetEnabled()),
		ClientSecret: dcl.StringOrNil(p.GetClientSecret()),
		ResponseType: ProtoToIdentitytoolkitAlphaOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// OAuthIdpConfigResponseTypeToProto converts a OAuthIdpConfigResponseType object to its proto representation.
func IdentitytoolkitAlphaOAuthIdpConfigResponseTypeToProto(o *alpha.OAuthIdpConfigResponseType) *alphapb.IdentitytoolkitAlphaOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaOAuthIdpConfigResponseType{}
	p.SetIdToken(dcl.ValueOrEmptyBool(o.IdToken))
	p.SetCode(dcl.ValueOrEmptyBool(o.Code))
	p.SetToken(dcl.ValueOrEmptyBool(o.Token))
	return p
}

// OAuthIdpConfigToProto converts a OAuthIdpConfig resource to its proto representation.
func OAuthIdpConfigToProto(resource *alpha.OAuthIdpConfig) *alphapb.IdentitytoolkitAlphaOAuthIdpConfig {
	p := &alphapb.IdentitytoolkitAlphaOAuthIdpConfig{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetClientId(dcl.ValueOrEmptyString(resource.ClientId))
	p.SetIssuer(dcl.ValueOrEmptyString(resource.Issuer))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetEnabled(dcl.ValueOrEmptyBool(resource.Enabled))
	p.SetClientSecret(dcl.ValueOrEmptyString(resource.ClientSecret))
	p.SetResponseType(IdentitytoolkitAlphaOAuthIdpConfigResponseTypeToProto(resource.ResponseType))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) applyOAuthIdpConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIdentitytoolkitAlphaOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaOAuthIdpConfig, error) {
	p := ProtoToOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OAuthIdpConfigToProto(res)
	return r, nil
}

// applyIdentitytoolkitAlphaOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Apply() method.
func (s *OAuthIdpConfigServer) ApplyIdentitytoolkitAlphaOAuthIdpConfig(ctx context.Context, request *alphapb.ApplyIdentitytoolkitAlphaOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaOAuthIdpConfig, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOAuthIdpConfig(ctx, cl, request)
}

// DeleteOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfig Delete() method.
func (s *OAuthIdpConfigServer) DeleteIdentitytoolkitAlphaOAuthIdpConfig(ctx context.Context, request *alphapb.DeleteIdentitytoolkitAlphaOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOAuthIdpConfig(ctx, ProtoToOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitAlphaOAuthIdpConfig handles the gRPC request by passing it to the underlying OAuthIdpConfigList() method.
func (s *OAuthIdpConfigServer) ListIdentitytoolkitAlphaOAuthIdpConfig(ctx context.Context, request *alphapb.ListIdentitytoolkitAlphaOAuthIdpConfigRequest) (*alphapb.ListIdentitytoolkitAlphaOAuthIdpConfigResponse, error) {
	cl, err := createConfigOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOAuthIdpConfig(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IdentitytoolkitAlphaOAuthIdpConfig
	for _, r := range resources.Items {
		rp := OAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIdentitytoolkitAlphaOAuthIdpConfigResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOAuthIdpConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
