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

// Server implements the gRPC interface for NetworkEndpointGroup.
type NetworkEndpointGroupServer struct{}

// ProtoToNetworkEndpointGroupNetworkEndpointTypeEnum converts a NetworkEndpointGroupNetworkEndpointTypeEnum enum from its proto representation.
func ProtoToComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum(e betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum) *beta.NetworkEndpointGroupNetworkEndpointTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum_name[int32(e)]; ok {
		e := beta.NetworkEndpointGroupNetworkEndpointTypeEnum(n[len("ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkEndpointGroupCloudRun converts a NetworkEndpointGroupCloudRun resource from its proto representation.
func ProtoToComputeBetaNetworkEndpointGroupCloudRun(p *betapb.ComputeBetaNetworkEndpointGroupCloudRun) *beta.NetworkEndpointGroupCloudRun {
	if p == nil {
		return nil
	}
	obj := &beta.NetworkEndpointGroupCloudRun{
		Service: dcl.StringOrNil(p.Service),
		Tag:     dcl.StringOrNil(p.Tag),
		UrlMask: dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroupAppEngine converts a NetworkEndpointGroupAppEngine resource from its proto representation.
func ProtoToComputeBetaNetworkEndpointGroupAppEngine(p *betapb.ComputeBetaNetworkEndpointGroupAppEngine) *beta.NetworkEndpointGroupAppEngine {
	if p == nil {
		return nil
	}
	obj := &beta.NetworkEndpointGroupAppEngine{
		Service: dcl.StringOrNil(p.Service),
		Version: dcl.StringOrNil(p.Version),
		UrlMask: dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroupCloudFunction converts a NetworkEndpointGroupCloudFunction resource from its proto representation.
func ProtoToComputeBetaNetworkEndpointGroupCloudFunction(p *betapb.ComputeBetaNetworkEndpointGroupCloudFunction) *beta.NetworkEndpointGroupCloudFunction {
	if p == nil {
		return nil
	}
	obj := &beta.NetworkEndpointGroupCloudFunction{
		Function: dcl.StringOrNil(p.Function),
		UrlMask:  dcl.StringOrNil(p.UrlMask),
	}
	return obj
}

// ProtoToNetworkEndpointGroup converts a NetworkEndpointGroup resource from its proto representation.
func ProtoToNetworkEndpointGroup(p *betapb.ComputeBetaNetworkEndpointGroup) *beta.NetworkEndpointGroup {
	obj := &beta.NetworkEndpointGroup{
		Id:                  dcl.Int64OrNil(p.Id),
		SelfLink:            dcl.StringOrNil(p.SelfLink),
		SelfLinkWithId:      dcl.StringOrNil(p.SelfLinkWithId),
		Name:                dcl.StringOrNil(p.Name),
		Description:         dcl.StringOrNil(p.Description),
		NetworkEndpointType: ProtoToComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum(p.GetNetworkEndpointType()),
		Size:                dcl.Int64OrNil(p.Size),
		Location:            dcl.StringOrNil(p.Location),
		Network:             dcl.StringOrNil(p.Network),
		Subnetwork:          dcl.StringOrNil(p.Subnetwork),
		DefaultPort:         dcl.Int64OrNil(p.DefaultPort),
		CloudRun:            ProtoToComputeBetaNetworkEndpointGroupCloudRun(p.GetCloudRun()),
		AppEngine:           ProtoToComputeBetaNetworkEndpointGroupAppEngine(p.GetAppEngine()),
		CloudFunction:       ProtoToComputeBetaNetworkEndpointGroupCloudFunction(p.GetCloudFunction()),
		Project:             dcl.StringOrNil(p.Project),
	}
	return obj
}

// NetworkEndpointGroupNetworkEndpointTypeEnumToProto converts a NetworkEndpointGroupNetworkEndpointTypeEnum enum to its proto representation.
func ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnumToProto(e *beta.NetworkEndpointGroupNetworkEndpointTypeEnum) betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum {
	if e == nil {
		return betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum_value["NetworkEndpointGroupNetworkEndpointTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum(v)
	}
	return betapb.ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnum(0)
}

// NetworkEndpointGroupCloudRunToProto converts a NetworkEndpointGroupCloudRun resource to its proto representation.
func ComputeBetaNetworkEndpointGroupCloudRunToProto(o *beta.NetworkEndpointGroupCloudRun) *betapb.ComputeBetaNetworkEndpointGroupCloudRun {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaNetworkEndpointGroupCloudRun{
		Service: dcl.ValueOrEmptyString(o.Service),
		Tag:     dcl.ValueOrEmptyString(o.Tag),
		UrlMask: dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupAppEngineToProto converts a NetworkEndpointGroupAppEngine resource to its proto representation.
func ComputeBetaNetworkEndpointGroupAppEngineToProto(o *beta.NetworkEndpointGroupAppEngine) *betapb.ComputeBetaNetworkEndpointGroupAppEngine {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaNetworkEndpointGroupAppEngine{
		Service: dcl.ValueOrEmptyString(o.Service),
		Version: dcl.ValueOrEmptyString(o.Version),
		UrlMask: dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupCloudFunctionToProto converts a NetworkEndpointGroupCloudFunction resource to its proto representation.
func ComputeBetaNetworkEndpointGroupCloudFunctionToProto(o *beta.NetworkEndpointGroupCloudFunction) *betapb.ComputeBetaNetworkEndpointGroupCloudFunction {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaNetworkEndpointGroupCloudFunction{
		Function: dcl.ValueOrEmptyString(o.Function),
		UrlMask:  dcl.ValueOrEmptyString(o.UrlMask),
	}
	return p
}

// NetworkEndpointGroupToProto converts a NetworkEndpointGroup resource to its proto representation.
func NetworkEndpointGroupToProto(resource *beta.NetworkEndpointGroup) *betapb.ComputeBetaNetworkEndpointGroup {
	p := &betapb.ComputeBetaNetworkEndpointGroup{
		Id:                  dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:            dcl.ValueOrEmptyString(resource.SelfLink),
		SelfLinkWithId:      dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		Name:                dcl.ValueOrEmptyString(resource.Name),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		NetworkEndpointType: ComputeBetaNetworkEndpointGroupNetworkEndpointTypeEnumToProto(resource.NetworkEndpointType),
		Size:                dcl.ValueOrEmptyInt64(resource.Size),
		Location:            dcl.ValueOrEmptyString(resource.Location),
		Network:             dcl.ValueOrEmptyString(resource.Network),
		Subnetwork:          dcl.ValueOrEmptyString(resource.Subnetwork),
		DefaultPort:         dcl.ValueOrEmptyInt64(resource.DefaultPort),
		CloudRun:            ComputeBetaNetworkEndpointGroupCloudRunToProto(resource.CloudRun),
		AppEngine:           ComputeBetaNetworkEndpointGroupAppEngineToProto(resource.AppEngine),
		CloudFunction:       ComputeBetaNetworkEndpointGroupCloudFunctionToProto(resource.CloudFunction),
		Project:             dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Apply() method.
func (s *NetworkEndpointGroupServer) applyNetworkEndpointGroup(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaNetworkEndpointGroupRequest) (*betapb.ComputeBetaNetworkEndpointGroup, error) {
	p := ProtoToNetworkEndpointGroup(request.GetResource())
	res, err := c.ApplyNetworkEndpointGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkEndpointGroupToProto(res)
	return r, nil
}

// ApplyNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Apply() method.
func (s *NetworkEndpointGroupServer) ApplyComputeBetaNetworkEndpointGroup(ctx context.Context, request *betapb.ApplyComputeBetaNetworkEndpointGroupRequest) (*betapb.ComputeBetaNetworkEndpointGroup, error) {
	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNetworkEndpointGroup(ctx, cl, request)
}

// DeleteNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroup Delete() method.
func (s *NetworkEndpointGroupServer) DeleteComputeBetaNetworkEndpointGroup(ctx context.Context, request *betapb.DeleteComputeBetaNetworkEndpointGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkEndpointGroup(ctx, ProtoToNetworkEndpointGroup(request.GetResource()))

}

// ListComputeBetaNetworkEndpointGroup handles the gRPC request by passing it to the underlying NetworkEndpointGroupList() method.
func (s *NetworkEndpointGroupServer) ListComputeBetaNetworkEndpointGroup(ctx context.Context, request *betapb.ListComputeBetaNetworkEndpointGroupRequest) (*betapb.ListComputeBetaNetworkEndpointGroupResponse, error) {
	cl, err := createConfigNetworkEndpointGroup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkEndpointGroup(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaNetworkEndpointGroup
	for _, r := range resources.Items {
		rp := NetworkEndpointGroupToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaNetworkEndpointGroupResponse{Items: protos}, nil
}

func createConfigNetworkEndpointGroup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
