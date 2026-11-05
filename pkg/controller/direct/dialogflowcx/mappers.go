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

// DialogflowCXIntentSpec_v1alpha1_FromProto maps pb.Intent to DialogflowCXIntentSpec.
func DialogflowCXIntentSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.DialogflowCXIntentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXIntentSpec{}
	out.DisplayName = in.GetDisplayName()
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, IntentTrainingPhrases_v1alpha1_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, IntentParameters_v1alpha1_FromProto)
	if in.Priority != 0 {
		p := int(in.Priority)
		out.Priority = &p
	}
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

// DialogflowCXIntentSpec_v1alpha1_ToProto maps DialogflowCXIntentSpec to pb.Intent.
func DialogflowCXIntentSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXIntentSpec) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.DisplayName = in.DisplayName
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, IntentTrainingPhrases_v1alpha1_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, IntentParameters_v1alpha1_ToProto)
	if in.Priority != nil {
		out.Priority = int32(*in.Priority)
	}
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.Description = direct.ValueOf(in.Description)
	return out
}

// IntentTrainingPhrases_v1alpha1_FromProto maps pb.Intent_TrainingPhrase to IntentTrainingPhrases.
func IntentTrainingPhrases_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.IntentTrainingPhrases {
	if in == nil {
		return nil
	}
	out := &krm.IntentTrainingPhrases{}
	out.Id = direct.LazyPtr(in.GetId())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, IntentParts_v1alpha1_FromProto)
	if in.RepeatCount != 0 {
		rc := int(in.RepeatCount)
		out.RepeatCount = &rc
	}
	return out
}

// IntentTrainingPhrases_v1alpha1_ToProto maps IntentTrainingPhrases to pb.Intent_TrainingPhrase.
func IntentTrainingPhrases_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.IntentTrainingPhrases) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Id = direct.ValueOf(in.Id)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, IntentParts_v1alpha1_ToProto)
	if in.RepeatCount != nil {
		out.RepeatCount = int32(*in.RepeatCount)
	}
	return out
}

// IntentParameters_v1alpha1_FromProto maps pb.Intent_Parameter to IntentParameters.
func IntentParameters_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.IntentParameters {
	if in == nil {
		return nil
	}
	out := &krm.IntentParameters{}
	out.Id = in.GetId()
	out.EntityType = in.GetEntityType()
	out.IsList = direct.LazyPtr(in.GetIsList())
	out.Redact = direct.LazyPtr(in.GetRedact())
	return out
}

// IntentParameters_v1alpha1_ToProto maps IntentParameters to pb.Intent_Parameter.
func IntentParameters_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.IntentParameters) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.Id = in.Id
	out.EntityType = in.EntityType
	out.IsList = direct.ValueOf(in.IsList)
	out.Redact = direct.ValueOf(in.Redact)
	return out
}

// IntentParts_v1alpha1_FromProto maps pb.Intent_TrainingPhrase_Part to IntentParts.
func IntentParts_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.IntentParts {
	if in == nil {
		return nil
	}
	out := &krm.IntentParts{}
	out.ParameterId = direct.LazyPtr(in.GetParameterId())
	out.Text = in.GetText()
	return out
}

// IntentParts_v1alpha1_ToProto maps IntentParts to pb.Intent_TrainingPhrase_Part.
func IntentParts_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.IntentParts) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.ParameterId = direct.ValueOf(in.ParameterId)
	out.Text = in.Text
	return out
}
