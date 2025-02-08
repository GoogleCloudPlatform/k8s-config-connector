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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)
func ConversationModel_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModel) *krm.ConversationModel {
	if in == nil {
		return nil
	}
	out := &krm.ConversationModel{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	out.Datasets = direct.Slice_FromProto(mapCtx, in.Datasets, InputDataset_FromProto)
	// MISSING: State
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.ArticleSuggestionModelMetadata = ArticleSuggestionModelMetadata_FromProto(mapCtx, in.GetArticleSuggestionModelMetadata())
	out.SmartReplyModelMetadata = SmartReplyModelMetadata_FromProto(mapCtx, in.GetSmartReplyModelMetadata())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ConversationModel_ToProto(mapCtx *direct.MapContext, in *krm.ConversationModel) *pb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModel{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	out.Datasets = direct.Slice_ToProto(mapCtx, in.Datasets, InputDataset_ToProto)
	// MISSING: State
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	if oneof := ArticleSuggestionModelMetadata_ToProto(mapCtx, in.ArticleSuggestionModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.ConversationModel_ArticleSuggestionModelMetadata{ArticleSuggestionModelMetadata: oneof}
	}
	if oneof := SmartReplyModelMetadata_ToProto(mapCtx, in.SmartReplyModelMetadata); oneof != nil {
		out.ModelMetadata = &pb.ConversationModel_SmartReplyModelMetadata{SmartReplyModelMetadata: oneof}
	}
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ConversationModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModel) *krm.ConversationModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Datasets
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func ConversationModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationModelObservedState) *pb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModel{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Datasets
	out.State = direct.Enum_ToProto[pb.ConversationModel_State](mapCtx, in.State)
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	return out
}
func DialogflowConversationModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModel) *krm.DialogflowConversationModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Datasets
	// MISSING: State
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DialogflowConversationModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelObservedState) *pb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModel{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Datasets
	// MISSING: State
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DialogflowConversationModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversationModel) *krm.DialogflowConversationModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationModelSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Datasets
	// MISSING: State
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DialogflowConversationModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationModelSpec) *pb.ConversationModel {
	if in == nil {
		return nil
	}
	out := &pb.ConversationModel{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: Datasets
	// MISSING: State
	// MISSING: LanguageCode
	// MISSING: ArticleSuggestionModelMetadata
	// MISSING: SmartReplyModelMetadata
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func InputDataset_FromProto(mapCtx *direct.MapContext, in *pb.InputDataset) *krm.InputDataset {
	if in == nil {
		return nil
	}
	out := &krm.InputDataset{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	return out
}
func InputDataset_ToProto(mapCtx *direct.MapContext, in *krm.InputDataset) *pb.InputDataset {
	if in == nil {
		return nil
	}
	out := &pb.InputDataset{}
	out.Dataset = direct.ValueOf(in.Dataset)
	return out
}
