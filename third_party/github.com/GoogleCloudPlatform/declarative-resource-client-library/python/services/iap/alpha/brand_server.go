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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/alpha/iap_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha"
)

// BrandServer implements the gRPC interface for Brand.
type BrandServer struct{}

// ProtoToBrand converts a Brand resource from its proto representation.
func ProtoToBrand(p *alphapb.IapAlphaBrand) *alpha.Brand {
	obj := &alpha.Brand{
		ApplicationTitle: dcl.StringOrNil(p.GetApplicationTitle()),
		Name:             dcl.StringOrNil(p.GetName()),
		OrgInternalOnly:  dcl.Bool(p.GetOrgInternalOnly()),
		SupportEmail:     dcl.StringOrNil(p.GetSupportEmail()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// BrandToProto converts a Brand resource to its proto representation.
func BrandToProto(resource *alpha.Brand) *alphapb.IapAlphaBrand {
	p := &alphapb.IapAlphaBrand{}
	p.SetApplicationTitle(dcl.ValueOrEmptyString(resource.ApplicationTitle))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetOrgInternalOnly(dcl.ValueOrEmptyBool(resource.OrgInternalOnly))
	p.SetSupportEmail(dcl.ValueOrEmptyString(resource.SupportEmail))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) applyBrand(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIapAlphaBrandRequest) (*alphapb.IapAlphaBrand, error) {
	p := ProtoToBrand(request.GetResource())
	res, err := c.ApplyBrand(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BrandToProto(res)
	return r, nil
}

// applyIapAlphaBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) ApplyIapAlphaBrand(ctx context.Context, request *alphapb.ApplyIapAlphaBrandRequest) (*alphapb.IapAlphaBrand, error) {
	cl, err := createConfigBrand(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyBrand(ctx, cl, request)
}

// DeleteBrand handles the gRPC request by passing it to the underlying Brand Delete() method.
func (s *BrandServer) DeleteIapAlphaBrand(ctx context.Context, request *alphapb.DeleteIapAlphaBrandRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Brand")

}

// ListIapAlphaBrand handles the gRPC request by passing it to the underlying BrandList() method.
func (s *BrandServer) ListIapAlphaBrand(ctx context.Context, request *alphapb.ListIapAlphaBrandRequest) (*alphapb.ListIapAlphaBrandResponse, error) {
	cl, err := createConfigBrand(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBrand(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IapAlphaBrand
	for _, r := range resources.Items {
		rp := BrandToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIapAlphaBrandResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigBrand(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
