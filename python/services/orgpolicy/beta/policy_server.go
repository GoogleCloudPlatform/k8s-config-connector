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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/beta/orgpolicy_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpec converts a PolicySpec object from its proto representation.
func ProtoToOrgpolicyBetaPolicySpec(p *betapb.OrgpolicyBetaPolicySpec) *beta.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyBetaPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules object from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRules(p *betapb.OrgpolicyBetaPolicySpecRules) *beta.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRules{
		Values:    ProtoToOrgpolicyBetaPolicySpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyBetaPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues object from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesValues(p *betapb.OrgpolicyBetaPolicySpecRulesValues) *beta.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyBetaPolicySpecRulesCondition(p *betapb.OrgpolicyBetaPolicySpecRulesCondition) *beta.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &beta.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicyDryRunSpec converts a PolicyDryRunSpec object from its proto representation.
func ProtoToOrgpolicyBetaPolicyDryRunSpec(p *betapb.OrgpolicyBetaPolicyDryRunSpec) *beta.PolicyDryRunSpec {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDryRunSpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyBetaPolicyDryRunSpecRules(r))
	}
	return obj
}

// ProtoToPolicyDryRunSpecRules converts a PolicyDryRunSpecRules object from its proto representation.
func ProtoToOrgpolicyBetaPolicyDryRunSpecRules(p *betapb.OrgpolicyBetaPolicyDryRunSpecRules) *beta.PolicyDryRunSpecRules {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDryRunSpecRules{
		Values:    ProtoToOrgpolicyBetaPolicyDryRunSpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyBetaPolicyDryRunSpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesValues converts a PolicyDryRunSpecRulesValues object from its proto representation.
func ProtoToOrgpolicyBetaPolicyDryRunSpecRulesValues(p *betapb.OrgpolicyBetaPolicyDryRunSpecRulesValues) *beta.PolicyDryRunSpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDryRunSpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesCondition converts a PolicyDryRunSpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyBetaPolicyDryRunSpecRulesCondition(p *betapb.OrgpolicyBetaPolicyDryRunSpecRulesCondition) *beta.PolicyDryRunSpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &beta.PolicyDryRunSpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *betapb.OrgpolicyBetaPolicy) *beta.Policy {
	obj := &beta.Policy{
		Name:       dcl.StringOrNil(p.GetName()),
		Spec:       ProtoToOrgpolicyBetaPolicySpec(p.GetSpec()),
		DryRunSpec: ProtoToOrgpolicyBetaPolicyDryRunSpec(p.GetDryRunSpec()),
		Etag:       dcl.StringOrNil(p.GetEtag()),
		Parent:     dcl.StringOrNil(p.GetParent()),
	}
	return obj
}

// PolicySpecToProto converts a PolicySpec object to its proto representation.
func OrgpolicyBetaPolicySpecToProto(o *beta.PolicySpec) *betapb.OrgpolicyBetaPolicySpec {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*betapb.OrgpolicyBetaPolicySpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyBetaPolicySpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules object to its proto representation.
func OrgpolicyBetaPolicySpecRulesToProto(o *beta.PolicySpecRules) *betapb.OrgpolicyBetaPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRules{}
	p.SetValues(OrgpolicyBetaPolicySpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyBetaPolicySpecRulesConditionToProto(o.Condition))
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues object to its proto representation.
func OrgpolicyBetaPolicySpecRulesValuesToProto(o *beta.PolicySpecRulesValues) *betapb.OrgpolicyBetaPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRulesValues{}
	sAllowedValues := make([]string, len(o.AllowedValues))
	for i, r := range o.AllowedValues {
		sAllowedValues[i] = r
	}
	p.SetAllowedValues(sAllowedValues)
	sDeniedValues := make([]string, len(o.DeniedValues))
	for i, r := range o.DeniedValues {
		sDeniedValues[i] = r
	}
	p.SetDeniedValues(sDeniedValues)
	return p
}

// PolicySpecRulesConditionToProto converts a PolicySpecRulesCondition object to its proto representation.
func OrgpolicyBetaPolicySpecRulesConditionToProto(o *beta.PolicySpecRulesCondition) *betapb.OrgpolicyBetaPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicySpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyDryRunSpecToProto converts a PolicyDryRunSpec object to its proto representation.
func OrgpolicyBetaPolicyDryRunSpecToProto(o *beta.PolicyDryRunSpec) *betapb.OrgpolicyBetaPolicyDryRunSpec {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicyDryRunSpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*betapb.OrgpolicyBetaPolicyDryRunSpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyBetaPolicyDryRunSpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicyDryRunSpecRulesToProto converts a PolicyDryRunSpecRules object to its proto representation.
func OrgpolicyBetaPolicyDryRunSpecRulesToProto(o *beta.PolicyDryRunSpecRules) *betapb.OrgpolicyBetaPolicyDryRunSpecRules {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicyDryRunSpecRules{}
	p.SetValues(OrgpolicyBetaPolicyDryRunSpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyBetaPolicyDryRunSpecRulesConditionToProto(o.Condition))
	return p
}

// PolicyDryRunSpecRulesValuesToProto converts a PolicyDryRunSpecRulesValues object to its proto representation.
func OrgpolicyBetaPolicyDryRunSpecRulesValuesToProto(o *beta.PolicyDryRunSpecRulesValues) *betapb.OrgpolicyBetaPolicyDryRunSpecRulesValues {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicyDryRunSpecRulesValues{}
	sAllowedValues := make([]string, len(o.AllowedValues))
	for i, r := range o.AllowedValues {
		sAllowedValues[i] = r
	}
	p.SetAllowedValues(sAllowedValues)
	sDeniedValues := make([]string, len(o.DeniedValues))
	for i, r := range o.DeniedValues {
		sDeniedValues[i] = r
	}
	p.SetDeniedValues(sDeniedValues)
	return p
}

// PolicyDryRunSpecRulesConditionToProto converts a PolicyDryRunSpecRulesCondition object to its proto representation.
func OrgpolicyBetaPolicyDryRunSpecRulesConditionToProto(o *beta.PolicyDryRunSpecRulesCondition) *betapb.OrgpolicyBetaPolicyDryRunSpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &betapb.OrgpolicyBetaPolicyDryRunSpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *beta.Policy) *betapb.OrgpolicyBetaPolicy {
	p := &betapb.OrgpolicyBetaPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSpec(OrgpolicyBetaPolicySpecToProto(resource.Spec))
	p.SetDryRunSpec(OrgpolicyBetaPolicyDryRunSpecToProto(resource.DryRunSpec))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyOrgpolicyBetaPolicyRequest) (*betapb.OrgpolicyBetaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyOrgpolicyBetaPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyBetaPolicy(ctx context.Context, request *betapb.ApplyOrgpolicyBetaPolicyRequest) (*betapb.OrgpolicyBetaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyBetaPolicy(ctx context.Context, request *betapb.DeleteOrgpolicyBetaPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyBetaPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyBetaPolicy(ctx context.Context, request *betapb.ListOrgpolicyBetaPolicyRequest) (*betapb.ListOrgpolicyBetaPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OrgpolicyBetaPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListOrgpolicyBetaPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
