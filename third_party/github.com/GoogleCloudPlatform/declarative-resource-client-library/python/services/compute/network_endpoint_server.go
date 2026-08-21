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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for NetworkEndpoint.
type NetworkEndpointServer struct{}

// ProtoToNetworkEndpoint converts a NetworkEndpoint resource from its proto representation.
func ProtoToNetworkEndpoint(p *computepb.ComputeNetworkEndpoint) *compute.NetworkEndpoint {
	obj := &compute.NetworkEndpoint{
		Port:      dcl.Int64OrNil(p.Port),
		IPAddress: dcl.StringOrNil(p.IpAddress),
		Fqdn:      dcl.StringOrNil(p.Fqdn),
		Instance:  dcl.StringOrNil(p.Instance),
		Project:   dcl.StringOrNil(p.Project),
		Location:  dcl.StringOrNil(p.Location),
		Group:     dcl.StringOrNil(p.Group),
	}
	return obj
}

// NetworkEndpointToProto converts a NetworkEndpoint resource to its proto representation.
func NetworkEndpointToProto(resource *compute.NetworkEndpoint) *computepb.ComputeNetworkEndpoint {
	p := &computepb.ComputeNetworkEndpoint{
		Port:      dcl.ValueOrEmptyInt64(resource.Port),
		IpAddress: dcl.ValueOrEmptyString(resource.IPAddress),
		Fqdn:      dcl.ValueOrEmptyString(resource.Fqdn),
		Instance:  dcl.ValueOrEmptyString(resource.Instance),
		Project:   dcl.ValueOrEmptyString(resource.Project),
		Location:  dcl.ValueOrEmptyString(resource.Location),
		Group:     dcl.ValueOrEmptyString(resource.Group),
	}

	return p
}

// ApplyNetworkEndpoint handles the gRPC request by passing it to the underlying NetworkEndpoint Apply() method.
func (s *NetworkEndpointServer) applyNetworkEndpoint(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeNetworkEndpointRequest) (*computepb.ComputeNetworkEndpoint, error) {
	p := ProtoToNetworkEndpoint(request.GetResource())
	res, err := c.ApplyNetworkEndpoint(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkEndpointToProto(res)
	return r, nil
}

// ApplyNetworkEndpoint handles the gRPC request by passing it to the underlying NetworkEndpoint Apply() method.
func (s *NetworkEndpointServer) ApplyComputeNetworkEndpoint(ctx context.Context, request *computepb.ApplyComputeNetworkEndpointRequest) (*computepb.ComputeNetworkEndpoint, error) {
	cl, err := createConfigNetworkEndpoint(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNetworkEndpoint(ctx, cl, request)
}

// DeleteNetworkEndpoint handles the gRPC request by passing it to the underlying NetworkEndpoint Delete() method.
func (s *NetworkEndpointServer) DeleteComputeNetworkEndpoint(ctx context.Context, request *computepb.DeleteComputeNetworkEndpointRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkEndpoint(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkEndpoint(ctx, ProtoToNetworkEndpoint(request.GetResource()))

}

// ListComputeNetworkEndpoint handles the gRPC request by passing it to the underlying NetworkEndpointList() method.
func (s *NetworkEndpointServer) ListComputeNetworkEndpoint(ctx context.Context, request *computepb.ListComputeNetworkEndpointRequest) (*computepb.ListComputeNetworkEndpointResponse, error) {
	cl, err := createConfigNetworkEndpoint(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkEndpoint(ctx, request.Project, request.Location, request.Group)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeNetworkEndpoint
	for _, r := range resources.Items {
		rp := NetworkEndpointToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeNetworkEndpointResponse{Items: protos}, nil
}

func createConfigNetworkEndpoint(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
