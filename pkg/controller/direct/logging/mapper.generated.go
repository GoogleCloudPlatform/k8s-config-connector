// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:mapper
// krm.group: logging.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.logging.v2

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krmloggingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apipb "google.golang.org/genproto/googleapis/api"
	distributionpb "google.golang.org/genproto/googleapis/api/distribution"
	labelpb "google.golang.org/genproto/googleapis/api/label"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
)

func BigQueryDataset_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krmloggingv1alpha1.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &krmloggingv1alpha1.BigQueryDataset{}
	// MISSING: DatasetID
	return out
}
func BigQueryDataset_ToProto(mapCtx *direct.MapContext, in *krmloggingv1alpha1.BigQueryDataset) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	// MISSING: DatasetID
	return out
}
func BigQueryDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDataset) *krmloggingv1alpha1.BigQueryDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krmloggingv1alpha1.BigQueryDatasetObservedState{}
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	return out
}
func BigQueryDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krmloggingv1alpha1.BigQueryDatasetObservedState) *pb.BigQueryDataset {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDataset{}
	out.DatasetId = direct.ValueOf(in.DatasetID)
	return out
}
func Distribution_BucketOptions_FromProto(mapCtx *direct.MapContext, in *distributionpb.Distribution_BucketOptions) *krm.Distribution_BucketOptions {
	if in == nil {
		return nil
	}
	out := &krm.Distribution_BucketOptions{}
	out.LinearBuckets = Distribution_BucketOptions_Linear_FromProto(mapCtx, in.GetLinearBuckets())
	out.ExponentialBuckets = Distribution_BucketOptions_Exponential_FromProto(mapCtx, in.GetExponentialBuckets())
	out.ExplicitBuckets = Distribution_BucketOptions_Explicit_FromProto(mapCtx, in.GetExplicitBuckets())
	return out
}
func Distribution_BucketOptions_ToProto(mapCtx *direct.MapContext, in *krm.Distribution_BucketOptions) *distributionpb.Distribution_BucketOptions {
	if in == nil {
		return nil
	}
	out := &distributionpb.Distribution_BucketOptions{}
	if oneof := Distribution_BucketOptions_Linear_ToProto(mapCtx, in.LinearBuckets); oneof != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_LinearBuckets{LinearBuckets: oneof}
	}
	if oneof := Distribution_BucketOptions_Exponential_ToProto(mapCtx, in.ExponentialBuckets); oneof != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_ExponentialBuckets{ExponentialBuckets: oneof}
	}
	if oneof := Distribution_BucketOptions_Explicit_ToProto(mapCtx, in.ExplicitBuckets); oneof != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_ExplicitBuckets{ExplicitBuckets: oneof}
	}
	return out
}
func Distribution_BucketOptions_Explicit_FromProto(mapCtx *direct.MapContext, in *distributionpb.Distribution_BucketOptions_Explicit) *krm.Distribution_BucketOptions_Explicit {
	if in == nil {
		return nil
	}
	out := &krm.Distribution_BucketOptions_Explicit{}
	out.Bounds = in.Bounds
	return out
}
func Distribution_BucketOptions_Explicit_ToProto(mapCtx *direct.MapContext, in *krm.Distribution_BucketOptions_Explicit) *distributionpb.Distribution_BucketOptions_Explicit {
	if in == nil {
		return nil
	}
	out := &distributionpb.Distribution_BucketOptions_Explicit{}
	out.Bounds = in.Bounds
	return out
}
func Distribution_BucketOptions_Exponential_FromProto(mapCtx *direct.MapContext, in *distributionpb.Distribution_BucketOptions_Exponential) *krm.Distribution_BucketOptions_Exponential {
	if in == nil {
		return nil
	}
	out := &krm.Distribution_BucketOptions_Exponential{}
	out.NumFiniteBuckets = direct.LazyPtr(in.GetNumFiniteBuckets())
	out.GrowthFactor = direct.LazyPtr(in.GetGrowthFactor())
	out.Scale = direct.LazyPtr(in.GetScale())
	return out
}
func Distribution_BucketOptions_Exponential_ToProto(mapCtx *direct.MapContext, in *krm.Distribution_BucketOptions_Exponential) *distributionpb.Distribution_BucketOptions_Exponential {
	if in == nil {
		return nil
	}
	out := &distributionpb.Distribution_BucketOptions_Exponential{}
	out.NumFiniteBuckets = direct.ValueOf(in.NumFiniteBuckets)
	out.GrowthFactor = direct.ValueOf(in.GrowthFactor)
	out.Scale = direct.ValueOf(in.Scale)
	return out
}
func Distribution_BucketOptions_Linear_FromProto(mapCtx *direct.MapContext, in *distributionpb.Distribution_BucketOptions_Linear) *krm.Distribution_BucketOptions_Linear {
	if in == nil {
		return nil
	}
	out := &krm.Distribution_BucketOptions_Linear{}
	out.NumFiniteBuckets = direct.LazyPtr(in.GetNumFiniteBuckets())
	out.Width = direct.LazyPtr(in.GetWidth())
	out.Offset = direct.LazyPtr(in.GetOffset())
	return out
}
func Distribution_BucketOptions_Linear_ToProto(mapCtx *direct.MapContext, in *krm.Distribution_BucketOptions_Linear) *distributionpb.Distribution_BucketOptions_Linear {
	if in == nil {
		return nil
	}
	out := &distributionpb.Distribution_BucketOptions_Linear{}
	out.NumFiniteBuckets = direct.ValueOf(in.NumFiniteBuckets)
	out.Width = direct.ValueOf(in.Width)
	out.Offset = direct.ValueOf(in.Offset)
	return out
}
func LabelDescriptor_FromProto(mapCtx *direct.MapContext, in *labelpb.LabelDescriptor) *krm.LabelDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.LabelDescriptor{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.ValueType = direct.Enum_FromProto(mapCtx, in.GetValueType())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func LabelDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.LabelDescriptor) *labelpb.LabelDescriptor {
	if in == nil {
		return nil
	}
	out := &labelpb.LabelDescriptor{}
	out.Key = direct.ValueOf(in.Key)
	out.ValueType = direct.Enum_ToProto[labelpb.LabelDescriptor_ValueType](mapCtx, in.ValueType)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func LoggingLinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krmloggingv1alpha1.LoggingLinkObservedState {
	if in == nil {
		return nil
	}
	out := &krmloggingv1alpha1.LoggingLinkObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	out.BigqueryDataset = BigQueryDatasetObservedState_FromProto(mapCtx, in.GetBigqueryDataset())
	return out
}
func LoggingLinkObservedState_ToProto(mapCtx *direct.MapContext, in *krmloggingv1alpha1.LoggingLinkObservedState) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LifecycleState = direct.Enum_ToProto[pb.LifecycleState](mapCtx, in.LifecycleState)
	out.BigqueryDataset = BigQueryDatasetObservedState_ToProto(mapCtx, in.BigqueryDataset)
	return out
}
func LoggingLinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Link) *krmloggingv1alpha1.LoggingLinkSpec {
	if in == nil {
		return nil
	}
	out := &krmloggingv1alpha1.LoggingLinkSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func LoggingLinkSpec_ToProto(mapCtx *direct.MapContext, in *krmloggingv1alpha1.LoggingLinkSpec) *pb.Link {
	if in == nil {
		return nil
	}
	out := &pb.Link{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	return out
}
func LoggingLogMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LoggingLogMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	// MISSING: BucketName
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MetricDescriptor = MetricDescriptor_FromProto(mapCtx, in.GetMetricDescriptor())
	out.ValueExtractor = direct.LazyPtr(in.GetValueExtractor())
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = Distribution_BucketOptions_FromProto(mapCtx, in.GetBucketOptions())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Version
	return out
}
func LoggingLogMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogMetricSpec) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	// MISSING: BucketName
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MetricDescriptor = MetricDescriptor_ToProto(mapCtx, in.MetricDescriptor)
	out.ValueExtractor = direct.ValueOf(in.ValueExtractor)
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = Distribution_BucketOptions_ToProto(mapCtx, in.BucketOptions)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Version
	return out
}
func LoggingLogMetricStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LoggingLogMetricStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricStatus{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	out.MetricDescriptor = LogmetricMetricDescriptorStatus_FromProto(mapCtx, in.GetMetricDescriptor())
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Version
	return out
}
func LoggingLogMetricStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogMetricStatus) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	out.MetricDescriptor = LogmetricMetricDescriptorStatus_ToProto(mapCtx, in.MetricDescriptor)
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Version
	return out
}
func MetricDescriptor_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.MetricDescriptor{}
	// MISSING: Name
	// MISSING: Type
	out.Labels = direct.Slice_FromProto(mapCtx, in.Labels, LabelDescriptor_FromProto)
	out.MetricKind = direct.Enum_FromProto(mapCtx, in.GetMetricKind())
	out.ValueType = direct.Enum_FromProto(mapCtx, in.GetValueType())
	out.Unit = direct.LazyPtr(in.GetUnit())
	// MISSING: Description
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Metadata = MetricDescriptor_MetricDescriptorMetadata_FromProto(mapCtx, in.GetMetadata())
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	// MISSING: MonitoredResourceTypes
	return out
}
func MetricDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.MetricDescriptor) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	// MISSING: Name
	// MISSING: Type
	out.Labels = direct.Slice_ToProto(mapCtx, in.Labels, LabelDescriptor_ToProto)
	out.MetricKind = direct.Enum_ToProto[metricpb.MetricDescriptor_MetricKind](mapCtx, in.MetricKind)
	out.ValueType = direct.Enum_ToProto[metricpb.MetricDescriptor_ValueType](mapCtx, in.ValueType)
	out.Unit = direct.ValueOf(in.Unit)
	// MISSING: Description
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Metadata = MetricDescriptor_MetricDescriptorMetadata_ToProto(mapCtx, in.Metadata)
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	// MISSING: MonitoredResourceTypes
	return out
}
func MetricDescriptor_MetricDescriptorMetadata_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor_MetricDescriptorMetadata) *krm.MetricDescriptor_MetricDescriptorMetadata {
	if in == nil {
		return nil
	}
	out := &krm.MetricDescriptor_MetricDescriptorMetadata{}
	// MISSING: LaunchStage
	out.SamplePeriod = direct.StringDuration_FromProto(mapCtx, in.GetSamplePeriod())
	out.IngestDelay = direct.StringDuration_FromProto(mapCtx, in.GetIngestDelay())
	// MISSING: TimeSeriesResourceHierarchyLevel
	return out
}
func MetricDescriptor_MetricDescriptorMetadata_ToProto(mapCtx *direct.MapContext, in *krm.MetricDescriptor_MetricDescriptorMetadata) *metricpb.MetricDescriptor_MetricDescriptorMetadata {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor_MetricDescriptorMetadata{}
	// MISSING: LaunchStage
	out.SamplePeriod = direct.StringDuration_ToProto(mapCtx, in.SamplePeriod)
	out.IngestDelay = direct.StringDuration_ToProto(mapCtx, in.IngestDelay)
	// MISSING: TimeSeriesResourceHierarchyLevel
	return out
}
