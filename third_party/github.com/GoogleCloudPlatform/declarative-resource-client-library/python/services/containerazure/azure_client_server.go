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
	containerazurepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/containerazure_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure"
)

// ClientServer implements the gRPC interface for Client.
type ClientServer struct{}

// ProtoToClient converts a Client resource from its proto representation.
func ProtoToClient(p *containerazurepb.ContainerazureClient) *containerazure.AzureClient {
	obj := &containerazure.AzureClient{
		Name:          dcl.StringOrNil(p.GetName()),
		TenantId:      dcl.StringOrNil(p.GetTenantId()),
		ApplicationId: dcl.StringOrNil(p.GetApplicationId()),
		Certificate:   dcl.StringOrNil(p.GetCertificate()),
		Uid:           dcl.StringOrNil(p.GetUid()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ClientToProto converts a Client resource to its proto representation.
func ClientToProto(resource *containerazure.AzureClient) *containerazurepb.ContainerazureClient {
	p := &containerazurepb.ContainerazureClient{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTenantId(dcl.ValueOrEmptyString(resource.TenantId))
	p.SetApplicationId(dcl.ValueOrEmptyString(resource.ApplicationId))
	p.SetCertificate(dcl.ValueOrEmptyString(resource.Certificate))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyClient handles the gRPC request by passing it to the underlying Client Apply() method.
func (s *ClientServer) applyClient(ctx context.Context, c *containerazure.Client, request *containerazurepb.ApplyContainerazureClientRequest) (*containerazurepb.ContainerazureClient, error) {
	p := ProtoToClient(request.GetResource())
	res, err := c.ApplyClient(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClientToProto(res)
	return r, nil
}

// applyContainerazureClient handles the gRPC request by passing it to the underlying Client Apply() method.
func (s *ClientServer) ApplyContainerazureClient(ctx context.Context, request *containerazurepb.ApplyContainerazureClientRequest) (*containerazurepb.ContainerazureClient, error) {
	cl, err := createConfigClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyClient(ctx, cl, request)
}

// DeleteClient handles the gRPC request by passing it to the underlying Client Delete() method.
func (s *ClientServer) DeleteContainerazureClient(ctx context.Context, request *containerazurepb.DeleteContainerazureClientRequest) (*emptypb.Empty, error) {

	cl, err := createConfigClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteClient(ctx, ProtoToClient(request.GetResource()))

}

// ListContainerazureClient handles the gRPC request by passing it to the underlying ClientList() method.
func (s *ClientServer) ListContainerazureClient(ctx context.Context, request *containerazurepb.ListContainerazureClientRequest) (*containerazurepb.ListContainerazureClientResponse, error) {
	cl, err := createConfigClient(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListClient(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*containerazurepb.ContainerazureClient
	for _, r := range resources.Items {
		rp := ClientToProto(r)
		protos = append(protos, rp)
	}
	p := &containerazurepb.ListContainerazureClientResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigClient(ctx context.Context, service_account_file string) (*containerazure.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containerazure.NewClient(conf), nil
}
