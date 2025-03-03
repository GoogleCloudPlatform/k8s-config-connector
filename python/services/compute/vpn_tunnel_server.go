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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// VpnTunnelServer implements the gRPC interface for VpnTunnel.
type VpnTunnelServer struct{}

// ProtoToVpnTunnelStatusEnum converts a VpnTunnelStatusEnum enum from its proto representation.
func ProtoToComputeVpnTunnelStatusEnum(e computepb.ComputeVpnTunnelStatusEnum) *compute.VpnTunnelStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeVpnTunnelStatusEnum_name[int32(e)]; ok {
		e := compute.VpnTunnelStatusEnum(n[len("ComputeVpnTunnelStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToVpnTunnel converts a VpnTunnel resource from its proto representation.
func ProtoToVpnTunnel(p *computepb.ComputeVpnTunnel) *compute.VpnTunnel {
	obj := &compute.VpnTunnel{
		Id:                           dcl.Int64OrNil(p.GetId()),
		Name:                         dcl.StringOrNil(p.GetName()),
		Description:                  dcl.StringOrNil(p.GetDescription()),
		Location:                     dcl.StringOrNil(p.GetLocation()),
		TargetVpnGateway:             dcl.StringOrNil(p.GetTargetVpnGateway()),
		VpnGateway:                   dcl.StringOrNil(p.GetVpnGateway()),
		VpnGatewayInterface:          dcl.Int64OrNil(p.GetVpnGatewayInterface()),
		PeerExternalGateway:          dcl.StringOrNil(p.GetPeerExternalGateway()),
		PeerExternalGatewayInterface: dcl.Int64OrNil(p.GetPeerExternalGatewayInterface()),
		PeerGcpGateway:               dcl.StringOrNil(p.GetPeerGcpGateway()),
		Router:                       dcl.StringOrNil(p.GetRouter()),
		PeerIP:                       dcl.StringOrNil(p.GetPeerIp()),
		SharedSecret:                 dcl.StringOrNil(p.GetSharedSecret()),
		SharedSecretHash:             dcl.StringOrNil(p.GetSharedSecretHash()),
		Status:                       ProtoToComputeVpnTunnelStatusEnum(p.GetStatus()),
		SelfLink:                     dcl.StringOrNil(p.GetSelfLink()),
		IkeVersion:                   dcl.Int64OrNil(p.GetIkeVersion()),
		DetailedStatus:               dcl.StringOrNil(p.GetDetailedStatus()),
		Project:                      dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetLocalTrafficSelector() {
		obj.LocalTrafficSelector = append(obj.LocalTrafficSelector, r)
	}
	for _, r := range p.GetRemoteTrafficSelector() {
		obj.RemoteTrafficSelector = append(obj.RemoteTrafficSelector, r)
	}
	return obj
}

// VpnTunnelStatusEnumToProto converts a VpnTunnelStatusEnum enum to its proto representation.
func ComputeVpnTunnelStatusEnumToProto(e *compute.VpnTunnelStatusEnum) computepb.ComputeVpnTunnelStatusEnum {
	if e == nil {
		return computepb.ComputeVpnTunnelStatusEnum(0)
	}
	if v, ok := computepb.ComputeVpnTunnelStatusEnum_value["VpnTunnelStatusEnum"+string(*e)]; ok {
		return computepb.ComputeVpnTunnelStatusEnum(v)
	}
	return computepb.ComputeVpnTunnelStatusEnum(0)
}

// VpnTunnelToProto converts a VpnTunnel resource to its proto representation.
func VpnTunnelToProto(resource *compute.VpnTunnel) *computepb.ComputeVpnTunnel {
	p := &computepb.ComputeVpnTunnel{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetTargetVpnGateway(dcl.ValueOrEmptyString(resource.TargetVpnGateway))
	p.SetVpnGateway(dcl.ValueOrEmptyString(resource.VpnGateway))
	p.SetVpnGatewayInterface(dcl.ValueOrEmptyInt64(resource.VpnGatewayInterface))
	p.SetPeerExternalGateway(dcl.ValueOrEmptyString(resource.PeerExternalGateway))
	p.SetPeerExternalGatewayInterface(dcl.ValueOrEmptyInt64(resource.PeerExternalGatewayInterface))
	p.SetPeerGcpGateway(dcl.ValueOrEmptyString(resource.PeerGcpGateway))
	p.SetRouter(dcl.ValueOrEmptyString(resource.Router))
	p.SetPeerIp(dcl.ValueOrEmptyString(resource.PeerIP))
	p.SetSharedSecret(dcl.ValueOrEmptyString(resource.SharedSecret))
	p.SetSharedSecretHash(dcl.ValueOrEmptyString(resource.SharedSecretHash))
	p.SetStatus(ComputeVpnTunnelStatusEnumToProto(resource.Status))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetIkeVersion(dcl.ValueOrEmptyInt64(resource.IkeVersion))
	p.SetDetailedStatus(dcl.ValueOrEmptyString(resource.DetailedStatus))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sLocalTrafficSelector := make([]string, len(resource.LocalTrafficSelector))
	for i, r := range resource.LocalTrafficSelector {
		sLocalTrafficSelector[i] = r
	}
	p.SetLocalTrafficSelector(sLocalTrafficSelector)
	sRemoteTrafficSelector := make([]string, len(resource.RemoteTrafficSelector))
	for i, r := range resource.RemoteTrafficSelector {
		sRemoteTrafficSelector[i] = r
	}
	p.SetRemoteTrafficSelector(sRemoteTrafficSelector)

	return p
}

// applyVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Apply() method.
func (s *VpnTunnelServer) applyVpnTunnel(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeVpnTunnelRequest) (*computepb.ComputeVpnTunnel, error) {
	p := ProtoToVpnTunnel(request.GetResource())
	res, err := c.ApplyVpnTunnel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VpnTunnelToProto(res)
	return r, nil
}

// applyComputeVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Apply() method.
func (s *VpnTunnelServer) ApplyComputeVpnTunnel(ctx context.Context, request *computepb.ApplyComputeVpnTunnelRequest) (*computepb.ComputeVpnTunnel, error) {
	cl, err := createConfigVpnTunnel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyVpnTunnel(ctx, cl, request)
}

// DeleteVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnel Delete() method.
func (s *VpnTunnelServer) DeleteComputeVpnTunnel(ctx context.Context, request *computepb.DeleteComputeVpnTunnelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVpnTunnel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVpnTunnel(ctx, ProtoToVpnTunnel(request.GetResource()))

}

// ListComputeVpnTunnel handles the gRPC request by passing it to the underlying VpnTunnelList() method.
func (s *VpnTunnelServer) ListComputeVpnTunnel(ctx context.Context, request *computepb.ListComputeVpnTunnelRequest) (*computepb.ListComputeVpnTunnelResponse, error) {
	cl, err := createConfigVpnTunnel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVpnTunnel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeVpnTunnel
	for _, r := range resources.Items {
		rp := VpnTunnelToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeVpnTunnelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigVpnTunnel(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
