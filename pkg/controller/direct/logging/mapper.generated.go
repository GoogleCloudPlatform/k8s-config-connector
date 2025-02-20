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

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LogMetric_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LogMetric {
	if in == nil {
		return nil
	}
	out := &krm.LogMetric{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.BucketName = direct.LazyPtr(in.GetBucketName())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.MetricDescriptor = MetricDescriptor_FromProto(mapCtx, in.GetMetricDescriptor())
	out.ValueExtractor = direct.LazyPtr(in.GetValueExtractor())
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = Distribution_BucketOptions_FromProto(mapCtx, in.GetBucketOptions())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Version = direct.Enum_FromProto(mapCtx, in.GetVersion())
	return out
}
func LogMetric_ToProto(mapCtx *direct.MapContext, in *krm.LogMetric) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	out.BucketName = direct.ValueOf(in.BucketName)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.MetricDescriptor = MetricDescriptor_ToProto(mapCtx, in.MetricDescriptor)
	out.ValueExtractor = direct.ValueOf(in.ValueExtractor)
	out.LabelExtractors = in.LabelExtractors
	out.BucketOptions = Distribution_BucketOptions_ToProto(mapCtx, in.BucketOptions)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Version = direct.Enum_ToProto[pb.LogMetric_ApiVersion](mapCtx, in.Version)
	return out
}
func LogMetricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LogMetricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogMetricObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Version
	return out
}
func LogMetricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogMetricObservedState) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Version
	return out
}
func LoggingLogMetricObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LoggingLogMetricObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Version
	return out
}
func LoggingLogMetricObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogMetricObservedState) *pb.LogMetric {
	if in == nil {
		return nil
	}
	out := &pb.LogMetric{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Version
	return out
}
func LoggingLogMetricSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogMetric) *krm.LoggingLogMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
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
	// MISSING: Description
	// MISSING: Filter
	// MISSING: BucketName
	// MISSING: Disabled
	// MISSING: MetricDescriptor
	// MISSING: ValueExtractor
	// MISSING: LabelExtractors
	// MISSING: BucketOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Version
	return out
}
