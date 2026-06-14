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
	"time"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	apipb "google.golang.org/genproto/googleapis/api"
	distributionpb "google.golang.org/genproto/googleapis/api/distribution"
	labelpb "google.golang.org/genproto/googleapis/api/label"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	"google.golang.org/protobuf/types/known/durationpb"
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

func LoggingLogMetricSpec_Filter_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func LogmetricMetricDescriptor_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.LogmetricMetricDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricMetricDescriptor{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if in.GetLaunchStage() != apipb.LaunchStage_LAUNCH_STAGE_UNSPECIFIED {
		out.LaunchStage = direct.LazyPtr(in.GetLaunchStage().String())
	}
	if in.GetMetricKind() != metricpb.MetricDescriptor_METRIC_KIND_UNSPECIFIED {
		out.MetricKind = direct.LazyPtr(in.GetMetricKind().String())
	}
	if in.GetValueType() != metricpb.MetricDescriptor_VALUE_TYPE_UNSPECIFIED {
		out.ValueType = direct.LazyPtr(in.GetValueType().String())
	}
	out.Unit = direct.LazyPtr(in.GetUnit())
	if in.GetMetadata() != nil {
		metadata := &krm.LogmetricMetadata{}
		hasMetadata := false
		if in.GetMetadata().GetIngestDelay() != nil {
			metadata.IngestDelay = direct.LazyPtr(in.GetMetadata().GetIngestDelay().AsDuration().String())
			hasMetadata = true
		}
		if in.GetMetadata().GetSamplePeriod() != nil {
			metadata.SamplePeriod = direct.LazyPtr(in.GetMetadata().GetSamplePeriod().AsDuration().String())
			hasMetadata = true
		}
		if hasMetadata {
			out.Metadata = metadata
		}
	}
	if len(in.GetLabels()) > 0 {
		out.Labels = make([]krm.LogmetricLabels, len(in.GetLabels()))
		for i, label := range in.GetLabels() {
			out.Labels[i] = krm.LogmetricLabels{
				Key:         direct.LazyPtr(label.GetKey()),
				Description: direct.LazyPtr(label.GetDescription()),
			}
			out.Labels[i].ValueType = direct.LazyPtr(label.GetValueType().String())
		}
	}
	return out
}

func LogmetricMetricDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricMetricDescriptor) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.LaunchStage != nil {
		out.LaunchStage = apipb.LaunchStage(apipb.LaunchStage_value[*in.LaunchStage])
	}
	if in.MetricKind != nil {
		out.MetricKind = metricpb.MetricDescriptor_MetricKind(metricpb.MetricDescriptor_MetricKind_value[*in.MetricKind])
	}
	if in.ValueType != nil {
		out.ValueType = metricpb.MetricDescriptor_ValueType(metricpb.MetricDescriptor_ValueType_value[*in.ValueType])
	}
	out.Unit = direct.ValueOf(in.Unit)
	if in.Metadata != nil {
		out.Metadata = &metricpb.MetricDescriptor_MetricDescriptorMetadata{}
		if in.Metadata.IngestDelay != nil {
			d, err := time.ParseDuration(*in.Metadata.IngestDelay)
			if err == nil {
				out.Metadata.IngestDelay = durationpb.New(d)
			}
		}
		if in.Metadata.SamplePeriod != nil {
			d, err := time.ParseDuration(*in.Metadata.SamplePeriod)
			if err == nil {
				out.Metadata.SamplePeriod = durationpb.New(d)
			}
		}
	}
	if len(in.Labels) > 0 {
		out.Labels = make([]*labelpb.LabelDescriptor, len(in.Labels))
		for i, label := range in.Labels {
			out.Labels[i] = &labelpb.LabelDescriptor{}
			out.Labels[i].Key = direct.ValueOf(label.Key)
			out.Labels[i].Description = direct.ValueOf(label.Description)
			if label.ValueType != nil {
				out.Labels[i].ValueType = labelpb.LabelDescriptor_ValueType(labelpb.LabelDescriptor_ValueType_value[*label.ValueType])
			}
		}
	}
	return out
}

func LogmetricBucketOptions_FromProto(mapCtx *direct.MapContext, in *distributionpb.Distribution_BucketOptions) *krm.LogmetricBucketOptions {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricBucketOptions{}
	if in.GetExplicitBuckets() != nil {
		out.ExplicitBuckets = &krm.LogmetricExplicitBuckets{
			Bounds: in.GetExplicitBuckets().GetBounds(),
		}
	}
	if in.GetExponentialBuckets() != nil {
		out.ExponentialBuckets = &krm.LogmetricExponentialBuckets{
			NumFiniteBuckets: direct.LazyPtr(int64(in.GetExponentialBuckets().GetNumFiniteBuckets())),
			GrowthFactor:     direct.LazyPtr(in.GetExponentialBuckets().GetGrowthFactor()),
			Scale:            direct.LazyPtr(in.GetExponentialBuckets().GetScale()),
		}
	}
	if in.GetLinearBuckets() != nil {
		out.LinearBuckets = &krm.LogmetricLinearBuckets{
			NumFiniteBuckets: direct.LazyPtr(int64(in.GetLinearBuckets().GetNumFiniteBuckets())),
			Offset:           direct.LazyPtr(in.GetLinearBuckets().GetOffset()),
			Width:            direct.LazyPtr(in.GetLinearBuckets().GetWidth()),
		}
	}
	return out
}

func LogmetricBucketOptions_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricBucketOptions) *distributionpb.Distribution_BucketOptions {
	if in == nil {
		return nil
	}
	out := &distributionpb.Distribution_BucketOptions{}
	if in.ExplicitBuckets != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_ExplicitBuckets{
			ExplicitBuckets: &distributionpb.Distribution_BucketOptions_Explicit{
				Bounds: in.ExplicitBuckets.Bounds,
			},
		}
	}
	if in.ExponentialBuckets != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_ExponentialBuckets{
			ExponentialBuckets: &distributionpb.Distribution_BucketOptions_Exponential{
				NumFiniteBuckets: int32(direct.ValueOf(in.ExponentialBuckets.NumFiniteBuckets)),
				GrowthFactor:     direct.ValueOf(in.ExponentialBuckets.GrowthFactor),
				Scale:            direct.ValueOf(in.ExponentialBuckets.Scale),
			},
		}
	}
	if in.LinearBuckets != nil {
		out.Options = &distributionpb.Distribution_BucketOptions_LinearBuckets{
			LinearBuckets: &distributionpb.Distribution_BucketOptions_Linear{
				NumFiniteBuckets: int32(direct.ValueOf(in.LinearBuckets.NumFiniteBuckets)),
				Offset:           direct.ValueOf(in.LinearBuckets.Offset),
				Width:            direct.ValueOf(in.LinearBuckets.Width),
			},
		}
	}
	return out
}

func LogmetricMetricDescriptorStatus_FromProto(mapCtx *direct.MapContext, in *metricpb.MetricDescriptor) *krm.LogmetricMetricDescriptorStatus {
	if in == nil {
		return nil
	}
	out := &krm.LogmetricMetricDescriptorStatus{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MonitoredResourceTypes = in.GetMonitoredResourceTypes()
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}

func LogmetricMetricDescriptorStatus_ToProto(mapCtx *direct.MapContext, in *krm.LogmetricMetricDescriptorStatus) *metricpb.MetricDescriptor {
	if in == nil {
		return nil
	}
	out := &metricpb.MetricDescriptor{}
	out.Description = direct.ValueOf(in.Description)
	out.MonitoredResourceTypes = in.MonitoredResourceTypes
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	return out
}
