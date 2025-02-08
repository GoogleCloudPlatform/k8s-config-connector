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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DiscoveryEngineDataStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DiscoveryEngineDataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreObservedState{}
	// MISSING: Name
	out.DefaultSchemaID = direct.LazyPtr(in.GetDefaultSchemaId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.BillingEstimation = DataStore_BillingEstimation_FromProto(mapCtx, in.GetBillingEstimation())
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreObservedState) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	out.DefaultSchemaId = direct.ValueOf(in.DefaultSchemaID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.BillingEstimation = DataStore_BillingEstimation_ToProto(mapCtx, in.BillingEstimation)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DiscoveryEngineDataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IndustryVertical = direct.Enum_FromProto(mapCtx, in.GetIndustryVertical())
	out.SolutionTypes = direct.EnumSlice_FromProto(mapCtx, in.SolutionTypes)
	out.ContentConfig = direct.Enum_FromProto(mapCtx, in.GetContentConfig())
	out.WorkspaceConfig = WorkspaceConfig_FromProto(mapCtx, in.GetWorkspaceConfig())
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreSpec) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IndustryVertical = direct.Enum_ToProto[pb.IndustryVertical](mapCtx, in.IndustryVertical)
	out.SolutionTypes = direct.EnumSlice_ToProto[pb.SolutionType](mapCtx, in.SolutionTypes)
	out.ContentConfig = direct.Enum_ToProto[pb.DataStore_ContentConfig](mapCtx, in.ContentConfig)
	out.WorkspaceConfig = WorkspaceConfig_ToProto(mapCtx, in.WorkspaceConfig)
	// MISSING: DocumentProcessingConfig
	// MISSING: StartingSchema
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteObservedState{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// MISSING: GeneratedURIPattern
	// (near miss): "GeneratedURIPattern" vs "GeneratedUriPattern"
	// MISSING: RootDomainURI
	// (near miss): "RootDomainURI" vs "RootDomainUri"
	out.SiteVerificationInfo = SiteVerificationInfo_FromProto(mapCtx, in.GetSiteVerificationInfo())
	out.IndexingStatus = direct.Enum_FromProto(mapCtx, in.GetIndexingStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.FailureReason = TargetSite_FailureReason_FromProto(mapCtx, in.GetFailureReason())
	return out
}
func DiscoveryEngineDataStoreTargetSiteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteObservedState) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// MISSING: GeneratedURIPattern
	// (near miss): "GeneratedURIPattern" vs "GeneratedUriPattern"
	// MISSING: RootDomainURI
	// (near miss): "RootDomainURI" vs "RootDomainUri"
	out.SiteVerificationInfo = SiteVerificationInfo_ToProto(mapCtx, in.SiteVerificationInfo)
	out.IndexingStatus = direct.Enum_ToProto[pb.TargetSite_IndexingStatus](mapCtx, in.IndexingStatus)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.FailureReason = TargetSite_FailureReason_ToProto(mapCtx, in.FailureReason)
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetSite) *krm.DiscoveryEngineDataStoreTargetSiteSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineDataStoreTargetSiteSpec{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// (near miss): "ProvidedURIPattern" vs "ProvidedUriPattern"
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExactMatch = direct.LazyPtr(in.GetExactMatch())
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	return out
}
func DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineDataStoreTargetSiteSpec) *pb.TargetSite {
	if in == nil {
		return nil
	}
	out := &pb.TargetSite{}
	// MISSING: Name
	// MISSING: ProvidedURIPattern
	// (near miss): "ProvidedURIPattern" vs "ProvidedUriPattern"
	out.Type = direct.Enum_ToProto[pb.TargetSite_Type](mapCtx, in.Type)
	out.ExactMatch = direct.ValueOf(in.ExactMatch)
	// MISSING: GeneratedURIPattern
	// MISSING: RootDomainURI
	return out
}
func DiscoveryEngineEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Engine) *krm.DiscoveryEngineEngineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEngineEngineObservedState{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataStoreIds
	return out
}
func DiscoveryEngineEngineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryEngineEngineObservedState) *pb.Engine {
	if in == nil {
		return nil
	}
	out := &pb.Engine{}
	// MISSING: ChatEngineMetadata
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataStoreIds
	return out
}
func DiscoveryengineChunkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.DiscoveryengineChunkObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineChunkObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	// MISSING: RelevanceScore
	// MISSING: DocumentMetadata
	// MISSING: DerivedStructData
	// MISSING: PageSpan
	// MISSING: ChunkMetadata
	return out
}
func DiscoveryengineChunkObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineChunkObservedState) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	// MISSING: RelevanceScore
	// MISSING: DocumentMetadata
	// MISSING: DerivedStructData
	// MISSING: PageSpan
	// MISSING: ChunkMetadata
	return out
}
func DiscoveryengineChunkSpec_FromProto(mapCtx *direct.MapContext, in *pb.Chunk) *krm.DiscoveryengineChunkSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineChunkSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	// MISSING: RelevanceScore
	// MISSING: DocumentMetadata
	// MISSING: DerivedStructData
	// MISSING: PageSpan
	// MISSING: ChunkMetadata
	return out
}
func DiscoveryengineChunkSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineChunkSpec) *pb.Chunk {
	if in == nil {
		return nil
	}
	out := &pb.Chunk{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: Content
	// MISSING: RelevanceScore
	// MISSING: DocumentMetadata
	// MISSING: DerivedStructData
	// MISSING: PageSpan
	// MISSING: ChunkMetadata
	return out
}
