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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Reservation.
type ReservationServer struct{}

// ProtoToReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum converts a ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum enum from its proto representation.
func ProtoToComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(e computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) *compute.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum_name[int32(e)]; ok {
		e := compute.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(n[len("ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToReservationStatusEnum converts a ReservationStatusEnum enum from its proto representation.
func ProtoToComputeReservationStatusEnum(e computepb.ComputeReservationStatusEnum) *compute.ReservationStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeReservationStatusEnum_name[int32(e)]; ok {
		e := compute.ReservationStatusEnum(n[len("ComputeReservationStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToReservationSpecificReservation converts a ReservationSpecificReservation resource from its proto representation.
func ProtoToComputeReservationSpecificReservation(p *computepb.ComputeReservationSpecificReservation) *compute.ReservationSpecificReservation {
	if p == nil {
		return nil
	}
	obj := &compute.ReservationSpecificReservation{
		InstanceProperties: ProtoToComputeReservationSpecificReservationInstanceProperties(p.GetInstanceProperties()),
		Count:              dcl.Int64OrNil(p.Count),
		InUseCount:         dcl.Int64OrNil(p.InUseCount),
	}
	return obj
}

// ProtoToReservationSpecificReservationInstanceProperties converts a ReservationSpecificReservationInstanceProperties resource from its proto representation.
func ProtoToComputeReservationSpecificReservationInstanceProperties(p *computepb.ComputeReservationSpecificReservationInstanceProperties) *compute.ReservationSpecificReservationInstanceProperties {
	if p == nil {
		return nil
	}
	obj := &compute.ReservationSpecificReservationInstanceProperties{
		MachineType:    dcl.StringOrNil(p.MachineType),
		MinCpuPlatform: dcl.StringOrNil(p.MinCpuPlatform),
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(r))
	}
	for _, r := range p.GetLocalSsds() {
		obj.LocalSsds = append(obj.LocalSsds, *ProtoToComputeReservationSpecificReservationInstancePropertiesLocalSsds(r))
	}
	return obj
}

// ProtoToReservationSpecificReservationInstancePropertiesGuestAccelerators converts a ReservationSpecificReservationInstancePropertiesGuestAccelerators resource from its proto representation.
func ProtoToComputeReservationSpecificReservationInstancePropertiesGuestAccelerators(p *computepb.ComputeReservationSpecificReservationInstancePropertiesGuestAccelerators) *compute.ReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &compute.ReservationSpecificReservationInstancePropertiesGuestAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
	}
	return obj
}

// ProtoToReservationSpecificReservationInstancePropertiesLocalSsds converts a ReservationSpecificReservationInstancePropertiesLocalSsds resource from its proto representation.
func ProtoToComputeReservationSpecificReservationInstancePropertiesLocalSsds(p *computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsds) *compute.ReservationSpecificReservationInstancePropertiesLocalSsds {
	if p == nil {
		return nil
	}
	obj := &compute.ReservationSpecificReservationInstancePropertiesLocalSsds{
		DiskSizeGb: dcl.Int64OrNil(p.DiskSizeGb),
		Interface:  ProtoToComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(p.GetInterface()),
	}
	return obj
}

// ProtoToReservation converts a Reservation resource from its proto representation.
func ProtoToReservation(p *computepb.ComputeReservation) *compute.Reservation {
	obj := &compute.Reservation{
		Id:                          dcl.Int64OrNil(p.Id),
		SelfLink:                    dcl.StringOrNil(p.SelfLink),
		Zone:                        dcl.StringOrNil(p.Zone),
		Description:                 dcl.StringOrNil(p.Description),
		Name:                        dcl.StringOrNil(p.Name),
		SpecificReservation:         ProtoToComputeReservationSpecificReservation(p.GetSpecificReservation()),
		Commitment:                  dcl.StringOrNil(p.Commitment),
		SpecificReservationRequired: dcl.Bool(p.SpecificReservationRequired),
		Status:                      ProtoToComputeReservationStatusEnum(p.GetStatus()),
		Project:                     dcl.StringOrNil(p.Project),
	}
	return obj
}

// ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto converts a ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum enum to its proto representation.
func ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto(e *compute.ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum) computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum {
	if e == nil {
		return computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(0)
	}
	if v, ok := computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum_value["ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum"+string(*e)]; ok {
		return computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(v)
	}
	return computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum(0)
}

// ReservationStatusEnumToProto converts a ReservationStatusEnum enum to its proto representation.
func ComputeReservationStatusEnumToProto(e *compute.ReservationStatusEnum) computepb.ComputeReservationStatusEnum {
	if e == nil {
		return computepb.ComputeReservationStatusEnum(0)
	}
	if v, ok := computepb.ComputeReservationStatusEnum_value["ReservationStatusEnum"+string(*e)]; ok {
		return computepb.ComputeReservationStatusEnum(v)
	}
	return computepb.ComputeReservationStatusEnum(0)
}

// ReservationSpecificReservationToProto converts a ReservationSpecificReservation resource to its proto representation.
func ComputeReservationSpecificReservationToProto(o *compute.ReservationSpecificReservation) *computepb.ComputeReservationSpecificReservation {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeReservationSpecificReservation{
		InstanceProperties: ComputeReservationSpecificReservationInstancePropertiesToProto(o.InstanceProperties),
		Count:              dcl.ValueOrEmptyInt64(o.Count),
		InUseCount:         dcl.ValueOrEmptyInt64(o.InUseCount),
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesToProto converts a ReservationSpecificReservationInstanceProperties resource to its proto representation.
func ComputeReservationSpecificReservationInstancePropertiesToProto(o *compute.ReservationSpecificReservationInstanceProperties) *computepb.ComputeReservationSpecificReservationInstanceProperties {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeReservationSpecificReservationInstanceProperties{
		MachineType:    dcl.ValueOrEmptyString(o.MachineType),
		MinCpuPlatform: dcl.ValueOrEmptyString(o.MinCpuPlatform),
	}
	for _, r := range o.GuestAccelerators {
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto(&r))
	}
	for _, r := range o.LocalSsds {
		p.LocalSsds = append(p.LocalSsds, ComputeReservationSpecificReservationInstancePropertiesLocalSsdsToProto(&r))
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto converts a ReservationSpecificReservationInstancePropertiesGuestAccelerators resource to its proto representation.
func ComputeReservationSpecificReservationInstancePropertiesGuestAcceleratorsToProto(o *compute.ReservationSpecificReservationInstancePropertiesGuestAccelerators) *computepb.ComputeReservationSpecificReservationInstancePropertiesGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeReservationSpecificReservationInstancePropertiesGuestAccelerators{
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
	}
	return p
}

// ReservationSpecificReservationInstancePropertiesLocalSsdsToProto converts a ReservationSpecificReservationInstancePropertiesLocalSsds resource to its proto representation.
func ComputeReservationSpecificReservationInstancePropertiesLocalSsdsToProto(o *compute.ReservationSpecificReservationInstancePropertiesLocalSsds) *computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsds {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeReservationSpecificReservationInstancePropertiesLocalSsds{
		DiskSizeGb: dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		Interface:  ComputeReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnumToProto(o.Interface),
	}
	return p
}

// ReservationToProto converts a Reservation resource to its proto representation.
func ReservationToProto(resource *compute.Reservation) *computepb.ComputeReservation {
	p := &computepb.ComputeReservation{
		Id:                          dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:                    dcl.ValueOrEmptyString(resource.SelfLink),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		SpecificReservation:         ComputeReservationSpecificReservationToProto(resource.SpecificReservation),
		Commitment:                  dcl.ValueOrEmptyString(resource.Commitment),
		SpecificReservationRequired: dcl.ValueOrEmptyBool(resource.SpecificReservationRequired),
		Status:                      ComputeReservationStatusEnumToProto(resource.Status),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) applyReservation(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeReservationRequest) (*computepb.ComputeReservation, error) {
	p := ProtoToReservation(request.GetResource())
	res, err := c.ApplyReservation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ReservationToProto(res)
	return r, nil
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) ApplyComputeReservation(ctx context.Context, request *computepb.ApplyComputeReservationRequest) (*computepb.ComputeReservation, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyReservation(ctx, cl, request)
}

// DeleteReservation handles the gRPC request by passing it to the underlying Reservation Delete() method.
func (s *ReservationServer) DeleteComputeReservation(ctx context.Context, request *computepb.DeleteComputeReservationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteReservation(ctx, ProtoToReservation(request.GetResource()))

}

// ListComputeReservation handles the gRPC request by passing it to the underlying ReservationList() method.
func (s *ReservationServer) ListComputeReservation(ctx context.Context, request *computepb.ListComputeReservationRequest) (*computepb.ListComputeReservationResponse, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListReservation(ctx, request.Project, request.Zone)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeReservation
	for _, r := range resources.Items {
		rp := ReservationToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeReservationResponse{Items: protos}, nil
}

func createConfigReservation(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
