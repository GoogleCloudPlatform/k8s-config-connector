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

// GatewayServer implements the gRPC interface for Gateway.
type GatewayServer struct{}

// ProtoToGatewayTypeEnum converts a GatewayTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaGatewayTypeEnum(e alphapb.NetworkservicesAlphaGatewayTypeEnum) *alpha.GatewayTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaGatewayTypeEnum_name[int32(e)]; ok {
		e := alpha.GatewayTypeEnum(n[len("NetworkservicesAlphaGatewayTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGateway converts a Gateway resource from its proto representation.
func ProtoToGateway(p *alphapb.NetworkservicesAlphaGateway) *alpha.Gateway {
	obj := &alpha.Gateway{
		Name:            dcl.StringOrNil(p.GetName()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		Type:            ProtoToNetworkservicesAlphaGatewayTypeEnum(p.GetType()),
		Scope:           dcl.StringOrNil(p.GetScope()),
		ServerTlsPolicy: dcl.StringOrNil(p.GetServerTlsPolicy()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
		SelfLink:        dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetAddresses() {
		obj.Addresses = append(obj.Addresses, r)
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// GatewayTypeEnumToProto converts a GatewayTypeEnum enum to its proto representation.
func NetworkservicesAlphaGatewayTypeEnumToProto(e *alpha.GatewayTypeEnum) alphapb.NetworkservicesAlphaGatewayTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaGatewayTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaGatewayTypeEnum_value["GatewayTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaGatewayTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaGatewayTypeEnum(0)
}

// GatewayToProto converts a Gateway resource to its proto representation.
func GatewayToProto(resource *alpha.Gateway) *alphapb.NetworkservicesAlphaGateway {
	p := &alphapb.NetworkservicesAlphaGateway{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(NetworkservicesAlphaGatewayTypeEnumToProto(resource.Type))
	p.SetScope(dcl.ValueOrEmptyString(resource.Scope))
	p.SetServerTlsPolicy(dcl.ValueOrEmptyString(resource.ServerTlsPolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sAddresses := make([]string, len(resource.Addresses))
	for i, r := range resource.Addresses {
		sAddresses[i] = r
	}
	p.SetAddresses(sAddresses)
	sPorts := make([]int64, len(resource.Ports))
	for i, r := range resource.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)

	return p
}

// applyGateway handles the gRPC request by passing it to the underlying Gateway Apply() method.
func (s *GatewayServer) applyGateway(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaGatewayRequest) (*alphapb.NetworkservicesAlphaGateway, error) {
	p := ProtoToGateway(request.GetResource())
	res, err := c.ApplyGateway(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GatewayToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaGateway handles the gRPC request by passing it to the underlying Gateway Apply() method.
func (s *GatewayServer) ApplyNetworkservicesAlphaGateway(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaGatewayRequest) (*alphapb.NetworkservicesAlphaGateway, error) {
	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGateway(ctx, cl, request)
}

// DeleteGateway handles the gRPC request by passing it to the underlying Gateway Delete() method.
func (s *GatewayServer) DeleteNetworkservicesAlphaGateway(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaGatewayRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGateway(ctx, ProtoToGateway(request.GetResource()))

}

// ListNetworkservicesAlphaGateway handles the gRPC request by passing it to the underlying GatewayList() method.
func (s *GatewayServer) ListNetworkservicesAlphaGateway(ctx context.Context, request *alphapb.ListNetworkservicesAlphaGatewayRequest) (*alphapb.ListNetworkservicesAlphaGatewayResponse, error) {
	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGateway(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaGateway
	for _, r := range resources.Items {
		rp := GatewayToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaGatewayResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGateway(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
