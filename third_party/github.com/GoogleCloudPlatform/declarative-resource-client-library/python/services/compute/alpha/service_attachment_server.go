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

// ServiceAttachmentServer implements the gRPC interface for ServiceAttachment.
type ServiceAttachmentServer struct{}

// ProtoToServiceAttachmentConnectionPreferenceEnum converts a ServiceAttachmentConnectionPreferenceEnum enum from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectionPreferenceEnum(e alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum) *alpha.ServiceAttachmentConnectionPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum_name[int32(e)]; ok {
		e := alpha.ServiceAttachmentConnectionPreferenceEnum(n[len("ComputeAlphaServiceAttachmentConnectionPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpointsStatusEnum converts a ServiceAttachmentConnectedEndpointsStatusEnum enum from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(e alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum) *alpha.ServiceAttachmentConnectedEndpointsStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum_name[int32(e)]; ok {
		e := alpha.ServiceAttachmentConnectedEndpointsStatusEnum(n[len("ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceAttachmentConnectedEndpoints converts a ServiceAttachmentConnectedEndpoints object from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConnectedEndpoints(p *alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints) *alpha.ServiceAttachmentConnectedEndpoints {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentConnectedEndpoints{
		Status:          ProtoToComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(p.GetStatus()),
		PscConnectionId: dcl.Int64OrNil(p.GetPscConnectionId()),
		Endpoint:        dcl.StringOrNil(p.GetEndpoint()),
	}
	return obj
}

// ProtoToServiceAttachmentConsumerAcceptLists converts a ServiceAttachmentConsumerAcceptLists object from its proto representation.
func ProtoToComputeAlphaServiceAttachmentConsumerAcceptLists(p *alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists) *alpha.ServiceAttachmentConsumerAcceptLists {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentConsumerAcceptLists{
		ProjectIdOrNum:  dcl.StringOrNil(p.GetProjectIdOrNum()),
		ConnectionLimit: dcl.Int64OrNil(p.GetConnectionLimit()),
	}
	return obj
}

// ProtoToServiceAttachmentPscServiceAttachmentId converts a ServiceAttachmentPscServiceAttachmentId object from its proto representation.
func ProtoToComputeAlphaServiceAttachmentPscServiceAttachmentId(p *alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId) *alpha.ServiceAttachmentPscServiceAttachmentId {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceAttachmentPscServiceAttachmentId{
		High: dcl.Int64OrNil(p.GetHigh()),
		Low:  dcl.Int64OrNil(p.GetLow()),
	}
	return obj
}

// ProtoToServiceAttachment converts a ServiceAttachment resource from its proto representation.
func ProtoToServiceAttachment(p *alphapb.ComputeAlphaServiceAttachment) *alpha.ServiceAttachment {
	obj := &alpha.ServiceAttachment{
		Id:                     dcl.Int64OrNil(p.GetId()),
		Name:                   dcl.StringOrNil(p.GetName()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		SelfLink:               dcl.StringOrNil(p.GetSelfLink()),
		Region:                 dcl.StringOrNil(p.GetRegion()),
		TargetService:          dcl.StringOrNil(p.GetTargetService()),
		ConnectionPreference:   ProtoToComputeAlphaServiceAttachmentConnectionPreferenceEnum(p.GetConnectionPreference()),
		EnableProxyProtocol:    dcl.Bool(p.GetEnableProxyProtocol()),
		PscServiceAttachmentId: ProtoToComputeAlphaServiceAttachmentPscServiceAttachmentId(p.GetPscServiceAttachmentId()),
		Fingerprint:            dcl.StringOrNil(p.GetFingerprint()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Location:               dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetConnectedEndpoints() {
		obj.ConnectedEndpoints = append(obj.ConnectedEndpoints, *ProtoToComputeAlphaServiceAttachmentConnectedEndpoints(r))
	}
	for _, r := range p.GetNatSubnets() {
		obj.NatSubnets = append(obj.NatSubnets, r)
	}
	for _, r := range p.GetConsumerRejectLists() {
		obj.ConsumerRejectLists = append(obj.ConsumerRejectLists, r)
	}
	for _, r := range p.GetConsumerAcceptLists() {
		obj.ConsumerAcceptLists = append(obj.ConsumerAcceptLists, *ProtoToComputeAlphaServiceAttachmentConsumerAcceptLists(r))
	}
	return obj
}

// ServiceAttachmentConnectionPreferenceEnumToProto converts a ServiceAttachmentConnectionPreferenceEnum enum to its proto representation.
func ComputeAlphaServiceAttachmentConnectionPreferenceEnumToProto(e *alpha.ServiceAttachmentConnectionPreferenceEnum) alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum {
	if e == nil {
		return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum_value["ServiceAttachmentConnectionPreferenceEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(v)
	}
	return alphapb.ComputeAlphaServiceAttachmentConnectionPreferenceEnum(0)
}

// ServiceAttachmentConnectedEndpointsStatusEnumToProto converts a ServiceAttachmentConnectedEndpointsStatusEnum enum to its proto representation.
func ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnumToProto(e *alpha.ServiceAttachmentConnectedEndpointsStatusEnum) alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum_value["ServiceAttachmentConnectedEndpointsStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(v)
	}
	return alphapb.ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnum(0)
}

// ServiceAttachmentConnectedEndpointsToProto converts a ServiceAttachmentConnectedEndpoints object to its proto representation.
func ComputeAlphaServiceAttachmentConnectedEndpointsToProto(o *alpha.ServiceAttachmentConnectedEndpoints) *alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints{}
	p.SetStatus(ComputeAlphaServiceAttachmentConnectedEndpointsStatusEnumToProto(o.Status))
	p.SetPscConnectionId(dcl.ValueOrEmptyInt64(o.PscConnectionId))
	p.SetEndpoint(dcl.ValueOrEmptyString(o.Endpoint))
	return p
}

// ServiceAttachmentConsumerAcceptListsToProto converts a ServiceAttachmentConsumerAcceptLists object to its proto representation.
func ComputeAlphaServiceAttachmentConsumerAcceptListsToProto(o *alpha.ServiceAttachmentConsumerAcceptLists) *alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists{}
	p.SetProjectIdOrNum(dcl.ValueOrEmptyString(o.ProjectIdOrNum))
	p.SetConnectionLimit(dcl.ValueOrEmptyInt64(o.ConnectionLimit))
	return p
}

// ServiceAttachmentPscServiceAttachmentIdToProto converts a ServiceAttachmentPscServiceAttachmentId object to its proto representation.
func ComputeAlphaServiceAttachmentPscServiceAttachmentIdToProto(o *alpha.ServiceAttachmentPscServiceAttachmentId) *alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaServiceAttachmentPscServiceAttachmentId{}
	p.SetHigh(dcl.ValueOrEmptyInt64(o.High))
	p.SetLow(dcl.ValueOrEmptyInt64(o.Low))
	return p
}

// ServiceAttachmentToProto converts a ServiceAttachment resource to its proto representation.
func ServiceAttachmentToProto(resource *alpha.ServiceAttachment) *alphapb.ComputeAlphaServiceAttachment {
	p := &alphapb.ComputeAlphaServiceAttachment{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetTargetService(dcl.ValueOrEmptyString(resource.TargetService))
	p.SetConnectionPreference(ComputeAlphaServiceAttachmentConnectionPreferenceEnumToProto(resource.ConnectionPreference))
	p.SetEnableProxyProtocol(dcl.ValueOrEmptyBool(resource.EnableProxyProtocol))
	p.SetPscServiceAttachmentId(ComputeAlphaServiceAttachmentPscServiceAttachmentIdToProto(resource.PscServiceAttachmentId))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sConnectedEndpoints := make([]*alphapb.ComputeAlphaServiceAttachmentConnectedEndpoints, len(resource.ConnectedEndpoints))
	for i, r := range resource.ConnectedEndpoints {
		sConnectedEndpoints[i] = ComputeAlphaServiceAttachmentConnectedEndpointsToProto(&r)
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
	sConsumerAcceptLists := make([]*alphapb.ComputeAlphaServiceAttachmentConsumerAcceptLists, len(resource.ConsumerAcceptLists))
	for i, r := range resource.ConsumerAcceptLists {
		sConsumerAcceptLists[i] = ComputeAlphaServiceAttachmentConsumerAcceptListsToProto(&r)
	}
	p.SetConsumerAcceptLists(sConsumerAcceptLists)

	return p
}

// applyServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) applyServiceAttachment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaServiceAttachmentRequest) (*alphapb.ComputeAlphaServiceAttachment, error) {
	p := ProtoToServiceAttachment(request.GetResource())
	res, err := c.ApplyServiceAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceAttachmentToProto(res)
	return r, nil
}

// applyComputeAlphaServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Apply() method.
func (s *ServiceAttachmentServer) ApplyComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.ApplyComputeAlphaServiceAttachmentRequest) (*alphapb.ComputeAlphaServiceAttachment, error) {
	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceAttachment(ctx, cl, request)
}

// DeleteServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachment Delete() method.
func (s *ServiceAttachmentServer) DeleteComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.DeleteComputeAlphaServiceAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceAttachment(ctx, ProtoToServiceAttachment(request.GetResource()))

}

// ListComputeAlphaServiceAttachment handles the gRPC request by passing it to the underlying ServiceAttachmentList() method.
func (s *ServiceAttachmentServer) ListComputeAlphaServiceAttachment(ctx context.Context, request *alphapb.ListComputeAlphaServiceAttachmentRequest) (*alphapb.ListComputeAlphaServiceAttachmentResponse, error) {
	cl, err := createConfigServiceAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceAttachment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaServiceAttachment
	for _, r := range resources.Items {
		rp := ServiceAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaServiceAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceAttachment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
