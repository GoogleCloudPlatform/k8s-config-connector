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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/beta/networksecurity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
)

// AddressGroupServer implements the gRPC interface for AddressGroup.
type AddressGroupServer struct{}

// ProtoToAddressGroupTypeEnum converts a AddressGroupTypeEnum enum from its proto representation.
func ProtoToNetworksecurityBetaAddressGroupTypeEnum(e betapb.NetworksecurityBetaAddressGroupTypeEnum) *beta.AddressGroupTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworksecurityBetaAddressGroupTypeEnum_name[int32(e)]; ok {
		e := beta.AddressGroupTypeEnum(n[len("NetworksecurityBetaAddressGroupTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressGroup converts a AddressGroup resource from its proto representation.
func ProtoToAddressGroup(p *betapb.NetworksecurityBetaAddressGroup) *beta.AddressGroup {
	obj := &beta.AddressGroup{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Type:        ProtoToNetworksecurityBetaAddressGroupTypeEnum(p.GetType()),
		Capacity:    dcl.Int64OrNil(p.GetCapacity()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, r)
	}
	return obj
}

// AddressGroupTypeEnumToProto converts a AddressGroupTypeEnum enum to its proto representation.
func NetworksecurityBetaAddressGroupTypeEnumToProto(e *beta.AddressGroupTypeEnum) betapb.NetworksecurityBetaAddressGroupTypeEnum {
	if e == nil {
		return betapb.NetworksecurityBetaAddressGroupTypeEnum(0)
	}
	if v, ok := betapb.NetworksecurityBetaAddressGroupTypeEnum_value["AddressGroupTypeEnum"+string(*e)]; ok {
		return betapb.NetworksecurityBetaAddressGroupTypeEnum(v)
	}
	return betapb.NetworksecurityBetaAddressGroupTypeEnum(0)
}

// AddressGroupToProto converts a AddressGroup resource to its proto representation.
func AddressGroupToProto(resource *beta.AddressGroup) *betapb.NetworksecurityBetaAddressGroup {
	p := &betapb.NetworksecurityBetaAddressGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(NetworksecurityBetaAddressGroupTypeEnumToProto(resource.Type))
	p.SetCapacity(dcl.ValueOrEmptyInt64(resource.Capacity))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sItems := make([]string, len(resource.Items))
	for i, r := range resource.Items {
		sItems[i] = r
	}
	p.SetItems(sItems)

	return p
}

// applyAddressGroup handles the gRPC request by passing it to the underlying AddressGroup Apply() method.
func (s *AddressGroupServer) applyAddressGroup(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaAddressGroupRequest) (*betapb.NetworksecurityBetaAddressGroup, error) {
	p := ProtoToAddressGroup(request.GetResource())
	res, err := c.ApplyAddressGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AddressGroupToProto(res)
	return r, nil
}

// applyNetworksecurityBetaAddressGroup handles the gRPC request by passing it to the underlying AddressGroup Apply() method.
func (s *AddressGroupServer) ApplyNetworksecurityBetaAddressGroup(ctx context.Context, request *betapb.ApplyNetworksecurityBetaAddressGroupRequest) (*betapb.NetworksecurityBetaAddressGroup, error) {
	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAddressGroup(ctx, cl, request)
}

// DeleteAddressGroup handles the gRPC request by passing it to the underlying AddressGroup Delete() method.
func (s *AddressGroupServer) DeleteNetworksecurityBetaAddressGroup(ctx context.Context, request *betapb.DeleteNetworksecurityBetaAddressGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAddressGroup(ctx, ProtoToAddressGroup(request.GetResource()))

}

// ListNetworksecurityBetaAddressGroup handles the gRPC request by passing it to the underlying AddressGroupList() method.
func (s *AddressGroupServer) ListNetworksecurityBetaAddressGroup(ctx context.Context, request *betapb.ListNetworksecurityBetaAddressGroupRequest) (*betapb.ListNetworksecurityBetaAddressGroupResponse, error) {
	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAddressGroup(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaAddressGroup
	for _, r := range resources.Items {
		rp := AddressGroupToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworksecurityBetaAddressGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAddressGroup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
