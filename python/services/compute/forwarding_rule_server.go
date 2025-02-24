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

// ForwardingRuleServer implements the gRPC interface for ForwardingRule.
type ForwardingRuleServer struct{}

// ProtoToForwardingRuleIPProtocolEnum converts a ForwardingRuleIPProtocolEnum enum from its proto representation.
func ProtoToComputeForwardingRuleIPProtocolEnum(e computepb.ComputeForwardingRuleIPProtocolEnum) *compute.ForwardingRuleIPProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleIPProtocolEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleIPProtocolEnum(n[len("ComputeForwardingRuleIPProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleIPVersionEnum converts a ForwardingRuleIPVersionEnum enum from its proto representation.
func ProtoToComputeForwardingRuleIPVersionEnum(e computepb.ComputeForwardingRuleIPVersionEnum) *compute.ForwardingRuleIPVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleIPVersionEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleIPVersionEnum(n[len("ComputeForwardingRuleIPVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleLoadBalancingSchemeEnum converts a ForwardingRuleLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeForwardingRuleLoadBalancingSchemeEnum(e computepb.ComputeForwardingRuleLoadBalancingSchemeEnum) *compute.ForwardingRuleLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleLoadBalancingSchemeEnum(n[len("ComputeForwardingRuleLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilterFilterMatchCriteriaEnum converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(e computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum) *compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleNetworkTierEnum converts a ForwardingRuleNetworkTierEnum enum from its proto representation.
func ProtoToComputeForwardingRuleNetworkTierEnum(e computepb.ComputeForwardingRuleNetworkTierEnum) *compute.ForwardingRuleNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRuleNetworkTierEnum_name[int32(e)]; ok {
		e := compute.ForwardingRuleNetworkTierEnum(n[len("ComputeForwardingRuleNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRulePscConnectionStatusEnum converts a ForwardingRulePscConnectionStatusEnum enum from its proto representation.
func ProtoToComputeForwardingRulePscConnectionStatusEnum(e computepb.ComputeForwardingRulePscConnectionStatusEnum) *compute.ForwardingRulePscConnectionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeForwardingRulePscConnectionStatusEnum_name[int32(e)]; ok {
		e := compute.ForwardingRulePscConnectionStatusEnum(n[len("ComputeForwardingRulePscConnectionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToForwardingRuleMetadataFilter converts a ForwardingRuleMetadataFilter object from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilter(p *computepb.ComputeForwardingRuleMetadataFilter) *compute.ForwardingRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &compute.ForwardingRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeForwardingRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToForwardingRuleMetadataFilterFilterLabel converts a ForwardingRuleMetadataFilterFilterLabel object from its proto representation.
func ProtoToComputeForwardingRuleMetadataFilterFilterLabel(p *computepb.ComputeForwardingRuleMetadataFilterFilterLabel) *compute.ForwardingRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &compute.ForwardingRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.GetName()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToForwardingRuleServiceDirectoryRegistrations converts a ForwardingRuleServiceDirectoryRegistrations object from its proto representation.
func ProtoToComputeForwardingRuleServiceDirectoryRegistrations(p *computepb.ComputeForwardingRuleServiceDirectoryRegistrations) *compute.ForwardingRuleServiceDirectoryRegistrations {
	if p == nil {
		return nil
	}
	obj := &compute.ForwardingRuleServiceDirectoryRegistrations{
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ProtoToForwardingRule converts a ForwardingRule resource from its proto representation.
func ProtoToForwardingRule(p *computepb.ComputeForwardingRule) *compute.ForwardingRule {
	obj := &compute.ForwardingRule{
		AllPorts:             dcl.Bool(p.GetAllPorts()),
		AllowGlobalAccess:    dcl.Bool(p.GetAllowGlobalAccess()),
		LabelFingerprint:     dcl.StringOrNil(p.GetLabelFingerprint()),
		BackendService:       dcl.StringOrNil(p.GetBackendService()),
		CreationTimestamp:    dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:          dcl.StringOrNil(p.GetDescription()),
		IPAddress:            dcl.StringOrNil(p.GetIpAddress()),
		IPProtocol:           ProtoToComputeForwardingRuleIPProtocolEnum(p.GetIpProtocol()),
		IPVersion:            ProtoToComputeForwardingRuleIPVersionEnum(p.GetIpVersion()),
		IsMirroringCollector: dcl.Bool(p.GetIsMirroringCollector()),
		LoadBalancingScheme:  ProtoToComputeForwardingRuleLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Network:              dcl.StringOrNil(p.GetNetwork()),
		NetworkTier:          ProtoToComputeForwardingRuleNetworkTierEnum(p.GetNetworkTier()),
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
		PscConnectionStatus:  ProtoToComputeForwardingRulePscConnectionStatusEnum(p.GetPscConnectionStatus()),
		BaseForwardingRule:   dcl.StringOrNil(p.GetBaseForwardingRule()),
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeForwardingRuleMetadataFilter(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetServiceDirectoryRegistrations() {
		obj.ServiceDirectoryRegistrations = append(obj.ServiceDirectoryRegistrations, *ProtoToComputeForwardingRuleServiceDirectoryRegistrations(r))
	}
	for _, r := range p.GetSourceIpRanges() {
		obj.SourceIPRanges = append(obj.SourceIPRanges, r)
	}
	return obj
}

// ForwardingRuleIPProtocolEnumToProto converts a ForwardingRuleIPProtocolEnum enum to its proto representation.
func ComputeForwardingRuleIPProtocolEnumToProto(e *compute.ForwardingRuleIPProtocolEnum) computepb.ComputeForwardingRuleIPProtocolEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleIPProtocolEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleIPProtocolEnum_value["ForwardingRuleIPProtocolEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleIPProtocolEnum(v)
	}
	return computepb.ComputeForwardingRuleIPProtocolEnum(0)
}

// ForwardingRuleIPVersionEnumToProto converts a ForwardingRuleIPVersionEnum enum to its proto representation.
func ComputeForwardingRuleIPVersionEnumToProto(e *compute.ForwardingRuleIPVersionEnum) computepb.ComputeForwardingRuleIPVersionEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleIPVersionEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleIPVersionEnum_value["ForwardingRuleIPVersionEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleIPVersionEnum(v)
	}
	return computepb.ComputeForwardingRuleIPVersionEnum(0)
}

// ForwardingRuleLoadBalancingSchemeEnumToProto converts a ForwardingRuleLoadBalancingSchemeEnum enum to its proto representation.
func ComputeForwardingRuleLoadBalancingSchemeEnumToProto(e *compute.ForwardingRuleLoadBalancingSchemeEnum) computepb.ComputeForwardingRuleLoadBalancingSchemeEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleLoadBalancingSchemeEnum_value["ForwardingRuleLoadBalancingSchemeEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(v)
	}
	return computepb.ComputeForwardingRuleLoadBalancingSchemeEnum(0)
}

// ForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a ForwardingRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *compute.ForwardingRuleMetadataFilterFilterMatchCriteriaEnum) computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum_value["ForwardingRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return computepb.ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// ForwardingRuleNetworkTierEnumToProto converts a ForwardingRuleNetworkTierEnum enum to its proto representation.
func ComputeForwardingRuleNetworkTierEnumToProto(e *compute.ForwardingRuleNetworkTierEnum) computepb.ComputeForwardingRuleNetworkTierEnum {
	if e == nil {
		return computepb.ComputeForwardingRuleNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRuleNetworkTierEnum_value["ForwardingRuleNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRuleNetworkTierEnum(v)
	}
	return computepb.ComputeForwardingRuleNetworkTierEnum(0)
}

// ForwardingRulePscConnectionStatusEnumToProto converts a ForwardingRulePscConnectionStatusEnum enum to its proto representation.
func ComputeForwardingRulePscConnectionStatusEnumToProto(e *compute.ForwardingRulePscConnectionStatusEnum) computepb.ComputeForwardingRulePscConnectionStatusEnum {
	if e == nil {
		return computepb.ComputeForwardingRulePscConnectionStatusEnum(0)
	}
	if v, ok := computepb.ComputeForwardingRulePscConnectionStatusEnum_value["ForwardingRulePscConnectionStatusEnum"+string(*e)]; ok {
		return computepb.ComputeForwardingRulePscConnectionStatusEnum(v)
	}
	return computepb.ComputeForwardingRulePscConnectionStatusEnum(0)
}

// ForwardingRuleMetadataFilterToProto converts a ForwardingRuleMetadataFilter object to its proto representation.
func ComputeForwardingRuleMetadataFilterToProto(o *compute.ForwardingRuleMetadataFilter) *computepb.ComputeForwardingRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeForwardingRuleMetadataFilter{}
	p.SetFilterMatchCriteria(ComputeForwardingRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria))
	sFilterLabel := make([]*computepb.ComputeForwardingRuleMetadataFilterFilterLabel, len(o.FilterLabel))
	for i, r := range o.FilterLabel {
		sFilterLabel[i] = ComputeForwardingRuleMetadataFilterFilterLabelToProto(&r)
	}
	p.SetFilterLabel(sFilterLabel)
	return p
}

// ForwardingRuleMetadataFilterFilterLabelToProto converts a ForwardingRuleMetadataFilterFilterLabel object to its proto representation.
func ComputeForwardingRuleMetadataFilterFilterLabelToProto(o *compute.ForwardingRuleMetadataFilterFilterLabel) *computepb.ComputeForwardingRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeForwardingRuleMetadataFilterFilterLabel{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// ForwardingRuleServiceDirectoryRegistrationsToProto converts a ForwardingRuleServiceDirectoryRegistrations object to its proto representation.
func ComputeForwardingRuleServiceDirectoryRegistrationsToProto(o *compute.ForwardingRuleServiceDirectoryRegistrations) *computepb.ComputeForwardingRuleServiceDirectoryRegistrations {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeForwardingRuleServiceDirectoryRegistrations{}
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	return p
}

// ForwardingRuleToProto converts a ForwardingRule resource to its proto representation.
func ForwardingRuleToProto(resource *compute.ForwardingRule) *computepb.ComputeForwardingRule {
	p := &computepb.ComputeForwardingRule{}
	p.SetAllPorts(dcl.ValueOrEmptyBool(resource.AllPorts))
	p.SetAllowGlobalAccess(dcl.ValueOrEmptyBool(resource.AllowGlobalAccess))
	p.SetLabelFingerprint(dcl.ValueOrEmptyString(resource.LabelFingerprint))
	p.SetBackendService(dcl.ValueOrEmptyString(resource.BackendService))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetIpAddress(dcl.ValueOrEmptyString(resource.IPAddress))
	p.SetIpProtocol(ComputeForwardingRuleIPProtocolEnumToProto(resource.IPProtocol))
	p.SetIpVersion(ComputeForwardingRuleIPVersionEnumToProto(resource.IPVersion))
	p.SetIsMirroringCollector(dcl.ValueOrEmptyBool(resource.IsMirroringCollector))
	p.SetLoadBalancingScheme(ComputeForwardingRuleLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetNetworkTier(ComputeForwardingRuleNetworkTierEnumToProto(resource.NetworkTier))
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
	p.SetPscConnectionStatus(ComputeForwardingRulePscConnectionStatusEnumToProto(resource.PscConnectionStatus))
	p.SetBaseForwardingRule(dcl.ValueOrEmptyString(resource.BaseForwardingRule))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sMetadataFilter := make([]*computepb.ComputeForwardingRuleMetadataFilter, len(resource.MetadataFilter))
	for i, r := range resource.MetadataFilter {
		sMetadataFilter[i] = ComputeForwardingRuleMetadataFilterToProto(&r)
	}
	p.SetMetadataFilter(sMetadataFilter)
	sPorts := make([]string, len(resource.Ports))
	for i, r := range resource.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	sServiceDirectoryRegistrations := make([]*computepb.ComputeForwardingRuleServiceDirectoryRegistrations, len(resource.ServiceDirectoryRegistrations))
	for i, r := range resource.ServiceDirectoryRegistrations {
		sServiceDirectoryRegistrations[i] = ComputeForwardingRuleServiceDirectoryRegistrationsToProto(&r)
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
func (s *ForwardingRuleServer) applyForwardingRule(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeForwardingRuleRequest) (*computepb.ComputeForwardingRule, error) {
	p := ProtoToForwardingRule(request.GetResource())
	res, err := c.ApplyForwardingRule(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ForwardingRuleToProto(res)
	return r, nil
}

// applyComputeForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Apply() method.
func (s *ForwardingRuleServer) ApplyComputeForwardingRule(ctx context.Context, request *computepb.ApplyComputeForwardingRuleRequest) (*computepb.ComputeForwardingRule, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyForwardingRule(ctx, cl, request)
}

// DeleteForwardingRule handles the gRPC request by passing it to the underlying ForwardingRule Delete() method.
func (s *ForwardingRuleServer) DeleteComputeForwardingRule(ctx context.Context, request *computepb.DeleteComputeForwardingRuleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteForwardingRule(ctx, ProtoToForwardingRule(request.GetResource()))

}

// ListComputeForwardingRule handles the gRPC request by passing it to the underlying ForwardingRuleList() method.
func (s *ForwardingRuleServer) ListComputeForwardingRule(ctx context.Context, request *computepb.ListComputeForwardingRuleRequest) (*computepb.ListComputeForwardingRuleResponse, error) {
	cl, err := createConfigForwardingRule(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListForwardingRule(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeForwardingRule
	for _, r := range resources.Items {
		rp := ForwardingRuleToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeForwardingRuleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigForwardingRule(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
