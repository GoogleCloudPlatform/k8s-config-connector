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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/beta/networkconnectivity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/beta"
)

// SpokeServer implements the gRPC interface for Spoke.
type SpokeServer struct{}

// ProtoToSpokeStateEnum converts a SpokeStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeStateEnum(e betapb.NetworkconnectivityBetaSpokeStateEnum) *beta.SpokeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkconnectivityBetaSpokeStateEnum_name[int32(e)]; ok {
		e := beta.SpokeStateEnum(n[len("NetworkconnectivityBetaSpokeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToSpokeLinkedVpnTunnels converts a SpokeLinkedVpnTunnels object from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeLinkedVpnTunnels(p *betapb.NetworkconnectivityBetaSpokeLinkedVpnTunnels) *beta.SpokeLinkedVpnTunnels {
	if p == nil {
		return nil
	}
	obj := &beta.SpokeLinkedVpnTunnels{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedInterconnectAttachments converts a SpokeLinkedInterconnectAttachments object from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeLinkedInterconnectAttachments(p *betapb.NetworkconnectivityBetaSpokeLinkedInterconnectAttachments) *beta.SpokeLinkedInterconnectAttachments {
	if p == nil {
		return nil
	}
	obj := &beta.SpokeLinkedInterconnectAttachments{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstances converts a SpokeLinkedRouterApplianceInstances object from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeLinkedRouterApplianceInstances(p *betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstances) *beta.SpokeLinkedRouterApplianceInstances {
	if p == nil {
		return nil
	}
	obj := &beta.SpokeLinkedRouterApplianceInstances{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToNetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances(r))
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstancesInstances converts a SpokeLinkedRouterApplianceInstancesInstances object from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances(p *betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances) *beta.SpokeLinkedRouterApplianceInstancesInstances {
	if p == nil {
		return nil
	}
	obj := &beta.SpokeLinkedRouterApplianceInstancesInstances{
		VirtualMachine: dcl.StringOrNil(p.GetVirtualMachine()),
		IPAddress:      dcl.StringOrNil(p.GetIpAddress()),
	}
	return obj
}

// ProtoToSpokeLinkedVPCNetwork converts a SpokeLinkedVPCNetwork object from its proto representation.
func ProtoToNetworkconnectivityBetaSpokeLinkedVPCNetwork(p *betapb.NetworkconnectivityBetaSpokeLinkedVPCNetwork) *beta.SpokeLinkedVPCNetwork {
	if p == nil {
		return nil
	}
	obj := &beta.SpokeLinkedVPCNetwork{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	for _, r := range p.GetExcludeExportRanges() {
		obj.ExcludeExportRanges = append(obj.ExcludeExportRanges, r)
	}
	return obj
}

// ProtoToSpoke converts a Spoke resource from its proto representation.
func ProtoToSpoke(p *betapb.NetworkconnectivityBetaSpoke) *beta.Spoke {
	obj := &beta.Spoke{
		Name:                           dcl.StringOrNil(p.GetName()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                     dcl.StringOrNil(p.GetUpdateTime()),
		Description:                    dcl.StringOrNil(p.GetDescription()),
		Hub:                            dcl.StringOrNil(p.GetHub()),
		LinkedVpnTunnels:               ProtoToNetworkconnectivityBetaSpokeLinkedVpnTunnels(p.GetLinkedVpnTunnels()),
		LinkedInterconnectAttachments:  ProtoToNetworkconnectivityBetaSpokeLinkedInterconnectAttachments(p.GetLinkedInterconnectAttachments()),
		LinkedRouterApplianceInstances: ProtoToNetworkconnectivityBetaSpokeLinkedRouterApplianceInstances(p.GetLinkedRouterApplianceInstances()),
		LinkedVPCNetwork:               ProtoToNetworkconnectivityBetaSpokeLinkedVPCNetwork(p.GetLinkedVpcNetwork()),
		UniqueId:                       dcl.StringOrNil(p.GetUniqueId()),
		State:                          ProtoToNetworkconnectivityBetaSpokeStateEnum(p.GetState()),
		Project:                        dcl.StringOrNil(p.GetProject()),
		Location:                       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// SpokeStateEnumToProto converts a SpokeStateEnum enum to its proto representation.
func NetworkconnectivityBetaSpokeStateEnumToProto(e *beta.SpokeStateEnum) betapb.NetworkconnectivityBetaSpokeStateEnum {
	if e == nil {
		return betapb.NetworkconnectivityBetaSpokeStateEnum(0)
	}
	if v, ok := betapb.NetworkconnectivityBetaSpokeStateEnum_value["SpokeStateEnum"+string(*e)]; ok {
		return betapb.NetworkconnectivityBetaSpokeStateEnum(v)
	}
	return betapb.NetworkconnectivityBetaSpokeStateEnum(0)
}

// SpokeLinkedVpnTunnelsToProto converts a SpokeLinkedVpnTunnels object to its proto representation.
func NetworkconnectivityBetaSpokeLinkedVpnTunnelsToProto(o *beta.SpokeLinkedVpnTunnels) *betapb.NetworkconnectivityBetaSpokeLinkedVpnTunnels {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaSpokeLinkedVpnTunnels{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedInterconnectAttachmentsToProto converts a SpokeLinkedInterconnectAttachments object to its proto representation.
func NetworkconnectivityBetaSpokeLinkedInterconnectAttachmentsToProto(o *beta.SpokeLinkedInterconnectAttachments) *betapb.NetworkconnectivityBetaSpokeLinkedInterconnectAttachments {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaSpokeLinkedInterconnectAttachments{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedRouterApplianceInstancesToProto converts a SpokeLinkedRouterApplianceInstances object to its proto representation.
func NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesToProto(o *beta.SpokeLinkedRouterApplianceInstances) *betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstances {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstances{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sInstances := make([]*betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstancesToProto(&r)
	}
	p.SetInstances(sInstances)
	return p
}

// SpokeLinkedRouterApplianceInstancesInstancesToProto converts a SpokeLinkedRouterApplianceInstancesInstances object to its proto representation.
func NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstancesToProto(o *beta.SpokeLinkedRouterApplianceInstancesInstances) *betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesInstances{}
	p.SetVirtualMachine(dcl.ValueOrEmptyString(o.VirtualMachine))
	p.SetIpAddress(dcl.ValueOrEmptyString(o.IPAddress))
	return p
}

// SpokeLinkedVPCNetworkToProto converts a SpokeLinkedVPCNetwork object to its proto representation.
func NetworkconnectivityBetaSpokeLinkedVPCNetworkToProto(o *beta.SpokeLinkedVPCNetwork) *betapb.NetworkconnectivityBetaSpokeLinkedVPCNetwork {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkconnectivityBetaSpokeLinkedVPCNetwork{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	sExcludeExportRanges := make([]string, len(o.ExcludeExportRanges))
	for i, r := range o.ExcludeExportRanges {
		sExcludeExportRanges[i] = r
	}
	p.SetExcludeExportRanges(sExcludeExportRanges)
	return p
}

// SpokeToProto converts a Spoke resource to its proto representation.
func SpokeToProto(resource *beta.Spoke) *betapb.NetworkconnectivityBetaSpoke {
	p := &betapb.NetworkconnectivityBetaSpoke{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHub(dcl.ValueOrEmptyString(resource.Hub))
	p.SetLinkedVpnTunnels(NetworkconnectivityBetaSpokeLinkedVpnTunnelsToProto(resource.LinkedVpnTunnels))
	p.SetLinkedInterconnectAttachments(NetworkconnectivityBetaSpokeLinkedInterconnectAttachmentsToProto(resource.LinkedInterconnectAttachments))
	p.SetLinkedRouterApplianceInstances(NetworkconnectivityBetaSpokeLinkedRouterApplianceInstancesToProto(resource.LinkedRouterApplianceInstances))
	p.SetLinkedVpcNetwork(NetworkconnectivityBetaSpokeLinkedVPCNetworkToProto(resource.LinkedVPCNetwork))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetState(NetworkconnectivityBetaSpokeStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applySpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) applySpoke(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkconnectivityBetaSpokeRequest) (*betapb.NetworkconnectivityBetaSpoke, error) {
	p := ProtoToSpoke(request.GetResource())
	res, err := c.ApplySpoke(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SpokeToProto(res)
	return r, nil
}

// applyNetworkconnectivityBetaSpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) ApplyNetworkconnectivityBetaSpoke(ctx context.Context, request *betapb.ApplyNetworkconnectivityBetaSpokeRequest) (*betapb.NetworkconnectivityBetaSpoke, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySpoke(ctx, cl, request)
}

// DeleteSpoke handles the gRPC request by passing it to the underlying Spoke Delete() method.
func (s *SpokeServer) DeleteNetworkconnectivityBetaSpoke(ctx context.Context, request *betapb.DeleteNetworkconnectivityBetaSpokeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSpoke(ctx, ProtoToSpoke(request.GetResource()))

}

// ListNetworkconnectivityBetaSpoke handles the gRPC request by passing it to the underlying SpokeList() method.
func (s *SpokeServer) ListNetworkconnectivityBetaSpoke(ctx context.Context, request *betapb.ListNetworkconnectivityBetaSpokeRequest) (*betapb.ListNetworkconnectivityBetaSpokeResponse, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSpoke(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkconnectivityBetaSpoke
	for _, r := range resources.Items {
		rp := SpokeToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkconnectivityBetaSpokeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSpoke(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
