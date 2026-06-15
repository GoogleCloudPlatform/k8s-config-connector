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

package examplestore

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAIExampleStoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExampleStore) *krm.VertexAIExampleStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIExampleStoreObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func VertexAIExampleStoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIExampleStoreObservedState) *pb.ExampleStore {
	if in == nil {
		return nil
	}
	out := &pb.ExampleStore{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func VertexAIExampleStoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExampleStore) *krm.VertexAIExampleStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIExampleStoreSpec{}
	out.DisplayName = in.GetDisplayName()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ExampleStoreConfig = ExampleStoreConfig_FromProto(mapCtx, in.GetExampleStoreConfig())
	return out
}

func VertexAIExampleStoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIExampleStoreSpec) *pb.ExampleStore {
	if in == nil {
		return nil
	}
	out := &pb.ExampleStore{}
	out.DisplayName = in.DisplayName
	out.Description = direct.ValueOf(in.Description)
	out.ExampleStoreConfig = ExampleStoreConfig_ToProto(mapCtx, in.ExampleStoreConfig)
	return out
}

func ExampleStoreConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExampleStoreConfig) *krm.ExampleStoreConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExampleStoreConfig{}
	out.VertexEmbeddingModel = direct.LazyPtr(in.GetVertexEmbeddingModel())
	return out
}

func ExampleStoreConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExampleStoreConfig) *pb.ExampleStoreConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExampleStoreConfig{}
	out.VertexEmbeddingModel = direct.ValueOf(in.VertexEmbeddingModel)
	return out
}
