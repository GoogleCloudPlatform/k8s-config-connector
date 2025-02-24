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

// InterconnectAttachmentServer implements the gRPC interface for InterconnectAttachment.
type InterconnectAttachmentServer struct{}

// ProtoToInterconnectAttachmentOperationalStatusEnum converts a InterconnectAttachmentOperationalStatusEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentOperationalStatusEnum(e alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum) *alpha.InterconnectAttachmentOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentOperationalStatusEnum(n[len("ComputeAlphaInterconnectAttachmentOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentTypeEnum converts a InterconnectAttachmentTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentTypeEnum(e alphapb.ComputeAlphaInterconnectAttachmentTypeEnum) *alpha.InterconnectAttachmentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentTypeEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentTypeEnum(n[len("ComputeAlphaInterconnectAttachmentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEdgeAvailabilityDomainEnum converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum(e alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum) *alpha.InterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentEdgeAvailabilityDomainEnum(n[len("ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentBandwidthEnum converts a InterconnectAttachmentBandwidthEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentBandwidthEnum(e alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum) *alpha.InterconnectAttachmentBandwidthEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentBandwidthEnum(n[len("ComputeAlphaInterconnectAttachmentBandwidthEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentStateEnum converts a InterconnectAttachmentStateEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentStateEnum(e alphapb.ComputeAlphaInterconnectAttachmentStateEnum) *alpha.InterconnectAttachmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentStateEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentStateEnum(n[len("ComputeAlphaInterconnectAttachmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentEncryptionEnum converts a InterconnectAttachmentEncryptionEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentEncryptionEnum(e alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum) *alpha.InterconnectAttachmentEncryptionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum_name[int32(e)]; ok {
		e := alpha.InterconnectAttachmentEncryptionEnum(n[len("ComputeAlphaInterconnectAttachmentEncryptionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectAttachmentPrivateInterconnectInfo converts a InterconnectAttachmentPrivateInterconnectInfo object from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentPrivateInterconnectInfo(p *alphapb.ComputeAlphaInterconnectAttachmentPrivateInterconnectInfo) *alpha.InterconnectAttachmentPrivateInterconnectInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InterconnectAttachmentPrivateInterconnectInfo{
		Tag8021q: dcl.Int64OrNil(p.GetTag8021Q()),
	}
	return obj
}

// ProtoToInterconnectAttachmentPartnerMetadata converts a InterconnectAttachmentPartnerMetadata object from its proto representation.
func ProtoToComputeAlphaInterconnectAttachmentPartnerMetadata(p *alphapb.ComputeAlphaInterconnectAttachmentPartnerMetadata) *alpha.InterconnectAttachmentPartnerMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.InterconnectAttachmentPartnerMetadata{
		PartnerName:      dcl.StringOrNil(p.GetPartnerName()),
		InterconnectName: dcl.StringOrNil(p.GetInterconnectName()),
		PortalUrl:        dcl.StringOrNil(p.GetPortalUrl()),
	}
	return obj
}

// ProtoToInterconnectAttachment converts a InterconnectAttachment resource from its proto representation.
func ProtoToInterconnectAttachment(p *alphapb.ComputeAlphaInterconnectAttachment) *alpha.InterconnectAttachment {
	obj := &alpha.InterconnectAttachment{
		Description:             dcl.StringOrNil(p.GetDescription()),
		SelfLink:                dcl.StringOrNil(p.GetSelfLink()),
		Id:                      dcl.Int64OrNil(p.GetId()),
		Name:                    dcl.StringOrNil(p.GetName()),
		Interconnect:            dcl.StringOrNil(p.GetInterconnect()),
		Router:                  dcl.StringOrNil(p.GetRouter()),
		Region:                  dcl.StringOrNil(p.GetRegion()),
		Mtu:                     dcl.Int64OrNil(p.GetMtu()),
		PrivateInterconnectInfo: ProtoToComputeAlphaInterconnectAttachmentPrivateInterconnectInfo(p.GetPrivateInterconnectInfo()),
		OperationalStatus:       ProtoToComputeAlphaInterconnectAttachmentOperationalStatusEnum(p.GetOperationalStatus()),
		CloudRouterIPAddress:    dcl.StringOrNil(p.GetCloudRouterIpAddress()),
		CustomerRouterIPAddress: dcl.StringOrNil(p.GetCustomerRouterIpAddress()),
		Type:                    ProtoToComputeAlphaInterconnectAttachmentTypeEnum(p.GetType()),
		PairingKey:              dcl.StringOrNil(p.GetPairingKey()),
		AdminEnabled:            dcl.Bool(p.GetAdminEnabled()),
		VlanTag8021q:            dcl.Int64OrNil(p.GetVlanTag8021Q()),
		EdgeAvailabilityDomain:  ProtoToComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum(p.GetEdgeAvailabilityDomain()),
		Bandwidth:               ProtoToComputeAlphaInterconnectAttachmentBandwidthEnum(p.GetBandwidth()),
		PartnerMetadata:         ProtoToComputeAlphaInterconnectAttachmentPartnerMetadata(p.GetPartnerMetadata()),
		State:                   ProtoToComputeAlphaInterconnectAttachmentStateEnum(p.GetState()),
		PartnerAsn:              dcl.Int64OrNil(p.GetPartnerAsn()),
		Encryption:              ProtoToComputeAlphaInterconnectAttachmentEncryptionEnum(p.GetEncryption()),
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
func ComputeAlphaInterconnectAttachmentOperationalStatusEnumToProto(e *alpha.InterconnectAttachmentOperationalStatusEnum) alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum_value["InterconnectAttachmentOperationalStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentOperationalStatusEnum(0)
}

// InterconnectAttachmentTypeEnumToProto converts a InterconnectAttachmentTypeEnum enum to its proto representation.
func ComputeAlphaInterconnectAttachmentTypeEnumToProto(e *alpha.InterconnectAttachmentTypeEnum) alphapb.ComputeAlphaInterconnectAttachmentTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentTypeEnum_value["InterconnectAttachmentTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentTypeEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentTypeEnum(0)
}

// InterconnectAttachmentEdgeAvailabilityDomainEnumToProto converts a InterconnectAttachmentEdgeAvailabilityDomainEnum enum to its proto representation.
func ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(e *alpha.InterconnectAttachmentEdgeAvailabilityDomainEnum) alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum_value["InterconnectAttachmentEdgeAvailabilityDomainEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnum(0)
}

// InterconnectAttachmentBandwidthEnumToProto converts a InterconnectAttachmentBandwidthEnum enum to its proto representation.
func ComputeAlphaInterconnectAttachmentBandwidthEnumToProto(e *alpha.InterconnectAttachmentBandwidthEnum) alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum_value["InterconnectAttachmentBandwidthEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentBandwidthEnum(0)
}

// InterconnectAttachmentStateEnumToProto converts a InterconnectAttachmentStateEnum enum to its proto representation.
func ComputeAlphaInterconnectAttachmentStateEnumToProto(e *alpha.InterconnectAttachmentStateEnum) alphapb.ComputeAlphaInterconnectAttachmentStateEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentStateEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentStateEnum_value["InterconnectAttachmentStateEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentStateEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentStateEnum(0)
}

// InterconnectAttachmentEncryptionEnumToProto converts a InterconnectAttachmentEncryptionEnum enum to its proto representation.
func ComputeAlphaInterconnectAttachmentEncryptionEnumToProto(e *alpha.InterconnectAttachmentEncryptionEnum) alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum_value["InterconnectAttachmentEncryptionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectAttachmentEncryptionEnum(0)
}

// InterconnectAttachmentPrivateInterconnectInfoToProto converts a InterconnectAttachmentPrivateInterconnectInfo object to its proto representation.
func ComputeAlphaInterconnectAttachmentPrivateInterconnectInfoToProto(o *alpha.InterconnectAttachmentPrivateInterconnectInfo) *alphapb.ComputeAlphaInterconnectAttachmentPrivateInterconnectInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInterconnectAttachmentPrivateInterconnectInfo{}
	p.SetTag8021Q(dcl.ValueOrEmptyInt64(o.Tag8021q))
	return p
}

// InterconnectAttachmentPartnerMetadataToProto converts a InterconnectAttachmentPartnerMetadata object to its proto representation.
func ComputeAlphaInterconnectAttachmentPartnerMetadataToProto(o *alpha.InterconnectAttachmentPartnerMetadata) *alphapb.ComputeAlphaInterconnectAttachmentPartnerMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInterconnectAttachmentPartnerMetadata{}
	p.SetPartnerName(dcl.ValueOrEmptyString(o.PartnerName))
	p.SetInterconnectName(dcl.ValueOrEmptyString(o.InterconnectName))
	p.SetPortalUrl(dcl.ValueOrEmptyString(o.PortalUrl))
	return p
}

// InterconnectAttachmentToProto converts a InterconnectAttachment resource to its proto representation.
func InterconnectAttachmentToProto(resource *alpha.InterconnectAttachment) *alphapb.ComputeAlphaInterconnectAttachment {
	p := &alphapb.ComputeAlphaInterconnectAttachment{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetInterconnect(dcl.ValueOrEmptyString(resource.Interconnect))
	p.SetRouter(dcl.ValueOrEmptyString(resource.Router))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetMtu(dcl.ValueOrEmptyInt64(resource.Mtu))
	p.SetPrivateInterconnectInfo(ComputeAlphaInterconnectAttachmentPrivateInterconnectInfoToProto(resource.PrivateInterconnectInfo))
	p.SetOperationalStatus(ComputeAlphaInterconnectAttachmentOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetCloudRouterIpAddress(dcl.ValueOrEmptyString(resource.CloudRouterIPAddress))
	p.SetCustomerRouterIpAddress(dcl.ValueOrEmptyString(resource.CustomerRouterIPAddress))
	p.SetType(ComputeAlphaInterconnectAttachmentTypeEnumToProto(resource.Type))
	p.SetPairingKey(dcl.ValueOrEmptyString(resource.PairingKey))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetVlanTag8021Q(dcl.ValueOrEmptyInt64(resource.VlanTag8021q))
	p.SetEdgeAvailabilityDomain(ComputeAlphaInterconnectAttachmentEdgeAvailabilityDomainEnumToProto(resource.EdgeAvailabilityDomain))
	p.SetBandwidth(ComputeAlphaInterconnectAttachmentBandwidthEnumToProto(resource.Bandwidth))
	p.SetPartnerMetadata(ComputeAlphaInterconnectAttachmentPartnerMetadataToProto(resource.PartnerMetadata))
	p.SetState(ComputeAlphaInterconnectAttachmentStateEnumToProto(resource.State))
	p.SetPartnerAsn(dcl.ValueOrEmptyInt64(resource.PartnerAsn))
	p.SetEncryption(ComputeAlphaInterconnectAttachmentEncryptionEnumToProto(resource.Encryption))
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
func (s *InterconnectAttachmentServer) applyInterconnectAttachment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaInterconnectAttachmentRequest) (*alphapb.ComputeAlphaInterconnectAttachment, error) {
	p := ProtoToInterconnectAttachment(request.GetResource())
	res, err := c.ApplyInterconnectAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectAttachmentToProto(res)
	return r, nil
}

// applyComputeAlphaInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Apply() method.
func (s *InterconnectAttachmentServer) ApplyComputeAlphaInterconnectAttachment(ctx context.Context, request *alphapb.ApplyComputeAlphaInterconnectAttachmentRequest) (*alphapb.ComputeAlphaInterconnectAttachment, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnectAttachment(ctx, cl, request)
}

// DeleteInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachment Delete() method.
func (s *InterconnectAttachmentServer) DeleteComputeAlphaInterconnectAttachment(ctx context.Context, request *alphapb.DeleteComputeAlphaInterconnectAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnectAttachment(ctx, ProtoToInterconnectAttachment(request.GetResource()))

}

// ListComputeAlphaInterconnectAttachment handles the gRPC request by passing it to the underlying InterconnectAttachmentList() method.
func (s *InterconnectAttachmentServer) ListComputeAlphaInterconnectAttachment(ctx context.Context, request *alphapb.ListComputeAlphaInterconnectAttachmentRequest) (*alphapb.ListComputeAlphaInterconnectAttachmentResponse, error) {
	cl, err := createConfigInterconnectAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnectAttachment(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaInterconnectAttachment
	for _, r := range resources.Items {
		rp := InterconnectAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaInterconnectAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnectAttachment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
