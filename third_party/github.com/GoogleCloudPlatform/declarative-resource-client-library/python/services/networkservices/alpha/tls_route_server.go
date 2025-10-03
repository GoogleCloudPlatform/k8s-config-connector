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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// TlsRouteServer implements the gRPC interface for TlsRoute.
type TlsRouteServer struct{}

// ProtoToTlsRouteRules converts a TlsRouteRules object from its proto representation.
func ProtoToNetworkservicesAlphaTlsRouteRules(p *alphapb.NetworkservicesAlphaTlsRouteRules) *alpha.TlsRouteRules {
	if p == nil {
		return nil
	}
	obj := &alpha.TlsRouteRules{
		Action: ProtoToNetworkservicesAlphaTlsRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesAlphaTlsRouteRulesMatches(r))
	}
	return obj
}

// ProtoToTlsRouteRulesMatches converts a TlsRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesAlphaTlsRouteRulesMatches(p *alphapb.NetworkservicesAlphaTlsRouteRulesMatches) *alpha.TlsRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &alpha.TlsRouteRulesMatches{}
	for _, r := range p.GetSniHost() {
		obj.SniHost = append(obj.SniHost, r)
	}
	for _, r := range p.GetAlpn() {
		obj.Alpn = append(obj.Alpn, r)
	}
	return obj
}

// ProtoToTlsRouteRulesAction converts a TlsRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesAlphaTlsRouteRulesAction(p *alphapb.NetworkservicesAlphaTlsRouteRulesAction) *alpha.TlsRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &alpha.TlsRouteRulesAction{}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesAlphaTlsRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToTlsRouteRulesActionDestinations converts a TlsRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesAlphaTlsRouteRulesActionDestinations(p *alphapb.NetworkservicesAlphaTlsRouteRulesActionDestinations) *alpha.TlsRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &alpha.TlsRouteRulesActionDestinations{
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
		Weight:      dcl.Int64OrNil(p.GetWeight()),
	}
	return obj
}

// ProtoToTlsRoute converts a TlsRoute resource from its proto representation.
func ProtoToTlsRoute(p *alphapb.NetworkservicesAlphaTlsRoute) *alpha.TlsRoute {
	obj := &alpha.TlsRoute{
		Name:        dcl.StringOrNil(p.GetName()),
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesAlphaTlsRouteRules(r))
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
func NetworkservicesAlphaTlsRouteRulesToProto(o *alpha.TlsRouteRules) *alphapb.NetworkservicesAlphaTlsRouteRules {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaTlsRouteRules{}
	p.SetAction(NetworkservicesAlphaTlsRouteRulesActionToProto(o.Action))
	sMatches := make([]*alphapb.NetworkservicesAlphaTlsRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesAlphaTlsRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// TlsRouteRulesMatchesToProto converts a TlsRouteRulesMatches object to its proto representation.
func NetworkservicesAlphaTlsRouteRulesMatchesToProto(o *alpha.TlsRouteRulesMatches) *alphapb.NetworkservicesAlphaTlsRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaTlsRouteRulesMatches{}
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
func NetworkservicesAlphaTlsRouteRulesActionToProto(o *alpha.TlsRouteRulesAction) *alphapb.NetworkservicesAlphaTlsRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaTlsRouteRulesAction{}
	sDestinations := make([]*alphapb.NetworkservicesAlphaTlsRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesAlphaTlsRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// TlsRouteRulesActionDestinationsToProto converts a TlsRouteRulesActionDestinations object to its proto representation.
func NetworkservicesAlphaTlsRouteRulesActionDestinationsToProto(o *alpha.TlsRouteRulesActionDestinations) *alphapb.NetworkservicesAlphaTlsRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaTlsRouteRulesActionDestinations{}
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	return p
}

// TlsRouteToProto converts a TlsRoute resource to its proto representation.
func TlsRouteToProto(resource *alpha.TlsRoute) *alphapb.NetworkservicesAlphaTlsRoute {
	p := &alphapb.NetworkservicesAlphaTlsRoute{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sRules := make([]*alphapb.NetworkservicesAlphaTlsRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesAlphaTlsRouteRulesToProto(&r)
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
func (s *TlsRouteServer) applyTlsRoute(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaTlsRouteRequest) (*alphapb.NetworkservicesAlphaTlsRoute, error) {
	p := ProtoToTlsRoute(request.GetResource())
	res, err := c.ApplyTlsRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TlsRouteToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaTlsRoute handles the gRPC request by passing it to the underlying TlsRoute Apply() method.
func (s *TlsRouteServer) ApplyNetworkservicesAlphaTlsRoute(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaTlsRouteRequest) (*alphapb.NetworkservicesAlphaTlsRoute, error) {
	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTlsRoute(ctx, cl, request)
}

// DeleteTlsRoute handles the gRPC request by passing it to the underlying TlsRoute Delete() method.
func (s *TlsRouteServer) DeleteNetworkservicesAlphaTlsRoute(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaTlsRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTlsRoute(ctx, ProtoToTlsRoute(request.GetResource()))

}

// ListNetworkservicesAlphaTlsRoute handles the gRPC request by passing it to the underlying TlsRouteList() method.
func (s *TlsRouteServer) ListNetworkservicesAlphaTlsRoute(ctx context.Context, request *alphapb.ListNetworkservicesAlphaTlsRouteRequest) (*alphapb.ListNetworkservicesAlphaTlsRouteResponse, error) {
	cl, err := createConfigTlsRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTlsRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaTlsRoute
	for _, r := range resources.Items {
		rp := TlsRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaTlsRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTlsRoute(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
