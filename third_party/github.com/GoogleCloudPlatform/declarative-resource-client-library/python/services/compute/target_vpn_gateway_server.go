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

// Server implements the gRPC interface for TargetVpnGateway.
type TargetVpnGatewayServer struct{}

// ProtoToTargetVpnGatewayStatusEnum converts a TargetVpnGatewayStatusEnum enum from its proto representation.
func ProtoToComputeTargetVpnGatewayStatusEnum(e computepb.ComputeTargetVpnGatewayStatusEnum) *compute.TargetVpnGatewayStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeTargetVpnGatewayStatusEnum_name[int32(e)]; ok {
		e := compute.TargetVpnGatewayStatusEnum(n[len("ComputeTargetVpnGatewayStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetVpnGateway converts a TargetVpnGateway resource from its proto representation.
func ProtoToTargetVpnGateway(p *computepb.ComputeTargetVpnGateway) *compute.TargetVpnGateway {
	obj := &compute.TargetVpnGateway{
		Id:          dcl.Int64OrNil(p.Id),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Region:      dcl.StringOrNil(p.Region),
		Network:     dcl.StringOrNil(p.Network),
		Status:      ProtoToComputeTargetVpnGatewayStatusEnum(p.GetStatus()),
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetTunnel() {
		obj.Tunnel = append(obj.Tunnel, r)
	}
	for _, r := range p.GetForwardingRule() {
		obj.ForwardingRule = append(obj.ForwardingRule, r)
	}
	return obj
}

// TargetVpnGatewayStatusEnumToProto converts a TargetVpnGatewayStatusEnum enum to its proto representation.
func ComputeTargetVpnGatewayStatusEnumToProto(e *compute.TargetVpnGatewayStatusEnum) computepb.ComputeTargetVpnGatewayStatusEnum {
	if e == nil {
		return computepb.ComputeTargetVpnGatewayStatusEnum(0)
	}
	if v, ok := computepb.ComputeTargetVpnGatewayStatusEnum_value["TargetVpnGatewayStatusEnum"+string(*e)]; ok {
		return computepb.ComputeTargetVpnGatewayStatusEnum(v)
	}
	return computepb.ComputeTargetVpnGatewayStatusEnum(0)
}

// TargetVpnGatewayToProto converts a TargetVpnGateway resource to its proto representation.
func TargetVpnGatewayToProto(resource *compute.TargetVpnGateway) *computepb.ComputeTargetVpnGateway {
	p := &computepb.ComputeTargetVpnGateway{
		Id:          dcl.ValueOrEmptyInt64(resource.Id),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Region:      dcl.ValueOrEmptyString(resource.Region),
		Network:     dcl.ValueOrEmptyString(resource.Network),
		Status:      ComputeTargetVpnGatewayStatusEnumToProto(resource.Status),
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Tunnel {
		p.Tunnel = append(p.Tunnel, r)
	}
	for _, r := range resource.ForwardingRule {
		p.ForwardingRule = append(p.ForwardingRule, r)
	}

	return p
}

// ApplyTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGateway Apply() method.
func (s *TargetVpnGatewayServer) applyTargetVpnGateway(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeTargetVpnGatewayRequest) (*computepb.ComputeTargetVpnGateway, error) {
	p := ProtoToTargetVpnGateway(request.GetResource())
	res, err := c.ApplyTargetVpnGateway(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetVpnGatewayToProto(res)
	return r, nil
}

// ApplyTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGateway Apply() method.
func (s *TargetVpnGatewayServer) ApplyComputeTargetVpnGateway(ctx context.Context, request *computepb.ApplyComputeTargetVpnGatewayRequest) (*computepb.ComputeTargetVpnGateway, error) {
	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetVpnGateway(ctx, cl, request)
}

// DeleteTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGateway Delete() method.
func (s *TargetVpnGatewayServer) DeleteComputeTargetVpnGateway(ctx context.Context, request *computepb.DeleteComputeTargetVpnGatewayRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetVpnGateway(ctx, ProtoToTargetVpnGateway(request.GetResource()))

}

// ListComputeTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGatewayList() method.
func (s *TargetVpnGatewayServer) ListComputeTargetVpnGateway(ctx context.Context, request *computepb.ListComputeTargetVpnGatewayRequest) (*computepb.ListComputeTargetVpnGatewayResponse, error) {
	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetVpnGateway(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeTargetVpnGateway
	for _, r := range resources.Items {
		rp := TargetVpnGatewayToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeTargetVpnGatewayResponse{Items: protos}, nil
}

func createConfigTargetVpnGateway(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
