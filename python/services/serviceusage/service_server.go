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

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceStateEnum converts a ServiceStateEnum enum from its proto representation.
func ProtoToServiceusageServiceStateEnum(e serviceusagepb.ServiceusageServiceStateEnum) *serviceusage.ServiceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := serviceusagepb.ServiceusageServiceStateEnum_name[int32(e)]; ok {
		e := serviceusage.ServiceStateEnum(n[len("ServiceusageServiceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *serviceusagepb.ServiceusageService) *serviceusage.Service {
	obj := &serviceusage.Service{
		Name:    dcl.StringOrNil(p.Name),
		State:   ProtoToServiceusageServiceStateEnum(p.GetState()),
		Project: dcl.StringOrNil(p.Project),
	}
	return obj
}

// ServiceStateEnumToProto converts a ServiceStateEnum enum to its proto representation.
func ServiceusageServiceStateEnumToProto(e *serviceusage.ServiceStateEnum) serviceusagepb.ServiceusageServiceStateEnum {
	if e == nil {
		return serviceusagepb.ServiceusageServiceStateEnum(0)
	}
	if v, ok := serviceusagepb.ServiceusageServiceStateEnum_value["ServiceStateEnum"+string(*e)]; ok {
		return serviceusagepb.ServiceusageServiceStateEnum(v)
	}
	return serviceusagepb.ServiceusageServiceStateEnum(0)
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *serviceusage.Service) *serviceusagepb.ServiceusageService {
	p := &serviceusagepb.ServiceusageService{
		Name:    dcl.ValueOrEmptyString(resource.Name),
		State:   ServiceusageServiceStateEnumToProto(resource.State),
		Project: dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *serviceusage.Client, request *serviceusagepb.ApplyServiceusageServiceRequest) (*serviceusagepb.ServiceusageService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// ApplyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyServiceusageService(ctx context.Context, request *serviceusagepb.ApplyServiceusageServiceRequest) (*serviceusagepb.ServiceusageService, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteServiceusageService(ctx context.Context, request *serviceusagepb.DeleteServiceusageServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListServiceusageService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListServiceusageService(ctx context.Context, request *serviceusagepb.ListServiceusageServiceRequest) (*serviceusagepb.ListServiceusageServiceResponse, error) {
	cl, err := createConfigService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*serviceusagepb.ServiceusageService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	return &serviceusagepb.ListServiceusageServiceResponse{Items: protos}, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*serviceusage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return serviceusage.NewClient(conf), nil
}
