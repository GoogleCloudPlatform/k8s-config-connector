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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Interconnect.
type InterconnectServer struct{}

// ProtoToInterconnectLinkTypeEnum converts a InterconnectLinkTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectLinkTypeEnum(e computepb.ComputeInterconnectLinkTypeEnum) *compute.InterconnectLinkTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectLinkTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectLinkTypeEnum(n[len("ComputeInterconnectLinkTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectInterconnectTypeEnum converts a InterconnectInterconnectTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectInterconnectTypeEnum(e computepb.ComputeInterconnectInterconnectTypeEnum) *compute.InterconnectInterconnectTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectInterconnectTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectInterconnectTypeEnum(n[len("ComputeInterconnectInterconnectTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectOperationalStatusEnum converts a InterconnectOperationalStatusEnum enum from its proto representation.
func ProtoToComputeInterconnectOperationalStatusEnum(e computepb.ComputeInterconnectOperationalStatusEnum) *compute.InterconnectOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectOperationalStatusEnum_name[int32(e)]; ok {
		e := compute.InterconnectOperationalStatusEnum(n[len("ComputeInterconnectOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesSourceEnum converts a InterconnectExpectedOutagesSourceEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesSourceEnum(e computepb.ComputeInterconnectExpectedOutagesSourceEnum) *compute.InterconnectExpectedOutagesSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesSourceEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesSourceEnum(n[len("ComputeInterconnectExpectedOutagesSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesStateEnum converts a InterconnectExpectedOutagesStateEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesStateEnum(e computepb.ComputeInterconnectExpectedOutagesStateEnum) *compute.InterconnectExpectedOutagesStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesStateEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesStateEnum(n[len("ComputeInterconnectExpectedOutagesStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesIssueTypeEnum converts a InterconnectExpectedOutagesIssueTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesIssueTypeEnum(e computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum) *compute.InterconnectExpectedOutagesIssueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesIssueTypeEnum(n[len("ComputeInterconnectExpectedOutagesIssueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectStateEnum converts a InterconnectStateEnum enum from its proto representation.
func ProtoToComputeInterconnectStateEnum(e computepb.ComputeInterconnectStateEnum) *compute.InterconnectStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectStateEnum_name[int32(e)]; ok {
		e := compute.InterconnectStateEnum(n[len("ComputeInterconnectStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutages converts a InterconnectExpectedOutages object from its proto representation.
func ProtoToComputeInterconnectExpectedOutages(p *computepb.ComputeInterconnectExpectedOutages) *compute.InterconnectExpectedOutages {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectExpectedOutages{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Source:      ProtoToComputeInterconnectExpectedOutagesSourceEnum(p.GetSource()),
		State:       ProtoToComputeInterconnectExpectedOutagesStateEnum(p.GetState()),
		IssueType:   ProtoToComputeInterconnectExpectedOutagesIssueTypeEnum(p.GetIssueType()),
		StartTime:   dcl.Int64OrNil(p.GetStartTime()),
		EndTime:     dcl.Int64OrNil(p.GetEndTime()),
	}
	for _, r := range p.GetAffectedCircuits() {
		obj.AffectedCircuits = append(obj.AffectedCircuits, r)
	}
	return obj
}

// ProtoToInterconnectCircuitInfos converts a InterconnectCircuitInfos object from its proto representation.
func ProtoToComputeInterconnectCircuitInfos(p *computepb.ComputeInterconnectCircuitInfos) *compute.InterconnectCircuitInfos {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectCircuitInfos{
		GoogleCircuitId:  dcl.StringOrNil(p.GetGoogleCircuitId()),
		GoogleDemarcId:   dcl.StringOrNil(p.GetGoogleDemarcId()),
		CustomerDemarcId: dcl.StringOrNil(p.GetCustomerDemarcId()),
	}
	return obj
}

// ProtoToInterconnect converts a Interconnect resource from its proto representation.
func ProtoToInterconnect(p *computepb.ComputeInterconnect) *compute.Interconnect {
	obj := &compute.Interconnect{
		Description:          dcl.StringOrNil(p.GetDescription()),
		SelfLink:             dcl.StringOrNil(p.GetSelfLink()),
		Id:                   dcl.Int64OrNil(p.GetId()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		LinkType:             ProtoToComputeInterconnectLinkTypeEnum(p.GetLinkType()),
		RequestedLinkCount:   dcl.Int64OrNil(p.GetRequestedLinkCount()),
		InterconnectType:     ProtoToComputeInterconnectInterconnectTypeEnum(p.GetInterconnectType()),
		AdminEnabled:         dcl.Bool(p.GetAdminEnabled()),
		NocContactEmail:      dcl.StringOrNil(p.GetNocContactEmail()),
		CustomerName:         dcl.StringOrNil(p.GetCustomerName()),
		OperationalStatus:    ProtoToComputeInterconnectOperationalStatusEnum(p.GetOperationalStatus()),
		ProvisionedLinkCount: dcl.Int64OrNil(p.GetProvisionedLinkCount()),
		PeerIPAddress:        dcl.StringOrNil(p.GetPeerIpAddress()),
		GoogleIPAddress:      dcl.StringOrNil(p.GetGoogleIpAddress()),
		GoogleReferenceId:    dcl.StringOrNil(p.GetGoogleReferenceId()),
		State:                ProtoToComputeInterconnectStateEnum(p.GetState()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetInterconnectAttachments() {
		obj.InterconnectAttachments = append(obj.InterconnectAttachments, r)
	}
	for _, r := range p.GetExpectedOutages() {
		obj.ExpectedOutages = append(obj.ExpectedOutages, *ProtoToComputeInterconnectExpectedOutages(r))
	}
	for _, r := range p.GetCircuitInfos() {
		obj.CircuitInfos = append(obj.CircuitInfos, *ProtoToComputeInterconnectCircuitInfos(r))
	}
	return obj
}

// InterconnectLinkTypeEnumToProto converts a InterconnectLinkTypeEnum enum to its proto representation.
func ComputeInterconnectLinkTypeEnumToProto(e *compute.InterconnectLinkTypeEnum) computepb.ComputeInterconnectLinkTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectLinkTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectLinkTypeEnum_value["InterconnectLinkTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectLinkTypeEnum(v)
	}
	return computepb.ComputeInterconnectLinkTypeEnum(0)
}

// InterconnectInterconnectTypeEnumToProto converts a InterconnectInterconnectTypeEnum enum to its proto representation.
func ComputeInterconnectInterconnectTypeEnumToProto(e *compute.InterconnectInterconnectTypeEnum) computepb.ComputeInterconnectInterconnectTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectInterconnectTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectInterconnectTypeEnum_value["InterconnectInterconnectTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectInterconnectTypeEnum(v)
	}
	return computepb.ComputeInterconnectInterconnectTypeEnum(0)
}

// InterconnectOperationalStatusEnumToProto converts a InterconnectOperationalStatusEnum enum to its proto representation.
func ComputeInterconnectOperationalStatusEnumToProto(e *compute.InterconnectOperationalStatusEnum) computepb.ComputeInterconnectOperationalStatusEnum {
	if e == nil {
		return computepb.ComputeInterconnectOperationalStatusEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectOperationalStatusEnum_value["InterconnectOperationalStatusEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectOperationalStatusEnum(v)
	}
	return computepb.ComputeInterconnectOperationalStatusEnum(0)
}

// InterconnectExpectedOutagesSourceEnumToProto converts a InterconnectExpectedOutagesSourceEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesSourceEnumToProto(e *compute.InterconnectExpectedOutagesSourceEnum) computepb.ComputeInterconnectExpectedOutagesSourceEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesSourceEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesSourceEnum_value["InterconnectExpectedOutagesSourceEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesSourceEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesSourceEnum(0)
}

// InterconnectExpectedOutagesStateEnumToProto converts a InterconnectExpectedOutagesStateEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesStateEnumToProto(e *compute.InterconnectExpectedOutagesStateEnum) computepb.ComputeInterconnectExpectedOutagesStateEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesStateEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesStateEnum_value["InterconnectExpectedOutagesStateEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesStateEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesStateEnum(0)
}

// InterconnectExpectedOutagesIssueTypeEnumToProto converts a InterconnectExpectedOutagesIssueTypeEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesIssueTypeEnumToProto(e *compute.InterconnectExpectedOutagesIssueTypeEnum) computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum_value["InterconnectExpectedOutagesIssueTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(0)
}

// InterconnectStateEnumToProto converts a InterconnectStateEnum enum to its proto representation.
func ComputeInterconnectStateEnumToProto(e *compute.InterconnectStateEnum) computepb.ComputeInterconnectStateEnum {
	if e == nil {
		return computepb.ComputeInterconnectStateEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectStateEnum_value["InterconnectStateEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectStateEnum(v)
	}
	return computepb.ComputeInterconnectStateEnum(0)
}

// InterconnectExpectedOutagesToProto converts a InterconnectExpectedOutages object to its proto representation.
func ComputeInterconnectExpectedOutagesToProto(o *compute.InterconnectExpectedOutages) *computepb.ComputeInterconnectExpectedOutages {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectExpectedOutages{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetSource(ComputeInterconnectExpectedOutagesSourceEnumToProto(o.Source))
	p.SetState(ComputeInterconnectExpectedOutagesStateEnumToProto(o.State))
	p.SetIssueType(ComputeInterconnectExpectedOutagesIssueTypeEnumToProto(o.IssueType))
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
func ComputeInterconnectCircuitInfosToProto(o *compute.InterconnectCircuitInfos) *computepb.ComputeInterconnectCircuitInfos {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectCircuitInfos{}
	p.SetGoogleCircuitId(dcl.ValueOrEmptyString(o.GoogleCircuitId))
	p.SetGoogleDemarcId(dcl.ValueOrEmptyString(o.GoogleDemarcId))
	p.SetCustomerDemarcId(dcl.ValueOrEmptyString(o.CustomerDemarcId))
	return p
}

// InterconnectToProto converts a Interconnect resource to its proto representation.
func InterconnectToProto(resource *compute.Interconnect) *computepb.ComputeInterconnect {
	p := &computepb.ComputeInterconnect{}
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetLinkType(ComputeInterconnectLinkTypeEnumToProto(resource.LinkType))
	p.SetRequestedLinkCount(dcl.ValueOrEmptyInt64(resource.RequestedLinkCount))
	p.SetInterconnectType(ComputeInterconnectInterconnectTypeEnumToProto(resource.InterconnectType))
	p.SetAdminEnabled(dcl.ValueOrEmptyBool(resource.AdminEnabled))
	p.SetNocContactEmail(dcl.ValueOrEmptyString(resource.NocContactEmail))
	p.SetCustomerName(dcl.ValueOrEmptyString(resource.CustomerName))
	p.SetOperationalStatus(ComputeInterconnectOperationalStatusEnumToProto(resource.OperationalStatus))
	p.SetProvisionedLinkCount(dcl.ValueOrEmptyInt64(resource.ProvisionedLinkCount))
	p.SetPeerIpAddress(dcl.ValueOrEmptyString(resource.PeerIPAddress))
	p.SetGoogleIpAddress(dcl.ValueOrEmptyString(resource.GoogleIPAddress))
	p.SetGoogleReferenceId(dcl.ValueOrEmptyString(resource.GoogleReferenceId))
	p.SetState(ComputeInterconnectStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sInterconnectAttachments := make([]string, len(resource.InterconnectAttachments))
	for i, r := range resource.InterconnectAttachments {
		sInterconnectAttachments[i] = r
	}
	p.SetInterconnectAttachments(sInterconnectAttachments)
	sExpectedOutages := make([]*computepb.ComputeInterconnectExpectedOutages, len(resource.ExpectedOutages))
	for i, r := range resource.ExpectedOutages {
		sExpectedOutages[i] = ComputeInterconnectExpectedOutagesToProto(&r)
	}
	p.SetExpectedOutages(sExpectedOutages)
	sCircuitInfos := make([]*computepb.ComputeInterconnectCircuitInfos, len(resource.CircuitInfos))
	for i, r := range resource.CircuitInfos {
		sCircuitInfos[i] = ComputeInterconnectCircuitInfosToProto(&r)
	}
	p.SetCircuitInfos(sCircuitInfos)

	return p
}

// applyInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) applyInterconnect(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInterconnectRequest) (*computepb.ComputeInterconnect, error) {
	p := ProtoToInterconnect(request.GetResource())
	res, err := c.ApplyInterconnect(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectToProto(res)
	return r, nil
}

// applyComputeInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) ApplyComputeInterconnect(ctx context.Context, request *computepb.ApplyComputeInterconnectRequest) (*computepb.ComputeInterconnect, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInterconnect(ctx, cl, request)
}

// DeleteInterconnect handles the gRPC request by passing it to the underlying Interconnect Delete() method.
func (s *InterconnectServer) DeleteComputeInterconnect(ctx context.Context, request *computepb.DeleteComputeInterconnectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnect(ctx, ProtoToInterconnect(request.GetResource()))

}

// ListComputeInterconnect handles the gRPC request by passing it to the underlying InterconnectList() method.
func (s *InterconnectServer) ListComputeInterconnect(ctx context.Context, request *computepb.ListComputeInterconnectRequest) (*computepb.ListComputeInterconnectResponse, error) {
	cl, err := createConfigInterconnect(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnect(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInterconnect
	for _, r := range resources.Items {
		rp := InterconnectToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeInterconnectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInterconnect(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
