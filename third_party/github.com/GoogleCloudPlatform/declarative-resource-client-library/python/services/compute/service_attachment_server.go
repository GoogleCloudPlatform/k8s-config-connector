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

// ServiceAttachmentServer implements the gRPC interface for ServiceAttachment.
type ServiceAttachmentServer struct{}

// ProtoToServiceAttachmentConnectionPreferenceEnum converts a ServiceAttachmentConnectionPreferenceEnum enum from its proto representation.
func ProtoToComputeServiceAttachmentConnectionPreferenceEnum(e computepb.ComputeServiceAttachmentConnectionPreferenceEnum) *compute.ServiceAttachmentConnectionPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeServiceAttachmentConnectionPreferenceEnum_name[int32(e)]; ok {
		e := compute.ServiceAttachmentConnectionPreferenceEnum(n[len("ComputeServiceAttachmentConnectionPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpointsStatusEnum converts a ServiceAttachmentConnectedEndpointsStatusEnum enum from its proto representation.
func ProtoToComputeServiceAttachmentConnectedEndpointsStatusEnum(e computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum) *compute.ServiceAttachmentConnectedEndpointsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum_name[int32(e)]; ok {
		e := compute.ServiceAttachmentConnectedEndpointsStatusEnum(n[len("ComputeServiceAttachmentConnectedEndpointsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpoints converts a ServiceAttachmentConnectedEndpoints object from its proto representation.
func ProtoToComputeServiceAttachmentConnectedEndpoints(p *computepb.ComputeServiceAttachmentConnectedEndpoints) *compute.ServiceAttachmentConnectedEndpoints {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentConnectedEndpoints{
		Status:          ProtoToComputeServiceAttachmentConnectedEndpointsStatusEnum(p.GetStatus()),
		PscConnectionId: dcl.Int64OrNil(p.GetPscConnectionId()),
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
	}
	return obj
}

// ProtoToServiceAttachmentConsumerAcceptLists converts a ServiceAttachmentConsumerAcceptLists object from its proto representation.
func ProtoToComputeServiceAttachmentConsumerAcceptLists(p *computepb.ComputeServiceAttachmentConsumerAcceptLists) *compute.ServiceAttachmentConsumerAcceptLists {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.StringOrNil(p.GetProjectIdOrNum()),
		ConnectionLimit: dcl.Int64OrNil(p.GetConnectionLimit()),
	}
	return obj
}

// ProtoToServiceAttachmentPscServiceAttachmentId converts a ServiceAttachmentPscServiceAttachmentId object from its proto representation.
func ProtoToComputeServiceAttachmentPscServiceAttachmentId(p *computepb.ComputeServiceAttachmentPscServiceAttachmentId) *compute.ServiceAttachmentPscServiceAttachmentId {
	if p == nil {
		return nil
	}
	obj := &compute.ServiceAttachmentPscServiceAttachmentId{
		High: dcl.Int64OrNil(p.GetHigh()),
		Low:  dcl.Int64OrNil(p.GetLow()),
	}
	return obj
}

// ProtoToServiceAttachment converts a ServiceAttachment resource from its proto representation.
func ProtoToServiceAttachment(p *computepb.ComputeServiceAttachment) *compute.ServiceAttachment {
	obj := &compute.ServiceAttachment{
		Id:                     dcl.Int64OrNil(p.GetId()),
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		SelfLink:               dcl.StringOrNil(p.GetSelfLink()),
		Region:                 dcl.StringOrNil(p.GetRegion()),
		TargetService:          dcl.StringOrNil(p.GetTargetService()),
		ConnectionPreference:   ProtoToComputeServiceAttachmentConnectionPreferenceEnum(p.GetConnectionPreference()),
		EnableProxyProtocol:    dcl.Bool(p.GetEnableProxyProtocol()),
		PscServiceAttachmentId: ProtoToComputeServiceAttachmentPscServiceAttachmentId(p.GetPscServiceAttachmentId()),
		Fingerprint:            dcl.StringOrNil(p.GetFingerprint()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetConnectedEndpoints() {
		obj.ConnectedEndpoints = append(obj.ConnectedEndpoints, *ProtoToComputeServiceAttachmentConnectedEndpoints(r))
	}
	for _, r := range p.GetNatSubnets() {
		obj.NatSubnets = append(obj.NatSubnets, r)
	}
	for _, r := range p.GetConsumerRejectLists() {
		obj.ConsumerRejectLists = append(obj.ConsumerRejectLists, r)
	}
	for _, r := range p.GetConsumerAcceptLists() {
		obj.ConsumerAcceptLists = append(obj.ConsumerAcceptLists, *ProtoToComputeServiceAttachmentConsumerAcceptLists(r))
	}
	return obj
}

// ServiceAttachmentConnectionPreferenceEnumToProto converts a ServiceAttachmentConnectionPreferenceEnum enum to its proto representation.
func ComputeServiceAttachmentConnectionPreferenceEnumToProto(e *compute.ServiceAttachmentConnectionPreferenceEnum) computepb.ComputeServiceAttachmentConnectionPreferenceEnum {
	if e == nil {
		return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(0)
	}
	if v, ok := computepb.ComputeServiceAttachmentConnectionPreferenceEnum_value["ServiceAttachmentConnectionPreferenceEnum"+string(*e)]; ok {
		return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(v)
	}
	return computepb.ComputeServiceAttachmentConnectionPreferenceEnum(0)
}

// ServiceAttachmentConnectedEndpointsStatusEnumToProto converts a ServiceAttachmentConnectedEndpointsStatusEnum enum to its proto representation.
func ComputeServiceAttachmentConnectedEndpointsStatusEnumToProto(e *compute.ServiceAttachmentConnectedEndpointsStatusEnum) computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum {
	if e == nil {
		return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(0)
	}
	if v, ok := computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum_value["ServiceAttachmentConnectedEndpointsStatusEnum"+string(*e)]; ok {
		return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(v)
	}
	return computepb.ComputeServiceAttachmentConnectedEndpointsStatusEnum(0)
}

// ServiceAttachmentConnectedEndpointsToProto converts a ServiceAttachmentConnectedEndpoints object to its proto representation.
func ComputeServiceAttachmentConnectedEndpointsToProto(o *compute.ServiceAttachmentConnectedEndpoints) *computepb.ComputeServiceAttachmentConnectedEndpoints {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentConnectedEndpoints{}
	p.SetStatus(ComputeServiceAttachmentConnectedEndpointsStatusEnumToProto(o.Status))
	p.SetPscConnectionId(dcl.ValueOrEmptyInt64(o.PscConnectionId))
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	return p
}

// ServiceAttachmentConsumerAcceptListsToProto converts a ServiceAttachmentConsumerAcceptLists object to its proto representation.
func ComputeServiceAttachmentConsumerAcceptListsToProto(o *compute.ServiceAttachmentConsumerAcceptLists) *computepb.ComputeServiceAttachmentConsumerAcceptLists {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentConsumerAcceptLists{}
	p.SetProjectIdOrNum(dcl.ValueOrEmptyString(o.ProjectIdOrNum))
	p.SetConnectionLimit(dcl.ValueOrEmptyInt64(o.ConnectionLimit))
	return p
}

// ServiceAttachmentPscServiceAttachmentIdToProto converts a ServiceAttachmentPscServiceAttachmentId object to its proto representation.
func ComputeServiceAttachmentPscServiceAttachmentIdToProto(o *compute.ServiceAttachmentPscServiceAttachmentId) *computepb.ComputeServiceAttachmentPscServiceAttachmentId {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeServiceAttachmentPscServiceAttachmentId{}
	p.SetHigh(dcl.ValueOrEmptyInt64(o.High))
	p.SetLow(dcl.ValueOrEmptyInt64(o.Low))
	return p
}

// ServiceAttachmentToProto converts a ServiceAttachment resource to its proto representation.
func ServiceAttachmentToProto(resource *compute.ServiceAttachment) *computepb.ComputeServiceAttachment {
	p := &computepb.ComputeServiceAttachment{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetTargetService(dcl.ValueOrEmptyString(resource.TargetService))
	p.SetConnectionPreference(ComputeServiceAttachmentConnectionPreferenceEnumToProto(resource.ConnectionPreference))
	p.SetEnableProxyProtocol(dcl.ValueOrEmptyBool(resource.EnableProxyProtocol))
	p.SetPscServiceAttachmentId(ComputeServiceAttachmentPscServiceAttachmentIdToProto(resource.PscServiceAttachmentId))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sConnectedEndpoints := make([]*computepb.ComputeServiceAttachmentConnectedEndpoints, len(resource.ConnectedEndpoints))
	for i, r := range resource.ConnectedEndpoints {
		sConnectedEndpoints[i] = ComputeServiceAttachmentConnectedEndpointsToProto(&r)
	}
	p.SetConnectedEndpoints(sConnectedEndpoints)
	sNatSubnets := make([]string, len(resource.NatSubnets))
	for i, r := range resource.NatSubnets {
		sNatSubnets[i] = r
	}
	p.SetNatSubnets(sNatSubnets)
	sConsumerRejectLists := make([]string, len(resource.ConsumerRejectLists))
	for i, r := range resource.ConsumerRejectLists {
		sConsumerRejectLists[i] = r
	}
	p.SetConsumerRejectLists(sConsumerRejectLists)
	sConsumerAcceptLists := make([]*computepb.ComputeServiceAttachmentConsumerAcceptLists, len(resource.ConsumerAcceptLists))
	for i, r := range resource.ConsumerAcceptLists {
		sConsumerAcceptLists[i] = ComputeServiceAttachmentConsumerAcceptListsToProto(&r)
	}
	p.SetConsumerAcceptLists(sConsumerAcceptLists)

	return p
}

// applyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) applyServiceAttachment(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeServiceAttachmentRequest) (*computepb.ComputeServiceAttachment, error) {
	p := ProtoToServiceAttachment(request.GetResource())
	res, err := c.ApplyServiceAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAttachmentToProto(res)
	return r, nil
}

// applyComputeServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) ApplyComputeServiceAttachment(ctx context.Context, request *computepb.ApplyComputeServiceAttachmentRequest) (*computepb.ComputeServiceAttachment, error) {
	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceAttachment(ctx, cl, request)
}

// DeleteServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Delete() method.
func (s *ServiceAttachmentServer) DeleteComputeServiceAttachment(ctx context.Context, request *computepb.DeleteComputeServiceAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))

}

// ListComputeServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachmentList() method.
func (s *ServiceAttachmentServer) ListComputeServiceAttachment(ctx context.Context, request *computepb.ListComputeServiceAttachmentRequest) (*computepb.ListComputeServiceAttachmentResponse, error) {
	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAttachment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeServiceAttachment
	for _, r := range resources.Items {
		rp := ServiceAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeServiceAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceAttachment(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
