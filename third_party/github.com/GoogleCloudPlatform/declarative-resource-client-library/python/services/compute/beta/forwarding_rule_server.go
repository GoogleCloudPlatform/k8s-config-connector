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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// ForwardingRuleServer implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleIPProtocolEnum(e betapb.ComputeBetaForwardingRuleIPProtocolEnum) *beta.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleIPProtocolEnum(n[len("ComputeBetaForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleIPVersionEnum(e betapb.ComputeBetaForwardingRuleIPVersionEnum) *beta.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleIPVersionEnum(n[len("ComputeBetaForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleLoadBalancingSchemeEnum(e betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum) *beta.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeBetaForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRuleNetworkTierEnum(e betapb.ComputeBetaForwardingRuleNetworkTierEnum) *beta.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := beta.ForwardingRuleNetworkTierEnum(n[len("ComputeBetaForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRulePscConnectionStatusEnum converts a ForwardingRulePscConnectionStatusEnum enum from its proto representation.
func ProtoToComputeBetaForwardingRulePscConnectionStatusEnum(e betapb.ComputeBetaForwardingRulePscConnectionStatusEnum) *beta.ForwardingRulePscConnectionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaForwardingRulePscConnectionStatusEnum_name[int32(e)]; ok {
		e := beta.ForwardingRulePscConnectionStatusEnum(n[len("ComputeBetaForwardingRulePscConnectionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter object from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilter(p *betapb.ComputeBetaForwardingRuleMetadataFilter) *beta.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &beta.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeBetaForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel object from its proto representation.
func ProtoToComputeBetaForwardingRuleMetadataFilterFilterLabel(p *betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel) *beta.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &beta.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToForwardingRuleServiceDirectoryRegistrations converts a ForwardingRuleServiceDirectoryRegistrations object from its proto representation.
func ProtoToComputeBetaForwardingRuleServiceDirectoryRegistrations(p *betapb.ComputeBetaForwardingRuleServiceDirectoryRegistrations) *beta.ForwardingRuleServiceDirectoryRegistrations {
	if p == nil {
		return nil
	}
	obj := &beta.ForwardingRuleServiceDirectoryRegistrations{
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *betapb.ComputeBetaForwardingRule) *beta.ForwardingRule {
	obj := &beta.ForwardingRule{
		AllPorts:             dcl.Bool(p.GetAllPorts()),
		AllowGlobalAccess:    dcl.Bool(p.GetAllowGlobalAccess()),
		LabelFingerprint:     dcl.StringOrNil(p.GetLabelFingerprint()),
		BackendService:       dcl.StringOrNil(p.GetBackendService()),
		CreationTimestamp:    dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		IPAddress:            dcl.StringOrNil(p.GetIpAddress()),
		IPProtocol:           ProtoToComputeBetaForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeBetaForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.GetIsMirroringCollector()),
		LoadBalancingScheme:  ProtoToComputeBetaForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Network:              dcl.StringOrNil(p.GetNetwork()),
		NetworkTier:          ProtoToComputeBetaForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
		PortRange:            dcl.StringOrNil(p.GetPortRange()),
		Region:               dcl.StringOrNil(p.GetRegion()),
		SelfLink:             dcl.StringOrNil(p.GetSelfLink()),
		ServiceLabel:         dcl.StringOrNil(p.GetServiceLabel()),
		ServiceName:          dcl.StringOrNil(p.GetServiceName()),
		Subnetwork:           dcl.StringOrNil(p.GetSubnetwork()),
		Target:               dcl.StringOrNil(p.GetTarget()),
		Project:              dcl.StringOrNil(p.GetProject()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		PscConnectionId:      dcl.StringOrNil(p.GetPscConnectionId()),
		PscConnectionStatus:  ProtoToComputeBetaForwardingRulePscConnectionStatusEnum(p.GetPscConnectionStatus()),
		BaseForwardingRule:   dcl.StringOrNil(p.GetBaseForwardingRule()),
		AllowPscGlobalAccess: dcl.Bool(p.GetAllowPscGlobalAccess()),
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeBetaForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetServiceDirectoryRegistrations() {
		obj.ServiceDirectoryRegistrations = append(obj.ServiceDirectoryRegistrations, *ProtoToComputeBetaForwardingRuleServiceDirectoryRegistrations(r))
	}
	for _, r := range p.GetSourceIpRanges() {
		obj.SourceIPRanges = append(obj.SourceIPRanges, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeBetaForwardingRuleIPProtocolEnumToProto(e *beta.ForwardingRuleIPProtocolEnum) betapb.ComputeBetaForwardingRuleIPProtocolEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleIPProtocolEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeBetaForwardingRuleIPVersionEnumToProto(e *beta.ForwardingRuleIPVersionEnum) betapb.ComputeBetaForwardingRuleIPVersionEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleIPVersionEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleIPVersionEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeBetaForwardingRuleLoadBalancingSchemeEnumToProto(e *beta.ForwardingRuleLoadBalancingSchemeEnum) betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *beta.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeBetaForwardingRuleNetworkTierEnumToProto(e *beta.ForwardingRuleNetworkTierEnum) betapb.ComputeBetaForwardingRuleNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRuleNetworkTierEnum(v)
	}
	return betapb.ComputeBetaForwardingRuleNetworkTierEnum(0)
}

// ForwardingRulePscConnectionStatusEnumToProto converts a ForwardingRulePscConnectionStatusEnum enum to its proto representation.
func ComputeBetaForwardingRulePscConnectionStatusEnumToProto(e *beta.ForwardingRulePscConnectionStatusEnum) betapb.ComputeBetaForwardingRulePscConnectionStatusEnum {
	if e == nil {
		return betapb.ComputeBetaForwardingRulePscConnectionStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaForwardingRulePscConnectionStatusEnum_value["ForwardingRulePscConnectionStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaForwardingRulePscConnectionStatusEnum(v)
	}
	return betapb.ComputeBetaForwardingRulePscConnectionStatusEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter object to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterToProto(o *beta.ForwardingRuleMetadataFilter) *betapb.ComputeBetaForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaForwardingRuleMetadataFilter{}
	p.SetFilterMatchCriteria(ComputeBetaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria))
	sFilterLabel := make([]*betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel, len(o.FilterLabel))
	for i, r := range o.FilterLabel {
		sFilterLabel[i] = ComputeBetaForwardingRuleMetadataFilterFilterLabelToProto(&r)
	}
	p.SetFilterLabel(sFilterLabel)
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel object to its proto representation.
func ComputeBetaForwardingRuleMetadataFilterFilterLabelToProto(o *beta.ForwardingRuleMetadataFilterFilterLabel) *betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaForwardingRuleMetadataFilterFilterLabel{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ForwardingRuleServiceDirectoryRegistrationsToProto converts a ForwardingRuleServiceDirectoryRegistrations object to its proto representation.
func ComputeBetaForwardingRuleServiceDirectoryRegistrationsToProto(o *beta.ForwardingRuleServiceDirectoryRegistrations) *betapb.ComputeBetaForwardingRuleServiceDirectoryRegistrations {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaForwardingRuleServiceDirectoryRegistrations{}
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *beta.ForwardingRule) *betapb.ComputeBetaForwardingRule {
	p := &betapb.ComputeBetaForwardingRule{}
	p.SetAllPorts(dcl.ValueOrEmptyBool(resource.AllPorts))
	p.SetAllowGlobalAccess(dcl.ValueOrEmptyBool(resource.AllowGlobalAccess))
	p.SetLabelFingerprint(dcl.ValueOrEmptyString(resource.LabelFingerprint))
	p.SetBackendService(dcl.ValueOrEmptyString(resource.BackendService))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetIpAddress(dcl.ValueOrEmptyString(resource.IPAddress))
	p.SetIpProtocol(ComputeBetaForwardingRuleIPProtocolEnumToProto(resource.IPProtocol))
	p.SetIpVersion(ComputeBetaForwardingRuleIPVersionEnumToProto(resource.IPVersion))
	p.SetIsMirroringCollector(dcl.ValueOrEmptyBool(resource.IsMirroringCollector))
	p.SetLoadBalancingScheme(ComputeBetaForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetNetworkTier(ComputeBetaForwardingRuleNetworkTierEnumToProto(resource.NetworkTier))
	p.SetPortRange(dcl.ValueOrEmptyString(resource.PortRange))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetServiceLabel(dcl.ValueOrEmptyString(resource.ServiceLabel))
	p.SetServiceName(dcl.ValueOrEmptyString(resource.ServiceName))
	p.SetSubnetwork(dcl.ValueOrEmptyString(resource.Subnetwork))
	p.SetTarget(dcl.ValueOrEmptyString(resource.Target))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPscConnectionId(dcl.ValueOrEmptyString(resource.PscConnectionId))
	p.SetPscConnectionStatus(ComputeBetaForwardingRulePscConnectionStatusEnumToProto(resource.PscConnectionStatus))
	p.SetBaseForwardingRule(dcl.ValueOrEmptyString(resource.BaseForwardingRule))
	p.SetAllowPscGlobalAccess(dcl.ValueOrEmptyBool(resource.AllowPscGlobalAccess))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sMetadataFilter := make([]*betapb.ComputeBetaForwardingRuleMetadataFilter, len(resource.MetadataFilter))
	for i, r := range resource.MetadataFilter {
		sMetadataFilter[i] = ComputeBetaForwardingRuleMetadataFilterToProto(&r)
	}
	p.SetMetadataFilter(sMetadataFilter)
	sPorts := make([]string, len(resource.Ports))
	for i, r := range resource.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	sServiceDirectoryRegistrations := make([]*betapb.ComputeBetaForwardingRuleServiceDirectoryRegistrations, len(resource.ServiceDirectoryRegistrations))
	for i, r := range resource.ServiceDirectoryRegistrations {
		sServiceDirectoryRegistrations[i] = ComputeBetaForwardingRuleServiceDirectoryRegistrationsToProto(&r)
	}
	p.SetServiceDirectoryRegistrations(sServiceDirectoryRegistrations)
	sSourceIPRanges := make([]string, len(resource.SourceIPRanges))
	for i, r := range resource.SourceIPRanges {
		sSourceIPRanges[i] = r
	}
	p.SetSourceIpRanges(sSourceIPRanges)

	return p
}

// applyForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaForwardingRuleRequest) (*betapb.ComputeBetaForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// applyComputeBetaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeBetaForwardingRule(ctx context.Context, request *betapb.ApplyComputeBetaForwardingRuleRequest) (*betapb.ComputeBetaForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeBetaForwardingRule(ctx context.Context, request *betapb.DeleteComputeBetaForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeBetaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeBetaForwardingRule(ctx context.Context, request *betapb.ListComputeBetaForwardingRuleRequest) (*betapb.ListComputeBetaForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaForwardingRuleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
