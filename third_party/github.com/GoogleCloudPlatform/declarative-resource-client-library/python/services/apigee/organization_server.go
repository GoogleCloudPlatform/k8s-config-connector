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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// OrganizationServer implements the gRPC interface for Organization.
type OrganizationServer struct{}

// ProtoToOrganizationRuntimeTypeEnum converts a OrganizationRuntimeTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationRuntimeTypeEnum(e apigeepb.ApigeeOrganizationRuntimeTypeEnum) *apigee.OrganizationRuntimeTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationRuntimeTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationRuntimeTypeEnum(n[len("ApigeeOrganizationRuntimeTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationSubscriptionTypeEnum converts a OrganizationSubscriptionTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationSubscriptionTypeEnum(e apigeepb.ApigeeOrganizationSubscriptionTypeEnum) *apigee.OrganizationSubscriptionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationSubscriptionTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationSubscriptionTypeEnum(n[len("ApigeeOrganizationSubscriptionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationBillingTypeEnum converts a OrganizationBillingTypeEnum enum from its proto representation.
func ProtoToApigeeOrganizationBillingTypeEnum(e apigeepb.ApigeeOrganizationBillingTypeEnum) *apigee.OrganizationBillingTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationBillingTypeEnum_name[int32(e)]; ok {
		e := apigee.OrganizationBillingTypeEnum(n[len("ApigeeOrganizationBillingTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationStateEnum converts a OrganizationStateEnum enum from its proto representation.
func ProtoToApigeeOrganizationStateEnum(e apigeepb.ApigeeOrganizationStateEnum) *apigee.OrganizationStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeOrganizationStateEnum_name[int32(e)]; ok {
		e := apigee.OrganizationStateEnum(n[len("ApigeeOrganizationStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOrganizationAddonsConfig converts a OrganizationAddonsConfig object from its proto representation.
func ProtoToApigeeOrganizationAddonsConfig(p *apigeepb.ApigeeOrganizationAddonsConfig) *apigee.OrganizationAddonsConfig {
	if p == nil {
		return nil
	}
	obj := &apigee.OrganizationAddonsConfig{
		AdvancedApiOpsConfig: ProtoToApigeeOrganizationAddonsConfigAdvancedApiOpsConfig(p.GetAdvancedApiOpsConfig()),
		MonetizationConfig:   ProtoToApigeeOrganizationAddonsConfigMonetizationConfig(p.GetMonetizationConfig()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigAdvancedApiOpsConfig converts a OrganizationAddonsConfigAdvancedApiOpsConfig object from its proto representation.
func ProtoToApigeeOrganizationAddonsConfigAdvancedApiOpsConfig(p *apigeepb.ApigeeOrganizationAddonsConfigAdvancedApiOpsConfig) *apigee.OrganizationAddonsConfigAdvancedApiOpsConfig {
	if p == nil {
		return nil
	}
	obj := &apigee.OrganizationAddonsConfigAdvancedApiOpsConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganizationAddonsConfigMonetizationConfig converts a OrganizationAddonsConfigMonetizationConfig object from its proto representation.
func ProtoToApigeeOrganizationAddonsConfigMonetizationConfig(p *apigeepb.ApigeeOrganizationAddonsConfigMonetizationConfig) *apigee.OrganizationAddonsConfigMonetizationConfig {
	if p == nil {
		return nil
	}
	obj := &apigee.OrganizationAddonsConfigMonetizationConfig{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToOrganization converts a Organization resource from its proto representation.
func ProtoToOrganization(p *apigeepb.ApigeeOrganization) *apigee.Organization {
	obj := &apigee.Organization{
		Name:                             dcl.StringOrNil(p.GetName()),
		DisplayName:                      dcl.StringOrNil(p.GetDisplayName()),
		Description:                      dcl.StringOrNil(p.GetDescription()),
		CreatedAt:                        dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:                   dcl.Int64OrNil(p.GetLastModifiedAt()),
		ExpiresAt:                        dcl.Int64OrNil(p.GetExpiresAt()),
		AnalyticsRegion:                  dcl.StringOrNil(p.GetAnalyticsRegion()),
		AuthorizedNetwork:                dcl.StringOrNil(p.GetAuthorizedNetwork()),
		RuntimeType:                      ProtoToApigeeOrganizationRuntimeTypeEnum(p.GetRuntimeType()),
		SubscriptionType:                 ProtoToApigeeOrganizationSubscriptionTypeEnum(p.GetSubscriptionType()),
		BillingType:                      ProtoToApigeeOrganizationBillingTypeEnum(p.GetBillingType()),
		AddonsConfig:                     ProtoToApigeeOrganizationAddonsConfig(p.GetAddonsConfig()),
		CaCertificate:                    dcl.StringOrNil(p.GetCaCertificate()),
		RuntimeDatabaseEncryptionKeyName: dcl.StringOrNil(p.GetRuntimeDatabaseEncryptionKeyName()),
		ProjectId:                        dcl.StringOrNil(p.GetProjectId()),
		State:                            ProtoToApigeeOrganizationStateEnum(p.GetState()),
		Project:                          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetEnvironments() {
		obj.Environments = append(obj.Environments, r)
	}
	return obj
}

// OrganizationRuntimeTypeEnumToProto converts a OrganizationRuntimeTypeEnum enum to its proto representation.
func ApigeeOrganizationRuntimeTypeEnumToProto(e *apigee.OrganizationRuntimeTypeEnum) apigeepb.ApigeeOrganizationRuntimeTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationRuntimeTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationRuntimeTypeEnum_value["OrganizationRuntimeTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationRuntimeTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationRuntimeTypeEnum(0)
}

// OrganizationSubscriptionTypeEnumToProto converts a OrganizationSubscriptionTypeEnum enum to its proto representation.
func ApigeeOrganizationSubscriptionTypeEnumToProto(e *apigee.OrganizationSubscriptionTypeEnum) apigeepb.ApigeeOrganizationSubscriptionTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationSubscriptionTypeEnum_value["OrganizationSubscriptionTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationSubscriptionTypeEnum(0)
}

// OrganizationBillingTypeEnumToProto converts a OrganizationBillingTypeEnum enum to its proto representation.
func ApigeeOrganizationBillingTypeEnumToProto(e *apigee.OrganizationBillingTypeEnum) apigeepb.ApigeeOrganizationBillingTypeEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationBillingTypeEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationBillingTypeEnum_value["OrganizationBillingTypeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationBillingTypeEnum(v)
	}
	return apigeepb.ApigeeOrganizationBillingTypeEnum(0)
}

// OrganizationStateEnumToProto converts a OrganizationStateEnum enum to its proto representation.
func ApigeeOrganizationStateEnumToProto(e *apigee.OrganizationStateEnum) apigeepb.ApigeeOrganizationStateEnum {
	if e == nil {
		return apigeepb.ApigeeOrganizationStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeOrganizationStateEnum_value["OrganizationStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeOrganizationStateEnum(v)
	}
	return apigeepb.ApigeeOrganizationStateEnum(0)
}

// OrganizationAddonsConfigToProto converts a OrganizationAddonsConfig object to its proto representation.
func ApigeeOrganizationAddonsConfigToProto(o *apigee.OrganizationAddonsConfig) *apigeepb.ApigeeOrganizationAddonsConfig {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeOrganizationAddonsConfig{}
	p.SetAdvancedApiOpsConfig(ApigeeOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o.AdvancedApiOpsConfig))
	p.SetMonetizationConfig(ApigeeOrganizationAddonsConfigMonetizationConfigToProto(o.MonetizationConfig))
	return p
}

// OrganizationAddonsConfigAdvancedApiOpsConfigToProto converts a OrganizationAddonsConfigAdvancedApiOpsConfig object to its proto representation.
func ApigeeOrganizationAddonsConfigAdvancedApiOpsConfigToProto(o *apigee.OrganizationAddonsConfigAdvancedApiOpsConfig) *apigeepb.ApigeeOrganizationAddonsConfigAdvancedApiOpsConfig {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeOrganizationAddonsConfigAdvancedApiOpsConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationAddonsConfigMonetizationConfigToProto converts a OrganizationAddonsConfigMonetizationConfig object to its proto representation.
func ApigeeOrganizationAddonsConfigMonetizationConfigToProto(o *apigee.OrganizationAddonsConfigMonetizationConfig) *apigeepb.ApigeeOrganizationAddonsConfigMonetizationConfig {
	if o == nil {
		return nil
	}
	p := &apigeepb.ApigeeOrganizationAddonsConfigMonetizationConfig{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// OrganizationToProto converts a Organization resource to its proto representation.
func OrganizationToProto(resource *apigee.Organization) *apigeepb.ApigeeOrganization {
	p := &apigeepb.ApigeeOrganization{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetExpiresAt(dcl.ValueOrEmptyInt64(resource.ExpiresAt))
	p.SetAnalyticsRegion(dcl.ValueOrEmptyString(resource.AnalyticsRegion))
	p.SetAuthorizedNetwork(dcl.ValueOrEmptyString(resource.AuthorizedNetwork))
	p.SetRuntimeType(ApigeeOrganizationRuntimeTypeEnumToProto(resource.RuntimeType))
	p.SetSubscriptionType(ApigeeOrganizationSubscriptionTypeEnumToProto(resource.SubscriptionType))
	p.SetBillingType(ApigeeOrganizationBillingTypeEnumToProto(resource.BillingType))
	p.SetAddonsConfig(ApigeeOrganizationAddonsConfigToProto(resource.AddonsConfig))
	p.SetCaCertificate(dcl.ValueOrEmptyString(resource.CaCertificate))
	p.SetRuntimeDatabaseEncryptionKeyName(dcl.ValueOrEmptyString(resource.RuntimeDatabaseEncryptionKeyName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetState(ApigeeOrganizationStateEnumToProto(resource.State))
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
func (s *OrganizationServer) applyOrganization(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeOrganizationRequest) (*apigeepb.ApigeeOrganization, error) {
	p := ProtoToOrganization(request.GetResource())
	res, err := c.ApplyOrganization(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OrganizationToProto(res)
	return r, nil
}

// applyApigeeOrganization handles the gRPC request by passing it to the underlying Organization Apply() method.
func (s *OrganizationServer) ApplyApigeeOrganization(ctx context.Context, request *apigeepb.ApplyApigeeOrganizationRequest) (*apigeepb.ApigeeOrganization, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOrganization(ctx, cl, request)
}

// DeleteOrganization handles the gRPC request by passing it to the underlying Organization Delete() method.
func (s *OrganizationServer) DeleteApigeeOrganization(ctx context.Context, request *apigeepb.DeleteApigeeOrganizationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOrganization(ctx, ProtoToOrganization(request.GetResource()))

}

// ListApigeeOrganization handles the gRPC request by passing it to the underlying OrganizationList() method.
func (s *OrganizationServer) ListApigeeOrganization(ctx context.Context, request *apigeepb.ListApigeeOrganizationRequest) (*apigeepb.ListApigeeOrganizationResponse, error) {
	cl, err := createConfigOrganization(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOrganization(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeOrganization
	for _, r := range resources.Items {
		rp := OrganizationToProto(r)
		protos = append(protos, rp)
	}
	p := &apigeepb.ListApigeeOrganizationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOrganization(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
