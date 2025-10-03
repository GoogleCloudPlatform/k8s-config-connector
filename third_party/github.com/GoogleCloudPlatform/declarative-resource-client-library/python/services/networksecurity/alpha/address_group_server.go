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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/alpha/networksecurity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
)

// AddressGroupServer implements the gRPC interface for AddressGroup.
type AddressGroupServer struct{}

// ProtoToAddressGroupTypeEnum converts a AddressGroupTypeEnum enum from its proto representation.
func ProtoToNetworksecurityAlphaAddressGroupTypeEnum(e alphapb.NetworksecurityAlphaAddressGroupTypeEnum) *alpha.AddressGroupTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworksecurityAlphaAddressGroupTypeEnum_name[int32(e)]; ok {
		e := alpha.AddressGroupTypeEnum(n[len("NetworksecurityAlphaAddressGroupTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressGroup converts a AddressGroup resource from its proto representation.
func ProtoToAddressGroup(p *alphapb.NetworksecurityAlphaAddressGroup) *alpha.AddressGroup {
	obj := &alpha.AddressGroup{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Type:        ProtoToNetworksecurityAlphaAddressGroupTypeEnum(p.GetType()),
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
func NetworksecurityAlphaAddressGroupTypeEnumToProto(e *alpha.AddressGroupTypeEnum) alphapb.NetworksecurityAlphaAddressGroupTypeEnum {
	if e == nil {
		return alphapb.NetworksecurityAlphaAddressGroupTypeEnum(0)
	}
	if v, ok := alphapb.NetworksecurityAlphaAddressGroupTypeEnum_value["AddressGroupTypeEnum"+string(*e)]; ok {
		return alphapb.NetworksecurityAlphaAddressGroupTypeEnum(v)
	}
	return alphapb.NetworksecurityAlphaAddressGroupTypeEnum(0)
}

// AddressGroupToProto converts a AddressGroup resource to its proto representation.
func AddressGroupToProto(resource *alpha.AddressGroup) *alphapb.NetworksecurityAlphaAddressGroup {
	p := &alphapb.NetworksecurityAlphaAddressGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(NetworksecurityAlphaAddressGroupTypeEnumToProto(resource.Type))
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
func (s *AddressGroupServer) applyAddressGroup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaAddressGroupRequest) (*alphapb.NetworksecurityAlphaAddressGroup, error) {
	p := ProtoToAddressGroup(request.GetResource())
	res, err := c.ApplyAddressGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AddressGroupToProto(res)
	return r, nil
}

// applyNetworksecurityAlphaAddressGroup handles the gRPC request by passing it to the underlying AddressGroup Apply() method.
func (s *AddressGroupServer) ApplyNetworksecurityAlphaAddressGroup(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaAddressGroupRequest) (*alphapb.NetworksecurityAlphaAddressGroup, error) {
	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAddressGroup(ctx, cl, request)
}

// DeleteAddressGroup handles the gRPC request by passing it to the underlying AddressGroup Delete() method.
func (s *AddressGroupServer) DeleteNetworksecurityAlphaAddressGroup(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaAddressGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAddressGroup(ctx, ProtoToAddressGroup(request.GetResource()))

}

// ListNetworksecurityAlphaAddressGroup handles the gRPC request by passing it to the underlying AddressGroupList() method.
func (s *AddressGroupServer) ListNetworksecurityAlphaAddressGroup(ctx context.Context, request *alphapb.ListNetworksecurityAlphaAddressGroupRequest) (*alphapb.ListNetworksecurityAlphaAddressGroupResponse, error) {
	cl, err := createConfigAddressGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAddressGroup(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaAddressGroup
	for _, r := range resources.Items {
		rp := AddressGroupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworksecurityAlphaAddressGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAddressGroup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
