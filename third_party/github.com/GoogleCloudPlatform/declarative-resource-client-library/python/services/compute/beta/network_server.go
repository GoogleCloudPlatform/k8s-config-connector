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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// NetworkServer implements the gRPC interface for Network.
type NetworkServer struct{}

// ProtoToNetworkRoutingConfigRoutingModeEnum converts a NetworkRoutingConfigRoutingModeEnum enum from its proto representation.
func ProtoToComputeBetaNetworkRoutingConfigRoutingModeEnum(e betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum) *beta.NetworkRoutingConfigRoutingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum_name[int32(e)]; ok {
		e := beta.NetworkRoutingConfigRoutingModeEnum(n[len("ComputeBetaNetworkRoutingConfigRoutingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkRoutingConfig converts a NetworkRoutingConfig object from its proto representation.
func ProtoToComputeBetaNetworkRoutingConfig(p *betapb.ComputeBetaNetworkRoutingConfig) *beta.NetworkRoutingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NetworkRoutingConfig{
		RoutingMode: ProtoToComputeBetaNetworkRoutingConfigRoutingModeEnum(p.GetRoutingMode()),
	}
	return obj
}

// ProtoToNetwork converts a Network resource from its proto representation.
func ProtoToNetwork(p *betapb.ComputeBetaNetwork) *beta.Network {
	obj := &beta.Network{
		Description:           dcl.StringOrNil(p.GetDescription()),
		GatewayIPv4:           dcl.StringOrNil(p.GetGatewayIpv4()),
		Name:                  dcl.StringOrNil(p.GetName()),
		AutoCreateSubnetworks: dcl.Bool(p.GetAutoCreateSubnetworks()),
		RoutingConfig:         ProtoToComputeBetaNetworkRoutingConfig(p.GetRoutingConfig()),
		Mtu:                   dcl.Int64OrNil(p.GetMtu()),
		Project:               dcl.StringOrNil(p.GetProject()),
		SelfLink:              dcl.StringOrNil(p.GetSelfLink()),
		SelfLinkWithId:        dcl.StringOrNil(p.GetSelfLinkWithId()),
	}
	return obj
}

// NetworkRoutingConfigRoutingModeEnumToProto converts a NetworkRoutingConfigRoutingModeEnum enum to its proto representation.
func ComputeBetaNetworkRoutingConfigRoutingModeEnumToProto(e *beta.NetworkRoutingConfigRoutingModeEnum) betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum {
	if e == nil {
		return betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum_value["NetworkRoutingConfigRoutingModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum(v)
	}
	return betapb.ComputeBetaNetworkRoutingConfigRoutingModeEnum(0)
}

// NetworkRoutingConfigToProto converts a NetworkRoutingConfig object to its proto representation.
func ComputeBetaNetworkRoutingConfigToProto(o *beta.NetworkRoutingConfig) *betapb.ComputeBetaNetworkRoutingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaNetworkRoutingConfig{}
	p.SetRoutingMode(ComputeBetaNetworkRoutingConfigRoutingModeEnumToProto(o.RoutingMode))
	return p
}

// NetworkToProto converts a Network resource to its proto representation.
func NetworkToProto(resource *beta.Network) *betapb.ComputeBetaNetwork {
	p := &betapb.ComputeBetaNetwork{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGatewayIpv4(dcl.ValueOrEmptyString(resource.GatewayIPv4))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAutoCreateSubnetworks(dcl.ValueOrEmptyBool(resource.AutoCreateSubnetworks))
	p.SetRoutingConfig(ComputeBetaNetworkRoutingConfigToProto(resource.RoutingConfig))
	p.SetMtu(dcl.ValueOrEmptyInt64(resource.Mtu))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetSelfLinkWithId(dcl.ValueOrEmptyString(resource.SelfLinkWithId))

	return p
}

// applyNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) applyNetwork(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaNetworkRequest) (*betapb.ComputeBetaNetwork, error) {
	p := ProtoToNetwork(request.GetResource())
	res, err := c.ApplyNetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkToProto(res)
	return r, nil
}

// applyComputeBetaNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) ApplyComputeBetaNetwork(ctx context.Context, request *betapb.ApplyComputeBetaNetworkRequest) (*betapb.ComputeBetaNetwork, error) {
	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNetwork(ctx, cl, request)
}

// DeleteNetwork handles the gRPC request by passing it to the underlying Network Delete() method.
func (s *NetworkServer) DeleteComputeBetaNetwork(ctx context.Context, request *betapb.DeleteComputeBetaNetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetwork(ctx, ProtoToNetwork(request.GetResource()))

}

// ListComputeBetaNetwork handles the gRPC request by passing it to the underlying NetworkList() method.
func (s *NetworkServer) ListComputeBetaNetwork(ctx context.Context, request *betapb.ListComputeBetaNetworkRequest) (*betapb.ListComputeBetaNetworkResponse, error) {
	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetwork(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaNetwork
	for _, r := range resources.Items {
		rp := NetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaNetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNetwork(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
