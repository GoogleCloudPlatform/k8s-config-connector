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

// TenantOAuthIdpConfigServer implements the gRPC interface for TenantOAuthIdpConfig.
type TenantOAuthIdpConfigServer struct{}

// ProtoToTenantOAuthIdpConfigResponseType converts a TenantOAuthIdpConfigResponseType object from its proto representation.
func ProtoToIdentitytoolkitAlphaTenantOAuthIdpConfigResponseType(p *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType) *alpha.TenantOAuthIdpConfigResponseType {
	if p == nil {
		return nil
	}
	obj := &alpha.TenantOAuthIdpConfigResponseType{
		IdToken: dcl.Bool(p.GetIdToken()),
		Code:    dcl.Bool(p.GetCode()),
		Token:   dcl.Bool(p.GetToken()),
	}
	return obj
}

// ProtoToTenantOAuthIdpConfig converts a TenantOAuthIdpConfig resource from its proto representation.
func ProtoToTenantOAuthIdpConfig(p *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig) *alpha.TenantOAuthIdpConfig {
	obj := &alpha.TenantOAuthIdpConfig{
		Name:         dcl.StringOrNil(p.GetName()),
		ClientId:     dcl.StringOrNil(p.GetClientId()),
		Issuer:       dcl.StringOrNil(p.GetIssuer()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
		Enabled:      dcl.Bool(p.GetEnabled()),
		ClientSecret: dcl.StringOrNil(p.GetClientSecret()),
		ResponseType: ProtoToIdentitytoolkitAlphaTenantOAuthIdpConfigResponseType(p.GetResponseType()),
		Project:      dcl.StringOrNil(p.GetProject()),
		Tenant:       dcl.StringOrNil(p.GetTenant()),
	}
	return obj
}

// TenantOAuthIdpConfigResponseTypeToProto converts a TenantOAuthIdpConfigResponseType object to its proto representation.
func IdentitytoolkitAlphaTenantOAuthIdpConfigResponseTypeToProto(o *alpha.TenantOAuthIdpConfigResponseType) *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfigResponseType{}
	p.SetIdToken(dcl.ValueOrEmptyBool(o.IdToken))
	p.SetCode(dcl.ValueOrEmptyBool(o.Code))
	p.SetToken(dcl.ValueOrEmptyBool(o.Token))
	return p
}

// TenantOAuthIdpConfigToProto converts a TenantOAuthIdpConfig resource to its proto representation.
func TenantOAuthIdpConfigToProto(resource *alpha.TenantOAuthIdpConfig) *alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig {
	p := &alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetClientId(dcl.ValueOrEmptyString(resource.ClientId))
	p.SetIssuer(dcl.ValueOrEmptyString(resource.Issuer))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetEnabled(dcl.ValueOrEmptyBool(resource.Enabled))
	p.SetClientSecret(dcl.ValueOrEmptyString(resource.ClientSecret))
	p.SetResponseType(IdentitytoolkitAlphaTenantOAuthIdpConfigResponseTypeToProto(resource.ResponseType))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetTenant(dcl.ValueOrEmptyString(resource.Tenant))

	return p
}

// applyTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) applyTenantOAuthIdpConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig, error) {
	p := ProtoToTenantOAuthIdpConfig(request.GetResource())
	res, err := c.ApplyTenantOAuthIdpConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantOAuthIdpConfigToProto(res)
	return r, nil
}

// applyIdentitytoolkitAlphaTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Apply() method.
func (s *TenantOAuthIdpConfigServer) ApplyIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.ApplyIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTenantOAuthIdpConfig(ctx, cl, request)
}

// DeleteTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfig Delete() method.
func (s *TenantOAuthIdpConfigServer) DeleteIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.DeleteIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenantOAuthIdpConfig(ctx, ProtoToTenantOAuthIdpConfig(request.GetResource()))

}

// ListIdentitytoolkitAlphaTenantOAuthIdpConfig handles the gRPC request by passing it to the underlying TenantOAuthIdpConfigList() method.
func (s *TenantOAuthIdpConfigServer) ListIdentitytoolkitAlphaTenantOAuthIdpConfig(ctx context.Context, request *alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigRequest) (*alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigResponse, error) {
	cl, err := createConfigTenantOAuthIdpConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenantOAuthIdpConfig(ctx, request.GetProject(), request.GetTenant())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IdentitytoolkitAlphaTenantOAuthIdpConfig
	for _, r := range resources.Items {
		rp := TenantOAuthIdpConfigToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIdentitytoolkitAlphaTenantOAuthIdpConfigResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTenantOAuthIdpConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
