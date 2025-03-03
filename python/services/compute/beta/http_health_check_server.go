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

// Server implements the gRPC interface for HttpHealthCheck.
type HttpHealthCheckServer struct{}

// ProtoToHttpHealthCheck converts a HttpHealthCheck resource from its proto representation.
func ProtoToHttpHealthCheck(p *betapb.ComputeBetaHttpHealthCheck) *beta.HttpHealthCheck {
	obj := &beta.HttpHealthCheck{
		CheckIntervalSec:   dcl.Int64OrNil(p.CheckIntervalSec),
		Description:        dcl.StringOrNil(p.Description),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		Host:               dcl.StringOrNil(p.Host),
		Name:               dcl.StringOrNil(p.Name),
		Port:               dcl.Int64OrNil(p.Port),
		RequestPath:        dcl.StringOrNil(p.RequestPath),
		TimeoutSec:         dcl.Int64OrNil(p.TimeoutSec),
		UnhealthyThreshold: dcl.Int64OrNil(p.UnhealthyThreshold),
		CreationTimestamp:  dcl.StringOrNil(p.CreationTimestamp),
		Project:            dcl.StringOrNil(p.Project),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
	}
	return obj
}

// HttpHealthCheckToProto converts a HttpHealthCheck resource to its proto representation.
func HttpHealthCheckToProto(resource *beta.HttpHealthCheck) *betapb.ComputeBetaHttpHealthCheck {
	p := &betapb.ComputeBetaHttpHealthCheck{
		CheckIntervalSec:   dcl.ValueOrEmptyInt64(resource.CheckIntervalSec),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(resource.HealthyThreshold),
		Host:               dcl.ValueOrEmptyString(resource.Host),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Port:               dcl.ValueOrEmptyInt64(resource.Port),
		RequestPath:        dcl.ValueOrEmptyString(resource.RequestPath),
		TimeoutSec:         dcl.ValueOrEmptyInt64(resource.TimeoutSec),
		UnhealthyThreshold: dcl.ValueOrEmptyInt64(resource.UnhealthyThreshold),
		CreationTimestamp:  dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
	}

	return p
}

// ApplyHttpHealthCheck handles the gRPC request by passing it to the underlying HttpHealthCheck Apply() method.
func (s *HttpHealthCheckServer) applyHttpHealthCheck(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaHttpHealthCheckRequest) (*betapb.ComputeBetaHttpHealthCheck, error) {
	p := ProtoToHttpHealthCheck(request.GetResource())
	res, err := c.ApplyHttpHealthCheck(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpHealthCheckToProto(res)
	return r, nil
}

// ApplyHttpHealthCheck handles the gRPC request by passing it to the underlying HttpHealthCheck Apply() method.
func (s *HttpHealthCheckServer) ApplyComputeBetaHttpHealthCheck(ctx context.Context, request *betapb.ApplyComputeBetaHttpHealthCheckRequest) (*betapb.ComputeBetaHttpHealthCheck, error) {
	cl, err := createConfigHttpHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHttpHealthCheck(ctx, cl, request)
}

// DeleteHttpHealthCheck handles the gRPC request by passing it to the underlying HttpHealthCheck Delete() method.
func (s *HttpHealthCheckServer) DeleteComputeBetaHttpHealthCheck(ctx context.Context, request *betapb.DeleteComputeBetaHttpHealthCheckRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpHealthCheck(ctx, ProtoToHttpHealthCheck(request.GetResource()))

}

// ListComputeBetaHttpHealthCheck handles the gRPC request by passing it to the underlying HttpHealthCheckList() method.
func (s *HttpHealthCheckServer) ListComputeBetaHttpHealthCheck(ctx context.Context, request *betapb.ListComputeBetaHttpHealthCheckRequest) (*betapb.ListComputeBetaHttpHealthCheckResponse, error) {
	cl, err := createConfigHttpHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpHealthCheck(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaHttpHealthCheck
	for _, r := range resources.Items {
		rp := HttpHealthCheckToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaHttpHealthCheckResponse{Items: protos}, nil
}

func createConfigHttpHealthCheck(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
