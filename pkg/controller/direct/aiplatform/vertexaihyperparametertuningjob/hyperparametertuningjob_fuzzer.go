// Copyright 2025 Google LLC
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

package vertexaihyperparametertuningjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(hyperparameterTuningJobFuzzer())
}

func hyperparameterTuningJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.HyperparameterTuningJob{},
		VertexAIHyperparameterTuningJobSpec_FromProto, VertexAIHyperparameterTuningJobSpec_ToProto,
		VertexAIHyperparameterTuningJobObservedState_FromProto, VertexAIHyperparameterTuningJobObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")          // mapped to metadata.name / externalRef
	f.UnimplementedFields.Insert(".satisfies_pzi") // output only reserved
	f.UnimplementedFields.Insert(".satisfies_pzs") // output only reserved
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".start_time")
	f.UnimplementedFields.Insert(".end_time")
	f.UnimplementedFields.Insert(".error")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".study_spec")
	f.UnimplementedFields.Insert(".study_spec.study_stopping_config.min_num_trials")
	f.UnimplementedFields.Insert(".study_spec.transfer_learning_config")
	f.UnimplementedFields.Insert(".study_spec.convex_stop_config.use_seconds")
	f.UnimplementedFields.Insert(".trials")
	f.UnimplementedFields.Insert(".max_trial_count")
	f.UnimplementedFields.Insert(".max_failed_trial_count")
	f.UnimplementedFields.Insert(".parallel_trial_count")
	f.UnimplementedFields.Insert(".trial_job_spec")
	f.UnimplementedFields.Insert(".encryption_spec")
	f.UnimplementedFields.Insert(".display_name")
	f.UnimplementedFields.Insert(".state")

	return f
}
