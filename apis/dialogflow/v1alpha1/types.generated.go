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


// +kcc:proto=google.cloud.dialogflow.v2.AnnotatedMessagePart
type AnnotatedMessagePart struct {
	// A part of a message possibly annotated with an entity.
	// +kcc:proto:field=google.cloud.dialogflow.v2.AnnotatedMessagePart.text
	Text *string `json:"text,omitempty"`

	// The [Dialogflow system entity
	//  type](https://cloud.google.com/dialogflow/docs/reference/system-entities)
	//  of this message part. If this is empty, Dialogflow could not annotate the
	//  phrase part with a system entity.
	// +kcc:proto:field=google.cloud.dialogflow.v2.AnnotatedMessagePart.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// The [Dialogflow system entity formatted value
	//  ](https://cloud.google.com/dialogflow/docs/reference/system-entities) of
	//  this message part. For example for a system entity of type
	//  `@sys.unit-currency`, this may contain:
	//  <pre>
	//  {
	//    "amount": 5,
	//    "currency": "USD"
	//  }
	//  </pre>
	// +kcc:proto:field=google.cloud.dialogflow.v2.AnnotatedMessagePart.formatted_value
	FormattedValue *Value `json:"formattedValue,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Message
type Message struct {
	// Optional. The unique identifier of the message.
	//  Format: `projects/<Project ID>/locations/<Location
	//  ID>/conversations/<Conversation ID>/messages/<Message ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.name
	Name *string `json:"name,omitempty"`

	// Required. The message content.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.content
	Content *string `json:"content,omitempty"`

	// Optional. The message language.
	//  This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt)
	//  language tag. Example: "en-US".
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. The time when the message was sent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.send_time
	SendTime *string `json:"sendTime,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.MessageAnnotation
type MessageAnnotation struct {
	// The collection of annotated message parts ordered by their
	//  position in the message. You can recover the annotated message by
	//  concatenating [AnnotatedMessagePart.text].
	// +kcc:proto:field=google.cloud.dialogflow.v2.MessageAnnotation.parts
	Parts []AnnotatedMessagePart `json:"parts,omitempty"`

	// Indicates whether the text message contains entities.
	// +kcc:proto:field=google.cloud.dialogflow.v2.MessageAnnotation.contain_entities
	ContainEntities *bool `json:"containEntities,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Sentiment
type Sentiment struct {
	// Sentiment score between -1.0 (negative sentiment) and 1.0 (positive
	//  sentiment).
	// +kcc:proto:field=google.cloud.dialogflow.v2.Sentiment.score
	Score *float32 `json:"score,omitempty"`

	// A non-negative number in the [0, +inf) range, which represents the absolute
	//  magnitude of sentiment, regardless of score (positive or negative).
	// +kcc:proto:field=google.cloud.dialogflow.v2.Sentiment.magnitude
	Magnitude *float32 `json:"magnitude,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.SentimentAnalysisResult
type SentimentAnalysisResult struct {
	// The sentiment analysis result for `query_text`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.SentimentAnalysisResult.query_text_sentiment
	QueryTextSentiment *Sentiment `json:"queryTextSentiment,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Message
type MessageObservedState struct {
	// Output only. The participant that sends this message.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.participant
	Participant *string `json:"participant,omitempty"`

	// Output only. The role of the participant.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.participant_role
	ParticipantRole *string `json:"participantRole,omitempty"`

	// Output only. The time when the message was created in Contact Center AI.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The annotation for the message.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.message_annotation
	MessageAnnotation *MessageAnnotation `json:"messageAnnotation,omitempty"`

	// Output only. The sentiment analysis result for the message.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Message.sentiment_analysis
	SentimentAnalysis *SentimentAnalysisResult `json:"sentimentAnalysis,omitempty"`
}
