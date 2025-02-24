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
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// MonitoredProjectServer implements the gRPC interface for MonitoredProject.
type MonitoredProjectServer struct{}

// ProtoToMonitoredProject converts a MonitoredProject resource from its proto representation.
func ProtoToMonitoredProject(p *monitoringpb.MonitoringMonitoredProject) *monitoring.MonitoredProject {
	obj := &monitoring.MonitoredProject{
		Name:         dcl.StringOrNil(p.GetName()),
		CreateTime:   dcl.StringOrNil(p.GetCreateTime()),
		MetricsScope: dcl.StringOrNil(p.GetMetricsScope()),
	}
	return obj
}

// MonitoredProjectToProto converts a MonitoredProject resource to its proto representation.
func MonitoredProjectToProto(resource *monitoring.MonitoredProject) *monitoringpb.MonitoringMonitoredProject {
	p := &monitoringpb.MonitoringMonitoredProject{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetMetricsScope(dcl.ValueOrEmptyString(resource.MetricsScope))

	return p
}

// applyMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Apply() method.
func (s *MonitoredProjectServer) applyMonitoredProject(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringMonitoredProjectRequest) (*monitoringpb.MonitoringMonitoredProject, error) {
	p := ProtoToMonitoredProject(request.GetResource())
	res, err := c.ApplyMonitoredProject(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MonitoredProjectToProto(res)
	return r, nil
}

// applyMonitoringMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Apply() method.
func (s *MonitoredProjectServer) ApplyMonitoringMonitoredProject(ctx context.Context, request *monitoringpb.ApplyMonitoringMonitoredProjectRequest) (*monitoringpb.MonitoringMonitoredProject, error) {
	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMonitoredProject(ctx, cl, request)
}

// DeleteMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProject Delete() method.
func (s *MonitoredProjectServer) DeleteMonitoringMonitoredProject(ctx context.Context, request *monitoringpb.DeleteMonitoringMonitoredProjectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMonitoredProject(ctx, ProtoToMonitoredProject(request.GetResource()))

}

// ListMonitoringMonitoredProject handles the gRPC request by passing it to the underlying MonitoredProjectList() method.
func (s *MonitoredProjectServer) ListMonitoringMonitoredProject(ctx context.Context, request *monitoringpb.ListMonitoringMonitoredProjectRequest) (*monitoringpb.ListMonitoringMonitoredProjectResponse, error) {
	cl, err := createConfigMonitoredProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMonitoredProject(ctx, request.GetMetricsScope())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringMonitoredProject
	for _, r := range resources.Items {
		rp := MonitoredProjectToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringMonitoredProjectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMonitoredProject(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
