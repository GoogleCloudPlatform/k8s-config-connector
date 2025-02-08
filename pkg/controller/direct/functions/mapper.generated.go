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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/functions/apiv2beta/functionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/functions/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutomaticUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.AutomaticUpdatePolicy) *krm.AutomaticUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.AutomaticUpdatePolicy{}
	return out
}
func AutomaticUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.AutomaticUpdatePolicy) *pb.AutomaticUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutomaticUpdatePolicy{}
	return out
}
func BuildConfig_FromProto(mapCtx *direct.MapContext, in *pb.BuildConfig) *krm.BuildConfig {
	if in == nil {
		return nil
	}
	out := &krm.BuildConfig{}
	out.AutomaticUpdatePolicy = AutomaticUpdatePolicy_FromProto(mapCtx, in.GetAutomaticUpdatePolicy())
	out.OnDeployUpdatePolicy = OnDeployUpdatePolicy_FromProto(mapCtx, in.GetOnDeployUpdatePolicy())
	// MISSING: Build
	out.Runtime = direct.LazyPtr(in.GetRuntime())
	out.EntryPoint = direct.LazyPtr(in.GetEntryPoint())
	out.Source = Source_FromProto(mapCtx, in.GetSource())
	// MISSING: SourceProvenance
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.EnvironmentVariables = in.EnvironmentVariables
	out.DockerRegistry = direct.Enum_FromProto(mapCtx, in.GetDockerRegistry())
	out.DockerRepository = direct.LazyPtr(in.GetDockerRepository())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func BuildConfig_ToProto(mapCtx *direct.MapContext, in *krm.BuildConfig) *pb.BuildConfig {
	if in == nil {
		return nil
	}
	out := &pb.BuildConfig{}
	if oneof := AutomaticUpdatePolicy_ToProto(mapCtx, in.AutomaticUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.BuildConfig_AutomaticUpdatePolicy{AutomaticUpdatePolicy: oneof}
	}
	if oneof := OnDeployUpdatePolicy_ToProto(mapCtx, in.OnDeployUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.BuildConfig_OnDeployUpdatePolicy{OnDeployUpdatePolicy: oneof}
	}
	// MISSING: Build
	out.Runtime = direct.ValueOf(in.Runtime)
	out.EntryPoint = direct.ValueOf(in.EntryPoint)
	out.Source = Source_ToProto(mapCtx, in.Source)
	// MISSING: SourceProvenance
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.EnvironmentVariables = in.EnvironmentVariables
	out.DockerRegistry = direct.Enum_ToProto[pb.BuildConfig_DockerRegistry](mapCtx, in.DockerRegistry)
	out.DockerRepository = direct.ValueOf(in.DockerRepository)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	return out
}
func BuildConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildConfig) *krm.BuildConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildConfigObservedState{}
	// MISSING: AutomaticUpdatePolicy
	out.OnDeployUpdatePolicy = OnDeployUpdatePolicyObservedState_FromProto(mapCtx, in.GetOnDeployUpdatePolicy())
	out.Build = direct.LazyPtr(in.GetBuild())
	// MISSING: Runtime
	// MISSING: EntryPoint
	// MISSING: Source
	out.SourceProvenance = SourceProvenance_FromProto(mapCtx, in.GetSourceProvenance())
	// MISSING: WorkerPool
	// MISSING: EnvironmentVariables
	// MISSING: DockerRegistry
	// MISSING: DockerRepository
	// MISSING: ServiceAccount
	return out
}
func BuildConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildConfigObservedState) *pb.BuildConfig {
	if in == nil {
		return nil
	}
	out := &pb.BuildConfig{}
	// MISSING: AutomaticUpdatePolicy
	if oneof := OnDeployUpdatePolicyObservedState_ToProto(mapCtx, in.OnDeployUpdatePolicy); oneof != nil {
		out.RuntimeUpdatePolicy = &pb.BuildConfig_OnDeployUpdatePolicy{OnDeployUpdatePolicy: oneof}
	}
	out.Build = direct.ValueOf(in.Build)
	// MISSING: Runtime
	// MISSING: EntryPoint
	// MISSING: Source
	out.SourceProvenance = SourceProvenance_ToProto(mapCtx, in.SourceProvenance)
	// MISSING: WorkerPool
	// MISSING: EnvironmentVariables
	// MISSING: DockerRegistry
	// MISSING: DockerRepository
	// MISSING: ServiceAccount
	return out
}
func EventFilter_FromProto(mapCtx *direct.MapContext, in *pb.EventFilter) *krm.EventFilter {
	if in == nil {
		return nil
	}
	out := &krm.EventFilter{}
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Operator = direct.LazyPtr(in.GetOperator())
	return out
}
func EventFilter_ToProto(mapCtx *direct.MapContext, in *krm.EventFilter) *pb.EventFilter {
	if in == nil {
		return nil
	}
	out := &pb.EventFilter{}
	out.Attribute = direct.ValueOf(in.Attribute)
	out.Value = direct.ValueOf(in.Value)
	out.Operator = direct.ValueOf(in.Operator)
	return out
}
func EventTrigger_FromProto(mapCtx *direct.MapContext, in *pb.EventTrigger) *krm.EventTrigger {
	if in == nil {
		return nil
	}
	out := &krm.EventTrigger{}
	// MISSING: Trigger
	out.TriggerRegion = direct.LazyPtr(in.GetTriggerRegion())
	out.EventType = direct.LazyPtr(in.GetEventType())
	out.EventFilters = direct.Slice_FromProto(mapCtx, in.EventFilters, EventFilter_FromProto)
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.RetryPolicy = direct.Enum_FromProto(mapCtx, in.GetRetryPolicy())
	out.Channel = direct.LazyPtr(in.GetChannel())
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func EventTrigger_ToProto(mapCtx *direct.MapContext, in *krm.EventTrigger) *pb.EventTrigger {
	if in == nil {
		return nil
	}
	out := &pb.EventTrigger{}
	// MISSING: Trigger
	out.TriggerRegion = direct.ValueOf(in.TriggerRegion)
	out.EventType = direct.ValueOf(in.EventType)
	out.EventFilters = direct.Slice_ToProto(mapCtx, in.EventFilters, EventFilter_ToProto)
	out.PubsubTopic = direct.ValueOf(in.PubsubTopic)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.RetryPolicy = direct.Enum_ToProto[pb.EventTrigger_RetryPolicy](mapCtx, in.RetryPolicy)
	out.Channel = direct.ValueOf(in.Channel)
	out.Service = direct.ValueOf(in.Service)
	return out
}
func EventTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EventTrigger) *krm.EventTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventTriggerObservedState{}
	out.Trigger = direct.LazyPtr(in.GetTrigger())
	// MISSING: TriggerRegion
	// MISSING: EventType
	// MISSING: EventFilters
	// MISSING: PubsubTopic
	// MISSING: ServiceAccountEmail
	// MISSING: RetryPolicy
	// MISSING: Channel
	// MISSING: Service
	return out
}
func EventTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventTriggerObservedState) *pb.EventTrigger {
	if in == nil {
		return nil
	}
	out := &pb.EventTrigger{}
	out.Trigger = direct.ValueOf(in.Trigger)
	// MISSING: TriggerRegion
	// MISSING: EventType
	// MISSING: EventFilters
	// MISSING: PubsubTopic
	// MISSING: ServiceAccountEmail
	// MISSING: RetryPolicy
	// MISSING: Channel
	// MISSING: Service
	return out
}
func Function_FromProto(mapCtx *direct.MapContext, in *pb.Function) *krm.Function {
	if in == nil {
		return nil
	}
	out := &krm.Function{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BuildConfig = BuildConfig_FromProto(mapCtx, in.GetBuildConfig())
	out.ServiceConfig = ServiceConfig_FromProto(mapCtx, in.GetServiceConfig())
	out.EventTrigger = EventTrigger_FromProto(mapCtx, in.GetEventTrigger())
	// MISSING: State
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: StateMessages
	out.Environment = direct.Enum_FromProto(mapCtx, in.GetEnvironment())
	// MISSING: URL
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: SatisfiesPzs
	// MISSING: CreateTime
	return out
}
func Function_ToProto(mapCtx *direct.MapContext, in *krm.Function) *pb.Function {
	if in == nil {
		return nil
	}
	out := &pb.Function{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.BuildConfig = BuildConfig_ToProto(mapCtx, in.BuildConfig)
	out.ServiceConfig = ServiceConfig_ToProto(mapCtx, in.ServiceConfig)
	out.EventTrigger = EventTrigger_ToProto(mapCtx, in.EventTrigger)
	// MISSING: State
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: StateMessages
	out.Environment = direct.Enum_ToProto[pb.Environment](mapCtx, in.Environment)
	// MISSING: URL
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: SatisfiesPzs
	// MISSING: CreateTime
	return out
}
func FunctionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Function) *krm.FunctionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FunctionObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.BuildConfig = BuildConfigObservedState_FromProto(mapCtx, in.GetBuildConfig())
	out.ServiceConfig = ServiceConfigObservedState_FromProto(mapCtx, in.GetServiceConfig())
	out.EventTrigger = EventTriggerObservedState_FromProto(mapCtx, in.GetEventTrigger())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.StateMessages = direct.Slice_FromProto(mapCtx, in.StateMessages, StateMessage_FromProto)
	// MISSING: Environment
	out.URL = direct.LazyPtr(in.GetUrl())
	// MISSING: KMSKeyName
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func FunctionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FunctionObservedState) *pb.Function {
	if in == nil {
		return nil
	}
	out := &pb.Function{}
	// MISSING: Name
	// MISSING: Description
	out.BuildConfig = BuildConfigObservedState_ToProto(mapCtx, in.BuildConfig)
	out.ServiceConfig = ServiceConfigObservedState_ToProto(mapCtx, in.ServiceConfig)
	out.EventTrigger = EventTriggerObservedState_ToProto(mapCtx, in.EventTrigger)
	out.State = direct.Enum_ToProto[pb.Function_State](mapCtx, in.State)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.StateMessages = direct.Slice_ToProto(mapCtx, in.StateMessages, StateMessage_ToProto)
	// MISSING: Environment
	out.Url = direct.ValueOf(in.URL)
	// MISSING: KMSKeyName
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func OnDeployUpdatePolicy_FromProto(mapCtx *direct.MapContext, in *pb.OnDeployUpdatePolicy) *krm.OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &krm.OnDeployUpdatePolicy{}
	// MISSING: RuntimeVersion
	return out
}
func OnDeployUpdatePolicy_ToProto(mapCtx *direct.MapContext, in *krm.OnDeployUpdatePolicy) *pb.OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.OnDeployUpdatePolicy{}
	// MISSING: RuntimeVersion
	return out
}
func OnDeployUpdatePolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OnDeployUpdatePolicy) *krm.OnDeployUpdatePolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OnDeployUpdatePolicyObservedState{}
	out.RuntimeVersion = direct.LazyPtr(in.GetRuntimeVersion())
	return out
}
func OnDeployUpdatePolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OnDeployUpdatePolicyObservedState) *pb.OnDeployUpdatePolicy {
	if in == nil {
		return nil
	}
	out := &pb.OnDeployUpdatePolicy{}
	out.RuntimeVersion = direct.ValueOf(in.RuntimeVersion)
	return out
}
func RepoSource_FromProto(mapCtx *direct.MapContext, in *pb.RepoSource) *krm.RepoSource {
	if in == nil {
		return nil
	}
	out := &krm.RepoSource{}
	out.BranchName = direct.LazyPtr(in.GetBranchName())
	out.TagName = direct.LazyPtr(in.GetTagName())
	out.CommitSha = direct.LazyPtr(in.GetCommitSha())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.RepoName = direct.LazyPtr(in.GetRepoName())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	return out
}
func RepoSource_ToProto(mapCtx *direct.MapContext, in *krm.RepoSource) *pb.RepoSource {
	if in == nil {
		return nil
	}
	out := &pb.RepoSource{}
	if oneof := RepoSource_BranchName_ToProto(mapCtx, in.BranchName); oneof != nil {
		out.Revision = oneof
	}
	if oneof := RepoSource_TagName_ToProto(mapCtx, in.TagName); oneof != nil {
		out.Revision = oneof
	}
	if oneof := RepoSource_CommitSha_ToProto(mapCtx, in.CommitSha); oneof != nil {
		out.Revision = oneof
	}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.RepoName = direct.ValueOf(in.RepoName)
	out.Dir = direct.ValueOf(in.Dir)
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
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
func ServiceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConfig) *krm.ServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConfig{}
	// MISSING: Service
	out.TimeoutSeconds = direct.LazyPtr(in.GetTimeoutSeconds())
	out.AvailableMemory = direct.LazyPtr(in.GetAvailableMemory())
	out.AvailableCpu = direct.LazyPtr(in.GetAvailableCpu())
	out.EnvironmentVariables = in.EnvironmentVariables
	out.MaxInstanceCount = direct.LazyPtr(in.GetMaxInstanceCount())
	out.MinInstanceCount = direct.LazyPtr(in.GetMinInstanceCount())
	out.VpcConnector = direct.LazyPtr(in.GetVpcConnector())
	out.VpcConnectorEgressSettings = direct.Enum_FromProto(mapCtx, in.GetVpcConnectorEgressSettings())
	out.IngressSettings = direct.Enum_FromProto(mapCtx, in.GetIngressSettings())
	// MISSING: URI
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.AllTrafficOnLatestRevision = direct.LazyPtr(in.GetAllTrafficOnLatestRevision())
	out.SecretEnvironmentVariables = direct.Slice_FromProto(mapCtx, in.SecretEnvironmentVariables, SecretEnvVar_FromProto)
	out.SecretVolumes = direct.Slice_FromProto(mapCtx, in.SecretVolumes, SecretVolume_FromProto)
	// MISSING: Revision
	out.MaxInstanceRequestConcurrency = direct.LazyPtr(in.GetMaxInstanceRequestConcurrency())
	out.SecurityLevel = direct.Enum_FromProto(mapCtx, in.GetSecurityLevel())
	out.BinaryAuthorizationPolicy = direct.LazyPtr(in.GetBinaryAuthorizationPolicy())
	return out
}
func ServiceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConfig) *pb.ServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConfig{}
	// MISSING: Service
	out.TimeoutSeconds = direct.ValueOf(in.TimeoutSeconds)
	out.AvailableMemory = direct.ValueOf(in.AvailableMemory)
	out.AvailableCpu = direct.ValueOf(in.AvailableCpu)
	out.EnvironmentVariables = in.EnvironmentVariables
	out.MaxInstanceCount = direct.ValueOf(in.MaxInstanceCount)
	out.MinInstanceCount = direct.ValueOf(in.MinInstanceCount)
	out.VpcConnector = direct.ValueOf(in.VpcConnector)
	out.VpcConnectorEgressSettings = direct.Enum_ToProto[pb.ServiceConfig_VpcConnectorEgressSettings](mapCtx, in.VpcConnectorEgressSettings)
	out.IngressSettings = direct.Enum_ToProto[pb.ServiceConfig_IngressSettings](mapCtx, in.IngressSettings)
	// MISSING: URI
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.AllTrafficOnLatestRevision = direct.ValueOf(in.AllTrafficOnLatestRevision)
	out.SecretEnvironmentVariables = direct.Slice_ToProto(mapCtx, in.SecretEnvironmentVariables, SecretEnvVar_ToProto)
	out.SecretVolumes = direct.Slice_ToProto(mapCtx, in.SecretVolumes, SecretVolume_ToProto)
	// MISSING: Revision
	out.MaxInstanceRequestConcurrency = direct.ValueOf(in.MaxInstanceRequestConcurrency)
	out.SecurityLevel = direct.Enum_ToProto[pb.ServiceConfig_SecurityLevel](mapCtx, in.SecurityLevel)
	out.BinaryAuthorizationPolicy = direct.ValueOf(in.BinaryAuthorizationPolicy)
	return out
}
func ServiceConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConfig) *krm.ServiceConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConfigObservedState{}
	out.Service = direct.LazyPtr(in.GetService())
	// MISSING: TimeoutSeconds
	// MISSING: AvailableMemory
	// MISSING: AvailableCpu
	// MISSING: EnvironmentVariables
	// MISSING: MaxInstanceCount
	// MISSING: MinInstanceCount
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	out.URI = direct.LazyPtr(in.GetUri())
	// MISSING: ServiceAccountEmail
	// MISSING: AllTrafficOnLatestRevision
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	out.Revision = direct.LazyPtr(in.GetRevision())
	// MISSING: MaxInstanceRequestConcurrency
	// MISSING: SecurityLevel
	// MISSING: BinaryAuthorizationPolicy
	return out
}
func ServiceConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceConfigObservedState) *pb.ServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConfig{}
	out.Service = direct.ValueOf(in.Service)
	// MISSING: TimeoutSeconds
	// MISSING: AvailableMemory
	// MISSING: AvailableCpu
	// MISSING: EnvironmentVariables
	// MISSING: MaxInstanceCount
	// MISSING: MinInstanceCount
	// MISSING: VpcConnector
	// MISSING: VpcConnectorEgressSettings
	// MISSING: IngressSettings
	out.Uri = direct.ValueOf(in.URI)
	// MISSING: ServiceAccountEmail
	// MISSING: AllTrafficOnLatestRevision
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	out.Revision = direct.ValueOf(in.Revision)
	// MISSING: MaxInstanceRequestConcurrency
	// MISSING: SecurityLevel
	// MISSING: BinaryAuthorizationPolicy
	return out
}
func Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.Source {
	if in == nil {
		return nil
	}
	out := &krm.Source{}
	out.StorageSource = StorageSource_FromProto(mapCtx, in.GetStorageSource())
	out.RepoSource = RepoSource_FromProto(mapCtx, in.GetRepoSource())
	out.GitURI = direct.LazyPtr(in.GetGitUri())
	return out
}
func Source_ToProto(mapCtx *direct.MapContext, in *krm.Source) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	if oneof := StorageSource_ToProto(mapCtx, in.StorageSource); oneof != nil {
		out.Source = &pb.Source_StorageSource{StorageSource: oneof}
	}
	if oneof := RepoSource_ToProto(mapCtx, in.RepoSource); oneof != nil {
		out.Source = &pb.Source_RepoSource{RepoSource: oneof}
	}
	if oneof := Source_GitUri_ToProto(mapCtx, in.GitURI); oneof != nil {
		out.Source = oneof
	}
	return out
}
func SourceProvenance_FromProto(mapCtx *direct.MapContext, in *pb.SourceProvenance) *krm.SourceProvenance {
	if in == nil {
		return nil
	}
	out := &krm.SourceProvenance{}
	out.ResolvedStorageSource = StorageSource_FromProto(mapCtx, in.GetResolvedStorageSource())
	out.ResolvedRepoSource = RepoSource_FromProto(mapCtx, in.GetResolvedRepoSource())
	out.GitURI = direct.LazyPtr(in.GetGitUri())
	return out
}
func SourceProvenance_ToProto(mapCtx *direct.MapContext, in *krm.SourceProvenance) *pb.SourceProvenance {
	if in == nil {
		return nil
	}
	out := &pb.SourceProvenance{}
	out.ResolvedStorageSource = StorageSource_ToProto(mapCtx, in.ResolvedStorageSource)
	out.ResolvedRepoSource = RepoSource_ToProto(mapCtx, in.ResolvedRepoSource)
	out.GitUri = direct.ValueOf(in.GitURI)
	return out
}
func StateMessage_FromProto(mapCtx *direct.MapContext, in *pb.StateMessage) *krm.StateMessage {
	if in == nil {
		return nil
	}
	out := &krm.StateMessage{}
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Type = direct.LazyPtr(in.GetType())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}
func StateMessage_ToProto(mapCtx *direct.MapContext, in *krm.StateMessage) *pb.StateMessage {
	if in == nil {
		return nil
	}
	out := &pb.StateMessage{}
	out.Severity = direct.Enum_ToProto[pb.StateMessage_Severity](mapCtx, in.Severity)
	out.Type = direct.ValueOf(in.Type)
	out.Message = direct.ValueOf(in.Message)
	return out
}
func StorageSource_FromProto(mapCtx *direct.MapContext, in *pb.StorageSource) *krm.StorageSource {
	if in == nil {
		return nil
	}
	out := &krm.StorageSource{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.Object = direct.LazyPtr(in.GetObject())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	out.SourceUploadURL = direct.LazyPtr(in.GetSourceUploadUrl())
	return out
}
func StorageSource_ToProto(mapCtx *direct.MapContext, in *krm.StorageSource) *pb.StorageSource {
	if in == nil {
		return nil
	}
	out := &pb.StorageSource{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.Object = direct.ValueOf(in.Object)
	out.Generation = direct.ValueOf(in.Generation)
	out.SourceUploadUrl = direct.ValueOf(in.SourceUploadURL)
	return out
}
