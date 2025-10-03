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

// Server implements the gRPC interface for HttpsHealthCheck.
type HttpsHealthCheckServer struct{}

// ProtoToHttpsHealthCheck converts a HttpsHealthCheck resource from its proto representation.
func ProtoToHttpsHealthCheck(p *computepb.ComputeHttpsHealthCheck) *compute.HttpsHealthCheck {
	obj := &compute.HttpsHealthCheck{
		CheckIntervalSec:   dcl.Int64OrNil(p.CheckIntervalSec),
		Description:        dcl.StringOrNil(p.Description),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		Host:               dcl.StringOrNil(p.Host),
		Name:               dcl.StringOrNil(p.Name),
		Port:               dcl.Int64OrNil(p.Port),
		RequestPath:        dcl.StringOrNil(p.RequestPath),
		TimeoutSec:         dcl.Int64OrNil(p.TimeoutSec),
		UnhealthyThreshold: dcl.Int64OrNil(p.UnhealthyThreshold),
		Project:            dcl.StringOrNil(p.Project),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		CreationTimestamp:  dcl.StringOrNil(p.CreationTimestamp),
	}
	return obj
}

// HttpsHealthCheckToProto converts a HttpsHealthCheck resource to its proto representation.
func HttpsHealthCheckToProto(resource *compute.HttpsHealthCheck) *computepb.ComputeHttpsHealthCheck {
	p := &computepb.ComputeHttpsHealthCheck{
		CheckIntervalSec:   dcl.ValueOrEmptyInt64(resource.CheckIntervalSec),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(resource.HealthyThreshold),
		Host:               dcl.ValueOrEmptyString(resource.Host),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Port:               dcl.ValueOrEmptyInt64(resource.Port),
		RequestPath:        dcl.ValueOrEmptyString(resource.RequestPath),
		TimeoutSec:         dcl.ValueOrEmptyInt64(resource.TimeoutSec),
		UnhealthyThreshold: dcl.ValueOrEmptyInt64(resource.UnhealthyThreshold),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
	}

	return p
}

// ApplyHttpsHealthCheck handles the gRPC request by passing it to the underlying HttpsHealthCheck Apply() method.
func (s *HttpsHealthCheckServer) applyHttpsHealthCheck(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeHttpsHealthCheckRequest) (*computepb.ComputeHttpsHealthCheck, error) {
	p := ProtoToHttpsHealthCheck(request.GetResource())
	res, err := c.ApplyHttpsHealthCheck(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpsHealthCheckToProto(res)
	return r, nil
}

// ApplyHttpsHealthCheck handles the gRPC request by passing it to the underlying HttpsHealthCheck Apply() method.
func (s *HttpsHealthCheckServer) ApplyComputeHttpsHealthCheck(ctx context.Context, request *computepb.ApplyComputeHttpsHealthCheckRequest) (*computepb.ComputeHttpsHealthCheck, error) {
	cl, err := createConfigHttpsHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHttpsHealthCheck(ctx, cl, request)
}

// DeleteHttpsHealthCheck handles the gRPC request by passing it to the underlying HttpsHealthCheck Delete() method.
func (s *HttpsHealthCheckServer) DeleteComputeHttpsHealthCheck(ctx context.Context, request *computepb.DeleteComputeHttpsHealthCheckRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpsHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpsHealthCheck(ctx, ProtoToHttpsHealthCheck(request.GetResource()))

}

// ListComputeHttpsHealthCheck handles the gRPC request by passing it to the underlying HttpsHealthCheckList() method.
func (s *HttpsHealthCheckServer) ListComputeHttpsHealthCheck(ctx context.Context, request *computepb.ListComputeHttpsHealthCheckRequest) (*computepb.ListComputeHttpsHealthCheckResponse, error) {
	cl, err := createConfigHttpsHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpsHealthCheck(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeHttpsHealthCheck
	for _, r := range resources.Items {
		rp := HttpsHealthCheckToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeHttpsHealthCheckResponse{Items: protos}, nil
}

func createConfigHttpsHealthCheck(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
