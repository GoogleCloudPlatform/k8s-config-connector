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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/alpha/apigee_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha"
)

// OrganizationServer implements the gRPC interface for Organization.
type OrganizationServer struct{}

// ProtoToOrganizationRuntimeTypeEnum converts a OrganizationRuntimeTypeEnum enum from its proto representation.
func ProtoToApigeeAlphaOrganizationRuntimeTypeEnum(e alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum) *alpha.OrganizationRuntimeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum_name[int32(e)]; ok {
		e := alpha.OrganizationRuntimeTypeEnum(n[len("ApigeeAlphaOrganizationRuntimeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationSubscriptionTypeEnum converts a OrganizationSubscriptionTypeEnum enum from its proto representation.
func ProtoToApigeeAlphaOrganizationSubscriptionTypeEnum(e alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum) *alpha.OrganizationSubscriptionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum_name[int32(e)]; ok {
		e := alpha.OrganizationSubscriptionTypeEnum(n[len("ApigeeAlphaOrganizationSubscriptionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationBillingTypeEnum converts a OrganizationBillingTypeEnum enum from its proto representation.
func ProtoToApigeeAlphaOrganizationBillingTypeEnum(e alphapb.ApigeeAlphaOrganizationBillingTypeEnum) *alpha.OrganizationBillingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaOrganizationBillingTypeEnum_name[int32(e)]; ok {
		e := alpha.OrganizationBillingTypeEnum(n[len("ApigeeAlphaOrganizationBillingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationStateEnum converts a OrganizationStateEnum enum from its proto representation.
func ProtoToApigeeAlphaOrganizationStateEnum(e alphapb.ApigeeAlphaOrganizationStateEnum) *alpha.OrganizationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaOrganizationStateEnum_name[int32(e)]; ok {
		e := alpha.OrganizationStateEnum(n[len("ApigeeAlphaOrganizationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationAddonsConfig converts a OrganizationAddonsConfig object from its proto representation.
func ProtoToApigeeAlphaOrganizationAddonsConfig(p *alphapb.ApigeeAlphaOrganizationAddonsConfig) *alpha.OrganizationAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.OrganizationAddonsConfig{
		AdvancedApiOpsConfig: ProtoToApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig(p.GetAdvancedApiOpsConfig()),
		MonetizationConfig:   ProtoToApigeeAlphaOrganizationAddonsConfigMonetizationConfig(p.GetMonetizationConfig()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigAdvancedApiOpsConfig converts a OrganizationAddonsConfigAdvancedApiOpsConfig object from its proto representation.
func ProtoToApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig(p *alphapb.ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig) *alpha.OrganizationAddonsConfigAdvancedApiOpsConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.OrganizationAddonsConfigAdvancedApiOpsConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigMonetizationConfig converts a OrganizationAddonsConfigMonetizationConfig object from its proto representation.
func ProtoToApigeeAlphaOrganizationAddonsConfigMonetizationConfig(p *alphapb.ApigeeAlphaOrganizationAddonsConfigMonetizationConfig) *alpha.OrganizationAddonsConfigMonetizationConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.OrganizationAddonsConfigMonetizationConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganization converts a Organization resource from its proto representation.
func ProtoToOrganization(p *alphapb.ApigeeAlphaOrganization) *alpha.Organization {
	obj := &alpha.Organization{
		Name:                             dcl.StringOrNil(p.GetName()),
		DisplayName:                      dcl.StringOrNil(p.GetDisplayName()),
		Description:                      dcl.StringOrNil(p.GetDescription()),
		CreatedAt:                        dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:                   dcl.Int64OrNil(p.GetLastModifiedAt()),
		ExpiresAt:                        dcl.Int64OrNil(p.GetExpiresAt()),
		AnalyticsRegion:                  dcl.StringOrNil(p.GetAnalyticsRegion()),
		AuthorizedNetwork:                dcl.StringOrNil(p.GetAuthorizedNetwork()),
		RuntimeType:                      ProtoToApigeeAlphaOrganizationRuntimeTypeEnum(p.GetRuntimeType()),
		SubscriptionType:                 ProtoToApigeeAlphaOrganizationSubscriptionTypeEnum(p.GetSubscriptionType()),
		BillingType:                      ProtoToApigeeAlphaOrganizationBillingTypeEnum(p.GetBillingType()),
		AddonsConfig:                     ProtoToApigeeAlphaOrganizationAddonsConfig(p.GetAddonsConfig()),
		CaCertificate:                    dcl.StringOrNil(p.GetCaCertificate()),
		RuntimeDatabaseEncryptionKeyName: dcl.StringOrNil(p.GetRuntimeDatabaseEncryptionKeyName()),
		ProjectId:                        dcl.StringOrNil(p.GetProjectId()),
		State:                            ProtoToApigeeAlphaOrganizationStateEnum(p.GetState()),
		Project:                          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetEnvironments() {
		obj.Environments = append(obj.Environments, r)
	}
	return obj
}

// OrganizationRuntimeTypeEnumToProto converts a OrganizationRuntimeTypeEnum enum to its proto representation.
func ApigeeAlphaOrganizationRuntimeTypeEnumToProto(e *alpha.OrganizationRuntimeTypeEnum) alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum {
	if e == nil {
		return alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum_value["OrganizationRuntimeTypeEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum(v)
	}
	return alphapb.ApigeeAlphaOrganizationRuntimeTypeEnum(0)
}

// OrganizationSubscriptionTypeEnumToProto converts a OrganizationSubscriptionTypeEnum enum to its proto representation.
func ApigeeAlphaOrganizationSubscriptionTypeEnumToProto(e *alpha.OrganizationSubscriptionTypeEnum) alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum {
	if e == nil {
		return alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum_value["OrganizationSubscriptionTypeEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum(v)
	}
	return alphapb.ApigeeAlphaOrganizationSubscriptionTypeEnum(0)
}

// OrganizationBillingTypeEnumToProto converts a OrganizationBillingTypeEnum enum to its proto representation.
func ApigeeAlphaOrganizationBillingTypeEnumToProto(e *alpha.OrganizationBillingTypeEnum) alphapb.ApigeeAlphaOrganizationBillingTypeEnum {
	if e == nil {
		return alphapb.ApigeeAlphaOrganizationBillingTypeEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaOrganizationBillingTypeEnum_value["OrganizationBillingTypeEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaOrganizationBillingTypeEnum(v)
	}
	return alphapb.ApigeeAlphaOrganizationBillingTypeEnum(0)
}

// OrganizationStateEnumToProto converts a OrganizationStateEnum enum to its proto representation.
func ApigeeAlphaOrganizationStateEnumToProto(e *alpha.OrganizationStateEnum) alphapb.ApigeeAlphaOrganizationStateEnum {
	if e == nil {
		return alphapb.ApigeeAlphaOrganizationStateEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaOrganizationStateEnum_value["OrganizationStateEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaOrganizationStateEnum(v)
	}
	return alphapb.ApigeeAlphaOrganizationStateEnum(0)
}

// OrganizationAddonsConfigToProto converts a OrganizationAddonsConfig object to its proto representation.
func ApigeeAlphaOrganizationAddonsConfigToProto(o *alpha.OrganizationAddonsConfig) *alphapb.ApigeeAlphaOrganizationAddonsConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ApigeeAlphaOrganizationAddonsConfig{}
	p.SetAdvancedApiOpsConfig(ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o.AdvancedApiOpsConfig))
	p.SetMonetizationConfig(ApigeeAlphaOrganizationAddonsConfigMonetizationConfigToProto(o.MonetizationConfig))
	return p
}

// OrganizationAddonsConfigAdvancedApiOpsConfigToProto converts a OrganizationAddonsConfigAdvancedApiOpsConfig object to its proto representation.
func ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o *alpha.OrganizationAddonsConfigAdvancedApiOpsConfig) *alphapb.ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ApigeeAlphaOrganizationAddonsConfigAdvancedApiOpsConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationAddonsConfigMonetizationConfigToProto converts a OrganizationAddonsConfigMonetizationConfig object to its proto representation.
func ApigeeAlphaOrganizationAddonsConfigMonetizationConfigToProto(o *alpha.OrganizationAddonsConfigMonetizationConfig) *alphapb.ApigeeAlphaOrganizationAddonsConfigMonetizationConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ApigeeAlphaOrganizationAddonsConfigMonetizationConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationToProto converts a Organization resource to its proto representation.
func OrganizationToProto(resource *alpha.Organization) *alphapb.ApigeeAlphaOrganization {
	p := &alphapb.ApigeeAlphaOrganization{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetExpiresAt(dcl.ValueOrEmptyInt64(resource.ExpiresAt))
	p.SetAnalyticsRegion(dcl.ValueOrEmptyString(resource.AnalyticsRegion))
	p.SetAuthorizedNetwork(dcl.ValueOrEmptyString(resource.AuthorizedNetwork))
	p.SetRuntimeType(ApigeeAlphaOrganizationRuntimeTypeEnumToProto(resource.RuntimeType))
	p.SetSubscriptionType(ApigeeAlphaOrganizationSubscriptionTypeEnumToProto(resource.SubscriptionType))
	p.SetBillingType(ApigeeAlphaOrganizationBillingTypeEnumToProto(resource.BillingType))
	p.SetAddonsConfig(ApigeeAlphaOrganizationAddonsConfigToProto(resource.AddonsConfig))
	p.SetCaCertificate(dcl.ValueOrEmptyString(resource.CaCertificate))
	p.SetRuntimeDatabaseEncryptionKeyName(dcl.ValueOrEmptyString(resource.RuntimeDatabaseEncryptionKeyName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetState(ApigeeAlphaOrganizationStateEnumToProto(resource.State))
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
func (s *OrganizationServer) applyOrganization(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaOrganizationRequest) (*alphapb.ApigeeAlphaOrganization, error) {
	p := ProtoToOrganization(request.GetResource())
	res, err := c.ApplyOrganization(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OrganizationToProto(res)
	return r, nil
}

// applyApigeeAlphaOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) ApplyApigeeAlphaOrganization(ctx context.Context, request *alphapb.ApplyApigeeAlphaOrganizationRequest) (*alphapb.ApigeeAlphaOrganization, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOrganization(ctx, cl, request)
}

// DeleteOrganization handles the gRPC request by passing it to the underlying Organization Delete() method.
func (s *OrganizationServer) DeleteApigeeAlphaOrganization(ctx context.Context, request *alphapb.DeleteApigeeAlphaOrganizationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOrganization(ctx, ProtoToOrganization(request.GetResource()))

}

// ListApigeeAlphaOrganization handles the gRPC request by passing it to the underlying OrganizationList() method.
func (s *OrganizationServer) ListApigeeAlphaOrganization(ctx context.Context, request *alphapb.ListApigeeAlphaOrganizationRequest) (*alphapb.ListApigeeAlphaOrganizationResponse, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOrganization(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaOrganization
	for _, r := range resources.Items {
		rp := OrganizationToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaOrganizationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOrganization(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
