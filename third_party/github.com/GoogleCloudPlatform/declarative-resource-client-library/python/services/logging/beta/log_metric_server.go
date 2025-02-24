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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/beta/logging_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta"
)

// LogMetricServer implements the gRPC interface for LogMetric.
type LogMetricServer struct{}

// ProtoToLogMetricMetricDescriptorLabelsValueTypeEnum converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum(e betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum) *beta.LogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum_name[int32(e)]; ok {
		e := beta.LogMetricMetricDescriptorLabelsValueTypeEnum(n[len("LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorMetricKindEnum converts a LogMetricMetricDescriptorMetricKindEnum enum from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorMetricKindEnum(e betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum) *beta.LogMetricMetricDescriptorMetricKindEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum_name[int32(e)]; ok {
		e := beta.LogMetricMetricDescriptorMetricKindEnum(n[len("LoggingBetaLogMetricMetricDescriptorMetricKindEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorValueTypeEnum converts a LogMetricMetricDescriptorValueTypeEnum enum from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorValueTypeEnum(e betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum) *beta.LogMetricMetricDescriptorValueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum_name[int32(e)]; ok {
		e := beta.LogMetricMetricDescriptorValueTypeEnum(n[len("LoggingBetaLogMetricMetricDescriptorValueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptorLaunchStageEnum converts a LogMetricMetricDescriptorLaunchStageEnum enum from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorLaunchStageEnum(e betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum) *beta.LogMetricMetricDescriptorLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum_name[int32(e)]; ok {
		e := beta.LogMetricMetricDescriptorLaunchStageEnum(n[len("LoggingBetaLogMetricMetricDescriptorLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToLogMetricMetricDescriptor converts a LogMetricMetricDescriptor object from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptor(p *betapb.LoggingBetaLogMetricMetricDescriptor) *beta.LogMetricMetricDescriptor {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricMetricDescriptor{
		Name:        dcl.StringOrNil(p.GetName()),
		Type:        dcl.StringOrNil(p.GetType()),
		MetricKind:  ProtoToLoggingBetaLogMetricMetricDescriptorMetricKindEnum(p.GetMetricKind()),
		ValueType:   ProtoToLoggingBetaLogMetricMetricDescriptorValueTypeEnum(p.GetValueType()),
		Unit:        dcl.StringOrNil(p.GetUnit()),
		Description: dcl.StringOrNil(p.GetDescription()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Metadata:    ProtoToLoggingBetaLogMetricMetricDescriptorMetadata(p.GetMetadata()),
		LaunchStage: ProtoToLoggingBetaLogMetricMetricDescriptorLaunchStageEnum(p.GetLaunchStage()),
	}
	for _, r := range p.GetLabels() {
		obj.Labels = append(obj.Labels, *ProtoToLoggingBetaLogMetricMetricDescriptorLabels(r))
	}
	for _, r := range p.GetMonitoredResourceTypes() {
		obj.MonitoredResourceTypes = append(obj.MonitoredResourceTypes, r)
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorLabels converts a LogMetricMetricDescriptorLabels object from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorLabels(p *betapb.LoggingBetaLogMetricMetricDescriptorLabels) *beta.LogMetricMetricDescriptorLabels {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricMetricDescriptorLabels{
		Key:         dcl.StringOrNil(p.GetKey()),
		ValueType:   ProtoToLoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum(p.GetValueType()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	return obj
}

// ProtoToLogMetricMetricDescriptorMetadata converts a LogMetricMetricDescriptorMetadata object from its proto representation.
func ProtoToLoggingBetaLogMetricMetricDescriptorMetadata(p *betapb.LoggingBetaLogMetricMetricDescriptorMetadata) *beta.LogMetricMetricDescriptorMetadata {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricMetricDescriptorMetadata{
		SamplePeriod: dcl.StringOrNil(p.GetSamplePeriod()),
		IngestDelay:  dcl.StringOrNil(p.GetIngestDelay()),
	}
	return obj
}

// ProtoToLogMetricBucketOptions converts a LogMetricBucketOptions object from its proto representation.
func ProtoToLoggingBetaLogMetricBucketOptions(p *betapb.LoggingBetaLogMetricBucketOptions) *beta.LogMetricBucketOptions {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricBucketOptions{
		LinearBuckets:      ProtoToLoggingBetaLogMetricBucketOptionsLinearBuckets(p.GetLinearBuckets()),
		ExponentialBuckets: ProtoToLoggingBetaLogMetricBucketOptionsExponentialBuckets(p.GetExponentialBuckets()),
		ExplicitBuckets:    ProtoToLoggingBetaLogMetricBucketOptionsExplicitBuckets(p.GetExplicitBuckets()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsLinearBuckets converts a LogMetricBucketOptionsLinearBuckets object from its proto representation.
func ProtoToLoggingBetaLogMetricBucketOptionsLinearBuckets(p *betapb.LoggingBetaLogMetricBucketOptionsLinearBuckets) *beta.LogMetricBucketOptionsLinearBuckets {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricBucketOptionsLinearBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		Width:            dcl.Float64OrNil(p.GetWidth()),
		Offset:           dcl.Float64OrNil(p.GetOffset()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExponentialBuckets converts a LogMetricBucketOptionsExponentialBuckets object from its proto representation.
func ProtoToLoggingBetaLogMetricBucketOptionsExponentialBuckets(p *betapb.LoggingBetaLogMetricBucketOptionsExponentialBuckets) *beta.LogMetricBucketOptionsExponentialBuckets {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricBucketOptionsExponentialBuckets{
		NumFiniteBuckets: dcl.Int64OrNil(p.GetNumFiniteBuckets()),
		GrowthFactor:     dcl.Float64OrNil(p.GetGrowthFactor()),
		Scale:            dcl.Float64OrNil(p.GetScale()),
	}
	return obj
}

// ProtoToLogMetricBucketOptionsExplicitBuckets converts a LogMetricBucketOptionsExplicitBuckets object from its proto representation.
func ProtoToLoggingBetaLogMetricBucketOptionsExplicitBuckets(p *betapb.LoggingBetaLogMetricBucketOptionsExplicitBuckets) *beta.LogMetricBucketOptionsExplicitBuckets {
	if p == nil {
		return nil
	}
	obj := &beta.LogMetricBucketOptionsExplicitBuckets{}
	for _, r := range p.GetBounds() {
		obj.Bounds = append(obj.Bounds, r)
	}
	return obj
}

// ProtoToLogMetric converts a LogMetric resource from its proto representation.
func ProtoToLogMetric(p *betapb.LoggingBetaLogMetric) *beta.LogMetric {
	obj := &beta.LogMetric{
		Name:             dcl.StringOrNil(p.GetName()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		Filter:           dcl.StringOrNil(p.GetFilter()),
		Disabled:         dcl.Bool(p.GetDisabled()),
		MetricDescriptor: ProtoToLoggingBetaLogMetricMetricDescriptor(p.GetMetricDescriptor()),
		ValueExtractor:   dcl.StringOrNil(p.GetValueExtractor()),
		BucketOptions:    ProtoToLoggingBetaLogMetricBucketOptions(p.GetBucketOptions()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Project:          dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// LogMetricMetricDescriptorLabelsValueTypeEnumToProto converts a LogMetricMetricDescriptorLabelsValueTypeEnum enum to its proto representation.
func LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnumToProto(e *beta.LogMetricMetricDescriptorLabelsValueTypeEnum) betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum {
	if e == nil {
		return betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum(0)
	}
	if v, ok := betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum_value["LogMetricMetricDescriptorLabelsValueTypeEnum"+string(*e)]; ok {
		return betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum(v)
	}
	return betapb.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum(0)
}

// LogMetricMetricDescriptorMetricKindEnumToProto converts a LogMetricMetricDescriptorMetricKindEnum enum to its proto representation.
func LoggingBetaLogMetricMetricDescriptorMetricKindEnumToProto(e *beta.LogMetricMetricDescriptorMetricKindEnum) betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum {
	if e == nil {
		return betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum(0)
	}
	if v, ok := betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum_value["LogMetricMetricDescriptorMetricKindEnum"+string(*e)]; ok {
		return betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum(v)
	}
	return betapb.LoggingBetaLogMetricMetricDescriptorMetricKindEnum(0)
}

// LogMetricMetricDescriptorValueTypeEnumToProto converts a LogMetricMetricDescriptorValueTypeEnum enum to its proto representation.
func LoggingBetaLogMetricMetricDescriptorValueTypeEnumToProto(e *beta.LogMetricMetricDescriptorValueTypeEnum) betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum {
	if e == nil {
		return betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum(0)
	}
	if v, ok := betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum_value["LogMetricMetricDescriptorValueTypeEnum"+string(*e)]; ok {
		return betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum(v)
	}
	return betapb.LoggingBetaLogMetricMetricDescriptorValueTypeEnum(0)
}

// LogMetricMetricDescriptorLaunchStageEnumToProto converts a LogMetricMetricDescriptorLaunchStageEnum enum to its proto representation.
func LoggingBetaLogMetricMetricDescriptorLaunchStageEnumToProto(e *beta.LogMetricMetricDescriptorLaunchStageEnum) betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum {
	if e == nil {
		return betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum(0)
	}
	if v, ok := betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum_value["LogMetricMetricDescriptorLaunchStageEnum"+string(*e)]; ok {
		return betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum(v)
	}
	return betapb.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum(0)
}

// LogMetricMetricDescriptorToProto converts a LogMetricMetricDescriptor object to its proto representation.
func LoggingBetaLogMetricMetricDescriptorToProto(o *beta.LogMetricMetricDescriptor) *betapb.LoggingBetaLogMetricMetricDescriptor {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricMetricDescriptor{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetMetricKind(LoggingBetaLogMetricMetricDescriptorMetricKindEnumToProto(o.MetricKind))
	p.SetValueType(LoggingBetaLogMetricMetricDescriptorValueTypeEnumToProto(o.ValueType))
	p.SetUnit(dcl.ValueOrEmptyString(o.Unit))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetMetadata(LoggingBetaLogMetricMetricDescriptorMetadataToProto(o.Metadata))
	p.SetLaunchStage(LoggingBetaLogMetricMetricDescriptorLaunchStageEnumToProto(o.LaunchStage))
	sLabels := make([]*betapb.LoggingBetaLogMetricMetricDescriptorLabels, len(o.Labels))
	for i, r := range o.Labels {
		sLabels[i] = LoggingBetaLogMetricMetricDescriptorLabelsToProto(&r)
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
func LoggingBetaLogMetricMetricDescriptorLabelsToProto(o *beta.LogMetricMetricDescriptorLabels) *betapb.LoggingBetaLogMetricMetricDescriptorLabels {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricMetricDescriptorLabels{}
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValueType(LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnumToProto(o.ValueType))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	return p
}

// LogMetricMetricDescriptorMetadataToProto converts a LogMetricMetricDescriptorMetadata object to its proto representation.
func LoggingBetaLogMetricMetricDescriptorMetadataToProto(o *beta.LogMetricMetricDescriptorMetadata) *betapb.LoggingBetaLogMetricMetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricMetricDescriptorMetadata{}
	p.SetSamplePeriod(dcl.ValueOrEmptyString(o.SamplePeriod))
	p.SetIngestDelay(dcl.ValueOrEmptyString(o.IngestDelay))
	return p
}

// LogMetricBucketOptionsToProto converts a LogMetricBucketOptions object to its proto representation.
func LoggingBetaLogMetricBucketOptionsToProto(o *beta.LogMetricBucketOptions) *betapb.LoggingBetaLogMetricBucketOptions {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricBucketOptions{}
	p.SetLinearBuckets(LoggingBetaLogMetricBucketOptionsLinearBucketsToProto(o.LinearBuckets))
	p.SetExponentialBuckets(LoggingBetaLogMetricBucketOptionsExponentialBucketsToProto(o.ExponentialBuckets))
	p.SetExplicitBuckets(LoggingBetaLogMetricBucketOptionsExplicitBucketsToProto(o.ExplicitBuckets))
	return p
}

// LogMetricBucketOptionsLinearBucketsToProto converts a LogMetricBucketOptionsLinearBuckets object to its proto representation.
func LoggingBetaLogMetricBucketOptionsLinearBucketsToProto(o *beta.LogMetricBucketOptionsLinearBuckets) *betapb.LoggingBetaLogMetricBucketOptionsLinearBuckets {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricBucketOptionsLinearBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetWidth(dcl.ValueOrEmptyDouble(o.Width))
	p.SetOffset(dcl.ValueOrEmptyDouble(o.Offset))
	return p
}

// LogMetricBucketOptionsExponentialBucketsToProto converts a LogMetricBucketOptionsExponentialBuckets object to its proto representation.
func LoggingBetaLogMetricBucketOptionsExponentialBucketsToProto(o *beta.LogMetricBucketOptionsExponentialBuckets) *betapb.LoggingBetaLogMetricBucketOptionsExponentialBuckets {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricBucketOptionsExponentialBuckets{}
	p.SetNumFiniteBuckets(dcl.ValueOrEmptyInt64(o.NumFiniteBuckets))
	p.SetGrowthFactor(dcl.ValueOrEmptyDouble(o.GrowthFactor))
	p.SetScale(dcl.ValueOrEmptyDouble(o.Scale))
	return p
}

// LogMetricBucketOptionsExplicitBucketsToProto converts a LogMetricBucketOptionsExplicitBuckets object to its proto representation.
func LoggingBetaLogMetricBucketOptionsExplicitBucketsToProto(o *beta.LogMetricBucketOptionsExplicitBuckets) *betapb.LoggingBetaLogMetricBucketOptionsExplicitBuckets {
	if o == nil {
		return nil
	}
	p := &betapb.LoggingBetaLogMetricBucketOptionsExplicitBuckets{}
	sBounds := make([]float64, len(o.Bounds))
	for i, r := range o.Bounds {
		sBounds[i] = r
	}
	p.SetBounds(sBounds)
	return p
}

// LogMetricToProto converts a LogMetric resource to its proto representation.
func LogMetricToProto(resource *beta.LogMetric) *betapb.LoggingBetaLogMetric {
	p := &betapb.LoggingBetaLogMetric{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetMetricDescriptor(LoggingBetaLogMetricMetricDescriptorToProto(resource.MetricDescriptor))
	p.SetValueExtractor(dcl.ValueOrEmptyString(resource.ValueExtractor))
	p.SetBucketOptions(LoggingBetaLogMetricBucketOptionsToProto(resource.BucketOptions))
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
func (s *LogMetricServer) applyLogMetric(ctx context.Context, c *beta.Client, request *betapb.ApplyLoggingBetaLogMetricRequest) (*betapb.LoggingBetaLogMetric, error) {
	p := ProtoToLogMetric(request.GetResource())
	res, err := c.ApplyLogMetric(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogMetricToProto(res)
	return r, nil
}

// applyLoggingBetaLogMetric handles the gRPC request by passing it to the underlying LogMetric Apply() method.
func (s *LogMetricServer) ApplyLoggingBetaLogMetric(ctx context.Context, request *betapb.ApplyLoggingBetaLogMetricRequest) (*betapb.LoggingBetaLogMetric, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogMetric(ctx, cl, request)
}

// DeleteLogMetric handles the gRPC request by passing it to the underlying LogMetric Delete() method.
func (s *LogMetricServer) DeleteLoggingBetaLogMetric(ctx context.Context, request *betapb.DeleteLoggingBetaLogMetricRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogMetric(ctx, ProtoToLogMetric(request.GetResource()))

}

// ListLoggingBetaLogMetric handles the gRPC request by passing it to the underlying LogMetricList() method.
func (s *LogMetricServer) ListLoggingBetaLogMetric(ctx context.Context, request *betapb.ListLoggingBetaLogMetricRequest) (*betapb.ListLoggingBetaLogMetricResponse, error) {
	cl, err := createConfigLogMetric(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogMetric(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.LoggingBetaLogMetric
	for _, r := range resources.Items {
		rp := LogMetricToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListLoggingBetaLogMetricResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogMetric(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
