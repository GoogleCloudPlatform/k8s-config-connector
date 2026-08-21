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

// Server implements the gRPC interface for TargetPool.
type TargetPoolServer struct{}

// ProtoToTargetPoolSessionAffinityEnum converts a TargetPoolSessionAffinityEnum enum from its proto representation.
func ProtoToComputeBetaTargetPoolSessionAffinityEnum(e betapb.ComputeBetaTargetPoolSessionAffinityEnum) *beta.TargetPoolSessionAffinityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaTargetPoolSessionAffinityEnum_name[int32(e)]; ok {
		e := beta.TargetPoolSessionAffinityEnum(n[len("ComputeBetaTargetPoolSessionAffinityEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetPool converts a TargetPool resource from its proto representation.
func ProtoToTargetPool(p *betapb.ComputeBetaTargetPool) *beta.TargetPool {
	obj := &beta.TargetPool{
		BackupPool:      dcl.StringOrNil(p.BackupPool),
		Description:     dcl.StringOrNil(p.Description),
		FailoverRatio:   dcl.Float64OrNil(p.FailoverRatio),
		Name:            dcl.StringOrNil(p.Name),
		Region:          dcl.StringOrNil(p.Region),
		SelfLink:        dcl.StringOrNil(p.SelfLink),
		SessionAffinity: ProtoToComputeBetaTargetPoolSessionAffinityEnum(p.GetSessionAffinity()),
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
func ComputeBetaTargetPoolSessionAffinityEnumToProto(e *beta.TargetPoolSessionAffinityEnum) betapb.ComputeBetaTargetPoolSessionAffinityEnum {
	if e == nil {
		return betapb.ComputeBetaTargetPoolSessionAffinityEnum(0)
	}
	if v, ok := betapb.ComputeBetaTargetPoolSessionAffinityEnum_value["TargetPoolSessionAffinityEnum"+string(*e)]; ok {
		return betapb.ComputeBetaTargetPoolSessionAffinityEnum(v)
	}
	return betapb.ComputeBetaTargetPoolSessionAffinityEnum(0)
}

// TargetPoolToProto converts a TargetPool resource to its proto representation.
func TargetPoolToProto(resource *beta.TargetPool) *betapb.ComputeBetaTargetPool {
	p := &betapb.ComputeBetaTargetPool{
		BackupPool:      dcl.ValueOrEmptyString(resource.BackupPool),
		Description:     dcl.ValueOrEmptyString(resource.Description),
		FailoverRatio:   dcl.ValueOrEmptyDouble(resource.FailoverRatio),
		Name:            dcl.ValueOrEmptyString(resource.Name),
		Region:          dcl.ValueOrEmptyString(resource.Region),
		SelfLink:        dcl.ValueOrEmptyString(resource.SelfLink),
		SessionAffinity: ComputeBetaTargetPoolSessionAffinityEnumToProto(resource.SessionAffinity),
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
func (s *TargetPoolServer) applyTargetPool(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaTargetPoolRequest) (*betapb.ComputeBetaTargetPool, error) {
	p := ProtoToTargetPool(request.GetResource())
	res, err := c.ApplyTargetPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetPoolToProto(res)
	return r, nil
}

// ApplyTargetPool handles the gRPC request by passing it to the underlying TargetPool Apply() method.
func (s *TargetPoolServer) ApplyComputeBetaTargetPool(ctx context.Context, request *betapb.ApplyComputeBetaTargetPoolRequest) (*betapb.ComputeBetaTargetPool, error) {
	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetPool(ctx, cl, request)
}

// DeleteTargetPool handles the gRPC request by passing it to the underlying TargetPool Delete() method.
func (s *TargetPoolServer) DeleteComputeBetaTargetPool(ctx context.Context, request *betapb.DeleteComputeBetaTargetPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetPool(ctx, ProtoToTargetPool(request.GetResource()))

}

// ListComputeBetaTargetPool handles the gRPC request by passing it to the underlying TargetPoolList() method.
func (s *TargetPoolServer) ListComputeBetaTargetPool(ctx context.Context, request *betapb.ListComputeBetaTargetPoolRequest) (*betapb.ListComputeBetaTargetPoolResponse, error) {
	cl, err := createConfigTargetPool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetPool(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaTargetPool
	for _, r := range resources.Items {
		rp := TargetPoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaTargetPoolResponse{Items: protos}, nil
}

func createConfigTargetPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
