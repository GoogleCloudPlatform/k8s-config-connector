// Copyright 2021 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkemulticloud/alpha/gkemulticloud_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/alpha"
)

// Server implements the gRPC interface for AzureClient.
type AzureClientServer struct{}

// ProtoToAzureClient converts a AzureClient resource from its proto representation.
func ProtoToAzureClient(p *alphapb.GkemulticloudAlphaAzureClient) *alpha.AzureClient {
	obj := &alpha.AzureClient{
		Name:          dcl.StringOrNil(p.Name),
		TenantId:      dcl.StringOrNil(p.TenantId),
		ApplicationId: dcl.StringOrNil(p.ApplicationId),
		Certificate:   dcl.StringOrNil(p.Certificate),
		Uid:           dcl.StringOrNil(p.Uid),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	return obj
}

// AzureClientToProto converts a AzureClient resource to its proto representation.
func AzureClientToProto(resource *alpha.AzureClient) *alphapb.GkemulticloudAlphaAzureClient {
	p := &alphapb.GkemulticloudAlphaAzureClient{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		TenantId:      dcl.ValueOrEmptyString(resource.TenantId),
		ApplicationId: dcl.ValueOrEmptyString(resource.ApplicationId),
		Certificate:   dcl.ValueOrEmptyString(resource.Certificate),
		Uid:           dcl.ValueOrEmptyString(resource.Uid),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyAzureClient handles the gRPC request by passing it to the underlying AzureClient Apply() method.
func (s *AzureClientServer) applyAzureClient(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkemulticloudAlphaAzureClientRequest) (*alphapb.GkemulticloudAlphaAzureClient, error) {
	p := ProtoToAzureClient(request.GetResource())
	res, err := c.ApplyAzureClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AzureClientToProto(res)
	return r, nil
}

// ApplyAzureClient handles the gRPC request by passing it to the underlying AzureClient Apply() method.
func (s *AzureClientServer) ApplyGkemulticloudAlphaAzureClient(ctx context.Context, request *alphapb.ApplyGkemulticloudAlphaAzureClientRequest) (*alphapb.GkemulticloudAlphaAzureClient, error) {
	cl, err := createConfigAzureClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAzureClient(ctx, cl, request)
}

// DeleteAzureClient handles the gRPC request by passing it to the underlying AzureClient Delete() method.
func (s *AzureClientServer) DeleteGkemulticloudAlphaAzureClient(ctx context.Context, request *alphapb.DeleteGkemulticloudAlphaAzureClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAzureClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAzureClient(ctx, ProtoToAzureClient(request.GetResource()))

}

// ListGkemulticloudAlphaAzureClient handles the gRPC request by passing it to the underlying AzureClientList() method.
func (s *AzureClientServer) ListGkemulticloudAlphaAzureClient(ctx context.Context, request *alphapb.ListGkemulticloudAlphaAzureClientRequest) (*alphapb.ListGkemulticloudAlphaAzureClientResponse, error) {
	cl, err := createConfigAzureClient(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAzureClient(ctx, ProtoToAzureClient(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkemulticloudAlphaAzureClient
	for _, r := range resources.Items {
		rp := AzureClientToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListGkemulticloudAlphaAzureClientResponse{Items: protos}, nil
}

func createConfigAzureClient(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
