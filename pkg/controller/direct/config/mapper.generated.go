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
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/config/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func ConfigRevisionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Revision) *krm.ConfigRevisionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigRevisionObservedState{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	// MISSING: QuotaValidation
	return out
}
func ConfigRevisionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigRevisionObservedState) *pb.Revision {
	if in == nil {
		return nil
	}
	out := &pb.Revision{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	// MISSING: QuotaValidation
	return out
}
func ConfigRevisionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Revision) *krm.ConfigRevisionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigRevisionSpec{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	// MISSING: QuotaValidation
	return out
}
func ConfigRevisionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigRevisionSpec) *pb.Revision {
	if in == nil {
		return nil
	}
	out := &pb.Revision{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	// MISSING: QuotaValidation
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
func Revision_FromProto(mapCtx *direct.MapContext, in *pb.Revision) *krm.Revision {
	if in == nil {
		return nil
	}
	out := &krm.Revision{}
	// MISSING: TerraformBlueprint
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	out.QuotaValidation = direct.Enum_FromProto(mapCtx, in.GetQuotaValidation())
	return out
}
func Revision_ToProto(mapCtx *direct.MapContext, in *krm.Revision) *pb.Revision {
	if in == nil {
		return nil
	}
	out := &pb.Revision{}
	// MISSING: TerraformBlueprint
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Action
	// MISSING: State
	// MISSING: ApplyResults
	// MISSING: StateDetail
	// MISSING: ErrorCode
	// MISSING: Build
	// MISSING: Logs
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: ServiceAccount
	// MISSING: ImportExistingResources
	// MISSING: WorkerPool
	// MISSING: TfVersionConstraint
	// MISSING: TfVersion
	// MISSING: QuotaValidationResults
	out.QuotaValidation = direct.Enum_ToProto[pb.QuotaValidation](mapCtx, in.QuotaValidation)
	return out
}
func RevisionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Revision) *krm.RevisionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RevisionObservedState{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ApplyResults = ApplyResults_FromProto(mapCtx, in.GetApplyResults())
	out.StateDetail = direct.LazyPtr(in.GetStateDetail())
	out.ErrorCode = direct.Enum_FromProto(mapCtx, in.GetErrorCode())
	out.Build = direct.LazyPtr(in.GetBuild())
	out.Logs = direct.LazyPtr(in.GetLogs())
	out.TfErrors = direct.Slice_FromProto(mapCtx, in.TfErrors, TerraformError_FromProto)
	out.ErrorLogs = direct.LazyPtr(in.GetErrorLogs())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ImportExistingResources = direct.LazyPtr(in.GetImportExistingResources())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.TfVersionConstraint = direct.LazyPtr(in.GetTfVersionConstraint())
	out.TfVersion = direct.LazyPtr(in.GetTfVersion())
	out.QuotaValidationResults = direct.LazyPtr(in.GetQuotaValidationResults())
	// MISSING: QuotaValidation
	return out
}
func RevisionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RevisionObservedState) *pb.Revision {
	if in == nil {
		return nil
	}
	out := &pb.Revision{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Revision_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Action = direct.Enum_ToProto[pb.Revision_Action](mapCtx, in.Action)
	out.State = direct.Enum_ToProto[pb.Revision_State](mapCtx, in.State)
	out.ApplyResults = ApplyResults_ToProto(mapCtx, in.ApplyResults)
	out.StateDetail = direct.ValueOf(in.StateDetail)
	out.ErrorCode = direct.Enum_ToProto[pb.Revision_ErrorCode](mapCtx, in.ErrorCode)
	out.Build = direct.ValueOf(in.Build)
	out.Logs = direct.ValueOf(in.Logs)
	out.TfErrors = direct.Slice_ToProto(mapCtx, in.TfErrors, TerraformError_ToProto)
	out.ErrorLogs = direct.ValueOf(in.ErrorLogs)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ImportExistingResources = direct.ValueOf(in.ImportExistingResources)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.TfVersionConstraint = direct.ValueOf(in.TfVersionConstraint)
	out.TfVersion = direct.ValueOf(in.TfVersion)
	out.QuotaValidationResults = direct.ValueOf(in.QuotaValidationResults)
	// MISSING: QuotaValidation
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
