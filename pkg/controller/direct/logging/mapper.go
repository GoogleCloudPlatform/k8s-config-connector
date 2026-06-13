// Copyright 2026 Google LLC
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

package logging

import (
	"strings"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gapi "google.golang.org/genproto/googleapis/api"
	gdist "google.golang.org/genproto/googleapis/api/distribution"
	glabel "google.golang.org/genproto/googleapis/api/label"
	gmetric "google.golang.org/genproto/googleapis/api/metric"
)

func LoggingLogBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.RetentionDays != 0 {
		days := int64(in.RetentionDays)
		out.RetentionDays = &days
	}
	out.Locked = direct.LazyPtr(in.GetLocked())
	out.EnableAnalytics = direct.LazyPtr(in.GetAnalyticsEnabled())
	return out
}

func LoggingLogBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketSpec) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	out.Description = direct.ValueOf(in.Description)
	if in.RetentionDays != nil {
		out.RetentionDays = int32(*in.RetentionDays)
	}
	out.Locked = direct.ValueOf(in.Locked)
	out.AnalyticsEnabled = direct.ValueOf(in.EnableAnalytics)
	return out
}

func LoggingLogBucketStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	return out
}

func LoggingLogBucketStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketStatus) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LifecycleState = direct.Enum_ToProto[pb.LifecycleState](mapCtx, in.LifecycleState)
	return out
}

func LoggingLogViewStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogView) *krm.LoggingLogViewStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogViewStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func LoggingLogViewStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogViewStatus) *pb.LogView {
	if in == nil {
		return nil
	}
	out := &pb.LogView{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func LoggingLogExclusionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LoggingLogExclusionSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogExclusionSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Filter = in.GetFilter()
	return out
}

func LoggingLogExclusionSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogExclusionSpec) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	out.Description = direct.ValueOf(in.Description)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Filter = in.Filter
	return out
}

func LoggingLogExclusionStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LoggingLogExclusionStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogExclusionStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func LoggingLogExclusionStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogExclusionStatus) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func LoggingLogSinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LoggingLogSinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogSinkSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.IncludeChildren = direct.LazyPtr(in.GetIncludeChildren())

	if in.GetBigqueryOptions() != nil {
		out.BigqueryOptions = &krm.LogSinkBigQueryOptions{
			UsePartitionedTables: in.GetBigqueryOptions().GetUsePartitionedTables(),
		}
	}

	dest := in.GetDestination()
	if dest != "" {
		if strings.HasPrefix(dest, "bigquery.googleapis.com/") {
			val := strings.TrimPrefix(dest, "bigquery.googleapis.com/")
			out.Destination.BigQueryDatasetRef = &bigqueryv1beta1.DatasetRef{External: val}
		} else if strings.HasPrefix(dest, "logging.googleapis.com/") {
			val := strings.TrimPrefix(dest, "logging.googleapis.com/")
			out.Destination.LoggingLogBucketRef = &krm.LoggingLogBucketRef{External: val}
		} else if strings.HasPrefix(dest, "pubsub.googleapis.com/") {
			val := strings.TrimPrefix(dest, "pubsub.googleapis.com/")
			out.Destination.PubSubTopicRef = &pubsubv1beta1.PubSubTopicRef{External: val}
		} else if strings.HasPrefix(dest, "storage.googleapis.com/") {
			val := strings.TrimPrefix(dest, "storage.googleapis.com/")
			out.Destination.StorageBucketRef = &storagev1beta1.StorageBucketRef{External: val}
		}
	}

	if len(in.GetExclusions()) > 0 {
		out.Exclusions = make([]krm.LogSinkExclusions, len(in.GetExclusions()))
		for i, excl := range in.GetExclusions() {
			out.Exclusions[i] = krm.LogSinkExclusions{
				Name:        excl.GetName(),
				Description: direct.LazyPtr(excl.GetDescription()),
				Filter:      excl.GetFilter(),
				Disabled:    direct.LazyPtr(excl.GetDisabled()),
			}
		}
	}

	return out
}

func LoggingLogSinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogSinkSpec) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	out.Description = direct.ValueOf(in.Description)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Filter = direct.ValueOf(in.Filter)
	out.IncludeChildren = direct.ValueOf(in.IncludeChildren)

	if in.BigqueryOptions != nil {
		out.Options = &pb.LogSink_BigqueryOptions{
			BigqueryOptions: &pb.BigQueryOptions{
				UsePartitionedTables: in.BigqueryOptions.UsePartitionedTables,
			},
		}
	}

	if in.Destination.BigQueryDatasetRef != nil && in.Destination.BigQueryDatasetRef.External != "" {
		ext := in.Destination.BigQueryDatasetRef.External
		if !strings.HasPrefix(ext, "bigquery.googleapis.com/") {
			ext = "bigquery.googleapis.com/" + ext
		}
		out.Destination = ext
	} else if in.Destination.LoggingLogBucketRef != nil && in.Destination.LoggingLogBucketRef.External != "" {
		ext := in.Destination.LoggingLogBucketRef.External
		if !strings.HasPrefix(ext, "logging.googleapis.com/") {
			ext = "logging.googleapis.com/" + ext
		}
		out.Destination = ext
	} else if in.Destination.PubSubTopicRef != nil && in.Destination.PubSubTopicRef.External != "" {
		ext := in.Destination.PubSubTopicRef.External
		if !strings.HasPrefix(ext, "pubsub.googleapis.com/") {
			ext = "pubsub.googleapis.com/" + ext
		}
		out.Destination = ext
	} else if in.Destination.StorageBucketRef != nil && in.Destination.StorageBucketRef.External != "" {
		ext := in.Destination.StorageBucketRef.External
		if strings.Contains(ext, "/buckets/") {
			parts := strings.Split(ext, "/buckets/")
			ext = parts[len(parts)-1]
		}
		if !strings.HasPrefix(ext, "storage.googleapis.com/") {
			ext = "storage.googleapis.com/" + ext
		}
		out.Destination = ext
	}

	if len(in.Exclusions) > 0 {
		out.Exclusions = make([]*pb.LogExclusion, len(in.Exclusions))
		for i, excl := range in.Exclusions {
			out.Exclusions[i] = &pb.LogExclusion{
				Name:        excl.Name,
				Description: direct.ValueOf(excl.Description),
				Filter:      excl.Filter,
				Disabled:    direct.ValueOf(excl.Disabled),
			}
		}
	}

	return out
}

func LoggingLogSinkStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LoggingLogSinkStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogSinkStatus{}
	out.WriterIdentity = direct.LazyPtr(in.GetWriterIdentity())
	return out
}

func LoggingLogSinkStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogSinkStatus) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	out.WriterIdentity = direct.ValueOf(in.WriterIdentity)
	return out
}

func LogmetricMetricDescriptor_FromProto(mapCtx *direct.MapContext, in *gmetric.MetricDescriptor) *krm.LogmetricMetricDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricMetricDescriptor{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if in.Labels != nil {
		out.Labels = make([]krm.LogmetricLabels, len(in.Labels))
		for i, lbl := range in.Labels {
			out.Labels[i] = LogmetricLabels_FromProto(mapCtx, lbl)
		}
	}
	if in.LaunchStage != gapi.LaunchStage_LAUNCH_STAGE_UNSPECIFIED {
		out.LaunchStage = direct.LazyPtr(in.LaunchStage.String())
	}
	out.Metadata = LogmetricMetadata_FromProto(mapCtx, in.GetMetadata())
	if in.MetricKind != gmetric.MetricDescriptor_METRIC_KIND_UNSPECIFIED {
		out.MetricKind = direct.LazyPtr(in.MetricKind.String())
	}
	out.Unit = direct.LazyPtr(in.GetUnit())
	if in.ValueType != gmetric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED {
		out.ValueType = direct.LazyPtr(in.ValueType.String())
	}
	return out
}

func LogmetricMetricDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricMetricDescriptor) *gmetric.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &gmetric.MetricDescriptor{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.Labels != nil {
		out.Labels = make([]*glabel.LabelDescriptor, len(in.Labels))
		for i, lbl := range in.Labels {
			out.Labels[i] = LogmetricLabels_ToProto(mapCtx, lbl)
		}
	}
	if in.LaunchStage != nil {
		if val, ok := gapi.LaunchStage_value[*in.LaunchStage]; ok {
			out.LaunchStage = gapi.LaunchStage(val)
		}
	}
	out.Metadata = LogmetricMetadata_ToProto(mapCtx, in.Metadata)
	if in.MetricKind != nil {
		if val, ok := gmetric.MetricDescriptor_MetricKind_value[*in.MetricKind]; ok {
			out.MetricKind = gmetric.MetricDescriptor_MetricKind(val)
		}
	}
	out.Unit = direct.ValueOf(in.Unit)
	if in.ValueType != nil {
		if val, ok := gmetric.MetricDescriptor_ValueType_value[*in.ValueType]; ok {
			out.ValueType = gmetric.MetricDescriptor_ValueType(val)
		}
	}
	return out
}

func LogmetricLabels_FromProto(mapCtx *direct.MapContext, in *glabel.LabelDescriptor) krm.LogmetricLabels {
	if in == nil {
		return krm.LogmetricLabels{}
	}
	out := krm.LogmetricLabels{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Key = direct.LazyPtr(in.GetKey())
	out.ValueType = direct.LazyPtr(in.ValueType.String())
	return out
}

func LogmetricLabels_ToProto(mapCtx *direct.MapContext, in krm.LogmetricLabels) *glabel.LabelDescriptor {
	out := &glabel.LabelDescriptor{}
	out.Description = direct.ValueOf(in.Description)
	out.Key = direct.ValueOf(in.Key)
	if in.ValueType != nil {
		if val, ok := glabel.LabelDescriptor_ValueType_value[*in.ValueType]; ok {
			out.ValueType = glabel.LabelDescriptor_ValueType(val)
		}
	}
	return out
}

func LogmetricMetadata_FromProto(mapCtx *direct.MapContext, in *gmetric.MetricDescriptor_MetricDescriptorMetadata) *krm.LogmetricMetadata {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricMetadata{}
	out.IngestDelay = direct.StringDuration_FromProto(mapCtx, in.GetIngestDelay())
	out.SamplePeriod = direct.StringDuration_FromProto(mapCtx, in.GetSamplePeriod())
	return out
}

func LogmetricMetadata_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricMetadata) *gmetric.MetricDescriptor_MetricDescriptorMetadata {
	if in == nil {
		return nil
	}
	out := &gmetric.MetricDescriptor_MetricDescriptorMetadata{}
	out.IngestDelay = direct.StringDuration_ToProto(mapCtx, in.IngestDelay)
	out.SamplePeriod = direct.StringDuration_ToProto(mapCtx, in.SamplePeriod)
	return out
}

func LogmetricBucketOptions_FromProto(mapCtx *direct.MapContext, in *gdist.Distribution_BucketOptions) *krm.LogmetricBucketOptions {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricBucketOptions{}
	out.ExplicitBuckets = LogmetricExplicitBuckets_FromProto(mapCtx, in.GetExplicitBuckets())
	out.ExponentialBuckets = LogmetricExponentialBuckets_FromProto(mapCtx, in.GetExponentialBuckets())
	out.LinearBuckets = LogmetricLinearBuckets_FromProto(mapCtx, in.GetLinearBuckets())
	return out
}

func LogmetricBucketOptions_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricBucketOptions) *gdist.Distribution_BucketOptions {
	if in == nil {
		return nil
	}
	out := &gdist.Distribution_BucketOptions{}
	if in.ExplicitBuckets != nil {
		out.Options = &gdist.Distribution_BucketOptions_ExplicitBuckets{
			ExplicitBuckets: LogmetricExplicitBuckets_ToProto(mapCtx, in.ExplicitBuckets),
		}
	} else if in.ExponentialBuckets != nil {
		out.Options = &gdist.Distribution_BucketOptions_ExponentialBuckets{
			ExponentialBuckets: LogmetricExponentialBuckets_ToProto(mapCtx, in.ExponentialBuckets),
		}
	} else if in.LinearBuckets != nil {
		out.Options = &gdist.Distribution_BucketOptions_LinearBuckets{
			LinearBuckets: LogmetricLinearBuckets_ToProto(mapCtx, in.LinearBuckets),
		}
	}
	return out
}

func LogmetricExplicitBuckets_FromProto(mapCtx *direct.MapContext, in *gdist.Distribution_BucketOptions_Explicit) *krm.LogmetricExplicitBuckets {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricExplicitBuckets{}
	out.Bounds = in.GetBounds()
	return out
}

func LogmetricExplicitBuckets_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricExplicitBuckets) *gdist.Distribution_BucketOptions_Explicit {
	if in == nil {
		return nil
	}
	out := &gdist.Distribution_BucketOptions_Explicit{}
	out.Bounds = in.Bounds
	return out
}

func LogmetricExponentialBuckets_FromProto(mapCtx *direct.MapContext, in *gdist.Distribution_BucketOptions_Exponential) *krm.LogmetricExponentialBuckets {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricExponentialBuckets{}
	out.GrowthFactor = direct.LazyPtr(in.GetGrowthFactor())
	out.NumFiniteBuckets = direct.LazyPtr(int64(in.GetNumFiniteBuckets()))
	out.Scale = direct.LazyPtr(in.GetScale())
	return out
}

func LogmetricExponentialBuckets_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricExponentialBuckets) *gdist.Distribution_BucketOptions_Exponential {
	if in == nil {
		return nil
	}
	out := &gdist.Distribution_BucketOptions_Exponential{}
	out.GrowthFactor = direct.ValueOf(in.GrowthFactor)
	out.NumFiniteBuckets = int32(direct.ValueOf(in.NumFiniteBuckets))
	out.Scale = direct.ValueOf(in.Scale)
	return out
}

func LogmetricLinearBuckets_FromProto(mapCtx *direct.MapContext, in *gdist.Distribution_BucketOptions_Linear) *krm.LogmetricLinearBuckets {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricLinearBuckets{}
	out.NumFiniteBuckets = direct.LazyPtr(int64(in.GetNumFiniteBuckets()))
	out.Offset = direct.LazyPtr(in.GetOffset())
	out.Width = direct.LazyPtr(in.GetWidth())
	return out
}

func LogmetricLinearBuckets_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricLinearBuckets) *gdist.Distribution_BucketOptions_Linear {
	if in == nil {
		return nil
	}
	out := &gdist.Distribution_BucketOptions_Linear{}
	out.NumFiniteBuckets = int32(direct.ValueOf(in.NumFiniteBuckets))
	out.Offset = direct.ValueOf(in.Offset)
	out.Width = direct.ValueOf(in.Width)
	return out
}

func LoggingLogMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LoggingLogMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = in.GetFilter()
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MetricDescriptor = LogmetricMetricDescriptor_FromProto(mapCtx, in.GetMetricDescriptor())
	out.ValueExtractor = direct.LazyPtr(in.GetValueExtractor())
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = LogmetricBucketOptions_FromProto(mapCtx, in.GetBucketOptions())
	if in.BucketName != "" {
		out.LoggingLogBucketRef = &v1alpha1.ResourceRef{
			External: in.BucketName,
		}
	}
	return out
}

func LoggingLogMetricSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogMetricSpec) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	out.Description = direct.ValueOf(in.Description)
	out.Filter = in.Filter
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MetricDescriptor = LogmetricMetricDescriptor_ToProto(mapCtx, in.MetricDescriptor)
	out.ValueExtractor = direct.ValueOf(in.ValueExtractor)
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = LogmetricBucketOptions_ToProto(mapCtx, in.BucketOptions)
	if in.LoggingLogBucketRef != nil {
		out.BucketName = in.LoggingLogBucketRef.External
	}
	return out
}
