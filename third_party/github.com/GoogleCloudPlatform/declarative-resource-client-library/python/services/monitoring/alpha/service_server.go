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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// ServiceServer implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceCustom converts a ServiceCustom object from its proto representation.
func ProtoToMonitoringAlphaServiceCustom(p *alphapb.MonitoringAlphaServiceCustom) *alpha.ServiceCustom {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceCustom{}
	return obj
}

// ProtoToServiceTelemetry converts a ServiceTelemetry object from its proto representation.
func ProtoToMonitoringAlphaServiceTelemetry(p *alphapb.MonitoringAlphaServiceTelemetry) *alpha.ServiceTelemetry {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTelemetry{
		ResourceName: dcl.StringOrNil(p.GetResourceName()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *alphapb.MonitoringAlphaService) *alpha.Service {
	obj := &alpha.Service{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Custom:      ProtoToMonitoringAlphaServiceCustom(p.GetCustom()),
		Telemetry:   ProtoToMonitoringAlphaServiceTelemetry(p.GetTelemetry()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ServiceCustomToProto converts a ServiceCustom object to its proto representation.
func MonitoringAlphaServiceCustomToProto(o *alpha.ServiceCustom) *alphapb.MonitoringAlphaServiceCustom {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceCustom{}
	return p
}

// ServiceTelemetryToProto converts a ServiceTelemetry object to its proto representation.
func MonitoringAlphaServiceTelemetryToProto(o *alpha.ServiceTelemetry) *alphapb.MonitoringAlphaServiceTelemetry {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceTelemetry{}
	p.SetResourceName(dcl.ValueOrEmptyString(o.ResourceName))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *alpha.Service) *alphapb.MonitoringAlphaService {
	p := &alphapb.MonitoringAlphaService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCustom(MonitoringAlphaServiceCustomToProto(resource.Custom))
	p.SetTelemetry(MonitoringAlphaServiceTelemetryToProto(resource.Telemetry))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mUserLabels := make(map[string]string, len(resource.UserLabels))
	for k, r := range resource.UserLabels {
		mUserLabels[k] = r
	}
	p.SetUserLabels(mUserLabels)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaServiceRequest) (*alphapb.MonitoringAlphaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyMonitoringAlphaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyMonitoringAlphaService(ctx context.Context, request *alphapb.ApplyMonitoringAlphaServiceRequest) (*alphapb.MonitoringAlphaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteMonitoringAlphaService(ctx context.Context, request *alphapb.DeleteMonitoringAlphaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListMonitoringAlphaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListMonitoringAlphaService(ctx context.Context, request *alphapb.ListMonitoringAlphaServiceRequest) (*alphapb.ListMonitoringAlphaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
