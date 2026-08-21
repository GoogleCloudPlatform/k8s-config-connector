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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// NetworkFirewallPolicyRuleServer implements the gRPC interface for NetworkFirewallPolicyRule.
type NetworkFirewallPolicyRuleServer struct{}

// ProtoToNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum converts a NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum enum from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(e computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum) *compute.NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum_name[int32(e)]; ok {
		e := compute.NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(n[len("ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkFirewallPolicyRuleDirectionEnum converts a NetworkFirewallPolicyRuleDirectionEnum enum from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleDirectionEnum(e computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum) *compute.NetworkFirewallPolicyRuleDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum_name[int32(e)]; ok {
		e := compute.NetworkFirewallPolicyRuleDirectionEnum(n[len("ComputeNetworkFirewallPolicyRuleDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkFirewallPolicyRuleTargetSecureTagsStateEnum converts a NetworkFirewallPolicyRuleTargetSecureTagsStateEnum enum from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum(e computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum) *compute.NetworkFirewallPolicyRuleTargetSecureTagsStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum_name[int32(e)]; ok {
		e := compute.NetworkFirewallPolicyRuleTargetSecureTagsStateEnum(n[len("ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNetworkFirewallPolicyRuleMatch converts a NetworkFirewallPolicyRuleMatch object from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleMatch(p *computepb.ComputeNetworkFirewallPolicyRuleMatch) *compute.NetworkFirewallPolicyRuleMatch {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkFirewallPolicyRuleMatch{}
	for _, r := range p.GetSrcIpRanges() {
		obj.SrcIPRanges = append(obj.SrcIPRanges, r)
	}
	for _, r := range p.GetDestIpRanges() {
		obj.DestIPRanges = append(obj.DestIPRanges, r)
	}
	for _, r := range p.GetLayer4Configs() {
		obj.Layer4Configs = append(obj.Layer4Configs, *ProtoToComputeNetworkFirewallPolicyRuleMatchLayer4Configs(r))
	}
	for _, r := range p.GetSrcSecureTags() {
		obj.SrcSecureTags = append(obj.SrcSecureTags, *ProtoToComputeNetworkFirewallPolicyRuleMatchSrcSecureTags(r))
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
	for _, r := range p.GetSrcAddressGroups() {
		obj.SrcAddressGroups = append(obj.SrcAddressGroups, r)
	}
	for _, r := range p.GetDestAddressGroups() {
		obj.DestAddressGroups = append(obj.DestAddressGroups, r)
	}
	return obj
}

// ProtoToNetworkFirewallPolicyRuleMatchLayer4Configs converts a NetworkFirewallPolicyRuleMatchLayer4Configs object from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleMatchLayer4Configs(p *computepb.ComputeNetworkFirewallPolicyRuleMatchLayer4Configs) *compute.NetworkFirewallPolicyRuleMatchLayer4Configs {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkFirewallPolicyRuleMatchLayer4Configs{
		IPProtocol: dcl.StringOrNil(p.GetIpProtocol()),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToNetworkFirewallPolicyRuleMatchSrcSecureTags converts a NetworkFirewallPolicyRuleMatchSrcSecureTags object from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleMatchSrcSecureTags(p *computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTags) *compute.NetworkFirewallPolicyRuleMatchSrcSecureTags {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkFirewallPolicyRuleMatchSrcSecureTags{
		Name:  dcl.StringOrNil(p.GetName()),
		State: ProtoToComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToNetworkFirewallPolicyRuleTargetSecureTags converts a NetworkFirewallPolicyRuleTargetSecureTags object from its proto representation.
func ProtoToComputeNetworkFirewallPolicyRuleTargetSecureTags(p *computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTags) *compute.NetworkFirewallPolicyRuleTargetSecureTags {
	if p == nil {
		return nil
	}
	obj := &compute.NetworkFirewallPolicyRuleTargetSecureTags{
		Name:  dcl.StringOrNil(p.GetName()),
		State: ProtoToComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToNetworkFirewallPolicyRule converts a NetworkFirewallPolicyRule resource from its proto representation.
func ProtoToNetworkFirewallPolicyRule(p *computepb.ComputeNetworkFirewallPolicyRule) *compute.NetworkFirewallPolicyRule {
	obj := &compute.NetworkFirewallPolicyRule{
		Description:    dcl.StringOrNil(p.GetDescription()),
		RuleName:       dcl.StringOrNil(p.GetRuleName()),
		Priority:       dcl.Int64OrNil(p.GetPriority()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Match:          ProtoToComputeNetworkFirewallPolicyRuleMatch(p.GetMatch()),
		Action:         dcl.StringOrNil(p.GetAction()),
		Direction:      ProtoToComputeNetworkFirewallPolicyRuleDirectionEnum(p.GetDirection()),
		EnableLogging:  dcl.Bool(p.GetEnableLogging()),
		RuleTupleCount: dcl.Int64OrNil(p.GetRuleTupleCount()),
		Disabled:       dcl.Bool(p.GetDisabled()),
		Kind:           dcl.StringOrNil(p.GetKind()),
		FirewallPolicy: dcl.StringOrNil(p.GetFirewallPolicy()),
		Project:        dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetTargetServiceAccounts() {
		obj.TargetServiceAccounts = append(obj.TargetServiceAccounts, r)
	}
	for _, r := range p.GetTargetSecureTags() {
		obj.TargetSecureTags = append(obj.TargetSecureTags, *ProtoToComputeNetworkFirewallPolicyRuleTargetSecureTags(r))
	}
	return obj
}

// NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnumToProto converts a NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum enum to its proto representation.
func ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnumToProto(e *compute.NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum) computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum {
	if e == nil {
		return computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(0)
	}
	if v, ok := computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum_value["NetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum"+string(*e)]; ok {
		return computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(v)
	}
	return computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnum(0)
}

// NetworkFirewallPolicyRuleDirectionEnumToProto converts a NetworkFirewallPolicyRuleDirectionEnum enum to its proto representation.
func ComputeNetworkFirewallPolicyRuleDirectionEnumToProto(e *compute.NetworkFirewallPolicyRuleDirectionEnum) computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum {
	if e == nil {
		return computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum(0)
	}
	if v, ok := computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum_value["NetworkFirewallPolicyRuleDirectionEnum"+string(*e)]; ok {
		return computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum(v)
	}
	return computepb.ComputeNetworkFirewallPolicyRuleDirectionEnum(0)
}

// NetworkFirewallPolicyRuleTargetSecureTagsStateEnumToProto converts a NetworkFirewallPolicyRuleTargetSecureTagsStateEnum enum to its proto representation.
func ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnumToProto(e *compute.NetworkFirewallPolicyRuleTargetSecureTagsStateEnum) computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum {
	if e == nil {
		return computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum(0)
	}
	if v, ok := computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum_value["NetworkFirewallPolicyRuleTargetSecureTagsStateEnum"+string(*e)]; ok {
		return computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum(v)
	}
	return computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnum(0)
}

// NetworkFirewallPolicyRuleMatchToProto converts a NetworkFirewallPolicyRuleMatch object to its proto representation.
func ComputeNetworkFirewallPolicyRuleMatchToProto(o *compute.NetworkFirewallPolicyRuleMatch) *computepb.ComputeNetworkFirewallPolicyRuleMatch {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkFirewallPolicyRuleMatch{}
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
	sLayer4Configs := make([]*computepb.ComputeNetworkFirewallPolicyRuleMatchLayer4Configs, len(o.Layer4Configs))
	for i, r := range o.Layer4Configs {
		sLayer4Configs[i] = ComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsToProto(&r)
	}
	p.SetLayer4Configs(sLayer4Configs)
	sSrcSecureTags := make([]*computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTags, len(o.SrcSecureTags))
	for i, r := range o.SrcSecureTags {
		sSrcSecureTags[i] = ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsToProto(&r)
	}
	p.SetSrcSecureTags(sSrcSecureTags)
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

// NetworkFirewallPolicyRuleMatchLayer4ConfigsToProto converts a NetworkFirewallPolicyRuleMatchLayer4Configs object to its proto representation.
func ComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsToProto(o *compute.NetworkFirewallPolicyRuleMatchLayer4Configs) *computepb.ComputeNetworkFirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkFirewallPolicyRuleMatchLayer4Configs{}
	p.SetIpProtocol(dcl.ValueOrEmptyString(o.IPProtocol))
	sPorts := make([]string, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	return p
}

// NetworkFirewallPolicyRuleMatchSrcSecureTagsToProto converts a NetworkFirewallPolicyRuleMatchSrcSecureTags object to its proto representation.
func ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsToProto(o *compute.NetworkFirewallPolicyRuleMatchSrcSecureTags) *computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTags {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkFirewallPolicyRuleMatchSrcSecureTags{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetState(ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsStateEnumToProto(o.State))
	return p
}

// NetworkFirewallPolicyRuleTargetSecureTagsToProto converts a NetworkFirewallPolicyRuleTargetSecureTags object to its proto representation.
func ComputeNetworkFirewallPolicyRuleTargetSecureTagsToProto(o *compute.NetworkFirewallPolicyRuleTargetSecureTags) *computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTags {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTags{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetState(ComputeNetworkFirewallPolicyRuleTargetSecureTagsStateEnumToProto(o.State))
	return p
}

// NetworkFirewallPolicyRuleToProto converts a NetworkFirewallPolicyRule resource to its proto representation.
func NetworkFirewallPolicyRuleToProto(resource *compute.NetworkFirewallPolicyRule) *computepb.ComputeNetworkFirewallPolicyRule {
	p := &computepb.ComputeNetworkFirewallPolicyRule{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRuleName(dcl.ValueOrEmptyString(resource.RuleName))
	p.SetPriority(dcl.ValueOrEmptyInt64(resource.Priority))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetMatch(ComputeNetworkFirewallPolicyRuleMatchToProto(resource.Match))
	p.SetAction(dcl.ValueOrEmptyString(resource.Action))
	p.SetDirection(ComputeNetworkFirewallPolicyRuleDirectionEnumToProto(resource.Direction))
	p.SetEnableLogging(dcl.ValueOrEmptyBool(resource.EnableLogging))
	p.SetRuleTupleCount(dcl.ValueOrEmptyInt64(resource.RuleTupleCount))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetKind(dcl.ValueOrEmptyString(resource.Kind))
	p.SetFirewallPolicy(dcl.ValueOrEmptyString(resource.FirewallPolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sTargetServiceAccounts := make([]string, len(resource.TargetServiceAccounts))
	for i, r := range resource.TargetServiceAccounts {
		sTargetServiceAccounts[i] = r
	}
	p.SetTargetServiceAccounts(sTargetServiceAccounts)
	sTargetSecureTags := make([]*computepb.ComputeNetworkFirewallPolicyRuleTargetSecureTags, len(resource.TargetSecureTags))
	for i, r := range resource.TargetSecureTags {
		sTargetSecureTags[i] = ComputeNetworkFirewallPolicyRuleTargetSecureTagsToProto(&r)
	}
	p.SetTargetSecureTags(sTargetSecureTags)

	return p
}

// applyNetworkFirewallPolicyRule handles the gRPC request by passing it to the underlying NetworkFirewallPolicyRule Apply() method.
func (s *NetworkFirewallPolicyRuleServer) applyNetworkFirewallPolicyRule(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeNetworkFirewallPolicyRuleRequest) (*computepb.ComputeNetworkFirewallPolicyRule, error) {
	p := ProtoToNetworkFirewallPolicyRule(request.GetResource())
	res, err := c.ApplyNetworkFirewallPolicyRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NetworkFirewallPolicyRuleToProto(res)
	return r, nil
}

// applyComputeNetworkFirewallPolicyRule handles the gRPC request by passing it to the underlying NetworkFirewallPolicyRule Apply() method.
func (s *NetworkFirewallPolicyRuleServer) ApplyComputeNetworkFirewallPolicyRule(ctx context.Context, request *computepb.ApplyComputeNetworkFirewallPolicyRuleRequest) (*computepb.ComputeNetworkFirewallPolicyRule, error) {
	cl, err := createConfigNetworkFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNetworkFirewallPolicyRule(ctx, cl, request)
}

// DeleteNetworkFirewallPolicyRule handles the gRPC request by passing it to the underlying NetworkFirewallPolicyRule Delete() method.
func (s *NetworkFirewallPolicyRuleServer) DeleteComputeNetworkFirewallPolicyRule(ctx context.Context, request *computepb.DeleteComputeNetworkFirewallPolicyRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNetworkFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNetworkFirewallPolicyRule(ctx, ProtoToNetworkFirewallPolicyRule(request.GetResource()))

}

// ListComputeNetworkFirewallPolicyRule handles the gRPC request by passing it to the underlying NetworkFirewallPolicyRuleList() method.
func (s *NetworkFirewallPolicyRuleServer) ListComputeNetworkFirewallPolicyRule(ctx context.Context, request *computepb.ListComputeNetworkFirewallPolicyRuleRequest) (*computepb.ListComputeNetworkFirewallPolicyRuleResponse, error) {
	cl, err := createConfigNetworkFirewallPolicyRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNetworkFirewallPolicyRule(ctx, request.GetProject(), request.GetLocation(), request.GetFirewallPolicy())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeNetworkFirewallPolicyRule
	for _, r := range resources.Items {
		rp := NetworkFirewallPolicyRuleToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeNetworkFirewallPolicyRuleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNetworkFirewallPolicyRule(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
