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
	bigqueryreservationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/bigqueryreservation_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
)

// AssignmentServer implements the gRPC interface for Assignment.
type AssignmentServer struct{}

// ProtoToAssignmentJobTypeEnum converts a AssignmentJobTypeEnum enum from its proto representation.
func ProtoToBigqueryreservationAssignmentJobTypeEnum(e bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum) *bigqueryreservation.AssignmentJobTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum_name[int32(e)]; ok {
		e := bigqueryreservation.AssignmentJobTypeEnum(n[len("BigqueryreservationAssignmentJobTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignmentStateEnum converts a AssignmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationAssignmentStateEnum(e bigqueryreservationpb.BigqueryreservationAssignmentStateEnum) *bigqueryreservation.AssignmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationAssignmentStateEnum_name[int32(e)]; ok {
		e := bigqueryreservation.AssignmentStateEnum(n[len("BigqueryreservationAssignmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignment converts a Assignment resource from its proto representation.
func ProtoToAssignment(p *bigqueryreservationpb.BigqueryreservationAssignment) *bigqueryreservation.Assignment {
	obj := &bigqueryreservation.Assignment{
		Name:        dcl.StringOrNil(p.GetName()),
		Assignee:    dcl.StringOrNil(p.GetAssignee()),
		JobType:     ProtoToBigqueryreservationAssignmentJobTypeEnum(p.GetJobType()),
		State:       ProtoToBigqueryreservationAssignmentStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		Reservation: dcl.StringOrNil(p.GetReservation()),
	}
	return obj
}

// AssignmentJobTypeEnumToProto converts a AssignmentJobTypeEnum enum to its proto representation.
func BigqueryreservationAssignmentJobTypeEnumToProto(e *bigqueryreservation.AssignmentJobTypeEnum) bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum_value["AssignmentJobTypeEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationAssignmentJobTypeEnum(0)
}

// AssignmentStateEnumToProto converts a AssignmentStateEnum enum to its proto representation.
func BigqueryreservationAssignmentStateEnumToProto(e *bigqueryreservation.AssignmentStateEnum) bigqueryreservationpb.BigqueryreservationAssignmentStateEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationAssignmentStateEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationAssignmentStateEnum_value["AssignmentStateEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationAssignmentStateEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationAssignmentStateEnum(0)
}

// AssignmentToProto converts a Assignment resource to its proto representation.
func AssignmentToProto(resource *bigqueryreservation.Assignment) *bigqueryreservationpb.BigqueryreservationAssignment {
	p := &bigqueryreservationpb.BigqueryreservationAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAssignee(dcl.ValueOrEmptyString(resource.Assignee))
	p.SetJobType(BigqueryreservationAssignmentJobTypeEnumToProto(resource.JobType))
	p.SetState(BigqueryreservationAssignmentStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetReservation(dcl.ValueOrEmptyString(resource.Reservation))

	return p
}

// applyAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) applyAssignment(ctx context.Context, c *bigqueryreservation.Client, request *bigqueryreservationpb.ApplyBigqueryreservationAssignmentRequest) (*bigqueryreservationpb.BigqueryreservationAssignment, error) {
	p := ProtoToAssignment(request.GetResource())
	res, err := c.ApplyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssignmentToProto(res)
	return r, nil
}

// applyBigqueryreservationAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) ApplyBigqueryreservationAssignment(ctx context.Context, request *bigqueryreservationpb.ApplyBigqueryreservationAssignmentRequest) (*bigqueryreservationpb.BigqueryreservationAssignment, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAssignment(ctx, cl, request)
}

// DeleteAssignment handles the gRPC request by passing it to the underlying Assignment Delete() method.
func (s *AssignmentServer) DeleteBigqueryreservationAssignment(ctx context.Context, request *bigqueryreservationpb.DeleteBigqueryreservationAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAssignment(ctx, ProtoToAssignment(request.GetResource()))

}

// ListBigqueryreservationAssignment handles the gRPC request by passing it to the underlying AssignmentList() method.
func (s *AssignmentServer) ListBigqueryreservationAssignment(ctx context.Context, request *bigqueryreservationpb.ListBigqueryreservationAssignmentRequest) (*bigqueryreservationpb.ListBigqueryreservationAssignmentResponse, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAssignment(ctx, request.GetProject(), request.GetLocation(), request.GetReservation())
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryreservationpb.BigqueryreservationAssignment
	for _, r := range resources.Items {
		rp := AssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &bigqueryreservationpb.ListBigqueryreservationAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAssignment(ctx context.Context, service_account_file string) (*bigqueryreservation.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryreservation.NewClient(conf), nil
}
