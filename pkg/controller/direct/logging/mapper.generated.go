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

func LogEntry_FromProto(mapCtx *direct.MapContext, in *pb.LogEntry) *krm.LogEntry {
	if in == nil {
		return nil
	}
	out := &krm.LogEntry{}
	out.LogName = direct.LazyPtr(in.GetLogName())
	out.Resource = MonitoredResource_FromProto(mapCtx, in.GetResource())
	out.ProtoPayload = Any_FromProto(mapCtx, in.GetProtoPayload())
	out.TextPayload = direct.LazyPtr(in.GetTextPayload())
	out.JsonPayload = JsonPayload_FromProto(mapCtx, in.GetJsonPayload())
	out.Timestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetTimestamp())
	// MISSING: ReceiveTimestamp
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.InsertID = direct.LazyPtr(in.GetInsertId())
	out.HTTPRequest = HTTPRequest_FromProto(mapCtx, in.GetHttpRequest())
	out.Labels = in.Labels
	out.Operation = LogEntryOperation_FromProto(mapCtx, in.GetOperation())
	out.Trace = direct.LazyPtr(in.GetTrace())
	out.SpanID = direct.LazyPtr(in.GetSpanId())
	out.TraceSampled = direct.LazyPtr(in.GetTraceSampled())
	out.SourceLocation = LogEntrySourceLocation_FromProto(mapCtx, in.GetSourceLocation())
	out.Split = LogSplit_FromProto(mapCtx, in.GetSplit())
	return out
}
func LogEntry_ToProto(mapCtx *direct.MapContext, in *krm.LogEntry) *pb.LogEntry {
	if in == nil {
		return nil
	}
	out := &pb.LogEntry{}
	out.LogName = direct.ValueOf(in.LogName)
	out.Resource = MonitoredResource_ToProto(mapCtx, in.Resource)
	if oneof := Any_ToProto(mapCtx, in.ProtoPayload); oneof != nil {
		out.Payload = &pb.LogEntry_ProtoPayload{ProtoPayload: oneof}
	}
	if oneof := LogEntry_TextPayload_ToProto(mapCtx, in.TextPayload); oneof != nil {
		out.Payload = oneof
	}
	if oneof := JsonPayload_ToProto(mapCtx, in.JsonPayload); oneof != nil {
		out.Payload = &pb.LogEntry_JsonPayload{JsonPayload: oneof}
	}
	out.Timestamp = direct.StringTimestamp_ToProto(mapCtx, in.Timestamp)
	// MISSING: ReceiveTimestamp
	out.Severity = direct.Enum_ToProto[pb.LogSeverity](mapCtx, in.Severity)
	out.InsertId = direct.ValueOf(in.InsertID)
	out.HttpRequest = HTTPRequest_ToProto(mapCtx, in.HTTPRequest)
	out.Labels = in.Labels
	out.Operation = LogEntryOperation_ToProto(mapCtx, in.Operation)
	out.Trace = direct.ValueOf(in.Trace)
	out.SpanId = direct.ValueOf(in.SpanID)
	out.TraceSampled = direct.ValueOf(in.TraceSampled)
	out.SourceLocation = LogEntrySourceLocation_ToProto(mapCtx, in.SourceLocation)
	out.Split = LogSplit_ToProto(mapCtx, in.Split)
	return out
}
func LogEntryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogEntry) *krm.LogEntryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogEntryObservedState{}
	// MISSING: LogName
	// MISSING: Resource
	// MISSING: ProtoPayload
	// MISSING: TextPayload
	// MISSING: JsonPayload
	// MISSING: Timestamp
	out.ReceiveTimestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetReceiveTimestamp())
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
func LogEntryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogEntryObservedState) *pb.LogEntry {
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
	out.ReceiveTimestamp = direct.StringTimestamp_ToProto(mapCtx, in.ReceiveTimestamp)
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
