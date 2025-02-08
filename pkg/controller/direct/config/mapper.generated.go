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

package config

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/config/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApplyResults_FromProto(mapCtx *direct.MapContext, in *pb.ApplyResults) *krm.ApplyResults {
	if in == nil {
		return nil
	}
	out := &krm.ApplyResults{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.Artifacts = direct.LazyPtr(in.GetArtifacts())
	// MISSING: Outputs
	return out
}
func ApplyResults_ToProto(mapCtx *direct.MapContext, in *krm.ApplyResults) *pb.ApplyResults {
	if in == nil {
		return nil
	}
	out := &pb.ApplyResults{}
	out.Content = direct.ValueOf(in.Content)
	out.Artifacts = direct.ValueOf(in.Artifacts)
	// MISSING: Outputs
	return out
}
func ConfigDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ConfigDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentObservedState{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: LockState
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func ConfigDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: LockState
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func ConfigDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ConfigDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentSpec{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: LockState
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func ConfigDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: LockState
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func Deployment_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.Deployment {
	if in == nil {
		return nil
	}
	out := &krm.Deployment{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	out.ArtifactsGcsBucket = in.ArtifactsGcsBucket
	out.ServiceAccount = in.ServiceAccount
	out.ImportExistingResources = in.ImportExistingResources
	out.WorkerPool = in.WorkerPool
	// MISSING: LockState
	out.TfVersionConstraint = in.TfVersionConstraint
	// MISSING: TfVersion
	out.QuotaValidation = direct.Enum_FromProto(mapCtx, in.GetQuotaValidation())
	out.Annotations = in.Annotations
	return out
}
func Deployment_ToProto(mapCtx *direct.MapContext, in *krm.Deployment) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Deployment_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: State
	// MISSING: LatestRevision
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: DeleteResults
	// MISSING: DeleteBuild
	// MISSING: DeleteLogs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	out.ArtifactsGcsBucket = in.ArtifactsGcsBucket
	out.ServiceAccount = in.ServiceAccount
	out.ImportExistingResources = in.ImportExistingResources
	out.WorkerPool = in.WorkerPool
	// MISSING: LockState
	out.TfVersionConstraint = in.TfVersionConstraint
	// MISSING: TfVersion
	out.QuotaValidation = direct.Enum_ToProto[pb.QuotaValidation](mapCtx, in.QuotaValidation)
	out.Annotations = in.Annotations
	return out
}
func DeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.DeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentObservedState{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.LatestRevision = direct.LazyPtr(in.GetLatestRevision())
	out.StateDetail = direct.LazyPtr(in.GetStateDetail())
	out.ErrorCode = direct.Enum_FromProto(mapCtx, in.GetErrorCode())
	out.DeleteResults = ApplyResults_FromProto(mapCtx, in.GetDeleteResults())
	out.DeleteBuild = direct.LazyPtr(in.GetDeleteBuild())
	out.DeleteLogs = direct.LazyPtr(in.GetDeleteLogs())
	out.TfErrors = direct.Slice_FromProto(mapCtx, in.TfErrors, TerraformError_FromProto)
	out.ErrorLogs = direct.LazyPtr(in.GetErrorLogs())
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	out.LockState = direct.Enum_FromProto(mapCtx, in.GetLockState())
	// MISSING: TfVersionConstraint
	out.TfVersion = direct.LazyPtr(in.GetTfVersion())
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func DeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Deployment_State](mapCtx, in.State)
	out.LatestRevision = direct.ValueOf(in.LatestRevision)
	out.StateDetail = direct.ValueOf(in.StateDetail)
	out.ErrorCode = direct.Enum_ToProto[pb.Deployment_ErrorCode](mapCtx, in.ErrorCode)
	out.DeleteResults = ApplyResults_ToProto(mapCtx, in.DeleteResults)
	out.DeleteBuild = direct.ValueOf(in.DeleteBuild)
	out.DeleteLogs = direct.ValueOf(in.DeleteLogs)
	out.TfErrors = direct.Slice_ToProto(mapCtx, in.TfErrors, TerraformError_ToProto)
	out.ErrorLogs = direct.ValueOf(in.ErrorLogs)
	// MISSING: ArtifactsGcsBucket
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	out.LockState = direct.Enum_ToProto[pb.Deployment_LockState](mapCtx, in.LockState)
	// MISSING: TfVersionConstraint
	out.TfVersion = direct.ValueOf(in.TfVersion)
	// MISSING: QuotaValidation
	// MISSING: Annotations
	return out
}
func GitSource_FromProto(mapCtx *direct.MapContext, in *pb.GitSource) *krm.GitSource {
	if in == nil {
		return nil
	}
	out := &krm.GitSource{}
	out.Repo = in.Repo
	out.Directory = in.Directory
	out.Ref = in.Ref
	return out
}
func GitSource_ToProto(mapCtx *direct.MapContext, in *krm.GitSource) *pb.GitSource {
	if in == nil {
		return nil
	}
	out := &pb.GitSource{}
	out.Repo = in.Repo
	out.Directory = in.Directory
	out.Ref = in.Ref
	return out
}
func TerraformBlueprint_FromProto(mapCtx *direct.MapContext, in *pb.TerraformBlueprint) *krm.TerraformBlueprint {
	if in == nil {
		return nil
	}
	out := &krm.TerraformBlueprint{}
	out.GcsSource = direct.LazyPtr(in.GetGcsSource())
	out.GitSource = GitSource_FromProto(mapCtx, in.GetGitSource())
	// MISSING: InputValues
	return out
}
func TerraformBlueprint_ToProto(mapCtx *direct.MapContext, in *krm.TerraformBlueprint) *pb.TerraformBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.TerraformBlueprint{}
	if oneof := TerraformBlueprint_GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = oneof
	}
	if oneof := GitSource_ToProto(mapCtx, in.GitSource); oneof != nil {
		out.Source = &pb.TerraformBlueprint_GitSource{GitSource: oneof}
	}
	// MISSING: InputValues
	return out
}
func TerraformError_FromProto(mapCtx *direct.MapContext, in *pb.TerraformError) *krm.TerraformError {
	if in == nil {
		return nil
	}
	out := &krm.TerraformError{}
	out.ResourceAddress = direct.LazyPtr(in.GetResourceAddress())
	out.HTTPResponseCode = direct.LazyPtr(in.GetHttpResponseCode())
	out.ErrorDescription = direct.LazyPtr(in.GetErrorDescription())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func TerraformError_ToProto(mapCtx *direct.MapContext, in *krm.TerraformError) *pb.TerraformError {
	if in == nil {
		return nil
	}
	out := &pb.TerraformError{}
	out.ResourceAddress = direct.ValueOf(in.ResourceAddress)
	out.HttpResponseCode = direct.ValueOf(in.HTTPResponseCode)
	out.ErrorDescription = direct.ValueOf(in.ErrorDescription)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func TerraformOutput_FromProto(mapCtx *direct.MapContext, in *pb.TerraformOutput) *krm.TerraformOutput {
	if in == nil {
		return nil
	}
	out := &krm.TerraformOutput{}
	out.Sensitive = direct.LazyPtr(in.GetSensitive())
	out.Value = Value_FromProto(mapCtx, in.GetValue())
	return out
}
func TerraformOutput_ToProto(mapCtx *direct.MapContext, in *krm.TerraformOutput) *pb.TerraformOutput {
	if in == nil {
		return nil
	}
	out := &pb.TerraformOutput{}
	out.Sensitive = direct.ValueOf(in.Sensitive)
	out.Value = Value_ToProto(mapCtx, in.Value)
	return out
}
func TerraformVariable_FromProto(mapCtx *direct.MapContext, in *pb.TerraformVariable) *krm.TerraformVariable {
	if in == nil {
		return nil
	}
	out := &krm.TerraformVariable{}
	out.InputValue = Value_FromProto(mapCtx, in.GetInputValue())
	return out
}
func TerraformVariable_ToProto(mapCtx *direct.MapContext, in *krm.TerraformVariable) *pb.TerraformVariable {
	if in == nil {
		return nil
	}
	out := &pb.TerraformVariable{}
	out.InputValue = Value_ToProto(mapCtx, in.InputValue)
	return out
}
