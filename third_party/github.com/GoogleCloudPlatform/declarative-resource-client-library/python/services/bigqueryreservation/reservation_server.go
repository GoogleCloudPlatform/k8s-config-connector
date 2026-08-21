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

// ReservationServer implements the gRPC interface for Reservation.
type ReservationServer struct{}

// ProtoToReservation converts a Reservation resource from its proto representation.
func ProtoToReservation(p *bigqueryreservationpb.BigqueryreservationReservation) *bigqueryreservation.Reservation {
	obj := &bigqueryreservation.Reservation{
		Name:            dcl.StringOrNil(p.GetName()),
		SlotCapacity:    dcl.Int64OrNil(p.GetSlotCapacity()),
		IgnoreIdleSlots: dcl.Bool(p.GetIgnoreIdleSlots()),
		CreationTime:    dcl.StringOrNil(p.GetCreationTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Project:         dcl.StringOrNil(p.GetProject()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ReservationToProto converts a Reservation resource to its proto representation.
func ReservationToProto(resource *bigqueryreservation.Reservation) *bigqueryreservationpb.BigqueryreservationReservation {
	p := &bigqueryreservationpb.BigqueryreservationReservation{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSlotCapacity(dcl.ValueOrEmptyInt64(resource.SlotCapacity))
	p.SetIgnoreIdleSlots(dcl.ValueOrEmptyBool(resource.IgnoreIdleSlots))
	p.SetCreationTime(dcl.ValueOrEmptyString(resource.CreationTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) applyReservation(ctx context.Context, c *bigqueryreservation.Client, request *bigqueryreservationpb.ApplyBigqueryreservationReservationRequest) (*bigqueryreservationpb.BigqueryreservationReservation, error) {
	p := ProtoToReservation(request.GetResource())
	res, err := c.ApplyReservation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ReservationToProto(res)
	return r, nil
}

// applyBigqueryreservationReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) ApplyBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.ApplyBigqueryreservationReservationRequest) (*bigqueryreservationpb.BigqueryreservationReservation, error) {
	cl, err := createConfigReservation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyReservation(ctx, cl, request)
}

// DeleteReservation handles the gRPC request by passing it to the underlying Reservation Delete() method.
func (s *ReservationServer) DeleteBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.DeleteBigqueryreservationReservationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigReservation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteReservation(ctx, ProtoToReservation(request.GetResource()))

}

// ListBigqueryreservationReservation handles the gRPC request by passing it to the underlying ReservationList() method.
func (s *ReservationServer) ListBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.ListBigqueryreservationReservationRequest) (*bigqueryreservationpb.ListBigqueryreservationReservationResponse, error) {
	cl, err := createConfigReservation(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListReservation(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryreservationpb.BigqueryreservationReservation
	for _, r := range resources.Items {
		rp := ReservationToProto(r)
		protos = append(protos, rp)
	}
	p := &bigqueryreservationpb.ListBigqueryreservationReservationResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigReservation(ctx context.Context, service_account_file string) (*bigqueryreservation.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryreservation.NewClient(conf), nil
}
