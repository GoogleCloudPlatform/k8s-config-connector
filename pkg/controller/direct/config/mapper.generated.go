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
func ConfigPreviewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Preview) *krm.ConfigPreviewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigPreviewObservedState{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	// MISSING: TfVersionConstraint
	// MISSING: Annotations
	return out
}
func ConfigPreviewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigPreviewObservedState) *pb.Preview {
	if in == nil {
		return nil
	}
	out := &pb.Preview{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	// MISSING: TfVersionConstraint
	// MISSING: Annotations
	return out
}
func ConfigPreviewSpec_FromProto(mapCtx *direct.MapContext, in *pb.Preview) *krm.ConfigPreviewSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigPreviewSpec{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	// MISSING: TfVersionConstraint
	// MISSING: Annotations
	return out
}
func ConfigPreviewSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigPreviewSpec) *pb.Preview {
	if in == nil {
		return nil
	}
	out := &pb.Preview{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	// MISSING: TfVersionConstraint
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
func Preview_FromProto(mapCtx *direct.MapContext, in *pb.Preview) *krm.Preview {
	if in == nil {
		return nil
	}
	out := &krm.Preview{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: State
	out.Deployment = direct.LazyPtr(in.GetDeployment())
	out.PreviewMode = direct.Enum_FromProto(mapCtx, in.GetPreviewMode())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactsGcsBucket = in.ArtifactsGcsBucket
	out.WorkerPool = in.WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	out.TfVersionConstraint = in.TfVersionConstraint
	out.Annotations = in.Annotations
	return out
}
func Preview_ToProto(mapCtx *direct.MapContext, in *krm.Preview) *pb.Preview {
	if in == nil {
		return nil
	}
	out := &pb.Preview{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Preview_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: State
	out.Deployment = direct.ValueOf(in.Deployment)
	out.PreviewMode = direct.Enum_ToProto[pb.Preview_PreviewMode](mapCtx, in.PreviewMode)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactsGcsBucket = in.ArtifactsGcsBucket
	out.WorkerPool = in.WorkerPool
	// MISSING: ErrorCode
	// MISSING: ErrorStatus
	// MISSING: Build
	// MISSING: TfErrors
	// MISSING: ErrorLogs
	// MISSING: PreviewArtifacts
	// MISSING: Logs
	// MISSING: TfVersion
	out.TfVersionConstraint = in.TfVersionConstraint
	out.Annotations = in.Annotations
	return out
}
func PreviewArtifacts_FromProto(mapCtx *direct.MapContext, in *pb.PreviewArtifacts) *krm.PreviewArtifacts {
	if in == nil {
		return nil
	}
	out := &krm.PreviewArtifacts{}
	// MISSING: Content
	// MISSING: Artifacts
	return out
}
func PreviewArtifacts_ToProto(mapCtx *direct.MapContext, in *krm.PreviewArtifacts) *pb.PreviewArtifacts {
	if in == nil {
		return nil
	}
	out := &pb.PreviewArtifacts{}
	// MISSING: Content
	// MISSING: Artifacts
	return out
}
func PreviewArtifactsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PreviewArtifacts) *krm.PreviewArtifactsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PreviewArtifactsObservedState{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.Artifacts = direct.LazyPtr(in.GetArtifacts())
	return out
}
func PreviewArtifactsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PreviewArtifactsObservedState) *pb.PreviewArtifacts {
	if in == nil {
		return nil
	}
	out := &pb.PreviewArtifacts{}
	out.Content = direct.ValueOf(in.Content)
	out.Artifacts = direct.ValueOf(in.Artifacts)
	return out
}
func PreviewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Preview) *krm.PreviewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PreviewObservedState{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	out.ErrorCode = direct.Enum_FromProto(mapCtx, in.GetErrorCode())
	out.ErrorStatus = Status_FromProto(mapCtx, in.GetErrorStatus())
	out.Build = direct.LazyPtr(in.GetBuild())
	out.TfErrors = direct.Slice_FromProto(mapCtx, in.TfErrors, TerraformError_FromProto)
	out.ErrorLogs = direct.LazyPtr(in.GetErrorLogs())
	out.PreviewArtifacts = PreviewArtifacts_FromProto(mapCtx, in.GetPreviewArtifacts())
	out.Logs = direct.LazyPtr(in.GetLogs())
	out.TfVersion = direct.LazyPtr(in.GetTfVersion())
	// MISSING: TfVersionConstraint
	// MISSING: Annotations
	return out
}
func PreviewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PreviewObservedState) *pb.Preview {
	if in == nil {
		return nil
	}
	out := &pb.Preview{}
	// MISSING: TerraformBlueprint
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	out.State = direct.Enum_ToProto[pb.Preview_State](mapCtx, in.State)
	// MISSING: Deployment
	// MISSING: PreviewMode
	// MISSING: ServiceAccount
	// MISSING: ArtifactsGcsBucket
	// MISSING: WorkerPool
	out.ErrorCode = direct.Enum_ToProto[pb.Preview_ErrorCode](mapCtx, in.ErrorCode)
	out.ErrorStatus = Status_ToProto(mapCtx, in.ErrorStatus)
	out.Build = direct.ValueOf(in.Build)
	out.TfErrors = direct.Slice_ToProto(mapCtx, in.TfErrors, TerraformError_ToProto)
	out.ErrorLogs = direct.ValueOf(in.ErrorLogs)
	out.PreviewArtifacts = PreviewArtifacts_ToProto(mapCtx, in.PreviewArtifacts)
	out.Logs = direct.ValueOf(in.Logs)
	out.TfVersion = direct.ValueOf(in.TfVersion)
	// MISSING: TfVersionConstraint
	// MISSING: Annotations
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
