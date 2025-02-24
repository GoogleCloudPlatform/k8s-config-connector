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
	serviceusagepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/serviceusage/serviceusage_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/serviceusage"
)

// Server implements the gRPC interface for ServiceusageService.
type ServiceusageServiceServer struct{}

// ProtoToServiceusageServiceStateEnum converts a ServiceusageServiceStateEnum enum from its proto representation.
func ProtoToServiceusageServiceusageServiceStateEnum(e serviceusagepb.ServiceusageServiceusageServiceStateEnum) *serviceusage.ServiceusageServiceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := serviceusagepb.ServiceusageServiceusageServiceStateEnum_name[int32(e)]; ok {
		e := serviceusage.ServiceusageServiceStateEnum(n[len("ServiceusageServiceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceusageService converts a ServiceusageService resource from its proto representation.
func ProtoToServiceusageService(p *serviceusagepb.ServiceusageServiceusageService) *serviceusage.ServiceusageService {
	obj := &serviceusage.ServiceusageService{
		Name:    dcl.StringOrNil(p.Name),
		State:   ProtoToServiceusageServiceusageServiceStateEnum(p.GetState()),
		Project: dcl.StringOrNil(p.Project),
	}
	return obj
}

// ServiceusageServiceStateEnumToProto converts a ServiceusageServiceStateEnum enum to its proto representation.
func ServiceusageServiceusageServiceStateEnumToProto(e *serviceusage.ServiceusageServiceStateEnum) serviceusagepb.ServiceusageServiceusageServiceStateEnum {
	if e == nil {
		return serviceusagepb.ServiceusageServiceusageServiceStateEnum(0)
	}
	if v, ok := serviceusagepb.ServiceusageServiceusageServiceStateEnum_value["ServiceusageServiceStateEnum"+string(*e)]; ok {
		return serviceusagepb.ServiceusageServiceusageServiceStateEnum(v)
	}
	return serviceusagepb.ServiceusageServiceusageServiceStateEnum(0)
}

// ServiceusageServiceToProto converts a ServiceusageService resource to its proto representation.
func ServiceusageServiceToProto(resource *serviceusage.ServiceusageService) *serviceusagepb.ServiceusageServiceusageService {
	p := &serviceusagepb.ServiceusageServiceusageService{
		Name:    dcl.ValueOrEmptyString(resource.Name),
		State:   ServiceusageServiceusageServiceStateEnumToProto(resource.State),
		Project: dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyServiceusageService handles the gRPC request by passing it to the underlying ServiceusageService Apply() method.
func (s *ServiceusageServiceServer) applyServiceusageService(ctx context.Context, c *serviceusage.Client, request *serviceusagepb.ApplyServiceusageServiceusageServiceRequest) (*serviceusagepb.ServiceusageServiceusageService, error) {
	p := ProtoToServiceusageService(request.GetResource())
	res, err := c.ApplyServiceusageService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceusageServiceToProto(res)
	return r, nil
}

// ApplyServiceusageService handles the gRPC request by passing it to the underlying ServiceusageService Apply() method.
func (s *ServiceusageServiceServer) ApplyServiceusageServiceusageService(ctx context.Context, request *serviceusagepb.ApplyServiceusageServiceusageServiceRequest) (*serviceusagepb.ServiceusageServiceusageService, error) {
	cl, err := createConfigServiceusageService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServiceusageService(ctx, cl, request)
}

// DeleteServiceusageService handles the gRPC request by passing it to the underlying ServiceusageService Delete() method.
func (s *ServiceusageServiceServer) DeleteServiceusageServiceusageService(ctx context.Context, request *serviceusagepb.DeleteServiceusageServiceusageServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceusageService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceusageService(ctx, ProtoToServiceusageService(request.GetResource()))

}

// ListServiceusageService handles the gRPC request by passing it to the underlying ServiceusageServiceList() method.
func (s *ServiceusageServiceServer) ListServiceusageServiceusageService(ctx context.Context, request *serviceusagepb.ListServiceusageServiceusageServiceRequest) (*serviceusagepb.ListServiceusageServiceusageServiceResponse, error) {
	cl, err := createConfigServiceusageService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceusageService(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*serviceusagepb.ServiceusageServiceusageService
	for _, r := range resources.Items {
		rp := ServiceusageServiceToProto(r)
		protos = append(protos, rp)
	}
	return &serviceusagepb.ListServiceusageServiceusageServiceResponse{Items: protos}, nil
}

func createConfigServiceusageService(ctx context.Context, service_account_file string) (*serviceusage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return serviceusage.NewClient(conf), nil
}
