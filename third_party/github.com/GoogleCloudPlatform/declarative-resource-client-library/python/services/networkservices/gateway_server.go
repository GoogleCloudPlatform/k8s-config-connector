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
	networkservicespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/networkservices_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices"
)

// GatewayServer implements the gRPC interface for Gateway.
type GatewayServer struct{}

// ProtoToGatewayTypeEnum converts a GatewayTypeEnum enum from its proto representation.
func ProtoToNetworkservicesGatewayTypeEnum(e networkservicespb.NetworkservicesGatewayTypeEnum) *networkservices.GatewayTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkservicespb.NetworkservicesGatewayTypeEnum_name[int32(e)]; ok {
		e := networkservices.GatewayTypeEnum(n[len("NetworkservicesGatewayTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGateway converts a Gateway resource from its proto representation.
func ProtoToGateway(p *networkservicespb.NetworkservicesGateway) *networkservices.Gateway {
	obj := &networkservices.Gateway{
		Name:            dcl.StringOrNil(p.GetName()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		Type:            ProtoToNetworkservicesGatewayTypeEnum(p.GetType()),
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
func NetworkservicesGatewayTypeEnumToProto(e *networkservices.GatewayTypeEnum) networkservicespb.NetworkservicesGatewayTypeEnum {
	if e == nil {
		return networkservicespb.NetworkservicesGatewayTypeEnum(0)
	}
	if v, ok := networkservicespb.NetworkservicesGatewayTypeEnum_value["GatewayTypeEnum"+string(*e)]; ok {
		return networkservicespb.NetworkservicesGatewayTypeEnum(v)
	}
	return networkservicespb.NetworkservicesGatewayTypeEnum(0)
}

// GatewayToProto converts a Gateway resource to its proto representation.
func GatewayToProto(resource *networkservices.Gateway) *networkservicespb.NetworkservicesGateway {
	p := &networkservicespb.NetworkservicesGateway{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(NetworkservicesGatewayTypeEnumToProto(resource.Type))
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
func (s *GatewayServer) applyGateway(ctx context.Context, c *networkservices.Client, request *networkservicespb.ApplyNetworkservicesGatewayRequest) (*networkservicespb.NetworkservicesGateway, error) {
	p := ProtoToGateway(request.GetResource())
	res, err := c.ApplyGateway(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GatewayToProto(res)
	return r, nil
}

// applyNetworkservicesGateway handles the gRPC request by passing it to the underlying Gateway Apply() method.
func (s *GatewayServer) ApplyNetworkservicesGateway(ctx context.Context, request *networkservicespb.ApplyNetworkservicesGatewayRequest) (*networkservicespb.NetworkservicesGateway, error) {
	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGateway(ctx, cl, request)
}

// DeleteGateway handles the gRPC request by passing it to the underlying Gateway Delete() method.
func (s *GatewayServer) DeleteNetworkservicesGateway(ctx context.Context, request *networkservicespb.DeleteNetworkservicesGatewayRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGateway(ctx, ProtoToGateway(request.GetResource()))

}

// ListNetworkservicesGateway handles the gRPC request by passing it to the underlying GatewayList() method.
func (s *GatewayServer) ListNetworkservicesGateway(ctx context.Context, request *networkservicespb.ListNetworkservicesGatewayRequest) (*networkservicespb.ListNetworkservicesGatewayResponse, error) {
	cl, err := createConfigGateway(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGateway(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*networkservicespb.NetworkservicesGateway
	for _, r := range resources.Items {
		rp := GatewayToProto(r)
		protos = append(protos, rp)
	}
	p := &networkservicespb.ListNetworkservicesGatewayResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGateway(ctx context.Context, service_account_file string) (*networkservices.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkservices.NewClient(conf), nil
}
