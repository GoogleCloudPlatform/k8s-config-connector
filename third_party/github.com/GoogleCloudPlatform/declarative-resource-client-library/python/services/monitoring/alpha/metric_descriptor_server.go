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

// MetricDescriptorServer implements the gRPC interface for MetricDescriptor.
type MetricDescriptorServer struct{}

// ProtoToMetricDescriptorLabelsValueTypeEnum converts a MetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLabelsValueTypeEnum(e alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum) *alpha.MetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorLabelsValueTypeEnum(n[len("MonitoringAlphaMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetricKindEnum converts a MetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetricKindEnum(e alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum) *alpha.MetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorMetricKindEnum(n[len("MonitoringAlphaMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorValueTypeEnum converts a MetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorValueTypeEnum(e alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum) *alpha.MetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorValueTypeEnum(n[len("MonitoringAlphaMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorMetadataLaunchStageEnum converts a MetricDescriptorMetadataLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(e alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum) *alpha.MetricDescriptorMetadataLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorMetadataLaunchStageEnum(n[len("MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLaunchStageEnum converts a MetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLaunchStageEnum(e alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum) *alpha.MetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.MetricDescriptorLaunchStageEnum(n[len("MonitoringAlphaMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetricDescriptorLabels converts a MetricDescriptorLabels object from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorLabels(p *alphapb.MonitoringAlphaMetricDescriptorLabels) *alpha.MetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.MetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToMonitoringAlphaMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToMetricDescriptorMetadata converts a MetricDescriptorMetadata object from its proto representation.
func ProtoToMonitoringAlphaMetricDescriptorMetadata(p *alphapb.MonitoringAlphaMetricDescriptorMetadata) *alpha.MetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.MetricDescriptorMetadata{
		LaunchStage:  ProtoToMonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(p.GetLaunchStage()),
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToMetricDescriptor converts a MetricDescriptor resource from its proto representation.
func ProtoToMetricDescriptor(p *alphapb.MonitoringAlphaMetricDescriptor) *alpha.MetricDescriptor {
	obj := &alpha.MetricDescriptor{
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToMonitoringAlphaMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToMonitoringAlphaMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToMonitoringAlphaMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToMonitoringAlphaMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToMonitoringAlphaMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// MetricDescriptorLabelsValueTypeEnumToProto converts a MetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorLabelsValueTypeEnumToProto(e *alpha.MetricDescriptorLabelsValueTypeEnum) alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum_value["MetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorLabelsValueTypeEnum(0)
}

// MetricDescriptorMetricKindEnumToProto converts a MetricDescriptorMetricKindEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorMetricKindEnumToProto(e *alpha.MetricDescriptorMetricKindEnum) alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum_value["MetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorMetricKindEnum(0)
}

// MetricDescriptorValueTypeEnumToProto converts a MetricDescriptorValueTypeEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorValueTypeEnumToProto(e *alpha.MetricDescriptorValueTypeEnum) alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum_value["MetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorValueTypeEnum(0)
}

// MetricDescriptorMetadataLaunchStageEnumToProto converts a MetricDescriptorMetadataLaunchStageEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorMetadataLaunchStageEnumToProto(e *alpha.MetricDescriptorMetadataLaunchStageEnum) alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum_value["MetricDescriptorMetadataLaunchStageEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorMetadataLaunchStageEnum(0)
}

// MetricDescriptorLaunchStageEnumToProto converts a MetricDescriptorLaunchStageEnum enum to its proto representation.
func MonitoringAlphaMetricDescriptorLaunchStageEnumToProto(e *alpha.MetricDescriptorLaunchStageEnum) alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum {
	if e == nil {
		return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum_value["MetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(v)
	}
	return alphapb.MonitoringAlphaMetricDescriptorLaunchStageEnum(0)
}

// MetricDescriptorLabelsToProto converts a MetricDescriptorLabels object to its proto representation.
func MonitoringAlphaMetricDescriptorLabelsToProto(o *alpha.MetricDescriptorLabels) *alphapb.MonitoringAlphaMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(MonitoringAlphaMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// MetricDescriptorMetadataToProto converts a MetricDescriptorMetadata object to its proto representation.
func MonitoringAlphaMetricDescriptorMetadataToProto(o *alpha.MetricDescriptorMetadata) *alphapb.MonitoringAlphaMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaMetricDescriptorMetadata{}
	p.SetLaunchStage(MonitoringAlphaMetricDescriptorMetadataLaunchStageEnumToProto(o.LaunchStage))
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// MetricDescriptorToProto converts a MetricDescriptor resource to its proto representation.
func MetricDescriptorToProto(resource *alpha.MetricDescriptor) *alphapb.MonitoringAlphaMetricDescriptor {
	p := &alphapb.MonitoringAlphaMetricDescriptor{}
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetType(dcl.ValueOrEmptyString(resource.Type))
	p.SetMetricKind(MonitoringAlphaMetricDescriptorMetricKindEnumToProto(resource.MetricKind))
	p.SetValueType(MonitoringAlphaMetricDescriptorValueTypeEnumToProto(resource.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(resource.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetMetadata(MonitoringAlphaMetricDescriptorMetadataToProto(resource.Metadata))
	p.SetLaunchStage(MonitoringAlphaMetricDescriptorLaunchStageEnumToProto(resource.LaunchStage))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sLabels := make([]*alphapb.MonitoringAlphaMetricDescriptorLabels, len(resource.Labels))
	for i, r := range resource.Labels {
		sLabels[i] = MonitoringAlphaMetricDescriptorLabelsToProto(&r)
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
func (s *MetricDescriptorServer) applyMetricDescriptor(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaMetricDescriptorRequest) (*alphapb.MonitoringAlphaMetricDescriptor, error) {
	p := ProtoToMetricDescriptor(request.GetResource())
	res, err := c.ApplyMetricDescriptor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetricDescriptorToProto(res)
	return r, nil
}

// applyMonitoringAlphaMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Apply() method.
func (s *MetricDescriptorServer) ApplyMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.ApplyMonitoringAlphaMetricDescriptorRequest) (*alphapb.MonitoringAlphaMetricDescriptor, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetricDescriptor(ctx, cl, request)
}

// DeleteMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptor Delete() method.
func (s *MetricDescriptorServer) DeleteMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.DeleteMonitoringAlphaMetricDescriptorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetricDescriptor(ctx, ProtoToMetricDescriptor(request.GetResource()))

}

// ListMonitoringAlphaMetricDescriptor handles the gRPC request by passing it to the underlying MetricDescriptorList() method.
func (s *MetricDescriptorServer) ListMonitoringAlphaMetricDescriptor(ctx context.Context, request *alphapb.ListMonitoringAlphaMetricDescriptorRequest) (*alphapb.ListMonitoringAlphaMetricDescriptorResponse, error) {
	cl, err := createConfigMetricDescriptor(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetricDescriptor(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaMetricDescriptor
	for _, r := range resources.Items {
		rp := MetricDescriptorToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaMetricDescriptorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetricDescriptor(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
