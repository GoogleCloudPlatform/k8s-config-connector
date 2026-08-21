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
	servicenetworkingpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/servicenetworking/servicenetworking_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicenetworking"
)

// Server implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *servicenetworkingpb.ServicenetworkingConnection) *servicenetworking.Connection {
	obj := &servicenetworking.Connection{
		Network: dcl.StringOrNil(p.Network),
		Project: dcl.StringOrNil(p.Project),
		Name:    dcl.StringOrNil(p.Name),
		Service: dcl.StringOrNil(p.Service),
	}
	for _, r := range p.GetReservedPeeringRanges() {
		obj.ReservedPeeringRanges = append(obj.ReservedPeeringRanges, r)
	}
	return obj
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *servicenetworking.Connection) *servicenetworkingpb.ServicenetworkingConnection {
	p := &servicenetworkingpb.ServicenetworkingConnection{
		Network: dcl.ValueOrEmptyString(resource.Network),
		Project: dcl.ValueOrEmptyString(resource.Project),
		Name:    dcl.ValueOrEmptyString(resource.Name),
		Service: dcl.ValueOrEmptyString(resource.Service),
	}
	for _, r := range resource.ReservedPeeringRanges {
		p.ReservedPeeringRanges = append(p.ReservedPeeringRanges, r)
	}

	return p
}

// ApplyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) applyConnection(ctx context.Context, c *servicenetworking.Client, request *servicenetworkingpb.ApplyServicenetworkingConnectionRequest) (*servicenetworkingpb.ServicenetworkingConnection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// ApplyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyServicenetworkingConnection(ctx context.Context, request *servicenetworkingpb.ApplyServicenetworkingConnectionRequest) (*servicenetworkingpb.ServicenetworkingConnection, error) {
	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteServicenetworkingConnection(ctx context.Context, request *servicenetworkingpb.DeleteServicenetworkingConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListServicenetworkingConnection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListServicenetworkingConnection(ctx context.Context, request *servicenetworkingpb.ListServicenetworkingConnectionRequest) (*servicenetworkingpb.ListServicenetworkingConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.Project, request.Network, request.Service)
	if err != nil {
		return nil, err
	}
	var protos []*servicenetworkingpb.ServicenetworkingConnection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	return &servicenetworkingpb.ListServicenetworkingConnectionResponse{Items: protos}, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*servicenetworking.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return servicenetworking.NewClient(conf), nil
}
