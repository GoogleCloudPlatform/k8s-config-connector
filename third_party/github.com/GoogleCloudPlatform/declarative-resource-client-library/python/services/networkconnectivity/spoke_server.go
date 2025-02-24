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
	networkconnectivitypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/networkconnectivity_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity"
)

// SpokeServer implements the gRPC interface for Spoke.
type SpokeServer struct{}

// ProtoToSpokeStateEnum converts a SpokeStateEnum enum from its proto representation.
func ProtoToNetworkconnectivitySpokeStateEnum(e networkconnectivitypb.NetworkconnectivitySpokeStateEnum) *networkconnectivity.SpokeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkconnectivitypb.NetworkconnectivitySpokeStateEnum_name[int32(e)]; ok {
		e := networkconnectivity.SpokeStateEnum(n[len("NetworkconnectivitySpokeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToSpokeLinkedVpnTunnels converts a SpokeLinkedVpnTunnels object from its proto representation.
func ProtoToNetworkconnectivitySpokeLinkedVpnTunnels(p *networkconnectivitypb.NetworkconnectivitySpokeLinkedVpnTunnels) *networkconnectivity.SpokeLinkedVpnTunnels {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.SpokeLinkedVpnTunnels{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedInterconnectAttachments converts a SpokeLinkedInterconnectAttachments object from its proto representation.
func ProtoToNetworkconnectivitySpokeLinkedInterconnectAttachments(p *networkconnectivitypb.NetworkconnectivitySpokeLinkedInterconnectAttachments) *networkconnectivity.SpokeLinkedInterconnectAttachments {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.SpokeLinkedInterconnectAttachments{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstances converts a SpokeLinkedRouterApplianceInstances object from its proto representation.
func ProtoToNetworkconnectivitySpokeLinkedRouterApplianceInstances(p *networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstances) *networkconnectivity.SpokeLinkedRouterApplianceInstances {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.SpokeLinkedRouterApplianceInstances{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToNetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances(r))
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstancesInstances converts a SpokeLinkedRouterApplianceInstancesInstances object from its proto representation.
func ProtoToNetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances(p *networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances) *networkconnectivity.SpokeLinkedRouterApplianceInstancesInstances {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.SpokeLinkedRouterApplianceInstancesInstances{
		VirtualMachine: dcl.StringOrNil(p.GetVirtualMachine()),
		IPAddress:      dcl.StringOrNil(p.GetIpAddress()),
	}
	return obj
}

// ProtoToSpokeLinkedVPCNetwork converts a SpokeLinkedVPCNetwork object from its proto representation.
func ProtoToNetworkconnectivitySpokeLinkedVPCNetwork(p *networkconnectivitypb.NetworkconnectivitySpokeLinkedVPCNetwork) *networkconnectivity.SpokeLinkedVPCNetwork {
	if p == nil {
		return nil
	}
	obj := &networkconnectivity.SpokeLinkedVPCNetwork{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	for _, r := range p.GetExcludeExportRanges() {
		obj.ExcludeExportRanges = append(obj.ExcludeExportRanges, r)
	}
	return obj
}

// ProtoToSpoke converts a Spoke resource from its proto representation.
func ProtoToSpoke(p *networkconnectivitypb.NetworkconnectivitySpoke) *networkconnectivity.Spoke {
	obj := &networkconnectivity.Spoke{
		Name:                           dcl.StringOrNil(p.GetName()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                     dcl.StringOrNil(p.GetUpdateTime()),
		Description:                    dcl.StringOrNil(p.GetDescription()),
		Hub:                            dcl.StringOrNil(p.GetHub()),
		LinkedVpnTunnels:               ProtoToNetworkconnectivitySpokeLinkedVpnTunnels(p.GetLinkedVpnTunnels()),
		LinkedInterconnectAttachments:  ProtoToNetworkconnectivitySpokeLinkedInterconnectAttachments(p.GetLinkedInterconnectAttachments()),
		LinkedRouterApplianceInstances: ProtoToNetworkconnectivitySpokeLinkedRouterApplianceInstances(p.GetLinkedRouterApplianceInstances()),
		LinkedVPCNetwork:               ProtoToNetworkconnectivitySpokeLinkedVPCNetwork(p.GetLinkedVpcNetwork()),
		UniqueId:                       dcl.StringOrNil(p.GetUniqueId()),
		State:                          ProtoToNetworkconnectivitySpokeStateEnum(p.GetState()),
		Project:                        dcl.StringOrNil(p.GetProject()),
		Location:                       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// SpokeStateEnumToProto converts a SpokeStateEnum enum to its proto representation.
func NetworkconnectivitySpokeStateEnumToProto(e *networkconnectivity.SpokeStateEnum) networkconnectivitypb.NetworkconnectivitySpokeStateEnum {
	if e == nil {
		return networkconnectivitypb.NetworkconnectivitySpokeStateEnum(0)
	}
	if v, ok := networkconnectivitypb.NetworkconnectivitySpokeStateEnum_value["SpokeStateEnum"+string(*e)]; ok {
		return networkconnectivitypb.NetworkconnectivitySpokeStateEnum(v)
	}
	return networkconnectivitypb.NetworkconnectivitySpokeStateEnum(0)
}

// SpokeLinkedVpnTunnelsToProto converts a SpokeLinkedVpnTunnels object to its proto representation.
func NetworkconnectivitySpokeLinkedVpnTunnelsToProto(o *networkconnectivity.SpokeLinkedVpnTunnels) *networkconnectivitypb.NetworkconnectivitySpokeLinkedVpnTunnels {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivitySpokeLinkedVpnTunnels{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedInterconnectAttachmentsToProto converts a SpokeLinkedInterconnectAttachments object to its proto representation.
func NetworkconnectivitySpokeLinkedInterconnectAttachmentsToProto(o *networkconnectivity.SpokeLinkedInterconnectAttachments) *networkconnectivitypb.NetworkconnectivitySpokeLinkedInterconnectAttachments {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivitySpokeLinkedInterconnectAttachments{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedRouterApplianceInstancesToProto converts a SpokeLinkedRouterApplianceInstances object to its proto representation.
func NetworkconnectivitySpokeLinkedRouterApplianceInstancesToProto(o *networkconnectivity.SpokeLinkedRouterApplianceInstances) *networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstances {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstances{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sInstances := make([]*networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstancesToProto(&r)
	}
	p.SetInstances(sInstances)
	return p
}

// SpokeLinkedRouterApplianceInstancesInstancesToProto converts a SpokeLinkedRouterApplianceInstancesInstances object to its proto representation.
func NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstancesToProto(o *networkconnectivity.SpokeLinkedRouterApplianceInstancesInstances) *networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivitySpokeLinkedRouterApplianceInstancesInstances{}
	p.SetVirtualMachine(dcl.ValueOrEmptyString(o.VirtualMachine))
	p.SetIpAddress(dcl.ValueOrEmptyString(o.IPAddress))
	return p
}

// SpokeLinkedVPCNetworkToProto converts a SpokeLinkedVPCNetwork object to its proto representation.
func NetworkconnectivitySpokeLinkedVPCNetworkToProto(o *networkconnectivity.SpokeLinkedVPCNetwork) *networkconnectivitypb.NetworkconnectivitySpokeLinkedVPCNetwork {
	if o == nil {
		return nil
	}
	p := &networkconnectivitypb.NetworkconnectivitySpokeLinkedVPCNetwork{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	sExcludeExportRanges := make([]string, len(o.ExcludeExportRanges))
	for i, r := range o.ExcludeExportRanges {
		sExcludeExportRanges[i] = r
	}
	p.SetExcludeExportRanges(sExcludeExportRanges)
	return p
}

// SpokeToProto converts a Spoke resource to its proto representation.
func SpokeToProto(resource *networkconnectivity.Spoke) *networkconnectivitypb.NetworkconnectivitySpoke {
	p := &networkconnectivitypb.NetworkconnectivitySpoke{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHub(dcl.ValueOrEmptyString(resource.Hub))
	p.SetLinkedVpnTunnels(NetworkconnectivitySpokeLinkedVpnTunnelsToProto(resource.LinkedVpnTunnels))
	p.SetLinkedInterconnectAttachments(NetworkconnectivitySpokeLinkedInterconnectAttachmentsToProto(resource.LinkedInterconnectAttachments))
	p.SetLinkedRouterApplianceInstances(NetworkconnectivitySpokeLinkedRouterApplianceInstancesToProto(resource.LinkedRouterApplianceInstances))
	p.SetLinkedVpcNetwork(NetworkconnectivitySpokeLinkedVPCNetworkToProto(resource.LinkedVPCNetwork))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetState(NetworkconnectivitySpokeStateEnumToProto(resource.State))
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
func (s *SpokeServer) applySpoke(ctx context.Context, c *networkconnectivity.Client, request *networkconnectivitypb.ApplyNetworkconnectivitySpokeRequest) (*networkconnectivitypb.NetworkconnectivitySpoke, error) {
	p := ProtoToSpoke(request.GetResource())
	res, err := c.ApplySpoke(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SpokeToProto(res)
	return r, nil
}

// applyNetworkconnectivitySpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) ApplyNetworkconnectivitySpoke(ctx context.Context, request *networkconnectivitypb.ApplyNetworkconnectivitySpokeRequest) (*networkconnectivitypb.NetworkconnectivitySpoke, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySpoke(ctx, cl, request)
}

// DeleteSpoke handles the gRPC request by passing it to the underlying Spoke Delete() method.
func (s *SpokeServer) DeleteNetworkconnectivitySpoke(ctx context.Context, request *networkconnectivitypb.DeleteNetworkconnectivitySpokeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSpoke(ctx, ProtoToSpoke(request.GetResource()))

}

// ListNetworkconnectivitySpoke handles the gRPC request by passing it to the underlying SpokeList() method.
func (s *SpokeServer) ListNetworkconnectivitySpoke(ctx context.Context, request *networkconnectivitypb.ListNetworkconnectivitySpokeRequest) (*networkconnectivitypb.ListNetworkconnectivitySpokeResponse, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSpoke(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*networkconnectivitypb.NetworkconnectivitySpoke
	for _, r := range resources.Items {
		rp := SpokeToProto(r)
		protos = append(protos, rp)
	}
	p := &networkconnectivitypb.ListNetworkconnectivitySpokeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSpoke(ctx context.Context, service_account_file string) (*networkconnectivity.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkconnectivity.NewClient(conf), nil
}
