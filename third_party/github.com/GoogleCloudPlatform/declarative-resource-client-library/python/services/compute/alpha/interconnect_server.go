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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// Server implements the gRPC interface for Interconnect.
type InterconnectServer struct{}

// ProtoToInterconnectLinkTypeEnum converts a InterconnectLinkTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectLinkTypeEnum(e alphapb.ComputeAlphaInterconnectLinkTypeEnum) *alpha.InterconnectLinkTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectLinkTypeEnum_name[int32(e)]; ok {
		e := alpha.InterconnectLinkTypeEnum(n[len("ComputeAlphaInterconnectLinkTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectInterconnectTypeEnum converts a InterconnectInterconnectTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectInterconnectTypeEnum(e alphapb.ComputeAlphaInterconnectInterconnectTypeEnum) *alpha.InterconnectInterconnectTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectInterconnectTypeEnum_name[int32(e)]; ok {
		e := alpha.InterconnectInterconnectTypeEnum(n[len("ComputeAlphaInterconnectInterconnectTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectOperationalStatusEnum converts a InterconnectOperationalStatusEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectOperationalStatusEnum(e alphapb.ComputeAlphaInterconnectOperationalStatusEnum) *alpha.InterconnectOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectOperationalStatusEnum_name[int32(e)]; ok {
		e := alpha.InterconnectOperationalStatusEnum(n[len("ComputeAlphaInterconnectOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesSourceEnum converts a InterconnectExpectedOutagesSourceEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectExpectedOutagesSourceEnum(e alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum) *alpha.InterconnectExpectedOutagesSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum_name[int32(e)]; ok {
		e := alpha.InterconnectExpectedOutagesSourceEnum(n[len("ComputeAlphaInterconnectExpectedOutagesSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesStateEnum converts a InterconnectExpectedOutagesStateEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectExpectedOutagesStateEnum(e alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum) *alpha.InterconnectExpectedOutagesStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum_name[int32(e)]; ok {
		e := alpha.InterconnectExpectedOutagesStateEnum(n[len("ComputeAlphaInterconnectExpectedOutagesStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesIssueTypeEnum converts a InterconnectExpectedOutagesIssueTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectExpectedOutagesIssueTypeEnum(e alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum) *alpha.InterconnectExpectedOutagesIssueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum_name[int32(e)]; ok {
		e := alpha.InterconnectExpectedOutagesIssueTypeEnum(n[len("ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectStateEnum converts a InterconnectStateEnum enum from its proto representation.
func ProtoToComputeAlphaInterconnectStateEnum(e alphapb.ComputeAlphaInterconnectStateEnum) *alpha.InterconnectStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInterconnectStateEnum_name[int32(e)]; ok {
		e := alpha.InterconnectStateEnum(n[len("ComputeAlphaInterconnectStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutages converts a InterconnectExpectedOutages object from its proto representation.
func ProtoToComputeAlphaInterconnectExpectedOutages(p *alphapb.ComputeAlphaInterconnectExpectedOutages) *alpha.InterconnectExpectedOutages {
	if p == nil {
		return nil
	}
	obj := &alpha.InterconnectExpectedOutages{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Source:      ProtoToComputeAlphaInterconnectExpectedOutagesSourceEnum(p.GetSource()),
		State:       ProtoToComputeAlphaInterconnectExpectedOutagesStateEnum(p.GetState()),
		IssueType:   ProtoToComputeAlphaInterconnectExpectedOutagesIssueTypeEnum(p.GetIssueType()),
		StartTime:   dcl.Int64OrNil(p.GetStartTime()),
		EndTime:     dcl.Int64OrNil(p.GetEndTime()),
	}
	for _, r := range p.GetAffectedCircuits() {
		obj.AffectedCircuits = append(obj.AffectedCircuits, r)
	}
	return obj
}

// ProtoToInterconnectCircuitInfos converts a InterconnectCircuitInfos object from its proto representation.
func ProtoToComputeAlphaInterconnectCircuitInfos(p *alphapb.ComputeAlphaInterconnectCircuitInfos) *alpha.InterconnectCircuitInfos {
	if p == nil {
		return nil
	}
	obj := &alpha.InterconnectCircuitInfos{
		GoogleCircuitId:  dcl.StringOrNil(p.GetGoogleCircuitId()),
		GoogleDemarcId:   dcl.StringOrNil(p.GetGoogleDemarcId()),
		CustomerDemarcId: dcl.StringOrNil(p.GetCustomerDemarcId()),
	}
	return obj
}

// ProtoToInterconnect converts a Interconnect resource from its proto representation.
func ProtoToInterconnect(p *alphapb.ComputeAlphaInterconnect) *alpha.Interconnect {
	obj := &alpha.Interconnect{
		Description:          dcl.StringOrNil(p.GetDescription()),
		SelfLink:             dcl.StringOrNil(p.GetSelfLink()),
		Id:                   dcl.Int64OrNil(p.GetId()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		LinkType:             ProtoToComputeAlphaInterconnectLinkTypeEnum(p.GetLinkType()),
		RequestedLinkCount:   dcl.Int64OrNil(p.GetRequestedLinkCount()),
		InterconnectType:     ProtoToComputeAlphaInterconnectInterconnectTypeEnum(p.GetInterconnectType()),
		AdminEnabled:         dcl.Bool(p.GetAdminEnabled()),
		NocContactEmail:      dcl.StringOrNil(p.GetNocContactEmail()),
		CustomerName:         dcl.StringOrNil(p.GetCustomerName()),
		OperationalStatus:    ProtoToComputeAlphaInterconnectOperationalStatusEnum(p.GetOperationalStatus()),
		ProvisionedLinkCount: dcl.Int64OrNil(p.GetProvisionedLinkCount()),
		PeerIPAddress:        dcl.StringOrNil(p.GetPeerIpAddress()),
		GoogleIPAddress:      dcl.StringOrNil(p.GetGoogleIpAddress()),
		GoogleReferenceId:    dcl.StringOrNil(p.GetGoogleReferenceId()),
		State:                ProtoToComputeAlphaInterconnectStateEnum(p.GetState()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetInterconnectAttachments() {
		obj.InterconnectAttachments = append(obj.InterconnectAttachments, r)
	}
	for _, r := range p.GetExpectedOutages() {
		obj.ExpectedOutages = append(obj.ExpectedOutages, *ProtoToComputeAlphaInterconnectExpectedOutages(r))
	}
	for _, r := range p.GetCircuitInfos() {
		obj.CircuitInfos = append(obj.CircuitInfos, *ProtoToComputeAlphaInterconnectCircuitInfos(r))
	}
	return obj
}

// InterconnectLinkTypeEnumToProto converts a InterconnectLinkTypeEnum enum to its proto representation.
func ComputeAlphaInterconnectLinkTypeEnumToProto(e *alpha.InterconnectLinkTypeEnum) alphapb.ComputeAlphaInterconnectLinkTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectLinkTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectLinkTypeEnum_value["InterconnectLinkTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectLinkTypeEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectLinkTypeEnum(0)
}

// InterconnectInterconnectTypeEnumToProto converts a InterconnectInterconnectTypeEnum enum to its proto representation.
func ComputeAlphaInterconnectInterconnectTypeEnumToProto(e *alpha.InterconnectInterconnectTypeEnum) alphapb.ComputeAlphaInterconnectInterconnectTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectInterconnectTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectInterconnectTypeEnum_value["InterconnectInterconnectTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectInterconnectTypeEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectInterconnectTypeEnum(0)
}

// InterconnectOperationalStatusEnumToProto converts a InterconnectOperationalStatusEnum enum to its proto representation.
func ComputeAlphaInterconnectOperationalStatusEnumToProto(e *alpha.InterconnectOperationalStatusEnum) alphapb.ComputeAlphaInterconnectOperationalStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectOperationalStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectOperationalStatusEnum_value["InterconnectOperationalStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectOperationalStatusEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectOperationalStatusEnum(0)
}

// InterconnectExpectedOutagesSourceEnumToProto converts a InterconnectExpectedOutagesSourceEnum enum to its proto representation.
func ComputeAlphaInterconnectExpectedOutagesSourceEnumToProto(e *alpha.InterconnectExpectedOutagesSourceEnum) alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum_value["InterconnectExpectedOutagesSourceEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectExpectedOutagesSourceEnum(0)
}

// InterconnectExpectedOutagesStateEnumToProto converts a InterconnectExpectedOutagesStateEnum enum to its proto representation.
func ComputeAlphaInterconnectExpectedOutagesStateEnumToProto(e *alpha.InterconnectExpectedOutagesStateEnum) alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum_value["InterconnectExpectedOutagesStateEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectExpectedOutagesStateEnum(0)
}

// InterconnectExpectedOutagesIssueTypeEnumToProto converts a InterconnectExpectedOutagesIssueTypeEnum enum to its proto representation.
func ComputeAlphaInterconnectExpectedOutagesIssueTypeEnumToProto(e *alpha.InterconnectExpectedOutagesIssueTypeEnum) alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum_value["InterconnectExpectedOutagesIssueTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectExpectedOutagesIssueTypeEnum(0)
}

// InterconnectStateEnumToProto converts a InterconnectStateEnum enum to its proto representation.
func ComputeAlphaInterconnectStateEnumToProto(e *alpha.InterconnectStateEnum) alphapb.ComputeAlphaInterconnectStateEnum {
	if e == nil {
		return alphapb.ComputeAlphaInterconnectStateEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInterconnectStateEnum_value["InterconnectStateEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInterconnectStateEnum(v)
	}
	return alphapb.ComputeAlphaInterconnectStateEnum(0)
}

// InterconnectExpectedOutagesToProto converts a InterconnectExpectedOutages object to its proto representation.
func ComputeAlphaInterconnectExpectedOutagesToProto(o *alpha.InterconnectExpectedOutages) *alphapb.ComputeAlphaInterconnectExpectedOutages {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInterconnectExpectedOutages{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetSource(ComputeAlphaInterconnectExpectedOutagesSourceEnumToProto(o.Source))
	p.SetState(ComputeAlphaInterconnectExpectedOutagesStateEnumToProto(o.State))
	p.SetIssueType(ComputeAlphaInterconnectExpectedOutagesIssueTypeEnumToProto(o.IssueType))
	p.SetStartTime(dcl.ValueOrEmptyInt64(o.StartTime))
	p.SetEndTime(dcl.ValueOrEmptyInt64(o.EndTime))
	sAffectedCircuits := make([]string, len(o.AffectedCircuits))
	for i, r := range o.AffectedCircuits {
		sAffectedCircuits[i] = r
	}
	p.SetAffectedCircuits(sAffectedCircuits)
	return p
}

// InterconnectCircuitInfosToProto converts a InterconnectCircuitInfos object to its proto representation.
func ComputeAlphaInterconnectCircuitInfosToProto(o *alpha.InterconnectCircuitInfos) *alphapb.ComputeAlphaInterconnectCircuitInfos {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInterconnectCircuitInfos{}
	p.SetGoogleCircuitId(dcl.ValueOrEmptyString(o.GoogleCircuitId))
	p.SetGoogleDemarcId(dcl.ValueOrEmptyString(o.GoogleDemarcId))
	p.SetCustomerDemarcId(dcl.ValueOrEmptyString(o.CustomerDemarcId))
	return p
}

// InterconnectToProto converts a Interconnect resource to its proto representation.
func InterconnectToProto(resource *alpha.Interconnect) *alphapb.ComputeAlphaInterconnect {
	p := &alphapb.ComputeAlphaInterconnect{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetLinkType(ComputeAlphaInterconnectLinkTypeEnumToProto(resource.LinkType))
	p.SetRequestedLinkCount(dcl.ValueOrEmptyInt64(resource.RequestedLinkCount))
	p.SetInterconnectType(ComputeAlphaInterconnectInterconnectTypeEnumToProto(resource.InterconnectType))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetNocContactEmail(dcl.ValueOrEmptyString(resource.NocContactEmail))
	p.SetCustomerName(dcl.ValueOrEmptyString(resource.CustomerName))
	p.SetOperationalStatus(ComputeAlphaInterconnectOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetProvisionedLinkCount(dcl.ValueOrEmptyInt64(resource.ProvisionedLinkCount))
	p.SetPeerIpAddress(dcl.ValueOrEmptyString(resource.PeerIPAddress))
	p.SetGoogleIpAddress(dcl.ValueOrEmptyString(resource.GoogleIPAddress))
	p.SetGoogleReferenceId(dcl.ValueOrEmptyString(resource.GoogleReferenceId))
	p.SetState(ComputeAlphaInterconnectStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sInterconnectAttachments := make([]string, len(resource.InterconnectAttachments))
	for i, r := range resource.InterconnectAttachments {
		sInterconnectAttachments[i] = r
	}
	p.SetInterconnectAttachments(sInterconnectAttachments)
	sExpectedOutages := make([]*alphapb.ComputeAlphaInterconnectExpectedOutages, len(resource.ExpectedOutages))
	for i, r := range resource.ExpectedOutages {
		sExpectedOutages[i] = ComputeAlphaInterconnectExpectedOutagesToProto(&r)
	}
	p.SetExpectedOutages(sExpectedOutages)
	sCircuitInfos := make([]*alphapb.ComputeAlphaInterconnectCircuitInfos, len(resource.CircuitInfos))
	for i, r := range resource.CircuitInfos {
		sCircuitInfos[i] = ComputeAlphaInterconnectCircuitInfosToProto(&r)
	}
	p.SetCircuitInfos(sCircuitInfos)

	return p
}

// applyInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) applyInterconnect(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaInterconnectRequest) (*alphapb.ComputeAlphaInterconnect, error) {
	p := ProtoToInterconnect(request.GetResource())
	res, err := c.ApplyInterconnect(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectToProto(res)
	return r, nil
}

// applyComputeAlphaInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) ApplyComputeAlphaInterconnect(ctx context.Context, request *alphapb.ApplyComputeAlphaInterconnectRequest) (*alphapb.ComputeAlphaInterconnect, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnect(ctx, cl, request)
}

// DeleteInterconnect handles the gRPC request by passing it to the underlying Interconnect Delete() method.
func (s *InterconnectServer) DeleteComputeAlphaInterconnect(ctx context.Context, request *alphapb.DeleteComputeAlphaInterconnectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnect(ctx, ProtoToInterconnect(request.GetResource()))

}

// ListComputeAlphaInterconnect handles the gRPC request by passing it to the underlying InterconnectList() method.
func (s *InterconnectServer) ListComputeAlphaInterconnect(ctx context.Context, request *alphapb.ListComputeAlphaInterconnectRequest) (*alphapb.ListComputeAlphaInterconnectResponse, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnect(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaInterconnect
	for _, r := range resources.Items {
		rp := InterconnectToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaInterconnectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnect(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
