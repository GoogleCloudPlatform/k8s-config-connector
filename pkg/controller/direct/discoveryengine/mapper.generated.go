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
func DiscoveryengineDocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DiscoveryengineDocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineDocumentObservedState{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func DiscoveryengineDocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineDocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func DiscoveryengineDocumentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DiscoveryengineDocumentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineDocumentSpec{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func DiscoveryengineDocumentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineDocumentSpec) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func Document_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.Document {
	if in == nil {
		return nil
	}
	out := &krm.Document{}
	out.StructData = StructData_FromProto(mapCtx, in.GetStructData())
	out.JsonData = direct.LazyPtr(in.GetJsonData())
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.SchemaID = direct.LazyPtr(in.GetSchemaId())
	out.Content = Document_Content_FromProto(mapCtx, in.GetContent())
	out.ParentDocumentID = direct.LazyPtr(in.GetParentDocumentId())
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func Document_ToProto(mapCtx *direct.MapContext, in *krm.Document) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	if oneof := StructData_ToProto(mapCtx, in.StructData); oneof != nil {
		out.Data = &pb.Document_StructData{StructData: oneof}
	}
	if oneof := Document_JsonData_ToProto(mapCtx, in.JsonData); oneof != nil {
		out.Data = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.SchemaId = direct.ValueOf(in.SchemaID)
	out.Content = Document_Content_ToProto(mapCtx, in.Content)
	out.ParentDocumentId = direct.ValueOf(in.ParentDocumentID)
	// MISSING: DerivedStructData
	// MISSING: IndexTime
	// MISSING: IndexStatus
	return out
}
func DocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentObservedState{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	out.DerivedStructData = DerivedStructData_FromProto(mapCtx, in.GetDerivedStructData())
	out.IndexTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIndexTime())
	out.IndexStatus = Document_IndexStatus_FromProto(mapCtx, in.GetIndexStatus())
	return out
}
func DocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: StructData
	// MISSING: JsonData
	// MISSING: Name
	// MISSING: ID
	// MISSING: SchemaID
	// MISSING: Content
	// MISSING: ParentDocumentID
	out.DerivedStructData = DerivedStructData_ToProto(mapCtx, in.DerivedStructData)
	out.IndexTime = direct.StringTimestamp_ToProto(mapCtx, in.IndexTime)
	out.IndexStatus = Document_IndexStatus_ToProto(mapCtx, in.IndexStatus)
	return out
}
func Document_Content_FromProto(mapCtx *direct.MapContext, in *pb.Document_Content) *krm.Document_Content {
	if in == nil {
		return nil
	}
	out := &krm.Document_Content{}
	out.RawBytes = in.GetRawBytes()
	out.URI = direct.LazyPtr(in.GetUri())
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	return out
}
func Document_Content_ToProto(mapCtx *direct.MapContext, in *krm.Document_Content) *pb.Document_Content {
	if in == nil {
		return nil
	}
	out := &pb.Document_Content{}
	if oneof := Document_Content_RawBytes_ToProto(mapCtx, in.RawBytes); oneof != nil {
		out.Content = oneof
	}
	if oneof := Document_Content_Uri_ToProto(mapCtx, in.URI); oneof != nil {
		out.Content = oneof
	}
	out.MimeType = direct.ValueOf(in.MimeType)
	return out
}
func Document_IndexStatus_FromProto(mapCtx *direct.MapContext, in *pb.Document_IndexStatus) *krm.Document_IndexStatus {
	if in == nil {
		return nil
	}
	out := &krm.Document_IndexStatus{}
	out.IndexTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIndexTime())
	out.ErrorSamples = direct.Slice_FromProto(mapCtx, in.ErrorSamples, Status_FromProto)
	return out
}
func Document_IndexStatus_ToProto(mapCtx *direct.MapContext, in *krm.Document_IndexStatus) *pb.Document_IndexStatus {
	if in == nil {
		return nil
	}
	out := &pb.Document_IndexStatus{}
	out.IndexTime = direct.StringTimestamp_ToProto(mapCtx, in.IndexTime)
	out.ErrorSamples = direct.Slice_ToProto(mapCtx, in.ErrorSamples, Status_ToProto)
	return out
}
