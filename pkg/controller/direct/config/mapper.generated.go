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
