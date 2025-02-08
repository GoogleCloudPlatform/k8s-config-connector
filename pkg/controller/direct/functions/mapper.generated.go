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

package functions

import (
	pb "cloud.google.com/go/functions/apiv1/functionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/functions/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func CloudFunction_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction) *krm.CloudFunction {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunction{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.SourceArchiveURL = direct.LazyPtr(in.GetSourceArchiveUrl())
	out.SourceRepository = SourceRepository_FromProto(mapCtx, in.GetSourceRepository())
	out.SourceUploadURL = direct.LazyPtr(in.GetSourceUploadUrl())
	out.HTTPSTrigger = HttpsTrigger_FromProto(mapCtx, in.GetHttpsTrigger())
	out.EventTrigger = EventTrigger_FromProto(mapCtx, in.GetEventTrigger())
	// MISSING: Status
	out.EntryPoint = direct.LazyPtr(in.GetEntryPoint())
	out.Runtime = direct.LazyPtr(in.GetRuntime())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.AvailableMemoryMb = direct.LazyPtr(in.GetAvailableMemoryMb())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	// MISSING: UpdateTime
	// MISSING: VersionID
	out.Labels = in.Labels
	out.EnvironmentVariables = in.EnvironmentVariables
	out.BuildEnvironmentVariables = in.BuildEnvironmentVariables
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.MaxInstances = direct.LazyPtr(in.GetMaxInstances())
	out.MinInstances = direct.LazyPtr(in.GetMinInstances())
	out.VpcConnector = direct.LazyPtr(in.GetVpcConnector())
	out.VpcConnectorEgressSettings = direct.Enum_FromProto(mapCtx, in.GetVpcConnectorEgressSettings())
	out.IngressSettings = direct.Enum_FromProto(mapCtx, in.GetIngressSettings())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.BuildWorkerPool = direct.LazyPtr(in.GetBuildWorkerPool())
	// MISSING: BuildID
	// MISSING: BuildName
	out.SecretEnvironmentVariables = direct.Slice_FromProto(mapCtx, in.SecretEnvironmentVariables, SecretEnvVar_FromProto)
	out.SecretVolumes = direct.Slice_FromProto(mapCtx, in.SecretVolumes, SecretVolume_FromProto)
	out.SourceToken = direct.LazyPtr(in.GetSourceToken())
	out.DockerRepository = direct.LazyPtr(in.GetDockerRepository())
	out.DockerRegistry = direct.Enum_FromProto(mapCtx, in.GetDockerRegistry())
	out.AutomaticUpdatePolicy = CloudFunction_AutomaticUpdatePolicy_FromProto(mapCtx, in.GetAutomaticUpdatePolicy())
	out.OnDeployUpdatePolicy = CloudFunction_OnDeployUpdatePolicy_FromProto(mapCtx, in.GetOnDeployUpdatePolicy())
	out.BuildServiceAccount = direct.LazyPtr(in.GetBuildServiceAccount())
	return out
}
func CloudFunction_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunction) *pb.CloudFunction {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if oneof := CloudFunction_SourceArchiveUrl_ToProto(mapCtx, in.SourceArchiveURL); oneof != nil {
		out.SourceCode = oneof
	}
	if oneof := SourceRepository_ToProto(mapCtx, in.SourceRepository); oneof != nil {
		out.SourceCode = &pb.CloudFunction_SourceRepository{SourceRepository: oneof}
	}
	if oneof := CloudFunction_SourceUploadUrl_ToProto(mapCtx, in.SourceUploadURL); oneof != nil {
		out.SourceCode = oneof
	}
	if oneof := HttpsTrigger_ToProto(mapCtx, in.HTTPSTrigger); oneof != nil {
		out.Trigger = &pb.CloudFunction_HttpsTrigger{HttpsTrigger: oneof}
	}
	if oneof := EventTrigger_ToProto(mapCtx, in.EventTrigger); oneof != nil {
		out.Trigger = &pb.CloudFunction_EventTrigger{EventTrigger: oneof}
	}
	// MISSING: Status
	out.EntryPoint = direct.ValueOf(in.EntryPoint)
	out.Runtime = direct.ValueOf(in.Runtime)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.AvailableMemoryMb = direct.ValueOf(in.AvailableMemoryMb)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	// MISSING: UpdateTime
	// MISSING: VersionID
	out.Labels = in.Labels
	out.EnvironmentVariables = in.EnvironmentVariables
	out.BuildEnvironmentVariables = in.BuildEnvironmentVariables
	out.Network = direct.ValueOf(in.Network)
	out.MaxInstances = direct.ValueOf(in.MaxInstances)
	out.MinInstances = direct.ValueOf(in.MinInstances)
	out.VpcConnector = direct.ValueOf(in.VpcConnector)
	out.VpcConnectorEgressSettings = direct.Enum_ToProto[pb.CloudFunction_VpcConnectorEgressSettings](mapCtx, in.VpcConnectorEgressSettings)
	out.IngressSettings = direct.Enum_ToProto[pb.CloudFunction_IngressSettings](mapCtx, in.IngressSettings)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.BuildWorkerPool = direct.ValueOf(in.BuildWorkerPool)
	// MISSING: BuildID
	// MISSING: BuildName
	out.SecretEnvironmentVariables = direct.Slice_ToProto(mapCtx, in.SecretEnvironmentVariables, SecretEnvVar_ToProto)
	out.SecretVolumes = direct.Slice_ToProto(mapCtx, in.SecretVolumes, SecretVolume_ToProto)
	out.SourceToken = direct.ValueOf(in.SourceToken)
	out.DockerRepository = direct.ValueOf(in.DockerRepository)
	out.DockerRegistry = direct.Enum_ToProto[pb.CloudFunction_DockerRegistry](mapCtx, in.DockerRegistry)
	if oneof := CloudFunction_AutomaticUpdatePolicy_ToProto(mapCtx, in.AutomaticUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.CloudFunction_AutomaticUpdatePolicy_{AutomaticUpdatePolicy: oneof}
	}
	if oneof := CloudFunction_OnDeployUpdatePolicy_ToProto(mapCtx, in.OnDeployUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.CloudFunction_OnDeployUpdatePolicy_{OnDeployUpdatePolicy: oneof}
	}
	out.BuildServiceAccount = direct.ValueOf(in.BuildServiceAccount)
	return out
}
func CloudFunctionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction) *krm.CloudFunctionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunctionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	out.SourceRepository = SourceRepositoryObservedState_FromProto(mapCtx, in.GetSourceRepository())
	// MISSING: SourceUploadURL
	out.HTTPSTrigger = HttpsTriggerObservedState_FromProto(mapCtx, in.GetHttpsTrigger())
	// MISSING: EventTrigger
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	out.BuildID = direct.LazyPtr(in.GetBuildId())
	out.BuildName = direct.LazyPtr(in.GetBuildName())
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	out.OnDeployUpdatePolicy = CloudFunction_OnDeployUpdatePolicyObservedState_FromProto(mapCtx, in.GetOnDeployUpdatePolicy())
	// MISSING: BuildServiceAccount
	return out
}
func CloudFunctionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunctionObservedState) *pb.CloudFunction {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	if oneof := SourceRepositoryObservedState_ToProto(mapCtx, in.SourceRepository); oneof != nil {
		out.SourceCode = &pb.CloudFunction_SourceRepository{SourceRepository: oneof}
	}
	// MISSING: SourceUploadURL
	if oneof := HttpsTriggerObservedState_ToProto(mapCtx, in.HTTPSTrigger); oneof != nil {
		out.Trigger = &pb.CloudFunction_HttpsTrigger{HttpsTrigger: oneof}
	}
	// MISSING: EventTrigger
	out.Status = direct.Enum_ToProto[pb.CloudFunctionStatus](mapCtx, in.Status)
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.VersionId = direct.ValueOf(in.VersionID)
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	out.BuildId = direct.ValueOf(in.BuildID)
	out.BuildName = direct.ValueOf(in.BuildName)
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	if oneof := CloudFunction_OnDeployUpdatePolicyObservedState_ToProto(mapCtx, in.OnDeployUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.CloudFunction_OnDeployUpdatePolicy_{OnDeployUpdatePolicy: oneof}
	}
	// MISSING: BuildServiceAccount
	return out
}
func CloudFunction_AutomaticUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction_AutomaticUpdatePolicy) *krm.CloudFunction_AutomaticUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunction_AutomaticUpdatePolicy{}
	return out
}
func CloudFunction_AutomaticUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunction_AutomaticUpdatePolicy) *pb.CloudFunction_AutomaticUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction_AutomaticUpdatePolicy{}
	return out
}
func CloudFunction_OnDeployUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction_OnDeployUpdatePolicy) *krm.CloudFunction_OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunction_OnDeployUpdatePolicy{}
	// MISSING: RuntimeVersion
	return out
}
func CloudFunction_OnDeployUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunction_OnDeployUpdatePolicy) *pb.CloudFunction_OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction_OnDeployUpdatePolicy{}
	// MISSING: RuntimeVersion
	return out
}
func CloudFunction_OnDeployUpdatePolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction_OnDeployUpdatePolicy) *krm.CloudFunction_OnDeployUpdatePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunction_OnDeployUpdatePolicyObservedState{}
	out.RuntimeVersion = direct.LazyPtr(in.GetRuntimeVersion())
	return out
}
func CloudFunction_OnDeployUpdatePolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunction_OnDeployUpdatePolicyObservedState) *pb.CloudFunction_OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction_OnDeployUpdatePolicy{}
	out.RuntimeVersion = direct.ValueOf(in.RuntimeVersion)
	return out
}
func EventTrigger_FromProto(mapCtx *direct.MapContext, in *pb.EventTrigger) *krm.EventTrigger {
	if in == nil {
		return nil
	}
	out := &krm.EventTrigger{}
	out.EventType = direct.LazyPtr(in.GetEventType())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.Service = direct.LazyPtr(in.GetService())
	out.FailurePolicy = FailurePolicy_FromProto(mapCtx, in.GetFailurePolicy())
	return out
}
func EventTrigger_ToProto(mapCtx *direct.MapContext, in *krm.EventTrigger) *pb.EventTrigger {
	if in == nil {
		return nil
	}
	out := &pb.EventTrigger{}
	out.EventType = direct.ValueOf(in.EventType)
	out.Resource = direct.ValueOf(in.Resource)
	out.Service = direct.ValueOf(in.Service)
	out.FailurePolicy = FailurePolicy_ToProto(mapCtx, in.FailurePolicy)
	return out
}
func FailurePolicy_FromProto(mapCtx *direct.MapContext, in *pb.FailurePolicy) *krm.FailurePolicy {
	if in == nil {
		return nil
	}
	out := &krm.FailurePolicy{}
	out.Retry = FailurePolicy_Retry_FromProto(mapCtx, in.GetRetry())
	return out
}
func FailurePolicy_ToProto(mapCtx *direct.MapContext, in *krm.FailurePolicy) *pb.FailurePolicy {
	if in == nil {
		return nil
	}
	out := &pb.FailurePolicy{}
	if oneof := FailurePolicy_Retry_ToProto(mapCtx, in.Retry); oneof != nil {
		out.Action = &pb.FailurePolicy_Retry_{Retry: oneof}
	}
	return out
}
func FailurePolicy_Retry_FromProto(mapCtx *direct.MapContext, in *pb.FailurePolicy_Retry) *krm.FailurePolicy_Retry {
	if in == nil {
		return nil
	}
	out := &krm.FailurePolicy_Retry{}
	return out
}
func FailurePolicy_Retry_ToProto(mapCtx *direct.MapContext, in *krm.FailurePolicy_Retry) *pb.FailurePolicy_Retry {
	if in == nil {
		return nil
	}
	out := &pb.FailurePolicy_Retry{}
	return out
}
func FunctionsCloudFunctionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction) *krm.FunctionsCloudFunctionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FunctionsCloudFunctionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	// MISSING: SourceRepository
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// MISSING: EventTrigger
	// MISSING: Status
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}
func FunctionsCloudFunctionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FunctionsCloudFunctionObservedState) *pb.CloudFunction {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	// MISSING: SourceRepository
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// MISSING: EventTrigger
	// MISSING: Status
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}
func FunctionsCloudFunctionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction) *krm.FunctionsCloudFunctionSpec {
	if in == nil {
		return nil
	}
	out := &krm.FunctionsCloudFunctionSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	// MISSING: SourceRepository
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// MISSING: EventTrigger
	// MISSING: Status
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}
func FunctionsCloudFunctionSpec_ToProto(mapCtx *direct.MapContext, in *krm.FunctionsCloudFunctionSpec) *pb.CloudFunction {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: SourceArchiveURL
	// MISSING: SourceRepository
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// MISSING: EventTrigger
	// MISSING: Status
	// MISSING: EntryPoint
	// MISSING: Runtime
	// MISSING: Timeout
	// MISSING: AvailableMemoryMb
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	// MISSING: EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	// MISSING: MaxInstances
	// MISSING: MinInstances
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}
func HttpsTrigger_FromProto(mapCtx *direct.MapContext, in *pb.HttpsTrigger) *krm.HttpsTrigger {
	if in == nil {
		return nil
	}
	out := &krm.HttpsTrigger{}
	// MISSING: URL
	out.SecurityLevel = direct.Enum_FromProto(mapCtx, in.GetSecurityLevel())
	return out
}
func HttpsTrigger_ToProto(mapCtx *direct.MapContext, in *krm.HttpsTrigger) *pb.HttpsTrigger {
	if in == nil {
		return nil
	}
	out := &pb.HttpsTrigger{}
	// MISSING: URL
	out.SecurityLevel = direct.Enum_ToProto[pb.HttpsTrigger_SecurityLevel](mapCtx, in.SecurityLevel)
	return out
}
func HttpsTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HttpsTrigger) *krm.HttpsTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HttpsTriggerObservedState{}
	out.URL = direct.LazyPtr(in.GetUrl())
	// MISSING: SecurityLevel
	return out
}
func HttpsTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HttpsTriggerObservedState) *pb.HttpsTrigger {
	if in == nil {
		return nil
	}
	out := &pb.HttpsTrigger{}
	out.Url = direct.ValueOf(in.URL)
	// MISSING: SecurityLevel
	return out
}
func SecretEnvVar_FromProto(mapCtx *direct.MapContext, in *pb.SecretEnvVar) *krm.SecretEnvVar {
	if in == nil {
		return nil
	}
	out := &krm.SecretEnvVar{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func SecretEnvVar_ToProto(mapCtx *direct.MapContext, in *krm.SecretEnvVar) *pb.SecretEnvVar {
	if in == nil {
		return nil
	}
	out := &pb.SecretEnvVar{}
	out.Key = direct.ValueOf(in.Key)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Secret = direct.ValueOf(in.Secret)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func SecretVolume_FromProto(mapCtx *direct.MapContext, in *pb.SecretVolume) *krm.SecretVolume {
	if in == nil {
		return nil
	}
	out := &krm.SecretVolume{}
	out.MountPath = direct.LazyPtr(in.GetMountPath())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.Versions = direct.Slice_FromProto(mapCtx, in.Versions, SecretVolume_SecretVersion_FromProto)
	return out
}
func SecretVolume_ToProto(mapCtx *direct.MapContext, in *krm.SecretVolume) *pb.SecretVolume {
	if in == nil {
		return nil
	}
	out := &pb.SecretVolume{}
	out.MountPath = direct.ValueOf(in.MountPath)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Secret = direct.ValueOf(in.Secret)
	out.Versions = direct.Slice_ToProto(mapCtx, in.Versions, SecretVolume_SecretVersion_ToProto)
	return out
}
func SecretVolume_SecretVersion_FromProto(mapCtx *direct.MapContext, in *pb.SecretVolume_SecretVersion) *krm.SecretVolume_SecretVersion {
	if in == nil {
		return nil
	}
	out := &krm.SecretVolume_SecretVersion{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func SecretVolume_SecretVersion_ToProto(mapCtx *direct.MapContext, in *krm.SecretVolume_SecretVersion) *pb.SecretVolume_SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVolume_SecretVersion{}
	out.Version = direct.ValueOf(in.Version)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func SourceRepository_FromProto(mapCtx *direct.MapContext, in *pb.SourceRepository) *krm.SourceRepository {
	if in == nil {
		return nil
	}
	out := &krm.SourceRepository{}
	out.URL = direct.LazyPtr(in.GetUrl())
	// MISSING: DeployedURL
	return out
}
func SourceRepository_ToProto(mapCtx *direct.MapContext, in *krm.SourceRepository) *pb.SourceRepository {
	if in == nil {
		return nil
	}
	out := &pb.SourceRepository{}
	out.Url = direct.ValueOf(in.URL)
	// MISSING: DeployedURL
	return out
}
func SourceRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SourceRepository) *krm.SourceRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SourceRepositoryObservedState{}
	// MISSING: URL
	out.DeployedURL = direct.LazyPtr(in.GetDeployedUrl())
	return out
}
func SourceRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SourceRepositoryObservedState) *pb.SourceRepository {
	if in == nil {
		return nil
	}
	out := &pb.SourceRepository{}
	// MISSING: URL
	out.DeployedUrl = direct.ValueOf(in.DeployedURL)
	return out
}
