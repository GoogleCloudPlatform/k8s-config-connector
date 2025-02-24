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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/beta/apigee_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
)

// OrganizationServer implements the gRPC interface for Organization.
type OrganizationServer struct{}

// ProtoToOrganizationRuntimeTypeEnum converts a OrganizationRuntimeTypeEnum enum from its proto representation.
func ProtoToApigeeBetaOrganizationRuntimeTypeEnum(e betapb.ApigeeBetaOrganizationRuntimeTypeEnum) *beta.OrganizationRuntimeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaOrganizationRuntimeTypeEnum_name[int32(e)]; ok {
		e := beta.OrganizationRuntimeTypeEnum(n[len("ApigeeBetaOrganizationRuntimeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationSubscriptionTypeEnum converts a OrganizationSubscriptionTypeEnum enum from its proto representation.
func ProtoToApigeeBetaOrganizationSubscriptionTypeEnum(e betapb.ApigeeBetaOrganizationSubscriptionTypeEnum) *beta.OrganizationSubscriptionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaOrganizationSubscriptionTypeEnum_name[int32(e)]; ok {
		e := beta.OrganizationSubscriptionTypeEnum(n[len("ApigeeBetaOrganizationSubscriptionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationBillingTypeEnum converts a OrganizationBillingTypeEnum enum from its proto representation.
func ProtoToApigeeBetaOrganizationBillingTypeEnum(e betapb.ApigeeBetaOrganizationBillingTypeEnum) *beta.OrganizationBillingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaOrganizationBillingTypeEnum_name[int32(e)]; ok {
		e := beta.OrganizationBillingTypeEnum(n[len("ApigeeBetaOrganizationBillingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationStateEnum converts a OrganizationStateEnum enum from its proto representation.
func ProtoToApigeeBetaOrganizationStateEnum(e betapb.ApigeeBetaOrganizationStateEnum) *beta.OrganizationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaOrganizationStateEnum_name[int32(e)]; ok {
		e := beta.OrganizationStateEnum(n[len("ApigeeBetaOrganizationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationAddonsConfig converts a OrganizationAddonsConfig object from its proto representation.
func ProtoToApigeeBetaOrganizationAddonsConfig(p *betapb.ApigeeBetaOrganizationAddonsConfig) *beta.OrganizationAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.OrganizationAddonsConfig{
		AdvancedApiOpsConfig: ProtoToApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfig(p.GetAdvancedApiOpsConfig()),
		MonetizationConfig:   ProtoToApigeeBetaOrganizationAddonsConfigMonetizationConfig(p.GetMonetizationConfig()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigAdvancedApiOpsConfig converts a OrganizationAddonsConfigAdvancedApiOpsConfig object from its proto representation.
func ProtoToApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfig(p *betapb.ApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfig) *beta.OrganizationAddonsConfigAdvancedApiOpsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.OrganizationAddonsConfigAdvancedApiOpsConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigMonetizationConfig converts a OrganizationAddonsConfigMonetizationConfig object from its proto representation.
func ProtoToApigeeBetaOrganizationAddonsConfigMonetizationConfig(p *betapb.ApigeeBetaOrganizationAddonsConfigMonetizationConfig) *beta.OrganizationAddonsConfigMonetizationConfig {
	if p == nil {
		return nil
	}
	obj := &beta.OrganizationAddonsConfigMonetizationConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganization converts a Organization resource from its proto representation.
func ProtoToOrganization(p *betapb.ApigeeBetaOrganization) *beta.Organization {
	obj := &beta.Organization{
		Name:                             dcl.StringOrNil(p.GetName()),
		DisplayName:                      dcl.StringOrNil(p.GetDisplayName()),
		Description:                      dcl.StringOrNil(p.GetDescription()),
		CreatedAt:                        dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:                   dcl.Int64OrNil(p.GetLastModifiedAt()),
		ExpiresAt:                        dcl.Int64OrNil(p.GetExpiresAt()),
		AnalyticsRegion:                  dcl.StringOrNil(p.GetAnalyticsRegion()),
		AuthorizedNetwork:                dcl.StringOrNil(p.GetAuthorizedNetwork()),
		RuntimeType:                      ProtoToApigeeBetaOrganizationRuntimeTypeEnum(p.GetRuntimeType()),
		SubscriptionType:                 ProtoToApigeeBetaOrganizationSubscriptionTypeEnum(p.GetSubscriptionType()),
		BillingType:                      ProtoToApigeeBetaOrganizationBillingTypeEnum(p.GetBillingType()),
		AddonsConfig:                     ProtoToApigeeBetaOrganizationAddonsConfig(p.GetAddonsConfig()),
		CaCertificate:                    dcl.StringOrNil(p.GetCaCertificate()),
		RuntimeDatabaseEncryptionKeyName: dcl.StringOrNil(p.GetRuntimeDatabaseEncryptionKeyName()),
		ProjectId:                        dcl.StringOrNil(p.GetProjectId()),
		State:                            ProtoToApigeeBetaOrganizationStateEnum(p.GetState()),
		Project:                          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetEnvironments() {
		obj.Environments = append(obj.Environments, r)
	}
	return obj
}

// OrganizationRuntimeTypeEnumToProto converts a OrganizationRuntimeTypeEnum enum to its proto representation.
func ApigeeBetaOrganizationRuntimeTypeEnumToProto(e *beta.OrganizationRuntimeTypeEnum) betapb.ApigeeBetaOrganizationRuntimeTypeEnum {
	if e == nil {
		return betapb.ApigeeBetaOrganizationRuntimeTypeEnum(0)
	}
	if v, ok := betapb.ApigeeBetaOrganizationRuntimeTypeEnum_value["OrganizationRuntimeTypeEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaOrganizationRuntimeTypeEnum(v)
	}
	return betapb.ApigeeBetaOrganizationRuntimeTypeEnum(0)
}

// OrganizationSubscriptionTypeEnumToProto converts a OrganizationSubscriptionTypeEnum enum to its proto representation.
func ApigeeBetaOrganizationSubscriptionTypeEnumToProto(e *beta.OrganizationSubscriptionTypeEnum) betapb.ApigeeBetaOrganizationSubscriptionTypeEnum {
	if e == nil {
		return betapb.ApigeeBetaOrganizationSubscriptionTypeEnum(0)
	}
	if v, ok := betapb.ApigeeBetaOrganizationSubscriptionTypeEnum_value["OrganizationSubscriptionTypeEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaOrganizationSubscriptionTypeEnum(v)
	}
	return betapb.ApigeeBetaOrganizationSubscriptionTypeEnum(0)
}

// OrganizationBillingTypeEnumToProto converts a OrganizationBillingTypeEnum enum to its proto representation.
func ApigeeBetaOrganizationBillingTypeEnumToProto(e *beta.OrganizationBillingTypeEnum) betapb.ApigeeBetaOrganizationBillingTypeEnum {
	if e == nil {
		return betapb.ApigeeBetaOrganizationBillingTypeEnum(0)
	}
	if v, ok := betapb.ApigeeBetaOrganizationBillingTypeEnum_value["OrganizationBillingTypeEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaOrganizationBillingTypeEnum(v)
	}
	return betapb.ApigeeBetaOrganizationBillingTypeEnum(0)
}

// OrganizationStateEnumToProto converts a OrganizationStateEnum enum to its proto representation.
func ApigeeBetaOrganizationStateEnumToProto(e *beta.OrganizationStateEnum) betapb.ApigeeBetaOrganizationStateEnum {
	if e == nil {
		return betapb.ApigeeBetaOrganizationStateEnum(0)
	}
	if v, ok := betapb.ApigeeBetaOrganizationStateEnum_value["OrganizationStateEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaOrganizationStateEnum(v)
	}
	return betapb.ApigeeBetaOrganizationStateEnum(0)
}

// OrganizationAddonsConfigToProto converts a OrganizationAddonsConfig object to its proto representation.
func ApigeeBetaOrganizationAddonsConfigToProto(o *beta.OrganizationAddonsConfig) *betapb.ApigeeBetaOrganizationAddonsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ApigeeBetaOrganizationAddonsConfig{}
	p.SetAdvancedApiOpsConfig(ApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o.AdvancedApiOpsConfig))
	p.SetMonetizationConfig(ApigeeBetaOrganizationAddonsConfigMonetizationConfigToProto(o.MonetizationConfig))
	return p
}

// OrganizationAddonsConfigAdvancedApiOpsConfigToProto converts a OrganizationAddonsConfigAdvancedApiOpsConfig object to its proto representation.
func ApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o *beta.OrganizationAddonsConfigAdvancedApiOpsConfig) *betapb.ApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ApigeeBetaOrganizationAddonsConfigAdvancedApiOpsConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationAddonsConfigMonetizationConfigToProto converts a OrganizationAddonsConfigMonetizationConfig object to its proto representation.
func ApigeeBetaOrganizationAddonsConfigMonetizationConfigToProto(o *beta.OrganizationAddonsConfigMonetizationConfig) *betapb.ApigeeBetaOrganizationAddonsConfigMonetizationConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ApigeeBetaOrganizationAddonsConfigMonetizationConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationToProto converts a Organization resource to its proto representation.
func OrganizationToProto(resource *beta.Organization) *betapb.ApigeeBetaOrganization {
	p := &betapb.ApigeeBetaOrganization{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetExpiresAt(dcl.ValueOrEmptyInt64(resource.ExpiresAt))
	p.SetAnalyticsRegion(dcl.ValueOrEmptyString(resource.AnalyticsRegion))
	p.SetAuthorizedNetwork(dcl.ValueOrEmptyString(resource.AuthorizedNetwork))
	p.SetRuntimeType(ApigeeBetaOrganizationRuntimeTypeEnumToProto(resource.RuntimeType))
	p.SetSubscriptionType(ApigeeBetaOrganizationSubscriptionTypeEnumToProto(resource.SubscriptionType))
	p.SetBillingType(ApigeeBetaOrganizationBillingTypeEnumToProto(resource.BillingType))
	p.SetAddonsConfig(ApigeeBetaOrganizationAddonsConfigToProto(resource.AddonsConfig))
	p.SetCaCertificate(dcl.ValueOrEmptyString(resource.CaCertificate))
	p.SetRuntimeDatabaseEncryptionKeyName(dcl.ValueOrEmptyString(resource.RuntimeDatabaseEncryptionKeyName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetState(ApigeeBetaOrganizationStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sEnvironments := make([]string, len(resource.Environments))
	for i, r := range resource.Environments {
		sEnvironments[i] = r
	}
	p.SetEnvironments(sEnvironments)
	mProperties := make(map[string]string, len(resource.Properties))
	for k, r := range resource.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)

	return p
}

// applyOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) applyOrganization(ctx context.Context, c *beta.Client, request *betapb.ApplyApigeeBetaOrganizationRequest) (*betapb.ApigeeBetaOrganization, error) {
	p := ProtoToOrganization(request.GetResource())
	res, err := c.ApplyOrganization(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OrganizationToProto(res)
	return r, nil
}

// applyApigeeBetaOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) ApplyApigeeBetaOrganization(ctx context.Context, request *betapb.ApplyApigeeBetaOrganizationRequest) (*betapb.ApigeeBetaOrganization, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOrganization(ctx, cl, request)
}

// DeleteOrganization handles the gRPC request by passing it to the underlying Organization Delete() method.
func (s *OrganizationServer) DeleteApigeeBetaOrganization(ctx context.Context, request *betapb.DeleteApigeeBetaOrganizationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOrganization(ctx, ProtoToOrganization(request.GetResource()))

}

// ListApigeeBetaOrganization handles the gRPC request by passing it to the underlying OrganizationList() method.
func (s *OrganizationServer) ListApigeeBetaOrganization(ctx context.Context, request *betapb.ListApigeeBetaOrganizationRequest) (*betapb.ListApigeeBetaOrganizationResponse, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOrganization(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ApigeeBetaOrganization
	for _, r := range resources.Items {
		rp := OrganizationToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListApigeeBetaOrganizationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOrganization(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
