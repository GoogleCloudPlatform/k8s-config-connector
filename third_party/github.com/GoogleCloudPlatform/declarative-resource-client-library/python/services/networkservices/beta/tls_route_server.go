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

// TlsRouteServer implements the gRPC interface for TlsRoute.
type TlsRouteServer struct{}

// ProtoToTlsRouteRules converts a TlsRouteRules object from its proto representation.
func ProtoToNetworkservicesBetaTlsRouteRules(p *betapb.NetworkservicesBetaTlsRouteRules) *beta.TlsRouteRules {
	if p == nil {
		return nil
	}
	obj := &beta.TlsRouteRules{
		Action: ProtoToNetworkservicesBetaTlsRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesBetaTlsRouteRulesMatches(r))
	}
	return obj
}

// ProtoToTlsRouteRulesMatches converts a TlsRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesBetaTlsRouteRulesMatches(p *betapb.NetworkservicesBetaTlsRouteRulesMatches) *beta.TlsRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &beta.TlsRouteRulesMatches{}
	for _, r := range p.GetSniHost() {
		obj.SniHost = append(obj.SniHost, r)
	}
	for _, r := range p.GetAlpn() {
		obj.Alpn = append(obj.Alpn, r)
	}
	return obj
}

// ProtoToTlsRouteRulesAction converts a TlsRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesBetaTlsRouteRulesAction(p *betapb.NetworkservicesBetaTlsRouteRulesAction) *beta.TlsRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &beta.TlsRouteRulesAction{}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesBetaTlsRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToTlsRouteRulesActionDestinations converts a TlsRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesBetaTlsRouteRulesActionDestinations(p *betapb.NetworkservicesBetaTlsRouteRulesActionDestinations) *beta.TlsRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.TlsRouteRulesActionDestinations{
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
		Weight:      dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToTlsRoute converts a TlsRoute resource from its proto representation.
func ProtoToTlsRoute(p *betapb.NetworkservicesBetaTlsRoute) *beta.TlsRoute {
	obj := &beta.TlsRoute{
		Name:        dcl.StringOrNil(p.GetName()),
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesBetaTlsRouteRules(r))
	}
	for _, r := range p.GetMeshes() {
		obj.Meshes = append(obj.Meshes, r)
	}
	for _, r := range p.GetGateways() {
		obj.Gateways = append(obj.Gateways, r)
	}
	return obj
}

// TlsRouteRulesToProto converts a TlsRouteRules object to its proto representation.
func NetworkservicesBetaTlsRouteRulesToProto(o *beta.TlsRouteRules) *betapb.NetworkservicesBetaTlsRouteRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTlsRouteRules{}
	p.SetAction(NetworkservicesBetaTlsRouteRulesActionToProto(o.Action))
	sMatches := make([]*betapb.NetworkservicesBetaTlsRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesBetaTlsRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// TlsRouteRulesMatchesToProto converts a TlsRouteRulesMatches object to its proto representation.
func NetworkservicesBetaTlsRouteRulesMatchesToProto(o *beta.TlsRouteRulesMatches) *betapb.NetworkservicesBetaTlsRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTlsRouteRulesMatches{}
	sSniHost := make([]string, len(o.SniHost))
	for i, r := range o.SniHost {
		sSniHost[i] = r
	}
	p.SetSniHost(sSniHost)
	sAlpn := make([]string, len(o.Alpn))
	for i, r := range o.Alpn {
		sAlpn[i] = r
	}
	p.SetAlpn(sAlpn)
	return p
}

// TlsRouteRulesActionToProto converts a TlsRouteRulesAction object to its proto representation.
func NetworkservicesBetaTlsRouteRulesActionToProto(o *beta.TlsRouteRulesAction) *betapb.NetworkservicesBetaTlsRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTlsRouteRulesAction{}
	sDestinations := make([]*betapb.NetworkservicesBetaTlsRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesBetaTlsRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// TlsRouteRulesActionDestinationsToProto converts a TlsRouteRulesActionDestinations object to its proto representation.
func NetworkservicesBetaTlsRouteRulesActionDestinationsToProto(o *beta.TlsRouteRulesActionDestinations) *betapb.NetworkservicesBetaTlsRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaTlsRouteRulesActionDestinations{}
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// TlsRouteToProto converts a TlsRoute resource to its proto representation.
func TlsRouteToProto(resource *beta.TlsRoute) *betapb.NetworkservicesBetaTlsRoute {
	p := &betapb.NetworkservicesBetaTlsRoute{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sRules := make([]*betapb.NetworkservicesBetaTlsRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesBetaTlsRouteRulesToProto(&r)
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

	return p
}

// applyTlsRoute handles the gRPC request by passing it to the underlying TlsRoute Apply() method.
func (s *TlsRouteServer) applyTlsRoute(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaTlsRouteRequest) (*betapb.NetworkservicesBetaTlsRoute, error) {
	p := ProtoToTlsRoute(request.GetResource())
	res, err := c.ApplyTlsRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TlsRouteToProto(res)
	return r, nil
}

// applyNetworkservicesBetaTlsRoute handles the gRPC request by passing it to the underlying TlsRoute Apply() method.
func (s *TlsRouteServer) ApplyNetworkservicesBetaTlsRoute(ctx context.Context, request *betapb.ApplyNetworkservicesBetaTlsRouteRequest) (*betapb.NetworkservicesBetaTlsRoute, error) {
	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTlsRoute(ctx, cl, request)
}

// DeleteTlsRoute handles the gRPC request by passing it to the underlying TlsRoute Delete() method.
func (s *TlsRouteServer) DeleteNetworkservicesBetaTlsRoute(ctx context.Context, request *betapb.DeleteNetworkservicesBetaTlsRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTlsRoute(ctx, ProtoToTlsRoute(request.GetResource()))

}

// ListNetworkservicesBetaTlsRoute handles the gRPC request by passing it to the underlying TlsRouteList() method.
func (s *TlsRouteServer) ListNetworkservicesBetaTlsRoute(ctx context.Context, request *betapb.ListNetworkservicesBetaTlsRouteRequest) (*betapb.ListNetworkservicesBetaTlsRouteResponse, error) {
	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTlsRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaTlsRoute
	for _, r := range resources.Items {
		rp := TlsRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaTlsRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTlsRoute(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
