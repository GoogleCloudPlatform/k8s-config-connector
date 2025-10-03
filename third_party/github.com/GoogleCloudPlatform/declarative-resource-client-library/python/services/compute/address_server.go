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

// Server implements the gRPC interface for Address.
type AddressServer struct{}

// ProtoToAddressStatusEnum converts a AddressStatusEnum enum from its proto representation.
func ProtoToComputeAddressStatusEnum(e computepb.ComputeAddressStatusEnum) *compute.AddressStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAddressStatusEnum_name[int32(e)]; ok {
		e := compute.AddressStatusEnum(n[len("ComputeAddressStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressNetworkTierEnum converts a AddressNetworkTierEnum enum from its proto representation.
func ProtoToComputeAddressNetworkTierEnum(e computepb.ComputeAddressNetworkTierEnum) *compute.AddressNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAddressNetworkTierEnum_name[int32(e)]; ok {
		e := compute.AddressNetworkTierEnum(n[len("ComputeAddressNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressIPVersionEnum converts a AddressIPVersionEnum enum from its proto representation.
func ProtoToComputeAddressIPVersionEnum(e computepb.ComputeAddressIPVersionEnum) *compute.AddressIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAddressIPVersionEnum_name[int32(e)]; ok {
		e := compute.AddressIPVersionEnum(n[len("ComputeAddressIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressAddressTypeEnum converts a AddressAddressTypeEnum enum from its proto representation.
func ProtoToComputeAddressAddressTypeEnum(e computepb.ComputeAddressAddressTypeEnum) *compute.AddressAddressTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAddressAddressTypeEnum_name[int32(e)]; ok {
		e := compute.AddressAddressTypeEnum(n[len("ComputeAddressAddressTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddressPurposeEnum converts a AddressPurposeEnum enum from its proto representation.
func ProtoToComputeAddressPurposeEnum(e computepb.ComputeAddressPurposeEnum) *compute.AddressPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeAddressPurposeEnum_name[int32(e)]; ok {
		e := compute.AddressPurposeEnum(n[len("ComputeAddressPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAddress converts a Address resource from its proto representation.
func ProtoToAddress(p *computepb.ComputeAddress) *compute.Address {
	obj := &compute.Address{
		Id:                dcl.Int64OrNil(p.Id),
		Name:              dcl.StringOrNil(p.Name),
		Description:       dcl.StringOrNil(p.Description),
		Address:           dcl.StringOrNil(p.Address),
		PrefixLength:      dcl.Int64OrNil(p.PrefixLength),
		Status:            ProtoToComputeAddressStatusEnum(p.GetStatus()),
		Region:            dcl.StringOrNil(p.Region),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		NetworkTier:       ProtoToComputeAddressNetworkTierEnum(p.GetNetworkTier()),
		IPVersion:         ProtoToComputeAddressIPVersionEnum(p.GetIpVersion()),
		AddressType:       ProtoToComputeAddressAddressTypeEnum(p.GetAddressType()),
		Purpose:           ProtoToComputeAddressPurposeEnum(p.GetPurpose()),
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
func ComputeAddressStatusEnumToProto(e *compute.AddressStatusEnum) computepb.ComputeAddressStatusEnum {
	if e == nil {
		return computepb.ComputeAddressStatusEnum(0)
	}
	if v, ok := computepb.ComputeAddressStatusEnum_value["AddressStatusEnum"+string(*e)]; ok {
		return computepb.ComputeAddressStatusEnum(v)
	}
	return computepb.ComputeAddressStatusEnum(0)
}

// AddressNetworkTierEnumToProto converts a AddressNetworkTierEnum enum to its proto representation.
func ComputeAddressNetworkTierEnumToProto(e *compute.AddressNetworkTierEnum) computepb.ComputeAddressNetworkTierEnum {
	if e == nil {
		return computepb.ComputeAddressNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeAddressNetworkTierEnum_value["AddressNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeAddressNetworkTierEnum(v)
	}
	return computepb.ComputeAddressNetworkTierEnum(0)
}

// AddressIPVersionEnumToProto converts a AddressIPVersionEnum enum to its proto representation.
func ComputeAddressIPVersionEnumToProto(e *compute.AddressIPVersionEnum) computepb.ComputeAddressIPVersionEnum {
	if e == nil {
		return computepb.ComputeAddressIPVersionEnum(0)
	}
	if v, ok := computepb.ComputeAddressIPVersionEnum_value["AddressIPVersionEnum"+string(*e)]; ok {
		return computepb.ComputeAddressIPVersionEnum(v)
	}
	return computepb.ComputeAddressIPVersionEnum(0)
}

// AddressAddressTypeEnumToProto converts a AddressAddressTypeEnum enum to its proto representation.
func ComputeAddressAddressTypeEnumToProto(e *compute.AddressAddressTypeEnum) computepb.ComputeAddressAddressTypeEnum {
	if e == nil {
		return computepb.ComputeAddressAddressTypeEnum(0)
	}
	if v, ok := computepb.ComputeAddressAddressTypeEnum_value["AddressAddressTypeEnum"+string(*e)]; ok {
		return computepb.ComputeAddressAddressTypeEnum(v)
	}
	return computepb.ComputeAddressAddressTypeEnum(0)
}

// AddressPurposeEnumToProto converts a AddressPurposeEnum enum to its proto representation.
func ComputeAddressPurposeEnumToProto(e *compute.AddressPurposeEnum) computepb.ComputeAddressPurposeEnum {
	if e == nil {
		return computepb.ComputeAddressPurposeEnum(0)
	}
	if v, ok := computepb.ComputeAddressPurposeEnum_value["AddressPurposeEnum"+string(*e)]; ok {
		return computepb.ComputeAddressPurposeEnum(v)
	}
	return computepb.ComputeAddressPurposeEnum(0)
}

// AddressToProto converts a Address resource to its proto representation.
func AddressToProto(resource *compute.Address) *computepb.ComputeAddress {
	p := &computepb.ComputeAddress{
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Address:           dcl.ValueOrEmptyString(resource.Address),
		PrefixLength:      dcl.ValueOrEmptyInt64(resource.PrefixLength),
		Status:            ComputeAddressStatusEnumToProto(resource.Status),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		NetworkTier:       ComputeAddressNetworkTierEnumToProto(resource.NetworkTier),
		IpVersion:         ComputeAddressIPVersionEnumToProto(resource.IPVersion),
		AddressType:       ComputeAddressAddressTypeEnumToProto(resource.AddressType),
		Purpose:           ComputeAddressPurposeEnumToProto(resource.Purpose),
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
func (s *AddressServer) applyAddress(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeAddressRequest) (*computepb.ComputeAddress, error) {
	p := ProtoToAddress(request.GetResource())
	res, err := c.ApplyAddress(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AddressToProto(res)
	return r, nil
}

// ApplyAddress handles the gRPC request by passing it to the underlying Address Apply() method.
func (s *AddressServer) ApplyComputeAddress(ctx context.Context, request *computepb.ApplyComputeAddressRequest) (*computepb.ComputeAddress, error) {
	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAddress(ctx, cl, request)
}

// DeleteAddress handles the gRPC request by passing it to the underlying Address Delete() method.
func (s *AddressServer) DeleteComputeAddress(ctx context.Context, request *computepb.DeleteComputeAddressRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAddress(ctx, ProtoToAddress(request.GetResource()))

}

// ListComputeAddress handles the gRPC request by passing it to the underlying AddressList() method.
func (s *AddressServer) ListComputeAddress(ctx context.Context, request *computepb.ListComputeAddressRequest) (*computepb.ListComputeAddressResponse, error) {
	cl, err := createConfigAddress(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAddress(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeAddress
	for _, r := range resources.Items {
		rp := AddressToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeAddressResponse{Items: protos}, nil
}

func createConfigAddress(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
