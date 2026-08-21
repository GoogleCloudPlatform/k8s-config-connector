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

// ForwardingRuleServer implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleIPProtocolEnum(e alphapb.ComputeAlphaForwardingRuleIPProtocolEnum) *alpha.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleIPProtocolEnum(n[len("ComputeAlphaForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleIPVersionEnum(e alphapb.ComputeAlphaForwardingRuleIPVersionEnum) *alpha.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleIPVersionEnum(n[len("ComputeAlphaForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleLoadBalancingSchemeEnum(e alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum) *alpha.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeAlphaForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRuleNetworkTierEnum(e alphapb.ComputeAlphaForwardingRuleNetworkTierEnum) *alpha.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRuleNetworkTierEnum(n[len("ComputeAlphaForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRulePscConnectionStatusEnum converts a ForwardingRulePscConnectionStatusEnum enum from its proto representation.
func ProtoToComputeAlphaForwardingRulePscConnectionStatusEnum(e alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum) *alpha.ForwardingRulePscConnectionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum_name[int32(e)]; ok {
		e := alpha.ForwardingRulePscConnectionStatusEnum(n[len("ComputeAlphaForwardingRulePscConnectionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter object from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilter(p *alphapb.ComputeAlphaForwardingRuleMetadataFilter) *alpha.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeAlphaForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel object from its proto representation.
func ProtoToComputeAlphaForwardingRuleMetadataFilterFilterLabel(p *alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel) *alpha.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &alpha.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToForwardingRuleServiceDirectoryRegistrations converts a ForwardingRuleServiceDirectoryRegistrations object from its proto representation.
func ProtoToComputeAlphaForwardingRuleServiceDirectoryRegistrations(p *alphapb.ComputeAlphaForwardingRuleServiceDirectoryRegistrations) *alpha.ForwardingRuleServiceDirectoryRegistrations {
	if p == nil {
		return nil
	}
	obj := &alpha.ForwardingRuleServiceDirectoryRegistrations{
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *alphapb.ComputeAlphaForwardingRule) *alpha.ForwardingRule {
	obj := &alpha.ForwardingRule{
		AllPorts:             dcl.Bool(p.GetAllPorts()),
		AllowGlobalAccess:    dcl.Bool(p.GetAllowGlobalAccess()),
		LabelFingerprint:     dcl.StringOrNil(p.GetLabelFingerprint()),
		BackendService:       dcl.StringOrNil(p.GetBackendService()),
		CreationTimestamp:    dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		IPAddress:            dcl.StringOrNil(p.GetIpAddress()),
		IPProtocol:           ProtoToComputeAlphaForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeAlphaForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.GetIsMirroringCollector()),
		LoadBalancingScheme:  ProtoToComputeAlphaForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Network:              dcl.StringOrNil(p.GetNetwork()),
		NetworkTier:          ProtoToComputeAlphaForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
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
		PscConnectionStatus:  ProtoToComputeAlphaForwardingRulePscConnectionStatusEnum(p.GetPscConnectionStatus()),
		BaseForwardingRule:   dcl.StringOrNil(p.GetBaseForwardingRule()),
		AllowPscGlobalAccess: dcl.Bool(p.GetAllowPscGlobalAccess()),
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeAlphaForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetServiceDirectoryRegistrations() {
		obj.ServiceDirectoryRegistrations = append(obj.ServiceDirectoryRegistrations, *ProtoToComputeAlphaForwardingRuleServiceDirectoryRegistrations(r))
	}
	for _, r := range p.GetSourceIpRanges() {
		obj.SourceIPRanges = append(obj.SourceIPRanges, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeAlphaForwardingRuleIPProtocolEnumToProto(e *alpha.ForwardingRuleIPProtocolEnum) alphapb.ComputeAlphaForwardingRuleIPProtocolEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeAlphaForwardingRuleIPVersionEnumToProto(e *alpha.ForwardingRuleIPVersionEnum) alphapb.ComputeAlphaForwardingRuleIPVersionEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeAlphaForwardingRuleLoadBalancingSchemeEnumToProto(e *alpha.ForwardingRuleLoadBalancingSchemeEnum) alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *alpha.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeAlphaForwardingRuleNetworkTierEnumToProto(e *alpha.ForwardingRuleNetworkTierEnum) alphapb.ComputeAlphaForwardingRuleNetworkTierEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRuleNetworkTierEnum(0)
}

// ForwardingRulePscConnectionStatusEnumToProto converts a ForwardingRulePscConnectionStatusEnum enum to its proto representation.
func ComputeAlphaForwardingRulePscConnectionStatusEnumToProto(e *alpha.ForwardingRulePscConnectionStatusEnum) alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum_value["ForwardingRulePscConnectionStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum(v)
	}
	return alphapb.ComputeAlphaForwardingRulePscConnectionStatusEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter object to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterToProto(o *alpha.ForwardingRuleMetadataFilter) *alphapb.ComputeAlphaForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaForwardingRuleMetadataFilter{}
	p.SetFilterMatchCriteria(ComputeAlphaForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria))
	sFilterLabel := make([]*alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel, len(o.FilterLabel))
	for i, r := range o.FilterLabel {
		sFilterLabel[i] = ComputeAlphaForwardingRuleMetadataFilterFilterLabelToProto(&r)
	}
	p.SetFilterLabel(sFilterLabel)
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel object to its proto representation.
func ComputeAlphaForwardingRuleMetadataFilterFilterLabelToProto(o *alpha.ForwardingRuleMetadataFilterFilterLabel) *alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaForwardingRuleMetadataFilterFilterLabel{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ForwardingRuleServiceDirectoryRegistrationsToProto converts a ForwardingRuleServiceDirectoryRegistrations object to its proto representation.
func ComputeAlphaForwardingRuleServiceDirectoryRegistrationsToProto(o *alpha.ForwardingRuleServiceDirectoryRegistrations) *alphapb.ComputeAlphaForwardingRuleServiceDirectoryRegistrations {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaForwardingRuleServiceDirectoryRegistrations{}
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *alpha.ForwardingRule) *alphapb.ComputeAlphaForwardingRule {
	p := &alphapb.ComputeAlphaForwardingRule{}
	p.SetAllPorts(dcl.ValueOrEmptyBool(resource.AllPorts))
	p.SetAllowGlobalAccess(dcl.ValueOrEmptyBool(resource.AllowGlobalAccess))
	p.SetLabelFingerprint(dcl.ValueOrEmptyString(resource.LabelFingerprint))
	p.SetBackendService(dcl.ValueOrEmptyString(resource.BackendService))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetIpAddress(dcl.ValueOrEmptyString(resource.IPAddress))
	p.SetIpProtocol(ComputeAlphaForwardingRuleIPProtocolEnumToProto(resource.IPProtocol))
	p.SetIpVersion(ComputeAlphaForwardingRuleIPVersionEnumToProto(resource.IPVersion))
	p.SetIsMirroringCollector(dcl.ValueOrEmptyBool(resource.IsMirroringCollector))
	p.SetLoadBalancingScheme(ComputeAlphaForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetNetworkTier(ComputeAlphaForwardingRuleNetworkTierEnumToProto(resource.NetworkTier))
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
	p.SetPscConnectionStatus(ComputeAlphaForwardingRulePscConnectionStatusEnumToProto(resource.PscConnectionStatus))
	p.SetBaseForwardingRule(dcl.ValueOrEmptyString(resource.BaseForwardingRule))
	p.SetAllowPscGlobalAccess(dcl.ValueOrEmptyBool(resource.AllowPscGlobalAccess))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sMetadataFilter := make([]*alphapb.ComputeAlphaForwardingRuleMetadataFilter, len(resource.MetadataFilter))
	for i, r := range resource.MetadataFilter {
		sMetadataFilter[i] = ComputeAlphaForwardingRuleMetadataFilterToProto(&r)
	}
	p.SetMetadataFilter(sMetadataFilter)
	sPorts := make([]string, len(resource.Ports))
	for i, r := range resource.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	sServiceDirectoryRegistrations := make([]*alphapb.ComputeAlphaForwardingRuleServiceDirectoryRegistrations, len(resource.ServiceDirectoryRegistrations))
	for i, r := range resource.ServiceDirectoryRegistrations {
		sServiceDirectoryRegistrations[i] = ComputeAlphaForwardingRuleServiceDirectoryRegistrationsToProto(&r)
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
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaForwardingRuleRequest) (*alphapb.ComputeAlphaForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// applyComputeAlphaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeAlphaForwardingRule(ctx context.Context, request *alphapb.ApplyComputeAlphaForwardingRuleRequest) (*alphapb.ComputeAlphaForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeAlphaForwardingRule(ctx context.Context, request *alphapb.DeleteComputeAlphaForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeAlphaForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeAlphaForwardingRule(ctx context.Context, request *alphapb.ListComputeAlphaForwardingRuleRequest) (*alphapb.ListComputeAlphaForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaForwardingRuleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
