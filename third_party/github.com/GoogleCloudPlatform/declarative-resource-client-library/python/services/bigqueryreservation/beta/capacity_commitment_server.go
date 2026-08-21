// Copyright 2022 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/beta/bigqueryreservation_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
)

// CapacityCommitmentServer implements the gRPC interface for CapacityCommitment.
type CapacityCommitmentServer struct{}

// ProtoToCapacityCommitmentPlanEnum converts a CapacityCommitmentPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationBetaCapacityCommitmentPlanEnum(e betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum) *beta.CapacityCommitmentPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum_name[int32(e)]; ok {
		e := beta.CapacityCommitmentPlanEnum(n[len("BigqueryreservationBetaCapacityCommitmentPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentStateEnum converts a CapacityCommitmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationBetaCapacityCommitmentStateEnum(e betapb.BigqueryreservationBetaCapacityCommitmentStateEnum) *beta.CapacityCommitmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryreservationBetaCapacityCommitmentStateEnum_name[int32(e)]; ok {
		e := beta.CapacityCommitmentStateEnum(n[len("BigqueryreservationBetaCapacityCommitmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentRenewalPlanEnum converts a CapacityCommitmentRenewalPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationBetaCapacityCommitmentRenewalPlanEnum(e betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum) *beta.CapacityCommitmentRenewalPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum_name[int32(e)]; ok {
		e := beta.CapacityCommitmentRenewalPlanEnum(n[len("BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentFailureStatus converts a CapacityCommitmentFailureStatus object from its proto representation.
func ProtoToBigqueryreservationBetaCapacityCommitmentFailureStatus(p *betapb.BigqueryreservationBetaCapacityCommitmentFailureStatus) *beta.CapacityCommitmentFailureStatus {
	if p == nil {
		return nil
	}
	obj := &beta.CapacityCommitmentFailureStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToBigqueryreservationBetaCapacityCommitmentFailureStatusDetails(r))
	}
	return obj
}

// ProtoToCapacityCommitmentFailureStatusDetails converts a CapacityCommitmentFailureStatusDetails object from its proto representation.
func ProtoToBigqueryreservationBetaCapacityCommitmentFailureStatusDetails(p *betapb.BigqueryreservationBetaCapacityCommitmentFailureStatusDetails) *beta.CapacityCommitmentFailureStatusDetails {
	if p == nil {
		return nil
	}
	obj := &beta.CapacityCommitmentFailureStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCapacityCommitment converts a CapacityCommitment resource from its proto representation.
func ProtoToCapacityCommitment(p *betapb.BigqueryreservationBetaCapacityCommitment) *beta.CapacityCommitment {
	obj := &beta.CapacityCommitment{
		Name:                dcl.StringOrNil(p.GetName()),
		SlotCount:           dcl.Int64OrNil(p.GetSlotCount()),
		Plan:                ProtoToBigqueryreservationBetaCapacityCommitmentPlanEnum(p.GetPlan()),
		State:               ProtoToBigqueryreservationBetaCapacityCommitmentStateEnum(p.GetState()),
		CommitmentStartTime: dcl.StringOrNil(p.GetCommitmentStartTime()),
		CommitmentEndTime:   dcl.StringOrNil(p.GetCommitmentEndTime()),
		FailureStatus:       ProtoToBigqueryreservationBetaCapacityCommitmentFailureStatus(p.GetFailureStatus()),
		RenewalPlan:         ProtoToBigqueryreservationBetaCapacityCommitmentRenewalPlanEnum(p.GetRenewalPlan()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CapacityCommitmentPlanEnumToProto converts a CapacityCommitmentPlanEnum enum to its proto representation.
func BigqueryreservationBetaCapacityCommitmentPlanEnumToProto(e *beta.CapacityCommitmentPlanEnum) betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum {
	if e == nil {
		return betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum(0)
	}
	if v, ok := betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum_value["CapacityCommitmentPlanEnum"+string(*e)]; ok {
		return betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum(v)
	}
	return betapb.BigqueryreservationBetaCapacityCommitmentPlanEnum(0)
}

// CapacityCommitmentStateEnumToProto converts a CapacityCommitmentStateEnum enum to its proto representation.
func BigqueryreservationBetaCapacityCommitmentStateEnumToProto(e *beta.CapacityCommitmentStateEnum) betapb.BigqueryreservationBetaCapacityCommitmentStateEnum {
	if e == nil {
		return betapb.BigqueryreservationBetaCapacityCommitmentStateEnum(0)
	}
	if v, ok := betapb.BigqueryreservationBetaCapacityCommitmentStateEnum_value["CapacityCommitmentStateEnum"+string(*e)]; ok {
		return betapb.BigqueryreservationBetaCapacityCommitmentStateEnum(v)
	}
	return betapb.BigqueryreservationBetaCapacityCommitmentStateEnum(0)
}

// CapacityCommitmentRenewalPlanEnumToProto converts a CapacityCommitmentRenewalPlanEnum enum to its proto representation.
func BigqueryreservationBetaCapacityCommitmentRenewalPlanEnumToProto(e *beta.CapacityCommitmentRenewalPlanEnum) betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum {
	if e == nil {
		return betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum(0)
	}
	if v, ok := betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum_value["CapacityCommitmentRenewalPlanEnum"+string(*e)]; ok {
		return betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum(v)
	}
	return betapb.BigqueryreservationBetaCapacityCommitmentRenewalPlanEnum(0)
}

// CapacityCommitmentFailureStatusToProto converts a CapacityCommitmentFailureStatus object to its proto representation.
func BigqueryreservationBetaCapacityCommitmentFailureStatusToProto(o *beta.CapacityCommitmentFailureStatus) *betapb.BigqueryreservationBetaCapacityCommitmentFailureStatus {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryreservationBetaCapacityCommitmentFailureStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*betapb.BigqueryreservationBetaCapacityCommitmentFailureStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = BigqueryreservationBetaCapacityCommitmentFailureStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// CapacityCommitmentFailureStatusDetailsToProto converts a CapacityCommitmentFailureStatusDetails object to its proto representation.
func BigqueryreservationBetaCapacityCommitmentFailureStatusDetailsToProto(o *beta.CapacityCommitmentFailureStatusDetails) *betapb.BigqueryreservationBetaCapacityCommitmentFailureStatusDetails {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryreservationBetaCapacityCommitmentFailureStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CapacityCommitmentToProto converts a CapacityCommitment resource to its proto representation.
func CapacityCommitmentToProto(resource *beta.CapacityCommitment) *betapb.BigqueryreservationBetaCapacityCommitment {
	p := &betapb.BigqueryreservationBetaCapacityCommitment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSlotCount(dcl.ValueOrEmptyInt64(resource.SlotCount))
	p.SetPlan(BigqueryreservationBetaCapacityCommitmentPlanEnumToProto(resource.Plan))
	p.SetState(BigqueryreservationBetaCapacityCommitmentStateEnumToProto(resource.State))
	p.SetCommitmentStartTime(dcl.ValueOrEmptyString(resource.CommitmentStartTime))
	p.SetCommitmentEndTime(dcl.ValueOrEmptyString(resource.CommitmentEndTime))
	p.SetFailureStatus(BigqueryreservationBetaCapacityCommitmentFailureStatusToProto(resource.FailureStatus))
	p.SetRenewalPlan(BigqueryreservationBetaCapacityCommitmentRenewalPlanEnumToProto(resource.RenewalPlan))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) applyCapacityCommitment(ctx context.Context, c *beta.Client, request *betapb.ApplyBigqueryreservationBetaCapacityCommitmentRequest) (*betapb.BigqueryreservationBetaCapacityCommitment, error) {
	p := ProtoToCapacityCommitment(request.GetResource())
	res, err := c.ApplyCapacityCommitment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CapacityCommitmentToProto(res)
	return r, nil
}

// applyBigqueryreservationBetaCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) ApplyBigqueryreservationBetaCapacityCommitment(ctx context.Context, request *betapb.ApplyBigqueryreservationBetaCapacityCommitmentRequest) (*betapb.BigqueryreservationBetaCapacityCommitment, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCapacityCommitment(ctx, cl, request)
}

// DeleteCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Delete() method.
func (s *CapacityCommitmentServer) DeleteBigqueryreservationBetaCapacityCommitment(ctx context.Context, request *betapb.DeleteBigqueryreservationBetaCapacityCommitmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCapacityCommitment(ctx, ProtoToCapacityCommitment(request.GetResource()))

}

// ListBigqueryreservationBetaCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitmentList() method.
func (s *CapacityCommitmentServer) ListBigqueryreservationBetaCapacityCommitment(ctx context.Context, request *betapb.ListBigqueryreservationBetaCapacityCommitmentRequest) (*betapb.ListBigqueryreservationBetaCapacityCommitmentResponse, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCapacityCommitment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BigqueryreservationBetaCapacityCommitment
	for _, r := range resources.Items {
		rp := CapacityCommitmentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBigqueryreservationBetaCapacityCommitmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCapacityCommitment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
