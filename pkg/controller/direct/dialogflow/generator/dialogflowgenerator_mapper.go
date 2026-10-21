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

package generator

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ArticleSuggestionModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ArticleSuggestionModelMetadata) *krm.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func ArticleSuggestionModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ArticleSuggestionModelMetadata) *pb.ArticleSuggestionModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ArticleSuggestionModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[pb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}

func SmartReplyModelMetadata_FromProto(mapCtx *direct.MapContext, in *pb.SmartReplyModelMetadata) *krm.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &krm.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_FromProto(mapCtx, in.GetTrainingModelType())
	return out
}

func SmartReplyModelMetadata_ToProto(mapCtx *direct.MapContext, in *krm.SmartReplyModelMetadata) *pb.SmartReplyModelMetadata {
	if in == nil {
		return nil
	}
	out := &pb.SmartReplyModelMetadata{}
	out.TrainingModelType = direct.Enum_ToProto[pb.ConversationModel_ModelType](mapCtx, in.TrainingModelType)
	return out
}
