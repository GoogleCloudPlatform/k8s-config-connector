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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzer())
}

func fuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.Tensorboard, krm.VertexAITensorBoardSpec, krm.VertexAITensorBoardObservedState](
		&pb.Tensorboard{},
		VertexAITensorBoardSpec_FromProto,
		VertexAITensorBoardSpec_ToProto,
		VertexAITensorBoardObservedState_FromProto,
		VertexAITensorBoardObservedState_ToProto,
	)

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

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()
	f.Unimplemented_LabelsAnnotations(".labels")

	return f
}
