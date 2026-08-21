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
	orgpolicypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/orgpolicy_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpec converts a PolicySpec object from its proto representation.
func ProtoToOrgpolicyPolicySpec(p *orgpolicypb.OrgpolicyPolicySpec) *orgpolicy.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules object from its proto representation.
func ProtoToOrgpolicyPolicySpecRules(p *orgpolicypb.OrgpolicyPolicySpecRules) *orgpolicy.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRules{
		Values:    ProtoToOrgpolicyPolicySpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues object from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesValues(p *orgpolicypb.OrgpolicyPolicySpecRulesValues) *orgpolicy.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyPolicySpecRulesCondition(p *orgpolicypb.OrgpolicyPolicySpecRulesCondition) *orgpolicy.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicyDryRunSpec converts a PolicyDryRunSpec object from its proto representation.
func ProtoToOrgpolicyPolicyDryRunSpec(p *orgpolicypb.OrgpolicyPolicyDryRunSpec) *orgpolicy.PolicyDryRunSpec {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicyDryRunSpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyPolicyDryRunSpecRules(r))
	}
	return obj
}

// ProtoToPolicyDryRunSpecRules converts a PolicyDryRunSpecRules object from its proto representation.
func ProtoToOrgpolicyPolicyDryRunSpecRules(p *orgpolicypb.OrgpolicyPolicyDryRunSpecRules) *orgpolicy.PolicyDryRunSpecRules {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicyDryRunSpecRules{
		Values:    ProtoToOrgpolicyPolicyDryRunSpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyPolicyDryRunSpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesValues converts a PolicyDryRunSpecRulesValues object from its proto representation.
func ProtoToOrgpolicyPolicyDryRunSpecRulesValues(p *orgpolicypb.OrgpolicyPolicyDryRunSpecRulesValues) *orgpolicy.PolicyDryRunSpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicyDryRunSpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesCondition converts a PolicyDryRunSpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyPolicyDryRunSpecRulesCondition(p *orgpolicypb.OrgpolicyPolicyDryRunSpecRulesCondition) *orgpolicy.PolicyDryRunSpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &orgpolicy.PolicyDryRunSpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *orgpolicypb.OrgpolicyPolicy) *orgpolicy.Policy {
	obj := &orgpolicy.Policy{
		Name:       dcl.StringOrNil(p.GetName()),
		Spec:       ProtoToOrgpolicyPolicySpec(p.GetSpec()),
		DryRunSpec: ProtoToOrgpolicyPolicyDryRunSpec(p.GetDryRunSpec()),
		Etag:       dcl.StringOrNil(p.GetEtag()),
		Parent:     dcl.StringOrNil(p.GetParent()),
	}
	return obj
}

// PolicySpecToProto converts a PolicySpec object to its proto representation.
func OrgpolicyPolicySpecToProto(o *orgpolicy.PolicySpec) *orgpolicypb.OrgpolicyPolicySpec {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*orgpolicypb.OrgpolicyPolicySpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyPolicySpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules object to its proto representation.
func OrgpolicyPolicySpecRulesToProto(o *orgpolicy.PolicySpecRules) *orgpolicypb.OrgpolicyPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRules{}
	p.SetValues(OrgpolicyPolicySpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyPolicySpecRulesConditionToProto(o.Condition))
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues object to its proto representation.
func OrgpolicyPolicySpecRulesValuesToProto(o *orgpolicy.PolicySpecRulesValues) *orgpolicypb.OrgpolicyPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRulesValues{}
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
func OrgpolicyPolicySpecRulesConditionToProto(o *orgpolicy.PolicySpecRulesCondition) *orgpolicypb.OrgpolicyPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicySpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyDryRunSpecToProto converts a PolicyDryRunSpec object to its proto representation.
func OrgpolicyPolicyDryRunSpecToProto(o *orgpolicy.PolicyDryRunSpec) *orgpolicypb.OrgpolicyPolicyDryRunSpec {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicyDryRunSpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*orgpolicypb.OrgpolicyPolicyDryRunSpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyPolicyDryRunSpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicyDryRunSpecRulesToProto converts a PolicyDryRunSpecRules object to its proto representation.
func OrgpolicyPolicyDryRunSpecRulesToProto(o *orgpolicy.PolicyDryRunSpecRules) *orgpolicypb.OrgpolicyPolicyDryRunSpecRules {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicyDryRunSpecRules{}
	p.SetValues(OrgpolicyPolicyDryRunSpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyPolicyDryRunSpecRulesConditionToProto(o.Condition))
	return p
}

// PolicyDryRunSpecRulesValuesToProto converts a PolicyDryRunSpecRulesValues object to its proto representation.
func OrgpolicyPolicyDryRunSpecRulesValuesToProto(o *orgpolicy.PolicyDryRunSpecRulesValues) *orgpolicypb.OrgpolicyPolicyDryRunSpecRulesValues {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicyDryRunSpecRulesValues{}
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
func OrgpolicyPolicyDryRunSpecRulesConditionToProto(o *orgpolicy.PolicyDryRunSpecRulesCondition) *orgpolicypb.OrgpolicyPolicyDryRunSpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &orgpolicypb.OrgpolicyPolicyDryRunSpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *orgpolicy.Policy) *orgpolicypb.OrgpolicyPolicy {
	p := &orgpolicypb.OrgpolicyPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSpec(OrgpolicyPolicySpecToProto(resource.Spec))
	p.SetDryRunSpec(OrgpolicyPolicyDryRunSpecToProto(resource.DryRunSpec))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *orgpolicy.Client, request *orgpolicypb.ApplyOrgpolicyPolicyRequest) (*orgpolicypb.OrgpolicyPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyOrgpolicyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.ApplyOrgpolicyPolicyRequest) (*orgpolicypb.OrgpolicyPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.DeleteOrgpolicyPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyPolicy(ctx context.Context, request *orgpolicypb.ListOrgpolicyPolicyRequest) (*orgpolicypb.ListOrgpolicyPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*orgpolicypb.OrgpolicyPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &orgpolicypb.ListOrgpolicyPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*orgpolicy.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return orgpolicy.NewClient(conf), nil
}
