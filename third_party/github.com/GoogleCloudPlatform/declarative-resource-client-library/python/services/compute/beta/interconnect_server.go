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

// Server implements the gRPC interface for Interconnect.
type InterconnectServer struct{}

// ProtoToInterconnectLinkTypeEnum converts a InterconnectLinkTypeEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectLinkTypeEnum(e betapb.ComputeBetaInterconnectLinkTypeEnum) *beta.InterconnectLinkTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectLinkTypeEnum_name[int32(e)]; ok {
		e := beta.InterconnectLinkTypeEnum(n[len("ComputeBetaInterconnectLinkTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectInterconnectTypeEnum converts a InterconnectInterconnectTypeEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectInterconnectTypeEnum(e betapb.ComputeBetaInterconnectInterconnectTypeEnum) *beta.InterconnectInterconnectTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectInterconnectTypeEnum_name[int32(e)]; ok {
		e := beta.InterconnectInterconnectTypeEnum(n[len("ComputeBetaInterconnectInterconnectTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectOperationalStatusEnum converts a InterconnectOperationalStatusEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectOperationalStatusEnum(e betapb.ComputeBetaInterconnectOperationalStatusEnum) *beta.InterconnectOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectOperationalStatusEnum_name[int32(e)]; ok {
		e := beta.InterconnectOperationalStatusEnum(n[len("ComputeBetaInterconnectOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesSourceEnum converts a InterconnectExpectedOutagesSourceEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectExpectedOutagesSourceEnum(e betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum) *beta.InterconnectExpectedOutagesSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum_name[int32(e)]; ok {
		e := beta.InterconnectExpectedOutagesSourceEnum(n[len("ComputeBetaInterconnectExpectedOutagesSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesStateEnum converts a InterconnectExpectedOutagesStateEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectExpectedOutagesStateEnum(e betapb.ComputeBetaInterconnectExpectedOutagesStateEnum) *beta.InterconnectExpectedOutagesStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectExpectedOutagesStateEnum_name[int32(e)]; ok {
		e := beta.InterconnectExpectedOutagesStateEnum(n[len("ComputeBetaInterconnectExpectedOutagesStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesIssueTypeEnum converts a InterconnectExpectedOutagesIssueTypeEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectExpectedOutagesIssueTypeEnum(e betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum) *beta.InterconnectExpectedOutagesIssueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum_name[int32(e)]; ok {
		e := beta.InterconnectExpectedOutagesIssueTypeEnum(n[len("ComputeBetaInterconnectExpectedOutagesIssueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectStateEnum converts a InterconnectStateEnum enum from its proto representation.
func ProtoToComputeBetaInterconnectStateEnum(e betapb.ComputeBetaInterconnectStateEnum) *beta.InterconnectStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInterconnectStateEnum_name[int32(e)]; ok {
		e := beta.InterconnectStateEnum(n[len("ComputeBetaInterconnectStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutages converts a InterconnectExpectedOutages object from its proto representation.
func ProtoToComputeBetaInterconnectExpectedOutages(p *betapb.ComputeBetaInterconnectExpectedOutages) *beta.InterconnectExpectedOutages {
	if p == nil {
		return nil
	}
	obj := &beta.InterconnectExpectedOutages{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Source:      ProtoToComputeBetaInterconnectExpectedOutagesSourceEnum(p.GetSource()),
		State:       ProtoToComputeBetaInterconnectExpectedOutagesStateEnum(p.GetState()),
		IssueType:   ProtoToComputeBetaInterconnectExpectedOutagesIssueTypeEnum(p.GetIssueType()),
		StartTime:   dcl.Int64OrNil(p.GetStartTime()),
		EndTime:     dcl.Int64OrNil(p.GetEndTime()),
	}
	for _, r := range p.GetAffectedCircuits() {
		obj.AffectedCircuits = append(obj.AffectedCircuits, r)
	}
	return obj
}

// ProtoToInterconnectCircuitInfos converts a InterconnectCircuitInfos object from its proto representation.
func ProtoToComputeBetaInterconnectCircuitInfos(p *betapb.ComputeBetaInterconnectCircuitInfos) *beta.InterconnectCircuitInfos {
	if p == nil {
		return nil
	}
	obj := &beta.InterconnectCircuitInfos{
		GoogleCircuitId:  dcl.StringOrNil(p.GetGoogleCircuitId()),
		GoogleDemarcId:   dcl.StringOrNil(p.GetGoogleDemarcId()),
		CustomerDemarcId: dcl.StringOrNil(p.GetCustomerDemarcId()),
	}
	return obj
}

// ProtoToInterconnect converts a Interconnect resource from its proto representation.
func ProtoToInterconnect(p *betapb.ComputeBetaInterconnect) *beta.Interconnect {
	obj := &beta.Interconnect{
		Description:          dcl.StringOrNil(p.GetDescription()),
		SelfLink:             dcl.StringOrNil(p.GetSelfLink()),
		Id:                   dcl.Int64OrNil(p.GetId()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		LinkType:             ProtoToComputeBetaInterconnectLinkTypeEnum(p.GetLinkType()),
		RequestedLinkCount:   dcl.Int64OrNil(p.GetRequestedLinkCount()),
		InterconnectType:     ProtoToComputeBetaInterconnectInterconnectTypeEnum(p.GetInterconnectType()),
		AdminEnabled:         dcl.Bool(p.GetAdminEnabled()),
		NocContactEmail:      dcl.StringOrNil(p.GetNocContactEmail()),
		CustomerName:         dcl.StringOrNil(p.GetCustomerName()),
		OperationalStatus:    ProtoToComputeBetaInterconnectOperationalStatusEnum(p.GetOperationalStatus()),
		ProvisionedLinkCount: dcl.Int64OrNil(p.GetProvisionedLinkCount()),
		PeerIPAddress:        dcl.StringOrNil(p.GetPeerIpAddress()),
		GoogleIPAddress:      dcl.StringOrNil(p.GetGoogleIpAddress()),
		GoogleReferenceId:    dcl.StringOrNil(p.GetGoogleReferenceId()),
		State:                ProtoToComputeBetaInterconnectStateEnum(p.GetState()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetInterconnectAttachments() {
		obj.InterconnectAttachments = append(obj.InterconnectAttachments, r)
	}
	for _, r := range p.GetExpectedOutages() {
		obj.ExpectedOutages = append(obj.ExpectedOutages, *ProtoToComputeBetaInterconnectExpectedOutages(r))
	}
	for _, r := range p.GetCircuitInfos() {
		obj.CircuitInfos = append(obj.CircuitInfos, *ProtoToComputeBetaInterconnectCircuitInfos(r))
	}
	return obj
}

// InterconnectLinkTypeEnumToProto converts a InterconnectLinkTypeEnum enum to its proto representation.
func ComputeBetaInterconnectLinkTypeEnumToProto(e *beta.InterconnectLinkTypeEnum) betapb.ComputeBetaInterconnectLinkTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectLinkTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectLinkTypeEnum_value["InterconnectLinkTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectLinkTypeEnum(v)
	}
	return betapb.ComputeBetaInterconnectLinkTypeEnum(0)
}

// InterconnectInterconnectTypeEnumToProto converts a InterconnectInterconnectTypeEnum enum to its proto representation.
func ComputeBetaInterconnectInterconnectTypeEnumToProto(e *beta.InterconnectInterconnectTypeEnum) betapb.ComputeBetaInterconnectInterconnectTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectInterconnectTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectInterconnectTypeEnum_value["InterconnectInterconnectTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectInterconnectTypeEnum(v)
	}
	return betapb.ComputeBetaInterconnectInterconnectTypeEnum(0)
}

// InterconnectOperationalStatusEnumToProto converts a InterconnectOperationalStatusEnum enum to its proto representation.
func ComputeBetaInterconnectOperationalStatusEnumToProto(e *beta.InterconnectOperationalStatusEnum) betapb.ComputeBetaInterconnectOperationalStatusEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectOperationalStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectOperationalStatusEnum_value["InterconnectOperationalStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectOperationalStatusEnum(v)
	}
	return betapb.ComputeBetaInterconnectOperationalStatusEnum(0)
}

// InterconnectExpectedOutagesSourceEnumToProto converts a InterconnectExpectedOutagesSourceEnum enum to its proto representation.
func ComputeBetaInterconnectExpectedOutagesSourceEnumToProto(e *beta.InterconnectExpectedOutagesSourceEnum) betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum_value["InterconnectExpectedOutagesSourceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum(v)
	}
	return betapb.ComputeBetaInterconnectExpectedOutagesSourceEnum(0)
}

// InterconnectExpectedOutagesStateEnumToProto converts a InterconnectExpectedOutagesStateEnum enum to its proto representation.
func ComputeBetaInterconnectExpectedOutagesStateEnumToProto(e *beta.InterconnectExpectedOutagesStateEnum) betapb.ComputeBetaInterconnectExpectedOutagesStateEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectExpectedOutagesStateEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectExpectedOutagesStateEnum_value["InterconnectExpectedOutagesStateEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectExpectedOutagesStateEnum(v)
	}
	return betapb.ComputeBetaInterconnectExpectedOutagesStateEnum(0)
}

// InterconnectExpectedOutagesIssueTypeEnumToProto converts a InterconnectExpectedOutagesIssueTypeEnum enum to its proto representation.
func ComputeBetaInterconnectExpectedOutagesIssueTypeEnumToProto(e *beta.InterconnectExpectedOutagesIssueTypeEnum) betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum_value["InterconnectExpectedOutagesIssueTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum(v)
	}
	return betapb.ComputeBetaInterconnectExpectedOutagesIssueTypeEnum(0)
}

// InterconnectStateEnumToProto converts a InterconnectStateEnum enum to its proto representation.
func ComputeBetaInterconnectStateEnumToProto(e *beta.InterconnectStateEnum) betapb.ComputeBetaInterconnectStateEnum {
	if e == nil {
		return betapb.ComputeBetaInterconnectStateEnum(0)
	}
	if v, ok := betapb.ComputeBetaInterconnectStateEnum_value["InterconnectStateEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInterconnectStateEnum(v)
	}
	return betapb.ComputeBetaInterconnectStateEnum(0)
}

// InterconnectExpectedOutagesToProto converts a InterconnectExpectedOutages object to its proto representation.
func ComputeBetaInterconnectExpectedOutagesToProto(o *beta.InterconnectExpectedOutages) *betapb.ComputeBetaInterconnectExpectedOutages {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInterconnectExpectedOutages{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetSource(ComputeBetaInterconnectExpectedOutagesSourceEnumToProto(o.Source))
	p.SetState(ComputeBetaInterconnectExpectedOutagesStateEnumToProto(o.State))
	p.SetIssueType(ComputeBetaInterconnectExpectedOutagesIssueTypeEnumToProto(o.IssueType))
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
func ComputeBetaInterconnectCircuitInfosToProto(o *beta.InterconnectCircuitInfos) *betapb.ComputeBetaInterconnectCircuitInfos {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInterconnectCircuitInfos{}
	p.SetGoogleCircuitId(dcl.ValueOrEmptyString(o.GoogleCircuitId))
	p.SetGoogleDemarcId(dcl.ValueOrEmptyString(o.GoogleDemarcId))
	p.SetCustomerDemarcId(dcl.ValueOrEmptyString(o.CustomerDemarcId))
	return p
}

// InterconnectToProto converts a Interconnect resource to its proto representation.
func InterconnectToProto(resource *beta.Interconnect) *betapb.ComputeBetaInterconnect {
	p := &betapb.ComputeBetaInterconnect{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetLinkType(ComputeBetaInterconnectLinkTypeEnumToProto(resource.LinkType))
	p.SetRequestedLinkCount(dcl.ValueOrEmptyInt64(resource.RequestedLinkCount))
	p.SetInterconnectType(ComputeBetaInterconnectInterconnectTypeEnumToProto(resource.InterconnectType))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetNocContactEmail(dcl.ValueOrEmptyString(resource.NocContactEmail))
	p.SetCustomerName(dcl.ValueOrEmptyString(resource.CustomerName))
	p.SetOperationalStatus(ComputeBetaInterconnectOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetProvisionedLinkCount(dcl.ValueOrEmptyInt64(resource.ProvisionedLinkCount))
	p.SetPeerIpAddress(dcl.ValueOrEmptyString(resource.PeerIPAddress))
	p.SetGoogleIpAddress(dcl.ValueOrEmptyString(resource.GoogleIPAddress))
	p.SetGoogleReferenceId(dcl.ValueOrEmptyString(resource.GoogleReferenceId))
	p.SetState(ComputeBetaInterconnectStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sInterconnectAttachments := make([]string, len(resource.InterconnectAttachments))
	for i, r := range resource.InterconnectAttachments {
		sInterconnectAttachments[i] = r
	}
	p.SetInterconnectAttachments(sInterconnectAttachments)
	sExpectedOutages := make([]*betapb.ComputeBetaInterconnectExpectedOutages, len(resource.ExpectedOutages))
	for i, r := range resource.ExpectedOutages {
		sExpectedOutages[i] = ComputeBetaInterconnectExpectedOutagesToProto(&r)
	}
	p.SetExpectedOutages(sExpectedOutages)
	sCircuitInfos := make([]*betapb.ComputeBetaInterconnectCircuitInfos, len(resource.CircuitInfos))
	for i, r := range resource.CircuitInfos {
		sCircuitInfos[i] = ComputeBetaInterconnectCircuitInfosToProto(&r)
	}
	p.SetCircuitInfos(sCircuitInfos)

	return p
}

// applyInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) applyInterconnect(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInterconnectRequest) (*betapb.ComputeBetaInterconnect, error) {
	p := ProtoToInterconnect(request.GetResource())
	res, err := c.ApplyInterconnect(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectToProto(res)
	return r, nil
}

// applyComputeBetaInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) ApplyComputeBetaInterconnect(ctx context.Context, request *betapb.ApplyComputeBetaInterconnectRequest) (*betapb.ComputeBetaInterconnect, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnect(ctx, cl, request)
}

// DeleteInterconnect handles the gRPC request by passing it to the underlying Interconnect Delete() method.
func (s *InterconnectServer) DeleteComputeBetaInterconnect(ctx context.Context, request *betapb.DeleteComputeBetaInterconnectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnect(ctx, ProtoToInterconnect(request.GetResource()))

}

// ListComputeBetaInterconnect handles the gRPC request by passing it to the underlying InterconnectList() method.
func (s *InterconnectServer) ListComputeBetaInterconnect(ctx context.Context, request *betapb.ListComputeBetaInterconnectRequest) (*betapb.ListComputeBetaInterconnectResponse, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnect(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInterconnect
	for _, r := range resources.Items {
		rp := InterconnectToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaInterconnectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnect(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
