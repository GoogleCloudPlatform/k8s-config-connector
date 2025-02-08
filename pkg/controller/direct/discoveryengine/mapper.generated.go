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
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
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
func Reply_FromProto(mapCtx *direct.MapContext, in *pb.Reply) *krm.Reply {
	if in == nil {
		return nil
	}
	out := &krm.Reply{}
	out.Reply = direct.LazyPtr(in.GetReply())
	out.References = direct.Slice_FromProto(mapCtx, in.References, Reply_Reference_FromProto)
	out.Summary = SearchResponse_Summary_FromProto(mapCtx, in.GetSummary())
	return out
}
func Reply_ToProto(mapCtx *direct.MapContext, in *krm.Reply) *pb.Reply {
	if in == nil {
		return nil
	}
	out := &pb.Reply{}
	out.Reply = direct.ValueOf(in.Reply)
	out.References = direct.Slice_ToProto(mapCtx, in.References, Reply_Reference_ToProto)
	out.Summary = SearchResponse_Summary_ToProto(mapCtx, in.Summary)
	return out
}
func Reply_Reference_FromProto(mapCtx *direct.MapContext, in *pb.Reply_Reference) *krm.Reply_Reference {
	if in == nil {
		return nil
	}
	out := &krm.Reply_Reference{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.AnchorText = direct.LazyPtr(in.GetAnchorText())
	out.Start = direct.LazyPtr(in.GetStart())
	out.End = direct.LazyPtr(in.GetEnd())
	return out
}
func Reply_Reference_ToProto(mapCtx *direct.MapContext, in *krm.Reply_Reference) *pb.Reply_Reference {
	if in == nil {
		return nil
	}
	out := &pb.Reply_Reference{}
	out.Uri = direct.ValueOf(in.URI)
	out.AnchorText = direct.ValueOf(in.AnchorText)
	out.Start = direct.ValueOf(in.Start)
	out.End = direct.ValueOf(in.End)
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
