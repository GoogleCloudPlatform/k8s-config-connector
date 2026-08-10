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

package tensorboardexperiment

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAITensorboardExperimentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TensorboardExperiment) *krm.VertexAITensorboardExperimentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAITensorboardExperimentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func VertexAITensorboardExperimentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITensorboardExperimentObservedState) *pb.TensorboardExperiment {
	if in == nil {
		return nil
	}
	out := &pb.TensorboardExperiment{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func VertexAITensorboardExperimentSpec_FromProto(mapCtx *direct.MapContext, in *pb.TensorboardExperiment) *krm.VertexAITensorboardExperimentSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAITensorboardExperimentSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Source = direct.LazyPtr(in.GetSource())
	return out
}

func VertexAITensorboardExperimentSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITensorboardExperimentSpec) *pb.TensorboardExperiment {
	if in == nil {
		return nil
	}
	out := &pb.TensorboardExperiment{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Source = direct.ValueOf(in.Source)
	return out
}
