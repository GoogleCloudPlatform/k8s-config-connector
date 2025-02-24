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
	loggingpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/logging_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
)

// LogMetricServer implements the gRPC interface for LogMetric.
type LogMetricServer struct{}

// ProtoToLogMetricMetricDescriptorLabelsValueTypeEnum converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorLabelsValueTypeEnum(e loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum) *logging.LogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorLabelsValueTypeEnum(n[len("LoggingLogMetricMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorMetricKindEnum converts a LogMetricMetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorMetricKindEnum(e loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum) *logging.LogMetricMetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorMetricKindEnum(n[len("LoggingLogMetricMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorValueTypeEnum converts a LogMetricMetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorValueTypeEnum(e loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum) *logging.LogMetricMetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorValueTypeEnum(n[len("LoggingLogMetricMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorLaunchStageEnum converts a LogMetricMetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorLaunchStageEnum(e loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum) *logging.LogMetricMetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := logging.LogMetricMetricDescriptorLaunchStageEnum(n[len("LoggingLogMetricMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptor converts a LogMetricMetricDescriptor object from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptor(p *loggingpb.LoggingLogMetricMetricDescriptor) *logging.LogMetricMetricDescriptor {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptor{
		Name:        dcl.StringOrNil(p.GetName()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToLoggingLogMetricMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToLoggingLogMetricMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToLoggingLogMetricMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToLoggingLogMetricMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToLoggingLogMetricMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorLabels converts a LogMetricMetricDescriptorLabels object from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorLabels(p *loggingpb.LoggingLogMetricMetricDescriptorLabels) *logging.LogMetricMetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToLoggingLogMetricMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorMetadata converts a LogMetricMetricDescriptorMetadata object from its proto representation.
func ProtoToLoggingLogMetricMetricDescriptorMetadata(p *loggingpb.LoggingLogMetricMetricDescriptorMetadata) *logging.LogMetricMetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricMetricDescriptorMetadata{
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToLogMetricBucketOptions converts a LogMetricBucketOptions object from its proto representation.
func ProtoToLoggingLogMetricBucketOptions(p *loggingpb.LoggingLogMetricBucketOptions) *logging.LogMetricBucketOptions {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptions{
		LinearBuckets:      ProtoToLoggingLogMetricBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToLoggingLogMetricBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToLoggingLogMetricBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsLinearBuckets converts a LogMetricBucketOptionsLinearBuckets object from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsLinearBuckets(p *loggingpb.LoggingLogMetricBucketOptionsLinearBuckets) *logging.LogMetricBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		Width:            dcl.Float64OrNil(p.GetWidth()),
		Offset:           dcl.Float64OrNil(p.GetOffset()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExponentialBuckets converts a LogMetricBucketOptionsExponentialBuckets object from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsExponentialBuckets(p *loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets) *logging.LogMetricBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		GrowthFactor:     dcl.Float64OrNil(p.GetGrowthFactor()),
		Scale:            dcl.Float64OrNil(p.GetScale()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExplicitBuckets converts a LogMetricBucketOptionsExplicitBuckets object from its proto representation.
func ProtoToLoggingLogMetricBucketOptionsExplicitBuckets(p *loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets) *logging.LogMetricBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &logging.LogMetricBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToLogMetric converts a LogMetric resource from its proto representation.
func ProtoToLogMetric(p *loggingpb.LoggingLogMetric) *logging.LogMetric {
	obj := &logging.LogMetric{
		Name:             dcl.StringOrNil(p.GetName()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		Filter:           dcl.StringOrNil(p.GetFilter()),
		Disabled:         dcl.Bool(p.GetDisabled()),
		MetricDescriptor: ProtoToLoggingLogMetricMetricDescriptor(p.GetMetricDescriptor()),
		ValueExtractor:   dcl.StringOrNil(p.GetValueExtractor()),
		BucketOptions:    ProtoToLoggingLogMetricBucketOptions(p.GetBucketOptions()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// LogMetricMetricDescriptorLabelsValueTypeEnumToProto converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorLabelsValueTypeEnumToProto(e *logging.LogMetricMetricDescriptorLabelsValueTypeEnum) loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum_value["LogMetricMetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorLabelsValueTypeEnum(0)
}

// LogMetricMetricDescriptorMetricKindEnumToProto converts a LogMetricMetricDescriptorMetricKindEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorMetricKindEnumToProto(e *logging.LogMetricMetricDescriptorMetricKindEnum) loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum_value["LogMetricMetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorMetricKindEnum(0)
}

// LogMetricMetricDescriptorValueTypeEnumToProto converts a LogMetricMetricDescriptorValueTypeEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorValueTypeEnumToProto(e *logging.LogMetricMetricDescriptorValueTypeEnum) loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum_value["LogMetricMetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorValueTypeEnum(0)
}

// LogMetricMetricDescriptorLaunchStageEnumToProto converts a LogMetricMetricDescriptorLaunchStageEnum enum to its proto representation.
func LoggingLogMetricMetricDescriptorLaunchStageEnumToProto(e *logging.LogMetricMetricDescriptorLaunchStageEnum) loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum {
	if e == nil {
		return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum_value["LogMetricMetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(v)
	}
	return loggingpb.LoggingLogMetricMetricDescriptorLaunchStageEnum(0)
}

// LogMetricMetricDescriptorToProto converts a LogMetricMetricDescriptor object to its proto representation.
func LoggingLogMetricMetricDescriptorToProto(o *logging.LogMetricMetricDescriptor) *loggingpb.LoggingLogMetricMetricDescriptor {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptor{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMetricKind(LoggingLogMetricMetricDescriptorMetricKindEnumToProto(o.MetricKind))
	p.SetValueType(LoggingLogMetricMetricDescriptorValueTypeEnumToProto(o.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(o.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetMetadata(LoggingLogMetricMetricDescriptorMetadataToProto(o.Metadata))
	p.SetLaunchStage(LoggingLogMetricMetricDescriptorLaunchStageEnumToProto(o.LaunchStage))
	sLabels := make([]*loggingpb.LoggingLogMetricMetricDescriptorLabels, len(o.Labels))
	for i, r := range o.Labels {
		sLabels[i] = LoggingLogMetricMetricDescriptorLabelsToProto(&r)
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
func LoggingLogMetricMetricDescriptorLabelsToProto(o *logging.LogMetricMetricDescriptorLabels) *loggingpb.LoggingLogMetricMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(LoggingLogMetricMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// LogMetricMetricDescriptorMetadataToProto converts a LogMetricMetricDescriptorMetadata object to its proto representation.
func LoggingLogMetricMetricDescriptorMetadataToProto(o *logging.LogMetricMetricDescriptorMetadata) *loggingpb.LoggingLogMetricMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricMetricDescriptorMetadata{}
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// LogMetricBucketOptionsToProto converts a LogMetricBucketOptions object to its proto representation.
func LoggingLogMetricBucketOptionsToProto(o *logging.LogMetricBucketOptions) *loggingpb.LoggingLogMetricBucketOptions {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptions{}
	p.SetLinearBuckets(LoggingLogMetricBucketOptionsLinearBucketsToProto(o.LinearBuckets))
	p.SetExponentialBuckets(LoggingLogMetricBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets))
	p.SetExplicitBuckets(LoggingLogMetricBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets))
	return p
}

// LogMetricBucketOptionsLinearBucketsToProto converts a LogMetricBucketOptionsLinearBuckets object to its proto representation.
func LoggingLogMetricBucketOptionsLinearBucketsToProto(o *logging.LogMetricBucketOptionsLinearBuckets) *loggingpb.LoggingLogMetricBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsLinearBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetWidth(dcl.ValueOrEmptyDouble(o.Width))
	p.SetOffset(dcl.ValueOrEmptyDouble(o.Offset))
	return p
}

// LogMetricBucketOptionsExponentialBucketsToProto converts a LogMetricBucketOptionsExponentialBuckets object to its proto representation.
func LoggingLogMetricBucketOptionsExponentialBucketsToProto(o *logging.LogMetricBucketOptionsExponentialBuckets) *loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsExponentialBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetGrowthFactor(dcl.ValueOrEmptyDouble(o.GrowthFactor))
	p.SetScale(dcl.ValueOrEmptyDouble(o.Scale))
	return p
}

// LogMetricBucketOptionsExplicitBucketsToProto converts a LogMetricBucketOptionsExplicitBuckets object to its proto representation.
func LoggingLogMetricBucketOptionsExplicitBucketsToProto(o *logging.LogMetricBucketOptionsExplicitBuckets) *loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &loggingpb.LoggingLogMetricBucketOptionsExplicitBuckets{}
	sBounds := make([]float64, len(o.Bounds))
	for i, r := range o.Bounds {
		sBounds[i] = r
	}
	p.SetBounds(sBounds)
	return p
}

// LogMetricToProto converts a LogMetric resource to its proto representation.
func LogMetricToProto(resource *logging.LogMetric) *loggingpb.LoggingLogMetric {
	p := &loggingpb.LoggingLogMetric{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetMetricDescriptor(LoggingLogMetricMetricDescriptorToProto(resource.MetricDescriptor))
	p.SetValueExtractor(dcl.ValueOrEmptyString(resource.ValueExtractor))
	p.SetBucketOptions(LoggingLogMetricBucketOptionsToProto(resource.BucketOptions))
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
func (s *LogMetricServer) applyLogMetric(ctx context.Context, c *logging.Client, request *loggingpb.ApplyLoggingLogMetricRequest) (*loggingpb.LoggingLogMetric, error) {
	p := ProtoToLogMetric(request.GetResource())
	res, err := c.ApplyLogMetric(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogMetricToProto(res)
	return r, nil
}

// applyLoggingLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) ApplyLoggingLogMetric(ctx context.Context, request *loggingpb.ApplyLoggingLogMetricRequest) (*loggingpb.LoggingLogMetric, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogMetric(ctx, cl, request)
}

// DeleteLogMetric handles the gRPC request by passing it to the underlying LogMetric Delete() method.
func (s *LogMetricServer) DeleteLoggingLogMetric(ctx context.Context, request *loggingpb.DeleteLoggingLogMetricRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogMetric(ctx, ProtoToLogMetric(request.GetResource()))

}

// ListLoggingLogMetric handles the gRPC request by passing it to the underlying LogMetricList() method.
func (s *LogMetricServer) ListLoggingLogMetric(ctx context.Context, request *loggingpb.ListLoggingLogMetricRequest) (*loggingpb.ListLoggingLogMetricResponse, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogMetric(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*loggingpb.LoggingLogMetric
	for _, r := range resources.Items {
		rp := LogMetricToProto(r)
		protos = append(protos, rp)
	}
	p := &loggingpb.ListLoggingLogMetricResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogMetric(ctx context.Context, service_account_file string) (*logging.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return logging.NewClient(conf), nil
}
