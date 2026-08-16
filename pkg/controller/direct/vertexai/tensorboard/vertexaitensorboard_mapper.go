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

package tensorboard

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/vertexai"
)

func VertexAITensorboardSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITensorboardSpec) *pb.Tensorboard {
	return vertexai.VertexAITensorboardSpec_v1alpha1_ToProto(mapCtx, in)
}

func VertexAITensorboardSpec_FromProto(mapCtx *direct.MapContext, in *pb.Tensorboard) *krm.VertexAITensorboardSpec {
	return vertexai.VertexAITensorboardSpec_v1alpha1_FromProto(mapCtx, in)
}

func VertexAITensorboardObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITensorboardObservedState) *pb.Tensorboard {
	return vertexai.VertexAITensorboardObservedState_v1alpha1_ToProto(mapCtx, in)
}

func VertexAITensorboardObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Tensorboard) *krm.VertexAITensorboardObservedState {
	return vertexai.VertexAITensorboardObservedState_v1alpha1_FromProto(mapCtx, in)
}
