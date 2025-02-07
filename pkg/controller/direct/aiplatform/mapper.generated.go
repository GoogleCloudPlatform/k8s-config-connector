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

package aiplatform

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiplatformAnnotationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationObservedState{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationObservedState) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Annotation) *krm.AiplatformAnnotationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpec{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpec) *pb.Annotation {
	if in == nil {
		return nil
	}
	out := &pb.Annotation{}
	// MISSING: Name
	// MISSING: PayloadSchemaURI
	// MISSING: Payload
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: AnnotationSource
	// MISSING: Labels
	return out
}
func AiplatformAnnotationSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecObservedState) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.AnnotationSpec) *krm.AiplatformAnnotationSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformAnnotationSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformAnnotationSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformAnnotationSpecSpec) *pb.AnnotationSpec {
	if in == nil {
		return nil
	}
	out := &pb.AnnotationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func AiplatformArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactObservedState) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.AiplatformArtifactSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiplatformArtifactSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func AiplatformArtifactSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiplatformArtifactSpec) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func Artifact_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.Artifact {
	if in == nil {
		return nil
	}
	out := &krm.Artifact{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SchemaTitle = direct.LazyPtr(in.GetSchemaTitle())
	out.SchemaVersion = direct.LazyPtr(in.GetSchemaVersion())
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Artifact_ToProto(mapCtx *direct.MapContext, in *krm.Artifact) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.State = direct.Enum_ToProto[pb.Artifact_State](mapCtx, in.State)
	out.SchemaTitle = direct.ValueOf(in.SchemaTitle)
	out.SchemaVersion = direct.ValueOf(in.SchemaVersion)
	out.Metadata = Metadata_ToProto(mapCtx, in.Metadata)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func ArtifactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Artifact) *krm.ArtifactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
func ArtifactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactObservedState) *pb.Artifact {
	if in == nil {
		return nil
	}
	out := &pb.Artifact{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: URI
	// MISSING: Etag
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: State
	// MISSING: SchemaTitle
	// MISSING: SchemaVersion
	// MISSING: Metadata
	// MISSING: Description
	return out
}
