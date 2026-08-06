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
// proto.message: google.cloud.aiplatform.v1beta1.Tensorboard
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAITensorboardFuzzer())
}

func vertexAITensorboardFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Tensorboard{},
		VertexAITensorboardSpec_v1alpha1_FromProto, VertexAITensorboardSpec_v1alpha1_ToProto,
		VertexAITensorboardObservedState_v1alpha1_FromProto, VertexAITensorboardObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()
	f.Unimplemented_NotYetTriaged(".labels")

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".encryption_spec")
	f.SpecField(".is_default")

	f.StatusField(".blob_storage_path_prefix")
	f.StatusField(".run_count")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	return f
}
