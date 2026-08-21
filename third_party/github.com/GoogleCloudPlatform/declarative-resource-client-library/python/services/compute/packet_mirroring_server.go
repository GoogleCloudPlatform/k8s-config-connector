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

// PacketMirroringServer implements the gRPC interface for PacketMirroring.
type PacketMirroringServer struct{}

// ProtoToPacketMirroringFilterDirectionEnum converts a PacketMirroringFilterDirectionEnum enum from its proto representation.
func ProtoToComputePacketMirroringFilterDirectionEnum(e computepb.ComputePacketMirroringFilterDirectionEnum) *compute.PacketMirroringFilterDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputePacketMirroringFilterDirectionEnum_name[int32(e)]; ok {
		e := compute.PacketMirroringFilterDirectionEnum(n[len("ComputePacketMirroringFilterDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringEnableEnum converts a PacketMirroringEnableEnum enum from its proto representation.
func ProtoToComputePacketMirroringEnableEnum(e computepb.ComputePacketMirroringEnableEnum) *compute.PacketMirroringEnableEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputePacketMirroringEnableEnum_name[int32(e)]; ok {
		e := compute.PacketMirroringEnableEnum(n[len("ComputePacketMirroringEnableEnum"):])
		return &e
	}
	return nil
}

// ProtoToPacketMirroringNetwork converts a PacketMirroringNetwork object from its proto representation.
func ProtoToComputePacketMirroringNetwork(p *computepb.ComputePacketMirroringNetwork) *compute.PacketMirroringNetwork {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringNetwork{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringCollectorIlb converts a PacketMirroringCollectorIlb object from its proto representation.
func ProtoToComputePacketMirroringCollectorIlb(p *computepb.ComputePacketMirroringCollectorIlb) *compute.PacketMirroringCollectorIlb {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringCollectorIlb{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResources converts a PacketMirroringMirroredResources object from its proto representation.
func ProtoToComputePacketMirroringMirroredResources(p *computepb.ComputePacketMirroringMirroredResources) *compute.PacketMirroringMirroredResources {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResources{}
	for _, r := range p.GetSubnetworks() {
		obj.Subnetworks = append(obj.Subnetworks, *ProtoToComputePacketMirroringMirroredResourcesSubnetworks(r))
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, *ProtoToComputePacketMirroringMirroredResourcesInstances(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesSubnetworks converts a PacketMirroringMirroredResourcesSubnetworks object from its proto representation.
func ProtoToComputePacketMirroringMirroredResourcesSubnetworks(p *computepb.ComputePacketMirroringMirroredResourcesSubnetworks) *compute.PacketMirroringMirroredResourcesSubnetworks {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResourcesSubnetworks{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringMirroredResourcesInstances converts a PacketMirroringMirroredResourcesInstances object from its proto representation.
func ProtoToComputePacketMirroringMirroredResourcesInstances(p *computepb.ComputePacketMirroringMirroredResourcesInstances) *compute.PacketMirroringMirroredResourcesInstances {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringMirroredResourcesInstances{
		Url:          dcl.StringOrNil(p.GetUrl()),
		CanonicalUrl: dcl.StringOrNil(p.GetCanonicalUrl()),
	}
	return obj
}

// ProtoToPacketMirroringFilter converts a PacketMirroringFilter object from its proto representation.
func ProtoToComputePacketMirroringFilter(p *computepb.ComputePacketMirroringFilter) *compute.PacketMirroringFilter {
	if p == nil {
		return nil
	}
	obj := &compute.PacketMirroringFilter{
		Direction: ProtoToComputePacketMirroringFilterDirectionEnum(p.GetDirection()),
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
func ProtoToPacketMirroring(p *computepb.ComputePacketMirroring) *compute.PacketMirroring {
	obj := &compute.PacketMirroring{
		Id:                dcl.Int64OrNil(p.GetId()),
		SelfLink:          dcl.StringOrNil(p.GetSelfLink()),
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Region:            dcl.StringOrNil(p.GetRegion()),
		Network:           ProtoToComputePacketMirroringNetwork(p.GetNetwork()),
		Priority:          dcl.Int64OrNil(p.GetPriority()),
		CollectorIlb:      ProtoToComputePacketMirroringCollectorIlb(p.GetCollectorIlb()),
		MirroredResources: ProtoToComputePacketMirroringMirroredResources(p.GetMirroredResources()),
		Filter:            ProtoToComputePacketMirroringFilter(p.GetFilter()),
		Enable:            ProtoToComputePacketMirroringEnableEnum(p.GetEnable()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// PacketMirroringFilterDirectionEnumToProto converts a PacketMirroringFilterDirectionEnum enum to its proto representation.
func ComputePacketMirroringFilterDirectionEnumToProto(e *compute.PacketMirroringFilterDirectionEnum) computepb.ComputePacketMirroringFilterDirectionEnum {
	if e == nil {
		return computepb.ComputePacketMirroringFilterDirectionEnum(0)
	}
	if v, ok := computepb.ComputePacketMirroringFilterDirectionEnum_value["PacketMirroringFilterDirectionEnum"+string(*e)]; ok {
		return computepb.ComputePacketMirroringFilterDirectionEnum(v)
	}
	return computepb.ComputePacketMirroringFilterDirectionEnum(0)
}

// PacketMirroringEnableEnumToProto converts a PacketMirroringEnableEnum enum to its proto representation.
func ComputePacketMirroringEnableEnumToProto(e *compute.PacketMirroringEnableEnum) computepb.ComputePacketMirroringEnableEnum {
	if e == nil {
		return computepb.ComputePacketMirroringEnableEnum(0)
	}
	if v, ok := computepb.ComputePacketMirroringEnableEnum_value["PacketMirroringEnableEnum"+string(*e)]; ok {
		return computepb.ComputePacketMirroringEnableEnum(v)
	}
	return computepb.ComputePacketMirroringEnableEnum(0)
}

// PacketMirroringNetworkToProto converts a PacketMirroringNetwork object to its proto representation.
func ComputePacketMirroringNetworkToProto(o *compute.PacketMirroringNetwork) *computepb.ComputePacketMirroringNetwork {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringNetwork{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringCollectorIlbToProto converts a PacketMirroringCollectorIlb object to its proto representation.
func ComputePacketMirroringCollectorIlbToProto(o *compute.PacketMirroringCollectorIlb) *computepb.ComputePacketMirroringCollectorIlb {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringCollectorIlb{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringMirroredResourcesToProto converts a PacketMirroringMirroredResources object to its proto representation.
func ComputePacketMirroringMirroredResourcesToProto(o *compute.PacketMirroringMirroredResources) *computepb.ComputePacketMirroringMirroredResources {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResources{}
	sSubnetworks := make([]*computepb.ComputePacketMirroringMirroredResourcesSubnetworks, len(o.Subnetworks))
	for i, r := range o.Subnetworks {
		sSubnetworks[i] = ComputePacketMirroringMirroredResourcesSubnetworksToProto(&r)
	}
	p.SetSubnetworks(sSubnetworks)
	sInstances := make([]*computepb.ComputePacketMirroringMirroredResourcesInstances, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = ComputePacketMirroringMirroredResourcesInstancesToProto(&r)
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
func ComputePacketMirroringMirroredResourcesSubnetworksToProto(o *compute.PacketMirroringMirroredResourcesSubnetworks) *computepb.ComputePacketMirroringMirroredResourcesSubnetworks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResourcesSubnetworks{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringMirroredResourcesInstancesToProto converts a PacketMirroringMirroredResourcesInstances object to its proto representation.
func ComputePacketMirroringMirroredResourcesInstancesToProto(o *compute.PacketMirroringMirroredResourcesInstances) *computepb.ComputePacketMirroringMirroredResourcesInstances {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringMirroredResourcesInstances{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetCanonicalUrl(dcl.ValueOrEmptyString(o.CanonicalUrl))
	return p
}

// PacketMirroringFilterToProto converts a PacketMirroringFilter object to its proto representation.
func ComputePacketMirroringFilterToProto(o *compute.PacketMirroringFilter) *computepb.ComputePacketMirroringFilter {
	if o == nil {
		return nil
	}
	p := &computepb.ComputePacketMirroringFilter{}
	p.SetDirection(ComputePacketMirroringFilterDirectionEnumToProto(o.Direction))
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
func PacketMirroringToProto(resource *compute.PacketMirroring) *computepb.ComputePacketMirroring {
	p := &computepb.ComputePacketMirroring{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetNetwork(ComputePacketMirroringNetworkToProto(resource.Network))
	p.SetPriority(dcl.ValueOrEmptyInt64(resource.Priority))
	p.SetCollectorIlb(ComputePacketMirroringCollectorIlbToProto(resource.CollectorIlb))
	p.SetMirroredResources(ComputePacketMirroringMirroredResourcesToProto(resource.MirroredResources))
	p.SetFilter(ComputePacketMirroringFilterToProto(resource.Filter))
	p.SetEnable(ComputePacketMirroringEnableEnumToProto(resource.Enable))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyPacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) applyPacketMirroring(ctx context.Context, c *compute.Client, request *computepb.ApplyComputePacketMirroringRequest) (*computepb.ComputePacketMirroring, error) {
	p := ProtoToPacketMirroring(request.GetResource())
	res, err := c.ApplyPacketMirroring(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PacketMirroringToProto(res)
	return r, nil
}

// applyComputePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Apply() method.
func (s *PacketMirroringServer) ApplyComputePacketMirroring(ctx context.Context, request *computepb.ApplyComputePacketMirroringRequest) (*computepb.ComputePacketMirroring, error) {
	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPacketMirroring(ctx, cl, request)
}

// DeletePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroring Delete() method.
func (s *PacketMirroringServer) DeleteComputePacketMirroring(ctx context.Context, request *computepb.DeleteComputePacketMirroringRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePacketMirroring(ctx, ProtoToPacketMirroring(request.GetResource()))

}

// ListComputePacketMirroring handles the gRPC request by passing it to the underlying PacketMirroringList() method.
func (s *PacketMirroringServer) ListComputePacketMirroring(ctx context.Context, request *computepb.ListComputePacketMirroringRequest) (*computepb.ListComputePacketMirroringResponse, error) {
	cl, err := createConfigPacketMirroring(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPacketMirroring(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputePacketMirroring
	for _, r := range resources.Items {
		rp := PacketMirroringToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputePacketMirroringResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPacketMirroring(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
