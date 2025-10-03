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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Reservation.
type ReservationServer struct{}

// ProtoToReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum converts a ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum enum from its proto representation.
func ProtoToComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(e betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) *beta.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum_name[int32(e)]; ok {
		e := beta.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(n[len("ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToReservationStatusEnum converts a ReservationStatusEnum enum from its proto representation.
func ProtoToComputeBetaReservationStatusEnum(e betapb.ComputeBetaReservationStatusEnum) *beta.ReservationStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaReservationStatusEnum_name[int32(e)]; ok {
		e := beta.ReservationStatusEnum(n[len("ComputeBetaReservationStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToReservationSpecificReservation converts a ReservationSpecificReservation resource from its proto representation.
func ProtoToComputeBetaReservationSpecificReservation(p *betapb.ComputeBetaReservationSpecificReservation) *beta.ReservationSpecificReservation {
	if p == nil {
		return nil
	}
	obj := &beta.ReservationSpecificReservation{
		InstanceProperties: ProtoToComputeBetaReservationSpecificReservationInstanceProperties(p.GetInstanceProperties()),
		Count:              dcl.Int64OrNil(p.Count),
		InUseCount:         dcl.Int64OrNil(p.InUseCount),
	}
	return obj
}

// ProtoToReservationSpecificReservationInstanceProperties converts a ReservationSpecificReservationInstanceProperties resource from its proto representation.
func ProtoToComputeBetaReservationSpecificReservationInstanceProperties(p *betapb.ComputeBetaReservationSpecificReservationInstanceProperties) *beta.ReservationSpecificReservationInstanceProperties {
	if p == nil {
		return nil
	}
	obj := &beta.ReservationSpecificReservationInstanceProperties{
		MachineType:    dcl.StringOrNil(p.MachineType),
		MinCpuPlatform: dcl.StringOrNil(p.MinCpuPlatform),
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators(r))
	}
	for _, r := range p.GetLocalSsds() {
		obj.LocalSsds = append(obj.LocalSsds, *ProtoToComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds(r))
	}
	return obj
}

// ProtoToReservationSpecificReservationInstancePropertiesGuestAccelerators converts a ReservationSpecificReservationInstancePropertiesGuestAccelerators resource from its proto representation.
func ProtoToComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators(p *betapb.ComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators) *beta.ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.ReservationSpecificReservationInstancePropertiesGuestAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
	}
	return obj
}

// ProtoToReservationSpecificReservationInstancePropertiesLocalSsds converts a ReservationSpecificReservationInstancePropertiesLocalSsds resource from its proto representation.
func ProtoToComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds(p *betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds) *beta.ReservationSpecificReservationInstancePropertiesLocalSsds {
	if p == nil {
		return nil
	}
	obj := &beta.ReservationSpecificReservationInstancePropertiesLocalSsds{
		DiskSizeGb: dcl.Int64OrNil(p.DiskSizeGb),
		Interface:  ProtoToComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(p.GetInterface()),
	}
	return obj
}

// ProtoToReservation converts a Reservation resource from its proto representation.
func ProtoToReservation(p *betapb.ComputeBetaReservation) *beta.Reservation {
	obj := &beta.Reservation{
		Id:                          dcl.Int64OrNil(p.Id),
		SelfLink:                    dcl.StringOrNil(p.SelfLink),
		Zone:                        dcl.StringOrNil(p.Zone),
		Description:                 dcl.StringOrNil(p.Description),
		Name:                        dcl.StringOrNil(p.Name),
		SpecificReservation:         ProtoToComputeBetaReservationSpecificReservation(p.GetSpecificReservation()),
		Commitment:                  dcl.StringOrNil(p.Commitment),
		SpecificReservationRequired: dcl.Bool(p.SpecificReservationRequired),
		Status:                      ProtoToComputeBetaReservationStatusEnum(p.GetStatus()),
		Project:                     dcl.StringOrNil(p.Project),
	}
	return obj
}

// ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto converts a ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum enum to its proto representation.
func ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto(e *beta.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	if e == nil {
		return betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(0)
	}
	if v, ok := betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum_value["ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(v)
	}
	return betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(0)
}

// ReservationStatusEnumToProto converts a ReservationStatusEnum enum to its proto representation.
func ComputeBetaReservationStatusEnumToProto(e *beta.ReservationStatusEnum) betapb.ComputeBetaReservationStatusEnum {
	if e == nil {
		return betapb.ComputeBetaReservationStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaReservationStatusEnum_value["ReservationStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaReservationStatusEnum(v)
	}
	return betapb.ComputeBetaReservationStatusEnum(0)
}

// ReservationSpecificReservationToProto converts a ReservationSpecificReservation resource to its proto representation.
func ComputeBetaReservationSpecificReservationToProto(o *beta.ReservationSpecificReservation) *betapb.ComputeBetaReservationSpecificReservation {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaReservationSpecificReservation{
		InstanceProperties: ComputeBetaReservationSpecificReservationInstancePropertiesToProto(o.InstanceProperties),
		Count:              dcl.ValueOrEmptyInt64(o.Count),
		InUseCount:         dcl.ValueOrEmptyInt64(o.InUseCount),
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesToProto converts a ReservationSpecificReservationInstanceProperties resource to its proto representation.
func ComputeBetaReservationSpecificReservationInstancePropertiesToProto(o *beta.ReservationSpecificReservationInstanceProperties) *betapb.ComputeBetaReservationSpecificReservationInstanceProperties {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaReservationSpecificReservationInstanceProperties{
		MachineType:    dcl.ValueOrEmptyString(o.MachineType),
		MinCpuPlatform: dcl.ValueOrEmptyString(o.MinCpuPlatform),
	}
	for _, r := range o.GuestAccelerators {
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeBetaReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto(&r))
	}
	for _, r := range o.LocalSsds {
		p.LocalSsds = append(p.LocalSsds, ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsToProto(&r))
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto converts a ReservationSpecificReservationInstancePropertiesGuestAccelerators resource to its proto representation.
func ComputeBetaReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto(o *beta.ReservationSpecificReservationInstancePropertiesGuestAccelerators) *betapb.ComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaReservationSpecificReservationInstancePropertiesGuestAccelerators{
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesLocalSsdsToProto converts a ReservationSpecificReservationInstancePropertiesLocalSsds resource to its proto representation.
func ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsToProto(o *beta.ReservationSpecificReservationInstancePropertiesLocalSsds) *betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsds{
		DiskSizeGb: dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		Interface:  ComputeBetaReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto(o.Interface),
	}
	return p
}

// ReservationToProto converts a Reservation resource to its proto representation.
func ReservationToProto(resource *beta.Reservation) *betapb.ComputeBetaReservation {
	p := &betapb.ComputeBetaReservation{
		Id:                          dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:                    dcl.ValueOrEmptyString(resource.SelfLink),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		SpecificReservation:         ComputeBetaReservationSpecificReservationToProto(resource.SpecificReservation),
		Commitment:                  dcl.ValueOrEmptyString(resource.Commitment),
		SpecificReservationRequired: dcl.ValueOrEmptyBool(resource.SpecificReservationRequired),
		Status:                      ComputeBetaReservationStatusEnumToProto(resource.Status),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) applyReservation(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaReservationRequest) (*betapb.ComputeBetaReservation, error) {
	p := ProtoToReservation(request.GetResource())
	res, err := c.ApplyReservation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ReservationToProto(res)
	return r, nil
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) ApplyComputeBetaReservation(ctx context.Context, request *betapb.ApplyComputeBetaReservationRequest) (*betapb.ComputeBetaReservation, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyReservation(ctx, cl, request)
}

// DeleteReservation handles the gRPC request by passing it to the underlying Reservation Delete() method.
func (s *ReservationServer) DeleteComputeBetaReservation(ctx context.Context, request *betapb.DeleteComputeBetaReservationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteReservation(ctx, ProtoToReservation(request.GetResource()))

}

// ListComputeBetaReservation handles the gRPC request by passing it to the underlying ReservationList() method.
func (s *ReservationServer) ListComputeBetaReservation(ctx context.Context, request *betapb.ListComputeBetaReservationRequest) (*betapb.ListComputeBetaReservationResponse, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListReservation(ctx, request.Project, request.Zone)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaReservation
	for _, r := range resources.Items {
		rp := ReservationToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaReservationResponse{Items: protos}, nil
}

func createConfigReservation(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
