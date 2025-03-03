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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/alpha/bigqueryreservation_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/alpha"
)

// AssignmentServer implements the gRPC interface for Assignment.
type AssignmentServer struct{}

// ProtoToAssignmentJobTypeEnum converts a AssignmentJobTypeEnum enum from its proto representation.
func ProtoToBigqueryreservationAlphaAssignmentJobTypeEnum(e alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum) *alpha.AssignmentJobTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum_name[int32(e)]; ok {
		e := alpha.AssignmentJobTypeEnum(n[len("BigqueryreservationAlphaAssignmentJobTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignmentStateEnum converts a AssignmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationAlphaAssignmentStateEnum(e alphapb.BigqueryreservationAlphaAssignmentStateEnum) *alpha.AssignmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BigqueryreservationAlphaAssignmentStateEnum_name[int32(e)]; ok {
		e := alpha.AssignmentStateEnum(n[len("BigqueryreservationAlphaAssignmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToAssignment converts a Assignment resource from its proto representation.
func ProtoToAssignment(p *alphapb.BigqueryreservationAlphaAssignment) *alpha.Assignment {
	obj := &alpha.Assignment{
		Name:        dcl.StringOrNil(p.GetName()),
		Assignee:    dcl.StringOrNil(p.GetAssignee()),
		JobType:     ProtoToBigqueryreservationAlphaAssignmentJobTypeEnum(p.GetJobType()),
		State:       ProtoToBigqueryreservationAlphaAssignmentStateEnum(p.GetState()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		Reservation: dcl.StringOrNil(p.GetReservation()),
	}
	return obj
}

// AssignmentJobTypeEnumToProto converts a AssignmentJobTypeEnum enum to its proto representation.
func BigqueryreservationAlphaAssignmentJobTypeEnumToProto(e *alpha.AssignmentJobTypeEnum) alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum {
	if e == nil {
		return alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum(0)
	}
	if v, ok := alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum_value["AssignmentJobTypeEnum"+string(*e)]; ok {
		return alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum(v)
	}
	return alphapb.BigqueryreservationAlphaAssignmentJobTypeEnum(0)
}

// AssignmentStateEnumToProto converts a AssignmentStateEnum enum to its proto representation.
func BigqueryreservationAlphaAssignmentStateEnumToProto(e *alpha.AssignmentStateEnum) alphapb.BigqueryreservationAlphaAssignmentStateEnum {
	if e == nil {
		return alphapb.BigqueryreservationAlphaAssignmentStateEnum(0)
	}
	if v, ok := alphapb.BigqueryreservationAlphaAssignmentStateEnum_value["AssignmentStateEnum"+string(*e)]; ok {
		return alphapb.BigqueryreservationAlphaAssignmentStateEnum(v)
	}
	return alphapb.BigqueryreservationAlphaAssignmentStateEnum(0)
}

// AssignmentToProto converts a Assignment resource to its proto representation.
func AssignmentToProto(resource *alpha.Assignment) *alphapb.BigqueryreservationAlphaAssignment {
	p := &alphapb.BigqueryreservationAlphaAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAssignee(dcl.ValueOrEmptyString(resource.Assignee))
	p.SetJobType(BigqueryreservationAlphaAssignmentJobTypeEnumToProto(resource.JobType))
	p.SetState(BigqueryreservationAlphaAssignmentStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetReservation(dcl.ValueOrEmptyString(resource.Reservation))

	return p
}

// applyAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) applyAssignment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBigqueryreservationAlphaAssignmentRequest) (*alphapb.BigqueryreservationAlphaAssignment, error) {
	p := ProtoToAssignment(request.GetResource())
	res, err := c.ApplyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AssignmentToProto(res)
	return r, nil
}

// applyBigqueryreservationAlphaAssignment handles the gRPC request by passing it to the underlying Assignment Apply() method.
func (s *AssignmentServer) ApplyBigqueryreservationAlphaAssignment(ctx context.Context, request *alphapb.ApplyBigqueryreservationAlphaAssignmentRequest) (*alphapb.BigqueryreservationAlphaAssignment, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAssignment(ctx, cl, request)
}

// DeleteAssignment handles the gRPC request by passing it to the underlying Assignment Delete() method.
func (s *AssignmentServer) DeleteBigqueryreservationAlphaAssignment(ctx context.Context, request *alphapb.DeleteBigqueryreservationAlphaAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAssignment(ctx, ProtoToAssignment(request.GetResource()))

}

// ListBigqueryreservationAlphaAssignment handles the gRPC request by passing it to the underlying AssignmentList() method.
func (s *AssignmentServer) ListBigqueryreservationAlphaAssignment(ctx context.Context, request *alphapb.ListBigqueryreservationAlphaAssignmentRequest) (*alphapb.ListBigqueryreservationAlphaAssignmentResponse, error) {
	cl, err := createConfigAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAssignment(ctx, request.GetProject(), request.GetLocation(), request.GetReservation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BigqueryreservationAlphaAssignment
	for _, r := range resources.Items {
		rp := AssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBigqueryreservationAlphaAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAssignment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
