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

// MetricDescriptorServer implements the gRPC interface for MetricDescriptor.
type MetricDescriptorServer struct{}

// ProtoToMetricDescriptorLabelsValueTypeEnum converts a MetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorLabelsValueTypeEnum(e monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum) *monitoring.MetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorLabelsValueTypeEnum(n[len("MonitoringMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetricKindEnum converts a MetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorMetricKindEnum(e monitoringpb.MonitoringMetricDescriptorMetricKindEnum) *monitoring.MetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorMetricKindEnum(n[len("MonitoringMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorValueTypeEnum converts a MetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorValueTypeEnum(e monitoringpb.MonitoringMetricDescriptorValueTypeEnum) *monitoring.MetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorValueTypeEnum(n[len("MonitoringMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetadataLaunchStageEnum converts a MetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorMetadataLaunchStageEnum(e monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum) *monitoring.MetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorMetadataLaunchStageEnum(n[len("MonitoringMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLaunchStageEnum converts a MetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringMetricDescriptorLaunchStageEnum(e monitoringpb.MonitoringMetricDescriptorLaunchStageEnum) *monitoring.MetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := monitoring.MetricDescriptorLaunchStageEnum(n[len("MonitoringMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLabels converts a MetricDescriptorLabels object from its proto representation.
func ProtoToMonitoringMetricDescriptorLabels(p *monitoringpb.MonitoringMetricDescriptorLabels) *monitoring.MetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &monitoring.MetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToMonitoringMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToMetricDescriptorMetadata converts a MetricDescriptorMetadata object from its proto representation.
func ProtoToMonitoringMetricDescriptorMetadata(p *monitoringpb.MonitoringMetricDescriptorMetadata) *monitoring.MetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &monitoring.MetricDescriptorMetadata{
		LaunchStage:  ProtoToMonitoringMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToMetricDescriptor converts a MetricDescriptor resource from its proto representation.
func ProtoToMetricDescriptor(p *monitoringpb.MonitoringMetricDescriptor) *monitoring.MetricDescriptor {
	obj := &monitoring.MetricDescriptor{
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToMonitoringMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToMonitoringMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToMonitoringMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToMonitoringMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToMonitoringMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// MetricDescriptorLabelsValueTypeEnumToProto converts a MetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func MonitoringMetricDescriptorLabelsValueTypeEnumToProto(e *monitoring.MetricDescriptorLabelsValueTypeEnum) monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum_value["MetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorLabelsValueTypeEnum(0)
}

// MetricDescriptorMetricKindEnumToProto converts a MetricDescriptorMetricKindEnum enum to its proto representation.
func MonitoringMetricDescriptorMetricKindEnumToProto(e *monitoring.MetricDescriptorMetricKindEnum) monitoringpb.MonitoringMetricDescriptorMetricKindEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorMetricKindEnum_value["MetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorMetricKindEnum(0)
}

// MetricDescriptorValueTypeEnumToProto converts a MetricDescriptorValueTypeEnum enum to its proto representation.
func MonitoringMetricDescriptorValueTypeEnumToProto(e *monitoring.MetricDescriptorValueTypeEnum) monitoringpb.MonitoringMetricDescriptorValueTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorValueTypeEnum_value["MetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorValueTypeEnum(0)
}

// MetricDescriptorMetadataLaunchStageEnumToProto converts a MetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func MonitoringMetricDescriptorMetadataLaunchStageEnumToProto(e *monitoring.MetricDescriptorMetadataLaunchStageEnum) monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum_value["MetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorMetadataLaunchStageEnum(0)
}

// MetricDescriptorLaunchStageEnumToProto converts a MetricDescriptorLaunchStageEnum enum to its proto representation.
func MonitoringMetricDescriptorLaunchStageEnumToProto(e *monitoring.MetricDescriptorLaunchStageEnum) monitoringpb.MonitoringMetricDescriptorLaunchStageEnum {
	if e == nil {
		return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := monitoringpb.MonitoringMetricDescriptorLaunchStageEnum_value["MetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(v)
	}
	return monitoringpb.MonitoringMetricDescriptorLaunchStageEnum(0)
}

// MetricDescriptorLabelsToProto converts a MetricDescriptorLabels object to its proto representation.
func MonitoringMetricDescriptorLabelsToProto(o *monitoring.MetricDescriptorLabels) *monitoringpb.MonitoringMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(MonitoringMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// MetricDescriptorMetadataToProto converts a MetricDescriptorMetadata object to its proto representation.
func MonitoringMetricDescriptorMetadataToProto(o *monitoring.MetricDescriptorMetadata) *monitoringpb.MonitoringMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringMetricDescriptorMetadata{}
	p.SetLaunchStage(MonitoringMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage))
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// MetricDescriptorToProto converts a MetricDescriptor resource to its proto representation.
func MetricDescriptorToProto(resource *monitoring.MetricDescriptor) *monitoringpb.MonitoringMetricDescriptor {
	p := &monitoringpb.MonitoringMetricDescriptor{}
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetMetricKind(MonitoringMetricDescriptorMetricKindEnumToProto(resource.MetricKind))
	p.SetValueType(MonitoringMetricDescriptorValueTypeEnumToProto(resource.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(resource.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetMetadata(MonitoringMetricDescriptorMetadataToProto(resource.Metadata))
	p.SetLaunchStage(MonitoringMetricDescriptorLaunchStageEnumToProto(resource.LaunchStage))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sLabels := make([]*monitoringpb.MonitoringMetricDescriptorLabels, len(resource.Labels))
	for i, r := range resource.Labels {
		sLabels[i] = MonitoringMetricDescriptorLabelsToProto(&r)
	}
	p.SetLabels(sLabels)
	sMonitoredResourceTypes := make([]string, len(resource.MonitoredResourceTypes))
	for i, r := range resource.MonitoredResourceTypes {
		sMonitoredResourceTypes[i] = r
	}
	p.SetMonitoredResourceTypes(sMonitoredResourceTypes)

	return p
}

// applyMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) applyMetricDescriptor(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringMetricDescriptorRequest) (*monitoringpb.MonitoringMetricDescriptor, error) {
	p := ProtoToMetricDescriptor(request.GetResource())
	res, err := c.ApplyMetricDescriptor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricDescriptorToProto(res)
	return r, nil
}

// applyMonitoringMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) ApplyMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.ApplyMonitoringMetricDescriptorRequest) (*monitoringpb.MonitoringMetricDescriptor, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetricDescriptor(ctx, cl, request)
}

// DeleteMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Delete() method.
func (s *MetricDescriptorServer) DeleteMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.DeleteMonitoringMetricDescriptorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetricDescriptor(ctx, ProtoToMetricDescriptor(request.GetResource()))

}

// ListMonitoringMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptorList() method.
func (s *MetricDescriptorServer) ListMonitoringMetricDescriptor(ctx context.Context, request *monitoringpb.ListMonitoringMetricDescriptorRequest) (*monitoringpb.ListMonitoringMetricDescriptorResponse, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetricDescriptor(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringMetricDescriptor
	for _, r := range resources.Items {
		rp := MetricDescriptorToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringMetricDescriptorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetricDescriptor(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
