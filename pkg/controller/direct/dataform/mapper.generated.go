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
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func CompilationResult_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult) *krm.CompilationResult {
	if in == nil {
		return nil
	}
	out := &krm.CompilationResult{}
	// MISSING: Name
	out.GitCommitish = direct.LazyPtr(in.GetGitCommitish())
	out.Workspace = direct.LazyPtr(in.GetWorkspace())
	out.ReleaseConfig = direct.LazyPtr(in.GetReleaseConfig())
	out.CodeCompilationConfig = CodeCompilationConfig_FromProto(mapCtx, in.GetCodeCompilationConfig())
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
	return out
}
func CompilationResult_ToProto(mapCtx *direct.MapContext, in *krm.CompilationResult) *pb.CompilationResult {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult{}
	// MISSING: Name
	if oneof := CompilationResult_GitCommitish_ToProto(mapCtx, in.GitCommitish); oneof != nil {
		out.Source = oneof
	}
	if oneof := CompilationResult_Workspace_ToProto(mapCtx, in.Workspace); oneof != nil {
		out.Source = oneof
	}
	if oneof := CompilationResult_ReleaseConfig_ToProto(mapCtx, in.ReleaseConfig); oneof != nil {
		out.Source = oneof
	}
	out.CodeCompilationConfig = CodeCompilationConfig_ToProto(mapCtx, in.CodeCompilationConfig)
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
	return out
}
func CompilationResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult) *krm.CompilationResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompilationResultObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	out.ResolvedGitCommitSha = direct.LazyPtr(in.GetResolvedGitCommitSha())
	out.DataformCoreVersion = direct.LazyPtr(in.GetDataformCoreVersion())
	out.CompilationErrors = direct.Slice_FromProto(mapCtx, in.CompilationErrors, CompilationResult_CompilationError_FromProto)
	return out
}
func CompilationResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompilationResultObservedState) *pb.CompilationResult {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	out.ResolvedGitCommitSha = direct.ValueOf(in.ResolvedGitCommitSha)
	out.DataformCoreVersion = direct.ValueOf(in.DataformCoreVersion)
	out.CompilationErrors = direct.Slice_ToProto(mapCtx, in.CompilationErrors, CompilationResult_CompilationError_ToProto)
	return out
}
func CompilationResult_CompilationError_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult_CompilationError) *krm.CompilationResult_CompilationError {
	if in == nil {
		return nil
	}
	out := &krm.CompilationResult_CompilationError{}
	// MISSING: Message
	// MISSING: Stack
	// MISSING: Path
	// MISSING: ActionTarget
	return out
}
func CompilationResult_CompilationError_ToProto(mapCtx *direct.MapContext, in *krm.CompilationResult_CompilationError) *pb.CompilationResult_CompilationError {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult_CompilationError{}
	// MISSING: Message
	// MISSING: Stack
	// MISSING: Path
	// MISSING: ActionTarget
	return out
}
func CompilationResult_CompilationErrorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult_CompilationError) *krm.CompilationResult_CompilationErrorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompilationResult_CompilationErrorObservedState{}
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Stack = direct.LazyPtr(in.GetStack())
	out.Path = direct.LazyPtr(in.GetPath())
	out.ActionTarget = Target_FromProto(mapCtx, in.GetActionTarget())
	return out
}
func CompilationResult_CompilationErrorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompilationResult_CompilationErrorObservedState) *pb.CompilationResult_CompilationError {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult_CompilationError{}
	out.Message = direct.ValueOf(in.Message)
	out.Stack = direct.ValueOf(in.Stack)
	out.Path = direct.ValueOf(in.Path)
	out.ActionTarget = Target_ToProto(mapCtx, in.ActionTarget)
	return out
}
func DataformCompilationResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult) *krm.DataformCompilationResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformCompilationResultObservedState{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
	return out
}
func DataformCompilationResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformCompilationResultObservedState) *pb.CompilationResult {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
	return out
}
func DataformCompilationResultSpec_FromProto(mapCtx *direct.MapContext, in *pb.CompilationResult) *krm.DataformCompilationResultSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformCompilationResultSpec{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
	return out
}
func DataformCompilationResultSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformCompilationResultSpec) *pb.CompilationResult {
	if in == nil {
		return nil
	}
	out := &pb.CompilationResult{}
	// MISSING: Name
	// MISSING: GitCommitish
	// MISSING: Workspace
	// MISSING: ReleaseConfig
	// MISSING: CodeCompilationConfig
	// MISSING: ResolvedGitCommitSha
	// MISSING: DataformCoreVersion
	// MISSING: CompilationErrors
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
func Target_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.Target {
	if in == nil {
		return nil
	}
	out := &krm.Target{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Target_ToProto(mapCtx *direct.MapContext, in *krm.Target) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	out.Database = direct.ValueOf(in.Database)
	out.Schema = direct.ValueOf(in.Schema)
	out.Name = direct.ValueOf(in.Name)
	return out
}
