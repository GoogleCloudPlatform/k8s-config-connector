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
// proto.message: google.cloud.aiplatform.v1.CustomJob
// api.group: vertexai.cnrm.cloud.google.com

package customjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(customJobFuzzer())
}

func customJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CustomJob{},
		VertexAICustomJobSpec_FromProto, VertexAICustomJobSpec_ToProto,
		VertexAICustomJobObservedState_FromProto, VertexAICustomJobObservedState_ToProto,
	)

	// Special identifiers
	f.UnimplementedFields.Insert(".name")

	f.SpecField(".display_name")
	f.SpecField(".job_spec")
	f.SpecField(".labels")
	f.SpecField(".encryption_spec")

	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".update_time")
	f.StatusField(".error")
	f.StatusField(".web_access_uris")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.gpu_partition_size")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.tpu_topology")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.reservation_affinity")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts[].instance_ip")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts[].mount_point")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts[].filesystem")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts[].volume_handle")
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".error.details[].type_url")
	f.Unimplemented_NotYetTriaged(".error.details[].value")

	return f
}
