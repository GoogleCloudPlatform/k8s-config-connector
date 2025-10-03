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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkconnectivity/alpha/networkconnectivity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/alpha"
)

// SpokeServer implements the gRPC interface for Spoke.
type SpokeServer struct{}

// ProtoToSpokeStateEnum converts a SpokeStateEnum enum from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeStateEnum(e alphapb.NetworkconnectivityAlphaSpokeStateEnum) *alpha.SpokeStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkconnectivityAlphaSpokeStateEnum_name[int32(e)]; ok {
		e := alpha.SpokeStateEnum(n[len("NetworkconnectivityAlphaSpokeStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToSpokeLinkedVpnTunnels converts a SpokeLinkedVpnTunnels object from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedVpnTunnels(p *alphapb.NetworkconnectivityAlphaSpokeLinkedVpnTunnels) *alpha.SpokeLinkedVpnTunnels {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedVpnTunnels{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedInterconnectAttachments converts a SpokeLinkedInterconnectAttachments object from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedInterconnectAttachments(p *alphapb.NetworkconnectivityAlphaSpokeLinkedInterconnectAttachments) *alpha.SpokeLinkedInterconnectAttachments {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedInterconnectAttachments{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetUris() {
		obj.Uris = append(obj.Uris, r)
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstances converts a SpokeLinkedRouterApplianceInstances object from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances(p *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances) *alpha.SpokeLinkedRouterApplianceInstances {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedRouterApplianceInstances{
		SiteToSiteDataTransfer: dcl.Bool(p.GetSiteToSiteDataTransfer()),
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances(r))
	}
	return obj
}

// ProtoToSpokeLinkedRouterApplianceInstancesInstances converts a SpokeLinkedRouterApplianceInstancesInstances object from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances(p *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances) *alpha.SpokeLinkedRouterApplianceInstancesInstances {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedRouterApplianceInstancesInstances{
		VirtualMachine: dcl.StringOrNil(p.GetVirtualMachine()),
		IPAddress:      dcl.StringOrNil(p.GetIpAddress()),
	}
	return obj
}

// ProtoToSpokeLinkedVPCNetwork converts a SpokeLinkedVPCNetwork object from its proto representation.
func ProtoToNetworkconnectivityAlphaSpokeLinkedVPCNetwork(p *alphapb.NetworkconnectivityAlphaSpokeLinkedVPCNetwork) *alpha.SpokeLinkedVPCNetwork {
	if p == nil {
		return nil
	}
	obj := &alpha.SpokeLinkedVPCNetwork{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	for _, r := range p.GetExcludeExportRanges() {
		obj.ExcludeExportRanges = append(obj.ExcludeExportRanges, r)
	}
	return obj
}

// ProtoToSpoke converts a Spoke resource from its proto representation.
func ProtoToSpoke(p *alphapb.NetworkconnectivityAlphaSpoke) *alpha.Spoke {
	obj := &alpha.Spoke{
		Name:                           dcl.StringOrNil(p.GetName()),
		CreateTime:                     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                     dcl.StringOrNil(p.GetUpdateTime()),
		Description:                    dcl.StringOrNil(p.GetDescription()),
		Hub:                            dcl.StringOrNil(p.GetHub()),
		LinkedVpnTunnels:               ProtoToNetworkconnectivityAlphaSpokeLinkedVpnTunnels(p.GetLinkedVpnTunnels()),
		LinkedInterconnectAttachments:  ProtoToNetworkconnectivityAlphaSpokeLinkedInterconnectAttachments(p.GetLinkedInterconnectAttachments()),
		LinkedRouterApplianceInstances: ProtoToNetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances(p.GetLinkedRouterApplianceInstances()),
		LinkedVPCNetwork:               ProtoToNetworkconnectivityAlphaSpokeLinkedVPCNetwork(p.GetLinkedVpcNetwork()),
		UniqueId:                       dcl.StringOrNil(p.GetUniqueId()),
		State:                          ProtoToNetworkconnectivityAlphaSpokeStateEnum(p.GetState()),
		Project:                        dcl.StringOrNil(p.GetProject()),
		Location:                       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// SpokeStateEnumToProto converts a SpokeStateEnum enum to its proto representation.
func NetworkconnectivityAlphaSpokeStateEnumToProto(e *alpha.SpokeStateEnum) alphapb.NetworkconnectivityAlphaSpokeStateEnum {
	if e == nil {
		return alphapb.NetworkconnectivityAlphaSpokeStateEnum(0)
	}
	if v, ok := alphapb.NetworkconnectivityAlphaSpokeStateEnum_value["SpokeStateEnum"+string(*e)]; ok {
		return alphapb.NetworkconnectivityAlphaSpokeStateEnum(v)
	}
	return alphapb.NetworkconnectivityAlphaSpokeStateEnum(0)
}

// SpokeLinkedVpnTunnelsToProto converts a SpokeLinkedVpnTunnels object to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedVpnTunnelsToProto(o *alpha.SpokeLinkedVpnTunnels) *alphapb.NetworkconnectivityAlphaSpokeLinkedVpnTunnels {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedVpnTunnels{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedInterconnectAttachmentsToProto converts a SpokeLinkedInterconnectAttachments object to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedInterconnectAttachmentsToProto(o *alpha.SpokeLinkedInterconnectAttachments) *alphapb.NetworkconnectivityAlphaSpokeLinkedInterconnectAttachments {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedInterconnectAttachments{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sUris := make([]string, len(o.Uris))
	for i, r := range o.Uris {
		sUris[i] = r
	}
	p.SetUris(sUris)
	return p
}

// SpokeLinkedRouterApplianceInstancesToProto converts a SpokeLinkedRouterApplianceInstances object to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesToProto(o *alpha.SpokeLinkedRouterApplianceInstances) *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances{}
	p.SetSiteToSiteDataTransfer(dcl.ValueOrEmptyBool(o.SiteToSiteDataTransfer))
	sInstances := make([]*alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstancesToProto(&r)
	}
	p.SetInstances(sInstances)
	return p
}

// SpokeLinkedRouterApplianceInstancesInstancesToProto converts a SpokeLinkedRouterApplianceInstancesInstances object to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstancesToProto(o *alpha.SpokeLinkedRouterApplianceInstancesInstances) *alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesInstances{}
	p.SetVirtualMachine(dcl.ValueOrEmptyString(o.VirtualMachine))
	p.SetIpAddress(dcl.ValueOrEmptyString(o.IPAddress))
	return p
}

// SpokeLinkedVPCNetworkToProto converts a SpokeLinkedVPCNetwork object to its proto representation.
func NetworkconnectivityAlphaSpokeLinkedVPCNetworkToProto(o *alpha.SpokeLinkedVPCNetwork) *alphapb.NetworkconnectivityAlphaSpokeLinkedVPCNetwork {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkconnectivityAlphaSpokeLinkedVPCNetwork{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	sExcludeExportRanges := make([]string, len(o.ExcludeExportRanges))
	for i, r := range o.ExcludeExportRanges {
		sExcludeExportRanges[i] = r
	}
	p.SetExcludeExportRanges(sExcludeExportRanges)
	return p
}

// SpokeToProto converts a Spoke resource to its proto representation.
func SpokeToProto(resource *alpha.Spoke) *alphapb.NetworkconnectivityAlphaSpoke {
	p := &alphapb.NetworkconnectivityAlphaSpoke{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHub(dcl.ValueOrEmptyString(resource.Hub))
	p.SetLinkedVpnTunnels(NetworkconnectivityAlphaSpokeLinkedVpnTunnelsToProto(resource.LinkedVpnTunnels))
	p.SetLinkedInterconnectAttachments(NetworkconnectivityAlphaSpokeLinkedInterconnectAttachmentsToProto(resource.LinkedInterconnectAttachments))
	p.SetLinkedRouterApplianceInstances(NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstancesToProto(resource.LinkedRouterApplianceInstances))
	p.SetLinkedVpcNetwork(NetworkconnectivityAlphaSpokeLinkedVPCNetworkToProto(resource.LinkedVPCNetwork))
	p.SetUniqueId(dcl.ValueOrEmptyString(resource.UniqueId))
	p.SetState(NetworkconnectivityAlphaSpokeStateEnumToProto(resource.State))
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
func (s *SpokeServer) applySpoke(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkconnectivityAlphaSpokeRequest) (*alphapb.NetworkconnectivityAlphaSpoke, error) {
	p := ProtoToSpoke(request.GetResource())
	res, err := c.ApplySpoke(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SpokeToProto(res)
	return r, nil
}

// applyNetworkconnectivityAlphaSpoke handles the gRPC request by passing it to the underlying Spoke Apply() method.
func (s *SpokeServer) ApplyNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.ApplyNetworkconnectivityAlphaSpokeRequest) (*alphapb.NetworkconnectivityAlphaSpoke, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySpoke(ctx, cl, request)
}

// DeleteSpoke handles the gRPC request by passing it to the underlying Spoke Delete() method.
func (s *SpokeServer) DeleteNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.DeleteNetworkconnectivityAlphaSpokeRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSpoke(ctx, ProtoToSpoke(request.GetResource()))

}

// ListNetworkconnectivityAlphaSpoke handles the gRPC request by passing it to the underlying SpokeList() method.
func (s *SpokeServer) ListNetworkconnectivityAlphaSpoke(ctx context.Context, request *alphapb.ListNetworkconnectivityAlphaSpokeRequest) (*alphapb.ListNetworkconnectivityAlphaSpokeResponse, error) {
	cl, err := createConfigSpoke(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSpoke(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkconnectivityAlphaSpoke
	for _, r := range resources.Items {
		rp := SpokeToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkconnectivityAlphaSpokeResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSpoke(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
