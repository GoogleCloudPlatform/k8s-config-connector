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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DataformWorkflowConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig) *krm.DataformWorkflowConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkflowConfigObservedState{}
	// MISSING: Name
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledExecutionRecords
	return out
}
func DataformWorkflowConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkflowConfigObservedState) *pb.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig{}
	// MISSING: Name
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledExecutionRecords
	return out
}
func DataformWorkflowConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig) *krm.DataformWorkflowConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkflowConfigSpec{}
	// MISSING: Name
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledExecutionRecords
	return out
}
func DataformWorkflowConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkflowConfigSpec) *pb.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig{}
	// MISSING: Name
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	// MISSING: RecentScheduledExecutionRecords
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
func WorkflowConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig) *krm.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowConfig{}
	// MISSING: Name
	out.ReleaseConfig = direct.LazyPtr(in.GetReleaseConfig())
	out.InvocationConfig = InvocationConfig_FromProto(mapCtx, in.GetInvocationConfig())
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	// MISSING: RecentScheduledExecutionRecords
	return out
}
func WorkflowConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowConfig) *pb.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig{}
	// MISSING: Name
	out.ReleaseConfig = direct.ValueOf(in.ReleaseConfig)
	out.InvocationConfig = InvocationConfig_ToProto(mapCtx, in.InvocationConfig)
	out.CronSchedule = direct.ValueOf(in.CronSchedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	// MISSING: RecentScheduledExecutionRecords
	return out
}
func WorkflowConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig) *krm.WorkflowConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	out.RecentScheduledExecutionRecords = direct.Slice_FromProto(mapCtx, in.RecentScheduledExecutionRecords, WorkflowConfig_ScheduledExecutionRecord_FromProto)
	return out
}
func WorkflowConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowConfigObservedState) *pb.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ReleaseConfig
	// MISSING: InvocationConfig
	// MISSING: CronSchedule
	// MISSING: TimeZone
	out.RecentScheduledExecutionRecords = direct.Slice_ToProto(mapCtx, in.RecentScheduledExecutionRecords, WorkflowConfig_ScheduledExecutionRecord_ToProto)
	return out
}
func WorkflowConfig_ScheduledExecutionRecord_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig_ScheduledExecutionRecord) *krm.WorkflowConfig_ScheduledExecutionRecord {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowConfig_ScheduledExecutionRecord{}
	out.ExecutionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExecutionTime())
	out.WorkflowInvocation = direct.LazyPtr(in.GetWorkflowInvocation())
	out.ErrorStatus = Status_FromProto(mapCtx, in.GetErrorStatus())
	return out
}
func WorkflowConfig_ScheduledExecutionRecord_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowConfig_ScheduledExecutionRecord) *pb.WorkflowConfig_ScheduledExecutionRecord {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig_ScheduledExecutionRecord{}
	out.ExecutionTime = direct.StringTimestamp_ToProto(mapCtx, in.ExecutionTime)
	if oneof := WorkflowConfig_ScheduledExecutionRecord_WorkflowInvocation_ToProto(mapCtx, in.WorkflowInvocation); oneof != nil {
		out.Result = oneof
	}
	if oneof := Status_ToProto(mapCtx, in.ErrorStatus); oneof != nil {
		out.Result = &pb.WorkflowConfig_ScheduledExecutionRecord_ErrorStatus{ErrorStatus: oneof}
	}
	return out
}
