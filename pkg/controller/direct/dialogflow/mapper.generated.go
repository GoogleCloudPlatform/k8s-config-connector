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
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Intent_FromProto(mapCtx *direct.MapContext, in *pb.Intent) *krm.Intent {
	if in == nil {
		return nil
	}
	out := &krm.Intent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.TrainingPhrases = direct.Slice_FromProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.Parameters, Intent_Parameter_FromProto)
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.IsFallback = direct.LazyPtr(in.GetIsFallback())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Intent_ToProto(mapCtx *direct.MapContext, in *krm.Intent) *pb.Intent {
	if in == nil {
		return nil
	}
	out := &pb.Intent{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.TrainingPhrases = direct.Slice_ToProto(mapCtx, in.TrainingPhrases, Intent_TrainingPhrase_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, Intent_Parameter_ToProto)
	out.Priority = direct.ValueOf(in.Priority)
	out.IsFallback = direct.ValueOf(in.IsFallback)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	return out
}
func Intent_Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Intent_Parameter) *krm.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Intent_Parameter{}
	out.ID = direct.LazyPtr(in.GetId())
	out.EntityType = direct.LazyPtr(in.GetEntityType())
	out.IsList = direct.LazyPtr(in.GetIsList())
	out.Redact = direct.LazyPtr(in.GetRedact())
	return out
}
func Intent_Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Intent_Parameter) *pb.Intent_Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Intent_Parameter{}
	out.Id = direct.ValueOf(in.ID)
	out.EntityType = direct.ValueOf(in.EntityType)
	out.IsList = direct.ValueOf(in.IsList)
	out.Redact = direct.ValueOf(in.Redact)
	return out
}
func Intent_TrainingPhrase_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase) *krm.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Parts = direct.Slice_FromProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_FromProto)
	out.RepeatCount = direct.LazyPtr(in.GetRepeatCount())
	return out
}
func Intent_TrainingPhrase_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase) *pb.Intent_TrainingPhrase {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase{}
	out.Id = direct.ValueOf(in.ID)
	out.Parts = direct.Slice_ToProto(mapCtx, in.Parts, Intent_TrainingPhrase_Part_ToProto)
	out.RepeatCount = direct.ValueOf(in.RepeatCount)
	return out
}
func Intent_TrainingPhrase_Part_FromProto(mapCtx *direct.MapContext, in *pb.Intent_TrainingPhrase_Part) *krm.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &krm.Intent_TrainingPhrase_Part{}
	out.Text = direct.LazyPtr(in.GetText())
	out.ParameterID = direct.LazyPtr(in.GetParameterId())
	return out
}
func Intent_TrainingPhrase_Part_ToProto(mapCtx *direct.MapContext, in *krm.Intent_TrainingPhrase_Part) *pb.Intent_TrainingPhrase_Part {
	if in == nil {
		return nil
	}
	out := &pb.Intent_TrainingPhrase_Part{}
	out.Text = direct.ValueOf(in.Text)
	out.ParameterId = direct.ValueOf(in.ParameterID)
	return out
}
