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

// Server implements the gRPC interface for TargetPool.
type TargetPoolServer struct{}

// ProtoToTargetPoolSessionAffinityEnum converts a TargetPoolSessionAffinityEnum enum from its proto representation.
func ProtoToComputeTargetPoolSessionAffinityEnum(e computepb.ComputeTargetPoolSessionAffinityEnum) *compute.TargetPoolSessionAffinityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeTargetPoolSessionAffinityEnum_name[int32(e)]; ok {
		e := compute.TargetPoolSessionAffinityEnum(n[len("ComputeTargetPoolSessionAffinityEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetPool converts a TargetPool resource from its proto representation.
func ProtoToTargetPool(p *computepb.ComputeTargetPool) *compute.TargetPool {
	obj := &compute.TargetPool{
		BackupPool:      dcl.StringOrNil(p.BackupPool),
		Description:     dcl.StringOrNil(p.Description),
		FailoverRatio:   dcl.Float64OrNil(p.FailoverRatio),
		Name:            dcl.StringOrNil(p.Name),
		Region:          dcl.StringOrNil(p.Region),
		SelfLink:        dcl.StringOrNil(p.SelfLink),
		SessionAffinity: ProtoToComputeTargetPoolSessionAffinityEnum(p.GetSessionAffinity()),
		Project:         dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetHealthChecks() {
		obj.HealthChecks = append(obj.HealthChecks, r)
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	return obj
}

// TargetPoolSessionAffinityEnumToProto converts a TargetPoolSessionAffinityEnum enum to its proto representation.
func ComputeTargetPoolSessionAffinityEnumToProto(e *compute.TargetPoolSessionAffinityEnum) computepb.ComputeTargetPoolSessionAffinityEnum {
	if e == nil {
		return computepb.ComputeTargetPoolSessionAffinityEnum(0)
	}
	if v, ok := computepb.ComputeTargetPoolSessionAffinityEnum_value["TargetPoolSessionAffinityEnum"+string(*e)]; ok {
		return computepb.ComputeTargetPoolSessionAffinityEnum(v)
	}
	return computepb.ComputeTargetPoolSessionAffinityEnum(0)
}

// TargetPoolToProto converts a TargetPool resource to its proto representation.
func TargetPoolToProto(resource *compute.TargetPool) *computepb.ComputeTargetPool {
	p := &computepb.ComputeTargetPool{
		BackupPool:      dcl.ValueOrEmptyString(resource.BackupPool),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		FailoverRatio:   dcl.ValueOrEmptyDouble(resource.FailoverRatio),
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Region:          dcl.ValueOrEmptyString(resource.Region),
		SelfLink:        dcl.ValueOrEmptyString(resource.SelfLink),
		SessionAffinity: ComputeTargetPoolSessionAffinityEnumToProto(resource.SessionAffinity),
		Project:         dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.HealthChecks {
		p.HealthChecks = append(p.HealthChecks, r)
	}
	for _, r := range resource.Instances {
		p.Instances = append(p.Instances, r)
	}

	return p
}

// ApplyTargetPool handles the gRPC request by passing it to the underlying TargetPool Apply() method.
func (s *TargetPoolServer) applyTargetPool(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeTargetPoolRequest) (*computepb.ComputeTargetPool, error) {
	p := ProtoToTargetPool(request.GetResource())
	res, err := c.ApplyTargetPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetPoolToProto(res)
	return r, nil
}

// ApplyTargetPool handles the gRPC request by passing it to the underlying TargetPool Apply() method.
func (s *TargetPoolServer) ApplyComputeTargetPool(ctx context.Context, request *computepb.ApplyComputeTargetPoolRequest) (*computepb.ComputeTargetPool, error) {
	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetPool(ctx, cl, request)
}

// DeleteTargetPool handles the gRPC request by passing it to the underlying TargetPool Delete() method.
func (s *TargetPoolServer) DeleteComputeTargetPool(ctx context.Context, request *computepb.DeleteComputeTargetPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetPool(ctx, ProtoToTargetPool(request.GetResource()))

}

// ListComputeTargetPool handles the gRPC request by passing it to the underlying TargetPoolList() method.
func (s *TargetPoolServer) ListComputeTargetPool(ctx context.Context, request *computepb.ListComputeTargetPoolRequest) (*computepb.ListComputeTargetPoolResponse, error) {
	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetPool(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeTargetPool
	for _, r := range resources.Items {
		rp := TargetPoolToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeTargetPoolResponse{Items: protos}, nil
}

func createConfigTargetPool(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
