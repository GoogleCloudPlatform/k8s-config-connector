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
	identitytoolkitpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/identitytoolkit_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
)

// TenantServer implements the gRPC interface for Tenant.
type TenantServer struct{}

// ProtoToTenantMfaConfigStateEnum converts a TenantMfaConfigStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitTenantMfaConfigStateEnum(e identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum) *identitytoolkit.TenantMfaConfigStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum_name[int32(e)]; ok {
		e := identitytoolkit.TenantMfaConfigStateEnum(n[len("IdentitytoolkitTenantMfaConfigStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfigEnabledProvidersEnum converts a TenantMfaConfigEnabledProvidersEnum enum from its proto representation.
func ProtoToIdentitytoolkitTenantMfaConfigEnabledProvidersEnum(e identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum) *identitytoolkit.TenantMfaConfigEnabledProvidersEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum_name[int32(e)]; ok {
		e := identitytoolkit.TenantMfaConfigEnabledProvidersEnum(n[len("IdentitytoolkitTenantMfaConfigEnabledProvidersEnum"):])
		return &e
	}
	return nil
}

// ProtoToTenantMfaConfig converts a TenantMfaConfig object from its proto representation.
func ProtoToIdentitytoolkitTenantMfaConfig(p *identitytoolkitpb.IdentitytoolkitTenantMfaConfig) *identitytoolkit.TenantMfaConfig {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.TenantMfaConfig{
		State: ProtoToIdentitytoolkitTenantMfaConfigStateEnum(p.GetState()),
	}
	for _, r := range p.GetEnabledProviders() {
		obj.EnabledProviders = append(obj.EnabledProviders, *ProtoToIdentitytoolkitTenantMfaConfigEnabledProvidersEnum(r))
	}
	return obj
}

// ProtoToTenant converts a Tenant resource from its proto representation.
func ProtoToTenant(p *identitytoolkitpb.IdentitytoolkitTenant) *identitytoolkit.Tenant {
	obj := &identitytoolkit.Tenant{
		Name:                  dcl.StringOrNil(p.GetName()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		AllowPasswordSignup:   dcl.Bool(p.GetAllowPasswordSignup()),
		EnableEmailLinkSignin: dcl.Bool(p.GetEnableEmailLinkSignin()),
		DisableAuth:           dcl.Bool(p.GetDisableAuth()),
		EnableAnonymousUser:   dcl.Bool(p.GetEnableAnonymousUser()),
		MfaConfig:             ProtoToIdentitytoolkitTenantMfaConfig(p.GetMfaConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// TenantMfaConfigStateEnumToProto converts a TenantMfaConfigStateEnum enum to its proto representation.
func IdentitytoolkitTenantMfaConfigStateEnumToProto(e *identitytoolkit.TenantMfaConfigStateEnum) identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum_value["TenantMfaConfigStateEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitTenantMfaConfigStateEnum(0)
}

// TenantMfaConfigEnabledProvidersEnumToProto converts a TenantMfaConfigEnabledProvidersEnum enum to its proto representation.
func IdentitytoolkitTenantMfaConfigEnabledProvidersEnumToProto(e *identitytoolkit.TenantMfaConfigEnabledProvidersEnum) identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum_value["TenantMfaConfigEnabledProvidersEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum(0)
}

// TenantMfaConfigToProto converts a TenantMfaConfig object to its proto representation.
func IdentitytoolkitTenantMfaConfigToProto(o *identitytoolkit.TenantMfaConfig) *identitytoolkitpb.IdentitytoolkitTenantMfaConfig {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitTenantMfaConfig{}
	p.SetState(IdentitytoolkitTenantMfaConfigStateEnumToProto(o.State))
	sEnabledProviders := make([]identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum, len(o.EnabledProviders))
	for i, r := range o.EnabledProviders {
		sEnabledProviders[i] = identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum(identitytoolkitpb.IdentitytoolkitTenantMfaConfigEnabledProvidersEnum_value[string(r)])
	}
	p.SetEnabledProviders(sEnabledProviders)
	return p
}

// TenantToProto converts a Tenant resource to its proto representation.
func TenantToProto(resource *identitytoolkit.Tenant) *identitytoolkitpb.IdentitytoolkitTenant {
	p := &identitytoolkitpb.IdentitytoolkitTenant{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetAllowPasswordSignup(dcl.ValueOrEmptyBool(resource.AllowPasswordSignup))
	p.SetEnableEmailLinkSignin(dcl.ValueOrEmptyBool(resource.EnableEmailLinkSignin))
	p.SetDisableAuth(dcl.ValueOrEmptyBool(resource.DisableAuth))
	p.SetEnableAnonymousUser(dcl.ValueOrEmptyBool(resource.EnableAnonymousUser))
	p.SetMfaConfig(IdentitytoolkitTenantMfaConfigToProto(resource.MfaConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mTestPhoneNumbers := make(map[string]string, len(resource.TestPhoneNumbers))
	for k, r := range resource.TestPhoneNumbers {
		mTestPhoneNumbers[k] = r
	}
	p.SetTestPhoneNumbers(mTestPhoneNumbers)

	return p
}

// applyTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) applyTenant(ctx context.Context, c *identitytoolkit.Client, request *identitytoolkitpb.ApplyIdentitytoolkitTenantRequest) (*identitytoolkitpb.IdentitytoolkitTenant, error) {
	p := ProtoToTenant(request.GetResource())
	res, err := c.ApplyTenant(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TenantToProto(res)
	return r, nil
}

// applyIdentitytoolkitTenant handles the gRPC request by passing it to the underlying Tenant Apply() method.
func (s *TenantServer) ApplyIdentitytoolkitTenant(ctx context.Context, request *identitytoolkitpb.ApplyIdentitytoolkitTenantRequest) (*identitytoolkitpb.IdentitytoolkitTenant, error) {
	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTenant(ctx, cl, request)
}

// DeleteTenant handles the gRPC request by passing it to the underlying Tenant Delete() method.
func (s *TenantServer) DeleteIdentitytoolkitTenant(ctx context.Context, request *identitytoolkitpb.DeleteIdentitytoolkitTenantRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTenant(ctx, ProtoToTenant(request.GetResource()))

}

// ListIdentitytoolkitTenant handles the gRPC request by passing it to the underlying TenantList() method.
func (s *TenantServer) ListIdentitytoolkitTenant(ctx context.Context, request *identitytoolkitpb.ListIdentitytoolkitTenantRequest) (*identitytoolkitpb.ListIdentitytoolkitTenantResponse, error) {
	cl, err := createConfigTenant(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTenant(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*identitytoolkitpb.IdentitytoolkitTenant
	for _, r := range resources.Items {
		rp := TenantToProto(r)
		protos = append(protos, rp)
	}
	p := &identitytoolkitpb.ListIdentitytoolkitTenantResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTenant(ctx context.Context, service_account_file string) (*identitytoolkit.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return identitytoolkit.NewClient(conf), nil
}
