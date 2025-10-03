// Copyright 2024 Google LLC
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

package dataflow

import (
	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoscalingEvent_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingEvent) *krm.AutoscalingEvent {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingEvent{}
	out.CurrentNumWorkers = direct.LazyPtr(in.GetCurrentNumWorkers())
	out.TargetNumWorkers = direct.LazyPtr(in.GetTargetNumWorkers())
	out.EventType = direct.Enum_FromProto(mapCtx, in.GetEventType())
	out.Description = StructuredMessage_FromProto(mapCtx, in.GetDescription())
	out.Time = AutoscalingEvent_Time_FromProto(mapCtx, in.GetTime())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	return out
}
func AutoscalingEvent_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingEvent) *pb.AutoscalingEvent {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingEvent{}
	out.CurrentNumWorkers = direct.ValueOf(in.CurrentNumWorkers)
	out.TargetNumWorkers = direct.ValueOf(in.TargetNumWorkers)
	out.EventType = direct.Enum_ToProto[pb.AutoscalingEvent_AutoscalingEventType](mapCtx, in.EventType)
	out.Description = StructuredMessage_ToProto(mapCtx, in.Description)
	out.Time = AutoscalingEvent_Time_ToProto(mapCtx, in.Time)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	return out
}
func AutoscalingSettings_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingSettings) *krm.AutoscalingSettings {
	if in == nil {
		return nil
	}
	out := &krm.AutoscalingSettings{}
	out.Algorithm = direct.Enum_FromProto(mapCtx, in.GetAlgorithm())
	out.MaxNumWorkers = direct.LazyPtr(in.GetMaxNumWorkers())
	return out
}
func AutoscalingSettings_ToProto(mapCtx *direct.MapContext, in *krm.AutoscalingSettings) *pb.AutoscalingSettings {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingSettings{}
	out.Algorithm = direct.Enum_ToProto[pb.AutoscalingAlgorithm](mapCtx, in.Algorithm)
	out.MaxNumWorkers = direct.ValueOf(in.MaxNumWorkers)
	return out
}
func BigQueryIODetails_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryIODetails) *krm.BigQueryIODetails {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryIODetails{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Query = direct.LazyPtr(in.GetQuery())
	return out
}
func BigQueryIODetails_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryIODetails) *pb.BigQueryIODetails {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryIODetails{}
	out.Table = direct.ValueOf(in.Table)
	out.Dataset = direct.ValueOf(in.Dataset)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Query = direct.ValueOf(in.Query)
	return out
}
func BigTableIODetails_FromProto(mapCtx *direct.MapContext, in *pb.BigTableIODetails) *krm.BigTableIODetails {
	if in == nil {
		return nil
	}
	out := &krm.BigTableIODetails{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.TableID = direct.LazyPtr(in.GetTableId())
	return out
}
func BigTableIODetails_ToProto(mapCtx *direct.MapContext, in *krm.BigTableIODetails) *pb.BigTableIODetails {
	if in == nil {
		return nil
	}
	out := &pb.BigTableIODetails{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.TableId = direct.ValueOf(in.TableID)
	return out
}
func ComputationTopology_FromProto(mapCtx *direct.MapContext, in *pb.ComputationTopology) *krm.ComputationTopology {
	if in == nil {
		return nil
	}
	out := &krm.ComputationTopology{}
	out.SystemStageName = direct.LazyPtr(in.GetSystemStageName())
	out.ComputationID = direct.LazyPtr(in.GetComputationId())
	out.KeyRanges = direct.Slice_FromProto(mapCtx, in.KeyRanges, KeyRangeLocation_FromProto)
	out.Inputs = direct.Slice_FromProto(mapCtx, in.Inputs, StreamLocation_FromProto)
	out.Outputs = direct.Slice_FromProto(mapCtx, in.Outputs, StreamLocation_FromProto)
	out.StateFamilies = direct.Slice_FromProto(mapCtx, in.StateFamilies, StateFamilyConfig_FromProto)
	return out
}
func ComputationTopology_ToProto(mapCtx *direct.MapContext, in *krm.ComputationTopology) *pb.ComputationTopology {
	if in == nil {
		return nil
	}
	out := &pb.ComputationTopology{}
	out.SystemStageName = direct.ValueOf(in.SystemStageName)
	out.ComputationId = direct.ValueOf(in.ComputationID)
	out.KeyRanges = direct.Slice_ToProto(mapCtx, in.KeyRanges, KeyRangeLocation_ToProto)
	out.Inputs = direct.Slice_ToProto(mapCtx, in.Inputs, StreamLocation_ToProto)
	out.Outputs = direct.Slice_ToProto(mapCtx, in.Outputs, StreamLocation_ToProto)
	out.StateFamilies = direct.Slice_ToProto(mapCtx, in.StateFamilies, StateFamilyConfig_ToProto)
	return out
}
func CustomSourceLocation_FromProto(mapCtx *direct.MapContext, in *pb.CustomSourceLocation) *krm.CustomSourceLocation {
	if in == nil {
		return nil
	}
	out := &krm.CustomSourceLocation{}
	out.Stateful = direct.LazyPtr(in.GetStateful())
	return out
}
func CustomSourceLocation_ToProto(mapCtx *direct.MapContext, in *krm.CustomSourceLocation) *pb.CustomSourceLocation {
	if in == nil {
		return nil
	}
	out := &pb.CustomSourceLocation{}
	out.Stateful = direct.ValueOf(in.Stateful)
	return out
}
func DataDiskAssignment_FromProto(mapCtx *direct.MapContext, in *pb.DataDiskAssignment) *krm.DataDiskAssignment {
	if in == nil {
		return nil
	}
	out := &krm.DataDiskAssignment{}
	out.VmInstance = direct.LazyPtr(in.GetVmInstance())
	out.DataDisks = in.DataDisks
	return out
}
func DataDiskAssignment_ToProto(mapCtx *direct.MapContext, in *krm.DataDiskAssignment) *pb.DataDiskAssignment {
	if in == nil {
		return nil
	}
	out := &pb.DataDiskAssignment{}
	out.VmInstance = direct.ValueOf(in.VmInstance)
	out.DataDisks = in.DataDisks
	return out
}
func DatastoreIODetails_FromProto(mapCtx *direct.MapContext, in *pb.DatastoreIODetails) *krm.DatastoreIODetails {
	if in == nil {
		return nil
	}
	out := &krm.DatastoreIODetails{}
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func DatastoreIODetails_ToProto(mapCtx *direct.MapContext, in *krm.DatastoreIODetails) *pb.DatastoreIODetails {
	if in == nil {
		return nil
	}
	out := &pb.DatastoreIODetails{}
	out.Namespace = direct.ValueOf(in.Namespace)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	return out
}
func DebugOptions_FromProto(mapCtx *direct.MapContext, in *pb.DebugOptions) *krm.DebugOptions {
	if in == nil {
		return nil
	}
	out := &krm.DebugOptions{}
	out.EnableHotKeyLogging = direct.LazyPtr(in.GetEnableHotKeyLogging())
	return out
}
func DebugOptions_ToProto(mapCtx *direct.MapContext, in *krm.DebugOptions) *pb.DebugOptions {
	if in == nil {
		return nil
	}
	out := &pb.DebugOptions{}
	out.EnableHotKeyLogging = direct.ValueOf(in.EnableHotKeyLogging)
	return out
}
func Disk_FromProto(mapCtx *direct.MapContext, in *pb.Disk) *krm.Disk {
	if in == nil {
		return nil
	}
	out := &krm.Disk{}
	out.SizeGb = direct.LazyPtr(in.GetSizeGb())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.MountPoint = direct.LazyPtr(in.GetMountPoint())
	return out
}
func Disk_ToProto(mapCtx *direct.MapContext, in *krm.Disk) *pb.Disk {
	if in == nil {
		return nil
	}
	out := &pb.Disk{}
	out.SizeGb = direct.ValueOf(in.SizeGb)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.MountPoint = direct.ValueOf(in.MountPoint)
	return out
}
func DynamicTemplateLaunchParams_FromProto(mapCtx *direct.MapContext, in *pb.DynamicTemplateLaunchParams) *krm.DynamicTemplateLaunchParams {
	if in == nil {
		return nil
	}
	out := &krm.DynamicTemplateLaunchParams{}
	out.GcsPath = direct.LazyPtr(in.GetGcsPath())
	out.StagingLocation = direct.LazyPtr(in.GetStagingLocation())
	return out
}
func DynamicTemplateLaunchParams_ToProto(mapCtx *direct.MapContext, in *krm.DynamicTemplateLaunchParams) *pb.DynamicTemplateLaunchParams {
	if in == nil {
		return nil
	}
	out := &pb.DynamicTemplateLaunchParams{}
	out.GcsPath = direct.ValueOf(in.GcsPath)
	out.StagingLocation = direct.ValueOf(in.StagingLocation)
	return out
}
func ExecutionStageState_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionStageState) *krm.ExecutionStageState {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionStageState{}
	out.ExecutionStageName = direct.LazyPtr(in.GetExecutionStageName())
	out.ExecutionStageState = direct.Enum_FromProto(mapCtx, in.GetExecutionStageState())
	out.CurrentStateTime = ExecutionStageState_CurrentStateTime_FromProto(mapCtx, in.GetCurrentStateTime())
	return out
}
func ExecutionStageState_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionStageState) *pb.ExecutionStageState {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionStageState{}
	out.ExecutionStageName = direct.ValueOf(in.ExecutionStageName)
	out.ExecutionStageState = direct.Enum_ToProto[pb.JobState](mapCtx, in.ExecutionStageState)
	out.CurrentStateTime = ExecutionStageState_CurrentStateTime_ToProto(mapCtx, in.CurrentStateTime)
	return out
}
func ExecutionStageSummary_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionStageSummary) *krm.ExecutionStageSummary {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionStageSummary{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.InputSource = direct.Slice_FromProto(mapCtx, in.InputSource, ExecutionStageSummary_StageSource_FromProto)
	out.OutputSource = direct.Slice_FromProto(mapCtx, in.OutputSource, ExecutionStageSummary_StageSource_FromProto)
	out.PrerequisiteStage = in.PrerequisiteStage
	out.ComponentTransform = direct.Slice_FromProto(mapCtx, in.ComponentTransform, ExecutionStageSummary_ComponentTransform_FromProto)
	out.ComponentSource = direct.Slice_FromProto(mapCtx, in.ComponentSource, ExecutionStageSummary_ComponentSource_FromProto)
	return out
}
func ExecutionStageSummary_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionStageSummary) *pb.ExecutionStageSummary {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionStageSummary{}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.Kind = direct.Enum_ToProto[pb.KindType](mapCtx, in.Kind)
	out.InputSource = direct.Slice_ToProto(mapCtx, in.InputSource, ExecutionStageSummary_StageSource_ToProto)
	out.OutputSource = direct.Slice_ToProto(mapCtx, in.OutputSource, ExecutionStageSummary_StageSource_ToProto)
	out.PrerequisiteStage = in.PrerequisiteStage
	out.ComponentTransform = direct.Slice_ToProto(mapCtx, in.ComponentTransform, ExecutionStageSummary_ComponentTransform_ToProto)
	out.ComponentSource = direct.Slice_ToProto(mapCtx, in.ComponentSource, ExecutionStageSummary_ComponentSource_ToProto)
	return out
}
func ExecutionStageSummary_ComponentSource_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionStageSummary_ComponentSource) *krm.ExecutionStageSummary_ComponentSource {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionStageSummary_ComponentSource{}
	out.UserName = direct.LazyPtr(in.GetUserName())
	out.Name = direct.LazyPtr(in.GetName())
	out.OriginalTransformOrCollection = direct.LazyPtr(in.GetOriginalTransformOrCollection())
	return out
}
func ExecutionStageSummary_ComponentSource_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionStageSummary_ComponentSource) *pb.ExecutionStageSummary_ComponentSource {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionStageSummary_ComponentSource{}
	out.UserName = direct.ValueOf(in.UserName)
	out.Name = direct.ValueOf(in.Name)
	out.OriginalTransformOrCollection = direct.ValueOf(in.OriginalTransformOrCollection)
	return out
}
func ExecutionStageSummary_ComponentTransform_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionStageSummary_ComponentTransform) *krm.ExecutionStageSummary_ComponentTransform {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionStageSummary_ComponentTransform{}
	out.UserName = direct.LazyPtr(in.GetUserName())
	out.Name = direct.LazyPtr(in.GetName())
	out.OriginalTransform = direct.LazyPtr(in.GetOriginalTransform())
	return out
}
func ExecutionStageSummary_ComponentTransform_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionStageSummary_ComponentTransform) *pb.ExecutionStageSummary_ComponentTransform {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionStageSummary_ComponentTransform{}
	out.UserName = direct.ValueOf(in.UserName)
	out.Name = direct.ValueOf(in.Name)
	out.OriginalTransform = direct.ValueOf(in.OriginalTransform)
	return out
}
func ExecutionStageSummary_StageSource_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionStageSummary_StageSource) *krm.ExecutionStageSummary_StageSource {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionStageSummary_StageSource{}
	out.UserName = direct.LazyPtr(in.GetUserName())
	out.Name = direct.LazyPtr(in.GetName())
	out.OriginalTransformOrCollection = direct.LazyPtr(in.GetOriginalTransformOrCollection())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	return out
}
func ExecutionStageSummary_StageSource_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionStageSummary_StageSource) *pb.ExecutionStageSummary_StageSource {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionStageSummary_StageSource{}
	out.UserName = direct.ValueOf(in.UserName)
	out.Name = direct.ValueOf(in.Name)
	out.OriginalTransformOrCollection = direct.ValueOf(in.OriginalTransformOrCollection)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	return out
}
func FailedLocation_FromProto(mapCtx *direct.MapContext, in *pb.FailedLocation) *krm.FailedLocation {
	if in == nil {
		return nil
	}
	out := &krm.FailedLocation{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func FailedLocation_ToProto(mapCtx *direct.MapContext, in *krm.FailedLocation) *pb.FailedLocation {
	if in == nil {
		return nil
	}
	out := &pb.FailedLocation{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func FileIODetails_FromProto(mapCtx *direct.MapContext, in *pb.FileIODetails) *krm.FileIODetails {
	if in == nil {
		return nil
	}
	out := &krm.FileIODetails{}
	out.FilePattern = direct.LazyPtr(in.GetFilePattern())
	return out
}
func FileIODetails_ToProto(mapCtx *direct.MapContext, in *krm.FileIODetails) *pb.FileIODetails {
	if in == nil {
		return nil
	}
	out := &pb.FileIODetails{}
	out.FilePattern = direct.ValueOf(in.FilePattern)
	return out
}
func InvalidTemplateParameters_FromProto(mapCtx *direct.MapContext, in *pb.InvalidTemplateParameters) *krm.InvalidTemplateParameters {
	if in == nil {
		return nil
	}
	out := &krm.InvalidTemplateParameters{}
	out.ParameterViolations = direct.Slice_FromProto(mapCtx, in.ParameterViolations, InvalidTemplateParameters_ParameterViolation_FromProto)
	return out
}
func InvalidTemplateParameters_ToProto(mapCtx *direct.MapContext, in *krm.InvalidTemplateParameters) *pb.InvalidTemplateParameters {
	if in == nil {
		return nil
	}
	out := &pb.InvalidTemplateParameters{}
	out.ParameterViolations = direct.Slice_ToProto(mapCtx, in.ParameterViolations, InvalidTemplateParameters_ParameterViolation_ToProto)
	return out
}
func InvalidTemplateParameters_ParameterViolation_FromProto(mapCtx *direct.MapContext, in *pb.InvalidTemplateParameters_ParameterViolation) *krm.InvalidTemplateParameters_ParameterViolation {
	if in == nil {
		return nil
	}
	out := &krm.InvalidTemplateParameters_ParameterViolation{}
	out.Parameter = direct.LazyPtr(in.GetParameter())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func InvalidTemplateParameters_ParameterViolation_ToProto(mapCtx *direct.MapContext, in *krm.InvalidTemplateParameters_ParameterViolation) *pb.InvalidTemplateParameters_ParameterViolation {
	if in == nil {
		return nil
	}
	out := &pb.InvalidTemplateParameters_ParameterViolation{}
	out.Parameter = direct.ValueOf(in.Parameter)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func JobExecutionDetails_FromProto(mapCtx *direct.MapContext, in *pb.JobExecutionDetails) *krm.JobExecutionDetails {
	if in == nil {
		return nil
	}
	out := &krm.JobExecutionDetails{}
	out.Stages = direct.Slice_FromProto(mapCtx, in.Stages, StageSummary_FromProto)
	out.NextPageToken = direct.LazyPtr(in.GetNextPageToken())
	return out
}
func JobExecutionDetails_ToProto(mapCtx *direct.MapContext, in *krm.JobExecutionDetails) *pb.JobExecutionDetails {
	if in == nil {
		return nil
	}
	out := &pb.JobExecutionDetails{}
	out.Stages = direct.Slice_ToProto(mapCtx, in.Stages, StageSummary_ToProto)
	out.NextPageToken = direct.ValueOf(in.NextPageToken)
	return out
}
func JobExecutionInfo_FromProto(mapCtx *direct.MapContext, in *pb.JobExecutionInfo) *krm.JobExecutionInfo {
	if in == nil {
		return nil
	}
	out := &krm.JobExecutionInfo{}
	// MISSING: Stages
	return out
}
func JobExecutionInfo_ToProto(mapCtx *direct.MapContext, in *krm.JobExecutionInfo) *pb.JobExecutionInfo {
	if in == nil {
		return nil
	}
	out := &pb.JobExecutionInfo{}
	// MISSING: Stages
	return out
}
func JobExecutionStageInfo_FromProto(mapCtx *direct.MapContext, in *pb.JobExecutionStageInfo) *krm.JobExecutionStageInfo {
	if in == nil {
		return nil
	}
	out := &krm.JobExecutionStageInfo{}
	out.StepName = in.StepName
	return out
}
func JobExecutionStageInfo_ToProto(mapCtx *direct.MapContext, in *krm.JobExecutionStageInfo) *pb.JobExecutionStageInfo {
	if in == nil {
		return nil
	}
	out := &pb.JobExecutionStageInfo{}
	out.StepName = in.StepName
	return out
}
func JobMessage_FromProto(mapCtx *direct.MapContext, in *pb.JobMessage) *krm.JobMessage {
	if in == nil {
		return nil
	}
	out := &krm.JobMessage{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Time = JobMessage_Time_FromProto(mapCtx, in.GetTime())
	out.MessageText = direct.LazyPtr(in.GetMessageText())
	out.MessageImportance = direct.Enum_FromProto(mapCtx, in.GetMessageImportance())
	return out
}
func JobMessage_ToProto(mapCtx *direct.MapContext, in *krm.JobMessage) *pb.JobMessage {
	if in == nil {
		return nil
	}
	out := &pb.JobMessage{}
	out.Id = direct.ValueOf(in.ID)
	out.Time = JobMessage_Time_ToProto(mapCtx, in.Time)
	out.MessageText = direct.ValueOf(in.MessageText)
	out.MessageImportance = direct.Enum_ToProto[pb.JobMessageImportance](mapCtx, in.MessageImportance)
	return out
}
func KeyRangeDataDiskAssignment_FromProto(mapCtx *direct.MapContext, in *pb.KeyRangeDataDiskAssignment) *krm.KeyRangeDataDiskAssignment {
	if in == nil {
		return nil
	}
	out := &krm.KeyRangeDataDiskAssignment{}
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	out.DataDisk = direct.LazyPtr(in.GetDataDisk())
	return out
}
func KeyRangeDataDiskAssignment_ToProto(mapCtx *direct.MapContext, in *krm.KeyRangeDataDiskAssignment) *pb.KeyRangeDataDiskAssignment {
	if in == nil {
		return nil
	}
	out := &pb.KeyRangeDataDiskAssignment{}
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	out.DataDisk = direct.ValueOf(in.DataDisk)
	return out
}
func KeyRangeLocation_FromProto(mapCtx *direct.MapContext, in *pb.KeyRangeLocation) *krm.KeyRangeLocation {
	if in == nil {
		return nil
	}
	out := &krm.KeyRangeLocation{}
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	out.DeliveryEndpoint = direct.LazyPtr(in.GetDeliveryEndpoint())
	out.DataDisk = direct.LazyPtr(in.GetDataDisk())
	out.DeprecatedPersistentDirectory = direct.LazyPtr(in.GetDeprecatedPersistentDirectory())
	return out
}
func KeyRangeLocation_ToProto(mapCtx *direct.MapContext, in *krm.KeyRangeLocation) *pb.KeyRangeLocation {
	if in == nil {
		return nil
	}
	out := &pb.KeyRangeLocation{}
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
	out.DeliveryEndpoint = direct.ValueOf(in.DeliveryEndpoint)
	out.DataDisk = direct.ValueOf(in.DataDisk)
	out.DeprecatedPersistentDirectory = direct.ValueOf(in.DeprecatedPersistentDirectory)
	return out
}
func LaunchTemplateParameters_FromProto(mapCtx *direct.MapContext, in *pb.LaunchTemplateParameters) *krm.LaunchTemplateParameters {
	if in == nil {
		return nil
	}
	out := &krm.LaunchTemplateParameters{}
	out.JobName = direct.LazyPtr(in.GetJobName())
	out.Parameters = in.Parameters
	out.Environment = RuntimeEnvironment_FromProto(mapCtx, in.GetEnvironment())
	out.Update = direct.LazyPtr(in.GetUpdate())
	out.TransformNameMapping = in.TransformNameMapping
	return out
}
func LaunchTemplateParameters_ToProto(mapCtx *direct.MapContext, in *krm.LaunchTemplateParameters) *pb.LaunchTemplateParameters {
	if in == nil {
		return nil
	}
	out := &pb.LaunchTemplateParameters{}
	out.JobName = direct.ValueOf(in.JobName)
	out.Parameters = in.Parameters
	out.Environment = RuntimeEnvironment_ToProto(mapCtx, in.Environment)
	out.Update = direct.ValueOf(in.Update)
	out.TransformNameMapping = in.TransformNameMapping
	return out
}
func MetricStructuredName_FromProto(mapCtx *direct.MapContext, in *pb.MetricStructuredName) *krm.MetricStructuredName {
	if in == nil {
		return nil
	}
	out := &krm.MetricStructuredName{}
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.Name = direct.LazyPtr(in.GetName())
	out.Context = in.Context
	return out
}
func MetricStructuredName_ToProto(mapCtx *direct.MapContext, in *krm.MetricStructuredName) *pb.MetricStructuredName {
	if in == nil {
		return nil
	}
	out := &pb.MetricStructuredName{}
	out.Origin = direct.ValueOf(in.Origin)
	out.Name = direct.ValueOf(in.Name)
	out.Context = in.Context
	return out
}
func MountedDataDisk_FromProto(mapCtx *direct.MapContext, in *pb.MountedDataDisk) *krm.MountedDataDisk {
	if in == nil {
		return nil
	}
	out := &krm.MountedDataDisk{}
	out.DataDisk = direct.LazyPtr(in.GetDataDisk())
	return out
}
func MountedDataDisk_ToProto(mapCtx *direct.MapContext, in *krm.MountedDataDisk) *pb.MountedDataDisk {
	if in == nil {
		return nil
	}
	out := &pb.MountedDataDisk{}
	out.DataDisk = direct.ValueOf(in.DataDisk)
	return out
}
func Package_FromProto(mapCtx *direct.MapContext, in *pb.Package) *krm.Package {
	if in == nil {
		return nil
	}
	out := &krm.Package{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func Package_ToProto(mapCtx *direct.MapContext, in *krm.Package) *pb.Package {
	if in == nil {
		return nil
	}
	out := &pb.Package{}
	out.Name = direct.ValueOf(in.Name)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func PipelineDescription_FromProto(mapCtx *direct.MapContext, in *pb.PipelineDescription) *krm.PipelineDescription {
	if in == nil {
		return nil
	}
	out := &krm.PipelineDescription{}
	out.OriginalPipelineTransform = direct.Slice_FromProto(mapCtx, in.OriginalPipelineTransform, TransformSummary_FromProto)
	out.ExecutionPipelineStage = direct.Slice_FromProto(mapCtx, in.ExecutionPipelineStage, ExecutionStageSummary_FromProto)
	out.DisplayData = direct.Slice_FromProto(mapCtx, in.DisplayData, DisplayData_FromProto)
	return out
}
func PipelineDescription_ToProto(mapCtx *direct.MapContext, in *krm.PipelineDescription) *pb.PipelineDescription {
	if in == nil {
		return nil
	}
	out := &pb.PipelineDescription{}
	out.OriginalPipelineTransform = direct.Slice_ToProto(mapCtx, in.OriginalPipelineTransform, TransformSummary_ToProto)
	out.ExecutionPipelineStage = direct.Slice_ToProto(mapCtx, in.ExecutionPipelineStage, ExecutionStageSummary_ToProto)
	out.DisplayData = direct.Slice_ToProto(mapCtx, in.DisplayData, DisplayData_ToProto)
	return out
}
func ProgressTimeseries_FromProto(mapCtx *direct.MapContext, in *pb.ProgressTimeseries) *krm.ProgressTimeseries {
	if in == nil {
		return nil
	}
	out := &krm.ProgressTimeseries{}
	out.CurrentProgress = direct.LazyPtr(in.GetCurrentProgress())
	out.DataPoints = direct.Slice_FromProto(mapCtx, in.DataPoints, ProgressTimeseries_Point_FromProto)
	return out
}
func ProgressTimeseries_ToProto(mapCtx *direct.MapContext, in *krm.ProgressTimeseries) *pb.ProgressTimeseries {
	if in == nil {
		return nil
	}
	out := &pb.ProgressTimeseries{}
	out.CurrentProgress = direct.ValueOf(in.CurrentProgress)
	out.DataPoints = direct.Slice_ToProto(mapCtx, in.DataPoints, ProgressTimeseries_Point_ToProto)
	return out
}
func ProgressTimeseries_Point_FromProto(mapCtx *direct.MapContext, in *pb.ProgressTimeseries_Point) *krm.ProgressTimeseries_Point {
	if in == nil {
		return nil
	}
	out := &krm.ProgressTimeseries_Point{}
	out.Time = Point_Time_FromProto(mapCtx, in.GetTime())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func ProgressTimeseries_Point_ToProto(mapCtx *direct.MapContext, in *krm.ProgressTimeseries_Point) *pb.ProgressTimeseries_Point {
	if in == nil {
		return nil
	}
	out := &pb.ProgressTimeseries_Point{}
	out.Time = Point_Time_ToProto(mapCtx, in.Time)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func PubSubIODetails_FromProto(mapCtx *direct.MapContext, in *pb.PubSubIODetails) *krm.PubSubIODetails {
	if in == nil {
		return nil
	}
	out := &krm.PubSubIODetails{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	return out
}
func PubSubIODetails_ToProto(mapCtx *direct.MapContext, in *krm.PubSubIODetails) *pb.PubSubIODetails {
	if in == nil {
		return nil
	}
	out := &pb.PubSubIODetails{}
	out.Topic = direct.ValueOf(in.Topic)
	out.Subscription = direct.ValueOf(in.Subscription)
	return out
}
func PubsubLocation_FromProto(mapCtx *direct.MapContext, in *pb.PubsubLocation) *krm.PubsubLocation {
	if in == nil {
		return nil
	}
	out := &krm.PubsubLocation{}
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	out.TimestampLabel = direct.LazyPtr(in.GetTimestampLabel())
	out.IDLabel = direct.LazyPtr(in.GetIdLabel())
	out.DropLateData = direct.LazyPtr(in.GetDropLateData())
	out.TrackingSubscription = direct.LazyPtr(in.GetTrackingSubscription())
	out.WithAttributes = direct.LazyPtr(in.GetWithAttributes())
	return out
}
func PubsubLocation_ToProto(mapCtx *direct.MapContext, in *krm.PubsubLocation) *pb.PubsubLocation {
	if in == nil {
		return nil
	}
	out := &pb.PubsubLocation{}
	out.Topic = direct.ValueOf(in.Topic)
	out.Subscription = direct.ValueOf(in.Subscription)
	out.TimestampLabel = direct.ValueOf(in.TimestampLabel)
	out.IdLabel = direct.ValueOf(in.IDLabel)
	out.DropLateData = direct.ValueOf(in.DropLateData)
	out.TrackingSubscription = direct.ValueOf(in.TrackingSubscription)
	out.WithAttributes = direct.ValueOf(in.WithAttributes)
	return out
}
func RuntimeEnvironment_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeEnvironment) *krm.RuntimeEnvironment {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeEnvironment{}
	out.NumWorkers = direct.LazyPtr(in.GetNumWorkers())
	out.MaxWorkers = direct.LazyPtr(in.GetMaxWorkers())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.TempLocation = direct.LazyPtr(in.GetTempLocation())
	out.BypassTempDirValidation = direct.LazyPtr(in.GetBypassTempDirValidation())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AdditionalExperiments = in.AdditionalExperiments
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.AdditionalUserLabels = in.AdditionalUserLabels
	out.KmsKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.IpConfiguration = direct.Enum_FromProto(mapCtx, in.GetIpConfiguration())
	out.WorkerRegion = direct.LazyPtr(in.GetWorkerRegion())
	out.WorkerZone = direct.LazyPtr(in.GetWorkerZone())
	out.EnableStreamingEngine = direct.LazyPtr(in.GetEnableStreamingEngine())
	return out
}
func RuntimeEnvironment_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeEnvironment) *pb.RuntimeEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeEnvironment{}
	out.NumWorkers = direct.ValueOf(in.NumWorkers)
	out.MaxWorkers = direct.ValueOf(in.MaxWorkers)
	out.Zone = direct.ValueOf(in.Zone)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.TempLocation = direct.ValueOf(in.TempLocation)
	out.BypassTempDirValidation = direct.ValueOf(in.BypassTempDirValidation)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AdditionalExperiments = in.AdditionalExperiments
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.AdditionalUserLabels = in.AdditionalUserLabels
	out.KmsKeyName = direct.ValueOf(in.KmsKeyName)
	out.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IpConfiguration)
	out.WorkerRegion = direct.ValueOf(in.WorkerRegion)
	out.WorkerZone = direct.ValueOf(in.WorkerZone)
	out.EnableStreamingEngine = direct.ValueOf(in.EnableStreamingEngine)
	return out
}
func SDKInfo_FromProto(mapCtx *direct.MapContext, in *pb.SDKInfo) *krm.SDKInfo {
	if in == nil {
		return nil
	}
	out := &krm.SDKInfo{}
	out.Language = direct.Enum_FromProto(mapCtx, in.GetLanguage())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func SDKInfo_ToProto(mapCtx *direct.MapContext, in *krm.SDKInfo) *pb.SDKInfo {
	if in == nil {
		return nil
	}
	out := &pb.SDKInfo{}
	out.Language = direct.Enum_ToProto[pb.SDKInfo_Language](mapCtx, in.Language)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func SdkHarnessContainerImage_FromProto(mapCtx *direct.MapContext, in *pb.SdkHarnessContainerImage) *krm.SdkHarnessContainerImage {
	if in == nil {
		return nil
	}
	out := &krm.SdkHarnessContainerImage{}
	out.ContainerImage = direct.LazyPtr(in.GetContainerImage())
	out.UseSingleCorePerContainer = direct.LazyPtr(in.GetUseSingleCorePerContainer())
	out.EnvironmentID = direct.LazyPtr(in.GetEnvironmentId())
	out.Capabilities = in.Capabilities
	return out
}
func SdkHarnessContainerImage_ToProto(mapCtx *direct.MapContext, in *krm.SdkHarnessContainerImage) *pb.SdkHarnessContainerImage {
	if in == nil {
		return nil
	}
	out := &pb.SdkHarnessContainerImage{}
	out.ContainerImage = direct.ValueOf(in.ContainerImage)
	out.UseSingleCorePerContainer = direct.ValueOf(in.UseSingleCorePerContainer)
	out.EnvironmentId = direct.ValueOf(in.EnvironmentID)
	out.Capabilities = in.Capabilities
	return out
}
func SdkVersion_FromProto(mapCtx *direct.MapContext, in *pb.SdkVersion) *krm.SdkVersion {
	if in == nil {
		return nil
	}
	out := &krm.SdkVersion{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.VersionDisplayName = direct.LazyPtr(in.GetVersionDisplayName())
	out.SdkSupportStatus = direct.Enum_FromProto(mapCtx, in.GetSdkSupportStatus())
	return out
}
func SdkVersion_ToProto(mapCtx *direct.MapContext, in *krm.SdkVersion) *pb.SdkVersion {
	if in == nil {
		return nil
	}
	out := &pb.SdkVersion{}
	out.Version = direct.ValueOf(in.Version)
	out.VersionDisplayName = direct.ValueOf(in.VersionDisplayName)
	out.SdkSupportStatus = direct.Enum_ToProto[pb.SdkVersion_SdkSupportStatus](mapCtx, in.SdkSupportStatus)
	return out
}
func Snapshot_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.Snapshot {
	if in == nil {
		return nil
	}
	out := &krm.Snapshot{}
	out.ID = direct.LazyPtr(in.GetId())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.SourceJobID = direct.LazyPtr(in.GetSourceJobId())
	out.CreationTime = Snapshot_CreationTime_FromProto(mapCtx, in.GetCreationTime())
	out.Ttl = Snapshot_Ttl_FromProto(mapCtx, in.GetTtl())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PubsubMetadata = direct.Slice_FromProto(mapCtx, in.PubsubMetadata, PubsubSnapshotMetadata_FromProto)
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DiskSizeBytes = direct.LazyPtr(in.GetDiskSizeBytes())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func Snapshot_ToProto(mapCtx *direct.MapContext, in *krm.Snapshot) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	out.Id = direct.ValueOf(in.ID)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.SourceJobId = direct.ValueOf(in.SourceJobID)
	out.CreationTime = Snapshot_CreationTime_ToProto(mapCtx, in.CreationTime)
	out.Ttl = Snapshot_Ttl_ToProto(mapCtx, in.Ttl)
	out.State = direct.Enum_ToProto[pb.SnapshotState](mapCtx, in.State)
	out.PubsubMetadata = direct.Slice_ToProto(mapCtx, in.PubsubMetadata, PubsubSnapshotMetadata_ToProto)
	out.Description = direct.ValueOf(in.Description)
	out.DiskSizeBytes = direct.ValueOf(in.DiskSizeBytes)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func SpannerIODetails_FromProto(mapCtx *direct.MapContext, in *pb.SpannerIODetails) *krm.SpannerIODetails {
	if in == nil {
		return nil
	}
	out := &krm.SpannerIODetails{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.InstanceID = direct.LazyPtr(in.GetInstanceId())
	out.DatabaseID = direct.LazyPtr(in.GetDatabaseId())
	return out
}
func SpannerIODetails_ToProto(mapCtx *direct.MapContext, in *krm.SpannerIODetails) *pb.SpannerIODetails {
	if in == nil {
		return nil
	}
	out := &pb.SpannerIODetails{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.InstanceId = direct.ValueOf(in.InstanceID)
	out.DatabaseId = direct.ValueOf(in.DatabaseID)
	return out
}
func StageExecutionDetails_FromProto(mapCtx *direct.MapContext, in *pb.StageExecutionDetails) *krm.StageExecutionDetails {
	if in == nil {
		return nil
	}
	out := &krm.StageExecutionDetails{}
	out.Workers = direct.Slice_FromProto(mapCtx, in.Workers, WorkerDetails_FromProto)
	out.NextPageToken = direct.LazyPtr(in.GetNextPageToken())
	return out
}
func StageExecutionDetails_ToProto(mapCtx *direct.MapContext, in *krm.StageExecutionDetails) *pb.StageExecutionDetails {
	if in == nil {
		return nil
	}
	out := &pb.StageExecutionDetails{}
	out.Workers = direct.Slice_ToProto(mapCtx, in.Workers, WorkerDetails_ToProto)
	out.NextPageToken = direct.ValueOf(in.NextPageToken)
	return out
}
func StageSummary_FromProto(mapCtx *direct.MapContext, in *pb.StageSummary) *krm.StageSummary {
	if in == nil {
		return nil
	}
	out := &krm.StageSummary{}
	out.StageID = direct.LazyPtr(in.GetStageId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StartTime = StageSummary_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = StageSummary_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.Progress = ProgressTimeseries_FromProto(mapCtx, in.GetProgress())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, MetricUpdate_FromProto)
	return out
}
func StageSummary_ToProto(mapCtx *direct.MapContext, in *krm.StageSummary) *pb.StageSummary {
	if in == nil {
		return nil
	}
	out := &pb.StageSummary{}
	out.StageId = direct.ValueOf(in.StageID)
	out.State = direct.Enum_ToProto[pb.ExecutionState](mapCtx, in.State)
	out.StartTime = StageSummary_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = StageSummary_EndTime_ToProto(mapCtx, in.EndTime)
	out.Progress = ProgressTimeseries_ToProto(mapCtx, in.Progress)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, MetricUpdate_ToProto)
	return out
}
func StateFamilyConfig_FromProto(mapCtx *direct.MapContext, in *pb.StateFamilyConfig) *krm.StateFamilyConfig {
	if in == nil {
		return nil
	}
	out := &krm.StateFamilyConfig{}
	out.StateFamily = direct.LazyPtr(in.GetStateFamily())
	out.IsRead = direct.LazyPtr(in.GetIsRead())
	return out
}
func StateFamilyConfig_ToProto(mapCtx *direct.MapContext, in *krm.StateFamilyConfig) *pb.StateFamilyConfig {
	if in == nil {
		return nil
	}
	out := &pb.StateFamilyConfig{}
	out.StateFamily = direct.ValueOf(in.StateFamily)
	out.IsRead = direct.ValueOf(in.IsRead)
	return out
}
func StreamLocation_FromProto(mapCtx *direct.MapContext, in *pb.StreamLocation) *krm.StreamLocation {
	if in == nil {
		return nil
	}
	out := &krm.StreamLocation{}
	out.StreamingStageLocation = StreamingStageLocation_FromProto(mapCtx, in.GetStreamingStageLocation())
	out.PubsubLocation = PubsubLocation_FromProto(mapCtx, in.GetPubsubLocation())
	out.SideInputLocation = StreamingSideInputLocation_FromProto(mapCtx, in.GetSideInputLocation())
	out.CustomSourceLocation = CustomSourceLocation_FromProto(mapCtx, in.GetCustomSourceLocation())
	return out
}
func StreamLocation_ToProto(mapCtx *direct.MapContext, in *krm.StreamLocation) *pb.StreamLocation {
	if in == nil {
		return nil
	}
	out := &pb.StreamLocation{}
	if oneof := StreamingStageLocation_ToProto(mapCtx, in.StreamingStageLocation); oneof != nil {
		out.Location = &pb.StreamLocation_StreamingStageLocation{StreamingStageLocation: oneof}
	}
	if oneof := PubsubLocation_ToProto(mapCtx, in.PubsubLocation); oneof != nil {
		out.Location = &pb.StreamLocation_PubsubLocation{PubsubLocation: oneof}
	}
	if oneof := StreamingSideInputLocation_ToProto(mapCtx, in.SideInputLocation); oneof != nil {
		out.Location = &pb.StreamLocation_SideInputLocation{SideInputLocation: oneof}
	}
	if oneof := CustomSourceLocation_ToProto(mapCtx, in.CustomSourceLocation); oneof != nil {
		out.Location = &pb.StreamLocation_CustomSourceLocation{CustomSourceLocation: oneof}
	}
	return out
}
func StreamingApplianceSnapshotConfig_FromProto(mapCtx *direct.MapContext, in *pb.StreamingApplianceSnapshotConfig) *krm.StreamingApplianceSnapshotConfig {
	if in == nil {
		return nil
	}
	out := &krm.StreamingApplianceSnapshotConfig{}
	out.SnapshotID = direct.LazyPtr(in.GetSnapshotId())
	out.ImportStateEndpoint = direct.LazyPtr(in.GetImportStateEndpoint())
	return out
}
func StreamingApplianceSnapshotConfig_ToProto(mapCtx *direct.MapContext, in *krm.StreamingApplianceSnapshotConfig) *pb.StreamingApplianceSnapshotConfig {
	if in == nil {
		return nil
	}
	out := &pb.StreamingApplianceSnapshotConfig{}
	out.SnapshotId = direct.ValueOf(in.SnapshotID)
	out.ImportStateEndpoint = direct.ValueOf(in.ImportStateEndpoint)
	return out
}
func StreamingComputationRanges_FromProto(mapCtx *direct.MapContext, in *pb.StreamingComputationRanges) *krm.StreamingComputationRanges {
	if in == nil {
		return nil
	}
	out := &krm.StreamingComputationRanges{}
	out.ComputationID = direct.LazyPtr(in.GetComputationId())
	out.RangeAssignments = direct.Slice_FromProto(mapCtx, in.RangeAssignments, KeyRangeDataDiskAssignment_FromProto)
	return out
}
func StreamingComputationRanges_ToProto(mapCtx *direct.MapContext, in *krm.StreamingComputationRanges) *pb.StreamingComputationRanges {
	if in == nil {
		return nil
	}
	out := &pb.StreamingComputationRanges{}
	out.ComputationId = direct.ValueOf(in.ComputationID)
	out.RangeAssignments = direct.Slice_ToProto(mapCtx, in.RangeAssignments, KeyRangeDataDiskAssignment_ToProto)
	return out
}
func StreamingSideInputLocation_FromProto(mapCtx *direct.MapContext, in *pb.StreamingSideInputLocation) *krm.StreamingSideInputLocation {
	if in == nil {
		return nil
	}
	out := &krm.StreamingSideInputLocation{}
	out.Tag = direct.LazyPtr(in.GetTag())
	out.StateFamily = direct.LazyPtr(in.GetStateFamily())
	return out
}
func StreamingSideInputLocation_ToProto(mapCtx *direct.MapContext, in *krm.StreamingSideInputLocation) *pb.StreamingSideInputLocation {
	if in == nil {
		return nil
	}
	out := &pb.StreamingSideInputLocation{}
	out.Tag = direct.ValueOf(in.Tag)
	out.StateFamily = direct.ValueOf(in.StateFamily)
	return out
}
func StreamingStageLocation_FromProto(mapCtx *direct.MapContext, in *pb.StreamingStageLocation) *krm.StreamingStageLocation {
	if in == nil {
		return nil
	}
	out := &krm.StreamingStageLocation{}
	out.StreamID = direct.LazyPtr(in.GetStreamId())
	return out
}
func StreamingStageLocation_ToProto(mapCtx *direct.MapContext, in *krm.StreamingStageLocation) *pb.StreamingStageLocation {
	if in == nil {
		return nil
	}
	out := &pb.StreamingStageLocation{}
	out.StreamId = direct.ValueOf(in.StreamID)
	return out
}
func StructuredMessage_FromProto(mapCtx *direct.MapContext, in *pb.StructuredMessage) *krm.StructuredMessage {
	if in == nil {
		return nil
	}
	out := &krm.StructuredMessage{}
	out.MessageText = direct.LazyPtr(in.GetMessageText())
	out.MessageKey = direct.LazyPtr(in.GetMessageKey())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, StructuredMessage_Parameter_FromProto)
	return out
}
func StructuredMessage_ToProto(mapCtx *direct.MapContext, in *krm.StructuredMessage) *pb.StructuredMessage {
	if in == nil {
		return nil
	}
	out := &pb.StructuredMessage{}
	out.MessageText = direct.ValueOf(in.MessageText)
	out.MessageKey = direct.ValueOf(in.MessageKey)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, StructuredMessage_Parameter_ToProto)
	return out
}
func TaskRunnerSettings_FromProto(mapCtx *direct.MapContext, in *pb.TaskRunnerSettings) *krm.TaskRunnerSettings {
	if in == nil {
		return nil
	}
	out := &krm.TaskRunnerSettings{}
	out.TaskUser = direct.LazyPtr(in.GetTaskUser())
	out.TaskGroup = direct.LazyPtr(in.GetTaskGroup())
	out.OauthScopes = in.OauthScopes
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.DataflowApiVersion = direct.LazyPtr(in.GetDataflowApiVersion())
	out.ParallelWorkerSettings = WorkerSettings_FromProto(mapCtx, in.GetParallelWorkerSettings())
	out.BaseTaskDir = direct.LazyPtr(in.GetBaseTaskDir())
	out.ContinueOnException = direct.LazyPtr(in.GetContinueOnException())
	out.LogToSerialconsole = direct.LazyPtr(in.GetLogToSerialconsole())
	out.Alsologtostderr = direct.LazyPtr(in.GetAlsologtostderr())
	out.LogUploadLocation = direct.LazyPtr(in.GetLogUploadLocation())
	out.LogDir = direct.LazyPtr(in.GetLogDir())
	out.TempStoragePrefix = direct.LazyPtr(in.GetTempStoragePrefix())
	out.HarnessCommand = direct.LazyPtr(in.GetHarnessCommand())
	out.WorkflowFileName = direct.LazyPtr(in.GetWorkflowFileName())
	out.CommandlinesFileName = direct.LazyPtr(in.GetCommandlinesFileName())
	out.VmID = direct.LazyPtr(in.GetVmId())
	out.LanguageHint = direct.LazyPtr(in.GetLanguageHint())
	out.StreamingWorkerMainClass = direct.LazyPtr(in.GetStreamingWorkerMainClass())
	return out
}
func TaskRunnerSettings_ToProto(mapCtx *direct.MapContext, in *krm.TaskRunnerSettings) *pb.TaskRunnerSettings {
	if in == nil {
		return nil
	}
	out := &pb.TaskRunnerSettings{}
	out.TaskUser = direct.ValueOf(in.TaskUser)
	out.TaskGroup = direct.ValueOf(in.TaskGroup)
	out.OauthScopes = in.OauthScopes
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.DataflowApiVersion = direct.ValueOf(in.DataflowApiVersion)
	out.ParallelWorkerSettings = WorkerSettings_ToProto(mapCtx, in.ParallelWorkerSettings)
	out.BaseTaskDir = direct.ValueOf(in.BaseTaskDir)
	out.ContinueOnException = direct.ValueOf(in.ContinueOnException)
	out.LogToSerialconsole = direct.ValueOf(in.LogToSerialconsole)
	out.Alsologtostderr = direct.ValueOf(in.Alsologtostderr)
	out.LogUploadLocation = direct.ValueOf(in.LogUploadLocation)
	out.LogDir = direct.ValueOf(in.LogDir)
	out.TempStoragePrefix = direct.ValueOf(in.TempStoragePrefix)
	out.HarnessCommand = direct.ValueOf(in.HarnessCommand)
	out.WorkflowFileName = direct.ValueOf(in.WorkflowFileName)
	out.CommandlinesFileName = direct.ValueOf(in.CommandlinesFileName)
	out.VmId = direct.ValueOf(in.VmID)
	out.LanguageHint = direct.ValueOf(in.LanguageHint)
	out.StreamingWorkerMainClass = direct.ValueOf(in.StreamingWorkerMainClass)
	return out
}
func TopologyConfig_FromProto(mapCtx *direct.MapContext, in *pb.TopologyConfig) *krm.TopologyConfig {
	if in == nil {
		return nil
	}
	out := &krm.TopologyConfig{}
	out.Computations = direct.Slice_FromProto(mapCtx, in.Computations, ComputationTopology_FromProto)
	out.DataDiskAssignments = direct.Slice_FromProto(mapCtx, in.DataDiskAssignments, DataDiskAssignment_FromProto)
	out.UserStageToComputationNameMap = in.UserStageToComputationNameMap
	out.ForwardingKeyBits = direct.LazyPtr(in.GetForwardingKeyBits())
	out.PersistentStateVersion = direct.LazyPtr(in.GetPersistentStateVersion())
	return out
}
func TopologyConfig_ToProto(mapCtx *direct.MapContext, in *krm.TopologyConfig) *pb.TopologyConfig {
	if in == nil {
		return nil
	}
	out := &pb.TopologyConfig{}
	out.Computations = direct.Slice_ToProto(mapCtx, in.Computations, ComputationTopology_ToProto)
	out.DataDiskAssignments = direct.Slice_ToProto(mapCtx, in.DataDiskAssignments, DataDiskAssignment_ToProto)
	out.UserStageToComputationNameMap = in.UserStageToComputationNameMap
	out.ForwardingKeyBits = direct.ValueOf(in.ForwardingKeyBits)
	out.PersistentStateVersion = direct.ValueOf(in.PersistentStateVersion)
	return out
}
func TransformSummary_FromProto(mapCtx *direct.MapContext, in *pb.TransformSummary) *krm.TransformSummary {
	if in == nil {
		return nil
	}
	out := &krm.TransformSummary{}
	out.Kind = direct.Enum_FromProto(mapCtx, in.GetKind())
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayData = direct.Slice_FromProto(mapCtx, in.DisplayData, DisplayData_FromProto)
	out.OutputCollectionName = in.OutputCollectionName
	out.InputCollectionName = in.InputCollectionName
	return out
}
func TransformSummary_ToProto(mapCtx *direct.MapContext, in *krm.TransformSummary) *pb.TransformSummary {
	if in == nil {
		return nil
	}
	out := &pb.TransformSummary{}
	out.Kind = direct.Enum_ToProto[pb.KindType](mapCtx, in.Kind)
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	out.DisplayData = direct.Slice_ToProto(mapCtx, in.DisplayData, DisplayData_ToProto)
	out.OutputCollectionName = in.OutputCollectionName
	out.InputCollectionName = in.InputCollectionName
	return out
}
func WorkItemDetails_FromProto(mapCtx *direct.MapContext, in *pb.WorkItemDetails) *krm.WorkItemDetails {
	if in == nil {
		return nil
	}
	out := &krm.WorkItemDetails{}
	out.TaskID = direct.LazyPtr(in.GetTaskId())
	out.AttemptID = direct.LazyPtr(in.GetAttemptId())
	out.StartTime = WorkItemDetails_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = WorkItemDetails_EndTime_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Progress = ProgressTimeseries_FromProto(mapCtx, in.GetProgress())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, MetricUpdate_FromProto)
	return out
}
func WorkItemDetails_ToProto(mapCtx *direct.MapContext, in *krm.WorkItemDetails) *pb.WorkItemDetails {
	if in == nil {
		return nil
	}
	out := &pb.WorkItemDetails{}
	out.TaskId = direct.ValueOf(in.TaskID)
	out.AttemptId = direct.ValueOf(in.AttemptID)
	out.StartTime = WorkItemDetails_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = WorkItemDetails_EndTime_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.ExecutionState](mapCtx, in.State)
	out.Progress = ProgressTimeseries_ToProto(mapCtx, in.Progress)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, MetricUpdate_ToProto)
	return out
}
func WorkerDetails_FromProto(mapCtx *direct.MapContext, in *pb.WorkerDetails) *krm.WorkerDetails {
	if in == nil {
		return nil
	}
	out := &krm.WorkerDetails{}
	out.WorkerName = direct.LazyPtr(in.GetWorkerName())
	out.WorkItems = direct.Slice_FromProto(mapCtx, in.WorkItems, WorkItemDetails_FromProto)
	return out
}
func WorkerDetails_ToProto(mapCtx *direct.MapContext, in *krm.WorkerDetails) *pb.WorkerDetails {
	if in == nil {
		return nil
	}
	out := &pb.WorkerDetails{}
	out.WorkerName = direct.ValueOf(in.WorkerName)
	out.WorkItems = direct.Slice_ToProto(mapCtx, in.WorkItems, WorkItemDetails_ToProto)
	return out
}
func WorkerSettings_FromProto(mapCtx *direct.MapContext, in *pb.WorkerSettings) *krm.WorkerSettings {
	if in == nil {
		return nil
	}
	out := &krm.WorkerSettings{}
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.ReportingEnabled = direct.LazyPtr(in.GetReportingEnabled())
	out.ServicePath = direct.LazyPtr(in.GetServicePath())
	out.ShuffleServicePath = direct.LazyPtr(in.GetShuffleServicePath())
	out.WorkerID = direct.LazyPtr(in.GetWorkerId())
	out.TempStoragePrefix = direct.LazyPtr(in.GetTempStoragePrefix())
	return out
}
func WorkerSettings_ToProto(mapCtx *direct.MapContext, in *krm.WorkerSettings) *pb.WorkerSettings {
	if in == nil {
		return nil
	}
	out := &pb.WorkerSettings{}
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.ReportingEnabled = direct.ValueOf(in.ReportingEnabled)
	out.ServicePath = direct.ValueOf(in.ServicePath)
	out.ShuffleServicePath = direct.ValueOf(in.ShuffleServicePath)
	out.WorkerId = direct.ValueOf(in.WorkerID)
	out.TempStoragePrefix = direct.ValueOf(in.TempStoragePrefix)
	return out
}
