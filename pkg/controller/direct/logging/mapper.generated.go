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

func BigQueryOptions_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryOptions) *krm.BigQueryOptions {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryOptions{}
	out.UsePartitionedTables = direct.LazyPtr(in.GetUsePartitionedTables())
	// MISSING: UsesTimestampColumnPartitioning
	return out
}
func BigQueryOptions_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryOptions) *pb.BigQueryOptions {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryOptions{}
	out.UsePartitionedTables = direct.ValueOf(in.UsePartitionedTables)
	// MISSING: UsesTimestampColumnPartitioning
	return out
}
func BigQueryOptionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryOptions) *krm.BigQueryOptionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryOptionsObservedState{}
	// MISSING: UsePartitionedTables
	out.UsesTimestampColumnPartitioning = direct.LazyPtr(in.GetUsesTimestampColumnPartitioning())
	return out
}
func BigQueryOptionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryOptionsObservedState) *pb.BigQueryOptions {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryOptions{}
	// MISSING: UsePartitionedTables
	out.UsesTimestampColumnPartitioning = direct.ValueOf(in.UsesTimestampColumnPartitioning)
	return out
}
func LogExclusion_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LogExclusion {
	if in == nil {
		return nil
	}
	out := &krm.LogExclusion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogExclusion_ToProto(mapCtx *direct.MapContext, in *krm.LogExclusion) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	out.Disabled = direct.ValueOf(in.Disabled)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogExclusionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LogExclusionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogExclusionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func LogExclusionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogExclusionObservedState) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func LogSink_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LogSink {
	if in == nil {
		return nil
	}
	out := &krm.LogSink{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Destination = direct.LazyPtr(in.GetDestination())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Exclusions = direct.Slice_FromProto(mapCtx, in.Exclusions, LogExclusion_FromProto)
	out.OutputVersionFormat = direct.Enum_FromProto(mapCtx, in.GetOutputVersionFormat())
	// MISSING: WriterIdentity
	out.IncludeChildren = direct.LazyPtr(in.GetIncludeChildren())
	out.BigqueryOptions = BigQueryOptions_FromProto(mapCtx, in.GetBigqueryOptions())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogSink_ToProto(mapCtx *direct.MapContext, in *krm.LogSink) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	out.Name = direct.ValueOf(in.Name)
	out.Destination = direct.ValueOf(in.Destination)
	out.Filter = direct.ValueOf(in.Filter)
	out.Description = direct.ValueOf(in.Description)
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Exclusions = direct.Slice_ToProto(mapCtx, in.Exclusions, LogExclusion_ToProto)
	out.OutputVersionFormat = direct.Enum_ToProto[pb.LogSink_VersionFormat](mapCtx, in.OutputVersionFormat)
	// MISSING: WriterIdentity
	out.IncludeChildren = direct.ValueOf(in.IncludeChildren)
	if oneof := BigQueryOptions_ToProto(mapCtx, in.BigqueryOptions); oneof != nil {
		out.Options = &pb.LogSink_BigqueryOptions{BigqueryOptions: oneof}
	}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogSinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LogSinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogSinkObservedState{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	out.Exclusions = direct.Slice_FromProto(mapCtx, in.Exclusions, LogExclusionObservedState_FromProto)
	// MISSING: OutputVersionFormat
	out.WriterIdentity = direct.LazyPtr(in.GetWriterIdentity())
	// MISSING: IncludeChildren
	out.BigqueryOptions = BigQueryOptionsObservedState_FromProto(mapCtx, in.GetBigqueryOptions())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func LogSinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogSinkObservedState) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	out.Exclusions = direct.Slice_ToProto(mapCtx, in.Exclusions, LogExclusionObservedState_ToProto)
	// MISSING: OutputVersionFormat
	out.WriterIdentity = direct.ValueOf(in.WriterIdentity)
	// MISSING: IncludeChildren
	if oneof := BigQueryOptionsObservedState_ToProto(mapCtx, in.BigqueryOptions); oneof != nil {
		out.Options = &pb.LogSink_BigqueryOptions{BigqueryOptions: oneof}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func LoggingLogSinkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LoggingLogSinkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogSinkObservedState{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	// MISSING: Exclusions
	// MISSING: OutputVersionFormat
	// MISSING: WriterIdentity
	// MISSING: IncludeChildren
	// MISSING: BigqueryOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogSinkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogSinkObservedState) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	// MISSING: Exclusions
	// MISSING: OutputVersionFormat
	// MISSING: WriterIdentity
	// MISSING: IncludeChildren
	// MISSING: BigqueryOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogSinkSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogSink) *krm.LoggingLogSinkSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogSinkSpec{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	// MISSING: Exclusions
	// MISSING: OutputVersionFormat
	// MISSING: WriterIdentity
	// MISSING: IncludeChildren
	// MISSING: BigqueryOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogSinkSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogSinkSpec) *pb.LogSink {
	if in == nil {
		return nil
	}
	out := &pb.LogSink{}
	// MISSING: Name
	// MISSING: Destination
	// MISSING: Filter
	// MISSING: Description
	// MISSING: Disabled
	// MISSING: Exclusions
	// MISSING: OutputVersionFormat
	// MISSING: WriterIdentity
	// MISSING: IncludeChildren
	// MISSING: BigqueryOptions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
