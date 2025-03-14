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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DriverSchedulingConfig_FromProto(mapCtx *direct.MapContext, in *pb.DriverSchedulingConfig) *krm.DriverSchedulingConfig {
	if in == nil {
		return nil
	}
	out := &krm.DriverSchedulingConfig{}
	out.MemoryMb = direct.LazyPtr(in.GetMemoryMb())
	out.Vcores = direct.LazyPtr(in.GetVcores())
	return out
}
func DriverSchedulingConfig_ToProto(mapCtx *direct.MapContext, in *krm.DriverSchedulingConfig) *pb.DriverSchedulingConfig {
	if in == nil {
		return nil
	}
	out := &pb.DriverSchedulingConfig{}
	out.MemoryMb = direct.ValueOf(in.MemoryMb)
	out.Vcores = direct.ValueOf(in.Vcores)
	return out
}
func JobPlacement_FromProto(mapCtx *direct.MapContext, in *pb.JobPlacement) *krm.JobPlacement {
	if in == nil {
		return nil
	}
	out := &krm.JobPlacement{}
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	// MISSING: ClusterUuid
	out.ClusterLabels = in.ClusterLabels
	return out
}
func JobPlacement_ToProto(mapCtx *direct.MapContext, in *krm.JobPlacement) *pb.JobPlacement {
	if in == nil {
		return nil
	}
	out := &pb.JobPlacement{}
	out.ClusterName = direct.ValueOf(in.ClusterName)
	// MISSING: ClusterUuid
	out.ClusterLabels = in.ClusterLabels
	return out
}
func JobPlacementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobPlacement) *krm.JobPlacementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobPlacementObservedState{}
	// MISSING: ClusterName
	out.ClusterUuid = direct.LazyPtr(in.GetClusterUuid())
	// MISSING: ClusterLabels
	return out
}
func JobPlacementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobPlacementObservedState) *pb.JobPlacement {
	if in == nil {
		return nil
	}
	out := &pb.JobPlacement{}
	// MISSING: ClusterName
	out.ClusterUuid = direct.ValueOf(in.ClusterUuid)
	// MISSING: ClusterLabels
	return out
}
func JobReference_FromProto(mapCtx *direct.MapContext, in *pb.JobReference) *krm.JobReference {
	if in == nil {
		return nil
	}
	out := &krm.JobReference{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.JobID = direct.LazyPtr(in.GetJobId())
	return out
}
func JobReference_ToProto(mapCtx *direct.MapContext, in *krm.JobReference) *pb.JobReference {
	if in == nil {
		return nil
	}
	out := &pb.JobReference{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.JobId = direct.ValueOf(in.JobID)
	return out
}
func JobScheduling_FromProto(mapCtx *direct.MapContext, in *pb.JobScheduling) *krm.JobScheduling {
	if in == nil {
		return nil
	}
	out := &krm.JobScheduling{}
	out.MaxFailuresPerHour = direct.LazyPtr(in.GetMaxFailuresPerHour())
	out.MaxFailuresTotal = direct.LazyPtr(in.GetMaxFailuresTotal())
	return out
}
func JobScheduling_ToProto(mapCtx *direct.MapContext, in *krm.JobScheduling) *pb.JobScheduling {
	if in == nil {
		return nil
	}
	out := &pb.JobScheduling{}
	out.MaxFailuresPerHour = direct.ValueOf(in.MaxFailuresPerHour)
	out.MaxFailuresTotal = direct.ValueOf(in.MaxFailuresTotal)
	return out
}
func JobStatus_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatus {
	if in == nil {
		return nil
	}
	out := &krm.JobStatus{}
	// MISSING: State
	// MISSING: Details
	// MISSING: StateStartTime
	// MISSING: Substate
	return out
}
func JobStatus_ToProto(mapCtx *direct.MapContext, in *krm.JobStatus) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	// MISSING: State
	// MISSING: Details
	// MISSING: StateStartTime
	// MISSING: Substate
	return out
}
func JobStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.JobStatus) *krm.JobStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.JobStatusObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Details = direct.LazyPtr(in.GetDetails())
	out.StateStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateStartTime())
	out.Substate = direct.Enum_FromProto(mapCtx, in.GetSubstate())
	return out
}
func JobStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.JobStatusObservedState) *pb.JobStatus {
	if in == nil {
		return nil
	}
	out := &pb.JobStatus{}
	out.State = direct.Enum_ToProto[pb.JobStatus_State](mapCtx, in.State)
	out.Details = direct.ValueOf(in.Details)
	out.StateStartTime = direct.StringTimestamp_ToProto(mapCtx, in.StateStartTime)
	out.Substate = direct.Enum_ToProto[pb.JobStatus_Substate](mapCtx, in.Substate)
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	// MISSING: DriverLogLevels
	return out
}
func PrestoJob_FromProto(mapCtx *direct.MapContext, in *pb.PrestoJob) *krm.PrestoJob {
	if in == nil {
		return nil
	}
	out := &krm.PrestoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func PrestoJob_ToProto(mapCtx *direct.MapContext, in *krm.PrestoJob) *pb.PrestoJob {
	if in == nil {
		return nil
	}
	out := &pb.PrestoJob{}
	if oneof := PrestoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.PrestoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func QueryList_FromProto(mapCtx *direct.MapContext, in *pb.QueryList) *krm.QueryList {
	if in == nil {
		return nil
	}
	out := &krm.QueryList{}
	out.Queries = in.Queries
	return out
}
func QueryList_ToProto(mapCtx *direct.MapContext, in *krm.QueryList) *pb.QueryList {
	if in == nil {
		return nil
	}
	out := &pb.QueryList{}
	out.Queries = in.Queries
	return out
}
func TrinoJob_FromProto(mapCtx *direct.MapContext, in *pb.TrinoJob) *krm.TrinoJob {
	if in == nil {
		return nil
	}
	out := &krm.TrinoJob{}
	out.QueryFileURI = direct.LazyPtr(in.GetQueryFileUri())
	out.QueryList = QueryList_FromProto(mapCtx, in.GetQueryList())
	out.ContinueOnFailure = direct.LazyPtr(in.GetContinueOnFailure())
	out.OutputFormat = direct.LazyPtr(in.GetOutputFormat())
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func TrinoJob_ToProto(mapCtx *direct.MapContext, in *krm.TrinoJob) *pb.TrinoJob {
	if in == nil {
		return nil
	}
	out := &pb.TrinoJob{}
	if oneof := TrinoJob_QueryFileUri_ToProto(mapCtx, in.QueryFileURI); oneof != nil {
		out.Queries = oneof
	}
	if oneof := QueryList_ToProto(mapCtx, in.QueryList); oneof != nil {
		out.Queries = &pb.TrinoJob_QueryList{QueryList: oneof}
	}
	out.ContinueOnFailure = direct.ValueOf(in.ContinueOnFailure)
	out.OutputFormat = direct.ValueOf(in.OutputFormat)
	out.ClientTags = in.ClientTags
	out.Properties = in.Properties
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func YarnApplication_FromProto(mapCtx *direct.MapContext, in *pb.YarnApplication) *krm.YarnApplication {
	if in == nil {
		return nil
	}
	out := &krm.YarnApplication{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Progress = direct.LazyPtr(in.GetProgress())
	out.TrackingURL = direct.LazyPtr(in.GetTrackingUrl())
	return out
}
func YarnApplication_ToProto(mapCtx *direct.MapContext, in *krm.YarnApplication) *pb.YarnApplication {
	if in == nil {
		return nil
	}
	out := &pb.YarnApplication{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.YarnApplication_State](mapCtx, in.State)
	out.Progress = direct.ValueOf(in.Progress)
	out.TrackingUrl = direct.ValueOf(in.TrackingURL)
	return out
}
