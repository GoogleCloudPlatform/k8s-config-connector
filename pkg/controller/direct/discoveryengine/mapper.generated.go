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
func DiscoveryengineSessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DiscoveryengineSessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSessionObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineSessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineSessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.DiscoveryengineSessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineSessionSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineSessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineSessionSpec) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func Query_FromProto(mapCtx *direct.MapContext, in *pb.Query) *krm.Query {
	if in == nil {
		return nil
	}
	out := &krm.Query{}
	out.Text = direct.LazyPtr(in.GetText())
	out.QueryID = direct.LazyPtr(in.GetQueryId())
	return out
}
func Query_ToProto(mapCtx *direct.MapContext, in *krm.Query) *pb.Query {
	if in == nil {
		return nil
	}
	out := &pb.Query{}
	if oneof := Query_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Content = oneof
	}
	out.QueryId = direct.ValueOf(in.QueryID)
	return out
}
func Session_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.Session {
	if in == nil {
		return nil
	}
	out := &krm.Session{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UserPseudoID = direct.LazyPtr(in.GetUserPseudoId())
	out.Turns = direct.Slice_FromProto(mapCtx, in.Turns, Session_Turn_FromProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func Session_ToProto(mapCtx *direct.MapContext, in *krm.Session) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Session_State](mapCtx, in.State)
	out.UserPseudoId = direct.ValueOf(in.UserPseudoID)
	out.Turns = direct.Slice_ToProto(mapCtx, in.Turns, Session_Turn_ToProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func SessionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Session) *krm.SessionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SessionObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func SessionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SessionObservedState) *pb.Session {
	if in == nil {
		return nil
	}
	out := &pb.Session{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Turns
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func Session_Turn_FromProto(mapCtx *direct.MapContext, in *pb.Session_Turn) *krm.Session_Turn {
	if in == nil {
		return nil
	}
	out := &krm.Session_Turn{}
	out.Query = Query_FromProto(mapCtx, in.GetQuery())
	out.Answer = direct.LazyPtr(in.GetAnswer())
	return out
}
func Session_Turn_ToProto(mapCtx *direct.MapContext, in *krm.Session_Turn) *pb.Session_Turn {
	if in == nil {
		return nil
	}
	out := &pb.Session_Turn{}
	out.Query = Query_ToProto(mapCtx, in.Query)
	out.Answer = direct.ValueOf(in.Answer)
	return out
}
