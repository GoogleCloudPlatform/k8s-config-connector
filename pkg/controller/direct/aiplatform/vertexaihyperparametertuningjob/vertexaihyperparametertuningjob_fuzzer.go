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
// proto.message: google.cloud.aiplatform.v1.HyperparameterTuningJob
// api.group: aiplatform.cnrm.cloud.google.com

package vertexaihyperparametertuningjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(hyperparameterTuningJobFuzzer())
}

func hyperparameterTuningJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HyperparameterTuningJob{},
		aiplatform.VertexAIHyperparameterTuningJobSpec_FromProto, aiplatform.VertexAIHyperparameterTuningJobSpec_ToProto,
		aiplatform.VertexAIHyperparameterTuningJobObservedState_FromProto, aiplatform.VertexAIHyperparameterTuningJobObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".study_spec")
	f.SpecField(".max_trial_count")
	f.SpecField(".parallel_trial_count")
	f.SpecField(".max_failed_trial_count")
	f.SpecField(".trial_job_spec")
	f.SpecField(".labels")
	f.SpecField(".encryption_spec")

	// Status fields (ObservedState)
	f.StatusField(".trials")
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".update_time")
	f.StatusField(".error")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	// Unmapped fields or details to ignore in fuzz testing
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".trials[].error.details")
	f.Unimplemented_NotYetTriaged(".trials[].web_access_uris")

	return f
}
