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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/beta/networkconnectivity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/beta"
)

// HubServer implements the gRPC interface for Hub.
type HubServer struct{}

// ProtoToHubStateEnum converts a HubStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityBetaHubStateEnum(e betapb.NetworkconnectivityBetaHubStateEnum) *beta.HubStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkconnectivityBetaHubStateEnum_name[int32(e)]; ok {
		e := beta.HubStateEnum(n[len("NetworkconnectivityBetaHubStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToHubRoutingVpcs converts a HubRoutingVpcs object from its proto representation.
func ProtoToNetworkconnectivityBetaHubRoutingVpcs(p *betapb.NetworkconnectivityBetaHubRoutingVpcs) *beta.HubRoutingVpcs {
	if p == nil {
		return nil
	}
	obj := &beta.HubRoutingVpcs{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToHub converts a Hub resource from its proto representation.
func ProtoToHub(p *betapb.NetworkconnectivityBetaHub) *beta.Hub {
	obj := &beta.Hub{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		UniqueId:    dcl.StringOrNil(p.GetUniqueId()),
		State:       ProtoToNetworkconnectivityBetaHubStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetRoutingVpcs() {
		obj.RoutingVpcs = append(obj.RoutingVpcs, *ProtoToNetworkconnectivityBetaHubRoutingVpcs(r))
	}
	return obj
}

// HubStateEnumToProto converts a HubStateEnum enum to its proto representation.
func NetworkconnectivityBetaHubStateEnumToProto(e *beta.HubStateEnum) betapb.NetworkconnectivityBetaHubStateEnum {
	if e == nil {
		return betapb.NetworkconnectivityBetaHubStateEnum(0)
	}
	if v, ok := betapb.NetworkconnectivityBetaHubStateEnum_value["HubStateEnum"+string(*e)]; ok {
		return betapb.NetworkconnectivityBetaHubStateEnum(v)
	}
	return betapb.NetworkconnectivityBetaHubStateEnum(0)
}

// HubRoutingVpcsToProto converts a HubRoutingVpcs object to its proto representation.
func NetworkconnectivityBetaHubRoutingVpcsToProto(o *beta.HubRoutingVpcs) *betapb.NetworkconnectivityBetaHubRoutingVpcs {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaHubRoutingVpcs{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// HubToProto converts a Hub resource to its proto representation.
func HubToProto(resource *beta.Hub) *betapb.NetworkconnectivityBetaHub {
	p := &betapb.NetworkconnectivityBetaHub{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetState(NetworkconnectivityBetaHubStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sRoutingVpcs := make([]*betapb.NetworkconnectivityBetaHubRoutingVpcs, len(resource.RoutingVpcs))
	for i, r := range resource.RoutingVpcs {
		sRoutingVpcs[i] = NetworkconnectivityBetaHubRoutingVpcsToProto(&r)
	}
	p.SetRoutingVpcs(sRoutingVpcs)

	return p
}

// applyHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) applyHub(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkconnectivityBetaHubRequest) (*betapb.NetworkconnectivityBetaHub, error) {
	p := ProtoToHub(request.GetResource())
	res, err := c.ApplyHub(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HubToProto(res)
	return r, nil
}

// applyNetworkconnectivityBetaHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) ApplyNetworkconnectivityBetaHub(ctx context.Context, request *betapb.ApplyNetworkconnectivityBetaHubRequest) (*betapb.NetworkconnectivityBetaHub, error) {
	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHub(ctx, cl, request)
}

// DeleteHub handles the gRPC request by passing it to the underlying Hub Delete() method.
func (s *HubServer) DeleteNetworkconnectivityBetaHub(ctx context.Context, request *betapb.DeleteNetworkconnectivityBetaHubRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHub(ctx, ProtoToHub(request.GetResource()))

}

// ListNetworkconnectivityBetaHub handles the gRPC request by passing it to the underlying HubList() method.
func (s *HubServer) ListNetworkconnectivityBetaHub(ctx context.Context, request *betapb.ListNetworkconnectivityBetaHubRequest) (*betapb.ListNetworkconnectivityBetaHubResponse, error) {
	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHub(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkconnectivityBetaHub
	for _, r := range resources.Items {
		rp := HubToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkconnectivityBetaHubResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHub(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
