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
