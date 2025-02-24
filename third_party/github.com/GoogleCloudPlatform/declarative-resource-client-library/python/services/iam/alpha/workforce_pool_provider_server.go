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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/alpha/iam_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
)

// WorkforcePoolProviderServer implements the gRPC interface for WorkforcePoolProvider.
type WorkforcePoolProviderServer struct{}

// ProtoToWorkforcePoolProviderStateEnum converts a WorkforcePoolProviderStateEnum enum from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderStateEnum(e alphapb.IamAlphaWorkforcePoolProviderStateEnum) *alpha.WorkforcePoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkforcePoolProviderStateEnum_name[int32(e)]; ok {
		e := alpha.WorkforcePoolProviderStateEnum(n[len("IamAlphaWorkforcePoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum converts a WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum enum from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(e alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum) *alpha.WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(n[len("IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum converts a WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum enum from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(e alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum) *alpha.WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum_name[int32(e)]; ok {
		e := alpha.WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(n[len("IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePoolProviderSaml converts a WorkforcePoolProviderSaml object from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderSaml(p *alphapb.IamAlphaWorkforcePoolProviderSaml) *alpha.WorkforcePoolProviderSaml {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkforcePoolProviderSaml{
		IdpMetadataXml: dcl.StringOrNil(p.GetIdpMetadataXml()),
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidc converts a WorkforcePoolProviderOidc object from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidc(p *alphapb.IamAlphaWorkforcePoolProviderOidc) *alpha.WorkforcePoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkforcePoolProviderOidc{
		IssuerUri:    dcl.StringOrNil(p.GetIssuerUri()),
		ClientId:     dcl.StringOrNil(p.GetClientId()),
		JwksJson:     dcl.StringOrNil(p.GetJwksJson()),
		WebSsoConfig: ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfig(p.GetWebSsoConfig()),
		ClientSecret: ProtoToIamAlphaWorkforcePoolProviderOidcClientSecret(p.GetClientSecret()),
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidcWebSsoConfig converts a WorkforcePoolProviderOidcWebSsoConfig object from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfig(p *alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfig) *alpha.WorkforcePoolProviderOidcWebSsoConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkforcePoolProviderOidcWebSsoConfig{
		ResponseType:            ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(p.GetResponseType()),
		AssertionClaimsBehavior: ProtoToIamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(p.GetAssertionClaimsBehavior()),
	}
	for _, r := range p.GetAdditionalScopes() {
		obj.AdditionalScopes = append(obj.AdditionalScopes, r)
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidcClientSecret converts a WorkforcePoolProviderOidcClientSecret object from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidcClientSecret(p *alphapb.IamAlphaWorkforcePoolProviderOidcClientSecret) *alpha.WorkforcePoolProviderOidcClientSecret {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkforcePoolProviderOidcClientSecret{
		Value: ProtoToIamAlphaWorkforcePoolProviderOidcClientSecretValue(p.GetValue()),
	}
	return obj
}

// ProtoToWorkforcePoolProviderOidcClientSecretValue converts a WorkforcePoolProviderOidcClientSecretValue object from its proto representation.
func ProtoToIamAlphaWorkforcePoolProviderOidcClientSecretValue(p *alphapb.IamAlphaWorkforcePoolProviderOidcClientSecretValue) *alpha.WorkforcePoolProviderOidcClientSecretValue {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkforcePoolProviderOidcClientSecretValue{
		PlainText:  dcl.StringOrNil(p.GetPlainText()),
		Thumbprint: dcl.StringOrNil(p.GetThumbprint()),
	}
	return obj
}

// ProtoToWorkforcePoolProvider converts a WorkforcePoolProvider resource from its proto representation.
func ProtoToWorkforcePoolProvider(p *alphapb.IamAlphaWorkforcePoolProvider) *alpha.WorkforcePoolProvider {
	obj := &alpha.WorkforcePoolProvider{
		Name:               dcl.StringOrNil(p.GetName()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToIamAlphaWorkforcePoolProviderStateEnum(p.GetState()),
		Disabled:           dcl.Bool(p.GetDisabled()),
		AttributeCondition: dcl.StringOrNil(p.GetAttributeCondition()),
		Saml:               ProtoToIamAlphaWorkforcePoolProviderSaml(p.GetSaml()),
		Oidc:               ProtoToIamAlphaWorkforcePoolProviderOidc(p.GetOidc()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		WorkforcePool:      dcl.StringOrNil(p.GetWorkforcePool()),
	}
	return obj
}

// WorkforcePoolProviderStateEnumToProto converts a WorkforcePoolProviderStateEnum enum to its proto representation.
func IamAlphaWorkforcePoolProviderStateEnumToProto(e *alpha.WorkforcePoolProviderStateEnum) alphapb.IamAlphaWorkforcePoolProviderStateEnum {
	if e == nil {
		return alphapb.IamAlphaWorkforcePoolProviderStateEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkforcePoolProviderStateEnum_value["WorkforcePoolProviderStateEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkforcePoolProviderStateEnum(v)
	}
	return alphapb.IamAlphaWorkforcePoolProviderStateEnum(0)
}

// WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnumToProto converts a WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum enum to its proto representation.
func IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnumToProto(e *alpha.WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum) alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum {
	if e == nil {
		return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum_value["WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(v)
	}
	return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum(0)
}

// WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnumToProto converts a WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum enum to its proto representation.
func IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnumToProto(e *alpha.WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum) alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum {
	if e == nil {
		return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum_value["WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(v)
	}
	return alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum(0)
}

// WorkforcePoolProviderSamlToProto converts a WorkforcePoolProviderSaml object to its proto representation.
func IamAlphaWorkforcePoolProviderSamlToProto(o *alpha.WorkforcePoolProviderSaml) *alphapb.IamAlphaWorkforcePoolProviderSaml {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkforcePoolProviderSaml{}
	p.SetIdpMetadataXml(dcl.ValueOrEmptyString(o.IdpMetadataXml))
	return p
}

// WorkforcePoolProviderOidcToProto converts a WorkforcePoolProviderOidc object to its proto representation.
func IamAlphaWorkforcePoolProviderOidcToProto(o *alpha.WorkforcePoolProviderOidc) *alphapb.IamAlphaWorkforcePoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkforcePoolProviderOidc{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	p.SetClientId(dcl.ValueOrEmptyString(o.ClientId))
	p.SetJwksJson(dcl.ValueOrEmptyString(o.JwksJson))
	p.SetWebSsoConfig(IamAlphaWorkforcePoolProviderOidcWebSsoConfigToProto(o.WebSsoConfig))
	p.SetClientSecret(IamAlphaWorkforcePoolProviderOidcClientSecretToProto(o.ClientSecret))
	return p
}

// WorkforcePoolProviderOidcWebSsoConfigToProto converts a WorkforcePoolProviderOidcWebSsoConfig object to its proto representation.
func IamAlphaWorkforcePoolProviderOidcWebSsoConfigToProto(o *alpha.WorkforcePoolProviderOidcWebSsoConfig) *alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkforcePoolProviderOidcWebSsoConfig{}
	p.SetResponseType(IamAlphaWorkforcePoolProviderOidcWebSsoConfigResponseTypeEnumToProto(o.ResponseType))
	p.SetAssertionClaimsBehavior(IamAlphaWorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnumToProto(o.AssertionClaimsBehavior))
	sAdditionalScopes := make([]string, len(o.AdditionalScopes))
	for i, r := range o.AdditionalScopes {
		sAdditionalScopes[i] = r
	}
	p.SetAdditionalScopes(sAdditionalScopes)
	return p
}

// WorkforcePoolProviderOidcClientSecretToProto converts a WorkforcePoolProviderOidcClientSecret object to its proto representation.
func IamAlphaWorkforcePoolProviderOidcClientSecretToProto(o *alpha.WorkforcePoolProviderOidcClientSecret) *alphapb.IamAlphaWorkforcePoolProviderOidcClientSecret {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkforcePoolProviderOidcClientSecret{}
	p.SetValue(IamAlphaWorkforcePoolProviderOidcClientSecretValueToProto(o.Value))
	return p
}

// WorkforcePoolProviderOidcClientSecretValueToProto converts a WorkforcePoolProviderOidcClientSecretValue object to its proto representation.
func IamAlphaWorkforcePoolProviderOidcClientSecretValueToProto(o *alpha.WorkforcePoolProviderOidcClientSecretValue) *alphapb.IamAlphaWorkforcePoolProviderOidcClientSecretValue {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkforcePoolProviderOidcClientSecretValue{}
	p.SetPlainText(dcl.ValueOrEmptyString(o.PlainText))
	p.SetThumbprint(dcl.ValueOrEmptyString(o.Thumbprint))
	return p
}

// WorkforcePoolProviderToProto converts a WorkforcePoolProvider resource to its proto representation.
func WorkforcePoolProviderToProto(resource *alpha.WorkforcePoolProvider) *alphapb.IamAlphaWorkforcePoolProvider {
	p := &alphapb.IamAlphaWorkforcePoolProvider{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamAlphaWorkforcePoolProviderStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetAttributeCondition(dcl.ValueOrEmptyString(resource.AttributeCondition))
	p.SetSaml(IamAlphaWorkforcePoolProviderSamlToProto(resource.Saml))
	p.SetOidc(IamAlphaWorkforcePoolProviderOidcToProto(resource.Oidc))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetWorkforcePool(dcl.ValueOrEmptyString(resource.WorkforcePool))
	mAttributeMapping := make(map[string]string, len(resource.AttributeMapping))
	for k, r := range resource.AttributeMapping {
		mAttributeMapping[k] = r
	}
	p.SetAttributeMapping(mAttributeMapping)

	return p
}

// applyWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Apply() method.
func (s *WorkforcePoolProviderServer) applyWorkforcePoolProvider(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaWorkforcePoolProviderRequest) (*alphapb.IamAlphaWorkforcePoolProvider, error) {
	p := ProtoToWorkforcePoolProvider(request.GetResource())
	res, err := c.ApplyWorkforcePoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolProviderToProto(res)
	return r, nil
}

// applyIamAlphaWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Apply() method.
func (s *WorkforcePoolProviderServer) ApplyIamAlphaWorkforcePoolProvider(ctx context.Context, request *alphapb.ApplyIamAlphaWorkforcePoolProviderRequest) (*alphapb.IamAlphaWorkforcePoolProvider, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePoolProvider(ctx, cl, request)
}

// DeleteWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProvider Delete() method.
func (s *WorkforcePoolProviderServer) DeleteIamAlphaWorkforcePoolProvider(ctx context.Context, request *alphapb.DeleteIamAlphaWorkforcePoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePoolProvider(ctx, ProtoToWorkforcePoolProvider(request.GetResource()))

}

// ListIamAlphaWorkforcePoolProvider handles the gRPC request by passing it to the underlying WorkforcePoolProviderList() method.
func (s *WorkforcePoolProviderServer) ListIamAlphaWorkforcePoolProvider(ctx context.Context, request *alphapb.ListIamAlphaWorkforcePoolProviderRequest) (*alphapb.ListIamAlphaWorkforcePoolProviderResponse, error) {
	cl, err := createConfigWorkforcePoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePoolProvider(ctx, request.GetLocation(), request.GetWorkforcePool())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaWorkforcePoolProvider
	for _, r := range resources.Items {
		rp := WorkforcePoolProviderToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIamAlphaWorkforcePoolProviderResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePoolProvider(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
