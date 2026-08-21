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

// PacketMirroringServer implements the gRPC interface for PacketMirroring.
type PacketMirroringServer struct{}

// ProtoToPacketMirroringFilterDirectionEnum converts a PacketMirroringFilterDirectionEnum enum from its proto representation.
func ProtoToComputeAlphaPacketMirroringFilterDirectionEnum(e alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum) *alpha.PacketMirroringFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum_name[int32(e)]; ok {
		e := alpha.PacketMirroringFilterDirectionEnum(n[len("ComputeAlphaPacketMirroringFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringEnableEnum converts a PacketMirroringEnableEnum enum from its proto representation.
func ProtoToComputeAlphaPacketMirroringEnableEnum(e alphapb.ComputeAlphaPacketMirroringEnableEnum) *alpha.PacketMirroringEnableEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaPacketMirroringEnableEnum_name[int32(e)]; ok {
		e := alpha.PacketMirroringEnableEnum(n[len("ComputeAlphaPacketMirroringEnableEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringNetwork converts a PacketMirroringNetwork object from its proto representation.
func ProtoToComputeAlphaPacketMirroringNetwork(p *alphapb.ComputeAlphaPacketMirroringNetwork) *alpha.PacketMirroringNetwork {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringNetwork{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringCollectorIlb converts a PacketMirroringCollectorIlb object from its proto representation.
func ProtoToComputeAlphaPacketMirroringCollectorIlb(p *alphapb.ComputeAlphaPacketMirroringCollectorIlb) *alpha.PacketMirroringCollectorIlb {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringCollectorIlb{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResources converts a PacketMirroringMirroredResources object from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResources(p *alphapb.ComputeAlphaPacketMirroringMirroredResources) *alpha.PacketMirroringMirroredResources {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResources{}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputeAlphaPacketMirroringMirroredResourcesSubnetworks(r))
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToComputeAlphaPacketMirroringMirroredResourcesInstances(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesSubnetworks converts a PacketMirroringMirroredResourcesSubnetworks object from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResourcesSubnetworks(p *alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks) *alpha.PacketMirroringMirroredResourcesSubnetworks {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesInstances converts a PacketMirroringMirroredResourcesInstances object from its proto representation.
func ProtoToComputeAlphaPacketMirroringMirroredResourcesInstances(p *alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances) *alpha.PacketMirroringMirroredResourcesInstances {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringMirroredResourcesInstances{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringFilter converts a PacketMirroringFilter object from its proto representation.
func ProtoToComputeAlphaPacketMirroringFilter(p *alphapb.ComputeAlphaPacketMirroringFilter) *alpha.PacketMirroringFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.PacketMirroringFilter{
		Direction: ProtoToComputeAlphaPacketMirroringFilterDirectionEnum(p.GetDirection()),
	}
	for _, r := range p.GetCidrRanges() {
		obj.CidrRanges = append(obj.CidrRanges, r)
	}
	for _, r := range p.GetIpProtocols() {
		obj.IPProtocols = append(obj.IPProtocols, r)
	}
	return obj
}

// ProtoToPacketMirroring converts a PacketMirroring resource from its proto representation.
func ProtoToPacketMirroring(p *alphapb.ComputeAlphaPacketMirroring) *alpha.PacketMirroring {
	obj := &alpha.PacketMirroring{
		Id:                dcl.Int64OrNil(p.GetId()),
		SelfLink:          dcl.StringOrNil(p.GetSelfLink()),
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Region:            dcl.StringOrNil(p.GetRegion()),
		Network:           ProtoToComputeAlphaPacketMirroringNetwork(p.GetNetwork()),
		Priority:          dcl.Int64OrNil(p.GetPriority()),
		CollectorIlb:      ProtoToComputeAlphaPacketMirroringCollectorIlb(p.GetCollectorIlb()),
		MirroredResources: ProtoToComputeAlphaPacketMirroringMirroredResources(p.GetMirroredResources()),
		Filter:            ProtoToComputeAlphaPacketMirroringFilter(p.GetFilter()),
		Enable:            ProtoToComputeAlphaPacketMirroringEnableEnum(p.GetEnable()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// PacketMirroringFilterDirectionEnumToProto converts a PacketMirroringFilterDirectionEnum enum to its proto representation.
func ComputeAlphaPacketMirroringFilterDirectionEnumToProto(e *alpha.PacketMirroringFilterDirectionEnum) alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum {
	if e == nil {
		return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum_value["PacketMirroringFilterDirectionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(v)
	}
	return alphapb.ComputeAlphaPacketMirroringFilterDirectionEnum(0)
}

// PacketMirroringEnableEnumToProto converts a PacketMirroringEnableEnum enum to its proto representation.
func ComputeAlphaPacketMirroringEnableEnumToProto(e *alpha.PacketMirroringEnableEnum) alphapb.ComputeAlphaPacketMirroringEnableEnum {
	if e == nil {
		return alphapb.ComputeAlphaPacketMirroringEnableEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaPacketMirroringEnableEnum_value["PacketMirroringEnableEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaPacketMirroringEnableEnum(v)
	}
	return alphapb.ComputeAlphaPacketMirroringEnableEnum(0)
}

// PacketMirroringNetworkToProto converts a PacketMirroringNetwork object to its proto representation.
func ComputeAlphaPacketMirroringNetworkToProto(o *alpha.PacketMirroringNetwork) *alphapb.ComputeAlphaPacketMirroringNetwork {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringNetwork{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringCollectorIlbToProto converts a PacketMirroringCollectorIlb object to its proto representation.
func ComputeAlphaPacketMirroringCollectorIlbToProto(o *alpha.PacketMirroringCollectorIlb) *alphapb.ComputeAlphaPacketMirroringCollectorIlb {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringCollectorIlb{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringMirroredResourcesToProto converts a PacketMirroringMirroredResources object to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesToProto(o *alpha.PacketMirroringMirroredResources) *alphapb.ComputeAlphaPacketMirroringMirroredResources {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResources{}
	sSubnetworks := make([]*alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks, len(o.Subnetworks))
	for i, r := range o.Subnetworks {
		sSubnetworks[i] = ComputeAlphaPacketMirroringMirroredResourcesSubnetworksToProto(&r)
	}
	p.SetSubnetworks(sSubnetworks)
	sInstances := make([]*alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = ComputeAlphaPacketMirroringMirroredResourcesInstancesToProto(&r)
	}
	p.SetInstances(sInstances)
	sTags := make([]string, len(o.Tags))
	for i, r := range o.Tags {
		sTags[i] = r
	}
	p.SetTags(sTags)
	return p
}

// PacketMirroringMirroredResourcesSubnetworksToProto converts a PacketMirroringMirroredResourcesSubnetworks object to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesSubnetworksToProto(o *alpha.PacketMirroringMirroredResourcesSubnetworks) *alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResourcesSubnetworks{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringMirroredResourcesInstancesToProto converts a PacketMirroringMirroredResourcesInstances object to its proto representation.
func ComputeAlphaPacketMirroringMirroredResourcesInstancesToProto(o *alpha.PacketMirroringMirroredResourcesInstances) *alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringMirroredResourcesInstances{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringFilterToProto converts a PacketMirroringFilter object to its proto representation.
func ComputeAlphaPacketMirroringFilterToProto(o *alpha.PacketMirroringFilter) *alphapb.ComputeAlphaPacketMirroringFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaPacketMirroringFilter{}
	p.SetDirection(ComputeAlphaPacketMirroringFilterDirectionEnumToProto(o.Direction))
	sCidrRanges := make([]string, len(o.CidrRanges))
	for i, r := range o.CidrRanges {
		sCidrRanges[i] = r
	}
	p.SetCidrRanges(sCidrRanges)
	sIPProtocols := make([]string, len(o.IPProtocols))
	for i, r := range o.IPProtocols {
		sIPProtocols[i] = r
	}
	p.SetIpProtocols(sIPProtocols)
	return p
}

// PacketMirroringToProto converts a PacketMirroring resource to its proto representation.
func PacketMirroringToProto(resource *alpha.PacketMirroring) *alphapb.ComputeAlphaPacketMirroring {
	p := &alphapb.ComputeAlphaPacketMirroring{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetNetwork(ComputeAlphaPacketMirroringNetworkToProto(resource.Network))
	p.SetPriority(dcl.ValueOrEmptyInt64(resource.Priority))
	p.SetCollectorIlb(ComputeAlphaPacketMirroringCollectorIlbToProto(resource.CollectorIlb))
	p.SetMirroredResources(ComputeAlphaPacketMirroringMirroredResourcesToProto(resource.MirroredResources))
	p.SetFilter(ComputeAlphaPacketMirroringFilterToProto(resource.Filter))
	p.SetEnable(ComputeAlphaPacketMirroringEnableEnumToProto(resource.Enable))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) applyPacketMirroring(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaPacketMirroringRequest) (*alphapb.ComputeAlphaPacketMirroring, error) {
	p := ProtoToPacketMirroring(request.GetResource())
	res, err := c.ApplyPacketMirroring(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PacketMirroringToProto(res)
	return r, nil
}

// applyComputeAlphaPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) ApplyComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.ApplyComputeAlphaPacketMirroringRequest) (*alphapb.ComputeAlphaPacketMirroring, error) {
	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPacketMirroring(ctx, cl, request)
}

// DeletePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Delete() method.
func (s *PacketMirroringServer) DeleteComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.DeleteComputeAlphaPacketMirroringRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePacketMirroring(ctx, ProtoToPacketMirroring(request.GetResource()))

}

// ListComputeAlphaPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroringList() method.
func (s *PacketMirroringServer) ListComputeAlphaPacketMirroring(ctx context.Context, request *alphapb.ListComputeAlphaPacketMirroringRequest) (*alphapb.ListComputeAlphaPacketMirroringResponse, error) {
	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPacketMirroring(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaPacketMirroring
	for _, r := range resources.Items {
		rp := PacketMirroringToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaPacketMirroringResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPacketMirroring(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
