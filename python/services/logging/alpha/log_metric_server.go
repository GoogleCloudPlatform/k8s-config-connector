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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/alpha/logging_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha"
)

// LogMetricServer implements the gRPC interface for LogMetric.
type LogMetricServer struct{}

// ProtoToLogMetricMetricDescriptorLabelsValueTypeEnum converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum(e alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum) *alpha.LogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := alpha.LogMetricMetricDescriptorLabelsValueTypeEnum(n[len("LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorMetricKindEnum converts a LogMetricMetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorMetricKindEnum(e alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum) *alpha.LogMetricMetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := alpha.LogMetricMetricDescriptorMetricKindEnum(n[len("LoggingAlphaLogMetricMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorValueTypeEnum converts a LogMetricMetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorValueTypeEnum(e alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum) *alpha.LogMetricMetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := alpha.LogMetricMetricDescriptorValueTypeEnum(n[len("LoggingAlphaLogMetricMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorLaunchStageEnum converts a LogMetricMetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorLaunchStageEnum(e alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum) *alpha.LogMetricMetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.LogMetricMetricDescriptorLaunchStageEnum(n[len("LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptor converts a LogMetricMetricDescriptor object from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptor(p *alphapb.LoggingAlphaLogMetricMetricDescriptor) *alpha.LogMetricMetricDescriptor {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricMetricDescriptor{
		Name:        dcl.StringOrNil(p.GetName()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToLoggingAlphaLogMetricMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToLoggingAlphaLogMetricMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToLoggingAlphaLogMetricMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToLoggingAlphaLogMetricMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToLoggingAlphaLogMetricMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorLabels converts a LogMetricMetricDescriptorLabels object from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorLabels(p *alphapb.LoggingAlphaLogMetricMetricDescriptorLabels) *alpha.LogMetricMetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricMetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToLoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorMetadata converts a LogMetricMetricDescriptorMetadata object from its proto representation.
func ProtoToLoggingAlphaLogMetricMetricDescriptorMetadata(p *alphapb.LoggingAlphaLogMetricMetricDescriptorMetadata) *alpha.LogMetricMetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricMetricDescriptorMetadata{
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToLogMetricBucketOptions converts a LogMetricBucketOptions object from its proto representation.
func ProtoToLoggingAlphaLogMetricBucketOptions(p *alphapb.LoggingAlphaLogMetricBucketOptions) *alpha.LogMetricBucketOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricBucketOptions{
		LinearBuckets:      ProtoToLoggingAlphaLogMetricBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToLoggingAlphaLogMetricBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToLoggingAlphaLogMetricBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsLinearBuckets converts a LogMetricBucketOptionsLinearBuckets object from its proto representation.
func ProtoToLoggingAlphaLogMetricBucketOptionsLinearBuckets(p *alphapb.LoggingAlphaLogMetricBucketOptionsLinearBuckets) *alpha.LogMetricBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		Width:            dcl.Float64OrNil(p.GetWidth()),
		Offset:           dcl.Float64OrNil(p.GetOffset()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExponentialBuckets converts a LogMetricBucketOptionsExponentialBuckets object from its proto representation.
func ProtoToLoggingAlphaLogMetricBucketOptionsExponentialBuckets(p *alphapb.LoggingAlphaLogMetricBucketOptionsExponentialBuckets) *alpha.LogMetricBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		GrowthFactor:     dcl.Float64OrNil(p.GetGrowthFactor()),
		Scale:            dcl.Float64OrNil(p.GetScale()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExplicitBuckets converts a LogMetricBucketOptionsExplicitBuckets object from its proto representation.
func ProtoToLoggingAlphaLogMetricBucketOptionsExplicitBuckets(p *alphapb.LoggingAlphaLogMetricBucketOptionsExplicitBuckets) *alpha.LogMetricBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &alpha.LogMetricBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToLogMetric converts a LogMetric resource from its proto representation.
func ProtoToLogMetric(p *alphapb.LoggingAlphaLogMetric) *alpha.LogMetric {
	obj := &alpha.LogMetric{
		Name:             dcl.StringOrNil(p.GetName()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		Filter:           dcl.StringOrNil(p.GetFilter()),
		Disabled:         dcl.Bool(p.GetDisabled()),
		MetricDescriptor: ProtoToLoggingAlphaLogMetricMetricDescriptor(p.GetMetricDescriptor()),
		ValueExtractor:   dcl.StringOrNil(p.GetValueExtractor()),
		BucketOptions:    ProtoToLoggingAlphaLogMetricBucketOptions(p.GetBucketOptions()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// LogMetricMetricDescriptorLabelsValueTypeEnumToProto converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnumToProto(e *alpha.LogMetricMetricDescriptorLabelsValueTypeEnum) alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum_value["LogMetricMetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum(v)
	}
	return alphapb.LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnum(0)
}

// LogMetricMetricDescriptorMetricKindEnumToProto converts a LogMetricMetricDescriptorMetricKindEnum enum to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorMetricKindEnumToProto(e *alpha.LogMetricMetricDescriptorMetricKindEnum) alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum {
	if e == nil {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum_value["LogMetricMetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum(v)
	}
	return alphapb.LoggingAlphaLogMetricMetricDescriptorMetricKindEnum(0)
}

// LogMetricMetricDescriptorValueTypeEnumToProto converts a LogMetricMetricDescriptorValueTypeEnum enum to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorValueTypeEnumToProto(e *alpha.LogMetricMetricDescriptorValueTypeEnum) alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum {
	if e == nil {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum_value["LogMetricMetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum(v)
	}
	return alphapb.LoggingAlphaLogMetricMetricDescriptorValueTypeEnum(0)
}

// LogMetricMetricDescriptorLaunchStageEnumToProto converts a LogMetricMetricDescriptorLaunchStageEnum enum to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorLaunchStageEnumToProto(e *alpha.LogMetricMetricDescriptorLaunchStageEnum) alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum {
	if e == nil {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum_value["LogMetricMetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum(v)
	}
	return alphapb.LoggingAlphaLogMetricMetricDescriptorLaunchStageEnum(0)
}

// LogMetricMetricDescriptorToProto converts a LogMetricMetricDescriptor object to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorToProto(o *alpha.LogMetricMetricDescriptor) *alphapb.LoggingAlphaLogMetricMetricDescriptor {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricMetricDescriptor{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMetricKind(LoggingAlphaLogMetricMetricDescriptorMetricKindEnumToProto(o.MetricKind))
	p.SetValueType(LoggingAlphaLogMetricMetricDescriptorValueTypeEnumToProto(o.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(o.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetMetadata(LoggingAlphaLogMetricMetricDescriptorMetadataToProto(o.Metadata))
	p.SetLaunchStage(LoggingAlphaLogMetricMetricDescriptorLaunchStageEnumToProto(o.LaunchStage))
	sLabels := make([]*alphapb.LoggingAlphaLogMetricMetricDescriptorLabels, len(o.Labels))
	for i, r := range o.Labels {
		sLabels[i] = LoggingAlphaLogMetricMetricDescriptorLabelsToProto(&r)
	}
	p.SetLabels(sLabels)
	sMonitoredResourceTypes := make([]string, len(o.MonitoredResourceTypes))
	for i, r := range o.MonitoredResourceTypes {
		sMonitoredResourceTypes[i] = r
	}
	p.SetMonitoredResourceTypes(sMonitoredResourceTypes)
	return p
}

// LogMetricMetricDescriptorLabelsToProto converts a LogMetricMetricDescriptorLabels object to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorLabelsToProto(o *alpha.LogMetricMetricDescriptorLabels) *alphapb.LoggingAlphaLogMetricMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(LoggingAlphaLogMetricMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// LogMetricMetricDescriptorMetadataToProto converts a LogMetricMetricDescriptorMetadata object to its proto representation.
func LoggingAlphaLogMetricMetricDescriptorMetadataToProto(o *alpha.LogMetricMetricDescriptorMetadata) *alphapb.LoggingAlphaLogMetricMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricMetricDescriptorMetadata{}
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// LogMetricBucketOptionsToProto converts a LogMetricBucketOptions object to its proto representation.
func LoggingAlphaLogMetricBucketOptionsToProto(o *alpha.LogMetricBucketOptions) *alphapb.LoggingAlphaLogMetricBucketOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricBucketOptions{}
	p.SetLinearBuckets(LoggingAlphaLogMetricBucketOptionsLinearBucketsToProto(o.LinearBuckets))
	p.SetExponentialBuckets(LoggingAlphaLogMetricBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets))
	p.SetExplicitBuckets(LoggingAlphaLogMetricBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets))
	return p
}

// LogMetricBucketOptionsLinearBucketsToProto converts a LogMetricBucketOptionsLinearBuckets object to its proto representation.
func LoggingAlphaLogMetricBucketOptionsLinearBucketsToProto(o *alpha.LogMetricBucketOptionsLinearBuckets) *alphapb.LoggingAlphaLogMetricBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricBucketOptionsLinearBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetWidth(dcl.ValueOrEmptyDouble(o.Width))
	p.SetOffset(dcl.ValueOrEmptyDouble(o.Offset))
	return p
}

// LogMetricBucketOptionsExponentialBucketsToProto converts a LogMetricBucketOptionsExponentialBuckets object to its proto representation.
func LoggingAlphaLogMetricBucketOptionsExponentialBucketsToProto(o *alpha.LogMetricBucketOptionsExponentialBuckets) *alphapb.LoggingAlphaLogMetricBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricBucketOptionsExponentialBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetGrowthFactor(dcl.ValueOrEmptyDouble(o.GrowthFactor))
	p.SetScale(dcl.ValueOrEmptyDouble(o.Scale))
	return p
}

// LogMetricBucketOptionsExplicitBucketsToProto converts a LogMetricBucketOptionsExplicitBuckets object to its proto representation.
func LoggingAlphaLogMetricBucketOptionsExplicitBucketsToProto(o *alpha.LogMetricBucketOptionsExplicitBuckets) *alphapb.LoggingAlphaLogMetricBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &alphapb.LoggingAlphaLogMetricBucketOptionsExplicitBuckets{}
	sBounds := make([]float64, len(o.Bounds))
	for i, r := range o.Bounds {
		sBounds[i] = r
	}
	p.SetBounds(sBounds)
	return p
}

// LogMetricToProto converts a LogMetric resource to its proto representation.
func LogMetricToProto(resource *alpha.LogMetric) *alphapb.LoggingAlphaLogMetric {
	p := &alphapb.LoggingAlphaLogMetric{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetMetricDescriptor(LoggingAlphaLogMetricMetricDescriptorToProto(resource.MetricDescriptor))
	p.SetValueExtractor(dcl.ValueOrEmptyString(resource.ValueExtractor))
	p.SetBucketOptions(LoggingAlphaLogMetricBucketOptionsToProto(resource.BucketOptions))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabelExtractors := make(map[string]string, len(resource.LabelExtractors))
	for k, r := range resource.LabelExtractors {
		mLabelExtractors[k] = r
	}
	p.SetLabelExtractors(mLabelExtractors)

	return p
}

// applyLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) applyLogMetric(ctx context.Context, c *alpha.Client, request *alphapb.ApplyLoggingAlphaLogMetricRequest) (*alphapb.LoggingAlphaLogMetric, error) {
	p := ProtoToLogMetric(request.GetResource())
	res, err := c.ApplyLogMetric(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogMetricToProto(res)
	return r, nil
}

// applyLoggingAlphaLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) ApplyLoggingAlphaLogMetric(ctx context.Context, request *alphapb.ApplyLoggingAlphaLogMetricRequest) (*alphapb.LoggingAlphaLogMetric, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogMetric(ctx, cl, request)
}

// DeleteLogMetric handles the gRPC request by passing it to the underlying LogMetric Delete() method.
func (s *LogMetricServer) DeleteLoggingAlphaLogMetric(ctx context.Context, request *alphapb.DeleteLoggingAlphaLogMetricRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogMetric(ctx, ProtoToLogMetric(request.GetResource()))

}

// ListLoggingAlphaLogMetric handles the gRPC request by passing it to the underlying LogMetricList() method.
func (s *LogMetricServer) ListLoggingAlphaLogMetric(ctx context.Context, request *alphapb.ListLoggingAlphaLogMetricRequest) (*alphapb.ListLoggingAlphaLogMetricResponse, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogMetric(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.LoggingAlphaLogMetric
	for _, r := range resources.Items {
		rp := LogMetricToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListLoggingAlphaLogMetricResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogMetric(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
