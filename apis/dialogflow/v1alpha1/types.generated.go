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

package v1alpha1


// +kcc:proto=google.cloud.dialogflow.v2beta1.Conversation
type Conversation struct {

	// Required. The Conversation Profile to be used to configure this
	//  Conversation. This field cannot be updated.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/conversationProfiles/<Conversation Profile ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.conversation_profile
	ConversationProfile *string `json:"conversationProfile,omitempty"`

	// Optional. The stage of a conversation. It indicates whether the virtual
	//  agent or a human agent is handling the conversation.
	//
	//  If the conversation is created with the conversation profile that has
	//  Dialogflow config set, defaults to
	//  [ConversationStage.VIRTUAL_AGENT_STAGE][google.cloud.dialogflow.v2beta1.Conversation.ConversationStage.VIRTUAL_AGENT_STAGE];
	//  Otherwise, defaults to
	//  [ConversationStage.HUMAN_ASSIST_STAGE][google.cloud.dialogflow.v2beta1.Conversation.ConversationStage.HUMAN_ASSIST_STAGE].
	//
	//  If the conversation is created with the conversation profile that has
	//  Dialogflow config set but explicitly sets conversation_stage to
	//  [ConversationStage.HUMAN_ASSIST_STAGE][google.cloud.dialogflow.v2beta1.Conversation.ConversationStage.HUMAN_ASSIST_STAGE],
	//  it skips
	//  [ConversationStage.VIRTUAL_AGENT_STAGE][google.cloud.dialogflow.v2beta1.Conversation.ConversationStage.VIRTUAL_AGENT_STAGE]
	//  stage and directly goes to
	//  [ConversationStage.HUMAN_ASSIST_STAGE][google.cloud.dialogflow.v2beta1.Conversation.ConversationStage.HUMAN_ASSIST_STAGE].
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.conversation_stage
	ConversationStage *string `json:"conversationStage,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.ConversationPhoneNumber
type ConversationPhoneNumber struct {
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Conversation
type ConversationObservedState struct {
	// Output only. Identifier. The unique identifier of this conversation.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/conversations/<Conversation ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.name
	Name *string `json:"name,omitempty"`

	// Output only. The current state of the Conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// Output only. Required if the conversation is to be connected over
	//  telephony.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.phone_number
	PhoneNumber *ConversationPhoneNumber `json:"phoneNumber,omitempty"`

	// Output only. The time the conversation was started.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the conversation was finished.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Conversation.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.ConversationPhoneNumber
type ConversationPhoneNumberObservedState struct {
	// Output only. The phone number to connect to this conversation.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.ConversationPhoneNumber.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}
