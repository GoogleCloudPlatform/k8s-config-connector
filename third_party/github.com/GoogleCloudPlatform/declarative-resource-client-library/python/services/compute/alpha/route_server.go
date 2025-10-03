// Copyright 2023 Google LLC. All Rights Reserved.
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

// RouteServer implements the gRPC interface for Route.
type RouteServer struct{}

// ProtoToRouteWarningCodeEnum converts a RouteWarningCodeEnum enum from its proto representation.
func ProtoToComputeAlphaRouteWarningCodeEnum(e alphapb.ComputeAlphaRouteWarningCodeEnum) *alpha.RouteWarningCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaRouteWarningCodeEnum_name[int32(e)]; ok {
		e := alpha.RouteWarningCodeEnum(n[len("ComputeAlphaRouteWarningCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouteWarning converts a RouteWarning object from its proto representation.
func ProtoToComputeAlphaRouteWarning(p *alphapb.ComputeAlphaRouteWarning) *alpha.RouteWarning {
	if p == nil {
		return nil
	}
	obj := &alpha.RouteWarning{
		Code:    ProtoToComputeAlphaRouteWarningCodeEnum(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	return obj
}

// ProtoToRoute converts a Route resource from its proto representation.
func ProtoToRoute(p *alphapb.ComputeAlphaRoute) *alpha.Route {
	obj := &alpha.Route{
		Id:               dcl.Int64OrNil(p.GetId()),
		Name:             dcl.StringOrNil(p.GetName()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		Network:          dcl.StringOrNil(p.GetNetwork()),
		DestRange:        dcl.StringOrNil(p.GetDestRange()),
		Priority:         dcl.Int64OrNil(p.GetPriority()),
		NextHopInstance:  dcl.StringOrNil(p.GetNextHopInstance()),
		NextHopIP:        dcl.StringOrNil(p.GetNextHopIp()),
		NextHopNetwork:   dcl.StringOrNil(p.GetNextHopNetwork()),
		NextHopGateway:   dcl.StringOrNil(p.GetNextHopGateway()),
		NextHopPeering:   dcl.StringOrNil(p.GetNextHopPeering()),
		NextHopIlb:       dcl.StringOrNil(p.GetNextHopIlb()),
		NextHopVpnTunnel: dcl.StringOrNil(p.GetNextHopVpnTunnel()),
		SelfLink:         dcl.StringOrNil(p.GetSelfLink()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetTag() {
		obj.Tag = append(obj.Tag, r)
	}
	for _, r := range p.GetWarning() {
		obj.Warning = append(obj.Warning, *ProtoToComputeAlphaRouteWarning(r))
	}
	return obj
}

// RouteWarningCodeEnumToProto converts a RouteWarningCodeEnum enum to its proto representation.
func ComputeAlphaRouteWarningCodeEnumToProto(e *alpha.RouteWarningCodeEnum) alphapb.ComputeAlphaRouteWarningCodeEnum {
	if e == nil {
		return alphapb.ComputeAlphaRouteWarningCodeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaRouteWarningCodeEnum_value["RouteWarningCodeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaRouteWarningCodeEnum(v)
	}
	return alphapb.ComputeAlphaRouteWarningCodeEnum(0)
}

// RouteWarningToProto converts a RouteWarning object to its proto representation.
func ComputeAlphaRouteWarningToProto(o *alpha.RouteWarning) *alphapb.ComputeAlphaRouteWarning {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaRouteWarning{}
	p.SetCode(ComputeAlphaRouteWarningCodeEnumToProto(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	mData := make(map[string]string, len(o.Data))
	for k, r := range o.Data {
		mData[k] = r
	}
	p.SetData(mData)
	return p
}

// RouteToProto converts a Route resource to its proto representation.
func RouteToProto(resource *alpha.Route) *alphapb.ComputeAlphaRoute {
	p := &alphapb.ComputeAlphaRoute{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetDestRange(dcl.ValueOrEmptyString(resource.DestRange))
	p.SetPriority(dcl.ValueOrEmptyInt64(resource.Priority))
	p.SetNextHopInstance(dcl.ValueOrEmptyString(resource.NextHopInstance))
	p.SetNextHopIp(dcl.ValueOrEmptyString(resource.NextHopIP))
	p.SetNextHopNetwork(dcl.ValueOrEmptyString(resource.NextHopNetwork))
	p.SetNextHopGateway(dcl.ValueOrEmptyString(resource.NextHopGateway))
	p.SetNextHopPeering(dcl.ValueOrEmptyString(resource.NextHopPeering))
	p.SetNextHopIlb(dcl.ValueOrEmptyString(resource.NextHopIlb))
	p.SetNextHopVpnTunnel(dcl.ValueOrEmptyString(resource.NextHopVpnTunnel))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sTag := make([]string, len(resource.Tag))
	for i, r := range resource.Tag {
		sTag[i] = r
	}
	p.SetTag(sTag)
	sWarning := make([]*alphapb.ComputeAlphaRouteWarning, len(resource.Warning))
	for i, r := range resource.Warning {
		sWarning[i] = ComputeAlphaRouteWarningToProto(&r)
	}
	p.SetWarning(sWarning)

	return p
}

// applyRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) applyRoute(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaRouteRequest) (*alphapb.ComputeAlphaRoute, error) {
	p := ProtoToRoute(request.GetResource())
	res, err := c.ApplyRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouteToProto(res)
	return r, nil
}

// applyComputeAlphaRoute handles the gRPC request by passing it to the underlying Route Apply() method.
func (s *RouteServer) ApplyComputeAlphaRoute(ctx context.Context, request *alphapb.ApplyComputeAlphaRouteRequest) (*alphapb.ComputeAlphaRoute, error) {
	cl, err := createConfigRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRoute(ctx, cl, request)
}

// DeleteRoute handles the gRPC request by passing it to the underlying Route Delete() method.
func (s *RouteServer) DeleteComputeAlphaRoute(ctx context.Context, request *alphapb.DeleteComputeAlphaRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRoute(ctx, ProtoToRoute(request.GetResource()))

}

// ListComputeAlphaRoute handles the gRPC request by passing it to the underlying RouteList() method.
func (s *RouteServer) ListComputeAlphaRoute(ctx context.Context, request *alphapb.ListComputeAlphaRouteRequest) (*alphapb.ListComputeAlphaRouteResponse, error) {
	cl, err := createConfigRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRoute(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaRoute
	for _, r := range resources.Items {
		rp := RouteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRoute(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
