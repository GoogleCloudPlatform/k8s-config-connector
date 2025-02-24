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
	dnspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dns/dns_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns"
)

// Server implements the gRPC interface for ManagedZone.
type ManagedZoneServer struct{}

// ProtoToManagedZoneDnssecConfigNonExistenceEnum converts a ManagedZoneDnssecConfigNonExistenceEnum enum from its proto representation.
func ProtoToDnsManagedZoneDnssecConfigNonExistenceEnum(e dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum) *dns.ManagedZoneDnssecConfigNonExistenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneDnssecConfigNonExistenceEnum(n[len("DnsManagedZoneDnssecConfigNonExistenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneDnssecConfigStateEnum converts a ManagedZoneDnssecConfigStateEnum enum from its proto representation.
func ProtoToDnsManagedZoneDnssecConfigStateEnum(e dnspb.DnsManagedZoneDnssecConfigStateEnum) *dns.ManagedZoneDnssecConfigStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneDnssecConfigStateEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneDnssecConfigStateEnum(n[len("DnsManagedZoneDnssecConfigStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum converts a ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum enum from its proto representation.
func ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(e dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum) *dns.ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(n[len("DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum converts a ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum enum from its proto representation.
func ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(e dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum) *dns.ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(n[len("DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneVisibilityEnum converts a ManagedZoneVisibilityEnum enum from its proto representation.
func ProtoToDnsManagedZoneVisibilityEnum(e dnspb.DnsManagedZoneVisibilityEnum) *dns.ManagedZoneVisibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneVisibilityEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneVisibilityEnum(n[len("DnsManagedZoneVisibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneForwardingConfigTargetNameServersForwardingPathEnum converts a ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum enum from its proto representation.
func ProtoToDnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(e dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum) *dns.ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum_name[int32(e)]; ok {
		e := dns.ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(n[len("DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedZoneDnssecConfig converts a ManagedZoneDnssecConfig resource from its proto representation.
func ProtoToDnsManagedZoneDnssecConfig(p *dnspb.DnsManagedZoneDnssecConfig) *dns.ManagedZoneDnssecConfig {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZoneDnssecConfig{
		Kind:         dcl.StringOrNil(p.Kind),
		NonExistence: ProtoToDnsManagedZoneDnssecConfigNonExistenceEnum(p.GetNonExistence()),
		State:        ProtoToDnsManagedZoneDnssecConfigStateEnum(p.GetState()),
	}
	for _, r := range p.GetDefaultKeySpecs() {
		obj.DefaultKeySpecs = append(obj.DefaultKeySpecs, *ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecs(r))
	}
	return obj
}

// ProtoToManagedZoneDnssecConfigDefaultKeySpecs converts a ManagedZoneDnssecConfigDefaultKeySpecs resource from its proto representation.
func ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecs(p *dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecs) *dns.ManagedZoneDnssecConfigDefaultKeySpecs {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZoneDnssecConfigDefaultKeySpecs{
		Algorithm: ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(p.GetAlgorithm()),
		KeyLength: dcl.Int64OrNil(p.KeyLength),
		KeyType:   ProtoToDnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(p.GetKeyType()),
		Kind:      dcl.StringOrNil(p.Kind),
	}
	return obj
}

// ProtoToManagedZonePrivateVisibilityConfig converts a ManagedZonePrivateVisibilityConfig resource from its proto representation.
func ProtoToDnsManagedZonePrivateVisibilityConfig(p *dnspb.DnsManagedZonePrivateVisibilityConfig) *dns.ManagedZonePrivateVisibilityConfig {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZonePrivateVisibilityConfig{}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToDnsManagedZonePrivateVisibilityConfigNetworks(r))
	}
	return obj
}

// ProtoToManagedZonePrivateVisibilityConfigNetworks converts a ManagedZonePrivateVisibilityConfigNetworks resource from its proto representation.
func ProtoToDnsManagedZonePrivateVisibilityConfigNetworks(p *dnspb.DnsManagedZonePrivateVisibilityConfigNetworks) *dns.ManagedZonePrivateVisibilityConfigNetworks {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZonePrivateVisibilityConfigNetworks{
		NetworkUrl: dcl.StringOrNil(p.NetworkUrl),
	}
	return obj
}

// ProtoToManagedZoneForwardingConfig converts a ManagedZoneForwardingConfig resource from its proto representation.
func ProtoToDnsManagedZoneForwardingConfig(p *dnspb.DnsManagedZoneForwardingConfig) *dns.ManagedZoneForwardingConfig {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZoneForwardingConfig{}
	for _, r := range p.GetTargetNameServers() {
		obj.TargetNameServers = append(obj.TargetNameServers, *ProtoToDnsManagedZoneForwardingConfigTargetNameServers(r))
	}
	return obj
}

// ProtoToManagedZoneForwardingConfigTargetNameServers converts a ManagedZoneForwardingConfigTargetNameServers resource from its proto representation.
func ProtoToDnsManagedZoneForwardingConfigTargetNameServers(p *dnspb.DnsManagedZoneForwardingConfigTargetNameServers) *dns.ManagedZoneForwardingConfigTargetNameServers {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZoneForwardingConfigTargetNameServers{
		IPv4Address:    dcl.StringOrNil(p.Ipv4Address),
		ForwardingPath: ProtoToDnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(p.GetForwardingPath()),
	}
	return obj
}

// ProtoToManagedZonePeeringConfig converts a ManagedZonePeeringConfig resource from its proto representation.
func ProtoToDnsManagedZonePeeringConfig(p *dnspb.DnsManagedZonePeeringConfig) *dns.ManagedZonePeeringConfig {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZonePeeringConfig{
		TargetNetwork: ProtoToDnsManagedZonePeeringConfigTargetNetwork(p.GetTargetNetwork()),
	}
	return obj
}

// ProtoToManagedZonePeeringConfigTargetNetwork converts a ManagedZonePeeringConfigTargetNetwork resource from its proto representation.
func ProtoToDnsManagedZonePeeringConfigTargetNetwork(p *dnspb.DnsManagedZonePeeringConfigTargetNetwork) *dns.ManagedZonePeeringConfigTargetNetwork {
	if p == nil {
		return nil
	}
	obj := &dns.ManagedZonePeeringConfigTargetNetwork{
		NetworkUrl: dcl.StringOrNil(p.NetworkUrl),
	}
	return obj
}

// ProtoToManagedZone converts a ManagedZone resource from its proto representation.
func ProtoToManagedZone(p *dnspb.DnsManagedZone) *dns.ManagedZone {
	obj := &dns.ManagedZone{
		Description:             dcl.StringOrNil(p.Description),
		DnsName:                 dcl.StringOrNil(p.DnsName),
		DnssecConfig:            ProtoToDnsManagedZoneDnssecConfig(p.GetDnssecConfig()),
		Name:                    dcl.StringOrNil(p.Name),
		Visibility:              ProtoToDnsManagedZoneVisibilityEnum(p.GetVisibility()),
		PrivateVisibilityConfig: ProtoToDnsManagedZonePrivateVisibilityConfig(p.GetPrivateVisibilityConfig()),
		ForwardingConfig:        ProtoToDnsManagedZoneForwardingConfig(p.GetForwardingConfig()),
		ReverseLookup:           dcl.Bool(p.ReverseLookup),
		PeeringConfig:           ProtoToDnsManagedZonePeeringConfig(p.GetPeeringConfig()),
		Project:                 dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetNameServers() {
		obj.NameServers = append(obj.NameServers, r)
	}
	return obj
}

// ManagedZoneDnssecConfigNonExistenceEnumToProto converts a ManagedZoneDnssecConfigNonExistenceEnum enum to its proto representation.
func DnsManagedZoneDnssecConfigNonExistenceEnumToProto(e *dns.ManagedZoneDnssecConfigNonExistenceEnum) dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum {
	if e == nil {
		return dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum_value["ManagedZoneDnssecConfigNonExistenceEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum(v)
	}
	return dnspb.DnsManagedZoneDnssecConfigNonExistenceEnum(0)
}

// ManagedZoneDnssecConfigStateEnumToProto converts a ManagedZoneDnssecConfigStateEnum enum to its proto representation.
func DnsManagedZoneDnssecConfigStateEnumToProto(e *dns.ManagedZoneDnssecConfigStateEnum) dnspb.DnsManagedZoneDnssecConfigStateEnum {
	if e == nil {
		return dnspb.DnsManagedZoneDnssecConfigStateEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneDnssecConfigStateEnum_value["ManagedZoneDnssecConfigStateEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneDnssecConfigStateEnum(v)
	}
	return dnspb.DnsManagedZoneDnssecConfigStateEnum(0)
}

// ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumToProto converts a ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum enum to its proto representation.
func DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumToProto(e *dns.ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum) dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum {
	if e == nil {
		return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum_value["ManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(v)
	}
	return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnum(0)
}

// ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumToProto converts a ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum enum to its proto representation.
func DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumToProto(e *dns.ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum) dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum {
	if e == nil {
		return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum_value["ManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(v)
	}
	return dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnum(0)
}

// ManagedZoneVisibilityEnumToProto converts a ManagedZoneVisibilityEnum enum to its proto representation.
func DnsManagedZoneVisibilityEnumToProto(e *dns.ManagedZoneVisibilityEnum) dnspb.DnsManagedZoneVisibilityEnum {
	if e == nil {
		return dnspb.DnsManagedZoneVisibilityEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneVisibilityEnum_value["ManagedZoneVisibilityEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneVisibilityEnum(v)
	}
	return dnspb.DnsManagedZoneVisibilityEnum(0)
}

// ManagedZoneForwardingConfigTargetNameServersForwardingPathEnumToProto converts a ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum enum to its proto representation.
func DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnumToProto(e *dns.ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum) dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum {
	if e == nil {
		return dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(0)
	}
	if v, ok := dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum_value["ManagedZoneForwardingConfigTargetNameServersForwardingPathEnum"+string(*e)]; ok {
		return dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(v)
	}
	return dnspb.DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnum(0)
}

// ManagedZoneDnssecConfigToProto converts a ManagedZoneDnssecConfig resource to its proto representation.
func DnsManagedZoneDnssecConfigToProto(o *dns.ManagedZoneDnssecConfig) *dnspb.DnsManagedZoneDnssecConfig {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZoneDnssecConfig{
		Kind:         dcl.ValueOrEmptyString(o.Kind),
		NonExistence: DnsManagedZoneDnssecConfigNonExistenceEnumToProto(o.NonExistence),
		State:        DnsManagedZoneDnssecConfigStateEnumToProto(o.State),
	}
	for _, r := range o.DefaultKeySpecs {
		p.DefaultKeySpecs = append(p.DefaultKeySpecs, DnsManagedZoneDnssecConfigDefaultKeySpecsToProto(&r))
	}
	return p
}

// ManagedZoneDnssecConfigDefaultKeySpecsToProto converts a ManagedZoneDnssecConfigDefaultKeySpecs resource to its proto representation.
func DnsManagedZoneDnssecConfigDefaultKeySpecsToProto(o *dns.ManagedZoneDnssecConfigDefaultKeySpecs) *dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecs {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZoneDnssecConfigDefaultKeySpecs{
		Algorithm: DnsManagedZoneDnssecConfigDefaultKeySpecsAlgorithmEnumToProto(o.Algorithm),
		KeyLength: dcl.ValueOrEmptyInt64(o.KeyLength),
		KeyType:   DnsManagedZoneDnssecConfigDefaultKeySpecsKeyTypeEnumToProto(o.KeyType),
		Kind:      dcl.ValueOrEmptyString(o.Kind),
	}
	return p
}

// ManagedZonePrivateVisibilityConfigToProto converts a ManagedZonePrivateVisibilityConfig resource to its proto representation.
func DnsManagedZonePrivateVisibilityConfigToProto(o *dns.ManagedZonePrivateVisibilityConfig) *dnspb.DnsManagedZonePrivateVisibilityConfig {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZonePrivateVisibilityConfig{}
	for _, r := range o.Networks {
		p.Networks = append(p.Networks, DnsManagedZonePrivateVisibilityConfigNetworksToProto(&r))
	}
	return p
}

// ManagedZonePrivateVisibilityConfigNetworksToProto converts a ManagedZonePrivateVisibilityConfigNetworks resource to its proto representation.
func DnsManagedZonePrivateVisibilityConfigNetworksToProto(o *dns.ManagedZonePrivateVisibilityConfigNetworks) *dnspb.DnsManagedZonePrivateVisibilityConfigNetworks {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZonePrivateVisibilityConfigNetworks{
		NetworkUrl: dcl.ValueOrEmptyString(o.NetworkUrl),
	}
	return p
}

// ManagedZoneForwardingConfigToProto converts a ManagedZoneForwardingConfig resource to its proto representation.
func DnsManagedZoneForwardingConfigToProto(o *dns.ManagedZoneForwardingConfig) *dnspb.DnsManagedZoneForwardingConfig {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZoneForwardingConfig{}
	for _, r := range o.TargetNameServers {
		p.TargetNameServers = append(p.TargetNameServers, DnsManagedZoneForwardingConfigTargetNameServersToProto(&r))
	}
	return p
}

// ManagedZoneForwardingConfigTargetNameServersToProto converts a ManagedZoneForwardingConfigTargetNameServers resource to its proto representation.
func DnsManagedZoneForwardingConfigTargetNameServersToProto(o *dns.ManagedZoneForwardingConfigTargetNameServers) *dnspb.DnsManagedZoneForwardingConfigTargetNameServers {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZoneForwardingConfigTargetNameServers{
		Ipv4Address:    dcl.ValueOrEmptyString(o.IPv4Address),
		ForwardingPath: DnsManagedZoneForwardingConfigTargetNameServersForwardingPathEnumToProto(o.ForwardingPath),
	}
	return p
}

// ManagedZonePeeringConfigToProto converts a ManagedZonePeeringConfig resource to its proto representation.
func DnsManagedZonePeeringConfigToProto(o *dns.ManagedZonePeeringConfig) *dnspb.DnsManagedZonePeeringConfig {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZonePeeringConfig{
		TargetNetwork: DnsManagedZonePeeringConfigTargetNetworkToProto(o.TargetNetwork),
	}
	return p
}

// ManagedZonePeeringConfigTargetNetworkToProto converts a ManagedZonePeeringConfigTargetNetwork resource to its proto representation.
func DnsManagedZonePeeringConfigTargetNetworkToProto(o *dns.ManagedZonePeeringConfigTargetNetwork) *dnspb.DnsManagedZonePeeringConfigTargetNetwork {
	if o == nil {
		return nil
	}
	p := &dnspb.DnsManagedZonePeeringConfigTargetNetwork{
		NetworkUrl: dcl.ValueOrEmptyString(o.NetworkUrl),
	}
	return p
}

// ManagedZoneToProto converts a ManagedZone resource to its proto representation.
func ManagedZoneToProto(resource *dns.ManagedZone) *dnspb.DnsManagedZone {
	p := &dnspb.DnsManagedZone{
		Description:             dcl.ValueOrEmptyString(resource.Description),
		DnsName:                 dcl.ValueOrEmptyString(resource.DnsName),
		DnssecConfig:            DnsManagedZoneDnssecConfigToProto(resource.DnssecConfig),
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Visibility:              DnsManagedZoneVisibilityEnumToProto(resource.Visibility),
		PrivateVisibilityConfig: DnsManagedZonePrivateVisibilityConfigToProto(resource.PrivateVisibilityConfig),
		ForwardingConfig:        DnsManagedZoneForwardingConfigToProto(resource.ForwardingConfig),
		ReverseLookup:           dcl.ValueOrEmptyBool(resource.ReverseLookup),
		PeeringConfig:           DnsManagedZonePeeringConfigToProto(resource.PeeringConfig),
		Project:                 dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.NameServers {
		p.NameServers = append(p.NameServers, r)
	}

	return p
}

// ApplyManagedZone handles the gRPC request by passing it to the underlying ManagedZone Apply() method.
func (s *ManagedZoneServer) applyManagedZone(ctx context.Context, c *dns.Client, request *dnspb.ApplyDnsManagedZoneRequest) (*dnspb.DnsManagedZone, error) {
	p := ProtoToManagedZone(request.GetResource())
	res, err := c.ApplyManagedZone(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ManagedZoneToProto(res)
	return r, nil
}

// ApplyManagedZone handles the gRPC request by passing it to the underlying ManagedZone Apply() method.
func (s *ManagedZoneServer) ApplyDnsManagedZone(ctx context.Context, request *dnspb.ApplyDnsManagedZoneRequest) (*dnspb.DnsManagedZone, error) {
	cl, err := createConfigManagedZone(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyManagedZone(ctx, cl, request)
}

// DeleteManagedZone handles the gRPC request by passing it to the underlying ManagedZone Delete() method.
func (s *ManagedZoneServer) DeleteDnsManagedZone(ctx context.Context, request *dnspb.DeleteDnsManagedZoneRequest) (*emptypb.Empty, error) {

	cl, err := createConfigManagedZone(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteManagedZone(ctx, ProtoToManagedZone(request.GetResource()))

}

// ListDnsManagedZone handles the gRPC request by passing it to the underlying ManagedZoneList() method.
func (s *ManagedZoneServer) ListDnsManagedZone(ctx context.Context, request *dnspb.ListDnsManagedZoneRequest) (*dnspb.ListDnsManagedZoneResponse, error) {
	cl, err := createConfigManagedZone(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListManagedZone(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*dnspb.DnsManagedZone
	for _, r := range resources.Items {
		rp := ManagedZoneToProto(r)
		protos = append(protos, rp)
	}
	return &dnspb.ListDnsManagedZoneResponse{Items: protos}, nil
}

func createConfigManagedZone(ctx context.Context, service_account_file string) (*dns.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dns.NewClient(conf), nil
}
