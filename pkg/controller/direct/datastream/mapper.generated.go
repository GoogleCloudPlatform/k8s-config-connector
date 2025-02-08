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

package datastream

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
)
func BackfillJob_FromProto(mapCtx *direct.MapContext, in *pb.BackfillJob) *krm.BackfillJob {
	if in == nil {
		return nil
	}
	out := &krm.BackfillJob{}
	// MISSING: State
	out.Trigger = direct.Enum_FromProto(mapCtx, in.GetTrigger())
	// MISSING: LastStartTime
	// MISSING: LastEndTime
	// MISSING: Errors
	return out
}
func BackfillJob_ToProto(mapCtx *direct.MapContext, in *krm.BackfillJob) *pb.BackfillJob {
	if in == nil {
		return nil
	}
	out := &pb.BackfillJob{}
	// MISSING: State
	out.Trigger = direct.Enum_ToProto[pb.BackfillJob_Trigger](mapCtx, in.Trigger)
	// MISSING: LastStartTime
	// MISSING: LastEndTime
	// MISSING: Errors
	return out
}
func BackfillJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackfillJob) *krm.BackfillJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackfillJobObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Trigger
	out.LastStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastStartTime())
	out.LastEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastEndTime())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Error_FromProto)
	return out
}
func BackfillJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackfillJobObservedState) *pb.BackfillJob {
	if in == nil {
		return nil
	}
	out := &pb.BackfillJob{}
	out.State = direct.Enum_ToProto[pb.BackfillJob_State](mapCtx, in.State)
	// MISSING: Trigger
	out.LastStartTime = direct.StringTimestamp_ToProto(mapCtx, in.LastStartTime)
	out.LastEndTime = direct.StringTimestamp_ToProto(mapCtx, in.LastEndTime)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Error_ToProto)
	return out
}
func DatastreamStreamObjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StreamObject) *krm.DatastreamStreamObjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamObjectObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Errors
	// MISSING: BackfillJob
	// MISSING: SourceObject
	return out
}
func DatastreamStreamObjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamObjectObservedState) *pb.StreamObject {
	if in == nil {
		return nil
	}
	out := &pb.StreamObject{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Errors
	// MISSING: BackfillJob
	// MISSING: SourceObject
	return out
}
func DatastreamStreamObjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.StreamObject) *krm.DatastreamStreamObjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamObjectSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Errors
	// MISSING: BackfillJob
	// MISSING: SourceObject
	return out
}
func DatastreamStreamObjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamObjectSpec) *pb.StreamObject {
	if in == nil {
		return nil
	}
	out := &pb.StreamObject{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Errors
	// MISSING: BackfillJob
	// MISSING: SourceObject
	return out
}
func Error_FromProto(mapCtx *direct.MapContext, in *pb.Error) *krm.Error {
	if in == nil {
		return nil
	}
	out := &krm.Error{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.ErrorUuid = direct.LazyPtr(in.GetErrorUuid())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetErrorTime())
	out.Details = in.Details
	return out
}
func Error_ToProto(mapCtx *direct.MapContext, in *krm.Error) *pb.Error {
	if in == nil {
		return nil
	}
	out := &pb.Error{}
	out.Reason = direct.ValueOf(in.Reason)
	out.ErrorUuid = direct.ValueOf(in.ErrorUuid)
	out.Message = direct.ValueOf(in.Message)
	out.ErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.ErrorTime)
	out.Details = in.Details
	return out
}
func SourceObjectIdentifier_FromProto(mapCtx *direct.MapContext, in *pb.SourceObjectIdentifier) *krm.SourceObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &krm.SourceObjectIdentifier{}
	out.OracleIdentifier = SourceObjectIdentifier_OracleObjectIdentifier_FromProto(mapCtx, in.GetOracleIdentifier())
	out.MysqlIdentifier = SourceObjectIdentifier_MysqlObjectIdentifier_FromProto(mapCtx, in.GetMysqlIdentifier())
	out.PostgresqlIdentifier = SourceObjectIdentifier_PostgresqlObjectIdentifier_FromProto(mapCtx, in.GetPostgresqlIdentifier())
	out.SqlServerIdentifier = SourceObjectIdentifier_SqlServerObjectIdentifier_FromProto(mapCtx, in.GetSqlServerIdentifier())
	return out
}
func SourceObjectIdentifier_ToProto(mapCtx *direct.MapContext, in *krm.SourceObjectIdentifier) *pb.SourceObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &pb.SourceObjectIdentifier{}
	if oneof := SourceObjectIdentifier_OracleObjectIdentifier_ToProto(mapCtx, in.OracleIdentifier); oneof != nil {
		out.SourceIdentifier = &pb.SourceObjectIdentifier_OracleIdentifier{OracleIdentifier: oneof}
	}
	if oneof := SourceObjectIdentifier_MysqlObjectIdentifier_ToProto(mapCtx, in.MysqlIdentifier); oneof != nil {
		out.SourceIdentifier = &pb.SourceObjectIdentifier_MysqlIdentifier{MysqlIdentifier: oneof}
	}
	if oneof := SourceObjectIdentifier_PostgresqlObjectIdentifier_ToProto(mapCtx, in.PostgresqlIdentifier); oneof != nil {
		out.SourceIdentifier = &pb.SourceObjectIdentifier_PostgresqlIdentifier{PostgresqlIdentifier: oneof}
	}
	if oneof := SourceObjectIdentifier_SqlServerObjectIdentifier_ToProto(mapCtx, in.SqlServerIdentifier); oneof != nil {
		out.SourceIdentifier = &pb.SourceObjectIdentifier_SqlServerIdentifier{SqlServerIdentifier: oneof}
	}
	return out
}
func SourceObjectIdentifier_MysqlObjectIdentifier_FromProto(mapCtx *direct.MapContext, in *pb.SourceObjectIdentifier_MysqlObjectIdentifier) *krm.SourceObjectIdentifier_MysqlObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &krm.SourceObjectIdentifier_MysqlObjectIdentifier{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Table = direct.LazyPtr(in.GetTable())
	return out
}
func SourceObjectIdentifier_MysqlObjectIdentifier_ToProto(mapCtx *direct.MapContext, in *krm.SourceObjectIdentifier_MysqlObjectIdentifier) *pb.SourceObjectIdentifier_MysqlObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &pb.SourceObjectIdentifier_MysqlObjectIdentifier{}
	out.Database = direct.ValueOf(in.Database)
	out.Table = direct.ValueOf(in.Table)
	return out
}
func SourceObjectIdentifier_OracleObjectIdentifier_FromProto(mapCtx *direct.MapContext, in *pb.SourceObjectIdentifier_OracleObjectIdentifier) *krm.SourceObjectIdentifier_OracleObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &krm.SourceObjectIdentifier_OracleObjectIdentifier{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Table = direct.LazyPtr(in.GetTable())
	return out
}
func SourceObjectIdentifier_OracleObjectIdentifier_ToProto(mapCtx *direct.MapContext, in *krm.SourceObjectIdentifier_OracleObjectIdentifier) *pb.SourceObjectIdentifier_OracleObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &pb.SourceObjectIdentifier_OracleObjectIdentifier{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Table = direct.ValueOf(in.Table)
	return out
}
func SourceObjectIdentifier_PostgresqlObjectIdentifier_FromProto(mapCtx *direct.MapContext, in *pb.SourceObjectIdentifier_PostgresqlObjectIdentifier) *krm.SourceObjectIdentifier_PostgresqlObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &krm.SourceObjectIdentifier_PostgresqlObjectIdentifier{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Table = direct.LazyPtr(in.GetTable())
	return out
}
func SourceObjectIdentifier_PostgresqlObjectIdentifier_ToProto(mapCtx *direct.MapContext, in *krm.SourceObjectIdentifier_PostgresqlObjectIdentifier) *pb.SourceObjectIdentifier_PostgresqlObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &pb.SourceObjectIdentifier_PostgresqlObjectIdentifier{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Table = direct.ValueOf(in.Table)
	return out
}
func SourceObjectIdentifier_SqlServerObjectIdentifier_FromProto(mapCtx *direct.MapContext, in *pb.SourceObjectIdentifier_SqlServerObjectIdentifier) *krm.SourceObjectIdentifier_SqlServerObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &krm.SourceObjectIdentifier_SqlServerObjectIdentifier{}
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Table = direct.LazyPtr(in.GetTable())
	return out
}
func SourceObjectIdentifier_SqlServerObjectIdentifier_ToProto(mapCtx *direct.MapContext, in *krm.SourceObjectIdentifier_SqlServerObjectIdentifier) *pb.SourceObjectIdentifier_SqlServerObjectIdentifier {
	if in == nil {
		return nil
	}
	out := &pb.SourceObjectIdentifier_SqlServerObjectIdentifier{}
	out.Schema = direct.ValueOf(in.Schema)
	out.Table = direct.ValueOf(in.Table)
	return out
}
func StreamObject_FromProto(mapCtx *direct.MapContext, in *pb.StreamObject) *krm.StreamObject {
	if in == nil {
		return nil
	}
	out := &krm.StreamObject{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Errors
	out.BackfillJob = BackfillJob_FromProto(mapCtx, in.GetBackfillJob())
	out.SourceObject = SourceObjectIdentifier_FromProto(mapCtx, in.GetSourceObject())
	return out
}
func StreamObject_ToProto(mapCtx *direct.MapContext, in *krm.StreamObject) *pb.StreamObject {
	if in == nil {
		return nil
	}
	out := &pb.StreamObject{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Errors
	out.BackfillJob = BackfillJob_ToProto(mapCtx, in.BackfillJob)
	out.SourceObject = SourceObjectIdentifier_ToProto(mapCtx, in.SourceObject)
	return out
}
func StreamObjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StreamObject) *krm.StreamObjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StreamObjectObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, Error_FromProto)
	out.BackfillJob = BackfillJobObservedState_FromProto(mapCtx, in.GetBackfillJob())
	// MISSING: SourceObject
	return out
}
func StreamObjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StreamObjectObservedState) *pb.StreamObject {
	if in == nil {
		return nil
	}
	out := &pb.StreamObject{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, Error_ToProto)
	out.BackfillJob = BackfillJobObservedState_ToProto(mapCtx, in.BackfillJob)
	// MISSING: SourceObject
	return out
}
