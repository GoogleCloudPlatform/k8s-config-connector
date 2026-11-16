// Copyright 2026 Google LLC
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
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowConversationModelObservedState_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ConversationModel) *krm.DialogflowConversationModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func DialogflowConversationModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelObservedState) *dialogflowpb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ConversationModel{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[dialogflowpb.ConversationModel_State](mapCtx, in.State)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}

func DialogflowConversationModelSpec_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.ConversationModel) *krm.DialogflowConversationModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Datasets = direct.Slice_FromProto(mapCtx, in.Datasets, InputDataset_FromProto)
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ArticleSuggestionModelMetadata = ArticleSuggestionModelMetadata_FromProto(mapCtx, in.GetArticleSuggestionModelMetadata())
	out.SmartReplyModelMetadata = SmartReplyModelMetadata_FromProto(mapCtx, in.GetSmartReplyModelMetadata())
	return out
}

func DialogflowConversationModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelSpec) *dialogflowpb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.ConversationModel{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Datasets = direct.Slice_ToProto(mapCtx, in.Datasets, InputDataset_ToProto)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	if oneof := ArticleSuggestionModelMetadata_ToProto(mapCtx, in.ArticleSuggestionModelMetadata); oneof != nil {
		out.ModelMetadata = &dialogflowpb.ConversationModel_ArticleSuggestionModelMetadata{ArticleSuggestionModelMetadata: oneof}
	}
	if oneof := SmartReplyModelMetadata_ToProto(mapCtx, in.SmartReplyModelMetadata); oneof != nil {
		out.ModelMetadata = &dialogflowpb.ConversationModel_SmartReplyModelMetadata{SmartReplyModelMetadata: oneof}
	}
	return out
}

func InputDataset_FromProto(mapCtx *direct.MapContext, in *dialogflowpb.InputDataset) *krm.InputDataset {
	if in == nil {
		return nil
	}
	out := &krm.InputDataset{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	return out
}

func InputDataset_ToProto(mapCtx *direct.MapContext, in *krm.InputDataset) *dialogflowpb.InputDataset {
	if in == nil {
		return nil
	}
	out := &dialogflowpb.InputDataset{}
	out.Dataset = direct.ValueOf(in.Dataset)
	return out
}
