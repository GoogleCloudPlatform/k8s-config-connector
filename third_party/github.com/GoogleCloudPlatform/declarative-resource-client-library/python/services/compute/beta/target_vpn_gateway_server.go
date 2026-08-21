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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for TargetVpnGateway.
type TargetVpnGatewayServer struct{}

// ProtoToTargetVpnGatewayStatusEnum converts a TargetVpnGatewayStatusEnum enum from its proto representation.
func ProtoToComputeBetaTargetVpnGatewayStatusEnum(e betapb.ComputeBetaTargetVpnGatewayStatusEnum) *beta.TargetVpnGatewayStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaTargetVpnGatewayStatusEnum_name[int32(e)]; ok {
		e := beta.TargetVpnGatewayStatusEnum(n[len("ComputeBetaTargetVpnGatewayStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetVpnGateway converts a TargetVpnGateway resource from its proto representation.
func ProtoToTargetVpnGateway(p *betapb.ComputeBetaTargetVpnGateway) *beta.TargetVpnGateway {
	obj := &beta.TargetVpnGateway{
		Id:          dcl.Int64OrNil(p.Id),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Region:      dcl.StringOrNil(p.Region),
		Network:     dcl.StringOrNil(p.Network),
		Status:      ProtoToComputeBetaTargetVpnGatewayStatusEnum(p.GetStatus()),
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
func ComputeBetaTargetVpnGatewayStatusEnumToProto(e *beta.TargetVpnGatewayStatusEnum) betapb.ComputeBetaTargetVpnGatewayStatusEnum {
	if e == nil {
		return betapb.ComputeBetaTargetVpnGatewayStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaTargetVpnGatewayStatusEnum_value["TargetVpnGatewayStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaTargetVpnGatewayStatusEnum(v)
	}
	return betapb.ComputeBetaTargetVpnGatewayStatusEnum(0)
}

// TargetVpnGatewayToProto converts a TargetVpnGateway resource to its proto representation.
func TargetVpnGatewayToProto(resource *beta.TargetVpnGateway) *betapb.ComputeBetaTargetVpnGateway {
	p := &betapb.ComputeBetaTargetVpnGateway{
		Id:          dcl.ValueOrEmptyInt64(resource.Id),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Region:      dcl.ValueOrEmptyString(resource.Region),
		Network:     dcl.ValueOrEmptyString(resource.Network),
		Status:      ComputeBetaTargetVpnGatewayStatusEnumToProto(resource.Status),
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
func (s *TargetVpnGatewayServer) applyTargetVpnGateway(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaTargetVpnGatewayRequest) (*betapb.ComputeBetaTargetVpnGateway, error) {
	p := ProtoToTargetVpnGateway(request.GetResource())
	res, err := c.ApplyTargetVpnGateway(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetVpnGatewayToProto(res)
	return r, nil
}

// ApplyTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGateway Apply() method.
func (s *TargetVpnGatewayServer) ApplyComputeBetaTargetVpnGateway(ctx context.Context, request *betapb.ApplyComputeBetaTargetVpnGatewayRequest) (*betapb.ComputeBetaTargetVpnGateway, error) {
	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetVpnGateway(ctx, cl, request)
}

// DeleteTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGateway Delete() method.
func (s *TargetVpnGatewayServer) DeleteComputeBetaTargetVpnGateway(ctx context.Context, request *betapb.DeleteComputeBetaTargetVpnGatewayRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetVpnGateway(ctx, ProtoToTargetVpnGateway(request.GetResource()))

}

// ListComputeBetaTargetVpnGateway handles the gRPC request by passing it to the underlying TargetVpnGatewayList() method.
func (s *TargetVpnGatewayServer) ListComputeBetaTargetVpnGateway(ctx context.Context, request *betapb.ListComputeBetaTargetVpnGatewayRequest) (*betapb.ListComputeBetaTargetVpnGatewayResponse, error) {
	cl, err := createConfigTargetVpnGateway(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetVpnGateway(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaTargetVpnGateway
	for _, r := range resources.Items {
		rp := TargetVpnGatewayToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaTargetVpnGatewayResponse{Items: protos}, nil
}

func createConfigTargetVpnGateway(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
