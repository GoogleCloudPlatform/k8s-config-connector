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

// MetricDescriptorServer implements the gRPC interface for MetricDescriptor.
type MetricDescriptorServer struct{}

// ProtoToMetricDescriptorLabelsValueTypeEnum converts a MetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorLabelsValueTypeEnum(e betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum) *beta.MetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := beta.MetricDescriptorLabelsValueTypeEnum(n[len("MonitoringBetaMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetricKindEnum converts a MetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorMetricKindEnum(e betapb.MonitoringBetaMetricDescriptorMetricKindEnum) *beta.MetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := beta.MetricDescriptorMetricKindEnum(n[len("MonitoringBetaMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorValueTypeEnum converts a MetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorValueTypeEnum(e betapb.MonitoringBetaMetricDescriptorValueTypeEnum) *beta.MetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := beta.MetricDescriptorValueTypeEnum(n[len("MonitoringBetaMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetadataLaunchStageEnum converts a MetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorMetadataLaunchStageEnum(e betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum) *beta.MetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := beta.MetricDescriptorMetadataLaunchStageEnum(n[len("MonitoringBetaMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLaunchStageEnum converts a MetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorLaunchStageEnum(e betapb.MonitoringBetaMetricDescriptorLaunchStageEnum) *beta.MetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := beta.MetricDescriptorLaunchStageEnum(n[len("MonitoringBetaMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLabels converts a MetricDescriptorLabels object from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorLabels(p *betapb.MonitoringBetaMetricDescriptorLabels) *beta.MetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &beta.MetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToMonitoringBetaMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToMetricDescriptorMetadata converts a MetricDescriptorMetadata object from its proto representation.
func ProtoToMonitoringBetaMetricDescriptorMetadata(p *betapb.MonitoringBetaMetricDescriptorMetadata) *beta.MetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.MetricDescriptorMetadata{
		LaunchStage:  ProtoToMonitoringBetaMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToMetricDescriptor converts a MetricDescriptor resource from its proto representation.
func ProtoToMetricDescriptor(p *betapb.MonitoringBetaMetricDescriptor) *beta.MetricDescriptor {
	obj := &beta.MetricDescriptor{
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToMonitoringBetaMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToMonitoringBetaMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToMonitoringBetaMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToMonitoringBetaMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToMonitoringBetaMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// MetricDescriptorLabelsValueTypeEnumToProto converts a MetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func MonitoringBetaMetricDescriptorLabelsValueTypeEnumToProto(e *beta.MetricDescriptorLabelsValueTypeEnum) betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum_value["MetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum(v)
	}
	return betapb.MonitoringBetaMetricDescriptorLabelsValueTypeEnum(0)
}

// MetricDescriptorMetricKindEnumToProto converts a MetricDescriptorMetricKindEnum enum to its proto representation.
func MonitoringBetaMetricDescriptorMetricKindEnumToProto(e *beta.MetricDescriptorMetricKindEnum) betapb.MonitoringBetaMetricDescriptorMetricKindEnum {
	if e == nil {
		return betapb.MonitoringBetaMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := betapb.MonitoringBetaMetricDescriptorMetricKindEnum_value["MetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaMetricDescriptorMetricKindEnum(v)
	}
	return betapb.MonitoringBetaMetricDescriptorMetricKindEnum(0)
}

// MetricDescriptorValueTypeEnumToProto converts a MetricDescriptorValueTypeEnum enum to its proto representation.
func MonitoringBetaMetricDescriptorValueTypeEnumToProto(e *beta.MetricDescriptorValueTypeEnum) betapb.MonitoringBetaMetricDescriptorValueTypeEnum {
	if e == nil {
		return betapb.MonitoringBetaMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := betapb.MonitoringBetaMetricDescriptorValueTypeEnum_value["MetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaMetricDescriptorValueTypeEnum(v)
	}
	return betapb.MonitoringBetaMetricDescriptorValueTypeEnum(0)
}

// MetricDescriptorMetadataLaunchStageEnumToProto converts a MetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func MonitoringBetaMetricDescriptorMetadataLaunchStageEnumToProto(e *beta.MetricDescriptorMetadataLaunchStageEnum) betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum_value["MetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return betapb.MonitoringBetaMetricDescriptorMetadataLaunchStageEnum(0)
}

// MetricDescriptorLaunchStageEnumToProto converts a MetricDescriptorLaunchStageEnum enum to its proto representation.
func MonitoringBetaMetricDescriptorLaunchStageEnumToProto(e *beta.MetricDescriptorLaunchStageEnum) betapb.MonitoringBetaMetricDescriptorLaunchStageEnum {
	if e == nil {
		return betapb.MonitoringBetaMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := betapb.MonitoringBetaMetricDescriptorLaunchStageEnum_value["MetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaMetricDescriptorLaunchStageEnum(v)
	}
	return betapb.MonitoringBetaMetricDescriptorLaunchStageEnum(0)
}

// MetricDescriptorLabelsToProto converts a MetricDescriptorLabels object to its proto representation.
func MonitoringBetaMetricDescriptorLabelsToProto(o *beta.MetricDescriptorLabels) *betapb.MonitoringBetaMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(MonitoringBetaMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// MetricDescriptorMetadataToProto converts a MetricDescriptorMetadata object to its proto representation.
func MonitoringBetaMetricDescriptorMetadataToProto(o *beta.MetricDescriptorMetadata) *betapb.MonitoringBetaMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaMetricDescriptorMetadata{}
	p.SetLaunchStage(MonitoringBetaMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage))
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// MetricDescriptorToProto converts a MetricDescriptor resource to its proto representation.
func MetricDescriptorToProto(resource *beta.MetricDescriptor) *betapb.MonitoringBetaMetricDescriptor {
	p := &betapb.MonitoringBetaMetricDescriptor{}
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetMetricKind(MonitoringBetaMetricDescriptorMetricKindEnumToProto(resource.MetricKind))
	p.SetValueType(MonitoringBetaMetricDescriptorValueTypeEnumToProto(resource.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(resource.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetMetadata(MonitoringBetaMetricDescriptorMetadataToProto(resource.Metadata))
	p.SetLaunchStage(MonitoringBetaMetricDescriptorLaunchStageEnumToProto(resource.LaunchStage))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sLabels := make([]*betapb.MonitoringBetaMetricDescriptorLabels, len(resource.Labels))
	for i, r := range resource.Labels {
		sLabels[i] = MonitoringBetaMetricDescriptorLabelsToProto(&r)
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
func (s *MetricDescriptorServer) applyMetricDescriptor(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaMetricDescriptorRequest) (*betapb.MonitoringBetaMetricDescriptor, error) {
	p := ProtoToMetricDescriptor(request.GetResource())
	res, err := c.ApplyMetricDescriptor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricDescriptorToProto(res)
	return r, nil
}

// applyMonitoringBetaMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) ApplyMonitoringBetaMetricDescriptor(ctx context.Context, request *betapb.ApplyMonitoringBetaMetricDescriptorRequest) (*betapb.MonitoringBetaMetricDescriptor, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetricDescriptor(ctx, cl, request)
}

// DeleteMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Delete() method.
func (s *MetricDescriptorServer) DeleteMonitoringBetaMetricDescriptor(ctx context.Context, request *betapb.DeleteMonitoringBetaMetricDescriptorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetricDescriptor(ctx, ProtoToMetricDescriptor(request.GetResource()))

}

// ListMonitoringBetaMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptorList() method.
func (s *MetricDescriptorServer) ListMonitoringBetaMetricDescriptor(ctx context.Context, request *betapb.ListMonitoringBetaMetricDescriptorRequest) (*betapb.ListMonitoringBetaMetricDescriptorResponse, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetricDescriptor(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaMetricDescriptor
	for _, r := range resources.Items {
		rp := MetricDescriptorToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaMetricDescriptorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetricDescriptor(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
