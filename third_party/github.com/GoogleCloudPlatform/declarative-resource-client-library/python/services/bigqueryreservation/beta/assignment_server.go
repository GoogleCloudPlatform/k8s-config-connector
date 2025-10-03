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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/beta/bigqueryreservation_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
)

// AssignmentServer implements the gRPC interface for Assignment.
type AssignmentServer struct{}

// ProtoToAssignmentJobTypeEnum converts a AssignmentJobTypeEnum enum from its proto representation.
func ProtoToBigqueryreservationBetaAssignmentJobTypeEnum(e betapb.BigqueryreservationBetaAssignmentJobTypeEnum) *beta.AssignmentJobTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryreservationBetaAssignmentJobTypeEnum_name[int32(e)]; ok {
		e := beta.AssignmentJobTypeEnum(n[len("BigqueryreservationBetaAssignmentJobTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignmentStateEnum converts a AssignmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationBetaAssignmentStateEnum(e betapb.BigqueryreservationBetaAssignmentStateEnum) *beta.AssignmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BigqueryreservationBetaAssignmentStateEnum_name[int32(e)]; ok {
		e := beta.AssignmentStateEnum(n[len("BigqueryreservationBetaAssignmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignment converts a Assignment resource from its proto representation.
func ProtoToAssignment(p *betapb.BigqueryreservationBetaAssignment) *beta.Assignment {
	obj := &beta.Assignment{
		Name:        dcl.StringOrNil(p.GetName()),
		Assignee:    dcl.StringOrNil(p.GetAssignee()),
		JobType:     ProtoToBigqueryreservationBetaAssignmentJobTypeEnum(p.GetJobType()),
		State:       ProtoToBigqueryreservationBetaAssignmentStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		Reservation: dcl.StringOrNil(p.GetReservation()),
	}
	return obj
}

// AssignmentJobTypeEnumToProto converts a AssignmentJobTypeEnum enum to its proto representation.
func BigqueryreservationBetaAssignmentJobTypeEnumToProto(e *beta.AssignmentJobTypeEnum) betapb.BigqueryreservationBetaAssignmentJobTypeEnum {
	if e == nil {
		return betapb.BigqueryreservationBetaAssignmentJobTypeEnum(0)
	}
	if v, ok := betapb.BigqueryreservationBetaAssignmentJobTypeEnum_value["AssignmentJobTypeEnum"+string(*e)]; ok {
		return betapb.BigqueryreservationBetaAssignmentJobTypeEnum(v)
	}
	return betapb.BigqueryreservationBetaAssignmentJobTypeEnum(0)
}

// AssignmentStateEnumToProto converts a AssignmentStateEnum enum to its proto representation.
func BigqueryreservationBetaAssignmentStateEnumToProto(e *beta.AssignmentStateEnum) betapb.BigqueryreservationBetaAssignmentStateEnum {
	if e == nil {
		return betapb.BigqueryreservationBetaAssignmentStateEnum(0)
	}
	if v, ok := betapb.BigqueryreservationBetaAssignmentStateEnum_value["AssignmentStateEnum"+string(*e)]; ok {
		return betapb.BigqueryreservationBetaAssignmentStateEnum(v)
	}
	return betapb.BigqueryreservationBetaAssignmentStateEnum(0)
}

// AssignmentToProto converts a Assignment resource to its proto representation.
func AssignmentToProto(resource *beta.Assignment) *betapb.BigqueryreservationBetaAssignment {
	p := &betapb.BigqueryreservationBetaAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAssignee(dcl.ValueOrEmptyString(resource.Assignee))
	p.SetJobType(BigqueryreservationBetaAssignmentJobTypeEnumToProto(resource.JobType))
	p.SetState(BigqueryreservationBetaAssignmentStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetReservation(dcl.ValueOrEmptyString(resource.Reservation))

	return p
}

// applyAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) applyAssignment(ctx context.Context, c *beta.Client, request *betapb.ApplyBigqueryreservationBetaAssignmentRequest) (*betapb.BigqueryreservationBetaAssignment, error) {
	p := ProtoToAssignment(request.GetResource())
	res, err := c.ApplyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssignmentToProto(res)
	return r, nil
}

// applyBigqueryreservationBetaAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) ApplyBigqueryreservationBetaAssignment(ctx context.Context, request *betapb.ApplyBigqueryreservationBetaAssignmentRequest) (*betapb.BigqueryreservationBetaAssignment, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAssignment(ctx, cl, request)
}

// DeleteAssignment handles the gRPC request by passing it to the underlying Assignment Delete() method.
func (s *AssignmentServer) DeleteBigqueryreservationBetaAssignment(ctx context.Context, request *betapb.DeleteBigqueryreservationBetaAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAssignment(ctx, ProtoToAssignment(request.GetResource()))

}

// ListBigqueryreservationBetaAssignment handles the gRPC request by passing it to the underlying AssignmentList() method.
func (s *AssignmentServer) ListBigqueryreservationBetaAssignment(ctx context.Context, request *betapb.ListBigqueryreservationBetaAssignmentRequest) (*betapb.ListBigqueryreservationBetaAssignmentResponse, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAssignment(ctx, request.GetProject(), request.GetLocation(), request.GetReservation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BigqueryreservationBetaAssignment
	for _, r := range resources.Items {
		rp := AssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBigqueryreservationBetaAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAssignment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
