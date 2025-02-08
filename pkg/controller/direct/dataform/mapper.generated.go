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

package dataform

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
)
func CodeCompilationConfig_FromProto(mapCtx *direct.MapContext, in *pb.CodeCompilationConfig) *krm.CodeCompilationConfig {
	if in == nil {
		return nil
	}
	out := &krm.CodeCompilationConfig{}
	out.DefaultDatabase = direct.LazyPtr(in.GetDefaultDatabase())
	out.DefaultSchema = direct.LazyPtr(in.GetDefaultSchema())
	out.DefaultLocation = direct.LazyPtr(in.GetDefaultLocation())
	out.AssertionSchema = direct.LazyPtr(in.GetAssertionSchema())
	out.Vars = in.Vars
	out.DatabaseSuffix = direct.LazyPtr(in.GetDatabaseSuffix())
	out.SchemaSuffix = direct.LazyPtr(in.GetSchemaSuffix())
	out.TablePrefix = direct.LazyPtr(in.GetTablePrefix())
	return out
}
func CodeCompilationConfig_ToProto(mapCtx *direct.MapContext, in *krm.CodeCompilationConfig) *pb.CodeCompilationConfig {
	if in == nil {
		return nil
	}
	out := &pb.CodeCompilationConfig{}
	out.DefaultDatabase = direct.ValueOf(in.DefaultDatabase)
	out.DefaultSchema = direct.ValueOf(in.DefaultSchema)
	out.DefaultLocation = direct.ValueOf(in.DefaultLocation)
	out.AssertionSchema = direct.ValueOf(in.AssertionSchema)
	out.Vars = in.Vars
	out.DatabaseSuffix = direct.ValueOf(in.DatabaseSuffix)
	out.SchemaSuffix = direct.ValueOf(in.SchemaSuffix)
	out.TablePrefix = direct.ValueOf(in.TablePrefix)
	return out
}
func DataformReleaseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig) *krm.DataformReleaseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformReleaseConfigObservedState{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledReleaseRecords
	// MISSING: ReleaseCompilationResult
	return out
}
func DataformReleaseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformReleaseConfigObservedState) *pb.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledReleaseRecords
	// MISSING: ReleaseCompilationResult
	return out
}
func DataformReleaseConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig) *krm.DataformReleaseConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformReleaseConfigSpec{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledReleaseRecords
	// MISSING: ReleaseCompilationResult
	return out
}
func DataformReleaseConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformReleaseConfigSpec) *pb.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledReleaseRecords
	// MISSING: ReleaseCompilationResult
	return out
}
func DataformRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.DataformRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositoryObservedState{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func DataformRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func ReleaseConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig) *krm.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReleaseConfig{}
	// MISSING: Name
	out.GitCommitish = direct.LazyPtr(in.GetGitCommitish())
	out.CodeCompilationConfig = CodeCompilationConfig_FromProto(mapCtx, in.GetCodeCompilationConfig())
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	// MISSING: RecentScheduledReleaseRecords
	out.ReleaseCompilationResult = direct.LazyPtr(in.GetReleaseCompilationResult())
	return out
}
func ReleaseConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReleaseConfig) *pb.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig{}
	// MISSING: Name
	out.GitCommitish = direct.ValueOf(in.GitCommitish)
	out.CodeCompilationConfig = CodeCompilationConfig_ToProto(mapCtx, in.CodeCompilationConfig)
	out.CronSchedule = direct.ValueOf(in.CronSchedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	// MISSING: RecentScheduledReleaseRecords
	out.ReleaseCompilationResult = direct.ValueOf(in.ReleaseCompilationResult)
	return out
}
func ReleaseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig) *krm.ReleaseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReleaseConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	out.RecentScheduledReleaseRecords = direct.Slice_FromProto(mapCtx, in.RecentScheduledReleaseRecords, ReleaseConfig_ScheduledReleaseRecord_FromProto)
	// MISSING: ReleaseCompilationResult
	return out
}
func ReleaseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReleaseConfigObservedState) *pb.ReleaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: GitCommitish
	// MISSING: CodeCompilationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	out.RecentScheduledReleaseRecords = direct.Slice_ToProto(mapCtx, in.RecentScheduledReleaseRecords, ReleaseConfig_ScheduledReleaseRecord_ToProto)
	// MISSING: ReleaseCompilationResult
	return out
}
func ReleaseConfig_ScheduledReleaseRecord_FromProto(mapCtx *direct.MapContext, in *pb.ReleaseConfig_ScheduledReleaseRecord) *krm.ReleaseConfig_ScheduledReleaseRecord {
	if in == nil {
		return nil
	}
	out := &krm.ReleaseConfig_ScheduledReleaseRecord{}
	out.ReleaseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetReleaseTime())
	out.CompilationResult = direct.LazyPtr(in.GetCompilationResult())
	out.ErrorStatus = Status_FromProto(mapCtx, in.GetErrorStatus())
	return out
}
func ReleaseConfig_ScheduledReleaseRecord_ToProto(mapCtx *direct.MapContext, in *krm.ReleaseConfig_ScheduledReleaseRecord) *pb.ReleaseConfig_ScheduledReleaseRecord {
	if in == nil {
		return nil
	}
	out := &pb.ReleaseConfig_ScheduledReleaseRecord{}
	out.ReleaseTime = direct.StringTimestamp_ToProto(mapCtx, in.ReleaseTime)
	if oneof := ReleaseConfig_ScheduledReleaseRecord_CompilationResult_ToProto(mapCtx, in.CompilationResult); oneof != nil {
		out.Result = oneof
	}
	if oneof := Status_ToProto(mapCtx, in.ErrorStatus); oneof != nil {
		out.Result = &pb.ReleaseConfig_ScheduledReleaseRecord_ErrorStatus{ErrorStatus: oneof}
	}
	return out
}
