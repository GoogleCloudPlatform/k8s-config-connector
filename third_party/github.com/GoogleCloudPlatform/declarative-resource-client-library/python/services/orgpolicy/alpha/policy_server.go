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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/orgpolicy/alpha/orgpolicy_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/alpha"
)

// PolicyServer implements the gRPC interface for Policy.
type PolicyServer struct{}

// ProtoToPolicySpec converts a PolicySpec object from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpec(p *alphapb.OrgpolicyAlphaPolicySpec) *alpha.PolicySpec {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyAlphaPolicySpecRules(r))
	}
	return obj
}

// ProtoToPolicySpecRules converts a PolicySpecRules object from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRules(p *alphapb.OrgpolicyAlphaPolicySpecRules) *alpha.PolicySpecRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRules{
		Values:    ProtoToOrgpolicyAlphaPolicySpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyAlphaPolicySpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicySpecRulesValues converts a PolicySpecRulesValues object from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesValues(p *alphapb.OrgpolicyAlphaPolicySpecRulesValues) *alpha.PolicySpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicySpecRulesCondition converts a PolicySpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyAlphaPolicySpecRulesCondition(p *alphapb.OrgpolicyAlphaPolicySpecRulesCondition) *alpha.PolicySpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicySpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicyDryRunSpec converts a PolicyDryRunSpec object from its proto representation.
func ProtoToOrgpolicyAlphaPolicyDryRunSpec(p *alphapb.OrgpolicyAlphaPolicyDryRunSpec) *alpha.PolicyDryRunSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyDryRunSpec{
		Etag:              dcl.StringOrNil(p.GetEtag()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		InheritFromParent: dcl.Bool(p.GetInheritFromParent()),
		Reset:             dcl.Bool(p.GetReset()),
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToOrgpolicyAlphaPolicyDryRunSpecRules(r))
	}
	return obj
}

// ProtoToPolicyDryRunSpecRules converts a PolicyDryRunSpecRules object from its proto representation.
func ProtoToOrgpolicyAlphaPolicyDryRunSpecRules(p *alphapb.OrgpolicyAlphaPolicyDryRunSpecRules) *alpha.PolicyDryRunSpecRules {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyDryRunSpecRules{
		Values:    ProtoToOrgpolicyAlphaPolicyDryRunSpecRulesValues(p.GetValues()),
		AllowAll:  dcl.Bool(p.GetAllowAll()),
		DenyAll:   dcl.Bool(p.GetDenyAll()),
		Enforce:   dcl.Bool(p.GetEnforce()),
		Condition: ProtoToOrgpolicyAlphaPolicyDryRunSpecRulesCondition(p.GetCondition()),
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesValues converts a PolicyDryRunSpecRulesValues object from its proto representation.
func ProtoToOrgpolicyAlphaPolicyDryRunSpecRulesValues(p *alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesValues) *alpha.PolicyDryRunSpecRulesValues {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyDryRunSpecRulesValues{}
	for _, r := range p.GetAllowedValues() {
		obj.AllowedValues = append(obj.AllowedValues, r)
	}
	for _, r := range p.GetDeniedValues() {
		obj.DeniedValues = append(obj.DeniedValues, r)
	}
	return obj
}

// ProtoToPolicyDryRunSpecRulesCondition converts a PolicyDryRunSpecRulesCondition object from its proto representation.
func ProtoToOrgpolicyAlphaPolicyDryRunSpecRulesCondition(p *alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesCondition) *alpha.PolicyDryRunSpecRulesCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.PolicyDryRunSpecRulesCondition{
		Expression:  dcl.StringOrNil(p.GetExpression()),
		Title:       dcl.StringOrNil(p.GetTitle()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ProtoToPolicy converts a Policy resource from its proto representation.
func ProtoToPolicy(p *alphapb.OrgpolicyAlphaPolicy) *alpha.Policy {
	obj := &alpha.Policy{
		Name:       dcl.StringOrNil(p.GetName()),
		Spec:       ProtoToOrgpolicyAlphaPolicySpec(p.GetSpec()),
		DryRunSpec: ProtoToOrgpolicyAlphaPolicyDryRunSpec(p.GetDryRunSpec()),
		Etag:       dcl.StringOrNil(p.GetEtag()),
		Parent:     dcl.StringOrNil(p.GetParent()),
	}
	return obj
}

// PolicySpecToProto converts a PolicySpec object to its proto representation.
func OrgpolicyAlphaPolicySpecToProto(o *alpha.PolicySpec) *alphapb.OrgpolicyAlphaPolicySpec {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*alphapb.OrgpolicyAlphaPolicySpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyAlphaPolicySpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicySpecRulesToProto converts a PolicySpecRules object to its proto representation.
func OrgpolicyAlphaPolicySpecRulesToProto(o *alpha.PolicySpecRules) *alphapb.OrgpolicyAlphaPolicySpecRules {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRules{}
	p.SetValues(OrgpolicyAlphaPolicySpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyAlphaPolicySpecRulesConditionToProto(o.Condition))
	return p
}

// PolicySpecRulesValuesToProto converts a PolicySpecRulesValues object to its proto representation.
func OrgpolicyAlphaPolicySpecRulesValuesToProto(o *alpha.PolicySpecRulesValues) *alphapb.OrgpolicyAlphaPolicySpecRulesValues {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRulesValues{}
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
func OrgpolicyAlphaPolicySpecRulesConditionToProto(o *alpha.PolicySpecRulesCondition) *alphapb.OrgpolicyAlphaPolicySpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicySpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyDryRunSpecToProto converts a PolicyDryRunSpec object to its proto representation.
func OrgpolicyAlphaPolicyDryRunSpecToProto(o *alpha.PolicyDryRunSpec) *alphapb.OrgpolicyAlphaPolicyDryRunSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicyDryRunSpec{}
	p.SetEtag(dcl.ValueOrEmptyString(o.Etag))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	p.SetInheritFromParent(dcl.ValueOrEmptyBool(o.InheritFromParent))
	p.SetReset(dcl.ValueOrEmptyBool(o.Reset))
	sRules := make([]*alphapb.OrgpolicyAlphaPolicyDryRunSpecRules, len(o.Rules))
	for i, r := range o.Rules {
		sRules[i] = OrgpolicyAlphaPolicyDryRunSpecRulesToProto(&r)
	}
	p.SetRules(sRules)
	return p
}

// PolicyDryRunSpecRulesToProto converts a PolicyDryRunSpecRules object to its proto representation.
func OrgpolicyAlphaPolicyDryRunSpecRulesToProto(o *alpha.PolicyDryRunSpecRules) *alphapb.OrgpolicyAlphaPolicyDryRunSpecRules {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicyDryRunSpecRules{}
	p.SetValues(OrgpolicyAlphaPolicyDryRunSpecRulesValuesToProto(o.Values))
	p.SetAllowAll(dcl.ValueOrEmptyBool(o.AllowAll))
	p.SetDenyAll(dcl.ValueOrEmptyBool(o.DenyAll))
	p.SetEnforce(dcl.ValueOrEmptyBool(o.Enforce))
	p.SetCondition(OrgpolicyAlphaPolicyDryRunSpecRulesConditionToProto(o.Condition))
	return p
}

// PolicyDryRunSpecRulesValuesToProto converts a PolicyDryRunSpecRulesValues object to its proto representation.
func OrgpolicyAlphaPolicyDryRunSpecRulesValuesToProto(o *alpha.PolicyDryRunSpecRulesValues) *alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesValues {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesValues{}
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
func OrgpolicyAlphaPolicyDryRunSpecRulesConditionToProto(o *alpha.PolicyDryRunSpecRulesCondition) *alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.OrgpolicyAlphaPolicyDryRunSpecRulesCondition{}
	p.SetExpression(dcl.ValueOrEmptyString(o.Expression))
	p.SetTitle(dcl.ValueOrEmptyString(o.Title))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	return p
}

// PolicyToProto converts a Policy resource to its proto representation.
func PolicyToProto(resource *alpha.Policy) *alphapb.OrgpolicyAlphaPolicy {
	p := &alphapb.OrgpolicyAlphaPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSpec(OrgpolicyAlphaPolicySpecToProto(resource.Spec))
	p.SetDryRunSpec(OrgpolicyAlphaPolicyDryRunSpecToProto(resource.DryRunSpec))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))

	return p
}

// applyPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) applyPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOrgpolicyAlphaPolicyRequest) (*alphapb.OrgpolicyAlphaPolicy, error) {
	p := ProtoToPolicy(request.GetResource())
	res, err := c.ApplyPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PolicyToProto(res)
	return r, nil
}

// applyOrgpolicyAlphaPolicy handles the gRPC request by passing it to the underlying Policy Apply() method.
func (s *PolicyServer) ApplyOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.ApplyOrgpolicyAlphaPolicyRequest) (*alphapb.OrgpolicyAlphaPolicy, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPolicy(ctx, cl, request)
}

// DeletePolicy handles the gRPC request by passing it to the underlying Policy Delete() method.
func (s *PolicyServer) DeleteOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.DeleteOrgpolicyAlphaPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePolicy(ctx, ProtoToPolicy(request.GetResource()))

}

// ListOrgpolicyAlphaPolicy handles the gRPC request by passing it to the underlying PolicyList() method.
func (s *PolicyServer) ListOrgpolicyAlphaPolicy(ctx context.Context, request *alphapb.ListOrgpolicyAlphaPolicyRequest) (*alphapb.ListOrgpolicyAlphaPolicyResponse, error) {
	cl, err := createConfigPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPolicy(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OrgpolicyAlphaPolicy
	for _, r := range resources.Items {
		rp := PolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListOrgpolicyAlphaPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
