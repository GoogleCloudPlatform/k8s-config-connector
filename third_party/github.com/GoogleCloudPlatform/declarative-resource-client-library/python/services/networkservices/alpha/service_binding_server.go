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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// ServiceBindingServer implements the gRPC interface for ServiceBinding.
type ServiceBindingServer struct{}

// ProtoToServiceBinding converts a ServiceBinding resource from its proto representation.
func ProtoToServiceBinding(p *alphapb.NetworkservicesAlphaServiceBinding) *alpha.ServiceBinding {
	obj := &alpha.ServiceBinding{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Service:     dcl.StringOrNil(p.GetService()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ServiceBindingToProto converts a ServiceBinding resource to its proto representation.
func ServiceBindingToProto(resource *alpha.ServiceBinding) *alphapb.NetworkservicesAlphaServiceBinding {
	p := &alphapb.NetworkservicesAlphaServiceBinding{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetService(dcl.ValueOrEmptyString(resource.Service))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyServiceBinding handles the gRPC request by passing it to the underlying ServiceBinding Apply() method.
func (s *ServiceBindingServer) applyServiceBinding(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaServiceBindingRequest) (*alphapb.NetworkservicesAlphaServiceBinding, error) {
	p := ProtoToServiceBinding(request.GetResource())
	res, err := c.ApplyServiceBinding(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceBindingToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaServiceBinding handles the gRPC request by passing it to the underlying ServiceBinding Apply() method.
func (s *ServiceBindingServer) ApplyNetworkservicesAlphaServiceBinding(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaServiceBindingRequest) (*alphapb.NetworkservicesAlphaServiceBinding, error) {
	cl, err := createConfigServiceBinding(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceBinding(ctx, cl, request)
}

// DeleteServiceBinding handles the gRPC request by passing it to the underlying ServiceBinding Delete() method.
func (s *ServiceBindingServer) DeleteNetworkservicesAlphaServiceBinding(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaServiceBindingRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceBinding(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceBinding(ctx, ProtoToServiceBinding(request.GetResource()))

}

// ListNetworkservicesAlphaServiceBinding handles the gRPC request by passing it to the underlying ServiceBindingList() method.
func (s *ServiceBindingServer) ListNetworkservicesAlphaServiceBinding(ctx context.Context, request *alphapb.ListNetworkservicesAlphaServiceBindingRequest) (*alphapb.ListNetworkservicesAlphaServiceBindingResponse, error) {
	cl, err := createConfigServiceBinding(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceBinding(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaServiceBinding
	for _, r := range resources.Items {
		rp := ServiceBindingToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaServiceBindingResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceBinding(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
