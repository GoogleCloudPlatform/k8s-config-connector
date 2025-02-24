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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/alpha/networksecurity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
)

// AuthorizationPolicyServer implements the gRPC interface for AuthorizationPolicy.
type AuthorizationPolicyServer struct{}

// ProtoToAuthorizationPolicyActionEnum converts a AuthorizationPolicyActionEnum enum from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyActionEnum(e alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum) *alpha.AuthorizationPolicyActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum_name[int32(e)]; ok {
		e := alpha.AuthorizationPolicyActionEnum(n[len("NetworksecurityAlphaAuthorizationPolicyActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToAuthorizationPolicyRules converts a AuthorizationPolicyRules object from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRules(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRules) *alpha.AuthorizationPolicyRules {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRules{}
	for _, r := range p.GetSources() {
		obj.Sources = append(obj.Sources, *ProtoToNetworksecurityAlphaAuthorizationPolicyRulesSources(r))
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinations(r))
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesSources converts a AuthorizationPolicyRulesSources object from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesSources(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources) *alpha.AuthorizationPolicyRulesSources {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesSources{}
	for _, r := range p.GetPrincipals() {
		obj.Principals = append(obj.Principals, r)
	}
	for _, r := range p.GetIpBlocks() {
		obj.IPBlocks = append(obj.IPBlocks, r)
	}
	return obj
}

// ProtoToAuthorizationPolicyRulesDestinations converts a AuthorizationPolicyRulesDestinations object from its proto representation.
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinations(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations) *alpha.AuthorizationPolicyRulesDestinations {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesDestinations{
		HttpHeaderMatch: ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p.GetHttpHeaderMatch()),
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
func ProtoToNetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch(p *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch) *alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch{
		HeaderName: dcl.StringOrNil(p.GetHeaderName()),
		RegexMatch: dcl.StringOrNil(p.GetRegexMatch()),
	}
	return obj
}

// ProtoToAuthorizationPolicy converts a AuthorizationPolicy resource from its proto representation.
func ProtoToAuthorizationPolicy(p *alphapb.NetworksecurityAlphaAuthorizationPolicy) *alpha.AuthorizationPolicy {
	obj := &alpha.AuthorizationPolicy{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Action:      ProtoToNetworksecurityAlphaAuthorizationPolicyActionEnum(p.GetAction()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworksecurityAlphaAuthorizationPolicyRules(r))
	}
	return obj
}

// AuthorizationPolicyActionEnumToProto converts a AuthorizationPolicyActionEnum enum to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyActionEnumToProto(e *alpha.AuthorizationPolicyActionEnum) alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum {
	if e == nil {
		return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(0)
	}
	if v, ok := alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum_value["AuthorizationPolicyActionEnum"+string(*e)]; ok {
		return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(v)
	}
	return alphapb.NetworksecurityAlphaAuthorizationPolicyActionEnum(0)
}

// AuthorizationPolicyRulesToProto converts a AuthorizationPolicyRules object to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyRulesToProto(o *alpha.AuthorizationPolicyRules) *alphapb.NetworksecurityAlphaAuthorizationPolicyRules {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRules{}
	sSources := make([]*alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources, len(o.Sources))
	for i, r := range o.Sources {
		sSources[i] = NetworksecurityAlphaAuthorizationPolicyRulesSourcesToProto(&r)
	}
	p.SetSources(sSources)
	sDestinations := make([]*alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworksecurityAlphaAuthorizationPolicyRulesDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// AuthorizationPolicyRulesSourcesToProto converts a AuthorizationPolicyRulesSources object to its proto representation.
func NetworksecurityAlphaAuthorizationPolicyRulesSourcesToProto(o *alpha.AuthorizationPolicyRulesSources) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesSources{}
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
func NetworksecurityAlphaAuthorizationPolicyRulesDestinationsToProto(o *alpha.AuthorizationPolicyRulesDestinations) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinations{}
	p.SetHttpHeaderMatch(NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o.HttpHeaderMatch))
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
func NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatchToProto(o *alpha.AuthorizationPolicyRulesDestinationsHttpHeaderMatch) *alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicyRulesDestinationsHttpHeaderMatch{}
	p.SetHeaderName(dcl.ValueOrEmptyString(o.HeaderName))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	return p
}

// AuthorizationPolicyToProto converts a AuthorizationPolicy resource to its proto representation.
func AuthorizationPolicyToProto(resource *alpha.AuthorizationPolicy) *alphapb.NetworksecurityAlphaAuthorizationPolicy {
	p := &alphapb.NetworksecurityAlphaAuthorizationPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAction(NetworksecurityAlphaAuthorizationPolicyActionEnumToProto(resource.Action))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sRules := make([]*alphapb.NetworksecurityAlphaAuthorizationPolicyRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworksecurityAlphaAuthorizationPolicyRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) applyAuthorizationPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.NetworksecurityAlphaAuthorizationPolicy, error) {
	p := ProtoToAuthorizationPolicy(request.GetResource())
	res, err := c.ApplyAuthorizationPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AuthorizationPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityAlphaAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Apply() method.
func (s *AuthorizationPolicyServer) ApplyNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.NetworksecurityAlphaAuthorizationPolicy, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAuthorizationPolicy(ctx, cl, request)
}

// DeleteAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicy Delete() method.
func (s *AuthorizationPolicyServer) DeleteNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaAuthorizationPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAuthorizationPolicy(ctx, ProtoToAuthorizationPolicy(request.GetResource()))

}

// ListNetworksecurityAlphaAuthorizationPolicy handles the gRPC request by passing it to the underlying AuthorizationPolicyList() method.
func (s *AuthorizationPolicyServer) ListNetworksecurityAlphaAuthorizationPolicy(ctx context.Context, request *alphapb.ListNetworksecurityAlphaAuthorizationPolicyRequest) (*alphapb.ListNetworksecurityAlphaAuthorizationPolicyResponse, error) {
	cl, err := createConfigAuthorizationPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAuthorizationPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaAuthorizationPolicy
	for _, r := range resources.Items {
		rp := AuthorizationPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworksecurityAlphaAuthorizationPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAuthorizationPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
