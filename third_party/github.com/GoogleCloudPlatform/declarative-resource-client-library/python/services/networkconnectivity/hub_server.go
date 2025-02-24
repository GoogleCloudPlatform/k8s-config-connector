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
	networkconnectivitypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/networkconnectivity_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity"
)

// HubServer implements the gRPC interface for Hub.
type HubServer struct{}

// ProtoToHubStateEnum converts a HubStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityHubStateEnum(e networkconnectivitypb.NetworkconnectivityHubStateEnum) *networkconnectivity.HubStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkconnectivitypb.NetworkconnectivityHubStateEnum_name[int32(e)]; ok {
		e := networkconnectivity.HubStateEnum(n[len("NetworkconnectivityHubStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToHubRoutingVpcs converts a HubRoutingVpcs object from its proto representation.
func ProtoToNetworkconnectivityHubRoutingVpcs(p *networkconnectivitypb.NetworkconnectivityHubRoutingVpcs) *networkconnectivity.HubRoutingVpcs {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.HubRoutingVpcs{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToHub converts a Hub resource from its proto representation.
func ProtoToHub(p *networkconnectivitypb.NetworkconnectivityHub) *networkconnectivity.Hub {
	obj := &networkconnectivity.Hub{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		UniqueId:    dcl.StringOrNil(p.GetUniqueId()),
		State:       ProtoToNetworkconnectivityHubStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetRoutingVpcs() {
		obj.RoutingVpcs = append(obj.RoutingVpcs, *ProtoToNetworkconnectivityHubRoutingVpcs(r))
	}
	return obj
}

// HubStateEnumToProto converts a HubStateEnum enum to its proto representation.
func NetworkconnectivityHubStateEnumToProto(e *networkconnectivity.HubStateEnum) networkconnectivitypb.NetworkconnectivityHubStateEnum {
	if e == nil {
		return networkconnectivitypb.NetworkconnectivityHubStateEnum(0)
	}
	if v, ok := networkconnectivitypb.NetworkconnectivityHubStateEnum_value["HubStateEnum"+string(*e)]; ok {
		return networkconnectivitypb.NetworkconnectivityHubStateEnum(v)
	}
	return networkconnectivitypb.NetworkconnectivityHubStateEnum(0)
}

// HubRoutingVpcsToProto converts a HubRoutingVpcs object to its proto representation.
func NetworkconnectivityHubRoutingVpcsToProto(o *networkconnectivity.HubRoutingVpcs) *networkconnectivitypb.NetworkconnectivityHubRoutingVpcs {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivityHubRoutingVpcs{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// HubToProto converts a Hub resource to its proto representation.
func HubToProto(resource *networkconnectivity.Hub) *networkconnectivitypb.NetworkconnectivityHub {
	p := &networkconnectivitypb.NetworkconnectivityHub{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetState(NetworkconnectivityHubStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sRoutingVpcs := make([]*networkconnectivitypb.NetworkconnectivityHubRoutingVpcs, len(resource.RoutingVpcs))
	for i, r := range resource.RoutingVpcs {
		sRoutingVpcs[i] = NetworkconnectivityHubRoutingVpcsToProto(&r)
	}
	p.SetRoutingVpcs(sRoutingVpcs)

	return p
}

// applyHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) applyHub(ctx context.Context, c *networkconnectivity.Client, request *networkconnectivitypb.ApplyNetworkconnectivityHubRequest) (*networkconnectivitypb.NetworkconnectivityHub, error) {
	p := ProtoToHub(request.GetResource())
	res, err := c.ApplyHub(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HubToProto(res)
	return r, nil
}

// applyNetworkconnectivityHub handles the gRPC request by passing it to the underlying Hub Apply() method.
func (s *HubServer) ApplyNetworkconnectivityHub(ctx context.Context, request *networkconnectivitypb.ApplyNetworkconnectivityHubRequest) (*networkconnectivitypb.NetworkconnectivityHub, error) {
	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHub(ctx, cl, request)
}

// DeleteHub handles the gRPC request by passing it to the underlying Hub Delete() method.
func (s *HubServer) DeleteNetworkconnectivityHub(ctx context.Context, request *networkconnectivitypb.DeleteNetworkconnectivityHubRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHub(ctx, ProtoToHub(request.GetResource()))

}

// ListNetworkconnectivityHub handles the gRPC request by passing it to the underlying HubList() method.
func (s *HubServer) ListNetworkconnectivityHub(ctx context.Context, request *networkconnectivitypb.ListNetworkconnectivityHubRequest) (*networkconnectivitypb.ListNetworkconnectivityHubResponse, error) {
	cl, err := createConfigHub(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHub(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*networkconnectivitypb.NetworkconnectivityHub
	for _, r := range resources.Items {
		rp := HubToProto(r)
		protos = append(protos, rp)
	}
	p := &networkconnectivitypb.ListNetworkconnectivityHubResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHub(ctx context.Context, service_account_file string) (*networkconnectivity.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkconnectivity.NewClient(conf), nil
}
