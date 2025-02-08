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

package run

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudSqlInstance_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlInstance) *krm.CloudSqlInstance {
	if in == nil {
		return nil
	}
	out := &krm.CloudSqlInstance{}
	out.Instances = in.Instances
	return out
}
func CloudSqlInstance_ToProto(mapCtx *direct.MapContext, in *krm.CloudSqlInstance) *pb.CloudSqlInstance {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlInstance{}
	out.Instances = in.Instances
	return out
}
func Condition_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krm.Condition {
	if in == nil {
		return nil
	}
	out := &krm.Condition{}
	out.Type = direct.LazyPtr(in.GetType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.LastTransitionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastTransitionTime())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	// MISSING: Reason
	// MISSING: RevisionReason
	// MISSING: ExecutionReason
	return out
}
func Condition_ToProto(mapCtx *direct.MapContext, in *krm.Condition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	out.Type = direct.ValueOf(in.Type)
	out.State = direct.Enum_ToProto[pb.Condition_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.LastTransitionTime = direct.StringTimestamp_ToProto(mapCtx, in.LastTransitionTime)
	out.Severity = direct.Enum_ToProto[pb.Condition_Severity](mapCtx, in.Severity)
	// MISSING: Reason
	// MISSING: RevisionReason
	// MISSING: ExecutionReason
	return out
}
func ConditionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Condition) *krm.ConditionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConditionObservedState{}
	// MISSING: Type
	// MISSING: State
	// MISSING: Message
	// MISSING: LastTransitionTime
	// MISSING: Severity
	out.Reason = direct.Enum_FromProto(mapCtx, in.GetReason())
	out.RevisionReason = direct.Enum_FromProto(mapCtx, in.GetRevisionReason())
	out.ExecutionReason = direct.Enum_FromProto(mapCtx, in.GetExecutionReason())
	return out
}
func ConditionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConditionObservedState) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	// MISSING: Type
	// MISSING: State
	// MISSING: Message
	// MISSING: LastTransitionTime
	// MISSING: Severity
	if oneof := ConditionObservedState_Reason_ToProto(mapCtx, in.Reason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := ConditionObservedState_RevisionReason_ToProto(mapCtx, in.RevisionReason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := ConditionObservedState_ExecutionReason_ToProto(mapCtx, in.ExecutionReason); oneof != nil {
		out.Reasons = oneof
	}
	return out
}
func Container_FromProto(mapCtx *direct.MapContext, in *pb.Container) *krm.Container {
	if in == nil {
		return nil
	}
	out := &krm.Container{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Image = direct.LazyPtr(in.GetImage())
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_FromProto)
	out.Resources = ResourceRequirements_FromProto(mapCtx, in.GetResources())
	out.Ports = direct.Slice_FromProto(mapCtx, in.Ports, ContainerPort_FromProto)
	out.VolumeMounts = direct.Slice_FromProto(mapCtx, in.VolumeMounts, VolumeMount_FromProto)
	out.WorkingDir = direct.LazyPtr(in.GetWorkingDir())
	out.LivenessProbe = Probe_FromProto(mapCtx, in.GetLivenessProbe())
	out.StartupProbe = Probe_FromProto(mapCtx, in.GetStartupProbe())
	out.DependsOn = in.DependsOn
	return out
}
func Container_ToProto(mapCtx *direct.MapContext, in *krm.Container) *pb.Container {
	if in == nil {
		return nil
	}
	out := &pb.Container{}
	out.Name = direct.ValueOf(in.Name)
	out.Image = direct.ValueOf(in.Image)
	out.Command = in.Command
	out.Args = in.Args
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
	out.Resources = ResourceRequirements_ToProto(mapCtx, in.Resources)
	out.Ports = direct.Slice_ToProto(mapCtx, in.Ports, ContainerPort_ToProto)
	out.VolumeMounts = direct.Slice_ToProto(mapCtx, in.VolumeMounts, VolumeMount_ToProto)
	out.WorkingDir = direct.ValueOf(in.WorkingDir)
	out.LivenessProbe = Probe_ToProto(mapCtx, in.LivenessProbe)
	out.StartupProbe = Probe_ToProto(mapCtx, in.StartupProbe)
	out.DependsOn = in.DependsOn
	return out
}
func ContainerPort_FromProto(mapCtx *direct.MapContext, in *pb.ContainerPort) *krm.ContainerPort {
	if in == nil {
		return nil
	}
	out := &krm.ContainerPort{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ContainerPort = direct.LazyPtr(in.GetContainerPort())
	return out
}
func ContainerPort_ToProto(mapCtx *direct.MapContext, in *krm.ContainerPort) *pb.ContainerPort {
	if in == nil {
		return nil
	}
	out := &pb.ContainerPort{}
	out.Name = direct.ValueOf(in.Name)
	out.ContainerPort = direct.ValueOf(in.ContainerPort)
	return out
}
func EmptyDirVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.EmptyDirVolumeSource) *krm.EmptyDirVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.EmptyDirVolumeSource{}
	out.Medium = direct.Enum_FromProto(mapCtx, in.GetMedium())
	out.SizeLimit = direct.LazyPtr(in.GetSizeLimit())
	return out
}
func EmptyDirVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.EmptyDirVolumeSource) *pb.EmptyDirVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.EmptyDirVolumeSource{}
	out.Medium = direct.Enum_ToProto[pb.EmptyDirVolumeSource_Medium](mapCtx, in.Medium)
	out.SizeLimit = direct.ValueOf(in.SizeLimit)
	return out
}
func EnvVar_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar) *krm.EnvVar {
	if in == nil {
		return nil
	}
	out := &krm.EnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	out.ValueSource = EnvVarSource_FromProto(mapCtx, in.GetValueSource())
	return out
}
func EnvVar_ToProto(mapCtx *direct.MapContext, in *krm.EnvVar) *pb.EnvVar {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := EnvVar_Value_ToProto(mapCtx, in.Value); oneof != nil {
		out.Values = oneof
	}
	if oneof := EnvVarSource_ToProto(mapCtx, in.ValueSource); oneof != nil {
		out.Values = &pb.EnvVar_ValueSource{ValueSource: oneof}
	}
	return out
}
func EnvVarSource_FromProto(mapCtx *direct.MapContext, in *pb.EnvVarSource) *krm.EnvVarSource {
	if in == nil {
		return nil
	}
	out := &krm.EnvVarSource{}
	out.SecretKeyRef = SecretKeySelector_FromProto(mapCtx, in.GetSecretKeyRef())
	return out
}
func EnvVarSource_ToProto(mapCtx *direct.MapContext, in *krm.EnvVarSource) *pb.EnvVarSource {
	if in == nil {
		return nil
	}
	out := &pb.EnvVarSource{}
	out.SecretKeyRef = SecretKeySelector_ToProto(mapCtx, in.SecretKeyRef)
	return out
}
func GCSVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.GCSVolumeSource) *krm.GCSVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.GCSVolumeSource{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.ReadOnly = direct.LazyPtr(in.GetReadOnly())
	out.MountOptions = in.MountOptions
	return out
}
func GCSVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.GCSVolumeSource) *pb.GCSVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.GCSVolumeSource{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.ReadOnly = direct.ValueOf(in.ReadOnly)
	out.MountOptions = in.MountOptions
	return out
}
func GRPCAction_FromProto(mapCtx *direct.MapContext, in *pb.GRPCAction) *krm.GRPCAction {
	if in == nil {
		return nil
	}
	out := &krm.GRPCAction{}
	out.Port = direct.LazyPtr(in.GetPort())
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func GRPCAction_ToProto(mapCtx *direct.MapContext, in *krm.GRPCAction) *pb.GRPCAction {
	if in == nil {
		return nil
	}
	out := &pb.GRPCAction{}
	out.Port = direct.ValueOf(in.Port)
	out.Service = direct.ValueOf(in.Service)
	return out
}
func HTTPGetAction_FromProto(mapCtx *direct.MapContext, in *pb.HTTPGetAction) *krm.HTTPGetAction {
	if in == nil {
		return nil
	}
	out := &krm.HTTPGetAction{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.HTTPHeaders = direct.Slice_FromProto(mapCtx, in.HTTPHeaders, HTTPHeader_FromProto)
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func HTTPGetAction_ToProto(mapCtx *direct.MapContext, in *krm.HTTPGetAction) *pb.HTTPGetAction {
	if in == nil {
		return nil
	}
	out := &pb.HTTPGetAction{}
	out.Path = direct.ValueOf(in.Path)
	out.HttpHeaders = direct.Slice_ToProto(mapCtx, in.HTTPHeaders, HTTPHeader_ToProto)
	out.Port = direct.ValueOf(in.Port)
	return out
}
func HTTPHeader_FromProto(mapCtx *direct.MapContext, in *pb.HTTPHeader) *krm.HTTPHeader {
	if in == nil {
		return nil
	}
	out := &krm.HTTPHeader{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func HTTPHeader_ToProto(mapCtx *direct.MapContext, in *krm.HTTPHeader) *pb.HTTPHeader {
	if in == nil {
		return nil
	}
	out := &pb.HTTPHeader{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func NFSVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.NFSVolumeSource) *krm.NFSVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.NFSVolumeSource{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.Path = direct.LazyPtr(in.GetPath())
	out.ReadOnly = direct.LazyPtr(in.GetReadOnly())
	return out
}
func NFSVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.NFSVolumeSource) *pb.NFSVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.NFSVolumeSource{}
	out.Server = direct.ValueOf(in.Server)
	out.Path = direct.ValueOf(in.Path)
	out.ReadOnly = direct.ValueOf(in.ReadOnly)
	return out
}
func Probe_FromProto(mapCtx *direct.MapContext, in *pb.Probe) *krm.Probe {
	if in == nil {
		return nil
	}
	out := &krm.Probe{}
	out.InitialDelaySeconds = direct.LazyPtr(in.GetInitialDelaySeconds())
	out.TimeoutSeconds = direct.LazyPtr(in.GetTimeoutSeconds())
	out.PeriodSeconds = direct.LazyPtr(in.GetPeriodSeconds())
	out.FailureThreshold = direct.LazyPtr(in.GetFailureThreshold())
	out.HTTPGet = HTTPGetAction_FromProto(mapCtx, in.GetHttpGet())
	out.TcpSocket = TCPSocketAction_FromProto(mapCtx, in.GetTcpSocket())
	out.Grpc = GRPCAction_FromProto(mapCtx, in.GetGrpc())
	return out
}
func Probe_ToProto(mapCtx *direct.MapContext, in *krm.Probe) *pb.Probe {
	if in == nil {
		return nil
	}
	out := &pb.Probe{}
	out.InitialDelaySeconds = direct.ValueOf(in.InitialDelaySeconds)
	out.TimeoutSeconds = direct.ValueOf(in.TimeoutSeconds)
	out.PeriodSeconds = direct.ValueOf(in.PeriodSeconds)
	out.FailureThreshold = direct.ValueOf(in.FailureThreshold)
	if oneof := HTTPGetAction_ToProto(mapCtx, in.HTTPGet); oneof != nil {
		out.ProbeType = &pb.Probe_HttpGet{HttpGet: oneof}
	}
	if oneof := TCPSocketAction_ToProto(mapCtx, in.TcpSocket); oneof != nil {
		out.ProbeType = &pb.Probe_TcpSocket{TcpSocket: oneof}
	}
	if oneof := GRPCAction_ToProto(mapCtx, in.Grpc); oneof != nil {
		out.ProbeType = &pb.Probe_Grpc{Grpc: oneof}
	}
	return out
}
func ResourceRequirements_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRequirements) *krm.ResourceRequirements {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRequirements{}
	out.Limits = in.Limits
	out.CpuIdle = direct.LazyPtr(in.GetCpuIdle())
	out.StartupCpuBoost = direct.LazyPtr(in.GetStartupCpuBoost())
	return out
}
func ResourceRequirements_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRequirements) *pb.ResourceRequirements {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRequirements{}
	out.Limits = in.Limits
	out.CpuIdle = direct.ValueOf(in.CpuIdle)
	out.StartupCpuBoost = direct.ValueOf(in.StartupCpuBoost)
	return out
}
func RunTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.RunTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RunTaskObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func RunTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RunTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func RunTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.RunTaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.RunTaskSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func RunTaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.RunTaskSpec) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func SecretKeySelector_FromProto(mapCtx *direct.MapContext, in *pb.SecretKeySelector) *krm.SecretKeySelector {
	if in == nil {
		return nil
	}
	out := &krm.SecretKeySelector{}
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func SecretKeySelector_ToProto(mapCtx *direct.MapContext, in *krm.SecretKeySelector) *pb.SecretKeySelector {
	if in == nil {
		return nil
	}
	out := &pb.SecretKeySelector{}
	out.Secret = direct.ValueOf(in.Secret)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func SecretVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.SecretVolumeSource) *krm.SecretVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.SecretVolumeSource{}
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, VersionToPath_FromProto)
	out.DefaultMode = direct.LazyPtr(in.GetDefaultMode())
	return out
}
func SecretVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.SecretVolumeSource) *pb.SecretVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.SecretVolumeSource{}
	out.Secret = direct.ValueOf(in.Secret)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, VersionToPath_ToProto)
	out.DefaultMode = direct.ValueOf(in.DefaultMode)
	return out
}
func TCPSocketAction_FromProto(mapCtx *direct.MapContext, in *pb.TCPSocketAction) *krm.TCPSocketAction {
	if in == nil {
		return nil
	}
	out := &krm.TCPSocketAction{}
	out.Port = direct.LazyPtr(in.GetPort())
	return out
}
func TCPSocketAction_ToProto(mapCtx *direct.MapContext, in *krm.TCPSocketAction) *pb.TCPSocketAction {
	if in == nil {
		return nil
	}
	out := &pb.TCPSocketAction{}
	out.Port = direct.ValueOf(in.Port)
	return out
}
func Task_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.Task {
	if in == nil {
		return nil
	}
	out := &krm.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	out.Containers = direct.Slice_FromProto(mapCtx, in.Containers, Container_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	out.MaxRetries = direct.LazyPtr(in.GetMaxRetries())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ExecutionEnvironment = direct.Enum_FromProto(mapCtx, in.GetExecutionEnvironment())
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func Task_ToProto(mapCtx *direct.MapContext, in *krm.Task) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Generation
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: CreateTime
	// MISSING: ScheduledTime
	// MISSING: StartTime
	// MISSING: CompletionTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: ExpireTime
	// MISSING: Job
	// MISSING: Execution
	out.Containers = direct.Slice_ToProto(mapCtx, in.Containers, Container_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	out.MaxRetries = direct.ValueOf(in.MaxRetries)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ExecutionEnvironment = direct.Enum_ToProto[pb.ExecutionEnvironment](mapCtx, in.ExecutionEnvironment)
	// MISSING: Reconciling
	// MISSING: Conditions
	// MISSING: ObservedGeneration
	// MISSING: Index
	// MISSING: Retried
	// MISSING: LastAttemptResult
	// MISSING: EncryptionKey
	// MISSING: VpcAccess
	// MISSING: LogURI
	// MISSING: SatisfiesPzs
	// MISSING: Etag
	return out
}
func TaskAttemptResult_FromProto(mapCtx *direct.MapContext, in *pb.TaskAttemptResult) *krm.TaskAttemptResult {
	if in == nil {
		return nil
	}
	out := &krm.TaskAttemptResult{}
	// MISSING: Status
	// MISSING: ExitCode
	return out
}
func TaskAttemptResult_ToProto(mapCtx *direct.MapContext, in *krm.TaskAttemptResult) *pb.TaskAttemptResult {
	if in == nil {
		return nil
	}
	out := &pb.TaskAttemptResult{}
	// MISSING: Status
	// MISSING: ExitCode
	return out
}
func TaskAttemptResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TaskAttemptResult) *krm.TaskAttemptResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaskAttemptResultObservedState{}
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	out.ExitCode = direct.LazyPtr(in.GetExitCode())
	return out
}
func TaskAttemptResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaskAttemptResultObservedState) *pb.TaskAttemptResult {
	if in == nil {
		return nil
	}
	out := &pb.TaskAttemptResult{}
	out.Status = Status_ToProto(mapCtx, in.Status)
	out.ExitCode = direct.ValueOf(in.ExitCode)
	return out
}
func TaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.TaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaskObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ScheduledTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduledTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.CompletionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompletionTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Job = direct.LazyPtr(in.GetJob())
	out.Execution = direct.LazyPtr(in.GetExecution())
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Conditions = direct.Slice_FromProto(mapCtx, in.Conditions, Condition_FromProto)
	out.ObservedGeneration = direct.LazyPtr(in.GetObservedGeneration())
	out.Index = direct.LazyPtr(in.GetIndex())
	out.Retried = direct.LazyPtr(in.GetRetried())
	out.LastAttemptResult = TaskAttemptResult_FromProto(mapCtx, in.GetLastAttemptResult())
	out.EncryptionKey = direct.LazyPtr(in.GetEncryptionKey())
	out.VpcAccess = VpcAccess_FromProto(mapCtx, in.GetVpcAccess())
	out.LogURI = direct.LazyPtr(in.GetLogUri())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func TaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.Generation = direct.ValueOf(in.Generation)
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.ScheduledTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduledTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.CompletionTime = direct.StringTimestamp_ToProto(mapCtx, in.CompletionTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Job = direct.ValueOf(in.Job)
	out.Execution = direct.ValueOf(in.Execution)
	// MISSING: Containers
	// MISSING: Volumes
	// MISSING: MaxRetries
	// MISSING: Timeout
	// MISSING: ServiceAccount
	// MISSING: ExecutionEnvironment
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.Conditions = direct.Slice_ToProto(mapCtx, in.Conditions, Condition_ToProto)
	out.ObservedGeneration = direct.ValueOf(in.ObservedGeneration)
	out.Index = direct.ValueOf(in.Index)
	out.Retried = direct.ValueOf(in.Retried)
	out.LastAttemptResult = TaskAttemptResult_ToProto(mapCtx, in.LastAttemptResult)
	out.EncryptionKey = direct.ValueOf(in.EncryptionKey)
	out.VpcAccess = VpcAccess_ToProto(mapCtx, in.VpcAccess)
	out.LogUri = direct.ValueOf(in.LogURI)
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func VersionToPath_FromProto(mapCtx *direct.MapContext, in *pb.VersionToPath) *krm.VersionToPath {
	if in == nil {
		return nil
	}
	out := &krm.VersionToPath{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Mode = direct.LazyPtr(in.GetMode())
	return out
}
func VersionToPath_ToProto(mapCtx *direct.MapContext, in *krm.VersionToPath) *pb.VersionToPath {
	if in == nil {
		return nil
	}
	out := &pb.VersionToPath{}
	out.Path = direct.ValueOf(in.Path)
	out.Version = direct.ValueOf(in.Version)
	out.Mode = direct.ValueOf(in.Mode)
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Secret = SecretVolumeSource_FromProto(mapCtx, in.GetSecret())
	out.CloudSqlInstance = CloudSqlInstance_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.EmptyDir = EmptyDirVolumeSource_FromProto(mapCtx, in.GetEmptyDir())
	out.Nfs = NFSVolumeSource_FromProto(mapCtx, in.GetNfs())
	out.Gcs = GCSVolumeSource_FromProto(mapCtx, in.GetGcs())
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := SecretVolumeSource_ToProto(mapCtx, in.Secret); oneof != nil {
		out.VolumeType = &pb.Volume_Secret{Secret: oneof}
	}
	if oneof := CloudSqlInstance_ToProto(mapCtx, in.CloudSqlInstance); oneof != nil {
		out.VolumeType = &pb.Volume_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := EmptyDirVolumeSource_ToProto(mapCtx, in.EmptyDir); oneof != nil {
		out.VolumeType = &pb.Volume_EmptyDir{EmptyDir: oneof}
	}
	if oneof := NFSVolumeSource_ToProto(mapCtx, in.Nfs); oneof != nil {
		out.VolumeType = &pb.Volume_Nfs{Nfs: oneof}
	}
	if oneof := GCSVolumeSource_ToProto(mapCtx, in.Gcs); oneof != nil {
		out.VolumeType = &pb.Volume_Gcs{Gcs: oneof}
	}
	return out
}
func VolumeMount_FromProto(mapCtx *direct.MapContext, in *pb.VolumeMount) *krm.VolumeMount {
	if in == nil {
		return nil
	}
	out := &krm.VolumeMount{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	return out
}
func VolumeMount_ToProto(mapCtx *direct.MapContext, in *krm.VolumeMount) *pb.VolumeMount {
	if in == nil {
		return nil
	}
	out := &pb.VolumeMount{}
	out.Name = direct.ValueOf(in.Name)
	out.MountPath = direct.ValueOf(in.MountPath)
	return out
}
func VpcAccess_FromProto(mapCtx *direct.MapContext, in *pb.VpcAccess) *krm.VpcAccess {
	if in == nil {
		return nil
	}
	out := &krm.VpcAccess{}
	out.Connector = direct.LazyPtr(in.GetConnector())
	out.Egress = direct.Enum_FromProto(mapCtx, in.GetEgress())
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, VpcAccess_NetworkInterface_FromProto)
	return out
}
func VpcAccess_ToProto(mapCtx *direct.MapContext, in *krm.VpcAccess) *pb.VpcAccess {
	if in == nil {
		return nil
	}
	out := &pb.VpcAccess{}
	out.Connector = direct.ValueOf(in.Connector)
	out.Egress = direct.Enum_ToProto[pb.VpcAccess_VpcEgress](mapCtx, in.Egress)
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, VpcAccess_NetworkInterface_ToProto)
	return out
}
func VpcAccess_NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.VpcAccess_NetworkInterface) *krm.VpcAccess_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.VpcAccess_NetworkInterface{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.Subnetwork = direct.LazyPtr(in.GetSubnetwork())
	out.Tags = in.Tags
	return out
}
func VpcAccess_NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.VpcAccess_NetworkInterface) *pb.VpcAccess_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.VpcAccess_NetworkInterface{}
	out.Network = direct.ValueOf(in.Network)
	out.Subnetwork = direct.ValueOf(in.Subnetwork)
	out.Tags = in.Tags
	return out
}
