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
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Generator_FromProto(mapCtx *direct.MapContext, in *pb.Generator) *krm.Generator {
	if in == nil {
		return nil
	}
	out := &krm.Generator{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.PromptText = Phrase_FromProto(mapCtx, in.GetPromptText())
	out.Placeholders = direct.Slice_FromProto(mapCtx, in.Placeholders, Generator_Placeholder_FromProto)
	out.LlmModelSettings = LlmModelSettings_FromProto(mapCtx, in.GetLlmModelSettings())
	out.ModelParameter = Generator_ModelParameter_FromProto(mapCtx, in.GetModelParameter())
	return out
}
func Generator_ToProto(mapCtx *direct.MapContext, in *krm.Generator) *pb.Generator {
	if in == nil {
		return nil
	}
	out := &pb.Generator{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PromptText = Phrase_ToProto(mapCtx, in.PromptText)
	out.Placeholders = direct.Slice_ToProto(mapCtx, in.Placeholders, Generator_Placeholder_ToProto)
	out.LlmModelSettings = LlmModelSettings_ToProto(mapCtx, in.LlmModelSettings)
	out.ModelParameter = Generator_ModelParameter_ToProto(mapCtx, in.ModelParameter)
	return out
}
func Generator_ModelParameter_FromProto(mapCtx *direct.MapContext, in *pb.Generator_ModelParameter) *krm.Generator_ModelParameter {
	if in == nil {
		return nil
	}
	out := &krm.Generator_ModelParameter{}
	out.Temperature = in.Temperature
	out.MaxDecodeSteps = in.MaxDecodeSteps
	out.TopP = in.TopP
	out.TopK = in.TopK
	return out
}
func Generator_ModelParameter_ToProto(mapCtx *direct.MapContext, in *krm.Generator_ModelParameter) *pb.Generator_ModelParameter {
	if in == nil {
		return nil
	}
	out := &pb.Generator_ModelParameter{}
	out.Temperature = in.Temperature
	out.MaxDecodeSteps = in.MaxDecodeSteps
	out.TopP = in.TopP
	out.TopK = in.TopK
	return out
}
func Generator_Placeholder_FromProto(mapCtx *direct.MapContext, in *pb.Generator_Placeholder) *krm.Generator_Placeholder {
	if in == nil {
		return nil
	}
	out := &krm.Generator_Placeholder{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Generator_Placeholder_ToProto(mapCtx *direct.MapContext, in *krm.Generator_Placeholder) *pb.Generator_Placeholder {
	if in == nil {
		return nil
	}
	out := &pb.Generator_Placeholder{}
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func LlmModelSettings_FromProto(mapCtx *direct.MapContext, in *pb.LlmModelSettings) *krm.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &krm.LlmModelSettings{}
	out.Model = direct.LazyPtr(in.GetModel())
	out.PromptText = direct.LazyPtr(in.GetPromptText())
	return out
}
func LlmModelSettings_ToProto(mapCtx *direct.MapContext, in *krm.LlmModelSettings) *pb.LlmModelSettings {
	if in == nil {
		return nil
	}
	out := &pb.LlmModelSettings{}
	out.Model = direct.ValueOf(in.Model)
	out.PromptText = direct.ValueOf(in.PromptText)
	return out
}
func Phrase_FromProto(mapCtx *direct.MapContext, in *pb.Phrase) *krm.Phrase {
	if in == nil {
		return nil
	}
	out := &krm.Phrase{}
	out.Text = direct.LazyPtr(in.GetText())
	return out
}
func Phrase_ToProto(mapCtx *direct.MapContext, in *krm.Phrase) *pb.Phrase {
	if in == nil {
		return nil
	}
	out := &pb.Phrase{}
	out.Text = direct.ValueOf(in.Text)
	return out
}
