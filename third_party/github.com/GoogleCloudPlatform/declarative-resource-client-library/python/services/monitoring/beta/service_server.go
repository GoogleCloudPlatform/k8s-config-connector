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

// ServiceServer implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceCustom converts a ServiceCustom object from its proto representation.
func ProtoToMonitoringBetaServiceCustom(p *betapb.MonitoringBetaServiceCustom) *beta.ServiceCustom {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceCustom{}
	return obj
}

// ProtoToServiceTelemetry converts a ServiceTelemetry object from its proto representation.
func ProtoToMonitoringBetaServiceTelemetry(p *betapb.MonitoringBetaServiceTelemetry) *beta.ServiceTelemetry {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTelemetry{
		ResourceName: dcl.StringOrNil(p.GetResourceName()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *betapb.MonitoringBetaService) *beta.Service {
	obj := &beta.Service{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Custom:      ProtoToMonitoringBetaServiceCustom(p.GetCustom()),
		Telemetry:   ProtoToMonitoringBetaServiceTelemetry(p.GetTelemetry()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ServiceCustomToProto converts a ServiceCustom object to its proto representation.
func MonitoringBetaServiceCustomToProto(o *beta.ServiceCustom) *betapb.MonitoringBetaServiceCustom {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceCustom{}
	return p
}

// ServiceTelemetryToProto converts a ServiceTelemetry object to its proto representation.
func MonitoringBetaServiceTelemetryToProto(o *beta.ServiceTelemetry) *betapb.MonitoringBetaServiceTelemetry {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceTelemetry{}
	p.SetResourceName(dcl.ValueOrEmptyString(o.ResourceName))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *beta.Service) *betapb.MonitoringBetaService {
	p := &betapb.MonitoringBetaService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCustom(MonitoringBetaServiceCustomToProto(resource.Custom))
	p.SetTelemetry(MonitoringBetaServiceTelemetryToProto(resource.Telemetry))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mUserLabels := make(map[string]string, len(resource.UserLabels))
	for k, r := range resource.UserLabels {
		mUserLabels[k] = r
	}
	p.SetUserLabels(mUserLabels)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaServiceRequest) (*betapb.MonitoringBetaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyMonitoringBetaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyMonitoringBetaService(ctx context.Context, request *betapb.ApplyMonitoringBetaServiceRequest) (*betapb.MonitoringBetaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteMonitoringBetaService(ctx context.Context, request *betapb.DeleteMonitoringBetaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListMonitoringBetaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListMonitoringBetaService(ctx context.Context, request *betapb.ListMonitoringBetaServiceRequest) (*betapb.ListMonitoringBetaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
