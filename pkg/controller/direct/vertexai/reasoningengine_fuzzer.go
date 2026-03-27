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

// +tool:fuzz-gen
// proto.message: google.cloud.aiplatform.v1beta1.ReasoningEngine
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIReasoningEngineFuzzer())
}

func vertexAIReasoningEngineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ReasoningEngine{},
		VertexAIReasoningEngineSpec_FromProto, VertexAIReasoningEngineSpec_ToProto,
		VertexAIReasoningEngineObservedState_FromProto, VertexAIReasoningEngineObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".spec")
	f.SpecFields.Insert(".etag")
	f.SpecFields.Insert(".context_spec")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
