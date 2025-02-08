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


// +kcc:proto=google.cloud.dialogflow.v2beta1.Context
type Context struct {
	// Required. The unique identifier of the context. Supported formats:
	//  - `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context
	//    ID>`,
	//  - `projects/<Project ID>/locations/<Location ID>/agent/sessions/<Session
	//    ID>/contexts/<Context ID>`,
	//  - `projects/<Project ID>/agent/environments/<Environment ID>/users/<User
	//    ID>/sessions/<Session ID>/contexts/<Context ID>`,
	//  - `projects/<Project ID>/locations/<Location
	//    ID>/agent/environments/<Environment ID>/users/<User ID>/sessions/<Session
	//    ID>/contexts/<Context ID>`,
	//
	//  The `Context ID` is always converted to lowercase, may only contain
	//  characters in `a-zA-Z0-9_-%` and may be at most 250 bytes long.
	//
	//  If `Environment ID` is not specified, we assume default 'draft'
	//  environment. If `User ID` is not specified, we assume default '-' user.
	//
	//  The following context names are reserved for internal use by Dialogflow.
	//  You should not use these contexts or create contexts with these names:
	//
	//  * `__system_counters__`
	//  * `*_id_dialog_context`
	//  * `*_dialog_params_size`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Context.name
	Name *string `json:"name,omitempty"`

	// Optional. The number of conversational query requests after which the
	//  context expires. The default is `0`. If set to `0`, the context expires
	//  immediately. Contexts expire automatically after 20 minutes if there
	//  are no matching queries.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Context.lifespan_count
	LifespanCount *int32 `json:"lifespanCount,omitempty"`

	// Optional. The collection of parameters associated with this context.
	//
	//  Depending on your protocol or client library language, this is a
	//  map, associative array, symbol table, dictionary, or JSON object
	//  composed of a collection of (MapKey, MapValue) pairs:
	//
	//  * MapKey type: string
	//  * MapKey value: parameter name
	//  * MapValue type: If parameter's entity type is a composite entity then use
	//  map, otherwise, depending on the parameter value type, it could be one of
	//  string, number, boolean, null, list or map.
	//  * MapValue value: If parameter's entity type is a composite entity then use
	//  map from composite entity property names to property values, otherwise,
	//  use parameter value.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Context.parameters
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent
type Intent struct {
	// Optional. The unique identifier of this intent.
	//  Required for
	//  [Intents.UpdateIntent][google.cloud.dialogflow.v2beta1.Intents.UpdateIntent]
	//  and
	//  [Intents.BatchUpdateIntents][google.cloud.dialogflow.v2beta1.Intents.BatchUpdateIntents]
	//  methods.
	//  Supported formats:
	//
	//  - `projects/<Project ID>/agent/intents/<Intent ID>`
	//  - `projects/<Project ID>/locations/<Location ID>/agent/intents/<Intent ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.name
	Name *string `json:"name,omitempty"`

	// Required. The name of this intent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Indicates whether webhooks are enabled for the intent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.webhook_state
	WebhookState *string `json:"webhookState,omitempty"`

	// Optional. The priority of this intent. Higher numbers represent higher
	//  priorities.
	//
	//  - If the supplied value is unspecified or 0, the service
	//    translates the value to 500,000, which corresponds to the
	//    `Normal` priority in the console.
	//  - If the supplied value is negative, the intent is ignored
	//    in runtime detect intent requests.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.priority
	Priority *int32 `json:"priority,omitempty"`

	// Optional. Indicates whether this is a fallback intent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.is_fallback
	IsFallback *bool `json:"isFallback,omitempty"`

	// Optional. Indicates whether Machine Learning is enabled for the intent.
	//  Note: If `ml_enabled` setting is set to false, then this intent is not
	//  taken into account during inference in `ML ONLY` match mode. Also,
	//  auto-markup in the UI is turned off.
	//  DEPRECATED! Please use `ml_disabled` field instead.
	//  NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false,
	//  then the default value is determined as follows:
	//
	//  - Before April 15th, 2018 the default is:
	//    ml_enabled = false / ml_disabled = true.
	//  - After April 15th, 2018 the default is:
	//    ml_enabled = true / ml_disabled = false.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.ml_enabled
	MlEnabled *bool `json:"mlEnabled,omitempty"`

	// Optional. Indicates whether Machine Learning is disabled for the intent.
	//  Note: If `ml_disabled` setting is set to true, then this intent is not
	//  taken into account during inference in `ML ONLY` match mode. Also,
	//  auto-markup in the UI is turned off.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.ml_disabled
	MlDisabled *bool `json:"mlDisabled,omitempty"`

	// Optional. Indicates that a live agent should be brought in to handle the
	//  interaction with the user. In most cases, when you set this flag to true,
	//  you would also want to set end_interaction to true as well. Default is
	//  false.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.live_agent_handoff
	LiveAgentHandoff *bool `json:"liveAgentHandoff,omitempty"`

	// Optional. Indicates that this intent ends an interaction. Some integrations
	//  (e.g., Actions on Google or Dialogflow phone gateway) use this information
	//  to close interaction with an end user. Default is false.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.end_interaction
	EndInteraction *bool `json:"endInteraction,omitempty"`

	// Optional. The list of context names required for this intent to be
	//  triggered.
	//  Formats:
	//
	//  - `projects/<Project ID>/agent/sessions/-/contexts/<Context ID>`
	//  - `projects/<Project ID>/locations/<Location
	//    ID>/agent/sessions/-/contexts/<Context ID>`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.input_context_names
	InputContextNames []string `json:"inputContextNames,omitempty"`

	// Optional. The collection of event names that trigger the intent.
	//  If the collection of input contexts is not empty, all of the contexts must
	//  be present in the active user session for an event to trigger this intent.
	//  Event names are limited to 150 characters.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.events
	Events []string `json:"events,omitempty"`

	// Optional. The collection of examples that the agent is
	//  trained on.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.training_phrases
	TrainingPhrases []Intent_TrainingPhrase `json:"trainingPhrases,omitempty"`

	// Optional. The name of the action associated with the intent.
	//  Note: The action name must not contain whitespaces.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.action
	Action *string `json:"action,omitempty"`

	// Optional. The collection of contexts that are activated when the intent
	//  is matched. Context messages in this collection should not set the
	//  parameters field. Setting the `lifespan_count` to 0 will reset the context
	//  when the intent is matched.
	//  Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.output_contexts
	OutputContexts []Context `json:"outputContexts,omitempty"`

	// Optional. Indicates whether to delete all contexts in the current
	//  session when this intent is matched.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.reset_contexts
	ResetContexts *bool `json:"resetContexts,omitempty"`

	// Optional. The collection of parameters associated with the intent.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.parameters
	Parameters []Intent_Parameter `json:"parameters,omitempty"`

	// Optional. The collection of rich messages corresponding to the
	//  `Response` field in the Dialogflow console.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.messages
	Messages []Intent_Message `json:"messages,omitempty"`

	// Optional. The list of platforms for which the first responses will be
	//  copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.default_response_platforms
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty"`

	// Optional. The unique identifier of the parent intent in the
	//  chain of followup intents. You can set this field when creating an intent,
	//  for example with
	//  [CreateIntent][google.cloud.dialogflow.v2beta1.Intents.CreateIntent] or
	//  [BatchUpdateIntents][google.cloud.dialogflow.v2beta1.Intents.BatchUpdateIntents],
	//  in order to make this intent a followup intent.
	//
	//  It identifies the parent followup intent.
	//  Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.parent_followup_intent_name
	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.FollowupIntentInfo
type Intent_FollowupIntentInfo struct {
	// The unique identifier of the followup intent.
	//  Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.FollowupIntentInfo.followup_intent_name
	FollowupIntentName *string `json:"followupIntentName,omitempty"`

	// The unique identifier of the followup intent's parent.
	//  Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.FollowupIntentInfo.parent_followup_intent_name
	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message
type Intent_Message struct {
	// Returns a text response.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.text
	Text *Intent_Message_Text `json:"text,omitempty"`

	// Displays an image.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.image
	Image *Intent_Message_Image `json:"image,omitempty"`

	// Displays quick replies.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.quick_replies
	QuickReplies *Intent_Message_QuickReplies `json:"quickReplies,omitempty"`

	// Displays a card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.card
	Card *Intent_Message_Card `json:"card,omitempty"`

	// A custom platform-specific response.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.payload
	Payload map[string]string `json:"payload,omitempty"`

	// Returns a voice or text-only response for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.simple_responses
	SimpleResponses *Intent_Message_SimpleResponses `json:"simpleResponses,omitempty"`

	// Displays a basic card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.basic_card
	BasicCard *Intent_Message_BasicCard `json:"basicCard,omitempty"`

	// Displays suggestion chips for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.suggestions
	Suggestions *Intent_Message_Suggestions `json:"suggestions,omitempty"`

	// Displays a link out suggestion chip for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.link_out_suggestion
	LinkOutSuggestion *Intent_Message_LinkOutSuggestion `json:"linkOutSuggestion,omitempty"`

	// Displays a list card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.list_select
	ListSelect *Intent_Message_ListSelect `json:"listSelect,omitempty"`

	// Displays a carousel card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.carousel_select
	CarouselSelect *Intent_Message_CarouselSelect `json:"carouselSelect,omitempty"`

	// Plays audio from a file in Telephony Gateway.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.telephony_play_audio
	TelephonyPlayAudio *Intent_Message_TelephonyPlayAudio `json:"telephonyPlayAudio,omitempty"`

	// Synthesizes speech in Telephony Gateway.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.telephony_synthesize_speech
	TelephonySynthesizeSpeech *Intent_Message_TelephonySynthesizeSpeech `json:"telephonySynthesizeSpeech,omitempty"`

	// Transfers the call in Telephony Gateway.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.telephony_transfer_call
	TelephonyTransferCall *Intent_Message_TelephonyTransferCall `json:"telephonyTransferCall,omitempty"`

	// Rich Business Messaging (RBM) text response.
	//
	//  RBM allows businesses to send enriched and branded versions of SMS. See
	//  https://jibe.google.com/business-messaging.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.rbm_text
	RbmText *Intent_Message_RbmText `json:"rbmText,omitempty"`

	// Standalone Rich Business Messaging (RBM) rich card response.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.rbm_standalone_rich_card
	RbmStandaloneRichCard *Intent_Message_RbmStandaloneCard `json:"rbmStandaloneRichCard,omitempty"`

	// Rich Business Messaging (RBM) carousel rich card response.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.rbm_carousel_rich_card
	RbmCarouselRichCard *Intent_Message_RbmCarouselCard `json:"rbmCarouselRichCard,omitempty"`

	// Browse carousel card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.browse_carousel_card
	BrowseCarouselCard *Intent_Message_BrowseCarouselCard `json:"browseCarouselCard,omitempty"`

	// Table card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.table_card
	TableCard *Intent_Message_TableCard `json:"tableCard,omitempty"`

	// The media content card for Actions on Google.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.media_content
	MediaContent *Intent_Message_MediaContent `json:"mediaContent,omitempty"`

	// Optional. The platform that this message is intended for.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.platform
	Platform *string `json:"platform,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard
type Intent_Message_BasicCard struct {
	// Optional. The title of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.title
	Title *string `json:"title,omitempty"`

	// Optional. The subtitle of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.subtitle
	Subtitle *string `json:"subtitle,omitempty"`

	// Required, unless image is present. The body text of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.formatted_text
	FormattedText *string `json:"formattedText,omitempty"`

	// Optional. The image for the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.image
	Image *Intent_Message_Image `json:"image,omitempty"`

	// Optional. The collection of card buttons.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.buttons
	Buttons []Intent_Message_BasicCard_Button `json:"buttons,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.Button
type Intent_Message_BasicCard_Button struct {
	// Required. The title of the button.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.Button.title
	Title *string `json:"title,omitempty"`

	// Required. Action to take when a user taps on the button.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.Button.open_uri_action
	OpenURIAction *Intent_Message_BasicCard_Button_OpenUriAction `json:"openURIAction,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.Button.OpenUriAction
type Intent_Message_BasicCard_Button_OpenUriAction struct {
	// Required. The HTTP or HTTPS scheme URI.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BasicCard.Button.OpenUriAction.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard
type Intent_Message_BrowseCarouselCard struct {
	// Required. List of items in the Browse Carousel Card. Minimum of two
	//  items, maximum of ten.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.items
	Items []Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem `json:"items,omitempty"`

	// Optional. Settings for displaying the image. Applies to every image in
	//  [items][google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.items].
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.image_display_options
	ImageDisplayOptions *string `json:"imageDisplayOptions,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem
type Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem struct {
	// Required. Action to present to the user.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.open_uri_action
	OpenURIAction *Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction `json:"openURIAction,omitempty"`

	// Required. Title of the carousel item. Maximum of two lines of text.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.title
	Title *string `json:"title,omitempty"`

	// Optional. Description of the carousel item. Maximum of four lines of
	//  text.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.description
	Description *string `json:"description,omitempty"`

	// Optional. Hero image for the carousel item.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.image
	Image *Intent_Message_Image `json:"image,omitempty"`

	// Optional. Text that appears at the bottom of the Browse Carousel
	//  Card. Maximum of one line of text.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.footer
	Footer *string `json:"footer,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.OpenUrlAction
type Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction struct {
	// Required. URL
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.OpenUrlAction.url
	URL *string `json:"url,omitempty"`

	// Optional. Specifies the type of viewer that is used when opening
	//  the URL. Defaults to opening via web browser.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.BrowseCarouselCard.BrowseCarouselCardItem.OpenUrlAction.url_type_hint
	URLTypeHint *string `json:"urlTypeHint,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Card
type Intent_Message_Card struct {
	// Optional. The title of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.title
	Title *string `json:"title,omitempty"`

	// Optional. The subtitle of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.subtitle
	Subtitle *string `json:"subtitle,omitempty"`

	// Optional. The public URI to an image file for the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// Optional. The collection of card buttons.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.buttons
	Buttons []Intent_Message_Card_Button `json:"buttons,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Card.Button
type Intent_Message_Card_Button struct {
	// Optional. The text to show on the button.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.Button.text
	Text *string `json:"text,omitempty"`

	// Optional. The text to send back to the Dialogflow API or a URI to
	//  open.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Card.Button.postback
	Postback *string `json:"postback,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect
type Intent_Message_CarouselSelect struct {
	// Required. Carousel items.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.items
	Items []Intent_Message_CarouselSelect_Item `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.Item
type Intent_Message_CarouselSelect_Item struct {
	// Required. Additional info about the option item.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.Item.info
	Info *Intent_Message_SelectItemInfo `json:"info,omitempty"`

	// Required. Title of the carousel item.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.Item.title
	Title *string `json:"title,omitempty"`

	// Optional. The body text of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.Item.description
	Description *string `json:"description,omitempty"`

	// Optional. The image to display.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.CarouselSelect.Item.image
	Image *Intent_Message_Image `json:"image,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.ColumnProperties
type Intent_Message_ColumnProperties struct {
	// Required. Column heading.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ColumnProperties.header
	Header *string `json:"header,omitempty"`

	// Optional. Defines text alignment for all cells in this column.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ColumnProperties.horizontal_alignment
	HorizontalAlignment *string `json:"horizontalAlignment,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Image
type Intent_Message_Image struct {
	// Optional. The public URI to an image file.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Image.image_uri
	ImageURI *string `json:"imageURI,omitempty"`

	// A text description of the image to be used for accessibility,
	//  e.g., screen readers. Required if image_uri is set for CarouselSelect.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Image.accessibility_text
	AccessibilityText *string `json:"accessibilityText,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.LinkOutSuggestion
type Intent_Message_LinkOutSuggestion struct {
	// Required. The name of the app or site this chip is linking to.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.LinkOutSuggestion.destination_name
	DestinationName *string `json:"destinationName,omitempty"`

	// Required. The URI of the app or site to open when the user taps the
	//  suggestion chip.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.LinkOutSuggestion.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect
type Intent_Message_ListSelect struct {
	// Optional. The overall title of the list.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.title
	Title *string `json:"title,omitempty"`

	// Required. List items.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.items
	Items []Intent_Message_ListSelect_Item `json:"items,omitempty"`

	// Optional. Subtitle of the list.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.subtitle
	Subtitle *string `json:"subtitle,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.Item
type Intent_Message_ListSelect_Item struct {
	// Required. Additional information about this option.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.Item.info
	Info *Intent_Message_SelectItemInfo `json:"info,omitempty"`

	// Required. The title of the list item.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.Item.title
	Title *string `json:"title,omitempty"`

	// Optional. The main text describing the item.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.Item.description
	Description *string `json:"description,omitempty"`

	// Optional. The image to display.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.ListSelect.Item.image
	Image *Intent_Message_Image `json:"image,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent
type Intent_Message_MediaContent struct {
	// Optional. What type of media is the content (ie "audio").
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.media_type
	MediaType *string `json:"mediaType,omitempty"`

	// Required. List of media objects.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.media_objects
	MediaObjects []Intent_Message_MediaContent_ResponseMediaObject `json:"mediaObjects,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject
type Intent_Message_MediaContent_ResponseMediaObject struct {
	// Required. Name of media card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject.name
	Name *string `json:"name,omitempty"`

	// Optional. Description of media card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject.description
	Description *string `json:"description,omitempty"`

	// Optional. Image to display above media content.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject.large_image
	LargeImage *Intent_Message_Image `json:"largeImage,omitempty"`

	// Optional. Icon to display above media content.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject.icon
	Icon *Intent_Message_Image `json:"icon,omitempty"`

	// Required. Url where the media is stored.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.MediaContent.ResponseMediaObject.content_url
	ContentURL *string `json:"contentURL,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.QuickReplies
type Intent_Message_QuickReplies struct {
	// Optional. The title of the collection of quick replies.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.QuickReplies.title
	Title *string `json:"title,omitempty"`

	// Optional. The collection of quick replies.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.QuickReplies.quick_replies
	QuickReplies []string `json:"quickReplies,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent
type Intent_Message_RbmCardContent struct {
	// Optional. Title of the card (at most 200 bytes).
	//
	//  At least one of the title, description or media must be set.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.title
	Title *string `json:"title,omitempty"`

	// Optional. Description of the card (at most 2000 bytes).
	//
	//  At least one of the title, description or media must be set.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.description
	Description *string `json:"description,omitempty"`

	// Optional. However at least one of the title, description or media must
	//  be set. Media (image, GIF or a video) to include in the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.media
	Media *Intent_Message_RbmCardContent_RbmMedia `json:"media,omitempty"`

	// Optional. List of suggestions to include in the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.suggestions
	Suggestions []Intent_Message_RbmSuggestion `json:"suggestions,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.RbmMedia
type Intent_Message_RbmCardContent_RbmMedia struct {
	// Required. Publicly reachable URI of the file. The RBM platform
	//  determines the MIME type of the file from the content-type field in
	//  the HTTP headers when the platform fetches the file. The content-type
	//  field must be present and accurate in the HTTP response from the URL.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.RbmMedia.file_uri
	FileURI *string `json:"fileURI,omitempty"`

	// Optional. Publicly reachable URI of the thumbnail.If you don't
	//  provide a thumbnail URI, the RBM platform displays a blank
	//  placeholder thumbnail until the user's device downloads the file.
	//  Depending on the user's setting, the file may not download
	//  automatically and may require the user to tap a download button.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.RbmMedia.thumbnail_uri
	ThumbnailURI *string `json:"thumbnailURI,omitempty"`

	// Required for cards with vertical orientation. The height of the media
	//  within a rich card with a vertical layout.
	//  For a standalone card with horizontal layout, height is not
	//  customizable, and this field is ignored.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCardContent.RbmMedia.height
	Height *string `json:"height,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCarouselCard
type Intent_Message_RbmCarouselCard struct {
	// Required. The width of the cards in the carousel.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCarouselCard.card_width
	CardWidth *string `json:"cardWidth,omitempty"`

	// Required. The cards in the carousel. A carousel must have at least
	//  2 cards and at most 10.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmCarouselCard.card_contents
	CardContents []Intent_Message_RbmCardContent `json:"cardContents,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmStandaloneCard
type Intent_Message_RbmStandaloneCard struct {
	// Required. Orientation of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmStandaloneCard.card_orientation
	CardOrientation *string `json:"cardOrientation,omitempty"`

	// Required if orientation is horizontal.
	//  Image preview alignment for standalone cards with horizontal layout.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmStandaloneCard.thumbnail_image_alignment
	ThumbnailImageAlignment *string `json:"thumbnailImageAlignment,omitempty"`

	// Required. Card content.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmStandaloneCard.card_content
	CardContent *Intent_Message_RbmCardContent `json:"cardContent,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction
type Intent_Message_RbmSuggestedAction struct {
	// Text to display alongside the action.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.text
	Text *string `json:"text,omitempty"`

	// Opaque payload that the Dialogflow receives in a user event
	//  when the user taps the suggested action. This data will be also
	//  forwarded to webhook to allow performing custom business logic.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.postback_data
	PostbackData *string `json:"postbackData,omitempty"`

	// Suggested client side action: Dial a phone number
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.dial
	Dial *Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial `json:"dial,omitempty"`

	// Suggested client side action: Open a URI on device
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.open_url
	OpenURL *Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri `json:"openURL,omitempty"`

	// Suggested client side action: Share user location
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.share_location
	ShareLocation *Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation `json:"shareLocation,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.RbmSuggestedActionDial
type Intent_Message_RbmSuggestedAction_RbmSuggestedActionDial struct {
	// Required. The phone number to fill in the default dialer app.
	//  This field should be in [E.164](https://en.wikipedia.org/wiki/E.164)
	//  format. An example of a correctly formatted phone number:
	//  +15556767888.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.RbmSuggestedActionDial.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.RbmSuggestedActionOpenUri
type Intent_Message_RbmSuggestedAction_RbmSuggestedActionOpenUri struct {
	// Required. The uri to open on the user device
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.RbmSuggestedActionOpenUri.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedAction.RbmSuggestedActionShareLocation
type Intent_Message_RbmSuggestedAction_RbmSuggestedActionShareLocation struct {
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedReply
type Intent_Message_RbmSuggestedReply struct {
	// Suggested reply text.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedReply.text
	Text *string `json:"text,omitempty"`

	// Opaque payload that the Dialogflow receives in a user event
	//  when the user taps the suggested reply. This data will be also
	//  forwarded to webhook to allow performing custom business logic.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestedReply.postback_data
	PostbackData *string `json:"postbackData,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestion
type Intent_Message_RbmSuggestion struct {
	// Predefined replies for user to select instead of typing
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestion.reply
	Reply *Intent_Message_RbmSuggestedReply `json:"reply,omitempty"`

	// Predefined client side actions that user can choose
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmSuggestion.action
	Action *Intent_Message_RbmSuggestedAction `json:"action,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.RbmText
type Intent_Message_RbmText struct {
	// Required. Text sent and displayed to the user.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmText.text
	Text *string `json:"text,omitempty"`

	// Optional. One or more suggestions to show to the user.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.RbmText.rbm_suggestion
	RbmSuggestion []Intent_Message_RbmSuggestion `json:"rbmSuggestion,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.SelectItemInfo
type Intent_Message_SelectItemInfo struct {
	// Required. A unique key that will be sent back to the agent if this
	//  response is given.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SelectItemInfo.key
	Key *string `json:"key,omitempty"`

	// Optional. A list of synonyms that can also be used to trigger this
	//  item in dialog.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SelectItemInfo.synonyms
	Synonyms []string `json:"synonyms,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponse
type Intent_Message_SimpleResponse struct {
	// One of text_to_speech or ssml must be provided. The plain text of the
	//  speech output. Mutually exclusive with ssml.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponse.text_to_speech
	TextToSpeech *string `json:"textToSpeech,omitempty"`

	// One of text_to_speech or ssml must be provided. Structured spoken
	//  response to the user in the SSML format. Mutually exclusive with
	//  text_to_speech.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponse.ssml
	Ssml *string `json:"ssml,omitempty"`

	// Optional. The text to display.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponse.display_text
	DisplayText *string `json:"displayText,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponses
type Intent_Message_SimpleResponses struct {
	// Required. The list of simple responses.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.SimpleResponses.simple_responses
	SimpleResponses []Intent_Message_SimpleResponse `json:"simpleResponses,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Suggestion
type Intent_Message_Suggestion struct {
	// Required. The text shown the in the suggestion chip.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Suggestion.title
	Title *string `json:"title,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Suggestions
type Intent_Message_Suggestions struct {
	// Required. The list of suggested replies.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Suggestions.suggestions
	Suggestions []Intent_Message_Suggestion `json:"suggestions,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard
type Intent_Message_TableCard struct {
	// Required. Title of the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.title
	Title *string `json:"title,omitempty"`

	// Optional. Subtitle to the title.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.subtitle
	Subtitle *string `json:"subtitle,omitempty"`

	// Optional. Image which should be displayed on the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.image
	Image *Intent_Message_Image `json:"image,omitempty"`

	// Optional. Display properties for the columns in this table.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.column_properties
	ColumnProperties []Intent_Message_ColumnProperties `json:"columnProperties,omitempty"`

	// Optional. Rows in this table of data.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.rows
	Rows []Intent_Message_TableCardRow `json:"rows,omitempty"`

	// Optional. List of buttons for the card.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCard.buttons
	Buttons []Intent_Message_BasicCard_Button `json:"buttons,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TableCardCell
type Intent_Message_TableCardCell struct {
	// Required. Text in this cell.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCardCell.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TableCardRow
type Intent_Message_TableCardRow struct {
	// Optional. List of cells that make up this row.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCardRow.cells
	Cells []Intent_Message_TableCardCell `json:"cells,omitempty"`

	// Optional. Whether to add a visual divider after this row.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TableCardRow.divider_after
	DividerAfter *bool `json:"dividerAfter,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonyPlayAudio
type Intent_Message_TelephonyPlayAudio struct {
	// Required. URI to a Google Cloud Storage object containing the audio to
	//  play, e.g., "gs://bucket/object". The object must contain a single
	//  channel (mono) of linear PCM audio (2 bytes / sample) at 8kHz.
	//
	//  This object must be readable by the `service-<Project
	//  Number>@gcp-sa-dialogflow.iam.gserviceaccount.com` service account
	//  where <Project Number> is the number of the Telephony Gateway project
	//  (usually the same as the Dialogflow agent project). If the Google Cloud
	//  Storage bucket is in the Telephony Gateway project, this permission is
	//  added by default when enabling the Dialogflow V2 API.
	//
	//  For audio from other sources, consider using the
	//  `TelephonySynthesizeSpeech` message with SSML.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonyPlayAudio.audio_uri
	AudioURI *string `json:"audioURI,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonySynthesizeSpeech
type Intent_Message_TelephonySynthesizeSpeech struct {
	// The raw text to be synthesized.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonySynthesizeSpeech.text
	Text *string `json:"text,omitempty"`

	// The SSML to be synthesized. For more information, see
	//  [SSML](https://developers.google.com/actions/reference/ssml).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonySynthesizeSpeech.ssml
	Ssml *string `json:"ssml,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonyTransferCall
type Intent_Message_TelephonyTransferCall struct {
	// Required. The phone number to transfer the call to
	//  in [E.164 format](https://en.wikipedia.org/wiki/E.164).
	//
	//  We currently only allow transferring to US numbers (+1xxxyyyzzzz).
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.TelephonyTransferCall.phone_number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Message.Text
type Intent_Message_Text struct {
	// Optional. The collection of the agent's responses.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Message.Text.text
	Text []string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.Parameter
type Intent_Parameter struct {
	// The unique identifier of this parameter.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The definition of the parameter value. It can be:
	//
	//  - a constant string,
	//  - a parameter value defined as `$parameter_name`,
	//  - an original parameter value defined as `$parameter_name.original`,
	//  - a parameter value from some context defined as
	//    `#context_name.parameter_name`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.value
	Value *string `json:"value,omitempty"`

	// Optional. The default value to use when the `value` yields an empty
	//  result.
	//  Default values can be extracted from contexts by using the following
	//  syntax: `#context_name.parameter_name`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.default_value
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Optional. The name of the entity type, prefixed with `@`, that
	//  describes values of the parameter. If the parameter is
	//  required, this must be provided.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.entity_type_display_name
	EntityTypeDisplayName *string `json:"entityTypeDisplayName,omitempty"`

	// Optional. Indicates whether the parameter is required. That is,
	//  whether the intent cannot be completed without collecting the parameter
	//  value.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.mandatory
	Mandatory *bool `json:"mandatory,omitempty"`

	// Optional. The collection of prompts that the agent can present to the
	//  user in order to collect a value for the parameter.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.prompts
	Prompts []string `json:"prompts,omitempty"`

	// Optional. Indicates whether the parameter represents a list of values.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.Parameter.is_list
	IsList *bool `json:"isList,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase
type Intent_TrainingPhrase struct {

	// Required. The type of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.type
	Type *string `json:"type,omitempty"`

	// Required. The ordered list of training phrase parts.
	//  The parts are concatenated in order to form the training phrase.
	//
	//  Note: The API does not automatically annotate training phrases like the
	//  Dialogflow Console does.
	//
	//  Note: Do not forget to include whitespace at part boundaries,
	//  so the training phrase is well formatted when the parts are concatenated.
	//
	//  If the training phrase does not need to be annotated with parameters,
	//  you just need a single part with only the
	//  [Part.text][google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part.text]
	//  field set.
	//
	//  If you want to annotate the training phrase, you must create multiple
	//  parts, where the fields of each part are populated in one of two ways:
	//
	//  -   `Part.text` is set to a part of the phrase that has no parameters.
	//  -   `Part.text` is set to a part of the phrase that you want to annotate,
	//      and the `entity_type`, `alias`, and `user_defined` fields are all
	//      set.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.parts
	Parts []Intent_TrainingPhrase_Part `json:"parts,omitempty"`

	// Optional. Indicates how many times this example was added to
	//  the intent. Each time a developer adds an existing sample by editing an
	//  intent or training, this counter is increased.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.times_added_count
	TimesAddedCount *int32 `json:"timesAddedCount,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part
type Intent_TrainingPhrase_Part struct {
	// Required. The text for this part.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part.text
	Text *string `json:"text,omitempty"`

	// Optional. The entity type name prefixed with `@`.
	//  This field is required for annotated parts of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part.entity_type
	EntityType *string `json:"entityType,omitempty"`

	// Optional. The parameter name for the value extracted from the
	//  annotated part of the example.
	//  This field is required for annotated parts of the training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part.alias
	Alias *string `json:"alias,omitempty"`

	// Optional. Indicates whether the text was manually annotated.
	//  This field is set to true when the Dialogflow Console is used to
	//  manually annotate the part. When creating an annotated part with the
	//  API, you must set this to true.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.Part.user_defined
	UserDefined *bool `json:"userDefined,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent
type IntentObservedState struct {
	// Optional. The collection of examples that the agent is
	//  trained on.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.training_phrases
	TrainingPhrases []Intent_TrainingPhraseObservedState `json:"trainingPhrases,omitempty"`

	// Output only. The unique identifier of the root intent in the chain of
	//  followup intents. It identifies the correct followup intents chain for
	//  this intent.
	//
	//  Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.root_followup_intent_name
	RootFollowupIntentName *string `json:"rootFollowupIntentName,omitempty"`

	// Output only. Information about all followup intents that have this intent
	//  as a direct or indirect parent. We populate this field only in the output.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.followup_intent_info
	FollowupIntentInfo []Intent_FollowupIntentInfo `json:"followupIntentInfo,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase
type Intent_TrainingPhraseObservedState struct {
	// Output only. The unique identifier of this training phrase.
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.Intent.TrainingPhrase.name
	Name *string `json:"name,omitempty"`
}
