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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// InterconnectAttachmentServer implements the gRPC interface for InterconnectAttachment.
type InterconnectAttachmentServer struct{}

// ProtoToInterconnectAttachmentOperationalStatusEnum converts a InterconnectAttachmentOperationalStatusEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentOperationalStatusEnum(e betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum) *beta.InterconnectAttachmentOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentOperationalStatusEnum(n[len("ComputeBetaInterconnectAttachmentOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentTypeEnum converts a InterconnectAttachmentTypeEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentTypeEnum(e betapb.ComputeBetaInterconnectAttachmentTypeEnum) *beta.InterconnectAttachmentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentTypeEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentTypeEnum(n[len("ComputeBetaInterconnectAttachmentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEdgeAvailabilityDomainEnum converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum(e betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum) *beta.InterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentEdgeAvailabilityDomainEnum(n[len("ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentBandwidthEnum converts a InterconnectAttachmentBandwidthEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentBandwidthEnum(e betapb.ComputeBetaInterconnectAttachmentBandwidthEnum) *beta.InterconnectAttachmentBandwidthEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentBandwidthEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentBandwidthEnum(n[len("ComputeBetaInterconnectAttachmentBandwidthEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentStateEnum converts a InterconnectAttachmentStateEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentStateEnum(e betapb.ComputeBetaInterconnectAttachmentStateEnum) *beta.InterconnectAttachmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentStateEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentStateEnum(n[len("ComputeBetaInterconnectAttachmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEncryptionEnum converts a InterconnectAttachmentEncryptionEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentEncryptionEnum(e betapb.ComputeBetaInterconnectAttachmentEncryptionEnum) *beta.InterconnectAttachmentEncryptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectAttachmentEncryptionEnum_name[int32(e)]; ok {
		e := beta.InterconnectAttachmentEncryptionEnum(n[len("ComputeBetaInterconnectAttachmentEncryptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentPrivateInterconnectInfo converts a InterconnectAttachmentPrivateInterconnectInfo object from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentPrivateInterconnectInfo(p *betapb.ComputeBetaInterconnectAttachmentPrivateInterconnectInfo) *beta.InterconnectAttachmentPrivateInterconnectInfo {
	if p == nil {
		return nil
	}
	obj := &beta.InterconnectAttachmentPrivateInterconnectInfo{
		Tag8021q: dcl.Int64OrNil(p.GetTag8021Q()),
	}
	return obj
}

// ProtoToInterconnectAttachmentPartnerMetadata converts a InterconnectAttachmentPartnerMetadata object from its proto representation.
func ProtoToComputeBetaInterconnectAttachmentPartnerMetadata(p *betapb.ComputeBetaInterconnectAttachmentPartnerMetadata) *beta.InterconnectAttachmentPartnerMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.InterconnectAttachmentPartnerMetadata{
		PartnerName:      dcl.StringOrNil(p.GetPartnerName()),
		InterconnectName: dcl.StringOrNil(p.GetInterconnectName()),
		PortalUrl:        dcl.StringOrNil(p.GetPortalUrl()),
	}
	return obj
}

// ProtoToInterconnectAttachment converts a InterconnectAttachment resource from its proto representation.
func ProtoToInterconnectAttachment(p *betapb.ComputeBetaInterconnectAttachment) *beta.InterconnectAttachment {
	obj := &beta.InterconnectAttachment{
		Description:             dcl.StringOrNil(p.GetDescription()),
		SelfLink:                dcl.StringOrNil(p.GetSelfLink()),
		Id:                      dcl.Int64OrNil(p.GetId()),
		Name:                    dcl.StringOrNil(p.GetName()),
		Interconnect:            dcl.StringOrNil(p.GetInterconnect()),
		Router:                  dcl.StringOrNil(p.GetRouter()),
		Region:                  dcl.StringOrNil(p.GetRegion()),
		Mtu:                     dcl.Int64OrNil(p.GetMtu()),
		PrivateInterconnectInfo: ProtoToComputeBetaInterconnectAttachmentPrivateInterconnectInfo(p.GetPrivateInterconnectInfo()),
		OperationalStatus:       ProtoToComputeBetaInterconnectAttachmentOperationalStatusEnum(p.GetOperationalStatus()),
		CloudRouterIPAddress:    dcl.StringOrNil(p.GetCloudRouterIpAddress()),
		CustomerRouterIPAddress: dcl.StringOrNil(p.GetCustomerRouterIpAddress()),
		Type:                    ProtoToComputeBetaInterconnectAttachmentTypeEnum(p.GetType()),
		PairingKey:              dcl.StringOrNil(p.GetPairingKey()),
		AdminEnabled:            dcl.Bool(p.GetAdminEnabled()),
		VlanTag8021q:            dcl.Int64OrNil(p.GetVlanTag8021Q()),
		EdgeAvailabilityDomain:  ProtoToComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum(p.GetEdgeAvailabilityDomain()),
		Bandwidth:               ProtoToComputeBetaInterconnectAttachmentBandwidthEnum(p.GetBandwidth()),
		PartnerMetadata:         ProtoToComputeBetaInterconnectAttachmentPartnerMetadata(p.GetPartnerMetadata()),
		State:                   ProtoToComputeBetaInterconnectAttachmentStateEnum(p.GetState()),
		PartnerAsn:              dcl.Int64OrNil(p.GetPartnerAsn()),
		Encryption:              ProtoToComputeBetaInterconnectAttachmentEncryptionEnum(p.GetEncryption()),
		DataplaneVersion:        dcl.Int64OrNil(p.GetDataplaneVersion()),
		SatisfiesPzs:            dcl.Bool(p.GetSatisfiesPzs()),
		LabelFingerprint:        dcl.StringOrNil(p.GetLabelFingerprint()),
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
func ComputeBetaInterconnectAttachmentOperationalStatusEnumToProto(e *beta.InterconnectAttachmentOperationalStatusEnum) betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum_value["InterconnectAttachmentOperationalStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentOperationalStatusEnum(0)
}

// InterconnectAttachmentTypeEnumToProto converts a InterconnectAttachmentTypeEnum enum to its proto representation.
func ComputeBetaInterconnectAttachmentTypeEnumToProto(e *beta.InterconnectAttachmentTypeEnum) betapb.ComputeBetaInterconnectAttachmentTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentTypeEnum_value["InterconnectAttachmentTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentTypeEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentTypeEnum(0)
}

// InterconnectAttachmentEdgeAvailabilityDomainEnumToProto converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum to its proto representation.
func ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(e *beta.InterconnectAttachmentEdgeAvailabilityDomainEnum) betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum_value["InterconnectAttachmentEdgeAvailabilityDomainEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
}

// InterconnectAttachmentBandwidthEnumToProto converts a InterconnectAttachmentBandwidthEnum enum to its proto representation.
func ComputeBetaInterconnectAttachmentBandwidthEnumToProto(e *beta.InterconnectAttachmentBandwidthEnum) betapb.ComputeBetaInterconnectAttachmentBandwidthEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentBandwidthEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentBandwidthEnum_value["InterconnectAttachmentBandwidthEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentBandwidthEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentBandwidthEnum(0)
}

// InterconnectAttachmentStateEnumToProto converts a InterconnectAttachmentStateEnum enum to its proto representation.
func ComputeBetaInterconnectAttachmentStateEnumToProto(e *beta.InterconnectAttachmentStateEnum) betapb.ComputeBetaInterconnectAttachmentStateEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentStateEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentStateEnum_value["InterconnectAttachmentStateEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentStateEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentStateEnum(0)
}

// InterconnectAttachmentEncryptionEnumToProto converts a InterconnectAttachmentEncryptionEnum enum to its proto representation.
func ComputeBetaInterconnectAttachmentEncryptionEnumToProto(e *beta.InterconnectAttachmentEncryptionEnum) betapb.ComputeBetaInterconnectAttachmentEncryptionEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectAttachmentEncryptionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectAttachmentEncryptionEnum_value["InterconnectAttachmentEncryptionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectAttachmentEncryptionEnum(v)
	}
	return betapb.ComputeBetaInterconnectAttachmentEncryptionEnum(0)
}

// InterconnectAttachmentPrivateInterconnectInfoToProto converts a InterconnectAttachmentPrivateInterconnectInfo object to its proto representation.
func ComputeBetaInterconnectAttachmentPrivateInterconnectInfoToProto(o *beta.InterconnectAttachmentPrivateInterconnectInfo) *betapb.ComputeBetaInterconnectAttachmentPrivateInterconnectInfo {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInterconnectAttachmentPrivateInterconnectInfo{}
	p.SetTag8021Q(dcl.ValueOrEmptyInt64(o.Tag8021q))
	return p
}

// InterconnectAttachmentPartnerMetadataToProto converts a InterconnectAttachmentPartnerMetadata object to its proto representation.
func ComputeBetaInterconnectAttachmentPartnerMetadataToProto(o *beta.InterconnectAttachmentPartnerMetadata) *betapb.ComputeBetaInterconnectAttachmentPartnerMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInterconnectAttachmentPartnerMetadata{}
	p.SetPartnerName(dcl.ValueOrEmptyString(o.PartnerName))
	p.SetInterconnectName(dcl.ValueOrEmptyString(o.InterconnectName))
	p.SetPortalUrl(dcl.ValueOrEmptyString(o.PortalUrl))
	return p
}

// InterconnectAttachmentToProto converts a InterconnectAttachment resource to its proto representation.
func InterconnectAttachmentToProto(resource *beta.InterconnectAttachment) *betapb.ComputeBetaInterconnectAttachment {
	p := &betapb.ComputeBetaInterconnectAttachment{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetInterconnect(dcl.ValueOrEmptyString(resource.Interconnect))
	p.SetRouter(dcl.ValueOrEmptyString(resource.Router))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetMtu(dcl.ValueOrEmptyInt64(resource.Mtu))
	p.SetPrivateInterconnectInfo(ComputeBetaInterconnectAttachmentPrivateInterconnectInfoToProto(resource.PrivateInterconnectInfo))
	p.SetOperationalStatus(ComputeBetaInterconnectAttachmentOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetCloudRouterIpAddress(dcl.ValueOrEmptyString(resource.CloudRouterIPAddress))
	p.SetCustomerRouterIpAddress(dcl.ValueOrEmptyString(resource.CustomerRouterIPAddress))
	p.SetType(ComputeBetaInterconnectAttachmentTypeEnumToProto(resource.Type))
	p.SetPairingKey(dcl.ValueOrEmptyString(resource.PairingKey))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetVlanTag8021Q(dcl.ValueOrEmptyInt64(resource.VlanTag8021q))
	p.SetEdgeAvailabilityDomain(ComputeBetaInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(resource.EdgeAvailabilityDomain))
	p.SetBandwidth(ComputeBetaInterconnectAttachmentBandwidthEnumToProto(resource.Bandwidth))
	p.SetPartnerMetadata(ComputeBetaInterconnectAttachmentPartnerMetadataToProto(resource.PartnerMetadata))
	p.SetState(ComputeBetaInterconnectAttachmentStateEnumToProto(resource.State))
	p.SetPartnerAsn(dcl.ValueOrEmptyInt64(resource.PartnerAsn))
	p.SetEncryption(ComputeBetaInterconnectAttachmentEncryptionEnumToProto(resource.Encryption))
	p.SetDataplaneVersion(dcl.ValueOrEmptyInt64(resource.DataplaneVersion))
	p.SetSatisfiesPzs(dcl.ValueOrEmptyBool(resource.SatisfiesPzs))
	p.SetLabelFingerprint(dcl.ValueOrEmptyString(resource.LabelFingerprint))
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
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Apply() method.
func (s *InterconnectAttachmentServer) applyInterconnectAttachment(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInterconnectAttachmentRequest) (*betapb.ComputeBetaInterconnectAttachment, error) {
	p := ProtoToInterconnectAttachment(request.GetResource())
	res, err := c.ApplyInterconnectAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectAttachmentToProto(res)
	return r, nil
}

// applyComputeBetaInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Apply() method.
func (s *InterconnectAttachmentServer) ApplyComputeBetaInterconnectAttachment(ctx context.Context, request *betapb.ApplyComputeBetaInterconnectAttachmentRequest) (*betapb.ComputeBetaInterconnectAttachment, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnectAttachment(ctx, cl, request)
}

// DeleteInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Delete() method.
func (s *InterconnectAttachmentServer) DeleteComputeBetaInterconnectAttachment(ctx context.Context, request *betapb.DeleteComputeBetaInterconnectAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnectAttachment(ctx, ProtoToInterconnectAttachment(request.GetResource()))

}

// ListComputeBetaInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachmentList() method.
func (s *InterconnectAttachmentServer) ListComputeBetaInterconnectAttachment(ctx context.Context, request *betapb.ListComputeBetaInterconnectAttachmentRequest) (*betapb.ListComputeBetaInterconnectAttachmentResponse, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnectAttachment(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInterconnectAttachment
	for _, r := range resources.Items {
		rp := InterconnectAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaInterconnectAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnectAttachment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
