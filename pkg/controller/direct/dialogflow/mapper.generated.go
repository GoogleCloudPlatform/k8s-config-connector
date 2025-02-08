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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Conversation_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.Conversation {
	if in == nil {
		return nil
	}
	out := &krm.Conversation{}
	// MISSING: Name
	// MISSING: LifecycleState
	out.ConversationProfile = direct.LazyPtr(in.GetConversationProfile())
	// MISSING: PhoneNumber
	// MISSING: StartTime
	// MISSING: EndTime
	out.ConversationStage = direct.Enum_FromProto(mapCtx, in.GetConversationStage())
	return out
}
func Conversation_ToProto(mapCtx *direct.MapContext, in *krm.Conversation) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	// MISSING: Name
	// MISSING: LifecycleState
	out.ConversationProfile = direct.ValueOf(in.ConversationProfile)
	// MISSING: PhoneNumber
	// MISSING: StartTime
	// MISSING: EndTime
	out.ConversationStage = direct.Enum_ToProto[pb.Conversation_ConversationStage](mapCtx, in.ConversationStage)
	return out
}
func ConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.ConversationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	// MISSING: ConversationProfile
	out.PhoneNumber = ConversationPhoneNumber_FromProto(mapCtx, in.GetPhoneNumber())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: ConversationStage
	return out
}
func ConversationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationObservedState) *pb.Conversation {
	if in == nil {
		return nil
	}
	out := &pb.Conversation{}
	out.Name = direct.ValueOf(in.Name)
	out.LifecycleState = direct.Enum_ToProto[pb.Conversation_LifecycleState](mapCtx, in.LifecycleState)
	// MISSING: ConversationProfile
	out.PhoneNumber = ConversationPhoneNumber_ToProto(mapCtx, in.PhoneNumber)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: ConversationStage
	return out
}
func ConversationPhoneNumber_FromProto(mapCtx *direct.MapContext, in *pb.ConversationPhoneNumber) *krm.ConversationPhoneNumber {
	if in == nil {
		return nil
	}
	out := &krm.ConversationPhoneNumber{}
	// MISSING: PhoneNumber
	return out
}
func ConversationPhoneNumber_ToProto(mapCtx *direct.MapContext, in *krm.ConversationPhoneNumber) *pb.ConversationPhoneNumber {
	if in == nil {
		return nil
	}
	out := &pb.ConversationPhoneNumber{}
	// MISSING: PhoneNumber
	return out
}
func ConversationPhoneNumberObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationPhoneNumber) *krm.ConversationPhoneNumberObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationPhoneNumberObservedState{}
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	return out
}
func ConversationPhoneNumberObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationPhoneNumberObservedState) *pb.ConversationPhoneNumber {
	if in == nil {
		return nil
	}
	out := &pb.ConversationPhoneNumber{}
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	return out
}
