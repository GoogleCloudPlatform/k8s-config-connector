// Copyright 2021 Google LLC. All Rights Reserved.
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

// Server implements the gRPC interface for Firewall.
type FirewallServer struct{}

// ProtoToFirewallDirectionEnum converts a FirewallDirectionEnum enum from its proto representation.
func ProtoToComputeBetaFirewallDirectionEnum(e betapb.ComputeBetaFirewallDirectionEnum) *beta.FirewallDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaFirewallDirectionEnum_name[int32(e)]; ok {
		e := beta.FirewallDirectionEnum(n[len("ComputeBetaFirewallDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirewallLogConfig converts a FirewallLogConfig resource from its proto representation.
func ProtoToComputeBetaFirewallLogConfig(p *betapb.ComputeBetaFirewallLogConfig) *beta.FirewallLogConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FirewallLogConfig{
		Enable: dcl.Bool(p.Enable),
	}
	return obj
}

// ProtoToFirewallAllowed converts a FirewallAllowed resource from its proto representation.
func ProtoToComputeBetaFirewallAllowed(p *betapb.ComputeBetaFirewallAllowed) *beta.FirewallAllowed {
	if p == nil {
		return nil
	}
	obj := &beta.FirewallAllowed{
		IPProtocol: dcl.StringOrNil(p.IpProtocol),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetIpProtocolAlt() {
		obj.IPProtocolAlt = append(obj.IPProtocolAlt, r)
	}
	return obj
}

// ProtoToFirewallDenied converts a FirewallDenied resource from its proto representation.
func ProtoToComputeBetaFirewallDenied(p *betapb.ComputeBetaFirewallDenied) *beta.FirewallDenied {
	if p == nil {
		return nil
	}
	obj := &beta.FirewallDenied{
		IPProtocol: dcl.StringOrNil(p.IpProtocol),
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	for _, r := range p.GetIpProtocolAlt() {
		obj.IPProtocolAlt = append(obj.IPProtocolAlt, r)
	}
	return obj
}

// ProtoToFirewall converts a Firewall resource from its proto representation.
func ProtoToFirewall(p *betapb.ComputeBetaFirewall) *beta.Firewall {
	obj := &beta.Firewall{
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:       dcl.StringOrNil(p.Description),
		Direction:         ProtoToComputeBetaFirewallDirectionEnum(p.GetDirection()),
		Disabled:          dcl.Bool(p.Disabled),
		Id:                dcl.StringOrNil(p.Id),
		LogConfig:         ProtoToComputeBetaFirewallLogConfig(p.GetLogConfig()),
		Name:              dcl.StringOrNil(p.Name),
		Network:           dcl.StringOrNil(p.Network),
		Priority:          dcl.Int64OrNil(p.Priority),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetAllowed() {
		obj.Allowed = append(obj.Allowed, *ProtoToComputeBetaFirewallAllowed(r))
	}
	for _, r := range p.GetDenied() {
		obj.Denied = append(obj.Denied, *ProtoToComputeBetaFirewallDenied(r))
	}
	for _, r := range p.GetDestinationRanges() {
		obj.DestinationRanges = append(obj.DestinationRanges, r)
	}
	for _, r := range p.GetSourceRanges() {
		obj.SourceRanges = append(obj.SourceRanges, r)
	}
	for _, r := range p.GetSourceServiceAccounts() {
		obj.SourceServiceAccounts = append(obj.SourceServiceAccounts, r)
	}
	for _, r := range p.GetSourceTags() {
		obj.SourceTags = append(obj.SourceTags, r)
	}
	for _, r := range p.GetTargetServiceAccounts() {
		obj.TargetServiceAccounts = append(obj.TargetServiceAccounts, r)
	}
	for _, r := range p.GetTargetTags() {
		obj.TargetTags = append(obj.TargetTags, r)
	}
	return obj
}

// FirewallDirectionEnumToProto converts a FirewallDirectionEnum enum to its proto representation.
func ComputeBetaFirewallDirectionEnumToProto(e *beta.FirewallDirectionEnum) betapb.ComputeBetaFirewallDirectionEnum {
	if e == nil {
		return betapb.ComputeBetaFirewallDirectionEnum(0)
	}
	if v, ok := betapb.ComputeBetaFirewallDirectionEnum_value["FirewallDirectionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaFirewallDirectionEnum(v)
	}
	return betapb.ComputeBetaFirewallDirectionEnum(0)
}

// FirewallLogConfigToProto converts a FirewallLogConfig resource to its proto representation.
func ComputeBetaFirewallLogConfigToProto(o *beta.FirewallLogConfig) *betapb.ComputeBetaFirewallLogConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaFirewallLogConfig{
		Enable: dcl.ValueOrEmptyBool(o.Enable),
	}
	return p
}

// FirewallAllowedToProto converts a FirewallAllowed resource to its proto representation.
func ComputeBetaFirewallAllowedToProto(o *beta.FirewallAllowed) *betapb.ComputeBetaFirewallAllowed {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaFirewallAllowed{
		IpProtocol: dcl.ValueOrEmptyString(o.IPProtocol),
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	for _, r := range o.IPProtocolAlt {
		p.IpProtocolAlt = append(p.IpProtocolAlt, r)
	}
	return p
}

// FirewallDeniedToProto converts a FirewallDenied resource to its proto representation.
func ComputeBetaFirewallDeniedToProto(o *beta.FirewallDenied) *betapb.ComputeBetaFirewallDenied {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaFirewallDenied{
		IpProtocol: dcl.ValueOrEmptyString(o.IPProtocol),
	}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	for _, r := range o.IPProtocolAlt {
		p.IpProtocolAlt = append(p.IpProtocolAlt, r)
	}
	return p
}

// FirewallToProto converts a Firewall resource to its proto representation.
func FirewallToProto(resource *beta.Firewall) *betapb.ComputeBetaFirewall {
	p := &betapb.ComputeBetaFirewall{
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Direction:         ComputeBetaFirewallDirectionEnumToProto(resource.Direction),
		Disabled:          dcl.ValueOrEmptyBool(resource.Disabled),
		Id:                dcl.ValueOrEmptyString(resource.Id),
		LogConfig:         ComputeBetaFirewallLogConfigToProto(resource.LogConfig),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Network:           dcl.ValueOrEmptyString(resource.Network),
		Priority:          dcl.ValueOrEmptyInt64(resource.Priority),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Allowed {
		p.Allowed = append(p.Allowed, ComputeBetaFirewallAllowedToProto(&r))
	}
	for _, r := range resource.Denied {
		p.Denied = append(p.Denied, ComputeBetaFirewallDeniedToProto(&r))
	}
	for _, r := range resource.DestinationRanges {
		p.DestinationRanges = append(p.DestinationRanges, r)
	}
	for _, r := range resource.SourceRanges {
		p.SourceRanges = append(p.SourceRanges, r)
	}
	for _, r := range resource.SourceServiceAccounts {
		p.SourceServiceAccounts = append(p.SourceServiceAccounts, r)
	}
	for _, r := range resource.SourceTags {
		p.SourceTags = append(p.SourceTags, r)
	}
	for _, r := range resource.TargetServiceAccounts {
		p.TargetServiceAccounts = append(p.TargetServiceAccounts, r)
	}
	for _, r := range resource.TargetTags {
		p.TargetTags = append(p.TargetTags, r)
	}

	return p
}

// ApplyFirewall handles the gRPC request by passing it to the underlying Firewall Apply() method.
func (s *FirewallServer) applyFirewall(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaFirewallRequest) (*betapb.ComputeBetaFirewall, error) {
	p := ProtoToFirewall(request.GetResource())
	res, err := c.ApplyFirewall(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirewallToProto(res)
	return r, nil
}

// ApplyFirewall handles the gRPC request by passing it to the underlying Firewall Apply() method.
func (s *FirewallServer) ApplyComputeBetaFirewall(ctx context.Context, request *betapb.ApplyComputeBetaFirewallRequest) (*betapb.ComputeBetaFirewall, error) {
	cl, err := createConfigFirewall(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFirewall(ctx, cl, request)
}

// DeleteFirewall handles the gRPC request by passing it to the underlying Firewall Delete() method.
func (s *FirewallServer) DeleteComputeBetaFirewall(ctx context.Context, request *betapb.DeleteComputeBetaFirewallRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFirewall(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFirewall(ctx, ProtoToFirewall(request.GetResource()))

}

// ListComputeBetaFirewall handles the gRPC request by passing it to the underlying FirewallList() method.
func (s *FirewallServer) ListComputeBetaFirewall(ctx context.Context, request *betapb.ListComputeBetaFirewallRequest) (*betapb.ListComputeBetaFirewallResponse, error) {
	cl, err := createConfigFirewall(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirewall(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaFirewall
	for _, r := range resources.Items {
		rp := FirewallToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaFirewallResponse{Items: protos}, nil
}

func createConfigFirewall(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
