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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/beta/monitoring_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
)

// MonitoredProjectServer implements the gRPC interface for MonitoredProject.
type MonitoredProjectServer struct{}

// ProtoToMonitoredProject converts a MonitoredProject resource from its proto representation.
func ProtoToMonitoredProject(p *betapb.MonitoringBetaMonitoredProject) *beta.MonitoredProject {
	obj := &beta.MonitoredProject{
		Name:         dcl.StringOrNil(p.GetName()),
		CreateTime:   dcl.StringOrNil(p.GetCreateTime()),
		MetricsScope: dcl.StringOrNil(p.GetMetricsScope()),
	}
	return obj
}

// MonitoredProjectToProto converts a MonitoredProject resource to its proto representation.
func MonitoredProjectToProto(resource *beta.MonitoredProject) *betapb.MonitoringBetaMonitoredProject {
	p := &betapb.MonitoringBetaMonitoredProject{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetMetricsScope(dcl.ValueOrEmptyString(resource.MetricsScope))

	return p
}

// applyMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Apply() method.
func (s *MonitoredProjectServer) applyMonitoredProject(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaMonitoredProjectRequest) (*betapb.MonitoringBetaMonitoredProject, error) {
	p := ProtoToMonitoredProject(request.GetResource())
	res, err := c.ApplyMonitoredProject(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MonitoredProjectToProto(res)
	return r, nil
}

// applyMonitoringBetaMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Apply() method.
func (s *MonitoredProjectServer) ApplyMonitoringBetaMonitoredProject(ctx context.Context, request *betapb.ApplyMonitoringBetaMonitoredProjectRequest) (*betapb.MonitoringBetaMonitoredProject, error) {
	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMonitoredProject(ctx, cl, request)
}

// DeleteMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Delete() method.
func (s *MonitoredProjectServer) DeleteMonitoringBetaMonitoredProject(ctx context.Context, request *betapb.DeleteMonitoringBetaMonitoredProjectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMonitoredProject(ctx, ProtoToMonitoredProject(request.GetResource()))

}

// ListMonitoringBetaMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProjectList() method.
func (s *MonitoredProjectServer) ListMonitoringBetaMonitoredProject(ctx context.Context, request *betapb.ListMonitoringBetaMonitoredProjectRequest) (*betapb.ListMonitoringBetaMonitoredProjectResponse, error) {
	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMonitoredProject(ctx, request.GetMetricsScope())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaMonitoredProject
	for _, r := range resources.Items {
		rp := MonitoredProjectToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaMonitoredProjectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMonitoredProject(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
