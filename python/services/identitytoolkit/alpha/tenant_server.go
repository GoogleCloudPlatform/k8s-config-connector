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

// TenantServer implements the gRPC interface for Tenant.
type TenantServer struct{}

// ProtoToTenantMfaConfigStateEnum converts a TenantMfaConfigStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaTenantMfaConfigStateEnum(e alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum) *alpha.TenantMfaConfigStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum_name[int32(e)]; ok {
		e := alpha.TenantMfaConfigStateEnum(n[len("IdentitytoolkitAlphaTenantMfaConfigStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfigEnabledProvidersEnum converts a TenantMfaConfigEnabledProvidersEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(e alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum) *alpha.TenantMfaConfigEnabledProvidersEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum_name[int32(e)]; ok {
		e := alpha.TenantMfaConfigEnabledProvidersEnum(n[len("IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfig converts a TenantMfaConfig object from its proto representation.
func ProtoToIdentitytoolkitAlphaTenantMfaConfig(p *alphapb.IdentitytoolkitAlphaTenantMfaConfig) *alpha.TenantMfaConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.TenantMfaConfig{
		State: ProtoToIdentitytoolkitAlphaTenantMfaConfigStateEnum(p.GetState()),
	}
	for _, r := range p.GetEnabledProviders() {
		obj.EnabledProviders = append(obj.EnabledProviders, *ProtoToIdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(r))
	}
	return obj
}

// ProtoToTenant converts a Tenant resource from its proto representation.
func ProtoToTenant(p *alphapb.IdentitytoolkitAlphaTenant) *alpha.Tenant {
	obj := &alpha.Tenant{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		AllowPasswordSignup:   dcl.Bool(p.GetAllowPasswordSignup()),
		EnableEmailLinkSignin: dcl.Bool(p.GetEnableEmailLinkSignin()),
		DisableAuth:           dcl.Bool(p.GetDisableAuth()),
		EnableAnonymousUser:   dcl.Bool(p.GetEnableAnonymousUser()),
		MfaConfig:             ProtoToIdentitytoolkitAlphaTenantMfaConfig(p.GetMfaConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// TenantMfaConfigStateEnumToProto converts a TenantMfaConfigStateEnum enum to its proto representation.
func IdentitytoolkitAlphaTenantMfaConfigStateEnumToProto(e *alpha.TenantMfaConfigStateEnum) alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum_value["TenantMfaConfigStateEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaTenantMfaConfigStateEnum(0)
}

// TenantMfaConfigEnabledProvidersEnumToProto converts a TenantMfaConfigEnabledProvidersEnum enum to its proto representation.
func IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnumToProto(e *alpha.TenantMfaConfigEnabledProvidersEnum) alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum_value["TenantMfaConfigEnabledProvidersEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(0)
}

// TenantMfaConfigToProto converts a TenantMfaConfig object to its proto representation.
func IdentitytoolkitAlphaTenantMfaConfigToProto(o *alpha.TenantMfaConfig) *alphapb.IdentitytoolkitAlphaTenantMfaConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaTenantMfaConfig{}
	p.SetState(IdentitytoolkitAlphaTenantMfaConfigStateEnumToProto(o.State))
	sEnabledProviders := make([]alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum, len(o.EnabledProviders))
	for i, r := range o.EnabledProviders {
		sEnabledProviders[i] = alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum(alphapb.IdentitytoolkitAlphaTenantMfaConfigEnabledProvidersEnum_value[string(r)])
	}
	p.SetEnabledProviders(sEnabledProviders)
	return p
}

// TenantToProto converts a Tenant resource to its proto representation.
func TenantToProto(resource *alpha.Tenant) *alphapb.IdentitytoolkitAlphaTenant {
	p := &alphapb.IdentitytoolkitAlphaTenant{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetAllowPasswordSignup(dcl.ValueOrEmptyBool(resource.AllowPasswordSignup))
	p.SetEnableEmailLinkSignin(dcl.ValueOrEmptyBool(resource.EnableEmailLinkSignin))
	p.SetDisableAuth(dcl.ValueOrEmptyBool(resource.DisableAuth))
	p.SetEnableAnonymousUser(dcl.ValueOrEmptyBool(resource.EnableAnonymousUser))
	p.SetMfaConfig(IdentitytoolkitAlphaTenantMfaConfigToProto(resource.MfaConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mTestPhoneNumbers := make(map[string]string, len(resource.TestPhoneNumbers))
	for k, r := range resource.TestPhoneNumbers {
		mTestPhoneNumbers[k] = r
	}
	p.SetTestPhoneNumbers(mTestPhoneNumbers)

	return p
}

// applyTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) applyTenant(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIdentitytoolkitAlphaTenantRequest) (*alphapb.IdentitytoolkitAlphaTenant, error) {
	p := ProtoToTenant(request.GetResource())
	res, err := c.ApplyTenant(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantToProto(res)
	return r, nil
}

// applyIdentitytoolkitAlphaTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) ApplyIdentitytoolkitAlphaTenant(ctx context.Context, request *alphapb.ApplyIdentitytoolkitAlphaTenantRequest) (*alphapb.IdentitytoolkitAlphaTenant, error) {
	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTenant(ctx, cl, request)
}

// DeleteTenant handles the gRPC request by passing it to the underlying Tenant Delete() method.
func (s *TenantServer) DeleteIdentitytoolkitAlphaTenant(ctx context.Context, request *alphapb.DeleteIdentitytoolkitAlphaTenantRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenant(ctx, ProtoToTenant(request.GetResource()))

}

// ListIdentitytoolkitAlphaTenant handles the gRPC request by passing it to the underlying TenantList() method.
func (s *TenantServer) ListIdentitytoolkitAlphaTenant(ctx context.Context, request *alphapb.ListIdentitytoolkitAlphaTenantRequest) (*alphapb.ListIdentitytoolkitAlphaTenantResponse, error) {
	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenant(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IdentitytoolkitAlphaTenant
	for _, r := range resources.Items {
		rp := TenantToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIdentitytoolkitAlphaTenantResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTenant(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
