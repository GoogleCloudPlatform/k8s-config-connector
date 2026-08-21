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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/alpha/iap_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha"
)

// IdentityAwareProxyClientServer implements the gRPC interface for IdentityAwareProxyClient.
type IdentityAwareProxyClientServer struct{}

// ProtoToIdentityAwareProxyClient converts a IdentityAwareProxyClient resource from its proto representation.
func ProtoToIdentityAwareProxyClient(p *alphapb.IapAlphaIdentityAwareProxyClient) *alpha.IdentityAwareProxyClient {
	obj := &alpha.IdentityAwareProxyClient{
		Name:        dcl.StringOrNil(p.GetName()),
		Secret:      dcl.StringOrNil(p.GetSecret()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Brand:       dcl.StringOrNil(p.GetBrand()),
	}
	return obj
}

// IdentityAwareProxyClientToProto converts a IdentityAwareProxyClient resource to its proto representation.
func IdentityAwareProxyClientToProto(resource *alpha.IdentityAwareProxyClient) *alphapb.IapAlphaIdentityAwareProxyClient {
	p := &alphapb.IapAlphaIdentityAwareProxyClient{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSecret(dcl.ValueOrEmptyString(resource.Secret))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetBrand(dcl.ValueOrEmptyString(resource.Brand))

	return p
}

// applyIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) applyIdentityAwareProxyClient(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIapAlphaIdentityAwareProxyClientRequest) (*alphapb.IapAlphaIdentityAwareProxyClient, error) {
	p := ProtoToIdentityAwareProxyClient(request.GetResource())
	res, err := c.ApplyIdentityAwareProxyClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := IdentityAwareProxyClientToProto(res)
	return r, nil
}

// applyIapAlphaIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Apply() method.
func (s *IdentityAwareProxyClientServer) ApplyIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.ApplyIapAlphaIdentityAwareProxyClientRequest) (*alphapb.IapAlphaIdentityAwareProxyClient, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyIdentityAwareProxyClient(ctx, cl, request)
}

// DeleteIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClient Delete() method.
func (s *IdentityAwareProxyClientServer) DeleteIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.DeleteIapAlphaIdentityAwareProxyClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigIdentityAwareProxyClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteIdentityAwareProxyClient(ctx, ProtoToIdentityAwareProxyClient(request.GetResource()))

}

// ListIapAlphaIdentityAwareProxyClient handles the gRPC request by passing it to the underlying IdentityAwareProxyClientList() method.
func (s *IdentityAwareProxyClientServer) ListIapAlphaIdentityAwareProxyClient(ctx context.Context, request *alphapb.ListIapAlphaIdentityAwareProxyClientRequest) (*alphapb.ListIapAlphaIdentityAwareProxyClientResponse, error) {
	cl, err := createConfigIdentityAwareProxyClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListIdentityAwareProxyClient(ctx, request.GetProject(), request.GetBrand())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IapAlphaIdentityAwareProxyClient
	for _, r := range resources.Items {
		rp := IdentityAwareProxyClientToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIapAlphaIdentityAwareProxyClientResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigIdentityAwareProxyClient(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
