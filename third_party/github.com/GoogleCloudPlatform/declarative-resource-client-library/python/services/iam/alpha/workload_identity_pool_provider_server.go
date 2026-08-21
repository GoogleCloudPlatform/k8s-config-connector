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

// WorkloadIdentityPoolProviderServer implements the gRPC interface for WorkloadIdentityPoolProvider.
type WorkloadIdentityPoolProviderServer struct{}

// ProtoToWorkloadIdentityPoolProviderStateEnum converts a WorkloadIdentityPoolProviderStateEnum enum from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderStateEnum(e alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum) *alpha.WorkloadIdentityPoolProviderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum_name[int32(e)]; ok {
		e := alpha.WorkloadIdentityPoolProviderStateEnum(n[len("IamAlphaWorkloadIdentityPoolProviderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPoolProviderAws converts a WorkloadIdentityPoolProviderAws object from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderAws(p *alphapb.IamAlphaWorkloadIdentityPoolProviderAws) *alpha.WorkloadIdentityPoolProviderAws {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadIdentityPoolProviderAws{
		AccountId: dcl.StringOrNil(p.GetAccountId()),
	}
	for _, r := range p.GetStsUri() {
		obj.StsUri = append(obj.StsUri, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProviderOidc converts a WorkloadIdentityPoolProviderOidc object from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolProviderOidc(p *alphapb.IamAlphaWorkloadIdentityPoolProviderOidc) *alpha.WorkloadIdentityPoolProviderOidc {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadIdentityPoolProviderOidc{
		IssuerUri: dcl.StringOrNil(p.GetIssuerUri()),
	}
	for _, r := range p.GetAllowedAudiences() {
		obj.AllowedAudiences = append(obj.AllowedAudiences, r)
	}
	return obj
}

// ProtoToWorkloadIdentityPoolProvider converts a WorkloadIdentityPoolProvider resource from its proto representation.
func ProtoToWorkloadIdentityPoolProvider(p *alphapb.IamAlphaWorkloadIdentityPoolProvider) *alpha.WorkloadIdentityPoolProvider {
	obj := &alpha.WorkloadIdentityPoolProvider{
		Name:                 dcl.StringOrNil(p.GetName()),
		DisplayName:          dcl.StringOrNil(p.GetDisplayName()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		State:                ProtoToIamAlphaWorkloadIdentityPoolProviderStateEnum(p.GetState()),
		Disabled:             dcl.Bool(p.GetDisabled()),
		AttributeCondition:   dcl.StringOrNil(p.GetAttributeCondition()),
		Aws:                  ProtoToIamAlphaWorkloadIdentityPoolProviderAws(p.GetAws()),
		Oidc:                 ProtoToIamAlphaWorkloadIdentityPoolProviderOidc(p.GetOidc()),
		Project:              dcl.StringOrNil(p.GetProject()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		WorkloadIdentityPool: dcl.StringOrNil(p.GetWorkloadIdentityPool()),
	}
	return obj
}

// WorkloadIdentityPoolProviderStateEnumToProto converts a WorkloadIdentityPoolProviderStateEnum enum to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderStateEnumToProto(e *alpha.WorkloadIdentityPoolProviderStateEnum) alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum {
	if e == nil {
		return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum_value["WorkloadIdentityPoolProviderStateEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(v)
	}
	return alphapb.IamAlphaWorkloadIdentityPoolProviderStateEnum(0)
}

// WorkloadIdentityPoolProviderAwsToProto converts a WorkloadIdentityPoolProviderAws object to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderAwsToProto(o *alpha.WorkloadIdentityPoolProviderAws) *alphapb.IamAlphaWorkloadIdentityPoolProviderAws {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkloadIdentityPoolProviderAws{}
	p.SetAccountId(dcl.ValueOrEmptyString(o.AccountId))
	sStsUri := make([]string, len(o.StsUri))
	for i, r := range o.StsUri {
		sStsUri[i] = r
	}
	p.SetStsUri(sStsUri)
	return p
}

// WorkloadIdentityPoolProviderOidcToProto converts a WorkloadIdentityPoolProviderOidc object to its proto representation.
func IamAlphaWorkloadIdentityPoolProviderOidcToProto(o *alpha.WorkloadIdentityPoolProviderOidc) *alphapb.IamAlphaWorkloadIdentityPoolProviderOidc {
	if o == nil {
		return nil
	}
	p := &alphapb.IamAlphaWorkloadIdentityPoolProviderOidc{}
	p.SetIssuerUri(dcl.ValueOrEmptyString(o.IssuerUri))
	sAllowedAudiences := make([]string, len(o.AllowedAudiences))
	for i, r := range o.AllowedAudiences {
		sAllowedAudiences[i] = r
	}
	p.SetAllowedAudiences(sAllowedAudiences)
	return p
}

// WorkloadIdentityPoolProviderToProto converts a WorkloadIdentityPoolProvider resource to its proto representation.
func WorkloadIdentityPoolProviderToProto(resource *alpha.WorkloadIdentityPoolProvider) *alphapb.IamAlphaWorkloadIdentityPoolProvider {
	p := &alphapb.IamAlphaWorkloadIdentityPoolProvider{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamAlphaWorkloadIdentityPoolProviderStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetAttributeCondition(dcl.ValueOrEmptyString(resource.AttributeCondition))
	p.SetAws(IamAlphaWorkloadIdentityPoolProviderAwsToProto(resource.Aws))
	p.SetOidc(IamAlphaWorkloadIdentityPoolProviderOidcToProto(resource.Oidc))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetWorkloadIdentityPool(dcl.ValueOrEmptyString(resource.WorkloadIdentityPool))
	mAttributeMapping := make(map[string]string, len(resource.AttributeMapping))
	for k, r := range resource.AttributeMapping {
		mAttributeMapping[k] = r
	}
	p.SetAttributeMapping(mAttributeMapping)

	return p
}

// applyWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) applyWorkloadIdentityPoolProvider(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.IamAlphaWorkloadIdentityPoolProvider, error) {
	p := ProtoToWorkloadIdentityPoolProvider(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPoolProvider(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolProviderToProto(res)
	return r, nil
}

// applyIamAlphaWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Apply() method.
func (s *WorkloadIdentityPoolProviderServer) ApplyIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.IamAlphaWorkloadIdentityPoolProvider, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPoolProvider(ctx, cl, request)
}

// DeleteWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProvider Delete() method.
func (s *WorkloadIdentityPoolProviderServer) DeleteIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.DeleteIamAlphaWorkloadIdentityPoolProviderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPoolProvider(ctx, ProtoToWorkloadIdentityPoolProvider(request.GetResource()))

}

// ListIamAlphaWorkloadIdentityPoolProvider handles the gRPC request by passing it to the underlying WorkloadIdentityPoolProviderList() method.
func (s *WorkloadIdentityPoolProviderServer) ListIamAlphaWorkloadIdentityPoolProvider(ctx context.Context, request *alphapb.ListIamAlphaWorkloadIdentityPoolProviderRequest) (*alphapb.ListIamAlphaWorkloadIdentityPoolProviderResponse, error) {
	cl, err := createConfigWorkloadIdentityPoolProvider(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPoolProvider(ctx, request.GetProject(), request.GetLocation(), request.GetWorkloadIdentityPool())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaWorkloadIdentityPoolProvider
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolProviderToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIamAlphaWorkloadIdentityPoolProviderResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkloadIdentityPoolProvider(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
