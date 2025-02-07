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
func AiChunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.AiChunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiChunkObservedState{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiChunkObservedState) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.AiChunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiChunkSpec{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiChunkSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiChunkSpec) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func AiCorpusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusObservedState) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_FromProto(mapCtx *direct.MapContext, in *pb.Corpus) *krm.AiCorpusSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiCorpusSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiCorpusSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiCorpusSpec) *pb.Corpus {
	if in == nil {
		return nil
	}
	out := &pb.Corpus{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func AiDocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.AiDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiDocumentSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiDocumentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomMetadata
	// MISSING: UpdateTime
	// MISSING: CreateTime
	return out
}
func AiPermissionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionObservedState{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionObservedState) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Permission) *krm.AiPermissionSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiPermissionSpec{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func AiPermissionSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiPermissionSpec) *pb.Permission {
	if in == nil {
		return nil
	}
	out := &pb.Permission{}
	// MISSING: Name
	// MISSING: GranteeType
	// MISSING: EmailAddress
	// MISSING: Role
	return out
}
func Chunk_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.Chunk {
	if in == nil {
		return nil
	}
	out := &krm.Chunk{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Data = ChunkData_FromProto(mapCtx, in.GetData())
	out.CustomMetadata = direct.Slice_FromProto(mapCtx, in.CustomMetadata, CustomMetadata_FromProto)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func Chunk_ToProto(mapCtx *direct.MapContext, in *krm.Chunk) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	out.Name = direct.ValueOf(in.Name)
	out.Data = ChunkData_ToProto(mapCtx, in.Data)
	out.CustomMetadata = direct.Slice_ToProto(mapCtx, in.CustomMetadata, CustomMetadata_ToProto)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	return out
}
func ChunkData_FromProto(mapCtx *direct.MapContext, in *pb.ChunkData) *krm.ChunkData {
	if in == nil {
		return nil
	}
	out := &krm.ChunkData{}
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	return out
}
func ChunkData_ToProto(mapCtx *direct.MapContext, in *krm.ChunkData) *pb.ChunkData {
	if in == nil {
		return nil
	}
	out := &pb.ChunkData{}
	if oneof := ChunkData_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Data = oneof
	}
	return out
}
func ChunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.ChunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChunkObservedState{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ChunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChunkObservedState) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: Data
	// MISSING: CustomMetadata
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Chunk_State](mapCtx, in.State)
	return out
}
func StringList_FromProto(mapCtx *direct.MapContext, in *pb.StringList) *krm.StringList {
	if in == nil {
		return nil
	}
	out := &krm.StringList{}
	out.Values = in.Values
	return out
}
func StringList_ToProto(mapCtx *direct.MapContext, in *krm.StringList) *pb.StringList {
	if in == nil {
		return nil
	}
	out := &pb.StringList{}
	out.Values = in.Values
	return out
}
