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

package customjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAIDataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobObservedState {
	return nil
}

func VertexAIDataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobObservedState) *pb.DataLabelingJob {
	return nil
}

func VertexAIDataLabelingJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobSpec {
	return nil
}

func VertexAIDataLabelingJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobSpec) *pb.DataLabelingJob {
	return nil
}
