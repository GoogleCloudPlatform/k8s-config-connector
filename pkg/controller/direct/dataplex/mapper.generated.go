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

package dataplex

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)
func DataplexMetadataJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob) *krm.DataplexMetadataJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexMetadataJobObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func DataplexMetadataJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexMetadataJobObservedState) *pb.MetadataJob {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func DataplexMetadataJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob) *krm.DataplexMetadataJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexMetadataJobSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func DataplexMetadataJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexMetadataJobSpec) *pb.MetadataJob {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func MetadataJob_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob) *krm.MetadataJob {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ImportSpec = MetadataJob_ImportJobSpec_FromProto(mapCtx, in.GetImportSpec())
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func MetadataJob_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob) *pb.MetadataJob {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.MetadataJob_Type](mapCtx, in.Type)
	if oneof := MetadataJob_ImportJobSpec_ToProto(mapCtx, in.ImportSpec); oneof != nil {
		out.Spec = &pb.MetadataJob_ImportSpec{ImportSpec: oneof}
	}
	// MISSING: ImportResult
	// MISSING: Status
	return out
}
func MetadataJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob) *krm.MetadataJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	out.ImportResult = MetadataJob_ImportJobResult_FromProto(mapCtx, in.GetImportResult())
	out.Status = MetadataJob_Status_FromProto(mapCtx, in.GetStatus())
	return out
}
func MetadataJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJobObservedState) *pb.MetadataJob {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Type
	// MISSING: ImportSpec
	if oneof := MetadataJob_ImportJobResult_ToProto(mapCtx, in.ImportResult); oneof != nil {
		out.Result = &pb.MetadataJob_ImportResult{ImportResult: oneof}
	}
	out.Status = MetadataJob_Status_ToProto(mapCtx, in.Status)
	return out
}
func MetadataJob_ImportJobResult_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_ImportJobResult) *krm.MetadataJob_ImportJobResult {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_ImportJobResult{}
	// MISSING: DeletedEntries
	// MISSING: UpdatedEntries
	// MISSING: CreatedEntries
	// MISSING: UnchangedEntries
	// MISSING: RecreatedEntries
	// MISSING: UpdateTime
	return out
}
func MetadataJob_ImportJobResult_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_ImportJobResult) *pb.MetadataJob_ImportJobResult {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_ImportJobResult{}
	// MISSING: DeletedEntries
	// MISSING: UpdatedEntries
	// MISSING: CreatedEntries
	// MISSING: UnchangedEntries
	// MISSING: RecreatedEntries
	// MISSING: UpdateTime
	return out
}
func MetadataJob_ImportJobResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_ImportJobResult) *krm.MetadataJob_ImportJobResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_ImportJobResultObservedState{}
	out.DeletedEntries = direct.LazyPtr(in.GetDeletedEntries())
	out.UpdatedEntries = direct.LazyPtr(in.GetUpdatedEntries())
	out.CreatedEntries = direct.LazyPtr(in.GetCreatedEntries())
	out.UnchangedEntries = direct.LazyPtr(in.GetUnchangedEntries())
	out.RecreatedEntries = direct.LazyPtr(in.GetRecreatedEntries())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func MetadataJob_ImportJobResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_ImportJobResultObservedState) *pb.MetadataJob_ImportJobResult {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_ImportJobResult{}
	out.DeletedEntries = direct.ValueOf(in.DeletedEntries)
	out.UpdatedEntries = direct.ValueOf(in.UpdatedEntries)
	out.CreatedEntries = direct.ValueOf(in.CreatedEntries)
	out.UnchangedEntries = direct.ValueOf(in.UnchangedEntries)
	out.RecreatedEntries = direct.ValueOf(in.RecreatedEntries)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func MetadataJob_ImportJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_ImportJobSpec) *krm.MetadataJob_ImportJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_ImportJobSpec{}
	out.SourceStorageURI = direct.LazyPtr(in.GetSourceStorageUri())
	out.SourceCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSourceCreateTime())
	out.Scope = MetadataJob_ImportJobSpec_ImportJobScope_FromProto(mapCtx, in.GetScope())
	out.EntrySyncMode = direct.Enum_FromProto(mapCtx, in.GetEntrySyncMode())
	out.AspectSyncMode = direct.Enum_FromProto(mapCtx, in.GetAspectSyncMode())
	out.LogLevel = direct.Enum_FromProto(mapCtx, in.GetLogLevel())
	return out
}
func MetadataJob_ImportJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_ImportJobSpec) *pb.MetadataJob_ImportJobSpec {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_ImportJobSpec{}
	out.SourceStorageUri = direct.ValueOf(in.SourceStorageURI)
	out.SourceCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.SourceCreateTime)
	out.Scope = MetadataJob_ImportJobSpec_ImportJobScope_ToProto(mapCtx, in.Scope)
	out.EntrySyncMode = direct.Enum_ToProto[pb.MetadataJob_ImportJobSpec_SyncMode](mapCtx, in.EntrySyncMode)
	out.AspectSyncMode = direct.Enum_ToProto[pb.MetadataJob_ImportJobSpec_SyncMode](mapCtx, in.AspectSyncMode)
	out.LogLevel = direct.Enum_ToProto[pb.MetadataJob_ImportJobSpec_LogLevel](mapCtx, in.LogLevel)
	return out
}
func MetadataJob_ImportJobSpec_ImportJobScope_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_ImportJobSpec_ImportJobScope) *krm.MetadataJob_ImportJobSpec_ImportJobScope {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_ImportJobSpec_ImportJobScope{}
	out.EntryGroups = in.EntryGroups
	out.EntryTypes = in.EntryTypes
	out.AspectTypes = in.AspectTypes
	return out
}
func MetadataJob_ImportJobSpec_ImportJobScope_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_ImportJobSpec_ImportJobScope) *pb.MetadataJob_ImportJobSpec_ImportJobScope {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_ImportJobSpec_ImportJobScope{}
	out.EntryGroups = in.EntryGroups
	out.EntryTypes = in.EntryTypes
	out.AspectTypes = in.AspectTypes
	return out
}
func MetadataJob_Status_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_Status) *krm.MetadataJob_Status {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_Status{}
	// MISSING: State
	// MISSING: Message
	// MISSING: CompletionPercent
	// MISSING: UpdateTime
	return out
}
func MetadataJob_Status_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_Status) *pb.MetadataJob_Status {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_Status{}
	// MISSING: State
	// MISSING: Message
	// MISSING: CompletionPercent
	// MISSING: UpdateTime
	return out
}
func MetadataJob_StatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataJob_Status) *krm.MetadataJob_StatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataJob_StatusObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.CompletionPercent = direct.LazyPtr(in.GetCompletionPercent())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func MetadataJob_StatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataJob_StatusObservedState) *pb.MetadataJob_Status {
	if in == nil {
		return nil
	}
	out := &pb.MetadataJob_Status{}
	out.State = direct.Enum_ToProto[pb.MetadataJob_Status_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.CompletionPercent = direct.ValueOf(in.CompletionPercent)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
