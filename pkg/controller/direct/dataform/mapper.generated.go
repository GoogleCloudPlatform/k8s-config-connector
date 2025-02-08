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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
)
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
func DataformWorkflowInvocationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowInvocation) *krm.DataformWorkflowInvocationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkflowInvocationObservedState{}
	// MISSING: Name
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func DataformWorkflowInvocationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkflowInvocationObservedState) *pb.WorkflowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowInvocation{}
	// MISSING: Name
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func DataformWorkflowInvocationSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowInvocation) *krm.DataformWorkflowInvocationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkflowInvocationSpec{}
	// MISSING: Name
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func DataformWorkflowInvocationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkflowInvocationSpec) *pb.WorkflowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowInvocation{}
	// MISSING: Name
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func InvocationConfig_FromProto(mapCtx *direct.MapContext, in *pb.InvocationConfig) *krm.InvocationConfig {
	if in == nil {
		return nil
	}
	out := &krm.InvocationConfig{}
	out.IncludedTargets = direct.Slice_FromProto(mapCtx, in.IncludedTargets, Target_FromProto)
	out.IncludedTags = in.IncludedTags
	out.TransitiveDependenciesIncluded = direct.LazyPtr(in.GetTransitiveDependenciesIncluded())
	out.TransitiveDependentsIncluded = direct.LazyPtr(in.GetTransitiveDependentsIncluded())
	out.FullyRefreshIncrementalTablesEnabled = direct.LazyPtr(in.GetFullyRefreshIncrementalTablesEnabled())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func InvocationConfig_ToProto(mapCtx *direct.MapContext, in *krm.InvocationConfig) *pb.InvocationConfig {
	if in == nil {
		return nil
	}
	out := &pb.InvocationConfig{}
	out.IncludedTargets = direct.Slice_ToProto(mapCtx, in.IncludedTargets, Target_ToProto)
	out.IncludedTags = in.IncludedTags
	out.TransitiveDependenciesIncluded = direct.ValueOf(in.TransitiveDependenciesIncluded)
	out.TransitiveDependentsIncluded = direct.ValueOf(in.TransitiveDependentsIncluded)
	out.FullyRefreshIncrementalTablesEnabled = direct.ValueOf(in.FullyRefreshIncrementalTablesEnabled)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
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
func WorkflowInvocation_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowInvocation) *krm.WorkflowInvocation {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowInvocation{}
	// MISSING: Name
	out.CompilationResult = direct.LazyPtr(in.GetCompilationResult())
	out.WorkflowConfig = direct.LazyPtr(in.GetWorkflowConfig())
	out.InvocationConfig = InvocationConfig_FromProto(mapCtx, in.GetInvocationConfig())
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func WorkflowInvocation_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowInvocation) *pb.WorkflowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowInvocation{}
	// MISSING: Name
	if oneof := WorkflowInvocation_CompilationResult_ToProto(mapCtx, in.CompilationResult); oneof != nil {
		out.CompilationSource = oneof
	}
	if oneof := WorkflowInvocation_WorkflowConfig_ToProto(mapCtx, in.WorkflowConfig); oneof != nil {
		out.CompilationSource = oneof
	}
	out.InvocationConfig = InvocationConfig_ToProto(mapCtx, in.InvocationConfig)
	// MISSING: State
	// MISSING: InvocationTiming
	return out
}
func WorkflowInvocationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowInvocation) *krm.WorkflowInvocationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowInvocationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.InvocationTiming = Interval_FromProto(mapCtx, in.GetInvocationTiming())
	return out
}
func WorkflowInvocationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowInvocationObservedState) *pb.WorkflowInvocation {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowInvocation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CompilationResult
	// MISSING: WorkflowConfig
	// MISSING: InvocationConfig
	out.State = direct.Enum_ToProto[pb.WorkflowInvocation_State](mapCtx, in.State)
	out.InvocationTiming = Interval_ToProto(mapCtx, in.InvocationTiming)
	return out
}
