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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Router.
type RouterServer struct{}

// ProtoToRouterNatsLogConfigFilterEnum converts a RouterNatsLogConfigFilterEnum enum from its proto representation.
func ProtoToComputeRouterNatsLogConfigFilterEnum(e computepb.ComputeRouterNatsLogConfigFilterEnum) *compute.RouterNatsLogConfigFilterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterNatsLogConfigFilterEnum_name[int32(e)]; ok {
		e := compute.RouterNatsLogConfigFilterEnum(n[len("ComputeRouterNatsLogConfigFilterEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNatsSourceSubnetworkIPRangesToNatEnum converts a RouterNatsSourceSubnetworkIPRangesToNatEnum enum from its proto representation.
func ProtoToComputeRouterNatsSourceSubnetworkIPRangesToNatEnum(e computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum) *compute.RouterNatsSourceSubnetworkIPRangesToNatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum_name[int32(e)]; ok {
		e := compute.RouterNatsSourceSubnetworkIPRangesToNatEnum(n[len("ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNatsNatIPAllocateOptionEnum converts a RouterNatsNatIPAllocateOptionEnum enum from its proto representation.
func ProtoToComputeRouterNatsNatIPAllocateOptionEnum(e computepb.ComputeRouterNatsNatIPAllocateOptionEnum) *compute.RouterNatsNatIPAllocateOptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterNatsNatIPAllocateOptionEnum_name[int32(e)]; ok {
		e := compute.RouterNatsNatIPAllocateOptionEnum(n[len("ComputeRouterNatsNatIPAllocateOptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterInterfacesManagementTypeEnum converts a RouterInterfacesManagementTypeEnum enum from its proto representation.
func ProtoToComputeRouterInterfacesManagementTypeEnum(e computepb.ComputeRouterInterfacesManagementTypeEnum) *compute.RouterInterfacesManagementTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterInterfacesManagementTypeEnum_name[int32(e)]; ok {
		e := compute.RouterInterfacesManagementTypeEnum(n[len("ComputeRouterInterfacesManagementTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterBgpPeersAdvertisedGroupsEnum converts a RouterBgpPeersAdvertisedGroupsEnum enum from its proto representation.
func ProtoToComputeRouterBgpPeersAdvertisedGroupsEnum(e computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum) *compute.RouterBgpPeersAdvertisedGroupsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum_name[int32(e)]; ok {
		e := compute.RouterBgpPeersAdvertisedGroupsEnum(n[len("ComputeRouterBgpPeersAdvertisedGroupsEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterBgpAdvertiseModeEnum converts a RouterBgpAdvertiseModeEnum enum from its proto representation.
func ProtoToComputeRouterBgpAdvertiseModeEnum(e computepb.ComputeRouterBgpAdvertiseModeEnum) *compute.RouterBgpAdvertiseModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeRouterBgpAdvertiseModeEnum_name[int32(e)]; ok {
		e := compute.RouterBgpAdvertiseModeEnum(n[len("ComputeRouterBgpAdvertiseModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNats converts a RouterNats resource from its proto representation.
func ProtoToComputeRouterNats(p *computepb.ComputeRouterNats) *compute.RouterNats {
	if p == nil {
		return nil
	}
	obj := &compute.RouterNats{
		Name:                          dcl.StringOrNil(p.Name),
		LogConfig:                     ProtoToComputeRouterNatsLogConfig(p.GetLogConfig()),
		SourceSubnetworkIPRangesToNat: ProtoToComputeRouterNatsSourceSubnetworkIPRangesToNatEnum(p.GetSourceSubnetworkIpRangesToNat()),
		MinPortsPerVm:                 dcl.Int64OrNil(p.MinPortsPerVm),
		UdpIdleTimeoutSec:             dcl.Int64OrNil(p.UdpIdleTimeoutSec),
		IcmpIdleTimeoutSec:            dcl.Int64OrNil(p.IcmpIdleTimeoutSec),
		TcpEstablishedIdleTimeoutSec:  dcl.Int64OrNil(p.TcpEstablishedIdleTimeoutSec),
		TcpTransitoryIdleTimeoutSec:   dcl.Int64OrNil(p.TcpTransitoryIdleTimeoutSec),
	}
	for _, r := range p.GetNatIps() {
		obj.NatIps = append(obj.NatIps, r)
	}
	for _, r := range p.GetDrainNatIps() {
		obj.DrainNatIps = append(obj.DrainNatIps, r)
	}
	for _, r := range p.GetNatIpAllocateOption() {
		obj.NatIPAllocateOption = append(obj.NatIPAllocateOption, *ProtoToComputeRouterNatsNatIPAllocateOptionEnum(r))
	}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputeRouterNatsSubnetworks(r))
	}
	return obj
}

// ProtoToRouterNatsLogConfig converts a RouterNatsLogConfig resource from its proto representation.
func ProtoToComputeRouterNatsLogConfig(p *computepb.ComputeRouterNatsLogConfig) *compute.RouterNatsLogConfig {
	if p == nil {
		return nil
	}
	obj := &compute.RouterNatsLogConfig{
		Enable: dcl.Bool(p.Enable),
		Filter: ProtoToComputeRouterNatsLogConfigFilterEnum(p.GetFilter()),
	}
	return obj
}

// ProtoToRouterNatsSubnetworks converts a RouterNatsSubnetworks resource from its proto representation.
func ProtoToComputeRouterNatsSubnetworks(p *computepb.ComputeRouterNatsSubnetworks) *compute.RouterNatsSubnetworks {
	if p == nil {
		return nil
	}
	obj := &compute.RouterNatsSubnetworks{
		Name:                  dcl.StringOrNil(p.Name),
		SourceIPRangesToNat:   dcl.StringOrNil(p.SourceIpRangesToNat),
		SecondaryIPRangeNames: dcl.StringOrNil(p.SecondaryIpRangeNames),
	}
	return obj
}

// ProtoToRouterInterfaces converts a RouterInterfaces resource from its proto representation.
func ProtoToComputeRouterInterfaces(p *computepb.ComputeRouterInterfaces) *compute.RouterInterfaces {
	if p == nil {
		return nil
	}
	obj := &compute.RouterInterfaces{
		Name:            dcl.StringOrNil(p.Name),
		LinkedVpnTunnel: dcl.StringOrNil(p.LinkedVpnTunnel),
		IPRange:         dcl.StringOrNil(p.IpRange),
		ManagementType:  ProtoToComputeRouterInterfacesManagementTypeEnum(p.GetManagementType()),
	}
	return obj
}

// ProtoToRouterBgpPeers converts a RouterBgpPeers resource from its proto representation.
func ProtoToComputeRouterBgpPeers(p *computepb.ComputeRouterBgpPeers) *compute.RouterBgpPeers {
	if p == nil {
		return nil
	}
	obj := &compute.RouterBgpPeers{
		Name:                    dcl.StringOrNil(p.Name),
		InterfaceName:           dcl.StringOrNil(p.InterfaceName),
		IPAddress:               dcl.StringOrNil(p.IpAddress),
		PeerIPAddress:           dcl.StringOrNil(p.PeerIpAddress),
		PeerAsn:                 dcl.Int64OrNil(p.PeerAsn),
		AdvertisedRoutePriority: dcl.Int64OrNil(p.AdvertisedRoutePriority),
		AdvertiseMode:           dcl.StringOrNil(p.AdvertiseMode),
		ManagementType:          dcl.StringOrNil(p.ManagementType),
	}
	for _, r := range p.GetAdvertisedGroups() {
		obj.AdvertisedGroups = append(obj.AdvertisedGroups, *ProtoToComputeRouterBgpPeersAdvertisedGroupsEnum(r))
	}
	for _, r := range p.GetAdvertisedIpRanges() {
		obj.AdvertisedIPRanges = append(obj.AdvertisedIPRanges, *ProtoToComputeRouterBgpPeersAdvertisedIPRanges(r))
	}
	return obj
}

// ProtoToRouterBgpPeersAdvertisedIPRanges converts a RouterBgpPeersAdvertisedIPRanges resource from its proto representation.
func ProtoToComputeRouterBgpPeersAdvertisedIPRanges(p *computepb.ComputeRouterBgpPeersAdvertisedIPRanges) *compute.RouterBgpPeersAdvertisedIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.RouterBgpPeersAdvertisedIPRanges{
		Range:       dcl.StringOrNil(p.Range),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToRouterBgp converts a RouterBgp resource from its proto representation.
func ProtoToComputeRouterBgp(p *computepb.ComputeRouterBgp) *compute.RouterBgp {
	if p == nil {
		return nil
	}
	obj := &compute.RouterBgp{
		Asn:           dcl.Int64OrNil(p.Asn),
		AdvertiseMode: ProtoToComputeRouterBgpAdvertiseModeEnum(p.GetAdvertiseMode()),
	}
	for _, r := range p.GetAdvertisedGroups() {
		obj.AdvertisedGroups = append(obj.AdvertisedGroups, r)
	}
	for _, r := range p.GetAdvertisedIpRanges() {
		obj.AdvertisedIPRanges = append(obj.AdvertisedIPRanges, *ProtoToComputeRouterBgpAdvertisedIPRanges(r))
	}
	return obj
}

// ProtoToRouterBgpAdvertisedIPRanges converts a RouterBgpAdvertisedIPRanges resource from its proto representation.
func ProtoToComputeRouterBgpAdvertisedIPRanges(p *computepb.ComputeRouterBgpAdvertisedIPRanges) *compute.RouterBgpAdvertisedIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.RouterBgpAdvertisedIPRanges{
		Range:       dcl.StringOrNil(p.Range),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToRouter converts a Router resource from its proto representation.
func ProtoToRouter(p *computepb.ComputeRouter) *compute.Router {
	obj := &compute.Router{
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:              dcl.StringOrNil(p.Name),
		Network:           dcl.StringOrNil(p.Network),
		Description:       dcl.StringOrNil(p.Description),
		Bgp:               ProtoToComputeRouterBgp(p.GetBgp()),
		Region:            dcl.StringOrNil(p.Region),
		Project:           dcl.StringOrNil(p.Project),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
	}
	for _, r := range p.GetNats() {
		obj.Nats = append(obj.Nats, *ProtoToComputeRouterNats(r))
	}
	for _, r := range p.GetInterfaces() {
		obj.Interfaces = append(obj.Interfaces, *ProtoToComputeRouterInterfaces(r))
	}
	for _, r := range p.GetBgpPeers() {
		obj.BgpPeers = append(obj.BgpPeers, *ProtoToComputeRouterBgpPeers(r))
	}
	return obj
}

// RouterNatsLogConfigFilterEnumToProto converts a RouterNatsLogConfigFilterEnum enum to its proto representation.
func ComputeRouterNatsLogConfigFilterEnumToProto(e *compute.RouterNatsLogConfigFilterEnum) computepb.ComputeRouterNatsLogConfigFilterEnum {
	if e == nil {
		return computepb.ComputeRouterNatsLogConfigFilterEnum(0)
	}
	if v, ok := computepb.ComputeRouterNatsLogConfigFilterEnum_value["RouterNatsLogConfigFilterEnum"+string(*e)]; ok {
		return computepb.ComputeRouterNatsLogConfigFilterEnum(v)
	}
	return computepb.ComputeRouterNatsLogConfigFilterEnum(0)
}

// RouterNatsSourceSubnetworkIPRangesToNatEnumToProto converts a RouterNatsSourceSubnetworkIPRangesToNatEnum enum to its proto representation.
func ComputeRouterNatsSourceSubnetworkIPRangesToNatEnumToProto(e *compute.RouterNatsSourceSubnetworkIPRangesToNatEnum) computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum {
	if e == nil {
		return computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum(0)
	}
	if v, ok := computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum_value["RouterNatsSourceSubnetworkIPRangesToNatEnum"+string(*e)]; ok {
		return computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum(v)
	}
	return computepb.ComputeRouterNatsSourceSubnetworkIPRangesToNatEnum(0)
}

// RouterNatsNatIPAllocateOptionEnumToProto converts a RouterNatsNatIPAllocateOptionEnum enum to its proto representation.
func ComputeRouterNatsNatIPAllocateOptionEnumToProto(e *compute.RouterNatsNatIPAllocateOptionEnum) computepb.ComputeRouterNatsNatIPAllocateOptionEnum {
	if e == nil {
		return computepb.ComputeRouterNatsNatIPAllocateOptionEnum(0)
	}
	if v, ok := computepb.ComputeRouterNatsNatIPAllocateOptionEnum_value["RouterNatsNatIPAllocateOptionEnum"+string(*e)]; ok {
		return computepb.ComputeRouterNatsNatIPAllocateOptionEnum(v)
	}
	return computepb.ComputeRouterNatsNatIPAllocateOptionEnum(0)
}

// RouterInterfacesManagementTypeEnumToProto converts a RouterInterfacesManagementTypeEnum enum to its proto representation.
func ComputeRouterInterfacesManagementTypeEnumToProto(e *compute.RouterInterfacesManagementTypeEnum) computepb.ComputeRouterInterfacesManagementTypeEnum {
	if e == nil {
		return computepb.ComputeRouterInterfacesManagementTypeEnum(0)
	}
	if v, ok := computepb.ComputeRouterInterfacesManagementTypeEnum_value["RouterInterfacesManagementTypeEnum"+string(*e)]; ok {
		return computepb.ComputeRouterInterfacesManagementTypeEnum(v)
	}
	return computepb.ComputeRouterInterfacesManagementTypeEnum(0)
}

// RouterBgpPeersAdvertisedGroupsEnumToProto converts a RouterBgpPeersAdvertisedGroupsEnum enum to its proto representation.
func ComputeRouterBgpPeersAdvertisedGroupsEnumToProto(e *compute.RouterBgpPeersAdvertisedGroupsEnum) computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum {
	if e == nil {
		return computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum(0)
	}
	if v, ok := computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum_value["RouterBgpPeersAdvertisedGroupsEnum"+string(*e)]; ok {
		return computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum(v)
	}
	return computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum(0)
}

// RouterBgpAdvertiseModeEnumToProto converts a RouterBgpAdvertiseModeEnum enum to its proto representation.
func ComputeRouterBgpAdvertiseModeEnumToProto(e *compute.RouterBgpAdvertiseModeEnum) computepb.ComputeRouterBgpAdvertiseModeEnum {
	if e == nil {
		return computepb.ComputeRouterBgpAdvertiseModeEnum(0)
	}
	if v, ok := computepb.ComputeRouterBgpAdvertiseModeEnum_value["RouterBgpAdvertiseModeEnum"+string(*e)]; ok {
		return computepb.ComputeRouterBgpAdvertiseModeEnum(v)
	}
	return computepb.ComputeRouterBgpAdvertiseModeEnum(0)
}

// RouterNatsToProto converts a RouterNats resource to its proto representation.
func ComputeRouterNatsToProto(o *compute.RouterNats) *computepb.ComputeRouterNats {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterNats{
		Name:                          dcl.ValueOrEmptyString(o.Name),
		LogConfig:                     ComputeRouterNatsLogConfigToProto(o.LogConfig),
		SourceSubnetworkIpRangesToNat: ComputeRouterNatsSourceSubnetworkIPRangesToNatEnumToProto(o.SourceSubnetworkIPRangesToNat),
		MinPortsPerVm:                 dcl.ValueOrEmptyInt64(o.MinPortsPerVm),
		UdpIdleTimeoutSec:             dcl.ValueOrEmptyInt64(o.UdpIdleTimeoutSec),
		IcmpIdleTimeoutSec:            dcl.ValueOrEmptyInt64(o.IcmpIdleTimeoutSec),
		TcpEstablishedIdleTimeoutSec:  dcl.ValueOrEmptyInt64(o.TcpEstablishedIdleTimeoutSec),
		TcpTransitoryIdleTimeoutSec:   dcl.ValueOrEmptyInt64(o.TcpTransitoryIdleTimeoutSec),
	}
	for _, r := range o.NatIps {
		p.NatIps = append(p.NatIps, r)
	}
	for _, r := range o.DrainNatIps {
		p.DrainNatIps = append(p.DrainNatIps, r)
	}
	for _, r := range o.NatIPAllocateOption {
		p.NatIpAllocateOption = append(p.NatIpAllocateOption, computepb.ComputeRouterNatsNatIPAllocateOptionEnum(computepb.ComputeRouterNatsNatIPAllocateOptionEnum_value[string(r)]))
	}
	for _, r := range o.Subnetworks {
		p.Subnetworks = append(p.Subnetworks, ComputeRouterNatsSubnetworksToProto(&r))
	}
	return p
}

// RouterNatsLogConfigToProto converts a RouterNatsLogConfig resource to its proto representation.
func ComputeRouterNatsLogConfigToProto(o *compute.RouterNatsLogConfig) *computepb.ComputeRouterNatsLogConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterNatsLogConfig{
		Enable: dcl.ValueOrEmptyBool(o.Enable),
		Filter: ComputeRouterNatsLogConfigFilterEnumToProto(o.Filter),
	}
	return p
}

// RouterNatsSubnetworksToProto converts a RouterNatsSubnetworks resource to its proto representation.
func ComputeRouterNatsSubnetworksToProto(o *compute.RouterNatsSubnetworks) *computepb.ComputeRouterNatsSubnetworks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterNatsSubnetworks{
		Name:                  dcl.ValueOrEmptyString(o.Name),
		SourceIpRangesToNat:   dcl.ValueOrEmptyString(o.SourceIPRangesToNat),
		SecondaryIpRangeNames: dcl.ValueOrEmptyString(o.SecondaryIPRangeNames),
	}
	return p
}

// RouterInterfacesToProto converts a RouterInterfaces resource to its proto representation.
func ComputeRouterInterfacesToProto(o *compute.RouterInterfaces) *computepb.ComputeRouterInterfaces {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterInterfaces{
		Name:            dcl.ValueOrEmptyString(o.Name),
		LinkedVpnTunnel: dcl.ValueOrEmptyString(o.LinkedVpnTunnel),
		IpRange:         dcl.ValueOrEmptyString(o.IPRange),
		ManagementType:  ComputeRouterInterfacesManagementTypeEnumToProto(o.ManagementType),
	}
	return p
}

// RouterBgpPeersToProto converts a RouterBgpPeers resource to its proto representation.
func ComputeRouterBgpPeersToProto(o *compute.RouterBgpPeers) *computepb.ComputeRouterBgpPeers {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterBgpPeers{
		Name:                    dcl.ValueOrEmptyString(o.Name),
		InterfaceName:           dcl.ValueOrEmptyString(o.InterfaceName),
		IpAddress:               dcl.ValueOrEmptyString(o.IPAddress),
		PeerIpAddress:           dcl.ValueOrEmptyString(o.PeerIPAddress),
		PeerAsn:                 dcl.ValueOrEmptyInt64(o.PeerAsn),
		AdvertisedRoutePriority: dcl.ValueOrEmptyInt64(o.AdvertisedRoutePriority),
		AdvertiseMode:           dcl.ValueOrEmptyString(o.AdvertiseMode),
		ManagementType:          dcl.ValueOrEmptyString(o.ManagementType),
	}
	for _, r := range o.AdvertisedGroups {
		p.AdvertisedGroups = append(p.AdvertisedGroups, computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum(computepb.ComputeRouterBgpPeersAdvertisedGroupsEnum_value[string(r)]))
	}
	for _, r := range o.AdvertisedIPRanges {
		p.AdvertisedIpRanges = append(p.AdvertisedIpRanges, ComputeRouterBgpPeersAdvertisedIPRangesToProto(&r))
	}
	return p
}

// RouterBgpPeersAdvertisedIPRangesToProto converts a RouterBgpPeersAdvertisedIPRanges resource to its proto representation.
func ComputeRouterBgpPeersAdvertisedIPRangesToProto(o *compute.RouterBgpPeersAdvertisedIPRanges) *computepb.ComputeRouterBgpPeersAdvertisedIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterBgpPeersAdvertisedIPRanges{
		Range:       dcl.ValueOrEmptyString(o.Range),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// RouterBgpToProto converts a RouterBgp resource to its proto representation.
func ComputeRouterBgpToProto(o *compute.RouterBgp) *computepb.ComputeRouterBgp {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterBgp{
		Asn:           dcl.ValueOrEmptyInt64(o.Asn),
		AdvertiseMode: ComputeRouterBgpAdvertiseModeEnumToProto(o.AdvertiseMode),
	}
	for _, r := range o.AdvertisedGroups {
		p.AdvertisedGroups = append(p.AdvertisedGroups, r)
	}
	for _, r := range o.AdvertisedIPRanges {
		p.AdvertisedIpRanges = append(p.AdvertisedIpRanges, ComputeRouterBgpAdvertisedIPRangesToProto(&r))
	}
	return p
}

// RouterBgpAdvertisedIPRangesToProto converts a RouterBgpAdvertisedIPRanges resource to its proto representation.
func ComputeRouterBgpAdvertisedIPRangesToProto(o *compute.RouterBgpAdvertisedIPRanges) *computepb.ComputeRouterBgpAdvertisedIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeRouterBgpAdvertisedIPRanges{
		Range:       dcl.ValueOrEmptyString(o.Range),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// RouterToProto converts a Router resource to its proto representation.
func RouterToProto(resource *compute.Router) *computepb.ComputeRouter {
	p := &computepb.ComputeRouter{
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Network:           dcl.ValueOrEmptyString(resource.Network),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Bgp:               ComputeRouterBgpToProto(resource.Bgp),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
	}
	for _, r := range resource.Nats {
		p.Nats = append(p.Nats, ComputeRouterNatsToProto(&r))
	}
	for _, r := range resource.Interfaces {
		p.Interfaces = append(p.Interfaces, ComputeRouterInterfacesToProto(&r))
	}
	for _, r := range resource.BgpPeers {
		p.BgpPeers = append(p.BgpPeers, ComputeRouterBgpPeersToProto(&r))
	}

	return p
}

// ApplyRouter handles the gRPC request by passing it to the underlying Router Apply() method.
func (s *RouterServer) applyRouter(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeRouterRequest) (*computepb.ComputeRouter, error) {
	p := ProtoToRouter(request.GetResource())
	res, err := c.ApplyRouter(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouterToProto(res)
	return r, nil
}

// ApplyRouter handles the gRPC request by passing it to the underlying Router Apply() method.
func (s *RouterServer) ApplyComputeRouter(ctx context.Context, request *computepb.ApplyComputeRouterRequest) (*computepb.ComputeRouter, error) {
	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRouter(ctx, cl, request)
}

// DeleteRouter handles the gRPC request by passing it to the underlying Router Delete() method.
func (s *RouterServer) DeleteComputeRouter(ctx context.Context, request *computepb.DeleteComputeRouterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRouter(ctx, ProtoToRouter(request.GetResource()))

}

// ListComputeRouter handles the gRPC request by passing it to the underlying RouterList() method.
func (s *RouterServer) ListComputeRouter(ctx context.Context, request *computepb.ListComputeRouterRequest) (*computepb.ListComputeRouterResponse, error) {
	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRouter(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeRouter
	for _, r := range resources.Items {
		rp := RouterToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeRouterResponse{Items: protos}, nil
}

func createConfigRouter(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
