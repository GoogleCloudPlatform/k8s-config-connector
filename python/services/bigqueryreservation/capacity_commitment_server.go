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
	bigqueryreservationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/bigqueryreservation_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
)

// CapacityCommitmentServer implements the gRPC interface for CapacityCommitment.
type CapacityCommitmentServer struct{}

// ProtoToCapacityCommitmentPlanEnum converts a CapacityCommitmentPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentPlanEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum) *bigqueryreservation.CapacityCommitmentPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentPlanEnum(n[len("BigqueryreservationCapacityCommitmentPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentStateEnum converts a CapacityCommitmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentStateEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum) *bigqueryreservation.CapacityCommitmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentStateEnum(n[len("BigqueryreservationCapacityCommitmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentRenewalPlanEnum converts a CapacityCommitmentRenewalPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentRenewalPlanEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum) *bigqueryreservation.CapacityCommitmentRenewalPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentRenewalPlanEnum(n[len("BigqueryreservationCapacityCommitmentRenewalPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentFailureStatus converts a CapacityCommitmentFailureStatus object from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentFailureStatus(p *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus) *bigqueryreservation.CapacityCommitmentFailureStatus {
	if p == nil {
		return nil
	}
	obj := &bigqueryreservation.CapacityCommitmentFailureStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToBigqueryreservationCapacityCommitmentFailureStatusDetails(r))
	}
	return obj
}

// ProtoToCapacityCommitmentFailureStatusDetails converts a CapacityCommitmentFailureStatusDetails object from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentFailureStatusDetails(p *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails) *bigqueryreservation.CapacityCommitmentFailureStatusDetails {
	if p == nil {
		return nil
	}
	obj := &bigqueryreservation.CapacityCommitmentFailureStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCapacityCommitment converts a CapacityCommitment resource from its proto representation.
func ProtoToCapacityCommitment(p *bigqueryreservationpb.BigqueryreservationCapacityCommitment) *bigqueryreservation.CapacityCommitment {
	obj := &bigqueryreservation.CapacityCommitment{
		Name:                dcl.StringOrNil(p.GetName()),
		SlotCount:           dcl.Int64OrNil(p.GetSlotCount()),
		Plan:                ProtoToBigqueryreservationCapacityCommitmentPlanEnum(p.GetPlan()),
		State:               ProtoToBigqueryreservationCapacityCommitmentStateEnum(p.GetState()),
		CommitmentStartTime: dcl.StringOrNil(p.GetCommitmentStartTime()),
		CommitmentEndTime:   dcl.StringOrNil(p.GetCommitmentEndTime()),
		FailureStatus:       ProtoToBigqueryreservationCapacityCommitmentFailureStatus(p.GetFailureStatus()),
		RenewalPlan:         ProtoToBigqueryreservationCapacityCommitmentRenewalPlanEnum(p.GetRenewalPlan()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CapacityCommitmentPlanEnumToProto converts a CapacityCommitmentPlanEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentPlanEnumToProto(e *bigqueryreservation.CapacityCommitmentPlanEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum_value["CapacityCommitmentPlanEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(0)
}

// CapacityCommitmentStateEnumToProto converts a CapacityCommitmentStateEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentStateEnumToProto(e *bigqueryreservation.CapacityCommitmentStateEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum_value["CapacityCommitmentStateEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(0)
}

// CapacityCommitmentRenewalPlanEnumToProto converts a CapacityCommitmentRenewalPlanEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentRenewalPlanEnumToProto(e *bigqueryreservation.CapacityCommitmentRenewalPlanEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum_value["CapacityCommitmentRenewalPlanEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(0)
}

// CapacityCommitmentFailureStatusToProto converts a CapacityCommitmentFailureStatus object to its proto representation.
func BigqueryreservationCapacityCommitmentFailureStatusToProto(o *bigqueryreservation.CapacityCommitmentFailureStatus) *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus {
	if o == nil {
		return nil
	}
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = BigqueryreservationCapacityCommitmentFailureStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// CapacityCommitmentFailureStatusDetailsToProto converts a CapacityCommitmentFailureStatusDetails object to its proto representation.
func BigqueryreservationCapacityCommitmentFailureStatusDetailsToProto(o *bigqueryreservation.CapacityCommitmentFailureStatusDetails) *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails {
	if o == nil {
		return nil
	}
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CapacityCommitmentToProto converts a CapacityCommitment resource to its proto representation.
func CapacityCommitmentToProto(resource *bigqueryreservation.CapacityCommitment) *bigqueryreservationpb.BigqueryreservationCapacityCommitment {
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSlotCount(dcl.ValueOrEmptyInt64(resource.SlotCount))
	p.SetPlan(BigqueryreservationCapacityCommitmentPlanEnumToProto(resource.Plan))
	p.SetState(BigqueryreservationCapacityCommitmentStateEnumToProto(resource.State))
	p.SetCommitmentStartTime(dcl.ValueOrEmptyString(resource.CommitmentStartTime))
	p.SetCommitmentEndTime(dcl.ValueOrEmptyString(resource.CommitmentEndTime))
	p.SetFailureStatus(BigqueryreservationCapacityCommitmentFailureStatusToProto(resource.FailureStatus))
	p.SetRenewalPlan(BigqueryreservationCapacityCommitmentRenewalPlanEnumToProto(resource.RenewalPlan))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) applyCapacityCommitment(ctx context.Context, c *bigqueryreservation.Client, request *bigqueryreservationpb.ApplyBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.BigqueryreservationCapacityCommitment, error) {
	p := ProtoToCapacityCommitment(request.GetResource())
	res, err := c.ApplyCapacityCommitment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CapacityCommitmentToProto(res)
	return r, nil
}

// applyBigqueryreservationCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) ApplyBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.ApplyBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.BigqueryreservationCapacityCommitment, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCapacityCommitment(ctx, cl, request)
}

// DeleteCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Delete() method.
func (s *CapacityCommitmentServer) DeleteBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.DeleteBigqueryreservationCapacityCommitmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCapacityCommitment(ctx, ProtoToCapacityCommitment(request.GetResource()))

}

// ListBigqueryreservationCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitmentList() method.
func (s *CapacityCommitmentServer) ListBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentResponse, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCapacityCommitment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryreservationpb.BigqueryreservationCapacityCommitment
	for _, r := range resources.Items {
		rp := CapacityCommitmentToProto(r)
		protos = append(protos, rp)
	}
	p := &bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCapacityCommitment(ctx context.Context, service_account_file string) (*bigqueryreservation.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryreservation.NewClient(conf), nil
}
