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

// +generated:mapper
// krm.group: run.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.run.v2

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	krmsecretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	krmvpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	apipb "google.golang.org/genproto/googleapis/api"
)

func BinaryAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.BinaryAuthorization) *krm.BinaryAuthorization {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorization{}
	out.UseDefault = direct.LazyPtr(in.GetUseDefault())
	// MISSING: Policy
	out.BreakglassJustification = direct.LazyPtr(in.GetBreakglassJustification())
	return out
}
func BinaryAuthorization_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorization) *pb.BinaryAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.BinaryAuthorization{}
	if oneof := BinaryAuthorization_UseDefault_ToProto(mapCtx, in.UseDefault); oneof != nil {
		out.BinauthzMethod = oneof
	}
	// MISSING: Policy
	out.BreakglassJustification = direct.ValueOf(in.BreakglassJustification)
	return out
}
func BinaryAuthorization_UseDefault_ToProto(mapCtx *direct.MapContext, in *bool) *pb.BinaryAuthorization_UseDefault {
	if in == nil {
		return nil
	}
	return &pb.BinaryAuthorization_UseDefault{UseDefault: *in}
}
func BuildInfo_FromProto(mapCtx *direct.MapContext, in *pb.BuildInfo) *krm.BuildInfo {
	if in == nil {
		return nil
	}
	out := &krm.BuildInfo{}
	// MISSING: FunctionTarget
	// MISSING: SourceLocation
	return out
}
func BuildInfo_ToProto(mapCtx *direct.MapContext, in *krm.BuildInfo) *pb.BuildInfo {
	if in == nil {
		return nil
	}
	out := &pb.BuildInfo{}
	// MISSING: FunctionTarget
	// MISSING: SourceLocation
	return out
}
func BuildInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildInfo) *krm.BuildInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildInfoObservedState{}
	out.FunctionTarget = direct.LazyPtr(in.GetFunctionTarget())
	out.SourceLocation = direct.LazyPtr(in.GetSourceLocation())
	return out
}
func BuildInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildInfoObservedState) *pb.BuildInfo {
	if in == nil {
		return nil
	}
	out := &pb.BuildInfo{}
	out.FunctionTarget = direct.ValueOf(in.FunctionTarget)
	out.SourceLocation = direct.ValueOf(in.SourceLocation)
	return out
}
func CloudSQLInstance_FromProto(mapCtx *direct.MapContext, in *pb.CloudSqlInstance) *krm.CloudSQLInstance {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLInstance{}

	if v := in.GetInstances(); len(v) != 0 {
		for i := range v {
			out.InstanceRefs = append(out.InstanceRefs, &refsv1beta1.SQLInstanceRef{External: v[i]})
		}
	}

	return out
}
func CloudSQLInstance_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLInstance) *pb.CloudSqlInstance {
	if in == nil {
		return nil
	}
	out := &pb.CloudSqlInstance{}

	if v := in.InstanceRefs; len(v) != 0 {
		for i := range v {
			out.Instances = append(out.Instances, v[i].External)
		}
	}

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
	out.Reason = direct.Enum_FromProto(mapCtx, in.GetReason())
	out.RevisionReason = direct.Enum_FromProto(mapCtx, in.GetRevisionReason())
	out.ExecutionReason = direct.Enum_FromProto(mapCtx, in.GetExecutionReason())
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
	if oneof := Condition_Reason_ToProto(mapCtx, in.Reason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := Condition_RevisionReason_ToProto(mapCtx, in.RevisionReason); oneof != nil {
		out.Reasons = oneof
	}
	if oneof := Condition_ExecutionReason_ToProto(mapCtx, in.ExecutionReason); oneof != nil {
		out.Reasons = oneof
	}
	return out
}
func Condition_Reason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_Reason {
	if in == nil {
		return nil
	}
	return &pb.Condition_Reason{Reason: direct.Enum_ToProto[pb.Condition_CommonReason](mapCtx, in)}
}
func Condition_RevisionReason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_RevisionReason_ {
	if in == nil {
		return nil
	}
	return &pb.Condition_RevisionReason_{RevisionReason: direct.Enum_ToProto[pb.Condition_RevisionReason](mapCtx, in)}
}
func Condition_ExecutionReason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_ExecutionReason_ {
	if in == nil {
		return nil
	}
	return &pb.Condition_ExecutionReason_{ExecutionReason: direct.Enum_ToProto[pb.Condition_ExecutionReason](mapCtx, in)}
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
	// MISSING: DependsOn
	// MISSING: BaseImageURI
	// MISSING: BuildInfo
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
	// MISSING: DependsOn
	// MISSING: BaseImageURI
	// MISSING: BuildInfo
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
func EnvVar_Value_ToProto(mapCtx *direct.MapContext, in *string) *pb.EnvVar_Value {
	if in == nil {
		return nil
	}
	return &pb.EnvVar_Value{Value: *in}
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
func ExecutionReference_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionReference) *krm.ExecutionReference {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionReference{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CompletionTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCompletionTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.CompletionStatus = direct.Enum_FromProto(mapCtx, in.GetCompletionStatus())
	return out
}
func ExecutionReference_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionReference) *pb.ExecutionReference {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionReference{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CompletionTime = direct.StringTimestamp_ToProto(mapCtx, in.CompletionTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.CompletionStatus = direct.Enum_ToProto[pb.ExecutionReference_CompletionStatus](mapCtx, in.CompletionStatus)
	return out
}
func ExecutionTemplate_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionTemplate) *krm.ExecutionTemplate {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionTemplate{}
	// MISSING: Labels
	out.Annotations = in.Annotations
	out.Parallelism = direct.LazyPtr(in.GetParallelism())
	out.TaskCount = direct.LazyPtr(in.GetTaskCount())
	out.Template = TaskTemplate_FromProto(mapCtx, in.GetTemplate())
	return out
}
func ExecutionTemplate_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionTemplate) *pb.ExecutionTemplate {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionTemplate{}
	// MISSING: Labels
	out.Annotations = in.Annotations
	out.Parallelism = direct.ValueOf(in.Parallelism)
	out.TaskCount = direct.ValueOf(in.TaskCount)
	out.Template = TaskTemplate_ToProto(mapCtx, in.Template)
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
func GrpcAction_FromProto(mapCtx *direct.MapContext, in *pb.GRPCAction) *krm.GrpcAction {
	if in == nil {
		return nil
	}
	out := &krm.GrpcAction{}
	out.Port = direct.LazyPtr(in.GetPort())
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func GrpcAction_ToProto(mapCtx *direct.MapContext, in *krm.GrpcAction) *pb.GRPCAction {
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
	// MISSING: HTTPHeaders
	// (near miss): "HTTPHeaders" vs "HttpHeaders"
	// MISSING: Port
	return out
}
func HTTPGetAction_ToProto(mapCtx *direct.MapContext, in *krm.HTTPGetAction) *pb.HTTPGetAction {
	if in == nil {
		return nil
	}
	out := &pb.HTTPGetAction{}
	out.Path = direct.ValueOf(in.Path)
	// MISSING: HTTPHeaders
	// (near miss): "HTTPHeaders" vs "HttpHeaders"
	// MISSING: Port
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
func NfsVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.NFSVolumeSource) *krm.NfsVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.NfsVolumeSource{}
	out.Server = direct.LazyPtr(in.GetServer())
	out.Path = direct.LazyPtr(in.GetPath())
	out.ReadOnly = direct.LazyPtr(in.GetReadOnly())
	return out
}
func NfsVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.NfsVolumeSource) *pb.NFSVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.NFSVolumeSource{}
	out.Server = direct.ValueOf(in.Server)
	out.Path = direct.ValueOf(in.Path)
	out.ReadOnly = direct.ValueOf(in.ReadOnly)
	return out
}
func NodeSelector_FromProto(mapCtx *direct.MapContext, in *pb.NodeSelector) *krm.NodeSelector {
	if in == nil {
		return nil
	}
	out := &krm.NodeSelector{}
	out.Accelerator = direct.LazyPtr(in.GetAccelerator())
	return out
}
func NodeSelector_ToProto(mapCtx *direct.MapContext, in *krm.NodeSelector) *pb.NodeSelector {
	if in == nil {
		return nil
	}
	out := &pb.NodeSelector{}
	out.Accelerator = direct.ValueOf(in.Accelerator)
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
	out.TCPSocket = TCPSocketAction_FromProto(mapCtx, in.GetTcpSocket())
	// MISSING: Grpc
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
	if oneof := TCPSocketAction_ToProto(mapCtx, in.TCPSocket); oneof != nil {
		out.ProbeType = &pb.Probe_TcpSocket{TcpSocket: oneof}
	}
	// MISSING: Grpc
	return out
}
func ResourceRequirements_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRequirements) *krm.ResourceRequirements {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRequirements{}
	out.Limits = in.Limits
	// MISSING: CPUIdle
	// MISSING: StartupCPUBoost
	return out
}
func ResourceRequirements_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRequirements) *pb.ResourceRequirements {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRequirements{}
	out.Limits = in.Limits
	// MISSING: CPUIdle
	// MISSING: StartupCPUBoost
	return out
}
func RunJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.RunJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RunJobObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Generation
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.LastModifier = direct.LazyPtr(in.GetLastModifier())
	// MISSING: ObservedGeneration
	if v := in.GetTerminalCondition(); v != nil {
		out.TerminalCondition = []*krm.Condition{Condition_FromProto(mapCtx, v)}
	}
	// MISSING: Conditions
	out.ExecutionCount = direct.LazyPtr(in.GetExecutionCount())
	if v := in.GetLatestCreatedExecution(); v != nil {
		out.LatestCreatedExecution = []*krm.ExecutionReference{ExecutionReference_FromProto(mapCtx, v)}
	}
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func RunJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RunJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Generation
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Creator = direct.ValueOf(in.Creator)
	out.LastModifier = direct.ValueOf(in.LastModifier)
	// MISSING: ObservedGeneration
	if len(in.TerminalCondition) > 0 && in.TerminalCondition[0] != nil {
		out.TerminalCondition = Condition_ToProto(mapCtx, in.TerminalCondition[0])
	}
	// MISSING: Conditions
	out.ExecutionCount = direct.ValueOf(in.ExecutionCount)
	if len(in.LatestCreatedExecution) > 0 && in.LatestCreatedExecution[0] != nil {
		out.LatestCreatedExecution = ExecutionReference_ToProto(mapCtx, in.LatestCreatedExecution[0])
	}
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func RunJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.RunJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.RunJobSpec{}
	// MISSING: Name
	// MISSING: Generation
	// MISSING: Labels
	out.Annotations = in.Annotations
	out.Client = direct.LazyPtr(in.GetClient())
	out.ClientVersion = direct.LazyPtr(in.GetClientVersion())
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.BinaryAuthorization = BinaryAuthorization_FromProto(mapCtx, in.GetBinaryAuthorization())
	out.Template = ExecutionTemplate_FromProto(mapCtx, in.GetTemplate())
	// MISSING: ObservedGeneration
	// MISSING: Conditions
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	return out
}
func RunJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.RunJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Generation
	// MISSING: Labels
	out.Annotations = in.Annotations
	out.Client = direct.ValueOf(in.Client)
	out.ClientVersion = direct.ValueOf(in.ClientVersion)
	out.LaunchStage = direct.Enum_ToProto[apipb.LaunchStage](mapCtx, in.LaunchStage)
	out.BinaryAuthorization = BinaryAuthorization_ToProto(mapCtx, in.BinaryAuthorization)
	out.Template = ExecutionTemplate_ToProto(mapCtx, in.Template)
	// MISSING: ObservedGeneration
	// MISSING: Conditions
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	return out
}
func SecretKeySelector_FromProto(mapCtx *direct.MapContext, in *pb.SecretKeySelector) *krm.SecretKeySelector {
	if in == nil {
		return nil
	}
	out := &krm.SecretKeySelector{}
	if in.GetSecret() != "" {
		out.SecretRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetSecret()}
	}
	if in.GetVersion() != "" {
		out.VersionRef = &krmsecretmanagerv1beta1.SecretVersionRef{External: in.GetVersion()}
	}
	return out
}
func SecretKeySelector_ToProto(mapCtx *direct.MapContext, in *krm.SecretKeySelector) *pb.SecretKeySelector {
	if in == nil {
		return nil
	}
	out := &pb.SecretKeySelector{}
	if in.SecretRef != nil {
		out.Secret = in.SecretRef.External
	}
	if in.VersionRef != nil {
		out.Version = in.VersionRef.External
	}
	return out
}
func SecretVolumeSource_FromProto(mapCtx *direct.MapContext, in *pb.SecretVolumeSource) *krm.SecretVolumeSource {
	if in == nil {
		return nil
	}
	out := &krm.SecretVolumeSource{}
	if in.GetSecret() != "" {
		out.SecretRef = &krmsecretmanagerv1beta1.SecretRef{External: in.GetSecret()}
	}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, VersionToPath_FromProto)
	out.DefaultMode = direct.LazyPtr(in.GetDefaultMode())
	return out
}
func SecretVolumeSource_ToProto(mapCtx *direct.MapContext, in *krm.SecretVolumeSource) *pb.SecretVolumeSource {
	if in == nil {
		return nil
	}
	out := &pb.SecretVolumeSource{}
	if in.SecretRef != nil {
		out.Secret = in.SecretRef.External
	}
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
func TaskTemplate_FromProto(mapCtx *direct.MapContext, in *pb.TaskTemplate) *krm.TaskTemplate {
	if in == nil {
		return nil
	}
	out := &krm.TaskTemplate{}
	out.Containers = direct.Slice_FromProto(mapCtx, in.Containers, Container_FromProto)
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	out.MaxRetries = direct.LazyPtr(in.GetMaxRetries())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ExecutionEnvironment = direct.Enum_FromProto(mapCtx, in.GetExecutionEnvironment())
	if in.GetEncryptionKey() != "" {
		out.EncryptionKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetEncryptionKey()}
	}
	out.VPCAccess = VPCAccess_FromProto(mapCtx, in.GetVpcAccess())
	// MISSING: NodeSelector
	// MISSING: GpuZonalRedundancyDisabled
	return out
}
func TaskTemplate_ToProto(mapCtx *direct.MapContext, in *krm.TaskTemplate) *pb.TaskTemplate {
	if in == nil {
		return nil
	}
	out := &pb.TaskTemplate{}
	out.Containers = direct.Slice_ToProto(mapCtx, in.Containers, Container_ToProto)
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	if oneof := TaskTemplate_MaxRetries_ToProto(mapCtx, in.MaxRetries); oneof != nil {
		out.Retries = oneof
	}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.ExecutionEnvironment = direct.Enum_ToProto[pb.ExecutionEnvironment](mapCtx, in.ExecutionEnvironment)
	if in.EncryptionKeyRef != nil {
		out.EncryptionKey = in.EncryptionKeyRef.External
	}
	out.VpcAccess = VPCAccess_ToProto(mapCtx, in.VPCAccess)
	// MISSING: NodeSelector
	// MISSING: GpuZonalRedundancyDisabled
	return out
}
func TaskTemplate_MaxRetries_ToProto(mapCtx *direct.MapContext, in *int32) *pb.TaskTemplate_MaxRetries {
	if in == nil {
		return nil
	}
	return &pb.TaskTemplate_MaxRetries{MaxRetries: *in}
}
func VPCAccess_FromProto(mapCtx *direct.MapContext, in *pb.VpcAccess) *krm.VPCAccess {
	if in == nil {
		return nil
	}
	out := &krm.VPCAccess{}
	if in.GetConnector() != "" {
		out.ConnectorRef = &krmvpcaccessv1beta1.VPCAccessConnectorRef{External: in.GetConnector()}
	}
	out.Egress = direct.Enum_FromProto(mapCtx, in.GetEgress())
	out.NetworkInterfaces = direct.Slice_FromProto(mapCtx, in.NetworkInterfaces, VPCAccess_NetworkInterface_FromProto)
	return out
}
func VPCAccess_ToProto(mapCtx *direct.MapContext, in *krm.VPCAccess) *pb.VpcAccess {
	if in == nil {
		return nil
	}
	out := &pb.VpcAccess{}
	if in.ConnectorRef != nil {
		out.Connector = in.ConnectorRef.External
	}
	out.Egress = direct.Enum_ToProto[pb.VpcAccess_VpcEgress](mapCtx, in.Egress)
	out.NetworkInterfaces = direct.Slice_ToProto(mapCtx, in.NetworkInterfaces, VPCAccess_NetworkInterface_ToProto)
	return out
}
func VPCAccess_NetworkInterface_FromProto(mapCtx *direct.MapContext, in *pb.VpcAccess_NetworkInterface) *krm.VPCAccess_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &krm.VPCAccess_NetworkInterface{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	out.Tags = in.Tags
	return out
}
func VPCAccess_NetworkInterface_ToProto(mapCtx *direct.MapContext, in *krm.VPCAccess_NetworkInterface) *pb.VpcAccess_NetworkInterface {
	if in == nil {
		return nil
	}
	out := &pb.VpcAccess_NetworkInterface{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.Tags = in.Tags
	return out
}
func VersionToPath_FromProto(mapCtx *direct.MapContext, in *pb.VersionToPath) *krm.VersionToPath {
	if in == nil {
		return nil
	}
	out := &krm.VersionToPath{}
	out.Path = direct.LazyPtr(in.GetPath())
	if in.GetVersion() != "" {
		out.VersionRef = &krmsecretmanagerv1beta1.SecretVersionRef{External: in.GetVersion()}
	}
	out.Mode = direct.LazyPtr(in.GetMode())
	return out
}
func VersionToPath_ToProto(mapCtx *direct.MapContext, in *krm.VersionToPath) *pb.VersionToPath {
	if in == nil {
		return nil
	}
	out := &pb.VersionToPath{}
	out.Path = direct.ValueOf(in.Path)
	if in.VersionRef != nil {
		out.Version = in.VersionRef.External
	}
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
	out.CloudSQLInstance = CloudSQLInstance_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.EmptyDir = EmptyDirVolumeSource_FromProto(mapCtx, in.GetEmptyDir())
	// MISSING: Nfs
	// MISSING: GCS
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
	if oneof := CloudSQLInstance_ToProto(mapCtx, in.CloudSQLInstance); oneof != nil {
		out.VolumeType = &pb.Volume_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := EmptyDirVolumeSource_ToProto(mapCtx, in.EmptyDir); oneof != nil {
		out.VolumeType = &pb.Volume_EmptyDir{EmptyDir: oneof}
	}
	// MISSING: Nfs
	// MISSING: GCS
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
