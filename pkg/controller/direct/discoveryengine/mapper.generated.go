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

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Chunk_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.Chunk {
	if in == nil {
		return nil
	}
	out := &krm.Chunk{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.Content = direct.LazyPtr(in.GetContent())
	// MISSING: RelevanceScore
	out.DocumentMetadata = Chunk_DocumentMetadata_FromProto(mapCtx, in.GetDocumentMetadata())
	// MISSING: DerivedStructData
	out.PageSpan = Chunk_PageSpan_FromProto(mapCtx, in.GetPageSpan())
	// MISSING: ChunkMetadata
	return out
}
func Chunk_ToProto(mapCtx *direct.MapContext, in *krm.Chunk) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.Content = direct.ValueOf(in.Content)
	// MISSING: RelevanceScore
	out.DocumentMetadata = Chunk_DocumentMetadata_ToProto(mapCtx, in.DocumentMetadata)
	// MISSING: DerivedStructData
	out.PageSpan = Chunk_PageSpan_ToProto(mapCtx, in.PageSpan)
	// MISSING: ChunkMetadata
	return out
}
func ChunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.ChunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChunkObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	out.RelevanceScore = in.RelevanceScore
	// MISSING: DocumentMetadata
	out.DerivedStructData = DerivedStructData_FromProto(mapCtx, in.GetDerivedStructData())
	// MISSING: PageSpan
	out.ChunkMetadata = Chunk_ChunkMetadata_FromProto(mapCtx, in.GetChunkMetadata())
	return out
}
func ChunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChunkObservedState) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	out.RelevanceScore = in.RelevanceScore
	// MISSING: DocumentMetadata
	out.DerivedStructData = DerivedStructData_ToProto(mapCtx, in.DerivedStructData)
	// MISSING: PageSpan
	out.ChunkMetadata = Chunk_ChunkMetadata_ToProto(mapCtx, in.ChunkMetadata)
	return out
}
func Chunk_PageSpan_FromProto(mapCtx *direct.MapContext, in *pb.Chunk_PageSpan) *krm.Chunk_PageSpan {
	if in == nil {
		return nil
	}
	out := &krm.Chunk_PageSpan{}
	out.PageStart = direct.LazyPtr(in.GetPageStart())
	out.PageEnd = direct.LazyPtr(in.GetPageEnd())
	return out
}
func Chunk_PageSpan_ToProto(mapCtx *direct.MapContext, in *krm.Chunk_PageSpan) *pb.Chunk_PageSpan {
	if in == nil {
		return nil
	}
	out := &pb.Chunk_PageSpan{}
	out.PageStart = direct.ValueOf(in.PageStart)
	out.PageEnd = direct.ValueOf(in.PageEnd)
	return out
}
