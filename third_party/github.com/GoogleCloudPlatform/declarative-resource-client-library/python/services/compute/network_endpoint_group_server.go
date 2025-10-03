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

// Server implements the gRPC interface for NetworkEndpointGroup.
type NetworkEndpointGroupServer struct{}

// ProtoToNetworkEndpointGroupNetworkEndpointTypeEnum converts a NetworkEndpointGroupNetworkEndpointTypeEnum enum from its proto representation.
func ProtoToComputeNetworkEndpointGroupNetworkEndpointTypeEnum(e computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum) *compute.NetworkEndpointGroupNetworkEndpointTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum_name[int32(e)]; ok {
		e := compute.NetworkEndpointGroupNetworkEndpointTypeEnum(n[len("ComputeNetworkEndpointGroupNetworkEndpointTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkEndpointGroupCloudRun converts a NetworkEndpointGroupCloudRun resource from its proto representation.
func ProtoToComputeNetworkEndpointGroupCloudRun(p *computepb.ComputeNetworkEndpointGroupCloudRun) *compute.NetworkEndpointGroupCloudRun {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkEndpointGroupCloudRun{
		Service: dcl.StringOrNil(p.Service),
		Tag:     dcl.StringOrNil(p.Tag),
		UrlMask: dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroupAppEngine converts a NetworkEndpointGroupAppEngine resource from its proto representation.
func ProtoToComputeNetworkEndpointGroupAppEngine(p *computepb.ComputeNetworkEndpointGroupAppEngine) *compute.NetworkEndpointGroupAppEngine {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkEndpointGroupAppEngine{
		Service: dcl.StringOrNil(p.Service),
		Version: dcl.StringOrNil(p.Version),
		UrlMask: dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroupCloudFunction converts a NetworkEndpointGroupCloudFunction resource from its proto representation.
func ProtoToComputeNetworkEndpointGroupCloudFunction(p *computepb.ComputeNetworkEndpointGroupCloudFunction) *compute.NetworkEndpointGroupCloudFunction {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkEndpointGroupCloudFunction{
		Function: dcl.StringOrNil(p.Function),
		UrlMask:  dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroup converts a NetworkEndpointGroup resource from its proto representation.
func ProtoToNetworkEndpointGroup(p *computepb.ComputeNetworkEndpointGroup) *compute.NetworkEndpointGroup {
	obj := &compute.NetworkEndpointGroup{
		Id:                  dcl.Int64OrNil(p.Id),
		SelfLink:            dcl.StringOrNil(p.SelfLink),
		SelfLinkWithId:      dcl.StringOrNil(p.SelfLinkWithId),
		Name:                dcl.StringOrNil(p.Name),
		Description:         dcl.StringOrNil(p.Description),
		NetworkEndpointType: ProtoToComputeNetworkEndpointGroupNetworkEndpointTypeEnum(p.GetNetworkEndpointType()),
		Size:                dcl.Int64OrNil(p.Size),
		Location:            dcl.StringOrNil(p.Location),
		Network:             dcl.StringOrNil(p.Network),
		Subnetwork:          dcl.StringOrNil(p.Subnetwork),
		DefaultPort:         dcl.Int64OrNil(p.DefaultPort),
		CloudRun:            ProtoToComputeNetworkEndpointGroupCloudRun(p.GetCloudRun()),
		AppEngine:           ProtoToComputeNetworkEndpointGroupAppEngine(p.GetAppEngine()),
		CloudFunction:       ProtoToComputeNetworkEndpointGroupCloudFunction(p.GetCloudFunction()),
		Project:             dcl.StringOrNil(p.Project),
	}
	return obj
}

// NetworkEndpointGroupNetworkEndpointTypeEnumToProto converts a NetworkEndpointGroupNetworkEndpointTypeEnum enum to its proto representation.
func ComputeNetworkEndpointGroupNetworkEndpointTypeEnumToProto(e *compute.NetworkEndpointGroupNetworkEndpointTypeEnum) computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum {
	if e == nil {
		return computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum(0)
	}
	if v, ok := computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum_value["NetworkEndpointGroupNetworkEndpointTypeEnum"+string(*e)]; ok {
		return computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum(v)
	}
	return computepb.ComputeNetworkEndpointGroupNetworkEndpointTypeEnum(0)
}

// NetworkEndpointGroupCloudRunToProto converts a NetworkEndpointGroupCloudRun resource to its proto representation.
func ComputeNetworkEndpointGroupCloudRunToProto(o *compute.NetworkEndpointGroupCloudRun) *computepb.ComputeNetworkEndpointGroupCloudRun {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkEndpointGroupCloudRun{
		Service: dcl.ValueOrEmptyString(o.Service),
		Tag:     dcl.ValueOrEmptyString(o.Tag),
		UrlMask: dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupAppEngineToProto converts a NetworkEndpointGroupAppEngine resource to its proto representation.
func ComputeNetworkEndpointGroupAppEngineToProto(o *compute.NetworkEndpointGroupAppEngine) *computepb.ComputeNetworkEndpointGroupAppEngine {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkEndpointGroupAppEngine{
		Service: dcl.ValueOrEmptyString(o.Service),
		Version: dcl.ValueOrEmptyString(o.Version),
		UrlMask: dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupCloudFunctionToProto converts a NetworkEndpointGroupCloudFunction resource to its proto representation.
func ComputeNetworkEndpointGroupCloudFunctionToProto(o *compute.NetworkEndpointGroupCloudFunction) *computepb.ComputeNetworkEndpointGroupCloudFunction {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkEndpointGroupCloudFunction{
		Function: dcl.ValueOrEmptyString(o.Function),
		UrlMask:  dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupToProto converts a NetworkEndpointGroup resource to its proto representation.
func NetworkEndpointGroupToProto(resource *compute.NetworkEndpointGroup) *computepb.ComputeNetworkEndpointGroup {
	p := &computepb.ComputeNetworkEndpointGroup{
		Id:                  dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:            dcl.ValueOrEmptyString(resource.SelfLink),
		SelfLinkWithId:      dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		Name:                dcl.ValueOrEmptyString(resource.Name),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		NetworkEndpointType: ComputeNetworkEndpointGroupNetworkEndpointTypeEnumToProto(resource.NetworkEndpointType),
		Size:                dcl.ValueOrEmptyInt64(resource.Size),
		Location:            dcl.ValueOrEmptyString(resource.Location),
		Network:             dcl.ValueOrEmptyString(resource.Network),
		Subnetwork:          dcl.ValueOrEmptyString(resource.Subnetwork),
		DefaultPort:         dcl.ValueOrEmptyInt64(resource.DefaultPort),
		CloudRun:            ComputeNetworkEndpointGroupCloudRunToProto(resource.CloudRun),
		AppEngine:           ComputeNetworkEndpointGroupAppEngineToProto(resource.AppEngine),
		CloudFunction:       ComputeNetworkEndpointGroupCloudFunctionToProto(resource.CloudFunction),
		Project:             dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Apply() method.
func (s *NetworkEndpointGroupServer) applyNetworkEndpointGroup(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeNetworkEndpointGroupRequest) (*computepb.ComputeNetworkEndpointGroup, error) {
	p := ProtoToNetworkEndpointGroup(request.GetResource())
	res, err := c.ApplyNetworkEndpointGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkEndpointGroupToProto(res)
	return r, nil
}

// ApplyNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Apply() method.
func (s *NetworkEndpointGroupServer) ApplyComputeNetworkEndpointGroup(ctx context.Context, request *computepb.ApplyComputeNetworkEndpointGroupRequest) (*computepb.ComputeNetworkEndpointGroup, error) {
	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNetworkEndpointGroup(ctx, cl, request)
}

// DeleteNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Delete() method.
func (s *NetworkEndpointGroupServer) DeleteComputeNetworkEndpointGroup(ctx context.Context, request *computepb.DeleteComputeNetworkEndpointGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkEndpointGroup(ctx, ProtoToNetworkEndpointGroup(request.GetResource()))

}

// ListComputeNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroupList() method.
func (s *NetworkEndpointGroupServer) ListComputeNetworkEndpointGroup(ctx context.Context, request *computepb.ListComputeNetworkEndpointGroupRequest) (*computepb.ListComputeNetworkEndpointGroupResponse, error) {
	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkEndpointGroup(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeNetworkEndpointGroup
	for _, r := range resources.Items {
		rp := NetworkEndpointGroupToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeNetworkEndpointGroupResponse{Items: protos}, nil
}

func createConfigNetworkEndpointGroup(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
