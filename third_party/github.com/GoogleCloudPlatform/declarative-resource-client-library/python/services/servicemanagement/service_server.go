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
	servicemanagementpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/servicemanagement/servicemanagement_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicemanagement"
)

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *servicemanagementpb.ServicemanagementService) *servicemanagement.Service {
	obj := &servicemanagement.Service{
		Name:    dcl.StringOrNil(p.Name),
		Project: dcl.StringOrNil(p.Project),
	}
	return obj
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *servicemanagement.Service) *servicemanagementpb.ServicemanagementService {
	p := &servicemanagementpb.ServicemanagementService{
		Name:    dcl.ValueOrEmptyString(resource.Name),
		Project: dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *servicemanagement.Client, request *servicemanagementpb.ApplyServicemanagementServiceRequest) (*servicemanagementpb.ServicemanagementService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyServicemanagementService(ctx context.Context, request *servicemanagementpb.ApplyServicemanagementServiceRequest) (*servicemanagementpb.ServicemanagementService, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteServicemanagementService(ctx context.Context, request *servicemanagementpb.DeleteServicemanagementServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListServicemanagementService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListServicemanagementService(ctx context.Context, request *servicemanagementpb.ListServicemanagementServiceRequest) (*servicemanagementpb.ListServicemanagementServiceResponse, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*servicemanagementpb.ServicemanagementService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	return &servicemanagementpb.ListServicemanagementServiceResponse{Items: protos}, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*servicemanagement.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return servicemanagement.NewClient(conf), nil
}
