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
// proto.message: google.cloud.aiplatform.v1.DataLabelingJob
// api.group: vertexai.cnrm.cloud.google.com

package datalabelingjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(datalabelingjobFuzzer())
}

func datalabelingjobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataLabelingJob{},
		VertexAIDataLabelingJobSpec_FromProto, VertexAIDataLabelingJobSpec_ToProto,
		VertexAIDataLabelingJobObservedState_FromProto, VertexAIDataLabelingJobObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".error.details[].value")
	f.Unimplemented_NotYetTriaged(".inputs")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".datasets")
	f.SpecFields.Insert(".annotation_labels")
	f.SpecFields.Insert(".labeler_count")
	f.SpecFields.Insert(".instruction_uri")
	f.SpecFields.Insert(".inputs_schema_uri")
	f.SpecFields.Insert(".inputs")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".specialist_pools")
	f.SpecFields.Insert(".encryption_spec")
	f.SpecFields.Insert(".active_learning_config")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".labeling_progress")
	f.StatusFields.Insert(".current_spend")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".error")

	return f
}
