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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/alpha/bigqueryreservation_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/alpha"
)

// CapacityCommitmentServer implements the gRPC interface for CapacityCommitment.
type CapacityCommitmentServer struct{}

// ProtoToCapacityCommitmentPlanEnum converts a CapacityCommitmentPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationAlphaCapacityCommitmentPlanEnum(e alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum) *alpha.CapacityCommitmentPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum_name[int32(e)]; ok {
		e := alpha.CapacityCommitmentPlanEnum(n[len("BigqueryreservationAlphaCapacityCommitmentPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentStateEnum converts a CapacityCommitmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationAlphaCapacityCommitmentStateEnum(e alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum) *alpha.CapacityCommitmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum_name[int32(e)]; ok {
		e := alpha.CapacityCommitmentStateEnum(n[len("BigqueryreservationAlphaCapacityCommitmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentRenewalPlanEnum converts a CapacityCommitmentRenewalPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum(e alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum) *alpha.CapacityCommitmentRenewalPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum_name[int32(e)]; ok {
		e := alpha.CapacityCommitmentRenewalPlanEnum(n[len("BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentFailureStatus converts a CapacityCommitmentFailureStatus object from its proto representation.
func ProtoToBigqueryreservationAlphaCapacityCommitmentFailureStatus(p *alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatus) *alpha.CapacityCommitmentFailureStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.CapacityCommitmentFailureStatus{
		Code:    dcl.Int64OrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToBigqueryreservationAlphaCapacityCommitmentFailureStatusDetails(r))
	}
	return obj
}

// ProtoToCapacityCommitmentFailureStatusDetails converts a CapacityCommitmentFailureStatusDetails object from its proto representation.
func ProtoToBigqueryreservationAlphaCapacityCommitmentFailureStatusDetails(p *alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatusDetails) *alpha.CapacityCommitmentFailureStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.CapacityCommitmentFailureStatusDetails{
		TypeUrl: dcl.StringOrNil(p.GetTypeUrl()),
		Value:   dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToCapacityCommitment converts a CapacityCommitment resource from its proto representation.
func ProtoToCapacityCommitment(p *alphapb.BigqueryreservationAlphaCapacityCommitment) *alpha.CapacityCommitment {
	obj := &alpha.CapacityCommitment{
		Name:                dcl.StringOrNil(p.GetName()),
		SlotCount:           dcl.Int64OrNil(p.GetSlotCount()),
		Plan:                ProtoToBigqueryreservationAlphaCapacityCommitmentPlanEnum(p.GetPlan()),
		State:               ProtoToBigqueryreservationAlphaCapacityCommitmentStateEnum(p.GetState()),
		CommitmentStartTime: dcl.StringOrNil(p.GetCommitmentStartTime()),
		CommitmentEndTime:   dcl.StringOrNil(p.GetCommitmentEndTime()),
		FailureStatus:       ProtoToBigqueryreservationAlphaCapacityCommitmentFailureStatus(p.GetFailureStatus()),
		RenewalPlan:         ProtoToBigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum(p.GetRenewalPlan()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// CapacityCommitmentPlanEnumToProto converts a CapacityCommitmentPlanEnum enum to its proto representation.
func BigqueryreservationAlphaCapacityCommitmentPlanEnumToProto(e *alpha.CapacityCommitmentPlanEnum) alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum {
	if e == nil {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum(0)
	}
	if v, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum_value["CapacityCommitmentPlanEnum"+string(*e)]; ok {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum(v)
	}
	return alphapb.BigqueryreservationAlphaCapacityCommitmentPlanEnum(0)
}

// CapacityCommitmentStateEnumToProto converts a CapacityCommitmentStateEnum enum to its proto representation.
func BigqueryreservationAlphaCapacityCommitmentStateEnumToProto(e *alpha.CapacityCommitmentStateEnum) alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum {
	if e == nil {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum(0)
	}
	if v, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum_value["CapacityCommitmentStateEnum"+string(*e)]; ok {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum(v)
	}
	return alphapb.BigqueryreservationAlphaCapacityCommitmentStateEnum(0)
}

// CapacityCommitmentRenewalPlanEnumToProto converts a CapacityCommitmentRenewalPlanEnum enum to its proto representation.
func BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnumToProto(e *alpha.CapacityCommitmentRenewalPlanEnum) alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum {
	if e == nil {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum(0)
	}
	if v, ok := alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum_value["CapacityCommitmentRenewalPlanEnum"+string(*e)]; ok {
		return alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum(v)
	}
	return alphapb.BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnum(0)
}

// CapacityCommitmentFailureStatusToProto converts a CapacityCommitmentFailureStatus object to its proto representation.
func BigqueryreservationAlphaCapacityCommitmentFailureStatusToProto(o *alpha.CapacityCommitmentFailureStatus) *alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatus{}
	p.SetCode(dcl.ValueOrEmptyInt64(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	sDetails := make([]*alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatusDetails, len(o.Details))
	for i, r := range o.Details {
		sDetails[i] = BigqueryreservationAlphaCapacityCommitmentFailureStatusDetailsToProto(&r)
	}
	p.SetDetails(sDetails)
	return p
}

// CapacityCommitmentFailureStatusDetailsToProto converts a CapacityCommitmentFailureStatusDetails object to its proto representation.
func BigqueryreservationAlphaCapacityCommitmentFailureStatusDetailsToProto(o *alpha.CapacityCommitmentFailureStatusDetails) *alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryreservationAlphaCapacityCommitmentFailureStatusDetails{}
	p.SetTypeUrl(dcl.ValueOrEmptyString(o.TypeUrl))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// CapacityCommitmentToProto converts a CapacityCommitment resource to its proto representation.
func CapacityCommitmentToProto(resource *alpha.CapacityCommitment) *alphapb.BigqueryreservationAlphaCapacityCommitment {
	p := &alphapb.BigqueryreservationAlphaCapacityCommitment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSlotCount(dcl.ValueOrEmptyInt64(resource.SlotCount))
	p.SetPlan(BigqueryreservationAlphaCapacityCommitmentPlanEnumToProto(resource.Plan))
	p.SetState(BigqueryreservationAlphaCapacityCommitmentStateEnumToProto(resource.State))
	p.SetCommitmentStartTime(dcl.ValueOrEmptyString(resource.CommitmentStartTime))
	p.SetCommitmentEndTime(dcl.ValueOrEmptyString(resource.CommitmentEndTime))
	p.SetFailureStatus(BigqueryreservationAlphaCapacityCommitmentFailureStatusToProto(resource.FailureStatus))
	p.SetRenewalPlan(BigqueryreservationAlphaCapacityCommitmentRenewalPlanEnumToProto(resource.RenewalPlan))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) applyCapacityCommitment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBigqueryreservationAlphaCapacityCommitmentRequest) (*alphapb.BigqueryreservationAlphaCapacityCommitment, error) {
	p := ProtoToCapacityCommitment(request.GetResource())
	res, err := c.ApplyCapacityCommitment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CapacityCommitmentToProto(res)
	return r, nil
}

// applyBigqueryreservationAlphaCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) ApplyBigqueryreservationAlphaCapacityCommitment(ctx context.Context, request *alphapb.ApplyBigqueryreservationAlphaCapacityCommitmentRequest) (*alphapb.BigqueryreservationAlphaCapacityCommitment, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCapacityCommitment(ctx, cl, request)
}

// DeleteCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Delete() method.
func (s *CapacityCommitmentServer) DeleteBigqueryreservationAlphaCapacityCommitment(ctx context.Context, request *alphapb.DeleteBigqueryreservationAlphaCapacityCommitmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCapacityCommitment(ctx, ProtoToCapacityCommitment(request.GetResource()))

}

// ListBigqueryreservationAlphaCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitmentList() method.
func (s *CapacityCommitmentServer) ListBigqueryreservationAlphaCapacityCommitment(ctx context.Context, request *alphapb.ListBigqueryreservationAlphaCapacityCommitmentRequest) (*alphapb.ListBigqueryreservationAlphaCapacityCommitmentResponse, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCapacityCommitment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BigqueryreservationAlphaCapacityCommitment
	for _, r := range resources.Items {
		rp := CapacityCommitmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBigqueryreservationAlphaCapacityCommitmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCapacityCommitment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
