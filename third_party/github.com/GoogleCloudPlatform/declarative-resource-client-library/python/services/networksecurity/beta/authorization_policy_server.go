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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/beta/networksecurity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
)

// AuthorizationPolicyServer implements the gRPC interface for AuthorizationPolicy.
type AuthorizationPolicyServer struct{}

// ProtoToAuthorizationPolicyActionEnum converts a AuthorizationPolicyActionEnum enum from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyActionEnum(e betapb.NetworksecurityBetaAuthorizationPolicyActionEnum) *beta.AuthorizationPolicyActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworksecurityBetaAuthorizationPolicyActionEnum_name[int32(e)]; ok {
		e := beta.AuthorizationPolicyActionEnum(n[len("NetworksecurityBetaAuthorizationPolicyActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAuthorizationPolicyRules converts a AuthorizationPolicyRules object from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRules(p *betapb.NetworksecurityBetaAuthorizationPolicyRules) *beta.AuthorizationPolicyRules {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRules{}
	for _, r := range p.GetSources() {
		obj.Sources = append(obj.Sources, *ProtoToNetworksecurityBetaAuthorizationPolicyRulesSources(r))
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinations(r))
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesSources converts a AuthorizationPolicyRulesSources object from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesSources(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesSources) *beta.AuthorizationPolicyRulesSources {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesSources{}
	for _, r := range p.GetPrincipals() {
		obj.Principals = append(obj.Principals, r)
	}
	for _, r := range p.GetIpBlocks() {
		obj.IPBlocks = append(obj.IPBlocks, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinations converts a AuthorizationPolicyRulesDestinations object from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinations(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations) *beta.AuthorizationPolicyRulesDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p.GetHttpHeaderMatch()),
	}
	for _, r := range p.GetHosts() {
		obj.Hosts = append(obj.Hosts, r)
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinationsHttpHeaderMatch converts a AuthorizationPolicyRulesDestinationsHttpHeaderMatch object from its proto representation.
func ProtoToNetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch) *beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.StringOrNil(p.GetHeaderName()),
		RegexMatch: dcl.StringOrNil(p.GetRegexMatch()),
	}
	return obj
}

// ProtoToAuthorizationPolicy converts a AuthorizationPolicy resource from its proto representation.
func ProtoToAuthorizationPolicy(p *betapb.NetworksecurityBetaAuthorizationPolicy) *beta.AuthorizationPolicy {
	obj := &beta.AuthorizationPolicy{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Action:      ProtoToNetworksecurityBetaAuthorizationPolicyActionEnum(p.GetAction()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworksecurityBetaAuthorizationPolicyRules(r))
	}
	return obj
}

// AuthorizationPolicyActionEnumToProto converts a AuthorizationPolicyActionEnum enum to its proto representation.
func NetworksecurityBetaAuthorizationPolicyActionEnumToProto(e *beta.AuthorizationPolicyActionEnum) betapb.NetworksecurityBetaAuthorizationPolicyActionEnum {
	if e == nil {
		return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(0)
	}
	if v, ok := betapb.NetworksecurityBetaAuthorizationPolicyActionEnum_value["AuthorizationPolicyActionEnum"+string(*e)]; ok {
		return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(v)
	}
	return betapb.NetworksecurityBetaAuthorizationPolicyActionEnum(0)
}

// AuthorizationPolicyRulesToProto converts a AuthorizationPolicyRules object to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesToProto(o *beta.AuthorizationPolicyRules) *betapb.NetworksecurityBetaAuthorizationPolicyRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRules{}
	sSources := make([]*betapb.NetworksecurityBetaAuthorizationPolicyRulesSources, len(o.Sources))
	for i, r := range o.Sources {
		sSources[i] = NetworksecurityBetaAuthorizationPolicyRulesSourcesToProto(&r)
	}
	p.SetSources(sSources)
	sDestinations := make([]*betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworksecurityBetaAuthorizationPolicyRulesDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// AuthorizationPolicyRulesSourcesToProto converts a AuthorizationPolicyRulesSources object to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesSourcesToProto(o *beta.AuthorizationPolicyRulesSources) *betapb.NetworksecurityBetaAuthorizationPolicyRulesSources {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesSources{}
	sPrincipals := make([]string, len(o.Principals))
	for i, r := range o.Principals {
		sPrincipals[i] = r
	}
	p.SetPrincipals(sPrincipals)
	sIPBlocks := make([]string, len(o.IPBlocks))
	for i, r := range o.IPBlocks {
		sIPBlocks[i] = r
	}
	p.SetIpBlocks(sIPBlocks)
	return p
}

// AuthorizationPolicyRulesDestinationsToProto converts a AuthorizationPolicyRulesDestinations object to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesDestinationsToProto(o *beta.AuthorizationPolicyRulesDestinations) *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinations{}
	p.SetHttpHeaderMatch(NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o.HttpHeaderMatch))
	sHosts := make([]string, len(o.Hosts))
	for i, r := range o.Hosts {
		sHosts[i] = r
	}
	p.SetHosts(sHosts)
	sPorts := make([]int64, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	sMethods := make([]string, len(o.Methods))
	for i, r := range o.Methods {
		sMethods[i] = r
	}
	p.SetMethods(sMethods)
	return p
}

// AuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto converts a AuthorizationPolicyRulesDestinationsHttpHeaderMatch object to its proto representation.
func NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o *beta.AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaAuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	p.SetHeaderName(dcl.ValueOrEmptyString(o.HeaderName))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	return p
}

// AuthorizationPolicyToProto converts a AuthorizationPolicy resource to its proto representation.
func AuthorizationPolicyToProto(resource *beta.AuthorizationPolicy) *betapb.NetworksecurityBetaAuthorizationPolicy {
	p := &betapb.NetworksecurityBetaAuthorizationPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAction(NetworksecurityBetaAuthorizationPolicyActionEnumToProto(resource.Action))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sRules := make([]*betapb.NetworksecurityBetaAuthorizationPolicyRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworksecurityBetaAuthorizationPolicyRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) applyAuthorizationPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.NetworksecurityBetaAuthorizationPolicy, error) {
	p := ProtoToAuthorizationPolicy(request.GetResource())
	res, err := c.ApplyAuthorizationPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AuthorizationPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityBetaAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) ApplyNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.ApplyNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.NetworksecurityBetaAuthorizationPolicy, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAuthorizationPolicy(ctx, cl, request)
}

// DeleteAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Delete() method.
func (s *AuthorizationPolicyServer) DeleteNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.DeleteNetworksecurityBetaAuthorizationPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAuthorizationPolicy(ctx, ProtoToAuthorizationPolicy(request.GetResource()))

}

// ListNetworksecurityBetaAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicyList() method.
func (s *AuthorizationPolicyServer) ListNetworksecurityBetaAuthorizationPolicy(ctx context.Context, request *betapb.ListNetworksecurityBetaAuthorizationPolicyRequest) (*betapb.ListNetworksecurityBetaAuthorizationPolicyResponse, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAuthorizationPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaAuthorizationPolicy
	for _, r := range resources.Items {
		rp := AuthorizationPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworksecurityBetaAuthorizationPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAuthorizationPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
