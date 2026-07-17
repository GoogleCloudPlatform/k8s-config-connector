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
// proto.message: google.cloud.aiplatform.v1.SpecialistPool
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAISpecialistPoolFuzzer())
}

func vertexAISpecialistPoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SpecialistPool{},
		VertexAISpecialistPoolSpec_FromProto, VertexAISpecialistPoolSpec_ToProto,
		VertexAISpecialistPoolObservedState_FromProto, VertexAISpecialistPoolObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Special resource name field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".specialist_manager_emails")
	f.SpecFields.Insert(".specialist_worker_emails")

	f.StatusFields.Insert(".specialist_managers_count")
	f.StatusFields.Insert(".pending_data_labeling_jobs")

	return f
}
