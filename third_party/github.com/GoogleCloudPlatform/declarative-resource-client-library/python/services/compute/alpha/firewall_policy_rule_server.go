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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// FirewallPolicyRuleServer implements the gRPC interface for FirewallPolicyRule.
type FirewallPolicyRuleServer struct{}

// ProtoToFirewallPolicyRuleDirectionEnum converts a FirewallPolicyRuleDirectionEnum enum from its proto representation.
func ProtoToComputeAlphaFirewallPolicyRuleDirectionEnum(e alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum) *alpha.FirewallPolicyRuleDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum_name[int32(e)]; ok {
		e := alpha.FirewallPolicyRuleDirectionEnum(n[len("ComputeAlphaFirewallPolicyRuleDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirewallPolicyRuleMatch converts a FirewallPolicyRuleMatch object from its proto representation.
func ProtoToComputeAlphaFirewallPolicyRuleMatch(p *alphapb.ComputeAlphaFirewallPolicyRuleMatch) *alpha.FirewallPolicyRuleMatch {
	if p == nil {
		return nil
	}
	obj := &alpha.FirewallPolicyRuleMatch{}
	for _, r := range p.GetSrcIpRanges() {
		obj.SrcIPRanges = append(obj.SrcIPRanges, r)
	}
	for _, r := range p.GetDestIpRanges() {
		obj.DestIPRanges = append(obj.DestIPRanges, r)
	}
	for _, r := range p.GetSrcRegionCodes() {
		obj.SrcRegionCodes = append(obj.SrcRegionCodes, r)
	}
	for _, r := range p.GetDestRegionCodes() {
		obj.DestRegionCodes = append(obj.DestRegionCodes, r)
	}
	for _, r := range p.GetSrcThreatIntelligences() {
		obj.SrcThreatIntelligences = append(obj.SrcThreatIntelligences, r)
	}
	for _, r := range p.GetDestThreatIntelligences() {
		obj.DestThreatIntelligences = append(obj.DestThreatIntelligences, r)
	}
	for _, r := range p.GetSrcFqdns() {
		obj.SrcFqdns = append(obj.SrcFqdns, r)
	}
	for _, r := range p.GetDestFqdns() {
		obj.DestFqdns = append(obj.DestFqdns, r)
	}
	for _, r := range p.GetLayer4Configs() {
		obj.Layer4Configs = append(obj.Layer4Configs, *ProtoToComputeAlphaFirewallPolicyRuleMatchLayer4Configs(r))
	}
	for _, r := range p.GetSrcAddressGroups() {
		obj.SrcAddressGroups = append(obj.SrcAddressGroups, r)
	}
	for _, r := range p.GetDestAddressGroups() {
		obj.DestAddressGroups = append(obj.DestAddressGroups, r)
	}
	return obj
}

// ProtoToFirewallPolicyRuleMatchLayer4Configs converts a FirewallPolicyRuleMatchLayer4Configs object from its proto representation.
func ProtoToComputeAlphaFirewallPolicyRuleMatchLayer4Configs(p *alphapb.ComputeAlphaFirewallPolicyRuleMatchLayer4Configs) *alpha.FirewallPolicyRuleMatchLayer4Configs {
	if p == nil {
		return nil
	}
	obj := &alpha.FirewallPolicyRuleMatchLayer4Configs{
		IPProtocol: dcl.StringOrNil(p.GetIpProtocol()),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToFirewallPolicyRule converts a FirewallPolicyRule resource from its proto representation.
func ProtoToFirewallPolicyRule(p *alphapb.ComputeAlphaFirewallPolicyRule) *alpha.FirewallPolicyRule {
	obj := &alpha.FirewallPolicyRule{
		Description:    dcl.StringOrNil(p.GetDescription()),
		Priority:       dcl.Int64OrNil(p.GetPriority()),
		Match:          ProtoToComputeAlphaFirewallPolicyRuleMatch(p.GetMatch()),
		Action:         dcl.StringOrNil(p.GetAction()),
		Direction:      ProtoToComputeAlphaFirewallPolicyRuleDirectionEnum(p.GetDirection()),
		EnableLogging:  dcl.Bool(p.GetEnableLogging()),
		RuleTupleCount: dcl.Int64OrNil(p.GetRuleTupleCount()),
		Disabled:       dcl.Bool(p.GetDisabled()),
		Kind:           dcl.StringOrNil(p.GetKind()),
		FirewallPolicy: dcl.StringOrNil(p.GetFirewallPolicy()),
	}
	for _, r := range p.GetTargetResources() {
		obj.TargetResources = append(obj.TargetResources, r)
	}
	for _, r := range p.GetTargetServiceAccounts() {
		obj.TargetServiceAccounts = append(obj.TargetServiceAccounts, r)
	}
	return obj
}

// FirewallPolicyRuleDirectionEnumToProto converts a FirewallPolicyRuleDirectionEnum enum to its proto representation.
func ComputeAlphaFirewallPolicyRuleDirectionEnumToProto(e *alpha.FirewallPolicyRuleDirectionEnum) alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum {
	if e == nil {
		return alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum_value["FirewallPolicyRuleDirectionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum(v)
	}
	return alphapb.ComputeAlphaFirewallPolicyRuleDirectionEnum(0)
}

// FirewallPolicyRuleMatchToProto converts a FirewallPolicyRuleMatch object to its proto representation.
func ComputeAlphaFirewallPolicyRuleMatchToProto(o *alpha.FirewallPolicyRuleMatch) *alphapb.ComputeAlphaFirewallPolicyRuleMatch {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaFirewallPolicyRuleMatch{}
	sSrcIPRanges := make([]string, len(o.SrcIPRanges))
	for i, r := range o.SrcIPRanges {
		sSrcIPRanges[i] = r
	}
	p.SetSrcIpRanges(sSrcIPRanges)
	sDestIPRanges := make([]string, len(o.DestIPRanges))
	for i, r := range o.DestIPRanges {
		sDestIPRanges[i] = r
	}
	p.SetDestIpRanges(sDestIPRanges)
	sSrcRegionCodes := make([]string, len(o.SrcRegionCodes))
	for i, r := range o.SrcRegionCodes {
		sSrcRegionCodes[i] = r
	}
	p.SetSrcRegionCodes(sSrcRegionCodes)
	sDestRegionCodes := make([]string, len(o.DestRegionCodes))
	for i, r := range o.DestRegionCodes {
		sDestRegionCodes[i] = r
	}
	p.SetDestRegionCodes(sDestRegionCodes)
	sSrcThreatIntelligences := make([]string, len(o.SrcThreatIntelligences))
	for i, r := range o.SrcThreatIntelligences {
		sSrcThreatIntelligences[i] = r
	}
	p.SetSrcThreatIntelligences(sSrcThreatIntelligences)
	sDestThreatIntelligences := make([]string, len(o.DestThreatIntelligences))
	for i, r := range o.DestThreatIntelligences {
		sDestThreatIntelligences[i] = r
	}
	p.SetDestThreatIntelligences(sDestThreatIntelligences)
	sSrcFqdns := make([]string, len(o.SrcFqdns))
	for i, r := range o.SrcFqdns {
		sSrcFqdns[i] = r
	}
	p.SetSrcFqdns(sSrcFqdns)
	sDestFqdns := make([]string, len(o.DestFqdns))
	for i, r := range o.DestFqdns {
		sDestFqdns[i] = r
	}
	p.SetDestFqdns(sDestFqdns)
	sLayer4Configs := make([]*alphapb.ComputeAlphaFirewallPolicyRuleMatchLayer4Configs, len(o.Layer4Configs))
	for i, r := range o.Layer4Configs {
		sLayer4Configs[i] = ComputeAlphaFirewallPolicyRuleMatchLayer4ConfigsToProto(&r)
	}
	p.SetLayer4Configs(sLayer4Configs)
	sSrcAddressGroups := make([]string, len(o.SrcAddressGroups))
	for i, r := range o.SrcAddressGroups {
		sSrcAddressGroups[i] = r
	}
	p.SetSrcAddressGroups(sSrcAddressGroups)
	sDestAddressGroups := make([]string, len(o.DestAddressGroups))
	for i, r := range o.DestAddressGroups {
		sDestAddressGroups[i] = r
	}
	p.SetDestAddressGroups(sDestAddressGroups)
	return p
}

// FirewallPolicyRuleMatchLayer4ConfigsToProto converts a FirewallPolicyRuleMatchLayer4Configs object to its proto representation.
func ComputeAlphaFirewallPolicyRuleMatchLayer4ConfigsToProto(o *alpha.FirewallPolicyRuleMatchLayer4Configs) *alphapb.ComputeAlphaFirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaFirewallPolicyRuleMatchLayer4Configs{}
	p.SetIpProtocol(dcl.ValueOrEmptyString(o.IPProtocol))
	sPorts := make([]string, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	return p
}

// FirewallPolicyRuleToProto converts a FirewallPolicyRule resource to its proto representation.
func FirewallPolicyRuleToProto(resource *alpha.FirewallPolicyRule) *alphapb.ComputeAlphaFirewallPolicyRule {
	p := &alphapb.ComputeAlphaFirewallPolicyRule{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetPriority(dcl.ValueOrEmptyInt64(resource.Priority))
	p.SetMatch(ComputeAlphaFirewallPolicyRuleMatchToProto(resource.Match))
	p.SetAction(dcl.ValueOrEmptyString(resource.Action))
	p.SetDirection(ComputeAlphaFirewallPolicyRuleDirectionEnumToProto(resource.Direction))
	p.SetEnableLogging(dcl.ValueOrEmptyBool(resource.EnableLogging))
	p.SetRuleTupleCount(dcl.ValueOrEmptyInt64(resource.RuleTupleCount))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetKind(dcl.ValueOrEmptyString(resource.Kind))
	p.SetFirewallPolicy(dcl.ValueOrEmptyString(resource.FirewallPolicy))
	sTargetResources := make([]string, len(resource.TargetResources))
	for i, r := range resource.TargetResources {
		sTargetResources[i] = r
	}
	p.SetTargetResources(sTargetResources)
	sTargetServiceAccounts := make([]string, len(resource.TargetServiceAccounts))
	for i, r := range resource.TargetServiceAccounts {
		sTargetServiceAccounts[i] = r
	}
	p.SetTargetServiceAccounts(sTargetServiceAccounts)

	return p
}

// applyFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Apply() method.
func (s *FirewallPolicyRuleServer) applyFirewallPolicyRule(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaFirewallPolicyRuleRequest) (*alphapb.ComputeAlphaFirewallPolicyRule, error) {
	p := ProtoToFirewallPolicyRule(request.GetResource())
	res, err := c.ApplyFirewallPolicyRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallPolicyRuleToProto(res)
	return r, nil
}

// applyComputeAlphaFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Apply() method.
func (s *FirewallPolicyRuleServer) ApplyComputeAlphaFirewallPolicyRule(ctx context.Context, request *alphapb.ApplyComputeAlphaFirewallPolicyRuleRequest) (*alphapb.ComputeAlphaFirewallPolicyRule, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFirewallPolicyRule(ctx, cl, request)
}

// DeleteFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRule Delete() method.
func (s *FirewallPolicyRuleServer) DeleteComputeAlphaFirewallPolicyRule(ctx context.Context, request *alphapb.DeleteComputeAlphaFirewallPolicyRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewallPolicyRule(ctx, ProtoToFirewallPolicyRule(request.GetResource()))

}

// ListComputeAlphaFirewallPolicyRule handles the gRPC request by passing it to the underlying FirewallPolicyRuleList() method.
func (s *FirewallPolicyRuleServer) ListComputeAlphaFirewallPolicyRule(ctx context.Context, request *alphapb.ListComputeAlphaFirewallPolicyRuleRequest) (*alphapb.ListComputeAlphaFirewallPolicyRuleResponse, error) {
	cl, err := createConfigFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewallPolicyRule(ctx, request.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaFirewallPolicyRule
	for _, r := range resources.Items {
		rp := FirewallPolicyRuleToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaFirewallPolicyRuleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFirewallPolicyRule(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
