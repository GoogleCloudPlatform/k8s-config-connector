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
func ContainerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ContainerSpec) *krm.ContainerSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerSpec{}
	out.Image = direct.LazyPtr(in.GetImage())
	out.Metadata = TemplateMetadata_FromProto(mapCtx, in.GetMetadata())
	out.SdkInfo = SDKInfo_FromProto(mapCtx, in.GetSdkInfo())
	out.DefaultEnvironment = FlexTemplateRuntimeEnvironment_FromProto(mapCtx, in.GetDefaultEnvironment())
	return out
}
func ContainerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerSpec) *pb.ContainerSpec {
	if in == nil {
		return nil
	}
	out := &pb.ContainerSpec{}
	out.Image = direct.ValueOf(in.Image)
	out.Metadata = TemplateMetadata_ToProto(mapCtx, in.Metadata)
	out.SdkInfo = SDKInfo_ToProto(mapCtx, in.SdkInfo)
	out.DefaultEnvironment = FlexTemplateRuntimeEnvironment_ToProto(mapCtx, in.DefaultEnvironment)
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
func DataFlowFlexTemplateJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.FlexTemplateRuntimeEnvironment) *krm.DataFlowFlexTemplateJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataFlowFlexTemplateJobSpec{}
	out.NumWorkers = direct.LazyPtr(in.GetNumWorkers())
	out.MaxWorkers = direct.LazyPtr(in.GetMaxWorkers())
	// MISSING: Zone
	// MISSING: ServiceAccountEmail
	out.TempLocation = direct.LazyPtr(in.GetTempLocation())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AdditionalExperiments = in.AdditionalExperiments
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: AdditionalUserLabels
	// MISSING: KmsKeyName
	out.IpConfiguration = direct.Enum_FromProto(mapCtx, in.GetIpConfiguration())
	// MISSING: WorkerRegion
	// MISSING: WorkerZone
	out.EnableStreamingEngine = direct.LazyPtr(in.GetEnableStreamingEngine())
	// MISSING: FlexrsGoal
	out.StagingLocation = direct.LazyPtr(in.GetStagingLocation())
	out.SdkContainerImage = direct.LazyPtr(in.GetSdkContainerImage())
	// MISSING: DiskSizeGb
	out.AutoscalingAlgorithm = direct.Enum_FromProto(mapCtx, in.GetAutoscalingAlgorithm())
	// MISSING: DumpHeapOnOom
	// MISSING: SaveHeapDumpsToGcsPath
	out.LauncherMachineType = direct.LazyPtr(in.GetLauncherMachineType())
	return out
}
func DataFlowFlexTemplateJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataFlowFlexTemplateJobSpec) *pb.FlexTemplateRuntimeEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.FlexTemplateRuntimeEnvironment{}
	out.NumWorkers = direct.ValueOf(in.NumWorkers)
	out.MaxWorkers = direct.ValueOf(in.MaxWorkers)
	// MISSING: Zone
	// MISSING: ServiceAccountEmail
	out.TempLocation = direct.ValueOf(in.TempLocation)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AdditionalExperiments = in.AdditionalExperiments
	// MISSING: Network
	// MISSING: Subnetwork
	// MISSING: AdditionalUserLabels
	// MISSING: KmsKeyName
	out.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IpConfiguration)
	// MISSING: WorkerRegion
	// MISSING: WorkerZone
	out.EnableStreamingEngine = direct.ValueOf(in.EnableStreamingEngine)
	// MISSING: FlexrsGoal
	out.StagingLocation = direct.ValueOf(in.StagingLocation)
	out.SdkContainerImage = direct.ValueOf(in.SdkContainerImage)
	// MISSING: DiskSizeGb
	out.AutoscalingAlgorithm = direct.Enum_ToProto[pb.AutoscalingAlgorithm](mapCtx, in.AutoscalingAlgorithm)
	// MISSING: DumpHeapOnOom
	// MISSING: SaveHeapDumpsToGcsPath
	out.LauncherMachineType = direct.ValueOf(in.LauncherMachineType)
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
func DisplayData_FromProto(mapCtx *direct.MapContext, in *pb.DisplayData) *krm.DisplayData {
	if in == nil {
		return nil
	}
	out := &krm.DisplayData{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Namespace = direct.LazyPtr(in.GetNamespace())
	out.StrValue = direct.LazyPtr(in.GetStrValue())
	out.Int64Value = direct.LazyPtr(in.GetInt64Value())
	out.FloatValue = direct.LazyPtr(in.GetFloatValue())
	out.JavaClassValue = direct.LazyPtr(in.GetJavaClassValue())
	out.TimestampValue = DisplayData_TimestampValue_FromProto(mapCtx, in.GetTimestampValue())
	out.DurationValue = DisplayData_DurationValue_FromProto(mapCtx, in.GetDurationValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.ShortStrValue = direct.LazyPtr(in.GetShortStrValue())
	out.URL = direct.LazyPtr(in.GetUrl())
	out.Label = direct.LazyPtr(in.GetLabel())
	return out
}
func DisplayData_ToProto(mapCtx *direct.MapContext, in *krm.DisplayData) *pb.DisplayData {
	if in == nil {
		return nil
	}
	out := &pb.DisplayData{}
	out.Key = direct.ValueOf(in.Key)
	out.Namespace = direct.ValueOf(in.Namespace)
	if oneof := DisplayData_StrValue_ToProto(mapCtx, in.StrValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := DisplayData_Int64Value_ToProto(mapCtx, in.Int64Value); oneof != nil {
		out.Value = oneof
	}
	if oneof := DisplayData_FloatValue_ToProto(mapCtx, in.FloatValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := DisplayData_JavaClassValue_ToProto(mapCtx, in.JavaClassValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := DisplayData_TimestampValue_ToProto(mapCtx, in.TimestampValue); oneof != nil {
		out.Value = &pb.DisplayData_TimestampValue{TimestampValue: oneof}
	}
	if oneof := DisplayData_DurationValue_ToProto(mapCtx, in.DurationValue); oneof != nil {
		out.Value = &pb.DisplayData_DurationValue{DurationValue: oneof}
	}
	if oneof := DisplayData_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Value = oneof
	}
	out.ShortStrValue = direct.ValueOf(in.ShortStrValue)
	out.Url = direct.ValueOf(in.URL)
	out.Label = direct.ValueOf(in.Label)
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
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.TempStoragePrefix = direct.LazyPtr(in.GetTempStoragePrefix())
	out.ClusterManagerApiService = direct.LazyPtr(in.GetClusterManagerApiService())
	out.Experiments = in.Experiments
	out.ServiceOptions = in.ServiceOptions
	out.ServiceKmsKeyName = direct.LazyPtr(in.GetServiceKmsKeyName())
	out.WorkerPools = direct.Slice_FromProto(mapCtx, in.WorkerPools, WorkerPool_FromProto)
	out.UserAgent = google_protobuf_Struct_FromProto(mapCtx, in.GetUserAgent())
	out.Version = google_protobuf_Struct_FromProto(mapCtx, in.GetVersion())
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.SdkPipelineOptions = google_protobuf_Struct_FromProto(mapCtx, in.GetSdkPipelineOptions())
	out.InternalExperiments = google_protobuf_Any_FromProto(mapCtx, in.GetInternalExperiments())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.FlexResourceSchedulingGoal = direct.Enum_FromProto(mapCtx, in.GetFlexResourceSchedulingGoal())
	out.WorkerRegion = direct.LazyPtr(in.GetWorkerRegion())
	out.WorkerZone = direct.LazyPtr(in.GetWorkerZone())
	out.ShuffleMode = direct.Enum_FromProto(mapCtx, in.GetShuffleMode())
	out.DebugOptions = DebugOptions_FromProto(mapCtx, in.GetDebugOptions())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.TempStoragePrefix = direct.ValueOf(in.TempStoragePrefix)
	out.ClusterManagerApiService = direct.ValueOf(in.ClusterManagerApiService)
	out.Experiments = in.Experiments
	out.ServiceOptions = in.ServiceOptions
	out.ServiceKmsKeyName = direct.ValueOf(in.ServiceKmsKeyName)
	out.WorkerPools = direct.Slice_ToProto(mapCtx, in.WorkerPools, WorkerPool_ToProto)
	out.UserAgent = google_protobuf_Struct_ToProto(mapCtx, in.UserAgent)
	out.Version = google_protobuf_Struct_ToProto(mapCtx, in.Version)
	out.Dataset = direct.ValueOf(in.Dataset)
	out.SdkPipelineOptions = google_protobuf_Struct_ToProto(mapCtx, in.SdkPipelineOptions)
	out.InternalExperiments = google_protobuf_Any_ToProto(mapCtx, in.InternalExperiments)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.FlexResourceSchedulingGoal = direct.Enum_ToProto[pb.FlexResourceSchedulingGoal](mapCtx, in.FlexResourceSchedulingGoal)
	out.WorkerRegion = direct.ValueOf(in.WorkerRegion)
	out.WorkerZone = direct.ValueOf(in.WorkerZone)
	out.ShuffleMode = direct.Enum_ToProto[pb.ShuffleMode](mapCtx, in.ShuffleMode)
	out.DebugOptions = DebugOptions_ToProto(mapCtx, in.DebugOptions)
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
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	out.ID = direct.LazyPtr(in.GetId())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Step_FromProto)
	out.StepsLocation = direct.LazyPtr(in.GetStepsLocation())
	out.CurrentState = direct.Enum_FromProto(mapCtx, in.GetCurrentState())
	out.CurrentStateTime = Job_CurrentStateTime_FromProto(mapCtx, in.GetCurrentStateTime())
	out.RequestedState = direct.Enum_FromProto(mapCtx, in.GetRequestedState())
	out.ExecutionInfo = JobExecutionInfo_FromProto(mapCtx, in.GetExecutionInfo())
	out.CreateTime = Job_CreateTime_FromProto(mapCtx, in.GetCreateTime())
	out.ReplaceJobID = direct.LazyPtr(in.GetReplaceJobId())
	out.TransformNameMapping = in.TransformNameMapping
	out.ClientRequestID = direct.LazyPtr(in.GetClientRequestId())
	out.ReplacedByJobID = direct.LazyPtr(in.GetReplacedByJobId())
	out.TempFiles = in.TempFiles
	out.Labels = in.Labels
	out.Location = direct.LazyPtr(in.GetLocation())
	out.PipelineDescription = PipelineDescription_FromProto(mapCtx, in.GetPipelineDescription())
	out.StageStates = direct.Slice_FromProto(mapCtx, in.StageStates, ExecutionStageState_FromProto)
	out.JobMetadata = JobMetadata_FromProto(mapCtx, in.GetJobMetadata())
	out.StartTime = Job_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.CreatedFromSnapshotID = direct.LazyPtr(in.GetCreatedFromSnapshotId())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Id = direct.ValueOf(in.ID)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.JobType](mapCtx, in.Type)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Step_ToProto)
	out.StepsLocation = direct.ValueOf(in.StepsLocation)
	out.CurrentState = direct.Enum_ToProto[pb.JobState](mapCtx, in.CurrentState)
	out.CurrentStateTime = Job_CurrentStateTime_ToProto(mapCtx, in.CurrentStateTime)
	out.RequestedState = direct.Enum_ToProto[pb.JobState](mapCtx, in.RequestedState)
	out.ExecutionInfo = JobExecutionInfo_ToProto(mapCtx, in.ExecutionInfo)
	out.CreateTime = Job_CreateTime_ToProto(mapCtx, in.CreateTime)
	out.ReplaceJobId = direct.ValueOf(in.ReplaceJobID)
	out.TransformNameMapping = in.TransformNameMapping
	out.ClientRequestId = direct.ValueOf(in.ClientRequestID)
	out.ReplacedByJobId = direct.ValueOf(in.ReplacedByJobID)
	out.TempFiles = in.TempFiles
	out.Labels = in.Labels
	out.Location = direct.ValueOf(in.Location)
	out.PipelineDescription = PipelineDescription_ToProto(mapCtx, in.PipelineDescription)
	out.StageStates = direct.Slice_ToProto(mapCtx, in.StageStates, ExecutionStageState_ToProto)
	out.JobMetadata = JobMetadata_ToProto(mapCtx, in.JobMetadata)
	out.StartTime = Job_StartTime_ToProto(mapCtx, in.StartTime)
	out.CreatedFromSnapshotId = direct.ValueOf(in.CreatedFromSnapshotID)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
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
func JobMetrics_FromProto(mapCtx *direct.MapContext, in *pb.JobMetrics) *krm.JobMetrics {
	if in == nil {
		return nil
	}
	out := &krm.JobMetrics{}
	out.MetricTime = JobMetrics_MetricTime_FromProto(mapCtx, in.GetMetricTime())
	out.Metrics = direct.Slice_FromProto(mapCtx, in.Metrics, MetricUpdate_FromProto)
	return out
}
func JobMetrics_ToProto(mapCtx *direct.MapContext, in *krm.JobMetrics) *pb.JobMetrics {
	if in == nil {
		return nil
	}
	out := &pb.JobMetrics{}
	out.MetricTime = JobMetrics_MetricTime_ToProto(mapCtx, in.MetricTime)
	out.Metrics = direct.Slice_ToProto(mapCtx, in.Metrics, MetricUpdate_ToProto)
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
func LaunchFlexTemplateParameter_FromProto(mapCtx *direct.MapContext, in *pb.LaunchFlexTemplateParameter) *krm.LaunchFlexTemplateParameter {
	if in == nil {
		return nil
	}
	out := &krm.LaunchFlexTemplateParameter{}
	out.JobName = direct.LazyPtr(in.GetJobName())
	out.ContainerSpec = ContainerSpec_FromProto(mapCtx, in.GetContainerSpec())
	out.ContainerSpecGcsPath = direct.LazyPtr(in.GetContainerSpecGcsPath())
	out.Parameters = in.Parameters
	out.LaunchOptions = in.LaunchOptions
	out.Environment = FlexTemplateRuntimeEnvironment_FromProto(mapCtx, in.GetEnvironment())
	out.Update = direct.LazyPtr(in.GetUpdate())
	out.TransformNameMappings = in.TransformNameMappings
	return out
}
func LaunchFlexTemplateParameter_ToProto(mapCtx *direct.MapContext, in *krm.LaunchFlexTemplateParameter) *pb.LaunchFlexTemplateParameter {
	if in == nil {
		return nil
	}
	out := &pb.LaunchFlexTemplateParameter{}
	out.JobName = direct.ValueOf(in.JobName)
	if oneof := ContainerSpec_ToProto(mapCtx, in.ContainerSpec); oneof != nil {
		out.Template = &pb.LaunchFlexTemplateParameter_ContainerSpec{ContainerSpec: oneof}
	}
	if oneof := LaunchFlexTemplateParameter_ContainerSpecGcsPath_ToProto(mapCtx, in.ContainerSpecGcsPath); oneof != nil {
		out.Template = oneof
	}
	out.Parameters = in.Parameters
	out.LaunchOptions = in.LaunchOptions
	out.Environment = FlexTemplateRuntimeEnvironment_ToProto(mapCtx, in.Environment)
	out.Update = direct.ValueOf(in.Update)
	out.TransformNameMappings = in.TransformNameMappings
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
func MetricUpdate_FromProto(mapCtx *direct.MapContext, in *pb.MetricUpdate) *krm.MetricUpdate {
	if in == nil {
		return nil
	}
	out := &krm.MetricUpdate{}
	out.Name = MetricStructuredName_FromProto(mapCtx, in.GetName())
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Cumulative = direct.LazyPtr(in.GetCumulative())
	out.Scalar = google_protobuf_Value_FromProto(mapCtx, in.GetScalar())
	out.MeanSum = google_protobuf_Value_FromProto(mapCtx, in.GetMeanSum())
	out.MeanCount = google_protobuf_Value_FromProto(mapCtx, in.GetMeanCount())
	out.Set = google_protobuf_Value_FromProto(mapCtx, in.GetSet())
	out.Distribution = google_protobuf_Value_FromProto(mapCtx, in.GetDistribution())
	out.Gauge = google_protobuf_Value_FromProto(mapCtx, in.GetGauge())
	out.Internal = google_protobuf_Value_FromProto(mapCtx, in.GetInternal())
	out.UpdateTime = MetricUpdate_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func MetricUpdate_ToProto(mapCtx *direct.MapContext, in *krm.MetricUpdate) *pb.MetricUpdate {
	if in == nil {
		return nil
	}
	out := &pb.MetricUpdate{}
	out.Name = MetricStructuredName_ToProto(mapCtx, in.Name)
	out.Kind = direct.ValueOf(in.Kind)
	out.Cumulative = direct.ValueOf(in.Cumulative)
	out.Scalar = google_protobuf_Value_ToProto(mapCtx, in.Scalar)
	out.MeanSum = google_protobuf_Value_ToProto(mapCtx, in.MeanSum)
	out.MeanCount = google_protobuf_Value_ToProto(mapCtx, in.MeanCount)
	out.Set = google_protobuf_Value_ToProto(mapCtx, in.Set)
	out.Distribution = google_protobuf_Value_ToProto(mapCtx, in.Distribution)
	out.Gauge = google_protobuf_Value_ToProto(mapCtx, in.Gauge)
	out.Internal = google_protobuf_Value_ToProto(mapCtx, in.Internal)
	out.UpdateTime = MetricUpdate_UpdateTime_ToProto(mapCtx, in.UpdateTime)
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
func Step_FromProto(mapCtx *direct.MapContext, in *pb.Step) *krm.Step {
	if in == nil {
		return nil
	}
	out := &krm.Step{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.Name = direct.LazyPtr(in.GetName())
	out.Properties = google_protobuf_Struct_FromProto(mapCtx, in.GetProperties())
	return out
}
func Step_ToProto(mapCtx *direct.MapContext, in *krm.Step) *pb.Step {
	if in == nil {
		return nil
	}
	out := &pb.Step{}
	out.Kind = direct.ValueOf(in.Kind)
	out.Name = direct.ValueOf(in.Name)
	out.Properties = google_protobuf_Struct_ToProto(mapCtx, in.Properties)
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
func StructuredMessage_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.StructuredMessage_Parameter) *krm.StructuredMessage_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.StructuredMessage_Parameter{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = google_protobuf_Value_FromProto(mapCtx, in.GetValue())
	return out
}
func StructuredMessage_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.StructuredMessage_Parameter) *pb.StructuredMessage_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.StructuredMessage_Parameter{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = google_protobuf_Value_ToProto(mapCtx, in.Value)
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
func WorkerPool_FromProto(mapCtx *direct.MapContext, in *pb.WorkerPool) *krm.WorkerPool {
	if in == nil {
		return nil
	}
	out := &krm.WorkerPool{}
	out.Kind = direct.LazyPtr(in.GetKind())
	out.NumWorkers = direct.LazyPtr(in.GetNumWorkers())
	out.Packages = direct.Slice_FromProto(mapCtx, in.Packages, Package_FromProto)
	out.DefaultPackageSet = direct.Enum_FromProto(mapCtx, in.GetDefaultPackageSet())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.TeardownPolicy = direct.Enum_FromProto(mapCtx, in.GetTeardownPolicy())
	out.DiskSizeGb = direct.LazyPtr(in.GetDiskSizeGb())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.DiskSourceImage = direct.LazyPtr(in.GetDiskSourceImage())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.TaskrunnerSettings = TaskRunnerSettings_FromProto(mapCtx, in.GetTaskrunnerSettings())
	out.OnHostMaintenance = direct.LazyPtr(in.GetOnHostMaintenance())
	out.DataDisks = direct.Slice_FromProto(mapCtx, in.DataDisks, Disk_FromProto)
	out.Metadata = in.Metadata
	out.AutoscalingSettings = AutoscalingSettings_FromProto(mapCtx, in.GetAutoscalingSettings())
	out.PoolArgs = google_protobuf_Any_FromProto(mapCtx, in.GetPoolArgs())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.WorkerHarnessContainerImage = direct.LazyPtr(in.GetWorkerHarnessContainerImage())
	out.NumThreadsPerWorker = direct.LazyPtr(in.GetNumThreadsPerWorker())
	out.IpConfiguration = direct.Enum_FromProto(mapCtx, in.GetIpConfiguration())
	out.SdkHarnessContainerImages = direct.Slice_FromProto(mapCtx, in.SdkHarnessContainerImages, SdkHarnessContainerImage_FromProto)
	return out
}
func WorkerPool_ToProto(mapCtx *direct.MapContext, in *krm.WorkerPool) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	out.Kind = direct.ValueOf(in.Kind)
	out.NumWorkers = direct.ValueOf(in.NumWorkers)
	out.Packages = direct.Slice_ToProto(mapCtx, in.Packages, Package_ToProto)
	out.DefaultPackageSet = direct.Enum_ToProto[pb.DefaultPackageSet](mapCtx, in.DefaultPackageSet)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.TeardownPolicy = direct.Enum_ToProto[pb.TeardownPolicy](mapCtx, in.TeardownPolicy)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGb)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.DiskSourceImage = direct.ValueOf(in.DiskSourceImage)
	out.Zone = direct.ValueOf(in.Zone)
	out.TaskrunnerSettings = TaskRunnerSettings_ToProto(mapCtx, in.TaskrunnerSettings)
	out.OnHostMaintenance = direct.ValueOf(in.OnHostMaintenance)
	out.DataDisks = direct.Slice_ToProto(mapCtx, in.DataDisks, Disk_ToProto)
	out.Metadata = in.Metadata
	out.AutoscalingSettings = AutoscalingSettings_ToProto(mapCtx, in.AutoscalingSettings)
	out.PoolArgs = google_protobuf_Any_ToProto(mapCtx, in.PoolArgs)
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.WorkerHarnessContainerImage = direct.ValueOf(in.WorkerHarnessContainerImage)
	out.NumThreadsPerWorker = direct.ValueOf(in.NumThreadsPerWorker)
	out.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IpConfiguration)
	out.SdkHarnessContainerImages = direct.Slice_ToProto(mapCtx, in.SdkHarnessContainerImages, SdkHarnessContainerImage_ToProto)
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
