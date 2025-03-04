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
	"time"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Override because of unhandled Any / Value / Struct fields
func DisplayData_FromProto(mapCtx *direct.MapContext, in *pb.DisplayData) *krm.DisplayData {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func DisplayData_ToProto(mapCtx *direct.MapContext, in *krm.DisplayData) *pb.DisplayData {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func LaunchFlexTemplateParameter_FromProto(mapCtx *direct.MapContext, in *pb.LaunchFlexTemplateParameter) *krm.LaunchFlexTemplateParameter {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func LaunchFlexTemplateParameter_ToProto(mapCtx *direct.MapContext, in *krm.LaunchFlexTemplateParameter) *pb.LaunchFlexTemplateParameter {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func MetricUpdate_FromProto(mapCtx *direct.MapContext, in *pb.MetricUpdate) *krm.MetricUpdate {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func MetricUpdate_ToProto(mapCtx *direct.MapContext, in *krm.MetricUpdate) *pb.MetricUpdate {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func ContainerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ContainerSpec) *krm.ContainerSpec {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func ContainerSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerSpec) *pb.ContainerSpec {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func StructuredMessage_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.StructuredMessage_Parameter) *krm.StructuredMessage_Parameter {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func StructuredMessage_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.StructuredMessage_Parameter) *pb.StructuredMessage_Parameter {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Step_FromProto(mapCtx *direct.MapContext, in *pb.Step) *krm.Step {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func Step_ToProto(mapCtx *direct.MapContext, in *krm.Step) *pb.Step {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func JobMetrics_FromProto(mapCtx *direct.MapContext, in *pb.JobMetrics) *krm.JobMetrics {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func JobMetrics_ToProto(mapCtx *direct.MapContext, in *krm.JobMetrics) *pb.JobMetrics {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func WorkerPool_ToProto(mapCtx *direct.MapContext, in *krm.WorkerPool) *pb.WorkerPool {
	mapCtx.NotImplemented()
	return nil
}

// Override because of unhandled Any / Value / Struct fields
func WorkerPool_FromProto(mapCtx *direct.MapContext, in *pb.WorkerPool) *krm.WorkerPool {
	mapCtx.NotImplemented()
	return nil
}

func PubsubSnapshotMetadata_FromProto(mapCtx *direct.MapContext, in *pb.PubsubSnapshotMetadata) *krm.PubsubSnapshotMetadata {
	mapCtx.NotImplemented()
	return nil
}

func PubsubSnapshotMetadata_ToProto(mapCtx *direct.MapContext, in *krm.PubsubSnapshotMetadata) *pb.PubsubSnapshotMetadata {
	mapCtx.NotImplemented()
	return nil
}

func AutoscalingEvent_Time_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func AutoscalingEvent_Time_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func ExecutionStageState_CurrentStateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func ExecutionStageState_CurrentStateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Job_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func Job_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Job_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func Job_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func JobMessage_Time_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func JobMessage_Time_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func JobMetrics_MetricTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func JobMetrics_MetricsTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Point_Time_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func Point_Time_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Snapshot_CreationTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func Snapshot_CreationTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Snapshot_Ttl_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}
func Snapshot_Ttl_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func StageSummary_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func StageSummary_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func StageSummary_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func StageSummary_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func WorkItemDetails_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func WorkItemDetails_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}
func WorkItemDetails_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return Timestamp_FromProto(mapCtx, in)
}
func WorkItemDetails_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return Timestamp_ToProto(mapCtx, in)
}

func Timestamp_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	if in == nil {
		return nil
	}
	t := in.AsTime()
	s := t.Format(time.RFC3339Nano)
	return &s
}

func Timestamp_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	if in == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *in)
	if err != nil {
		mapCtx.Errorf("parsing timestamp %q", *in)
		return nil
	}
	return timestamppb.New(t)
}

// Override because of refs
func DataflowFlexTemplateJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.FlexTemplateRuntimeEnvironment) *krm.DataflowFlexTemplateJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataflowFlexTemplateJobSpec{}
	out.NumWorkers = direct.LazyPtr(in.GetNumWorkers())
	out.MaxWorkers = direct.LazyPtr(in.GetMaxWorkers())
	// MISSING: Zone
	out.TempLocation = direct.LazyPtr(in.GetTempLocation())
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AdditionalExperiments = in.AdditionalExperiments

	if in.Network != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{
			External: in.Network,
		}
	}

	if in.Subnetwork != "" {
		out.SubnetworkRef = &refs.ComputeSubnetworkRef{
			External: in.Subnetwork,
		}
	}
	if in.KmsKeyName != "" {
		out.KmsKeyNameRef = &kmsv1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		}
	}
	if in.ServiceAccountEmail != "" {
		out.ServiceAccountEmailRef = &refs.IAMServiceAccountRef{
			External: in.ServiceAccountEmail,
		}
	}

	// MISSING: AdditionalUserLabels
	out.IPConfiguration = direct.Enum_FromProto(mapCtx, in.GetIpConfiguration())
	// MISSING: WorkerRegion
	// MISSING: WorkerZone
	out.EnableStreamingEngine = direct.LazyPtr(in.GetEnableStreamingEngine())
	// MISSING: FlexrsGoal
	out.StagingLocation = direct.LazyPtr(in.GetStagingLocation())
	out.SDKContainerImage = direct.LazyPtr(in.GetSdkContainerImage())
	// MISSING: DiskSizeGb
	out.AutoscalingAlgorithm = direct.Enum_FromProto(mapCtx, in.GetAutoscalingAlgorithm())
	// MISSING: DumpHeapOnOom
	// MISSING: SaveHeapDumpsToGcsPath
	out.LauncherMachineType = direct.LazyPtr(in.GetLauncherMachineType())
	return out
}
func DataflowFlexTemplateJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataflowFlexTemplateJobSpec) *pb.FlexTemplateRuntimeEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.FlexTemplateRuntimeEnvironment{}
	out.NumWorkers = direct.ValueOf(in.NumWorkers)
	out.MaxWorkers = direct.ValueOf(in.MaxWorkers)
	// MISSING: Zone
	out.TempLocation = direct.ValueOf(in.TempLocation)
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AdditionalExperiments = in.AdditionalExperiments

	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}

	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}

	if in.KmsKeyNameRef != nil {
		out.KmsKeyName = in.KmsKeyNameRef.External
	}

	if in.ServiceAccountEmailRef != nil {
		out.ServiceAccountEmail = in.ServiceAccountEmailRef.External
	}

	// MISSING: AdditionalUserLabels
	out.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IPConfiguration)
	// MISSING: WorkerRegion
	// MISSING: WorkerZone
	out.EnableStreamingEngine = direct.ValueOf(in.EnableStreamingEngine)
	// MISSING: FlexrsGoal
	out.StagingLocation = direct.ValueOf(in.StagingLocation)
	out.SdkContainerImage = direct.ValueOf(in.SDKContainerImage)
	// MISSING: DiskSizeGb
	out.AutoscalingAlgorithm = direct.Enum_ToProto[pb.AutoscalingAlgorithm](mapCtx, in.AutoscalingAlgorithm)
	// MISSING: DumpHeapOnOom
	// MISSING: SaveHeapDumpsToGcsPath
	out.LauncherMachineType = direct.ValueOf(in.LauncherMachineType)
	return out
}
