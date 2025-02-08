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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
)
func Conversation_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.Conversation {
	if in == nil {
		return nil
	}
	out := &krm.Conversation{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UserPseudoID = direct.LazyPtr(in.GetUserPseudoId())
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, ConversationMessage_FromProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func Conversation_ToProto(mapCtx *direct.MapContext, in *krm.Conversation) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Conversation_State](mapCtx, in.State)
	out.UserPseudoId = direct.ValueOf(in.UserPseudoID)
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, ConversationMessage_ToProto)
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func ConversationContext_FromProto(mapCtx *direct.MapContext, in *pb.ConversationContext) *krm.ConversationContext {
	if in == nil {
		return nil
	}
	out := &krm.ConversationContext{}
	out.ContextDocuments = in.ContextDocuments
	out.ActiveDocument = direct.LazyPtr(in.GetActiveDocument())
	return out
}
func ConversationContext_ToProto(mapCtx *direct.MapContext, in *krm.ConversationContext) *pb.ConversationContext {
	if in == nil {
		return nil
	}
	out := &pb.ConversationContext{}
	out.ContextDocuments = in.ContextDocuments
	out.ActiveDocument = direct.ValueOf(in.ActiveDocument)
	return out
}
func ConversationMessage_FromProto(mapCtx *direct.MapContext, in *pb.ConversationMessage) *krm.ConversationMessage {
	if in == nil {
		return nil
	}
	out := &krm.ConversationMessage{}
	out.UserInput = TextInput_FromProto(mapCtx, in.GetUserInput())
	out.Reply = Reply_FromProto(mapCtx, in.GetReply())
	// MISSING: CreateTime
	return out
}
func ConversationMessage_ToProto(mapCtx *direct.MapContext, in *krm.ConversationMessage) *pb.ConversationMessage {
	if in == nil {
		return nil
	}
	out := &pb.ConversationMessage{}
	if oneof := TextInput_ToProto(mapCtx, in.UserInput); oneof != nil {
		out.Message = &pb.ConversationMessage_UserInput{UserInput: oneof}
	}
	if oneof := Reply_ToProto(mapCtx, in.Reply); oneof != nil {
		out.Message = &pb.ConversationMessage_Reply{Reply: oneof}
	}
	// MISSING: CreateTime
	return out
}
func ConversationMessageObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationMessage) *krm.ConversationMessageObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationMessageObservedState{}
	// MISSING: UserInput
	// MISSING: Reply
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func ConversationMessageObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationMessageObservedState) *pb.ConversationMessage {
	if in == nil {
		return nil
	}
	out := &pb.ConversationMessage{}
	// MISSING: UserInput
	// MISSING: Reply
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func ConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.ConversationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, ConversationMessageObservedState_FromProto)
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func ConversationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationObservedState) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, ConversationMessageObservedState_ToProto)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
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
func DiscoveryengineConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.DiscoveryengineConversationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineConversationObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Messages
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineConversationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineConversationObservedState) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Messages
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineConversationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.DiscoveryengineConversationSpec {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryengineConversationSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Messages
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func DiscoveryengineConversationSpec_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryengineConversationSpec) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: State
	// MISSING: UserPseudoID
	// MISSING: Messages
	// MISSING: StartTime
	// MISSING: EndTime
	return out
}
func Reply_FromProto(mapCtx *direct.MapContext, in *pb.Reply) *krm.Reply {
	if in == nil {
		return nil
	}
	out := &krm.Reply{}
	out.Summary = SearchResponse_Summary_FromProto(mapCtx, in.GetSummary())
	return out
}
func Reply_ToProto(mapCtx *direct.MapContext, in *krm.Reply) *pb.Reply {
	if in == nil {
		return nil
	}
	out := &pb.Reply{}
	out.Summary = SearchResponse_Summary_ToProto(mapCtx, in.Summary)
	return out
}
func TextInput_FromProto(mapCtx *direct.MapContext, in *pb.TextInput) *krm.TextInput {
	if in == nil {
		return nil
	}
	out := &krm.TextInput{}
	out.Input = direct.LazyPtr(in.GetInput())
	out.Context = ConversationContext_FromProto(mapCtx, in.GetContext())
	return out
}
func TextInput_ToProto(mapCtx *direct.MapContext, in *krm.TextInput) *pb.TextInput {
	if in == nil {
		return nil
	}
	out := &pb.TextInput{}
	out.Input = direct.ValueOf(in.Input)
	out.Context = ConversationContext_ToProto(mapCtx, in.Context)
	return out
}
