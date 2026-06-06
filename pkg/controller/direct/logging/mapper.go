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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
			out.Destination.BigQueryDatasetRef = &bigqueryv1beta1.DatasetRef{External: dest}
		} else if strings.HasPrefix(dest, "logging.googleapis.com/") {
			out.Destination.LoggingLogBucketRef = &krm.LoggingLogBucketRef{External: dest}
		} else if strings.HasPrefix(dest, "pubsub.googleapis.com/") {
			out.Destination.PubSubTopicRef = &pubsubv1beta1.PubSubTopicRef{External: dest}
		} else if strings.HasPrefix(dest, "storage.googleapis.com/") {
			out.Destination.StorageBucketRef = &storagev1beta1.StorageBucketRef{External: dest}
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
		out.Destination = in.Destination.BigQueryDatasetRef.External
	} else if in.Destination.LoggingLogBucketRef != nil && in.Destination.LoggingLogBucketRef.External != "" {
		out.Destination = in.Destination.LoggingLogBucketRef.External
	} else if in.Destination.PubSubTopicRef != nil && in.Destination.PubSubTopicRef.External != "" {
		out.Destination = in.Destination.PubSubTopicRef.External
	} else if in.Destination.StorageBucketRef != nil && in.Destination.StorageBucketRef.External != "" {
		out.Destination = in.Destination.StorageBucketRef.External
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
