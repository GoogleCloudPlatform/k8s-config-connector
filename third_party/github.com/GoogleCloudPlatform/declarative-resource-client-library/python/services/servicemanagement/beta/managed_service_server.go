// Copyright 2022 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/servicemanagement/beta/servicemanagement_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicemanagement/beta"
)

// ManagedServiceServer implements the gRPC interface for ManagedService.
type ManagedServiceServer struct{}

// ProtoToManagedService converts a ManagedService resource from its proto representation.
func ProtoToManagedService(p *betapb.ServicemanagementBetaManagedService) *beta.ManagedService {
	obj := &beta.ManagedService{
		Name:    dcl.StringOrNil(p.GetName()),
		Project: dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ManagedServiceToProto converts a ManagedService resource to its proto representation.
func ManagedServiceToProto(resource *beta.ManagedService) *betapb.ServicemanagementBetaManagedService {
	p := &betapb.ServicemanagementBetaManagedService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyManagedService handles the gRPC request by passing it to the underlying ManagedService Apply() method.
func (s *ManagedServiceServer) applyManagedService(ctx context.Context, c *beta.Client, request *betapb.ApplyServicemanagementBetaManagedServiceRequest) (*betapb.ServicemanagementBetaManagedService, error) {
	p := ProtoToManagedService(request.GetResource())
	res, err := c.ApplyManagedService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ManagedServiceToProto(res)
	return r, nil
}

// applyServicemanagementBetaManagedService handles the gRPC request by passing it to the underlying ManagedService Apply() method.
func (s *ManagedServiceServer) ApplyServicemanagementBetaManagedService(ctx context.Context, request *betapb.ApplyServicemanagementBetaManagedServiceRequest) (*betapb.ServicemanagementBetaManagedService, error) {
	cl, err := createConfigManagedService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyManagedService(ctx, cl, request)
}

// DeleteManagedService handles the gRPC request by passing it to the underlying ManagedService Delete() method.
func (s *ManagedServiceServer) DeleteServicemanagementBetaManagedService(ctx context.Context, request *betapb.DeleteServicemanagementBetaManagedServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigManagedService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteManagedService(ctx, ProtoToManagedService(request.GetResource()))

}

// ListServicemanagementBetaManagedService handles the gRPC request by passing it to the underlying ManagedServiceList() method.
func (s *ManagedServiceServer) ListServicemanagementBetaManagedService(ctx context.Context, request *betapb.ListServicemanagementBetaManagedServiceRequest) (*betapb.ListServicemanagementBetaManagedServiceResponse, error) {
	cl, err := createConfigManagedService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListManagedService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ServicemanagementBetaManagedService
	for _, r := range resources.Items {
		rp := ManagedServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListServicemanagementBetaManagedServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigManagedService(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
