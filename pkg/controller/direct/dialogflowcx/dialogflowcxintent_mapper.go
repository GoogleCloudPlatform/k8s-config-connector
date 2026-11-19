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

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowCXIntentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.DialogflowCXIntentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXIntentObservedState{}
	return out
}

func DialogflowCXIntentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXIntentObservedState) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	return out
}

func DialogflowCXIntentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.DialogflowCXIntentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXIntentSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = in.GetDisplayName()
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.Parameters = direct.Slice_FromProto(mapCtx, in.GetParameters(), IntentParameters_FromProto)
	if in.GetPriority() != 0 {
		priority := int(in.GetPriority())
		out.Priority = &priority
	}
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.GetTrainingPhrases(), IntentTrainingPhrases_FromProto)
	return out
}

func DialogflowCXIntentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXIntentSpec) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = in.DisplayName
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, IntentParameters_ToProto)
	if in.Priority != nil {
		out.Priority = int32(*in.Priority)
	}
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, IntentTrainingPhrases_ToProto)
	return out
}

func IntentParameters_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.IntentParameters {
	if in == nil {
		return nil
	}
	out := &krm.IntentParameters{}
	out.EntityType = in.GetEntityType()
	out.Id = in.GetId()
	out.IsList = direct.LazyPtr(in.GetIsList())
	out.Redact = direct.LazyPtr(in.GetRedact())
	return out
}

func IntentParameters_ToProto(mapCtx *direct.MapContext, in *krm.IntentParameters) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.EntityType = in.EntityType
	out.Id = in.Id
	out.IsList = direct.ValueOf(in.IsList)
	out.Redact = direct.ValueOf(in.Redact)
	return out
}

func IntentParts_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.IntentParts {
	if in == nil {
		return nil
	}
	out := &krm.IntentParts{}
	out.ParameterId = direct.LazyPtr(in.GetParameterId())
	out.Text = in.GetText()
	return out
}

func IntentParts_ToProto(mapCtx *direct.MapContext, in *krm.IntentParts) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.ParameterId = direct.ValueOf(in.ParameterId)
	out.Text = in.Text
	return out
}

func IntentTrainingPhrases_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.IntentTrainingPhrases {
	if in == nil {
		return nil
	}
	out := &krm.IntentTrainingPhrases{}
	out.Id = direct.LazyPtr(in.GetId())
	out.Parts = direct.Slice_FromProto(mapCtx, in.GetParts(), IntentParts_FromProto)
	if in.GetRepeatCount() != 0 {
		repeatCount := int(in.GetRepeatCount())
		out.RepeatCount = &repeatCount
	}
	return out
}

func IntentTrainingPhrases_ToProto(mapCtx *direct.MapContext, in *krm.IntentTrainingPhrases) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Id = direct.ValueOf(in.Id)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, IntentParts_ToProto)
	if in.RepeatCount != nil {
		out.RepeatCount = int32(*in.RepeatCount)
	}
	return out
}
