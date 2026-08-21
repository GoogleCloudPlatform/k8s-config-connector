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

// Server implements the gRPC interface for Address.
type AddressServer struct{}

// ProtoToAddressStatusEnum converts a AddressStatusEnum enum from its proto representation.
func ProtoToComputeBetaAddressStatusEnum(e betapb.ComputeBetaAddressStatusEnum) *beta.AddressStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAddressStatusEnum_name[int32(e)]; ok {
		e := beta.AddressStatusEnum(n[len("ComputeBetaAddressStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressNetworkTierEnum converts a AddressNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaAddressNetworkTierEnum(e betapb.ComputeBetaAddressNetworkTierEnum) *beta.AddressNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAddressNetworkTierEnum_name[int32(e)]; ok {
		e := beta.AddressNetworkTierEnum(n[len("ComputeBetaAddressNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressIPVersionEnum converts a AddressIPVersionEnum enum from its proto representation.
func ProtoToComputeBetaAddressIPVersionEnum(e betapb.ComputeBetaAddressIPVersionEnum) *beta.AddressIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAddressIPVersionEnum_name[int32(e)]; ok {
		e := beta.AddressIPVersionEnum(n[len("ComputeBetaAddressIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressAddressTypeEnum converts a AddressAddressTypeEnum enum from its proto representation.
func ProtoToComputeBetaAddressAddressTypeEnum(e betapb.ComputeBetaAddressAddressTypeEnum) *beta.AddressAddressTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAddressAddressTypeEnum_name[int32(e)]; ok {
		e := beta.AddressAddressTypeEnum(n[len("ComputeBetaAddressAddressTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressPurposeEnum converts a AddressPurposeEnum enum from its proto representation.
func ProtoToComputeBetaAddressPurposeEnum(e betapb.ComputeBetaAddressPurposeEnum) *beta.AddressPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaAddressPurposeEnum_name[int32(e)]; ok {
		e := beta.AddressPurposeEnum(n[len("ComputeBetaAddressPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddress converts a Address resource from its proto representation.
func ProtoToAddress(p *betapb.ComputeBetaAddress) *beta.Address {
	obj := &beta.Address{
		Id:                dcl.Int64OrNil(p.Id),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Address:           dcl.StringOrNil(p.Address),
		PrefixLength:      dcl.Int64OrNil(p.PrefixLength),
		Status:            ProtoToComputeBetaAddressStatusEnum(p.GetStatus()),
		Region:            dcl.StringOrNil(p.Region),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		NetworkTier:       ProtoToComputeBetaAddressNetworkTierEnum(p.GetNetworkTier()),
		IPVersion:         ProtoToComputeBetaAddressIPVersionEnum(p.GetIpVersion()),
		AddressType:       ProtoToComputeBetaAddressAddressTypeEnum(p.GetAddressType()),
		Purpose:           ProtoToComputeBetaAddressPurposeEnum(p.GetPurpose()),
		Subnetwork:        dcl.StringOrNil(p.Subnetwork),
		Network:           dcl.StringOrNil(p.Network),
		Project:           dcl.StringOrNil(p.Project),
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		LabelFingerprint:  dcl.StringOrNil(p.LabelFingerprint),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetUsers() {
		obj.Users = append(obj.Users, r)
	}
	return obj
}

// AddressStatusEnumToProto converts a AddressStatusEnum enum to its proto representation.
func ComputeBetaAddressStatusEnumToProto(e *beta.AddressStatusEnum) betapb.ComputeBetaAddressStatusEnum {
	if e == nil {
		return betapb.ComputeBetaAddressStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaAddressStatusEnum_value["AddressStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAddressStatusEnum(v)
	}
	return betapb.ComputeBetaAddressStatusEnum(0)
}

// AddressNetworkTierEnumToProto converts a AddressNetworkTierEnum enum to its proto representation.
func ComputeBetaAddressNetworkTierEnumToProto(e *beta.AddressNetworkTierEnum) betapb.ComputeBetaAddressNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaAddressNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaAddressNetworkTierEnum_value["AddressNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAddressNetworkTierEnum(v)
	}
	return betapb.ComputeBetaAddressNetworkTierEnum(0)
}

// AddressIPVersionEnumToProto converts a AddressIPVersionEnum enum to its proto representation.
func ComputeBetaAddressIPVersionEnumToProto(e *beta.AddressIPVersionEnum) betapb.ComputeBetaAddressIPVersionEnum {
	if e == nil {
		return betapb.ComputeBetaAddressIPVersionEnum(0)
	}
	if v, ok := betapb.ComputeBetaAddressIPVersionEnum_value["AddressIPVersionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAddressIPVersionEnum(v)
	}
	return betapb.ComputeBetaAddressIPVersionEnum(0)
}

// AddressAddressTypeEnumToProto converts a AddressAddressTypeEnum enum to its proto representation.
func ComputeBetaAddressAddressTypeEnumToProto(e *beta.AddressAddressTypeEnum) betapb.ComputeBetaAddressAddressTypeEnum {
	if e == nil {
		return betapb.ComputeBetaAddressAddressTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaAddressAddressTypeEnum_value["AddressAddressTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAddressAddressTypeEnum(v)
	}
	return betapb.ComputeBetaAddressAddressTypeEnum(0)
}

// AddressPurposeEnumToProto converts a AddressPurposeEnum enum to its proto representation.
func ComputeBetaAddressPurposeEnumToProto(e *beta.AddressPurposeEnum) betapb.ComputeBetaAddressPurposeEnum {
	if e == nil {
		return betapb.ComputeBetaAddressPurposeEnum(0)
	}
	if v, ok := betapb.ComputeBetaAddressPurposeEnum_value["AddressPurposeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaAddressPurposeEnum(v)
	}
	return betapb.ComputeBetaAddressPurposeEnum(0)
}

// AddressToProto converts a Address resource to its proto representation.
func AddressToProto(resource *beta.Address) *betapb.ComputeBetaAddress {
	p := &betapb.ComputeBetaAddress{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Address:           dcl.ValueOrEmptyString(resource.Address),
		PrefixLength:      dcl.ValueOrEmptyInt64(resource.PrefixLength),
		Status:            ComputeBetaAddressStatusEnumToProto(resource.Status),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		NetworkTier:       ComputeBetaAddressNetworkTierEnumToProto(resource.NetworkTier),
		IpVersion:         ComputeBetaAddressIPVersionEnumToProto(resource.IPVersion),
		AddressType:       ComputeBetaAddressAddressTypeEnumToProto(resource.AddressType),
		Purpose:           ComputeBetaAddressPurposeEnumToProto(resource.Purpose),
		Subnetwork:        dcl.ValueOrEmptyString(resource.Subnetwork),
		Network:           dcl.ValueOrEmptyString(resource.Network),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		LabelFingerprint:  dcl.ValueOrEmptyString(resource.LabelFingerprint),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Users {
		p.Users = append(p.Users, r)
	}

	return p
}

// ApplyAddress handles the gRPC request by passing it to the underlying Address Apply() method.
func (s *AddressServer) applyAddress(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaAddressRequest) (*betapb.ComputeBetaAddress, error) {
	p := ProtoToAddress(request.GetResource())
	res, err := c.ApplyAddress(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AddressToProto(res)
	return r, nil
}

// ApplyAddress handles the gRPC request by passing it to the underlying Address Apply() method.
func (s *AddressServer) ApplyComputeBetaAddress(ctx context.Context, request *betapb.ApplyComputeBetaAddressRequest) (*betapb.ComputeBetaAddress, error) {
	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAddress(ctx, cl, request)
}

// DeleteAddress handles the gRPC request by passing it to the underlying Address Delete() method.
func (s *AddressServer) DeleteComputeBetaAddress(ctx context.Context, request *betapb.DeleteComputeBetaAddressRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAddress(ctx, ProtoToAddress(request.GetResource()))

}

// ListComputeBetaAddress handles the gRPC request by passing it to the underlying AddressList() method.
func (s *AddressServer) ListComputeBetaAddress(ctx context.Context, request *betapb.ListComputeBetaAddressRequest) (*betapb.ListComputeBetaAddressResponse, error) {
	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAddress(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaAddress
	for _, r := range resources.Items {
		rp := AddressToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaAddressResponse{Items: protos}, nil
}

func createConfigAddress(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
