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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/beta/monitoring_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
)

// MetricsScopeServer implements the gRPC interface for MetricsScope.
type MetricsScopeServer struct{}

// ProtoToMetricsScopeMonitoredProjects converts a MetricsScopeMonitoredProjects object from its proto representation.
func ProtoToMonitoringBetaMetricsScopeMonitoredProjects(p *betapb.MonitoringBetaMetricsScopeMonitoredProjects) *beta.MetricsScopeMonitoredProjects {
	if p == nil {
		return nil
	}
	obj := &beta.MetricsScopeMonitoredProjects{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToMetricsScope converts a MetricsScope resource from its proto representation.
func ProtoToMetricsScope(p *betapb.MonitoringBetaMetricsScope) *beta.MetricsScope {
	obj := &beta.MetricsScope{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMonitoredProjects() {
		obj.MonitoredProjects = append(obj.MonitoredProjects, *ProtoToMonitoringBetaMetricsScopeMonitoredProjects(r))
	}
	return obj
}

// MetricsScopeMonitoredProjectsToProto converts a MetricsScopeMonitoredProjects object to its proto representation.
func MonitoringBetaMetricsScopeMonitoredProjectsToProto(o *beta.MetricsScopeMonitoredProjects) *betapb.MonitoringBetaMetricsScopeMonitoredProjects {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaMetricsScopeMonitoredProjects{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// MetricsScopeToProto converts a MetricsScope resource to its proto representation.
func MetricsScopeToProto(resource *beta.MetricsScope) *betapb.MonitoringBetaMetricsScope {
	p := &betapb.MonitoringBetaMetricsScope{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	sMonitoredProjects := make([]*betapb.MonitoringBetaMetricsScopeMonitoredProjects, len(resource.MonitoredProjects))
	for i, r := range resource.MonitoredProjects {
		sMonitoredProjects[i] = MonitoringBetaMetricsScopeMonitoredProjectsToProto(&r)
	}
	p.SetMonitoredProjects(sMonitoredProjects)

	return p
}

// applyMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Apply() method.
func (s *MetricsScopeServer) applyMetricsScope(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaMetricsScopeRequest) (*betapb.MonitoringBetaMetricsScope, error) {
	p := ProtoToMetricsScope(request.GetResource())
	res, err := c.ApplyMetricsScope(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricsScopeToProto(res)
	return r, nil
}

// applyMonitoringBetaMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Apply() method.
func (s *MetricsScopeServer) ApplyMonitoringBetaMetricsScope(ctx context.Context, request *betapb.ApplyMonitoringBetaMetricsScopeRequest) (*betapb.MonitoringBetaMetricsScope, error) {
	cl, err := createConfigMetricsScope(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetricsScope(ctx, cl, request)
}

// DeleteMetricsScope handles the gRPC request by passing it to the underlying MetricsScope Delete() method.
func (s *MetricsScopeServer) DeleteMonitoringBetaMetricsScope(ctx context.Context, request *betapb.DeleteMonitoringBetaMetricsScopeRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for MetricsScope")

}

// ListMonitoringBetaMetricsScope is a no-op method because MetricsScope has no list method.
func (s *MetricsScopeServer) ListMonitoringBetaMetricsScope(_ context.Context, _ *betapb.ListMonitoringBetaMetricsScopeRequest) (*betapb.ListMonitoringBetaMetricsScopeResponse, error) {
	return nil, nil
}

func createConfigMetricsScope(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
