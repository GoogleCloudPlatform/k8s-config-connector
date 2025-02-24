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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/beta/networkservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
)

// TcpRouteServer implements the gRPC interface for TcpRoute.
type TcpRouteServer struct{}

// ProtoToTcpRouteRules converts a TcpRouteRules object from its proto representation.
func ProtoToNetworkservicesBetaTcpRouteRules(p *betapb.NetworkservicesBetaTcpRouteRules) *beta.TcpRouteRules {
	if p == nil {
		return nil
	}
	obj := &beta.TcpRouteRules{
		Action: ProtoToNetworkservicesBetaTcpRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesBetaTcpRouteRulesMatches(r))
	}
	return obj
}

// ProtoToTcpRouteRulesMatches converts a TcpRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesBetaTcpRouteRulesMatches(p *betapb.NetworkservicesBetaTcpRouteRulesMatches) *beta.TcpRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &beta.TcpRouteRulesMatches{
		Address: dcl.StringOrNil(p.GetAddress()),
		Port:    dcl.StringOrNil(p.GetPort()),
	}
	return obj
}

// ProtoToTcpRouteRulesAction converts a TcpRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesBetaTcpRouteRulesAction(p *betapb.NetworkservicesBetaTcpRouteRulesAction) *beta.TcpRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &beta.TcpRouteRulesAction{
		OriginalDestination: dcl.Bool(p.GetOriginalDestination()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesBetaTcpRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToTcpRouteRulesActionDestinations converts a TcpRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesBetaTcpRouteRulesActionDestinations(p *betapb.NetworkservicesBetaTcpRouteRulesActionDestinations) *beta.TcpRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.TcpRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToTcpRoute converts a TcpRoute resource from its proto representation.
func ProtoToTcpRoute(p *betapb.NetworkservicesBetaTcpRoute) *beta.TcpRoute {
	obj := &beta.TcpRoute{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesBetaTcpRouteRules(r))
	}
	for _, r := range p.GetMeshes() {
		obj.Meshes = append(obj.Meshes, r)
	}
	for _, r := range p.GetGateways() {
		obj.Gateways = append(obj.Gateways, r)
	}
	return obj
}

// TcpRouteRulesToProto converts a TcpRouteRules object to its proto representation.
func NetworkservicesBetaTcpRouteRulesToProto(o *beta.TcpRouteRules) *betapb.NetworkservicesBetaTcpRouteRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTcpRouteRules{}
	p.SetAction(NetworkservicesBetaTcpRouteRulesActionToProto(o.Action))
	sMatches := make([]*betapb.NetworkservicesBetaTcpRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesBetaTcpRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// TcpRouteRulesMatchesToProto converts a TcpRouteRulesMatches object to its proto representation.
func NetworkservicesBetaTcpRouteRulesMatchesToProto(o *beta.TcpRouteRulesMatches) *betapb.NetworkservicesBetaTcpRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTcpRouteRulesMatches{}
	p.SetAddress(dcl.ValueOrEmptyString(o.Address))
	p.SetPort(dcl.ValueOrEmptyString(o.Port))
	return p
}

// TcpRouteRulesActionToProto converts a TcpRouteRulesAction object to its proto representation.
func NetworkservicesBetaTcpRouteRulesActionToProto(o *beta.TcpRouteRulesAction) *betapb.NetworkservicesBetaTcpRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTcpRouteRulesAction{}
	p.SetOriginalDestination(dcl.ValueOrEmptyBool(o.OriginalDestination))
	sDestinations := make([]*betapb.NetworkservicesBetaTcpRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesBetaTcpRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// TcpRouteRulesActionDestinationsToProto converts a TcpRouteRulesActionDestinations object to its proto representation.
func NetworkservicesBetaTcpRouteRulesActionDestinationsToProto(o *beta.TcpRouteRulesActionDestinations) *betapb.NetworkservicesBetaTcpRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTcpRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// TcpRouteToProto converts a TcpRoute resource to its proto representation.
func TcpRouteToProto(resource *beta.TcpRoute) *betapb.NetworkservicesBetaTcpRoute {
	p := &betapb.NetworkservicesBetaTcpRoute{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	sRules := make([]*betapb.NetworkservicesBetaTcpRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesBetaTcpRouteRulesToProto(&r)
	}
	p.SetRules(sRules)
	sMeshes := make([]string, len(resource.Meshes))
	for i, r := range resource.Meshes {
		sMeshes[i] = r
	}
	p.SetMeshes(sMeshes)
	sGateways := make([]string, len(resource.Gateways))
	for i, r := range resource.Gateways {
		sGateways[i] = r
	}
	p.SetGateways(sGateways)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTcpRoute handles the gRPC request by passing it to the underlying TcpRoute Apply() method.
func (s *TcpRouteServer) applyTcpRoute(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaTcpRouteRequest) (*betapb.NetworkservicesBetaTcpRoute, error) {
	p := ProtoToTcpRoute(request.GetResource())
	res, err := c.ApplyTcpRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TcpRouteToProto(res)
	return r, nil
}

// applyNetworkservicesBetaTcpRoute handles the gRPC request by passing it to the underlying TcpRoute Apply() method.
func (s *TcpRouteServer) ApplyNetworkservicesBetaTcpRoute(ctx context.Context, request *betapb.ApplyNetworkservicesBetaTcpRouteRequest) (*betapb.NetworkservicesBetaTcpRoute, error) {
	cl, err := createConfigTcpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTcpRoute(ctx, cl, request)
}

// DeleteTcpRoute handles the gRPC request by passing it to the underlying TcpRoute Delete() method.
func (s *TcpRouteServer) DeleteNetworkservicesBetaTcpRoute(ctx context.Context, request *betapb.DeleteNetworkservicesBetaTcpRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTcpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTcpRoute(ctx, ProtoToTcpRoute(request.GetResource()))

}

// ListNetworkservicesBetaTcpRoute handles the gRPC request by passing it to the underlying TcpRouteList() method.
func (s *TcpRouteServer) ListNetworkservicesBetaTcpRoute(ctx context.Context, request *betapb.ListNetworkservicesBetaTcpRouteRequest) (*betapb.ListNetworkservicesBetaTcpRouteResponse, error) {
	cl, err := createConfigTcpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTcpRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaTcpRoute
	for _, r := range resources.Items {
		rp := TcpRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaTcpRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTcpRoute(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
