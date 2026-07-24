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
// proto.message: google.cloud.aiplatform.v1beta1.CustomJob
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(customJobFuzzer())
}

func customJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CustomJob{},
		VertexAICustomJobSpec_v1alpha1_FromProto, VertexAICustomJobSpec_v1alpha1_ToProto,
		VertexAICustomJobObservedState_v1alpha1_FromProto, VertexAICustomJobObservedState_v1alpha1_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].python_package_spec.env")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.min_gpu_driver_version")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".job_spec")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".end_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".error")
	f.StatusFields.Insert(".web_access_uris")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")

	return f
}
