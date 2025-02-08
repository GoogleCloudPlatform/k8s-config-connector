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
func Context_FromProto(mapCtx *direct.MapContext, in *pb.Context) *krm.Context {
	if in == nil {
		return nil
	}
	out := &krm.Context{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LifespanCount = direct.LazyPtr(in.GetLifespanCount())
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func Context_ToProto(mapCtx *direct.MapContext, in *krm.Context) *pb.Context {
	if in == nil {
		return nil
	}
	out := &pb.Context{}
	out.Name = direct.ValueOf(in.Name)
	out.LifespanCount = direct.ValueOf(in.LifespanCount)
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	return out
}
func Intent_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.Intent {
	if in == nil {
		return nil
	}
	out := &krm.Intent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.WebhookState = direct.Enum_FromProto(mapCtx, in.GetWebhookState())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.MlDisabled = direct.LazyPtr(in.GetMlDisabled())
	out.LiveAgentHandoff = direct.LazyPtr(in.GetLiveAgentHandoff())
	out.EndInteraction = direct.LazyPtr(in.GetEndInteraction())
	out.InputContextNames = in.InputContextNames
	out.Events = in.Events
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_FromProto)
	out.Action = direct.LazyPtr(in.GetAction())
	out.OutputContexts = direct.Slice_FromProto(mapCtx, in.OutputContexts, Context_FromProto)
	out.ResetContexts = direct.LazyPtr(in.GetResetContexts())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Intent_Parameter_FromProto)
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, Intent_Message_FromProto)
	out.DefaultResponsePlatforms = direct.EnumSlice_FromProto(mapCtx, in.DefaultResponsePlatforms)
	// MISSING: RootFollowupIntentName
	out.ParentFollowupIntentName = direct.LazyPtr(in.GetParentFollowupIntentName())
	// MISSING: FollowupIntentInfo
	return out
}
func Intent_ToProto(mapCtx *direct.MapContext, in *krm.Intent) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.WebhookState = direct.Enum_ToProto[pb.Intent_WebhookState](mapCtx, in.WebhookState)
	out.Priority = direct.ValueOf(in.Priority)
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.MlDisabled = direct.ValueOf(in.MlDisabled)
	out.LiveAgentHandoff = direct.ValueOf(in.LiveAgentHandoff)
	out.EndInteraction = direct.ValueOf(in.EndInteraction)
	out.InputContextNames = in.InputContextNames
	out.Events = in.Events
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_ToProto)
	out.Action = direct.ValueOf(in.Action)
	out.OutputContexts = direct.Slice_ToProto(mapCtx, in.OutputContexts, Context_ToProto)
	out.ResetContexts = direct.ValueOf(in.ResetContexts)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Intent_Parameter_ToProto)
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, Intent_Message_ToProto)
	out.DefaultResponsePlatforms = direct.EnumSlice_ToProto[pb.Intent_Message_Platform](mapCtx, in.DefaultResponsePlatforms)
	// MISSING: RootFollowupIntentName
	out.ParentFollowupIntentName = direct.ValueOf(in.ParentFollowupIntentName)
	// MISSING: FollowupIntentInfo
	return out
}
func IntentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.IntentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IntentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebhookState
	// MISSING: Priority
	// MISSING: IsFallback
	// MISSING: MlDisabled
	// MISSING: LiveAgentHandoff
	// MISSING: EndInteraction
	// MISSING: InputContextNames
	// MISSING: Events
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhraseObservedState_FromProto)
	// MISSING: Action
	// MISSING: OutputContexts
	// MISSING: ResetContexts
	// MISSING: Parameters
	// MISSING: Messages
	// MISSING: DefaultResponsePlatforms
	out.RootFollowupIntentName = direct.LazyPtr(in.GetRootFollowupIntentName())
	// MISSING: ParentFollowupIntentName
	out.FollowupIntentInfo = direct.Slice_FromProto(mapCtx, in.FollowupIntentInfo, Intent_FollowupIntentInfo_FromProto)
	return out
}
func IntentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IntentObservedState) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: WebhookState
	// MISSING: Priority
	// MISSING: IsFallback
	// MISSING: MlDisabled
	// MISSING: LiveAgentHandoff
	// MISSING: EndInteraction
	// MISSING: InputContextNames
	// MISSING: Events
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhraseObservedState_ToProto)
	// MISSING: Action
	// MISSING: OutputContexts
	// MISSING: ResetContexts
	// MISSING: Parameters
	// MISSING: Messages
	// MISSING: DefaultResponsePlatforms
	out.RootFollowupIntentName = direct.ValueOf(in.RootFollowupIntentName)
	// MISSING: ParentFollowupIntentName
	out.FollowupIntentInfo = direct.Slice_ToProto(mapCtx, in.FollowupIntentInfo, Intent_FollowupIntentInfo_ToProto)
	return out
}
func Intent_FollowupIntentInfo_FromProto(mapCtx *direct.MapContext, in *pb.Intent_FollowupIntentInfo) *krm.Intent_FollowupIntentInfo {
	if in == nil {
		return nil
	}
	out := &krm.Intent_FollowupIntentInfo{}
	out.FollowupIntentName = direct.LazyPtr(in.GetFollowupIntentName())
	out.ParentFollowupIntentName = direct.LazyPtr(in.GetParentFollowupIntentName())
	return out
}
func Intent_FollowupIntentInfo_ToProto(mapCtx *direct.MapContext, in *krm.Intent_FollowupIntentInfo) *pb.Intent_FollowupIntentInfo {
	if in == nil {
		return nil
	}
	out := &pb.Intent_FollowupIntentInfo{}
	out.FollowupIntentName = direct.ValueOf(in.FollowupIntentName)
	out.ParentFollowupIntentName = direct.ValueOf(in.ParentFollowupIntentName)
	return out
}
func Intent_Message_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message) *krm.Intent_Message {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message{}
	out.Text = Intent_Message_Text_FromProto(mapCtx, in.GetText())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.QuickReplies = Intent_Message_QuickReplies_FromProto(mapCtx, in.GetQuickReplies())
	out.Card = Intent_Message_Card_FromProto(mapCtx, in.GetCard())
	out.Payload = Payload_FromProto(mapCtx, in.GetPayload())
	out.SimpleResponses = Intent_Message_SimpleResponses_FromProto(mapCtx, in.GetSimpleResponses())
	out.BasicCard = Intent_Message_BasicCard_FromProto(mapCtx, in.GetBasicCard())
	out.Suggestions = Intent_Message_Suggestions_FromProto(mapCtx, in.GetSuggestions())
	out.LinkOutSuggestion = Intent_Message_LinkOutSuggestion_FromProto(mapCtx, in.GetLinkOutSuggestion())
	out.ListSelect = Intent_Message_ListSelect_FromProto(mapCtx, in.GetListSelect())
	out.CarouselSelect = Intent_Message_CarouselSelect_FromProto(mapCtx, in.GetCarouselSelect())
	out.BrowseCarouselCard = Intent_Message_BrowseCarouselCard_FromProto(mapCtx, in.GetBrowseCarouselCard())
	out.TableCard = Intent_Message_TableCard_FromProto(mapCtx, in.GetTableCard())
	out.MediaContent = Intent_Message_MediaContent_FromProto(mapCtx, in.GetMediaContent())
	out.Platform = direct.Enum_FromProto(mapCtx, in.GetPlatform())
	return out
}
func Intent_Message_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message) *pb.Intent_Message {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message{}
	if oneof := Intent_Message_Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Message = &pb.Intent_Message_Text_{Text: oneof}
	}
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.Image); oneof != nil {
		out.Message = &pb.Intent_Message_Image_{Image: oneof}
	}
	if oneof := Intent_Message_QuickReplies_ToProto(mapCtx, in.QuickReplies); oneof != nil {
		out.Message = &pb.Intent_Message_QuickReplies_{QuickReplies: oneof}
	}
	if oneof := Intent_Message_Card_ToProto(mapCtx, in.Card); oneof != nil {
		out.Message = &pb.Intent_Message_Card_{Card: oneof}
	}
	if oneof := Payload_ToProto(mapCtx, in.Payload); oneof != nil {
		out.Message = &pb.Intent_Message_Payload{Payload: oneof}
	}
	if oneof := Intent_Message_SimpleResponses_ToProto(mapCtx, in.SimpleResponses); oneof != nil {
		out.Message = &pb.Intent_Message_SimpleResponses_{SimpleResponses: oneof}
	}
	if oneof := Intent_Message_BasicCard_ToProto(mapCtx, in.BasicCard); oneof != nil {
		out.Message = &pb.Intent_Message_BasicCard_{BasicCard: oneof}
	}
	if oneof := Intent_Message_Suggestions_ToProto(mapCtx, in.Suggestions); oneof != nil {
		out.Message = &pb.Intent_Message_Suggestions_{Suggestions: oneof}
	}
	if oneof := Intent_Message_LinkOutSuggestion_ToProto(mapCtx, in.LinkOutSuggestion); oneof != nil {
		out.Message = &pb.Intent_Message_LinkOutSuggestion_{LinkOutSuggestion: oneof}
	}
	if oneof := Intent_Message_ListSelect_ToProto(mapCtx, in.ListSelect); oneof != nil {
		out.Message = &pb.Intent_Message_ListSelect_{ListSelect: oneof}
	}
	if oneof := Intent_Message_CarouselSelect_ToProto(mapCtx, in.CarouselSelect); oneof != nil {
		out.Message = &pb.Intent_Message_CarouselSelect_{CarouselSelect: oneof}
	}
	if oneof := Intent_Message_BrowseCarouselCard_ToProto(mapCtx, in.BrowseCarouselCard); oneof != nil {
		out.Message = &pb.Intent_Message_BrowseCarouselCard_{BrowseCarouselCard: oneof}
	}
	if oneof := Intent_Message_TableCard_ToProto(mapCtx, in.TableCard); oneof != nil {
		out.Message = &pb.Intent_Message_TableCard_{TableCard: oneof}
	}
	if oneof := Intent_Message_MediaContent_ToProto(mapCtx, in.MediaContent); oneof != nil {
		out.Message = &pb.Intent_Message_MediaContent_{MediaContent: oneof}
	}
	out.Platform = direct.Enum_ToProto[pb.Intent_Message_Platform](mapCtx, in.Platform)
	return out
}
func Intent_Message_BasicCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard) *krm.Intent_Message_BasicCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.FormattedText = direct.LazyPtr(in.GetFormattedText())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_FromProto)
	return out
}
func Intent_Message_BasicCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard) *pb.Intent_Message_BasicCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.FormattedText = direct.ValueOf(in.FormattedText)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_ToProto)
	return out
}
func Intent_Message_BasicCard_Button_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard_Button) *krm.Intent_Message_BasicCard_Button {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard_Button{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.OpenURIAction = Intent_Message_BasicCard_Button_OpenUriAction_FromProto(mapCtx, in.GetOpenUriAction())
	return out
}
func Intent_Message_BasicCard_Button_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard_Button) *pb.Intent_Message_BasicCard_Button {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard_Button{}
	out.Title = direct.ValueOf(in.Title)
	out.OpenUriAction = Intent_Message_BasicCard_Button_OpenUriAction_ToProto(mapCtx, in.OpenURIAction)
	return out
}
func Intent_Message_BasicCard_Button_OpenUriAction_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BasicCard_Button_OpenUriAction) *krm.Intent_Message_BasicCard_Button_OpenUriAction {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BasicCard_Button_OpenUriAction{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Intent_Message_BasicCard_Button_OpenUriAction_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BasicCard_Button_OpenUriAction) *pb.Intent_Message_BasicCard_Button_OpenUriAction {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BasicCard_Button_OpenUriAction{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Intent_Message_BrowseCarouselCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard) *krm.Intent_Message_BrowseCarouselCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_FromProto)
	out.ImageDisplayOptions = direct.Enum_FromProto(mapCtx, in.GetImageDisplayOptions())
	return out
}
func Intent_Message_BrowseCarouselCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard) *pb.Intent_Message_BrowseCarouselCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_ToProto)
	out.ImageDisplayOptions = direct.Enum_ToProto[pb.Intent_Message_BrowseCarouselCard_ImageDisplayOptions](mapCtx, in.ImageDisplayOptions)
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem) *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem{}
	out.OpenURIAction = Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_FromProto(mapCtx, in.GetOpenUriAction())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.Footer = direct.LazyPtr(in.GetFooter())
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem) *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem{}
	out.OpenUriAction = Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_ToProto(mapCtx, in.OpenURIAction)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.Footer = direct.ValueOf(in.Footer)
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction) *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction{}
	out.URL = direct.LazyPtr(in.GetUrl())
	out.URLTypeHint = direct.Enum_FromProto(mapCtx, in.GetUrlTypeHint())
	return out
}
func Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction) *pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction{}
	out.Url = direct.ValueOf(in.URL)
	out.UrlTypeHint = direct.Enum_ToProto[pb.Intent_Message_BrowseCarouselCard_BrowseCarouselCardItem_OpenUrlAction_UrlTypeHint](mapCtx, in.URLTypeHint)
	return out
}
func Intent_Message_Card_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Card) *krm.Intent_Message_Card {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Card{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_Card_Button_FromProto)
	return out
}
func Intent_Message_Card_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Card) *pb.Intent_Message_Card {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Card{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_Card_Button_ToProto)
	return out
}
func Intent_Message_Card_Button_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Card_Button) *krm.Intent_Message_Card_Button {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Card_Button{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Postback = direct.LazyPtr(in.GetPostback())
	return out
}
func Intent_Message_Card_Button_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Card_Button) *pb.Intent_Message_Card_Button {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Card_Button{}
	out.Text = direct.ValueOf(in.Text)
	out.Postback = direct.ValueOf(in.Postback)
	return out
}
func Intent_Message_CarouselSelect_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_CarouselSelect) *krm.Intent_Message_CarouselSelect {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_CarouselSelect{}
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_CarouselSelect_Item_FromProto)
	return out
}
func Intent_Message_CarouselSelect_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_CarouselSelect) *pb.Intent_Message_CarouselSelect {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_CarouselSelect{}
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_CarouselSelect_Item_ToProto)
	return out
}
func Intent_Message_CarouselSelect_Item_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_CarouselSelect_Item) *krm.Intent_Message_CarouselSelect_Item {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_CarouselSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_FromProto(mapCtx, in.GetInfo())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	return out
}
func Intent_Message_CarouselSelect_Item_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_CarouselSelect_Item) *pb.Intent_Message_CarouselSelect_Item {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_CarouselSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_ToProto(mapCtx, in.Info)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	return out
}
func Intent_Message_ColumnProperties_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ColumnProperties) *krm.Intent_Message_ColumnProperties {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ColumnProperties{}
	out.Header = direct.LazyPtr(in.GetHeader())
	out.HorizontalAlignment = direct.Enum_FromProto(mapCtx, in.GetHorizontalAlignment())
	return out
}
func Intent_Message_ColumnProperties_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ColumnProperties) *pb.Intent_Message_ColumnProperties {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ColumnProperties{}
	out.Header = direct.ValueOf(in.Header)
	out.HorizontalAlignment = direct.Enum_ToProto[pb.Intent_Message_ColumnProperties_HorizontalAlignment](mapCtx, in.HorizontalAlignment)
	return out
}
func Intent_Message_Image_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Image) *krm.Intent_Message_Image {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Image{}
	out.ImageURI = direct.LazyPtr(in.GetImageUri())
	out.AccessibilityText = direct.LazyPtr(in.GetAccessibilityText())
	return out
}
func Intent_Message_Image_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Image) *pb.Intent_Message_Image {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Image{}
	out.ImageUri = direct.ValueOf(in.ImageURI)
	out.AccessibilityText = direct.ValueOf(in.AccessibilityText)
	return out
}
func Intent_Message_LinkOutSuggestion_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_LinkOutSuggestion) *krm.Intent_Message_LinkOutSuggestion {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_LinkOutSuggestion{}
	out.DestinationName = direct.LazyPtr(in.GetDestinationName())
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Intent_Message_LinkOutSuggestion_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_LinkOutSuggestion) *pb.Intent_Message_LinkOutSuggestion {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_LinkOutSuggestion{}
	out.DestinationName = direct.ValueOf(in.DestinationName)
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Intent_Message_ListSelect_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ListSelect) *krm.Intent_Message_ListSelect {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ListSelect{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, Intent_Message_ListSelect_Item_FromProto)
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	return out
}
func Intent_Message_ListSelect_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ListSelect) *pb.Intent_Message_ListSelect {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ListSelect{}
	out.Title = direct.ValueOf(in.Title)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, Intent_Message_ListSelect_Item_ToProto)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	return out
}
func Intent_Message_ListSelect_Item_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_ListSelect_Item) *krm.Intent_Message_ListSelect_Item {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_ListSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_FromProto(mapCtx, in.GetInfo())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	return out
}
func Intent_Message_ListSelect_Item_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_ListSelect_Item) *pb.Intent_Message_ListSelect_Item {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_ListSelect_Item{}
	out.Info = Intent_Message_SelectItemInfo_ToProto(mapCtx, in.Info)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	return out
}
func Intent_Message_MediaContent_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_MediaContent) *krm.Intent_Message_MediaContent {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_MediaContent{}
	out.MediaType = direct.Enum_FromProto(mapCtx, in.GetMediaType())
	out.MediaObjects = direct.Slice_FromProto(mapCtx, in.MediaObjects, Intent_Message_MediaContent_ResponseMediaObject_FromProto)
	return out
}
func Intent_Message_MediaContent_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_MediaContent) *pb.Intent_Message_MediaContent {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_MediaContent{}
	out.MediaType = direct.Enum_ToProto[pb.Intent_Message_MediaContent_ResponseMediaType](mapCtx, in.MediaType)
	out.MediaObjects = direct.Slice_ToProto(mapCtx, in.MediaObjects, Intent_Message_MediaContent_ResponseMediaObject_ToProto)
	return out
}
func Intent_Message_MediaContent_ResponseMediaObject_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_MediaContent_ResponseMediaObject) *krm.Intent_Message_MediaContent_ResponseMediaObject {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_MediaContent_ResponseMediaObject{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.LargeImage = Intent_Message_Image_FromProto(mapCtx, in.GetLargeImage())
	out.Icon = Intent_Message_Image_FromProto(mapCtx, in.GetIcon())
	out.ContentURL = direct.LazyPtr(in.GetContentUrl())
	return out
}
func Intent_Message_MediaContent_ResponseMediaObject_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_MediaContent_ResponseMediaObject) *pb.Intent_Message_MediaContent_ResponseMediaObject {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_MediaContent_ResponseMediaObject{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.LargeImage); oneof != nil {
		out.Image = &pb.Intent_Message_MediaContent_ResponseMediaObject_LargeImage{LargeImage: oneof}
	}
	if oneof := Intent_Message_Image_ToProto(mapCtx, in.Icon); oneof != nil {
		out.Image = &pb.Intent_Message_MediaContent_ResponseMediaObject_Icon{Icon: oneof}
	}
	out.ContentUrl = direct.ValueOf(in.ContentURL)
	return out
}
func Intent_Message_QuickReplies_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_QuickReplies) *krm.Intent_Message_QuickReplies {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_QuickReplies{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.QuickReplies = in.QuickReplies
	return out
}
func Intent_Message_QuickReplies_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_QuickReplies) *pb.Intent_Message_QuickReplies {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_QuickReplies{}
	out.Title = direct.ValueOf(in.Title)
	out.QuickReplies = in.QuickReplies
	return out
}
func Intent_Message_SelectItemInfo_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_SelectItemInfo) *krm.Intent_Message_SelectItemInfo {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_SelectItemInfo{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Synonyms = in.Synonyms
	return out
}
func Intent_Message_SelectItemInfo_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_SelectItemInfo) *pb.Intent_Message_SelectItemInfo {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_SelectItemInfo{}
	out.Key = direct.ValueOf(in.Key)
	out.Synonyms = in.Synonyms
	return out
}
func Intent_Message_SimpleResponses_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_SimpleResponses) *krm.Intent_Message_SimpleResponses {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_SimpleResponses{}
	out.SimpleResponses = direct.Slice_FromProto(mapCtx, in.SimpleResponses, Intent_Message_SimpleResponse_FromProto)
	return out
}
func Intent_Message_SimpleResponses_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_SimpleResponses) *pb.Intent_Message_SimpleResponses {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_SimpleResponses{}
	out.SimpleResponses = direct.Slice_ToProto(mapCtx, in.SimpleResponses, Intent_Message_SimpleResponse_ToProto)
	return out
}
func Intent_Message_Suggestion_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Suggestion) *krm.Intent_Message_Suggestion {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Suggestion{}
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}
func Intent_Message_Suggestion_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Suggestion) *pb.Intent_Message_Suggestion {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Suggestion{}
	out.Title = direct.ValueOf(in.Title)
	return out
}
func Intent_Message_Suggestions_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Suggestions) *krm.Intent_Message_Suggestions {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Suggestions{}
	out.Suggestions = direct.Slice_FromProto(mapCtx, in.Suggestions, Intent_Message_Suggestion_FromProto)
	return out
}
func Intent_Message_Suggestions_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Suggestions) *pb.Intent_Message_Suggestions {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Suggestions{}
	out.Suggestions = direct.Slice_ToProto(mapCtx, in.Suggestions, Intent_Message_Suggestion_ToProto)
	return out
}
func Intent_Message_TableCard_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCard) *krm.Intent_Message_TableCard {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCard{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.Image = Intent_Message_Image_FromProto(mapCtx, in.GetImage())
	out.ColumnProperties = direct.Slice_FromProto(mapCtx, in.ColumnProperties, Intent_Message_ColumnProperties_FromProto)
	out.Rows = direct.Slice_FromProto(mapCtx, in.Rows, Intent_Message_TableCardRow_FromProto)
	out.Buttons = direct.Slice_FromProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_FromProto)
	return out
}
func Intent_Message_TableCard_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCard) *pb.Intent_Message_TableCard {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCard{}
	out.Title = direct.ValueOf(in.Title)
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.Image = Intent_Message_Image_ToProto(mapCtx, in.Image)
	out.ColumnProperties = direct.Slice_ToProto(mapCtx, in.ColumnProperties, Intent_Message_ColumnProperties_ToProto)
	out.Rows = direct.Slice_ToProto(mapCtx, in.Rows, Intent_Message_TableCardRow_ToProto)
	out.Buttons = direct.Slice_ToProto(mapCtx, in.Buttons, Intent_Message_BasicCard_Button_ToProto)
	return out
}
func Intent_Message_TableCardCell_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCardCell) *krm.Intent_Message_TableCardCell {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCardCell{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func Intent_Message_TableCardCell_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCardCell) *pb.Intent_Message_TableCardCell {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCardCell{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
func Intent_Message_TableCardRow_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_TableCardRow) *krm.Intent_Message_TableCardRow {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_TableCardRow{}
	out.Cells = direct.Slice_FromProto(mapCtx, in.Cells, Intent_Message_TableCardCell_FromProto)
	out.DividerAfter = direct.LazyPtr(in.GetDividerAfter())
	return out
}
func Intent_Message_TableCardRow_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_TableCardRow) *pb.Intent_Message_TableCardRow {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_TableCardRow{}
	out.Cells = direct.Slice_ToProto(mapCtx, in.Cells, Intent_Message_TableCardCell_ToProto)
	out.DividerAfter = direct.ValueOf(in.DividerAfter)
	return out
}
func Intent_Message_Text_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Message_Text) *krm.Intent_Message_Text {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Message_Text{}
	out.Text = in.Text
	return out
}
func Intent_Message_Text_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Message_Text) *pb.Intent_Message_Text {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Message_Text{}
	out.Text = in.Text
	return out
}
func Intent_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Parameter{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Value = direct.LazyPtr(in.GetValue())
	out.DefaultValue = direct.LazyPtr(in.GetDefaultValue())
	out.EntityTypeDisplayName = direct.LazyPtr(in.GetEntityTypeDisplayName())
	out.Mandatory = direct.LazyPtr(in.GetMandatory())
	out.Prompts = in.Prompts
	out.IsList = direct.LazyPtr(in.GetIsList())
	return out
}
func Intent_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Parameter) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Value = direct.ValueOf(in.Value)
	out.DefaultValue = direct.ValueOf(in.DefaultValue)
	out.EntityTypeDisplayName = direct.ValueOf(in.EntityTypeDisplayName)
	out.Mandatory = direct.ValueOf(in.Mandatory)
	out.Prompts = in.Prompts
	out.IsList = direct.ValueOf(in.IsList)
	return out
}
func Intent_TrainingPhrase_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase{}
	// MISSING: Name
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_FromProto)
	out.TimesAddedCount = direct.LazyPtr(in.GetTimesAddedCount())
	return out
}
func Intent_TrainingPhrase_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	// MISSING: Name
	out.Type = direct.Enum_ToProto[pb.Intent_TrainingPhrase_Type](mapCtx, in.Type)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_ToProto)
	out.TimesAddedCount = direct.ValueOf(in.TimesAddedCount)
	return out
}
func Intent_TrainingPhraseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhraseObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhraseObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	// MISSING: Parts
	// MISSING: TimesAddedCount
	return out
}
func Intent_TrainingPhraseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhraseObservedState) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	// MISSING: Parts
	// MISSING: TimesAddedCount
	return out
}
func Intent_TrainingPhrase_Part_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase_Part{}
	out.Text = direct.LazyPtr(in.GetText())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.UserDefined = direct.LazyPtr(in.GetUserDefined())
	return out
}
func Intent_TrainingPhrase_Part_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase_Part) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.Text = direct.ValueOf(in.Text)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.Alias = direct.ValueOf(in.Alias)
	out.UserDefined = direct.ValueOf(in.UserDefined)
	return out
}
