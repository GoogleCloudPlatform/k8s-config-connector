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

// Server implements the gRPC interface for Router.
type RouterServer struct{}

// ProtoToRouterNatsLogConfigFilterEnum converts a RouterNatsLogConfigFilterEnum enum from its proto representation.
func ProtoToComputeBetaRouterNatsLogConfigFilterEnum(e betapb.ComputeBetaRouterNatsLogConfigFilterEnum) *beta.RouterNatsLogConfigFilterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterNatsLogConfigFilterEnum_name[int32(e)]; ok {
		e := beta.RouterNatsLogConfigFilterEnum(n[len("ComputeBetaRouterNatsLogConfigFilterEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNatsSourceSubnetworkIPRangesToNatEnum converts a RouterNatsSourceSubnetworkIPRangesToNatEnum enum from its proto representation.
func ProtoToComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum(e betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum) *beta.RouterNatsSourceSubnetworkIPRangesToNatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum_name[int32(e)]; ok {
		e := beta.RouterNatsSourceSubnetworkIPRangesToNatEnum(n[len("ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNatsNatIPAllocateOptionEnum converts a RouterNatsNatIPAllocateOptionEnum enum from its proto representation.
func ProtoToComputeBetaRouterNatsNatIPAllocateOptionEnum(e betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum) *beta.RouterNatsNatIPAllocateOptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum_name[int32(e)]; ok {
		e := beta.RouterNatsNatIPAllocateOptionEnum(n[len("ComputeBetaRouterNatsNatIPAllocateOptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterInterfacesManagementTypeEnum converts a RouterInterfacesManagementTypeEnum enum from its proto representation.
func ProtoToComputeBetaRouterInterfacesManagementTypeEnum(e betapb.ComputeBetaRouterInterfacesManagementTypeEnum) *beta.RouterInterfacesManagementTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterInterfacesManagementTypeEnum_name[int32(e)]; ok {
		e := beta.RouterInterfacesManagementTypeEnum(n[len("ComputeBetaRouterInterfacesManagementTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterBgpPeersAdvertisedGroupsEnum converts a RouterBgpPeersAdvertisedGroupsEnum enum from its proto representation.
func ProtoToComputeBetaRouterBgpPeersAdvertisedGroupsEnum(e betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum) *beta.RouterBgpPeersAdvertisedGroupsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum_name[int32(e)]; ok {
		e := beta.RouterBgpPeersAdvertisedGroupsEnum(n[len("ComputeBetaRouterBgpPeersAdvertisedGroupsEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterBgpAdvertiseModeEnum converts a RouterBgpAdvertiseModeEnum enum from its proto representation.
func ProtoToComputeBetaRouterBgpAdvertiseModeEnum(e betapb.ComputeBetaRouterBgpAdvertiseModeEnum) *beta.RouterBgpAdvertiseModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaRouterBgpAdvertiseModeEnum_name[int32(e)]; ok {
		e := beta.RouterBgpAdvertiseModeEnum(n[len("ComputeBetaRouterBgpAdvertiseModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToRouterNats converts a RouterNats resource from its proto representation.
func ProtoToComputeBetaRouterNats(p *betapb.ComputeBetaRouterNats) *beta.RouterNats {
	if p == nil {
		return nil
	}
	obj := &beta.RouterNats{
		Name:                          dcl.StringOrNil(p.Name),
		LogConfig:                     ProtoToComputeBetaRouterNatsLogConfig(p.GetLogConfig()),
		SourceSubnetworkIPRangesToNat: ProtoToComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum(p.GetSourceSubnetworkIpRangesToNat()),
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
		obj.NatIPAllocateOption = append(obj.NatIPAllocateOption, *ProtoToComputeBetaRouterNatsNatIPAllocateOptionEnum(r))
	}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputeBetaRouterNatsSubnetworks(r))
	}
	return obj
}

// ProtoToRouterNatsLogConfig converts a RouterNatsLogConfig resource from its proto representation.
func ProtoToComputeBetaRouterNatsLogConfig(p *betapb.ComputeBetaRouterNatsLogConfig) *beta.RouterNatsLogConfig {
	if p == nil {
		return nil
	}
	obj := &beta.RouterNatsLogConfig{
		Enable: dcl.Bool(p.Enable),
		Filter: ProtoToComputeBetaRouterNatsLogConfigFilterEnum(p.GetFilter()),
	}
	return obj
}

// ProtoToRouterNatsSubnetworks converts a RouterNatsSubnetworks resource from its proto representation.
func ProtoToComputeBetaRouterNatsSubnetworks(p *betapb.ComputeBetaRouterNatsSubnetworks) *beta.RouterNatsSubnetworks {
	if p == nil {
		return nil
	}
	obj := &beta.RouterNatsSubnetworks{
		Name:                  dcl.StringOrNil(p.Name),
		SourceIPRangesToNat:   dcl.StringOrNil(p.SourceIpRangesToNat),
		SecondaryIPRangeNames: dcl.StringOrNil(p.SecondaryIpRangeNames),
	}
	return obj
}

// ProtoToRouterInterfaces converts a RouterInterfaces resource from its proto representation.
func ProtoToComputeBetaRouterInterfaces(p *betapb.ComputeBetaRouterInterfaces) *beta.RouterInterfaces {
	if p == nil {
		return nil
	}
	obj := &beta.RouterInterfaces{
		Name:            dcl.StringOrNil(p.Name),
		LinkedVpnTunnel: dcl.StringOrNil(p.LinkedVpnTunnel),
		IPRange:         dcl.StringOrNil(p.IpRange),
		ManagementType:  ProtoToComputeBetaRouterInterfacesManagementTypeEnum(p.GetManagementType()),
	}
	return obj
}

// ProtoToRouterBgpPeers converts a RouterBgpPeers resource from its proto representation.
func ProtoToComputeBetaRouterBgpPeers(p *betapb.ComputeBetaRouterBgpPeers) *beta.RouterBgpPeers {
	if p == nil {
		return nil
	}
	obj := &beta.RouterBgpPeers{
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
		obj.AdvertisedGroups = append(obj.AdvertisedGroups, *ProtoToComputeBetaRouterBgpPeersAdvertisedGroupsEnum(r))
	}
	for _, r := range p.GetAdvertisedIpRanges() {
		obj.AdvertisedIPRanges = append(obj.AdvertisedIPRanges, *ProtoToComputeBetaRouterBgpPeersAdvertisedIPRanges(r))
	}
	return obj
}

// ProtoToRouterBgpPeersAdvertisedIPRanges converts a RouterBgpPeersAdvertisedIPRanges resource from its proto representation.
func ProtoToComputeBetaRouterBgpPeersAdvertisedIPRanges(p *betapb.ComputeBetaRouterBgpPeersAdvertisedIPRanges) *beta.RouterBgpPeersAdvertisedIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.RouterBgpPeersAdvertisedIPRanges{
		Range:       dcl.StringOrNil(p.Range),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToRouterBgp converts a RouterBgp resource from its proto representation.
func ProtoToComputeBetaRouterBgp(p *betapb.ComputeBetaRouterBgp) *beta.RouterBgp {
	if p == nil {
		return nil
	}
	obj := &beta.RouterBgp{
		Asn:           dcl.Int64OrNil(p.Asn),
		AdvertiseMode: ProtoToComputeBetaRouterBgpAdvertiseModeEnum(p.GetAdvertiseMode()),
	}
	for _, r := range p.GetAdvertisedGroups() {
		obj.AdvertisedGroups = append(obj.AdvertisedGroups, r)
	}
	for _, r := range p.GetAdvertisedIpRanges() {
		obj.AdvertisedIPRanges = append(obj.AdvertisedIPRanges, *ProtoToComputeBetaRouterBgpAdvertisedIPRanges(r))
	}
	return obj
}

// ProtoToRouterBgpAdvertisedIPRanges converts a RouterBgpAdvertisedIPRanges resource from its proto representation.
func ProtoToComputeBetaRouterBgpAdvertisedIPRanges(p *betapb.ComputeBetaRouterBgpAdvertisedIPRanges) *beta.RouterBgpAdvertisedIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.RouterBgpAdvertisedIPRanges{
		Range:       dcl.StringOrNil(p.Range),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToRouter converts a Router resource from its proto representation.
func ProtoToRouter(p *betapb.ComputeBetaRouter) *beta.Router {
	obj := &beta.Router{
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:              dcl.StringOrNil(p.Name),
		Network:           dcl.StringOrNil(p.Network),
		Description:       dcl.StringOrNil(p.Description),
		Bgp:               ProtoToComputeBetaRouterBgp(p.GetBgp()),
		Region:            dcl.StringOrNil(p.Region),
		Project:           dcl.StringOrNil(p.Project),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
	}
	for _, r := range p.GetNats() {
		obj.Nats = append(obj.Nats, *ProtoToComputeBetaRouterNats(r))
	}
	for _, r := range p.GetInterfaces() {
		obj.Interfaces = append(obj.Interfaces, *ProtoToComputeBetaRouterInterfaces(r))
	}
	for _, r := range p.GetBgpPeers() {
		obj.BgpPeers = append(obj.BgpPeers, *ProtoToComputeBetaRouterBgpPeers(r))
	}
	return obj
}

// RouterNatsLogConfigFilterEnumToProto converts a RouterNatsLogConfigFilterEnum enum to its proto representation.
func ComputeBetaRouterNatsLogConfigFilterEnumToProto(e *beta.RouterNatsLogConfigFilterEnum) betapb.ComputeBetaRouterNatsLogConfigFilterEnum {
	if e == nil {
		return betapb.ComputeBetaRouterNatsLogConfigFilterEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterNatsLogConfigFilterEnum_value["RouterNatsLogConfigFilterEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterNatsLogConfigFilterEnum(v)
	}
	return betapb.ComputeBetaRouterNatsLogConfigFilterEnum(0)
}

// RouterNatsSourceSubnetworkIPRangesToNatEnumToProto converts a RouterNatsSourceSubnetworkIPRangesToNatEnum enum to its proto representation.
func ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnumToProto(e *beta.RouterNatsSourceSubnetworkIPRangesToNatEnum) betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum {
	if e == nil {
		return betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum_value["RouterNatsSourceSubnetworkIPRangesToNatEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum(v)
	}
	return betapb.ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnum(0)
}

// RouterNatsNatIPAllocateOptionEnumToProto converts a RouterNatsNatIPAllocateOptionEnum enum to its proto representation.
func ComputeBetaRouterNatsNatIPAllocateOptionEnumToProto(e *beta.RouterNatsNatIPAllocateOptionEnum) betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum {
	if e == nil {
		return betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum_value["RouterNatsNatIPAllocateOptionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum(v)
	}
	return betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum(0)
}

// RouterInterfacesManagementTypeEnumToProto converts a RouterInterfacesManagementTypeEnum enum to its proto representation.
func ComputeBetaRouterInterfacesManagementTypeEnumToProto(e *beta.RouterInterfacesManagementTypeEnum) betapb.ComputeBetaRouterInterfacesManagementTypeEnum {
	if e == nil {
		return betapb.ComputeBetaRouterInterfacesManagementTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterInterfacesManagementTypeEnum_value["RouterInterfacesManagementTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterInterfacesManagementTypeEnum(v)
	}
	return betapb.ComputeBetaRouterInterfacesManagementTypeEnum(0)
}

// RouterBgpPeersAdvertisedGroupsEnumToProto converts a RouterBgpPeersAdvertisedGroupsEnum enum to its proto representation.
func ComputeBetaRouterBgpPeersAdvertisedGroupsEnumToProto(e *beta.RouterBgpPeersAdvertisedGroupsEnum) betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum {
	if e == nil {
		return betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum_value["RouterBgpPeersAdvertisedGroupsEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum(v)
	}
	return betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum(0)
}

// RouterBgpAdvertiseModeEnumToProto converts a RouterBgpAdvertiseModeEnum enum to its proto representation.
func ComputeBetaRouterBgpAdvertiseModeEnumToProto(e *beta.RouterBgpAdvertiseModeEnum) betapb.ComputeBetaRouterBgpAdvertiseModeEnum {
	if e == nil {
		return betapb.ComputeBetaRouterBgpAdvertiseModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaRouterBgpAdvertiseModeEnum_value["RouterBgpAdvertiseModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaRouterBgpAdvertiseModeEnum(v)
	}
	return betapb.ComputeBetaRouterBgpAdvertiseModeEnum(0)
}

// RouterNatsToProto converts a RouterNats resource to its proto representation.
func ComputeBetaRouterNatsToProto(o *beta.RouterNats) *betapb.ComputeBetaRouterNats {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterNats{
		Name:                          dcl.ValueOrEmptyString(o.Name),
		LogConfig:                     ComputeBetaRouterNatsLogConfigToProto(o.LogConfig),
		SourceSubnetworkIpRangesToNat: ComputeBetaRouterNatsSourceSubnetworkIPRangesToNatEnumToProto(o.SourceSubnetworkIPRangesToNat),
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
		p.NatIpAllocateOption = append(p.NatIpAllocateOption, betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum(betapb.ComputeBetaRouterNatsNatIPAllocateOptionEnum_value[string(r)]))
	}
	for _, r := range o.Subnetworks {
		p.Subnetworks = append(p.Subnetworks, ComputeBetaRouterNatsSubnetworksToProto(&r))
	}
	return p
}

// RouterNatsLogConfigToProto converts a RouterNatsLogConfig resource to its proto representation.
func ComputeBetaRouterNatsLogConfigToProto(o *beta.RouterNatsLogConfig) *betapb.ComputeBetaRouterNatsLogConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterNatsLogConfig{
		Enable: dcl.ValueOrEmptyBool(o.Enable),
		Filter: ComputeBetaRouterNatsLogConfigFilterEnumToProto(o.Filter),
	}
	return p
}

// RouterNatsSubnetworksToProto converts a RouterNatsSubnetworks resource to its proto representation.
func ComputeBetaRouterNatsSubnetworksToProto(o *beta.RouterNatsSubnetworks) *betapb.ComputeBetaRouterNatsSubnetworks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterNatsSubnetworks{
		Name:                  dcl.ValueOrEmptyString(o.Name),
		SourceIpRangesToNat:   dcl.ValueOrEmptyString(o.SourceIPRangesToNat),
		SecondaryIpRangeNames: dcl.ValueOrEmptyString(o.SecondaryIPRangeNames),
	}
	return p
}

// RouterInterfacesToProto converts a RouterInterfaces resource to its proto representation.
func ComputeBetaRouterInterfacesToProto(o *beta.RouterInterfaces) *betapb.ComputeBetaRouterInterfaces {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterInterfaces{
		Name:            dcl.ValueOrEmptyString(o.Name),
		LinkedVpnTunnel: dcl.ValueOrEmptyString(o.LinkedVpnTunnel),
		IpRange:         dcl.ValueOrEmptyString(o.IPRange),
		ManagementType:  ComputeBetaRouterInterfacesManagementTypeEnumToProto(o.ManagementType),
	}
	return p
}

// RouterBgpPeersToProto converts a RouterBgpPeers resource to its proto representation.
func ComputeBetaRouterBgpPeersToProto(o *beta.RouterBgpPeers) *betapb.ComputeBetaRouterBgpPeers {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterBgpPeers{
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
		p.AdvertisedGroups = append(p.AdvertisedGroups, betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum(betapb.ComputeBetaRouterBgpPeersAdvertisedGroupsEnum_value[string(r)]))
	}
	for _, r := range o.AdvertisedIPRanges {
		p.AdvertisedIpRanges = append(p.AdvertisedIpRanges, ComputeBetaRouterBgpPeersAdvertisedIPRangesToProto(&r))
	}
	return p
}

// RouterBgpPeersAdvertisedIPRangesToProto converts a RouterBgpPeersAdvertisedIPRanges resource to its proto representation.
func ComputeBetaRouterBgpPeersAdvertisedIPRangesToProto(o *beta.RouterBgpPeersAdvertisedIPRanges) *betapb.ComputeBetaRouterBgpPeersAdvertisedIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterBgpPeersAdvertisedIPRanges{
		Range:       dcl.ValueOrEmptyString(o.Range),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// RouterBgpToProto converts a RouterBgp resource to its proto representation.
func ComputeBetaRouterBgpToProto(o *beta.RouterBgp) *betapb.ComputeBetaRouterBgp {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterBgp{
		Asn:           dcl.ValueOrEmptyInt64(o.Asn),
		AdvertiseMode: ComputeBetaRouterBgpAdvertiseModeEnumToProto(o.AdvertiseMode),
	}
	for _, r := range o.AdvertisedGroups {
		p.AdvertisedGroups = append(p.AdvertisedGroups, r)
	}
	for _, r := range o.AdvertisedIPRanges {
		p.AdvertisedIpRanges = append(p.AdvertisedIpRanges, ComputeBetaRouterBgpAdvertisedIPRangesToProto(&r))
	}
	return p
}

// RouterBgpAdvertisedIPRangesToProto converts a RouterBgpAdvertisedIPRanges resource to its proto representation.
func ComputeBetaRouterBgpAdvertisedIPRangesToProto(o *beta.RouterBgpAdvertisedIPRanges) *betapb.ComputeBetaRouterBgpAdvertisedIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaRouterBgpAdvertisedIPRanges{
		Range:       dcl.ValueOrEmptyString(o.Range),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// RouterToProto converts a Router resource to its proto representation.
func RouterToProto(resource *beta.Router) *betapb.ComputeBetaRouter {
	p := &betapb.ComputeBetaRouter{
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Network:           dcl.ValueOrEmptyString(resource.Network),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Bgp:               ComputeBetaRouterBgpToProto(resource.Bgp),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
	}
	for _, r := range resource.Nats {
		p.Nats = append(p.Nats, ComputeBetaRouterNatsToProto(&r))
	}
	for _, r := range resource.Interfaces {
		p.Interfaces = append(p.Interfaces, ComputeBetaRouterInterfacesToProto(&r))
	}
	for _, r := range resource.BgpPeers {
		p.BgpPeers = append(p.BgpPeers, ComputeBetaRouterBgpPeersToProto(&r))
	}

	return p
}

// ApplyRouter handles the gRPC request by passing it to the underlying Router Apply() method.
func (s *RouterServer) applyRouter(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaRouterRequest) (*betapb.ComputeBetaRouter, error) {
	p := ProtoToRouter(request.GetResource())
	res, err := c.ApplyRouter(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RouterToProto(res)
	return r, nil
}

// ApplyRouter handles the gRPC request by passing it to the underlying Router Apply() method.
func (s *RouterServer) ApplyComputeBetaRouter(ctx context.Context, request *betapb.ApplyComputeBetaRouterRequest) (*betapb.ComputeBetaRouter, error) {
	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRouter(ctx, cl, request)
}

// DeleteRouter handles the gRPC request by passing it to the underlying Router Delete() method.
func (s *RouterServer) DeleteComputeBetaRouter(ctx context.Context, request *betapb.DeleteComputeBetaRouterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRouter(ctx, ProtoToRouter(request.GetResource()))

}

// ListComputeBetaRouter handles the gRPC request by passing it to the underlying RouterList() method.
func (s *RouterServer) ListComputeBetaRouter(ctx context.Context, request *betapb.ListComputeBetaRouterRequest) (*betapb.ListComputeBetaRouterResponse, error) {
	cl, err := createConfigRouter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRouter(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaRouter
	for _, r := range resources.Items {
		rp := RouterToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaRouterResponse{Items: protos}, nil
}

func createConfigRouter(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
