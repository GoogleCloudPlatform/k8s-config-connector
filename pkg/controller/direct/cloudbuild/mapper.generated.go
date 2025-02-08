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

package cloudbuild

import (
	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ApprovalConfig_FromProto(mapCtx *direct.MapContext, in *pb.ApprovalConfig) *krm.ApprovalConfig {
	if in == nil {
		return nil
	}
	out := &krm.ApprovalConfig{}
	out.ApprovalRequired = direct.LazyPtr(in.GetApprovalRequired())
	return out
}
func ApprovalConfig_ToProto(mapCtx *direct.MapContext, in *krm.ApprovalConfig) *pb.ApprovalConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApprovalConfig{}
	out.ApprovalRequired = direct.ValueOf(in.ApprovalRequired)
	return out
}
func ApprovalResult_FromProto(mapCtx *direct.MapContext, in *pb.ApprovalResult) *krm.ApprovalResult {
	if in == nil {
		return nil
	}
	out := &krm.ApprovalResult{}
	// MISSING: ApproverAccount
	// MISSING: ApprovalTime
	out.Decision = direct.Enum_FromProto(mapCtx, in.GetDecision())
	out.Comment = direct.LazyPtr(in.GetComment())
	out.URL = direct.LazyPtr(in.GetUrl())
	return out
}
func ApprovalResult_ToProto(mapCtx *direct.MapContext, in *krm.ApprovalResult) *pb.ApprovalResult {
	if in == nil {
		return nil
	}
	out := &pb.ApprovalResult{}
	// MISSING: ApproverAccount
	// MISSING: ApprovalTime
	out.Decision = direct.Enum_ToProto[pb.ApprovalResult_Decision](mapCtx, in.Decision)
	out.Comment = direct.ValueOf(in.Comment)
	out.Url = direct.ValueOf(in.URL)
	return out
}
func ApprovalResultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApprovalResult) *krm.ApprovalResultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApprovalResultObservedState{}
	out.ApproverAccount = direct.LazyPtr(in.GetApproverAccount())
	out.ApprovalTime = direct.StringTimestamp_FromProto(mapCtx, in.GetApprovalTime())
	// MISSING: Decision
	// MISSING: Comment
	// MISSING: URL
	return out
}
func ApprovalResultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApprovalResultObservedState) *pb.ApprovalResult {
	if in == nil {
		return nil
	}
	out := &pb.ApprovalResult{}
	out.ApproverAccount = direct.ValueOf(in.ApproverAccount)
	out.ApprovalTime = direct.StringTimestamp_ToProto(mapCtx, in.ApprovalTime)
	// MISSING: Decision
	// MISSING: Comment
	// MISSING: URL
	return out
}
func Artifacts_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts) *krm.Artifacts {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts{}
	out.Images = in.Images
	out.Objects = Artifacts_ArtifactObjects_FromProto(mapCtx, in.GetObjects())
	out.MavenArtifacts = direct.Slice_FromProto(mapCtx, in.MavenArtifacts, Artifacts_MavenArtifact_FromProto)
	out.GoModules = direct.Slice_FromProto(mapCtx, in.GoModules, Artifacts_GoModule_FromProto)
	out.PythonPackages = direct.Slice_FromProto(mapCtx, in.PythonPackages, Artifacts_PythonPackage_FromProto)
	out.NpmPackages = direct.Slice_FromProto(mapCtx, in.NpmPackages, Artifacts_NpmPackage_FromProto)
	return out
}
func Artifacts_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts) *pb.Artifacts {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts{}
	out.Images = in.Images
	out.Objects = Artifacts_ArtifactObjects_ToProto(mapCtx, in.Objects)
	out.MavenArtifacts = direct.Slice_ToProto(mapCtx, in.MavenArtifacts, Artifacts_MavenArtifact_ToProto)
	out.GoModules = direct.Slice_ToProto(mapCtx, in.GoModules, Artifacts_GoModule_ToProto)
	out.PythonPackages = direct.Slice_ToProto(mapCtx, in.PythonPackages, Artifacts_PythonPackage_ToProto)
	out.NpmPackages = direct.Slice_ToProto(mapCtx, in.NpmPackages, Artifacts_NpmPackage_ToProto)
	return out
}
func ArtifactsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts) *krm.ArtifactsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactsObservedState{}
	// MISSING: Images
	out.Objects = Artifacts_ArtifactObjectsObservedState_FromProto(mapCtx, in.GetObjects())
	// MISSING: MavenArtifacts
	// MISSING: GoModules
	// MISSING: PythonPackages
	// MISSING: NpmPackages
	return out
}
func ArtifactsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactsObservedState) *pb.Artifacts {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts{}
	// MISSING: Images
	out.Objects = Artifacts_ArtifactObjectsObservedState_ToProto(mapCtx, in.Objects)
	// MISSING: MavenArtifacts
	// MISSING: GoModules
	// MISSING: PythonPackages
	// MISSING: NpmPackages
	return out
}
func Artifacts_ArtifactObjects_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_ArtifactObjects) *krm.Artifacts_ArtifactObjects {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_ArtifactObjects{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Paths = in.Paths
	// MISSING: Timing
	return out
}
func Artifacts_ArtifactObjects_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_ArtifactObjects) *pb.Artifacts_ArtifactObjects {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_ArtifactObjects{}
	out.Location = direct.ValueOf(in.Location)
	out.Paths = in.Paths
	// MISSING: Timing
	return out
}
func Artifacts_ArtifactObjectsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_ArtifactObjects) *krm.Artifacts_ArtifactObjectsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_ArtifactObjectsObservedState{}
	// MISSING: Location
	// MISSING: Paths
	out.Timing = TimeSpan_FromProto(mapCtx, in.GetTiming())
	return out
}
func Artifacts_ArtifactObjectsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_ArtifactObjectsObservedState) *pb.Artifacts_ArtifactObjects {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_ArtifactObjects{}
	// MISSING: Location
	// MISSING: Paths
	out.Timing = TimeSpan_ToProto(mapCtx, in.Timing)
	return out
}
func Artifacts_GoModule_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_GoModule) *krm.Artifacts_GoModule {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_GoModule{}
	out.RepositoryName = direct.LazyPtr(in.GetRepositoryName())
	out.RepositoryLocation = direct.LazyPtr(in.GetRepositoryLocation())
	out.RepositoryProjectID = direct.LazyPtr(in.GetRepositoryProjectId())
	out.SourcePath = direct.LazyPtr(in.GetSourcePath())
	out.ModulePath = direct.LazyPtr(in.GetModulePath())
	out.ModuleVersion = direct.LazyPtr(in.GetModuleVersion())
	return out
}
func Artifacts_GoModule_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_GoModule) *pb.Artifacts_GoModule {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_GoModule{}
	out.RepositoryName = direct.ValueOf(in.RepositoryName)
	out.RepositoryLocation = direct.ValueOf(in.RepositoryLocation)
	out.RepositoryProjectId = direct.ValueOf(in.RepositoryProjectID)
	out.SourcePath = direct.ValueOf(in.SourcePath)
	out.ModulePath = direct.ValueOf(in.ModulePath)
	out.ModuleVersion = direct.ValueOf(in.ModuleVersion)
	return out
}
func Artifacts_MavenArtifact_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_MavenArtifact) *krm.Artifacts_MavenArtifact {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_MavenArtifact{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Path = direct.LazyPtr(in.GetPath())
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.GroupID = direct.LazyPtr(in.GetGroupId())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Artifacts_MavenArtifact_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_MavenArtifact) *pb.Artifacts_MavenArtifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_MavenArtifact{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Path = direct.ValueOf(in.Path)
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	out.GroupId = direct.ValueOf(in.GroupID)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func Artifacts_NpmPackage_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_NpmPackage) *krm.Artifacts_NpmPackage {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_NpmPackage{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.PackagePath = direct.LazyPtr(in.GetPackagePath())
	return out
}
func Artifacts_NpmPackage_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_NpmPackage) *pb.Artifacts_NpmPackage {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_NpmPackage{}
	out.Repository = direct.ValueOf(in.Repository)
	out.PackagePath = direct.ValueOf(in.PackagePath)
	return out
}
func Artifacts_PythonPackage_FromProto(mapCtx *direct.MapContext, in *pb.Artifacts_PythonPackage) *krm.Artifacts_PythonPackage {
	if in == nil {
		return nil
	}
	out := &krm.Artifacts_PythonPackage{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Paths = in.Paths
	return out
}
func Artifacts_PythonPackage_ToProto(mapCtx *direct.MapContext, in *krm.Artifacts_PythonPackage) *pb.Artifacts_PythonPackage {
	if in == nil {
		return nil
	}
	out := &pb.Artifacts_PythonPackage{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Paths = in.Paths
	return out
}
func Build_FromProto(mapCtx *direct.MapContext, in *pb.Build) *krm.Build {
	if in == nil {
		return nil
	}
	out := &krm.Build{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: ProjectID
	// MISSING: Status
	// MISSING: StatusDetail
	out.Source = Source_FromProto(mapCtx, in.GetSource())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, BuildStep_FromProto)
	// MISSING: Results
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: FinishTime
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Images = in.Images
	out.QueueTtl = direct.StringDuration_FromProto(mapCtx, in.GetQueueTtl())
	out.Artifacts = Artifacts_FromProto(mapCtx, in.GetArtifacts())
	out.LogsBucket = direct.LazyPtr(in.GetLogsBucket())
	// MISSING: SourceProvenance
	// MISSING: BuildTriggerID
	out.Options = BuildOptions_FromProto(mapCtx, in.GetOptions())
	// MISSING: LogURL
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags
	out.Secrets = direct.Slice_FromProto(mapCtx, in.Secrets, Secret_FromProto)
	// MISSING: Timing
	// MISSING: Approval
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.AvailableSecrets = Secrets_FromProto(mapCtx, in.GetAvailableSecrets())
	// MISSING: Warnings
	// MISSING: FailureInfo
	return out
}
func Build_ToProto(mapCtx *direct.MapContext, in *krm.Build) *pb.Build {
	if in == nil {
		return nil
	}
	out := &pb.Build{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: ProjectID
	// MISSING: Status
	// MISSING: StatusDetail
	out.Source = Source_ToProto(mapCtx, in.Source)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, BuildStep_ToProto)
	// MISSING: Results
	// MISSING: CreateTime
	// MISSING: StartTime
	// MISSING: FinishTime
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Images = in.Images
	out.QueueTtl = direct.StringDuration_ToProto(mapCtx, in.QueueTtl)
	out.Artifacts = Artifacts_ToProto(mapCtx, in.Artifacts)
	out.LogsBucket = direct.ValueOf(in.LogsBucket)
	// MISSING: SourceProvenance
	// MISSING: BuildTriggerID
	out.Options = BuildOptions_ToProto(mapCtx, in.Options)
	// MISSING: LogURL
	out.Substitutions = in.Substitutions
	out.Tags = in.Tags
	out.Secrets = direct.Slice_ToProto(mapCtx, in.Secrets, Secret_ToProto)
	// MISSING: Timing
	// MISSING: Approval
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.AvailableSecrets = Secrets_ToProto(mapCtx, in.AvailableSecrets)
	// MISSING: Warnings
	// MISSING: FailureInfo
	return out
}
func BuildApproval_FromProto(mapCtx *direct.MapContext, in *pb.BuildApproval) *krm.BuildApproval {
	if in == nil {
		return nil
	}
	out := &krm.BuildApproval{}
	// MISSING: State
	// MISSING: Config
	// MISSING: Result
	return out
}
func BuildApproval_ToProto(mapCtx *direct.MapContext, in *krm.BuildApproval) *pb.BuildApproval {
	if in == nil {
		return nil
	}
	out := &pb.BuildApproval{}
	// MISSING: State
	// MISSING: Config
	// MISSING: Result
	return out
}
func BuildApprovalObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildApproval) *krm.BuildApprovalObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildApprovalObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Config = ApprovalConfig_FromProto(mapCtx, in.GetConfig())
	out.Result = ApprovalResult_FromProto(mapCtx, in.GetResult())
	return out
}
func BuildApprovalObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildApprovalObservedState) *pb.BuildApproval {
	if in == nil {
		return nil
	}
	out := &pb.BuildApproval{}
	out.State = direct.Enum_ToProto[pb.BuildApproval_State](mapCtx, in.State)
	out.Config = ApprovalConfig_ToProto(mapCtx, in.Config)
	out.Result = ApprovalResult_ToProto(mapCtx, in.Result)
	return out
}
func BuildObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Build) *krm.BuildObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	out.StatusDetail = direct.LazyPtr(in.GetStatusDetail())
	// MISSING: Source
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, BuildStepObservedState_FromProto)
	out.Results = Results_FromProto(mapCtx, in.GetResults())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.FinishTime = direct.StringTimestamp_FromProto(mapCtx, in.GetFinishTime())
	// MISSING: Timeout
	// MISSING: Images
	// MISSING: QueueTtl
	out.Artifacts = ArtifactsObservedState_FromProto(mapCtx, in.GetArtifacts())
	// MISSING: LogsBucket
	out.SourceProvenance = SourceProvenance_FromProto(mapCtx, in.GetSourceProvenance())
	out.BuildTriggerID = direct.LazyPtr(in.GetBuildTriggerId())
	// MISSING: Options
	out.LogURL = direct.LazyPtr(in.GetLogUrl())
	// MISSING: Substitutions
	// MISSING: Tags
	// MISSING: Secrets
	// MISSING: Timing
	out.Approval = BuildApproval_FromProto(mapCtx, in.GetApproval())
	// MISSING: ServiceAccount
	// MISSING: AvailableSecrets
	out.Warnings = direct.Slice_FromProto(mapCtx, in.Warnings, Build_Warning_FromProto)
	out.FailureInfo = Build_FailureInfo_FromProto(mapCtx, in.GetFailureInfo())
	return out
}
func BuildObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildObservedState) *pb.Build {
	if in == nil {
		return nil
	}
	out := &pb.Build{}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Status = direct.Enum_ToProto[pb.Build_Status](mapCtx, in.Status)
	out.StatusDetail = direct.ValueOf(in.StatusDetail)
	// MISSING: Source
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, BuildStepObservedState_ToProto)
	out.Results = Results_ToProto(mapCtx, in.Results)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.FinishTime = direct.StringTimestamp_ToProto(mapCtx, in.FinishTime)
	// MISSING: Timeout
	// MISSING: Images
	// MISSING: QueueTtl
	out.Artifacts = ArtifactsObservedState_ToProto(mapCtx, in.Artifacts)
	// MISSING: LogsBucket
	out.SourceProvenance = SourceProvenance_ToProto(mapCtx, in.SourceProvenance)
	out.BuildTriggerId = direct.ValueOf(in.BuildTriggerID)
	// MISSING: Options
	out.LogUrl = direct.ValueOf(in.LogURL)
	// MISSING: Substitutions
	// MISSING: Tags
	// MISSING: Secrets
	// MISSING: Timing
	out.Approval = BuildApproval_ToProto(mapCtx, in.Approval)
	// MISSING: ServiceAccount
	// MISSING: AvailableSecrets
	out.Warnings = direct.Slice_ToProto(mapCtx, in.Warnings, Build_Warning_ToProto)
	out.FailureInfo = Build_FailureInfo_ToProto(mapCtx, in.FailureInfo)
	return out
}
func BuildOptions_FromProto(mapCtx *direct.MapContext, in *pb.BuildOptions) *krm.BuildOptions {
	if in == nil {
		return nil
	}
	out := &krm.BuildOptions{}
	out.SourceProvenanceHash = direct.EnumSlice_FromProto(mapCtx, in.SourceProvenanceHash)
	out.RequestedVerifyOption = direct.Enum_FromProto(mapCtx, in.GetRequestedVerifyOption())
	out.MachineType = direct.Enum_FromProto(mapCtx, in.GetMachineType())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.SubstitutionOption = direct.Enum_FromProto(mapCtx, in.GetSubstitutionOption())
	out.DynamicSubstitutions = direct.LazyPtr(in.GetDynamicSubstitutions())
	out.AutomapSubstitutions = direct.LazyPtr(in.GetAutomapSubstitutions())
	out.LogStreamingOption = direct.Enum_FromProto(mapCtx, in.GetLogStreamingOption())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.Pool = BuildOptions_PoolOption_FromProto(mapCtx, in.GetPool())
	out.Logging = direct.Enum_FromProto(mapCtx, in.GetLogging())
	out.Env = in.Env
	out.SecretEnv = in.SecretEnv
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	out.DefaultLogsBucketBehavior = direct.Enum_FromProto(mapCtx, in.GetDefaultLogsBucketBehavior())
	out.EnableStructuredLogging = direct.LazyPtr(in.GetEnableStructuredLogging())
	return out
}
func BuildOptions_ToProto(mapCtx *direct.MapContext, in *krm.BuildOptions) *pb.BuildOptions {
	if in == nil {
		return nil
	}
	out := &pb.BuildOptions{}
	out.SourceProvenanceHash = direct.EnumSlice_ToProto[pb.Hash_HashType](mapCtx, in.SourceProvenanceHash)
	out.RequestedVerifyOption = direct.Enum_ToProto[pb.BuildOptions_VerifyOption](mapCtx, in.RequestedVerifyOption)
	out.MachineType = direct.Enum_ToProto[pb.BuildOptions_MachineType](mapCtx, in.MachineType)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.SubstitutionOption = direct.Enum_ToProto[pb.BuildOptions_SubstitutionOption](mapCtx, in.SubstitutionOption)
	out.DynamicSubstitutions = direct.ValueOf(in.DynamicSubstitutions)
	out.AutomapSubstitutions = direct.ValueOf(in.AutomapSubstitutions)
	out.LogStreamingOption = direct.Enum_ToProto[pb.BuildOptions_LogStreamingOption](mapCtx, in.LogStreamingOption)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.Pool = BuildOptions_PoolOption_ToProto(mapCtx, in.Pool)
	out.Logging = direct.Enum_ToProto[pb.BuildOptions_LoggingMode](mapCtx, in.Logging)
	out.Env = in.Env
	out.SecretEnv = in.SecretEnv
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	out.DefaultLogsBucketBehavior = direct.Enum_ToProto[pb.BuildOptions_DefaultLogsBucketBehavior](mapCtx, in.DefaultLogsBucketBehavior)
	out.EnableStructuredLogging = direct.ValueOf(in.EnableStructuredLogging)
	return out
}
func BuildOptions_PoolOption_FromProto(mapCtx *direct.MapContext, in *pb.BuildOptions_PoolOption) *krm.BuildOptions_PoolOption {
	if in == nil {
		return nil
	}
	out := &krm.BuildOptions_PoolOption{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func BuildOptions_PoolOption_ToProto(mapCtx *direct.MapContext, in *krm.BuildOptions_PoolOption) *pb.BuildOptions_PoolOption {
	if in == nil {
		return nil
	}
	out := &pb.BuildOptions_PoolOption{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func BuildStep_FromProto(mapCtx *direct.MapContext, in *pb.BuildStep) *krm.BuildStep {
	if in == nil {
		return nil
	}
	out := &krm.BuildStep{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Env = in.Env
	out.Args = in.Args
	out.Dir = direct.LazyPtr(in.GetDir())
	out.ID = direct.LazyPtr(in.GetId())
	out.WaitFor = in.WaitFor
	out.Entrypoint = direct.LazyPtr(in.GetEntrypoint())
	out.SecretEnv = in.SecretEnv
	out.Volumes = direct.Slice_FromProto(mapCtx, in.Volumes, Volume_FromProto)
	// MISSING: Timing
	// MISSING: PullTiming
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	// MISSING: Status
	out.AllowFailure = direct.LazyPtr(in.GetAllowFailure())
	// MISSING: ExitCode
	out.AllowExitCodes = in.AllowExitCodes
	out.Script = direct.LazyPtr(in.GetScript())
	out.AutomapSubstitutions = in.AutomapSubstitutions
	return out
}
func BuildStep_ToProto(mapCtx *direct.MapContext, in *krm.BuildStep) *pb.BuildStep {
	if in == nil {
		return nil
	}
	out := &pb.BuildStep{}
	out.Name = direct.ValueOf(in.Name)
	out.Env = in.Env
	out.Args = in.Args
	out.Dir = direct.ValueOf(in.Dir)
	out.Id = direct.ValueOf(in.ID)
	out.WaitFor = in.WaitFor
	out.Entrypoint = direct.ValueOf(in.Entrypoint)
	out.SecretEnv = in.SecretEnv
	out.Volumes = direct.Slice_ToProto(mapCtx, in.Volumes, Volume_ToProto)
	// MISSING: Timing
	// MISSING: PullTiming
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	// MISSING: Status
	out.AllowFailure = direct.ValueOf(in.AllowFailure)
	// MISSING: ExitCode
	out.AllowExitCodes = in.AllowExitCodes
	out.Script = direct.ValueOf(in.Script)
	out.AutomapSubstitutions = in.AutomapSubstitutions
	return out
}
func BuildStepObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildStep) *krm.BuildStepObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildStepObservedState{}
	// MISSING: Name
	// MISSING: Env
	// MISSING: Args
	// MISSING: Dir
	// MISSING: ID
	// MISSING: WaitFor
	// MISSING: Entrypoint
	// MISSING: SecretEnv
	// MISSING: Volumes
	out.Timing = TimeSpan_FromProto(mapCtx, in.GetTiming())
	out.PullTiming = TimeSpan_FromProto(mapCtx, in.GetPullTiming())
	// MISSING: Timeout
	out.Status = direct.Enum_FromProto(mapCtx, in.GetStatus())
	// MISSING: AllowFailure
	out.ExitCode = direct.LazyPtr(in.GetExitCode())
	// MISSING: AllowExitCodes
	// MISSING: Script
	// MISSING: AutomapSubstitutions
	return out
}
func BuildStepObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildStepObservedState) *pb.BuildStep {
	if in == nil {
		return nil
	}
	out := &pb.BuildStep{}
	// MISSING: Name
	// MISSING: Env
	// MISSING: Args
	// MISSING: Dir
	// MISSING: ID
	// MISSING: WaitFor
	// MISSING: Entrypoint
	// MISSING: SecretEnv
	// MISSING: Volumes
	out.Timing = TimeSpan_ToProto(mapCtx, in.Timing)
	out.PullTiming = TimeSpan_ToProto(mapCtx, in.PullTiming)
	// MISSING: Timeout
	out.Status = direct.Enum_ToProto[pb.Build_Status](mapCtx, in.Status)
	// MISSING: AllowFailure
	out.ExitCode = direct.ValueOf(in.ExitCode)
	// MISSING: AllowExitCodes
	// MISSING: Script
	// MISSING: AutomapSubstitutions
	return out
}
func BuildTrigger_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &krm.BuildTrigger{}
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	// MISSING: ID
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Name = direct.LazyPtr(in.GetName())
	out.Tags = in.Tags
	out.TriggerTemplate = RepoSource_FromProto(mapCtx, in.GetTriggerTemplate())
	out.Github = GitHubEventsConfig_FromProto(mapCtx, in.GetGithub())
	out.PubsubConfig = PubsubConfig_FromProto(mapCtx, in.GetPubsubConfig())
	out.WebhookConfig = WebhookConfig_FromProto(mapCtx, in.GetWebhookConfig())
	out.Autodetect = direct.LazyPtr(in.GetAutodetect())
	out.Build = Build_FromProto(mapCtx, in.GetBuild())
	out.Filename = direct.LazyPtr(in.GetFilename())
	out.GitFileSource = GitFileSource_FromProto(mapCtx, in.GetGitFileSource())
	// MISSING: CreateTime
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.Substitutions = in.Substitutions
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.SourceToBuild = GitRepoSource_FromProto(mapCtx, in.GetSourceToBuild())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.RepositoryEventConfig = RepositoryEventConfig_FromProto(mapCtx, in.GetRepositoryEventConfig())
	return out
}
func BuildTrigger_ToProto(mapCtx *direct.MapContext, in *krm.BuildTrigger) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	out.ResourceName = direct.ValueOf(in.ResourceName)
	// MISSING: ID
	out.Description = direct.ValueOf(in.Description)
	out.Name = direct.ValueOf(in.Name)
	out.Tags = in.Tags
	out.TriggerTemplate = RepoSource_ToProto(mapCtx, in.TriggerTemplate)
	out.Github = GitHubEventsConfig_ToProto(mapCtx, in.Github)
	out.PubsubConfig = PubsubConfig_ToProto(mapCtx, in.PubsubConfig)
	out.WebhookConfig = WebhookConfig_ToProto(mapCtx, in.WebhookConfig)
	if oneof := BuildTrigger_Autodetect_ToProto(mapCtx, in.Autodetect); oneof != nil {
		out.BuildTemplate = oneof
	}
	if oneof := Build_ToProto(mapCtx, in.Build); oneof != nil {
		out.BuildTemplate = &pb.BuildTrigger_Build{Build: oneof}
	}
	if oneof := BuildTrigger_Filename_ToProto(mapCtx, in.Filename); oneof != nil {
		out.BuildTemplate = oneof
	}
	if oneof := GitFileSource_ToProto(mapCtx, in.GitFileSource); oneof != nil {
		out.BuildTemplate = &pb.BuildTrigger_GitFileSource{GitFileSource: oneof}
	}
	// MISSING: CreateTime
	out.Disabled = direct.ValueOf(in.Disabled)
	out.Substitutions = in.Substitutions
	out.IgnoredFiles = in.IgnoredFiles
	out.IncludedFiles = in.IncludedFiles
	out.Filter = direct.ValueOf(in.Filter)
	out.SourceToBuild = GitRepoSource_ToProto(mapCtx, in.SourceToBuild)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.RepositoryEventConfig = RepositoryEventConfig_ToProto(mapCtx, in.RepositoryEventConfig)
	return out
}
func BuildTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.BuildTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuildTriggerObservedState{}
	// MISSING: ResourceName
	out.ID = direct.LazyPtr(in.GetId())
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	out.PubsubConfig = PubsubConfigObservedState_FromProto(mapCtx, in.GetPubsubConfig())
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	out.Build = BuildObservedState_FromProto(mapCtx, in.GetBuild())
	// MISSING: Filename
	// MISSING: GitFileSource
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	out.RepositoryEventConfig = RepositoryEventConfigObservedState_FromProto(mapCtx, in.GetRepositoryEventConfig())
	return out
}
func BuildTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuildTriggerObservedState) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	// MISSING: ResourceName
	out.Id = direct.ValueOf(in.ID)
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	out.PubsubConfig = PubsubConfigObservedState_ToProto(mapCtx, in.PubsubConfig)
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	if oneof := BuildObservedState_ToProto(mapCtx, in.Build); oneof != nil {
		out.BuildTemplate = &pb.BuildTrigger_Build{Build: oneof}
	}
	// MISSING: Filename
	// MISSING: GitFileSource
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	out.RepositoryEventConfig = RepositoryEventConfigObservedState_ToProto(mapCtx, in.RepositoryEventConfig)
	return out
}
func Build_FailureInfo_FromProto(mapCtx *direct.MapContext, in *pb.Build_FailureInfo) *krm.Build_FailureInfo {
	if in == nil {
		return nil
	}
	out := &krm.Build_FailureInfo{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Detail = direct.LazyPtr(in.GetDetail())
	return out
}
func Build_FailureInfo_ToProto(mapCtx *direct.MapContext, in *krm.Build_FailureInfo) *pb.Build_FailureInfo {
	if in == nil {
		return nil
	}
	out := &pb.Build_FailureInfo{}
	out.Type = direct.Enum_ToProto[pb.Build_FailureInfo_FailureType](mapCtx, in.Type)
	out.Detail = direct.ValueOf(in.Detail)
	return out
}
func Build_Warning_FromProto(mapCtx *direct.MapContext, in *pb.Build_Warning) *krm.Build_Warning {
	if in == nil {
		return nil
	}
	out := &krm.Build_Warning{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Priority = direct.Enum_FromProto(mapCtx, in.GetPriority())
	return out
}
func Build_Warning_ToProto(mapCtx *direct.MapContext, in *krm.Build_Warning) *pb.Build_Warning {
	if in == nil {
		return nil
	}
	out := &pb.Build_Warning{}
	out.Text = direct.ValueOf(in.Text)
	out.Priority = direct.Enum_ToProto[pb.Build_Warning_Priority](mapCtx, in.Priority)
	return out
}
func BuiltImage_FromProto(mapCtx *direct.MapContext, in *pb.BuiltImage) *krm.BuiltImage {
	if in == nil {
		return nil
	}
	out := &krm.BuiltImage{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Digest = direct.LazyPtr(in.GetDigest())
	// MISSING: PushTiming
	return out
}
func BuiltImage_ToProto(mapCtx *direct.MapContext, in *krm.BuiltImage) *pb.BuiltImage {
	if in == nil {
		return nil
	}
	out := &pb.BuiltImage{}
	out.Name = direct.ValueOf(in.Name)
	out.Digest = direct.ValueOf(in.Digest)
	// MISSING: PushTiming
	return out
}
func BuiltImageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuiltImage) *krm.BuiltImageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BuiltImageObservedState{}
	// MISSING: Name
	// MISSING: Digest
	out.PushTiming = TimeSpan_FromProto(mapCtx, in.GetPushTiming())
	return out
}
func BuiltImageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BuiltImageObservedState) *pb.BuiltImage {
	if in == nil {
		return nil
	}
	out := &pb.BuiltImage{}
	// MISSING: Name
	// MISSING: Digest
	out.PushTiming = TimeSpan_ToProto(mapCtx, in.PushTiming)
	return out
}
func CloudBuildWorkerPoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudBuildWorkerPoolObservedState) *pb.WorkerPool {
	if in == nil {
		return nil
	}
	out := &pb.WorkerPool{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Annotations
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DeleteTime
	// MISSING: State
	// MISSING: PrivatePoolV1Config
	// MISSING: Etag
	// (near miss): "Etag" vs "ETag"
	return out
}
func CloudbuildBuildTriggerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudbuildBuildTriggerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildBuildTriggerObservedState{}
	// MISSING: ResourceName
	// MISSING: ID
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	// MISSING: PubsubConfig
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	// MISSING: Build
	// MISSING: Filename
	// MISSING: GitFileSource
	// MISSING: CreateTime
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	// MISSING: RepositoryEventConfig
	return out
}
func CloudbuildBuildTriggerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildBuildTriggerObservedState) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	// MISSING: ResourceName
	// MISSING: ID
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	// MISSING: PubsubConfig
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	// MISSING: Build
	// MISSING: Filename
	// MISSING: GitFileSource
	// MISSING: CreateTime
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	// MISSING: RepositoryEventConfig
	return out
}
func CloudbuildBuildTriggerSpec_FromProto(mapCtx *direct.MapContext, in *pb.BuildTrigger) *krm.CloudbuildBuildTriggerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudbuildBuildTriggerSpec{}
	// MISSING: ResourceName
	// MISSING: ID
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	// MISSING: PubsubConfig
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	// MISSING: Build
	// MISSING: Filename
	// MISSING: GitFileSource
	// MISSING: CreateTime
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	// MISSING: RepositoryEventConfig
	return out
}
func CloudbuildBuildTriggerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudbuildBuildTriggerSpec) *pb.BuildTrigger {
	if in == nil {
		return nil
	}
	out := &pb.BuildTrigger{}
	// MISSING: ResourceName
	// MISSING: ID
	// MISSING: Description
	// MISSING: Name
	// MISSING: Tags
	// MISSING: TriggerTemplate
	// MISSING: Github
	// MISSING: PubsubConfig
	// MISSING: WebhookConfig
	// MISSING: Autodetect
	// MISSING: Build
	// MISSING: Filename
	// MISSING: GitFileSource
	// MISSING: CreateTime
	// MISSING: Disabled
	// MISSING: Substitutions
	// MISSING: IgnoredFiles
	// MISSING: IncludedFiles
	// MISSING: Filter
	// MISSING: SourceToBuild
	// MISSING: ServiceAccount
	// MISSING: RepositoryEventConfig
	return out
}
func FileHashes_FromProto(mapCtx *direct.MapContext, in *pb.FileHashes) *krm.FileHashes {
	if in == nil {
		return nil
	}
	out := &krm.FileHashes{}
	out.FileHash = direct.Slice_FromProto(mapCtx, in.FileHash, Hash_FromProto)
	return out
}
func FileHashes_ToProto(mapCtx *direct.MapContext, in *krm.FileHashes) *pb.FileHashes {
	if in == nil {
		return nil
	}
	out := &pb.FileHashes{}
	out.FileHash = direct.Slice_ToProto(mapCtx, in.FileHash, Hash_ToProto)
	return out
}
func GitFileSource_FromProto(mapCtx *direct.MapContext, in *pb.GitFileSource) *krm.GitFileSource {
	if in == nil {
		return nil
	}
	out := &krm.GitFileSource{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.RepoType = direct.Enum_FromProto(mapCtx, in.GetRepoType())
	out.Revision = direct.LazyPtr(in.GetRevision())
	out.GithubEnterpriseConfig = direct.LazyPtr(in.GetGithubEnterpriseConfig())
	return out
}
func GitFileSource_ToProto(mapCtx *direct.MapContext, in *krm.GitFileSource) *pb.GitFileSource {
	if in == nil {
		return nil
	}
	out := &pb.GitFileSource{}
	out.Path = direct.ValueOf(in.Path)
	out.Uri = direct.ValueOf(in.URI)
	if oneof := GitFileSource_Repository_ToProto(mapCtx, in.Repository); oneof != nil {
		out.Source = oneof
	}
	out.RepoType = direct.Enum_ToProto[pb.GitFileSource_RepoType](mapCtx, in.RepoType)
	out.Revision = direct.ValueOf(in.Revision)
	if oneof := GitFileSource_GithubEnterpriseConfig_ToProto(mapCtx, in.GithubEnterpriseConfig); oneof != nil {
		out.EnterpriseConfig = oneof
	}
	return out
}
func GitHubEventsConfig_FromProto(mapCtx *direct.MapContext, in *pb.GitHubEventsConfig) *krm.GitHubEventsConfig {
	if in == nil {
		return nil
	}
	out := &krm.GitHubEventsConfig{}
	out.InstallationID = direct.LazyPtr(in.GetInstallationId())
	out.Owner = direct.LazyPtr(in.GetOwner())
	out.Name = direct.LazyPtr(in.GetName())
	out.PullRequest = PullRequestFilter_FromProto(mapCtx, in.GetPullRequest())
	out.Push = PushFilter_FromProto(mapCtx, in.GetPush())
	return out
}
func GitHubEventsConfig_ToProto(mapCtx *direct.MapContext, in *krm.GitHubEventsConfig) *pb.GitHubEventsConfig {
	if in == nil {
		return nil
	}
	out := &pb.GitHubEventsConfig{}
	out.InstallationId = direct.ValueOf(in.InstallationID)
	out.Owner = direct.ValueOf(in.Owner)
	out.Name = direct.ValueOf(in.Name)
	if oneof := PullRequestFilter_ToProto(mapCtx, in.PullRequest); oneof != nil {
		out.Event = &pb.GitHubEventsConfig_PullRequest{PullRequest: oneof}
	}
	if oneof := PushFilter_ToProto(mapCtx, in.Push); oneof != nil {
		out.Event = &pb.GitHubEventsConfig_Push{Push: oneof}
	}
	return out
}
func GitRepoSource_FromProto(mapCtx *direct.MapContext, in *pb.GitRepoSource) *krm.GitRepoSource {
	if in == nil {
		return nil
	}
	out := &krm.GitRepoSource{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Ref = direct.LazyPtr(in.GetRef())
	out.RepoType = direct.Enum_FromProto(mapCtx, in.GetRepoType())
	out.GithubEnterpriseConfig = direct.LazyPtr(in.GetGithubEnterpriseConfig())
	return out
}
func GitRepoSource_ToProto(mapCtx *direct.MapContext, in *krm.GitRepoSource) *pb.GitRepoSource {
	if in == nil {
		return nil
	}
	out := &pb.GitRepoSource{}
	out.Uri = direct.ValueOf(in.URI)
	if oneof := GitRepoSource_Repository_ToProto(mapCtx, in.Repository); oneof != nil {
		out.Source = oneof
	}
	out.Ref = direct.ValueOf(in.Ref)
	out.RepoType = direct.Enum_ToProto[pb.GitFileSource_RepoType](mapCtx, in.RepoType)
	if oneof := GitRepoSource_GithubEnterpriseConfig_ToProto(mapCtx, in.GithubEnterpriseConfig); oneof != nil {
		out.EnterpriseConfig = oneof
	}
	return out
}
func GitSource_FromProto(mapCtx *direct.MapContext, in *pb.GitSource) *krm.GitSource {
	if in == nil {
		return nil
	}
	out := &krm.GitSource{}
	out.URL = direct.LazyPtr(in.GetUrl())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.Revision = direct.LazyPtr(in.GetRevision())
	return out
}
func GitSource_ToProto(mapCtx *direct.MapContext, in *krm.GitSource) *pb.GitSource {
	if in == nil {
		return nil
	}
	out := &pb.GitSource{}
	out.Url = direct.ValueOf(in.URL)
	out.Dir = direct.ValueOf(in.Dir)
	out.Revision = direct.ValueOf(in.Revision)
	return out
}
func Hash_FromProto(mapCtx *direct.MapContext, in *pb.Hash) *krm.Hash {
	if in == nil {
		return nil
	}
	out := &krm.Hash{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Value = in.GetValue()
	return out
}
func Hash_ToProto(mapCtx *direct.MapContext, in *krm.Hash) *pb.Hash {
	if in == nil {
		return nil
	}
	out := &pb.Hash{}
	out.Type = direct.Enum_ToProto[pb.Hash_HashType](mapCtx, in.Type)
	out.Value = in.Value
	return out
}
func InlineSecret_FromProto(mapCtx *direct.MapContext, in *pb.InlineSecret) *krm.InlineSecret {
	if in == nil {
		return nil
	}
	out := &krm.InlineSecret{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: EnvMap
	return out
}
func InlineSecret_ToProto(mapCtx *direct.MapContext, in *krm.InlineSecret) *pb.InlineSecret {
	if in == nil {
		return nil
	}
	out := &pb.InlineSecret{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: EnvMap
	return out
}
func PrivatePoolV1Config_NetworkConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePoolV1Config_NetworkConfig) *krm.PrivatePoolV1Config_NetworkConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePoolV1Config_NetworkConfigSpec{}
	if in.GetPeeredNetwork() != "" {
		out.PeeredNetworkRef = &refs.refv1beta1.ComputeNetworkRef{External: in.GetPeeredNetwork()}
	}
	out.EgressOption = direct.Enum_FromProto(mapCtx, in.GetEgressOption())
	out.PeeredNetworkIPRange = direct.LazyPtr(in.GetPeeredNetworkIpRange())
	return out
}
func PrivatePoolV1Config_NetworkConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePoolV1Config_NetworkConfigSpec) *pb.PrivatePoolV1Config_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_NetworkConfig{}
	if in.PeeredNetworkRef != nil {
		out.PeeredNetwork = in.PeeredNetworkRef.External
	}
	out.EgressOption = direct.Enum_ToProto[pb.PrivatePoolV1Config_NetworkConfig_EgressOption](mapCtx, in.EgressOption)
	out.PeeredNetworkIpRange = direct.ValueOf(in.PeeredNetworkIPRange)
	return out
}
func PrivatePoolV1Config_NetworkConfigStatus_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePoolV1Config_NetworkConfigStatus) *pb.PrivatePoolV1Config_NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePoolV1Config_NetworkConfig{}
	out.PeeredNetwork = direct.ValueOf(in.PeeredNetwork)
	out.EgressOption = direct.Enum_ToProto[pb.PrivatePoolV1Config_NetworkConfig_EgressOption](mapCtx, in.EgressOption)
	out.PeeredNetworkIpRange = direct.ValueOf(in.PeeredNetworkIPRange)
	return out
}
func PubsubConfig_FromProto(mapCtx *direct.MapContext, in *pb.PubsubConfig) *krm.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &krm.PubsubConfig{}
	// MISSING: Subscription
	out.Topic = direct.LazyPtr(in.GetTopic())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func PubsubConfig_ToProto(mapCtx *direct.MapContext, in *krm.PubsubConfig) *pb.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &pb.PubsubConfig{}
	// MISSING: Subscription
	out.Topic = direct.ValueOf(in.Topic)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.State = direct.Enum_ToProto[pb.PubsubConfig_State](mapCtx, in.State)
	return out
}
func PubsubConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PubsubConfig) *krm.PubsubConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PubsubConfigObservedState{}
	out.Subscription = direct.LazyPtr(in.GetSubscription())
	// MISSING: Topic
	// MISSING: ServiceAccountEmail
	// MISSING: State
	return out
}
func PubsubConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PubsubConfigObservedState) *pb.PubsubConfig {
	if in == nil {
		return nil
	}
	out := &pb.PubsubConfig{}
	out.Subscription = direct.ValueOf(in.Subscription)
	// MISSING: Topic
	// MISSING: ServiceAccountEmail
	// MISSING: State
	return out
}
func PullRequestFilter_FromProto(mapCtx *direct.MapContext, in *pb.PullRequestFilter) *krm.PullRequestFilter {
	if in == nil {
		return nil
	}
	out := &krm.PullRequestFilter{}
	out.Branch = direct.LazyPtr(in.GetBranch())
	out.CommentControl = direct.Enum_FromProto(mapCtx, in.GetCommentControl())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	return out
}
func PullRequestFilter_ToProto(mapCtx *direct.MapContext, in *krm.PullRequestFilter) *pb.PullRequestFilter {
	if in == nil {
		return nil
	}
	out := &pb.PullRequestFilter{}
	if oneof := PullRequestFilter_Branch_ToProto(mapCtx, in.Branch); oneof != nil {
		out.GitRef = oneof
	}
	out.CommentControl = direct.Enum_ToProto[pb.PullRequestFilter_CommentControl](mapCtx, in.CommentControl)
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
	return out
}
func PushFilter_FromProto(mapCtx *direct.MapContext, in *pb.PushFilter) *krm.PushFilter {
	if in == nil {
		return nil
	}
	out := &krm.PushFilter{}
	out.Branch = direct.LazyPtr(in.GetBranch())
	out.Tag = direct.LazyPtr(in.GetTag())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	return out
}
func PushFilter_ToProto(mapCtx *direct.MapContext, in *krm.PushFilter) *pb.PushFilter {
	if in == nil {
		return nil
	}
	out := &pb.PushFilter{}
	if oneof := PushFilter_Branch_ToProto(mapCtx, in.Branch); oneof != nil {
		out.GitRef = oneof
	}
	if oneof := PushFilter_Tag_ToProto(mapCtx, in.Tag); oneof != nil {
		out.GitRef = oneof
	}
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
	return out
}
func RepoSource_FromProto(mapCtx *direct.MapContext, in *pb.RepoSource) *krm.RepoSource {
	if in == nil {
		return nil
	}
	out := &krm.RepoSource{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.RepoName = direct.LazyPtr(in.GetRepoName())
	out.BranchName = direct.LazyPtr(in.GetBranchName())
	out.TagName = direct.LazyPtr(in.GetTagName())
	out.CommitSha = direct.LazyPtr(in.GetCommitSha())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.InvertRegex = direct.LazyPtr(in.GetInvertRegex())
	out.Substitutions = in.Substitutions
	return out
}
func RepoSource_ToProto(mapCtx *direct.MapContext, in *krm.RepoSource) *pb.RepoSource {
	if in == nil {
		return nil
	}
	out := &pb.RepoSource{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.RepoName = direct.ValueOf(in.RepoName)
	if oneof := RepoSource_BranchName_ToProto(mapCtx, in.BranchName); oneof != nil {
		out.Revision = oneof
	}
	if oneof := RepoSource_TagName_ToProto(mapCtx, in.TagName); oneof != nil {
		out.Revision = oneof
	}
	if oneof := RepoSource_CommitSha_ToProto(mapCtx, in.CommitSha); oneof != nil {
		out.Revision = oneof
	}
	out.Dir = direct.ValueOf(in.Dir)
	out.InvertRegex = direct.ValueOf(in.InvertRegex)
	out.Substitutions = in.Substitutions
	return out
}
func RepositoryEventConfig_FromProto(mapCtx *direct.MapContext, in *pb.RepositoryEventConfig) *krm.RepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryEventConfig{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	// MISSING: RepositoryType
	out.PullRequest = PullRequestFilter_FromProto(mapCtx, in.GetPullRequest())
	out.Push = PushFilter_FromProto(mapCtx, in.GetPush())
	return out
}
func RepositoryEventConfig_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryEventConfig) *pb.RepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepositoryEventConfig{}
	out.Repository = direct.ValueOf(in.Repository)
	// MISSING: RepositoryType
	if oneof := PullRequestFilter_ToProto(mapCtx, in.PullRequest); oneof != nil {
		out.Filter = &pb.RepositoryEventConfig_PullRequest{PullRequest: oneof}
	}
	if oneof := PushFilter_ToProto(mapCtx, in.Push); oneof != nil {
		out.Filter = &pb.RepositoryEventConfig_Push{Push: oneof}
	}
	return out
}
func RepositoryEventConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RepositoryEventConfig) *krm.RepositoryEventConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RepositoryEventConfigObservedState{}
	// MISSING: Repository
	out.RepositoryType = direct.Enum_FromProto(mapCtx, in.GetRepositoryType())
	// MISSING: PullRequest
	// MISSING: Push
	return out
}
func RepositoryEventConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RepositoryEventConfigObservedState) *pb.RepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := &pb.RepositoryEventConfig{}
	// MISSING: Repository
	out.RepositoryType = direct.Enum_ToProto[pb.RepositoryEventConfig_RepositoryType](mapCtx, in.RepositoryType)
	// MISSING: PullRequest
	// MISSING: Push
	return out
}
func Results_FromProto(mapCtx *direct.MapContext, in *pb.Results) *krm.Results {
	if in == nil {
		return nil
	}
	out := &krm.Results{}
	out.Images = direct.Slice_FromProto(mapCtx, in.Images, BuiltImage_FromProto)
	out.BuildStepImages = in.BuildStepImages
	out.ArtifactManifest = direct.LazyPtr(in.GetArtifactManifest())
	out.NumArtifacts = direct.LazyPtr(in.GetNumArtifacts())
	out.BuildStepOutputs = in.BuildStepOutputs
	out.ArtifactTiming = TimeSpan_FromProto(mapCtx, in.GetArtifactTiming())
	out.PythonPackages = direct.Slice_FromProto(mapCtx, in.PythonPackages, UploadedPythonPackage_FromProto)
	out.MavenArtifacts = direct.Slice_FromProto(mapCtx, in.MavenArtifacts, UploadedMavenArtifact_FromProto)
	out.GoModules = direct.Slice_FromProto(mapCtx, in.GoModules, UploadedGoModule_FromProto)
	out.NpmPackages = direct.Slice_FromProto(mapCtx, in.NpmPackages, UploadedNpmPackage_FromProto)
	return out
}
func Results_ToProto(mapCtx *direct.MapContext, in *krm.Results) *pb.Results {
	if in == nil {
		return nil
	}
	out := &pb.Results{}
	out.Images = direct.Slice_ToProto(mapCtx, in.Images, BuiltImage_ToProto)
	out.BuildStepImages = in.BuildStepImages
	out.ArtifactManifest = direct.ValueOf(in.ArtifactManifest)
	out.NumArtifacts = direct.ValueOf(in.NumArtifacts)
	out.BuildStepOutputs = in.BuildStepOutputs
	out.ArtifactTiming = TimeSpan_ToProto(mapCtx, in.ArtifactTiming)
	out.PythonPackages = direct.Slice_ToProto(mapCtx, in.PythonPackages, UploadedPythonPackage_ToProto)
	out.MavenArtifacts = direct.Slice_ToProto(mapCtx, in.MavenArtifacts, UploadedMavenArtifact_ToProto)
	out.GoModules = direct.Slice_ToProto(mapCtx, in.GoModules, UploadedGoModule_ToProto)
	out.NpmPackages = direct.Slice_ToProto(mapCtx, in.NpmPackages, UploadedNpmPackage_ToProto)
	return out
}
func ResultsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Results) *krm.ResultsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResultsObservedState{}
	out.Images = direct.Slice_FromProto(mapCtx, in.Images, BuiltImageObservedState_FromProto)
	// MISSING: BuildStepImages
	// MISSING: ArtifactManifest
	// MISSING: NumArtifacts
	// MISSING: BuildStepOutputs
	// MISSING: ArtifactTiming
	out.PythonPackages = direct.Slice_FromProto(mapCtx, in.PythonPackages, UploadedPythonPackageObservedState_FromProto)
	out.MavenArtifacts = direct.Slice_FromProto(mapCtx, in.MavenArtifacts, UploadedMavenArtifactObservedState_FromProto)
	out.GoModules = direct.Slice_FromProto(mapCtx, in.GoModules, UploadedGoModuleObservedState_FromProto)
	out.NpmPackages = direct.Slice_FromProto(mapCtx, in.NpmPackages, UploadedNpmPackageObservedState_FromProto)
	return out
}
func ResultsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResultsObservedState) *pb.Results {
	if in == nil {
		return nil
	}
	out := &pb.Results{}
	out.Images = direct.Slice_ToProto(mapCtx, in.Images, BuiltImageObservedState_ToProto)
	// MISSING: BuildStepImages
	// MISSING: ArtifactManifest
	// MISSING: NumArtifacts
	// MISSING: BuildStepOutputs
	// MISSING: ArtifactTiming
	out.PythonPackages = direct.Slice_ToProto(mapCtx, in.PythonPackages, UploadedPythonPackageObservedState_ToProto)
	out.MavenArtifacts = direct.Slice_ToProto(mapCtx, in.MavenArtifacts, UploadedMavenArtifactObservedState_ToProto)
	out.GoModules = direct.Slice_ToProto(mapCtx, in.GoModules, UploadedGoModuleObservedState_ToProto)
	out.NpmPackages = direct.Slice_ToProto(mapCtx, in.NpmPackages, UploadedNpmPackageObservedState_ToProto)
	return out
}
func Secret_FromProto(mapCtx *direct.MapContext, in *pb.Secret) *krm.Secret {
	if in == nil {
		return nil
	}
	out := &krm.Secret{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: SecretEnv
	return out
}
func Secret_ToProto(mapCtx *direct.MapContext, in *krm.Secret) *pb.Secret {
	if in == nil {
		return nil
	}
	out := &pb.Secret{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: SecretEnv
	return out
}
func SecretManagerSecret_FromProto(mapCtx *direct.MapContext, in *pb.SecretManagerSecret) *krm.SecretManagerSecret {
	if in == nil {
		return nil
	}
	out := &krm.SecretManagerSecret{}
	out.VersionName = direct.LazyPtr(in.GetVersionName())
	out.Env = direct.LazyPtr(in.GetEnv())
	return out
}
func SecretManagerSecret_ToProto(mapCtx *direct.MapContext, in *krm.SecretManagerSecret) *pb.SecretManagerSecret {
	if in == nil {
		return nil
	}
	out := &pb.SecretManagerSecret{}
	out.VersionName = direct.ValueOf(in.VersionName)
	out.Env = direct.ValueOf(in.Env)
	return out
}
func Secrets_FromProto(mapCtx *direct.MapContext, in *pb.Secrets) *krm.Secrets {
	if in == nil {
		return nil
	}
	out := &krm.Secrets{}
	out.SecretManager = direct.Slice_FromProto(mapCtx, in.SecretManager, SecretManagerSecret_FromProto)
	out.Inline = direct.Slice_FromProto(mapCtx, in.Inline, InlineSecret_FromProto)
	return out
}
func Secrets_ToProto(mapCtx *direct.MapContext, in *krm.Secrets) *pb.Secrets {
	if in == nil {
		return nil
	}
	out := &pb.Secrets{}
	out.SecretManager = direct.Slice_ToProto(mapCtx, in.SecretManager, SecretManagerSecret_ToProto)
	out.Inline = direct.Slice_ToProto(mapCtx, in.Inline, InlineSecret_ToProto)
	return out
}
func Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.Source {
	if in == nil {
		return nil
	}
	out := &krm.Source{}
	out.StorageSource = StorageSource_FromProto(mapCtx, in.GetStorageSource())
	out.RepoSource = RepoSource_FromProto(mapCtx, in.GetRepoSource())
	out.GitSource = GitSource_FromProto(mapCtx, in.GetGitSource())
	out.StorageSourceManifest = StorageSourceManifest_FromProto(mapCtx, in.GetStorageSourceManifest())
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
	if oneof := GitSource_ToProto(mapCtx, in.GitSource); oneof != nil {
		out.Source = &pb.Source_GitSource{GitSource: oneof}
	}
	if oneof := StorageSourceManifest_ToProto(mapCtx, in.StorageSourceManifest); oneof != nil {
		out.Source = &pb.Source_StorageSourceManifest{StorageSourceManifest: oneof}
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
	out.ResolvedStorageSourceManifest = StorageSourceManifest_FromProto(mapCtx, in.GetResolvedStorageSourceManifest())
	// MISSING: FileHashes
	return out
}
func SourceProvenance_ToProto(mapCtx *direct.MapContext, in *krm.SourceProvenance) *pb.SourceProvenance {
	if in == nil {
		return nil
	}
	out := &pb.SourceProvenance{}
	out.ResolvedStorageSource = StorageSource_ToProto(mapCtx, in.ResolvedStorageSource)
	out.ResolvedRepoSource = RepoSource_ToProto(mapCtx, in.ResolvedRepoSource)
	out.ResolvedStorageSourceManifest = StorageSourceManifest_ToProto(mapCtx, in.ResolvedStorageSourceManifest)
	// MISSING: FileHashes
	return out
}
func SourceProvenanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SourceProvenance) *krm.SourceProvenanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SourceProvenanceObservedState{}
	// MISSING: ResolvedStorageSource
	// MISSING: ResolvedRepoSource
	// MISSING: ResolvedStorageSourceManifest
	// MISSING: FileHashes
	return out
}
func SourceProvenanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SourceProvenanceObservedState) *pb.SourceProvenance {
	if in == nil {
		return nil
	}
	out := &pb.SourceProvenance{}
	// MISSING: ResolvedStorageSource
	// MISSING: ResolvedRepoSource
	// MISSING: ResolvedStorageSourceManifest
	// MISSING: FileHashes
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
	out.SourceFetcher = direct.Enum_FromProto(mapCtx, in.GetSourceFetcher())
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
	out.SourceFetcher = direct.Enum_ToProto[pb.StorageSource_SourceFetcher](mapCtx, in.SourceFetcher)
	return out
}
func StorageSourceManifest_FromProto(mapCtx *direct.MapContext, in *pb.StorageSourceManifest) *krm.StorageSourceManifest {
	if in == nil {
		return nil
	}
	out := &krm.StorageSourceManifest{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.Object = direct.LazyPtr(in.GetObject())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	return out
}
func StorageSourceManifest_ToProto(mapCtx *direct.MapContext, in *krm.StorageSourceManifest) *pb.StorageSourceManifest {
	if in == nil {
		return nil
	}
	out := &pb.StorageSourceManifest{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.Object = direct.ValueOf(in.Object)
	out.Generation = direct.ValueOf(in.Generation)
	return out
}
func TimeSpan_FromProto(mapCtx *direct.MapContext, in *pb.TimeSpan) *krm.TimeSpan {
	if in == nil {
		return nil
	}
	out := &krm.TimeSpan{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeSpan_ToProto(mapCtx *direct.MapContext, in *krm.TimeSpan) *pb.TimeSpan {
	if in == nil {
		return nil
	}
	out := &pb.TimeSpan{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func UploadedGoModule_FromProto(mapCtx *direct.MapContext, in *pb.UploadedGoModule) *krm.UploadedGoModule {
	if in == nil {
		return nil
	}
	out := &krm.UploadedGoModule{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.FileHashes = FileHashes_FromProto(mapCtx, in.GetFileHashes())
	// MISSING: PushTiming
	return out
}
func UploadedGoModule_ToProto(mapCtx *direct.MapContext, in *krm.UploadedGoModule) *pb.UploadedGoModule {
	if in == nil {
		return nil
	}
	out := &pb.UploadedGoModule{}
	out.Uri = direct.ValueOf(in.URI)
	out.FileHashes = FileHashes_ToProto(mapCtx, in.FileHashes)
	// MISSING: PushTiming
	return out
}
func UploadedGoModuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UploadedGoModule) *krm.UploadedGoModuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UploadedGoModuleObservedState{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_FromProto(mapCtx, in.GetPushTiming())
	return out
}
func UploadedGoModuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UploadedGoModuleObservedState) *pb.UploadedGoModule {
	if in == nil {
		return nil
	}
	out := &pb.UploadedGoModule{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_ToProto(mapCtx, in.PushTiming)
	return out
}
func UploadedMavenArtifact_FromProto(mapCtx *direct.MapContext, in *pb.UploadedMavenArtifact) *krm.UploadedMavenArtifact {
	if in == nil {
		return nil
	}
	out := &krm.UploadedMavenArtifact{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.FileHashes = FileHashes_FromProto(mapCtx, in.GetFileHashes())
	// MISSING: PushTiming
	return out
}
func UploadedMavenArtifact_ToProto(mapCtx *direct.MapContext, in *krm.UploadedMavenArtifact) *pb.UploadedMavenArtifact {
	if in == nil {
		return nil
	}
	out := &pb.UploadedMavenArtifact{}
	out.Uri = direct.ValueOf(in.URI)
	out.FileHashes = FileHashes_ToProto(mapCtx, in.FileHashes)
	// MISSING: PushTiming
	return out
}
func UploadedMavenArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UploadedMavenArtifact) *krm.UploadedMavenArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UploadedMavenArtifactObservedState{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_FromProto(mapCtx, in.GetPushTiming())
	return out
}
func UploadedMavenArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UploadedMavenArtifactObservedState) *pb.UploadedMavenArtifact {
	if in == nil {
		return nil
	}
	out := &pb.UploadedMavenArtifact{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_ToProto(mapCtx, in.PushTiming)
	return out
}
func UploadedNpmPackage_FromProto(mapCtx *direct.MapContext, in *pb.UploadedNpmPackage) *krm.UploadedNpmPackage {
	if in == nil {
		return nil
	}
	out := &krm.UploadedNpmPackage{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.FileHashes = FileHashes_FromProto(mapCtx, in.GetFileHashes())
	// MISSING: PushTiming
	return out
}
func UploadedNpmPackage_ToProto(mapCtx *direct.MapContext, in *krm.UploadedNpmPackage) *pb.UploadedNpmPackage {
	if in == nil {
		return nil
	}
	out := &pb.UploadedNpmPackage{}
	out.Uri = direct.ValueOf(in.URI)
	out.FileHashes = FileHashes_ToProto(mapCtx, in.FileHashes)
	// MISSING: PushTiming
	return out
}
func UploadedNpmPackageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UploadedNpmPackage) *krm.UploadedNpmPackageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UploadedNpmPackageObservedState{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_FromProto(mapCtx, in.GetPushTiming())
	return out
}
func UploadedNpmPackageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UploadedNpmPackageObservedState) *pb.UploadedNpmPackage {
	if in == nil {
		return nil
	}
	out := &pb.UploadedNpmPackage{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_ToProto(mapCtx, in.PushTiming)
	return out
}
func UploadedPythonPackage_FromProto(mapCtx *direct.MapContext, in *pb.UploadedPythonPackage) *krm.UploadedPythonPackage {
	if in == nil {
		return nil
	}
	out := &krm.UploadedPythonPackage{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.FileHashes = FileHashes_FromProto(mapCtx, in.GetFileHashes())
	// MISSING: PushTiming
	return out
}
func UploadedPythonPackage_ToProto(mapCtx *direct.MapContext, in *krm.UploadedPythonPackage) *pb.UploadedPythonPackage {
	if in == nil {
		return nil
	}
	out := &pb.UploadedPythonPackage{}
	out.Uri = direct.ValueOf(in.URI)
	out.FileHashes = FileHashes_ToProto(mapCtx, in.FileHashes)
	// MISSING: PushTiming
	return out
}
func UploadedPythonPackageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UploadedPythonPackage) *krm.UploadedPythonPackageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UploadedPythonPackageObservedState{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_FromProto(mapCtx, in.GetPushTiming())
	return out
}
func UploadedPythonPackageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UploadedPythonPackageObservedState) *pb.UploadedPythonPackage {
	if in == nil {
		return nil
	}
	out := &pb.UploadedPythonPackage{}
	// MISSING: URI
	// MISSING: FileHashes
	out.PushTiming = TimeSpan_ToProto(mapCtx, in.PushTiming)
	return out
}
func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func WebhookConfig_FromProto(mapCtx *direct.MapContext, in *pb.WebhookConfig) *krm.WebhookConfig {
	if in == nil {
		return nil
	}
	out := &krm.WebhookConfig{}
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func WebhookConfig_ToProto(mapCtx *direct.MapContext, in *krm.WebhookConfig) *pb.WebhookConfig {
	if in == nil {
		return nil
	}
	out := &pb.WebhookConfig{}
	if oneof := WebhookConfig_Secret_ToProto(mapCtx, in.Secret); oneof != nil {
		out.AuthMethod = oneof
	}
	out.State = direct.Enum_ToProto[pb.WebhookConfig_State](mapCtx, in.State)
	return out
}
