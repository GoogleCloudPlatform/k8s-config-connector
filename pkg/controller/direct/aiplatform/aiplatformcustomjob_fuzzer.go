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
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(aiplatformCustomJobFuzzer())
}

func aiplatformCustomJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CustomJob{},
		AIPlatformCustomJobSpec_FromProto, AIPlatformCustomJobSpec_ToProto,
		AIPlatformCustomJobObservedState_FromProto, AIPlatformCustomJobObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".job_spec")
	f.SpecField(".labels")
	f.SpecField(".encryption_spec")

	// Status fields
	f.Unimplemented_Identity(".name")
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".start_time")
	f.StatusField(".end_time")
	f.StatusField(".update_time")
	f.StatusField(".error")
	f.StatusField(".web_access_uris")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	// Unimplemented / Ignored fields
	f.Unimplemented_NotYetTriaged(".error.details")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].lustre_mounts")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].python_package_spec.env")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.min_gpu_driver_version")
	f.Unimplemented_NotYetTriaged(".job_spec.worker_pool_specs[].machine_spec.gpu_partition_size")

	return f
}
