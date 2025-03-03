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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// NetworkServer implements the gRPC interface for Network.
type NetworkServer struct{}

// ProtoToNetworkRoutingConfigRoutingModeEnum converts a NetworkRoutingConfigRoutingModeEnum enum from its proto representation.
func ProtoToComputeAlphaNetworkRoutingConfigRoutingModeEnum(e alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum) *alpha.NetworkRoutingConfigRoutingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum_name[int32(e)]; ok {
		e := alpha.NetworkRoutingConfigRoutingModeEnum(n[len("ComputeAlphaNetworkRoutingConfigRoutingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkRoutingConfig converts a NetworkRoutingConfig object from its proto representation.
func ProtoToComputeAlphaNetworkRoutingConfig(p *alphapb.ComputeAlphaNetworkRoutingConfig) *alpha.NetworkRoutingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NetworkRoutingConfig{
		RoutingMode: ProtoToComputeAlphaNetworkRoutingConfigRoutingModeEnum(p.GetRoutingMode()),
	}
	return obj
}

// ProtoToNetwork converts a Network resource from its proto representation.
func ProtoToNetwork(p *alphapb.ComputeAlphaNetwork) *alpha.Network {
	obj := &alpha.Network{
		Description:           dcl.StringOrNil(p.GetDescription()),
		GatewayIPv4:           dcl.StringOrNil(p.GetGatewayIpv4()),
		Name:                  dcl.StringOrNil(p.GetName()),
		AutoCreateSubnetworks: dcl.Bool(p.GetAutoCreateSubnetworks()),
		RoutingConfig:         ProtoToComputeAlphaNetworkRoutingConfig(p.GetRoutingConfig()),
		Mtu:                   dcl.Int64OrNil(p.GetMtu()),
		Project:               dcl.StringOrNil(p.GetProject()),
		SelfLink:              dcl.StringOrNil(p.GetSelfLink()),
		SelfLinkWithId:        dcl.StringOrNil(p.GetSelfLinkWithId()),
	}
	return obj
}

// NetworkRoutingConfigRoutingModeEnumToProto converts a NetworkRoutingConfigRoutingModeEnum enum to its proto representation.
func ComputeAlphaNetworkRoutingConfigRoutingModeEnumToProto(e *alpha.NetworkRoutingConfigRoutingModeEnum) alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum {
	if e == nil {
		return alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum_value["NetworkRoutingConfigRoutingModeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum(v)
	}
	return alphapb.ComputeAlphaNetworkRoutingConfigRoutingModeEnum(0)
}

// NetworkRoutingConfigToProto converts a NetworkRoutingConfig object to its proto representation.
func ComputeAlphaNetworkRoutingConfigToProto(o *alpha.NetworkRoutingConfig) *alphapb.ComputeAlphaNetworkRoutingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaNetworkRoutingConfig{}
	p.SetRoutingMode(ComputeAlphaNetworkRoutingConfigRoutingModeEnumToProto(o.RoutingMode))
	return p
}

// NetworkToProto converts a Network resource to its proto representation.
func NetworkToProto(resource *alpha.Network) *alphapb.ComputeAlphaNetwork {
	p := &alphapb.ComputeAlphaNetwork{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGatewayIpv4(dcl.ValueOrEmptyString(resource.GatewayIPv4))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAutoCreateSubnetworks(dcl.ValueOrEmptyBool(resource.AutoCreateSubnetworks))
	p.SetRoutingConfig(ComputeAlphaNetworkRoutingConfigToProto(resource.RoutingConfig))
	p.SetMtu(dcl.ValueOrEmptyInt64(resource.Mtu))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetSelfLinkWithId(dcl.ValueOrEmptyString(resource.SelfLinkWithId))

	return p
}

// applyNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) applyNetwork(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaNetworkRequest) (*alphapb.ComputeAlphaNetwork, error) {
	p := ProtoToNetwork(request.GetResource())
	res, err := c.ApplyNetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkToProto(res)
	return r, nil
}

// applyComputeAlphaNetwork handles the gRPC request by passing it to the underlying Network Apply() method.
func (s *NetworkServer) ApplyComputeAlphaNetwork(ctx context.Context, request *alphapb.ApplyComputeAlphaNetworkRequest) (*alphapb.ComputeAlphaNetwork, error) {
	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNetwork(ctx, cl, request)
}

// DeleteNetwork handles the gRPC request by passing it to the underlying Network Delete() method.
func (s *NetworkServer) DeleteComputeAlphaNetwork(ctx context.Context, request *alphapb.DeleteComputeAlphaNetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetwork(ctx, ProtoToNetwork(request.GetResource()))

}

// ListComputeAlphaNetwork handles the gRPC request by passing it to the underlying NetworkList() method.
func (s *NetworkServer) ListComputeAlphaNetwork(ctx context.Context, request *alphapb.ListComputeAlphaNetworkRequest) (*alphapb.ListComputeAlphaNetworkResponse, error) {
	cl, err := createConfigNetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetwork(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaNetwork
	for _, r := range resources.Items {
		rp := NetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaNetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNetwork(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
