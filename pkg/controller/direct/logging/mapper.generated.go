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

func LogEntryOperation_FromProto(mapCtx *direct.MapContext, in *pb.LogEntryOperation) *krm.LogEntryOperation {
	if in == nil {
		return nil
	}
	out := &krm.LogEntryOperation{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Producer = direct.LazyPtr(in.GetProducer())
	out.First = direct.LazyPtr(in.GetFirst())
	out.Last = direct.LazyPtr(in.GetLast())
	return out
}
func LogEntryOperation_ToProto(mapCtx *direct.MapContext, in *krm.LogEntryOperation) *pb.LogEntryOperation {
	if in == nil {
		return nil
	}
	out := &pb.LogEntryOperation{}
	out.Id = direct.ValueOf(in.ID)
	out.Producer = direct.ValueOf(in.Producer)
	out.First = direct.ValueOf(in.First)
	out.Last = direct.ValueOf(in.Last)
	return out
}
func LogEntrySourceLocation_FromProto(mapCtx *direct.MapContext, in *pb.LogEntrySourceLocation) *krm.LogEntrySourceLocation {
	if in == nil {
		return nil
	}
	out := &krm.LogEntrySourceLocation{}
	out.File = direct.LazyPtr(in.GetFile())
	out.Line = direct.LazyPtr(in.GetLine())
	out.Function = direct.LazyPtr(in.GetFunction())
	return out
}
func LogEntrySourceLocation_ToProto(mapCtx *direct.MapContext, in *krm.LogEntrySourceLocation) *pb.LogEntrySourceLocation {
	if in == nil {
		return nil
	}
	out := &pb.LogEntrySourceLocation{}
	out.File = direct.ValueOf(in.File)
	out.Line = direct.ValueOf(in.Line)
	out.Function = direct.ValueOf(in.Function)
	return out
}
func LogSplit_FromProto(mapCtx *direct.MapContext, in *pb.LogSplit) *krm.LogSplit {
	if in == nil {
		return nil
	}
	out := &krm.LogSplit{}
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Index = direct.LazyPtr(in.GetIndex())
	out.TotalSplits = direct.LazyPtr(in.GetTotalSplits())
	return out
}
func LogSplit_ToProto(mapCtx *direct.MapContext, in *krm.LogSplit) *pb.LogSplit {
	if in == nil {
		return nil
	}
	out := &pb.LogSplit{}
	out.Uid = direct.ValueOf(in.Uid)
	out.Index = direct.ValueOf(in.Index)
	out.TotalSplits = direct.ValueOf(in.TotalSplits)
	return out
}
func LoggingLogEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogEntry) *krm.LoggingLogEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogEntryObservedState{}
	// MISSING: LogName
	// MISSING: Resource
	// MISSING: ProtoPayload
	// MISSING: TextPayload
	// MISSING: JsonPayload
	// MISSING: Timestamp
	// MISSING: ReceiveTimestamp
	// MISSING: Severity
	// MISSING: InsertID
	// MISSING: HTTPRequest
	// MISSING: Labels
	// MISSING: Operation
	// MISSING: Trace
	// MISSING: SpanID
	// MISSING: TraceSampled
	// MISSING: SourceLocation
	// MISSING: Split
	return out
}
func LoggingLogEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogEntryObservedState) *pb.LogEntry {
	if in == nil {
		return nil
	}
	out := &pb.LogEntry{}
	// MISSING: LogName
	// MISSING: Resource
	// MISSING: ProtoPayload
	// MISSING: TextPayload
	// MISSING: JsonPayload
	// MISSING: Timestamp
	// MISSING: ReceiveTimestamp
	// MISSING: Severity
	// MISSING: InsertID
	// MISSING: HTTPRequest
	// MISSING: Labels
	// MISSING: Operation
	// MISSING: Trace
	// MISSING: SpanID
	// MISSING: TraceSampled
	// MISSING: SourceLocation
	// MISSING: Split
	return out
}
func LoggingLogEntrySpec_FromProto(mapCtx *direct.MapContext, in *pb.LogEntry) *krm.LoggingLogEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogEntrySpec{}
	// MISSING: LogName
	// MISSING: Resource
	// MISSING: ProtoPayload
	// MISSING: TextPayload
	// MISSING: JsonPayload
	// MISSING: Timestamp
	// MISSING: ReceiveTimestamp
	// MISSING: Severity
	// MISSING: InsertID
	// MISSING: HTTPRequest
	// MISSING: Labels
	// MISSING: Operation
	// MISSING: Trace
	// MISSING: SpanID
	// MISSING: TraceSampled
	// MISSING: SourceLocation
	// MISSING: Split
	return out
}
func LoggingLogEntrySpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogEntrySpec) *pb.LogEntry {
	if in == nil {
		return nil
	}
	out := &pb.LogEntry{}
	// MISSING: LogName
	// MISSING: Resource
	// MISSING: ProtoPayload
	// MISSING: TextPayload
	// MISSING: JsonPayload
	// MISSING: Timestamp
	// MISSING: ReceiveTimestamp
	// MISSING: Severity
	// MISSING: InsertID
	// MISSING: HTTPRequest
	// MISSING: Labels
	// MISSING: Operation
	// MISSING: Trace
	// MISSING: SpanID
	// MISSING: TraceSampled
	// MISSING: SourceLocation
	// MISSING: Split
	return out
}
