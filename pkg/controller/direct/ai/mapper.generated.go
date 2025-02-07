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

package ai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1/generativelanguagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AiModelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AiModelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AiModelObservedState{}
	// MISSING: Name
	// MISSING: BaseModelID
	// MISSING: Version
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InputTokenLimit
	// MISSING: OutputTokenLimit
	// MISSING: SupportedGenerationMethods
	// MISSING: Temperature
	// MISSING: MaxTemperature
	// MISSING: TopP
	// MISSING: TopK
	return out
}
func AiModelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AiModelObservedState) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: BaseModelID
	// MISSING: Version
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InputTokenLimit
	// MISSING: OutputTokenLimit
	// MISSING: SupportedGenerationMethods
	// MISSING: Temperature
	// MISSING: MaxTemperature
	// MISSING: TopP
	// MISSING: TopK
	return out
}
func AiModelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.AiModelSpec {
	if in == nil {
		return nil
	}
	out := &krm.AiModelSpec{}
	// MISSING: Name
	// MISSING: BaseModelID
	// MISSING: Version
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InputTokenLimit
	// MISSING: OutputTokenLimit
	// MISSING: SupportedGenerationMethods
	// MISSING: Temperature
	// MISSING: MaxTemperature
	// MISSING: TopP
	// MISSING: TopK
	return out
}
func AiModelSpec_ToProto(mapCtx *direct.MapContext, in *krm.AiModelSpec) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	// MISSING: Name
	// MISSING: BaseModelID
	// MISSING: Version
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: InputTokenLimit
	// MISSING: OutputTokenLimit
	// MISSING: SupportedGenerationMethods
	// MISSING: Temperature
	// MISSING: MaxTemperature
	// MISSING: TopP
	// MISSING: TopK
	return out
}
func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Name = direct.LazyPtr(in.GetName())
	out.BaseModelID = direct.LazyPtr(in.GetBaseModelId())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InputTokenLimit = direct.LazyPtr(in.GetInputTokenLimit())
	out.OutputTokenLimit = direct.LazyPtr(in.GetOutputTokenLimit())
	out.SupportedGenerationMethods = in.SupportedGenerationMethods
	out.Temperature = in.Temperature
	out.MaxTemperature = in.MaxTemperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	return out
}
func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Name = direct.ValueOf(in.Name)
	out.BaseModelId = direct.ValueOf(in.BaseModelID)
	out.Version = direct.ValueOf(in.Version)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.InputTokenLimit = direct.ValueOf(in.InputTokenLimit)
	out.OutputTokenLimit = direct.ValueOf(in.OutputTokenLimit)
	out.SupportedGenerationMethods = in.SupportedGenerationMethods
	out.Temperature = in.Temperature
	out.MaxTemperature = in.MaxTemperature
	out.TopP = in.TopP
	out.TopK = in.TopK
	return out
}
