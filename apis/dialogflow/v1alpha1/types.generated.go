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


// +kcc:proto=google.cloud.dialogflow.v2.Participant
type Participant struct {
	// Optional. The unique identifier of this participant.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/conversations/<Conversation ID>/participants/<Participant ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.name
	Name *string `json:"name,omitempty"`

	// Immutable. The role this participant plays in the conversation. This field
	//  must be set during participant creation and is then immutable.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.role
	Role *string `json:"role,omitempty"`

	// Optional. Label applied to streams representing this participant in SIPREC
	//  XML metadata and SDP. This is used to assign transcriptions from that
	//  media stream to this participant. This field can be updated.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.sip_recording_media_label
	SipRecordingMediaLabel *string `json:"sipRecordingMediaLabel,omitempty"`

	// Optional. Obfuscated user id that should be associated with the created
	//  participant.
	//
	//  You can specify a user id as follows:
	//
	//  1. If you set this field in
	//     [CreateParticipantRequest][google.cloud.dialogflow.v2.CreateParticipantRequest.participant]
	//     or
	//     [UpdateParticipantRequest][google.cloud.dialogflow.v2.UpdateParticipantRequest.participant],
	//     Dialogflow adds the obfuscated user id with the participant.
	//
	//  2. If you set this field in
	//     [AnalyzeContent][google.cloud.dialogflow.v2.AnalyzeContentRequest.participant]
	//     or
	//     [StreamingAnalyzeContent][google.cloud.dialogflow.v2.StreamingAnalyzeContentRequest.participant],
	//     Dialogflow will update
	//     [Participant.obfuscated_external_user_id][google.cloud.dialogflow.v2.Participant.obfuscated_external_user_id].
	//
	//  Dialogflow returns an error if you try to add a user id for a
	//  non-[END_USER][google.cloud.dialogflow.v2.Participant.Role.END_USER]
	//  participant.
	//
	//  Dialogflow uses this user id for billing and measurement purposes. For
	//  example, Dialogflow determines whether a user in one conversation returned
	//  in a later conversation.
	//
	//  Note:
	//
	//  * Please never pass raw user ids to Dialogflow. Always obfuscate your user
	//    id first.
	//  * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a
	//    hash function like SHA-512.
	//  * The length of the user id must be <= 256 characters.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.obfuscated_external_user_id
	ObfuscatedExternalUserID *string `json:"obfuscatedExternalUserID,omitempty"`

	// Optional. Key-value filters on the metadata of documents returned by
	//  article suggestion. If specified, article suggestion only returns suggested
	//  documents that match all filters in their
	//  [Document.metadata][google.cloud.dialogflow.v2.Document.metadata]. Multiple
	//  values for a metadata key should be concatenated by comma. For example,
	//  filters to match all documents that have 'US' or 'CA' in their market
	//  metadata values and 'agent' in their user metadata values will be
	//  ```
	//  documents_metadata_filters {
	//    key: "market"
	//    value: "US,CA"
	//  }
	//  documents_metadata_filters {
	//    key: "user"
	//    value: "agent"
	//  }
	//  ```
	// +kcc:proto:field=google.cloud.dialogflow.v2.Participant.documents_metadata_filters
	DocumentsMetadataFilters map[string]string `json:"documentsMetadataFilters,omitempty"`
}
