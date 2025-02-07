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

package ai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiCachedContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiCachedContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiCachedContentObservedState{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiCachedContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiCachedContentObservedState) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiCachedContentSpec_FromProto(mapCtx *direct.MapContext, in *pb.CachedContent) *krm.AiCachedContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiCachedContentSpec{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiCachedContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiCachedContentSpec) *pb.CachedContent {
	if in == nil {
		return nil
	}
	out := &pb.CachedContent{}
	// MISSING: ExpireTime
	// MISSING: Ttl
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Model
	// MISSING: SystemInstruction
	// MISSING: Contents
	// MISSING: Tools
	// MISSING: ToolConfig
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: UsageMetadata
	return out
}
func AiFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.AiFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiFileObservedState{}
	// MISSING: VideoMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func AiFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiFileObservedState) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	// MISSING: VideoMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func AiFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.AiFileSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiFileSpec{}
	// MISSING: VideoMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func AiFileSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiFileSpec) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	// MISSING: VideoMetadata
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func File_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.File {
	if in == nil {
		return nil
	}
	out := &krm.File{}
	// MISSING: VideoMetadata
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func File_ToProto(mapCtx *direct.MapContext, in *krm.File) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	// MISSING: VideoMetadata
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: MimeType
	// MISSING: SizeBytes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ExpirationTime
	// MISSING: Sha256Hash
	// MISSING: URI
	// MISSING: State
	// MISSING: Error
	return out
}
func FileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.FileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FileObservedState{}
	out.VideoMetadata = VideoMetadata_FromProto(mapCtx, in.GetVideoMetadata())
	// MISSING: Name
	// MISSING: DisplayName
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.SizeBytes = direct.LazyPtr(in.GetSizeBytes())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpirationTime())
	out.Sha256Hash = in.GetSha256Hash()
	out.URI = direct.LazyPtr(in.GetUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func FileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FileObservedState) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	if oneof := VideoMetadata_ToProto(mapCtx, in.VideoMetadata); oneof != nil {
		out.Metadata = &pb.File_VideoMetadata{VideoMetadata: oneof}
	}
	// MISSING: Name
	// MISSING: DisplayName
	out.MimeType = direct.ValueOf(in.MimeType)
	out.SizeBytes = direct.ValueOf(in.SizeBytes)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpirationTime)
	out.Sha256Hash = in.Sha256Hash
	out.Uri = direct.ValueOf(in.URI)
	out.State = direct.Enum_ToProto[pb.File_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
