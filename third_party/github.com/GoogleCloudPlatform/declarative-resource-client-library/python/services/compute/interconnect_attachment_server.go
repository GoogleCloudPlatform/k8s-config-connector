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

// InterconnectAttachmentServer implements the gRPC interface for InterconnectAttachment.
type InterconnectAttachmentServer struct{}

// ProtoToInterconnectAttachmentOperationalStatusEnum converts a InterconnectAttachmentOperationalStatusEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentOperationalStatusEnum(e computepb.ComputeInterconnectAttachmentOperationalStatusEnum) *compute.InterconnectAttachmentOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentOperationalStatusEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentOperationalStatusEnum(n[len("ComputeInterconnectAttachmentOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentTypeEnum converts a InterconnectAttachmentTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentTypeEnum(e computepb.ComputeInterconnectAttachmentTypeEnum) *compute.InterconnectAttachmentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentTypeEnum(n[len("ComputeInterconnectAttachmentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEdgeAvailabilityDomainEnum converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentEdgeAvailabilityDomainEnum(e computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum) *compute.InterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentEdgeAvailabilityDomainEnum(n[len("ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentBandwidthEnum converts a InterconnectAttachmentBandwidthEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentBandwidthEnum(e computepb.ComputeInterconnectAttachmentBandwidthEnum) *compute.InterconnectAttachmentBandwidthEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentBandwidthEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentBandwidthEnum(n[len("ComputeInterconnectAttachmentBandwidthEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentStateEnum converts a InterconnectAttachmentStateEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentStateEnum(e computepb.ComputeInterconnectAttachmentStateEnum) *compute.InterconnectAttachmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentStateEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentStateEnum(n[len("ComputeInterconnectAttachmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEncryptionEnum converts a InterconnectAttachmentEncryptionEnum enum from its proto representation.
func ProtoToComputeInterconnectAttachmentEncryptionEnum(e computepb.ComputeInterconnectAttachmentEncryptionEnum) *compute.InterconnectAttachmentEncryptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectAttachmentEncryptionEnum_name[int32(e)]; ok {
		e := compute.InterconnectAttachmentEncryptionEnum(n[len("ComputeInterconnectAttachmentEncryptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentPrivateInterconnectInfo converts a InterconnectAttachmentPrivateInterconnectInfo object from its proto representation.
func ProtoToComputeInterconnectAttachmentPrivateInterconnectInfo(p *computepb.ComputeInterconnectAttachmentPrivateInterconnectInfo) *compute.InterconnectAttachmentPrivateInterconnectInfo {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectAttachmentPrivateInterconnectInfo{
		Tag8021q: dcl.Int64OrNil(p.GetTag8021Q()),
	}
	return obj
}

// ProtoToInterconnectAttachmentPartnerMetadata converts a InterconnectAttachmentPartnerMetadata object from its proto representation.
func ProtoToComputeInterconnectAttachmentPartnerMetadata(p *computepb.ComputeInterconnectAttachmentPartnerMetadata) *compute.InterconnectAttachmentPartnerMetadata {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectAttachmentPartnerMetadata{
		PartnerName:      dcl.StringOrNil(p.GetPartnerName()),
		InterconnectName: dcl.StringOrNil(p.GetInterconnectName()),
		PortalUrl:        dcl.StringOrNil(p.GetPortalUrl()),
	}
	return obj
}

// ProtoToInterconnectAttachment converts a InterconnectAttachment resource from its proto representation.
func ProtoToInterconnectAttachment(p *computepb.ComputeInterconnectAttachment) *compute.InterconnectAttachment {
	obj := &compute.InterconnectAttachment{
		Description:             dcl.StringOrNil(p.GetDescription()),
		SelfLink:                dcl.StringOrNil(p.GetSelfLink()),
		Id:                      dcl.Int64OrNil(p.GetId()),
		Name:                    dcl.StringOrNil(p.GetName()),
		Interconnect:            dcl.StringOrNil(p.GetInterconnect()),
		Router:                  dcl.StringOrNil(p.GetRouter()),
		Region:                  dcl.StringOrNil(p.GetRegion()),
		Mtu:                     dcl.Int64OrNil(p.GetMtu()),
		PrivateInterconnectInfo: ProtoToComputeInterconnectAttachmentPrivateInterconnectInfo(p.GetPrivateInterconnectInfo()),
		OperationalStatus:       ProtoToComputeInterconnectAttachmentOperationalStatusEnum(p.GetOperationalStatus()),
		CloudRouterIPAddress:    dcl.StringOrNil(p.GetCloudRouterIpAddress()),
		CustomerRouterIPAddress: dcl.StringOrNil(p.GetCustomerRouterIpAddress()),
		Type:                    ProtoToComputeInterconnectAttachmentTypeEnum(p.GetType()),
		PairingKey:              dcl.StringOrNil(p.GetPairingKey()),
		AdminEnabled:            dcl.Bool(p.GetAdminEnabled()),
		VlanTag8021q:            dcl.Int64OrNil(p.GetVlanTag8021Q()),
		EdgeAvailabilityDomain:  ProtoToComputeInterconnectAttachmentEdgeAvailabilityDomainEnum(p.GetEdgeAvailabilityDomain()),
		Bandwidth:               ProtoToComputeInterconnectAttachmentBandwidthEnum(p.GetBandwidth()),
		PartnerMetadata:         ProtoToComputeInterconnectAttachmentPartnerMetadata(p.GetPartnerMetadata()),
		State:                   ProtoToComputeInterconnectAttachmentStateEnum(p.GetState()),
		PartnerAsn:              dcl.Int64OrNil(p.GetPartnerAsn()),
		Encryption:              ProtoToComputeInterconnectAttachmentEncryptionEnum(p.GetEncryption()),
		DataplaneVersion:        dcl.Int64OrNil(p.GetDataplaneVersion()),
		SatisfiesPzs:            dcl.Bool(p.GetSatisfiesPzs()),
		Project:                 dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetCandidateSubnets() {
		obj.CandidateSubnets = append(obj.CandidateSubnets, r)
	}
	for _, r := range p.GetIpsecInternalAddresses() {
		obj.IpsecInternalAddresses = append(obj.IpsecInternalAddresses, r)
	}
	return obj
}

// InterconnectAttachmentOperationalStatusEnumToProto converts a InterconnectAttachmentOperationalStatusEnum enum to its proto representation.
func ComputeInterconnectAttachmentOperationalStatusEnumToProto(e *compute.InterconnectAttachmentOperationalStatusEnum) computepb.ComputeInterconnectAttachmentOperationalStatusEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentOperationalStatusEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentOperationalStatusEnum_value["InterconnectAttachmentOperationalStatusEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentOperationalStatusEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentOperationalStatusEnum(0)
}

// InterconnectAttachmentTypeEnumToProto converts a InterconnectAttachmentTypeEnum enum to its proto representation.
func ComputeInterconnectAttachmentTypeEnumToProto(e *compute.InterconnectAttachmentTypeEnum) computepb.ComputeInterconnectAttachmentTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentTypeEnum_value["InterconnectAttachmentTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentTypeEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentTypeEnum(0)
}

// InterconnectAttachmentEdgeAvailabilityDomainEnumToProto converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum to its proto representation.
func ComputeInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(e *compute.InterconnectAttachmentEdgeAvailabilityDomainEnum) computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum_value["InterconnectAttachmentEdgeAvailabilityDomainEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
}

// InterconnectAttachmentBandwidthEnumToProto converts a InterconnectAttachmentBandwidthEnum enum to its proto representation.
func ComputeInterconnectAttachmentBandwidthEnumToProto(e *compute.InterconnectAttachmentBandwidthEnum) computepb.ComputeInterconnectAttachmentBandwidthEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentBandwidthEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentBandwidthEnum_value["InterconnectAttachmentBandwidthEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentBandwidthEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentBandwidthEnum(0)
}

// InterconnectAttachmentStateEnumToProto converts a InterconnectAttachmentStateEnum enum to its proto representation.
func ComputeInterconnectAttachmentStateEnumToProto(e *compute.InterconnectAttachmentStateEnum) computepb.ComputeInterconnectAttachmentStateEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentStateEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentStateEnum_value["InterconnectAttachmentStateEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentStateEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentStateEnum(0)
}

// InterconnectAttachmentEncryptionEnumToProto converts a InterconnectAttachmentEncryptionEnum enum to its proto representation.
func ComputeInterconnectAttachmentEncryptionEnumToProto(e *compute.InterconnectAttachmentEncryptionEnum) computepb.ComputeInterconnectAttachmentEncryptionEnum {
	if e == nil {
		return computepb.ComputeInterconnectAttachmentEncryptionEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectAttachmentEncryptionEnum_value["InterconnectAttachmentEncryptionEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectAttachmentEncryptionEnum(v)
	}
	return computepb.ComputeInterconnectAttachmentEncryptionEnum(0)
}

// InterconnectAttachmentPrivateInterconnectInfoToProto converts a InterconnectAttachmentPrivateInterconnectInfo object to its proto representation.
func ComputeInterconnectAttachmentPrivateInterconnectInfoToProto(o *compute.InterconnectAttachmentPrivateInterconnectInfo) *computepb.ComputeInterconnectAttachmentPrivateInterconnectInfo {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectAttachmentPrivateInterconnectInfo{}
	p.SetTag8021Q(dcl.ValueOrEmptyInt64(o.Tag8021q))
	return p
}

// InterconnectAttachmentPartnerMetadataToProto converts a InterconnectAttachmentPartnerMetadata object to its proto representation.
func ComputeInterconnectAttachmentPartnerMetadataToProto(o *compute.InterconnectAttachmentPartnerMetadata) *computepb.ComputeInterconnectAttachmentPartnerMetadata {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectAttachmentPartnerMetadata{}
	p.SetPartnerName(dcl.ValueOrEmptyString(o.PartnerName))
	p.SetInterconnectName(dcl.ValueOrEmptyString(o.InterconnectName))
	p.SetPortalUrl(dcl.ValueOrEmptyString(o.PortalUrl))
	return p
}

// InterconnectAttachmentToProto converts a InterconnectAttachment resource to its proto representation.
func InterconnectAttachmentToProto(resource *compute.InterconnectAttachment) *computepb.ComputeInterconnectAttachment {
	p := &computepb.ComputeInterconnectAttachment{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetInterconnect(dcl.ValueOrEmptyString(resource.Interconnect))
	p.SetRouter(dcl.ValueOrEmptyString(resource.Router))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetMtu(dcl.ValueOrEmptyInt64(resource.Mtu))
	p.SetPrivateInterconnectInfo(ComputeInterconnectAttachmentPrivateInterconnectInfoToProto(resource.PrivateInterconnectInfo))
	p.SetOperationalStatus(ComputeInterconnectAttachmentOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetCloudRouterIpAddress(dcl.ValueOrEmptyString(resource.CloudRouterIPAddress))
	p.SetCustomerRouterIpAddress(dcl.ValueOrEmptyString(resource.CustomerRouterIPAddress))
	p.SetType(ComputeInterconnectAttachmentTypeEnumToProto(resource.Type))
	p.SetPairingKey(dcl.ValueOrEmptyString(resource.PairingKey))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetVlanTag8021Q(dcl.ValueOrEmptyInt64(resource.VlanTag8021q))
	p.SetEdgeAvailabilityDomain(ComputeInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(resource.EdgeAvailabilityDomain))
	p.SetBandwidth(ComputeInterconnectAttachmentBandwidthEnumToProto(resource.Bandwidth))
	p.SetPartnerMetadata(ComputeInterconnectAttachmentPartnerMetadataToProto(resource.PartnerMetadata))
	p.SetState(ComputeInterconnectAttachmentStateEnumToProto(resource.State))
	p.SetPartnerAsn(dcl.ValueOrEmptyInt64(resource.PartnerAsn))
	p.SetEncryption(ComputeInterconnectAttachmentEncryptionEnumToProto(resource.Encryption))
	p.SetDataplaneVersion(dcl.ValueOrEmptyInt64(resource.DataplaneVersion))
	p.SetSatisfiesPzs(dcl.ValueOrEmptyBool(resource.SatisfiesPzs))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sCandidateSubnets := make([]string, len(resource.CandidateSubnets))
	for i, r := range resource.CandidateSubnets {
		sCandidateSubnets[i] = r
	}
	p.SetCandidateSubnets(sCandidateSubnets)
	sIpsecInternalAddresses := make([]string, len(resource.IpsecInternalAddresses))
	for i, r := range resource.IpsecInternalAddresses {
		sIpsecInternalAddresses[i] = r
	}
	p.SetIpsecInternalAddresses(sIpsecInternalAddresses)

	return p
}

// applyInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Apply() method.
func (s *InterconnectAttachmentServer) applyInterconnectAttachment(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInterconnectAttachmentRequest) (*computepb.ComputeInterconnectAttachment, error) {
	p := ProtoToInterconnectAttachment(request.GetResource())
	res, err := c.ApplyInterconnectAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectAttachmentToProto(res)
	return r, nil
}

// applyComputeInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Apply() method.
func (s *InterconnectAttachmentServer) ApplyComputeInterconnectAttachment(ctx context.Context, request *computepb.ApplyComputeInterconnectAttachmentRequest) (*computepb.ComputeInterconnectAttachment, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnectAttachment(ctx, cl, request)
}

// DeleteInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Delete() method.
func (s *InterconnectAttachmentServer) DeleteComputeInterconnectAttachment(ctx context.Context, request *computepb.DeleteComputeInterconnectAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnectAttachment(ctx, ProtoToInterconnectAttachment(request.GetResource()))

}

// ListComputeInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachmentList() method.
func (s *InterconnectAttachmentServer) ListComputeInterconnectAttachment(ctx context.Context, request *computepb.ListComputeInterconnectAttachmentRequest) (*computepb.ListComputeInterconnectAttachmentResponse, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnectAttachment(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInterconnectAttachment
	for _, r := range resources.Items {
		rp := InterconnectAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeInterconnectAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnectAttachment(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
